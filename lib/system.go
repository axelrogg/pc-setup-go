package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
)

type System struct {
	name     string
	distro   string
	packages []string
}

type Apps struct {
	Apps []App `json:"apps"`
}

type App struct {
	Os       string   `json:"os"`
	Packages []string `json:"packages"`
}

var supportedLinuxDistros = []string{
	"ubuntu",
}

func GetLinuxDistro() string {
	cmd := "lsb_release"
	arg := "-a"
	execOut, err := ExecuteCmd(cmd, arg)

	if err != nil {
		errStr := fmt.Sprintf("Execution of '%s' failed. Error: %s", arg, err)
		log.Fatal(errStr)
	}

	// Note the hard-coded `1` array access value in `distId`.
	// This is because the `lsb_release -a` command output is consistent. For example,
	// the following is the output of the same command in an Ubuntu 22.04 machine:

	// ```No LSB modules are available.
	//    Distributor ID:	Ubuntu            <------ This is the line that we want.
	//    Description:	Ubuntu 22.04 LTS
	//    Release:	22.04
	//    Codename:	jammy```

	distId := strings.Split(execOut, "\n")[1]

	for _, distro := range supportedLinuxDistros {
		if strings.Contains(strings.ToLower(distId), distro) {
			return distro
		}
	}
	return ""
}

func GetSystem() System {

	system := System{}
	os := runtime.GOOS

	switch os {

	case "windows":
		system.name = os
		system.distro = ""

	case "linux":
		system.name = os
		system.distro = GetLinuxDistro()

	default:
		log.Fatal("Unsupported Operating System error")
	}
	return system
}

func GetSystemPackages(fileName string, system *System) {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		errStr := fmt.Sprintf("os.Open(%s) resulted in error %s", fileName, err)
		log.Fatal(errStr)
	}
	defer jsonFile.Close()

	// Reads `jsonFile` into a byteArray
	byteVal, _ := ioutil.ReadAll(jsonFile)

	var apps Apps

	// Unmarshal `byteVal` into `apps`
	json.Unmarshal(byteVal, &apps)

	var currentOs string
	currentOs = system.distro
	if system.name == "windows" {
		currentOs = system.name
	}

	// Copy packages from `apps.json` file to the `System` struct.
	for i := 0; i < len(apps.Apps); i++ {
		// we only care about the packages for the current OS
		if apps.Apps[i].Os != currentOs {
			continue
		}
		system.packages = apps.Apps[i].Packages
	}
}
