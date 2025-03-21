package GUI

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strings"
)

func openNewWindow(a fyne.App, onUpdate func(string)) {
	newWindow := a.NewWindow("选择MCU")
	newWindow.Resize(fyne.NewSize(800, 600))

	// 变量存储用户选中的 MCU 型号
	selectedMCU := ""

	// 存储收藏的MCU
	favoriteMCUs := make(map[string]bool)

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

	// ✅ 添加开发环境支持数据
	mcuSupport := map[string]map[string]bool{
		"A-100": {"KEIL5": true, "IAR": true, "Arduino": false},
		"A-200": {"KEIL5": true, "IAR": true, "Arduino": true},
		"A-300": {"KEIL5": true, "IAR": false, "Arduino": false},
		"B-100": {"KEIL5": true, "IAR": true, "Arduino": false},
		"B-200": {"KEIL5": true, "IAR": true, "Arduino": true},
		"C-100": {"KEIL5": true, "IAR": false, "Arduino": true},
		"C-200": {"KEIL5": true, "IAR": true, "Arduino": false},
		"C-300": {"KEIL5": false, "IAR": true, "Arduino": false},
		"C-400": {"KEIL5": true, "IAR": false, "Arduino": false},
		"D-100": {"KEIL5": false, "IAR": false, "Arduino": true},
		"D-200": {"KEIL5": true, "IAR": true, "Arduino": false},
		"D-300": {"KEIL5": true, "IAR": false, "Arduino": false},
		"D-400": {"KEIL5": false, "IAR": true, "Arduino": true},
		"D-500": {"KEIL5": true, "IAR": true, "Arduino": true},
	}

	// 创建开始按钮，并默认禁用
	startButton := widget.NewButton("开始", func() {
		fmt.Println("程序开始运行，选中的 MCU: ", selectedMCU)
		onUpdate(selectedMCU) // 将选中的 MCU 传回主窗口
		newWindow.Close()     // 关闭子窗口
	})
	startButton.Disable() // 初始时禁用

	// ✅ 右侧信息展示区域
	infoLabel := widget.NewLabel("MCU 详情信息将显示在这里...")
	separator := widget.NewSeparator() // 分割线
	infoContainer := container.NewVBox(separator, infoLabel)

	// ✅ 右侧 MCU 详情列表（初始为空）
	rightData := []string{}

	// 创建带有星星图标的自定义列表项
	rightList := widget.NewList(
		func() int { return len(rightData) }, // 数据数量
		func() fyne.CanvasObject {
			// 创建一个水平容器包含星星按钮和文本标签
			starButton := widget.NewButton("☆", nil)
			starButton.Importance = widget.LowImportance
			starButton.Resize(fyne.NewSize(30, 30))
			label := widget.NewLabel("")
			return container.NewHBox(starButton, label)
		},
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			// 获取容器中的按钮和标签
			hbox := obj.(*fyne.Container)
			starButton := hbox.Objects[0].(*widget.Button)
			label := hbox.Objects[1].(*widget.Label)

			mcu := rightData[i]
			label.SetText(mcu)

			// 设置星星按钮的文本，根据收藏状态显示不同图标
			if favoriteMCUs[mcu] {
				starButton.SetText("★") // 实心星星表示已收藏
			} else {
				starButton.SetText("☆") // 空心星星表示未收藏
			}

			// 设置星星按钮的点击事件
			starButton.OnTapped = func() {
				// 切换收藏状态
				favoriteMCUs[mcu] = !favoriteMCUs[mcu]

				// 更新按钮显示
				if favoriteMCUs[mcu] {
					starButton.SetText("★")
				} else {
					starButton.SetText("☆")
				}
			}
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

	// 创建搜索框
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("搜索MCU...")

	// 使用固定宽度的容器包装搜索框
	searchEntryContainer := container.NewGridWrap(
		fyne.NewSize(200, 40), // 宽度为200，高度为40
		searchEntry,
	)

	// 添加筛选条件复选框
	keilCheck := widget.NewCheck("KEIL5", nil)
	iarCheck := widget.NewCheck("IAR", nil)
	arduinoCheck := widget.NewCheck("Arduino", nil)
	favoriteCheck := widget.NewCheck("仅显示收藏", nil) // 添加收藏筛选选项

	// 创建筛选条件容器
	filterContainer := container.NewHBox(
		widget.NewLabel("筛选:"),
		keilCheck,
		iarCheck,
		arduinoCheck,
		favoriteCheck,
	)

	// 搜索按钮
	searchButton := widget.NewButton("搜索", func() {
		searchText := searchEntry.Text

		// 获取筛选条件
		keilFilter := keilCheck.Checked
		iarFilter := iarCheck.Checked
		arduinoFilter := arduinoCheck.Checked
		favoriteFilter := favoriteCheck.Checked

		// 如果搜索框为空且没有选择筛选条件，恢复默认视图
		if searchText == "" && !keilFilter && !iarFilter && !arduinoFilter && !favoriteFilter {
			leftList.Select(0)
			return
		}

		// 清空当前选择
		leftList.UnselectAll()

		// 搜索匹配的MCU
		var matchedMCUs []string
		for _, category := range categories {
			for _, mcu := range mcuData[category] {
				// 检查MCU名称是否匹配搜索文本
				nameMatch := searchText == "" || strings.Contains(strings.ToLower(mcu), strings.ToLower(searchText))

				// 如果名称不匹配，则跳过此MCU
				if !nameMatch {
					continue
				}

				// 检查是否满足收藏筛选条件
				if favoriteFilter && !favoriteMCUs[mcu] {
					continue
				}

				// 检查是否满足筛选条件
				// 如果没有选择筛选条件，则视为通过筛选
				filterMatch := true

				// 如果选择了筛选条件，则检查MCU是否支持这些条件
				if keilFilter || iarFilter || arduinoFilter {
					// 检查MCU是否在支持数据中
					if support, exists := mcuSupport[mcu]; exists {
						// 初始假设为通过筛选
						filterMatch = true

						// 检查各个筛选条件
						if keilFilter && !support["KEIL5"] {
							filterMatch = false
						}
						if iarFilter && !support["IAR"] {
							filterMatch = false
						}
						if arduinoFilter && !support["Arduino"] {
							filterMatch = false
						}
					} else {
						// 如果MCU不在支持数据中，则默认为不通过筛选
						filterMatch = false
					}
				}

				// 如果同时匹配名称和筛选条件，则添加到结果中
				if nameMatch && filterMatch {
					matchedMCUs = append(matchedMCUs, mcu)
				}
			}
		}

		// 更新右侧列表显示搜索结果
		rightData = matchedMCUs
		rightList.Refresh()

		// 更新信息标签
		if len(matchedMCUs) == 0 {
			infoLabel.SetText("未找到匹配的MCU")
		} else {
			infoLabel.SetText(fmt.Sprintf("找到 %d 个匹配的MCU", len(matchedMCUs)))
		}
	})

	// 搜索区域布局 - 搜索框和筛选条件在同一行
	searchFilterContainer := container.NewHBox(
		searchEntryContainer,
		filterContainer,
		searchButton,
	)

	// ✅ 右侧操作区域，包含"开始"按钮
	rightControls := container.NewBorder(nil, nil, nil,
		container.NewGridWrap(fyne.NewSize(100, 40), startButton))

	// ✅ 顶部工具栏 - 包含搜索框、筛选条件和开始按钮
	topBar := container.NewBorder(nil, nil, searchFilterContainer, rightControls)

	// ✅ 当左侧分类被点击时，更新右侧 MCU 列表
	leftList.OnSelected = func(id widget.ListItemID) {
		category := categories[id]
		rightData = mcuData[category] // 更新数据
		rightList.Refresh()           // 刷新右侧列表
		rightList.UnselectAll()       // 清除右侧选择
		startButton.Disable()         // 禁用开始按钮
		infoLabel.SetText("MCU 详情信息将显示在这里...")
	}

	// ✅ 当右侧 MCU 详情列表被点击时，存储选中的 MCU 型号并更新详情信息
	rightList.OnSelected = func(id widget.ListItemID) {
		if id < len(rightData) {
			selectedMCU = rightData[id]
			fmt.Println("已选择 MCU: ", selectedMCU)

			// 构建包含开发环境支持信息的详情字符串
			var detailText string
			if detail, exists := mcuDetails[selectedMCU]; exists {
				detailText = detail
			} else {
				detailText = "未知的 MCU 详情"
			}

			// 添加开发环境支持信息
			if support, exists := mcuSupport[selectedMCU]; exists {
				detailText += "\n\n支持开发环境: "
				var supportedEnvs []string

				if support["KEIL5"] {
					supportedEnvs = append(supportedEnvs, "KEIL5")
				}
				if support["IAR"] {
					supportedEnvs = append(supportedEnvs, "IAR")
				}
				if support["Arduino"] {
					supportedEnvs = append(supportedEnvs, "Arduino")
				}

				if len(supportedEnvs) > 0 {
					detailText += strings.Join(supportedEnvs, ", ")
				} else {
					detailText += "无"
				}
			}

			// 添加收藏状态到详情
			if favoriteMCUs[selectedMCU] {
				detailText += "\n\n✓ 已收藏"
			}

			infoLabel.SetText(detailText)
			startButton.Enable() // 选定 MCU 后启用按钮
		}
	}

	// 添加搜索框的回车键监听
	searchEntry.OnSubmitted = func(text string) {
		searchButton.OnTapped()
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
