package main

import (
	"os"

	"github.com/biograph-health/interview-platform-external-asrtdd/cmd"
)

func main() {
	cmd.Execute(os.Args[1:])
}
