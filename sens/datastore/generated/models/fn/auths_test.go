package fn

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
)

func TestInsertAuth(t *testing.T) {
	os.Setenv("COCKROACH_HOST", "localhost")
	u, _ := uuid.NewRandom()
	data := []byte(fmt.Sprintf(`{
		"Id": "%s",
		"Email": "email9",
		"Mobile": "mobile9",
		"Social": "social9",
		"Properties": {
			"x": "y"
		}
	}`, u.String()))
	fmt.Println(InsertAuth(data))
}

func TestUpdateAuth(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	data := []byte(`{
		"Properties": {"a":"b"},
		"FirstName": "Mayank"
	}`)
	UpdateAuth("50db7414-e07c-4f2c-a6e8-a0d7c38a9dd1", data)
}

func TestFindAuth(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	a, e := FindAuth(
		[]string{"Email:mayank@dozee.io", "Email:emayank@gmail.com", "Mobile:mobile9"},
		// []string{"Social:social", "Mobile:8309284201"},
		// nil,
		nil,
		[]string{"Mobile:0:9000000000", "Email:a:email9"},
		"4", "id", "DESC",
	)
	j, ee := json.Marshal(a)
	fmt.Printf("%s, %v, %v\n", j, e, ee)
}
