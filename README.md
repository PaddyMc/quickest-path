# DID Document Authentication Path Finder

## Task Description

This project implements an algorithm to find the quickest authentication path between two Decentralized Identifier (DID) documents in a network of interconnected DIDs. The goal is to traverse the network using verification methods to establish a path from a start DID to a target DID.

## Step 1: Graph Structure

The network of DID documents is represented as a directed graph:

1. Each node in the graph is a DID document, containing:
   - An ID (the DID itself)
   - A list of verification methods (links to other DIDs)
   - A public key

2. The edges of the graph are the verification methods, pointing from one DID to another.

3. The graph is directed, meaning that if DID A has a verification method pointing to DID B, it doesn't necessarily mean that B points back to A.

4. The graph structure is defined in the `getDocuments()` function, which returns a map of DID documents.

Example structure:

```go
type Document struct {
    ID                 string
    VerificationMethod []string
    PubKey             string
}

func getDocuments() Documents {
    // ... (document definitions)
}
```

Goal of Step 1:

- Define the Documents data structure


## Step 2: Path Finding Algorithm

The algorithm to find the quickest authentication path is implemented in the findAuthenticationPath function:

- The algorithm starts from the start DID and explores all neighboring DIDs.
- The search continues until the target DID is found or all reachable DIDs have been explored.
- If a path is found, it's reconstructed using a prev map that stores the previous DID for each visited DID.

Key points of the algorithm:

1. It finds the shortest path in terms of the number of hops between DIDs.
2. It can handle cases where no path exists between the start and target DIDs.
3. The algorithm is efficient for large graphs as it doesn't explore unnecessary paths.

Usage:

```go
path := findAuthenticationPath(docs, startDID, targetDID)
```

Goal of Step 2:

- Write the findAuthenticationPath function

## Running the Tests
The main function includes multiple test cases to demonstrate the algorithm's functionality:

1. A case where a path exists between two DIDs.
2. A case where no path exists between two DIDs.
3. An additional case with a different start and target DID.

To run the tests, simply execute the main.go file:
```go
run main.go
```

The output will show the results for each test case, including the path found (if any) and the number of hops.

```bash
Test case 1 (Path found):
Start: did:example:1, Target: did:example:10
Quickest path: [did:example:1 did:example:3 did:example:5 did:example:8 did:example:10]
Number of hops: 4

Test case 2 (Path not found):
Start: did:example:1, Target: did:example:11
No path found

Test case 3 (Another path found):
Start: did:example:2, Target: did:example:9
Quickest path: [did:example:2 did:example:5 did:example:3 did:example:6 did:example:9]
Number of hops: 4
```
