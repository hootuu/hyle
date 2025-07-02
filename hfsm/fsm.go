package hfsm

import (
	"context"
	"fmt"
	"github.com/hootuu/hyle/data/dict"
	"sync"
)

type Machine struct {
	transitions map[State]map[Event]Transition
	mu          sync.RWMutex
}

func NewMachine() *Machine {
	return &Machine{
		transitions: make(map[State]map[Event]Transition),
	}
}

func (m *Machine) AddTransition(current State, event Event, trans Transition) *Machine {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.transitions[current]; !ok {
		m.transitions[current] = make(map[Event]Transition)
	}

	m.transitions[current][event] = trans
	return m
}

func (m *Machine) Handle(ctx context.Context, current State, event Event, data dict.Dict) (State, error) {
	m.mu.Lock()
	transitions, ok := m.transitions[current]
	if !ok {
		m.mu.Unlock()
		return current, fmt.Errorf("no transitions defined for state: %d", current)
	}
	trans, ok := transitions[event]
	if !ok {
		m.mu.Unlock()
		return current, fmt.Errorf("no transition for event %d in state %d", event, current)
	}
	m.mu.Unlock()

	target, err := trans(ctx, current, event, data)
	if err != nil {
		return current, err
	}
	return target, nil
}
