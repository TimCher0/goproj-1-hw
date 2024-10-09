package storage

import (
	"strconv"

	"go.uber.org/zap"
)

type Storage struct {
	ins    map[string]string
	logger *zap.Logger
}

func NewStorage() (Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Storage{}, err
	}
	defer logger.Sync()
	logger.Info("new storage created")
	return Storage{
		ins:    make(map[string]string),
		logger: logger,
	}, nil
}

func (rf Storage) Set(key, value string) {
	rf.ins[key] = value
	rf.logger.Info("key set")
	rf.logger.Sync()
}

func (rf Storage) Get(key string) *string {
	res, tr := rf.ins[key]
	if !tr {
		return nil
	}
	rf.logger.Info("key obtained", zap.String("value", res))
	rf.logger.Sync()
	return &res
}

func (rf Storage) GetType(key string) string {
	var res2 any
	res2, tr2 := strconv.Atoi(rf.ins[key])
	if tr2 != nil {
		res2 = rf.ins[key]
	}
	defer rf.logger.Sync()
	switch res2.(type) {
	case string:
		rf.logger.Info(rf.ins[key], zap.String("type", "string"))
		return "s"
	case int:
		rf.logger.Info(rf.ins[key], zap.String("type", "integer"))
		return "d"
	default:
		rf.logger.Info(rf.ins[key], zap.String("type", "unknown"))
		return "unidentified"
	}
}
