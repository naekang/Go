package main

func sum(a []int) int {
	var ans int = 0

	for i := 0; i < len(a); i++ {
		ans += a[i]
	}

	return ans
}
