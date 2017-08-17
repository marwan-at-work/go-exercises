# Interfaces

Interfaces in Go let you define behavior as types. This makes packages pluggable. 

Create a function that implements the `io.Writer` method that will do two things: 

1. It will write any incoming bytes to a logs.txt file on disk. 

2. If the `-v` flag is present, then it should also write the incoming bytes to stdout. 

When you finish writing your package, import the `chatty` package and call `chatty.Chat` passing the io.Writer implementation you just craeted.