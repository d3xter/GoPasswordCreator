GoPasswordCreator
=================

This program is used for generating passwords.
The user can choose which character-group will be used to generate the password.
For example: lower case letters, numbers, and so on.
It is also possible for the user to define his/her own set of characters which will be used to create the password.


Compilation
===========

The easiest way to compile the Password-Creator is just to use the new "go" tool and run "go build",
and it will generate a binary called "GoPasswordCreator"


Arguments
=========

At least one of these arguments must be passed to GoPasswordCreator:

- **all**	When this Flag is set, lower/upper-case letters, numbers, special characters and user defined characters are used to generate the password
- **lower**	Lower-Case Letters will be included
- **upper**	Upper-Case Letters will be included
- **numbers**	Numbers will be included
- **special**	Special Letters (like "-") will be included
- **own**	The User can pass a string to GoPasswordCreator and those characters in the string will be included


Options
=======

- **-length** 	Specifies the length of the generated password. Default is set to 8
- **-count**	User can specify how many passwords should be generated at the same time. Default is set to 1
- **-file**	If file is set, the passwords will be written into this file rather than printed out on stdout


Examples
========

	GoPasswordCreator lower upper numbers
This Password could contain lower-case letters, upper-case letters, and numbers.

	GoPasswordCreator lower own="?="
This password could contain lower-case letters and the two characters "?" and "=".

	GoPasswordCreator all lower=f
This password could contain any characters except for lower-case letters.

	GoPasswordCreator -length 8 -count 5 -file /home/d3xter/passwords.txt all
5 Passwords with 8 characters per password will be written into /home/d3xter/passwords.txt
