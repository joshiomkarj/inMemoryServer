set -x

docker build -t joshiomkarj/in-memory-server .

docker rm in-memory-server

docker run -p 8080:8080 --name in-memory-server joshiomkarj/in-memory-server
