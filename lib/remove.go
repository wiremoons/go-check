// remove.go : supporting library to remove the current installation of Go
//
// Copyright 2023 Simon Rowe <simon@wiremoons.com>.
// All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package lib

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Installed struct {
	installedVersion  string
	installedPath     string
	installPathExists bool
}

// CheckInstalled function obtains the data for any installed Go version and its install PATH (GOROOT)
func (i *Installed) RemoveInstalled() {
	if !i.localInstall() {
		fmt.Printf("\nABORT: failed to installed Go SDK:")
		fmt.Printf("Located Go version: '%s'\n", i.installedVersion)
		fmt.Printf("Installation path:  '%s'\n", i.installedPath)
		os.Exit(1)
	}
	fmt.Printf("\nLocated Go version: '%s'\n", i.installedVersion)
	fmt.Printf("Installation path:  '%s'\n\n", i.installedPath)
	i.pathExist()
	fmt.Printf("[i]  installation path exists: '%t'\n", i.installPathExists)
	if i.installPathExists {
		fmt.Printf("[*]  removing all files and directories in the path...\n")
		err := os.RemoveAll(i.installedPath)
		if err != nil {
			fmt.Printf("[!]  ERROR: %v\n", err)
			return
		}
		fmt.Printf("[✔]  removal completed successfully.\n")
		return
	}
}

// localInstall obtains the `go env GOVERSION` and `go env GOROOT` data
func (i *Installed) localInstall() bool {
	verOut, err := exec.Command("go", "env", "GOVERSION").Output()
	if err != nil {
		i.installedVersion = "not found"
		i.installedPath = "not found"
		return false
	}
	i.installedVersion = strings.TrimSpace(string(verOut))

	pathOut, err := exec.Command("go", "env", "GOROOT").Output()
	if err != nil {
		i.installedPath = "not found"
		return false
	}
	i.installedPath = strings.TrimSpace(string(pathOut))
	return true
}

// check if 'i.installedPath' exists?
func (i *Installed) pathExist() {
	_, err := os.Stat(i.installedPath)
	if err == nil {
		i.installPathExists = true
		return
	}
	if os.IsNotExist(err) {
		i.installPathExists = false
		return
	}
	i.installPathExists = false
	return
}
