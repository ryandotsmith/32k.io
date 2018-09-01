package mytls

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/crypto/acme/autocert"
)

var acmehost string

func LocalOrLets(dir string) (*tls.Config, error) {
	cert, key, err := findCertKey(dir)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if cert != "" {
		cert, err := tls.LoadX509KeyPair(cert, key)
		if err != nil {
			return nil, errors.New("unable to parse cert or key")
		}
		return &tls.Config{Certificates: []tls.Certificate{cert}}, nil
	} else {
		return LetsEncrypt(dir)
	}
}

func LetsEncrypt(dir string) (*tls.Config, error) {
	cache := filepath.Join(dir, "autocert")
	err := os.MkdirAll(cache, 0700)
	if err != nil {
		return nil, err
	}
	tlsConfig := (&autocert.Manager{
		Cache:      autocert.DirCache(cache),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autoHostWhitelist(),
	}).TLSConfig()

	// Security settings from Filippo's late-2016 blog post
	// https://blog.filippo.io/exposing-go-on-the-internet/.
	tlsConfig.PreferServerCipherSuites = true
	tlsConfig.CurvePreferences = []tls.CurveID{
		tls.CurveP256,
		tls.X25519,
	}
	tlsConfig.MinVersion = tls.VersionTLS12
	tlsConfig.CipherSuites = []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	}

	return tlsConfig, nil
}

func findCertKey(dir string) (cert, key string, err error) {
	cert = filepath.Join(dir, "localhost.pem")
	key = filepath.Join(dir, "localhost-key.pem")
	_, err = os.Stat(cert)
	if err != nil {
		return "", "", err
	}
	fi, err := os.Stat(key)
	if err != nil {
		return "", "", err
	}
	if fi.Mode()&077 != 0 {
		return "", "", fmt.Errorf(key, "must be accessible only to current user")
	}
	return cert, key, nil
}

// autoHostWhitelist provides a TOFU-like mechanism as an
// autocert host policy. It whitelists the first-requested
// name and rejects all subsequent names.
func autoHostWhitelist() autocert.HostPolicy {
	return func(ctx context.Context, host string) error {
		if acmehost == "" {
			fmt.Println("adding %s to acmehost", host)
			acmehost = host
		}
		if acmehost == host {
			return nil
		}
		return errors.New("host name mismatch")
	}
}
