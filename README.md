This is an example of the problem with strict Minimum Version Selection

This repo contains two modules, uses_old and uses_new

uses_old has a package which is a library that has one function which returns a static string. However that module also depends on runc version 1.0.0-93 which it just uses nominally, to print the result of a default struct called Device.

uses_new has a main package which includes uses_old and prints the result of its static function. However uses_old also includes runc version 1.0.2 which it just uses nominally.

Runc 1.0.2 has breaking changes, it deletes the struct called device.

uses_new cannot compile
```
$ go build .
# github.com/benbuzbee/module_example/uses_old
../../go/pkg/mod/github.com/benbuzbee/module_example/uses_old@v0.0.0-20220119185220-96d52e523360/main.go:9:14: undefined: configs.Device
```


This is because go belows 1.0.2 MUST be compatible with 1.0.0-93 and forces uses_old to compile with the new version.
However it is very often the case that minor version changes include breaking changes, even if we wish it were not.
