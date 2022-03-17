wait-for "localhost:3306" -- "$@"

# Watch your .go files and invoke go build if the files changed.
CompileDaemon --build="go build -o main"  --command=./main