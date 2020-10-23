package xml

import (
	"bytes"
	"encoding/xml"
	"io"
)

func EncodeXml(i interface{}) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	buf.WriteString(xml.Header)
	enc := xml.NewEncoder(buf)
	enc.Indent("", "\t")
	if e := enc.Encode(i); e != nil {
		return nil, e
	}
	return buf, nil

}
func EncodeXmlToString(i interface{}) (string, error) {
	buf, e := EncodeXml(i)
	if e != nil {
		return "", e
	}
	return buf.String(), nil
}
func DecodeXml(r io.Reader, v interface{}) error {
	dec := xml.NewDecoder(r)
	return dec.Decode(v)
}
func DecodeXmlCustom(r io.Reader, v interface{}, f func(d *xml.Decoder)) error {
	dec := xml.NewDecoder(r)
	f(dec)
	return dec.Decode(v)
}
