package Studienarbeit_src

import (
	"crypto/tls"
	"crypto/x509"
	"log"
)

func Tcp_client() {
	log.SetFlags(log.Lshortfile)

	rootPEM := `-----BEGIN CERTIFICATE-----
MIICBjCCAYwCFAZEHGuxiEpA/rmLo4FL/JUOEMs9MAoGCCqGSM49BAMCMGcxCzAJ
BgNVBAYTAkRFMRwwGgYDVQQIDBNCYWRlbi1XdWVydHRtZW5iZXJnMRIwEAYDVQQH
DAlTdHV0dGdhcnQxDTALBgNVBAoMBERIQlcxCzAJBgNVBAsMAklOMQowCAYDVQQD
DAExMB4XDTIwMDIyNjE0MDA0MVoXDTMwMDIyMzE0MDA0MVowZzELMAkGA1UEBhMC
REUxHDAaBgNVBAgME0JhZGVuLVd1ZXJ0dG1lbmJlcmcxEjAQBgNVBAcMCVN0dXR0
Z2FydDENMAsGA1UECgwEREhCVzELMAkGA1UECwwCSU4xCjAIBgNVBAMMATEwdjAQ
BgcqhkjOPQIBBgUrgQQAIgNiAAT9YEH1DcESbPSPwzhtUMjw208CKdOXQoQvRNbm
0YRJlRwduM1vSn3Uxbqapz0Oqd+HBQAHcTaRP8eUWefy7dGdPRe3r6WKiqiPC9Cl
UL2d5xKuMNntlLbZyUERINsOzFEwCgYIKoZIzj0EAwIDaAAwZQIwNaavUgu1FTxw
Suz6rD1XC3TS+QAkhsy/tsa19Bh6F68k+6cs5FDOip0jlYxqeIGgAjEA+qNO32/5
pC6pylq+BDUjRll7oPZSe8moWKDsZsFO36Mx1sZU2siaP5nvz6csJD4P
-----END CERTIFICATE-----`

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}

	tlsConf := &tls.Config{RootCAs: roots}

	conn, err := tls.Dial("tcp", "127.0.0.1:443", tlsConf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	println(string(buf[:n]))
}
