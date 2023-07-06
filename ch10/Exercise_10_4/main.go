//“Exercise 10.4:
//Construct a tool that reports the set of all packages in the workspace
//that transitively depend on the packages specified by the arguments.
//
//Hint: you will need to run go list twice, once for the initial
//packages and once for all packages.
//
//You may want to parse its JSON output using
//the encoding/json package.”

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Package represents a Go package
type Package struct {
	ImportPath string   `json:"ImportPath"`
	Deps       []string `json:"Deps"`
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("No packages specified.")
		os.Exit(1)
	}

	initialPackages := strings.Join(args, " ")

	// Run go list for the initial packages
	initialOutput, err := runGoList(initialPackages)
	if err != nil {
		log.Fatal("Error running go list:", err)
	}

	var initialDeps []string
	err = json.Unmarshal(initialOutput, &initialDeps)
	if err != nil {
		log.Fatal("Error parsing go list output:", err)
	}

	// Run go list for all packages in the workspace
	allOutput, err := runGoList("./...")
	if err != nil {
		log.Fatal("Error running go list:", err)
	}

	var allPackages []Package
	err = json.Unmarshal(allOutput, &allPackages)
	if err != nil {
		log.Fatal("Error parsing go list output:", err)
	}

	// Find the transitive dependencies of the initial packages
	transitiveDeps := make(map[string]bool)
	for _, pkg := range allPackages {
		if isTransitiveDep(pkg, initialDeps) {
			transitiveDeps[pkg.ImportPath] = true
		}
	}

	// Print the packages that transitively depend on the initial packages
	fmt.Println("Packages that transitively depend on the specified packages:")
	for pkg := range transitiveDeps {
		fmt.Println(pkg)
	}
}

func runGoList(args string) ([]byte, error) {
	cmd := exec.Command("go", "list", "-json", args)
	return cmd.Output()
}

func isTransitiveDep(pkg Package, initialDeps []string) bool {
	for _, dep := range pkg.Deps {
		for _, initialDep := range initialDeps {
			if dep == initialDep {
				return true
			}
		}
	}
	return false
}
