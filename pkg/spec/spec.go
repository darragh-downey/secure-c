package spec

// ... other imports

type SecurityIssue struct {
	ID             string `xml:"id,attr"`
	Type           string `xml:"Type"`
	Description    string `xml:"Description"`
	Severity       string `xml:"Severity"`
	Recommendation string `xml:"Recommendation"`
	URL            string `xml:"URL"` // Add the URL field
	// ... other fields (Likelihood, RemediationCost, etc.)
}

// ... (rest of your specifications package code)
