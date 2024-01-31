run:
	@go run golang_notes.go list
compile:
	go build -o build/notes golang_notes.go
install:
	go install

