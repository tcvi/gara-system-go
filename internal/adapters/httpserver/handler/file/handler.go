package file

import "garasystem/internal/core/ports"

type Handler struct {
	service ports.FileService
}

func NewHandler(fileService ports.FileService) *Handler {
	return &Handler{fileService}
}
