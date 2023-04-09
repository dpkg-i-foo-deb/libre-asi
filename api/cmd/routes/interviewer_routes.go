package routes

import (
	"libre-asi-api/cmd/handlers"
	"libre-asi-api/pkg/auth"
)

func interviewerRoutes() {
	server.Get("/interviewers/:id", auth.ValidateAccessToken, auth.ValidateAdministratorOrInterviewerRole, handlers.GetInterviewer)
	server.Get("/interviewers", auth.ValidateAccessToken, auth.ValidateAdministratorOrInterviewerRole, handlers.GetInterviewers)

}
