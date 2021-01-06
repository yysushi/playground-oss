from flask import Flask, request

app = Flask(__name__)


@app.route('/')
def hello():
  token = request.args.get('token')
  return f'authorized with {token}'


if __name__ == "__main__":
  app.run(host='0.0.0.0', port=8080, debug=False)
