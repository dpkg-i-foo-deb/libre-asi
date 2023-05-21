package routes

import (
	"libre-asi-api/cmd/handlers"
	"libre-asi-api/pkg/auth"
)

func interviewRoutes() {
	server.Post("/interviews/start", auth.ValidateAccessToken, auth.ValidateAdministratorOrInterviewerRole, handlers.StartInterview)
	server.Get("interviews", auth.ValidateAccessToken, auth.ValidateAdministratorOrInterviewerRole, handlers.GetInterviews)
	server.Get("/interviews/:id", auth.ValidateAccessToken, auth.ValidateAdministratorOrInterviewerRole, handlers.GetInterview)
	server.Post("/interviews/next-question", auth.ValidateAccessToken, auth.ValidateAdministratorOrInterviewerRole, handlers.NextQuestion)
	server.Get("/questions/:code", handlers.GetQuestion)
	server.Post("/interviews/answer/:id", auth.ValidateAccessToken, auth.ValidateAdministratorOrInterviewerRole, handlers.AnswerQuestion)
}
