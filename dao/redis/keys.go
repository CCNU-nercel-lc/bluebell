package redis

// redis key

const (
	KeyPrefix          = "bluebell:"   // key前缀
	KeyPostTimeZSet    = "post:time"   // 帖子及发帖时间
	KeyPostScoreZSet   = "post:score"  // 帖子及投票分数
	KeyPostVotedZSetPF = "post:voted:" // 记录用户及投票类型
	KeyCommunitySetPF  = "community:"  // 保存每个分区下帖子的id
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
