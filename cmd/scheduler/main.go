package main

import (
	"fmt"

	_ "github.com/NoTuxAllowed/dragon-scheduler/internal/providers/aws"
	_ "github.com/NoTuxAllowed/dragon-scheduler/internal/providers/azure"
	_ "github.com/NoTuxAllowed/dragon-scheduler/internal/providers/gcp"
)

func main() {
	fmt.Println("this is scheduler")
}
