package lib

import "fmt"

func InstallPkg(cmd string, pkg string) bool {
	iCmd := fmt.Sprintf("%s %s", cmd, pkg)
	_, err := ExecuteCmd(iCmd, "-y")
	return err == nil
}

func UbuntuInstallPkg(pkg string) bool {
	return InstallPkg("apt install", pkg)
}

func WindowsInstallPkg(pkg string) bool {
	return InstallPkg("winget install", pkg)
}
