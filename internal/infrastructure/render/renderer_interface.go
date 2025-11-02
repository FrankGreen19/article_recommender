package render

import "net/http"

type Renderer interface {
	Render(http.ResponseWriter, interface{})
}
