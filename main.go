package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type SettingsStruct struct {
	ProjectName string   `json:"ProjectName"`
	Req         string   `json:"REQ"`
	Sprint      string   `json:"Sprint"`
	Hu          string   `json:"HU"`
	Rfc         string   `json:"RFC"`
	Type        []string `json:"Type"`
	Scope       []string `json:"Scope"`
}

func evalBox(Str string, defaultValue string, prefix string, end string) string {
	if Str == "" {
		return defaultValue
	} else {
		return prefix + Str + end
	}
}

func gcommitgrid(settings SettingsStruct) fyne.CanvasObject {
	txtProName := widget.NewEntry()
	txtReq := widget.NewEntry()
	txtSprint := widget.NewEntry()
	txtHU := widget.NewEntry()
	txtRFC := widget.NewEntry()
	cmbType := widget.NewSelect(settings.Type, func(value string) {
	})
	cmbScope := widget.NewSelect(settings.Scope, func(value string) {
	})
	txtDescription := widget.NewMultiLineEntry()
	txtBody := widget.NewMultiLineEntry()
	txtResult := widget.NewMultiLineEntry()

	txtProName.SetText(settings.ProjectName)
	txtReq.SetText(settings.Req)
	txtSprint.SetText(settings.Sprint)
	txtHU.SetText(settings.Hu)
	txtRFC.SetText(settings.Rfc)

	getMessage := func() {
		strType := evalBox(cmbType.Selected, "feat", "", "")
		strScope := evalBox(cmbScope.Selected, "", "(", ")")
		strDescription := evalBox(txtDescription.Text, "Commit 1", "", "")
		strBody := evalBox(txtBody.Text, "", "", "")
		strProject := evalBox(txtProName.Text, settings.ProjectName, "PROJECT:", "")
		strReq := evalBox(txtReq.Text, settings.Req, " REQ:", "")
		strHU := evalBox(txtHU.Text, settings.Hu, " HU:", "")
		strSprint := evalBox(txtSprint.Text, settings.Sprint, " SPRINT:", "")
		strRFC := evalBox(txtRFC.Text, settings.Rfc, " RFC:", "")
		message := fmt.Sprintf("%s%s: %s\n%s\n%s%s%s%s%s", strType, strScope, strDescription, strBody, strProject, strReq, strHU, strSprint, strRFC)
		txtResult.SetText(message)
	}

	form := &widget.Form{
		Items:    []*widget.FormItem{},
		OnSubmit: getMessage,
	}
	form.Append("Project name", txtProName)
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

func readSettings() SettingsStruct {
	jsonFile, err := os.Open("gcommit.json")
	if err != nil {
		log.Printf("Could not open settings")
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var settings SettingsStruct
	json.Unmarshal(byteValue, &settings)
	jsonFile.Close()
	return settings
}
func gwindow(app fyne.App, settings SettingsStruct) fyne.Window {
	myWindow := app.NewWindow("gcommit")
	tabs := container.NewAppTabs(
		container.NewTabItem("gcommit", gcommitgrid(settings)),
	)
	tabs.SetTabLocation(container.TabLocationTop)
	myWindow.SetContent(tabs)
	return myWindow
}

func main() {
	settings := readSettings()
	gcapp := app.New()
	gcw := gwindow(gcapp, settings)
	gcw.ShowAndRun()

}
