package gate

import (
	"gate/conf"
	"gate/errors"
	"gate/grpc"
	"gate/xframe/xlog"
)

// Ring ring proto buffer.
type Ring struct {
	// read
	rp   uint64 //读位置
	num  uint64
	mask uint64
	// write
	wp   uint64       //写位置
	data []grpc.Proto //缓冲区
}

func NewRing(num int) *Ring {
	r := new(Ring)
	r.init(uint64(num))
	return r
}

func (r *Ring) Init(num int) {
	r.init(uint64(num))
}

func (r *Ring) init(num uint64) {
	// 2^N
	if num&(num-1) != 0 {
		for num&(num-1) != 0 {
			num &= num - 1
		}
		num = num << 1
	}
	r.data = make([]grpc.Proto, num)
	r.num = num
	r.mask = r.num - 1
}

func (r *Ring) Get() (proto *grpc.Proto, err error) {
	if r.rp == r.wp {
		return nil, errors.ErrRingEmpty
	}
	proto = &r.data[r.rp&r.mask]
	return
}

func (r *Ring) GetAdv() {
	r.rp++
	if conf.Conf.Debug {
		xlog.Infof("ring rp: %d, idx: %d", r.rp, r.rp&r.mask)
	}
}

func (r *Ring) Set() (proto *grpc.Proto, err error) {
	if r.wp-r.rp >= r.num {
		return nil, errors.ErrRingFull
	}
	proto = &r.data[r.wp&r.mask]
	return
}

func (r *Ring) SetAdv() {
	r.wp++
	if conf.Conf.Debug {
		xlog.Infof("ring wp: %d, idx: %d", r.wp, r.wp&r.mask)
	}
}

func (r *Ring) Reset() {
	r.rp = 0
	r.wp = 0
}
