package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/step-security/secure-workflows/remediation/workflow"
)

func TestingSecureFlow() {
	queryPerms := map[string]string{
		"pinActions":        "true",
		"addHardenRunner":   "true",
		"addPermissions":    "true",
		"ignoreMissingKBs":  "true",
		"addProjectComment": "true",
	}
	input, _ := ioutil.ReadFile("./testfiles/secure/test1.yml")

	output, err := workflow.SecureWorkflow(queryPerms, string(input), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf(output.FinalOutput)
}
