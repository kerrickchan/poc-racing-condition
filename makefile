test: api call

api:
	cd server && go run cmd/app/main.go

call:
	cd client && node index.js
