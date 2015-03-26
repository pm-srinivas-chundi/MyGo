package mystr

func Strfill(ch string, n int) string {
	str := ""
	for i:=0; i<n; i++ {
		str = str + ch
	}

	return str
}
