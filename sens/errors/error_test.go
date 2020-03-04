package errors_test

import (
	"fmt"
	"testing"

	"github.com/senslabs/alpha/sens/errors"
)

func TestNew(t *testing.T) {
	err := errors.New(100, "Pata nahin")
	fmt.Println(errors.GetErrorCode(err), err.Error())
}
