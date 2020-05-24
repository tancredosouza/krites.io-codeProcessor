package main

import (
	"log"
	"net/http"
	"os"

	"bytes"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.POST("/evaluate", func (c *gin.Context) { evaluateCode(c) })

	router.Run(":" + port)
}

func evaluateCode(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])

	var submissionId int = rand.Intn(10000000)
	
	res := run([]byte(reqBody), submissionId)
	log.Println("answer ---- ", res)
	c.String(http.StatusOK, res)
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
	log.Println(string(dataReceived))
	createDirectory("./" + outputDirectory)
	log.Println("--------------- Directory created!!!!")
	ioutil.WriteFile("./"+ outputDirectory + "/codeTest.c", dataReceived, 0644)
	log.Println("--------------- Output file written!!!!")
	err := RunCommand("g++", "-o", "./" + outputDirectory + "/prog", "./"+ outputDirectory +"/codeTest.c");

	if (err != nil) {
		os.RemoveAll("./" + outputDirectory)
		return "Compilation or Execution error!"
	}

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
		os.RemoveAll("./" + outputDirectory)
		return "Time Limit Exceeded."
	}
}