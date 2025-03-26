package GUI

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// 定义结构体匹配 JSON 结构
type MCUInfo struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Type   string `json:"type"`
	Specs  struct {
		Flash int `json:"flash"` // Flash 大小（KB）
		Core  int `json:"core"`  // Core 频率（MHz）
		SRAM  int `json:"sram"`  // SRAM 大小（KB）
	} `json:"specs"`
	IDE struct {
		IAR     bool `json:"IAR"`
		KEIL    bool `json:"KEIL"`
		Arduino bool `json:"Arduino"`
	} `json:"IDE"`
}

type Template struct {
	Path string `json:"path"`
}

type Author struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

// 总的 JSON 结构
type Data struct {
	Info     MCUInfo  `json:"info"`
	Template Template `json:"template"`
	Author   Author   `json:"author"`
}

func openSelectWindow(a fyne.App, onUpdate func(string)) {
	newWindow := a.NewWindow("选择MCU")
	newWindow.Resize(fyne.NewSize(800, 600))

	// 变量存储用户选中的 MCU 型号
	selectedMCU := ""

	// 存储收藏的MCU
	favoriteMCUs := make(map[string]bool)

	// ✅ MCU 分类列表 - 添加"全部"选项
	categories := []string{"全部", "MCU A系列", "MCU B系列", "MCU C系列", "MCU D系列"}
	mcuData := map[string][]string{
		"MCU A系列": {"A-100", "A-200", "A-300"},
		"MCU B系列": {"B-100", "B-200"},
		"MCU C系列": {"C-100", "C-200", "C-300", "C-400"},
		"MCU D系列": {"D-100", "D-200", "D-300", "D-400", "D-500"},
	}

	var allMCUs []string
	for _, mcuList := range mcuData {
		allMCUs = append(allMCUs, mcuList...)
	}

	// 按字母数字顺序排序
	sort.Strings(allMCUs)

	// 将全部MCU添加到MCU数据中
	mcuData["全部"] = allMCUs

	// ✅ MCU 详情数据
	mcuDetails := map[string]string{
		"A-100": "A-100: 32-bit, 48MHz, 128KB Flash",
		"A-200": "A-200: 32-bit, 72MHz, 256KB Flash",
		"A-300": "A-300: 32-bit, 96MHz, 512KB Flash",
		"B-100": "B-100: 16-bit, 24MHz, 64KB Flash",
		"B-200": "B-200: 16-bit, 48MHz, 128KB Flash",
		"C-100": "C-100: ARM Cortex-M3, 120MHz, 512KB Flash",
		"C-200": "C-200: ARM Cortex-M4, 144MHz, 768KB Flash",
		"C-300": "C-300: ARM Cortex-M4, 168MHz, 1MB Flash",
		"C-400": "C-400: ARM Cortex-M7, 216MHz, 2MB Flash",
		"D-100": "D-100: RISC-V, 160MHz, 1MB Flash",
		"D-200": "D-200: RISC-V, 180MHz, 2MB Flash",
		"D-300": "D-300: RISC-V, 200MHz, 4MB Flash",
		"D-400": "D-400: RISC-V, 240MHz, 8MB Flash",
		"D-500": "D-500: RISC-V, 300MHz, 16MB Flash",
	}

	// ✅ 添加 MCU 规格信息映射
	mcuSpecs := map[string]struct {
		Flash int
		Core  int
		SRAM  int
	}{
		"A-100": {Flash: 128, Core: 48, SRAM: 32},
		"A-200": {Flash: 256, Core: 72, SRAM: 64},
		"A-300": {Flash: 512, Core: 96, SRAM: 128},
		"B-100": {Flash: 64, Core: 24, SRAM: 16},
		"B-200": {Flash: 128, Core: 48, SRAM: 32},
		"C-100": {Flash: 512, Core: 120, SRAM: 128},
		"C-200": {Flash: 768, Core: 144, SRAM: 256},
		"C-300": {Flash: 1024, Core: 168, SRAM: 512},
		"C-400": {Flash: 2048, Core: 216, SRAM: 1024},
		"D-100": {Flash: 1024, Core: 160, SRAM: 256},
		"D-200": {Flash: 2048, Core: 180, SRAM: 512},
		"D-300": {Flash: 4096, Core: 200, SRAM: 1024},
		"D-400": {Flash: 8192, Core: 240, SRAM: 2048},
		"D-500": {Flash: 16384, Core: 300, SRAM: 4096},
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

	// 指定要遍历的文件夹路径
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current directory:", err)
		return
	}
	rootDir := filepath.Join(currentDir, "Library")

	// 读取文件夹内容
	entries, err := os.ReadDir(rootDir)
	if err != nil {
		fmt.Println("读取目录失败:", err)
		return
	}

	// 遍历所有文件和文件夹
	fmt.Println("文件夹列表:")
	for _, entry := range entries {
		if entry.IsDir() { // 只打印文件夹
			fmt.Println(rootDir + "/" + entry.Name())
			files, err := os.ReadDir(rootDir + "/" + entry.Name())
			if err != nil {
				fmt.Println("无法读取文件夹:", err)
				return
			}
			for _, file := range files {
				if filepath.Ext(file.Name()) == ".json" { // 只处理 JSON 文件
					filePath := filepath.Join(rootDir+"/"+entry.Name(), file.Name())
					fmt.Println(filePath)
					data, err := os.ReadFile(filePath)
					if err != nil {
						fmt.Println("无法读取文件:", filePath, err)
						continue
					}

					var jsonData Data
					err = json.Unmarshal(data, &jsonData)
					if err != nil {
						fmt.Println("JSON 解析失败:", filePath, err)
						continue
					}

					// 更新数据
					categories = append(categories, jsonData.Info.Type)
					mcuData[jsonData.Info.Type] = append(mcuData[jsonData.Info.Type], jsonData.Info.Name)
					mcuDetails[jsonData.Info.Name] = jsonData.Info.Detail

					// 更新规格信息
					mcuSpecs[jsonData.Info.Name] = struct {
						Flash int
						Core  int
						SRAM  int
					}{
						Flash: jsonData.Info.Specs.Flash,
						Core:  jsonData.Info.Specs.Core,
						SRAM:  jsonData.Info.Specs.SRAM,
					}

					// 更新支持的开发环境
					mcuSupport[jsonData.Info.Name] = make(map[string]bool)
					mcuSupport[jsonData.Info.Name]["IAR"] = jsonData.Info.IDE.IAR
					mcuSupport[jsonData.Info.Name]["KEIL5"] = jsonData.Info.IDE.KEIL
					mcuSupport[jsonData.Info.Name]["Arduino"] = jsonData.Info.IDE.Arduino

					fmt.Printf("文件: %s, MCU 名称: %s\n", file.Name(), jsonData.Info.Name)
				}
			}
		}
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
			starButton := widget.NewButton("☆", nil)
			starButton.Importance = widget.LowImportance
			starButton.Resize(fyne.NewSize(30, 30))
			label := widget.NewLabel("")
			return container.NewHBox(starButton, label)
		},
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			hbox := obj.(*fyne.Container)
			starButton := hbox.Objects[0].(*widget.Button)
			label := hbox.Objects[1].(*widget.Label)

			mcu := rightData[i]
			label.SetText(mcu)

			if favoriteMCUs[mcu] {
				starButton.SetText("★")
			} else {
				starButton.SetText("☆")
			}

			starButton.OnTapped = func() {
				favoriteMCUs[mcu] = !favoriteMCUs[mcu]

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
			leftList.Select(0) // 选择"全部"分类
			return
		}

		// 清空当前选择
		leftList.UnselectAll()

		// 使用全部MCU数据作为基础进行搜索
		var matchedMCUs []string
		for _, mcu := range allMCUs {
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

		// 更新信息标签
		if category == "全部" {
			infoLabel.SetText(fmt.Sprintf("全部MCU，共 %d 个", len(allMCUs)))
		} else {
			infoLabel.SetText("MCU 详情信息将显示在这里...")
		}
	}

	// ✅ 当右侧 MCU 详情列表被点击时，存储选中的 MCU 型号并更新详情信息
	rightList.OnSelected = func(id widget.ListItemID) {
		if id < len(rightData) {
			selectedMCU = rightData[id]
			fmt.Println("已选择 MCU: ", selectedMCU)

			var detailText string
			if detail, exists := mcuDetails[selectedMCU]; exists {
				detailText = detail
			} else {
				detailText = "未知的 MCU 详情"
			}

			// 添加规格信息
			if specs, exists := mcuSpecs[selectedMCU]; exists {
				detailText += fmt.Sprintf("\n\nMCU 规格信息:\n")
				detailText += fmt.Sprintf("Flash 大小: %d KB\n", specs.Flash)
				detailText += fmt.Sprintf("Core 频率: %d MHz\n", specs.Core)
				detailText += fmt.Sprintf("SRAM 大小: %d KB", specs.SRAM)
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

	// 默认选择"全部"分类
	leftList.Select(0)

	newWindow.Show()
}
