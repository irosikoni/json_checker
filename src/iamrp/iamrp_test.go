package iamrp

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
					Sid: "",
					Effect: "",
					Action: []string{""},
					Resource: "*",
				},
			},
		},
	}
	badPolicy := goodPolicy
	badPolicy.PolicyDocument.Statement[0].Resource = "*sth"

	t.Run("Positive", func(t *testing.T){
		if goodPolicy.isResourceAsterisk() {
			t.Log("Founding asterisk resulted in True")
		}
		if !badPolicy.isResourceAsterisk() {
			t.Log("Founding non-asterisk resulted in False")
		}
	})
}
