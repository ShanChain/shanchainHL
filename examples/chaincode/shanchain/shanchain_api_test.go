package main

import (
  "encoding/json"
  "fmt"
  "testing"

  "github.com/hyperledger/fabric/core/chaincode/shim"
)

type Transaction2 struct {
  ID  string // 交易流水号
  Step int  //转化步数
  Integral int  //交易积分
  FromType int //交易发起方类型
  FromID string //交易发起方id
  ToType int //交易接受方类型
  ToID string //交易接收方id
}

func checkInit(t *testing.T, stub *shim.MockStub, args []string) (error) {
  _, err := stub.MockInit("1", "init", args)
  return err
}

func checkUnknown(t *testing.T, stub *shim.MockStub, args []string) (error) {
  _, err := stub.MockInvoke("1", "helloworld", args)
  return err
}

func checkState(t *testing.T, stub *shim.MockStub, name string, value string) {
  bytes := stub.State[name]
  if bytes == nil {
    fmt.Println("State", name, "failed to get value")
    t.FailNow()
  }
  if string(bytes) != value {
    fmt.Println("State value", name, "was not", value, "as expected")
    t.FailNow()
  }
}

func checkCreateUser(t *testing.T, stub *shim.MockStub, args []string) (error) {
  _, err := stub.MockInvoke("1", "createUser", args)
  return err
}

func checkGetUser(t *testing.T, stub *shim.MockStub, args []string) ([]byte, error) {
  bytes, err := stub.MockQuery("getUser", args)
  return bytes, err
}

func checkGetRoot(t *testing.T, stub *shim.MockStub, args []string) ([]byte, error) {
  bytes, err := stub.MockQuery("getRoot", args)
  return bytes, err
}

func checkAdditional(t *testing.T, stub *shim.MockStub, args []string) ([]byte, error) {
  ts, err := stub.MockInvoke("1", "additional", args)
  return ts, err
}

func checkExchange(t *testing.T, stub *shim.MockStub, args []string) ([]byte, error) {
  ts, err := stub.MockInvoke("1", "exchange", args)
  return ts, err
}

func checkTransfer(t *testing.T, stub *shim.MockStub, args []string) ([]byte, error) {
  ts, err := stub.MockInvoke("1", "transfer", args)
  return ts, err
}

func checkTransaction(t *testing.T, stub *shim.MockStub, args []string, expect string) {
  var transaction Transaction2
  bytes, err := stub.MockQuery("getTransaction", args)
  if err != nil {
    fmt.Println("Query", args, "failed", err)
    t.FailNow()
  }
  if bytes == nil {
    fmt.Println("Query", args, "failed to get result")
    t.FailNow()
  }
  err = json.Unmarshal(bytes, &transaction)
  if err != nil {
    fmt.Println("Error unmarshalling transaction")
  }
  bytes, err  = json.Marshal(&transaction)
  if err != nil {
    fmt.Println("Error retrieving rootBytes")
  }
  if string(bytes) != expect {
    fmt.Println("Query result ", string(bytes), "was not", expect, "as expected")
    t.FailNow()
  }
}


func TestShanchain_Init(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  // Init RootName=shanchain , InitIntegral=50000
  checkInit(t, stub, []string{"shanchain", "50000"})
  checkState(t, stub, "root", "{\"ID\":\"0001\",\"Name\":\"shanchain\",\"TotalIntegral\":50000,\"RestIntegral\":50000}")
}

func TestShanchain_Init_Incorrect_arguments(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  err := checkInit(t, stub, []string{"shanchain", "50000", "100"})
  if err != nil {
    if err.Error() == "Incorrect number of arguments. Expecting 2" {
      t.SkipNow()
    }
  } else {
    t.FailNow()
  }
}

func TestShanchain_Init_Invalid_argument(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  err := checkInit(t, stub, []string{"shanchain", "hello"})
  if err != nil {
    if err.Error() == "Expecting integer value for asset holding" {
      t.SkipNow()
    }
  } else {
    t.FailNow()
  }
}

func TestShanchain_Invoke_Unknown_function(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  err := checkUnknown(t, stub, []string{"shanchain", "hello"})
  if err != nil {
    if err.Error() == "Received unknown function invocation" {
      t.SkipNow()
    } else {
      t.FailNow()
    }
  }
}

func TestShanchain_GetRoot(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  // Init RootName=shanchain , InitIntegral=50000
  checkInit(t, stub, []string{"shanchain", "50000"})
  bytes, err := checkGetRoot(t, stub, []string{})
  if err != nil {
    t.FailNow()
  }
  if bytes == nil {
    t.FailNow()
  }
  if string(bytes) != "{\"ID\":\"0001\",\"Name\":\"shanchain\",\"TotalIntegral\":50000,\"RestIntegral\":50000}" {
    t.FailNow()
  }
}

func TestShanchain_GetRoot_Incorrect_arguments(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  checkInit(t, stub, []string{"shanchain", "50000"})
  _, err := checkGetRoot(t, stub, []string{"hello"})
  if err != nil {
    if err.Error() == "Incorrect number of arguments. Expecting 0" {
      t.SkipNow()
    } else {
      t.FailNow()
    }
  }
}

func TestShanchain_CreateUser(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  err := checkCreateUser(t, stub, []string{"10086", "china mobile", "100"})
  if err != nil {
    t.FailNow()
  }
}

func TestShanchain_CreateUser_Incorrect_arguments(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  err := checkCreateUser(t, stub, []string{"10086", "china mobile", "100", "name"})
  if err != nil {
    if err.Error() == "Incorrect number of arguments. Expecting 3" {
      t.SkipNow()
    }
  }
  t.FailNow()
}

func TestShanchain_CreateUser_Invalid_argument(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  err := checkCreateUser(t, stub, []string{"10086", "china mobile", "number"})
  if err != nil {
    if err.Error() == "Expecting integer value for asset holding" {
      t.SkipNow()
    }
  }
  t.FailNow()
}

func TestShanchain_GetUser(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)

  checkCreateUser(t, stub, []string{"10086", "china mobile", "100"})
  bytes, err := checkGetUser(t, stub, []string{"10086"})
  if err != nil {
    t.FailNow()
  }
  if bytes == nil {
    t.FailNow()
  }
  if string(bytes) != "{\"ID\":\"10086\",\"Name\":\"china mobile\",\"Integral\":100}" {
    t.FailNow()
  }
}

func TestShanchain_GetUser_Incorrect_arguments(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)

  checkCreateUser(t, stub, []string{"10086", "china mobile", "100"})
  _, err := checkGetUser(t, stub, []string{"10086", "hello"})
  if err != nil {
    if err.Error() == "Incorrect number of arguments. Expecting 1" {
      t.SkipNow()
    }
  }
  t.FailNow()
}

func TestShanchain_Additional(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  // Init RootName=shanchain , InitIntegral=50000
  checkInit(t, stub, []string{"shanchain", "50000"})
  _, err := checkAdditional(t, stub, []string{"10000"})
  if err !=nil {
    t.FailNow()
  }
  bytes, err := checkGetRoot(t, stub, []string{})
  if string(bytes) != "{\"ID\":\"0001\",\"Name\":\"shanchain\",\"TotalIntegral\":60000,\"RestIntegral\":60000}" {
    t.FailNow()
  }
}

func TestShanchain_Additional_Invalid_argument1(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  // Init RootName=shanchain , InitIntegral=50000
  checkInit(t, stub, []string{"shanchain", "50000"})
  _, err := checkAdditional(t, stub, []string{"d0000"})
  if err !=nil {
    if err.Error() == "want Integer number" {
      t.SkipNow()
    }
  }
  t.FailNow()
}

func TestShanchain_Additional_Invalid_argument2(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  // Init RootName=shanchain , InitIntegral=50000
  checkInit(t, stub, []string{"shanchain", "50000"})
  _, err := checkAdditional(t, stub, []string{"-9"})
  if err !=nil {
    if err.Error() == "want positive Integer number" {
      t.SkipNow()
    }
  }
  t.FailNow()
}

func TestShanchain_Exchange(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  // Init RootName=shanchain , InitIntegral=50000
  checkInit(t, stub, []string{"shanchain", "50000"})
  checkCreateUser(t, stub, []string{"10086", "china mobile", "100"})
  _, err := checkExchange(t, stub, []string{"10086", "900"})
  if err != nil {
    t.FailNow()
  }
  bytes, err := checkGetUser(t, stub, []string{"10086"})
  if err != nil {
    t.FailNow()
  }
  if bytes == nil {
    t.FailNow()
  }
  if string(bytes) != "{\"ID\":\"10086\",\"Name\":\"china mobile\",\"Integral\":1000}" {
    t.FailNow()
  }
}

func TestShanchain_Exchange_Incorrect_arguments(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  checkInit(t, stub, []string{"shanchain", "50000"})
  checkCreateUser(t, stub, []string{"10086", "china mobile", "100"})
  _, err := checkExchange(t, stub, []string{"10086", "900", "more"})
  if err != nil {
    if err.Error() == "Incorrect number of arguments. Expecting 2" {
      t.SkipNow()
    }
  }
  t.FailNow()
}

func TestShanchain_Exchange_Invalid_argument1(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  checkInit(t, stub, []string{"shanchain", "50000"})
  checkCreateUser(t, stub, []string{"10086", "china mobile", "100"})
  _, err := checkExchange(t, stub, []string{"10086", "f00"})
  if err != nil {
    if err.Error() == "want Integer number" {
      t.SkipNow()
    }
  }
  t.FailNow()
}

func TestShanchain_Exchange_Invalid_argument2(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  checkInit(t, stub, []string{"shanchain", "50000"})
  checkCreateUser(t, stub, []string{"10086", "china mobile", "100"})
  _, err := checkExchange(t, stub, []string{"10086", "60000"})
  if err != nil {
    if err.Error() == "Root 剩余善圆不足" {
      t.SkipNow()
    }
  }
  t.FailNow()
}

func TestShanchain_Transfer(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  checkInit(t, stub, []string{"shanchain", "50000"})
  checkCreateUser(t, stub, []string{"10086", "china mobile", "1000"})
  checkCreateUser(t, stub, []string{"10000", "china unicom", "1000"})
  _, err := checkTransfer(t, stub, []string{"10086", "10000", "200"})
  if err != nil {
    t.FailNow()
  }
  abytes, err := checkGetUser(t, stub, []string{"10086"})
  if string(abytes) != "{\"ID\":\"10086\",\"Name\":\"china mobile\",\"Integral\":800}" {
    t.FailNow()
  }
  bbytes, err := checkGetUser(t, stub, []string{"10000"})
  if string(bbytes) != "{\"ID\":\"10000\",\"Name\":\"china unicom\",\"Integral\":1200}" {
    t.FailNow()
  }
}

func TestShanchain_Transfer_Invalid_argument1(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  checkInit(t, stub, []string{"shanchain", "50000"})
  checkCreateUser(t, stub, []string{"10086", "china mobile", "1000"})
  checkCreateUser(t, stub, []string{"10000", "china unicom", "1000"})
  _, err := checkTransfer(t, stub, []string{"10086", "10000", "u00"})
  if err != nil {
    if err.Error() == "want Integer number" {
      t.SkipNow()
    }
  }
  t.FailNow()
}

func TestShanchain_Transfer_Invalid_argument2(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  checkInit(t, stub, []string{"shanchain", "50000"})
  checkCreateUser(t, stub, []string{"10086", "china mobile", "1000"})
  checkCreateUser(t, stub, []string{"10000", "china unicom", "1000"})
  _, err := checkTransfer(t, stub, []string{"10086", "10000", "3000"})
  if err != nil {
    if err.Error() == "用户 china mobile 剩余善圆不足本次交易" {
      t.SkipNow()
    }
  }
  t.FailNow()
}


func TestShanchain_Transaction(t *testing.T) {
  scc := new(ShanChainAPI)
  stub := shim.NewMockStub("shanchain_api", scc)
  checkInit(t, stub, []string{"shanchain", "50000"})
  checkCreateUser(t, stub, []string{"10086", "china mobile", "1000"})
  checkCreateUser(t, stub, []string{"10000", "china unicom", "1000"})
  checkTransfer(t, stub, []string{"10086", "10000", "200"})
  checkTransaction(t, stub, []string{"1"},"{\"ID\":\"1\",\"Step\":0,\"Integral\":200,\"FromType\":1,\"FromID\":\"10086\",\"ToType\":1,\"ToID\":\"10000\"}")
}
