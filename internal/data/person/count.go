package person

import "github.com/tomazcx/rinha-backend-go/internal/data/protocol"

type CountPersonUseCase struct {
	repo protocol.IPersonRepository
}

func (uc *CountPersonUseCase) Execute() (int, error) {
	count, err := uc.repo.Count()

	if err != nil {
		return 0, err
	}

	return count, nil
}

func NewCountPersonUseCase(repo protocol.IPersonRepository) *CountPersonUseCase{
	return &CountPersonUseCase{repo:repo}
}
