package expiretime

import (
	"math/rand"
	"time"
)

func SetExpireTime() time.Duration {
	// 使用当前时间的UnixNano作为随机源的种子
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	// 设置随机过期时间范围
	minPExpire := 1 * time.Minute
	maxPExpire := 10 * time.Minute

	// 生成[minPExpire, maxPExpire)范围内的随机时间.Duration
	randomDuration := time.Duration(rng.Int63n(int64(maxPExpire.Minutes()-minPExpire.Minutes()+1))) * time.Minute
	randomDuration += minPExpire // 确保范围是从minPExpire开始

	return randomDuration
}
