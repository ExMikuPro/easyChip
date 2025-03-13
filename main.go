package main

import (
	"easyChip/template"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// 选择文件夹路径
func createPathSelector(w fyne.Window) (*widget.Entry, *widget.Button) {
	pathEntry := widget.NewEntry()
	pathEntry.SetPlaceHolder("请选择保存文件的路径...")

	var selectedPath string

	selectPathButton := widget.NewButton("选择文件夹", func() {
		dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if uri == nil {
				fmt.Println("用户取消了选择")
				return
			}

			selectedPath = uri.Path()
			pathEntry.SetText(selectedPath)
			fmt.Println("用户选择的保存路径:", selectedPath)
		}, w).Show()
	})

	// 返回输入框和路径变量
	return pathEntry, selectPathButton
}

// 创建输入框
func createLabeledEntry(placeholder string, label string) (*widget.Entry, fyne.CanvasObject) {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(placeholder)
	return entry, container.NewVBox(widget.NewLabel(label), entry)
}

// 创建确认按钮
func createConfirmButton(w fyne.Window, workspaceEntry, projectEntry *widget.Entry, pathEntry *widget.Entry) *widget.Button {
	return widget.NewButton("确认", func() {
		workspaceName := workspaceEntry.Text
		projectName := projectEntry.Text
		basePath := pathEntry.Text + "/"

		if workspaceName == "" || projectName == "" || basePath == "" {
			dialog.ShowError(fmt.Errorf("所有字段均不能为空"), w)
			return
		}

		// 调用 template 相关方法
		template.Iar_eww(workspaceName, basePath, template.IarEwwType{
			Projects: []string{projectName},
		})
		template.Iar_ewp(projectName, basePath, template.IarEwpType{
			false,
			template.CCOptLevel.LowOptimization,
			template.IccLang.CLang,
			template.IccCDialect.C11,
		})

		// 控制台输出
		fmt.Printf("✅ 工作区名称: %s\n✅ 工程名称: %s\n✅ 文件保存路径: %s\n", workspaceName, projectName, basePath)

		// 显示弹窗
		dialog.NewInformation("提示", "工作区和工程已创建！", w).Show()
	})
}

func main() {
	a := app.NewWithID("org.sfnco.easychip")
	w := a.NewWindow("easy Chip")
	w.Resize(fyne.NewSize(800, 600))

	// 创建输入框
	workspaceEntry, workspaceContainer := createLabeledEntry("请输入工作区名称...", "工作区信息")
	projectEntry, projectContainer := createLabeledEntry("请输入工程名称...", "工程信息")

	// 创建路径选择器
	pathEntry, selectPathButton := createPathSelector(w)

	// 创建确认按钮
	buttonConfirm := createConfirmButton(w, workspaceEntry, projectEntry, pathEntry)

	// 布局
	form := container.NewVBox(
		workspaceContainer,
		projectContainer,
		widget.NewLabel("文件保存路径选择"),
		pathEntry,
		selectPathButton,
		buttonConfirm,
	)

	// 设置窗口内容
	w.SetContent(form)
	w.ShowAndRun()
}
