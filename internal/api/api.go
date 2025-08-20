package api

type APIServer struct {
	listenAddr string
	storage    Storage
}

func (s *APIServer) Run() {}

func NewAPIServer(listenAddr string, storage Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		storage:    storage,
	}
}
