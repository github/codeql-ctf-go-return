package minioExtensions

func errorSource() int {
        // The source of error values that must be checked.
        // We should treat this the same as `isReqAuthenticated` in the real `minio`.
        return 0
}

// The ErrNone value that errorSource() is compared against.
var ErrNone = 0

func doSomething() {
        // A no-op used as a stand-in for whatever code would normally follow an error check.
}

// A filler function, used to occupy a then- or else-block that would otherwise be empty, 
// where a return would be expected
func insteadOfReturn() {
}
