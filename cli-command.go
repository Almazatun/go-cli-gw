package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Almazatun/go-cli-gw/util"
)

func HandleGet(getCMD *flag.FlagSet, all *bool, platformName *string, characters *string) {
	getCMD.Parse(os.Args[2:])

	if *all == false && *platformName == "" && *characters == "" {
		fmt.Print("platformName is required or specify -all for all pass list \n")
		getCMD.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		passList := getPass()

		fmt.Printf("Email \t Password \t Username \t Platform \t Description \n")
		for _, passData := range passList {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", passData.Email, passData.Password, passData.Username, passData.Platform, passData.Description)
		}
	}

	if *platformName != "" {
		passList := getPass()

		for _, passData := range passList {
			if *platformName == passData.Platform {
				fmt.Printf("Email \t Username \t Username \t Platform \t Description \n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", passData.Email, passData.Password, passData.Username, passData.Platform, passData.Description)
			}
		}
	}

	if *characters != "" {
		passList := getPass()

		fmt.Printf("Email \t Username \t Username \t Platform \t Description \n")
		for _, passData := range passList {
			if strings.ContainsAny(passData.Platform, *characters) {
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", passData.Email, passData.Password, passData.Username, passData.Platform, passData.Description)
			}
		}
	}
}

func HandleAdd(addCMD *flag.FlagSet, email *string, password *string, username *string, platform *string, desc *string) {

	addCMD.Parse(os.Args[2:])

	util.ValidatePassInputData(addCMD, email, password, username, platform, desc)

	createPassInput := Pass{
		Email:       *email,
		Password:    *password,
		Username:    *username,
		Platform:    *platform,
		Description: *desc,
	}

	passList := getPass()
	passList = append(passList, createPassInput)

	savePass(passList)

}
