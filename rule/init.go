package rule

import (
	"errors"
	"fmt"
	dbt "github.com/boram-gong/db_tool"
	"github.com/boram-gong/json-decorator/common"
	json "github.com/json-iterator/go"
	"strings"
)

const (
	JsonRuleTable = "rule"
)

func query(querySql string, db dbt.DB) (result []map[string]interface{}, err error) {
	rows, err := db.QueryX(querySql)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		m := map[string]interface{}{}
		if e := rows.MapScan(m); e != nil {
			continue
		}
		for k, v := range m {
			if strings.Contains(dbt.Interface2String(v), "}") {
				temp := make(map[string]interface{})
				if json.UnmarshalFromString(dbt.Interface2String(v), &temp) == nil {
					m[k] = temp
				}
			}
		}
		result = append(result, m)

	}
	return result, nil
}

func ReAllRule(cli dbt.DB) []common.SaveRuleReq {
	ruleMap := NewAllRuleSafeMap()
	var respData []common.SaveRuleReq
	result, _ := query(dbt.SelectFieldsSql(JsonRuleTable, "*", ""), cli)
	for _, m := range result {
		var data []common.UserRule
		if err := json.UnmarshalFromString(dbt.Interface2String(m["rule"]), &data); err != nil {
			continue
		}
		respData = append(respData, common.SaveRuleReq{
			Id:        dbt.Interface2Int(m["id"]),
			Name:      dbt.Interface2String(m["rule_name"]),
			Rules:     data,
			Stat:      dbt.Interface2Int(m["stat"]),
			StartTime: dbt.Interface2String(m["start_time"]),
			EndTime:   dbt.Interface2String(m["end_time"]),
		})
		for _, d := range data {
			r := &Rule{
				Key:       d.Key,
				Operation: d.Operation,
				Content:   d.Content,
				Stat:      dbt.Interface2Int(m["stat"]),
				StartTime: dbt.Interface2String(m["start_time"]),
				EndTime:   dbt.Interface2String(m["end_time"]),
			}
			ruleName := fmt.Sprintf("%v", m["rule_name"])
			ruleMap.UnSafeStore(ruleName, r)
		}
	}
	AllRule.Store(ruleMap)
	return respData
}

func GetAllRule(cli dbt.DB) []common.SaveRuleReq {
	var respData []common.SaveRuleReq
	result, _ := query(dbt.SelectFieldsSql(JsonRuleTable, "*", ""), cli)
	for _, m := range result {
		var data []common.UserRule
		if err := json.UnmarshalFromString(dbt.Interface2String(m["rule"]), &data); err != nil {
			continue
		}
		respData = append(respData, common.SaveRuleReq{
			Id:        dbt.Interface2Int(m["id"]),
			Name:      dbt.Interface2String(m["rule_name"]),
			Rules:     data,
			Stat:      dbt.Interface2Int(m["stat"]),
			StartTime: dbt.Interface2String(m["start_time"]),
			EndTime:   dbt.Interface2String(m["end_time"]),
		})
	}
	return respData
}

func GetOneRule(id int, cli dbt.DB) common.SaveRuleReq {
	var respData common.SaveRuleReq
	result, _ := query(dbt.SelectFieldsSql(JsonRuleTable, "*", fmt.Sprintf("id=%v", id)), cli)
	for _, m := range result {
		var data []common.UserRule
		if err := json.UnmarshalFromString(dbt.Interface2String(m["rule"]), &data); err != nil {
			continue
		}
		respData = common.SaveRuleReq{
			Id:        dbt.Interface2Int(m["id"]),
			Name:      dbt.Interface2String(m["rule_name"]),
			Rules:     data,
			Stat:      dbt.Interface2Int(m["stat"]),
			StartTime: dbt.Interface2String(m["start_time"]),
			EndTime:   dbt.Interface2String(m["end_time"]),
		}
		break
	}
	return respData
}
func SaveRule(data *common.SaveRuleReq, cli dbt.DB) error {
	saveSql := ""
	rules, err := json.Marshal(data.Rules)
	if err != nil {
		return err
	}
	if data.Id != 0 {
		change := []string{
			fmt.Sprintf("rule_name='%v'", data.Name),
			fmt.Sprintf("rule='%s'", string(rules)),
			fmt.Sprintf("stat=%v", data.Stat),
			fmt.Sprintf("start_time='%v'", data.StartTime),
			fmt.Sprintf("end_time='%v'", data.EndTime),
		}
		saveSql = dbt.UpdateSql(JsonRuleTable, fmt.Sprintf("id=%v", data.Id), change)
	} else {
		fields := []string{"rule_name", "rule", "stat", "start_time", "end_time"}
		values := fmt.Sprintf("'%v','%v',%v,'%v','%v'",
			data.Name,
			string(rules),
			data.Stat,
			data.StartTime,
			data.EndTime,
		)
		saveSql = dbt.InsertSql(JsonRuleTable, fields, values)
	}
	_, err = cli.Exec(saveSql)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRule(id int, cli dbt.DB) error {
	if id == 0 {
		return errors.New("id is null")
	}
	_, err := cli.Exec(dbt.DeleteSql(JsonRuleTable, fmt.Sprintf("id=%v", id)))
	if err != nil {
		return err
	}
	return nil
}
