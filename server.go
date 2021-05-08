package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

type PlayerServer struct {
	store PlayerStore
}

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.processWin(w)
	case http.MethodGet:
		s.ShowScore(w, r)
	}
}

func (s *PlayerServer) processWin(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

func (s *PlayerServer) ShowScore(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	score := s.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

//func PlayerServer(w http.ResponseWriter, r *http.Request) {
//	player := r.URL.Path[len("/players/"):]
//	fmt.Fprint(w, GetPlayerScore(player))
//}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}
	if name == "Floyd" {
		return "10"
	}
	return ""
}
