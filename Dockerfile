# Frontend

FROM node:latest as frontend

WORKDIR /app

COPY frontend/package.json .

RUN npm i

COPY frontend/*.config.* ./
COPY frontend/src ./src
COPY frontend/static ./static

ENV VITE_BACKEND_URL=""

RUN npm run build

FROM golang:latest as backend

WORKDIR /app

COPY Backend/ ./

RUN go mod download

COPY --from=frontend /app/build ./frontend_build/

ENV CGO_ENABLED=0

RUN go build -ldflags="-w -s" -o main main.go

FROM scratch

COPY --from=backend /app/main /main

EXPOSE 8088

ENTRYPOINT ["/main"]
