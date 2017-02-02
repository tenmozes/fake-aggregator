package server

import "time"

type Option func(*Server)

func WithDelay(delay int) Option {
	return func(s *Server) {
		s.maxDelay = delay
	}
}

func WithRandomFactor(factor int) Option  {
	return func(s *Server) {
		s.randomFactor = factor
	}
}

func WithDeadline(d time.Duration) Option {
	return func(s *Server) {
		s.deadline = d
	}
}