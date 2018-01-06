package main

/*
	接口更高层次的抽象
*/

func main() {

	// 实现IInstance接口的对象必须有Instance方法
	type IInstance interface {
		Instance() float64
	}

}
