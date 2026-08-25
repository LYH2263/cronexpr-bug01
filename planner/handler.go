package planner

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/LYH2263/go-cronexpr/cronexpr"
)

// API implements planner HTTP endpoints.
type API struct {
	DefaultTZ *time.Location
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/validate":
		a.handleValidate(w, r)
	case "/api/next":
		a.handleNext(w, r)
	case "/api/describe":
		a.handleDescribe(w, r)
	default:
		http.NotFound(w, r)
	}
}

type reqBody struct {
	Spec     string `json:"spec"`
	From     string `json:"from"`
	Timezone string `json:"timezone"`
	Locale   string `json:"locale"`
	Count    int    `json:"count"`
}

func (a *API) handleValidate(w http.ResponseWriter, r *http.Request) {
	var body reqBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := cronexpr.Validate(body.Spec)
	writeJSON(w, map[string]any{"ok": err == nil, "error": errString(err)})
}

func (a *API) handleNext(w http.ResponseWriter, r *http.Request) {
	var body reqBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	loc := a.loc(body.Timezone)
	e, err := cronexpr.ParseWithOptions(body.Spec, cronexpr.Options{Location: loc})
	if err != nil {
		writeJSON(w, map[string]any{"error": err.Error()})
		return
	}
	from := time.Now()
	if body.From != "" {
		if t, err := time.Parse(time.RFC3339, body.From); err == nil {
			from = t
		}
	}
	n := body.Count
	if n <= 0 {
		n = 5
	}
	times, err := e.NextN(from, n)
	if err != nil {
		writeJSON(w, map[string]any{"error": err.Error()})
		return
	}
	rows := make([]string, len(times))
	for i, t := range times {
		rows[i] = t.Format(time.RFC3339)
	}
	writeJSON(w, map[string]any{"times": rows})
}

func (a *API) handleDescribe(w http.ResponseWriter, r *http.Request) {
	var body reqBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	loc := body.Locale
	if loc == "" {
		loc = "zh"
	}
	text, err := cronexpr.DescribeLocale(body.Spec, loc)
	writeJSON(w, map[string]any{"text": text, "error": errString(err)})
}

func (a *API) loc(tzName string) *time.Location {
	if tzName == "" {
		if a.DefaultTZ != nil {
			return a.DefaultTZ
		}
		return time.Local
	}
	loc, err := time.LoadLocation(tzName)
	if err != nil {
		return time.Local
	}
	return loc
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(v)
}

func errString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
