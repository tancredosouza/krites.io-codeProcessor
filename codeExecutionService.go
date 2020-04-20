package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func runCommand(exe string, args ...string) {
	cmd := exec.Command(exe, args...)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(out.String())
}

func main() {
	runCommand("CodeProcessor/execute_cpp_code",
		"CodeProcessor/input/codeTest.cpp",
		"CodeProcessor/input/input.txt",
		"CodeProcessor/input/expected_result.txt")
}
