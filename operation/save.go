package operation

import (
	"errors"
	dbt "github.com/boram-gong/db_tool"
	"github.com/boram-gong/json-decorator/common"
	"github.com/boram-gong/json-decorator/rule"
)

func SaveJsonMap(keyList []rule.KeyStruct, jsonMap interface{}, op string, split bool, rightValue interface{}) error {
	if len(keyList) == 0 || len(keyList) == 1 && keyList[0].Key == "" {
		// 表示直接在顶级操作
		if op == common.INSERT && dbt.Interface2Map(rightValue) != nil && dbt.Interface2Map(jsonMap) != nil {
			// 表示直接在顶级插入一个键值对
			for k, v := range dbt.Interface2Map(rightValue) {
				jsonMap.(map[string]interface{})[k] = v
			}
			return nil
		}
	}
	for i, key := range keyList { // 遍历各个级别的key
		switch (jsonMap).(type) {
		case map[string]interface{}:
			if i == len(keyList)-1 {
				// 最后级别的key表示要存储
				data, ok := jsonMap.(map[string]interface{})[key.Key]
				if ok {
					switch data.(type) {
					case map[string]interface{}:
						switch op {
						case common.REPLACE, common.RENAME:
							jsonMap.(map[string]interface{})[key.Key] = rightValue
						case common.INSERT, common.MOVE:
							kv := dbt.Interface2Map(rightValue)
							if kv == nil {
								return errors.New("insert/move content type err, must kv")
							} else {
								for k, v := range kv {
									jsonMap.(map[string]interface{})[key.Key].(map[string]interface{})[k] = v
								}
							}
						}
						return nil
					case []interface{}:
						if len(key.Index) == 0 {
							switch op {
							case common.REPLACE:
								jsonMap.(map[string]interface{})[key.Key] = rightValue
							case common.TAIL:
								if split {
									data = append(data.([]interface{}), dbt.Interface2Slice(rightValue)...)
								} else {
									data = append(data.([]interface{}), rightValue)
								}
								jsonMap.(map[string]interface{})[key.Key] = data
							case common.HEAD:
								temp := data.([]interface{})
								if split {
									data = append(dbt.Interface2Slice(rightValue), temp...)
								} else {
									data = []interface{}{rightValue}
									data = append(data.([]interface{}), temp...)
								}
								jsonMap.(map[string]interface{})[key.Key] = data
							}
							return nil
						} else {
							if len(key.Index) == 1 && op == common.REPLACE {
								if key.Index[0] < len(data.([]interface{})) {
									jsonMap.(map[string]interface{})[key.Key].([]interface{})[key.Index[0]] = rightValue
								} else {
									return errors.New(key.Key + " index out")
								}
							} else {
								source := dbt.Interface2Map(jsonMap)[key.Key]
								saveDeepSliceValue(&source, rightValue, key.Index, nil, op, split)
							}
							return nil
						}
					case string:
						switch op {
						case common.HEAD:
							jsonMap.(map[string]interface{})[key.Key] = dbt.Interface2String(rightValue) + dbt.Interface2String(data)
						case common.TAIL:
							jsonMap.(map[string]interface{})[key.Key] = dbt.Interface2String(data) + dbt.Interface2String(rightValue)
						case common.REPLACE, common.RENAME:
							jsonMap.(map[string]interface{})[key.Key] = rightValue
						}
						return nil
					default:
						jsonMap.(map[string]interface{})[key.Key] = rightValue
					}
				} else {
					switch len(key.Index) {
					case 0:
						switch op {
						case common.INSERT, common.MOVE, common.RENAME:
							jsonMap.(map[string]interface{})[key.Key] = rightValue
						case common.TAIL, common.HEAD:
							if split {
								jsonMap.(map[string]interface{})[key.Key] = rightValue
							} else {
								jsonMap.(map[string]interface{})[key.Key] = []interface{}{rightValue}
							}
						}
						return nil
					case 1:
						switch op {
						case common.TAIL, common.HEAD:
							if split {
								jsonMap.(map[string]interface{})[key.Key] = rightValue
							} else {
								jsonMap.(map[string]interface{})[key.Key] = []interface{}{rightValue}
							}
						}
						return nil
					default:
						return errors.New(key.Key + " index out")
					}
				}
			} else { // 非最后一级的key，就要get出来，进入下次嵌套
				if _, ok := (jsonMap).(map[string]interface{})[key.Key]; ok {
					// 如果存在
					switch len(key.Index) {
					case 0:
						(jsonMap) = (jsonMap).(map[string]interface{})[key.Key]
					case 1:
						index := key.Index[0]
						if index < 0 {
							index = len(dbt.Interface2Map(jsonMap)[key.Key].([]interface{})) - 1
							if index < 0 {
								return errors.New(key.Key + " index out")
							}
						}
						return SaveJsonMap(
							keyList[i+1:],
							dbt.Interface2Map(jsonMap)[key.Key].([]interface{})[index],
							op, split, rightValue)

					default:
						source := dbt.Interface2Map(jsonMap)[key.Key]
						if saveDeepSliceValue(&source, rightValue, key.Index, keyList[i+1:], op, split) {
							return nil
						} else {
							return errors.New(key.Key + " save fail")
						}
					}
				} else {
					return errors.New(key.Key + " not exist")
				}

			}
		}
	}
	return nil
}

func saveDeepSliceValue(source *interface{}, value interface{}, allIndex []int, nextKey []rule.KeyStruct, op string, split bool) bool {
	for i, n := range allIndex {
		switch (*source).(type) {
		case []interface{}:
			if i == len(allIndex)-1 {
				if len(nextKey) > 0 { // 最后一层的后面，要有新的kv， 继续嵌套存储
					return SaveJsonMap(nextKey, (*source).([]interface{})[n], op, split, value) == nil
				} else {
					if n == -1 {
						if len((*source).([]interface{}))-1 >= 0 {
							n = len((*source).([]interface{})) - 1
						} else {
							return false
						}
					}
					switch (*source).([]interface{})[n].(type) {
					case []interface{}:
						if op == common.TAIL {
							if split {
								(*source).([]interface{})[n] = append(
									(*source).([]interface{})[n].([]interface{}),
									dbt.Interface2Slice(value)...,
								)
							} else {
								(*source).([]interface{})[n] = append(
									(*source).([]interface{})[n].([]interface{}),
									value,
								)
							}
						} else if op == common.HEAD {
							temp := (*source).([]interface{})[n].([]interface{})
							if split {
								(*source).([]interface{})[n] = append(
									dbt.Interface2Slice(value),
									temp...,
								)
							} else {
								(*source).([]interface{})[n] = []interface{}{value}
								(*source).([]interface{})[n] = append(
									(*source).([]interface{})[n].([]interface{}),
									temp...,
								)
							}
						} else if op == common.REPLACE {
							(*source).([]interface{})[n] = value
						}
					case map[string]interface{}:
						return SaveJsonMap(nil, (*source).([]interface{})[n], op, split, value) == nil
					case string:
						if op == common.TAIL {
							(*source).([]interface{})[n] = (*source).([]interface{})[n].(string) + dbt.Interface2String(value)
						} else if op == common.HEAD {
							(*source).([]interface{})[n] = dbt.Interface2String(value) + (*source).([]interface{})[n].(string)
						} else if op == common.REPLACE {
							(*source).([]interface{})[n] = value
						}
					}
				}
			} else {
				if n == -1 {
					if len((*source).([]interface{})) > 0 {
						n = len((*source).([]interface{})) - 1
					} else {
						n = 0
					}
				}
				if len((*source).([]interface{})) > n {
					source = &(*source).([]interface{})[n]
				}
			}
		}
	}
	return false
}
