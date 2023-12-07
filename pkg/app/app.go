package app

import (
	"github.com/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	TotalRows int64 `json:"total_rows"`
}

type CommonResponse struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Details interface{} `json:"details,omitempty"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToSuccessResponse(data interface{}) {
	if data == nil {
		data = make([]string, 0)
	}
	commonResponse := CommonResponse{
		Code: errcode.Success.Code(),
		Msg:  errcode.Success.Msg(),
		Data: data,
	}
	r.Ctx.JSON(http.StatusOK, commonResponse)
}

func (r *Response) ToResponseList(list interface{}, totalRows int64) {
	commonResponse := CommonResponse{
		Code: errcode.Success.Code(),
		Msg:  errcode.Success.Msg(),
		Data: map[string]interface{}{
			"list": list,
			"pager": Pager{
				Page:      GetPage(r.Ctx),
				PageSize:  GetPageSize(r.Ctx),
				TotalRows: totalRows,
			},
		},
	}
	r.Ctx.JSON(http.StatusOK, commonResponse)
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(http.StatusOK, response)
}
