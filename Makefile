build:
	@go build

run: build
	@./hangman

clean:
	@rm ./hangman