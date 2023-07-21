package errcode

var (
	ErrorGetCategoryList = NewError(20010001, "获取标签列表失败")
	ErrorCreateCategory  = NewError(20010002, "创建标签失败")
	ErrorUpdateCategory  = NewError(20010003, "更新标签失败")
	ErrorDeleteCategory  = NewError(20010004, "删除标签失败")
	ErrorCountCategory   = NewError(20010005, "统计标签失败")

	ErrorUploadFile = NewError(20030001, "上传文件失败")

	ErrorCreateArticle   = NewError(20010010, "创建文章失败")
	ErrorUpdateArticle   = NewError(20010011, "更新文章失败")
	ErrorListArticle     = NewError(20010012, "获取文章列表失败")
	ErrorDetailArticle   = NewError(20010013, "获取文章详情失败")
	ErrorDeleteArticle   = NewError(20010014, "删除文章失败")
	ErrorNotFoundArticle = NewError(20010015, "文章不存在")

	ErrorCreateAuthor   = NewError(20010016, "创建作者失败")
	ErrorUpdateAuthor   = NewError(20010017, "更新作者失败")
	ErrorDeleteAuthor   = NewError(20010018, "删除作者失败")
	ErrorListAuthor     = NewError(20010019, "获取作者列表失败")
	ErrorNotFoundAuthor = NewError(20010020, "作者不存在")

	ErrorCreateComment = NewError(20010021, "创建评论失败")
	ErrorDeleteComment = NewError(20010022, "删除评论失败")
)
