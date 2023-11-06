package reqDto

/**
 * @ClassName authDto
 * @Description TODO
 * @Author khr
 * @Date 2023/8/1 14:15
 * @Version 1.0
 */
type Login struct {
	Email    string `json:"email" `
	Phone    string `json:"phone" `
	Password string `json:"password" binding:"required"`
	Method   string `json:"method" `
}

// 图形验证
type Captcha struct {
	Id      string `json:"id,omitempty" `
	Content string `json:"content,omitempty" `
}
