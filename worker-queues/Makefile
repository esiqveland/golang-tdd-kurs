
static:
	CGO_ENABLED=0 GOOS=linux go build -o main.linux

build:
	docker build -t tdufourd/queues .

run:
	docker run -p 8000:8000 tdufourd/queues

work:
	for i in {1..100}; do curl 192.168.99.100:8000/work -d name=$USER -d delay=$(expr $i % 11)s; done
