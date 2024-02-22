package main

func main() {
	defer print("urazmetov")
	defer print("lev ")
	defer println(printing())
	println("hello ")
	a := 2
	func(a int) {
		a *= a
		println(a)
	}(a)
	var slides []int
	slice := []int{1, 2, 3}
	slides = slices(slice)
	for idx := range slides {
		println(slides[idx])
	}
	return
}

func printing() string {
	defer println("hello brother")
	return "aboba"
}

func slices(arr []int) []int {
	arr[1] = 22
	return arr
}
