package template

type Config struct {
	App     AppInfo           `json:"app"`
	MCU     MCUInfo           `json:"mcu"`
	Pin     map[string]string `json:"pin"`
	PinMode map[string]string `json:"pin_mode"`
	RCC     map[string]any    `json:"rcc"`
	ADC     map[string]any    `json:"adc"`
	TIM     map[string]any    `json:"tim"`
	SPI     map[string]any    `json:"spi"`
	USART   map[string]any    `json:"usart"`
	UART    map[string]any    `json:"uart"`
	Other   map[string]any    `json:"other"`
}

// AppInfo 代表应用程序信息
type AppInfo struct {
	Version string `json:"version"`
}

// MCUInfo 代表 MCU 的基本信息
type MCUInfo struct {
	Name         string `json:"name"`
	Encapsulated string `json:"encapsulated"`
	Number       string `json:"number"`
}
