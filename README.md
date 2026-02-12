# Prerequisites

**Requirements**  
- Go  
- Docker  

## Install Goose
```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Make sure `$GOPATH/bin` is in your path in order to run goose cmd:
```shell
export PATH="$PATH:$(go env GOPATH)/bin"
```
or put it in `.zshrc`, `.bashrc` or alike.

# Run
## 1. Start up mysql db
```shell
docker compose up -d
```
## 2. Apply db migrations
```shell 
-- not done --
```

