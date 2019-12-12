package main

import (
	"fmt"
	"os"
	"time"
)

type Person struct {
	firstName string
	lastName string
	birthDay time.Time
}

type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	if p[i].birthDay.Unix() == p[j].birthDay.Unix() {
		if p[i].firstName == p[j].firstName {
			return p[i].lastName < p[j].lastName
		}
		return p[i].firstName < p[j].firstName
	}
	return p[i].birthDay.Unix() > p[j].birthDay.Unix()
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func GenerateDate(s string) time.Time {
	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return date
}
