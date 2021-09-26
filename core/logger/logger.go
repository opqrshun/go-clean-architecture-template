package logger

import (
	"go.uber.org/zap"
)

// type Logger zap.Logger
var logger *zap.SugaredLogger

func init() {
	l, _ := zap.NewProduction()
	defer l.Sync()
	logger = l.Sugar()

}

func GetLogger() *zap.SugaredLogger {
	return logger
}

// func (l *Logger) Fatal(message string , err error){
//   l.Fatal(message, zap.Error(err))
// }

// func (l *Logger) Error(message string , err error){
//   l.Error(message, zap.Error(err))
// }

// func InfoU(message string , usecase string,userID string){
//   Logger.Info(message, zap.String("usecase", usecase), zap.String("userID", userID))
// }

// func ErrorU(message string , usecase string,userID string){
//   Logger.Error(message, zap.String("usecase", usecase), zap.String("userID", userID))
// }

// func FatalU(message string , usecase string,userID string){
//   Logger.Fatal(message, zap.String("usecase", usecase), zap.String("userID", userID))
// }
