package bag01

// New 创建01背包问题
func New(itemsW []int, itemsV []int, w int) *Bag01 {
	n := len(itemsW)
	vn := len(itemsV)
	if n == 0 || vn != n || w == 0 {
		panic("error params")
	}
	// 初始化状态数组
	status := make([]int, w+1)
	for i := 0; i <= w; i++ {
		status[i] = -1
	}

	return &Bag01{itemsW, itemsV, w, status}
}

// Bag01 01背包问题
// 固定物品范围的前提下，不超过背包承重能力，求背包可放物品最大价值
type Bag01 struct {
	// 可放物品的重量
	itemsW []int
	// 可放物品的价值
	itemsV []int
	// 背包承重能力
	wCap int
	// 状态数组 [重量]价值
	status []int
}

// MaxValue 动态规划求解
func (b *Bag01) MaxValue() int {
	n := len(b.itemsW)
	// 赋值 第一层状态数组
	b.status[0] = 0
	b.status[b.itemsW[0]] = b.itemsV[0]
	// 求解 其他层状态数组
	for i := 1; i < n; i++ {
		for j := b.wCap - b.itemsW[i]; j >= 0; j-- {
			// 上一层存在此重量
			if b.status[j] != -1 {
				if b.status[j+b.itemsW[i]] < b.status[j]+b.itemsV[i] {
					b.status[j+b.itemsW[i]] = b.status[j] + b.itemsV[i]
				}
			}
		}
	}
	// 求解 最大价值
	maxV := 0
	for j := b.wCap; j >= 0; j-- {
		if maxV < b.status[j] {
			maxV = b.status[j]
		}
	}
	return maxV
}
