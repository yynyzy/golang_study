package main

import (
	"fmt"
	"time"
)

// time定时器
func main() {
	/*第一种
	 var timer = time.NewTimer(time.Second * 2)
		 fmt.Printf(" time.Now() = %v\n", time.Now())
		 t1 := <-timer.C //阻塞的，指定时间到了
	 	// var t1 time.Time = <-timer.C
		 fmt.Printf(" t1 = %v\n", t1)
	*/

	/*第二种
	fmt.Printf(" time1.Now() = %v\n", time.Now())
	var timer = time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Printf(" time2.Now() = %v\n", time.Now())
	*/

	/*第三种
	fmt.Printf(" time1.Now() = %v\n", time.Now())
	<-time.After(time.Second * 2)
	fmt.Printf(" time2.Now() = %v\n", time.Now())
	*/

	/*第四种
	timer := time.NewTimer(time.Second)
	go func() {
		<-timer.C
	}()
	s := timer.Stop()
	if s {
		fmt.Println("timer.Stop")
	}
	time.Sleep(time.Second * 3)
	fmt.Println("End")
	*/
	fmt.Println("before")
	timer := time.NewTimer(time.Second * 5) //设置为5s
	timer.Reset(time.Second)                //重置为 1s
	<-timer.C
	fmt.Println("end")
}
