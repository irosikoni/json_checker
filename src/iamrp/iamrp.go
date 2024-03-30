package iamrp

import (
	"encoding/json"
	"fmt"
	"os"
)


type Statement struct {
	Sid string
	Effect string
	Action []string
	Resource string
}

type PolicyDocument struct {
	Version string
	Statement []Statement
}

type IAMRolePolicy struct {
	PolicyName string
	PolicyDocument PolicyDocument
}

func IAMRolePolicyValidator(filepath string){
	policy := loadAndParseJSON(filepath)
	fmt.Println(policy.isResourceAsterisk())
}

func loadAndParseJSON(filepath string) IAMRolePolicy{
	content, loadingErr := os.ReadFile(filepath)
	if loadingErr != nil {
		panic(loadingErr)
	}
	parsedPolicy := IAMRolePolicy{}
	parsingError := json.Unmarshal([]byte(content), &parsedPolicy)
	if parsingError != nil {
		panic(parsingError)
	}
	return parsedPolicy
}

func (policy IAMRolePolicy) isResourceAsterisk() bool{
	return policy.PolicyDocument.Statement[0].Resource != "*"
}

