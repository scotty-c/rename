package cmd

import (
	"fmt"
	"os"

	"rename/internal/rename"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rename [file or directory]",
	Short: "Rename strings in files based on a list of replacements",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]

		// Load replacements from Viper
		replacements := rename.LoadConfig()
		if replacements == nil {
			return fmt.Errorf("failed to load config")
		}

		// Check if the path is a file or directory
		fileInfo, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("failed to access path: %w", err)
		}

		if fileInfo.IsDir() {
			// Process the directory
			if err := rename.ProcessDirectory(path, replacements); err != nil {
				return fmt.Errorf("error processing directory: %w", err)
			}
		} else {
			// Process the single file
			if err := rename.ProcessFile(path, replacements); err != nil {
				return fmt.Errorf("error processing file: %w", err)
			}
		}

		fmt.Println("Replacements applied successfully!")
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) // Exit gracefully with a non-zero status code
	}
}
