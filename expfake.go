package main

import (
	"crypto/rand"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/jhoblitt/expfake/filesizes"
	"github.com/jhoblitt/expfake/hostmap"
)

func main() {
	hostName := flag.String("hostname", "", "name of host to generate test file data for")
	hostMapPath := flag.String("hostmap", "hostmap.yaml", "Path to the hostmap file")
	fileSizesPath := flag.String("filesizes", "filesizes.yaml", "Path to the file sizes file")
	outputDir := flag.String("dir", "", "Path to the directory to write files to")
	flag.Parse()

	if *hostName == "" {
		log.Fatal("hostname flag is required")
	}

	if *hostMapPath == "" {
		log.Fatal("hostmap flag is required")
	}

	if *fileSizesPath == "" {
		log.Fatal("filesizes flag is required")
	}

	if *outputDir == "" {
		log.Fatal("filesizes flag is required")
	}

	hMap := hostmap.Parse(hostMapPath)
	// log.Println(hMap)

	fSizes := filesizes.Parse(fileSizesPath)
	// log.Println(fSizes)

	hostFiles, ok := hMap.Hosts[*hostName]
	if !ok {
		log.Fatalf("No files found for host %s", *hostName)
	}

	// create output directory if it doesn't exist
	if err := os.MkdirAll(*outputDir, 0o755); err != nil {
		log.Fatalf("mkdir %s: %v", *outputDir, err)
	}

	totalSize := int64(0)
	for _, file := range hostFiles {
		size, ok := fSizes.Hosts[file]["max"]
		if !ok {
			log.Fatalf("No size found for file %s on host %s", file, *hostName)
		}

		totalSize += size

		fullFilePath := filepath.Join(*outputDir, file)
		f, err := os.Create(fullFilePath)
		if err != nil {
			log.Fatalf("Error creating file %s: %v", file, err)
		}
		defer func() {
			err := f.Close()
			if err != nil {
				log.Fatalf("Error closing file %s: %v", file, err)
			}
		}()

		// fill file with random data to make it difficult to compress
		if _, err := io.CopyN(f, rand.Reader, size); err != nil && err != io.EOF {
			log.Fatalf("Error writing to file %s: %v", file, err)
		}

		log.Printf("Generated file %s with size %d bytes", fullFilePath, size)
	}

	log.Printf("Total size of files generated for host %s: %d bytes", *hostName, totalSize)
}
