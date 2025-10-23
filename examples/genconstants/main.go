package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chrisjoyce911/active-campaign-sdk-go/genconstants"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	base := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")
	if base == "" || token == "" {
		log.Fatalf("ACTIVE_URL and ACTIVE_TOKEN must be set")
	}

	gen := genconstants.NewGenerator(base, token)
	gen.SetOutputPath("examples/gen-constants/constants.go")
	gen.SetPackageName("active")
	gen.SetMapPath("examples/gen-constants/.gen-constants.map.json")

	if err := gen.Generate(); err != nil {
		fmt.Println("Error generating constants:", err)
	} else {
		fmt.Println("generated", gen.OutPath)
	}
}
