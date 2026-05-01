package main

import (
	"slices"
	"sync"
)

// registry, who is current listening

// {
// 27: [broswerChannel, terminal Channel]
// }

type Registry struct {
	mu      sync.RWMutex // multiple goroutine reads, one goroutine write
	clients map[string][]*Client
}

func NewRegistry() *Registry {
	return &Registry{
		clients: make(map[string][]*Client),
	}
}

func (r *Registry) GetClientsByUserId(userID string) []*Client {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.clients[userID]
}

func (r *Registry) Register(userID string, cli *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()

	existingClients := r.clients[userID]
	r.clients[userID] = append(existingClients, cli)
}

func (r *Registry) Unregister(userID string, cli *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()

	existingClients := r.clients[userID]

	r.clients[userID] = slices.DeleteFunc(existingClients, func(c *Client) bool {
		return c == cli
	})

	if len(r.clients[userID]) == 0 {
		delete(r.clients, userID)
	}
}
