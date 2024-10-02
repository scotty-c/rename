package cmd

import (
	"fmt"

	"github.com/scotty-c/rename/internal/rename"

	"github.com/spf13/cobra"
)

// changeCmd represents the change command
var changeCmd = &cobra.Command{
	Use:   "change [file or directory]",
	Short: "Replace strings in files based on a list from the config file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]

		// Load replacements from the config
		replacements, err := rename.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config replacements: %w", err)
		}

		// Perform replacements in the directory or file
		err = rename.ProcessDirectory(path, replacements)
		if err != nil {
			return fmt.Errorf("failed to process directory: %w", err)
		}

		fmt.Println("Replacement process completed.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(changeCmd)
}
