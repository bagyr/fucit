package issue

import (
	"testing"
)

func TestFromRefName(t *testing.T) {
	mainRef := "main"
	i, err := FromRefName(mainRef)
	if err != nil {
		t.Fatal(err)
	}
	if i.IssueType != Main {
		t.Errorf("main did not detected: %v", i)
	}

	issueRef := "feature/GH1-feature_detection"
	i, err = FromRefName(issueRef)
	if err != nil {
		t.Fatal(err)
	}
	if i.IssueType != "feature" {
		t.Errorf("feature type did not detected: %v", i)
	}
	if i.IssueName != "GH1" {
		t.Errorf("feature name did not detected: %v", i)
	}
	if i.Description != "feature_detection" {
		t.Errorf("feature description did not detected: %v", i)
	}
}
