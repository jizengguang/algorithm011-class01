package Week_04

func lemonadeChange(bills []int) bool {

	//时间复杂度O(N)
	//空间复杂度O(1)
	five, ten := 0, 0

	for _, v := range bills {
		if v == 5 {
			five++
		} else if v == 10 {
			if five > 1 {
				ten++
				five--
			} else {
				return false
			}
		} else {
			if five >= 1 && ten >= 1 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}

		}

	}
	return true

}

func main() {
	var bill = []int{5, 5, 10, 10, 20}
	lemonadeChange(bill)
}
