package markdown

import (
	"github.com/russross/blackfriday"
	"github.com/spiral/roadrunner/service/rpc"
)

const ID  = "markdown"

type Service struct {}

func (s *Service) Init(rpc *rpc.Service) (bool, error) {
	if err := rpc.Register(ID, &rpcService{}); err != nil {
		return false, err
	}

	return true, nil
}

type rpcService struct {}

func (s *rpcService) Convert(input []byte, output *[]byte) error {
	*output = blackfriday.MarkdownBasic(input)
	return nil
}
