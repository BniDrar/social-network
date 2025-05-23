package utils

import (
	"errors"

	"socialNetwork/entity"
)

func ValidGroupCredentials(g entity.Group) error {
	if  g.Type > entity.FakeGroup ||g.Type < entity.RealGroup ||
			g.Name == "" || g.Description == ""  {
		return errors.New("invalid credentials")
	}
	return nil
}

func ValidateEventCredentials(event entity.Event) bool {
	return event.GroupID > 0 && IsValidName(event.Location) && IsValidText(event.Description)
}