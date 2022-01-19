package main

import (
	"fmt"
	lconfigs "github.com/opencontainers/runc/libcontainer/configs"
	old "github.com/benbuzbee/module_example/uses_old"
) 

func main() {
	fmt.Println(lconfigs.Arg{}, old.HelloWorld())
}
func HelloWorld() string {
	return "Hello World - I am a module that uses runc 1.0.2"
}