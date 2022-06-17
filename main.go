package main

import (
	"fmt"

	"github.com/axelisvoid/pc-setup/lib"
)

func main() {

	sys := lib.GetSystem()

	lib.GetSystemPackages("settings.json", &sys)
	fmt.Println(sys)
}
