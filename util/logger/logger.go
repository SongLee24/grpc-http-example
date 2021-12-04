package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	LogFile = "./log/app.log"
)

var log *zap.Logger

func init() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:     "time",
		LevelKey:    "level",
		NameKey:     "name",
		MessageKey:  "msg",
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime:  zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
	}

	// 日志轮转
	writer := &lumberjack.Logger{
		// 日志名称
		Filename: LogFile,
		// 日志大小限制，单位MB
		MaxSize: 100,
		// 历史日志文件保留天数
		MaxAge: 30,
		// 最大保留历史日志数量
		MaxBackups: 10,
		// 本地时区
		LocalTime: true,
		// 历史日志文件压缩标识
		Compress: false,
	}

	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(writer),
		zap.DebugLevel,
	)
	log = zap.New(zapCore).With(zap.String("App", "grpc-http-example"),
		zap.String("Author", "https://github.com/SongLee24"))
}

func Warn(msg string, fields ...zap.Field) {
	defer log.Sync()
	log.Warn(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	defer log.Sync()
	log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	defer log.Sync()
	log.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	defer log.Sync()
	log.Fatal(msg, fields...)
}
