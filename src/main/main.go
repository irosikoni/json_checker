package main

import (
	"fmt"
	"os"

	"github.com/irosikoni/json_checker/src/policy"
)

func main() {
	filepath := os.Args[1]
	fmt.Println(policy.IAMRolePolicyValidator(filepath))
}
