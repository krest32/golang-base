package main


func main() {
	test15_3()
}


func test15_3()  {

	defer func() {
		if err := recover(); err != nil{
			println(err.(string))
		}
	}()
	panic("panic error")

}