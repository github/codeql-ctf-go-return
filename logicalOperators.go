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

func logicalAndThenBranchGood() {

	if errorSource() != ErrNone && someOtherCondition() {
		// Good: in the then-block, where we're certain there has been an error,
		// we return early. In the other branch an error is possible but not certain. 
		return
	}
	doSomething()

}

func logicalAndThenBranchBad() {

	if errorSource() != ErrNone && someOtherCondition() {
		// Bad: in the then-block, where we're certain there has been an error,
		// we do not return early.
		insteadOfReturn()
	}
	doSomething()

}

func logicalAndElseBranchUncertain() {

	if errorSource() == ErrNone && someOtherCondition() {
		doSomething()
	} else {
		// Uncertain: an error is not a precondition for entering either branch.
		// We should be conservative and NOT flag this function as a problem.
		insteadOfReturn()
	}
	doSomething()

}

func logicalOrElseBranchGood() {

	if errorSource() == ErrNone || someOtherCondition() {
		doSomething()
	} else {
		// Good: in the else-block, where we're certain there has been an error,
		// we return early. In the other branch an error is possible but not certain. 
		return
	}
	doSomething()

}

func logicalOrElseBranchBad() {

	if errorSource() == ErrNone || someOtherCondition() {
		doSomething()
	} else {
		// Bad: in the else-block, where we're certain there has been an error,
		// we do not return early.
		insteadOfReturn()
	}
	doSomething()

}

func logicalOrThenBranchUncertain() {

	if errorSource() != ErrNone || someOtherCondition() {
		// Uncertain: an error is not a precondition for entering either branch.
		// We should be conservative and NOT flag this function as a problem.
		insteadOfReturn()
	} else {
		doSomething()
	}
	doSomething()

}
