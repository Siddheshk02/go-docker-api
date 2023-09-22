/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// multiplyCmd represents the multiply command
var multiplyCmd = &cobra.Command{
	Use:   "multiply",
	Short: "Multiplication of two integers",
	Run: func(cmd *cobra.Command, args []string) {
		result := num1 * num2
		fmt.Printf("Result of %d * %d = %d\n", num1, num2, result)
	},
}

func init() {
	rootCmd.AddCommand(multiplyCmd)

	multiplyCmd.Flags().IntVar(&num1, "num1", 0, "First integer")
	multiplyCmd.Flags().IntVar(&num2, "num2", 0, "Second integer")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// multiplyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// multiplyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
