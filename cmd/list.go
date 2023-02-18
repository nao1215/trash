package cmd

import (
	"fmt"
	"os"

	"github.com/nao1215/trash/go-trash"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List file/directory at $XDG_DATA_HOME/Trash/files.",
	Long:    `List file/directory at $XDG_DATA_HOME/Trash/files.`,
	Example: `  trash list`,

	Run: func(cmd *cobra.Command, args []string) {
		OsExit(list(cmd, args))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) int {
	trash, err := trash.NewTrash()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	files, err := trash.List()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	for _, v := range files {
		fmt.Fprintln(os.Stdout, v)
	}
	return 0
}
