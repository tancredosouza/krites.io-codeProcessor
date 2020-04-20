#include "CodeProcessor.h"
#include <cstdlib>
#include <fstream>
#include <filesystem>

const std::string OUTPUT_DIRECTORY_NAME = "output";
const std::string PROGRAM_FILEPATH = OUTPUT_DIRECTORY_NAME + "/prog.cpp";
const std::string EXEC_FILEPATH = OUTPUT_DIRECTORY_NAME + "/prog";

void CodeProcessor::compileCodeToExec(const std::string &CODE_RAW_TEXT)
{
  std::filesystem::create_directory(OUTPUT_DIRECTORY_NAME);
  std::ofstream(PROGRAM_FILEPATH) << CODE_RAW_TEXT;

  const std::string compileCodeCmd = "g++ -std=c++17 -o " + EXEC_FILEPATH + " " + PROGRAM_FILEPATH;
  system(compileCodeCmd.c_str());
}

void CodeProcessor::runExecWithInput(const std::string &CODE_INPUT_FILEPATH)
{
  std::string runCompiledCodeCmd = (EXEC_FILEPATH + " < " + CODE_INPUT_FILEPATH + " > ./output/actual_result.txt").c_str();
  system(runCompiledCodeCmd.c_str());
}

bool isEmptyFile(const char *filename)
{
  std::ifstream file(filename);
  return !file || file.get() == EOF;
}

bool CodeProcessor::assertCorrectOutput(const std::string &expectedResultFilepath)
{
  std::string compareResultFilesCmd = "diff ./output/actual_result.txt " + expectedResultFilepath + " > ./output/d.txt";
  system(compareResultFilesCmd.c_str());

  return isEmptyFile("./output/d.txt");
}