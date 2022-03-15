package Service

import (
	"go.uber.org/zap"
	"golang-boilerplate/DTO/Request/Ticket"
	"golang-boilerplate/DTO/Request/User"
	Ticket2 "golang-boilerplate/DTO/Response/Ticket"
	"golang-boilerplate/Repository"
)

type TicketService struct {
	ticketRepository *Repository.TicketRepository
	userService      *UserService
	logger           *zap.SugaredLogger
}

func NewTicketService(logger *zap.SugaredLogger, userService *UserService, ticketRepository *Repository.TicketRepository) *TicketService {
	return &TicketService{logger: logger, userService: userService, ticketRepository: ticketRepository}
}

func (ticketService TicketService) CreateTicket(createTicketRequest Ticket.CreateTicketRequest) (Ticket2.CreateTicketResponse, error) {
	// validate username len and not empty
	validationError := ValidationCheck(createTicketRequest)

	if validationError != nil {
		return Ticket2.CreateTicketResponse{}, validationError
	}
	user, userError := ticketService.userService.GetUser(User.GetUserRequest{UserName: createTicketRequest.UserName})
	if userError != nil {
		return Ticket2.CreateTicketResponse{}, userError
	}

	createTicketRequest.UserId = user.UserId
	ticket, ticketError := ticketService.ticketRepository.Create(createTicketRequest)
	if ticketError != nil {
		return Ticket2.CreateTicketResponse{}, ticketError
	}
	// we need a transformer
	return Ticket2.CreateTicketResponse{UserId: ticket.UserId, Subject: ticket.Subject, Message: ticket.Message}, nil
}
