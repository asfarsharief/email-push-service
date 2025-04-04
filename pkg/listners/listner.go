package listners

type ListnerInterface interface {
	InitializeListner() error
	Listen(topic string)
}

func GetListnerObject(name string) ListnerInterface {
	switch name {
	case "nats":
		return NewNatsListner("default")
	}
	return nil
}
