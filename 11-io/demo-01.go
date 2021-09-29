package main

import (
	"io"
	"os"
)

type AlphaReader string

func (alphaReader AlphaReader) Read(p []byte) (n int, err error) {
	count := 0
	for i := 0; i < len(alphaReader); i++ {
		if (alphaReader[i] >= 'a' && alphaReader[i] <= 'z') || (alphaReader[i] >= 'A' && alphaReader[i] <= 'Z') {
			p[count] = alphaReader[i]
			count++
		}
	}
	return count, io.EOF
}
func main() {
	alphaReader := AlphaReader("378535 Consequat cillum^(*&(&* adipisicing ea nisi tempor consequat velit commodo nisi dolor culpa magna nulla. Aute in elit ea occaecat elit. Ut laborum dolor occaecat laborum pariatur ipsum adipisicing cupidatat irure cupidatat reprehenderit esse. Amet ut irure Lorem in mollit proident exercitation consectetur magna et cillum aliqua. Eu sit aliqua elit commodo quis amet enim. Id sint et proident nisi aliquip dolore consectetur ullamco ad.")
	io.Copy(os.Stdout, alphaReader)
}
