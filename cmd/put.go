package cmd

import (
	"fmt"
	"os"

	"github.com/nao1215/trash/go-trash"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:     "put",
	Short:   "put sub-command move file/directory under $XDG_DATA_HOME/Trash/files.",
	Long:    `put sub-command move file/directory under $XDG_DATA_HOME/Trash/files `,
	Example: `  trash put path/to/file`,

	Run: func(cmd *cobra.Command, args []string) {
		OsExit(put(cmd, args))
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}

func put(cmd *cobra.Command, args []string) int {
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
