package main

import (
	"fmt"
	plugin "github.com/kcmerrill/packs.go"
	"os"
)

func main() {
	plugin.Init("./plugins/", os.Args)
	my_password := "asd123"
	password := plugin.Filter("filter_password", my_password)
	fmt.Println("Before: ", "asd123")
	fmt.Println("Hashed: ", password)

	/* Displaying helper functions */
	fmt.Println("Enabled: ", plugin.IsEnabled("append.py"))
	fmt.Println("Enabled: ", plugin.IsEnabled("doesnotexist.py"))
}
