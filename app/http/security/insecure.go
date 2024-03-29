package security

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

const certPEM = `-----BEGIN CERTIFICATE-----
MIIF0zCCA7ugAwIBAgIUWvRa4/CCUzVY5I/YFTDfUuSs1v4wDQYJKoZIhvcNAQEL
BQAweTELMAkGA1UEBhMCSUQxEDAOBgNVBAgMB0pha2FydGExEDAOBgNVBAcMB0ph
a2FydGExITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDEjMCEGCSqG
SIb3DQEJARYUYWdhbXdvcmsyOEBnbWFpbC5jb20wHhcNMjIwMzE3MDUzNDU3WhcN
MjIwNDE2MDUzNDU3WjB5MQswCQYDVQQGEwJJRDEQMA4GA1UECAwHSmFrYXJ0YTEQ
MA4GA1UEBwwHSmFrYXJ0YTEhMB8GA1UECgwYSW50ZXJuZXQgV2lkZ2l0cyBQdHkg
THRkMSMwIQYJKoZIhvcNAQkBFhRhZ2Ftd29yazI4QGdtYWlsLmNvbTCCAiIwDQYJ
KoZIhvcNAQEBBQADggIPADCCAgoCggIBAN4oqMmCMmdw21qHK9e2V1jMCDV00vLR
W4ToHDxU/ALPiM1Gg8R65HONCJL70i5hHFhTJoh4jTHTtZ/JVqvr1XXzq/Wbg/zI
OzRg4Z3KSxAsejhynJIPXtsMmqYBZwUkYVymCcYdeK2W3WYzwzby6lByTXQKVtW3
gCiM9sFC+8XIq7wia8bgyKcGGmb3yQnxdRyZ8v2qdRS93k6ne0afC6H/Agnk9gyi
8cAJHhnygiPARIYPOJzzEGYowvi/lp9ebrdWIQ1lbQ/ULTUrd8prFh+rBNFcWEY5
4pFcz6laC6x9hq0F/rQ+s+0FWQ7rXvXvQG1KqI1gdG1aLoOcjHLaF2kUWOKSGnr+
VcBSGZElLv+D/ZhN8c3BA9B8K3jEMP2IT6UqXxM5/RDxVfF+LoJuQULTrEEPIN6U
awiRWOrEhne1yMLbKpAEkPE9tTuAOcOt1jPj3XNuw/fJN0epHSb0O2XEczbJIcnR
QE0Kd0d/1BtSgRPv5C/14aREC/vAe6o5pQKMw5F55FpGrce1eBpn/vsJBOFgMvfZ
h4dTg1Z88vlovcgTQj/FCiVqXoZ96Xl/yjTnsmZZ08PgkRO8SaQu6cuOcG0cSQDk
d9ruxyHZMg3epMleJnBXJTNzPQLu72RcXo6PvfPaGaCrtNjV9ZYxZdGWnPDq/GfR
n0KWL+SrE7JXAgMBAAGjUzBRMB0GA1UdDgQWBBQgTTivJeP4YvxNYQe39bFxvfO3
4DAfBgNVHSMEGDAWgBQgTTivJeP4YvxNYQe39bFxvfO34DAPBgNVHRMBAf8EBTAD
AQH/MA0GCSqGSIb3DQEBCwUAA4ICAQB8m1XZ6ULaNunlbft7MHAy3ROLwJFxJxYZ
GgxHyegyIpsc0NdXKTCPwPDNhFgjZWNjN6BCw7MK4Boi9qJ8hpCt7V+HeKJBj2Rc
W4QQ5mlsbyylPu7AjJjrhZIXrlg+vlySm5c2vCE7kHFXBqNsvZ6ymaJcYoUBY6/U
uIIxf0vKxNx3eWgP5P+2fBHb6Dc0CDZBOOgDRBg7uTUq6JBiaa/WLiJHllhMidFw
39j75QLtZezEWG+tS5mKnFGNG5yXLru7B4AX4mU6zATDj5Uft/eaOdI58Pdrmp5Q
24QvYgL63k0oAsAegdAH5oVjTjIi5jIcX+ByiyZPS5qbDJi2945q3yOEbzYQismz
mbC3Woy1fZzHGiN8TLsdK5+qY8z/5xda7yT6grU3SgChWIwWW6yPGwkVCFdwaUb7
vUxZwO0tlhfdbc2aJ0fnRQg0scB4fBU+PgEiQUtwwlS9iC5W1Mo/gmuJ5pGBt2Sy
p2iG5RHK2HpkwpPSoN2GfDi1qVxo7nJvzxh7ruNmJp9JAiFXIq4p+MNlRzAXioY2
2lCO+Qet/UMIk/JP0Xmr6Vf7rDoXtf6SRPGjggoMhxeGsHJHN1ZXAiv43k3RaT7s
QE9YUucg7Svu2yXCw+zkzo+s2XWIoHe2a9t9LzE57cS98DYM5O8uD59cHnNMJq5c
GcACpcJUTw==
-----END CERTIFICATE-----
`

const keyPEM = `-----BEGIN PRIVATE KEY-----
MIIJQgIBADANBgkqhkiG9w0BAQEFAASCCSwwggkoAgEAAoICAQDeKKjJgjJncNta
hyvXtldYzAg1dNLy0VuE6Bw8VPwCz4jNRoPEeuRzjQiS+9IuYRxYUyaIeI0x07Wf
yVar69V186v1m4P8yDs0YOGdyksQLHo4cpySD17bDJqmAWcFJGFcpgnGHXitlt1m
M8M28upQck10ClbVt4AojPbBQvvFyKu8ImvG4MinBhpm98kJ8XUcmfL9qnUUvd5O
p3tGnwuh/wIJ5PYMovHACR4Z8oIjwESGDzic8xBmKML4v5afXm63ViENZW0P1C01
K3fKaxYfqwTRXFhGOeKRXM+pWgusfYatBf60PrPtBVkO617170BtSqiNYHRtWi6D
nIxy2hdpFFjikhp6/lXAUhmRJS7/g/2YTfHNwQPQfCt4xDD9iE+lKl8TOf0Q8VXx
fi6CbkFC06xBDyDelGsIkVjqxIZ3tcjC2yqQBJDxPbU7gDnDrdYz491zbsP3yTdH
qR0m9DtlxHM2ySHJ0UBNCndHf9QbUoET7+Qv9eGkRAv7wHuqOaUCjMOReeRaRq3H
tXgaZ/77CQThYDL32YeHU4NWfPL5aL3IE0I/xQolal6Gfel5f8o057JmWdPD4JET
vEmkLunLjnBtHEkA5Hfa7sch2TIN3qTJXiZwVyUzcz0C7u9kXF6Oj73z2hmgq7TY
1fWWMWXRlpzw6vxn0Z9Cli/kqxOyVwIDAQABAoICAGJB5XVKqnRpy/rcMh30Pem+
RBswkmRnCG1eI67jf5rWo5D3ESyDlistQ4+O5UUyGlFOZYqSlxBkrL4BVN4xLnY+
0d2NbPEOtnDSMLeTU2PR6DSzLGDFf5SunG+zyxbRImbR4RzIyCkuwrmDR2I0SkJO
a8dZE1LvGDXXMwQDVkwCF0LcjoKlxNlqr56/FSHas4t0464iCCqmV36+QJ3ZllGN
8h7iL5kjMH1ZCwE66yhYz3XjoYx5TNWCTcmrIzlVkO5xn6jHJVeM5nnPwem7N0vG
XhkjGL/mtjjIregfNW6TDk/UiPS7juTUKrGYlKV9ricokCgHczcyXcREJzJdF2DE
8CZHvPtJmlSdLb/wnzE4VBJHfxugLGQQuCGRfqeZA7upWv7qU+UwkUH+xgGJ8eD+
ADf2UZAhlcmk8jZuNLHnwQOUQSdXLoKLlKZ+rYjHQV/TEjvVGxPa/7BteD/gancW
VyQepnDtmyIx0woe3CpsYuKi/E4YdkHMqBHmudJEvmn0oXroWD66vx3eIMlnBnku
maSxlkoKaBjRPlxpTRjBww7tC8HN5T//HoDfEPERfQL8SLAulI6p7i+ZqE779Mo+
9i35yJaCqNZWCxpiE58MkZpcc2amvtXU0X+E5N1bYJ2IlwO2yCSvsHR/VaySpPkO
yqbmZO3Xp8cMTNSMJ1phAoIBAQD/xko7Mig3vwwaHQyekNVqu8+0C5g3Mn2zM9/j
sLQ7InV4cBKPd/2/OYjgS9uvV7KZCBe7/wdVFUpzCVMYhnD5wmRVeDSMD75T60B/
p+do7KtS6spUG9bitZXXQ9Ssy84kp7xGKoQWCUkB6DmYIlOSAjbnuQu0HgtKr/CL
4cnu5euH37My0nwZuGX76Q8tftZSFrxs2nSf3X8yeMIA5BtspoUtdWFW3chPPo30
0WFRlAUclMNjbqDwr5g5q6CD92nZAp+0MTbxW4nvb725isykfXZZjcLUiI6njzOn
nrvpu67W4X3H49YiMmHNGxmCqULUQ0DR+hrQ7OnQ+6h531wJAoIBAQDeWsjhXLNN
fV9QSdbn89W0qx/hnomfuHNma27kxhVoj6fMcOYoqrPy8Q6Vp3CbYM2BaLSrPi/k
MOKj9KpNqn9ore7XVPSKl/hE2jr8b6sWIy2hKaPCAyQdUjO1VW6tKHuMqbyg765D
wTfitlbARDyBDbxtBRolf7+84VRBaOLLlFZsLLwwL2RgKr0Ee36J73DIW2aK7UwX
OABzgLNB+/LkFmhKcSbEKyVWoSVDiT82p3o8BmeZ60gCQt3YbaU8O5QI51+S9S7e
d66Sag7QslHq7Ri1qNAKm/eouni5FWThy0xO0jati8Pcm3UccOTuT+20IjV8FhB7
UkkMh/ijNPNfAoIBAQDv7u0XlrePg9IwdOuvm6H4G/J9rRq/PMgZ4SG3j72ic3LK
5icjGqaYSbqS3dZZJ+VMFQFew5/3Op+6RhcDPKqiQK6qVrbEZMhgJKE+nx0/mjQk
9keSmwL7bOgBJYpOBml8jTTahWue2Fr6RQQyKJFIuHcU7MuDzWM+TslSDx7E0j4e
GcUaecUcw5Q3uasZPJCENPvRCSS6CqgVip5DA7iONa9cs236dNZBMCcub3PheP9K
jmqzzcV/wtkD0HIlGWTbw9BY0on946cDS0ko0oGiFGd/HrKTeP6JRJZqOUeBvSQI
Q0MgziPDtIt2QmOt5GXihKw7zWh2cKft+QVUPikxAoIBAAm4bEm11LzRe0Y2RCfV
uATy6izh0mBiHuFeVvW3Vx5AEXTVhqpGdCFtulZ+gGS6gaFjo2pNCKjwHihOvNcy
WemImQVe2GXqqNbRt3xroU4RcaA3AsqrenuxqUMHHhSBzYzjqrRaZgLeQoQttnMn
SIHXyOH3NCJpiEdIjndhj3CY5oJKHlv45mlx1NuKTojXJ9YKcNwn7YiiURq7iJqc
2gxwklSWx7wRqssztj5SlKnd/nD7UrIe4AIBTFN2z59nd1xASSRShZ8CGmaaLlWf
RgeAH/FpFNcx5EzDAh4NOmDIk2hqtis2DfTp8t+wa8GjSt/yEuZ+zG+ntDvRAeSY
nv8CggEARrRpfwCOuYCvIGwCQ1+77xVxCSbp8FA1ATGltorOIyLWLMv6AAdJmTG7
iUlItCdSydT4mEAG2bqOQbXRENg/Lv0wtDyUAsoV1k/NB4/B5N3zzff3b03lqpO1
sDm6ydLzjv6DSFHu3Eck7S3en3jR6UG/upc89O9scbg8ySL4qW0XAphuvxRRGI0+
qf5a+N06toPMRFk0PhtKh6T6VwWS9VvHDZ6SV//4UpJ4wA8+jvfwXdqGvyYAVsZ+
y++y5ej/PejjdmWGR6tEbmwUzTTO/Edux7bBdH7riJNzX51SywYmG5T5kItbozVb
+/QcDCGAoLCTPZbXwaVf68j2/fOY9A==
-----END PRIVATE KEY-----

`

var (
	Cert                tls.Certificate
	CertPool            *x509.CertPool
	CredTransportServer credentials.TransportCredentials
	CredTransportClient credentials.TransportCredentials
)

func Init() {
	// it shows your line code while error
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var err error
	Cert, err = tls.X509KeyPair([]byte(certPEM), []byte(keyPEM))
	if err != nil {
		log.Fatalln("Failed to parse key pair:", err)
	}
	Cert.Leaf, err = x509.ParseCertificate(Cert.Certificate[0])
	if err != nil {
		log.Fatalln("Failed to parse certificate:", err)
	}
	CertPool = x509.NewCertPool()
	CertPool.AddCert(Cert.Leaf)

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("./app/http/security/cert/server-cert.pem", "./app/http/security/cert/server-key.pem")
	if err != nil {
		log.Fatalln(err)
	}

	// Create the credentials and return it
	configServer := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}
	CredTransportServer = credentials.NewTLS(configServer)

	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("./app/http/security/cert/ca-cert.pem")
	if err != nil {
		log.Fatalln(err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		log.Fatalln(fmt.Errorf("failed to add server CA's certificate"))
	}

	// Create the credentials and return it
	configClient := &tls.Config{
		RootCAs:   certPool,
		ClientCAs: certPool,
	}
	CredTransportClient = credentials.NewTLS(configClient)
}

func LoadTLSCredentialsServer() credentials.TransportCredentials {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("./app/http/security/cert/server-cert.pem", "./app/http/security/cert/server-key.pem")
	if err != nil {
		log.Println(err)
		return nil
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config)
}

func LoadTLSCredentialsClient() credentials.TransportCredentials {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("./app/http/security/cert/ca-cert.pem")
	if err != nil {
		log.Println(err)
		return nil
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		log.Println(fmt.Errorf("failed to add server CA's certificate"))
		return nil
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs:   certPool,
		ClientCAs: certPool,
	}

	return credentials.NewTLS(config)
}
