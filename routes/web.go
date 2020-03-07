package routes

import (
	"html/template"
	"net/http"
	// "io/ioutil"
	"io"
	"github.com/labstack/echo"
	// "fmt"
	"portofolio-go/src"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func Web() *echo.Echo {

	// echo.NotFoundHandler = func(c echo.Context) error {
 //        // render your 404 page
 //        return c.String(http.StatusNotFound, "not found page")
 //    }
    echo.NotFoundHandler = func(c echo.Context) error {
        user_input :=     c.Request().URL    //  http.URL
        msg :=  "page not found bro : " + user_input.String()
       // render your 404 page
       return c.String(http.StatusNotFound,  msg )
   } 

	e := echo.New()
	e.Debug = true
	renderer := &TemplateRenderer{
    	templates: template.Must(template.ParseFiles(
    		"frontend/templates/header.html",
    		"frontend/templates/footer.html",
    		"frontend/index.html",
    		"frontend/about.html",
    	)),
  	}
  	e.Renderer = renderer

	// home
	e.GET("/home", controllers.Home)

	e.GET("/", controllers.Home)

	e.GET("/about", controllers.About)

	return e
}

