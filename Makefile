.PHONY: all

build:
	sudo docker run -p 4223:4223 -p 8223:8223 nats-streaming -p 4223 -m 8223
run:
	go run cmd/main.go