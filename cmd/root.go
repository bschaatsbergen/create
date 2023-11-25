package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/bschaatsbergen/create/pkg/core"
	"github.com/bschaatsbergen/create/pkg/model"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type PlainFormatter struct{}

var (
	version   string
	flagStore model.Flagstore

	rootCmd = &cobra.Command{
		Use:     "create",
		Short:   "create - A modern UNIX file generation tool",
		Version: version,
		PreRun:  toggleDebug,
		Example: "create bar.txt " +
			"\n" + "create foo/bar.txt",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("error: provide a single argument")
				fmt.Println("See 'create -h' for help and examples")
				os.Exit(1)
			}
			fileName := args[0]
			core.CreateFile(fileName, flagStore)
		},
	}
)

func init() {
	setupCobraUsageTemplate()
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().BoolVarP(&flagStore.Debug, "debug", "d", false, "verbose logging")
}

func setupCobraUsageTemplate() {
	cobra.AddTemplateFunc("StyleHeading", color.New(color.FgWhite).SprintFunc())
	usageTemplate := rootCmd.UsageTemplate()
	usageTemplate = strings.NewReplacer(
		`Usage:`, `{{StyleHeading "Usage:"}}`,
		`Examples:`, `{{StyleHeading "Examples:"}}`,
		`Flags:`, `{{StyleHeading "Flags:"}}`,
	).Replace(usageTemplate)
	rootCmd.SetUsageTemplate(usageTemplate)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (f *PlainFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s\n", entry.Message)), nil
}

func toggleDebug(cmd *cobra.Command, args []string) {
	if flagStore.Debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug logs enabled")
		log.SetFormatter(&log.TextFormatter{})
	} else {
		plainFormatter := new(PlainFormatter)
		log.SetFormatter(plainFormatter)
	}
}
