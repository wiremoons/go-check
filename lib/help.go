package lib

import "fmt"

// PrintHelp function prints out some basic help information for the user
// when requested via the command line flag '-h'
func Help() {
	helptext := `
About Go.
  Go is an open-source programming language supported by Google, that has 
  built-in concurrency capabilities and a robust extensive standard library.

About the Application.
  This application provides a tool to check the current version of the
  Go (golang) langauge developement software and compiler that is freely 
  available from: https://go.dev/

Source Code Information.
  The source code and further information is available here: 
  https://github.com/wiremoons/go-check/

Licence and Copyright.
  The program is provided under the opensource MIT License.
  Copyright © 2023 Simon Rowe [wiremoons].
    `
	// now output the above to screen
	fmt.Println(helptext)
}