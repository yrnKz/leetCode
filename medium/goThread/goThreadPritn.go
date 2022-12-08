package main

import "sync"

//启动两个线程, 一个输出 1,3,5,7…99, 另一个输出 2,4,6,8…100 最后 STDOUT 中按序输出 1,2,3,4,5…100？

var c1 chan struct{}
var c2 chan struct{}
var i = 1
var s = sync.WaitGroup{}

func main() {
	c1 = make(chan struct{})
	c2 = make(chan struct{})
	s.Add(1)
	go func() { c2 <- struct{}{} }()
	go f1()
	go f2()
	s.Wait()
}

func f1() {
	for j := 0; j < 50; j++ {
		select {
		case _ = <-c2:
			print("thread1->")
			println(i)
			i++
			c1 <- struct{}{}
		}
	}
}

func f2() {
	for j := 0; j < 50; j++ {
		select {
		case _ = <-c1:
			print("thread2->")
			println(i)
			i++
			if j == 49 {
				break
			}
			c2 <- struct{}{}
		}
	}
	s.Done()
}
