package main

func main() {

	for n := 0; n < 100; n++ {
		if n%3 == 0 && n%5 == 0 {
			println("Pin Pan")
		} else if n%3 == 0 {
			println("Pin")
		} else if n%5 == 0 {
			println("Pan")
		}
	}

}
