package main

import "github.com/yanzay/log"

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Debugf("unable to execute command: %v", err)
	}
}
