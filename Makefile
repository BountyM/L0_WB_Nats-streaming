.PHONY: all

build:
	sudo docker network create nats-net

	sudo docker-compose -f nats&nats-streaming.yaml up -d
