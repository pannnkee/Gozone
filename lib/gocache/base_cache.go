package gocache

import (
	"github.com/astaxie/beego"
	cache "github.com/patrickmn/go-cache"
	"time"
)

// 基本的缓存
type BaseCache struct{}

// 所有缓存
var AllCaches *cache.Cache

//	创建一个具有给定默认过期时间和清除间隔的新缓存
//	如果过期时间小于1(或NoExpiration)，则缓存中的项永远不会过期(默认情况下)，必须手动删除。
//	如果清理间隔小于1，则在调用c.DeleteExpired()之前不会从缓存中删除过期的项。
func (this *BaseCache) newCache() {
	defaultTime := beego.AppConfig.DefaultInt("go_cache::defaultHour", 48)
	timeInt := time.Duration(defaultTime)
	AllCaches = cache.New(timeInt*time.Hour, 60*time.Second)
}

// 从缓存中删除key的对应项。如果key不在缓存中，则不执行任何操作。
// @param key 缓存键值
func (this *BaseCache) Del(key string) {
	if len(key) == 0 {
		return
	}
	if AllCaches == nil {
		this.newCache()
	}
	AllCaches.Delete(key)
}

// 从缓存中获取key的对应项。返回项或nil，以及是否找到键的bool
// @param key 缓存键值
// @return data 缓存值
// @return isExist 键值为key的value是否存在
func (this *BaseCache) Get(key string) (data interface{}, isExist bool) {
	if len(key) == 0 {
		return
	}
	if AllCaches == nil {
		this.newCache()
	}
	data, isExist = AllCaches.Get(key)
	return
}

// 向缓存中添加key的对应项或替换现有key的对应项。
// 如果持续时间为0 (DefaultExpiration)，则使用缓存的默认过期时间。如果是-1 (NoExpiration)，则该条目永远不会过期。
// @param key 缓存键值
// @param data 替换现有key的值
func (this *BaseCache) Set(key string, data interface{}) {
	if len(key) == 0 {
		return
	}
	if AllCaches == nil {
		this.newCache()
	}
	AllCaches.Set(key, data, cache.DefaultExpiration)
}

//	清空所有设置的缓存
func (this *BaseCache) DelAll() {
	if AllCaches == nil {
		this.newCache()
	}
	AllCaches.Flush()
}
