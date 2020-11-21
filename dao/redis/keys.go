package redis

const (
	// sort of like namespace
	KeyPrefix       = "bluebell:"
	KeyPostTime     = "post:time"   // ZSet
	KeyPostScore    = "post:score"  // ZSet
	KeyPostVotedPre = "post:voted:" // ZSet, need to process later(+post_id)
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
