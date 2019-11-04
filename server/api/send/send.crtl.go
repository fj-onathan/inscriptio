package send

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	"html/template"
	"inscriptio/database/models"
	"inscriptio/libraries/common"
	"net/http"
	"net/smtp"
	"os"
)

// Global variables envroiment
var mailDefault = os.Getenv("MAIL_DEFAULT")
var mailPassword = os.Getenv("MAIL_PASSWORD")
var mailHost = os.Getenv("MAIL_HOST")
var mailSmtpPort = os.Getenv("MAIL_SMTP_PORT")

// Smtp
var auth smtp.Auth

func normal(c *gin.Context) {
	// testing
	type RequestBody struct {
		Name    string `json:"name" binding:"required"`
		Email   string `json:"email"  binding:"required"`
		Phone   int64  `json:"phone"`
		Subject string `json:"subject"`
		Message string `json:"message" binding:"required"`
	}
	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   "Complete all required fields.",
			},
		)
		return
	}
	auth = smtp.PlainAuth("", mailDefault, mailPassword, mailHost)
	wd, _ := os.Getwd()
	templateData := struct {
		Name    string
		Email   string
		Phone   int64
		Subject string
		Message string
	}{
		Name:    body.Name,
		Email:   body.Email,
		Phone:   body.Phone,
		Subject: body.Subject,
		Message: body.Message,
	}
	code := c.Query("code")
	company := Company(c, code)
	if company == nil {
		return
	}
	r := NewRequest([]string{body.Email}, body.Name, company["name"], company["email"], body.Subject, "Contacto via formulário web")
	if err := r.ParseTemplate(wd+"/html/email/normal.html", templateData); err == nil {
		ok, _ := r.SendEmail()
		fmt.Println(ok)
	}
	c.JSON(http.StatusOK, common.JSON{
		"success": true,
		"data":    "message",
	})
}

//Company struct
func Company(context *gin.Context, code string) map[string]string {
	db := context.MustGet("db").(*gorm.DB)
	db = context.MustGet("db").(*gorm.DB)
	var company models.Company
	if err := db.Where("code = ?", code).First(&company).Error; err != nil {
		context.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"success": false,
				"error":   "Company " + code + " not exists.",
			},
		)
		return nil
	}
	companyDetails := make(map[string]string)
	companyDetails["name"] = company.Name
	companyDetails["email"] = company.Email
	companyDetails["code"] = code
	return companyDetails
}

//Request struct
type Request struct {
	from    string
	companyName string
	companyEmail string
	client  string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, client, companyName, companyEmail, subject, body string) *Request {
	return &Request{
		to:      to,
		companyName: companyName,
		companyEmail: companyEmail,
		client:  client,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	from := "From: " + r.client + " <" + mailDefault + ">\r\n"
	to := "To: " + r.companyName + " <"+ r.companyEmail +">\r\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(from + to + subject + mime + "\n" + r.body)
	addr := mailHost + ":" + mailSmtpPort

	if err := smtp.SendMail(addr, auth, mailDefault, r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
