/*Copyright (C) 2010 Andreas Sinz

This file is part of GoPasswordCreator.

GoPasswordCreator is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; only version 2 of the License.
This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; 
without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
You should have received a copy of the GNU General Public License along with this program; if not, see <http://www.gnu.org/licenses/>.
*/


package password_creator

import (
	"strings"
	"rand"
	"time"
	"bytes"
)

var (
	LETTERS = "abcdefghijklmnopqrstuvwxyz"
	NUMBERS = "0123456789"
	SPECIAL = ",.-"

	//Characters will store the selected characters, which will be used to generate the password
	CHARACTERS string

	//Error Message
	NOTENOUGHCHARACTERS = "Not enough Characters specified to generate a password"
)

func CreateCharacterArray(lowerCase, upperCase, numbers, specialCharacters bool) {

	if lowerCase {
		CHARACTERS += LETTERS
	}

	if upperCase {
		CHARACTERS += strings.ToUpper(LETTERS)
	}

	if numbers {
		CHARACTERS += NUMBERS
	}

	if specialCharacters {
		CHARACTERS += SPECIAL
	}
}

func AddUserDefinedCharacters(userCharacters string) {
	CHARACTERS += userCharacters
}

//Returns a bool, which describes if there are enough characters to generate a password
func EnoughCharacters() bool { return len(CHARACTERS) > 1 }

func GeneratePassword(length, count int, quit chan<- bool, output chan<- string) {
	passwordBuffer := new(bytes.Buffer)

	//For now, we use the actual time to set the seed, otherwise the password would be the same all the time
	rand.Seed(time.Nanoseconds())

	if !EnoughCharacters() {
		output <- NOTENOUGHCHARACTERS
	}

	for i := 0; i < count; i++ {
		passwordBuffer.Reset()
		
		for j := 0; j < length; j++ {
			char := CHARACTERS[rand.Intn(len(CHARACTERS))]

			//Append the character at the end of the password
			passwordBuffer.WriteString(string(char))
		}

		output <- passwordBuffer.String()
	}

	quit <- true
}

