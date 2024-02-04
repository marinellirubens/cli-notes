# bash -> /usr/share/bash-completion/completions

run:
	@go run golang_notes.go list
compile:
	go build -o build/notes golang_notes.go
install: compile
	cp -f ./completions/fish/notes.fish /usr/share/fish/completions/notes.fish
	cp -f ./completions/zsh/_notes /usr/share/zsh/site-functions/_notes
	cp -f ./build/notes /usr/bin/notes

