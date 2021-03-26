package minioExtensions

func negatedThenBranchGood() {

	if !(errorSource() == ErrNone) {
		// Good: an error means we return early
		return
	}
	doSomething()

}

func negatedThenBranchGoodWithElse() {

	if !(errorSource() == ErrNone) {
		// Good: an error means we return early
		return
	} else {
		doSomething()
	}
	doSomething()

}

func negatedElseBranchGood() {

	if !(errorSource() != ErrNone) {
		doSomething()
	} else {
		// Good: an error means we return early
		return
	}
	doSomething()

}

func negatedElseBranchBad() {

	if !(errorSource() != ErrNone) {
		doSomething()
	} else {
		// Bad: despite an error, we carry on to execute doSomething()
		insteadOfReturn()
	}
	doSomething()

}

func someOtherCondition() bool {
	return true
}

func logicalAndThenBranchSometimesBad() {

	if errorSource() != ErrNone && someOtherCondition() {
		// Bad: there is a route from a positive error test around the 'return' statement.
		return
	}
	doSomething()

}

func logicalAndThenBranchAlwaysBad() {

	if errorSource() != ErrNone && someOtherCondition() {
		// Bad: there is no return statement at all.
		insteadOfReturn()
	}
	doSomething()

}

func logicalAndElseBranchAlwaysBad2() {

	if errorSource() == ErrNone && someOtherCondition() {
		doSomething()
	} else {
		// Bad: there is no return statement at all.
		insteadOfReturn()
	}
	doSomething()

}

func logicalAndThenBranchGood() {

	if someOtherCondition() && errorSource() != ErrNone {
		// Good: whenever an error is indicated we return (note errorSource() is not called until someOtherCondition() passes)
		return
	}
	doSomething()

}

func logicalAndElseBranchGood() {

	if someOtherCondition() && errorSource() == ErrNone {
		// Good: whenever an error is indicated we return (note errorSource() is not called until someOtherCondition() passes)
		doSomething()
	} else {
		return
	}

}

func logicalAndElseBranchGood2() {

	if errorSource() == ErrNone && someOtherCondition() {
		// Good: whenever an error is indicated we return.
		doSomething()
	} else {
		return
	}
	doSomething()

}

func logicalOrElseBranchSometimesBad() {

	if errorSource() == ErrNone || someOtherCondition() {
		doSomething()
	} else {
		// Bad: there is a route from a failing error test that bypasses the return statement.
		return
	}
	doSomething()

}

func logicalOrElseBranchAlwaysBad() {

	if errorSource() == ErrNone || someOtherCondition() {
		doSomething()
	} else {
		// Bad: regardless of error status, we do not return.
		insteadOfReturn()
	}
	doSomething()

}

func logicalOrThenBranchAlwaysBad() {

	if errorSource() != ErrNone || someOtherCondition() {
		// Bad: regardless of error status, we do not return.
		insteadOfReturn()
	} else {
		doSomething()
	}
	doSomething()

}

func logicalOrThenBranchGood() {

	if someOtherCondition() || errorSource() != ErrNone {
		// Good: whenever an error is indicated we return. (note errorSource() is not called until someOtherCondition() fails)
		return
	}
	doSomething()

}

func logicalOrElseBranchGood() {

	if someOtherCondition() || errorSource() == ErrNone {
		// Good: whenever an error is indicated we return. (note errorSource() is not called until someOtherCondition() fails)
		doSomething()
	} else {
		return
	}

}

func logicalOrThenBranchGood2() {

	if errorSource() != ErrNone || someOtherCondition() {
		// Good: whenever an error is indicated we return.
		return
	}
	doSomething()

}
