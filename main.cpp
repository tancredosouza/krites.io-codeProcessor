#include <iostream>
#include <string>
#include "CodeProcessor.h"

using namespace std;

// Use raw string literal for easy coding
std::string prog = R"~(

#include <iostream>
using namespace std;
int fib(int n) {
  if (n <= 1) return 1;

  return fib(n-1) + fib(n-2);
}

int main()
{
  int n;
  while (cin >> n) {
    cout << fib(n) << '\n';
  }
}

)~"; // raw string literal stops here

int main()
{
  CodeProcessor::compileAndExecute(prog);

  if (CodeProcessor::assertOutput())
  {
    cout << "CORRECT" << endl;
  }
  else
  {
    cout << "INCORRECT" << endl;
  }

  return 0;
}