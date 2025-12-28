package main

import (
	"fmt"
	"os"
	"github.com/NoTuxAllowed/dragon-scheduler/internal/config"
)

func main() {
	fmt.Println("this is dragonctl")
	filePath  := "/home/bigpod/test.yaml"
	data, _ := os.ReadFile(filePath)
	test, err_:= config.LoadManifest(data)
	fmt.Println(test)
	fmt.Println(err_.Error())
}