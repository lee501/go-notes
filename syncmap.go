package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
)

var test1 = sync.Map{}
var del = sync.Map{}

type rule struct {
	Key   string
	Value string
}

var str11 = "[{\"Cron\":\"0 * * * *\",\"Dimension\":\"ip\",\"Expire\": 0,\"Params\":[\"2\",\"2\",\"2\",\"2\"],\"Risk\":\"ip_behavior\",\"RiskKey\":\"IP\",\"RiskType\":\"Crawler\",\"RuleDesc\":\"IP行为异常(24小时): 最近一天通过的URL TOP10序列化,以中位数以上IP为样本，如果其URL序列与全局序列相交\\u003e0.8。\",\"RuleName\":\"我自己的测试\",\"Sql\":\" SELECT toUnixTimestamp(max(ts)) AS _ts, ip, count() AS count, countIf(url_id = 1) AS u1, countIf(url_id = 2) AS u2, countIf(url_id = 3) AS u3, countIf(url_id = 4) AS u4, countIf(url) AS u5, countIf(url_id = 6) AS u6, countIf(url_id = 7) AS u7, countIf(url_id = 8) AS u8, countIf(url_id = 9) AS u9, countIf(url_id = 10) AS u10, array(u1, u2, u3, u4, u5, u6, u7, u8, u9, u10) AS l_vec, arraySum(l_vec) AS l_vec_sum, arrayMap(x -\\u003e(x / l_vec_sum), l_vec) AS l_feat, ( SELECT topK(url) FROM access_airline WHERE result='by_airline' and ts\\u003enow()-86400) AS g_top, topK(url) AS l_top, ( SELECT groupArray(g_count) FROM ( SELECT url, count() AS g_count, ( SELECT count() FROM access_airline WHERE ts\\u003enow()-86400 ) AS total FROM access_airline WHERE result='by_airline' and ts\\u003enow()-86400 GROUP BY url ORDER BY g_count DESC LIMIT 10 ) ) AS g_vec, arraySum(g_vec) AS g_vec_sum, arrayMap(x -\\u003e (x / g_vec_sum), g_vec) AS g_feat, sqrt(arraySum(arrayMap((x, y) -\\u003e ((x - y) * (x - y)), g_feat, l_feat))) AS euc, length(arrayIntersect(l_top, g_top)) AS int_cnt, 9 - ((((((((((u1 = 0) OR (u1 \\u003c u2)) + ((u2 = 0) OR (u2 \\u003c u3))) + ((u3 = 0) OR (u3 \\u003c u4))) + ((u4 = 0) OR (u4 \\u003c u5))) + ((u5 = 0) OR (u5 \\u003c u6))) + ((u6 = 0) OR (u6 \\u003c u7))) + ((u7 = 0) OR (u7 \\u003c u8))) + ((u8 = 0) OR (u8 \\u003c u9))) + ((u9 = 0) OR (u9 \\u003c u10))) AS ooo, ( SELECT quantiles(0.25, 0.5, 0.75)(count) FROM ( SELECT count() AS count FROM access_airline WHERE ts\\u003enow()-86400 GROUP BY ip ) ) AS Q, (%s) AND (count \\u003e Q[3]+3*(Q[3]-Q[1])) AND (length(g_top) = 10) AS _result FROM ( SELECT ts, ip, url FROM access_airline WHERE ts\\u003enow()-86400 ) AS access ANY LEFT JOIN ( SELECT *, rowNumberInAllBlocks() + 1 AS url_id FROM ( SELECT url, count() AS g_count FROM access_airline WHERE result='by_airline' and ts\\u003enow()-86400 GROUP BY url ORDER BY g_count DESC LIMIT 10 ) AS top_url ) AS top_url_id USING (url) GROUP BY ip \",\"Table\":\"access_airline_view\"}]"

type Rule struct {
	//Rule define
	RuleName   string //rule name
	RuleDesc   string
	RuleType   string
	RuleCustom string //custom rule
	Cron       string //cron
	Alert      int
	Status     int

	//SQL
	Insert    string        //insert
	Sql       string        //custom sql
	With      string        //with
	Dimension string        //group by
	Metric    string        //metric passed to measure
	Measure   string        //agg
	Params    []interface{} //default params
	TimeKey   string        //time key
	TimeRange string        //time range
	Filter    string
	Table     string //table name
	//Risk define
	Risk      string
	RiskKey   string
	RiskType  string //scenario
	RiskLevel string

	Expire int // default: 3600
	//Runtime
	Debug  string
	TaskId cron.EntryID

	//OutRiskMode
	RiskSink string
}

type Rules []Rule

var tests = "[{\"Cron\":\"0 * * * *\",\"Dimension\":\"ip\",\"Expire\": 0,\"Params\":[\"2\",\"2\",\"2\",\"2\"],\"Risk\":\"ip_behavior\",\"RiskKey\":\"IP\",\"RiskType\":\"Crawler\",\"RuleDesc\":\"IP行为异常(24小时): 最近一天通过的URL TOP10序列化,以中位数以上IP为样本，如果其URL序列与全局序列相交\\u003e0.8。\",\"RuleName\":\"我自己的测试1111\", \"Sql\": \"SELECT toUnixTimestamp(max(ts)) AS _ts, ip, count() AS count, countIf(url_id = 1) AS u1, countIf(url_id = 2) AS u2, countIf(url_id = 3) AS u3, countIf(url_id = 4) AS u4, countIf(url) AS u5, countIf(url_id = 6) AS u6, countIf(url_id = 7) AS u7, countIf(url_id = 8) AS u8, countIf(url_id = 9) AS u9, countIf(url_id = 10) AS u10, array(u1, u2, u3, u4, u5, u6, u7, u8, u9, u10) AS l_vec, arraySum(l_vec) AS l_vec_sum, arrayMap(x -\\u003e(x / l_vec_sum), l_vec) AS l_feat, ( SELECT topK(url) FROM access_airline WHERE result='by_airline' and ts\\u003enow()-86400) AS g_top, topK(url) AS l_top, ( SELECT groupArray(g_count) FROM ( SELECT url, count() AS g_count, ( SELECT count() FROM access_airline WHERE ts\\u003enow()-86400 ) AS total FROM access_airline WHERE result='by_airline' and ts\\u003enow()-86400 GROUP BY url ORDER BY g_count DESC LIMIT 10 ) ) AS g_vec, arraySum(g_vec) AS g_vec_sum, arrayMap(x -\\u003e (x / g_vec_sum), g_vec) AS g_feat, sqrt(arraySum(arrayMap((x, y) -\\u003e ((x - y) * (x - y)), g_feat, l_feat))) AS euc, length(arrayIntersect(l_top, g_top)) AS int_cnt, 9 - ((((((((((u1 = 0) OR (u1 \\u003c u2)) + ((u2 = 0) OR (u2 \\u003c u3))) + ((u3 = 0) OR (u3 \\u003c u4))) + ((u4 = 0) OR (u4 \\u003c u5))) + ((u5 = 0) OR (u5 \\u003c u6))) + ((u6 = 0) OR (u6 \\u003c u7))) + ((u7 = 0) OR (u7 \\u003c u8))) + ((u8 = 0) OR (u8 \\u003c u9))) + ((u9 = 0) OR (u9 \\u003c u10))) AS ooo, ( SELECT quantiles(0.25, 0.5, 0.75)(count) FROM ( SELECT count() AS count FROM access_airline WHERE ts\\u003enow()-86400 GROUP BY ip ) ) AS Q, (%s) AND (count \\u003e Q[3]+3*(Q[3]-Q[1])) AND (length(g_top) = 10) AS _result FROM ( SELECT ts, ip, url FROM access_airline WHERE ts\\u003enow()-86400 ) AS access ANY LEFT JOIN ( SELECT *, rowNumberInAllBlocks() + 1 AS url_id FROM ( SELECT url, count() AS g_count FROM access_airline WHERE result='by_airline' and ts\\u003enow()-86400 GROUP BY url ORDER BY g_count DESC LIMIT 10 ) AS top_url ) AS top_url_id USING (url) GROUP BY ip\",\"Table\":\"access_airline_view\"}, {\"Cron\":\"0 * * * *\",\"Dimension\":\"ip\",\"Expire\": 0,\"Params\":[\"2\",\"2\",\"2\",\"2\"],\"Risk\":\"ip_behavior\",\"RiskKey\":\"IP\",\"RiskType\":\"Crawler\",\"RuleDesc\":\"IP行为异常(24小时): 最近一天通过的URL TOP10序列化,以中位数以上IP为样本，如果其URL序列与全局序列相交\\u003e0.8。\",\"RuleName\":\"我自己的测试\", \"Sql\": \"SELECT toUnixTimestamp(max(ts)) AS _ts, ip, count() AS count, countIf(url_id = 1) AS u1, countIf(url_id = 2) AS u2, countIf(url_id = 3) AS u3, countIf(url_id = 4) AS u4, countIf(url) AS u5, countIf(url_id = 6) AS u6, countIf(url_id = 7) AS u7, countIf(url_id = 8) AS u8, countIf(url_id = 9) AS u9, countIf(url_id = 10) AS u10, array(u1, u2, u3, u4, u5, u6, u7, u8, u9, u10) AS l_vec, arraySum(l_vec) AS l_vec_sum, arrayMap(x -\\u003e(x / l_vec_sum), l_vec) AS l_feat, ( SELECT topK(url) FROM access_airline WHERE result='by_airline' and ts\\u003enow()-86400) AS g_top, topK(url) AS l_top, ( SELECT groupArray(g_count) FROM ( SELECT url, count() AS g_count, ( SELECT count() FROM access_airline WHERE ts\\u003enow()-86400 ) AS total FROM access_airline WHERE result='by_airline' and ts\\u003enow()-86400 GROUP BY url ORDER BY g_count DESC LIMIT 10 ) ) AS g_vec, arraySum(g_vec) AS g_vec_sum, arrayMap(x -\\u003e (x / g_vec_sum), g_vec) AS g_feat, sqrt(arraySum(arrayMap((x, y) -\\u003e ((x - y) * (x - y)), g_feat, l_feat))) AS euc, length(arrayIntersect(l_top, g_top)) AS int_cnt, 9 - ((((((((((u1 = 0) OR (u1 \\u003c u2)) + ((u2 = 0) OR (u2 \\u003c u3))) + ((u3 = 0) OR (u3 \\u003c u4))) + ((u4 = 0) OR (u4 \\u003c u5))) + ((u5 = 0) OR (u5 \\u003c u6))) + ((u6 = 0) OR (u6 \\u003c u7))) + ((u7 = 0) OR (u7 \\u003c u8))) + ((u8 = 0) OR (u8 \\u003c u9))) + ((u9 = 0) OR (u9 \\u003c u10))) AS ooo, ( SELECT quantiles(0.25, 0.5, 0.75)(count) FROM ( SELECT count() AS count FROM access_airline WHERE ts\\u003enow()-86400 GROUP BY ip ) ) AS Q, (%s) AND (count \\u003e Q[3]+3*(Q[3]-Q[1])) AND (length(g_top) = 10) AS _result FROM ( SELECT ts, ip, url FROM access_airline WHERE ts\\u003enow()-86400 ) AS access ANY LEFT JOIN ( SELECT *, rowNumberInAllBlocks() + 1 AS url_id FROM ( SELECT url, count() AS g_count FROM access_airline WHERE result='by_airline' and ts\\u003enow()-86400 GROUP BY url ORDER BY g_count DESC LIMIT 10 ) AS top_url ) AS top_url_id USING (url) GROUP BY ip\",\"Table\":\"access_airline_view\"}]"

func main() {
	var obj Rules
	fmt.Println(tests)
	err := json.Unmarshal([]byte(tests), &obj)
	if err != nil {
		fmt.Println("=========", err)
	}
}
