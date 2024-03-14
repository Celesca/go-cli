/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

// findCircleAreaCmd represents the findCircleArea command
var findCircleAreaCmd = &cobra.Command{
	Use:   "findCircleArea",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		res := findCircleArea(args)
		fmt.Println("Result Circle area is: ", res)
	},
}

func init() {
	rootCmd.AddCommand(findCircleAreaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCircleAreaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCircleAreaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func findCircleArea(args []string) float64 {
	a, _ := strconv.ParseFloat(args[0], 64)
	b := math.Pow(a, 2)

	return math.Pi * b
}
