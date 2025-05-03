package hlog

import (
	"github.com/hootuu/hyle/hcfg"
	"github.com/hootuu/hyle/hsys"
	"github.com/natefinch/lumberjack/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"sync"
)

type Level = string

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
)

func LevelOf(level Level) zapcore.Level {
	zapLevel := zapcore.InfoLevel
	switch level {
	case DebugLevel:
		zapLevel = zapcore.DebugLevel
	case InfoLevel:
		zapLevel = zapcore.InfoLevel
	case WarnLevel:
		zapLevel = zapcore.WarnLevel
	case ErrorLevel:
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.ErrorLevel
	}
	return zapLevel
}

var gLoggerMap = make(map[string]*zap.Logger)
var gLoggerMu sync.Mutex

func getLogger(key string) *zap.Logger {
	l, ok := gLoggerMap[key]
	if ok {
		return l
	}
	gLoggerMu.Lock()
	defer gLoggerMu.Unlock()
	l, ok = gLoggerMap[key]
	if ok {
		return l
	}
	l = newLogger(key)
	gLoggerMap[key] = l
	return l
}

func newLogger(key string) *zap.Logger {
	rootPath := hcfg.GetString("logger."+key+".root", "./.logs/"+key+"/")
	logLevel := hcfg.GetString("logger."+key+".level", "debug")
	fileName := rootPath + hcfg.GetString("logger."+key+".file", key+".jsonl")
	maxSize := hcfg.GetInt64("logger."+key+".size.max", 16)
	maxBackups := hcfg.GetInt("logger."+key+".backup.max", 30)
	maxAge := hcfg.GetDuration("logger."+key+".age.max", 7)
	compress := hcfg.GetBool("logger."+key+".compress", true)
	hsys.Success("# Initialize the " + key + " log system ..... #")
	hsys.Info(" * PATH: ", strings.ToUpper(rootPath), "      ${ logger."+key+".root }")
	hsys.Info(" * LEVEL: ", strings.ToUpper(logLevel), "      ${ logger."+key+".level }")
	hsys.Info(" * FILE: ", strings.ToUpper(fileName), "      ${ logger."+key+".file }")
	hsys.Info(" * MAX SIZE: ", maxSize, "      ${ logger."+key+".size.max }")
	hsys.Info(" * BACKUP MAX: ", maxBackups, "      ${ logger."+key+".backup.max }")
	hsys.Info(" * AGE MAX: ", maxAge, "      ${ logger."+key+".age.max }")
	hsys.Info(" * COMPRESS: ", compress, "      ${ logger."+key+".compress }")

	zapLogLevel := LevelOf(logLevel)

	hook, err := lumberjack.NewRoller(fileName, maxSize, &lumberjack.Options{
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		LocalTime:  true,
		Compress:   compress,
	})
	if err != nil {
		hsys.Exit(err)
		return nil
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        key,
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapLogLevel)

	writeSyncers := []zapcore.WriteSyncer{
		zapcore.AddSync(hook),
	}

	if hsys.RunMode().IsRd() {
		logWithStdout := hcfg.GetBool("logger.std.out", false)
		if logWithStdout {
			writeSyncers = append(writeSyncers, zapcore.AddSync(os.Stdout))
		}
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writeSyncers...),
		atomicLevel,
	)

	hsys.Success("# Initialize the " + key + " log system [OK] #")

	return zap.New(core)
}
