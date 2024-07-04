package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, " started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, " finished job ", j)
		results <- j * 2
	}
}

func main() {
	println(time.Now().String())
	const numOfJobs = 5
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= numOfJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numOfJobs; a++ {
		<-results
	}
	println(time.Now().String())
}
