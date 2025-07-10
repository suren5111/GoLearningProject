package cache

import (
	"StudentService/config"
	"StudentService/model"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

var RDB *redis.Client
var ctx = context.Background()

func InitRedis(cfg config.RedisConfig) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}

func GetStudentFromCache(id string) (*model.Student, error) {
	val, err := RDB.Get(ctx, id).Result()
	if err != nil {
		return nil, err
	}
	var student model.Student
	json.Unmarshal([]byte(val), &student)
	return &student, nil
}

func SetStudentToCache(student *model.Student) error {
	data, _ := json.Marshal(student)
	return RDB.Set(ctx, student.ID, data, 10*time.Minute).Err()
}

func DeleteStudentFromCache(id string) error {
	return RDB.Del(ctx, id).Err()
}

func SetToken(token string) error {
	return RDB.Set(ctx, "token:"+token, "1", 10*time.Minute).Err()
}

func IsTokenValid(token string) bool {
	val, err := RDB.Get(ctx, "token:"+token).Result()
	return err == nil && val == "1"
}

func DeleteToken(token string) error {
	return RDB.Del(ctx, "token:"+token).Err()
}
