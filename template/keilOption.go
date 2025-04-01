package template

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
)

// KeilProject 表示Keil5项目的整体结构
type KeilProject struct {
	XMLName                      xml.Name `xml:"Project" json:"Project"`
	XMLNSXSI                     string   `xml:"xmlns:xsi,attr" json:"attr_XMLNSXSI"`
	XSINoNamespaceSchemaLocation string   `xml:"xsi:noNamespaceSchemaLocation,attr" json:"attr_XSINoNamespaceSchemaLocation"`

	SchemaVersion string    `xml:"SchemaVersion" json:"SchemaVersion"`
	Header        string    `xml:"Header" json:"Header"`
	Targets       Targets   `xml:"Targets" json:"Targets"` // 包含的目标配置列表
	RTE           RTE       `xml:"RTE" json:"RTE"`
	LayerInfo     LayerInfo `xml:"LayerInfo" json:"LayerInfo"`
}

type Targets struct {
	Target Target `xml:"Target" json:"Target"`
}

// Targets 表示一个编译目标配置
type Target struct {
	TargetName    string       `xml:"TargetName" json:"TargetName"`       // 目标名称，如"Target 1"
	ToolsetNumber string       `xml:"ToolsetNumber" json:"ToolsetNumber"` // 工具集编号，如"0x4"代表ARM-ADS工具链
	ToolsetName   string       `xml:"ToolsetName" json:"ToolsetName"`     // 工具集名称，如"ARM-ADS"
	PCCUsed       string       `xml:"pCCUsed" json:"PCCUsed"`             // 使用的编译器版本，如"6190000::V6.19::ARMCLANG"
	UAC6          string       `xml:"uAC6" json:"UAC6"`                   // ARM编译器6标志，"1"表示使用ARMCLANG
	TargetOption  TargetOption `xml:"TargetOption" json:"TargetOption"`   // 目标详细选项配置
	Groups        []Group      `xml:"Groups>Group" json:"Groups"`         // 项目文件组织结构
}

// TargetOption 包含目标的所有配置选项
type TargetOption struct {
	TargetCommonOption TargetCommonOption `xml:"TargetCommonOption" json:"TargetCommonOption"` // 通用目标选项
	CommonProperty     CommonProperty     `xml:"CommonProperty" json:"CommonProperty"`         // 通用属性设置
	DllOption          DllOption          `xml:"DllOption" json:"DllOption"`                   // DLL选项，用于调试
	DebugOption        DebugOption        `xml:"DebugOption" json:"DebugOption"`               // 调试选项
	Utilities          Utilities          `xml:"Utilities" json:"Utilities"`                   // 实用工具选项，如Flash下载设置
	TargetArmAds       TargetArmAds       `xml:"TargetArmAds" json:"TargetArmAds"`             // ARM特定的编译器和链接器设置

}

// TargetCommonOption 包含目标的通用选项设置
type TargetCommonOption struct {
	Device                string        `xml:"Device" json:"Device"`   // 目标设备名称，如"AIR001Dev"
	Vendor                string        `xml:"Vendor" json:"Vendor"`   // 设备厂商，如"Generic"
	PackID                string        `xml:"PackID" json:"PackID"`   // 设备包ID，如"Keil.AIR001_DFP.1.1.1"
	PackURL               string        `xml:"PackURL" json:"PackURL"` // 设备包下载URL
	Cpu                   string        `xml:"Cpu" json:"Cpu"`         // CPU配置，包含内存布局、CPU类型和时钟频率
	FlashUtilSpec         string        `xml:"FlashUtilSpec" json:"FlashUtilSpec"`
	StartupFile           string        `xml:"StartupFile" json:"StartupFile"`
	FlashDriverDll        string        `xml:"FlashDriverDll" json:"FlashDriverDll"` // Flash下载器DLL配置
	DeviceId              string        `xml:"DeviceId" json:"DeviceId"`             // 设备ID
	RegisterFile          string        `xml:"RegisterFile" json:"RegisterFile"`     // 寄存器定义文件
	MemoryEnv             string        `xml:"MemoryEnv" json:"MemoryEnv"`
	Cmp                   string        `xml:"Cmp" json:"Cmp"`
	Asm                   string        `xml:"Asm" json:"Asm"`
	Linker                string        `xml:"Linker" json:"Linker"`
	OHString              string        `xml:"OHString" json:"OHString"`
	InfinionOptionDll     string        `xml:"InfinionOptionDll" json:"InfinionOptionDll"`
	SLE66CMisc            string        `xml:"SLE66CMisc" json:"SLE66CMisc"`
	SLE66AMisc            string        `xml:"SLE66AMisc" json:"SLE66AMisc"`
	SLE66LinkerMisc       string        `xml:"SLE66LinkerMisc" json:"SLE66LinkerMisc"`
	SFDFile               string        `xml:"SFDFile" json:"SFDFile"` // SVD描述文件，用于调试
	BCustSvd              string        `xml:"bCustSvd" json:"BCustSvd"`
	UseEnv                string        `xml:"UseEnv" json:"UseEnv"`
	BinPath               string        `xml:"BinPath" json:"BinPath"`
	IncludePath           string        `xml:"IncludePath" json:"IncludePath"`
	LibPath               string        `xml:"LibPath" json:"LibPath"`
	RegisterFilePath      string        `xml:"RegisterFilePath" json:"RegisterFilePath"`
	DBRegisterFilePath    string        `xml:"DBRegisterFilePath" json:"DBRegisterFilePath"`
	TargetStatus          TargetStatus  `xml:"TargetStatus" json:"TargetStatus"`
	OutputDirectory       string        `xml:"OutputDirectory" json:"OutputDirectory"`   // 输出目录路径，如".\Objects\"
	OutputName            string        `xml:"OutputName" json:"OutputName"`             // 输出文件名，如"Example_HAL"
	CreateExecutable      string        `xml:"CreateExecutable" json:"CreateExecutable"` // 是否创建可执行文件，"1"表示是
	CreateLib             string        `xml:"CreateLib" json:"CreateLib"`
	CreateHexFile         string        `xml:"CreateHexFile" json:"CreateHexFile"`         // 是否创建HEX文件，"1"表示是
	DebugInformation      string        `xml:"DebugInformation" json:"DebugInformation"`   // 是否包含调试信息，"1"表示是
	BrowseInformation     string        `xml:"BrowseInformation" json:"BrowseInformation"` // 是否生成浏览信息，"1"表示是
	ListingPath           string        `xml:"ListingPath" json:"ListingPath"`             // 列表文件输出路径
	HexFormatSelection    string        `xml:"HexFormatSelection" json:"HexFormatSelection"`
	Merge32K              string        `xml:"Merge32K" json:"Merge32K"`
	CreateBatchFile       string        `xml:"CreateBatchFile" json:"CreateBatchFile"`
	BeforeCompile         BeforeCompile `xml:"BeforeCompile" json:"BeforeCompile"`
	BeforeMake            BeforeMake    `xml:"BeforeMake" json:"BeforeMake"`
	AfterMake             AfterMake     `xml:"AfterMake" json:"AfterMake"`
	SelectedForBatchBuild string        `xml:"SelectedForBatchBuild" json:"SelectedForBatchBuild"`
	SVCSIdString          string        `xml:"SVCSIdString" json:"SVCSIdString"`
}

type TargetStatus struct {
	Error        string `xml:"Error" json:"Error"`               // 错误状态
	ExitCodeStop string `xml:"ExitCodeStop" json:"ExitCodeStop"` // 退出代码停止标志
	ButtonStop   string `xml:"ButtonStop" json:"ButtonStop"`     // 按钮停止标志
	NotGenerated string `xml:"NotGenerated" json:"NotGenerated"` // 未生成标志
	InvalidFlash string `xml:"InvalidFlash" json:"InvalidFlash"` // 无效Flash标志
}

type BeforeCompile struct {
	RunUserProg1       string `xml:"RunUserProg1" json:"RunUserProg1"`
	RunUserProg2       string `xml:"RunUserProg2" json:"RunUserProg2"`
	UserProg1Name      string `xml:"UserProg1Name" json:"UserProg1Name"`
	UserProg2Name      string `xml:"UserProg2Name" json:"UserProg2Name"`
	UserProg1Dos16Mode string `xml:"UserProg1Dos16Mode" json:"UserProg1Dos16Mode"`
	UserProg2Dos16Mode string `xml:"UserProg2Dos16Mode" json:"UserProg2Dos16Mode"`
	NStopU1X           string `xml:"nStopU1X" json:"NStopU1X"`
	NStopU2X           string `xml:"nStopU2X" json:"NStopU2X"`
}
type BeforeMake struct {
	RunUserProg1       string `xml:"RunUserProg1" json:"RunUserProg1"`
	RunUserProg2       string `xml:"RunUserProg2" json:"RunUserProg2"`
	UserProg1Name      string `xml:"UserProg1Name" json:"UserProg1Name"`
	UserProg2Name      string `xml:"UserProg2Name" json:"UserProg2Name"`
	UserProg1Dos16Mode string `xml:"UserProg1Dos16Mode" json:"UserProg1Dos16Mode"`
	UserProg2Dos16Mode string `xml:"UserProg2Dos16Mode" json:"UserProg2Dos16Mode"`
	NStopB1X           string `xml:"nStopB1X" json:"NStopB1X"`
	NStopB2X           string `xml:"nStopB2X" json:"NStopB2X"`
}

type AfterMake struct {
	RunUserProg1       string `xml:"RunUserProg1" json:"RunUserProg1"`
	RunUserProg2       string `xml:"RunUserProg2" json:"RunUserProg2"`
	UserProg1Name      string `xml:"UserProg1Name" json:"UserProg1Name"`
	UserProg2Name      string `xml:"UserProg2Name" json:"UserProg2Name"`
	UserProg1Dos16Mode string `xml:"UserProg1Dos16Mode" json:"UserProg1Dos16Mode"`
	UserProg2Dos16Mode string `xml:"UserProg2Dos16Mode" json:"UserProg2Dos16Mode"`
	NStopA1X           string `xml:"nStopA1X" json:"NStopA1X"`
	NStopA2X           string `xml:"nStopA2X" json:"NStopA2X"`
}

// CommonProperty 包含通用属性设置
type CommonProperty struct {
	UseCPPCompiler        string `xml:"UseCPPCompiler" json:"UseCPPCompiler"`             // 是否使用C++编译器，"0"表示不使用
	RVCTCodeConst         string `xml:"RVCTCodeConst" json:"RVCTCodeConst"`               // 代码常量设置
	RVCTZI                string `xml:"RVCTZI" json:"RVCTZI"`                             // 零初始化设置
	RVCTOtherData         string `xml:"RVCTOtherData" json:"RVCTOtherData"`               // 其他数据设置
	ModuleSelection       string `xml:"ModuleSelection" json:"ModuleSelection"`           // 模块选择
	IncludeInBuild        string `xml:"IncludeInBuild" json:"IncludeInBuild"`             // 是否包含在构建中，"1"表示是
	AlwaysBuild           string `xml:"AlwaysBuild" json:"AlwaysBuild"`                   // 是否总是构建，"0"表示否
	GenerateAssemblyFile  string `xml:"GenerateAssemblyFile" json:"GenerateAssemblyFile"` // 是否生成汇编文件，"0"表示否
	AssembleAssemblyFile  string `xml:"AssembleAssemblyFile" json:"AssembleAssemblyFile"`
	PublicsOnly           string `xml:"PublicsOnly" json:"PublicsOnly"`
	StopOnExitCode        string `xml:"StopOnExitCode" json:"StopOnExitCode"`
	CustomArgument        string `xml:"CustomArgument" json:"CustomArgument"`
	IncludeLibraryModules string `xml:"IncludeLibraryModules" json:"IncludeLibraryModules"`
	ComprImg              string `xml:"ComprImg" json:"ComprImg"` // 图像压缩设置
}

// DllOption 包含用于调试的DLL设置
type DllOption struct {
	SimDllName            string `xml:"SimDllName" json:"SimDllName"`                       // 模拟器DLL名称
	SimDllArguments       string `xml:"SimDllArguments" json:"SimDllArguments"`             // 模拟器DLL参数
	SimDlgDll             string `xml:"SimDlgDll" json:"SimDlgDll"`                         // 模拟器对话框DLL
	SimDlgDllArguments    string `xml:"SimDlgDllArguments" json:"SimDlgDllArguments"`       // 模拟器对话框DLL参数
	TargetDllName         string `xml:"TargetDllName" json:"TargetDllName"`                 // 目标DLL名称
	TargetDllArguments    string `xml:"TargetDllArguments" json:"TargetDllArguments"`       // 目标DLL参数
	TargetDlgDll          string `xml:"TargetDlgDll" json:"TargetDlgDll"`                   // 目标对话框DLL
	TargetDlgDllArguments string `xml:"TargetDlgDllArguments" json:"TargetDlgDllArguments"` // 目标对话框DLL参数
}

// DebugOption 包含调试选项设置
type DebugOption struct {
	OPTHX OPTHX `xml:"OPTHX" json:"OPTHX"` // 优化十六进制设置
}

type OPTHX struct {
	HexSelection        string `xml:"HexSelection" json:"HexSelection"`               // 十六进制选择
	HexRangeLowAddress  string `xml:"HexRangeLowAddress" json:"HexRangeLowAddress"`   // 十六进制范围低地址
	HexRangeHighAddress string `xml:"HexRangeHighAddress" json:"HexRangeHighAddress"` // 十六进制范围高地址
	HexOffset           string `xml:"HexOffset" json:"HexOffset"`                     // 十六进制偏移量
	Oh166RecLen         string `xml:"Oh166RecLen" json:"Oh166RecLen"`
}

// Utilities 包含实用工具设置
type Utilities struct {
	Flash1 struct {
		UseTargetDll               string `xml:"UseTargetDll" json:"UseTargetDll"`                             // 是否使用目标DLL，"1"表示是
		UseExternalTool            string `xml:"UseExternalTool" json:"UseExternalTool"`                       // 是否使用外部工具，"0"表示否
		RunIndependent             string `xml:"RunIndependent" json:"RunIndependent"`                         // 是否独立运行，"0"表示否
		UpdateFlashBeforeDebugging string `xml:"UpdateFlashBeforeDebugging" json:"UpdateFlashBeforeDebugging"` // 调试前更新Flash，"1"表示是
		Capability                 string `xml:"Capability" json:"Capability"`                                 // 能力设置
		DriverSelection            string `xml:"DriverSelection" json:"DriverSelection"`                       // 驱动选择
	} `xml:"Flash1" json:"Flash1"` // Flash工具1配置
	BUseTDR string `xml:"bUseTDR" json:"BUseTDR"`
	Flash2  struct {
		UseTargetDll               string `xml:"UseTargetDll" json:"UseTargetDll"`                             // 是否使用目标DLL，"1"表示是
		UseExternalTool            string `xml:"UseExternalTool" json:"UseExternalTool"`                       // 是否使用外部工具，"0"表示否
		RunIndependent             string `xml:"RunIndependent" json:"RunIndependent"`                         // 是否独立运行，"0"表示否
		UpdateFlashBeforeDebugging string `xml:"UpdateFlashBeforeDebugging" json:"UpdateFlashBeforeDebugging"` // 调试前更新Flash，"1"表示是
		Capability                 string `xml:"Capability" json:"Capability"`                                 // 能力设置
		DriverSelection            string `xml:"DriverSelection" json:"DriverSelection"`                       // 驱动选择
	} `xml:"Flash2" json:"Flash2"` // Flash工具1配置
	Flash3 struct {
		UseTargetDll               string `xml:"UseTargetDll" json:"UseTargetDll"`                             // 是否使用目标DLL，"1"表示是
		UseExternalTool            string `xml:"UseExternalTool" json:"UseExternalTool"`                       // 是否使用外部工具，"0"表示否
		RunIndependent             string `xml:"RunIndependent" json:"RunIndependent"`                         // 是否独立运行，"0"表示否
		UpdateFlashBeforeDebugging string `xml:"UpdateFlashBeforeDebugging" json:"UpdateFlashBeforeDebugging"` // 调试前更新Flash，"1"表示是
		Capability                 string `xml:"Capability" json:"Capability"`                                 // 能力设置
		DriverSelection            string `xml:"DriverSelection" json:"DriverSelection"`                       // 驱动选择
	} `xml:"Flash3" json:"Flash3"` // Flash工具1配置
	Flash4 struct {
		UseTargetDll               string `xml:"UseTargetDll" json:"UseTargetDll"`                             // 是否使用目标DLL，"1"表示是
		UseExternalTool            string `xml:"UseExternalTool" json:"UseExternalTool"`                       // 是否使用外部工具，"0"表示否
		RunIndependent             string `xml:"RunIndependent" json:"RunIndependent"`                         // 是否独立运行，"0"表示否
		UpdateFlashBeforeDebugging string `xml:"UpdateFlashBeforeDebugging" json:"UpdateFlashBeforeDebugging"` // 调试前更新Flash，"1"表示是
		Capability                 string `xml:"Capability" json:"Capability"`                                 // 能力设置
		DriverSelection            string `xml:"DriverSelection" json:"DriverSelection"`                       // 驱动选择
	} `xml:"Flash4"` // Flash工具1配置
	PFcarmOut  string `xml:"pFcarmOut" json:"PFcarmOut"`
	PFcarmGrp  string `xml:"pFcarmGrp" json:"PFcarmGrp"`
	PFcArmRoot string `xml:"pFcArmRoot" json:"PFcArmRoot"`
	FcArmLst   string `xml:"FcArmLst" json:"FcArmLst"`
}

// TargetArmAds 包含ARM特定的编译和链接设置
type TargetArmAds struct {
	ArmAdsMisc ArmAdsMisc `xml:"ArmAdsMisc" json:"ArmAdsMisc"` // ARM杂项设置
	Cads       Cads       `xml:"Cads" json:"Cads"`             // C编译器设置
	Aads       Aads       `xml:"Aads" json:"Aads"`             // 汇编器设置
	LDads      LDads      `xml:"LDads" json:"LDads"`           // 链接器设置
}

// ArmAdsMisc 包含ARM杂项设置
type ArmAdsMisc struct {
	GenerateListings string         `xml:"GenerateListings" json:"GenerateListings"`
	AsHll            string         `xml:"asHll" json:"AsHll"`
	AsAsm            string         `xml:"asAsm" json:"AsAsm"`
	AsMacX           string         `xml:"asMacX" json:"AsMacX"`
	AsSyms           string         `xml:"asSyms" json:"AsSyms"`
	AsFals           string         `xml:"asFals" json:"AsFals"`
	AsDbgD           string         `xml:"asDbgD" json:"AsDbgD"`
	AsForm           string         `xml:"asForm" json:"AsForm"`
	LdLst            string         `xml:"ldLst" json:"LdLst"`
	Ldmm             string         `xml:"ldmm" json:"Ldmm"`
	LdXref           string         `xml:"ldXref" json:"LdXref"`
	BigEnd           string         `xml:"BigEnd" json:"BigEnd"`
	AdsALst          string         `xml:"AdsALst" json:"AdsALst"`
	AdsACrf          string         `xml:"AdsACrf" json:"AdsACrf"`
	AdsANop          string         `xml:"AdsANop" json:"AdsANop"`
	AdsANot          string         `xml:"AdsANot" json:"AdsANot"`
	AdsLLst          string         `xml:"AdsLLst" json:"AdsLLst"`
	AdsLmap          string         `xml:"AdsLmap" json:"AdsLmap"`
	AdsLcgr          string         `xml:"AdsLcgr" json:"AdsLcgr"`
	AdsLsym          string         `xml:"AdsLsym" json:"AdsLsym"`
	AdsLszi          string         `xml:"AdsLszi" json:"AdsLszi"`
	AdsLtoi          string         `xml:"AdsLtoi" json:"AdsLtoi"`
	AdsLsun          string         `xml:"AdsLsun" json:"AdsLsun"`
	AdsLven          string         `xml:"AdsLven" json:"AdsLven"`
	AdsLsxf          string         `xml:"AdsLsxf" json:"AdsLsxf"`
	RvctClst         string         `xml:"RvctClst" json:"RvctClst"`
	GenPPlst         string         `xml:"GenPPlst" json:"GenPPlst"`
	AdsCpuType       string         `xml:"AdsCpuType" json:"AdsCpuType"`
	RvctDeviceName   string         `xml:"RvctDeviceName" json:"RvctDeviceName"`
	MOS              string         `xml:"mOS" json:"mOS"`
	UocRom           string         `xml:"uocRom" json:"uocRom"`
	UocRam           string         `xml:"uocRam" json:"uocRam"`
	HadIROM          string         `xml:"hadIROM" json:"hadIROM"`
	HadIRAM          string         `xml:"hadIRAM" json:"hadIRAM"`
	HadXRAM          string         `xml:"hadXRAM" json:"hadXRAM"`
	UocXRam          string         `xml:"uocXRam" json:"uocXRam"`
	RvdsVP           string         `xml:"RvdsVP" json:"RvdsVP"`
	RvdsMve          string         `xml:"RvdsMve" json:"RvdsMve"`
	RvdsCdeCp        string         `xml:"RvdsCdeCp" json:"RvdsCdeCp"`
	NBranchProt      string         `xml:"nBranchProt" json:"nBranchProt"`
	HadIRAM2         string         `xml:"hadIRAM2" json:"hadIRAM2"`
	HadIROM2         string         `xml:"hadIROM2" json:"hadIROM2"`
	StupSel          string         `xml:"StupSel" json:"StupSel"`
	UseUlib          string         `xml:"useUlib" json:"useUlib"`
	EndSel           string         `xml:"EndSel" json:"EndSel"`
	ULtcg            string         `xml:"uLtcg" json:"uLtcg"`
	NSecure          string         `xml:"nSecure" json:"nSecure"`
	RoSelD           string         `xml:"RoSelD" json:"RoSelD"`
	RwSelD           string         `xml:"RwSelD" json:"RwSelD"`
	CodeSel          string         `xml:"CodeSel" json:"CodeSel"`
	OptFeed          string         `xml:"OptFeed" json:"OptFeed"`
	NoZi1            string         `xml:"NoZi1" json:"NoZi1"`
	NoZi2            string         `xml:"NoZi2" json:"NoZi2"`
	NoZi3            string         `xml:"NoZi3" json:"NoZi3"`
	NoZi4            string         `xml:"NoZi4" json:"NoZi4"`
	NoZi5            string         `xml:"NoZi5" json:"NoZi5"`
	Ro1Chk           string         `xml:"Ro1Chk" json:"Ro1Chk"`
	Ro2Chk           string         `xml:"Ro2Chk" json:"Ro2Chk"`
	Ro3Chk           string         `xml:"Ro3Chk" json:"Ro3Chk"`
	Ir1Chk           string         `xml:"Ir1Chk" json:"Ir1Chk"`
	Ir2Chk           string         `xml:"Ir2Chk" json:"Ir2Chk"`
	Ra1Chk           string         `xml:"Ra1Chk" json:"Ra1Chk"`
	Ra2Chk           string         `xml:"Ra2Chk" json:"Ra2Chk"`
	Ra3Chk           string         `xml:"Ra3Chk" json:"Ra3Chk"`
	Im1Chk           string         `xml:"Im1Chk" json:"Im1Chk"`
	Im2Chk           string         `xml:"Im2Chk" json:"Im2Chk"`
	OnChipMemories   OnChipMemories `xml:"OnChipMemories" json:"OnChipMemories"` // 片上存储器配置
	RvctStartVector  string         `xml:"RvctStartVector" json:"RvctStartVector"`
}

type OnChipMemories struct {
	Ocm1       Ocm1       `xml:"Ocm1" json:"Ocm1"`
	Ocm2       Ocm2       `xml:"Ocm2" json:"Ocm2"`
	Ocm3       Ocm3       `xml:"Ocm3" json:"Ocm3"`
	Ocm4       Ocm4       `xml:"Ocm4" json:"Ocm4"`
	Ocm5       Ocm5       `xml:"Ocm5" json:"Ocm5"`
	Ocm6       Ocm6       `xml:"Ocm6" json:"Ocm6"`
	IRAM       IRAM       `xml:"IRAM" json:"IRAM"` // 内部RAM配置
	IROM       IROM       `xml:"IROM" json:"IROM"` // 内部ROM配置
	XRAM       XRAM       `xml:"XRAM" json:"XRAM"`
	OCR_RVCT1  OCR_RVCT1  `xml:"OCR_RVCT1" json:"OCR_RVCT1"`
	OCR_RVCT2  OCR_RVCT2  `xml:"OCR_RVCT2" json:"OCR_RVCT2"`
	OCR_RVCT3  OCR_RVCT3  `xml:"OCR_RVCT3" json:"OCR_RVCT3"`
	OCR_RVCT4  OCR_RVCT4  `xml:"OCR_RVCT4" json:"OCR_RVCT4"`   // RVCT4存储器配置
	OCR_RVCT5  OCR_RVCT5  `xml:"OCR_RVCT5" json:"OCR_RVCT5"`   // RVCT4存储器配置
	OCR_RVCT6  OCR_RVCT6  `xml:"OCR_RVCT6" json:"OCR_RVCT6"`   // RVCT4存储器配置
	OCR_RVCT7  OCR_RVCT7  `xml:"OCR_RVCT7" json:"OCR_RVCT7"`   // RVCT4存储器配置
	OCR_RVCT8  OCR_RVCT8  `xml:"OCR_RVCT8" json:"OCR_RVCT8"`   // RVCT4存储器配置
	OCR_RVCT9  OCR_RVCT9  `xml:"OCR_RVCT9" json:"OCR_RVCT9"`   // RVCT4存储器配置
	OCR_RVCT10 OCR_RVCT10 `xml:"OCR_RVCT10" json:"OCR_RVCT10"` // RVCT4存储器配置
}

type Ocm1 struct {
	Type         string `xml:"Type" json:"Type"`                 // RAM类型，"0"表示标准RAM
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址，如"0x20000000"
	Size         string `xml:"Size" json:"Size"`                 // 大小，如"0x1000"(4KB)
}

type Ocm2 struct {
	Type         string `xml:"Type" json:"Type"`                 // RAM类型，"0"表示标准RAM
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址，如"0x20000000"
	Size         string `xml:"Size" json:"Size"`                 // 大小，如"0x1000"(4KB)
}

type Ocm3 struct {
	Type         string `xml:"Type" json:"Type"`                 // RAM类型，"0"表示标准RAM
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址，如"0x20000000"
	Size         string `xml:"Size" json:"Size"`                 // 大小，如"0x1000"(4KB)
}

type Ocm4 struct {
	Type         string `xml:"Type" json:"Type"`                 // RAM类型，"0"表示标准RAM
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址，如"0x20000000"
	Size         string `xml:"Size" json:"Size"`                 // 大小，如"0x1000"(4KB)
}

type Ocm5 struct {
	Type         string `xml:"Type" json:"Type"`                 // RAM类型，"0"表示标准RAM
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址，如"0x20000000"
	Size         string `xml:"Size" json:"Size"`                 // 大小，如"0x1000"(4KB)
}

type Ocm6 struct {
	Type         string `xml:"Type" json:"Type"`                 // RAM类型，"0"表示标准RAM
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址，如"0x20000000"
	Size         string `xml:"Size" json:"Size"`                 // 大小，如"0x1000"(4KB)
}

type IRAM struct {
	Type         string `xml:"Type" json:"Type"`                 // RAM类型，"0"表示标准RAM
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址，如"0x20000000"
	Size         string `xml:"Size" json:"Size"`                 // 大小，如"0x1000"(4KB)
}

type IROM struct {
	Type         string `xml:"Type" json:"Type"`                 // ROM类型，"1"表示Flash
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址，如"0x8000000"
	Size         string `xml:"Size" json:"Size"`                 // 大小，如"0x8000"(32KB)
}

type XRAM struct {
	Type         string `xml:"Type" json:"Type"`                 // RAM类型，"0"表示标准RAM
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址，如"0x20000000"
	Size         string `xml:"Size" json:"Size"`                 // 大小，如"0x1000"(4KB)
}

type OCR_RVCT1 struct {
	Type         string `xml:"Type" json:"Type"`                 // 存储器类型
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址
	Size         string `xml:"Size" json:"Size"`                 // 大小
}

type OCR_RVCT2 struct {
	Type         string `xml:"Type" json:"Type"`                 // 存储器类型
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址
	Size         string `xml:"Size" json:"Size"`                 // 大小
}

type OCR_RVCT3 struct {
	Type         string `xml:"Type" json:"Type"`                 // 存储器类型
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址
	Size         string `xml:"Size" json:"Size"`                 // 大小
}

type OCR_RVCT4 struct {
	Type         string `xml:"Type" json:"Type"`                 // 存储器类型
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址
	Size         string `xml:"Size" json:"Size"`                 // 大小
}

type OCR_RVCT5 struct {
	Type         string `xml:"Type" json:"Type"`                 // 存储器类型
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址
	Size         string `xml:"Size" json:"Size"`                 // 大小
}

type OCR_RVCT6 struct {
	Type         string `xml:"Type" json:"Type"`                 // 存储器类型
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址
	Size         string `xml:"Size" json:"Size"`                 // 大小
}

type OCR_RVCT7 struct {
	Type         string `xml:"Type" json:"Type"`                 // 存储器类型
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址
	Size         string `xml:"Size" json:"Size"`                 // 大小
}

type OCR_RVCT8 struct {
	Type         string `xml:"Type" json:"Type"`                 // 存储器类型
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址
	Size         string `xml:"Size" json:"Size"`                 // 大小
}

type OCR_RVCT9 struct {
	Type         string `xml:"Type" json:"Type"`                 // 存储器类型
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址
	Size         string `xml:"Size" json:"Size"`                 // 大小
}

type OCR_RVCT10 struct {
	Type         string `xml:"Type" json:"Type"`                 // 存储器类型
	StartAddress string `xml:"StartAddress" json:"StartAddress"` // 起始地址
	Size         string `xml:"Size" json:"Size"`                 // 大小
}

// Cads 包含C编译器设置
type Cads struct {
	Interw          string          `xml:"interw" json:"interw"`
	Optim           string          `xml:"Optim" json:"Optim"` // 优化级别，"1"表示优化级别1
	OTime           string          `xml:"oTime" json:"oTime"` // 时间优化，"0"表示不启用
	SplitLS         string          `xml:"SplitLS" json:"SplitLS"`
	OneElfS         string          `xml:"OneElfS" json:"OneElfS"`
	Strict          string          `xml:"Strict" json:"Strict"`
	EnumInt         string          `xml:"EnumInt" json:"EnumInt"`
	PlainCh         string          `xml:"PlainCh" json:"PlainCh"`
	Ropi            string          `xml:"Ropi" json:"Ropi"`
	Rwpi            string          `xml:"Rwpi" json:"Rwpi"`
	WLevel          string          `xml:"wLevel" json:"wLevel"`
	UThumb          string          `xml:"uThumb" json:"uThumb"`
	USurpInc        string          `xml:"uSurpInc" json:"uSurpInc"`
	UC99            string          `xml:"uC99" json:"uC99"` // C99标准支持，"1"表示启用
	UGnu            string          `xml:"uGnu" json:"uGnu"` // GNU扩展支持，"1"表示启用
	UseXO           string          `xml:"useXO" json:"useXO"`
	V6Lang          string          `xml:"v6Lang" json:"v6Lang"`
	V6LangP         string          `xml:"v6LangP" json:"v6LangP"`
	VShortEn        string          `xml:"vShortEn" json:"vShortEn"`
	VShortWch       string          `xml:"vShortWch" json:"vShortWch"`
	V6Lto           string          `xml:"v6Lto" json:"v6Lto"`
	V6WtE           string          `xml:"v6WtE" json:"v6WtE"`
	V6Rtti          string          `xml:"v6Rtti" json:"v6Rtti"`
	VariousControls VariousControls `xml:"VariousControls" json:"VariousControls"` // 各种控制选项
}

type VariousControls struct {
	MiscControls string `xml:"MiscControls" json:"MiscControls"` // 杂项控制选项
	Define       string `xml:"Define" json:"Define"`             // 预定义宏，如"AIR001_DEV"
	Undefine     string `xml:"Undefine" json:"Undefine"`         // 取消定义的宏
	IncludePath  string `xml:"IncludePath" json:"IncludePath"`   // 包含路径列表
}

// Aads 包含汇编器设置
type Aads struct {
	Interw          string          `xml:"interw" json:"interw"` // 交互式警告，"1"表示启用
	Ropi            string          `xml:"Ropi" json:"Ropi"`
	Rwpi            string          `xml:"Rwpi" json:"Rwpi"`
	Thumb           string          `xml:"thumb" json:"thumb"`
	SplitLS         string          `xml:"SplitLS" json:"SplitLS"`
	SwStkChk        string          `xml:"SwStkChk" json:"SwStkChk"`
	NoWarn          string          `xml:"NoWarn" json:"NoWarn"`
	USurpInc        string          `xml:"uSurpInc" json:"uSurpInc"`
	UseXO           string          `xml:"useXO" json:"useXO"`
	ClangAsOpt      string          `xml:"ClangAsOpt" json:"ClangAsOpt"`
	VariousControls VariousControls `xml:"VariousControls" json:"VariousControls"` // 各种控制选项
}

// LDads 包含链接器设置
type LDads struct {
	UmfTarg          string `xml:"umfTarg" json:"umfTarg"` // 使用微控制器格式，"1"表示启用
	Ropi             string `xml:"Ropi" json:"Ropi"`
	Rwpi             string `xml:"Rwpi" json:"Rwpi"`
	NoStLib          string `xml:"noStLib" json:"noStLib"`
	RepFail          string `xml:"RepFail" json:"RepFail"`
	UseFile          string `xml:"useFile" json:"useFile"`
	TextAddressRange string `xml:"TextAddressRange" json:"TextAddressRange"` // 代码地址范围，如"0x08000000"
	DataAddressRange string `xml:"DataAddressRange" json:"DataAddressRange"` // 数据地址范围，如"0x20000000"
	PXoBase          string `xml:"pXoBase" json:"pXoBase"`
	ScatterFile      string `xml:"ScatterFile" json:"ScatterFile"`
	IncludeLibs      string `xml:"IncludeLibs" json:"IncludeLibs"`
	IncludeLibsPath  string `xml:"IncludeLibsPath" json:"IncludeLibsPath"`
	Misc             string `xml:"Misc" json:"Misc"`
	LinkerInputFile  string `xml:"LinkerInputFile" json:"LinkerInputFile"`
	DisabledWarnings string `xml:"DisabledWarnings" json:"DisabledWarnings"`
}

// Group 表示项目中的文件组
type Group struct {
	GroupName string `xml:"GroupName" json:"GroupName"` // 组名称，如"Source Group 1"或"HAL"
	Files     []File `xml:"Files>File" json:"Files"`    // 组内文件列表
}

// File 表示项目中的单个文件
type File struct {
	FileName string `xml:"FileName" json:"FileName"` // 文件名，如"main.c"
	FileType string `xml:"FileType" json:"FileType"` // 文件类型，"1"表示C源文件，"2"表示汇编文件等
	FilePath string `xml:"FilePath" json:"FilePath"` // 文件路径，如".\main.c"
}

type RTE struct {
	APIs       APIs       `xml:"apis" json:"apis"`
	Components Components `xml:"components" json:"components"`
	Files      FileRET    `xml:"files" json:"files"`
}

type APIs struct{}

type Components struct {
	Components []Component `xml:"component" json:"component"`
}

type Component struct {
	CClass     string     `xml:"Cclass,attr" json:"attr_Cclass"`
	CGroup     string     `xml:"Cgroup,attr" json:"attr_Cgroup"`
	CVendor    string     `xml:"Cvendor,attr" json:"attr_Cvendor"`
	CVersion   string     `xml:"Cversion,attr" json:"attr_Cversion"`
	Condition  string     `xml:"condition,attr" json:"attr_condition"`
	Package    Package    `xml:"package" json:"package"`
	TargetInfo TargetInfo `xml:"targetInfos" json:"targetInfos"`
}

type Package struct {
	Name          string `xml:"name,attr" json:"attr_name"`
	SchemaVersion string `xml:"schemaVersion,attr" json:"attr_schemaVersion"`
	URL           string `xml:"url,attr" json:"attr_url"`
	Vendor        string `xml:"vendor,attr" json:"attr_vendor"`
	Version       string `xml:"version,attr" json:"attr_version"`
}

type TargetInfo struct {
	TargetName string `xml:"targetInfo" json:"targetInfo"`
}

type FileRET struct {
	Files []FilesRET `xml:"file" json:"file"`
}

type FilesRET struct {
	Attr      string     `xml:"attr,attr" json:"attr_attr"`
	Category  string     `xml:"category,attr" json:"attr_category"`
	Condition string     `xml:"condition,attr,omitempty" json:"attr_condition,omitempty"`
	Name      string     `xml:"name,attr" json:"attr_name"`
	Version   string     `xml:"version,attr" json:"attr_version"`
	Instance  Instance   `xml:"instance" json:"instance"`
	Component Component  `xml:"component" json:"component"`
	Package   Package    `xml:"package" json:"package"`
	Target    TargetInfo `xml:"targetInfos" json:"targetInfos"`
}

type Instance struct {
	Index int    `xml:"index,attr" json:"attr_index"`
	Path  string `xml:",chardata" json:"path"`
}

type LayerInfo struct {
	Layers Layers `xml:"Layers" json:"Layers"`
}

type Layers struct {
	Layer Layer `xml:"Layer" json:"Layer"`
}

type Layer struct {
	LayName    string `xml:"LayName" json:"LayName"`
	LayPrjMark string `xml:"LayPrjMark" json:"LayPrjMark"`
}

// writeFile 将字符串写入文件
func writeFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// LoadNewKeilConfigFile 将keil5工程文件存储为json模版文件
func LoadNewKeilConfigFile(projectFile string, inputPath string, chipType string) ([]byte, error) {
	data, err := os.ReadFile(inputPath + projectFile)
	if err != nil {
		return nil, err
	}

	// 解析XML到结构体
	var project KeilProject
	err = xml.Unmarshal(data, &project)
	if err != nil {
		return nil, err
	}

	project.XMLNSXSI = "http://www.w3.org/2001/XMLSchema-instance"
	project.XSINoNamespaceSchemaLocation = "project_projx.xsd"

	jsonData, err := json.Marshal(project)
	if err != nil {
		return nil, err
	}
	err = writeFile("./template/raw/"+chipType+"/"+chipType+"-keil.json", jsonData, 0666)
	return jsonData, nil
}

// CreateNewKeilProjectFile 从json模版生成keil5工程
func CreateNewKeilProjectFile(projectFile string, outputPath string, chipType string) ([]byte, error) {
	data, err := os.ReadFile("./template/raw/" + chipType + "/" + chipType + "-keil.json")
	output := &KeilProject{}
	err = json.Unmarshal(data, output)
	if err != nil {
		return nil, err
	}

	data, err = xml.MarshalIndent(output, "  ", "    ")
	if err != nil {
		return nil, err
	}
	outputFile := "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\" ?> \n" + string(data)
	err = writeFile(outputPath+projectFile, []byte(outputFile), 0666)
	if err != nil {
		return nil, err
	}
	return []byte(outputFile), nil
}
