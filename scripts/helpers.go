package auth_gateway

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

// GenerateCert generates a self-signed certificate with the given subject.
func GenerateCert(subject string, expiresAt time.Time) ([]byte, error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err!= nil {
		return nil, err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:     pkix.Name{CommonName: subject},
		NotBefore:   time.Now(),
		NotAfter:    expiresAt,
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:    []string{subject},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err!= nil {
		return nil, err
	}

	return pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes}), nil
}

// GetLocalIP returns the IP address of the local machine.
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err!= nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok &&!ipnet.IP.IsLoopback() {
			if ipnet.IP.To4()!= nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", err
}

// GetLogger returns a logger instance.
func GetLogger() *logrus.Logger {
	return logrus.New()
}