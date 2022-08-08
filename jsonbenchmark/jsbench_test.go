package jsonbenchmark

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/goccy/go-json"
)

type AgentService struct {
	ServiceName    string
	Version        string
	ServiceId      string
	Address        string
	Port           int
	Metadata       map[string]string
	ConnectTimeOut int
	ConnectType    string
	ReadTimeOut    int
	WriteTimeOut   int
	Protocol       string
	Balance        string
	Idcs           string
	Converter      string
	Retry          int
}

var obj = AgentService{
	ServiceName:    "kaleidoscope_api",
	Version:        "1517558949087295000_1298498081",
	ServiceId:      "kaleidoscope.com_v1.2",
	Address:        "127.0.0.1",
	Port:           80,
	Metadata:       map[string]string{},
	ConnectTimeOut: 1000,
	ConnectType:    "LONG",
	ReadTimeOut:    1000,
	WriteTimeOut:   1000,
	Protocol:       "HTTP",
	Balance:        "Random",
	Idcs:           "hu,hd,hn",
	Converter:      "json",
	Retry:          3,
}

func BenchmarkSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := sonic.Marshal(obj)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGoJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(obj)
		if err != nil {
			b.Error(err)
		}
	}
}

var obj1 = `{"req":{"Url":{"Scheme":"","Opaque":"","User":null,"Host":"","Path":"/ibe/flightSearch","RawPath":"","ForceQuery":false,"RawQuery":"","Fragment":""},"Method":"POST","HeaderMap":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Accept-Language":["en-US,en;q=0.5"],"B2c-Api-User-Token":["UM5u3alZTs3VYCRIPnkvmwtFQ6CkzfJr/JSUDmuOe/H8+q1hxU7IFQ=="],"Connection":["close"],"Content-Length":["218"],"Content-Type":["application/json;charset=UTF-8"],"Origin":["https://www.donghaiair.com"],"Referer":["https://www.donghaiair.com/html/booking-manage/choose-flight-two.html?flightType=1\u0026orgCode=SZX\u0026destCode=DLC\u0026starCity=%E6%B7%B1%E5%9C%B3\u0026arrviceCity=%E5%A4%A7%E8%BF%9E\u0026departureDateStr=2020-09-17\u0026returnDateStr=2020-09-19\u0026adult=1\u0026child=0\u0026infant=0\u0026airCode=DZ\u0026direct=true\u0026noneStop=true"],"Sec-Fetch-Dest":["empty"],"Sec-Fetch-Mode":["cors"],"Sec-Fetch-Site":["cross-site"],"User-Agent":["Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36"],"X-Forwarded-For":["172.31.45.232"]},"CookieMap":{},"Timestamp":1600154854,"Referer":"https://www.donghaiair.com/html/booking-manage/choose-flight-two.html?flightType=1\u0026orgCode=SZX\u0026destCode=DLC\u0026starCity=%E6%B7%B1%E5%9C%B3\u0026arrviceCity=%E5%A4%A7%E8%BF%9E\u0026departureDateStr=2020-09-17\u0026returnDateStr=2020-09-19\u0026adult=1\u0026child=0\u0026infant=0\u0026airCode=DZ\u0026direct=true\u0026noneStop=true","UserAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36","ClientIP":"172.31.45.232","BodyMap":{"adult":"1","child":"0","departureDateStr":"2020-09-16","destCode":"DLC","flightType":"1","infant":"0","orgCode":"SZX","returnDateStr":"2020-09-19","riskToken":"5f5f1d84PI5d65iKhJB2Cd6OYmnpn4gvqn0ChaN1","secCode":"1"},"ParamMap":{},"Payload":"{\"flightType\":\"1\",\"orgCode\":\"SZX\",\"destCode\":\"DLC\",\"departureDateStr\":\"2020-09-16\",\"returnDateStr\":\"2020-09-19\",\"adult\":\"1\",\"child\":\"0\",\"infant\":\"0\",\"riskToken\":\"5f5f1d84PI5d65iKhJB2Cd6OYmnpn4gvqn0ChaN1\",\"secCode\":\"1\"}"},"resp":{"Data":{"data":[{"arriTerm":"","arridate":"2020-09-16 12:20:00","bottomPrice":"600.00","bottomSeat":"A","cabins":[{"airline":"DZ","airportTax":"50.00","basicCabin":"J","cabin":"头等/公务舱","cabinCode":"I","cabinType":0,"calendar":"","chBase":"J","childPrice":"3080.0","currency":"CNY","discountRate":"2.0","dstCity":"SZX","ei":"I舱浮不得签转","fareBasis":"I/CAF20P2","fltNo":"DZ6241","fuelTax":"0.00","fuel_ch":"0.0","gap_ad":"","gap_ch":"","gap_inf":"","infBase":"J","infantPrice":"620.0","onewayPrice":"1230.00","orgCity":"DLC","policePrice":"","roundtripPrice":"","rule":"","seat":"2","soldierPrice":""},{"airline":"DZ","airportTax":"50.00","basicCabin":"Y","cabin":"经济舱","cabinCode":"N","cabinType":1,"calendar":"","chBase":"Y","childPrice":"1230.0","currency":"CNY","discountRate":"2.4","dstCity":"SZX","ei":"不得签转","fareBasis":"TJ24","fltNo":"DZ6241","fuelTax":"0.00","fuel_ch":"0.0","gap_ad":"","gap_ch":"","gap_inf":"","infBase":"Y","infantPrice":"250.0","onewayPrice":"600.00","orgCity":"DLC","policePrice":"","roundtripPrice":"","rule":"","seat":"A","soldierPrice":""},{"airline":"","airportTax":"","basicCabin":"","cabin":"官网专享","cabinCode":"","cabinType":2,"calendar":"","chBase":"","childPrice":"","currency":"","discountRate":"","dstCity":"","ei":"","fareBasis":"","fltNo":"","fuelTax":"","fuel_ch":"","gap_ad":"","gap_ch":"","gap_inf":"","infBase":"","infantPrice":"","onewayPrice":"","orgCity":"","policePrice":"","roundtripPrice":"","rule":"","seat":"已售罄","soldierPrice":""}],"depTerm":"T3","depdate":"2020-09-16 06:55:00","dstAirport3Code":"DLC","dstAirport_ch":"大连周水子机场","dstcity":"DLC","dstcity_ch":"大连","flightType":"OW","flyNo":"DZ6241","flyTime":"05:25","flydistance":"2368","meal":true,"mealCode":"S","orgAirport3Code":"SZX","orgAirport_ch":"深圳宝安机场","orgcity":"SZX","orgcity_ch":"深圳","planestyle":"738","punctualRate":"13.33%","stopTime":"1:25","stopcity":"LYG","stopcity_ch":"连云港","stopnumber":1}],"message":"成功","status":"0"},"IsComplete":true,"Metadata":{"Headers":{"Access-Control-Allow-Headers":["X-Requested-With,b2c-api-user-token,Content-Type","X-Requested-With"],"Access-Control-Allow-Methods":["GET,POST,OPTIONS"],"Access-Control-Allow-Origin":["*"],"Content-Length":["2027"],"Content-Type":["application/json;charset=utf-8"],"Date":["Tue, 15 Sep 2020 07:27:38 GMT"],"Server":["nginx"]},"StatusCode":200}},"is_protected":false}`

type WhiskyMessage struct {
	Req              map[string]interface{} `json:"req"`
	Resp             interface{}            `json:"resp"`
	IsProtected      bool                   `json:"is_protected"`
	GlobalProtection bool                   `json:"global_protection"`
	ProtectionLevel  string                 `json:"protection_level"`
	IsRespOK         bool                   `json:"is_resp_ok"`
	Reason           string                 `json:"reason"`
	Device           map[string]interface{} `json:"device"`
	FinalTestingRes  map[string]interface{} `json:"final_testing_res"`
	Metrics          map[string]interface{} `json:"metrics"`
	DashboardConfig  map[string]interface{} `json:"dashboard_config"`
	KeyParam         map[string]interface{} `json:"key_param"`
	TraceID          string                 `json:"trace_id"`
	UserDefinedParam map[string]string      `json:"user_defined_param"`
	HostEnv          map[string]interface{} `json:"host_env"`
	Upstream         string                 `json:"upstream"`
}

func BenchmarkSonic1(b *testing.B) {
	var res WhiskyMessage
	for i := 0; i < b.N; i++ {
		err := sonic.UnmarshalString(obj1, &res)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGoJson1(b *testing.B) {
	var res WhiskyMessage
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(obj1), &res)
		if err != nil {
			b.Error(err)
		}
	}
}
