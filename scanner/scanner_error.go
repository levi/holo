package scanner

type ScannerError struct {
	s    string
	Line int
}

func (e *ScannerError) Error() string {
	return e.s
}