run:
	@go run golang_notes.go list
build:
	go build -o ./build/notes golang_notes.go
install:
	cp -f ./build/notes /usr/local/bin

