package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	programName = "shoot-watchdog"
)

var (
	rootCmd = &cobra.Command{
		Use:          programName,
		SilenceUsage: true,
	}
)

// Execute is the entrypoint of the cient-go application
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		if viper.GetBool("debug") {
			st := errors.WithStack(err)
			fmt.Printf("%+v", st)
		}
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().Duration("checkinterval", 60*time.Second, "time between nodeReady checks")
	rootCmd.PersistentFlags().Int("depl-max-fails", 5, "number of checks till shoot is marked as failed and restarted")

	rootCmd.AddCommand(checkShoot)

	err := viper.BindPFlags(rootCmd.PersistentFlags())
	if err != nil {
		log.Fatalf("error setup root cmd:%v", err)
	}
}

func initConfig() {
	viper.SetEnvPrefix(strings.ToUpper(programName))
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
}
