package delivery

import (
	"net/http"
	"strconv"

	"github.com/HosseinForouzan/E-Commerce-API/param"
	"github.com/labstack/echo/v4"
)

func (s Server) ListProducts(c echo.Context) error {
	var req param.ProductRequest
	categoryID := c.QueryParam("category_id")
	categoryIDInt, _ := strconv.Atoi(categoryID)
	searchParam := c.QueryParam("search")
	req.CategoryID = uint8(categoryIDInt)
	req.Search = searchParam

	resp, err := s.productSvc.Product(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (s Server) GetProductByID(c echo.Context) error {
	var req param.ProductByIDRequest
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	req.ProductID = uint8(idInt)
	resp, err := s.productSvc.ProductByID(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resp)
}