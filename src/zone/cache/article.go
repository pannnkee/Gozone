package cache

import (
	"gozone/library/gocache"
	"gozone/src/zone/dao"
	"gozone/src/zone/models"
	"fmt"
)

type ArticleCache struct{}

func init() {
	articleCache := new(ArticleCache)
	err := new(gocache.Helper).PushListCache(articleCache)
	if err != nil {
		panic(err)
	}
}

func (this *ArticleCache) CacheConfig() (cacheName string, needItem bool, itemKey string) {
	return "article_type", true, "article_type:%v"
}

func (this *ArticleCache) PrimaryKey(model interface{}) string {
	return fmt.Sprintf("%v", model.(*models.Article).Id)
}

func (this *ArticleCache) GetAllData() (data interface{}, err error) {
	data, err = new(dao.ArticleDao).GetAllData()
	return
}

func (this *ArticleCache) GetItemData(articleID int64) (data interface{}, err error) {
	article, err := new(dao.ArticleDao).Get(articleID)
	if err != nil {
		return nil, err
	}
	return article, nil
}
