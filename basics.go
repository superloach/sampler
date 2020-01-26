package sampler

import (
	"math"
	"math/rand"
)

type Func (func(*Sampler) float64)

func Sine(s *Sampler) float64 {
	return s.Volume * (math.Sin(2*math.Pi*(s.Frequency/float64(s.Rate))*float64(s.Index)) + 1) / 2
}

func Saw(s *Sampler) float64 {
	rf := float64(s.Rate) / s.Frequency
	return s.Volume * (math.Mod(float64(s.Index), rf) / rf)
}

func Pulse(duty float64) Func {
	return func(s *Sampler) float64 {
		sw := Saw(s) / s.Volume
		if sw > duty {
			return 0
		} else {
			return s.Volume
		}
	}
}

var Square Func = Pulse(0.5)

func Noise(s *Sampler) float64 {
	return rand.Float64() * s.Volume
}
