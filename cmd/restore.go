package cmd

import (
	"fmt"
	"os"

	"github.com/nao1215/trash/go-trash"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore file/directory from $XDG_DATA_HOME/Trash/files to original location.",
	Long: `Restore file/directory from $XDG_DATA_HOME/Trash/files to original location.
User specify trashed file name in trash can.`,
	Example: `  trash restore TRASHED_FILE_NAME(s)`,

	Run: func(cmd *cobra.Command, args []string) {
		OsExit(restore(cmd, args))
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}

func restore(cmd *cobra.Command, args []string) int {
	trash, err := trash.NewTrash()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	result := 0
	for _, v := range args {
		if err = trash.Restore(v, false); err != nil {
			fmt.Fprintln(os.Stderr, err)
			result = 1
		}
	}
	return result
}
