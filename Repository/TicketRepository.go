package Repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang-boilerplate/DTO/Request/Ticket"
	"golang-boilerplate/Entity"
)

type TicketRepository struct {
	logger   *zap.SugaredLogger
	database *sqlx.DB
}

func NewTicketRepository(logger *zap.SugaredLogger, database *sqlx.DB) *TicketRepository {
	return &TicketRepository{
		logger:   logger,
		database: database,
	}
}

func (TicketRepository *TicketRepository) Create(request Ticket.CreateTicketRequest) (Entity.Ticket, error) {
	ticket := Entity.Ticket{}
	queryError := TicketRepository.database.Get(&ticket, `SELECT * FROM newTicket($1 , $2 , $3 , $4, $5)`,
		request.UserId, request.Subject, request.Message, request.Image, request.Like)

	if queryError != nil {
		return Entity.Ticket{}, nil
	}
	return ticket, queryError

}
