package main

import (
	"github.com/mamemomonga/dpsd-import/dpsd"
	"github.com/spf13/viper"
)

func main() {
	viperConfigInit()

	dpsd.NewDPSD(dpsd.NewDPSDInput{
		Model:     "DP3M",
		InputDir:  "/Volumes/SIGMA DP3M",
		OutputDir: viper.GetString("OutputDir"),
	}).Run()

	dpsd.NewDPSD(dpsd.NewDPSDInput{
		Model:     "DP2M",
		InputDir:  "/Volumes/SIGMA DP2M",
		OutputDir: viper.GetString("OutputDir"),
	}).Run()

	dpsd.NewDPSD(dpsd.NewDPSDInput{
		Model:     "DP1M",
		InputDir:  "/Volumes/SIGMA DP1M",
		OutputDir: viper.GetString("OutputDir"),
	}).Run()

	dpsd.NewDPSD(dpsd.NewDPSDInput{
		Model:     "DP1Q",
		InputDir:  "/Volumes/SIGMA DP1Q",
		OutputDir: viper.GetString("OutputDir"),
	}).Run()

}
