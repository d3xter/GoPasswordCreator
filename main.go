/*Copyright (C) 2010 Andreas Sinz

This file is part of GoPasswordCreator.

GoPasswordCreator is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; only version 2 of the License.
This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
You should have received a copy of the GNU General Public License along with this program; if not, see <http://www.gnu.org/licenses/>.
*/


package main

import (
	"fmt"
	"flag"
	"os"
)

var (
	passwordLength = flag.Int("length", 8, "Length of the generated Password")

	//Define all available flags which are used to specify which characters to use to generate the password
	lowerCase         = flag.Bool("lower", false, "Should LowerCase characters be included?")
	upperCase         = flag.Bool("upper", false, "Should UpperCase characters be included?")
	numerals          = flag.Bool("numbers", false, "Should the Numbers be included?")
	specialCharacters = flag.Bool("special", false, "Should special characters be included?")
	usersCharacters   = flag.String("own", "", "Characters defined by the user which will be also be used to generate the password")

	//If --all Flag is set, then lower/upper-case letters, numbers, special characters, and user defined characters are used
	allGroups = flag.Bool("all", false, "Use lower/upper-case letters, numbers, special characters, and user defined characters to generate the password")

	//The user can determine how many passwords will be created
	passwordCount = flag.Int("count", 1, "Determine how many passwords will be created")

	//The user can define a file where the passwords will be written into.
	//If file is omitted, then it will print the passwords on Stdout.
	file = flag.String("file", "", "The File where the passwords should be written into")
)


func main() {
	flag.Parse()

	if *allGroups {
		*lowerCase = true
		*upperCase = true
		*numerals = true
		*specialCharacters = true
	}

	var output *os.File
	var fileErr error

	if *file != "" {
		if output, fileErr = os.Create(*file); fileErr != nil {
			printError(fileErr)
			output = os.Stdout
		}
	} else {
		output = os.Stdout
	}

	creator, err := NewCreator(output, *lowerCase, *upperCase, *numerals, *specialCharacters, *usersCharacters)
	defer output.Close()

	if err != nil {
		printError(err)
	} else {
		writeErr := creator.WritePasswords(*passwordLength, *passwordCount)

		if writeErr != nil  {
			printError(writeErr)
		}
	}
}

func printError(err error) {
	fmt.Println("Error: " + err.Error())
}
