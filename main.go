package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// get cmd
	getCMD := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCMD.Bool("all", false, "Get all pass that you saved")
	getByPlatform := getCMD.String("plm", "", "Platform name where you register")

	// add cmd
	addCMD := flag.NewFlagSet("add", flag.ExitOnError)
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
		addCMD.Parse(os.Args[2:])
		HandleAdd(addCMD, addEmail, addPassword, addUsername, addPlatform, addDesc)
	default:
		fmt.Println("<>|<>")
	}
}
