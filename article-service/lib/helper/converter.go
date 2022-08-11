package helper

func ToFloat32(in *float32) float32 {
	if in != nil {
		return *in
	}
	return 0
}

func ToString(in *string) string {
	if in != nil {
		return *in
	}
	return ""
}
