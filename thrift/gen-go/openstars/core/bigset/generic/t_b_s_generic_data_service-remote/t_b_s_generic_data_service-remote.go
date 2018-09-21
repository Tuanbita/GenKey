// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "context"
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "openstars/core/bigset/generic"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  TPutItemResult bsgPutItem(TContainerKey rootID, TItem item)")
  fmt.Fprintln(os.Stderr, "  bool bsgRemoveItem(TMetaKey key, TItemKey itemKey)")
  fmt.Fprintln(os.Stderr, "  TExistedResult bsgExisted(TContainerKey rootID, TItemKey itemKey)")
  fmt.Fprintln(os.Stderr, "  TItemResult bsgGetItem(TContainerKey rootID, TItemKey itemKey)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult bsgGetSlice(TContainerKey rootID, i32 fromIdx, i32 count)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult bsgGetSliceFromItem(TContainerKey rootID, TItemKey fromKey, i32 count)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult bsgGetSliceR(TContainerKey rootID, i32 fromIdx, i32 count)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult bsgGetSliceFromItemR(TContainerKey rootID, TItemKey fromKey, i32 count)")
  fmt.Fprintln(os.Stderr, "  TSplitBigSetResult splitBigSet(TContainerKey rootID, TContainerKey brotherRootID, i64 currentSize)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult bsgRangeQuery(TContainerKey rootID, TItemKey startKey, TItemKey endKey)")
  fmt.Fprintln(os.Stderr, "  bool bsgBulkLoad(TContainerKey rootID, TItemSet setData)")
  fmt.Fprintln(os.Stderr, "  TMultiPutItemResult bsgMultiPut(TContainerKey rootID, TItemSet setData, bool getAddedItems, bool getReplacedItems)")
  fmt.Fprintln(os.Stderr, "  TBigSetGenericData getSetGenData(TMetaKey metaID)")
  fmt.Fprintln(os.Stderr, "  void putSetGenData(TMetaKey metaID, TBigSetGenericData metadata)")
  fmt.Fprintln(os.Stderr, "  i64 getTotalCount(TContainerKey rootID)")
  fmt.Fprintln(os.Stderr, "  i64 removeAll(TContainerKey rootID)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := generic.NewTBSGenericDataServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "bsgPutItem":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BsgPutItem requires 2 args")
      flag.Usage()
    }
    argvalue0, err50 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err50 != nil {
      Usage()
      return
    }
    value0 := generic.TContainerKey(argvalue0)
    arg51 := flag.Arg(2)
    mbTrans52 := thrift.NewTMemoryBufferLen(len(arg51))
    defer mbTrans52.Close()
    _, err53 := mbTrans52.WriteString(arg51)
    if err53 != nil {
      Usage()
      return
    }
    factory54 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt55 := factory54.GetProtocol(mbTrans52)
    argvalue1 := generic.NewTItem()
    err56 := argvalue1.Read(jsProt55)
    if err56 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.BsgPutItem(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "bsgRemoveItem":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BsgRemoveItem requires 2 args")
      flag.Usage()
    }
    argvalue0, err57 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err57 != nil {
      Usage()
      return
    }
    value0 := generic.TMetaKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    fmt.Print(client.BsgRemoveItem(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "bsgExisted":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BsgExisted requires 2 args")
      flag.Usage()
    }
    argvalue0, err59 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err59 != nil {
      Usage()
      return
    }
    value0 := generic.TContainerKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    fmt.Print(client.BsgExisted(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "bsgGetItem":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BsgGetItem requires 2 args")
      flag.Usage()
    }
    argvalue0, err61 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err61 != nil {
      Usage()
      return
    }
    value0 := generic.TContainerKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    fmt.Print(client.BsgGetItem(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "bsgGetSlice":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BsgGetSlice requires 3 args")
      flag.Usage()
    }
    argvalue0, err63 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err63 != nil {
      Usage()
      return
    }
    value0 := generic.TContainerKey(argvalue0)
    tmp1, err64 := (strconv.Atoi(flag.Arg(2)))
    if err64 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    tmp2, err65 := (strconv.Atoi(flag.Arg(3)))
    if err65 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.BsgGetSlice(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "bsgGetSliceFromItem":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BsgGetSliceFromItem requires 3 args")
      flag.Usage()
    }
    argvalue0, err66 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err66 != nil {
      Usage()
      return
    }
    value0 := generic.TContainerKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    tmp2, err68 := (strconv.Atoi(flag.Arg(3)))
    if err68 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.BsgGetSliceFromItem(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "bsgGetSliceR":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BsgGetSliceR requires 3 args")
      flag.Usage()
    }
    argvalue0, err69 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err69 != nil {
      Usage()
      return
    }
    value0 := generic.TContainerKey(argvalue0)
    tmp1, err70 := (strconv.Atoi(flag.Arg(2)))
    if err70 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    tmp2, err71 := (strconv.Atoi(flag.Arg(3)))
    if err71 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.BsgGetSliceR(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "bsgGetSliceFromItemR":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BsgGetSliceFromItemR requires 3 args")
      flag.Usage()
    }
    argvalue0, err72 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err72 != nil {
      Usage()
      return
    }
    value0 := generic.TContainerKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    tmp2, err74 := (strconv.Atoi(flag.Arg(3)))
    if err74 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.BsgGetSliceFromItemR(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "splitBigSet":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "SplitBigSet requires 3 args")
      flag.Usage()
    }
    argvalue0, err75 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err75 != nil {
      Usage()
      return
    }
    value0 := generic.TContainerKey(argvalue0)
    argvalue1, err76 := (strconv.ParseInt(flag.Arg(2), 10, 64))
    if err76 != nil {
      Usage()
      return
    }
    value1 := generic.TContainerKey(argvalue1)
    argvalue2, err77 := (strconv.ParseInt(flag.Arg(3), 10, 64))
    if err77 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    fmt.Print(client.SplitBigSet(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "bsgRangeQuery":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BsgRangeQuery requires 3 args")
      flag.Usage()
    }
    argvalue0, err78 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err78 != nil {
      Usage()
      return
    }
    value0 := generic.TContainerKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    argvalue2 := []byte(flag.Arg(3))
    value2 := generic.TItemKey(argvalue2)
    fmt.Print(client.BsgRangeQuery(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "bsgBulkLoad":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BsgBulkLoad requires 2 args")
      flag.Usage()
    }
    argvalue0, err81 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err81 != nil {
      Usage()
      return
    }
    value0 := generic.TContainerKey(argvalue0)
    arg82 := flag.Arg(2)
    mbTrans83 := thrift.NewTMemoryBufferLen(len(arg82))
    defer mbTrans83.Close()
    _, err84 := mbTrans83.WriteString(arg82)
    if err84 != nil {
      Usage()
      return
    }
    factory85 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt86 := factory85.GetProtocol(mbTrans83)
    argvalue1 := generic.NewTItemSet()
    err87 := argvalue1.Read(jsProt86)
    if err87 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.BsgBulkLoad(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "bsgMultiPut":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "BsgMultiPut requires 4 args")
      flag.Usage()
    }
    argvalue0, err88 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err88 != nil {
      Usage()
      return
    }
    value0 := generic.TContainerKey(argvalue0)
    arg89 := flag.Arg(2)
    mbTrans90 := thrift.NewTMemoryBufferLen(len(arg89))
    defer mbTrans90.Close()
    _, err91 := mbTrans90.WriteString(arg89)
    if err91 != nil {
      Usage()
      return
    }
    factory92 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt93 := factory92.GetProtocol(mbTrans90)
    argvalue1 := generic.NewTItemSet()
    err94 := argvalue1.Read(jsProt93)
    if err94 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    argvalue2 := flag.Arg(3) == "true"
    value2 := argvalue2
    argvalue3 := flag.Arg(4) == "true"
    value3 := argvalue3
    fmt.Print(client.BsgMultiPut(context.Background(), value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "getSetGenData":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetSetGenData requires 1 args")
      flag.Usage()
    }
    argvalue0, err97 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err97 != nil {
      Usage()
      return
    }
    value0 := generic.TMetaKey(argvalue0)
    fmt.Print(client.GetSetGenData(context.Background(), value0))
    fmt.Print("\n")
    break
  case "putSetGenData":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "PutSetGenData requires 2 args")
      flag.Usage()
    }
    argvalue0, err98 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err98 != nil {
      Usage()
      return
    }
    value0 := generic.TMetaKey(argvalue0)
    arg99 := flag.Arg(2)
    mbTrans100 := thrift.NewTMemoryBufferLen(len(arg99))
    defer mbTrans100.Close()
    _, err101 := mbTrans100.WriteString(arg99)
    if err101 != nil {
      Usage()
      return
    }
    factory102 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt103 := factory102.GetProtocol(mbTrans100)
    argvalue1 := generic.NewTBigSetGenericData()
    err104 := argvalue1.Read(jsProt103)
    if err104 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.PutSetGenData(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "getTotalCount":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTotalCount requires 1 args")
      flag.Usage()
    }
    argvalue0, err105 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err105 != nil {
      Usage()
      return
    }
    value0 := generic.TContainerKey(argvalue0)
    fmt.Print(client.GetTotalCount(context.Background(), value0))
    fmt.Print("\n")
    break
  case "removeAll":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RemoveAll requires 1 args")
      flag.Usage()
    }
    argvalue0, err106 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err106 != nil {
      Usage()
      return
    }
    value0 := generic.TContainerKey(argvalue0)
    fmt.Print(client.RemoveAll(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
