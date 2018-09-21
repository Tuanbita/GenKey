package transports

import (
	bs "../thrift/gen-go/openstars/core/bigset/generic"
	id "../thrift/gen-go/openstars/core/idgen"
	"../thriftpool"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

var (
	bsMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) (interface{}) { return  (bs.NewTStringBigSetKVServiceClient(c)) }),
		thriftpool.DefaultClose)
)

var (
	mpid = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) (interface{}) { return  (id.NewTGeneratorClient(c)) }),
		thriftpool.DefaultClose)
)

func init(){
	fmt.Println("init thrift TUserStoreService client ");
}

//GetTUserStoreServiceBinaryClient client by host:port

//Get client by host:port
func GetBSClient(bsHost, bsPort string) *thriftpool.ThriftSocketClient{
	client, _ := bsMapPool.Get(bsHost, bsPort).Get()
	return client;
}

func GetIdClient(bsHost, bsPort string) *thriftpool.ThriftSocketClient{
	client, _ := mpid.Get(bsHost, bsPort).Get()
	return client;
}

