package chaincode

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type CounterNO struct {
	Counter int `json:"counter"`
}


type User struct {
	UserId      string 			`json:"userId"`
	UserCode    string 			`json:"userCode"`
	PhoneNumber string 			`json:"phoneNumber"`
	Email       string 			`json:"email"`
	Password    string 			`json:"password"`
	FullName    string 			`json:"fullName"`
	UserName    string 			`json:"userName"`
	Address     string 			`json:"address"`
	Avatar     	string 			`json:"avatar"`
	Role        string 			`json:"role"`
	RoleId      int 			`json:"roleId"`
	Status      string 			`json:"status"`
	Signature   string 			`json:"signature"`
	Cart		[]ProductIdItem `json:"cart" metadata:",optional"`
}



type Actor struct {
	UserId      string `json:"userId"`
	UserCode    string `json:"userCode"`
	PhoneNumber string `json:"phoneNumber"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	Avatar     	string `json:"avatar"`
	Role        string `json:"role"`
}

type ProductDate struct {
	Status     	string 	 `json:"status"`
	Time 		string	 `json:"time"`
	Actor  		Actor 	 `json:"actor"`
}

type Product struct {
	ProductId      string         `json:"productId"`
	ProductCode    string 		  `json:"productCode"`
	ProductName    string         `json:"productName"`
	Supplier 	   Actor          `json:"supplier"`
	Dates          []ProductDate  `json:"dates" metadata:",optional"`
	Image          []string       `json:"image" metadata:",optional"`
	Expired        string         `json:"expireTime"`
	Price          string         `json:"price"`
	Amount         string         `json:"amount"`
	Unit           string         `json:"unit"`
	Status         string         `json:"status"`
	Description    string         `json:"description"`
	CertificateUrl string         `json:"certificateUrl"`
	QRCode		   string		  `json:"qrCode"`
}

type ProductCommercial struct {
	ProductCommercialId string         `json:"productCommercialId"`
	ProductId      		string         `json:"productId"`
	ProductCode    		string 		   `json:"productCode"`
	ProductName    		string         `json:"productName"`
	Dates          		[]ProductDate  `json:"dates" metadata:",optional"`
	Image          		[]string       `json:"image" metadata:",optional"`
	Expired        		string         `json:"expireTime"`
	Price          		string         `json:"price"`
	Unit           		string         `json:"unit"`
	Status         		string         `json:"status"`
	Description    		string         `json:"description"`
	CertificateUrl 		string         `json:"certificateUrl"`
	QRCode		   		string		   `json:"qrCode"`
}

type ProductPayload struct {
	ProductName    string        `json:"productName"`
	ProductCode    string        `json:"productCode"`
	Image          []string      `json:"image" metadata:",optional"`
	Price          string        `json:"price"`
	Amount         string        `json:"amount"`
	Unit           string        `json:"unit"`
	Description    string        `json:"description"`
	CertificateUrl string        `json:"certificateUrl"`
}

type ProductHistory struct {
	Record    		*Product  			`json:"record"`
	TransactionId   string    			`json:"transactionId"`
	Timestamp 		time.Time 			`json:"timestamp"`
	IsDelete  		bool      			`json:"isDelete"`
}

type ProductCommercialHistory struct {
	Record    		*ProductCommercial  `json:"record"`
	TransactionId   string    			`json:"transactionId"`
	Timestamp 		time.Time 			`json:"timestamp"`
	IsDelete  		bool      			`json:"isDelete"`
}

type OrderHistory struct {
	Record    		*Order    `json:"record"`
	TransactionId   string    `json:"transactionId"`
	Timestamp 		time.Time `json:"timestamp"`
	IsDelete  		bool      `json:"isDelete"`
}

type ProductItem struct {
	Product  Product `json:"product"`
	Quantity string  `json:"quantity"`
}

type ProductCommercialItem struct {
	Product  ProductCommercial 	`json:"product"`
	Quantity string  			`json:"quantity"`
}

type ProductIdItem struct {
	ProductId  	string 	`json:"productId"`
	Quantity 	string  `json:"quantity"`
}

type ProductIdQRCodeItem struct {
	ProductId  	string 	`json:"productId"`
	Quantity 	string  `json:"quantity"`
	QRCode 		string  `json:"qrCode"`
}

type ProductItemPayload struct {
	ProductId  	string 	`json:"productId"`
	Quantity 	string  `json:"quantity"`
}

type DeliveryStatus struct {
	Status       	string    	`json:"status"`
	DeliveryDate 	string		`json:"deliveryDate"`
	Address			string    	`json:"address"`
	Actor 			Actor 		`json:"actor"`
}

type DeliveryStatusCreateOrder struct {
	Address			string    	`json:"address"`
}

type Order struct {
	OrderId 		string      	 		`json:"orderId"`
	ProductItemList []ProductCommercialItem	`json:"productItemList" metadata:",optional"`
	DeliveryStatuses[]DeliveryStatus 		`json:"deliveryStatuses" metadata:",optional"`
	Signatures 		[]string 		 		`json:"signatures"`
	Status          string     	 	 		`json:"status"`
	CreateDate 		string 			 		`json:"createDate"`
	UpdateDate 		string 			 		`json:"updateDate"`
	FinishDate   	string      	 		`json:"finishDate"`
	QRCode		   	string		 	 		`json:"qrCode"`
	Retailer     	Actor 			 		`json:"retailer"`
	Manufacturer  	Actor 			 		`json:"manufacturer"`
	Distributor  	Actor 			 		`json:"distributor"`
}

type OrderForCreate struct {
	ProductIdQRCodeItems 	[]ProductIdQRCodeItem 		`json:"productIdQRCodeItems" metadata:",optional"`
	DeliveryStatus 			DeliveryStatusCreateOrder 	`json:"deliveryStatus"`
	Signatures 				[]string 					`json:"signatures"`
	QRCode		   			string		 				`json:"qrCode"`
}

type OrderForUpdateFinish struct {
	OrderId 		string      	 			`json:"orderId"`
	DeliveryStatus 	DeliveryStatusCreateOrder 	`json:"deliveryStatus"`
	Signature 		string 						`json:"signature"`
}

// Initialize chaincode
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	error := initCounter(ctx)
	if error != nil {
		return fmt.Errorf("error init counter: %s", error.Error())
	}
	return nil
}


func (s *SmartContract) CreateUser(ctx contractapi.TransactionContextInterface, userId string, userCode string, phoneNumber string, email string, password string, fullName string, userName string, address string, avatar string, role string, roleId int, status string) error {
	user := User{
		UserId:      userId,
		UserCode:    userCode,
		PhoneNumber: phoneNumber,
		Email:       email,
		Password:    password,
		FullName:    fullName,
		UserName:    userName,
		Address:     address,
		Avatar:      avatar,
		Role:        role,
		RoleId:      roleId,
		Status:      status,
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("error marshalling user to JSON: %s", err.Error())
	}
	err = ctx.GetStub().PutState(user.UserId, userJson)
	if err != nil {
		return fmt.Errorf("error putting user to state: %s", err.Error())
	}
	return nil
}
func (s *SmartContract) GetUser(ctx contractapi.TransactionContextInterface, userId string) (*User, error) {
	userJson, err := ctx.GetStub().GetState(userId)
	if err != nil {
		return nil, fmt.Errorf("error getting user from state: %s", err.Error())
	}
	if userJson == nil {
		return nil, fmt.Errorf("user not found")
	}
	var user User
	err = json.Unmarshal(userJson, &user)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling user from JSON: %s", err.Error())
	}
	return &user, nil
}
func (s *SmartContract) UpdateUser(ctx contractapi.TransactionContextInterface, userId string, userCode string, phoneNumber string, email string, password string, fullName string, userName string, address string, avatar string, role string, roleId int, status string) error {
	userJson, err := ctx.GetStub().GetState(userId)
	if err != nil {
		return fmt.Errorf("error getting user from state: %s", err.Error())
	}
	if userJson == nil {
		return fmt.Errorf("user not found")
	}
	var user User
	err = json.Unmarshal(userJson, &user)
	if err != nil {
		return fmt.Errorf("error unmarshalling user from JSON: %s", err.Error())
	}
	user.UserCode = userCode
	user.PhoneNumber = phoneNumber
	user.Email = email
	user.Password = password
	user.FullName = fullName
	user.UserName = userName
	user.Address = address
	user.Avatar = avatar
	user.Role = role
	user.RoleId = roleId
	user.Status = status

	userJson, err = json.Marshal(user)
	if err != nil {
		return fmt.Errorf("error marshalling user to JSON: %s", err.Error())
	}
	err = ctx.GetStub().PutState(user.UserId, userJson)
	if err != nil {
		return fmt.Errorf("error putting user to state: %s", err.Error())
	}
	return nil
}
func (s *SmartContract) DeleteUser(ctx contractapi.TransactionContextInterface, userId string) error {
	userJson, err := ctx.GetStub().GetState(userId)
	if err != nil {
		return fmt.Errorf("error getting user from state: %s", err.Error())
	}
	if userJson == nil {
		return fmt.Errorf("user not found")
	}
	var user User
	err = json.Unmarshal(userJson, &user)
	if err != nil {
		return fmt.Errorf("error unmarshalling user from JSON: %s", err.Error())
	}
	user.Status = "deleted"
	userJson, err = json.Marshal(user)
	if err != nil {
		return fmt.Errorf("error marshalling user to JSON: %s", err.Error())
	}
	err = ctx.GetStub().PutState(user.UserId, userJson)
	if err != nil {
		return fmt.Errorf("error putting user to state: %s", err.Error())
	}
	return nil
}
func (s *SmartContract) GetAllUsers(ctx contractapi.TransactionContextInterface) ([]User, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"user"}}`)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("error getting all users: %s", err.Error())
	}
	defer resultsIterator.Close()

	var users []User
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("error getting next user: %s", err.Error())
		}
		var user User
		err = json.Unmarshal(queryResponse.Value, &user)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling user from JSON: %s", err.Error())
		}
		users = append(users, user)
	}
	return users, nil
}


func (s *SmartContract) CreateProduct(ctx contractapi.TransactionContextInterface, productId string, productCode string, productName string, supplier Actor, dates []ProductDate, image []string, expired string, price string, amount string, unit string, status string, description string, certificateUrl string) error {
	product := Product{
		ProductId:      productId,
		ProductCode:    productCode,
		ProductName:    productName,
		Supplier:       supplier,
		Dates:          dates,
		Image:          image,
		Expired:        expired,
		Price:          price,
		Amount:         amount,
		Unit:           unit,
		Status:         status,
		Description:    description,
		CertificateUrl: certificateUrl,
	}
	productJson, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("error marshalling product to JSON: %s", err.Error())
	}
	err = ctx.GetStub().PutState(product.ProductId, productJson)
	if err != nil {
		return fmt.Errorf("error putting product to state: %s", err.Error())
	}
	return nil
}
func (s *SmartContract) GetProduct(ctx contractapi.TransactionContextInterface, productId string) (*Product, error) {
	productJson, err := ctx.GetStub().GetState(productId)
	if err != nil {
		return nil, fmt.Errorf("error getting product from state: %s", err.Error())
	}
	if productJson == nil {
		return nil, fmt.Errorf("product not found")
	}
	var product Product
	err = json.Unmarshal(productJson, &product)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling product from JSON: %s", err.Error())
	}
	return &product, nil
}
func (s *SmartContract) UpdateProduct(ctx contractapi.TransactionContextInterface, productId string, productCode string, productName string, supplier Actor, dates []ProductDate, image []string, expired string, price string, amount string, unit string, status string, description string, certificateUrl string) error {
	productJson, err := ctx.GetStub().GetState(productId)
	if err != nil {
		return fmt.Errorf("error getting product from state: %s", err.Error())
	}
	if productJson == nil {
		return fmt.Errorf("product not found")
	}
	var product Product
	err = json.Unmarshal(productJson, &product)
	if err != nil {
		return fmt.Errorf("error unmarshalling product from JSON: %s", err.Error())
	}
	product.ProductCode = productCode
	product.ProductName = productName
	product.Supplier = supplier
	product.Dates = dates
	product.Image = image
	product.Expired = expired
	product.Price = price
	product.Amount = amount
	product.Unit = unit
	product.Status = status
	product.Description = description
	product.CertificateUrl = certificateUrl

	productJson, err = json.Marshal(product)
	if err != nil {
		return fmt.Errorf("error marshalling product to JSON: %s", err.Error())
	}
	err = ctx.GetStub().PutState(product.ProductId, productJson)
	if err != nil {
		return fmt.Errorf("error putting product to state: %s", err.Error())
	}
	return nil
}
func (s *SmartContract) OrderManagement(ctx contractapi.TransactionContextInterface, orderId string, productItemList []ProductCommercialItem, deliveryStatuses []DeliveryStatus, signatures []string, status string, createDate string, updateDate string, finishDate string, qrCode string, retailer Actor, manufacturer Actor, distributor Actor) error {
	order := Order{
		OrderId:		 orderId,
		ProductItemList: productItemList,
		DeliveryStatuses: devliveryStatuses,
		Signatures:        signatures,
		Status:         status,
		CreateDate:   createDate,
		UpdateDate:   updateDate,
		FinishDate:   finishDate,
		QRCode:     qrCode,
		Retailer:   retailer,
		Manufacturer: manufacturer,
		Distributor: distributor,
	}
	orderJson , err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("errormarshalling order to json: %s", err.Error())
	}
	err = ctx.getStub().putState(order.OrderId, orderJson)
	if err != nil {
		return fmt.Errorf("error putting order to state : %s", err.Error())
	}
	return nil
}

func (s *SmartContract) GetOrder(ctx contractapi.TransactionContextInterface, orderId string) (*Order, error) {
	orderJson, err := ctx.GetStub().GetState(orderId)
	if err != nil {
		return nil, fmt.Errorf("error getting order from state: %s", err.Error())
	}
	if orderJson == nil {
		return nil, fmt.Errorf("order not found")
	}
	var order Order
	err = json.Unmarshal(orderJson, &order)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling order from JSON: %s", err.Error())
	}
	return &order, nil
}


func (s *SmartContract) UpdateOrder(ctx contractapi.TransactionContextInterface, orderId string, productItemList []ProductCommercialItem, deliveryStatuses []DeliveryStatus, signatures []string, status string, createDate string, updateDate string, finishDate string, qrCode string, retailer Actor, manufacturer Actor, distributor Actor) error {
	orderJson, err := ctx.GetStub().GetState(orderId)
	if err != nil {
		return fmt.Errorf("error getting order from state: %s", err.Error())
	}
	if orderJson == nil {
		return fmt.Errorf("order not found")
	}
	var order Order
	err = json.Unmarshal(orderJson, &order)
	if err != nil {
		return fmt.Errorf("error unmarshalling order from JSON: %s", err.Error())
	}
	order.ProductItemList = productItemList
	order.DeliveryStatuses = deliveryStatuses
	order.Signatures = signatures
	order.Status = status
	order.CreateDate = createDate
	order.UpdateDate = updateDate
	order.FinishDate = finishDate
	order.QRCode = qrCode
	order.Retailer = retailer
	order.Manufacturer = manufacturer
	order.Distributor = distributor

	orderJson, err = json.Marshal(order)
	if err != nil {
		return fmt.Errorf("error marshalling order to JSON: %s", err.Error())
	}
	err = ctx.GetStub().PutState(order.OrderId, orderJson)
	if err != nil {
		return fmt.Errorf("error putting order to state: %s", err.Error())
	}
	return nil
}
func (s *SmartContract) DeleteOrder(ctx contractapi.TransactionContextInterface, orderId string) error {
	orderJson, err := ctx.GetStub().GetState(orderId)
	if err != nil {
		return fmt.Errorf("error getting order from state: %s", err.Error())
	}
	if orderJson == nil {
		return fmt.Errorf("order not found")
	}
	var order Order
	err = json.Unmarshal(orderJson, &order)
	if err != nil {
		return fmt.Errorf("error unmarshalling order from JSON: %s", err.Error())
	}
	order.Status = "deleted"
	orderJson, err = json.Marshal(order)
	if err != nil {
		return fmt.Errorf("error marshalling order to JSON: %s", err.Error())
	}
	err = ctx.GetStub().PutState(order.OrderId, orderJson)
	if err != nil {
		return fmt.Errorf("error putting order to state: %s", err.Error())
	}
	return nil
}
func (s *SmartContract) GetAllOrders(ctx contractapi.TransactionContextInterface) ([]Order, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"order"}}`)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("error getting all orders: %s", err.Error())
	}
	defer resultsIterator.Close()

	var orders []Order
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("error getting next order: %s", err.Error())
		}
		var order Order
		err = json.Unmarshal(queryResponse.Value, &order)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling order from JSON: %s", err.Error())
		}
		orders = append(orders, order)
	}
	return orders, nil
}
func (s *SmartContract) CreateProductCommercial(ctx contractapi.TransactionContextInterface, productCommercialId string, productId string, productCode string, productName string, dates []ProductDate, image []string, expired string, price string, unit string, status string, description string, certificateUrl string) error {
	productCommercial := ProductCommercial{
		ProductCommercialId: productCommercialId,
		ProductId:           productId,
		ProductCode:         productCode,
		ProductName:         productName,
		Dates:               dates,
		Image:               image,
		Expired:             expired,
		Price:               price,
		Unit:                unit,
		Status:              status,
		Description:         description,
		CertificateUrl:      certificateUrl,
	}
	productCommercialJson, err := json.Marshal(productCommercial)
	if err != nil {
		return fmt.Errorf("error marshalling product commercial to JSON: %s", err.Error())
	}
	err = ctx.GetStub().PutState(productCommercial.ProductCommercialId, productCommercialJson)
	if err != nil {
		return fmt.Errorf("error putting product commercial to state: %s", err.Error())
	}
	return nil
}