package helper

import (
	"strconv"
	"strings"
)

func GetString(in *string) string {
	res := ""

	if in != nil {
		res = *in
	}

	return res
}

func GetInt32(in *int32) int32 {
	var res int32
	res = 0

	if in != nil {
		res = *in
	}

	return res
}

func GetInt(in *int) int {
	res := 0

	if in != nil {
		res = *in
	}

	return res
}

func ToString(in *string) string {
	if in != nil {
		return *in
	}
	return ""
}

func ToInt32(in *int32) int32 {
	if in != nil {
		return *in
	}
	return 0
}

func ToFloat64(in *float64) float64 {
	if in != nil {
		return *in
	}
	return 0
}

func ToBool(in *bool) bool {
	if in != nil {
		return *in
	}
	return false
}

func MakeArray(in *string, sep string) []string {
	if ToString(in) != "" {
		return strings.Split(*in, sep)
	}
	return nil
}

func MakeArrayInt32(in *string, sep string) []int32 {
	if ToString(in) != "" {
		ret := []int32{}
		splitted := strings.Split(*in, sep)
		for _, i := range splitted {
			ii, _ := strconv.Atoi(i)
			ret = append(ret, int32(ii))
		}
		return ret
	}
	return nil
}

func ToFloat32(in *float32) float32 {
	if in != nil {
		return *in
	}
	return 0
}
