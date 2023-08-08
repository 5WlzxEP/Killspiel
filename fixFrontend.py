import os

BASE = "Backend/frontend_build"

for (dirpath, dirnames, filenames) in os.walk(BASE):
    if "_app" in dirpath:
        continue
    # print(dirpath, dirnames, filenames)
    for file in filenames:
        file: str
        if not file.endswith(".html") or file == "index.html" or file == "r.html":
            continue

        f: str = file.removesuffix(".html")
        if f not in dirnames:
            os.mkdir(f"{dirpath}/{f}")
        if "index.html" in os.listdir(f"{dirpath}/{f}"):
            raise IndexError(f"Unexpected index.html in {dirpath}/{f}")
        os.rename(f"{dirpath}/{file}", f"{dirpath}/{f}/index.html")
