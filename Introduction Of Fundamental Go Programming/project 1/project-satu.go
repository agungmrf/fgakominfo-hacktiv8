package main

import (
	"fmt"
	"sync"
	"time"
)

func printData1(data interface{}, no int) {
	fmt.Println(data, " ", no)
}

func printData2(data interface{}, no int) {
	fmt.Println(data, " ", no)
}

func printDataConcurrently(data1 interface{}, data2 interface{}, printDataFunc func(interface{}, int)) {
	for i := 1; i <= 4; i++ {
		go printDataFunc(data1, i)
		go printDataFunc(data2, i)
	}
}

func printDataSynchronized(data interface{}, no int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	fmt.Println(data, " ", no)
	mutex.Unlock()
}

func printDataSynchronizedConcurrently(data1 interface{}, data2 interface{}, printDataFunc func(interface{}, int),
	mutex *sync.Mutex, wg *sync.WaitGroup) {

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go printDataSynchronized(data1, i, mutex, wg)
		go printDataSynchronized(data2, i, mutex, wg)
		wg.Wait()
	}
}

func main() {
	data1 := []interface{}{"bisa1", "bisa2", "bisa3"}
	data2 := []interface{}{"coba1", "coba2", "coba3"}

	var mutex = &sync.Mutex{}
	var wg = &sync.WaitGroup{}

	fmt.Println("---------------------------------------------")
	fmt.Println("Goroutine Menampilkan Secara Acak")
	fmt.Println("---------------------------------------------")
	printDataConcurrently(data1, data2, printData1)
	time.Sleep(3 * time.Second)

	fmt.Println("---------------------------------------------")
	fmt.Println("Goroutine Menampilkan Secara Rapih")
	fmt.Println("---------------------------------------------")
	printDataSynchronizedConcurrently(data1, data2, printData2, mutex, wg)
	time.Sleep(3 * time.Second)
}
