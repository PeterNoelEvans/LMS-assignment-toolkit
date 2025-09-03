package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "assignment-toolkit",
	Short: "A powerful CLI tool for creating and managing portable assignments",
	Long: `Assignment Toolkit is a comprehensive CLI tool for creating, validating, 
and managing portable assignments for your LMS. Create assignments offline 
and sync them when you have connectivity.

Features:
- Create assignments with interactive wizards
- Validate assignment packages
- Export/import assignment bundles
- Sync with remote LMS
- Template management`,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
