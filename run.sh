set -x

docker build -t joshiomkarj/in-memory-server .

docker run -p 8080:8080 joshiomkarj/in-memory-server