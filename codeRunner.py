import subprocess

from flask import Flask, request, jsonify

app = Flask(__name__)


@app.route('/', methods=["POST"])
def handleRequest():
    response = {"msg": request.args.get('text', '')}
    content = request.get_json(silent=True)
    tempFilepath = temporarilyStoreCodeFile(content["submitted_code"])
    compileAndRun(tempFilepath)
    return jsonify(response)


def temporarilyStoreCodeFile(codeText):
    # TODO: add random directory to be deleted after execution
    codeFilepath = "sample.cpp"
    textFile = open(codeFilepath, "w")
    textFile.write(codeText)
    textFile.close()
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

    subprocess.run(
        "diff -u ./input/expected_result.txt ./input/output.txt > dif.txt", shell=True)


if __name__ == '__main__':
    app.run(debug=True)
