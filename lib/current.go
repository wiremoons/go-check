// current.go : supporting library to display the current installed and available
// Go versions.
//
// Copyright 2023 Simon Rowe <simon@wiremoons.com>.
// All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package lib

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Current struct {
	localPath        string
	localVersion     string
	currentUrl       string
	availableVersion string
	executingVersion string
}

func (c *Current) Populate() {
	c.execVer()
	c.localVer()
	c.availVersion()
}

func (c *Current) execVer() {
	c.executingVersion = runtime.Version()
	if c.executingVersion == "" {
		fmt.Fprintf(os.Stderr, "no executing Go version found")
	}
}

// go env GOVERSION
func (c *Current) localVer() {
	out, err := exec.Command("go", "env", "GOVERSION").Output()
	if err != nil {
		// fmt.Fprintf(os.Stderr, "failed to execute command 'go env GOVERSION': %v", err)
		c.localVersion = "not found"
		return
	}
	c.localVersion = strings.TrimSpace(string(out))
}

// https://go.dev/VERSION?m=text
func (c *Current) availVersion() {
	url := "https://go.dev/VERSION?m=text"

	// configure the web request
	var webClient = &http.Client{Timeout: 10 * time.Second}
	resp, err := webClient.Get(url)
	// exit app if web request errors
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nWARNING HTTP ERROR: %v\n", err)
		runtime.Goexit()
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nUnable to extract webpage content: %v\n", err)
			c.availableVersion = "UNKNOWN"
			return
		}
		c.availableVersion = string(bodyBytes)
		return
	}
	fmt.Fprintf(os.Stderr, "failed to obtain available Go version from: https://go.dev/VERSION?m=text")
	c.availableVersion = "UNKNOWN"
}

func (c Current) VersionString() string {
	return fmt.Sprintf("\nGo Language Versions\n\nAvailable: '%s'\nInstalled: '%s'\nExecuting: '%s'\n", c.availableVersion, c.localVersion, c.executingVersion)
}
