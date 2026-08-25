package cronexpr

import "errors"

var (
	ErrInvalid   = errors.New("cronexpr: invalid expression")
	ErrEmpty     = errors.New("cronexpr: empty expression")
	ErrConflict  = errors.New("cronexpr: day-of-month and day-of-week conflict")
	ErrNoMatch   = errors.New("cronexpr: no matching fire time")
	ErrCanceled  = errors.New("cronexpr: canceled")
	ErrNilExpr   = errors.New("cronexpr: nil expression")
	ErrBadLocale = errors.New("cronexpr: unknown locale")
)
