package _758_Minimum_Changes_To_Make_Alternating_Binary_String

func minOperations(s string) int {
	count := 0
	for k, v := range s {
		if k%2 == 1 && v == '1' {
			count++
		} else if k%2 == 0 && v == '0' {
			count++
		}
	}
	if len(s)-count > count {
		return count
	}
	return len(s) - count
}
