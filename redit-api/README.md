# Reddit Topics

This exercise is based on the [Get Started with Go](https://www.youtube.com/watch?v=2KmHtgtEZ1s) tutorial by Andrew Gerrand.

I encourage you to go through the exercise first, try it yourself, then watch the video and compare.

Write a reddit client that will ping https://reddit.com/r/{topic-name}.json

where {topic-name} is any given topic from the command line. 

1. Create a struct type that resembles the JSON data coming from the request. You don't need to create every field, just the ones you create about: i.e. `children` and `data`.

2. Unmarshal the json in the struct, and print out every `topic` and `url` for that topic in the sub reddit. 

3. Add another field to the struct that represents the number of comments each topic has. 

4. Include the number of comments in the print out *only* if there is at least one comment. 

5. Implement the `fmt.Stringer` interface to do your printing. 

6. Extract the reddit logic into its own package, and have the main package import the reddit library. 