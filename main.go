package main

import (

	"net/http"

	"github.com/stevehebert/go-getting-started/controllers"

)





func main() { 
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}