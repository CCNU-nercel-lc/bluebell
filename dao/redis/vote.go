package redis

import (
	"errors"
	redis "github.com/redis/go-redis/v9"
	"math"
	"strconv"
	"time"
)

const (
	oneWeekInSeconds = 60 * 60 * 24 * 7
	scorePerVote     = 432 * 100000 // 每次投票的分数
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过！")
	ErrVoteRepeated   = errors.New("不允许重复投票！")
)

func VoteForPost(userID, postID string, value float64) (err error) {
	// 1. 判断投票限制
	// 去redis取帖子发布时间，投票时间限制为7天
	postTime := rdb.ZScore(ctx, getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		// 超过7天不能投票
		return ErrVoteTimeExpire
	}
	// 2. 更新帖子的分数
	// 先查之前的投票记录
	oldValue := rdb.ZScore(ctx, getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	var op float64
	// 如果跟之前的投票结果一致，则报错
	if oldValue == value {
		return ErrVoteRepeated
	}
	// 根据value与oldValue的大小计算操作算子为1还是-1
	if value > oldValue {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(oldValue - value)
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(ctx, getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	if err != nil {
		return err
	}
	// 3. 记录用户为该帖子投票的数据
	if value == 0 {
		// 如果是取消投票，则删除投票记录
		pipeline.ZRem(ctx, getRedisKey(KeyPostVotedZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(ctx, getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err = pipeline.Exec(ctx)
	return
}

func CreatePost(postID, communityID int64) error {
	pipeline := rdb.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(ctx, getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 帖子分数
	pipeline.ZAdd(ctx, getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 帖子ID加入到社区set
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID))) // 社区key ，前缀+社区ID
	pipeline.SAdd(ctx, cKey, postID)

	_, err := pipeline.Exec(ctx)
	return err
}
