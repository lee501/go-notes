package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var CA_CERT = []byte(`-----BEGIN CERTIFICATE-----
MIIDADCCAegCCQC2TZdnue2EVzANBgkqhkiG9w0BAQsFADBCMQswCQYDVQQGEwJY
WDEVMBMGA1UEBwwMRGVmYXVsdCBDaXR5MRwwGgYDVQQKDBNEZWZhdWx0IENvbXBh
bnkgTHRkMB4XDTIxMDgyNjA2MjAwMloXDTMxMDgyNDA2MjAwMlowQjELMAkGA1UE
BhMCWFgxFTATBgNVBAcMDERlZmF1bHQgQ2l0eTEcMBoGA1UECgwTRGVmYXVsdCBD
b21wYW55IEx0ZDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMnUhaJK
oawAa1cTtFONT7Bst6pdUE+GwmCEsSiwqq+jVuBy8Vf6Qmnx7UREnXu4lhjdiUmZ
+yzP0SrP6tQpSnrUItFvV8xypzmdQ5KWdLpZ439/1pB0bUV17ha90npBTJDHs6wi
ou1BIgGPjBhpJIOCPsvzzUb3YNTrX6kHCnb017DPeo/I7lHfAIwcE6buzWpGaaS6
+H29v+9dxrCHQve3e8gMH9Y1L7xKxkexdt3idapCSQUOBAgIiDTtKYo7sWxiDelA
CenQ/ws6NLvovHFuI8CAh2+Jdc0ojY9T92B8Qispbsxr68kWw37QGVCjp4uwWLvo
NeVzf90V1qtiCv8CAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAorjPH+r67JabuKS5
U3rs4QZ37NnRLN+ohrubl8y8TbaSDVPdZ3Z4zq1QNtd19pTbXzNY8XPQo/oHYGPx
Xr9JYNY9wbUasz0c4FA7OgVarxSg6xHl1aYEDNnyBTsYm45mmB8Hex30cd7TH/aj
ABKIcuN1VlsyXcLP9aCNDyPxqF9fYzZ25s6QES03+DEW34IXruPMslBVe2cQzTdS
17r1AzAV/UWjyXaKm4bKuorbouIgrrP1LcrRYYXXm7gxXIIg1Q9YOhH8cBytrLkq
E05NV5IiwEKFBod7gDu3TlPQMLvs3FAbPWvvxMqpKTbRDSKmJQ4mDT9M7GQr1M4D
x0MX9g==
-----END CERTIFICATE-----`)

func welcome(res http.ResponseWriter, req *http.Request) {
	//res.Header().Set("Content-Type", "text/html;charset=usf-8")
	//fmt.Fprintln(res, "Welcome to <b>go</b>")
	url := "https://www.jianshu.com/favicon.ico"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	res.Write(b)
}

func cert(res http.ResponseWriter, req *http.Request) {
	res.Write(CA_CERT)
}
func main() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/cert", cert)
	http.ListenAndServe(":3301", nil)
	fmt.Println("服务已启动")
}
