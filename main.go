package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Printf("Starting server for HTTP POST...\n")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleRequest(resWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.Error(resWriter, "404 not found.", http.StatusNotFound)
		return
	}

	setCorsHeaders(&resWriter)

	switch request.Method {
	case "POST":
		log.Println("Request received!")

		bodyBytes, _ := ioutil.ReadAll(request.Body)
		rand.Seed(time.Now().Unix())
		var submissionId int = rand.Intn(10000000)
		res := run(bodyBytes, submissionId)
		fmt.Fprintf(resWriter, res)
	default:
		fmt.Fprintf(resWriter, "Sorry, only the POST method is supported.")
	}
}


func createDirectory(directoryName string) {
	_, err := os.Stat(directoryName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(directoryName, 0755)
		if errDir != nil {
			log.Fatal("Error creating directory", err)
		}
	}
}

func run(dataReceived []byte, submissionId int) string {
	var outputDirectory string = "submission_" + strconv.Itoa(submissionId)

	createDirectory("./" + outputDirectory)
	ioutil.WriteFile("./"+ outputDirectory + "/codeTest.cpp", dataReceived, 0644)
	RunCommand("g++", "-std=c++17", "-o", "./" + outputDirectory + "/prog", "./"+ outputDirectory +"/codeTest.cpp");

	cmd := exec.Command("./"+ outputDirectory +"/prog")
	cmd.Stdin, _ = os.Open("./input/input.txt")
	var out bytes.Buffer
	cmd.Stdout = &out
	var done = make(chan struct{})

	go func(){
		err := cmd.Start()
		if err != nil {
			log.Fatal("Error running command", cmd, err)
		}
		cmd.Wait()
		done <- struct{}{}
	}()

	select {
	case <-done:
		ioutil.WriteFile("./" + outputDirectory + "/actualResult.txt", out.Bytes(), 0644)

		cmd := exec.Command("diff", "./" + outputDirectory + "/actualResult.txt", "./input/expected_result.txt")

		var out bytes.Buffer
		cmd.Stdout = &out

		cmd.Run()

		ioutil.WriteFile("./" + outputDirectory + "/d.txt", out.Bytes(), 0644)
		x := isFileEmpty( "./" + outputDirectory + "/d.txt")
		os.RemoveAll("./" + outputDirectory)

		if(x) {
			return "Correct!"
		} else {
			return "Incorrect."
		}
	case <-time.After(2 * time.Second):
		cmd.Process.Kill()
		return "Time Limit Exceeded."
	}
}

func setCorsHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}