package utils

func SOrX(s, x string) string {
	if s == "" {
		return x
	}
	return s
}

func SOrR(s string, run func() string) string {
	if s == "" {
		return run()
	}
	return s
}
