import subprocess
import os
from Error import Error
import filecmp


CODE_FILENAME = "solution.cpp"
THREE_SECONDS = 3


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
            [f'{self.submissionDir}/a.out', f'<{infile}', f'>{outfile}'],
            capture_output=True,
            check=True,
            timeout=THREE_SECONDS
        )
        return (Error.NO_ERROR if filecmp.cmp("input/expected_result.txt", outfile)
                else Error.WRONG_ANSWER), None
