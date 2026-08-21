package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"symmer/utils"
)

// Config type for json structure
type Config struct {
	CurrentConfigLocation string `json:"cfg"`
	FuturedSymLocation string `json:"des"`
}

// Mapped config of the json structure with a string acting as the title of the config
// Example: Opencode { "config": "location", "symlink": "location"}
type MapConfig map[string]Config

// Checks if the config for symmer exists, if not, it creates one in the current directory.
func ConfigCheck() *os.File {
	// Attempts to open the json file if it exists.
	file, err := os.Open("symmer.json")

	// Error check if it cannot be opened
	if err != nil {
		// If the file doesn't exist, then it'll create the file and exit
		if os.IsNotExist(err) {
			utils.SymmerPrint("Missing config, creating...")
			file = CreateConfig()
			utils.SymmerPrint("Config created, fill it out and rerun symmer.")
			os.Exit(0)
		}
	}
	return file
}

// Creates a symmer json config file when called.
func CreateConfig() *os.File {
	// Attempts to create config file
	file, err := os.Create("symmer.json")
	
	// if the file exists then it'll skip the error check
	if os.IsExist(err) {
		goto Return
	} else {
		s := "{}\n"
		file.Write([]byte(s))
	}

	// Presents the user with an error stating that the json config can't be created
	if err != nil {
		utils.SymmerError("Unable to create json config file in current directory. Review error below")
		utils.SymmerError(err.Error())
		os.Exit(1)
	}
	
	Return:
	return file
}

// Opens the json config file, decodes it using the Config type template, returns it as a typed object for parsing.
func OpenJSON(file *os.File) MapConfig {
	// Creates a config object to hold the future decoded json
	var cfg MapConfig
	
	// Opens and reads the json config
	data, err := os.ReadFile(file.Name())
	if err != nil {
		// If unable to read, symmer will present an error and quit.
		utils.SymmerError("Unable to read symmer config")
		utils.SymmerError(err.Error())
		os.Exit(1)
	}
	

	// Attempts to decode the json and presents the user with an error message if unable to create it.
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		utils.SymmerError("Error decoding json")
		utils.SymmerPrint("Check if your config has been filled out with applications and sym locations")
		os.Exit(1)
	}


	// Success message and returns the decoded config
	// Checks if the file is an empty config or not. If so, it'll display a different confirmation message.
	if string(data) == "{}\n" {
		utils.SymmerPrint("Config loaded successfully but it's empty.")
	} else {
		utils.SymmerPrint("Config loaded successfully.")
	}

	return cfg
}

// Creates symlinks based on the configuration options given in the json config file
func CreateSymlinks(configName string, configLocation string, configDestination string) error {

	// Searches through files given
	matches, err := filepath.Glob(configLocation)
	if err != nil {
		return fmt.Errorf("Invalid location given %q: %w", configLocation, err)
	}
	
	// If there is no files, it'll return an error.
	if len(matches) == 0 {
		return fmt.Errorf("no files matched %q", configLocation)
	}
	
	// Attempt to make a directory for the destination if it doesn't exist
	if err := os.MkdirAll(configDestination, 0755); err != nil {
		if os.IsExist(err) {
			utils.SymmerPrint("file/directory already exists")
		}
	}
	
	// Loop through all files and creates absolute paths for each found file/directory
	for _, source := range matches {
		absSource, err := filepath.Abs(source)
		if err != nil {
			return err
		}
		
		// Joins the absolute path and joins it with the config destination 
		destination := filepath.Join(
			configDestination,
			filepath.Base(source),
		)
		
		// Attempts to create a symlink for a config, returns an error if unable to create one
		if err := os.Symlink(absSource, destination); err != nil {
			utils.SymmerError(err)
		}
	}
	
	// Return nil to satisfy the functions return requirement
	return nil
}

// Runs the symlinking process for symmer.
func RunSymmer() {
	// Opens the file after it checks if it exists
	file := ConfigCheck()
	defer file.Close()
	
	// Assigns cfg as the parsed config file
	cfg := OpenJSON(file)

	// Loop over every entry in config file 
	for k, v := range cfg {
		// Attempts to create a symlink via the given paths
		err := CreateSymlinks(k, v.CurrentConfigLocation, v.FuturedSymLocation)
		
		// If an error occurs, prompts the user that symmer cant make the symlink and provides the error for the user to review
		if err != nil {
			utils.SymmerError("Unable to create symlink for " + k)
		} else {
			utils.SymmerPrint("Symlink for " + k + " created successfully!")
		}
	}
}

// Prints the symmer version to the user.
func GetSymmerVersion() {
	fmt.Printf("symmer version: %v\n", utils.GreenText(utils.SYMMER_VERSION))
}

// Prints the help menu for symmer.
func PrintSymmerHelp() {
	helpMenu := "usage: symmer [-h | --help] [-v | --version]"

	fmt.Println(helpMenu)
}

// Prints a usage message to the user if they input an unsupported argument.
func PrintMissInput(arg string) {
	fmt.Printf("Unknown argument %s.\n", arg)
	fmt.Println("usage: symmer [-h | --help] [-v | --version]")
}

func main() {
	args := os.Args

	if len(args) < 2 {
		RunSymmer()
		os.Exit(0)
	}

	switch args[1] {
	case "-v", "--version":
		GetSymmerVersion()
	case "-h", "--help":
		PrintSymmerHelp()
	default:
		PrintMissInput(args[1])
	}

}
