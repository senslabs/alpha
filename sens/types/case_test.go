package types

import (
	"fmt"
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	fmt.Println(ToSnakeCase("UserAuth"))
}
