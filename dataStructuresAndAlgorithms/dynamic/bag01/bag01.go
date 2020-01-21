package bag01

// New 创建01背包问题
func New(itemsW []int, w int) *Bag01 {
	n := len(itemsW)
	if n == 0 || w == 0 {
		panic("error params")
	}
	// 初始化状态数组
	status := make([][]bool, n)
	for i := 0; i < n; i++ {
		status[i] = make([]bool, w+1)
	}
	return &Bag01{itemsW, w, status}
}

// Bag01 01背包问题
// 固定物品范围的前提下，求背包可放物品最大重量
type Bag01 struct {
	// 可放物品的重量
	itemsW []int
	// 背包承重能力
	wCap int
	// 状态数组 [层数][重量]
	status [][]bool
}

// MaxWeight 动态规划求解
func (b *Bag01) MaxWeight() int {
	n := len(b.itemsW)
	// 赋值 第一层状态数组
	b.status[0][0] = true
	b.status[0][b.itemsW[0]] = true
	// 求解 其他层状态数组
	for i := 1; i < n; i++ {
		for j := 0; j <= b.wCap; j++ {
			// 上一层存在此重量
			if b.status[i-1][j] {
				b.status[i][j] = true
				if j+b.itemsW[i] <= b.wCap {
					b.status[i][j+b.itemsW[i]] = true
				}
			}
		}
	}
	// 求解 最大重量
	for j := b.wCap; j >= 0; j-- {
		if b.status[n-1][j] {
			return j
		}
	}
	return 0
}

// MaxWeightItems 满足条件的物品
func (b *Bag01) MaxWeightItems() []int {
	res := []int{}
	j := b.MaxWeight()
	if j == 0 {
		return res
	}
	for i := len(b.status) - 1; i > 0; i-- {
		if j-b.itemsW[i] >= 0 && b.status[i-1][j-b.itemsW[i]] {
			res = append(res, i)
			j = j - b.itemsW[i]
		}
	}
	if j > 0 {
		res = append(res, 0)
	}
	return res
}
