package password_creator

import (
	"strings"
	"os"
	"rand"
	"time"
	"buffer"
)

var (
	LETTERS = "abcdefghijklmnopqrstuvwxyz"
	NUMBERS = "0123456789"
	SPECIAL = ",.-"

	//Characters will store the selected characters, which will be used to generate the password
	CHARACTERS string
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

func CreatePassword(length int) (password string, err os.Error) {
	passwordBuffer := new(bytes.Buffer)

	//For now, we use the actual time to set the seed, otherwise the password would be the same all the time
	rand.Seed(time.Seconds())

	if len(CHARACTERS) <= 1 {
		return "", os.NewError("Not enough Characters specified to generate a password")
	}

	for i := 0; i < length; i++ {
		char := CHARACTERS[rand.Intn(len(CHARACTERS))]

		//Append the character at the end of the password
		passwordBuffer.WriteString(string(char))
	}

	return passwordBuffer.String(), nil
}

