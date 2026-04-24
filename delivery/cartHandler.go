package delivery

import (
	"net/http"
	"strconv"

	"github.com/HosseinForouzan/E-Commerce-API/param"
	"github.com/HosseinForouzan/E-Commerce-API/service/authservice"
	"github.com/labstack/echo/v4"
)

func (s Server) AddItem(c echo.Context) error {
	claims := c.Get("claims").(*authservice.Claims)
	var req param.AddItemRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	req.UserID = claims.UserID

	err := s.cartSvc.AddItem(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, echo.Map{"message" : "item created."})
}

func (s Server) GetCart(c echo.Context) error {
	var req param.CartRequest
	claims := c.Get("claims").(*authservice.Claims)
	req.UserID = claims.UserID

	resp, err := s.cartSvc.GetCart(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (s Server) UpdateCart(c echo.Context) error {
	var req param.UpdateItemRequest
	claims := c.Get("claims").(*authservice.Claims)
	productID := c.Param("productId")
	ProductIDUint, _ := strconv.Atoi(productID)

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	req.UserID = claims.UserID
	req.ProductID = uint(ProductIDUint)
	
	err := s.cartSvc.UpdateItem(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "item updated."})

}

func (s Server) DeleteCart(c echo.Context) error {
	var req param.DeleteItemRequest
	claims := c.Get("claims").(*authservice.Claims)
	productID := c.Param("productId")
	ProductIDUint, _ := strconv.Atoi(productID)

	if err := c.Bind(&req); err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	req.UserID = claims.UserID
	req.ProductID = uint(ProductIDUint)

	err := s.cartSvc.DeleteItem(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusNoContent, "")
}

func (s Server) ClearCart(c echo.Context) error {
	claims := c.Get("claims").(*authservice.Claims)

	err := s.cartSvc.Clear(claims.UserID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusNoContent, "")

}