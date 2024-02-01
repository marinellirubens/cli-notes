run:
	@go run golang_notes.go list
compile:
	go build -o build/notes golang_notes.go
install:
	cp -f ./completions/fish/notes.fish /usr/share/fish/completions/notes.fish
	go install

