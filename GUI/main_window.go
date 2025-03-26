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
			openSelectWindow(a, func(updatedData string) {
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

	// ✅ 创建主页面的三个按钮（纵向排列）
	button1 := widget.NewButton("新建工程", func() {
		openSelectWindow(a, func(updatedData string) {
			w.SetTitle("easy Chip - " + updatedData)
		})
	})
	button2 := widget.NewButton("打开工程", func() { openFileDialog() })
	button3 := widget.NewButton("设置", func() {})
	buttonContainer := container.NewVBox(button1, button2, button3)

	// ✅ 历史文件数据（可以动态更新）
	historyFiles := []string{}
	// ✅ 创建历史文件列表（右侧）
	historyList := widget.NewList(
		func() int {
			if len(historyFiles) == 0 {
				return 1 // 显示“暂无历史工程”
			}
			return len(historyFiles)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel(""),
				widget.NewButton("打开", func() {}),
				widget.NewButton("删除", func() {}),
			)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			hBox := obj.(*fyne.Container)
			label := hBox.Objects[0].(*widget.Label)
			openButton := hBox.Objects[1].(*widget.Button)
			deleteButton := hBox.Objects[2].(*widget.Button)

			if len(historyFiles) == 0 {
				label.SetText("暂无历史工程")
				openButton.Hide()
				deleteButton.Hide()
			} else {
				label.SetText(historyFiles[id])
				openButton.Show()
				deleteButton.Show()
			}
		},
	)
	historyScroll := container.NewVScroll(historyList)
	historyScroll.SetMinSize(fyne.NewSize(200, 150)) // 设置历史文件列表的最小大小

	// ✅ 使用 HSplit 布局，将按钮和历史文件列表水平排列
	mainContent := container.NewHSplit(buttonContainer, historyScroll)
	mainContent.SetOffset(0.3) // 按钮占 30%，历史文件占 70%

	// ✅ 使用 Border 布局，将菜单栏固定在顶部，主要内容在中间
	return container.NewBorder(menuBar, nil, nil, nil, mainContent)
}

func MainWindows(w fyne.Window, a fyne.App) *fyne.Container {
	widget.NewLabel("文件保存路径选择")

	return startPage(w, a)
}
