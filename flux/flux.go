package flux

import (
	"github.com/jjeffcaii/reactor-go"
	"github.com/jjeffcaii/reactor-go/scheduler"
)

type OverflowStrategy int8

const (
	OverflowBuffer OverflowStrategy = iota
	OverflowIgnore
	OverflowError
	OverflowDrop
	OverflowLatest
)

type FnSwitchOnFirst = func(s Signal, f Flux) Flux

type Flux interface {
	rs.Publisher
	Filter(rs.Predicate) Flux
	Map(rs.Transformer) Flux
	DoOnDiscard(rs.FnOnDiscard) Flux
	DoOnNext(rs.FnOnNext) Flux
	DoOnComplete(rs.FnOnComplete) Flux
	DoOnCancel(rs.FnOnCancel) Flux
	DoOnRequest(rs.FnOnRequest) Flux
	DoFinally(rs.FnOnFinally) Flux
	SwitchOnFirst(FnSwitchOnFirst) Flux
	SubscribeOn(scheduler.Scheduler) Flux
}

type Sink interface {
	Complete()
	Error(error)
	Next(interface{})
}

type Processor interface {
	Flux
	Sink
}
