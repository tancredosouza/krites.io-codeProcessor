import filecmp
from Error import Error
import os
import subprocess


CODE_FILENAME = "solution.py"
THREE_SECONDS = 3


class PythonEvaluator:
    def __init__(self, submissionDir):
        self.submissionDir = submissionDir

    def tryCompileAndRun(self):
        try:
            return self.compileAndRun()
        except subprocess.TimeoutExpired as e:
            return Error.TIMEOUT, None
        except subprocess.CalledProcessError as e:
            return Error.EXECUTION, e.stderr

    def compileAndRun(self):
        codeFilepath = os.path.join(self.submissionDir, CODE_FILENAME)
        outfile = f'{self.submissionDir}/output.txt'

        subprocess.run(
            ['python', codeFilepath, '< input/input.txt', f'> {outfile}'],
            capture_output=True,
            check=True,
            timeout=THREE_SECONDS
        )
        return (Error.NO_ERROR if filecmp.cmp("input/expected_result.txt", outfile)
                else Error.WRONG_ANSWER), None
