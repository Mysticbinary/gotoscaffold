package gtsorder

// 找切片中的最大值 , 返回最大值和对应下标
func FindSliceMax(s []int) (int, int) {
	maxVal := s[0]
	maxValIndex := 0
	for i := 0; i < len(s); i++ {
		//从第二个元素开始循环比较，如果发现有更大的数，则交换
		if maxVal < s[i] {
			maxVal = s[i]
			maxValIndex = i
			return maxVal, maxValIndex
		}
	}
	return 0, 0
}

// 找切片最大和第二大的值，并返回
// []int{1, 2, 3, 3, 2, 5, 5} 如果最大值有重复，那么也将返回： 5，3
func FindMaxAndSecond(s []int) (int, int) {
	max := s[0]
	secondmax := 0
	for i := 0; i < len(s); i++ {
		if max < s[i] {
			max = s[i]
		}
	}
	for j := 0; j < len(s); j++ {
		if secondmax < s[j] {
			if s[j] < max {
				secondmax = s[j]
			}
		}
	}
	return max, secondmax
}

// int 元素去重，判断长度来决定用什么方式去重
func RemoveIntRepeat(slc []int) []int {
	if len(slc) < 1024 {
		// 切片长度小于1024的时候，循环来过滤
		return RemoveRepByLoop(slc)
	} else {
		// 大于的时候，通过map来过滤
		return RemoveRepByMap(slc)
	}
}

// 通过双重循环来过滤重复元素，属于时间换空间
func RemoveRepByLoop(slc []int) []int {
	result := []int{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// 通过map主键唯一的特性过滤重复元素，通过字典来过滤（空间换时间）
func RemoveRepByMap(slc []int) []int {
	result := []int{}
	tempMap := map[int]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// Slice string 去重复
func RemoveStrRepeat(oldS []string) []string {
	newS := []string{}
	for i := 0; i < len(oldS); i++ {
		repeat := false
		for j := i + 1; j < len(oldS); j++ {
			if oldS[i] == oldS[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newS = append(newS, oldS[i])
		}
	}
	return newS
}

// 排序
func SmallToLargeOrder(arr []int) []int {
	m := len(arr)
	for i := 0; i < m-1; i++ {
		positon := i
		for j := i + 1; j < m; j++ {
			if arr[j] < arr[positon] {
				positon = j
			}
		}
		arr[i], arr[positon] = arr[positon], arr[i]
	}
	return arr
}
