package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "dog",
	}).Info("一条舔狗出现了。")
}