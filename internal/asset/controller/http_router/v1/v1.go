package v1

import "github.com/iNgredie/file-service/internal/asset/usecase"

type Handlers struct {
	uc *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handlers {
	return &Handlers{
		uc: uc,
	}
}
