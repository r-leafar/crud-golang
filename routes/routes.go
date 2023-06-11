package routes

import (
	"net/http"

	"github.com/r-leafar/controller"
)

func CarregaRotas() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Delete)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/update", controller.Update)

	http.HandleFunc("/MapTest", controller.MapTest)
}
