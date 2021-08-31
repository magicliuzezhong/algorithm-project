//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/24 09:04
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//defer_call()

	//pase_student()

	//testGo()

	//var teacher = Teacher{}
	//teacher.ShowA()

	//exceptionTest()

	//testCalc()

	//testArrPrint()

	var u = userAges{
		ages: make(map[string]int),
	}
	go func() {
		u.Add("aa", 1)
		fmt.Println("添加aa")
	}()
	go func() {
		u.Add("aa1", 2)
		fmt.Println("添加aa1")
	}()
	time.Sleep(time.Second * 3)
	var aa = u.Get("aa")
	fmt.Println(aa)

	//var aa = iterTest()
	//for val := range aa {
	//	fmt.Println(val)
	//}

	//var poe IPeople = &StudentPeople{}
	//var aa = poe.Speaking("bitch")
	//fmt.Println(aa)

	//var peo = live()
	//if peo == nil {
	//	fmt.Println("aa")
	//} else {
	//	fmt.Println("bb")
	//}
	//var a = peo.Speaking("bitch")
	//fmt.Println("a", a)

	//var b *StudentPeople
	//fmt.Println(b.Speaking("bitch"))
}

func defer_call() {
	defer func() {
		fmt.Println("调用1")
	}()
	defer func() {
		fmt.Println("调用2")
	}()
	defer func() {
		fmt.Println("调用3")
	}()
	panic("触发异常")
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	var m = make(map[string]*student)
	var stus = []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Println("a")
}

func testGo() {
	runtime.GOMAXPROCS(1)
	var wg = sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("show A")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func exceptionTest() {
	runtime.GOMAXPROCS(1)
	var int_chan = make(chan int, 1)
	var string_chan = make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println("value: ", value)
	case value := <-string_chan:
		panic(value)
	}
}

func testCalc() {
	var a = 1
	var b = 2
	defer calc("1", a, b)
	a = 0
	defer calc("2", a, b)
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b)
	return ret
}

func testArrPrint() {
	var s = make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

type userAges struct {
	ages map[string]int
	sync.Mutex
}

func (u *userAges) Add(name string, age int) {
	u.Lock()
	time.Sleep(time.Second)
	defer u.Unlock()
	u.ages[name] = age
}
func (u *userAges) Get(name string) int {
	if age, ok := u.ages[name]; ok {
		return age
	}
	return -1
}

func iterTest() <-chan int {
	var ch = make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

type IPeople interface {
	Speaking(string) string
}

type StudentPeople struct {
}

func (s *StudentPeople) Speaking(thinking string) (talk string) {
	if thinking == "bitch" {
		talk = "you are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func live() IPeople {
	var stu *StudentPeople
	return stu
}
