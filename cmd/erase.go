package cmd

import (
	"fmt"
	"os"

	"github.com/nao1215/trash/go-trash"
	"github.com/spf13/cobra"
)

var eraseCmd = &cobra.Command{
	Use:     "erase",
	Short:   "erase file/directory at $XDG_DATA_HOME/Trash/files.",
	Long:    `erase file/directory at $XDG_DATA_HOME/Trash/files.`,
	Example: `  trash erase FILE_NAME(s)`,

	Run: func(cmd *cobra.Command, args []string) {
		OsExit(erase(cmd, args))
	},
}

func init() {
	rootCmd.AddCommand(eraseCmd)
}

func erase(cmd *cobra.Command, args []string) int {
	trash, err := trash.NewTrash()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	result := 0
	for _, v := range args {
		if err := trash.Erase(v); err != nil {
			fmt.Fprintln(os.Stderr, err)
			result = 1
		}
	}
	return result
}
