package ex

import (
	"github.com/hootuu/hyle/data/ctrl"
	"github.com/hootuu/hyle/data/dict"
	"github.com/hootuu/hyle/data/hjson"
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

func MustFix(src *Ex, dst *Ex) *Ex {

	if src != nil {
		if dst == nil {
			return src
		}
		newEx := EmptyEx()
		if len(src.Tag) > 0 {
			newEx.Tag.Append(src.Tag...)
		}
		if len(dst.Tag) > 0 {
			newEx.Tag.Append(dst.Tag...)
		}
		if src.Meta != nil {
			for k, v := range src.Meta {
				newEx.Meta[k] = v
			}
		}
		if dst.Meta != nil {
			for k, v := range dst.Meta {
				newEx.Meta[k] = v
			}
		}
		return newEx
		//todo add ctrl info

	} else {
		if dst == nil {
			return EmptyEx()
		}
		return dst
	}
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

func WithBytes(ctrl Ctrl, tag []byte, meta []byte) *Ex {
	ex := &Ex{
		Ctrl: ctrl,
		Tag:  nil,
		Meta: nil,
	}
	if len(tag) > 0 {
		ex.Tag = *hjson.MustFromBytes[Tag](tag)
	}
	if len(meta) > 0 {
		ex.Meta = *hjson.MustFromBytes[Meta](meta)
	}
	return ex
}
