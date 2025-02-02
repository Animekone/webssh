package core

import (
	"github.com/gorilla/websocket"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"unicode/utf8"
)

// SSHClient 结构体
type SSHClient struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	IPAddress string `json:"ipaddress"`
	Port      int    `json:"port"`
	LoginType int    `json:"logintype"`
	Client    *ssh.Client
	Sftp      *sftp.Client
	StdinPipe io.WriteCloser
	Session   *ssh.Session
}

type wsOutput struct {
	ws *websocket.Conn
}

// Write: implement Write interface to write bytes from ssh server into bytes.Buffer.
func (w *wsOutput) Write(p []byte) (int, error) {
	// 处理非utf8字符
	if !utf8.Valid(p) {
		bufStr := string(p)
		buf := make([]rune, 0, len(bufStr))
		for _, r := range bufStr {
			if r == utf8.RuneError {
				buf = append(buf, []rune("@")...)
			} else {
				buf = append(buf, r)
			}
		}
		p = []byte(string(buf))
	}
	err := w.ws.WriteMessage(websocket.TextMessage, p)
	return len(p), err
}

// NewSSHClient 返回默认ssh信息
func NewSSHClient() SSHClient {
	client := SSHClient{}
	client.Username = "root"
	client.Port = 22
	return client
}
