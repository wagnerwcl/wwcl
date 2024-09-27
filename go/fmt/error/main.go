package main

import "fmt"

func main() {
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())
}

/*
func main() { ... }: This defines the main function, which is the entry point of every Go program. The code within the curly braces will be executed when the program runs.

const name, id = "bueller", 17: This line declares two constants:

name: A string constant with the value "bueller".
id: An integer constant with the value 17.
const ensures that these values cannot be changed later in the program.
err := fmt.Errorf("user %q (id %d) not found", name, id):

fmt.Errorf: This function from the fmt package creates a new error value. It takes a format string and a list of arguments.
"user %q (id %d) not found": This is the format string.
%q is a placeholder for a quoted string (the name variable).
%d is a placeholder for a decimal integer (the id variable).
name, id: These are the values that will be substituted into the placeholders in the format string.
The result is an error value containing the formatted string: "user "bueller" (id 17) not found".
This error value is assigned to the variable err using the short variable declaration operator :=.
fmt.Println(err.Error()):

err.Error(): This method on the error value returns the error message as a string.
fmt.Println: This function prints the error message to the console, followed by a newline character.
*/
