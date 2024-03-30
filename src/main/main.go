package main

import (
	"os"

	"github.com/irosikoni/json_checker/src/iamrp"
)
func main(){
	filepath := os.Args[1]
	iamrp.IAMRolePolicyValidator(filepath)
}