// Auto-generated by avdl-compiler v1.3.9 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/chat_ui.avdl

package chat1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type ChatAttachmentUploadStartArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type ChatAttachmentUploadProgressArg struct {
	SessionID     int `codec:"sessionID" json:"sessionID"`
	BytesComplete int `codec:"bytesComplete" json:"bytesComplete"`
	BytesTotal    int `codec:"bytesTotal" json:"bytesTotal"`
}

type ChatAttachmentUploadDoneArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type ChatAttachmentPreviewUploadStartArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type ChatAttachmentPreviewUploadDoneArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type ChatAttachmentDownloadStartArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type ChatAttachmentDownloadProgressArg struct {
	SessionID     int `codec:"sessionID" json:"sessionID"`
	BytesComplete int `codec:"bytesComplete" json:"bytesComplete"`
	BytesTotal    int `codec:"bytesTotal" json:"bytesTotal"`
}

type ChatAttachmentDownloadDoneArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type ChatUiInterface interface {
	ChatAttachmentUploadStart(context.Context, int) error
	ChatAttachmentUploadProgress(context.Context, ChatAttachmentUploadProgressArg) error
	ChatAttachmentUploadDone(context.Context, int) error
	ChatAttachmentPreviewUploadStart(context.Context, int) error
	ChatAttachmentPreviewUploadDone(context.Context, int) error
	ChatAttachmentDownloadStart(context.Context, int) error
	ChatAttachmentDownloadProgress(context.Context, ChatAttachmentDownloadProgressArg) error
	ChatAttachmentDownloadDone(context.Context, int) error
}

func ChatUiProtocol(i ChatUiInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "chat.1.chatUi",
		Methods: map[string]rpc.ServeHandlerDescription{
			"chatAttachmentUploadStart": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentUploadStartArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentUploadStartArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentUploadStartArg)(nil), args)
						return
					}
					err = i.ChatAttachmentUploadStart(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatAttachmentUploadProgress": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentUploadProgressArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentUploadProgressArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentUploadProgressArg)(nil), args)
						return
					}
					err = i.ChatAttachmentUploadProgress(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatAttachmentUploadDone": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentUploadDoneArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentUploadDoneArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentUploadDoneArg)(nil), args)
						return
					}
					err = i.ChatAttachmentUploadDone(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatAttachmentPreviewUploadStart": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentPreviewUploadStartArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentPreviewUploadStartArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentPreviewUploadStartArg)(nil), args)
						return
					}
					err = i.ChatAttachmentPreviewUploadStart(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatAttachmentPreviewUploadDone": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentPreviewUploadDoneArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentPreviewUploadDoneArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentPreviewUploadDoneArg)(nil), args)
						return
					}
					err = i.ChatAttachmentPreviewUploadDone(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatAttachmentDownloadStart": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentDownloadStartArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentDownloadStartArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentDownloadStartArg)(nil), args)
						return
					}
					err = i.ChatAttachmentDownloadStart(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatAttachmentDownloadProgress": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentDownloadProgressArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentDownloadProgressArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentDownloadProgressArg)(nil), args)
						return
					}
					err = i.ChatAttachmentDownloadProgress(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatAttachmentDownloadDone": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentDownloadDoneArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentDownloadDoneArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentDownloadDoneArg)(nil), args)
						return
					}
					err = i.ChatAttachmentDownloadDone(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type ChatUiClient struct {
	Cli rpc.GenericClient
}

func (c ChatUiClient) ChatAttachmentUploadStart(ctx context.Context, sessionID int) (err error) {
	__arg := ChatAttachmentUploadStartArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatAttachmentUploadStart", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatAttachmentUploadProgress(ctx context.Context, __arg ChatAttachmentUploadProgressArg) (err error) {
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatAttachmentUploadProgress", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatAttachmentUploadDone(ctx context.Context, sessionID int) (err error) {
	__arg := ChatAttachmentUploadDoneArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatAttachmentUploadDone", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatAttachmentPreviewUploadStart(ctx context.Context, sessionID int) (err error) {
	__arg := ChatAttachmentPreviewUploadStartArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatAttachmentPreviewUploadStart", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatAttachmentPreviewUploadDone(ctx context.Context, sessionID int) (err error) {
	__arg := ChatAttachmentPreviewUploadDoneArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatAttachmentPreviewUploadDone", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatAttachmentDownloadStart(ctx context.Context, sessionID int) (err error) {
	__arg := ChatAttachmentDownloadStartArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatAttachmentDownloadStart", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatAttachmentDownloadProgress(ctx context.Context, __arg ChatAttachmentDownloadProgressArg) (err error) {
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatAttachmentDownloadProgress", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatAttachmentDownloadDone(ctx context.Context, sessionID int) (err error) {
	__arg := ChatAttachmentDownloadDoneArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatAttachmentDownloadDone", []interface{}{__arg}, nil)
	return
}
