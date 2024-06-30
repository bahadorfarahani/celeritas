package session

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alexedwards/scs/v2"
)

func TestSession_IntSession(t *testing.T) {
	c := &Session{
		CookieLifetime: "100",
		CookiePersist:  "true",
		CookieName:     "test",
		CookieDomain:   "localhost",
		SessionType:    "cookie",
	}

	var sm *scs.SessionManager
	ses := c.InitSession()

	var sesKind reflect.Kind
	var sesType reflect.Type

	rv := reflect.ValueOf(ses)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		fmt.Println("For loop::", rv.Kind(), rv.Type(), rv)
		sesKind = rv.Kind()
		sesType = rv.Type()
		rv = rv.Elem()
	}

	if !rv.IsValid() {
		t.Error("invalid type or kind; kind:", rv.Kind(), "type:", rv.Type())
	}

	if sesKind != reflect.ValueOf(sm).Kind() {
		t.Error("invalid kind; expected:", reflect.ValueOf(sm).Kind(), "got:", sesKind)
	}

	if sesType != reflect.ValueOf(sm).Type() {
		t.Error("invalid type; expected:", reflect.ValueOf(sm).Type(), "got:", sesType)
	}
}
