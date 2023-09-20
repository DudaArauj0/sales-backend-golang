package contactusecases

import (
	"context"

	personentity "github.com/willjrcom/sales-backend-go/internal/domain/person"
	contactdto "github.com/willjrcom/sales-backend-go/internal/infra/dto/contact"
	entitydto "github.com/willjrcom/sales-backend-go/internal/infra/dto/entity"
)

type Service struct {
	r personentity.ContactRepository
}

func NewService(c personentity.ContactRepository) *Service {
	return &Service{r: c}
}

func (s *Service) GetContactsBy(ctx context.Context, dto *contactdto.FilterContact) ([]contactdto.ContactOutput, error) {
	contact, err := dto.ToModel()

	if err != nil {
		return nil, err
	}

	contacts, err := s.r.GetContactsBy(ctx, contact)

	if err != nil {
		return nil, err
	}

	dtos := contactsToDtos(contacts)
	return dtos, nil
}

func (s *Service) GetContactById(ctx context.Context, dto *entitydto.IdRequest) (*contactdto.ContactOutput, error) {
	if contact, err := s.r.GetContactById(ctx, dto.ID.String()); err != nil {
		return nil, err
	} else {
		output := &contactdto.ContactOutput{}
		output.FromModel(contact)
		return output, nil
	}
}

func (s *Service) GetAllContacts(ctx context.Context) ([]contactdto.ContactOutput, error) {
	if contacts, err := s.r.GetAllContacts(ctx); err != nil {
		return nil, err
	} else {
		dtos := contactsToDtos(contacts)
		return dtos, nil
	}
}

func contactsToDtos(contacts []personentity.Contact) []contactdto.ContactOutput {
	dtos := make([]contactdto.ContactOutput, len(contacts))
	for i, contact := range contacts {
		dtos[i].FromModel(&contact)
	}

	return dtos
}
