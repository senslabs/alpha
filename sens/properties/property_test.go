package properties_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/senslabs/alpha/sens/properties"
)

func TestGetProperty(t *testing.T) {
	os.Setenv("ENV", "dev")
	r := properties.GetProperties().Get("test")
	fmt.Println(r)
}
