package main

import (
	"github.com/sirupsen/logrus"
)
var log1 = logrus.New()

func main() {
	log1.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 1,
		"size":   10,
	}).Info("A walrus appears")
}