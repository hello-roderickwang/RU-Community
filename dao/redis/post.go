package redis

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"strconv"
	"time"
	"web_app/models"
)

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	zap.L().Debug("Params in GetPostINsInOrder()", zap.String("key", key), zap.Int64("start", start), zap.Int64("end", end))
	return rdb.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTime)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScore)
	}
	return getIDsFormKey(key, p.Page, p.Size)
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedPre + id)
	//	v := rdb.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}

	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedPre + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

func GetCommunityPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTime)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScore)
	}

	key := orderKey + strconv.Itoa(int(p.CommunityID))
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))
	if rdb.Exists(key).Val() < 1 {
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey)
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return getIDsFormKey(key, p.Page, p.Size)
}
