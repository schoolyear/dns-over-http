package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"image/png"
	"log"
	"os"

	"github.com/auyer/steganography"
)

type ResultEncoder interface {
	Encode(data []byte) ([]byte, error)
}

type ResultEncoderFn func(data []byte) ([]byte, error)

func (r ResultEncoderFn) Encode(data []byte) ([]byte, error) {
	return r(data)
}

var encoders = map[string]ResultEncoder{
	"plain": ResultEncoderFn(func(data []byte) ([]byte, error) {
		return data, nil
	}),
	"base64": ResultEncoderFn(func(data []byte) ([]byte, error) {
		out := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
		base64.StdEncoding.Encode(out, data)
		return out, nil
	}),
	"meme": mustNewSteganographyEncoder("img.png"),
}

func mustNewSteganographyEncoder(pngFilePath string) ResultEncoder {
	// load file
	inFile, _ := os.Open(pngFilePath)
	reader := bufio.NewReader(inFile)
	img, err := png.Decode(reader)
	if err != nil {
		log.Fatalln("failed to load steganography image", pngFilePath, err)
	}

	return ResultEncoderFn(func(data []byte) ([]byte, error) {
		out := new(bytes.Buffer)
		err := steganography.Encode(out, img, data)
		if err != nil {
			return nil, err
		}

		return out.Bytes(), nil
	})
}
