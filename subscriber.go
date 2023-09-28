package main

import "fmt"

type Subscriber struct {
	name string
}

func (s *Subscriber) handleEvent(vacancies []string) {
	fmt.Printf("Hello %s we changed our vacancy list to %s\n\n", s.name, vacancies)
}
