package config

import "fmt"

var G int = 999
var g int = 999

func init() {
	fmt.Println("config package initialized")
}
