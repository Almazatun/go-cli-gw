package util

import (
	"flag"
	"fmt"
	"os"
)

func ValidatePassInputData(addCMD *flag.FlagSet, email *string, password *string, username *string, platform *string, desc *string) {
	if *email == "" || *password == "" || *username == "" || *platform == "" || *desc == "" {
		fmt.Print("All fields are required for adding a pass list \n")
		addCMD.PrintDefaults()
		os.Exit(1)
	}
}
