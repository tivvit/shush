package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tivvit/shush/server"
	"os"
)

var (
	GitCommit string
	GitTag    = "unknown"
)

const confFileFlag = "confFile"
const shushName = "Self Hosted Url SHortener"

var rootCmd = &cobra.Command{
	Use:     "shush",
	Short:   shushName,
	Long:    shushName + " (https://github.com/tivvit/shush)",
	Args:    cobra.OnlyValidArgs,
	Version: GitTag + " " + GitCommit,
	RunE: func(cmd *cobra.Command, args []string) error {
		confFile, err := cmd.Flags().GetString(confFileFlag)
		if err != nil {
			return err
		}
		server.Server(confFile)
		return nil
	},
}

func init() {
	rootCmd.Flags().StringP(confFileFlag, "c", "", "Configuration file path")
	err := rootCmd.MarkFlagFilename(confFileFlag)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
