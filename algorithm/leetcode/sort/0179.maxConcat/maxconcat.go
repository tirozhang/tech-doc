package _179_maxConcat

import (
	"math"
	"strconv"
)

// 比较函数  定义比较规则
//func compare(a, b int) bool {
//	// case 1 : 两数相等
//	if a == b {
//		return true
//	}
//	// ab 和ba比较 323*100+32  32*1000+323
//	if a*getNumlen(b)+b > b*getNumlen(a)+a {
//		return true
//	}
//	return false
//}

// 采用排序，排序规则是为高位取大胜出 排序算法冒泡
func maxConcat(nums []int) string {

	// 步骤一：排序算法 可以选择任何一种 降序排列
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			// 判断交换条件
			if !compare(nums[i], nums[j]) {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	// 步骤二：数组拼接
	var result int = 0
	for _, v := range nums {
		result = result*(int(math.Pow(10, float64(getLength(v))))) + v
	}
	return strconv.Itoa(result)
}

// 比较函数  定义比较规则
// case 1: 两数相等，{123,123} 直接返回true
// case 1.1: 两数循环重复{1212,12}或{333,3} 直接返回true
// case 2: 两数前缀不同{123,14} 匹配到第二位2<4即可判断
// case 3: 两数前缀相同{12345,12} 匹配前缀后继续循环比较{345,12}
// case 3.1: 两数前缀相同且循环比较后长度比匹配,{12121,12}=>{1,12} 需要补齐长度继续比较{112121,12}
// 问题转化为 将两数长度补全为倍数关系，循环比较至最后一位结束即{33233,332}=>{332333,332} 确认比较次数为6
func compare(a, b int) bool {
	if a == b {
		return true
	}
	// step1 获取数字位数和需要比较次数
	al, bl := getLength(a), getLength(b)
	compareTimes := getCompareTimes(a, b)

	// step2: 初始化比较变量
	aLast, aDivisor := a, int(math.Pow(10, float64(al-1)))
	bLast, bDivisor := b, int(math.Pow(10, float64(bl-1)))

	// step3: 循环比较
	for compareTimes > 0 {
		aTop := aLast / aDivisor
		bTop := bLast / bDivisor
		if aTop != bTop {
			return aTop > bTop
		}
		aLast, bLast = aLast%aDivisor, bLast%bDivisor
		aDivisor, bDivisor = aDivisor/10, bDivisor/10

		// 当数字遍历完成后，继续循环
		if aDivisor == 0 {
			aLast, aDivisor = a, int(math.Pow(10, float64(al-1)))
		}
		if bDivisor == 0 {
			bLast, bDivisor = b, int(math.Pow(10, float64(bl-1)))
		}
		compareTimes--
	}

	return true
}

// 获取数字位数  987 -> 1000
func getLength(num int) int {
	i := 0
	for num > 0 {
		num /= 10
		i++
	}
	return i
}

func getCompareTimes(al, bl int) int {
	if al > bl {
		return al/bl + 1
	} else {
		return bl/al + 1
	}
}
