package handler

type ThriftHandler struct {

}

func(p *ThriftHandler) Add(addStart int32, addEnd int32) (int32, error) {
	return addStart+addEnd, nil
}
