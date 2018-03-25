package main

import "io"

const MY_COPY_BUF_SIZE = 1024 * 4

func MyCopy(dst io.Writer, src io.Reader) (written int64, err error) {
	buf := make([]byte, MY_COPY_BUF_SIZE)
	for {
		n, err := src.Read(buf)
		written += int64(n)
		nn := 0
		for nn < n {
			nw, ew := dst.Write(buf[nn:n])
			nn += nw
			if ew != nil {
				break
			}
		}
		if err != nil {
			break
		}
	}
	return
}
