package types

func Ptr(s string) *string {
	return &s
}

func IPtr(v interface{}) interface{} {
	return &v
}
