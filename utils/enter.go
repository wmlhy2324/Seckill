package utils

func List(list []string, key string) (ok bool) {
	for _, v := range list {
		if v == key {
			return true
		}
	}
	return false

}
