package triangle

// New 类杨辉三角问题
func New(pointsV [][]int) *Tri {
	n := len(pointsV)
	if n == 0 {
		panic("error params")
	}
	// 初始化状态数组
	status := make([][]int, n)
	for i := 0; i < n; i++ {
		status[i] = make([]int, n)
	}
	return &Tri{pointsV, status}
}

// Tri 类杨辉三角问题
// 类杨辉三角中，求长度最短路径
type Tri struct {
	// 节点值
	pointsV [][]int
	// 状态数组 [层数][位置]值
	status [][]int
}

// MinValue 动态规划求解
func (t *Tri) MinValue() int {
	n := len(t.pointsV)
	// 赋值 第一层状态数组
	t.status[0][0] = t.pointsV[0][0]
	// 求解 其他层状态数组
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				t.status[i][j] = t.status[i-1][j] + t.pointsV[i][j]
			} else if j == i {
				t.status[i][j] = t.status[i-1][j-1] + t.pointsV[i][j]
			} else {
				lv := t.status[i-1][j-1]
				rv := t.status[i-1][j]
				t.status[i][j] = min(lv, rv) + t.pointsV[i][j]
			}
		}
	}
	// 求解 最小值
	minV := t.status[n-1][0]
	for j := 1; j < n; j++ {
		if t.status[n-1][j] < minV {
			minV = t.status[n-1][j]
		}
	}
	return minV
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
