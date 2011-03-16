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
)

var (
	passwordLength = flag.Int("length", 8, "Length of the generated Password")

	//Define all available flags, which are used to specify, which characters to use to generate the password
	lowerCase         = flag.Bool("lower", false, "Should LowerCase characters be included?")
	upperCase         = flag.Bool("upper", false, "Should UpperCase characters be included?")
	numerals          = flag.Bool("numbers", false, "Should the Numbers be included?")
	specialCharacters = flag.Bool("special", false, "Should special characters be included?")
	usersCharacters   = flag.String("own", "", "Characters defined by the user, which will be also be used to generate the password")

	//If --all Flag is set, then lower/upper-case letters, numbers and special characters and by user defined characters are used
	allGroups = flag.Bool("all", false, "Use lower/upper-case letters, numbers, special characters and by user defined characters to generate the password")

	//The user can determine, how many passwords will be created
	passwordCount = flag.Int("count", 1, "Determine, how many passwords will be created")
)


func main() {
	flag.Parse()

	if *allGroups {
		*lowerCase = true
		*upperCase = true
		*numerals = true
		*specialCharacters = true
	}

	creator, err := NewCreator(*lowerCase, *upperCase, *numerals, *specialCharacters, *usersCharacters)

	if err == nil {
		fmt.Println("Your password(s):")

		for i := 0; i < *passwordCount; i++ {
			if password, createErr := creator.CreatePassword(*passwordLength); createErr == nil {
				fmt.Println(password)
			} else {
				fmt.Println(createErr)
			}
		}
	} else {
		fmt.Println(err)
	}
}
