package apierror

import (
	"fmt"
	"strings"
)

type MissingParameters struct {
	Params []string
}

func (err *MissingParameters) Error() string {
	missing := strings.Join(err.Params, ",")
	return fmt.Sprintf("Missing required params: %s", missing)
}
