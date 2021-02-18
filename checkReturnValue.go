package minioExtensions

import "errors"

func thenBranchGoodError() error {

	if errorSource() != ErrNone {
		// Good: returning an error
		return errors.New("failed")
	}
	doSomething()
	return nil	

}

func thenBranchGoodWithElseError() error {

	if errorSource() != ErrNone {
		// Good: an error means we return an error
		return errors.New("failed")
	} else {
		doSomething()
	}
	doSomething()
	return nil

}

func thenBranchBadError() error {

	if errorSource() != ErrNone {
		// Bad: despite an error, we return nil
		return nil
	}
	doSomething()
	return nil

}

func thenBranchBadWithElseError() error {

	if errorSource() != ErrNone {
		// Bad: despite an error, we return nil
		return nil
	} else {
		doSomething()
	}
	doSomething()
	return nil

}

func elseBranchGoodError() error {

	if errorSource() == ErrNone {
		doSomething()
	} else {
		// Good: returning an error
		return errors.New("failed")
	}
	doSomething()
	return nil
}

func elseBranchBadError() error {

	if errorSource() == ErrNone {
		doSomething()
	} else {
		// Bad: despite an error, we return nil
		return nil
	}
	doSomething()
	return nil

}

func multiReturnBad() (string, error) {

	if errorSource() != ErrNone {
		// Bad: despite an error, we return a nil error
		return "", nil
	}
	doSomething()
	return "Result", nil

}

func getNil() error {
	return nil
}

func getError(s string) error {
	return errors.New(s)
}

func thenBranchGoodInterproceduralError() error {

	if errorSource() != ErrNone {
		// Good: returning an error
		return getError("failed")
	}
	doSomething()
	return getNil()
	
}

func thenBranchBadInterproceduralError() error {

	if errorSource() != ErrNone {
		// Good: returning nil despite an error
		return getNil()
	}
	doSomething()
	return getNil()
	
}

