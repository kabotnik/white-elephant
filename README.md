# white-elephant
Give it a list of players and it will select the next player until all have had a turn.

## Building

On Windows
```
go build -o white-elephant.exe
```

Elsewhere
```
go build -o white-elephant
```

## Running
```
white-elephant play -p players
```

Players should be a file containing a list of names, each name separated by a newline
```
Bob
Arnold
Mary
Alice
James
```
