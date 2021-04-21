package main

import (
    "log"
    "time"
)

func main() {
    log.Println("Shipping Service Worker Started")
    ticker := time.NewTicker(10 * time.Second)
    for range ticker.C {
        log.Println("Checking for pending shipments...")
    }
}
