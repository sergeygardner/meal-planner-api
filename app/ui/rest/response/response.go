package response

import (
	"github.com/go-chi/render"
)

type Response interface {
	GetStatus() int
	render.Renderer
}
