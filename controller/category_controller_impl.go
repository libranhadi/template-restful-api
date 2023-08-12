package controller
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"template-restful-api/service"
)


type CategorControllerImpl struct {
	CategoryService service.CategoryService
}

func (c *CategorControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func (c *CategorControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	
}

func (c *CategorControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	
}

func (c *CategorControllerImpl) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	
}

func (c *CategorControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	
}
