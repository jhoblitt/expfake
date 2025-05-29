package filesizes

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type FileSizes struct {
	Hosts map[string]map[string]int64 `yaml:"files"`
}

func Parse(fileSizesPath *string) *FileSizes {
	fileSizesRaw, err := os.ReadFile(*fileSizesPath)
	if err != nil {
		log.Fatal(err)
	}

	fSizes := &FileSizes{}

	err = yaml.Unmarshal([]byte(fileSizesRaw), &fSizes)
	if err != nil {
		log.Fatal(err)
	}

	return fSizes
}
