package transport

type Transport interface {
	SendUrls(request []byte) []byte
}