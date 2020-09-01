import subprocess

from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/', methods=["POST"])

def fff():
  response = {"msg": "oporra"}
  return jsonify(response)

if __name__ == '__main__':
  app.run(debug=True)

# cmd = "cpy.cpp"

# Flow
# 0. Estabelecer conexão rest
# 1. Receber código por REST
# 2. Salvar código em arquivo (dir. temporario)
# 3. Rodar código em C++
# 4. Retornar resultado por REST 



#try:
#  subprocess.check_output(["g++", cmd], stderr=subprocess.STDOUT)
#except subprocess.CalledProcessError as e:
#  print e.output

#subprocess.call("./a.out")