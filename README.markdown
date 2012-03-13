GoPasswordCreator
=================

This program is used for generating passwords.
The user can choose, which character-group will be used to generate the password.
For example: lower case letters, numbers and so on.
It is also possible for the user to define his/her own set of characters, which will be used to create the password.


Compilation
===========

The easiest way to compile the Password-Creator is, just to use the new "go" tool and run "go build" 
and it will generate a binary called "GoPasswordCreator"


Arguments
=========

There are several arguments, which can be passed to the Password-Creator:

--all		When this Flag is set, lower/upper-case letters, numbers, special characters and by user defined characters are used to generate the password
--lower		Lower-Case Letters will be included
--upper		Upper-Case Letters will be included
--numbers   	Numbers will be included
--special	Special Letters (like "-") will be included
--own		The User can pass a string to the Password-Creator and those characters in the string will be included
--length 	Specifies the length of the generated password. Default is set to 8
--count		User can specify with --count, how many passwords should be generated at the same. Default is set to 1
--file		If file is set, the passwords will be written into this file, rather than printed out on stdout


Examples
========

	password_creator --lower --upper --numbers
This Password could contain lower-case letters, upper-case letters and numbers

	password_creator --lower --own "?="
This password could contain lower-case letters and those two characters ("?" and "="), which the user has passed to the Password-Creator

	password_creator --lower=false --upper=true
This password could contain upper-case letters. Lower-case letters will not be used to generate the password

	password_creator --all --length 8 --count 5 --file /home/d3xter/passwords.txt
5 Passwords with 8 characters per password will be written into /home/d3xter/passwords.txt
