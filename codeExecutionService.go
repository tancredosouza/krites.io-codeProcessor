package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("CodeProcessor/execute_cpp_code", "CodeProcessor/input/codeTest.cpp", "CodeProcessor/input/input.txt", "CodeProcessor/input/expected_result.txt")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(out.String())
}
