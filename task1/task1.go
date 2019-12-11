package task1

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type server struct {
}

type Query struct {
	Key   interface{} `json:"key"`
	Value interface{} `json:"value"`
}

type Headers struct {
	Accept    []string `json:"Accept"`
	UserAgent []string `json:"User-Agent"`
}

type Response struct {
	Host       string  `json:"host"`
	UserAgent  string  `json:"user_agent"`
	RequestURI string  `json:"request_uri"`
	Headers    Headers `json:"headers"`
	Querys     []Query `json:"querys"`
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var qr []Query
	for key, value := range r.URL.Query() {
		var res interface{} = value
		if len(value) == 1 {
			res = value[0]
		}
		qr = append(qr, Query{
			Key:   key,
			Value: res,
		})
	}
	resp := Response{
		Host:       r.Host,
		UserAgent:  r.UserAgent(),
		RequestURI: r.RequestURI,
		Headers: Headers{
			Accept:    strings.Split(r.Header.Get("Accept"), ","),
			UserAgent: []string{r.Header.Get("User-Agent")},
		},
		Querys: qr,
	}

	rJson, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(rJson)
}

func StartServer() {
	s := server{}
	if err := http.ListenAndServe(":5050", &s); err != nil {
		log.Panic(err)
	}
}

func main() {
	StartServer()
}
