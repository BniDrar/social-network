package notification

import (
	"context"
	"errors"
	"net/http"

	"socialNetwork/entity"
)

func (n *Notification) ProcessVoteOnNotification(ctx context.Context, notf entity.Notification) (int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	exist, err := n.GroupContainsMember(notf.GroupId, userId)
	if !exist || err != nil {
		if err!= nil {
			return http.StatusInternalServerError, err
		}
		return http.StatusForbidden, errors.New("forbidden access to this action")
	}

	switch notf.Type {
	case entity.FollowingNotification:
		if notf.Accepted {
			// add the user to the following list in data base

		}
		// n.addResponseToNotification()
	case entity.EventNotification:
		// add the react to the event 
	case entity.GroupInvitationNotification  :
		if notf.Accepted {
			// add the invited to the group 
		}
	default:
		return http.StatusBadRequest, errors.New("invalid data")
	}

	return http.StatusOK, nil
}
