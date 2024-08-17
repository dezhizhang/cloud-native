package logs

import "go.uber.org/zap"

var LG *zap.Logger

type LogConfig struct {
	DebugFileName string `json:"debug_file_fame"`
	InfoFileName  string `json:"info_file_fame"`
	WarnFileName  string `json:"warn_file_fame"`
	MaxSize       int    `json:"max_size"`
	MaxAge        int    `json:"max_age"`
	MaxBackups    int    `json:"max_backups"`
}

// InitLogger 初始化日志
func InitLogger(cfg LogConfig) (err error) {
	return nil
	//writeSyncerDebug := getLogWriter(cfg.DebugFileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
}

func getLogWriter() {

}
