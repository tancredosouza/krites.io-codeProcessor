import subprocess
import os
from enum import Enum
import filecmp


CODE_FILENAME = "solution.cpp"
THREE_SECONDS = 3


class Error(Enum):
    NO_ERROR = 0
    COMPILE = 1
    EXECUTION = 2
    TIMEOUT = 3
    WRONG_ANSWER = 4


class CppEvaluator:
    def __init__(self, submissionDir):
        self.submissionDir = submissionDir

    def tryCompileAndRun(self):
        try:
            self.compile()
        except subprocess.CalledProcessError as e:
            return Error.COMPILE, e.output

        try:
            return self.run()
        except subprocess.TimeoutExpired as e:
            return Error.TIMEOUT, None
        except subprocess.CalledProcessError as e:
            return Error.EXECUTION, e.stderr

    def compile(self):
        codeFilepath = os.path.join(self.submissionDir, CODE_FILENAME)
        subprocess.check_output(
            ('g++', '-w', '-o', f'{self.submissionDir}/a.out', codeFilepath),
            stdin=subprocess.DEVNULL,
            stderr=subprocess.STDOUT)

    def run(self):
        infile = 'input/input.txt'
        outfile = f'{self.submissionDir}/output.txt'
        subprocess.run(
            f'{self.submissionDir}/a.out <{infile} >{outfile}',
            capture_output=True,
            shell=True,
            check=True,
            timeout=THREE_SECONDS
        )
        return (Error.NO_ERROR if filecmp.cmp("input/expected_result.txt", outfile)
                else Error.WRONG_ANSWER), None
