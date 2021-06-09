package main

import "time"

func getDuration(seconds int) time.Duration {
	return time.Duration(seconds) * time.Second
}
