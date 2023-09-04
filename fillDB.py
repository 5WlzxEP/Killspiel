import dataclasses
import random
import sqlite3
from faker import Faker


@dataclasses.dataclass
class Vote:
    game: int
    vote: float


class User:
    def __init__(self):
        self._id: int = random.randint(0, 1_000_000)
        self._name: str = faker.name()
        self._guesses: int = 0
        self._points: int = 0
        self._latest: int = 0

        self._votes: [Vote] = []

    def saveDB(self, cursor: sqlite3.Cursor):
        cursor.execute(
            f"INSERT INTO Users VALUES ({self._id}, '{self._name}', {self._guesses}, {self._points}, {self._latest})")

        for vote in self._votes:
            cursor.execute(f"INSERT INTO Votes VALUES ({vote.game}, {self._id}, {vote.vote})")

    def addVote(self, game: int, correct: int):
        vote = random.randint(0, 15)
        self._votes.append(Vote(game=game, vote=vote))
        self._guesses += 1
        self._latest <<= 1
        self._latest += 1 if vote == correct else 0
        self._points += 1 if vote == correct else 0
        self._latest %= 256


faker = Faker()
with sqlite3.connect("config/db.sqlite") as db:
    cursor = db.cursor()

    users = []
    for _ in range(100):
        users.append(User())

    for i in range(200):
        correct = random.randint(0, 15)
        cursor.execute(f"INSERT INTO Game (correct) VALUES ({float(correct)});")
        cursor.execute("SELECT last_insert_rowid()")
        game = cursor.fetchone()[0]

        for user in users:
            user.addVote(game, correct)

    for user in users:
        user.saveDB(cursor)

    # for i in range(100):
    #     g = random.randint(10, 150)
    #     cursor.execute(f"INSERT INTO Users VALUES ({random.randint(0, 1_000_000)}, '{faker.name()}', {g}, {random.randint(0, g)}, {random.randint(0, 256)})")

    db.commit()
