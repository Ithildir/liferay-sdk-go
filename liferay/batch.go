package liferay

type BatchSession struct {
	cmds []map[string]interface{}
	*session
}

func NewBatchSession(server string, username string, password string) *BatchSession {
	return &BatchSession{session: NewSession(server, username, password).(*session)}
}

func (s *BatchSession) Clear() {
	s.cmds = nil
}

func (s *BatchSession) Invoke(cmd map[string]interface{}) (interface{}, error) {
	s.cmds = append(s.cmds, cmd)

	return nil, nil
}

func (s *BatchSession) InvokeAll() ([]interface{}, error) {
	defer s.Clear()

	return post(s, s.cmds)
}
