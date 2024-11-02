package output

var _ Output = Error{}

type Error struct {
	Message string
}

func (e Error) Render(renderer Renderer) string {
	return renderer.RenderError(e)
}
