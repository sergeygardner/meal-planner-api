go get -d -v ./... && go install -v ./...
env $(cat .env | sed -E 's/^\#.*$//g' | xargs)
sh