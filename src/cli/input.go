package cli

import (
	ui "github.com/gizak/termui/v3"
	"regexp"

	"github.com/gizak/termui/v3/widgets"
)

type Input struct {
	input *widgets.Paragraph
}

func (i Input) Draw() {
	ui.Render(i.input)
}

func (i Input) Digit(keyboard string) {
	isSpecial := regexp.MustCompile(`<*>`)
	if isSpecial.MatchString(keyboard) {
		switch keyboard {
		case "<Backspace>":
			if len(i.input.Text) == 0 { return }
			current := i.input.Text
			current = current[:len(current)-1]
			i.input.Text = current
		case "<Space>":
			i.input.Text += " "
		case "<Enter>":
			return
		}
	} else {
		i.input.Text += keyboard
	}

}

func initInput(x1, y1, x2, y2 int) Input {
	input := widgets.NewParagraph()
	input.SetRect(x1, y1, x2, y2)
	return Input{input: input}
}
