package service

import "github.com/tencent-connect/botgo/log"

var logger log.Logger

func init() {
	logger = log.DefaultLogger
}
