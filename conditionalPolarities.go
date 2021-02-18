package minioExtensions

func thenBranchGood() {

	if errorSource() != ErrNone {
		// Good: an error means we return early
		return
	}
	doSomething()

}

func thenBranchGoodWithElse() {

	if errorSource() != ErrNone {
		// Good: an error means we return early
		return
	} else {
		doSomething()
	}
	doSomething()

}

func thenBranchBad() {

	if errorSource() != ErrNone {
		// Bad: despite an error, we carry on to execute doSomething()
		insteadOfReturn()
	}
	doSomething()

}

func thenBranchBadWithElse() {

	if errorSource() != ErrNone {
		// Bad: despite an error, we carry on to execute doSomething()
		insteadOfReturn()
	} else {
		doSomething()
	}
	doSomething()

}

func elseBranchGood() {

	if errorSource() == ErrNone {
		doSomething()
	} else {
		// Good: an error means we return early
		return
	}
	doSomething()

}

func elseBranchBad() {

	if errorSource() == ErrNone {
		doSomething()
	} else {
		// Bad: despite an error, we carry on to execute doSomething()
		insteadOfReturn()
	}
	doSomething()

}
