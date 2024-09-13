package main

type Message struct {
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
}
