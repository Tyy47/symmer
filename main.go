package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)


type Config struct {
	CurrentConfigLocation string `json:"config_location"`
	FuturedSymLocation string `json:"symlink_location"`
}

type MapConfig map[string]Config

func ConfigCheck() *os.File {
	file, err := os.Open("symmer.json")
	if err != nil {
		fmt.Println("Missing config, creating...")
		file, _ = os.Create("symmer.json")
		fmt.Println("Config created, fill it out and rerun symmer.")
		os.Exit(0)
	}

	return file
}

func OpenJSON(file *os.File) MapConfig {
	var cfg MapConfig
	
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		log.Fatalf("Error decoding json: %+v\n", err)
		fmt.Println("Check if your config has been filled out with applications and sym locations")
	}

	fmt.Printf("Config loaded: %+v\n", cfg)

	return cfg
}

func main() {

	file := ConfigCheck()
	defer file.Close()

	cfg := OpenJSON(file)

	for k, v := range cfg {
		fmt.Printf("Config: %v has been symmed to %v\n", k, v.FuturedSymLocation)
	}
}
