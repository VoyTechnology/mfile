package mfile

// ReadFile returns the whole data from the path.
func ReadFile(path string) ([]byte, error) {
	h, path, err := load(path)
	if err != nil {
		return nil, err
	}
	
	return h.ReadFile(path)
}