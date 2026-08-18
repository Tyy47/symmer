package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)


type Config struct {
	CurrentConfigLocation string `json:"config"`
	FuturedSymLocation string `json:"symlink"`
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

func CreateSymLink(config string, symlinkLocation string) error {
	if err := os.Symlink(config, symlinkLocation); err != nil {
		return err
	} else {
		return nil
	}
}

func main() {

	file := ConfigCheck()
	defer file.Close()

	cfg := OpenJSON(file)

	for k, v := range cfg {
		err := CreateSymLink(v.CurrentConfigLocation, v.FuturedSymLocation)

		if err != nil {
			fmt.Printf("symmer: Unable to create symlink for %s. Check both locations for potential errors.", k)
			continue
		} else {
			fmt.Printf("symmer: Config %s symlinked successfully!\n%s ==> %s\n", k, v.CurrentConfigLocation, v.FuturedSymLocation)
		}
	}

}
