package template

import (
	"os"
	"text/template"
)

// Iar_eww 生成 IAR Embedded Workbench Workspace (EWW) 文件
// name: 项目名称
// path: 生成路径
// data: 需要填充的 IAR EWW 数据结构
func Iar_eww(name string, path string, data IarEwwType) {
	// 解析 IAR 工作区 (EWW) 模板
	tmpl, err := template.ParseFiles("template/raw/iar/iar-eww.txt")
	if err != nil {
		panic(err)
	}

	// 创建 .eww 输出文件
	outputFile, err := os.Create(path + name + ".eww")
	if err != nil {
		panic(err)
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			// 处理文件关闭时的错误
		}
	}(outputFile)

	// 执行模板并填充数据
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		panic(err)
	}
}

// Iar_ewp 生成 IAR Embedded Workbench Project (EWP) 文件，并创建必要的目录结构
// name: 项目名称
// path: 生成路径
// data: 需要填充的 IAR EWP 数据结构
func Iar_ewp(name string, path string, data IarEwpType) {
	// 解析 IAR 项目 (EWP) 模板
	tmpl, err := template.ParseFiles("template/raw/cc2530/cc2530-iar-ewp.txt")
	if err != nil {
		panic(err)
	}

	// 创建项目目录
	err = os.MkdirAll(path+"/"+name, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// 创建代码目录
	err = os.MkdirAll(path+"/"+name+"/code", os.ModePerm)
	if err != nil {
		panic(err)
	}

	// 生成 main.c 文件
	Iar_main(name, path)

	// 创建 .ewp 输出文件
	outputFile, err := os.Create(path + "/" + name + "/" + name + ".ewp")
	if err != nil {
		panic(err)
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			// 处理文件关闭时的错误
		}
	}(outputFile)

	// 执行模板并填充数据
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		panic(err)
	}
}

// Iar_main 生成 IAR 项目的 main.c 文件
// name: 项目名称
// path: 生成路径
func Iar_main(name string, path string) {
	// 解析 main.c 模板
	tmpl, err := template.ParseFiles("template/raw/cc2530/cc2530-iar-main.c.txt")
	if err != nil {
		panic(err)
	}

	// 创建 main.c 文件
	outputFile, err := os.Create(path + "/" + name + "/code/main.c")
	if err != nil {
		panic(err)
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			// 处理文件关闭时的错误
		}
	}(outputFile)

	// 执行模板生成 main.c
	err = tmpl.Execute(outputFile, nil)
	if err != nil {
		panic(err)
	}
}

// APPConfigECH 预留的配置文件处理函数
func APPConfigECH() {
	// 未来可添加配置解析或生成逻辑
}
