package handlers

import (
	"errors"
	"fmt"
	"github.com/vietthangc1/booking/pkg/config"
	"github.com/vietthangc1/booking/pkg/render"
	"github.com/vietthangc1/booking/pkg/models"
	"net/http"
)

var Repo *Repository

type Repository struct {
	app *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.app.Session.Put(r.Context(), "remote_IP", remoteIP)

	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
	// fmt.Fprintf(w, "Home page!")
	// if (err != nil) {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(fmt.Sprintf("Bytes: %d", n))
	// }
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	var stringMap = make(map[string]string)
	stringMap["message"] = "Hello again!"

	remote_IP := m.app.Session.GetString(r.Context(), "remote_IP")
	stringMap["remote_IP"] = remote_IP

	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
	// sum:=addValues(2,3)
	// fmt.Fprintf(w, fmt.Sprintf("About page!. 2 + 3 = %d", sum))
}

func (m *Repository) Divide(w http.ResponseWriter, r *http.Request) {
	var x, y float32
	x = 100
	y = 0

	result, err := divideValues(x, y)

	if err != nil {
		fmt.Fprint(w, "Error")
	} else {
		fmt.Fprint(w, fmt.Sprintf("%.2f divided by %.2f is %.2f", x, y, result))
	}
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{
	})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{
	})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
	})
}

// Get
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{
	})
}

// POST
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start_date")
	end := r.Form.Get("end_date")
	
	w.Write([]byte(fmt.Sprintf("Start: %s and End: %s", start,end)))
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{
	})
}


func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("Cannot divded by 0")
	}
	return x / y, nil
}
