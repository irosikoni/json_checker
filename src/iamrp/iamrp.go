package iamrp

import (
	"bytes"
	"encoding/json"
	"os"
)


type Statement struct {
	Sid string
	Effect string
	Principal map[string]string
	Action []string
	Resource string
	Condition map[string]map[string]string
}

type PolicyDocument struct {
	Version string
	Statement []Statement
}

type IAMRolePolicy struct {
	PolicyName string
	PolicyDocument PolicyDocument
}

func IAMRolePolicyValidator(filepath string) bool{
	policy := loadAndParseJSON(filepath)
	return policy.isResourceAsterisk()
}

func loadAndParseJSON(filepath string) IAMRolePolicy{
	content, loadingErr := os.ReadFile(filepath)
	if loadingErr != nil {
		if os.IsNotExist(loadingErr) {
			panic("File does not exist")
		} else {
		panic(loadingErr)
		}
	}
	var parsedPolicy IAMRolePolicy
	parsingError := StrictUnmarshal([]byte(content), &parsedPolicy)
	if parsingError != nil {
		panic("Error parsing JSON, check if file has AWS::IAM::Role Policy structure")
	}
	return parsedPolicy
}

func (policy IAMRolePolicy) isResourceAsterisk() bool{
	if len(policy.PolicyDocument.Statement) == 0 {
		panic("No resource field in the policy")
	}
	if len(policy.PolicyDocument.Statement) > 1 {
		panic("Multi-statement policy, cannot validate exact resource field")
	}
	return policy.PolicyDocument.Statement[0].Resource != "*"
}

func StrictUnmarshal(data []byte, v interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	return dec.Decode(v)
}

