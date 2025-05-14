package utils

import "socialNetwork/entity"

func SetOnlineStatus(contacts []entity.User, isOnline func(uint) bool) []entity.User {
	for i := range contacts {
		if isOnline(contacts[i].ID) {
			contacts[i].Online = true
		} 
	}
	return contacts
}
