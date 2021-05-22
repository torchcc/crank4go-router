package crank_router

import (
	"crypto/tls"
	"path"

	"github.com/torchcc/crank4go/test"
	"github.com/torchcc/crank4go/util"
)

func GetTLSConfig() *tls.Config {
	// return nil
	cert, err := tls.LoadX509KeyPair(path.Join(test.TestDir, "static/cert/server.pem"), path.Join(test.TestDir, "static/cert/server.key"))
	if err != nil {
		util.LOG.Errorf("failed to load ssl certificate: err: %s", err)
		return nil
	}
	return &tls.Config{Certificates: []tls.Certificate{cert}}
}
