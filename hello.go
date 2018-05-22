package main

/*
bin/
    hello                 # command executable
pkg/
    linux_amd64/          # this will reflect your OS and architecture
        github.com/user/
            stringutil.a  # package object
src/
    github.com/user/
        hello/
            hello.go      # command source
        stringutil/
            reverse.go    # package source
 */

import (
	"fmt"
	"github.com/samkeen/go-stringUtils"
)


func main() {
	fmt.Println(stringutil.Reverse("Hello, world."))
}