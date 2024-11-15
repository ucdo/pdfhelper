package main

import (
	"Octopus/PdfHelper/utils"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestShot(t *testing.T) {
	a := []string{"png", "pdf"}
	a = append(a, append(a, append(a, a...)...)...)
	urlStr := "https://www.example.com"
	for i, s := range a {
		wg.Add(1)
		go func() {
			utils.Do(urlStr, s, i)
			defer wg.Done()
		}()

	}
	wg.Wait()
}
