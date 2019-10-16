package main

import (
"fmt"
"github.com/oleiade/lane"
"sync"
)

func worker(item interface{}, wg *sync.WaitGroup) {
fmt.Println(item)
wg.Done()
}


func main() {

queue := lane.NewQueue()
queue.Enqueue("grumpyClient")
queue.Enqueue("happyClient")
queue.Enqueue("ecstaticClient")

var wg sync.WaitGroup

// Let's handle the clients asynchronously
for queue.Head() != nil {
item := queue.Dequeue()

wg.Add(1)
go worker(item, &wg)
}

// Wait until everything is printed
wg.Wait()
}