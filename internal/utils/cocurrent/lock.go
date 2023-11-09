package cocurrent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Yeuoly/coshub/internal/db"
	"github.com/Yeuoly/coshub/internal/utils/math"
)

type lock struct {
	key string
}

func GetDistributedLock() *lock {
	return &lock{
		key: math.RandomString(32),
	}
}

func (l *lock) TryLock() bool {
	// get redis connection
	cacheManager := db.NewCacheManager()
	err := cacheManager.SetNX(cacheManager.GetRedisKey("distribute_lock", l.key), 1, 0)
	return err == nil
}

func (l *lock) Lock(context context.Context) error {
	cacheManager := db.NewCacheManager()
	key := cacheManager.GetRedisKey("distribute_lock", l.key)
	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()
	for {
		select {
		case <-context.Done():
			return fmt.Errorf("lock timeout")
		case <-ticker.C:
			err := cacheManager.SetNX(key, 1, 0)
			if err == nil {
				return nil
			}
			if !errors.Is(err, db.ErrCacheUsedAlready) {
				return err
			}
		}
	}
}

func (l *lock) Unlock() {
	cacheManager := db.NewCacheManager()
	cacheManager.Del(cacheManager.GetRedisKey("distribute_lock", l.key))
}

/*
return false when locked, true when unlocked
*/
type highGranularityMutex struct {
	key string
}

func GetRandomHighGranularityMutex() *highGranularityMutex {
	return &highGranularityMutex{
		key: math.RandomString(32),
	}
}

func GetHighGranularityMutex(key string) *highGranularityMutex {
	return &highGranularityMutex{
		key: key,
	}
}

/*
TryLock method lock the thread owned by id,
returns true if Locked Successfully
*/
func (c *highGranularityMutex) TryLock(id int) bool {
	cacheManager := db.NewCacheManager()
	return cacheManager.SetNX(cacheManager.GetRedisKey("high_granularity_mutex", fmt.Sprintf("%s_%d", c.key, id)), 1, 0) == nil
}

/*
Lock method lock the thread owned by id,
returns error if Locked Failed
*/
func (c *highGranularityMutex) Lock(id int, context context.Context) error {
	cacheManager := db.NewCacheManager()
	key := cacheManager.GetRedisKey("high_granularity_mutex", fmt.Sprintf("%s_%d", c.key, id))
	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()
	for {
		select {
		case <-context.Done():
			return fmt.Errorf("lock timeout")
		case <-ticker.C:
			err := cacheManager.SetNX(key, 1, 0)
			if err == nil {
				return nil
			}
			if !errors.Is(err, db.ErrCacheUsedAlready) {
				return err
			}
		}
	}
}

func (c *highGranularityMutex) Unlock(id int) {
	cacheManager := db.NewCacheManager()
	cacheManager.Del(cacheManager.GetRedisKey("high_granularity_mutex", fmt.Sprintf("%s_%d", c.key, id)))
}
