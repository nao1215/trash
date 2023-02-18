// Package cmd define subcommands provided by the gup command
package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "trash",
	Short: `The trash command move(delete) file/directory at $XDG_DATA_HOME/Trash/files.
These files can be listed, restore, or cleaned from the trash can.`,
}

// OsExit is wrapper for  os.Exit(). It's for unit test.
var OsExit = os.Exit

// Execute run gup process.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

var (
	// Version value is set by ldflags
	Version string
	// Name is cli command name
	Name = "trash"
)

// getVersion return gup command version.
// Version global variable is set by ldflags.
func getVersion() string {
	version := "unknown"
	if Version != "" {
		version = Version
	} else if buildInfo, ok := debug.ReadBuildInfo(); ok {
		version = buildInfo.Main.Version
	}
	return fmt.Sprintf("%s version %s (under MIT LICENSE)", Name, version)
}
