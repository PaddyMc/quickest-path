package main

// Document represents a data structure for representing a collection of public keys
type Document struct {
	ID                 string
	VerificationMethod []string
	PubKey             string
}

type Documents map[string]Document

/* These are example documents to use to populate your get documents function
var exampleDocument = Document{
	ID:                 "did:example:1",
	VerificationMethod: []string{"did:example:2", "did:example:3"},
	PubKey:             "pubkey01",
}

var exampleDocument2 = Document{
	ID:                 "did:example:2",
	VerificationMethod: []string{"did:example:3", "did:example:4"},
	PubKey:             "pubkey01",
}
*/

func getDocuments() Documents {
	docs := make(Documents)

	docs["did:example:1"] = Document{
		ID:                 "did:example:1",
		VerificationMethod: []string{"did:example:2", "did:example:3", "did:example:5"},
		PubKey:             "pubkey01",
	}
	docs["did:example:2"] = Document{
		ID:                 "did:example:2",
		VerificationMethod: []string{"did:example:1", "did:example:4", "did:example:6"},
		PubKey:             "pubkey02",
	}
	docs["did:example:3"] = Document{
		ID:                 "did:example:3",
		VerificationMethod: []string{"did:example:1", "did:example:5", "did:example:7"},
		PubKey:             "pubkey03",
	}
	docs["did:example:4"] = Document{
		ID:                 "did:example:4",
		VerificationMethod: []string{"did:example:2", "did:example:8", "did:example:9"},
		PubKey:             "pubkey04",
	}
	docs["did:example:5"] = Document{
		ID:                 "did:example:5",
		VerificationMethod: []string{"did:example:1", "did:example:3", "did:example:10"},
		PubKey:             "pubkey05",
	}
	docs["did:example:6"] = Document{
		ID:                 "did:example:6",
		VerificationMethod: []string{"did:example:2", "did:example:8", "did:example:10"},
		PubKey:             "pubkey06",
	}
	docs["did:example:7"] = Document{
		ID:                 "did:example:7",
		VerificationMethod: []string{"did:example:3", "did:example:9", "did:example:10"},
		PubKey:             "pubkey07",
	}
	docs["did:example:8"] = Document{
		ID:                 "did:example:8",
		VerificationMethod: []string{"did:example:4", "did:example:6", "did:example:9"},
		PubKey:             "pubkey08",
	}
	docs["did:example:9"] = Document{
		ID:                 "did:example:9",
		VerificationMethod: []string{"did:example:4", "did:example:7", "did:example:8"},
		PubKey:             "pubkey09",
	}
	docs["did:example:10"] = Document{
		ID:                 "did:example:10",
		VerificationMethod: []string{"did:example:5", "did:example:6", "did:example:7"},
		PubKey:             "pubkey10",
	}
	docs["did:example:11"] = Document{
		ID:                 "did:example:11",
		VerificationMethod: []string{"did:example:12"},
		PubKey:             "pubkey11",
	}
	docs["did:example:12"] = Document{
		ID:                 "did:example:12",
		VerificationMethod: []string{"did:example:11"},
		PubKey:             "pubkey12",
	}

	return docs
}
