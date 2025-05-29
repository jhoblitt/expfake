package hostmap

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type HostMap struct {
	Hosts map[string][]string `yaml:"hosts"`
}

func Parse(hostMapPath *string) *HostMap {
	hostMapRaw, err := os.ReadFile(*hostMapPath)
	if err != nil {
		log.Fatal(err)
	}

	hMap := &HostMap{}

	err = yaml.Unmarshal([]byte(hostMapRaw), &hMap)
	if err != nil {
		log.Fatal(err)
	}

	return hMap
}
