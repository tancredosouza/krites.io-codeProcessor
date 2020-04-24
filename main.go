package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func handleSubmission(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	setCorsHeaders(&w)
	switch r.Method {
	case "POST":
		log.Println("Recebi uma request!")
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		ioutil.WriteFile("./input/codeTest.cpp", bodyBytes, 0644)

		res := run()
		fmt.Fprintf(w, res)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func setCorsHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func run() string {
	CompileCodeExecutable("./input/codeTest.cpp")
	cmd := exec.Command("./output/prog")
	cmd.Stdin, _ = os.Open("./input/input.txt")
	var out bytes.Buffer
	cmd.Stdout = &out
	var done = make(chan string)

	go func(){
		err := cmd.Start()
		if err != nil {
			log.Fatal("Error running command", cmd, err)
		}
		cmd.Wait()
		ioutil.WriteFile("./output/actualResult.txt", out.Bytes(), 0644)

		cmd := exec.Command("diff", "./output/actualResult.txt", "./input/expected_result.txt")

		var out bytes.Buffer
		cmd.Stdout = &out

		cmd.Run()

		ioutil.WriteFile("./output/d.txt", out.Bytes(), 0644)

		if(isFileEmpty( "./output/d.txt")) {
			done <- "CORRETO"
		} else {
			done <- "INCORRETO"
		}

		os.RemoveAll("./output")
	}()

	select {
	case res := <-done:
		return res
	case <-time.After(2 * time.Second):
		cmd.Process.Kill()
		return "TIME LIMIT EXCEEDED"
	}
}

func main() {
	http.HandleFunc("/", handleSubmission)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}