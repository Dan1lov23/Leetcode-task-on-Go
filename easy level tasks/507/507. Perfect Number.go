package main

import "fmt"

func checkPerfectNumber(num int) bool {

	sum := 0

    if num % 2 == 0 {
        halfNum := num / 2
        for i := 1; i <= halfNum ; i++ {
		    if num % i == 0 {
			    sum += i
		    }
	    }
    } else {
        for i := 1; i < num - 1; i++ {
            if num % i == 0 {
                sum += i
            }
        }
    }

	if sum == num {
		return true
	}

	return false

}

func main() {
	fmt.Println(checkPerfectNumber(28))
}
