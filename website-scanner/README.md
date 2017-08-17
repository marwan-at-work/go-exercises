# Website Scanner

Given the `urls.txt` file, write a program that:

1. Reads each line in that file.
2. Makes a GET request to the url.
3. Outputs to terminal if given url has the text `hello` in it. 

## Further enhancements.

Make sure you do these one at a time and only after you finished the basic exercise.

1. Change the search text from `hello` to any given cmd line argument: `go run main.go welcome`
2. Make the http requests concurrent with Goroutines. 
3. Limit the Goroutines to 10 Goroutines at any given time, hint: worker pools.