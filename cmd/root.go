package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rename [file]",
	Short: "Rename strings in a file based on a list of replacements",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]

		// Load replacements from Viper
		replacements := viper.GetStringMapString("replacements")
		if len(replacements) == 0 {
			log.Fatal("No replacements found in config")
		}

		// Read the file content
		content, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}

		// Perform replacements
		updatedContent := string(content)
		for old, new := range replacements {
			updatedContent = strings.ReplaceAll(updatedContent, old, new)
		}

		// Save the updated content back to the file
		err = os.WriteFile(fileName, []byte(updatedContent), 0644)
		if err != nil {
			log.Fatalf("Failed to write updated file: %v", err)
		}

		fmt.Println("Replacements applied successfully!")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
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
}
