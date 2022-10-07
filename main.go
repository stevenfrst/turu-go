/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"gitag-request/cmd"
	"gitag-request/config"
	"gitag-request/repository"
)

func main() {
	config.SetupConfig()
	//fmt.Println(viper.GetString("host.server"))
	repository.SetupServiceData()
	//repository.ReadEnv()
	cmd.Execute()
}
