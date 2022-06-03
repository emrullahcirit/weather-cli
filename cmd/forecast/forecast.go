/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package forecast

import (
	"fmt"

	"github.com/emrullahcirit/weather-cli/controller"
	"github.com/spf13/cobra"
)

var (
	city string
)

// forecastCmd represents the forecast command
var ForecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "A brief description of your city's forecast",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		controller.PrintWeatherInfos(city)
	},
}

func init() {
	ForecastCmd.Flags().StringVarP(&city, "city", "c", "", "Name of the city to forecast.")

	if err := ForecastCmd.MarkFlagRequired("city"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forecastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forecastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
