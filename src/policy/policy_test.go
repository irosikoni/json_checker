package policy

import (
	"testing"
)

func TestIsResourceAsterisk(t *testing.T) {
	goodPolicy := IAMRolePolicy{
		PolicyName: "",
		PolicyDocument: PolicyDocument{
			Version: "",
			Statement: []Statement{
				{
					Sid:      "",
					Effect:   "",
					Action:   []string{""},
					Resource: "*",
				},
			},
		},
	}
	badPolicy := goodPolicy
	badPolicy.PolicyDocument.Statement[0].Resource = "*sth"

	t.Run("Positive", func(t *testing.T) {
		if goodPolicy.isResourceAsterisk() {
			t.Log("Founding asterisk resulted in True")
		}
		if !badPolicy.isResourceAsterisk() {
			t.Log("Founding non-asterisk resulted in False")
		}
	})
}

func TestIsResourceAsteriskSub(t *testing.T) {
	policyWithMissingResource := IAMRolePolicy{
		PolicyName: "",
		PolicyDocument: PolicyDocument{
			Version:   "",
			Statement: []Statement{},
		},
	}
	policyWithTooManyStatements := IAMRolePolicy{
		PolicyName: "",
		PolicyDocument: PolicyDocument{
			Version: "",
			Statement: []Statement{
				{
					Sid:      "",
					Effect:   "",
					Action:   []string{""},
					Resource: "*",
				},
				{
					Sid:      "",
					Effect:   "",
					Action:   []string{""},
					Resource: "*",
				},
			},
		},
	}

	t.Run("Error checking - missing resource", func(t *testing.T) {
		defer func() {
			if recovered := recover(); recovered != nil {
				t.Log("Error checking for missing resource field")
			}
		}()
		policyWithMissingResource.isResourceAsterisk()
	})
	t.Run("Error checking - too many statements", func(t *testing.T) {
		defer func() {
			if recovered := recover(); recovered != nil {
				t.Log("Error checking for multi-statement policy")
			}
		}()
		policyWithTooManyStatements.isResourceAsterisk()
	})
}

func TestLoadAndParseJSON(t *testing.T) {
	nonexistingFilepath := "thereisnosuchfile"
	t.Run("Error checking - no file", func(t *testing.T) {
		defer func() {
			if recovered := recover(); recovered != nil {
				t.Log("Error checking for non-existing filepath")
			}
		}()
		loadAndParseJSON(nonexistingFilepath)
	})
}
