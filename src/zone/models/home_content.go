package models

import "gozone/library/enum"

// 主页内容
type HomeContent struct {

	//base_right.html
	Tags         []*Tag          // 标签
	Links        []*Link         // 友情链接
	ArticleClass []*ArticleClass // 文章分类

	//base_top
	TopContent TopContent // 分类内容

	Articles    []*Article
	CarouselArticle []*Article
	SortType    int64
	ContentType enum.ContentType
}
