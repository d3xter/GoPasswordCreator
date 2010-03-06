package password_creator

import (
	"strings"
	"os"
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

func GeneratePassword(length int) (password string, err os.Error) {
	passwordBuffer := new(bytes.Buffer)

	//For now, we use the actual time to set the seed, otherwise the password would be the same all the time
	rand.Seed(time.Nanoseconds())

	if !EnoughCharacters() {
		return "", os.NewError(NOTENOUGHCHARACTERS)
	}

	for i := 0; i < length; i++ {
		char := CHARACTERS[rand.Intn(len(CHARACTERS))]

		//Append the character at the end of the password
		passwordBuffer.WriteString(string(char))
	}

	return passwordBuffer.String(), nil
}

