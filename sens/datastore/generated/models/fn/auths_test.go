package fn

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestInsertAuth(t *testing.T) {
	os.Setenv("COCKROACH_HOST", "localhost")
	data := []byte(`[{
		"Email": "email3",
		"Mobile": "mobile3",
		"Social": "social3",
		"Psroperties": {
			"x": "y"
		}
	},
	{
		"Email": "email4",
		"Mobile": "mobile4",
		"Social": "social4",
		"Psroperties": {
			"x1": "y1"
		}
		}]`)
	fmt.Println(BatchInsertAuth(data))
}

func TestUpdateAuth(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	data := []byte(`{
		"Properties": {"a":"b"},
		"FirstName": "Mayank"
	}`)
	UpdateAuth("a43c8563-8a50-42a4-8493-77ee83526248", data)
}

func TestFindAuth(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	a, e := FindAuth(
		[]string{"Email:mayank@dozee.io", "Email:emayank@gmail.com", "Mobile:mobile3"},
		// []string{"Social:social", "Mobile:8309284201"},
		// nil,
		nil,
		[]string{"Mobile:0:9000000000", "Email:a:mayank@dozee.io"},
		"4", "id", "DESC",
	)
	j, ee := json.Marshal(a)
	fmt.Printf("%s, %v, %v\n", j, e, ee)
}
