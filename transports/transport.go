package transports

import (
	bs "GrpcChatServer/ThriftCommon/thrift/gen-go/openstars/core/bigset/generic"
	id "GrpcChatServer/ThriftCommon/thrift/gen-go/openstars/core/idgen"
	"GrpcChatServer/ThriftCommon/thriftpool"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/astaxie/beego/logs"
	"GrpcChatServer/ThriftCommon/thrift/gen-go/OpenStars/Platform/UserStore"
)

var (
	mTUserStoreServiceBinaryMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc( func( c thrift.TClient ) (interface{}) { return  (UserStore.NewTUserStoreServiceClient(c)) }),
		thriftpool.DefaultClose)

	mTUserStoreServiceCommpactMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol( func(c thrift.TClient) (interface{}) { return  (UserStore.NewTUserStoreServiceClient(c)) }),
		thriftpool.DefaultClose)
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
	logs.Info("init thrift TUserStoreService client ");
}

//GetTUserStoreServiceBinaryClient client by host:port
func GetTUserStoreServiceBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient{
	client, _ := mTUserStoreServiceBinaryMapPool.Get(aHost, aPort).Get()
	return client;
}

//GetTUserStoreServiceCompactClient get compact client by host:port
func GetTUserStoreServiceCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient{
	client, _ := mTUserStoreServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client;
}

//Get client by host:port
func GetBSClient(bsHost, bsPort string) *thriftpool.ThriftSocketClient{
	client, _ := bsMapPool.Get(bsHost, bsPort).Get()
	return client;
}

func GetIdClient(bsHost, bsPort string) *thriftpool.ThriftSocketClient{
	client, _ := mpid.Get(bsHost, bsPort).Get()
	return client;
}

