package module

import (
	"fmt"
	lconfigs "github.com/opencontainers/runc/libcontainer/configs"
) 

func init() {
	fmt.Println(lconfigs.Device{})
}
func HelloWorld() string {
	return "Hello World - I am a module that uses runc 1.0.0-rc93"
}