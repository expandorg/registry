package authorization

import (
	"os"
	"strconv"

	"github.com/gemsorg/registry/pkg/authentication"
)

type Authorizer interface {
	SetAuthData(data authentication.AuthData)
	IsModerator() (bool, error)
}

type authorizor struct {
	authData authentication.AuthData
}

func NewAuthorizer() Authorizer {
	return &authorizor{
		authentication.AuthData{},
	}
}

func (a *authorizor) SetAuthData(data authentication.AuthData) {
	a.authData = data
}

func (a *authorizor) IsModerator() (bool, error) {
	moderatorID, err := strconv.ParseUint(os.Getenv("MODERATOR_ID"), 10, 64)
	if err != nil || a.authData.UserID != moderatorID {
		return false, UnauthorizedAccess{}
	}
	return true, nil
}
