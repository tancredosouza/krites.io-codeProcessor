# Code Processor

This C++ program compiles and executes another C++17 program.
It then compares its output with the expected one.

This code is being developed as a submodule to the main (redis.io repo)[https://github.com/tancredosouza/redis.io]: an open-source online judge.

## Running

```
$ g++ -std=c++17 main.cpp CodeProcessor.cpp -o execute_cpp_code
$ execute_cpp_code [CODE FILE PATH] [CODE INPUT FILE PATH] [CODE EXPECTED OUTPUT FILE PATH]
```
