package mail

import (
	"bytes"
	"errors"
	"net/smtp"
)

type loginAuth struct {
	username, password string
	host               string
}

func LoginAuth(username, password, host string) smtp.Auth {
	return &loginAuth{username, password, host}
}

func (a loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	if server.Name != a.host {
		return "", nil, errors.New("wrong host name")
	}
	return "LOGIN", nil, nil
}

func (a loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		if bytes.EqualFold([]byte("username:"), fromServer) {
			return []byte(a.username), nil
		} else if bytes.EqualFold([]byte("password:"), fromServer) {
			return []byte(a.password), nil
		}
	}
	return nil, nil
}
