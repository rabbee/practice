package easy

// isRectangleOverlap
// 直角坐标系下判断两个平行于x轴y轴的矩形相交，等价于判断两个矩形在x轴和y轴上的投影线段是否相交
// 在一条直线上的两条线段是否相交，只需判断线段1的右端小于线段2的左端，或者线段2的右端小于线段1的左端
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return isLineOverLap([]int{rec1[0], rec1[2]}, []int{rec2[0], rec2[2]}) &&
		isLineOverLap([]int{rec1[1], rec1[3]}, []int{rec2[1], rec2[3]})
}

func isLineOverLap(l1, l2 []int) bool {
	return !(l1[1] <= l2[0] || l2[1] <= l1[0])
}
