package routes

import (
	"libre-asi-api/cmd/handlers"
	"libre-asi-api/pkg/auth"
)

func interviewRoutes() {
	server.Post("/interviews/create", auth.ValidateAccessToken, auth.ValidateInterviewerRole, handlers.StartInterview)
}
