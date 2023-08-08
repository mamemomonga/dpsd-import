package main

import (
	"os"

	"github.com/spf13/viper"
)

func viperConfigInit() {
	viperConfigLoad()
	if !viper.IsSet("OutputDir") {
		s := readStdinInputText("Output directory")
		viper.Set("OutputDir", s)
		viperConfigSave()
		os.Exit(0)
	}
}

func viperConfigLoad() bool {
	viper.SetConfigName("/.dpsdimport")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(os.Getenv("HOME"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return false
		}
	}
	return true
}

func viperConfigSave() {
	f := viper.ConfigFileUsed()
	if f == "" {
		viper.SafeWriteConfig()
	} else {
		viper.WriteConfig()
	}
	viper.ReadInConfig()
}
