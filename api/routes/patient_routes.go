package routes

import (
	"libre-asi-api/auth"
	"libre-asi-api/handlers"
)

func patientRoutes() {
	server.Get("/patients/:id", auth.ValidateAccessToken, auth.ValidateAdministratorOrInterviewerRole, handlers.GetPatient)
	server.Get("/patients", auth.ValidateAccessToken, auth.ValidateAdministratorOrInterviewerRole, handlers.GetPatients)
	server.Post("/patients", auth.ValidateAccessToken, auth.ValidateAdministratorOrInterviewerRole, handlers.RegisterPatient)
	server.Patch("/patients/", auth.ValidateAccessToken, auth.ValidateAdministratorOrInterviewerRole, handlers.UpdatePatient)
	server.Delete("/patients/:id", auth.ValidateAccessToken, auth.ValidateAdministratorOrInterviewerRole, handlers.DeletePatient)

}
