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
	"go2o/core/service/auto_gen/rpc/ttype"
        "go2o/core/service/auto_gen/rpc/order_service"
)

var _ = ttype.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  Result WholesaleCartV1(i64 memberId, string action,  data)")
  fmt.Fprintln(os.Stderr, "  Result RetailCartV1(i64 memberId, string action,  data)")
  fmt.Fprintln(os.Stderr, "   SubmitOrderV1(i64 buyerId, i32 cartType,  data)")
  fmt.Fprintln(os.Stderr, "  SComplexOrder GetOrder(string order_no, bool sub_order)")
  fmt.Fprintln(os.Stderr, "  SComplexOrder GetOrderAndItems(string order_no, bool sub_order)")
  fmt.Fprintln(os.Stderr, "  SComplexOrder GetSubOrder(i64 id)")
  fmt.Fprintln(os.Stderr, "  SComplexOrder GetSubOrderByNo(string orderNo)")
  fmt.Fprintln(os.Stderr, "   GetSubOrderItems(i64 subOrderId)")
  fmt.Fprintln(os.Stderr, "  Result SubmitTradeOrder(SComplexOrder o, double rate)")
  fmt.Fprintln(os.Stderr, "  Result TradeOrderCashPay(i64 orderId)")
  fmt.Fprintln(os.Stderr, "  Result TradeOrderUpdateTicket(i64 orderId, string img)")
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
  client := order_service.NewOrderServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "WholesaleCartV1":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "WholesaleCartV1 requires 3 args")
      flag.Usage()
    }
    argvalue0, err38 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err38 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    arg40 := flag.Arg(3)
    mbTrans41 := thrift.NewTMemoryBufferLen(len(arg40))
    defer mbTrans41.Close()
    _, err42 := mbTrans41.WriteString(arg40)
    if err42 != nil { 
      Usage()
      return
    }
    factory43 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt44 := factory43.GetProtocol(mbTrans41)
    containerStruct2 := order_service.NewOrderServiceWholesaleCartV1Args()
    err45 := containerStruct2.ReadField3(jsProt44)
    if err45 != nil {
      Usage()
      return
    }
    argvalue2 := containerStruct2.Data
    value2 := argvalue2
    fmt.Print(client.WholesaleCartV1(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "RetailCartV1":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "RetailCartV1 requires 3 args")
      flag.Usage()
    }
    argvalue0, err46 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err46 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    arg48 := flag.Arg(3)
    mbTrans49 := thrift.NewTMemoryBufferLen(len(arg48))
    defer mbTrans49.Close()
    _, err50 := mbTrans49.WriteString(arg48)
    if err50 != nil { 
      Usage()
      return
    }
    factory51 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt52 := factory51.GetProtocol(mbTrans49)
    containerStruct2 := order_service.NewOrderServiceRetailCartV1Args()
    err53 := containerStruct2.ReadField3(jsProt52)
    if err53 != nil {
      Usage()
      return
    }
    argvalue2 := containerStruct2.Data
    value2 := argvalue2
    fmt.Print(client.RetailCartV1(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "SubmitOrderV1":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "SubmitOrderV1 requires 3 args")
      flag.Usage()
    }
    argvalue0, err54 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err54 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err55 := (strconv.Atoi(flag.Arg(2)))
    if err55 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    arg56 := flag.Arg(3)
    mbTrans57 := thrift.NewTMemoryBufferLen(len(arg56))
    defer mbTrans57.Close()
    _, err58 := mbTrans57.WriteString(arg56)
    if err58 != nil { 
      Usage()
      return
    }
    factory59 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt60 := factory59.GetProtocol(mbTrans57)
    containerStruct2 := order_service.NewOrderServiceSubmitOrderV1Args()
    err61 := containerStruct2.ReadField3(jsProt60)
    if err61 != nil {
      Usage()
      return
    }
    argvalue2 := containerStruct2.Data
    value2 := argvalue2
    fmt.Print(client.SubmitOrderV1(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "GetOrder":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetOrder requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2) == "true"
    value1 := argvalue1
    fmt.Print(client.GetOrder(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "GetOrderAndItems":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetOrderAndItems requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2) == "true"
    value1 := argvalue1
    fmt.Print(client.GetOrderAndItems(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "GetSubOrder":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetSubOrder requires 1 args")
      flag.Usage()
    }
    argvalue0, err66 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err66 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetSubOrder(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetSubOrderByNo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetSubOrderByNo requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetSubOrderByNo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetSubOrderItems":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetSubOrderItems requires 1 args")
      flag.Usage()
    }
    argvalue0, err68 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err68 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetSubOrderItems(context.Background(), value0))
    fmt.Print("\n")
    break
  case "SubmitTradeOrder":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SubmitTradeOrder requires 2 args")
      flag.Usage()
    }
    arg69 := flag.Arg(1)
    mbTrans70 := thrift.NewTMemoryBufferLen(len(arg69))
    defer mbTrans70.Close()
    _, err71 := mbTrans70.WriteString(arg69)
    if err71 != nil {
      Usage()
      return
    }
    factory72 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt73 := factory72.GetProtocol(mbTrans70)
    argvalue0 := order_service.NewSComplexOrder()
    err74 := argvalue0.Read(jsProt73)
    if err74 != nil {
      Usage()
      return
    }
    value0 := order_service.SComplexOrder(argvalue0)
    argvalue1, err75 := (strconv.ParseFloat(flag.Arg(2), 64))
    if err75 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.SubmitTradeOrder(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "TradeOrderCashPay":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "TradeOrderCashPay requires 1 args")
      flag.Usage()
    }
    argvalue0, err76 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err76 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.TradeOrderCashPay(context.Background(), value0))
    fmt.Print("\n")
    break
  case "TradeOrderUpdateTicket":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "TradeOrderUpdateTicket requires 2 args")
      flag.Usage()
    }
    argvalue0, err77 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err77 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.TradeOrderUpdateTicket(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}