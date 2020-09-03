from flask import Flask, request, jsonify
import subprocess
import filecmp
import random
import os
import shutil

CODE_FILENAME = "solution.cpp"

app = Flask(__name__)


@app.route('/', methods=["POST"])
def handleSubmission():
    content = request.get_json()
    submissionCodeFile = temporarilyStoreCodeFile(content["submitted_code"])
    submissionDirectory = os.path.dirname(submissionCodeFile)
    result = compileAndRun(submissionCodeFile)
    shutil.rmtree(submissionDirectory)
    return jsonify(result)


def temporarilyStoreCodeFile(codeText):
    codeDirectory = f"submission_{random.randint(0,9999999)}"
    os.mkdir(codeDirectory)
    codeFilepath = os.path.join(codeDirectory, CODE_FILENAME)

    codeFile = open(codeFilepath, "w")
    codeFile.write(codeText)
    codeFile.close()
    return codeFilepath


def compileAndRun(codeFilepath):
    subprocess.check_call(
        ('g++', '-o', 'a.out', codeFilepath),
        stdin=subprocess.DEVNULL)

    with open('input/input.txt') as infile, open('input/output.txt', 'w') as outfile:
        subprocess.check_call(
            ('./a.out',),
            stdin=infile,
            stdout=outfile,
            universal_newlines=True)
        return filecmp.cmp(outfile.name, "input/expected_result.txt")


if __name__ == '__main__':
    app.run(debug=True)
