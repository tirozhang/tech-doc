package _170_Two_Sum_III

type TwoSum struct {
	hash map[int]int
}

func Constructor() TwoSum {
	return TwoSum{
		hash: make(map[int]int),
	}
}

func (this *TwoSum) Add(number int) {
	if _, ok := this.hash[number]; ok {
		this.hash[number]++
	} else {
		this.hash[number] = 1
	}
}

func (this *TwoSum) Find(value int) bool {
	if len(this.hash) < 2 {
		return false
	}
	for number, times := range this.hash {
		if number == value-number && times > 1 {
			return true
		}
		if _, ok := this.hash[value-number]; ok && number != value-number {
			return true
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
