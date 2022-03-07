package common

import (
	"fmt"
	"strconv"
)

func simpleCalculator(data1 interface{}, operation string, data2 interface{}) interface{} {
	switch operation {
	case "+":
		switch data1.(type) {
		case string:
			return Interface2String(data1) + Interface2String(data2)
		case int:
			return Interface2Int(data1) + Interface2Int(data2)
		case int64:
			return Interface2Int(data1) + Interface2Int(data2)
		case float64:
			return Interface2Float(data1) + Interface2Float(data2)
		default:
			return nil
		}
	case "-":
		switch data1.(type) {
		case int:
			return Interface2Int(data1) - Interface2Int(data2)
		case int64:
			return Interface2Int(data1) - Interface2Int(data2)
		case float64:
			return Interface2Float(data1) - Interface2Float(data2)
		default:
			return nil
		}
	case "*":
		switch data1.(type) {
		case int:
			return Interface2Int(data1) * Interface2Int(data2)
		case int64:
			return Interface2Int(data1) * Interface2Int(data2)
		case float64:
			return Interface2Float(data1) * Interface2Float(data2)
		default:
			return nil
		}
	case "/":
		if Interface2Int(data2) == 0 || Interface2Float(data2) == 0 {
			return nil
		}

		switch data1.(type) {
		case int:
			return Interface2Int(data1) / Interface2Int(data2)
		case int64:
			return Interface2Int(data1) / Interface2Int(data2)
		case float64:
			return Interface2Float(data1) / Interface2Float(data2)
		default:
			return nil
		}

	}
	return nil
}

func Interface2Map(data interface{}) map[string]interface{} {
	switch data.(type) {
	case map[string]interface{}:
		return data.(map[string]interface{})
	default:
		return nil
	}
}

func Interface2Slice(data interface{}) []interface{} {
	switch data.(type) {
	case []interface{}:
		return data.([]interface{})
	default:
		return nil
	}
}

func Interface2String(data interface{}) string {
	switch data.(type) {
	case string:
		return data.(string)
	default:
		return fmt.Sprintf("%v", data)
	}
}

func Interface2Int(data interface{}) int {
	switch data.(type) {
	case int64:
		return int(data.(int64))
	case int:
		return data.(int)
	case string:
		i, _ := strconv.Atoi(data.(string))
		return i
	case float64:
		return int(data.(float64))
	default:
		return 0
	}
}

func Interface2Float(data interface{}) float64 {
	switch data.(type) {
	case float64:
		return data.(float64)
	case int64:
		return float64(data.(int64))
	case int:
		return float64(data.(int))
	case string:
		i, _ := strconv.ParseFloat(data.(string), 64)
		return i
	default:
		return 0.0
	}
}
