package rule

import (
	"fmt"
	dbt "github.com/boram-gong/db_tool"
	dbt_pg "github.com/boram-gong/db_tool/pg"
	"github.com/boram-gong/json-decorator/common"
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

func Query(querySql string) (result []map[string]interface{}) {
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

func GetRuleCfg() {
	AllRule.Init()
	result := Query("SELECT * from cfg ORDER BY rule_name,weight")
	for _, m := range result {
		r := &Rule{
			Id:            common.Interface2Int(m["id"]),
			Key:           common.Interface2String(m["key"]),
			Operation:     common.Interface2String(m["operation"]),
			Content:       m["content"],
			RealOperation: "",
			KeyList:       nil,
			AT:            false,
			ATList:        nil,
			Split:         false,
			Del:           false,
		}
		AllRule.Store(common.Interface2String(m["rule_name"]), r)
	}
}
