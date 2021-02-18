package minioExtensions

const particularErrOne = 1
const particularErrTwo = 2

func subBranchGood() int {

	// Good: while the if-block's logic is complex, it always returns.
	err := errorSource()
	if err != ErrNone {
		if err == particularErrOne {
			return -1
		} else if err == particularErrTwo {
			return -2
		} else {
			return -3
		}
	}
	doSomething()
	return 0

}

func subBranchBad() int {

	// Bad: one of the if-block's branches falls through to execute `doSomething()`.
	err := errorSource()
	if err != ErrNone {
		if err == particularErrOne {
			return -1
		} else if err == particularErrTwo {
			err = ErrNone
		} else {
			return -3
		}
	}
	doSomething()
	return 0

}

func switchGood() int {

	// Good: while the if-block's logic is complex, it always returns.
	err := errorSource()
	if err != ErrNone {
		switch err {
		case particularErrOne:
			return -1
		case particularErrTwo:
			return -2
		default:
			return -3
		}
	}
	doSomething()
	return 0

}

func switchBad() int {

	// Bad: one of the if-block's branches falls through to execute `doSomething()`.
	err := errorSource()
	if err != ErrNone {
		switch err {
		case particularErrOne:
			err = ErrNone
		case particularErrTwo:
			return -2
		default:
			return -3
		}
	}
	doSomething()
	return 0

}