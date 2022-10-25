package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/step-security/secure-workflows/remediation/docker"
)

func TestingPinningDocker() {
	input, _ := ioutil.ReadFile("./testfiles/docker/test1.yml")
	resp, err := docker.SecureDockerFile(string(input))
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf(resp.FinalOutput)
}
