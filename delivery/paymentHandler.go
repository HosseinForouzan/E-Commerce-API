package delivery

import (
	"strconv"
	"fmt"

	"github.com/HosseinForouzan/E-Commerce-API/param"
	"github.com/HosseinForouzan/E-Commerce-API/service/authservice"
	"github.com/labstack/echo/v4"
)

func (s Server) StartPayment(c echo.Context) error {
	var req param.PaymentRequest
	fmt.Println("SALAM")
	claims := c.Get("claims").(*authservice.Claims)
	fmt.Println("BYE", claims)

	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, "invalid order id")
	}

	req.OrderID = uint(orderID)
	req.UserID = claims.UserID

	

	url, err := s.paymentSvc.StartPayment(
		c.Request().Context(), req)

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, map[string]any{
		"payment_url": url,
	})
}

func (s Server) MockPage(c echo.Context) error {

	id := c.Param("paymentID")

	html := fmt.Sprintf(`
	<h1>Mock Payment Gateway</h1>

	<form method="POST" action="/payment/mock/%s/success">
		<button>Pay Success</button>
	</form>

	<form method="POST" action="/payment/mock/%s/fail">
		<button>Pay Fail</button>
	</form>
	`, id, id)

	return c.HTML(200, html)
}

func (s Server) Success(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("paymentID"))

	err := s.paymentSvc.Success(
		c.Request().Context(),
		uint(id),
	)

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.Redirect(
		302,
		"/payment/callback?status=success",
	)
}

func (s Server) Fail(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("paymentID"))

	err := s.paymentSvc.Fail(
		c.Request().Context(),
		uint(id),
	)

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.Redirect(
		302,
		"/payment/callback?status=failed",
	)
}

func (s Server) Callback(c echo.Context) error {

	status := c.QueryParam("status")

	if status == "success" {
		return c.HTML(200, "<h1>Payment Successful</h1>")
	}

	return c.HTML(200, "<h1>Payment Failed</h1>")
}