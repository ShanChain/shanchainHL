/*
	author:krew
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	// "time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	// "github.com/satori/go.uuid"
)

/**
 * 
 */
type ShanChainAPI struct {
}

/**
 * 
 */
type Root struct {
	ID string  //id
	Name string  //名称
	TotalIntegral int  //总积分
	RestIntegral int //剩余积分
}

/**
 * 
 */
type User struct {
	ID string //用户id
	Name string //用户名称
	Integral int //积分
}

/**
 * 
 */
type Transaction struct {
	ID  string // 交易流水号
	Step int  //转化步数
	Integral int  //交易积分
	FromType int //交易发起方类型
	FromID string //交易发起方id
	ToType int //交易接受方类型
	ToID string //交易接收方id
	Time int64  //发起交易时间戳
}

/**
 * [main description]
 * @return {[type]} [description]
 */
func main()  {
	err := shim.Start(new(ShanChainAPI))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

/**
 * [func description]
 * @param  {[type]} t *ShanChainAPI) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error [description]
 * @return {[type]}   [description]
 */
func (t *ShanChainAPI) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error)  {
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}
	var root Root
	var totalIntegral int
	var restIntegral int
	var rootBytes []byte
	totalIntegral, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}
	restIntegral = totalIntegral
	root = Root{
		Name : args[0],
		TotalIntegral : totalIntegral,
		RestIntegral : restIntegral,
		ID : "0001"}
	err = writeRoot(stub, root)
	if err != nil {
		return nil, errors.New("writeRoot Error" + err.Error())
	}
	rootBytes, err  = json.Marshal(&root)
	if err != nil {
		return nil,errors.New("Error retrieving rootBytes")
	}
	return rootBytes, nil
}

/**
 * [func description]
 * @param  {[type]} t *ShanChainAPI) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error [description]
 * @return {[type]}   [description]
 */
func (t *ShanChainAPI) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "createUser" {
		return t.createUser(stub, args)
	} 
	// else if function == "transfer" {
	// 	return t.transfer(stub, args)
	// }
	return nil, errors.New("Received unknown function invocation")
}

/**
 * [func description]
 * @param  {[type]} t *ShanChainAPI) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error [description]
 * @return {[type]}   [description]
 */
func (t *ShanChainAPI) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error)  {
	argsLength := len(args)
	if function == "getRoot" {
		if argsLength != 0 {
			return nil, errors.New("Incorrect number of arguments. Expecting 0")
		}
		_, rootBytes, err := getRoot(stub)
		if err != nil {
			return nil, err
		}
		return rootBytes, nil
	} else if function == "getUser"{
		if argsLength != 1 {
			return nil, errors.New("Incorrect number of arguments. Expecting 1")
		}
		id := args[0]
		_, rootBytes, err := getUser(stub, id)
		if err != nil {
			return nil, err
		}
		return rootBytes, nil
	}
	return nil, nil
}

/**
 * [func description]
 * @param  {[type]} t *ShanChainAPI) createUser(stub *shim.ChaincodeStub, args []string) ([]byte, error [description]
 * @return {[type]}   [description]
 */
func (t *ShanChainAPI) createUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}
	var user User
	var id string
	var name string
	var integral int
	var userBytes []byte

	id = args[0]
	name = args[1]
	integral, err := strconv.Atoi(args[2])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}
	user = User{
		ID : id,
		Name : name,
		Integral : integral}
	err = writeUser(stub, user)
	if err != nil {
		return nil, errors.New("writeUser Error" + err.Error())
	}
	userBytes, err  = json.Marshal(&user)
	if err != nil {
		return nil, errors.New("Error retrieving userBytes")
	}
	return userBytes, nil
}


/**
 * [func description]
 * @param  {[type]} t *             ShanChainAPI) exchange(stub *shim.ChaincodeStub, args []string) ([]byte, error [description]
 * @return {[type]}   [description]
 */
// func (t * ShanChainAPI) exchange(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
// 	if len(args) != 3 {
// 		return nil, errors.New("Incorrect number of arguments. Expecting 3")
// 	}
// 	var (
// 		root Root
// 		user User
// 		number int
// 		accepterID string
// 		step int
// 		err error
// 		transaction Transaction
// 		tsBytes []byte
// 	)
// 	accepterID = args[0]
// 	number, err = strconv.Atoi(args[1])
// 	if err != nil {
// 		return nil,errors.New("want Integer number")
// 	}
// 	step, err = strconv.Atoi(args[2])
// 	if err != nil {
// 		return nil,errors.New("want Integer number")
// 	}
// 	root, _, err = getRoot(stub)
// 	if err != nil {
// 		return nil,errors.New("get root errors")
// 	}
// 	if root.RestIntegral < number {
// 		return nil,errors.New("Root 剩余善圆不足")
// 	}
// 	user, _, err = getUser(stub, accepterID)
// 	if err != nil {
// 		return nil,errors.New("get user errors")
// 	}

// 	root.RestIntegral = root.RestIntegral - number
// 	user.Integral = user.Integral + number

// 	err = writeRoot(stub, root)
// 	if err != nil {
// 		root.RestIntegral = root.RestIntegral - number
// 		user.Integral = user.Integral - number
// 		return nil, errors.New("write root errors" + err.Error())
// 	}
// 	err = writeUser(stub, user)
// 	if err != nil {
// 		root.RestIntegral = root.RestIntegral - number
// 		user.Integral = user.Integral - number
// 		err = writeRoot(stub, root)
// 		if err != nil {
// 			return nil, errors.New("roll down errors" + err.Error())
// 		}
// 		return nil, err
// 	}
	
// 	var domain byte
// 	id := uuid.NewV2(domain)
// 	transaction = Transaction{
// 		Step : step,
// 		Integral : number,
// 		FromType : 0,
// 		FromID : "0001",
// 		ToType : 1,
// 		ToID : accepterID,
// 		Time : time.Now().Unix()}
	
// 	err = writeTransaction(stub, transaction)
// 	if err != nil {
// 		return nil, errors.New("write transaction Error" + err.Error())
// 	}
// 	tsBytes, err = json.Marshal(&transaction)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return nil, nil
// }


/**
 * [func description]
 * @param  {[type]} t *ShanChainAPI) transfer(stub *shim.ChaincodeStub, args []string) ([]byte, error [description]
 * @return {[type]}   [description]
 */
// func (t *ShanChainAPI) transfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error){
// 	if len(args) != 3 {
// 		return nil, errors.New("Incorrect number of arguments. Expecting 3")
// 	}
// 	return nil, nil
// }


/**
 * [writeRoot description]
 * @param  {[type]} stub *shim.ChaincodeStub [description]
 * @param  {[type]} root Root)               (error        [description]
 * @return {[type]}      [description]
 */
func writeRoot(stub shim.ChaincodeStubInterface, root Root) (error) {
	rootBytes, err := json.Marshal(root)
	if err != nil {
		return err
	}
	err = stub.PutState("root", rootBytes)
	if err != nil {
		return errors.New("PutState Error" + err.Error())
	}
	return nil
}


/**
 * [getRoot description]
 * @param  {[type]} stub *shim.ChaincodeStub) (Root, []byte, error [description]
 * @return {[type]}      [description]
 */
func getRoot(stub shim.ChaincodeStubInterface) (Root, []byte, error) {
	var root Root
	rootBytyes, err := stub.GetState("root")
	if err != nil {
		fmt.Println("Error retrieving rootBytyes")
	}
	err = json.Unmarshal(rootBytyes, &root)
	if err != nil {
		fmt.Println("Error unmarshalling root")
	}
	return root, rootBytyes, nil
}


/**
 * [writeUser description]
 * @param  {[type]} stub *shim.ChaincodeStub [description]
 * @param  {[type]} user User)               (error        [description]
 * @return {[type]}      [description]
 */
func writeUser(stub shim.ChaincodeStubInterface, user User) (error) {
	id := user.ID
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = stub.PutState(id, userBytes)
	if err != nil {
		return errors.New("PutState Error" + err.Error())
	}
	return nil
}

/**
 * [getUser description]
 * @param  {[type]} stub *shim.ChaincodeStub [description]
 * @param  {[type]} id   string)             (User,        []byte, error [description]
 * @return {[type]}      [description]
 */
func getUser(stub shim.ChaincodeStubInterface, id string) (User, []byte, error) {
	var user User
	userBytes, err := stub.GetState(id)
	if err != nil {
		fmt.Println("Error retrieving userBytes")
	}
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		fmt.Println("Error unmarshalling user")
	}
	return user, userBytes, nil
}


/**
 * [writeTransaction description]
 * @param  {[type]} stub        *shim.ChaincodeStub [description]
 * @param  {[type]} transaction Transaction)        (error        [description]
 * @return {[type]}             [description]
 */
func writeTransaction(stub shim.ChaincodeStubInterface,transaction Transaction) (error) {
	var tsID string
	tsBytes, err := json.Marshal(&transaction)
	if err != nil {
		return err
	}
	tsID = transaction.ID
	err = stub.PutState(tsID, tsBytes)
	if err != nil {
		return errors.New("PutState Error" + err.Error())
	}
	return nil
}

