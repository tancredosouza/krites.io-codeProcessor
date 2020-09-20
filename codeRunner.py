from flask import Flask, request, jsonify
from flask_cors import CORS
import subprocess
import filecmp
import random
import os
import shutil
from CppEvaluator import CppEvaluator
from enum import Enum

CODE_FILENAME = "solution.cpp"
THREE_SECONDS = 3


class Error(Enum):
    NO_ERROR = 0
    COMPILE = 1
    EXECUTION = 2
    TIMEOUT = 3
    WRONG_ANSWER = 4


app = Flask(__name__)
CORS(app)


@app.route('/', methods=["POST"])
def handleSubmission():
    content = request.get_json()
    submissionDir = temporarilyStoreCodeFile(content["submitted_code"])
    evaluator = CppEvaluator(submissionDir)
    err_type, err_body = evaluator.tryCompileAndRun()
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


def buildMsg(error_type, error_body):
    return jsonify(
        err_type="correct" if error_type == Error.NO_ERROR
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
