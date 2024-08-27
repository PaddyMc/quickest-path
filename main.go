package main

import (
	"fmt"
)

// Find the quickest path to publickey
func findAuthenticationPath(docs Documents, start, target string) []string {
	// write your code here....

	return nil
}

// testFindPath is a simple helper function to allow multiple tests easily
func testFindPath(docs Documents, start, target, testName string) {
	fmt.Printf("\n%s:\n", testName)
	fmt.Printf("Start: %s, Target: %s\n", start, target)

	path := findAuthenticationPath(docs, start, target)

	if path == nil {
		fmt.Println("No path found")
	} else {
		fmt.Println("Quickest path:", path)
		fmt.Println("Number of hops:", len(path)-1)
	}
}

// runs 3 simple tests to find the authentication path for each document
func main() {
	docs := getDocuments()

	// Test case 1: Path found
	start1 := "did:example:1"
	target1 := "did:example:10"
	testFindPath(docs, start1, target1, "Test case 1 (Path found)")

	// Test case 2: Path not found
	start2 := "did:example:1"
	target2 := "did:example:11"
	testFindPath(docs, start2, target2, "Test case 2 (Path not found)")

	// Test case 3: Another path found
	start3 := "did:example:2"
	target3 := "did:example:9"
	testFindPath(docs, start3, target3, "Test case 3 (Another path found)")
}
