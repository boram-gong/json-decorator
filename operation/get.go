package operation

import (
	"github.com/boram-gong/json-decorator/common"
	"github.com/boram-gong/json-decorator/rule"
)

func GetJsonValue(keyList []rule.KeyStruct, jsonMap interface{}, del bool) interface{} {
	for i, key := range keyList {
		switch jsonMap.(type) {
		case map[string]interface{}:
			// 判断下一层的类型，主要是为了处理list
			switch common.Interface2Map(jsonMap)[key.Key].(type) {
			case []interface{}: // 当为list时
				switch len(key.Index) {
				case 0:
					// 不存在索引，意味着此key应该为最后一层
					if i == len(keyList)-1 && del {
						temp := common.Interface2Map(jsonMap)[key.Key]
						delete(common.Interface2Map(jsonMap), key.Key)
						return temp
					} else {
						jsonMap = common.Interface2Map(jsonMap)[key.Key]
						break
					}
				case 1:
					if len(common.Interface2Map(jsonMap)[key.Key].([]interface{})) == 0 {
						return nil
					}
					// 当只有一个索引的时候, 直接调key的值处理，不需要深层处理list
					if i == len(keyList)-1 { // 当为最后一个的时候可以直接返回
						index := key.Index[0]
						if index == -1 {
							index = len(common.Interface2Map(jsonMap)[key.Key].([]interface{})) - 1
						}
						if del {
							result := common.Interface2Map(jsonMap)[key.Key].([]interface{})[index]
							var temp []interface{}
							for n, v := range common.Interface2Map(jsonMap)[key.Key].([]interface{}) {
								if n == index {
									continue
								} else {
									temp = append(temp, v)
								}
							}
							common.Interface2Map(jsonMap)[key.Key] = temp
							return result
						} else {
							return common.Interface2Map(jsonMap)[key.Key].([]interface{})[index]
						}
					} else {
						// 非最后一个的时候我们进入新的循环
						return GetJsonValue(keyList[i+1:], common.Interface2Map(jsonMap)[key.Key].([]interface{})[key.Index[0]], del)
					}
				default:
					// 当有多个索引的时候, 深度处理list
					source := common.Interface2Map(jsonMap)[key.Key]
					return getDeepSliceValue(&source, key.Index, del, keyList[i+1:])
				}
			default:
				if i == len(keyList)-1 && del {
					temp := common.Interface2Map(jsonMap)[key.Key]
					delete(common.Interface2Map(jsonMap), key.Key)
					return temp
				} else {
					jsonMap = common.Interface2Map(jsonMap)[key.Key]
				}
			}
		default: // 当不为map类型时，就得终止循环
			if i == len(keyList)-1 {
				return jsonMap
			} else {
				return nil
			}
		}
	}
	return jsonMap
}

func getDeepSliceValue(source *interface{}, allIndex []int, del bool, nextKey []rule.KeyStruct) interface{} {
	for i, n := range allIndex {
		switch (*source).(type) {
		case []interface{}: // 必须为此类型
			// 处理索引
			if n == -1 {
				if len((*source).([]interface{})) > 0 {
					n = len((*source).([]interface{})) - 1
				} else {
					n = 0
				}
			}
			// 判断索引长度
			if len((*source).([]interface{})) > n { // 当前查询的值是在list里面的
				if i == len(allIndex)-1 { // 最后一层
					if len(nextKey) > 0 { // 最后一层的后面，要有新的kv， 继续嵌套查询
						return GetJsonValue(nextKey, (*source).([]interface{})[n], del)
					} else {
						result := (*source).([]interface{})[n]
						if del {
							var temp []interface{}
							for index, v := range (*source).([]interface{}) {
								if index == n {
									continue
								}
								temp = append(temp, v)
							}
							(*source) = temp
						}
						return result
					}
				} else {
					source = &(*source).([]interface{})[n]
				}
			} else {
				// 越界
				return nil
			}
		}
	}
	return nil
}
