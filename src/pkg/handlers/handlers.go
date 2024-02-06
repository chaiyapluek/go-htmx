package handlers

import (
	"fmt"
	"strings"

	"github.com/chaiyapluek/go-htmx/src/pkg/models"
	pages "github.com/chaiyapluek/go-htmx/templates/pages"
	"github.com/labstack/echo/v4"
)

var additionalHeaderLog = map[string]interface{}{
	"referer": nil,
}

func LogHeader(headers map[string][]string) {
	for k, header := range headers {
		if strings.Contains(strings.ToLower(k), "hx") {
			fmt.Println(k, header)
		} else if _, ok := additionalHeaderLog[strings.ToLower(k)]; ok {
			fmt.Println(k, header)
		}
	}
}

// func RenderView(c echo.Context, view templ.Component) error {
// 	LogHeader(c.Request().Header)
// 	if c.Request().Header.Get("Hx-Request") == "true" {
// 		log.Println("HTMX Request")
// 		return view.Render(c.Request().Context(), c.Response())
// 	}
// 	log.Println("Not htmx request")
// 	fmt.Println("====================================")
// 	return templates.Layout().Render(c.Request().Context(), c.Response())
// }

func HomeGetHandler(c echo.Context) error {
	return pages.Home().Render(c.Request().Context(), c.Response())
}

func AboutGetHandler(c echo.Context) error {
	return pages.About(models.Info{
		Test:     "Test",
		RemoteIP: c.RealIP(),
	}).Render(c.Request().Context(), c.Response())
}

func ContactGetHandler(c echo.Context) error {
	return pages.Contact().Render(c.Request().Context(), c.Response())
}

func ReservationGetHandler(c echo.Context) error {
	return pages.Reservation().Render(c.Request().Context(), c.Response())
}
