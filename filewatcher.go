package main

import (
	"fmt"
	"os"
	"time"
)

// WatchFile blocks until file is updated.
func WatchFile(path string) {
	s, err := os.Stat(path)
	if err != nil {
		panic("watched file does not exist")
	}
	for {
		time.Sleep(time.Second)
		s2, err := os.Stat(path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// file modified?
		if s2.Size() != s.Size() ||
			s2.ModTime() != s.ModTime() {
			return
		}
	}
}
