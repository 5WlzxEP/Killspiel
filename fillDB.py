import random
import sqlite3
from faker import Faker


faker = Faker()
with sqlite3.connect("config/db.sqlite") as db:

    cursor = db.cursor()

    for i in range(100):
        g = random.randint(10, 150)
        cursor.execute(f"INSERT INTO Users VALUES ({random.randint(0, 1_000_000)}, '{faker.name()}', {g}, {random.randint(0, g)}, {random.randint(0, 256)})")

    db.commit()
