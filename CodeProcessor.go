package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func CompileCodeExecutable(codeFilepath string) {
	createDirectory("./output")
	RunCommand("g++", "-std=c++17", "-o", "./output/prog", codeFilepath);
}

func createDirectory(directoryName string) {
	_, err := os.Stat(directoryName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(directoryName, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}

func RunCommand(exe string, args ...string) {
	cmd := exec.Command(exe, args...)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func RunCmdWithInput(execFilepath string, inputFilepath string, outputFilepath string) {
	cmd := exec.Command(execFilepath)
	cmd.Stdin, _ = os.Open(inputFilepath)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)

	}
	ioutil.WriteFile(outputFilepath, out.Bytes(), 0644)
}

func isFileEmpty(filepath string) bool {
	file, err := os.Stat(filepath)

	if err != nil {
		log.Fatal(err)
	}

	return file.Size() == 0
}

func AssertCorrectOutput(actualFilepath string, expectedFilepath string) bool {
	cmd := exec.Command("diff", actualFilepath, expectedFilepath)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)

	}
	ioutil.WriteFile("./output/d.txt", out.Bytes(), 0644)

	ans := isFileEmpty( "./output/d.txt");

	os.RemoveAll("./output")
	return ans
}