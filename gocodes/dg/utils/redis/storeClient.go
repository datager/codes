package redis

const (
	storePrefix = "loki-store"
)

type StoreClient struct {
	ConnPool
}

func NewStoreClient(baseAddr, password string) *StoreClient {
	if baseAddr == "" {
		panic("store server addr is empty")
	}

	conn := initConnPool(baseAddr, password, storePrefix)

	return &StoreClient{
		ConnPool: *conn,
	}
}
