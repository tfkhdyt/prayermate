/*
Copyright © 2023 Taufik Hidayat <tfkhdyt@proton.me>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// locationCmd represents the location command
var locationCmd = &cobra.Command{
	Use:   "location",
	Short: "Manage your location",
	Long:  "Command to manage your location",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("location called")
	// },
}

func init() {
	rootCmd.AddCommand(locationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// locationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// locationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
