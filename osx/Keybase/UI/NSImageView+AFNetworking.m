// UIImageView+AFNetworking.m
//
// Copyright (c) 2013-2015 AFNetworking (http://afnetworking.com)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

#import "NSImageView+AFNetworking.h"

#import <objc/runtime.h>

#import "AFHTTPRequestOperation.h"

@interface AFImageCache : NSCache <AFImageCache>
@end

#pragma mark -

@interface NSImageView (_AFNetworking)
@property (readwrite, nonatomic, strong, setter = af_setImageRequestOperation:) AFHTTPRequestOperation *af_imageRequestOperation;
@end

@implementation NSImageView (_AFNetworking)

+ (NSOperationQueue *)af_sharedImageRequestOperationQueue {
  static NSOperationQueue *_af_sharedImageRequestOperationQueue = nil;
  static dispatch_once_t onceToken;
  dispatch_once(&onceToken, ^{
    _af_sharedImageRequestOperationQueue = [[NSOperationQueue alloc] init];
    _af_sharedImageRequestOperationQueue.maxConcurrentOperationCount = NSOperationQueueDefaultMaxConcurrentOperationCount;
  });

  return _af_sharedImageRequestOperationQueue;
}

- (AFHTTPRequestOperation *)af_imageRequestOperation {
  return (AFHTTPRequestOperation *)objc_getAssociatedObject(self, @selector(af_imageRequestOperation));
}

- (void)af_setImageRequestOperation:(AFHTTPRequestOperation *)imageRequestOperation {
  objc_setAssociatedObject(self, @selector(af_imageRequestOperation), imageRequestOperation, OBJC_ASSOCIATION_RETAIN_NONATOMIC);
}

@end

#pragma mark -

@implementation NSImageView (AFNetworking)
@dynamic imageResponseSerializer;

+ (id <AFImageCache>)sharedImageCache {
  static AFImageCache *_af_defaultImageCache = nil;
  static dispatch_once_t oncePredicate;
  dispatch_once(&oncePredicate, ^{
    _af_defaultImageCache = [[AFImageCache alloc] init];

    // TODO: Low memory notification for OS X
//    [[NSNotificationCenter defaultCenter] addObserverForName:UIApplicationDidReceiveMemoryWarningNotification object:nil queue:[NSOperationQueue mainQueue] usingBlock:^(NSNotification * __unused notification) {
//      [_af_defaultImageCache removeAllObjects];
//    }];
  });

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wgnu"
  return objc_getAssociatedObject(self, @selector(sharedImageCache)) ?: _af_defaultImageCache;
#pragma clang diagnostic pop
}

+ (void)setSharedImageCache:(id <AFImageCache>)imageCache {
  objc_setAssociatedObject(self, @selector(sharedImageCache), imageCache, OBJC_ASSOCIATION_RETAIN_NONATOMIC);
}

#pragma mark -

- (id <AFURLResponseSerialization>)imageResponseSerializer {
  static id <AFURLResponseSerialization> _af_defaultImageResponseSerializer = nil;
  static dispatch_once_t onceToken;
  dispatch_once(&onceToken, ^{
    _af_defaultImageResponseSerializer = [AFImageResponseSerializer serializer];
  });

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wgnu"
  return objc_getAssociatedObject(self, @selector(imageResponseSerializer)) ?: _af_defaultImageResponseSerializer;
#pragma clang diagnostic pop
}

- (void)setImageResponseSerializer:(id <AFURLResponseSerialization>)serializer {
  objc_setAssociatedObject(self, @selector(imageResponseSerializer), serializer, OBJC_ASSOCIATION_RETAIN_NONATOMIC);
}

#pragma mark -

- (void)setImageWithURL:(NSURL *)url {
  [self setImageWithURL:url placeholderImage:nil];
}

- (void)setImageWithURL:(NSURL *)url
       placeholderImage:(NSImage *)placeholderImage
{
  if ([url.scheme isEqualTo:@"bundle"]) {
    self.image = [NSImage imageNamed:url.host];
    return;
  }

  NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:url];
  [request addValue:@"image/*" forHTTPHeaderField:@"Accept"];

  [self setImageWithURLRequest:request placeholderImage:placeholderImage success:nil failure:nil];
}

- (void)setImageWithURLRequest:(NSURLRequest *)urlRequest
              placeholderImage:(NSImage *)placeholderImage
                       success:(void (^)(NSURLRequest *request, NSHTTPURLResponse *response, NSImage *image))success
                       failure:(void (^)(NSURLRequest *request, NSHTTPURLResponse *response, NSError *error))failure
{
  [self cancelImageRequestOperation];

  NSImage *cachedImage = [[[self class] sharedImageCache] cachedImageForRequest:urlRequest];
  if (cachedImage) {
    if (success) {
      success(nil, nil, cachedImage);
    } else {
      self.image = cachedImage;
    }

    self.af_imageRequestOperation = nil;
  } else {
    if (placeholderImage) {
      self.image = placeholderImage;
    }

    __weak __typeof(self)weakSelf = self;
    self.af_imageRequestOperation = [[AFHTTPRequestOperation alloc] initWithRequest:urlRequest];
    self.af_imageRequestOperation.responseSerializer = self.imageResponseSerializer;
    [self.af_imageRequestOperation setCompletionBlockWithSuccess:^(AFHTTPRequestOperation *operation, id responseObject) {
      __strong __typeof(weakSelf)strongSelf = weakSelf;
      if ([[urlRequest URL] isEqual:[strongSelf.af_imageRequestOperation.request URL]]) {
        if (success) {
          success(urlRequest, operation.response, responseObject);
        } else if (responseObject) {
          strongSelf.image = responseObject;
        }

        if (operation == strongSelf.af_imageRequestOperation){
          strongSelf.af_imageRequestOperation = nil;
        }
      }

      [[[strongSelf class] sharedImageCache] cacheImage:responseObject forRequest:urlRequest];
    } failure:^(AFHTTPRequestOperation *operation, NSError *error) {
      __strong __typeof(weakSelf)strongSelf = weakSelf;
      if ([[urlRequest URL] isEqual:[strongSelf.af_imageRequestOperation.request URL]]) {
        if (failure) {
          failure(urlRequest, operation.response, error);
        }

        if (operation == strongSelf.af_imageRequestOperation){
          strongSelf.af_imageRequestOperation = nil;
        }
      }
    }];

    [[[self class] af_sharedImageRequestOperationQueue] addOperation:self.af_imageRequestOperation];
  }
}

- (void)cancelImageRequestOperation {
  [self.af_imageRequestOperation cancel];
  self.af_imageRequestOperation = nil;
}

@end

#pragma mark -

static inline NSString * AFImageCacheKeyFromURLRequest(NSURLRequest *request) {
  return [[request URL] absoluteString];
}

@implementation AFImageCache

- (NSImage *)cachedImageForRequest:(NSURLRequest *)request {
  switch ([request cachePolicy]) {
    case NSURLRequestReloadIgnoringCacheData:
    case NSURLRequestReloadIgnoringLocalAndRemoteCacheData:
      return nil;
    default:
      break;
  }

  return [self objectForKey:AFImageCacheKeyFromURLRequest(request)];
}

- (void)cacheImage:(NSImage *)image
        forRequest:(NSURLRequest *)request
{
  if (image && request) {
    [self setObject:image forKey:AFImageCacheKeyFromURLRequest(request)];
  }
}

@end
