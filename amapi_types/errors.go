package amapi_types

import "fmt"

type AmAPIError struct {
	Details string `json:"details"`
}

func (a AmAPIError) Error() string {
	return fmt.Sprintf("%s", a.Details)
}
