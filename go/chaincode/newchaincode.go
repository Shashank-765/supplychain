package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

// Structs as provided previously...
type User struct {
	UserID        string `json:"userId"`
	UserType      string `json:"userType"`
	UserName      string `json:"userName"`
	UserEmail     string `json:"userEmail"`
	UserPhone     string `json:"userPhone"`
	UserAddress   string `json:"userAddress"`
	UserStatus    string `json:"userStatus"`
	UserCreatedAt string `json:"userCreated"`
	UserUpdatedAt string `json:"userUpdated"`
	UserDeletedAt string `json:"userDeleted"`
	UserCreatedBy string `json:"userCreated"`
	UserUpdatedBy string `json:"userUpdated"`
	UserDeletedBy string `json:"userDeleted"`
}

type FarmInspector struct {
	FarmInspectionId string   `json:"farmInspectionId"`
	FarmInspectionName string `json:"farmInspectionName"`
	CertificateNo     string   `json:"certificateNo"`
	CertificateFrom   string   `json:"certificateFrom"`
	ProductName       string   `json:"productName"`
	TypeOfFertilizer  string   `json:"typeOfFertilizer"`
	FertilizerUsed    string   `json:"fertilizerUsed"`
	Image             []string `json:"image" metadata:",optional"`
}

type Harvester struct {
	HarvestId        string `json:"harvestId"`
	ImporterName     string `json:"importerName"`
	CropSampling     string `json:"cropSampling"`
	TempratureLevel  string `json:"temperatureLevel"`
	Humidity         string `json:"humidityLevel"`
	HarvestStatus    string `json:"harvestStatus"`
	HarvestCreatedAt string `json:"harvestCreated"`
	HarvestUpdatedAt string `json:"harvestUpdated"`
	HarvestDeletedAt string `json:"harvestDeleted"`
}

type Importer struct {
	ImporterId           string `json:"importerId"`
	Quantity             string `json:"quantity"`
	ShipStorage          string `json:"shipStorage"`
	ArrivalDate          string `json:"arrivalDate"`
	WarehouseLocation    string `json:"warehouseLocation"`
	WarehouseArrivalDate string `json:"warehouseArrivalDate"`
	ImporterAddress      string `json:"importerAddress"`
	ImporterStatus       string `json:"importerStatus"`
	ImporterCreatedAt    string `json:"importerCreated"`
	ImporterUpdatedAt    string `json:"importerUpdated"`
	ImporterDeletedAt    string `json:"importerDeleted"`
}

type Exporter struct {
	ExporterId          string `json:"exporterId"`
	CoordinationAddress string `json:"coordinationAddress"`
	ShipName            string `json:"shipName"`
	ShipNo              string `json:"shipNo"`
	DepartureDate       string `json:"departureDate"`
	EstimatedDate       string `json:"estimatedDate"`
	ExportedTo          string `json:"exportedTo"`
	ExporterStatus      string `json:"exporterStatus"`
	ExporterCreatedAt   string `json:"exporterCreated"`
	ExporterUpdatedAt   string `json:"exporterUpdated"`
	ExporterDeletedAt   string `json:"exporterDeleted"`
}

type Processor struct {
	ProcessorId        string   `json:"processorId"`
	Quantity           string   `json:"quantity"`
	ProcessingMethod   string   `json:"processingMethod"`
	Packaging          string   `json:"packaging"`
	PackagedDate       string   `json:"packagedDate"`
	Warehouse          string   `json:"warehouse"`
	WarehouseLocation  string   `json:"warehouseLocation"`
	Destination        string   `json:"destination"`
	ProcessorStatus    string   `json:"processorStatus"`
	ProcessorCreatedAt string   `json:"processorCreated"`
	ProcessorUpdatedAt string   `json:"processorUpdated"`
	ProcessorDeletedAt string   `json:"processorDeleted"`
	Image              []string `json:"image" metadata:",optional"`
}

// Generic CRUD for any struct
func createAsset(ctx contractapi.TransactionContextInterface, key string, asset interface{}) error {
	assetBytes, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal asset: %v", err)
	}
	return ctx.GetStub().PutState(key, assetBytes)
}

func readAsset(ctx contractapi.TransactionContextInterface, key string, asset interface{}) error {
	data, err := ctx.GetStub().GetState(key)
	if err != nil {
		return fmt.Errorf("failed to get asset from world state: %v", err)
	}
	if data == nil {
		return fmt.Errorf("asset not found")
	}
	return json.Unmarshal(data, asset)
}

func updateAsset(ctx contractapi.TransactionContextInterface, key string, asset interface{}) error {
	return createAsset(ctx, key, asset)
}

func deleteAsset(ctx contractapi.TransactionContextInterface, key string) error {
	return ctx.GetStub().DelState(key)
}

// CRUD wrappers for User
func (s *SmartContract) CreateUser(ctx contractapi.TransactionContextInterface, userJson string) error {
	var user User
	if err := json.Unmarshal([]byte(userJson), &user); err != nil {
		return err
	}
	return createAsset(ctx, "USER_"+user.UserID, user)
}

func (s *SmartContract) ReadUser(ctx contractapi.TransactionContextInterface, userID string) (*User, error) {
	var user User
	if err := readAsset(ctx, "USER_"+userID, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *SmartContract) UpdateUser(ctx contractapi.TransactionContextInterface, userJson string) error {
	var user User
	if err := json.Unmarshal([]byte(userJson), &user); err != nil {
		return err
	}
	return updateAsset(ctx, "USER_"+user.UserID, user)
}

func (s *SmartContract) DeleteUser(ctx contractapi.TransactionContextInterface, userID string) error {
	return deleteAsset(ctx, "USER_"+userID)
}

func (s *SmartContract) CreateFarmInspector(ctx contractapi.TransactionContextInterface, inspectorJson string) error {
	var inspector FarmInspector
	if err := json.Unmarshal([]byte(inspectorJson), &inspector); err != nil {
		return err
	}
	return createAsset(ctx, "FARM_INSPECTOR_"+inspector.FarmInspectionId, inspector)
}
func (s *SmartContract) ReadFarmInspector(ctx contractapi.TransactionContextInterface, inspectorID string) (*FarmInspector, error) {
	var inspector FarmInspector
	if err := readAsset(ctx, "FARM_INSPECTOR_"+inspectorID, &inspector); err != nil {
		return nil, err
	}
	return &inspector, nil
}
func (s *SmartContract) UpdateFarmInspector(ctx contractapi.TransactionContextInterface, inspectorJson string) error {
	var inspector FarmInspector
	if err := json.Unmarshal([]byte(inspectorJson), &inspector); err != nil {
		return err
	}
	return updateAsset(ctx, "FARM_INSPECTOR_"+inspector.FarmInspectionId, inspector)
}
func (s *SmartContract) DeleteFarmInspector(ctx contractapi.TransactionContextInterface, inspectorID string) error {
	return deleteAsset(ctx, "FARM_INSPECTOR_"+inspectorID)
}
func (s *SmartContract) CreateHarvester(ctx contractapi.TransactionContextInterface, harvesterJson string) error {
	var harvester Harvester
	if err := json.Unmarshal([]byte(harvesterJson), &harvester); err != nil {
		return err
	}
	return createAsset(ctx, "HARVESTER_"+harvester.HarvestId, harvester)
}
func (s *SmartContract) ReadHarvester(ctx contractapi.TransactionContextInterface, harvesterID string) (*Harvester, error) {
	var harvester Harvester
	if err := readAsset(ctx, "HARVESTER_"+harvesterID, &harvester); err != nil {
		return nil, err
	}
	return &harvester, nil
}
func (s *SmartContract) UpdateHarvester(ctx contractapi.TransactionContextInterface, harvesterJson string) error {
	var harvester Harvester
	if err := json.Unmarshal([]byte(harvesterJson), &harvester); err != nil {
		return err
	}
	return updateAsset(ctx, "HARVESTER_"+harvester.HarvestId, harvester)
}
func (s *SmartContract) DeleteHarvester(ctx contractapi.TransactionContextInterface, harvesterID string) error {
	return deleteAsset(ctx, "HARVESTER_"+harvesterID)
}
func (s *SmartContract) CreateImporter(ctx contractapi.TransactionContextInterface, importerJson string) error {
	var importer Importer
	if err := json.Unmarshal([]byte(importerJson), &importer); err != nil {
		return err
	}
	return createAsset(ctx, "IMPORTER_"+importer.ImporterId, importer)
}
func (s *SmartContract) ReadImporter(ctx contractapi.TransactionContextInterface, importerID string) (*Importer, error) {
	var importer Importer
	if err := readAsset(ctx, "IMPORTER_"+importerID, &importer); err != nil {
		return nil, err
	}
	return &importer, nil
}
func (s *SmartContract) UpdateImporter(ctx contractapi.TransactionContextInterface, importerJson string) error {
	var importer Importer
	if err := json.Unmarshal([]byte(importerJson), &importer); err != nil {
		return err
	}
	return updateAsset(ctx, "IMPORTER_"+importer.ImporterId, importer)
}
func (s *SmartContract) DeleteImporter(ctx contractapi.TransactionContextInterface, importerID string) error {
	return deleteAsset(ctx, "IMPORTER_"+importerID)
}
func (s *SmartContract) CreateExporter(ctx contractapi.TransactionContextInterface, exporterJson string) error {
	var exporter Exporter
	if err := json.Unmarshal([]byte(exporterJson), &exporter); err != nil {
		return err
	}
	return createAsset(ctx, "EXPORTER_"+exporter.ExporterId, exporter)
}
func (s *SmartContract) ReadExporter(ctx contractapi.TransactionContextInterface, exporterID string) (*Exporter, error) {
	var exporter Exporter
	if err := readAsset(ctx, "EXPORTER_"+exporterID, &exporter); err != nil {
		return nil, err
	}
	return &exporter, nil
}
func (s *SmartContract) UpdateExporter(ctx contractapi.TransactionContextInterface, exporterJson string) error {
	var exporter Exporter
	if err := json.Unmarshal([]byte(exporterJson), &exporter); err != nil {
		return err
	}
	return updateAsset(ctx, "EXPORTER_"+exporter.ExporterId, exporter)
}
func (s *SmartContract) DeleteExporter(ctx contractapi.TransactionContextInterface, exporterID string) error {
	return deleteAsset(ctx, "EXPORTER_"+exporterID)
}
func (s *SmartContract) CreateProcessor(ctx contractapi.TransactionContextInterface, processorJson string) error {
	var processor Processor
	if err := json.Unmarshal([]byte(processorJson), &processor); err != nil {
		return err
	}
	return createAsset(ctx, "PROCESSOR_"+processor.ProcessorId, processor)
}
func (s *SmartContract) ReadProcessor(ctx contractapi.TransactionContextInterface, processorID string) (*Processor, error) {
	var processor Processor
	if err := readAsset(ctx, "PROCESSOR_"+processorID, &processor); err != nil {
		return nil, err
	}
	return &processor, nil
}
func (s *SmartContract) UpdateProcessor(ctx contractapi.TransactionContextInterface, processorJson string) error {
	var processor Processor
	if err := json.Unmarshal([]byte(processorJson), &processor); err != nil {
		return err
	}
	return updateAsset(ctx, "PROCESSOR_"+processor.ProcessorId, processor)
}
func (s *SmartContract) DeleteProcessor(ctx contractapi.TransactionContextInterface, processorID string) error {
	return deleteAsset(ctx, "PROCESSOR_"+processorID)
}
func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Error create chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting chaincode: %s", err.Error())
	}
}

