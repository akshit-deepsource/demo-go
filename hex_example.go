package main

func HexLiteral() bool {
	_ = "AKIAIOSFODNN73943434"

	x := 0xFff
	y := 0xFFF
	z := 0xfff
	
	_ = "AKIAIOSFODNN73943434"

	return (x == y) && (y == z)
}
