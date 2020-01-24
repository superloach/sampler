package sampler

type Sampler struct {
	Index     int64
	Frequency float64
	Volume    float64
	Rate      int64
	Func     func(*Sampler) float64
}

func (s *Sampler) Read(p []byte) (n int, err error) {
	samples := len(p) / 2

	for i := 0; i < samples; i++ {
		s.Index = s.Index + 1
		sample := s.Sample(s)
		sample16bit := uint16(sample)
		p[2*i] = uint8(sample16bit & 0xff)
		p[2*i+1] = uint8(sample16bit >> 8)
	}

	return samples * 2, nil
}

func (s *Sampler) Seek(offset int64, whence int) (n int64, err error) {
	switch whence {
	case io.SeekStart:
		s.Index = offset
	case io.SeekCurrent:
		s.Index = s.Index + offset
	default:
		return 0, nil
	}

	return int64(s.Index), nil
}

func (s *Sampler) Close() error {
	return nil
}