package rule

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	//InitDB()
	//data := &body.SaveRuleReq{
	//	Id:   11,
	//	Name: "demo2",
	//	Rules: []body.UserRule{{
	//		Key:       "demo.key",
	//		Operation: "rename",
	//		Content:   "demo.newKey2",
	//	}},
	//	Stat:      1,
	//	StartTime: "",
	//	EndTime:   "",
	//}
	//err := SaveRule(data)
	//if err != nil {
	//	fmt.Println(err)
	//	t.Fail()
	//}
	//fmt.Println(GetAllRule())
	fmt.Println(time.Now().Format("15:04:05"))
}
