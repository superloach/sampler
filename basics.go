package sampler

import "math"

type Func (func(*Sampler) float64)

func Sine(s *Sampler) float64 {
	pk := s.Volume * (1 << 15)
	return pk * (math.Sin(2*math.Pi*(s.Frequency/float64(s.Rate))*float64(s.Index)) + 1) / 2
}

func Saw(s *Sampler) float64 {
	rf := float64(s.Rate) / s.Frequency
	pk := s.Volume * (1 << 15)
	return pk * (math.Mod(float64(s.Index), rf) / rf)
}

func Pulse(duty float64) Func {
	return func(s *Sampler) float64 {
		pk := s.Volume * (1 << 15)
		sw := Saw(s) / pk
		if sw > duty {
			return 0
		} else {
			return pk
		}
	}
}

var Square Func = Pulse(0.5)
