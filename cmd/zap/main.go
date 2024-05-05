package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	enConf := zap.NewProductionEncoderConfig()
	enConf.EncodeTime = zapcore.ISO8601TimeEncoder
	enConf.CallerKey = zapcore.OmitKey
	enConf.NameKey = zapcore.OmitKey
	enConf.TimeKey = "time"
	encoder := zapcore.NewJSONEncoder(enConf)

	c := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.InfoLevel)
	l := zap.New(c)

	m := make(map[int]OffPart)
	for i := 0; i < 10; i++ {
		m[i] = OffPart{
			Offset:    int64(i),
			NewOffset: int64(i + 1),
		}
	}

	l.Info("map", zap.Reflect("map", m))
}

type OffPart struct {
	Offset    int64
	NewOffset int64
}
