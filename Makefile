build:
	docker-compose build $(srv)
start:
	docker-compose up -d $(srv)
