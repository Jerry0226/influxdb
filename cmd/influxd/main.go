package main

import (
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/influxdata/flux"
	"github.com/influxdata/influxdb"
	"github.com/influxdata/influxdb/cmd/influxd/generate"
	"github.com/influxdata/influxdb/cmd/influxd/inspect"
	"github.com/influxdata/influxdb/cmd/influxd/launcher"
	_ "github.com/influxdata/influxdb/query/builtin"
	_ "github.com/influxdata/influxdb/tsdb/tsi1"
	_ "github.com/influxdata/influxdb/tsdb/tsm1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "influxd",
	Short: "Influx Server",
}

func init() {
	influxdb.SetBuildInfo(version, commit, date)
	viper.SetEnvPrefix("INFLUXD")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	rootCmd.InitDefaultHelpCmd()

	rootCmd.AddCommand(launcher.NewCommand())
	rootCmd.AddCommand(generate.Command)
	rootCmd.AddCommand(inspect.NewCommand())

	// TODO: this should be removed in the future: https://github.com/influxdata/influxdb/issues/16220
	if os.Getenv("QUERY_TRACING") == "1" {
		flux.EnableExperimentalTracing()
	}
}

// find determines the default behavior when running influxd.
// Specifically, find will return the influxd run command if no sub-command
// was specified.
func find(args []string) *cobra.Command {
	cmd, _, err := rootCmd.Find(args)
	if err == nil && cmd == rootCmd {
		// Execute the run command if no sub-command is specified
		return launcher.NewCommand()
	}

	return rootCmd
}

func main() {
	cmd := find(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
