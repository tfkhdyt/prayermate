/*
Copyright © 2023 Taufik Hidayat <tfkhdyt@proton.me>
*/package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "prayermate",
	Short: "Command line based Muslim prayer reminder",
	Long:  `PrayerMate is a CLI-based application to remind users of prayer times.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.prayermate.yaml)")
	rootCmd.PersistentFlags().
		StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/prayermate/config.toml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	// Find configDir directory.
	configDir, err := os.UserConfigDir()
	cobra.CheckErr(err)

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".cobra" (without extension).
		viper.SetConfigName("config") // name of config file (without extension)
		viper.SetConfigType("toml")   // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath(".")      // optionally look for config in the working directory
		viper.AddConfigPath(filepath.Join(configDir, "prayermate"))
	}
	viper.SetDefault("location.id", "1301")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Error: failed to read config")
	}

	if _, err := os.Stat(filepath.Join(configDir, "prayermate/config.toml")); os.IsNotExist(err) {
		createConfig()
	}
}

func createConfig() {
	// Create the directory and file paths.
	configDir, err := os.UserConfigDir()
	cobra.CheckErr(err)
	dirPath := filepath.Join(configDir, "prayermate")
	filePath := filepath.Join(dirPath, "config.toml")

	// Create the directory if it does not exist.
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0o755); err != nil {
			log.Fatalln("Error: failed to make config dir")
		}
	}

	// Create the file if it does not exist.
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
}
