package main

import (
	bs "./thrift/gen-go/openstars/core/bigset/generic"
	"./transports"
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"strconv"
)
type key struct {
	id hexutil.Bytes
	Priv  hexutil.Bytes
	Pub  hexutil.Bytes
	s hexutil.Bytes
}

func genKey() *ecdsa.PrivateKey{
	privateKey, _ := crypto.GenerateKey()
	return privateKey
}

var priv = make(map[int]string)
var pub = make(map[int]string)

func main() {
	client := transports.GetBSClient("127.0.0.1", "18407")
	defer client.BackToPool()

	for i:=1; i<=20; i++{
		fmt.Print(i,"  : ")
		privateKey :=  genKey()

		pubKey_byt := crypto.CompressPubkey(&privateKey.PublicKey)

		//pubKey_byt := crypto.FromECDSAPub(publicKey)

		fmt.Print ("priv: ",string(privateKey.D.Bytes()))
		fmt.Println("      pub: ",string(pubKey_byt))

		key := strconv.Itoa(i)
		client.Client.(*bs.TStringBigSetKVServiceClient).BsPutItem(context.Background(),"PrivateKey", &bs.TItem{[]byte(key), privateKey.D.Bytes()})
		client.Client.(*bs.TStringBigSetKVServiceClient).BsPutItem(context.Background(),"PublicKey", &bs.TItem{[]byte(key), pubKey_byt})
	}

	var Keys [21]key
	//neu nhu co thi return false
	for i := 1; i<=20; i++{

		priv,_ := client.Client.(*bs.TStringBigSetKVServiceClient).BsGetItem(context.Background(),"PrivateKey", []byte(strconv.Itoa(i)))
		pub,_ := client.Client.(*bs.TStringBigSetKVServiceClient).BsGetItem(context.Background(),"PublicKey", []byte(strconv.Itoa(i)))

		//abc,_ := client.Client.(*bs.TStringBigSetKVServiceClient).BsGetItem("abc", []byte(strconv.Itoa(i)))
		//time.Sleep(time.Second*2)
		if priv != nil && priv.Item != nil{
			fmt.Println(i)
			logs.Info("priv: ",string(priv.GetItem().GetValue()))
			logs.Info("pub: ",hex.EncodeToString(pub.GetItem().GetValue()))

			a := key{
					[]byte(strconv.Itoa(i)),
				priv.GetItem().GetValue(),
				pub.GetItem().GetValue(),
					[]byte("\n"),
			}
			Keys[i]=a
		}

	}
	//
	//
	rankingsJson, _ := json.Marshal(Keys)
	ioutil.WriteFile("output.json", rankingsJson, 0644)
	fmt.Printf("%+v", Keys)

}
