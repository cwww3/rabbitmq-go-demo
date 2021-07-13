package rabbit_demo

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func BodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func SeverityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}

func SeverityFrom2(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
	} else {
		s = os.Args[1]
		fmt.Println(s)
	}
	return s
}
