package routes

import (
	"libre-asi-api/cmd/handlers"
	"libre-asi-api/pkg/auth"
)

func interviewRoutes() {
	server.Post("/interviews/create", auth.ValidateAccessToken, auth.ValidateInterviewerRole, handlers.StartInterview)
	server.Get("interviews", auth.ValidateAccessToken, auth.ValidateAdministratorOrInterviewerRole, handlers.GetInterviews)
	server.Post("/interviews/next-question", auth.ValidateAccessToken, auth.ValidateInterviewerRole, handlers.NextQuestion)
	server.Get("/questions/:code", handlers.GetQuestion)
}
