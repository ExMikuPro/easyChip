package GUI

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
	"path/filepath"
)

type Peripheral struct {
	Peripheral []CategoryGroup `json:"peripheral"`
}

type CategoryGroup map[string][]string

func GetCategoryNames(p Peripheral) []string {
	var names []string
	for _, item := range p.Peripheral {
		for key := range item {
			names = append(names, key)
		}
	}
	return names
}

func GetPeripheralsByCategory(p Peripheral, category string) []string {
	for _, group := range p.Peripheral {
		if devices, exists := group[category]; exists {
			return devices
		}
	}
	return nil
}

func openConfigWindow(ctx AppContext) {
	configWindow := ctx.App.NewWindow("easyChip " + ctx.MCUConfig)
	configWindow.Resize(fyne.NewSize(800, 600))

	menuBar := container.NewHBox( // 顶部按钮
		widget.NewButton("Pinout&Configuration", nil),
		widget.NewButton("Project Manager", func() {}),
	)

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current directory:", err)
		return
	}
	rootDir := filepath.Join(currentDir, "Library")
	filePath := filepath.Join(rootDir + "/")
	fmt.Println(filePath)
	fmt.Println(filePath + ctx.McuConfigPath[ctx.MCUConfig] + "peripheral.json")
	jsonData, err := os.ReadFile(filePath + ctx.McuConfigPath[ctx.MCUConfig] + "peripheral.json")
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}

	var p Peripheral
	err = json.Unmarshal(jsonData, &p)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
		return
	}

	accordion := widget.NewAccordion()

	for _, list := range GetCategoryNames(p) {
		// 为每个类别创建一个新的容器
		buttonContainer := container.NewVBox()

		// 获取该类别的所有外设
		peripherals := GetPeripheralsByCategory(p, list)

		// 为每个外设创建按钮并添加到容器
		for _, kit := range peripherals {
			// 使用闭包来捕获当前外设值
			peripheral := kit // 重要：创建一个局部变量来捕获当前的值
			btn := widget.NewButton(peripheral, func() {
				println("按钮被点击:", peripheral)
			})
			buttonContainer.Add(btn) // 将按钮添加到当前类别的容器中
		}

		// 创建该类别的手风琴项并添加到手风琴中
		item := widget.NewAccordionItem(list, buttonContainer)
		accordion.Append(item)
	}

	accordion.MultiOpen = false // 允许多个展开

	// 如果需要控制内容宽度但允许高度自动调整
	leftScroll := container.NewGridWrap(fyne.NewSize(230, 500), container.NewScroll(accordion))

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
