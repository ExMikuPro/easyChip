package template

type IarEwwType struct {
	Projects []string
}
type IarEwpType struct {
	Release     bool
	CCOptLevel  int
	IccLang     int
	IccCDialect int
}

type IarEwpCCOptLevelType struct { // 代码优化等级
	NoOptimization     int
	LowOptimization    int
	MediumOptimization int
	HighOptimization   int
}

type IarEwpIccLang struct { // 编译器语言模式
	CLang   int
	CppLang int
}

type IarEwpIccCDialect struct { // C 语言标准版本
	ANSI int
	C99  int
	C11  int
}

var (
	CCOptLevel = IarEwpCCOptLevelType{
		NoOptimization: 0, LowOptimization: 1, MediumOptimization: 2, HighOptimization: 3,
	}
	IccLang     = IarEwpIccLang{CLang: 0, CppLang: 1}
	IccCDialect = IarEwpIccCDialect{ANSI: 0, C99: 1, C11: 2}
)

//func Iar_ewp() {
//	data := IarEwpType{false, CCOptLevel.LowOptimization, IccLang.CLang, IccCDialect.C11}
//}
