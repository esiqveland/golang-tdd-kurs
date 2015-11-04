type Request struct {
	// Method specifies the HTTP method (GET, POST, PUT, etc.).
	// For client requests an empty string means GET.
	Method string
	URL *url.URL
	Body io.ReadCloser
	Header Header
	...
}