package mutube

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/workos/workos-go/pkg/workos"
)

func InitializeAuthRoutes(e *echo.Echo) {
	workosApiKey := os.Getenv("WORKOS_API_KEY")
	workosClientID := os.Getenv("WORKOS_CLIENT_ID")
	mutubeRedirectURI := "http://localhost:5173/browse"
	workosOrgID := "org_test_idp"
	sso.Configure(workosApiKey, workosClientID)

	e.GET("/authenticate", func(c echo.Context) error {
		url, err := sso.GetAuthorizationURL(sso.GetAuthorizationURLOpts{
			Organization: workosOrgID,
			RedirectURI:  mutubeRedirectURI,
		})
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.Redirect(http.StatusFound, url.String())
	})

	e.GET("/callback", func(c echo.Context) error {
		opts := sso.GetProfileAndTokenOpts{
			Code: c.QueryParam("code"),
		}

		profileAndToken, err := sso.GetProfileAndToken(c.Request().Context(), opts)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		profile := profileAndToken.Profile

		if profile.OrganizationID != workosOrgID {
			return c.String(http.StatusForbidden, "Unauthorized")
		}

		return c.Redirect(http.StatusSeeOther, "http://localhost:5173/browse")
	})
}
