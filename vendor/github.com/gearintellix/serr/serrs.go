package serr

// SErrs type mutiple standard error
type SErrs []SErr

// CaptureSErr to capture standard error
func (ox *SErrs) CaptureSErr(errx SErr) {
	*ox = append(*ox, errx)
}
