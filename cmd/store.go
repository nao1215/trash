package cmd

import (
	"fmt"
	"os"

	"github.com/nao1215/trash/go-trash"
	"github.com/spf13/cobra"
)

var storeCmd = &cobra.Command{
	Use:   "put",
	Short: "Store file/directory at $XDG_DATA_HOME/Trash/files. only move, not delete",
	Long: `Store file/directory at $XDG_DATA_HOME/Trash/files.
Only move, not delete`,
	Example: `  trash put PATH/TO/FILE(S)`,

	Run: func(cmd *cobra.Command, args []string) {
		OsExit(store(cmd, args))
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}

func store(cmd *cobra.Command, args []string) int {
	trash, err := trash.NewTrash()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	result := 0
	for _, v := range args {
		if err = trash.Trash(v); err != nil {
			fmt.Fprintln(os.Stderr, err)
			result = 1
		}
	}
	return result
}
