package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type config struct {
	cfgFile  string
	logLevel string
	profile  string
}

var cfg config

var rootCmd = &cobra.Command{
	Use:   "ovh-cli",
	Short: "CLI to interact with OVH me API",
	Long: `This CLI allow to manipulate OVH me API.
	
ovh-cli is a tool for manage OVH resources.
To discover OVH API, see: https://api.ovh.com/console/#`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Load configuration file: ", viper.ConfigFileUsed())
		cmd.Help()
	},
}

// Execute root command
func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfg.cfgFile, "config", "", "config file (default is $HOME/.ovh-cli.toml)")
	rootCmd.PersistentFlags().StringVarP(&cfg.profile, "profile", "p", "default", "Configuration profile to use")
	rootCmd.PersistentFlags().StringVar(&cfg.logLevel, "log-level", "info", "Log level. Expect: debug, info, warn, error")

	viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("profile"))
	viper.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level"))
}

func configureLogLevel() {
	switch viper.GetString("log-level") {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}
}

func initConfig() {
	if cfg.cfgFile != "" {
		viper.SetConfigFile(cfg.cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.SetConfigName(".ovh-cli")
		viper.AddConfigPath("/etc/ovh-cli/")
		viper.AddConfigPath("/usr/local/etc/ovh-cli/")
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Debugln("Using config file:", viper.ConfigFileUsed())
	}
	configureLogLevel()
}
