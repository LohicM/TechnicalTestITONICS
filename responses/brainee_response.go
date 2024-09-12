package responses

import "TechnicalTestITONICS/models"

type BraineeResponse struct {
	Status int            `json:"status"`
	Data   models.Brainee `json:"data"`
}

type BraineesResponse struct {
	Status int              `json:"status"`
	Data   []models.Brainee `json:"data"`
}

type MessageResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
