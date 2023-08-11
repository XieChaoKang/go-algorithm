package dynamic_programming

// 动态规划之计算爬楼梯方式
// 每一层楼梯的攀爬方式都是前面两节楼梯的攀爬方式数量之和：F(n) -> 爬一层楼梯的方式，F(n) = F(n-1) + F(n-2)

// n 代表楼梯数量
// 暴力递归去推算
func GetClimBingWay(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return GetClimBingWay(n-1) + GetClimBingWay(n-2)
}

var record = map[int]int{}

// 稍微优化一下 有一些楼梯会反复计算
func GetClimBingWays(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	v, ok := record[n]
	if ok {
		return v
	}
	v = GetClimBingWays(n-1) + GetClimBingWays(n-2)
	record[n] = v
	return v
}

// 自底向上来计算 每一层楼梯都是前面两节楼梯的攀爬方式数量之和 那我们从最底的楼梯开始算起
func GetClimBingWays2(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	// 已知道 1，2节楼梯的攀爬方式 那就从第三节开始往上推
	a := 1
	b := 2
	temp := 0
	for i := 3; i <= n; i++ {
		// 当前楼梯就等于前面两节楼梯的攀爬方式数量之和
		temp = a + b
		// 更新当前的前两节楼梯的攀爬方式数量
		a = b
		b = temp
	}
	return temp
}
