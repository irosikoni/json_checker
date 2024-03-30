package main

import (
	"fmt"
	"os"

	"github.com/irosikoni/json_checker/src/iamrp"
)
func main(){
	filepath := os.Args[1]
	fmt.Println(iamrp.IAMRolePolicyValidator(filepath))
}