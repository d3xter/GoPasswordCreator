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

At least one of these arguments must be passed to GoPasswordCreator to specify what characters may be used in the password:

- **all**	Equivalent to 'alphanum special'
- **alphanum**  Equivalent to 'lower upper numbers'
- **lower**	Use lower-case letters
- **upper**	Use upper-case letters
- **numbers**	Use digits
- **special**	Use special characters (like '-')
- **own**	Takes a custom string that contains characters to use


Options
=======

- **-length** 	Specifies the length of the generated password. Default is 8.
- **-count**	Specifies how many passwords to generate. Default is 1.
- **-file**	Write passwords to the named file instead of standard output.


Examples
========

	GoPasswordCreator lower upper numbers
This generates passwords that could contain lower-case letters, upper-case letters, and numbers.

	GoPasswordCreator alphanum
This generates passwords that could contain lower-case letters, upper-case letters, and numbers.  This is just a short hand for the previous command.

	GoPasswordCreator lower own="?="
This generates passwords that could contain lower-case letters and the two characters "?" and "=".

	GoPasswordCreator all lower=f
This generates passwords that could contain any characters except for lower-case letters.

	GoPasswordCreator -length 8 -count 5 -file /home/d3xter/passwords.txt all
5 Passwords with 8 characters per password will be written into /home/d3xter/passwords.txt
