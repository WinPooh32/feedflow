package session

import "github.com/WinPooh32/feedflow/user/previlegies"

func (s *Session) SetUserID(id int64) {
	s.setNumber(FieldUserID, float64(id))
}

func (s *Session) GetUserID() int64 {
	v, ok := s.getInt64(FieldUserID)
	if !ok {
		return 0
	}
	return v
}

func (s *Session) SetUserRole(role previlegies.Role) {
	s.setNumber(FieldPageHits, float64(role))
}

func (s *Session) GetUserRole() previlegies.Role {
	v, ok := s.get(FieldUserID)
	value, valid := v.(previlegies.Role)

	if !ok {
		return previlegies.Guest
	}

	if !valid {
		logFieldError(FieldUserID)
		return previlegies.Guest
	}

	return value
}

func (s *Session) SetHits(hits int64) {
	s.setNumber(FieldPageHits, float64(hits))
}

func (s *Session) GetHits() int64 {
	v, ok := s.getInt64(FieldPageHits)
	if !ok {
		s.SetHits(0)
		return 0
	}
	return v
}
