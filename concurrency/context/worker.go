package main

import (
  "fmt"
  "sync"
  "time"

  "golang.org/x/net/context"
)

// global variable available throughout package
var wg sync.WaitGroup

func work(ctx context.Context) error {
  defer wg.Done()
  for i := 0; i < 1000; i++ {
    // select is used when dealing with channels
    // so you can select each time if done is signaled
    select {
      // go about your business as usual
      case <-time.After(2*time.Second):
        fmt.Printf("Work work work work work... %v\n", i)
      case <-ctx.Done():
        // received the signal to cancel in this channel
        fmt.Printf("Canceled the context %v\n", i)
        return ctx.Err()
    }
  }
  return nil
}

func main() {
  ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
  defer cancel()

  fmt.Println("Howdy, I'm about to do some work...")

  wg.Add(1)
  // passing in context as argument, so if we reach the 4 second timeout
  // ctx.Done() will be kicked off
  // and that select statment will be carried out
  go work(ctx)
  wg.Wait()

  fmt.Printf("Finished. Time to go home.")
}