package redis

import (
	"errors"
	"go.uber.org/zap"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600 * 10000
	scorePerVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("INVALID VOTE TIME")
)

func CreatePost(postID int64) error {
	pipeline := rdb.TxPipeline()

	zap.L().Debug("CreatePost in redis/vote")
	pipeline.ZAdd(getRedisKey(KeyPostTime), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	pipeline.ZAdd(getRedisKey(KeyPostScore), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID string, postID int64, value float64) error {
	postStr := strconv.FormatInt(postID, 10)
	zap.L().Debug("postStr in redis/vote.VoteForPost", zap.String("postStr", postStr))
	//zap.L().Debug("getRedisKey(KeyPostTime) in redis/vote.VoteForPost", zap.String("postStr", postStr))
	postTime := rdb.ZScore(getRedisKey(KeyPostTime), postStr).Val()
	zap.L().Debug("postTime in redis/vote.VoteForPost", zap.Float64("postTime", postTime))
	zap.L().Debug("time.Now().Unix() in redis/vote.VoteForPost", zap.Float64("time.Now().Unix()", float64(time.Now().Unix())))
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	ov := rdb.ZScore(getRedisKey(KeyPostVotedPre+postStr), userID).Val()

	// 更新：如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	//if value == ov {
	//	return ErrVoteRepeated
	//}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScore), op*diff*scorePerVote, postStr)

	// 3. 记录用户为该贴子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedPre+postStr), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedPre+postStr), redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err

}
