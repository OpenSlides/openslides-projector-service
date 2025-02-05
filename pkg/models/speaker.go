package models

func (m *Speaker) IsCurrent() bool {
	return m.BeginTime != nil && *m.BeginTime != 0 && (m.EndTime == nil || *m.EndTime == 0)
}
