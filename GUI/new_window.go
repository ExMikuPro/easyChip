package GUI

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func openNewWindow(a fyne.App, onUpdate func(string)) {
	newWindow := a.NewWindow("选择MCU")
	newWindow.Resize(fyne.NewSize(800, 600))

	// 变量存储用户选中的 MCU 型号
	selectedMCU := ""

	// 创建开始按钮，并默认禁用
	startButton := widget.NewButton("开始", func() {
		fmt.Println("程序开始运行，选中的 MCU: ", selectedMCU)
		onUpdate(selectedMCU) // 将选中的 MCU 传回主窗口
		newWindow.Close()     // 关闭子窗口
	})
	startButton.Disable() // 初始时禁用

	// ✅ 右侧操作区域，包含“开始”按钮
	rightControls := container.NewBorder(nil, nil, nil,
		container.NewGridWrap(fyne.NewSize(100, 40), startButton))

	// ✅ 让顶部工具栏不会过高
	topBar := container.NewBorder(nil, nil, nil, rightControls)

	// ✅ MCU 分类列表
	categories := []string{"MCU A系列", "MCU B系列", "MCU C系列", "MCU D系列"}
	mcuData := map[string][]string{
		"MCU A系列": {"A-100", "A-200", "A-300"},
		"MCU B系列": {"B-100", "B-200"},
		"MCU C系列": {"C-100", "C-200", "C-300", "C-400"},
		"MCU D系列": {"D-100", "D-200", "D-300", "D-400", "D-500"},
	}

	// ✅ MCU 详情数据
	mcuDetails := map[string]string{
		"A-100": "A-100: 32-bit, 48MHz, 128KB Flash",
		"A-200": "A-200: 32-bit, 72MHz, 256KB Flash",
		"B-100": "B-100: 16-bit, 24MHz, 64KB Flash",
		"B-200": "B-200: 16-bit, 48MHz, 128KB Flash",
		"C-100": "C-100: ARM Cortex-M3, 120MHz, 512KB Flash",
		"D-100": "D-100: RISC-V, 160MHz, 1MB Flash",
	}

	// ✅ 右侧信息展示区域
	infoLabel := widget.NewLabel("MCU 详情信息将显示在这里...")
	separator := widget.NewSeparator() // 分割线
	infoContainer := container.NewVBox(separator, infoLabel)

	// ✅ 右侧 MCU 详情列表（初始为空）
	rightData := []string{}
	rightList := widget.NewList(
		func() int { return len(rightData) }, // 数据数量
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(rightData[i]) // 绑定数据
		},
	)

	// ✅ 创建左侧分类列表
	leftList := widget.NewList(
		func() int { return len(categories) }, // 数据数量
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(categories[i]) // 绑定数据
		},
	)

	// ✅ 当左侧分类被点击时，更新右侧 MCU 列表
	leftList.OnSelected = func(id widget.ListItemID) {
		category := categories[id]
		rightData = mcuData[category] // 更新数据
		rightList.Refresh()           // 刷新右侧列表
	}

	// ✅ 当右侧 MCU 详情列表被点击时，存储选中的 MCU 型号并更新详情信息
	rightList.OnSelected = func(id widget.ListItemID) {
		selectedMCU = rightData[id]
		fmt.Println("已选择 MCU: ", selectedMCU)
		if detail, exists := mcuDetails[selectedMCU]; exists {
			infoLabel.SetText(detail)
		} else {
			infoLabel.SetText("未知的 MCU 详情")
		}
		startButton.Enable() // 选定 MCU 后启用按钮
	}

	// ✅ 滚动容器（左侧分类列表）
	leftScroll := container.NewGridWrap(fyne.NewSize(200, 500), container.NewScroll(leftList))

	// ✅ 右侧 MCU 详情列表（上部）
	rightScroll := container.NewGridWrap(fyne.NewSize(600, 250), container.NewScroll(rightList))

	// ✅ 右侧整体布局（列表 + 信息区域）
	rightContainer := container.NewVBox(rightScroll, separator, infoContainer)

	// ✅ 使用 `NewHSplit()` 控制 2:5 比例
	splitContainer := container.NewHSplit(leftScroll, rightContainer)
	splitContainer.SetOffset(0.3) // **左侧占 30%，右侧占 70%**

	// ✅ 设置窗口内容
	newWindow.SetContent(container.NewBorder(topBar, nil, nil, nil, splitContainer))
	newWindow.Show()
}
