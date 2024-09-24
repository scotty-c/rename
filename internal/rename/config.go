package rename

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// LoadConfig loads the replacement configuration from the YAML file
func LoadConfig() map[string]string {
	// Setup Viper
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to find home directory: %v", err)
	}

	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(home + "/.rename")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	replacements := viper.GetStringMapString("replacements")
	if len(replacements) == 0 {
		log.Fatal("No replacements found in config")
	}

	return replacements
}
