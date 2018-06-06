package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/mclark4386/buffalo_helpers"
	"github.com/rs/cors"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/packr"
	"mmgitl.mattclark.guru/Anchorspace/dashboard/models"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_dashboard_session",
			PreWares:    []buffalo.PreWare{buffalo_helpers.AutoSetContentType},
		})
		// Automatically redirect to SSL
		//		app.Use(forceSSL())

		if ENV == "development" {
			c := cors.New(cors.Options{
				AllowedOrigins:   []string{"*"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
				AllowedHeaders:   []string{"*"},
				AllowCredentials: true,
				ExposedHeaders:   []string{"Access-Token"},
			})

			app.PreWares = append(app.PreWares, c.Handler)

			app.Use(middleware.ParameterLogger)
		}

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		//		if ENV == "production" {
		app.Use(csrf.New)
		//		}

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		// Setup and use translations:
		app.Use(translations())

		app.GET("/", HomeHandler)
		app.Use(SetCurrentUser)
		app.Use(Authorize)
		app.GET("/signin", AuthNew)
		app.POST("/signin", AuthCreate)
		app.DELETE("/signout", AuthDestroy)
		usr := &UsersResource{}
		listing := &ListingsResource{}
		neighborhood := &NeighborhoodsResource{}
		app.Middleware.Skip(Authorize, HomeHandler, usr.New, usr.Create, AuthNew, AuthCreate, listing.List, listing.Show, neighborhood.List, neighborhood.Show)
		app.Resource("/users", UsersResource{&buffalo.BaseResource{}})

		app.Resource("/listings", ListingsResource{})

		app.Resource("/neighborhoods", NeighborhoodsResource{})
		app.Resource("/uses", UsesResource{})
		app.Resource("/terms", TermsResource{})

		app.Resource("/show_requests", ShowRequestsResource{})
		app.GET("/admin/panel", AdminPanel)
		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return ssl.ForceSSL(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
