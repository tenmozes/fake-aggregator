package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/tenmozes/fake-aggregator/aggregator"
)

type Server struct {
	mapper       aggregator.MapperInterface
	minDelay     int
	maxDelay     int
	randomFactor int
	deadline     time.Duration
	port         string
	baseURL      string

	mux *http.ServeMux
}

func NewServer(m aggregator.MapperInterface, options ...Option) *Server {
	s := &Server{
		mapper: m,
		mux:    http.NewServeMux(),
	}
	for _, o := range options {
		o(s)
	}
	for path, aggr := range m.Mappings() {
		s.mux.HandleFunc(path, s.GetMiddleware(s.DelayedMiddleware(s.Aggregator(aggr))))
	}
	s.mux.HandleFunc("/numbers", s.Numbers)
	return s
}

func (Server) GetMiddleware(next http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.NotFound(rw, r)
			return
		}
		next.ServeHTTP(rw, r)
	}
}

func (s *Server) DelayedMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		delay := rand.Intn(s.maxDelay)
		if delay < s.minDelay {
			delay = s.minDelay
		}
		log.Printf("%s delay for %d ms", r.RequestURI, delay)
		time.Sleep(time.Duration(delay) * time.Millisecond)
		next.ServeHTTP(rw, r)
	}
}

func (s *Server) Aggregator(a aggregator.AggregatorInterface) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		r.Header.Add("Content-Type", "application/json")
		result, err := a.Numbers(rand.Intn(s.randomFactor))
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		j, err := json.Marshal(numberResponse{Numbers: result})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		rw.Write(j)
	}
}

func (s *Server) Numbers(rw http.ResponseWriter, r *http.Request) {
	resp := &numberResponse{
		Numbers: make([]int, 0),
	}
	defer func() {
		j, err := json.Marshal(resp)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Write(j)
	}()
	urls := r.URL.Query()["u"]
	if len(urls) == 0 {
		return
	}
	var wg sync.WaitGroup
	ctx, cancel := context.WithDeadline(r.Context(), time.Now().Add(s.deadline))
	results := make(chan []int)
	for _, raw := range urls {
		wg.Add(1)
		go func(raw string, r chan []int) {
			defer wg.Done()
			sequence, err := s.processURL(ctx, raw)
			if err != nil {
				log.Println(err)
				return
			}
			r <- sequence
		}(raw, results)
	}
	go func() {
		wg.Wait()
		cancel()
		close(results)
	}()
	for i := range results {
		resp.Numbers = append(resp.Numbers, i...)
	}
	resp.Compile()
	return
}

func (s *Server) processURL(ctx context.Context, raw string) ([]int, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}
	if _, ok := s.mapper.Mappings()[u.Path]; !ok {
		return nil, fmt.Errorf("Aggregator %q not found", u.Path)
	}
	req, err := http.NewRequest("GET", s.baseURL+u.Path, nil)
	req = req.WithContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := (http.DefaultClient).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	numbers := numberResponse{
		Numbers: make([]int, 0),
	}
	if err := json.NewDecoder(resp.Body).Decode(&numbers); err != nil {
		return nil, err
	}
	return numbers.Numbers, nil
}

func (s *Server) Run(port int) error{
	s.port = fmt.Sprintf(":%d", port)
	s.baseURL = "http://127.0.0.1" + s.port
	log.Printf("Server starts on %d", port)
	return http.ListenAndServe(s.port, s.mux)
}
