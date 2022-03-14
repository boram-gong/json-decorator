package body

type JsonReq struct {
	Name      string                 `json:"name"`
	Data      interface{}            `json:"data"`
	JsonMap   map[string]interface{} `json:"-"`
	JsonSlice []interface{}          `json:"-"`
}

type SaveRuleReq struct {
	Id        int        `json:"id"`
	Name      string     `json:"rule_name"`
	Rules     []UserRule `json:"rules"`
	Stat      int        `json:"stat"`
	StartTime string     `json:"start_time"`
	EndTime   string     `json:"end_time"`
}

type UserRule struct {
	Key       string      `json:"key"`
	Operation string      `json:"operation"`
	Content   interface{} `json:"content"`
}
