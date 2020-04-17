#include "CodeProcessor.h"
#include <cstdlib>
#include <fstream>

void CodeProcessor::executeCode(const std::string &codeText)
{
  // save program to disk
  std::ofstream("output/prog.cpp") << codeText;

  system("g++ -o output/prog output/prog.cpp"); // compile
  system("./output/prog");                      // run
}
