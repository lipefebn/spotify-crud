package cli

import (
	"regexp"
	"strconv"

	ui "github.com/gizak/termui/v3"
)

type Forms struct {
	inputs []Input
	focus Input
}

func createForms(x1, y1 int, inputs ...string) (result Forms) {
	result = Forms{ inputs: []Input{} }
	for range(inputs) {
		result.inputs = append(result.inputs, initInput(x1, y1, x1 + 50, y1 + 3))
		y1+=4
	}
	for i, elem := range(result.inputs) {
		elem.input.Title = inputs[i]
	}
	return
	//form := Forms{ []Input{} }
}

func (f Forms) Draw() {
	ui.Clear()
	for _, input := range(f.inputs) { 
		input.Draw() 
	}
}

func (f Forms) ChooseUser() []string{
	intervalNumber := regexp.MustCompile("[1-"+ strconv.Itoa(len(f.inputs))+ "]")

	for {
		choose := (<-uiEvents).ID

		if intervalNumber.MatchString(choose) {
			i, _ := strconv.Atoi(choose)
			f.focus = f.inputs[i-1]
			f.Digit()
		}
		switch choose {
			case "<Enter>":
				return f.getTexts()
		}
	}
}

func (f Forms) Digit(){
	for {
		digit := (<-uiEvents).ID
		switch digit {
		case "<Enter>":
			return
		default:
			f.focus.Digit(digit)
			f.Draw()
	}
	}
}

func (f Forms) getTexts() (result []string){
	result = []string{}
	for _, current := range(f.inputs) {
		result = append(result, current.input.Text)
	}
	return
}