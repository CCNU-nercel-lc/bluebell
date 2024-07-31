package redis

import (
	"bluebell/models"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// 根据key到redis查询数据，范围为start~end
func getIDsFormKey(key string, page, size int64) (data []string, err error) {
	// 确定查询的起始和结束位置
	start := (page - 1) * size
	end := start + size - 1

	//  查询
	return rdb.ZRevRange(ctx, key, start, end).Result()
}

// GetPostIDsInOrder 根据制定的order方式查询帖子ids
func GetPostIDsInOrder(p *models.ParamPostList) (data []string, err error) {
	// 1. 根据用户请求中携带的order参数，决定返回的帖子是按分数排序的还是按创建时间排序的
	key := getRedisKey(KeyPostTimeZSet) // 默认按照时间排序
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	// 2. 确定查询的起始和结束位置
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	// 3. 查询
	return rdb.ZRevRange(ctx, key, start, end).Result()
}

// GetPostIDsInOrderByCommunity 根据社区查询帖子ids
func GetPostIDsInOrderByCommunity(p *models.ParamPostList) (data []string, err error) {
	orderkey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderkey = getRedisKey(KeyPostScoreZSet)
	}
	zap.L().Debug("orderkey = ", zap.String("orderkey", orderkey))
	ckey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID))) // 社区key
	zap.L().Debug("ckey = ", zap.String("ckey", ckey))
	key := orderkey + strconv.Itoa(int(p.CommunityID)) // 带score和社区id的key
	if rdb.Exists(ctx, key).Val() < 1 {                // 不存在，则计算出新的key（放在缓存内，超过60s后被删除）
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(ctx, key, &redis.ZStore{
			Aggregate: "MAX",
			Keys:      []string{ckey, orderkey}, // 计算ckey和orderkey的交集
			Weights:   []float64{1, 1},
		})
		pipeline.Expire(ctx, key, 60*time.Second) // 设置超时时间
		_, err = pipeline.Exec(ctx)
		if err != nil {
			zap.L().Error("pipeline.Exec(ctx) error", zap.Error(err))
			return nil, err
		}
	}

	// 存在，则根据key查询ids
	return getIDsFormKey(key, p.Page, p.Size)
}

// GetPostVoteData 根据ids查询每篇帖子的投票数
func GetPostVoteData(ids []string) (data []int64, err error) {
	// 使用pipeline减少RTT
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipeline.ZCount(ctx, key, "1", "1")
	}
	cmders, err := pipeline.Exec(ctx) // 返回包含赞成票数量的切片
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(ids))
	for _, cmder := range cmders {
		data = append(data, cmder.(*redis.IntCmd).Val())
	}
	return
}
