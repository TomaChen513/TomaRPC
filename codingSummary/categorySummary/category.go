package categorysummary

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
			if j<weight[i]{
				dp[i][j] = dp[i-1][j]
			}else{
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
