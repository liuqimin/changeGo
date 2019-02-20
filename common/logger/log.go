package loggercommon

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func baseLogConf(s string) (logger *zap.Logger) {
	InitialFieldsdata := make(map[string]interface{})
	InitialFieldsdata["type"] = s
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)
	myConfig := zap.Config{
		Level:    level,
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Msg",
			StacktraceKey:  "St",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		InitialFields:    InitialFieldsdata,
		OutputPaths:      []string{"stdout", fmt.Sprintf("./logs/%v.log", s)},
		ErrorOutputPaths: []string{"stderr", "./err-logs"},
	}
	logger, err := myConfig.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	return
}
func Example_basicConfiguration() (logger *zap.Logger) {
	// For some users, the presets offered by the NewProduction, NewDevelopment,
	// and NewExample constructors won't be appropriate. For most of those
	// users, the bundled Config struct offers the right balance of flexibility
	// and convenience. (For more complex needs, see the AdvancedConfiguration
	// example.)
	//
	// See the documentation for Config and zapcore.EncoderConfig for all the
	// available options.
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
      "DisableCaller": "false",
	  "DisableStacktrace": "false",
	  "outputPaths": ["stdout", "./logs/logs"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"type": "bar"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	return
	//logger.Info("logger construction succeeded")
	// Output:
	// {"level":"info","message":"logger construction succeeded","foo":"bar"}
}

func baseLogTest(s string) (logger *zap.Logger) {
	InitialFieldsdata := make(map[string]interface{})
	InitialFieldsdata["type"] = s
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)
	fw := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fmt.Sprintf("./logs/%v.log", s),
		MaxSize:    102400, // megabytes
		MaxBackups: 3,
		MaxAge:     7,    // days
		Compress:   true, // 是否压缩 disabled by default
	})
	myConfig := zapcore.EncoderConfig{
		TimeKey:        "Time",
		LevelKey:       "Level",
		NameKey:        "Name",
		CallerKey:      "Caller",
		MessageKey:     "Msg",
		StacktraceKey:  "St",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	//myConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(myConfig), fw, zap.InfoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(myConfig), os.Stdout, zap.InfoLevel),
	)
	logger = zap.New(core)
	return
}

func NewApplloLog() (logger *zap.Logger) {
	return baseLogConf("itsm")
}
