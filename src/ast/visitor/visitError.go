package visitor

type VisitorError struct {
	Line, Column       int
	Msg                string
	Raw                error
	Text               string
	EndLine, EndColumn int
}
