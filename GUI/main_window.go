package GUI

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func startPage(w fyne.Window, a fyne.App) *fyne.Container {
	// ✅ 文件打开对话框
	openFileDialog := func() {
		openDialog := dialog.NewFileOpen(
			func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, w)
					return
				}
				if reader == nil {
					return
				}

				// 获取用户选择的文件路径
				filePath := reader.URI().Path()

				fmt.Println(filePath)

				defer func(reader fyne.URIReadCloser) {
					err := reader.Close()
					if err != nil {

					}
				}(reader) // 关闭文件
			}, w)

		// 设定文件类型过滤器（只允许选择 .ech 文件）
		openDialog.SetFilter(storage.NewExtensionFileFilter([]string{".ech"}))

		// 显示文件选择对话框
		openDialog.Show()
	}

	// ✅ 创建菜单按钮
	menuButton := widget.NewButton("文件 ▼", nil)

	// ✅ 创建菜单
	inputEntry := widget.NewEntry()
	fileMenu := fyne.NewMenu("",
		fyne.NewMenuItem("新建", func() {
			openNewWindow(a, func(updatedData string) {
				inputEntry.SetText(updatedData)         // 更新主窗口的数据
				w.SetTitle("easyChip - " + updatedData) // 更新窗口标题
			})
		}),
		fyne.NewMenuItem("打开", func() { openFileDialog() }),
		fyne.NewMenuItem("保存", func() {}),
	)
	popUpMenu := widget.NewPopUpMenu(fileMenu, w.Canvas())

	// ✅ 绑定点击事件，调整菜单位置
	menuButton.OnTapped = func() {
		// 计算菜单显示位置（在按钮下方）
		buttonPos := menuButton.Position()
		popUpMenu.Move(fyne.NewPos(buttonPos.X+10, buttonPos.Y+menuButton.Size().Height))
		popUpMenu.Show() // 显示菜单
	}

	// ✅ 创建菜单栏（水平排列）
	menuBar := container.NewHBox(
		menuButton,
		widget.NewButton("编辑", func() {}),
		widget.NewButton("视图", func() {}),
	)

	// ✅ 使用 Border 布局，将菜单栏固定在顶部
	return container.NewBorder(menuBar, nil, nil, nil, nil)
}

func MainWindows(w fyne.Window, a fyne.App) *fyne.Container {
	widget.NewLabel("文件保存路径选择")

	return startPage(w, a)
}
