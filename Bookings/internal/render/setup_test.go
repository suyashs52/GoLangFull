package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/suyashs52/golang/bookings/internal/config"
	"github.com/suyashs52/golang/bookings/internal/models"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	//what am i going to put in the session
	gob.Register(models.Reservation{})

	testApp.InProduction = false
	//true when in production

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false //app.InProduction

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct {
}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}
func (tw *myWriter) WriteHeader(i int) {

}
func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
