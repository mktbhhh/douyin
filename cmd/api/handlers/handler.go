package handlers

//
//// type ResponseOld struct {
//// 	StatusCode int32  `json:"status_code"`
//// 	StatusMsg  string `json:"status_msg,omitempty"`
//// }
//
//type Resp struct {
//	Code    int64  `json:"status_code"`
//	Message string `json:"status_msg,omitempty"`
//}
//
//// SendResponse pack response
//// func SendResponse(c *gin.Context, err error, data interface{}) {
//// 	Err := errno.ConvertErr(err)
//// 	c.JSON(http.StatusOK, Resp{
//// 		Code:    Err.ErrCode,
//// 		Message: Err.ErrMsg,
//// 	})
//// }
//
//type UserRegisterResponse struct {
//	Resp
//	UserId int64  `json:"user_id"`
//	Token  string `json:"token"`
//}
//
//type NoteParam struct {
//	Title   string `json:"title"`
//	Content string `json:"content"`
//}
//
//type UserRegisterParam struct {
//	UserName string `form:"username"`
//	PassWord string `form:"password"`
//}
//
//type UserLoginParam struct {
//	UserName string `form:"username"`
//	PassWord string `form:"password"`
//}
