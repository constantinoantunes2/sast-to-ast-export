package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/checkmarxDev/ast-sast-export/internal"
	"github.com/checkmarxDev/ast-sast-export/internal/app/export"
	"github.com/checkmarxDev/ast-sast-export/internal/app/logging"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	userArg                = "user"
	passArg                = "pass"
	urlArg                 = "url"
	exportArg              = "export"
	projectsActiveSinceArg = "projects-active-since"
	debugArg               = "debug"
	verboseArg             = "verbose"
	projectsIds            = "project-id"
	teamName               = "project-team"

	projectsActiveSinceDefaultValue = 180
)

// productName is defined in Makefile and initialized during build
var productName string

// productVersion is defined in VERSION and initialized during build
var productVersion string

// productBuild is defined in Makefile and initialized during build
var productBuild string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   productName,
	Short: "Exports SAST triaged results for importing in AST",
	Long: `Exports encrypted SAST triaged results for importing in AST. Example usage:

cxsast_exporter --user username --pass password --url http://localhost

Produces:
1. a zip file containing the encrypted SAST triaged results, e.g. cxsast_exporter-2021-09-10-15-42-35.zip
2. a key file containing base64 encoded password for decrypting the contents of zip file, e.g. cxsast_exporter-2021-09-10-15-42-35.txt
3. a log file with diagnostic information, e.g. cxsast_exporter-2021-09-10-15-42-35.log

NOTE the minimum supported SAST version is 9.3. SAST installations below this version should be upgraded in order to run this export tool. 
`,
	Run: func(cmd *cobra.Command, args []string) {
		// setup logging
		verbose, flagErr := cmd.Flags().GetBool(verboseArg)
		if flagErr != nil {
			panic(flagErr)
		}

		now := time.Now()
		logFileName := fmt.Sprintf("%s-%s.log", productName, now.Format(export.DateTimeFormat))
		logFileWriter, err := os.Create(logFileName)
		if err != nil {
			panic(err)
		}
		defer func() {
			if closeErr := logFileWriter.Close(); closeErr != nil {
				log.Debug().Err(closeErr).Msg("closing log file writer")
			}
		}()

		levelWriter := logging.NewMultiLevelWriter(verbose, zerolog.InfoLevel, logging.GetNewConsoleWriter(), logFileWriter)

		logInitErr := logging.Init(zerolog.LevelTraceValue, &levelWriter)
		if logInitErr != nil {
			panic(logInitErr)
		}

		defer func() {
			if r := recover(); r != nil {
				log.Error().Msgf("execution failed: %v", r)
				os.Exit(1)
			}
		}()

		// start export
		allArgs := GetArgs(cmd, productName)
		allArgs.RunTime = now
		exportErr := internal.RunExport(&allArgs)
		if exportErr != nil {
			log.Error().Err(exportErr)
			panic(exportErr)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

//nolint:gochecknoinits
func init() {
	projectsActiveSinceUsage := "include only triaged results from projects active in the last N days"

	rootCmd.Flags().StringP(userArg, "", "", "SAST username")
	rootCmd.Flags().StringP(passArg, "", "", "SAST password")
	rootCmd.Flags().StringP(urlArg, "", "", "SAST url")
	rootCmd.Flags().StringP(teamName, "", "", "Team name filter")
	rootCmd.Flags().StringP(projectsIds, "", "", "Project ID filter")
	rootCmd.Flags().StringSliceP(exportArg, "", export.GetOptions(), "SAST export options")
	rootCmd.Flags().IntP(projectsActiveSinceArg, "", projectsActiveSinceDefaultValue, projectsActiveSinceUsage)
	rootCmd.Flags().Bool(debugArg, false, "activate debug mode")
	rootCmd.Flags().BoolP(verboseArg, "v", false, "enable verbose logging to console")

	if err := rootCmd.MarkFlagRequired(userArg); err != nil {
		panic(err)
	}
	if err := rootCmd.MarkFlagRequired(passArg); err != nil {
		panic(err)
	}
	if err := rootCmd.MarkFlagRequired(urlArg); err != nil {
		panic(err)
	}
	if err := rootCmd.MarkFlagCustom(projectsActiveSinceArg, projectsActiveSinceUsage); err != nil {
		panic(err)
	}
}
