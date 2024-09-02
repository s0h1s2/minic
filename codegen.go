package main

func GenAst(expr Expr) int {
	var leftReg, rightReg int
	switch e := expr.(type) {
	case *BinaryExpr:
		{
			switch e.Kind {
			case Plus:
				return cgAdd(leftReg, rightReg)
			case Minus:
				return cgSubtract(leftReg, rightReg)
			case Star:
				return cgSubtract(leftReg, rightReg)
			case Slash:
				return cgSubtract(leftReg, rightReg)

			}
		}
	case *IntegerExpr:
		{
			return cgLoad(e.Value)
		}
	}
	panic("Invalid ast in codegen")
}

func cgLoad(value string) int {
	panic("unimplemented")
}

func cgSubtract(leftReg, rightReg int) int {
	panic("unimplemented")
}

func cgAdd(leftReg, rightReg int) int {
	panic("unimplemented")
}

func allocateRegister() int {}
