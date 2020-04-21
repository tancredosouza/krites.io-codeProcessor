package main

import (
	"fmt"
)


func main() {
	CompileCodeExecutable("./input/codeTest.cpp")
	RunCmdWithInput("./output/prog","./input/input.txt", "./output/actualResult.txt")
	res := AssertCorrectOutput("./input/expected_result.txt", "./output/actualResult.txt")

	fmt.Println(res)
}
