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
	// localPath        string
	localVersion string
	// currentUrl       string
	availableVersion string
	executingVersion string
}

// Populate function obtains the data required to
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
// NB: in Aug 2023 the output for the above URL changes to add a second and third line:
//
//	1  go1.21.0
//	2  time 2023-08-04T20:14:06Z
//	3
func (c *Current) availVersion() {
	url := "https://go.dev/VERSION?m=text"

	// configure the web request
	var webClient = &http.Client{Timeout: 10 * time.Second}
	resp, err := webClient.Get(url)
	// exit app if web request errors
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "\nWARNING HTTP ERROR: %v\n", err)
		runtime.Goexit()
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "\nUnable to extract webpage content: %v\n", err)
			c.availableVersion = "UNKNOWN"
			return
		}
		fullVersion := string(bodyBytes)
		if len(fullVersion) > 0 {
			indexOfNewline := strings.Index(fullVersion, "\n")
			if indexOfNewline > 0 {
				c.availableVersion = fullVersion[:indexOfNewline]
				return
			}
		}
	}
	_, _ = fmt.Fprintf(os.Stderr, "failed to obtain available Go version from: https://go.dev/VERSION?m=text")
	c.availableVersion = "UNKNOWN"
}

func (c *Current) VersionString() string {
	return fmt.Sprintf("\nGo Language Versions\n\nAvailable: '%s'\nInstalled: '%s'\nExecuting: '%s'\n", c.availableVersion, c.localVersion, c.executingVersion)
}
