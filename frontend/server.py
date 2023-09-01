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
