package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Almazatun/go-cli-gw/util"
)

func main() {
	getCMD := flag.NewFlagSet("get", flag.ExitOnError)
	addCMD := flag.NewFlagSet("add", flag.ExitOnError)

	// get cmd
	getAll := getCMD.Bool("all", false, "Get all pass that you saved")
	getByPlatform := getCMD.String("plm", "", "Platform name where you register")

	// add cmd
	addEmail := addCMD.String("email", "", "Your email in platform")
	addPassword := addCMD.String("password", "", "Your password in platform")
	addUsername := addCMD.String("username", "", "Your username in platform")
	addPlatform := addCMD.String("platform", "", "Platform name where you sing up")
	addDesc := addCMD.String("des", "", "Some description")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCMD, getAll, getByPlatform)
	case "add":
		HandleAdd(addCMD, addEmail, addPassword, addUsername, addPlatform, addDesc)
	default:
		fmt.Println("END")
	}
}

func HandleGet(getCMD *flag.FlagSet, all *bool, platformName *string) {
	getCMD.Parse(os.Args[2:])

	if *all == false && *platformName == "" {
		fmt.Print("platformName is required or specify -all for all pass list")
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
}

func HandleAdd(addCMD *flag.FlagSet, email *string, password *string, username *string, platform *string, desc *string) {
	fmt.Println(&email, &password, &username, &platform, &desc)

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
