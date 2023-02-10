package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func evalBox(Str string, defaultValue string, prefix string, end string) string {
	if Str == "" {
		return defaultValue
	} else {
		return prefix + Str + end
	}
}

func gcommitgrid() fyne.CanvasObject {
	txtProName := widget.NewEntry()
	txtReq := widget.NewEntry()
	txtSprint := widget.NewEntry()
	txtHU := widget.NewEntry()
	txtRFC := widget.NewEntry()
	cmbType := widget.NewSelect([]string{"fix", "feat", "BREAKING CHANGE"}, func(value string) {
	})
	cmbScope := widget.NewSelect([]string{"release", "models", "controllers", "frontend"}, func(value string) {
	})
	txtDescription := widget.NewMultiLineEntry()
	txtBody := widget.NewMultiLineEntry()
	txtResult := widget.NewMultiLineEntry()

	getMessage := func() {
		strType := evalBox(cmbType.Selected, "feat", "", "")
		strScope := evalBox(cmbScope.Selected, "", "(", ")")
		strDescription := evalBox(txtDescription.Text, "Commit 1", "", "")
		strBody := evalBox(txtBody.Text, "", "", "")
		strProject := evalBox(txtProName.Text, "", "PROJECT:", "")
		strReq := evalBox(txtReq.Text, "", " REQ:", "")
		strHU := evalBox(txtHU.Text, "", " HU:", "")
		strSprint := evalBox(txtSprint.Text, "", " SPRINT:", "")
		strRFC := evalBox(txtRFC.Text, "", " RFC:", "")
		message := fmt.Sprintf("%s%s: %s\n%s\n%s%s%s%s%s", strType, strScope, strDescription, strBody, strProject, strReq, strHU, strSprint, strRFC)
		txtResult.SetText(message)
	}

	form := &widget.Form{
		Items:    []*widget.FormItem{},
		OnSubmit: getMessage,
	}
	form.Append("Projec name", txtProName)
	form.Append("REQ", txtReq)
	form.Append("Sprint", txtSprint)
	form.Append("HU", txtHU)
	form.Append("RFC", txtRFC)
	form.Append("Type", cmbType)
	form.Append("Scope", cmbScope)
	form.Append("Description", txtDescription)
	form.Append("Body", txtBody)
	form.Append("Result", txtResult)

	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(400, 550)), form)
	return grid
}

func gwindow(app fyne.App) fyne.Window {
	myWindow := app.NewWindow("gcommit")
	tabs := container.NewAppTabs(
		container.NewTabItem("gcommit", gcommitgrid()),
	)
	tabs.SetTabLocation(container.TabLocationTop)
	myWindow.SetContent(tabs)
	return myWindow
}

func main() {
	gcapp := app.New()
	gcw := gwindow(gcapp)
	gcw.ShowAndRun()

}
