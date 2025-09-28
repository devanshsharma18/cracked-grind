package main

import (
	"flag"
	"log"
	"os"
)

type ServicePackage struct {
	path string
}

func main() {
	var servicePackage ServicePackage
	flag.StringVar(&servicePackage.path, "file", "test.md", "usage for path")
	flag.Parse()
	fileInfo, err := os.Stat(servicePackage.path)
	if err != nil {
		log.Fatal("File does not exist")
	}
	log.Printf("The file name is: %s", fileInfo.Name())

}
