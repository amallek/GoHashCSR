package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"crypto/md5"
	"encoding/pem"
	"crypto/sha256"
)

func getSHA(data []byte) string {
	hash := sha256.Sum256(data)
	return strings.ToUpper(fmt.Sprintf("%x", hash[:]))
}

func getMD5(data []byte) string {
	hash := md5.Sum(data)
	return strings.ToUpper(fmt.Sprintf("%x", hash[:]))
}

func main() {

	var csr string

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if sc.Text() == "" {
			break
		}
		csr += sc.Text()+"\n"
	}
	csr = strings.Trim(csr,"\n")

	block, _ := pem.Decode([]byte(csr))
	if block == nil {
		panic("failed to parse PEM")
	}

	fmt.Println("CSR:\n", csr)
	fmt.Println("SHA-256: ", getSHA(block.Bytes))
	fmt.Println("MD5: ", getMD5(block.Bytes))
}
