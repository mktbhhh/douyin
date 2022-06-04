package handlers

// type ResponseOld struct {
// 	StatusCode int32  `json:"status_code"`
// 	StatusMsg  string `json:"status_msg,omitempty"`
// }

type Response struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg,omitempty"`
}

// SendResponse pack response
// func SendResponse(c *gin.Context, err error, data interface{}) {
// 	Err := errno.ConvertErr(err)
// 	c.JSON(http.StatusOK, Response{
// 		Code:    Err.ErrCode,
// 		Message: Err.ErrMsg,
// 	})
// }

type UserRegisterResponse struct {
	Response
	user_id int64
	token   string
}

type NoteParam struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UserRegisterParam struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}
