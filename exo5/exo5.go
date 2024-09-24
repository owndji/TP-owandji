package exo5

func Ft_max_substring(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := string(s[0])
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(res); j++ {
			if s[i] == res[j] {
				i = len(s) - 1
			}
		}
		if i != len(s)-1 {
			res += string(s[i])
		}
	}

	return len(res)
}
