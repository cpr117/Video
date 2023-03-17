// @User CPR
package utils

import (
	"MyBilibili/config"
	"context"
	"github.com/redis/go-redis/v9"
	_ "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"log"
	"time"
)

const REDIS_UTIL_ERR_PREFIX = "utils/redis.go ->"

var (
	ctx = context.Background()
	rdb *redis.Client
)

var Redis = new(_redis)

type _redis struct{}

func InitRedis() *redis.Client {
	redisCfg := config.Cfg.Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	// 测试连接状况
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Panic("Redis 连接失败: ", err)
	}
	log.Println("Redis 连接成功 ")
	return rdb
}

// 通用执行指令方法
// func (*redisUtil) Execute(cmd string, args ...any) (any, error) {
// 	return rdb.Do(ctx, args...).Result()
// }

// Keys
// redis 获取根据匹配项获取键名列表
func (*_redis) Keys(pattern string) []string {
	return rdb.Keys(ctx, pattern).Val()
}

// Del
// redis 删除值
func (*_redis) Del(key string) {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"Del: ", zap.Error(err))
		panic(err)
	}
}

// Set
// redis 设置值
func (*_redis) Set(key string, value interface{}, expiration time.Duration) {
	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"Set: ", zap.Error(err))
		panic(err)
	}
}

func (*_redis) Incr(key string) int64 {
	ret := rdb.Incr(ctx, key)
	if err := ret.Err(); err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"Incr: ", zap.Error(err))
		panic(err)
	}
	return ret.Val()
}

// GetVal
// redis string类型
func (*_redis) GetVal(key string) string {
	ret := rdb.Get(ctx, key)
	if err := ret.Err(); err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"GetVal: ", zap.Error(err))
	}
	return ret.Val()
}

// GetInt
// redis获取数字
func (*_redis) GetInt(key string) int {
	val, err := rdb.Get(ctx, key).Int()
	if err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"GetInt: ", zap.Error(err))
		panic(err)
	}
	return val
}

// 从redis中取值，不存在会有 redis:nil 的错误
func (*_redis) GetResult(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

// 往[set]中添加元素
func (*_redis) SAdd(key string, members ...any) {
	err := rdb.SAdd(ctx, key, members...).Err()
	if err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"SAdd: ", zap.Error(err))
		panic(err)
	}
}

// 判断 元素 是否是 [集合(Set)] 的成员
func (*_redis) SIsMember(key string, member any) bool {
	return rdb.SIsMember(ctx, key, member).Val()
}

// 获取 [集合(Set)] 的成员列表
func (*_redis) SMembers(key string) []string {
	return rdb.SMembers(ctx, key).Val()
}

// 移除 [集合(Set)] 中的元素
func (*_redis) SRem(key string, member any) {
	rdb.SRem(ctx, key, member)
}

// 为[哈希表(Hash)]中的字段值加上指定增量值 (可以为负)
// 如果 key 不存在, 自动创建哈希表并执行操作
// 如果 field 不存在, 创建该字段值并初始化为 0
func (*_redis) HIncrBy(key, field string, incr int64) {
	err := rdb.HIncrBy(ctx, key, field, incr).Err()
	if err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"HIncrBy: ", zap.Error(err))
		panic(err)
	}
}

// 获取[哈希表(Hash)]中 指定字段的值
func (*_redis) HGet(key, filed string) int {
	val, _ := rdb.HGet(ctx, key, filed).Int()
	return val
}

// 获取[哈希表(Hash)]中 所有的字段和值
func (*_redis) HGetAll(key string) map[string]string {
	return rdb.HGetAll(ctx, key).Val()
}

// func (*redisUtil) ZAdd(key string, me) {
// rdb.ZAdd()
// }

func (*_redis) ZRangeWithScores(key string, start, stop int64) []redis.Z {
	return rdb.ZRangeWithScores(ctx, key, start, stop).Val()
}

// 获取[有序集合]中, 成员的分数值
func (*_redis) ZScore(key, member string) int {
	return int(rdb.ZScore(ctx, key, member).Val())
}

// [有序集合]中 key 中指定字段的整数值加上增量 incr
func (*_redis) ZincrBy(key, member string, incr float64) {
	err := rdb.ZIncrBy(ctx, key, incr, member).Err()
	if err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"ZincrBy: ", zap.Error(err))
		panic(err)
	}
}

// 指定成员的剩余时间
func (*_redis) TTL(key string) time.Duration {
	return rdb.TTL(ctx, key).Val()
}
