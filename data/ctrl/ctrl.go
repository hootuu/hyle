package ctrl

import "errors"

type Ctrl []byte

func NewCtrl(length ...int) (Ctrl, error) {
	size := 128
	if len(length) > 0 {
		size = length[0]
	}
	if size <= 0 {
		return nil, errors.New("length must be greater than 0")
	}
	byteSize := (size + 7) / 8
	return make(Ctrl, byteSize), nil
}

func MustNewCtrl(length ...int) Ctrl {
	ctrl, _ := NewCtrl(length...)
	return ctrl
}

func (ctrl Ctrl) Set(pos int, value bool) error {
	if pos < 0 || pos >= len(ctrl)*8 {
		return errors.New("position out of bounds")
	}

	byteIndex := pos / 8
	bitIndex := pos % 8

	if value {
		// Set to 1, Use OR Opr
		ctrl[byteIndex] |= 1 << bitIndex
	} else {
		// Set to 0, Use AND Opr
		ctrl[byteIndex] &^= 1 << bitIndex
	}

	return nil
}

func (ctrl Ctrl) MustSet(pos int, value bool) Ctrl {
	_ = ctrl.Set(pos, value)
	return ctrl
}

func (ctrl Ctrl) Check(pos int) (bool, error) {
	if pos < 0 || pos >= len(ctrl)*8 {
		return false, errors.New("position out of bounds")
	}

	byteIndex := pos / 8
	bitIndex := pos % 8

	if ctrl[byteIndex]&(1<<bitIndex) != 0 {
		return true, nil
	}
	return false, nil
}

func (ctrl Ctrl) Iter(callback func(pos int, value bool)) {
	for i := 0; i < len(ctrl)*8; i++ {
		value, _ := ctrl.Check(i)
		callback(i, value)
	}
}
