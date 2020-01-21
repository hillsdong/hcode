package bag01

// New 创建01背包问题
func New(itemsW []int, w int) *Bag01 {
	n := len(itemsW)
	if n == 0 || w == 0 {
		panic("error params")
	}
	// 初始化状态数组
	status := make([]bool, w+1)
	return &Bag01{itemsW, w, status}
}

// Bag01 01背包问题
// 固定物品范围的前提下，求背包可放物品最大重量
type Bag01 struct {
	// 可放物品的重量
	itemsW []int
	// 背包承重能力
	wCap int
	// 状态数组 [重量]
	status []bool
}

// MaxWeight 动态规划求解
func (b *Bag01) MaxWeight() int {
	n := len(b.itemsW)
	// 赋值 第一层状态数组
	b.status[0] = true
	b.status[b.itemsW[0]] = true
	// 求解 其他层状态数组
	for i := 1; i < n; i++ {
		for j := b.wCap - b.itemsW[i]; j >= 0; j-- {
			// 上一层存在此重量
			if b.status[j] {
				b.status[j+b.itemsW[i]] = true
			}
		}
	}
	// 求解 最大重量
	for j := b.wCap; j >= 0; j-- {
		if b.status[j] {
			return j
		}
	}
	return 0
}
