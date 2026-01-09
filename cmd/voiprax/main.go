package main

import (
	"fmt"
	"os"

	"github.com/ismailtsdln/VoIPrax/internal/logger"
	"github.com/ismailtsdln/VoIPrax/internal/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "voiprax",
		Short: "VoIPrax - Modern VoIP Penetration & Analysis Toolkit",
		Long: `VoIPrax is a high-performance, modular toolkit for VoIP security analysis,
fuzzing, and penetration testing.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			ui.PrintBanner()
		},
	}
	cfgFile string
	verbose bool
	log     *logger.Logger
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.voiprax.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".voiprax")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	logLevel := "info"
	if verbose {
		logLevel = "debug"
	}
	log = logger.New(logLevel, true)
	logger.InitGlobal(logLevel, true)
}
