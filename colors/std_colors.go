package colors

var (
	Black          Color
	Blue           Color
	ElectricGreen  Color
	ElectricCyan   Color
	Red            Color
	Magenta        Color
	White          Color
	ElectricYellow Color
)

func init() {
	Black = New(0, 0, 0)
	White = New(255, 255, 255)

	Blue = New(0, 0, 255)

	ElectricGreen = New(0, 255, 0)
	ElectricCyan = New(0, 255, 255)

	Red = New(255, 0, 0)

	Magenta = New(255, 0, 255)
	ElectricYellow = New(255, 255, 0)
}
