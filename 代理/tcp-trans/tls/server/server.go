package main

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
)

var (
	cert = []byte(`-----BEGIN CERTIFICATE-----
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

	key = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAydSFokqhrABrVxO0U41PsGy3ql1QT4bCYISxKLCqr6NW4HLx
V/pCafHtRESde7iWGN2JSZn7LM/RKs/q1ClKetQi0W9XzHKnOZ1DkpZ0ulnjf3/W
kHRtRXXuFr3SekFMkMezrCKi7UEiAY+MGGkkg4I+y/PNRvdg1OtfqQcKdvTXsM96
j8juUd8AjBwTpu7NakZppLr4fb2/713GsIdC97d7yAwf1jUvvErGR7F23eJ1qkJJ
BQ4ECAiINO0pijuxbGIN6UAJ6dD/Czo0u+i8cW4jwICHb4l1zSiNj1P3YHxCKylu
zGvryRbDftAZUKOni7BYu+g15XN/3RXWq2IK/wIDAQABAoIBAQCMigr/xl4riiM0
gEkKARVjcWC3JcEdj2XclalS/ynroZ6GLiO2MlJ6uXaRGgJbO15xNBo8ARfwzUkG
Ob94OsRdIUiZc6G8gH3HVaXO9iuT87HBf4AmGJxaSNMgTM+6R6wMCLMpAvlWogsO
8SkJenVtiY02sa7YYruvp7J13axAxeFeiQtMenaqLrkfLtzm8G5s/hXXudtxRR7R
SetD6GDc5F8RfA16treG4FDXJ7Vrv+m3v+toTFB8RO7ZfNpq7rU3hEAu1jFTI7oA
lbr+MnaOLj/wsIHieXbF6UgN7LUfFZE+AKQC6ZZzGtvA2vXkJoGxRL3+IfctXjPV
ZhBMVrHBAoGBAPlClS34BDW3JWuq8sDBkQg9iXpziTlLXoWDOuM2fGe4foRB6S+t
7JOrI10d96o0grvJeZV08NvKytW78IiNRDo3sjdE11OV4YBmo6uM0FUf1j6sAUDU
+Rk3DPWl1EYrtI4EsRYWBLJZ1rCGdrwZa1hWQkjz8cqpRalPVXYHU8NfAoGBAM9J
nzaqxomxNOBHyqAgT7uNO9U8rGti/nRLslqzhv1kUX+bCDPNAEJ1oAWFNMdAUwMG
dJJI6UHWuInRCH/kVtjXhgIlKsNYn/4pWuD3JZ2q86nqIrZISLfqj9xgsKdW/t4y
lnh/71v21IIiVXpebJN92V5ipT6OrHFlTN/33XxhAoGBAOwU3CaF2bbmmFK5vuJC
c5NSdXu9IiNZ91SNTqEVYg769xldM/csy/2xAfWYBJL42TnzQW0FLXt/P03gi5lJ
820Qg4NFeIx6UPLERBCfdlKprMn/L9Cv5p6zPVsMjVlI2+IHH5LDl/80h9r6AEt3
5+vrBy0Bn+lLGeeadF8t4XARAoGABbmQzh+1UAQ246LX88gwq4thBaihUm3vfSLC
EEM5DTKCFbYgad3VjgBVpRFivcYHBORev9OHravSeOvvAik6RM7fApwoLDD8ajaW
LfRlc49PH3g58TLZ30p5IqsA9f4vF8/p2/YclDui7t0n4zNMaF5nbHXo/mtOU+tq
Th04XsECgYEAuIIjNYIANsCXNPxP+gaXvafa5rNZiZ5oQ9NqfMSt3TMyZ3eFymWt
tNw/x00bV1kGVD3RMkQY4mpiwpSrNqAG1UGvUm2Tis0NEiNTng64ayJhXAbZG4Je
Ccc6ZQbWLeqQh0gcbxqc4UUnTZI0sjcVtpLS5+y99oiBfkEY3iWFS+k=
-----END RSA PRIVATE KEY-----`)
)

func main() {
	cert, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Println(err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	p := make([]byte, 1024)
	for {
		n, err := r.Read(p)
		if err != nil {
			log.Println(err)
			return
		}
		println(string(p))
		_, err = conn.Write([]byte("hello client\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}
