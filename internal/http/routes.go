package http

const (
	orderBaseURL = "/orders"
)

func (s *Server) useRoutes() {
	v1 := s.echo.Group("/api/v1")

	v1.GET(orderBaseURL, s.queryController.getOrders)
	v1.GET(orderBaseURL+"/:id", s.queryController.getOrder)

	v1.POST(orderBaseURL, s.commandController.create)
	v1.PUT(orderBaseURL+"/pay"+"/:id", s.commandController.pay)
	v1.PUT(orderBaseURL+"/ship"+"/:id", s.commandController.ship)
	v1.PUT(orderBaseURL+"/cancel"+"/:id", s.commandController.cancel)
}
