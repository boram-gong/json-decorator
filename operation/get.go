package operation

import (
	dbt "github.com/boram-gong/db_tool"
	"github.com/boram-gong/json-decorator/common"
	"github.com/boram-gong/json-decorator/rule"
)

func GetJsonValue(keyList []rule.KeyStruct, jsonMap interface{}, del bool) interface{} {
	for i, key := range keyList {
		switch jsonMap.(type) {
		case map[string]interface{}:
			// 判断下一层的类型，主要是为了处理list
			switch dbt.Interface2Map(jsonMap)[key.Key].(type) {
			case []interface{}: // 当为list时
				switch len(key.Index) {
				case 0:
					// 不存在索引，意味着此key应该为最后一层
					if i == len(keyList)-1 && del {
						temp := dbt.Interface2Map(jsonMap)[key.Key]
						delete(dbt.Interface2Map(jsonMap), key.Key)
						return temp
					} else {
						jsonMap = dbt.Interface2Map(jsonMap)[key.Key]
						break
					}
				case 1: // 当只有一个索引的时候, 直接处理map的key值，不需要深层处理list
					if len(dbt.Interface2Map(jsonMap)[key.Key].([]interface{})) == 0 {
						// 越界取值取的结果为nil
						return nil
					}
					index := key.Index[0]
					if index == -1 {
						index = len(dbt.Interface2Map(jsonMap)[key.Key].([]interface{})) - 1
					}
					if i == len(keyList)-1 {
						// 当为最后一个的时，取到结果直接返回
						if del {
							result := dbt.Interface2Map(jsonMap)[key.Key].([]interface{})[index]
							dbt.Interface2Map(jsonMap)[key.Key] = common.MakeSlice(dbt.Interface2Map(jsonMap)[key.Key].([]interface{}), index)
							return result
						} else {
							return dbt.Interface2Map(jsonMap)[key.Key].([]interface{})[index]
						}
					} else {
						// 非最后一个的时候我们进入新的循环
						return GetJsonValue(keyList[i+1:], dbt.Interface2Map(jsonMap)[key.Key].([]interface{})[index], del)
					}
				default:
					// 当有多个索引的时候, 深度处理list
					source := dbt.Interface2Map(jsonMap)[key.Key]
					return getDeepSliceValue(&source, key.Index, del, keyList[i+1:])
				}
			default:
				if i == len(keyList)-1 && del {
					temp := dbt.Interface2Map(jsonMap)[key.Key]
					delete(dbt.Interface2Map(jsonMap), key.Key)
					return temp
				} else {
					jsonMap = dbt.Interface2Map(jsonMap)[key.Key]
				}
			}
		default:
			// 当不为map类型时，得终止循环
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
	for i, index := range allIndex {
		switch (*source).(type) {
		case []interface{}: // 必须为此类型
			if len((*source).([]interface{})) == 0 {
				return nil
			}
			// 处理索引
			if index == -1 {
				index = len((*source).([]interface{})) - 1
			}
			// 判断索引长度
			if len((*source).([]interface{})) > index { // 越界合法判断
				if i == len(allIndex)-1 {
					// 索引查询的最后一层
					if len(nextKey) > 0 {
						//  索引查询的最后一层后面还有key（list[1][0].key），继续嵌套查询
						return GetJsonValue(nextKey, (*source).([]interface{})[index], del)
					} else {
						result := (*source).([]interface{})[index]
						if del {
							(*source) = common.MakeSlice((*source).([]interface{}), index)
						}
						return result
					}
				} else {
					// 非最后一层
					source = &(*source).([]interface{})[index]
				}
			} else {
				// 越界
				return nil
			}
		}
	}
	return nil
}
