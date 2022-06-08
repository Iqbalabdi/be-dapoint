package request

import "dapoint-api/business/content/spec"

type CreateContentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (req *CreateContentRequest) ToSpec() *spec.UpsertContentSpec {
	return &spec.UpsertContentSpec{
		Name:        req.Name,
		Description: req.Description,
	}
}
