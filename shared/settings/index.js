// @flow
import * as React from 'react'
import * as ConfigGen from '../actions/config-gen'
import * as Types from '../constants/types/settings'
import SettingsContainer from './render'
import {connect} from '../util/container'
import {switchTo} from '../actions/route-tree'
import {type RouteProps} from '../route-tree/render-route'

type OwnProps = {|children: React.Node, ...$Exact<RouteProps<{}, {}>>|}

const mapStateToProps = (state, {routeLeafTags, routeSelected}: OwnProps) => ({
  _badgeNumbers: state.notifications.get('navBadges'),
  badgeNotifications: !state.push.hasPermissions,
  isModal: routeLeafTags.modal,
  selectedTab: ((routeSelected: any): Types.Tab),
})

const mapDispatchToProps = (dispatch, {routePath}: OwnProps) => ({
  onLogout: () => dispatch(ConfigGen.createLogout()),
  onTabChange: (tab: Types.Tab) => dispatch(switchTo(routePath.push(tab))),
})

const mergeProps = (stateProps, dispatchProps, ownProps) => ({
  badgeNotifications: stateProps.badgeNotifications,
  // $FlowIssue fix badgeNumbers
  badgeNumbers: stateProps._badgeNumbers.toObject(),
  children: ownProps.children,
  isModal: stateProps.isModal,
  onLogout: dispatchProps.onLogout,
  onTabChange: dispatchProps.onTabChange,
  selectedTab: stateProps.selectedTab,
})

export default connect<OwnProps, _, _, _, _>(
  mapStateToProps,
  mapDispatchToProps,
  mergeProps
)(SettingsContainer)
