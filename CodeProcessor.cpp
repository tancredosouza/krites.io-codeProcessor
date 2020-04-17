#include "CodeProcessor.h"
#include <cstdlib>
#include <fstream>

void CodeProcessor::executeCode(const std::string &codeText)
{
  // save program to disk
  std::ofstream("prog.cpp") << codeText;

  system("g++ -o prog prog.cpp"); // compile
  system("./prog");               // run
}
