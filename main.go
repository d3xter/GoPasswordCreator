package main

import (
	"fmt"
	"./password_creator"
	"flag"
)

var (
	passwordLength = flag.Int("length", 8, "Length of the generated Password")

	//Define all available flags, which are used to specify, which characters to use to generate the password
	lowerCase         = flag.Bool("lower", false, "Should LowerCase characters be included?")
	upperCase         = flag.Bool("upper", false, "Should UpperCase characters be included?")
	numbers           = flag.Bool("numbers", false, "Should the Numbers be included?")
	specialCharacters = flag.Bool("special", false, "Should special characters be included?")
	usersCharacters   = flag.String("own", "", "Characters defined by the user, which will be also be used to generate the password")

	//If --all Flag is set, then lower/upper-case letters, numbers and special characters and by user defined characters are used
	allGroups = flag.Bool("all", false, "Use lower/upper-case letters, numbers, special characters and by user defined characters to generate the password")

	//The user can determine, how many passwords will be created
	passwordsCount = flag.Int("count", 1, "Determine, how many passwords will be created")
)


func main() {
	flag.Parse()

	if *allGroups {
		password_creator.CreateCharacterArray(true, true, true, true)
	} else {
		password_creator.CreateCharacterArray(*lowerCase, *upperCase, *numbers, *specialCharacters)
	}

	password_creator.AddUserDefinedCharacters(*usersCharacters)

	if password_creator.EnoughCharacters() {
		fmt.Println("Your password(s):")

		for i := 0; i < *passwordsCount; i++ {
			password, error := password_creator.GeneratePassword(*passwordLength)

			if error == nil {
				//Everything went well
				fmt.Println(password)
			} else {
				fmt.Println("Error: ", error.String())
			}
		}
	} else {
		fmt.Println(password_creator.NOTENOUGHCHARACTERS)
	}
}

