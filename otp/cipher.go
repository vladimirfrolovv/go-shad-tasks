//go:build !solution

package otp

import (
	"io"
)

type CipherReader struct {
	r    io.Reader
	prng io.Reader
}

func (c *CipherReader) Read(p []byte) (n int, err error) {

	n, err = c.r.Read(p)

	key := make([]byte, n)
	c.prng.Read(key)
	for i := 0; i < n; i++ {
		p[i] ^= key[i]
	}
	if err != nil {
		return n, err
	}
	return n, nil
}

func NewReader(r io.Reader, prng io.Reader) io.Reader {
	return &CipherReader{r: r, prng: prng}
}

func NewWriter(w io.Writer, prng io.Reader) io.Writer {
	return &CipherWriter{w: w, prng: prng}
}

type CipherWriter struct {
	w    io.Writer
	prng io.Reader
}

func (c *CipherWriter) Write(p []byte) (n int, err error) {
	key := make([]byte, len(p))
	_, err = c.prng.Read(key)
	if err != nil {
		return n, err
	}
	ciphertext := make([]byte, len(p))
	for i := 0; i < len(p); i++ {
		ciphertext[i] = p[i] ^ key[i]
	}

	return c.w.Write(ciphertext)
}
