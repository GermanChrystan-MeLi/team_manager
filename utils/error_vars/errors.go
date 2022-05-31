package error_vars

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound           = errors.New("product not found")
	ErrRetrievingData     = errors.New("there was an error while retrieving data")
	ErrActionNotPerformed = errors.New("the action could not be performed")
	ErrDatabase           = errors.New("there has been an error with the Database")
	ErrWrongCredentials   = errors.New("wrong credentials")
)

func XNotFound(x string) error {
	errorStr := fmt.Sprintf("%s not found", x)
	return errors.New(errorStr)
}
