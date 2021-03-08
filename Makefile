build:
	go build -ldflags="-s -w" -o godiary .
	chmod +x ./godiary
