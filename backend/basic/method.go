package basic

import "fmt"

type Military struct {
	Title string
}

func (m Military) TitleJobs() {
	fmt.Println("Jobs name : ", m.Title)
}