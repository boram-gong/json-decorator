package rule

import (
	"errors"
	"fmt"
	dbt "github.com/boram-gong/db_tool"
	dbt_pg "github.com/boram-gong/db_tool/pg"
	"github.com/boram-gong/json-decorator/common"
	"github.com/boram-gong/json-decorator/common/body"
	json "github.com/json-iterator/go"
)

var (
	PClient dbt.DB
	pCfg    = &dbt.CfgDB{
		Host:        "114.67.78.94",
		Port:        5432,
		User:        "postgres",
		Password:    "Wayz2022",
		Database:    "response_adapter",
		MaxIdleConn: 5,
		MaxOpenConn: 20,
	}
)

func InitDB() {
	PClient = dbt_pg.NewPgClient(pCfg)
}

func query(querySql string) (result []map[string]interface{}) {
	rows, err := PClient.QueryX(querySql)
	if err != nil {
		fmt.Println(querySql, err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		m := map[string]interface{}{}
		if err := rows.MapScan(m); err != nil {
			continue
		}
		result = append(result, m)

	}
	return result
}

func ReAllRule() []body.SaveRuleReq {
	AllRule.Init()
	var respData []body.SaveRuleReq
	result := query("SELECT * from rule")
	for _, m := range result {
		var data []body.UserRule
		if err := json.UnmarshalFromString(common.Interface2String(m["rule"]), &data); err != nil {
			continue
		}
		respData = append(respData, body.SaveRuleReq{
			Id:        common.Interface2Int(m["id"]),
			Name:      common.Interface2String(m["rule_name"]),
			Rules:     data,
			Stat:      common.Interface2Int(m["stat"]),
			StartTime: common.Interface2String(m["start_time"]),
			EndTime:   common.Interface2String(m["end_time"]),
		})
		for _, d := range data {
			r := &Rule{
				Key:           d.Key,
				Operation:     d.Operation,
				Content:       d.Content,
				RealOperation: "",
				KeyList:       nil,
				AT:            false,
				ATList:        nil,
				Split:         false,
				Del:           false,
				Stat:          common.Interface2Int(m["stat"]),
				StartTime:     common.Interface2String(m["start_time"]),
				EndTime:       common.Interface2String(m["end_time"]),
			}
			ruleName := fmt.Sprintf("%v", m["rule_name"])
			AllRule.Store(ruleName, r)
		}
	}
	return respData
}

func GetAllRule() []body.SaveRuleReq {
	var respData []body.SaveRuleReq
	result := query("SELECT * from rule order by id")
	for _, m := range result {
		var data []body.UserRule
		if err := json.UnmarshalFromString(common.Interface2String(m["rule"]), &data); err != nil {
			continue
		}
		respData = append(respData, body.SaveRuleReq{
			Id:        common.Interface2Int(m["id"]),
			Name:      common.Interface2String(m["rule_name"]),
			Rules:     data,
			Stat:      common.Interface2Int(m["stat"]),
			StartTime: common.Interface2String(m["start_time"]),
			EndTime:   common.Interface2String(m["end_time"]),
		})
	}
	return respData
}

func GetOneRule(id int) body.SaveRuleReq {
	var respData body.SaveRuleReq
	result := query(fmt.Sprintf("SELECT * from rule where id=%v", id))
	for _, m := range result {
		var data []body.UserRule
		if err := json.UnmarshalFromString(common.Interface2String(m["rule"]), &data); err != nil {
			continue
		}
		respData = body.SaveRuleReq{
			Id:        common.Interface2Int(m["id"]),
			Name:      common.Interface2String(m["rule_name"]),
			Rules:     data,
			Stat:      common.Interface2Int(m["stat"]),
			StartTime: common.Interface2String(m["start_time"]),
			EndTime:   common.Interface2String(m["end_time"]),
		}
		break
	}
	return respData
}
func SaveRule(data *body.SaveRuleReq) error {
	saveSql := ""
	rules, err := json.Marshal(data.Rules)
	if err != nil {
		return err
	}
	if data.Id != 0 {
		saveSql = fmt.Sprintf("update rule set rule_name='%v',rule='%s',stat=%v,start_time='%v',end_time='%v' where id=%v",
			data.Name,
			string(rules),
			data.Stat,
			data.StartTime,
			data.EndTime,
			data.Id,
		)
	} else {
		saveSql = fmt.Sprintf("insert into rule (rule_name,rule,stat,start_time,end_time) values ('%v','%v',%v,'%v','%v')",
			data.Name,
			string(rules),
			data.Stat,
			data.StartTime,
			data.EndTime,
		)
	}
	_, err = PClient.Exec(saveSql)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRule(id int) error {
	if id == 0 {
		return errors.New("id is null")
	}
	deleteSql := fmt.Sprintf("delete from rule where id=%v", id)
	_, err := PClient.Exec(deleteSql)
	if err != nil {
		return err
	}
	return nil
}
