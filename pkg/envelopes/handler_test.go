package envelopes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lockwarr/Go-Server/common/pkg/cql"
	"github.com/gocql/gocql"
)

var app *App

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.ProtoVersion = 4
	cluster.Keyspace = "denislav"
	session, _ := cluster.CreateSession()

	repo := NewRepository(session, cql.NewKeyspaceBinder("denislav"))
	app = &App{&HandlerLatest{repo}, &HandlerDate{repo, ""}, &HandlerAnalyze{repo}}
}

func TestHandlerLatest_ServeHTTP(t *testing.T) {
	req, err := http.NewRequest("GET", "/rates/latest", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler := app.HandlerLatest
	rr := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *HandlerLatest
		args args
	}{
		{name: "test1", h: handler, args: args{w: rr, r: req}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}

//func TestHandlerDate_ServeHTTP(t *testing.T) {
//	//	req, err := http.NewRequest("GET", "/rates/2019-08-16", nil)
//	//if err != nil {
//	//		t.Fatal(err)
//	//	}
//	//handler := app.HandlerDate
//	//	rr := httptest.NewRecorder()
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name string
//		h    *HandlerDate
//		args args
//	}{
//		//{name: "test1", h: handler, args: args{w: rr, r: req}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.h.ServeHTTP(tt.args.w, tt.args.r)
//		})
//	}
//}
//
//func TestHandlerAnalyze_ServeHTTP(t *testing.T) {
//	req, err := http.NewRequest("GET", "/rates/analyze", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	handler := app.HandlerAnalyze
//	rr := httptest.NewRecorder()
//
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name string
//		h    *HandlerAnalyze
//		args args
//	}{
//		{name: "test1", h: handler, args: args{w: rr, r: req}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.h.ServeHTTP(tt.args.w, tt.args.r)
//		})
//	}
//}
//
func TestApp_ServeHTTP(t *testing.T) {
	req, err := http.NewRequest("GET", "/rates/latest", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		a    *App
		args args
	}{
		{name: "test1", a: app, args: args{w: rr, r: req}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}

func Test_byteResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	type args struct {
		w    http.ResponseWriter
		code int
		body []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test1", args: args{w: rr, code: 200, body: []byte("test")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			byteResponse(tt.args.w, tt.args.code, tt.args.body)
		})
	}
}

func Test_notFoundResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	type args struct {
		w    http.ResponseWriter
		code int
		body string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test1", args: args{w: rr, code: 200, body: "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notFoundResponse(tt.args.w, tt.args.code, tt.args.body)
		})
	}
}

func TestGetXML(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test1", args: args{url: "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetXML(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHandlerLatest_GetRatesLatest(t *testing.T) {
	req, err := http.NewRequest("GET", "/rates/latest", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler := app.HandlerLatest
	rr := httptest.NewRecorder()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *HandlerLatest
		args args
	}{
		{name: "test1", h: handler, args: args{w: rr, r: req}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetRatesLatest(tt.args.w, tt.args.r)
		})
	}
}

//func TestHandlerAnalyze_GetRatesMinMaxAverage(t *testing.T) {
//	req, err := http.NewRequest("GET", "/rates/analyze", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	handler := app.HandlerAnalyze
//	rr := httptest.NewRecorder()
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name string
//		h    *HandlerAnalyze
//		args args
//	}{
//		{name: "test1", h: handler, args: args{w: rr, r: req}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.h.GetRatesMinMaxAverage(tt.args.w, tt.args.r)
//		})
//	}
//}
//
//func TestHandlerDate_GetRatesOnDate(t *testing.T) {
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name string
//		h    *HandlerDate
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.h.GetRatesOnDate(tt.args.w, tt.args.r)
//		})
//	}
//}

func TestShiftPath(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name     string
		args     args
		wantHead string
		wantTail string
	}{
		{name: "test1", args: args{p: "test/1"}, wantHead: "test", wantTail: "/1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHead, gotTail := ShiftPath(tt.args.p)
			if gotHead != tt.wantHead {
				t.Errorf("ShiftPath() gotHead = %v, want %v", gotHead, tt.wantHead)
			}
			if gotTail != tt.wantTail {
				t.Errorf("ShiftPath() gotTail = %v, want %v", gotTail, tt.wantTail)
			}
		})
	}
}
