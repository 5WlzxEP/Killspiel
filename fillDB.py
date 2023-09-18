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
        self._id: int = random.randint(0, 100_000_000)
        self._name: str = (faker.name() + "_" + str(random.randint(0, 10000))).lower()
        self._guesses: int = 0
        self._points: int = 0
        self._latest: int = 0

        self._votes: [Vote] = []

    def save_db(self, cursor: sqlite3.Cursor):
        cursor.execute(
            f"INSERT INTO Users VALUES ({self._id}, '{self._name}', {self._guesses}, {self._points}, {self._latest})")

        for vote in self._votes:
            cursor.execute(f"INSERT INTO Votes VALUES ({vote.game}, {self._id}, {vote.vote})")

    def add_vote(self, game: int, correct: int):
        vote = random.randint(0, 15)
        self._votes.append(Vote(game=game, vote=vote))
        self._guesses += 1
        self._latest <<= 1
        self._latest += 1 if vote == correct else 0
        self._points += 1 if vote == correct else 0
        self._latest %= 256


faker = Faker(use_weighting=False)
while True:
    try:
        with sqlite3.connect("config/db.sqlite") as db:
            cursor = db.cursor()

            users = []
            print("Generating Users")
            for _ in range(2_000):
                users.append(User())

            for i in range(6_000):
                if i % 100 == 0:
                    print(f"Generating game {i}")
                correct = random.randint(0, 15)
                cursor.execute(f"INSERT INTO Game (correct, info) VALUES ({float(correct)}, 'generated test data');")
                cursor.execute("SELECT last_insert_rowid()")
                game = cursor.fetchone()[0]

                for user in users:
                    if random.randint(0, 10) < 5:
                        user.add_vote(game, correct)

            for i, user in enumerate(users):
                if i % 10 == 0:
                    print(f"saving user nr. {i}")
                user.save_db(cursor)

            db.commit()
    except sqlite3.IntegrityError as e:
        print(e)
    else:
        break
