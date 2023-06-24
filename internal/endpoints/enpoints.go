package endpoints

import (
	"fmt"
	"github.com/Bananad47/go_test_task_echo/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

// TimeHandler возвращает кол-во дней до 01.01.2025
func TimeHandler(c echo.Context) error {
	daysCountString := fmt.Sprintf("%d days left until January 1, 2025", service.DaysLeft())
	return c.String(http.StatusOK, daysCountString)
}
