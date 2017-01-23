package add

func Add(a, b int) int {
	GoCover_0.Count[0] = 1
	return a + b
}

var GoCover_0 = struct {
	Count     [1]uint32
	Pos       [3 * 1]uint32
	NumStmt   [1]uint16
} {
	Pos: [3 * 1]uint32{
		3, 5, 0x20018, // [0]
	},
	NumStmt: [1]uint16{
		1, // 0
	},
}
