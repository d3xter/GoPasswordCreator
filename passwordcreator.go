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
	"bytes"
	"crypto/rand"
	"errors"
	"os"
	"strings"
)

type Creator struct {
	characters string
	file       *os.File
}

const (
	letters = "abcdefghijklmnopqrstuvwxyz"
	numbers = "0123456789"
	special = ",.-"
)

func NewCreator(file *os.File, lowerCase, upperCase, numerals, specialCharacters bool, userCharacters string) (creator *Creator, err error) {
	if file == nil {
		return nil, errors.New("File is nil!")
	}

	characters := ""

	if lowerCase {
		characters += letters
	}

	if upperCase {
		characters += strings.ToUpper(letters)
	}

	if numerals {
		characters += numbers
	}

	if specialCharacters {
		characters += special
	}

	characters += userCharacters

	i := 0
	seen := map[byte]struct{}{}
	for i < len(characters) {
		char := characters[i]
		if _, ok := seen[char]; ok {
			characters = characters[:i] + characters[i+1:]
		} else {
			i++
			seen[char] = struct{}{}
		}
	}

	if len(characters) <= 1 {
		err = errors.New("Not enough Characters specified to generate passwords")
		return nil, err
	}

	return &Creator{characters, file}, err
}

func (creator *Creator) CreatePassword(length int) (password []byte, err error) {
	passwordBuffer := new(bytes.Buffer)

	//Create a byte-array and fill it with random bytes
	randbytes := make([]byte, length)
	if _, err := rand.Read(randbytes); err == nil {

		for j := 0; j < length; j++ {
			tmpindex := int(randbytes[j]) % len(creator.characters)
			char := creator.characters[tmpindex]

			//Append the character at the end of the password
			passwordBuffer.WriteString(string(char))
		}

		return passwordBuffer.Bytes(), nil
	}

	return nil, err
}

func (creator *Creator) WritePasswords(length, count int) error {
	for i := 0; i < count; i++ {
		if pass, err := creator.CreatePassword(length); err == nil {
			creator.file.Write(append(pass, byte('\n')))
		} else {
			return err
		}
	}

	return nil
}
