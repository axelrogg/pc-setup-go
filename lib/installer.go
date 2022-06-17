package lib

import (
	"fmt"
	"strings"
)

func InstallPkg(cmd string, pkg string) bool {
	icmd := fmt.Sprintf("%s %s", cmd, pkg)
	_, err := ExecuteCmd(icmd)
	return err == nil
}

func UbuntuInstallPkg(pkgs []string) bool {

	return InstallPkg("apt install", strings.Join(pkgs, " "))
}

func WindowsInstallPkg(pkg string) bool {
	return InstallPkg("winget install", pkg)
}
