package router

import (
	"net/http"

	"github.com/bryanljx/go-rest-api/internal/config"
	"github.com/bryanljx/go-rest-api/internal/handlers/notifications"
	"github.com/bryanljx/go-rest-api/internal/handlers/students"
	"github.com/bryanljx/go-rest-api/internal/handlers/teachers"
	"github.com/bryanljx/go-rest-api/internal/repositories"
	"github.com/bryanljx/go-rest-api/internal/service"
)

const (
	POST = "POST "

	routePrefix                   = "/api"
	registerEndPt                 = routePrefix + "/register"
	commonStudentsEndPt           = routePrefix + "/commonstudents"
	suspendEndPt                  = routePrefix + "/suspend"
	retrieveForNotificationsEndPt = routePrefix + "/retrievefornotifications"
)

func SetUpRoutes(s *service.Service, config *config.Config) *http.Handler {
	mux := http.NewServeMux()

	SetUpStudentRoutes(mux, s)
	SetUpTeacherRoutes(mux, s)
	SetUpNotificationRoutes(mux, s)

	return setUpMiddleware(mux, config)
}

func SetUpStudentRoutes(mux *http.ServeMux, s *service.Service) {
	studentRepo := repositories.NewStudentRepository(s.DB)
	studentHandler := students.New(studentRepo, s)
	mux.HandleFunc(POST+suspendEndPt, studentHandler.Suspend)
}

func SetUpTeacherRoutes(mux *http.ServeMux, s *service.Service) {
	// Set up teacher handlers
	teacherRepo := repositories.NewTeacherRepository(s.DB)
	teacherHandler := teachers.New(teacherRepo, s)
	mux.HandleFunc(POST+registerEndPt, teacherHandler.RegisterStudents)
	mux.HandleFunc(POST+commonStudentsEndPt, teacherHandler.CommonStudents)
}

func SetUpNotificationRoutes(mux *http.ServeMux, s *service.Service) {
	// Set up notification handlers
	studentRepo := repositories.NewStudentRepository(s.DB)
	teacherRepo := repositories.NewTeacherRepository(s.DB)
	notificationHandler := notifications.New(studentRepo, teacherRepo, s)
	mux.HandleFunc(POST+retrieveForNotificationsEndPt, notificationHandler.RetrieveForNotification)
}
