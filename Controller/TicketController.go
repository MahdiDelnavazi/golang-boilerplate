package Controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang-boilerplate/DTO/Request/Ticket"
	"golang-boilerplate/DTO/Response"
	Ticket2 "golang-boilerplate/DTO/Response/Ticket"
	"golang-boilerplate/Helper"
	"golang-boilerplate/Service"
	"net/http"
)

type TicketController struct {
	logger        *zap.SugaredLogger
	ticketService *Service.TicketService
}

func NewTicketController(logger *zap.SugaredLogger, ticketService *Service.TicketService) *TicketController {
	return &TicketController{logger: logger, ticketService: ticketService}
}

func (ticketControler *TicketController) CreateTicket(context *gin.Context) {
	var ticketRequest Ticket.CreateTicketRequest
	Helper.Decode(context.Request, &ticketRequest)

	//fmt.Println("before call controler", ticketRequest)
	ticketResponse, responseError := ticketControler.ticketService.CreateTicket(ticketRequest)

	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}

	// all ok
	// create general response
	response := Response.GeneralResponse{Error: false, Message: "ticket have been created", Data: Ticket2.CreateTicketResponse{UserName: ticketResponse.UserName, Message: ticketResponse.Message, Subject: ticketResponse.Subject}}
	context.JSON(http.StatusOK, gin.H{"response": response})
}
