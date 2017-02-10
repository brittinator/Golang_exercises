# Using go's context functionality when dealing with concurrency

This work is based heavily off of [this blog post](http://dahernan.github.io/2015/02/04/context-and-cancellation-of-goroutines/)  
Also a good StackOverflow answer to ['The term "Context" in programming'](http://stackoverflow.com/questions/6145091/the-term-context-in-programming)  

This exercise is mainly to understand the `context` functionality that can be used in concurrent programs.

**context** in this program allows slow requests to automatically get cancelled if they take
longer than 4 seconds.


## Usage

Start the server first inside a new terminal window that you can view with
    `go run server/server.go`
Then start the worker with
    `go run worker/worker.go`


Watch the worker program use the context to cancel outstanding requests to the flakey server.
