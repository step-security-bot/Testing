package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/step-security/secure-workflows/remediation/dependabot"
)

func TestingDependabot() {
	test := struct {
		fileName   string
		Ecosystems []dependabot.Ecosystem
		isChanged  bool
	}{
		fileName:   "Same-ecosystem-different-directory.yml",
		Ecosystems: []dependabot.Ecosystem{{PackageEcosystem: "github-actions", Directory: "/", Interval: "daily"}, {PackageEcosystem: "npm", Directory: "/sample", Interval: "daily"}},
		isChanged:  true,
	}
	// add more test using loop
	var updateDependabotConfigRequest dependabot.UpdateDependabotConfigRequest
	input, err := ioutil.ReadFile("./testfiles/dependabot/test1.yml")
	if err != nil {
		fmt.Printf("not reading file")
		return
	}
	updateDependabotConfigRequest.Content = string(input)
	updateDependabotConfigRequest.Ecosystems = test.Ecosystems
	inputRequest, err := json.Marshal(updateDependabotConfigRequest)
	if err != nil {
		log.Fatal(err)
		return
	}
	output, err := dependabot.UpdateDependabotConfig(string(inputRequest))
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf(output.FinalOutput)
}
