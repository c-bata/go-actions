package actions

const (
	// ExitSuccess means the action completed successfully.
	ExitSuccess = 0
	// ExitNeutral to terminates all concurrently running actions and prevents any future actions from starting.
	ExitNeutral = 78
)
