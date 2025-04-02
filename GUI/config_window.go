package GUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func openConfigWindow(ctx AppContext) {
	configWindow := ctx.App.NewWindow("easyChip " + ctx.MCUConfig)
	configWindow.Resize(fyne.NewSize(800, 600))

	menuBar := container.NewHBox( // 顶部按钮
		widget.NewButton("Pinout&Configuration", nil),
		widget.NewButton("Project Manager", func() {}),
	)

	accordion := widget.NewAccordion()
	btn := widget.NewButton("GPIO", func() {
		println("按钮被点击")
	})
	fixedContainer := container.NewVBox(btn)
	btn1 := widget.NewButton("ADC", func() {
		println("按钮被点击")
	})
	fixedContainer1 := container.NewVBox(btn1)
	item := widget.NewAccordionItem("System Core", fixedContainer)
	item1 := widget.NewAccordionItem("Analog", fixedContainer1)

	accordion.Append(item)
	accordion.Append(item1)

	accordion.MultiOpen = true // 允许多个展开

	// 如果需要控制内容宽度但允许高度自动调整
	leftScroll := container.NewGridWrap(fyne.NewSize(230, 500), container.NewScroll(accordion))
	// ✅ 创建主页面的三个按钮（纵向排列）

	//content := container.NewStack(scroll)
	//
	//button2 := widget.NewButton("Connectivity", func() {})

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
	mainContent := container.NewHSplit(leftScroll, historyScroll)
	mainContent.SetOffset(0.3) // 按钮占 30%，历史文件占 70%

	// ✅ 使用 Border 布局，将菜单栏固定在顶部，主要内容在中间

	configWindow.SetContent(container.NewBorder(menuBar, nil, nil, nil, mainContent))

	configWindow.Show()
}
