package main

import (
	"flag"
	"log"
	"time"
)

func memAllocateGb(num int) {

	// sleep 一会儿先让pod启动，不然会 起不来，演示不了oom
	log.Printf("go-mem-allocate sleep 30 second")
	time.Sleep(30 * time.Second)
	// arr代表100MB内存大小的arr，因为一个int64是8byte，128就是1024
	arr2 := make([][128 * 1024 * 1024]int64, num)
	for i := 0; i < num; i++ {
		arr2[i] = [128 * 1024 * 1024]int64{}
	}

	log.Printf("go-mem-allocate end  with gb_num GB:%v", num)
	time.Sleep(5 * time.Hour)

}

func main() {

	var gbNum int
	flag.IntVar(&gbNum, "gb_num", 1, "num of mem gb to allocate")
	flag.Parse()
	log.Printf("go-mem-allocate start with gb_num GB:%v", gbNum)
	memAllocateGb(gbNum)

}
