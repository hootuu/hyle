package ex

import (
	"github.com/hootuu/hyle/data/ctrl"
	"github.com/hootuu/hyle/data/dict"
	"github.com/hootuu/hyle/data/tag"
)

type Ex struct {
	Ctrl ctrl.Ctrl `json:"ctrl"`
	Tag  tag.Tag   `json:"tag"`
	Meta dict.Dict `json:"meta"`
}

func NewEx() *Ex {
	e := &Ex{}
	e.SetCtrl(nil).SetTag(nil).SetMeta(nil)
	return e
}

func MustEx(ex *Ex) *Ex {
	if ex == nil {
		return EmptyEx()
	}
	if ex.Ctrl == nil {
		ex.SetCtrl(ctrl.MustNewCtrl())
	}
	if ex.Tag == nil {
		ex.SetTag(tag.NewTag())
	}
	if ex.Meta == nil {
		ex.SetMeta(dict.NewDict())
	}
	return ex
}

func (e *Ex) SetCtrl(c ctrl.Ctrl) *Ex {
	if c == nil {
		c = ctrl.MustNewCtrl()
	}
	e.Ctrl = c
	return e
}

func (e *Ex) SetTag(t tag.Tag) *Ex {
	if t == nil {
		t = tag.NewTag()
	}
	e.Tag = t
	return e
}

func (e *Ex) SetMeta(meta dict.Dict) *Ex {
	if meta == nil {
		meta = dict.NewDict()
	}
	e.Meta = meta
	return e
}

func EmptyEx() *Ex {
	return NewEx().
		SetCtrl(ctrl.MustNewCtrl()).
		SetTag(tag.NewTag()).
		SetMeta(dict.NewDict())
}
