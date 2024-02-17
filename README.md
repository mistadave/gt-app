# Applitaction Game Tutor

Main is a Go Application, which does the api and also serves the ui.

## Client folder

UI for the api.

## Dev

Install air for hot reload

```bash
go install github.com/cosmtrek/air@latest
# may need to register alias in your bash, zsh
alias air='$(go env GOPATH)/bin/air'
air -v

# run go app with air
air main.go
```