package config

import (
	"net"
	"os"

	"github.com/pkg/errors"
)

const (
	authHostEnvName = "AUTH_HOST"
	authPortEnvName = "AUTH_PORT"
	authCertEnvName = "AUTH_CERT"
	authKeyEnvName  = "AUTH_KEY"
)

// AuthConfig grpc server config
type AuthConfig interface {
	Address() string
	Cert() []byte
	Key() []byte
}

type authConfig struct {
	host string
	port string
	cert []byte
	key  []byte
}

// NewAuthConfig create new grpc server config
func NewAuthConfig() (AuthConfig, error) {
	host := os.Getenv(authHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("auth host not found")
	}

	port := os.Getenv(authPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("auth port not found")
	}

	cert := os.Getenv(authCertEnvName)
	if len(cert) == 0 {
		return nil, errors.New("auth cert not found")
	}

	key := os.Getenv(authKeyEnvName)
	if len(key) == 0 {
		return nil, errors.New("auth key not found")
	}

	return &authConfig{
		host: host,
		port: port,
		cert: []byte(cert),
		key:  []byte(key),
	}, nil
}

// Address get grpc server address
func (cfg *authConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

// Cert get service pem cert
func (cfg *authConfig) Cert() []byte {
	return cfg.cert
}

// Key get service key
func (cfg *authConfig) Key() []byte {
	return cfg.key
}
