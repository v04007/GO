package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

type obj struct {
	name       chan string
	department chan string
}

type info struct {
	name       string
	department string
}

func main() {
	var eg errgroup.Group

	newobj := obj{
		name:       make(chan string),
		department: make(chan string),
	}

	i := new(info)
	eg.Go(func() error {
		for range time.Tick(time.Millisecond * 15) {
			select {
			case name := <-newobj.name:
				fmt.Println(name)
				i.name = name
			case department := <-newobj.department:
				fmt.Println(department)
				i.department = department
			default:
				fmt.Println(i)
				if len(i.department) != 0 && len(i.name) != 0 {
					return nil
				}
				continue
			}
		}
		return nil
	})

	eg.Go(func() error {
		time.Sleep(time.Second * 2)
		newobj.name <- "小明"
		return nil
	})

	eg.Go(func() error {
		time.Sleep(time.Second * 4)
		newobj.department <- "工程部"
		return nil
	})

	if err := eg.Wait(); err != nil {
		fmt.Printf("some error occur: %s\n", err.Error())
	}
	fmt.Print("over", i)

}
