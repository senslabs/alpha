package properties

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/senslabs/alpha/sens/types"
)

var instance *properties
var once sync.Once

type properties struct {
	*types.Map
}

func GetProperties() *properties {
	once.Do(func() {
		env := os.Getenv("ENV")
		files, err := filepath.Glob("/var/sens/config/*." + env + ".json")
		if err != nil {
			log.Fatal(err)
		}
		instance = &properties{types.Create()}
		for _, f := range files {
			content, err := ioutil.ReadFile(f)
			if err != nil {
				log.Fatal(err)
			}
			var m types.Map
			err = json.Unmarshal(content, &m)
			if err != nil {
				log.Fatal(err)
			}
			svcname := m["svcname"]
			if svcname == "" {
				log.Fatal("Service Name (\"svcname\" field) is mandatory in config")
			}
			instance.Put(svcname.(string), m)
		}
	})
	return instance
}

func GetServiceName() string {
	svcname := os.Getenv("SVCNAME")
	if svcname == "" {
		svcname = strings.ToLower(filepath.Base(os.Args[0]))

	}
	return svcname
}

func SetServiceName(svcname string) {
	os.Setenv("svcname", svcname)
}

func (this *properties) GetValue(keys ...string) interface{} {
	var k []string
	k = append(k, GetServiceName())
	k = append(k, keys...)
	return this.Get(k...)
}
