package hfsm

import (
	"context"
	"github.com/hootuu/hyle/data/dict"
)

type State int

type Event int

type Transition func(ctx context.Context, currentState State, event Event, data ...dict.Dict) (State, error)
