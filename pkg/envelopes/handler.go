package envelopes

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

type App struct {
	HandlerLatest  *HandlerLatest
	HandlerDate    *HandlerDate
	HandlerAnalyze *HandlerAnalyze
}

type HandlerLatest struct {
	Repo Repository
}

type HandlerDate struct {
	Repo Repository
	Date string
}

type HandlerAnalyze struct {
	Repo Repository
}

func (h *HandlerLatest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.GetRatesLatest(w, r)
	default:
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
	}
}

func (h *HandlerDate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.GetRatesOnDate(w, r)
	default:
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
	}
}

func (h *HandlerAnalyze) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.GetRatesMinMaxAverage(w, r)
	default:
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
	}
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)
	if head == "rates" {
		if r.URL.Path == "/latest" {
			a.HandlerLatest.ServeHTTP(w, r)
			return
		} else if t, err := time.Parse("2006-01-02", r.URL.Path[1:]); err == nil {
			a.HandlerDate.Date = t.String()
			a.HandlerDate.ServeHTTP(w, r)
			return
		} else if r.URL.Path == "/analyze" {
			a.HandlerAnalyze.ServeHTTP(w, r)
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

func byteResponse(w http.ResponseWriter, code int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	w.Write(body)
}

func notFoundResponse(w http.ResponseWriter, code int, body string) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(code)

	io.WriteString(w, body)
}

func GetXML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Read body: %v", err)
	}

	return string(data), nil
}

func (h *HandlerLatest) GetRatesLatest(w http.ResponseWriter, r *http.Request) {
	envelope, err := h.Repo.GetEnvelope()
	if err != nil {
		log.Fatal(err, "get")
	}
	currencies := envelope.Cube.Cube[0]
	rates := make(map[string]string, len(currencies.Cube))
	for _, cube3 := range currencies.Cube {
		rates[cube3.Currency] = cube3.Rate
	}
	resp := ResponseLatest{Base: "EUR", Rates: rates}
	b, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err, "get")
	}
	byteResponse(w, http.StatusOK, b)
}

func (h *HandlerAnalyze) GetRatesMinMaxAverage(w http.ResponseWriter, r *http.Request) {
	envelope, err := h.Repo.GetEnvelope()
	if err != nil {
		log.Fatal(err, "get")
	}
	analyzedRates := AnalyzedRates{Min: 99999999999.99999, Max: 0.00, Avg: 0.00}
	currencies := envelope.Cube.Cube[0]
	rates := make(map[string]AnalyzedRates, len(currencies.Cube))
	for _, cube3 := range currencies.Cube {
		rates[cube3.Currency] = analyzedRates
	}
	allRates := envelope.Cube
	for _, cube2 := range allRates.Cube {
		for _, cube3 := range cube2.Cube {
			analyzedRates = AnalyzedRates{Min: rates[cube3.Currency].Min, Max: rates[cube3.Currency].Max, Avg: rates[cube3.Currency].Avg}
			minFloat, err := strconv.ParseFloat(cube3.Rate, 64)
			if err != nil {
				log.Fatal(err)
			}
			maxFloat, err := strconv.ParseFloat(cube3.Rate, 64)
			if err != nil {
				log.Fatal(err)
			}
			avgFloat, err := strconv.ParseFloat(cube3.Rate, 64)
			if err != nil {
				log.Fatal(err)
			}
			if rates[cube3.Currency].Min > minFloat {
				analyzedRates.Min = minFloat
			}
			if rates[cube3.Currency].Max < maxFloat {
				analyzedRates.Max = maxFloat
			}
			if err != nil {
				log.Fatal(err)
			}
			analyzedRates.Avg += avgFloat
			rates[cube3.Currency] = analyzedRates
		}
	}
	for _, cube3 := range currencies.Cube {
		analyzedRates = rates[cube3.Currency]
		analyzedRates.Avg = analyzedRates.Avg / float64(len(currencies.Cube))
		rates[cube3.Currency] = analyzedRates
	}
	resp := ResponseAnalyze{Base: "EUR", Rates: rates}
	b, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err, "get")
	}
	byteResponse(w, http.StatusOK, b)
}

func (h *HandlerDate) GetRatesOnDate(w http.ResponseWriter, r *http.Request) {
	YMD := h.Date[0:10]
	var cubeRatesForDate []Cube3
	envelope, err := h.Repo.GetEnvelope()
	if err != nil {
		log.Fatal(err, "get")
	}
	currencies := envelope.Cube
	match := false
	for _, cube2 := range currencies.Cube {
		if cube2.Time == YMD {
			cubeRatesForDate = cube2.Cube
			match = true
			break
		}
	}
	if match != true {
		notFoundResponse(w, http.StatusNotFound, "No info for "+YMD+" date Found")
	} else {
		rates := make(map[string]string, len(cubeRatesForDate))
		for _, cube := range cubeRatesForDate {
			rates[cube.Currency] = cube.Rate
		}
		resp := ResponseLatest{Base: "EUR", Rates: rates}
		b, err := json.Marshal(resp)
		if err != nil {
			log.Fatal(err, "get")
		}
		byteResponse(w, http.StatusOK, b)
	}
}

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
