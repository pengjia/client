// @flow
import * as Chat2Gen from '../../../../actions/chat2-gen'
import * as Constants from '../../../../constants/chat2'
import * as Types from '../../../../constants/types/chat2'
import {SmallTeam} from '.'
import {connect, type TypedState, type Dispatch} from '../../../../util/container'

type OwnProps = {conversationIDKey: Types.ConversationIDKey}

const mapStateToProps = (state: TypedState, ownProps: OwnProps) => {
  const _conversationIDKey = ownProps.conversationIDKey
  const youAreReset = false // Constants.isResetConversationIDKey(state, _conversationIDKey)

  return {
    _meta: Constants.getMeta(state, _conversationIDKey),
    _username: state.config.username || '',
    hasBadge: Constants.getHasBadge(state, _conversationIDKey),
    hasUnread: Constants.getHasUnread(state, _conversationIDKey),
    isSelected: !state.chat2.pendingSelected && Constants.getIsSelected(state, _conversationIDKey),
    youAreReset,
  }
}

const mapDispatchToProps = (dispatch: Dispatch, {conversationIDKey}: OwnProps) => ({
  onSelectConversation: () =>
    dispatch(Chat2Gen.createSelectConversation({conversationIDKey, fromUser: true})),
})

const mergeProps = (stateProps, dispatchProps, ownProps) => {
  const isSelected = stateProps.isSelected
  const hasUnread = stateProps.hasUnread
  const styles = Constants.getRowStyles(stateProps._meta, isSelected, hasUnread)
  const participantNeedToRekey = stateProps._meta.rekeyers.size > 0
  const youNeedToRekey = !participantNeedToRekey && stateProps._meta.rekeyers.has(stateProps._username)

  return {
    backgroundColor: styles.backgroundColor,
    hasBadge: stateProps.hasBadge,
    hasResetUsers: !stateProps._meta.resetParticipants.isEmpty(),
    hasUnread,
    isMuted: stateProps._meta.isMuted,
    isSelected,
    // Don't allow you to select yourself
    onSelectConversation: isSelected ? () => {} : dispatchProps.onSelectConversation,
    participantNeedToRekey,
    participants: Constants.getRowParticipants(stateProps._meta, stateProps._username).toArray(),
    showBold: styles.showBold,
    snippet: stateProps._meta.snippet,
    subColor: styles.subColor,
    teamname: stateProps._meta.teamname,
    timestamp: Constants.timestampToString(stateProps._meta),
    usernameColor: styles.usernameColor,
    youAreReset: stateProps.youAreReset,
    youNeedToRekey,
  }
}

export default connect(mapStateToProps, mapDispatchToProps, mergeProps)(SmallTeam)
