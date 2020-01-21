package cart

// New 创建购物车问题
func New(itemsP []int, pLimit int) *Cart {
	n := len(itemsP)
	if n == 0 || pLimit == 0 {
		panic("error params")
	}
	pCap := pLimit * 2
	// 初始化状态数组
	status := make([][]bool, n)
	for i := 0; i < n; i++ {
		status[i] = make([]bool, pCap+1)
	}
	return &Cart{itemsP, pLimit, pCap, status}
}

// Cart  购物车问题
// 购物车中，求满足满减券商品的最小总额
type Cart struct {
	// 购物车中商品的金额
	itemsP []int
	// 满减券限制
	pLimit int
	// 最大金额
	pCap int
	// 状态数组 [层数][金额]
	status [][]bool
}

// MinPrice 动态规划求解
func (c *Cart) MinPrice() int {
	n := len(c.itemsP)
	// 赋值 第一层状态数组
	c.status[0][0] = true
	c.status[0][c.itemsP[0]] = true
	// 求解 其他层状态数组
	for i := 1; i < n; i++ {
		for j := 0; j <= c.pCap; j++ {
			// 上一层存在此金额
			if c.status[i-1][j] {
				c.status[i][j] = true
				if j+c.itemsP[i] <= c.pCap {
					c.status[i][j+c.itemsP[i]] = true
				}
			}
		}
	}
	// 求解 最小金额
	for j := c.pLimit; j <= c.pCap; j++ {
		if c.status[n-1][j] {
			return j
		}
	}
	return 0
}

// MinPriceItems 满足条件的商品
func (c *Cart) MinPriceItems() []int {
	res := []int{}
	j := c.MinPrice()
	if j == 0 {
		return res
	}
	for i := len(c.status) - 1; i > 0; i-- {
		if j-c.itemsP[i] >= 0 && c.status[i-1][j-c.itemsP[i]] {
			res = append(res, i)
			j = j - c.itemsP[i]
		}
	}
	if j > 0 {
		res = append(res, 0)
	}
	return res
}
