package cmd

import (
	"io"
	"log"
	"os"

	"github.com/GoToolSharing/htb-cli/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "htb-cli",
	Short: "CLI enhancing the HackTheBox user experience.",
	Long:  `This software, engineered using the Go programming language, serves to streamline and automate various tasks for the HackTheBox platform, enhancing user efficiency and productivity.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if config.GlobalConfig.Verbose {
			log.SetOutput(os.Stdout)
		} else {
			log.SetOutput(io.Discard)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().BoolVarP(&config.GlobalConfig.Verbose, "verbose", "v", false, "Verbose mode")
	rootCmd.PersistentFlags().StringVarP(&config.GlobalConfig.ProxyParam, "proxy", "p", "", "Configure a URL for an HTTP proxy")
	rootCmd.PersistentFlags().BoolVarP(&config.GlobalConfig.BatchParam, "batch", "b", false, "Don't ask questions")
}
