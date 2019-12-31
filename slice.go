package go_utils

// 切割一个字符串数组为多个数组
func SplitSliceString(strList []string, limit int) [][]string {
	index := 0
	length := len(strList)
	var uidList [][]string
	for length-index > limit {
		l := limit
		if index+l > length {
			l = length - index
		}
		uidList = append(uidList, strList[index:index+l])
		index += l
	}
	if length-index > 0 {
		uidList = append(uidList, strList[index:length])
	}
	return uidList
}
