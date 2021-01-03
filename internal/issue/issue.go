package issue

import (
	"errors"
	"strings"
)

const (
	Main = "mainBranch"
	Raw  = "rawBranch"
)

// Feature is model for git branch -> tracker issue relation
type Feature struct {
	IssueName   string
	IssueType   string
	Description string
	FullName    string
}

// FromRefName converts raw ref name to smth to work with
func FromRefName(refName string) (*Feature, error) {
	parts := strings.SplitN(refName, "/", 2)
	if len(parts) < 1 {
		return nil, errors.New("not renough refname parts")
	}

	out := Feature{
		FullName:  refName,
		IssueType: Raw,
	}
	if refName == "main" || refName == "master" {
		out.IssueType = Main
	}
	if len(parts) == 1 {
		return &out, nil
	}

	out.IssueType = strings.ToUpper(parts[0])
	nameParts := strings.SplitN(parts[1], "-", 2)
	if len(nameParts) == 2 {
		out.IssueName = nameParts[0]
		out.Description = nameParts[1]
	}

	return &out, nil
}
