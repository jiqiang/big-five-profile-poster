package utils

// BigFiveResultsTextSerializer class
type BigFiveResultsTextSerializer struct {
	profile string
}

// Read function
func (s *BigFiveResultsTextSerializer) Read(text string) {
	s.profile = text
}

// Hash function
func (s BigFiveResultsTextSerializer) Hash() string {
	return s.profile
}
