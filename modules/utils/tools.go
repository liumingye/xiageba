package utils

import (
	"fmt"
	"math"
	"reflect"
)

// convert any numeric value to int64
func ToInt64(value interface{}) (int64, error) {
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u := val.Uint()
		if u > math.MaxInt64 {
			return 0, fmt.Errorf("uint value is too large to fit in int64")
		}
		return int64(u), nil
	default:
		return 0, fmt.Errorf("ToInt64 needs a numeric type, got `%T`", value)
	}
}

// 格式化秒数为人类可读的字符串
func FormatSecondsAgo(seconds float64) string {
	timeUnits := []struct {
		Threshold float64
		Divisor   float64
		Format    string
	}{
		{60, 1, "%.0f秒前"},
		{3600, 60, "%.0f分钟前"},
		{86400, 3600, "%.0f小时前"},
		{math.MaxFloat64, 86400, "%.0f天前"}, // 使用math.MaxFloat64作为最后一个阈值
	}

	for _, unit := range timeUnits {
		if seconds < unit.Threshold {
			return fmt.Sprintf(unit.Format, seconds/unit.Divisor)
		}
	}
	return "" // 或根据逻辑返回一个默认值
}
