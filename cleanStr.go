package goreload

func CleanStr(byt []byte) []string {
	slice := []string{}
	str := ""
	for _, v := range byt {
		if v != ' ' {
			str += string(v)
		} else {
			slice = append(slice, str)
			str = ""
		}
	}
	slice = append(slice, str)
	return slice
}
