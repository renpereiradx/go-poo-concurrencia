package main

import (
	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
)

func main() {
	hellomod.SayHello()
	hellomod2.SayHello("V2")
}

// go mod init github.com/username/hello
// go mod tidy
// go mod vendor
// go mod verify
// go mod why
// go mod graph
// go mod edit -fmt
// go mod download	// download modules to local cache	// go mod download -json
// go mod edit		// edit go.mod from tools or scripts
