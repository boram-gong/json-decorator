package rule

import (
	json "github.com/json-iterator/go"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type KeyStruct struct {
	Key   string
	Index []int
}

type Rule struct {
	Key           string                 `json:"key"`
	Operation     string                 `json:"operation"`
	Content       interface{}            `json:"content"`
	ERR           string                 `json:"-"`
	RealOperation string                 `json:"-"`
	KeyList       []KeyStruct            `json:"-"`
	AT            bool                   `json:"-"`
	ATList        map[string][]KeyStruct `json:"-"`
	Split         bool                   `json:"-"`
	Stat          int                    `json:"-"`
	StartTime     string                 `json:"-"`
	EndTime       string                 `json:"-"`
	Del           bool                   `json:"-"`
}

func (r *Rule) MakeRule() {
	r.ATList = make(map[string][]KeyStruct)
	r.Split = false
	r.Del = false
	r.KeyList, _, r.ERR = makeRule(r.Key)
	if r.ERR != "" {
		return
	}
	switch r.Operation {
	case "rename", "delete":
		r.Del = true
		r.RealOperation = r.Operation
		r.AT = false
	case "move":
		r.Del = true
		r.RealOperation = r.Operation
	case "move-head":
		r.Del = true
		r.RealOperation = "head"
	case "move-tail":
		r.Del = true
		r.RealOperation = "tail"
	case "append-head":
		r.RealOperation = "head"
	case "append-tail":
		r.RealOperation = "tail"
	case "insert", "replace":
		r.RealOperation = r.Operation
	default:
		r.ERR = "operation err"
		return
	}

	switch r.Content.(type) {
	case string:
		var (
			at = r.Content.(string)
			kv map[string]interface{}
		)
		if r.Operation == "rename" {
			al, split, errStat := makeRule(at)
			r.Split = split
			r.ATList[""] = al
			r.ERR = errStat
		} else {
			if err := json.UnmarshalFromString(at, &kv); err != nil {
				if at != "" && at[0] == '@' {
					r.AT = true
					key := at[1:]
					al, split, errStat := makeRule(key)
					r.Split = split
					r.ATList[""] = al
					r.ERR = errStat
				}
			} else {
				for k, v := range kv {
					switch v.(type) {
					case string:
						s := v.(string)
						if s[0] == '@' {
							r.AT = true
							key := s[1:]
							al, _, errStat := makeRule(key)
							r.ATList[k] = al
							r.ERR = errStat
						}
					}
				}
				r.Content = kv
			}
		}
	case map[string]interface{}:
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

func makeRule(key string) (keys []KeyStruct, split bool, errStat string) {
	key = strings.ReplaceAll(key, "[...]", "[-2]")
	split = false
	var (
		l = strings.Split(key, ".")
	)
	for _, s := range l {
		if s == "" && key != "" {
			errStat = "rule err"
			return
		}
		sum := bracketSum(s)
		rs := "(.*)\\[(-?\\d+)\\]"
		for i := 1; i < sum; i++ {
			rs += `\[(-?\d+)\]`
		}
		re, _ := regexp.Compile(rs)
		result := re.FindStringSubmatch(s)
		var (
			indexList []int
			name      = s
		)
		if len(result) >= 3 {
			name = result[1]
			for _, i := range result[2:] {
				index, err := strconv.Atoi(i)
				if err != nil {
					indexList = nil
					name = s
					break
				} else {
					if index == -2 {
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

func (m *AllRuleSafeMap) Load(k string) []*Rule {
	m.RLock()
	defer m.RUnlock()
	return m.M[k]
}

var (
	AllRule = &AllRuleSafeMap{M: make(map[string][]*Rule)}
)

func bracketSum(s string) int {
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
	return sum
}
