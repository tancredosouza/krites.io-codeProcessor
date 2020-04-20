#include <string>

class CodeProcessor
{
public:
  static void compileCodeToExec(const std::string &codeText);
  static void runExecWithInput(const std::string &codeInputFilepath);
  static bool assertCorrectOutput(const std::string &expectedResultFilepath);

private:
  CodeProcessor() {}
};
