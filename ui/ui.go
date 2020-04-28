package ui

import (
	"fmt"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

func UiPrintf() (n int) {
	n = 0
	fmt.Println("Hello World!")
	return
}

func UiTest() (n int) {
	n = 0
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())

	logrus.Info("succeeded")
	logrus.Warn("not correct")
	logrus.Error("something error")
	logrus.Fatal("panic")
	return
}
