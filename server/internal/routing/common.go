package routing

import (
	"ttms/internal/model/config"
	limit "ttms/internal/pkg/limiter/api"

	"golang.org/x/time/rate"
)

func GetLimiters(buckets []config.Bucket) limit.RateLimiter {
	limiters := make([]limit.RateLimiter, len(buckets))
	for i := range limiters {
		limiters[i] = rate.NewLimiter(limit.Per(buckets[i].Count, buckets[i].Duration), buckets[i].Burst)
	}
	return limit.MultiLimiter(limiters...)
}
