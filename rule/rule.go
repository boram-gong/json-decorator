package rule

import (
	"github.com/boram-gong/json-decorator/common"
	json "github.com/json-iterator/go"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

var (
	AllRule atomic.Value
)

type KeyStruct struct {
	Key   string
	Index []int
}

type Rule struct {
	Key           string                 `json:"key"`
	Operation     string                 `json:"operation"`
	Content       interface{}            `json:"content"`
	Stat          int                    `json:"stat"`
	StartTime     string                 `json:"start_time"`
	EndTime       string                 `json:"end_time"`
	ERR           string                 `json:"-"`
	RealOperation string                 `json:"-"`
	KeyList       []KeyStruct            `json:"-"`
	AT            bool                   `json:"-"`
	ATList        map[string][]KeyStruct `json:"-"`
	Split         bool                   `json:"-"`
	Del           bool                   `json:"-"`
}

func (r *Rule) MakeRule() {
	// 初始化
	r.ATList = make(map[string][]KeyStruct)
	r.Split = false
	r.Del = false
	// 处理key的表达式，等到key的层级slice
	r.KeyList, _, r.ERR = makeRule(r.Key)
	if r.ERR != "" {
		return
	}
	// 处理operation的表达式，判断是否存在删除原始kv的操作
	switch r.Operation {
	case common.RENAME, common.DELETE:
		r.Del = true
		r.RealOperation = r.Operation
		r.AT = false
	case common.MOVE:
		r.Del = true
		r.RealOperation = r.Operation
	case "move-head":
		r.Del = true
		r.RealOperation = common.HEAD
	case "move-tail":
		r.Del = true
		r.RealOperation = common.TAIL
	case "append-head":
		r.RealOperation = common.HEAD
	case "append-tail":
		r.RealOperation = common.TAIL
	case common.INSERT, common.REPLACE:
		r.RealOperation = r.Operation
	default:
		r.ERR = "operation err"
		return
	}
	// 处理content的表达式，主要判断是否存在@以及是否市kv类型
	switch r.Content.(type) {
	case string:
		var (
			at = r.Content.(string)
			kv map[string]interface{}
		)
		if r.Operation == common.RENAME || r.Operation == common.DELETE {
			// 如果是rename的话，无视@
			al, split, errStat := makeRule(at)
			r.Split = split
			r.ATList[""] = al
			r.ERR = errStat
		} else {
			// 判断是否为kv类型
			if err := json.UnmarshalFromString(at, &kv); err != nil {
				// 非kv类型
				if at != "" && at[0] == '@' {
					r.AT = true
					key := at[1:]
					al, split, errStat := makeRule(key)
					r.Split = split
					r.ATList[""] = al
					r.ERR = errStat
				}
			} else {
				// kv 类型
				for k, v := range kv {
					switch v.(type) {
					case string:
						s := v.(string)
						if s[0] == '@' {
							r.AT = true
							key := s[1:]
							al, _, errStat := makeRule(key)
							// 因为每一个kv的v都可能含有@，所以这里以map的形式存储
							r.ATList[k] = al
							r.ERR = errStat
						}
					}
				}
				r.Content = kv
			}
		}
	case map[string]interface{}:
		// kv 类型 逻辑与上同理
		for k, v := range r.Content.(map[string]interface{}) {
			switch v.(type) {
			case string:
				at := v.(string)
				if at[0] == '@' {
					r.AT = true
					key := at[1:]
					al, _, errStat := makeRule(key)
					r.ATList[k] = al
					r.ERR = errStat
				}
			}
		}
	}
}

func makeRule(oKey string) (keys []KeyStruct, split bool, errStat string) {
	oKey = strings.ReplaceAll(oKey, "[...]", "[-2]")
	split = false
	var (
		l = strings.Split(oKey, ".")
	)
	for keySum, key := range l {
		if key == "" && oKey != "" {
			errStat = "rule err"
			return
		}
		// 获取中括号数量并生成正则表达式，检索中括号
		re, _ := regexp.Compile(makeBracketRegexp(key))
		result := re.FindStringSubmatch(key)
		var (
			indexList []int
			name      = key
		)
		if len(result) >= 3 {
			// 当正则匹配出有3个结果以上时，意味着此key存在[]
			name = result[1]
			for _, i := range result[2:] {
				index, err := strconv.Atoi(i)
				if err != nil {
					indexList = nil
					name = key
					break
				} else {
					if index == -2 {
						if keySum != len(l)-1 {
							// 不是最后一层的key
							errStat = "rule err: [...] mast last"
							return
						}
						split = true
						break
					}
					indexList = append(indexList, index)
				}
			}
			keys = append(keys, KeyStruct{
				Key:   name,
				Index: indexList,
			})
		} else {
			// 不存在[]
			keys = append(keys, KeyStruct{
				Key:   name,
				Index: indexList,
			})
		}
	}
	return
}

type AllRuleSafeMap struct {
	sync.RWMutex
	M map[string][]*Rule
}

func NewAllRuleSafeMap() *AllRuleSafeMap {
	return &AllRuleSafeMap{M: make(map[string][]*Rule)}
}

func (m *AllRuleSafeMap) Init() {
	m.Lock()
	defer m.Unlock()
	m.M = make(map[string][]*Rule)
}

func (m *AllRuleSafeMap) Store(ruleName string, r *Rule) {
	m.Lock()
	defer m.Unlock()
	r.MakeRule()
	v, ok := m.M[ruleName]
	if ok {
		v = append(v, r)
		m.M[ruleName] = v
	} else {
		m.M[ruleName] = []*Rule{r}
	}
}

func (m *AllRuleSafeMap) UnSafeStore(ruleName string, r *Rule) {
	r.MakeRule()
	v, ok := m.M[ruleName]
	if ok {
		v = append(v, r)
		m.M[ruleName] = v
	} else {
		m.M[ruleName] = []*Rule{r}
	}
}

func (m *AllRuleSafeMap) Load(k string) []*Rule {
	m.RLock()
	defer m.RUnlock()
	return m.M[k]
}

func makeBracketRegexp(s string) string {
	var (
		sum   = 0
		stack []byte
	)
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			stack = []byte{s[i]}
		} else if s[i] == ']' {
			if len(stack) == 1 && stack[0] == '[' {
				sum += 1
				stack = []byte{}
			}
		}
	}
	rs := "(.*)\\[(-?\\d+)\\]"
	for i := 1; i < sum; i++ {
		rs += `\[(-?\d+)\]`
	}
	return rs
}
