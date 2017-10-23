package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/shenwei356/bio/seqio/fastx"
	//"github.com/srynobio/go-vmc/vmc"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	// open fastq file
	reader, err := fastx.NewDefaultReader(os.Args[1])
	eCheck(err)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		eCheck(err)

	}
}

// ------------------------ //

func digest(bv []byte, truncate int) string {
	hasher := sha512.New()
	hasher.Write(bv)

	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil)[:truncate])
	vmcID := "VMC:GS_" + sha
	return vmcID
}

// -------------------------------------------- //

func eCheck(p error) {
	if p != nil {
		panic(p)
	}
}

// -------------------------------------------- //
