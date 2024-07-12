/*
 *	The name of the identifier must begin with a letter or an underscore(_).
 *	And the names may contain the letters ‘a-z’ or ’A-Z’ or digits 0-9 as well as the character ‘_’.
 *	The name of the identifier should not start with a digit.
 *	The name of the identifier is case-sensitive.
 *	Keywords are not allowed to be used as an identifier name.
 *	There is no limit on the length of the name of the identifier, but it is advisable to use an optimum length of 4 – 15 letters only.
 */
package main

import (
	"fmt"

	"github.com/AbdullahReda1/Golang_GeeksforGeeks/Fundamentals/packages/exportedpackage"
)

func main() {
	/* Valid identifiers */
	var _example23 = 0
	var example = 1
	var example23sd = 2
	var Example = 3
	var exampLe = 4
	var exa_mple = 5

	//var 212ID				/* Invalid identifier */
	//var if				/* Invalid identifier */
	//var default			/* Invalid identifier */

	fmt.Println(_example23, example, example23sd, Example, exampLe, exa_mple)
	fmt.Println(exportedpackage.ExportedVariable)
}
