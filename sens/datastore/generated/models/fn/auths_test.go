package fn

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestInsertAuth(t *testing.T) {
	data := []byte(`{
		"Email": "email3",
		"Mobile": "mobile3",
		"Social": "social3",
		"Properties": {
			"x": "y"
		}
	}`)
	fmt.Println(InsertAuth(data))
}

func TestUpdateAuth(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	data := []byte(`{
		"Email": "mayank@dozee.io",
		"Mobile": "8309284201"
	}`)
	UpdateAuth("cb6bc6d9-167c-48dd-b2fb-a794c58f1d92", data)
}

func TestFindAuth(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	a, e := FindAuth(
		// []string{"Email:mayank@dozee.io", "Email:emayank@gmail.com", "Mobile:7032806003"},
		// []string{"Social:social", "Mobile:8309284201"},
		nil,
		nil,
		[]string{"Mobile:0:9000000000", "Email:a:mayank@dozee.io"},
		"4", "id", "DESC",
	)
	j, ee := json.Marshal(a)
	fmt.Printf("%s, %v, %v\n", j, e, ee)
}
