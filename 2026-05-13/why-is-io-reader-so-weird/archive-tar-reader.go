package tar

import "io" // OMIT
// OMIT
func (tr *Reader) Read(b []byte) (int, error) {
	if tr.err != nil { // Abort if the inner Reader returned an error.
		return 0, tr.err
	}
	n, err := tr.curr.Read(b) // Read from the inner Reader // HL
	if err != nil && err != io.EOF {
		tr.err = err
	}
	return n, err
}
