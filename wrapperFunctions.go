package minioExtensions

func success() bool {
	return errorSource() == ErrNone
}

func failure() bool {
	return errorSource() != ErrNone
}

func succeeded(ret int) bool {
	return ret == ErrNone
}

func failed(ret int) bool {
	return ret != ErrNone 
}

func thenBranchGoodSourceWrapper() {

	if failure() {
		// Good: an error means we return early
		return
	}
	doSomething()

}

func thenBranchBadSourceWrapper() {

	if failure() {
		// Bad: despite an error, we carry on to execute doSomething()
		insteadOfReturn()
	}
	doSomething()

}

func elseBranchGoodSourceWrapper() {

	if success() {
		doSomething()
	} else {
		// Good: an error means we return early
		return
	}
	doSomething()

}

func elseBranchBadSourceWrapper() {

	if success() {
		doSomething()
	} else {
		// Bad: despite an error, we carry on to execute doSomething()
		insteadOfReturn()
	}
	doSomething()

}

func thenBranchGoodTestWrapper() {

	if failed(errorSource()) {
		// Good: an error means we return early
		return
	}
	doSomething()

}

func thenBranchBadTestWrapper() {

	if failed(errorSource()) {
		// Bad: despite an error, we carry on to execute doSomething()
		insteadOfReturn()
	}
	doSomething()

}

func elseBranchGoodTestWrapper() {

	if succeeded(errorSource()) {
		doSomething()
	} else {
		// Good: an error means we return early
		return
	}
	doSomething()

}

func elseBranchBadTestWrapper() {

	if succeeded(errorSource()) {
		doSomething()
	} else {
		// Bad: despite an error, we carry on to execute doSomething()
		insteadOfReturn()
	}
	doSomething()

}
