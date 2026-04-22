package delivery

import (
	"fmt"
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
	req.ProductID = uint(idInt)
	resp, err := s.productSvc.ProductByID(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (s Server) AddProduct (c echo.Context) error {
	var req param.AddProductRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := s.productSvc.AddProduct(req)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, resp)

}

func (s Server) UpdateProduct (c echo.Context) error {
	var req param.UpdateProductRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	id := c.Param("id")
	idUint, _ := strconv.Atoi(id)
	req.ID = uint(idUint)

	resp, err := s.productSvc.UpdateProduct(req)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, resp)

}


func (s Server) AddCategory (c echo.Context) error {
	var req param.AddCategoryRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := s.productSvc.AddCategory(req)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, resp)

}


func (s Server) DeleteProduct (c echo.Context) error {
	id := c.Param("id")
	idUint, _ := strconv.Atoi(id)
	req := param.ProductByIDRequest{ProductID: uint(idUint)}

	err := s.productSvc.DeleteProduct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusNoContent, "")
}