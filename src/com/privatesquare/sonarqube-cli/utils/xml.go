package utils

import (
	"bytes"
	"encoding/xml"
	"io"
)

func FormatXML(data []byte) ([]byte, error) {
	b := &bytes.Buffer{}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	encoder := xml.NewEncoder(b)
	encoder.Indent("", "  ")
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			encoder.Flush()
			return b.Bytes(), nil
		}
		if err != nil {
			return nil, err
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			return nil, err
		}
	}
}