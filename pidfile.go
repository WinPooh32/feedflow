package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//https://gist.github.com/davidnewhall/3627895a9fc8fa0affbd747183abca39
// Write a pid file, but first make sure it doesn't exist with a running pid.
func writePidFile(pidFile string) error {
	return ioutil.WriteFile(pidFile, []byte(fmt.Sprintf("%d", os.Getpid())), 0664)
}
