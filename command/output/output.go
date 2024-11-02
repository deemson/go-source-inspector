package output

type Output interface {
	Render(renderer Renderer) string
}
