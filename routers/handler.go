package routers

import (
	"net/http"

	control "github.com/thompsonmss/minificador-url-golang-with-gorm/controllers"
)

func Handler() {

	http.HandleFunc("/index", control.InitHomeController)
	http.HandleFunc("/registerhash", control.RegisterHashHomeController)

	http.HandleFunc("/", control.RedirectController)

}
