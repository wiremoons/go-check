//	go-check - an application to check current Go toolchain version.
//
// Copyright 2023 Simon Rowe <simon@wiremoons.com>.
// All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.
package main

import (
	"flag"
	gc "github.com/wiremoons/go-check/lib"
	"os"
	"path/filepath"
)

// GLOBAL VARIABLES

// set the version of the app here
var appversion = "0.1.0"
var appname string

// below used by flag to store command line arguments given by the user
var helpMe bool
var version bool

// init function always runs before main() so used here to
// set-up the required command line flag variables
func init() {
	// IntVar; StringVar; BoolVar options for flag
	// format required: variable, cmd line flag, initial value, description.
	flag.BoolVar(&helpMe, "h", false, "\tdisplay help for this program.")
	flag.BoolVar(&version, "v", false, "\tdisplay the applications version.")
	appname = filepath.Base(os.Args[0])
}

func main() {
	// get the command line args passed to the program
	flag.Parse()

	// was the command line flag '-h' used?
	if helpMe {
		// call function to display information about the application
		gc.Help()
		// call to display the standard command line usage info
		flag.Usage()
		os.Exit(0)
	}

	// was the command line flag '-v' used?
	if version {
		gc.Version(appname, appversion)
		// exit the application
		os.Exit(0)
	}

}