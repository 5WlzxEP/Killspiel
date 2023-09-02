import random
import string

from flask import Flask, request
from flask_cors import CORS
from faker import Faker
from random import randint

app = Flask(__name__)
CORS(app)

Leaderboard = {
    "meta": {
        "size": 0,
    },
    "data": [
    ]
}

faker = Faker()
data = []
points = 1_000
for i in range(1, 102):
    points = randint(points - 15, points)
    data.append({
        "id": i,
        "rank": i,
        "name": faker.name(),
        "points": points,
        "guesses": randint(points, 1_000),
        "latest": randint(0, 255)
    })

Leaderboard["meta"]["size"] = 100

chat_config = {
    "apiKey": "val.apiKey",
    "channel": "val.channel",
    "channelApi": "val.channelApi",
    "prefix": "val.prefix",
    "msgBeginn": "val.msgBeginn",
    "msgEnd": "val.msgEnd",
}


class Buffer:
    def __init__(self):
        self._buf = [i for i in range(10)]
        self._index = 0

    def add(self, val) -> None:
        self._buf[self._index] = val
        self._index = (self._index + 1) % 10

    def __getitem__(self, item) -> object | None:
        if item < 10:
            return self._buf[item]

    def __contains__(self, item) -> bool:
        return item in self._buf


buf = Buffer()


def genKey() -> str:
    return "".join(random.choices(string.ascii_lowercase, k=15))

def sortdata(sort: str, o: bool):
    global data
    if sort == "p":
        data.sort(key=lambda x: x['points'], reverse=not o)
    elif sort == "n":
        data.sort(key=lambda x: x['name'], reverse=o)
    elif sort == "g":
        data.sort(key=lambda x: x['guesses'], reverse=not o)


@app.route("/")
def test():
    return "<p>Hallo Welt</p>"


@app.route("/api/leaderboard")
def leaderboard():
    # time.sleep(.5)
    global Leaderboard
    global data
    limit = min(int(request.args.get("l", 25)), len(data) - 1)
    offset = int(request.args.get("p", 0))
    end = min(offset + limit, len(data) - 1)

    s = request.args.get("s", "p")
    o = request.args.get("o", 0) == "0"

    sortdata(s, o)

    Leaderboard["data"] = data[offset:end]
    return Leaderboard


@app.route("/api/collector/chat/", methods=["GET", "POST"])
def chat():
    global chat_config
    if request.method == "GET":
        return chat_config
    elif request.method == "POST":
        data = request.json
        chat_config = data
        return "", 200
    return "", 404


@app.route("/api/user/<int:id>/", methods=["GET", "DELETE"])
def user(id: int):
    if request.method == "GET":
        return {
            "name": "TestUser123",
            "id": id,
            "guesses": randint(100, 1_000),
            "points": randint(0, 100),
            "latest": randint(0, 255),
            "history": [
                {"game": i, "guess": randint(0, 10), "correct": randint(0, 10)}
                for i in range(10)
            ]
        }
    if request.is_json:
        res = request.get_json()
        if res["key"] not in buf:
            return "", 400
        return "", 204
    else:
        key = genKey()
        buf.add(key)
        return {"key": key}
