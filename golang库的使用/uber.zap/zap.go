package uber_zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)
//订制日志

/*
	通过zap.New(...)手动传递需要的配置
	func New(core zapcore.Core, options ...Option) *Logger
*/

//zapcore.Core需要三个配置
/*
	1.Encoder 日志写入格式 NewConsoleEncoder() 和 NewJSONEncoder()
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())

	2.WriteSyncer 日志写入的位置
		file, _ := os.Create(./log.log)
		writeSyncer := zapcore.AddSync(file)

	3. log level 日志等级
*/

//demo
var Logger *zap.Logger

func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
}

func getWriteSyncer() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}

func InitLogger() {
	writeSyncer := getWriteSyncer()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	Logger = zap.New(core, zap.AddCaller(), zap.Development())
}
