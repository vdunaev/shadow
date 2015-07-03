package system

import (
	"github.com/GeertJohan/go.rice"
	"github.com/kihamo/shadow/service/frontend"
)

func (s *SystemService) GetTemplateBox() *rice.Box {
	return rice.MustFindBox("../system/templates")
}

func (s *SystemService) GetFrontendMenu() *frontend.FrontendMenu {
	menu := []*frontend.FrontendMenu{
		&frontend.FrontendMenu{
			Name: "Config",
			Url:  "/system/config",
		},
	}

	if s.Application.HasResource("logger") {
		menu = append(menu, &frontend.FrontendMenu{
			Name: "Logs",
			Url:  "/system/logs",
		})
	}

	if s.Application.HasResource("tasks") {
		menu = append(menu, &frontend.FrontendMenu{
			Name: "Tasks",
			Url:  "/system/tasks",
		})
	}

	return &frontend.FrontendMenu{
		Name:    "System",
		SubMenu: menu,
	}
}

func (s *SystemService) SetFrontendHandlers(router *frontend.Router) {
	router.GET(s, "/system/config", &ConfigHandler{})

	if s.Application.HasResource("logger") {
		router.GET(s, "/system/logs", &LogsHandler{})
	}

	if s.Application.HasResource("tasks") {
		router.GET(s, "/system/tasks", &TasksHandler{})
	}
}
