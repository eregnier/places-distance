run:
	#https://github.com/gravityblast/fresh
	fresh

build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

docker-run: build
	docker-compose up --build

test:
	curl "http://localhost:8030/?pA=970,+boulevard+Albert+1er+59500+Douai&pB=rue+schweitzer+59139+wattignies"