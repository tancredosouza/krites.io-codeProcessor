#include <iostream>
#include <string>
#include "CodeProcessor.h"

using namespace std;

// Use raw string literal for easy coding
std::string prog = R"~(

#include <iostream>

int fib(int n) {
  if (n <= 1) return 1;

  return fib(n-1) + fib(n-2);
}

int main()
{
    std::cout << fib(10) << '\n';
}

)~"; // raw string literal stops here

int main()
{
  CodeProcessor::executeCode(prog);
  return 0;
}