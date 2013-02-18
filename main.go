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
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	passwordLength = flag.Int("length", 8, "Length of the generated Password")

	// Variables that define what characters to use in the password
	lowerCase         bool
	upperCase         bool
	numerals          bool
	specialCharacters bool
	usersCharacters   string

	passwordCount = flag.Int("count", 1, "Determine how many passwords to create")

	file = flag.String("file", "", "Write passwords to the named file instead of standard output")
)

func usage() {
	command := os.Args[0]
	fmt.Fprintf(os.Stderr,
		`Usage: %s [all] [alphanum] [lower] [upper] [numbers] [special] [own=CHARACTERS]
%s requires at least one of the following subcommands to specify what characters
may be used in the password:
  all: Equivalent to 'alphanum special'
  alphanum: Equivalent to 'lower upper numbers'
  lower: Use lower-case letters
  upper: Use upper-case letters
  numbers: Use digits
  special: Use special characters
  own: Specifies a custom set of characters to use
'all', 'alphanum', 'lower', 'upper', 'numbers', and 'special' may be followed by
'=f' to nullify that character set.
Options:
`,
		command, command)
	flag.PrintDefaults()
}

// Sets a list of bool variables to the same value. If len(args) == 1, use true.
// Otherwise, parse args[1] as a bool and use that.
func setBool(args []string, vars ...*bool) {
	on := true
	if len(args) > 1 {
		var err error
		on, err = strconv.ParseBool(args[1])
		if err != nil {
			printError(err)
		}
	}
	for _, bp := range vars {
		*bp = on
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()

	for _, arg := range flag.Args() {

		// Separate the subcommand from the value
		parsed := strings.SplitN(arg, "=", 2)

		// Group arguments by the data type of their value
		switch parsed[0] {
		case "own":
			// Need a string value
			if len(parsed) == 2 {
				usersCharacters = parsed[1]
			} else {
				printError(fmt.Errorf("'own' requires a '=' to specify characters"))
			}

			// All other arguments take boolean values
		case "all":
			setBool(parsed, &lowerCase, &upperCase, &numerals, &specialCharacters)
		case "alphanum":
			setBool(parsed, &lowerCase, &upperCase, &numerals)
		case "lower":
			setBool(parsed, &lowerCase)
		case "upper":
			setBool(parsed, &upperCase)
		case "numbers":
			setBool(parsed, &numerals)
		case "special":
			setBool(parsed, &specialCharacters)

		default:
			printError(fmt.Errorf("Invalid argument: %s", parsed[0]))
		}
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

	creator, err := NewCreator(output, lowerCase, upperCase, numerals, specialCharacters, usersCharacters)
	defer output.Close()

	if err != nil {
		printError(err)
	} else {
		writeErr := creator.WritePasswords(*passwordLength, *passwordCount)

		if writeErr != nil {
			printError(writeErr)
		}
	}
}

func printError(err error) {
	fmt.Fprintln(os.Stderr, "Error: "+err.Error())
}
