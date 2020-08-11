package Week_08

func isPerfectSquare(num int) bool {
	//x^2 = num
	if num < 2 {
		return true
	}

	var low, high = 0, num/2

	for low <= high {
		var medium = (high-low)/2 + low
		var tmp = medium*medium
		if tmp == num {
			return true
		} else if tmp < num {
			low = medium+1
		} else {
			high = medium-1
		}
	}
	return false

}
