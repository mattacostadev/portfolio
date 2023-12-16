package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomePageData struct {
	Contact interface{}
	About   interface{}
}

func Home(c echo.Context) error {
	bio :=
		`
		Hello! My name is Matt Acosta. I am a (self-identified) web developer/(parent-identified) human being 
		located in the DFW Metroplex, TX, USA.
		`

	data := HomePageData{
		Contact: map[string]interface{}{
			"Name":   "Matt Acosta",
			"Email":  "matthew.t.acost@gmail.com",
			"Github": "mattacostadev",
		},
		About: map[string]interface{}{
			"Title":   "About",
			"Content": bio,
		},
	}

	return c.Render(http.StatusOK, "index", data)
}
