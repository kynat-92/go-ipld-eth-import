package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mistyexhibit/go-ipld-eth-import/lib"
)

/*

## EVM CODE IPFS

Takes the files dumped from the geth database and imports them to IPFS

## EXAMPLE USAGE

make evmcode-ipfs && ./build/bin/evmcode-ipfs --ipfs-repo-path ~/.ipfs

*/

func main() {
	var (
		evmcodeDir   string
		ipfsRepoPath string
		prefix       string
	)

	// Command line options
	flag.StringVar(&evmcodeDir, "evmcode-directory", "/tmp/evmcode", "Directory where the EVM code files are")
	flag.StringVar(&ipfsRepoPath, "ipfs-repo-path", "~/.ipfs", "IPFS repository path")
	flag.StringVar(&prefix, "prefix", "", "If set, will only process the files which name starts with <prefix>. Only two characters supported")
	flag.Parse()

	// Param check
	if prefix != "" && len(prefix) != 2 {
		fmt.Printf("ERROR: Param '--prefix' only supports two characters. Exiting")
		os.Exit(1)
	}

	// IPFS
	ipfs := lib.InitIPFSNode(ipfsRepoPath)

	// Launch the main loop
	walker := lib.InitWalker(ipfs, evmcodeDir, prefix)
	walker.TraverseDirectory()

	// Print the metrics
	printReport()
}
