package Studienarbeit_src

import (
	"crypto/tls"
	"crypto/x509"
	"log"
)

func Tcp_client() {
	log.SetFlags(log.Lshortfile)

	rootPEM := `-----BEGIN CERTIFICATE-----
MIICKjCCAYygAwIBAgIQSNQ20G+zn//2Y09EgWn4jDAKBggqhkjOPQQDBDAuMRQw
EgYDVQQKEwtMb2cgQ291cmllcjEWMBQGA1UEAxMNMTkyLjE2OC4yLjExMTAeFw0y
MDAyMjgwNzQxMDJaFw0yMTAyMjcwNzQxMDJaMC4xFDASBgNVBAoTC0xvZyBDb3Vy
aWVyMRYwFAYDVQQDEw0xOTIuMTY4LjIuMTExMIGbMBAGByqGSM49AgEGBSuBBAAj
A4GGAAQAayh+so/Hvc3ivMr09cv0GmnxpOc3YxdOyaT6zxom92Zob2OyMee1dQfL
fmFMdLdLT2aW4gA2e8oAHfStNOHaQrwAx8DnuBeysdE+sfP8qAH5nxMGxxfU7u8H
jYGL14HZKtKLn+s5+bFFoon9+kQKZ7a685xWZg8q6p5j+UGZyWgHBLKjSTBHMA4G
A1UdDwEB/wQEAwICpDATBgNVHSUEDDAKBggrBgEFBQcDATAPBgNVHRMBAf8EBTAD
AQH/MA8GA1UdEQQIMAaHBMCoAm8wCgYIKoZIzj0EAwQDgYsAMIGHAkIAztMxt0qW
RsA0ImAjE9qpTRZ0Wx40uHJRHx6lyYczhf4DUQBZ3zEca4wGvI/q5kNuVH9N1gEU
UruJXF2D+sv8NjACQXPhXvWANi4DY8ijJwqJsQaWJ61Muo/QfllRuvu/XUG4cSqM
w5o5zhS94elVruAhIcOBA4Vnig2LCXPQiPyURSXf
-----END CERTIFICATE-----`

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}

	tlsConf := &tls.Config{RootCAs: roots}

	conn, err := tls.Dial("tcp", "192.168.2.111:443", tlsConf)
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
