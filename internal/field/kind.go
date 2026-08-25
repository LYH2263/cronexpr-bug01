package field

// Kind identifies a cron column.
type Kind int

const (
	Second Kind = iota
	Minute
	Hour
	DayOfMonth
	Month
	DayOfWeek
)

func (k Kind) String() string {
	switch k {
	case Second:
		return "second"
	case Minute:
		return "minute"
	case Hour:
		return "hour"
	case DayOfMonth:
		return "dom"
	case Month:
		return "month"
	case DayOfWeek:
		return "dow"
	default:
		return "unknown"
	}
}

// StandardKinds returns field order for 5 or 6 part cron.
func StandardKinds(withSecond bool) []Kind {
	if withSecond {
		return []Kind{Second, Minute, Hour, DayOfMonth, Month, DayOfWeek}
	}
	return []Kind{Minute, Hour, DayOfMonth, Month, DayOfWeek}
}
