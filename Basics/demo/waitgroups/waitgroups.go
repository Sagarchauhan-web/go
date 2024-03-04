package main

import (
	"fmt"
	"sync"
)

func main() {
    var wg sync.WaitGroup
    sum := 0

    for i := 0; i < 10; i++ {
        wg.Add(1);
        value := i

        go func() { 
            defer wg.Done()
            sum += value;
        }()
    }

    wg.Wait()
    fmt.Println("Sum:", sum)
}
