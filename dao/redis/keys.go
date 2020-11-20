package redis

const (
	// sort of like namespace
	KeyPrefix       = "whitebell:"
	KeyPostTime     = "post:time"   // ZSet
	KeyPostScore    = "post:score"  // ZSet
	KeyPostVotedPre = "post:voted:" // ZSet, need to process later(+post_id)
)
