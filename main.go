// This is a placeholder file that demonstrates how to use the fsutils package.
package main

import (
	"fmt"
	
	// Import the package with a custom import name
	"github.com/utsav-56/go_fs_utils/fsutils"
)

func main() {
	fmt.Println("FSUtils - Go Filesystem Utilities")
	fmt.Println("Please check the 'examples' directory for usage examples")
	fmt.Println("Import the package in your project:")
	fmt.Println("  import \"github.com/utsav-56/go_fs_utils/fsutils\"")
	
	// Show a simple example
	fmt.Println("\nExample: Checking if a directory exists")
	if fsutils.DirExists(".") {
		fmt.Println("Current directory exists")
	}
}
