package services

type ServerParams struct {
}

type Server struct {
}

func NewServer(p *ServerParams) (*Server, error) {
	return &Server{}, nil
}
