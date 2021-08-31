//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/19 18:02
// @Company cloud-ark.com
//
package main

//
// main
// 哲学家就餐问题
// 1 | 2 | 3 | 4 | 5 |
// 怎么让人能就餐，一个人必须要左手和右手都有叉子才能就行就餐
//
func main() {

}
/*
var state = make([]int, 5)

const HUNGRY = 1
const EATING = 2
const THINKLING = 3

func philosopher(i int) {
	for {
		think()
		take_forks(i)
		eat()
		put_forks(i)
	}
}

func think()  {
	p(mutex)          //进入临界区
	state[n] = THINKLING
	v(mutex) // 退出临界区
}

func eat()  {
	// TODO 吃东西
}

func take_forks(i int) {
	p(mutex)          //进入临界区
	state[i] = HUNGRY //表明饿了
	// 试图拿两把叉子
	test_take_left_right_forks(i)
	// 没有叉子拿需要阻塞
	v(mutex) // 退出临界区
	P(s[i])  //没有叉子便阻塞
}

func put_forks(i int) {
	p(mutex)                          //进入临界区
	state[i] = THINKLING              //进入思考状态
	test_take_left_right_forks(left)  //唤醒左邻居
	test_take_left_right_forks(right) // 唤醒右邻居
	v(mutex)                          // 退出临界区
}

func test_take_left_right_forks(i int) {
	if state[i] == HUNGRY && state[left] != EATING && state[right] != EATING { //自己处于饥饿状态
		state[i] = EATING
		V(s[i]) //通知第i个人可以吃饭了
	}
}
*/