package bot

import (
	"cmp"
	"context"
	"errors"
	"log/slog"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/teeworlds-go/protocol/examples/client_trivia/servers"
)

var (
	ErrNoActiveServers = errors.New("no active servers available")
)

func NewServerFactory() ServerFactory {
	return ServerFactory{}
}

type ServerFactory struct {
	mu          sync.Mutex
	lastServers []servers.Server
	lastIndex   int

	lastFetch time.Time
}

func (s *ServerFactory) Next(ctx context.Context) (servers.Server, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastIndex++
	fetchAgain := time.Since(s.lastFetch) > 5*time.Minute || s.lastIndex >= len(s.lastServers)
	if fetchAgain {
		serverList, err := s.fetch(ctx)
		if err != nil {
			return servers.Server{}, err
		}
		s.lastServers = serverList
		s.lastFetch = time.Now()
		s.lastIndex = 0

		if len(s.lastServers) == 0 {
			return servers.Server{}, ErrNoActiveServers
		}
	}

	return s.lastServers[s.lastIndex], nil
}

func (sf *ServerFactory) fetch(ctx context.Context) ([]servers.Server, error) {
	slog.Info("fetching servers")
	const sixupProto = "tw-0.7+"
	serverList, err := servers.GetAllServers(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]servers.Server, 0, len(serverList))

	// only filter servers with 0.7 protocol
	for _, server := range serverList {
		addresses := make([]string, 0, 1)
		for _, address := range server.Addresses {
			if strings.HasPrefix(address, sixupProto) {
				addresses = append(addresses, address)
			}
		}

		// no 0.7 address, skip
		if len(addresses) == 0 {
			continue
		}

		if len(server.Info.Clients) == 0 {
			// skip empty servers
			continue
		}

		// server full, skip
		if len(server.Info.Clients) == int(server.Info.MaxClients) {
			// server is full
			continue
		}

		server.Addresses = addresses
		result = append(result, server)
	}

	slices.SortFunc(result, func(a, b servers.Server) int {
		return cmp.Compare(len(b.Info.Clients), len(a.Info.Clients))
	})

	slog.Info("fetched servers", "count", len(result))
	return result, nil
}
