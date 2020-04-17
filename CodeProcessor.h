#include <string>

class CodeProcessor
{
public:
  static void compileAndExecute(const std::string &codeText);
  static bool assertOutput();

private:
  CodeProcessor() {}
};
