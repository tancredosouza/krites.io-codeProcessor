from flask import Flask, request, jsonify
import subprocess
import filecmp
import random
import os
import shutil
from enum import Enum

CODE_FILENAME = "solution.cpp"


class Error(Enum):
    NO_ERROR = 0
    COMPILE = 1
    EXECUTION = 2
    TIMEOUT = 3
    WRONG_ANSWER = 4


app = Flask(__name__)


@app.route('/', methods=["POST"])
def handleSubmission():
    content = request.get_json()
    submissionDir = temporarilyStoreCodeFile(content["submitted_code"])
    err_type, err_body = compileAndRun(submissionDir)
    shutil.rmtree(submissionDir)
    return buildMsg(err_type, err_body)


def temporarilyStoreCodeFile(codeText):
    codeDir = f"submission_{random.randint(0,9999999)}"
    os.mkdir(codeDir)
    codeFilepath = os.path.join(codeDir, CODE_FILENAME)

    codeFile = open(codeFilepath, "w")
    codeFile.write(codeText)
    codeFile.close()
    return codeDir


def compileAndRun(submissionDir):
    codeFilepath = os.path.join(submissionDir, CODE_FILENAME)
    try:
        subprocess.check_output(
            ('g++', '-o', 'a.out', codeFilepath),
            stdin=subprocess.DEVNULL,
            stderr=subprocess.STDOUT)
    except subprocess.CalledProcessError as e:
        return Error.COMPILE, e.output

    try:
        with open('input/input.txt') as infile, open(f'{submissionDir}/output.txt', 'w') as outfile:
            subprocess.run(
                './a.out',
                stdin=infile,
                stdout=outfile,
                stderr=subprocess.STDOUT,
                universal_newlines=True,
            )
            return (Error.NO_ERROR if filecmp.cmp("input/expected_result.txt", outfile.name) else Error.WRONG_ANSWER), None
    except subprocess.TimeoutExpired as e:
        return Error.TIMEOUT, None
    except subprocess.CalledProcessError as e:
        return Error.EXECUTION, e.output


def buildMsg(error_type, error_body):
    return jsonify(
        err_type="none" if error_type == Error.NO_ERROR
        else "wa" if error_type == Error.WRONG_ANSWER
        else "timeout" if error_type == Error.TIMEOUT
        else "compile" if error_type == Error.COMPILE
        else "exec" if error_type == Error.EXECUTION
        else "unknown",

        err_body=error_body.decode(
            "utf-8") if error_body != None else error_body
    )


if __name__ == '__main__':
    app.run(debug=True)
