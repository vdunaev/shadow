package system

import (
	"github.com/kihamo/shadow/resource/mail"
	"github.com/kihamo/shadow/service/frontend"
	"gopkg.in/gomail.v2"
)

type MailHandler struct {
	frontend.AbstractFrontendHandler
}

func (h *MailHandler) Handle() {
	if h.IsPost() {
		message := gomail.NewMessage()
		message.SetHeader("Subject", h.Input.FormValue("subject"))
		message.SetHeader("To", h.Input.FormValue("to"))

		if h.Input.FormValue("type") == "html" {
			message.SetBody("text/html", h.Input.FormValue("message"))
		} else {
			message.SetBody("text/plain", h.Input.FormValue("message"))
		}

		resourceMail, _ := h.Application.GetResource("mail")
		if err := resourceMail.(*mail.Resource).SendAndReturn(message); err != nil {
			h.SendJSON(map[string]interface{}{
				"error": err.Error(),
			})
		} else {
			h.SendJSON(map[string]string{})
		}

		return
	}

	h.SetTemplate("mail.tpl.html")
	h.SetPageTitle("Mail")
	h.SetPageHeader("Send mail")
}
