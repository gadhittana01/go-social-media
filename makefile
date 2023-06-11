# run db migartion
db-init:
	go run migration/main.go init
	go run migration/main.go up

db-up:
	go run migration/main.go up

db-down:
	go run migration/main.go down

db-reset:
	go run migration/main.go reset

# run http server
run-http-server-local:
	go build -o "./cmd/social-media-http/social-media-http" ./cmd/social-media-http && ./cmd/social-media-http/social-media-http