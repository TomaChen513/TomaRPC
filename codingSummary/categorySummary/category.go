package categorysummary

import (
	"sort"
	"strconv"
)

// 1.算法：模拟 枚举 排序 贪心 双指针 构造 分治 哈希 高精度 dfs bfs 最短路 二分 动态规划 位运算 记忆化搜索 递归 前缀和 差分 拓扑排序 快速幂
// 2.数据结构:数组 链表 栈 队列 堆(优先队列) 树 图论 字符串 并查集 单调栈
// 3.数学：基础数学 数论 计算几何 位运算 博弈 概率论 组合数学  线性代数

// 01 背包
func bagProblem(weight []int, limit int) int {
	// dp[i][j]代表从0~i背包中取，限制为j的情况
	// dp[i][j]=max(dp[i-1][j],dp[i-1][j-weight[i]]+weight[i])
	// 初始化：dp[0][0]=0
	size := len(weight)
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, limit+1)
	}
	for i := 0; i <= limit; i++ {
		if i >= weight[0] {
			dp[0][i] = weight[0]
		}
	}
	for i := 1; i < size; i++ {
		for j := 0; j <= limit; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+weight[i])
			}

		}
	}
	return dp[size-1][limit]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 哈希
// https://www.programmercarl.com/0015.%E4%B8%89%E6%95%B0%E4%B9%8B%E5%92%8C.html#%E5%93%88%E5%B8%8C%E8%A7%A3%E6%B3%95
// 三数之和

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return res
}

// gerddy
func monotoneIncreasingDigits(N int) int {
	s := strconv.Itoa(N)
	ss := []byte(s)
	n := len(ss)
	if n <= 1 {
		return N
	}
	for i := n - 1; i > 0; i-- {
		if ss[i-1] > ss[i] {
			ss[i-1] -= 1
			for j := i; j < n; j++ {
				ss[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}
