package main

import (
	"fmt"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	// Where your local node is running on localhost:5001
	sh := shell.NewShell("http://middleware-8de8.settlemint.com/bpaas-19faDE45dFE2ea995cCa257c764c669C580b2980")
	cid, err := sh.Add(strings.NewReader("RTA-Logo2.png"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s", cid)
}
