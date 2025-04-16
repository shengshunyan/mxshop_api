package forms

type SendSmsForm struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required,mobile"`
	Type   string `json:"type" form:"type" binding:"required,oneof=register login"`
}
