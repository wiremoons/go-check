// current.go : supporting library to display the current installed and available
// Go versions.
//
// Copyright 2023 Simon Rowe <simon@wiremoons.com>.
// All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package lib

import (
	"errors"
	"fmt"
	"os"
)

type Current struct {
	localPath        string
	localVersion     string
	currentUrl       string
	availableVersion string
}

func (c Current) localVer() error {
	c.localVersion = os.Getenv("GOVERSION")
	if c.localVersion != "" {
		fmt.Printf("Found: %s\n", c.localVersion)
		return nil
	}
	return errors.New("no local version of Go found")
}

func (c Current) String() (string, error) {
	if c.availableVersion != "" && c.localVersion != "" {
		return fmt.Sprintf("Avaiable: %s\nCurrent: %s\n", c.availableVersion, c.localVersion), nil
	}
	return "", errors.New("current version data not found")
}
