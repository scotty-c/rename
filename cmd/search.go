package cmd

import (
	"fmt"

	"github.com/scotty-c/rename/internal/rename"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [file or directory]",
	Short: "Search for strings in files based on a list from the config file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]

		// Load search keys from the config
		searchKeys, err := rename.LoadConfigKeys()
		if err != nil {
			return fmt.Errorf("failed to load config keys: %w", err)
		}

		// Search directory or file for files containing any of the keys
		files, err := rename.SearchFilesInDirectory(path, searchKeys)
		if err != nil {
			return fmt.Errorf("failed to search directory: %w", err)
		}

		// Print the list of files containing the keys
		fmt.Println("Files containing the search keys:")
		for _, file := range files {
			fmt.Println(file)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
