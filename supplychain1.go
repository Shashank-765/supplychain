package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type CoffeeData struct {
	ID          string `json:"id"`
	FarmerData  *FarmerData  `json:"farmerData,omitempty"`
	ProcessorData *ProcessorData `json:"processorData,omitempty"`
	ManufactureData *ManufactureData `json:"manufactureData,omitempty"`
	CertificationData *CertificationData `json:"certificationData,omitempty"`
	MarketData *MarketData `json:"marketData,omitempty"`
}

type FarmerData struct {
	FarmName       string `json:"farmName"`
	FarmAddress    string `json:"farmAddress"`
	GPSCoordinates string `json:"gpsCoordinates"`
	CoffeeTypes    string `json:"coffeeTypes"`
	SeedlingAge    string `json:"seedlingAge"`
	SeedlingBatch  string `json:"seedlingBatch"`
	GrowthDetails  string `json:"growthDetails"`
	HarvestingData string `json:"harvestingData"`
	TransactionDetails string `json:"transactionDetails"`
}

type ProcessorData struct {
	ProcessorName  string `json:"processorName"`
	Batch          string `json:"batch"`
	CoffeeBeanGrade string `json:"coffeeBeanGrade"`
	RoastingMethod string `json:"roastingMethod"`
	TransactionDetails string `json:"transactionDetails"`
}

type ManufactureData struct {
	ManufacturerName string `json:"manufacturerName"`
	ManufacturerAddress string `json:"manufacturerAddress"`
	ProcessDetails string `json:"processDetails"`
	ProcessingBatch string `json:"processingBatch"`
	CoffeeBeanGrades string `json:"coffeeBeanGrades"`
	Volume string `json:"volume"`
	ProductionCode string `json:"productionCode"`
}

type CertificationData struct {
	ProductionCode string `json:"productionCode"`
	CertificationNumber string `json:"certificationNumber"`
	ManufacturerName string `json:"manufacturerName"`
	Location string `json:"location"`
	DateIssued string `json:"dateIssued"`
	CoffeeBeanGrades string `json:"coffeeBeanGrades"`
	CertificationAgencyName string `json:"certificationAgencyName"`
}

type MarketData struct {
	ProductionCode string `json:"productionCode"`
	CertificationNumber string `json:"certificationNumber"`
	MarketName string `json:"marketName"`
	Location string `json:"location"`
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`
	SupplyDate string `json:"supplyDate"`
	ProcessingDate string `json:"processingDate"`
	TransactionDetails string `json:"transactionDetails"`
}

type UserRole struct {
	UserID string `json:"userId"`
	Role   string `json:"role"`
}

func (s *SmartContract) RegisterUserRole(ctx contractapi.TransactionContextInterface, role string) error {
	clientID, err := s.getClientID(ctx)
	if err != nil {
		return err
	}

	userRole := UserRole{
		UserID: clientID,
		Role:   role,
	}

	dataAsBytes, err := json.Marshal(userRole)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState("ROLE_"+clientID, dataAsBytes)
}

func (s *SmartContract) checkUserRole(ctx contractapi.TransactionContextInterface, expectedRole string) error {
	clientID, err := s.getClientID(ctx)
	if err != nil {
		return err
	}

	dataAsBytes, err := ctx.GetStub().GetState("ROLE_" + clientID)
	if err != nil {
		return err
	}
	if dataAsBytes == nil {
		return fmt.Errorf("role not registered for user: %s", clientID)
	}

	var userRole UserRole
	if err := json.Unmarshal(dataAsBytes, &userRole); err != nil {
		return err
	}

	if userRole.Role != expectedRole {
		return fmt.Errorf("unauthorized: expected role '%s', but user has role '%s'", expectedRole, userRole.Role)
	}

	return nil
}

func (s *SmartContract) getClientID(ctx contractapi.TransactionContextInterface) (string, error) {
	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return "", fmt.Errorf("failed to get client identity: %v", err)
	}
	return clientID, nil
}

func (s *SmartContract) InitCoffeeData(ctx contractapi.TransactionContextInterface, id string) error {
	data := CoffeeData{ID: id}
	dataAsBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, dataAsBytes)
}

func (s *SmartContract) AddFarmerData(ctx contractapi.TransactionContextInterface, id string, farmerJSON string) error {
	if err := s.checkUserRole(ctx, "FARM_INSPECTION"); err != nil {
		return err
	}
	data, err := s.getCoffeeData(ctx, id)
	if err != nil {
		return err
	}
	var farmer FarmerData
	if err := json.Unmarshal([]byte(farmerJSON), &farmer); err != nil {
		return err
	}
	data.FarmerData = &farmer
	return s.saveCoffeeData(ctx, id, data)
}

func (s *SmartContract) AddProcessorData(ctx contractapi.TransactionContextInterface, id string, processorJSON string) error {
	if err := s.checkUserRole(ctx, "PROCESSOR"); err != nil {
		return err
	}
	data, err := s.getCoffeeData(ctx, id)
	if err != nil {
		return err
	}
	var processor ProcessorData
	if err := json.Unmarshal([]byte(processorJSON), &processor); err != nil {
		return err
	}
	data.ProcessorData = &processor
	return s.saveCoffeeData(ctx, id, data)
}

func (s *SmartContract) AddManufactureData(ctx contractapi.TransactionContextInterface, id string, manufactureJSON string) error {
	if err := s.checkUserRole(ctx, "EXPORTER"); err != nil {
		return err
	}
	data, err := s.getCoffeeData(ctx, id)
	if err != nil {
		return err
	}
	var manufacture ManufactureData
	if err := json.Unmarshal([]byte(manufactureJSON), &manufacture); err != nil {
		return err
	}
	data.ManufactureData = &manufacture
	return s.saveCoffeeData(ctx, id, data)
}

func (s *SmartContract) AddCertificationData(ctx contractapi.TransactionContextInterface, id string, certificationJSON string) error {
	if err := s.checkUserRole(ctx, "IMPORTER"); err != nil {
		return err
	}
	data, err := s.getCoffeeData(ctx, id)
	if err != nil {
		return err
	}
	var certification CertificationData
	if err := json.Unmarshal([]byte(certificationJSON), &certification); err != nil {
		return err
	}
	data.CertificationData = &certification
	return s.saveCoffeeData(ctx, id, data)
}

func (s *SmartContract) AddMarketData(ctx contractapi.TransactionContextInterface, id string, marketJSON string) error {
	if err := s.checkUserRole(ctx, "HARVESTER"); err != nil {
		return err
	}
	data, err := s.getCoffeeData(ctx, id)
	if err != nil {
		return err
	}
	var market MarketData
	if err := json.Unmarshal([]byte(marketJSON), &market); err != nil {
		return err
	}
	data.MarketData = &market
	return s.saveCoffeeData(ctx, id, data)
}

func (s *SmartContract) GetCompleteCoffeeData(ctx contractapi.TransactionContextInterface, id string) (*CoffeeData, error) {
	return s.getCoffeeData(ctx, id)
}

func (s *SmartContract) getCoffeeData(ctx contractapi.TransactionContextInterface, id string) (*CoffeeData, error) {
	dataAsBytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state: %v", err)
	}
	if dataAsBytes == nil {
		return nil, fmt.Errorf("Data not found for ID: %s", id)
	}
	var data CoffeeData
	if err := json.Unmarshal(dataAsBytes, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *SmartContract) saveCoffeeData(ctx contractapi.TransactionContextInterface, id string, data *CoffeeData) error {
	dataAsBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, dataAsBytes)
}

func (s*)

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating coffee chaincode: %v", err)
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting coffee chaincode: %v", err)
	}
}