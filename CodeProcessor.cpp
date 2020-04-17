#include "CodeProcessor.h"
#include <cstdlib>
#include <fstream>
#include <filesystem>

void CodeProcessor::compileAndExecute(const std::string &codeText)
{
  // save program to disk
  std::ofstream("output/prog.cpp") << codeText;

  system("g++ -o output/prog output/prog.cpp");                           // compile
  system("./output/prog < input/input.txt > ./output/actual_result.txt"); // run
}

bool file_isempty(const char *filename)
{
  std::ifstream file(filename);
  return !file || file.get() == EOF;
}

bool CodeProcessor::assertOutput()
{
  system("diff ./output/actual_result.txt ./output/expected_result.txt > ./output/d.txt");

  return file_isempty("./output/d.txt");
}