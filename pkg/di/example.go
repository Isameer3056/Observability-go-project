package di

import (
	"github.com/Isameer3056/Observability-go-project/internals/domain/core/domain"
	"github.com/Isameer3056/Observability-go-project/internals/domain/usecase"
	"github.com/Isameer3056/Observability-go-project/internals/infra/controller"
	"github.com/Isameer3056/Observability-go-project/internals/infra/repository"
)

func NewExampleController() domain.ExampleController {
	exampleHttpRepository := repository.NewExampleHttpRepository()
	exampleUseCase := usecase.NewExampleHttpRepository(exampleHttpRepository)
	return controller.NewExampleController(exampleUseCase)
}
