from flask import Flask, request, jsonify
from flask_cors import CORS
import subprocess
import filecmp
import random
import os
import shutil
from CppEvaluator import CppEvaluator
from PythonEvalutor import PythonEvaluator
from Error import Error

CODE_FILENAME = "solution"
THREE_SECONDS = 3

app = Flask(__name__)
CORS(app)


@app.route('/', methods=["POST"])
def handleSubmission():
    content = request.get_json()

    submissionLanguage = content["submitted_language"]
    submissionDir = temporarilyStoreCodeFile(
        content["submitted_code"], submissionLanguage)

    evaluator = getEvaluator(submissionDir, submissionLanguage)
    err_type, err_body = evaluator.tryCompileAndRun()

    shutil.rmtree(submissionDir)

    return buildMsg(err_type, err_body)


def temporarilyStoreCodeFile(codeText, language):
    codeDir = f"submission_{random.randint(0,9999999)}"
    os.mkdir(codeDir)
    codeFilepath = os.path.join(codeDir, f'{CODE_FILENAME}.{language}')

    codeFile = open(codeFilepath, "w")
    codeFile.write(codeText)
    codeFile.close()
    return codeDir


def getEvaluator(submissionDir, submissionLanguage):
    return CppEvaluator(submissionDir) if submissionLanguage == "cpp" else PythonEvaluator(
        submissionDir) if submissionLanguage == "py" else None


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
