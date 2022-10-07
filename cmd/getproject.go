/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gitag-request/delivery/gitlab"
	"gitag-request/repository"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// getprojectCmd represents the getproject command
var getprojectCmd = &cobra.Command{
	Use:   "getproject",
	Short: "List All Project UUID",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		promptOptions()
	},
}

func promptOptions() {
	prompt := promptui.Select{
		Label: "Service List",
		Items: repository.SERVICE_NAME,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	gitlab.GetAPI(repository.SERVICE_DATA[result],repository.SERVICE_TAGS[result])
}

func init() {
	rootCmd.AddCommand(getprojectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getprojectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getprojectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type promptContent struct {
	errorMsg string
	label    string
}
