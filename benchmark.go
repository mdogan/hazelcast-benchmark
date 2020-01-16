package main

import (
    "context"
    "fmt"
    "math/rand"
    "sync"
    "time"

    "github.com/hazelcast/hazelcast-go-client/core"
    "github.com/mdogan/hdrhistogram"
)

type result struct {
    *hdrhistogram.Histogram
    duration time.Duration
}

func benchmark(m core.Map, wg *sync.WaitGroup, ctx context.Context, ch chan *result) {
    hist := hdrhistogram.New(1, int64(time.Second), 3)
    value := make([]byte, valueSize)
    rnd := rand.New(rand.NewSource(rand.Int63()))
    var err error
    result := result{Histogram: hist}

    defer func() {
        ch <- &result
        wg.Done()
    }()

    for i := 0; i < requests / clients; i++ {
        select {
        case <-ctx.Done():
            return
        default:

            k := rnd.Intn(setRatio + getRatio)
            start := time.Now()

            if k < getRatio {
                _, err = m.Get(i)
            } else {
                err = m.Set(i, value)
            }
            elapsed := time.Now().Sub(start)
            _ = hist.RecordValue(elapsed.Nanoseconds())
            result.duration += elapsed

            if err != nil {
                fmt.Println(err)
                return
            }
        }
    }
}
