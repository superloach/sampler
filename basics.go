package sampler

func Sine(s *Sampler) float64 {
	return s.Volume * (1 << 15) * (math.Sin(2*math.Pi*(s.Frequency/float64(s.Rate))*float64(s.Index)) + 1) / 2
}
