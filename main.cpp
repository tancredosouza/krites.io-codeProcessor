#include <iostream>
#include <string>
#include "CodeProcessor.h"
#include <fstream>
#include <sstream>

using namespace std;

void printHelpMessage()
{
  cout << "Invalid passed arguments." << endl;
  cout << "Make sure you're passing the arguments on the following format: " << endl;
  cout << "execute_cpp_code [CODE FILE PATH] [CODE INPUT FILE PATH] [CODE EXPECTED OUTPUT FILE PATH]" << endl;
}

string getTextFromFile(const string &textFilepath)
{
  std::ifstream t(textFilepath);
  std::stringstream buffer;
  buffer << t.rdbuf();

  return buffer.str();
}

int main(int argc, char *argv[])
{
  if (argc != 4)
  {
    printHelpMessage();
    exit(1);
  }

  CodeProcessor::compileCodeToExec(getTextFromFile(argv[1]));
  CodeProcessor::runExecWithInput(argv[2]);

  if (CodeProcessor::assertCorrectOutput(argv[3]))
  {
    cout << "CORRECT" << endl;
  }
  else
  {
    cout << "INCORRECT" << endl;
  }

  return 0;
}