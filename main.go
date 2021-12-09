package main

import (
	"fmt"
	"image"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ImgInfo struct {
	Name string
	Img  *image.Image
	code string
}

type Browser struct {
	Curr int
	Imgs []ImgInfo

	WgtSelImage    *widget.Select
	CanvasImg      *canvas.Image
	CanvasTextStat *canvas.Text
}

func (b *Browser) SetImge(index int) {
	if index < 0 {
		index += len(b.Imgs)
		b.SetImge(index)
	}
	if index > len(b.Imgs)-1 {
		index -= len(b.Imgs)
		b.SetImge(index)
	}
	b.Curr = index

	b.CanvasImg.Image = *b.Imgs[index].Img
	canvas.Refresh(b.CanvasImg)
	b.ShowStat()
}

func (b *Browser) ImgNameList() []string {
	lst := make([]string, len(b.Imgs))
	for i, img := range b.Imgs {
		lst[i] = img.Name
	}
	return lst
}
func (b *Browser) ShowStat() {
	b.CanvasTextStat.Text = fmt.Sprint(b.Imgs[b.Curr].Name)
	canvas.Refresh(b.CanvasTextStat)
}
func loadChartsIamges() []ImgInfo {
	return []ImgInfo{
		{"PieChart", PieChart(), ""},
		{"StockAnalysis", StockAnalysis(), ""},
		{"StackedBar", StackedBar(), ""},
		{"Annotation", Annotation(), ""},
		{"BenchmarkLine", BenchmarkLine(), ""},
		{"LinearRegression", LinearRegression(), ""},
		{"BarChart", BarChart(), ""}}
}

func imgSelScreen(_ fyne.Window) fyne.CanvasObject {

	b := &Browser{Curr: 0}

	b.Imgs = loadChartsIamges()

	b.CanvasImg = canvas.NewImageFromImage(*b.Imgs[b.Curr].Img)

	b.CanvasTextStat = canvas.NewText("", theme.TextColor())
	prev := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		b.SetImge(b.Curr - 1)
	})
	next := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
		b.SetImge(b.Curr + 1)
	})
	b.WgtSelImage = widget.NewSelect(b.ImgNameList(),
		func(selName string) {
			for i, icon := range b.Imgs {
				if icon.Name == selName {
					b.SetImge(i)
					break
				}
			}
		})

	b.WgtSelImage.SetSelected(b.ImgNameList()[b.Curr])

	selFillMod := widget.NewSelect(
		[]string{
			"ImageFillStretch",
			"ImageFillContain",
			"ImageFillOriginal"},
		func(selName string) {
			switch selName {
			case "ImageFillStretch":
				b.CanvasImg.FillMode = canvas.ImageFillStretch
				break
			case "ImageFillContain":
				b.CanvasImg.FillMode = canvas.ImageFillContain
				break
			case "ImageFillOriginal":
				b.CanvasImg.FillMode = canvas.ImageFillOriginal
				break
			}
			canvas.Refresh(b.CanvasImg)
		})

	selFillMod.SetSelected("ImageFillOriginal")
	buttons := container.NewHBox(prev, next, selFillMod)

	bar := container.NewBorder(nil, nil, buttons, nil, b.WgtSelImage)

	return container.NewBorder(bar, b.CanvasTextStat, nil, nil, b.CanvasImg)
}

func main() {
	os.Setenv("FYNE_SCALE", "1")
	a := app.New()
	a.SetIcon(resourceIconPng)
	a.Settings().SetTheme(theme.LightTheme())

	w := a.NewWindow("go-chart using fyne")
	w.Resize(fyne.NewSize(1100, 640))

	w.SetContent(imgSelScreen(w))
	w.ShowAndRun()

}
