package sampler

import (
	"math"
	"math/rand"
)

// Func is a function which generates waveforms.
type Func func(*Sampler) float64

// Sine is a Func that generates a sine wave.
func Sine(s *Sampler) float64 {
	return s.Volume * (math.Sin(2*math.Pi*(s.Frequency/float64(s.Rate))*float64(s.Index)) + 1) / 2
}

// Saw is a Func that generates a saw wave.
func Saw(s *Sampler) float64 {
	rf := float64(s.Rate) / s.Frequency
	return s.Volume * (math.Mod(float64(s.Index), rf) / rf)
}

// Pulse returns a Func that generates a pulse wave with a given duty.
func Pulse(duty float64) Func {
	return func(s *Sampler) float64 {
		sw := Saw(s) / s.Volume
		if sw > duty {
			return 0
		}

		return s.Volume
	}
}

// Square is a Func that generates a square wave.
var Square Func = Pulse(0.5)

// Noise is a Func that generates random noise.
func Noise(s *Sampler) float64 {
	return rand.Float64() * s.Volume
}
