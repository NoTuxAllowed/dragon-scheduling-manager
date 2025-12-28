package main

import (
	"fmt"
	"os"
	"github.com/NoTuxAllowed/dragon-scheduler/internal/config"
)

func main() {
	fmt.Println("this is dragonctl")
	filePath  := "/home/bigpod/test.yaml"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	test, err := config.LoadManifest(data)
	if err != nil {
		fmt.Printf("Error loading manifest: %v\n", err)
		return
	}
	fmt.Printf("Kind: %s\n", test.Kind)
	fmt.Printf("Name: %s\n", test.Metadata.Name)
	fmt.Printf("Full Object: %+v\n", test)
}
