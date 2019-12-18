package authorization

import "fmt"

type UnauthorizedAccess struct{}

func (err UnauthorizedAccess) Error() string {
	return fmt.Sprint("You are not authorized!")
}
