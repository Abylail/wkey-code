package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"wkey-core/src/data/models"
	"wkey-core/src/events/admin_session"
	"wkey-core/src/events/client_session"
)

var (
	adminSession  *admin_session.Event
	clientSession *client_session.Event
)

func AdminAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// получаем токен с headers
		token := ctx.Request().Header.Get("token")

		// если токена нет, возвращаем ошибку
		if token == "" {
			errorObject := &models.Response{
				Status:       "ERROR",
				Message:      "Произошла ошибка",
				InnerMessage: "Token required",
			}
			return ctx.JSON(http.StatusUnauthorized, errorObject)
		}

		// проверяем активна ли сессия
		active, err := adminSession.Active(token)
		if err != nil {
			errorObject := &models.Response{
				Status:       "ERROR",
				Message:      "Произошла ошибка",
				InnerMessage: err.Error(),
			}
			return ctx.JSON(http.StatusUnauthorized, errorObject)
		}

		// если сессия не активна, возвращаем ошибку
		if !active {
			errorObject := &models.Response{
				Status:       "ERROR",
				Message:      "Произошла ошибка",
				InnerMessage: "Пользователь не авторизован",
			}
			return ctx.JSON(http.StatusUnauthorized, errorObject)
		}

		// запуск следующего middleware или уже action
		if err = next(ctx); err != nil {
			ctx.Error(err)
			return nil
		}

		return nil
	}
}

func ClientAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// do it, babe!

		if err := next(ctx); err != nil {
			ctx.Error(err)
			return nil
		}

		return nil
	}
}
