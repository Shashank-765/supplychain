package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Structs for each entity
type User struct {
	UserId        string `json:"userId"`
	UserType      string `json:"userType"`
	UserRole      string `json:"userRole"`
	UserName      string `json:"userName"`
	UserEmail     string `json:"userEmail"`
	UserPhone     string `json:"userPhone"`
	UserAddress   string `json:"userAddress"`
	UserPassword   string `json:"userPassword"`
	UserStatus    string `json:"userStatus"`
	UserCreatedAt string `json:"userCreatedAt"`
	UserUpdatedAt string `json:"userUpdatedAt"`
	UserDeletedAt string `json:"userDeletedAt"`
	UserCreatedBy string `json:"userCreatedBy"`
	UserUpdatedBy string `json:"userUpdatedBy"`
	UserDeletedBy string `json:"userDeletedBy"`
}

type FarmInspector struct {
	FarmInspectionId   string   `json:"farmInspectionId"`
	FarmInspectionName string   `json:"farmInspectionName"`
	CertificateNo      string   `json:"certificateNo"`
	CertificateFrom    string   `json:"certificateFrom"`
	ProductName        string   `json:"productName"`
	TypeOfFertilizer   string   `json:"typeOfFertilizer"`
	FertilizerUsed     string   `json:"fertilizerUsed"`
	Image              []string `json:"image" metadata:",optional"`
	FarmInspectionStatus string   `json:"farmInspectionStatus"`
	FarmInspectionCreatedAt string   `json:"farmInspectionCreatedAt"`
	FarmInspectionUpdatedAt string   `json:"farmInspectionUpdatedAt"`
	FarmInspectionDeletedAt string   `json:"farmInspectionDeletedAt"`
	BatchId            string   `json:"batchId"` // Link to Batch
}

type Harvester struct {
	HarvestId        string `json:"harvestId"`
	HarvesterName     string `json:"harvesterName"`
	CropSampling     string `json:"cropSampling"`
	TempratureLevel  string `json:"temperatureLevel"`
	Humidity         string `json:"humidityLevel"`
	HarvestStatus    string `json:"harvestStatus"`
	HarvestCreatedAt string `json:"harvestCreatedAt"`
	HarvestUpdatedAt string `json:"harvestUpdatedAt"`
	HarvestDeletedAt string `json:"harvestDeletedAt"`
	BatchId          string `json:"batchId"` // Link to Batch
}

type Importer struct {
	ImporterId           string `json:"importerId"`
	ImporterName     string `json:"importerName"`
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
	BatchId              string `json:"batchId"` // Link to Batch
}

type Exporter struct {
	ExporterId          string `json:"exporterId"`
	CoordinationAddress string `json:"coordinationAddress"`
	ExporterName        string `json:"exporterName"`
	ShipName            string `json:"shipName"`
	ShipNo              string `json:"shipNo"`
	DepartureDate       string `json:"departureDate"`
	EstimatedDate       string `json:"estimatedDate"`
	ExportedTo          string `json:"exportedTo"`
	ExporterStatus      string `json:"exporterStatus"`
	ExporterCreatedAt   string `json:"exporterCreated"`
	ExporterUpdatedAt   string `json:"exporterUpdated"`
	ExporterDeletedAt   string `json:"exporterDeleted"`
	BatchId             string `json:"batchId"` // Link to Batch
}

type Processor struct {
	ProcessorId        string   `json:"processorId"`
	Quantity           string   `json:"quantity"`
	ProcessingMethod   string   `json:"processingMethod"`
	ProcessorName     string   `json:"processorName"`
	Price 		       string   `json:"price"`
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
	BatchId            string   `json:"batchId"` // Link to Batch
}

type Batch struct {
	BatchId             string `json:"batchId"`
	FarmerRegNo         string `json:"farmerRegNo"`
	FarmerName          string `json:"farmerName"`
	FarmerAddress       string `json:"farmerAddress"`
	FarmInspectionName  string `json:"farmInspectionName"`
	HarvesterName       string `json:"harvesterName"`
	ProcessorName       string `json:"processorName"`
	ExporterName        string `json:"exporterName"`
	ImporterName        string `json:"importerName"`
	CoffeeType          string `json:"coffeeType"`
	QRCode              string `json:"qrCode"`
	FarmInspectionId    string `json:"farmInspectionId"`
	HarvesterId         string `json:"harvesterId"`
	ProcessorId         string `json:"processorId"`
	ExporterId          string `json:"exporterId"`
	ImporterId          string `json:"importerId"`
}

// Smart Contract Structure
type SmartContract struct {
	contractapi.Contract
}


func (s *SmartContract) CreateUser(ctx contractapi.TransactionContextInterface, user User) error {
	// Check if user already exists
	userJSON, err := ctx.GetStub().GetState(user.UserId)
	if err != nil {
		return fmt.Errorf("Failed to check if user exists: %v", err)
	}
	if userJSON != nil {
		return fmt.Errorf("Batch with ID %s already exists", user.UserId)
	}

	// Add user to the ledger
	userJSON, err = json.Marshal(user)
	if err != nil {
		return fmt.Errorf("Failed to marshal user: %v", err)
	}

	err = ctx.GetStub().PutState(user.UserId, userJSON)
	if err != nil {
		return fmt.Errorf("Failed to create user: %v", err)
	}

	return nil
}

// Create Functions for Batch, FarmInspector, Harvester, Importer, Exporter, and Processor

func (s *SmartContract) CreateBatch(ctx contractapi.TransactionContextInterface, batch Batch) error {
	// Check if batch already exists
	batchJSON, err := ctx.GetStub().GetState(batch.BatchId)
	if err != nil {
		return fmt.Errorf("Failed to check if batch exists: %v", err)
	}
	if batchJSON != nil {
		return fmt.Errorf("Batch with ID %s already exists", batch.BatchId)
	}

	// Add batch to the ledger
	batchJSON, err = json.Marshal(batch)
	if err != nil {
		return fmt.Errorf("Failed to marshal batch: %v", err)
	}

	err = ctx.GetStub().PutState(batch.BatchId, batchJSON)
	if err != nil {
		return fmt.Errorf("Failed to create batch: %v", err)
	}

	return nil
}

func (s *SmartContract) CreateFarmInspector(ctx contractapi.TransactionContextInterface, farmInspector FarmInspector) error {
	// Check if farm inspector already exists
	farmInspectorJSON, err := ctx.GetStub().GetState(farmInspector.FarmInspectionId)
	if err != nil {
		return fmt.Errorf("Failed to check if farm inspector exists: %v", err)
	}
	if farmInspectorJSON != nil {
		return fmt.Errorf("Farm inspector with ID %s already exists", farmInspector.FarmInspectionId)
	}

	// Link to BatchId
	// farmInspector.BatchId = farmInspector.FarmInspectionId

	// Add farm inspector to the ledger
	farmInspectorJSON, err = json.Marshal(farmInspector)
	if err != nil {
		return fmt.Errorf("Failed to marshal farm inspector: %v", err)
	}

	err = ctx.GetStub().PutState(farmInspector.FarmInspectionId, farmInspectorJSON)
	if err != nil {
		return fmt.Errorf("Failed to create farm inspector: %v", err)
	}

	return nil
}

func (s *SmartContract) CreateHarvester(ctx contractapi.TransactionContextInterface, harvester Harvester) error {
	// Check if harvester already exists
	harvesterJSON, err := ctx.GetStub().GetState(harvester.HarvestId)
	if err != nil {
		return fmt.Errorf("Failed to check if harvester exists: %v", err)
	}
	if harvesterJSON != nil {
		return fmt.Errorf("Harvester with ID %s already exists", harvester.HarvestId)
	}

	// Link to BatchId
	// harvester.BatchId = harvester.HarvestId

	// Add harvester to the ledger
	harvesterJSON, err = json.Marshal(harvester)
	if err != nil {
		return fmt.Errorf("Failed to marshal harvester: %v", err)
	}

	err = ctx.GetStub().PutState(harvester.HarvestId, harvesterJSON)
	if err != nil {
		return fmt.Errorf("Failed to create harvester: %v", err)
	}

	return nil
}

func (s *SmartContract) CreateImporter(ctx contractapi.TransactionContextInterface, importer Importer) error {
	// Check if importer already exists
	importerJSON, err := ctx.GetStub().GetState(importer.ImporterId)
	if err != nil {
		return fmt.Errorf("Failed to check if importer exists: %v", err)
	}
	if importerJSON != nil {
		return fmt.Errorf("Importer with ID %s already exists", importer.ImporterId)
	}

	// Link to BatchId
	// importer.BatchId = importer.ImporterId

	// Add importer to the ledger
	importerJSON, err = json.Marshal(importer)
	if err != nil {
		return fmt.Errorf("Failed to marshal importer: %v", err)
	}

	err = ctx.GetStub().PutState(importer.ImporterId, importerJSON)
	if err != nil {
		return fmt.Errorf("Failed to create importer: %v", err)
	}

	return nil
}

func (s *SmartContract) CreateExporter(ctx contractapi.TransactionContextInterface, exporter Exporter) error {
	// Check if exporter already exists
	exporterJSON, err := ctx.GetStub().GetState(exporter.ExporterId)
	if err != nil {
		return fmt.Errorf("Failed to check if exporter exists: %v", err)
	}
	if exporterJSON != nil {
		return fmt.Errorf("Exporter with ID %s already exists", exporter.ExporterId)
	}

	// Link to BatchId
	// exporter.BatchId = exporter.ExporterId

	// Add exporter to the ledger
	exporterJSON, err = json.Marshal(exporter)
	if err != nil {
		return fmt.Errorf("Failed to marshal exporter: %v", err)
	}

	err = ctx.GetStub().PutState(exporter.ExporterId, exporterJSON)
	if err != nil {
		return fmt.Errorf("Failed to create exporter: %v", err)
	}

	return nil
}

func (s *SmartContract) CreateProcessor(ctx contractapi.TransactionContextInterface, processor Processor) error {
	// Check if processor already exists
	processorJSON, err := ctx.GetStub().GetState(processor.ProcessorId)
	if err != nil {
		return fmt.Errorf("Failed to check if processor exists: %v", err)
	}
	if processorJSON != nil {
		return fmt.Errorf("Processor with ID %s already exists", processor.ProcessorId)
	}

	// Link to BatchId
	// processor.BatchId = processor.ProcessorId

	// Add processor to the ledger
	processorJSON, err = json.Marshal(processor)
	if err != nil {
		return fmt.Errorf("Failed to marshal processor: %v", err)
	}

	err = ctx.GetStub().PutState(processor.ProcessorId, processorJSON)
	if err != nil {
		return fmt.Errorf("Failed to create processor: %v", err)
	}

	return nil
}

// Update Functions

func (s *SmartContract) UpdateUser(ctx contractapi.TransactionContextInterface, user User) error {
	// Check if user exists
	userJSON, err := ctx.GetStub().GetState(user.UserId)
	if err != nil || userJSON == nil {
		return fmt.Errorf("User with ID %s does not exist", user.UserId)
	}

	// Update user information
	updatedUserJSON, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("Failed to marshal updated user: %v", err)
	}

	err = ctx.GetStub().PutState(user.UserId, updatedUserJSON)
	if err != nil {
		return fmt.Errorf("Failed to update user: %v", err)
	}

	return nil
}


func (s *SmartContract) UpdateBatch(ctx contractapi.TransactionContextInterface, batch Batch) error {
	// Check if batch exists
	batchJSON, err := ctx.GetStub().GetState(batch.BatchId)
	if err != nil || batchJSON == nil {
		return fmt.Errorf("Batch with ID %s does not exist", batch.BatchId)
	}

	// Update batch information
	updatedBatchJSON, err := json.Marshal(batch)
	if err != nil {
		return fmt.Errorf("Failed to marshal updated batch: %v", err)
	}

	err = ctx.GetStub().PutState(batch.BatchId, updatedBatchJSON)
	if err != nil {
		return fmt.Errorf("Failed to update batch: %v", err)
	}

	return nil
}



// UpdateFarmInspector updates farm inspector information
func (s *SmartContract) UpdateFarmInspector(ctx contractapi.TransactionContextInterface, farmInspector FarmInspector) error {
	// Check if farm inspector exists
	farmInspectorJSON, err := ctx.GetStub().GetState(farmInspector.FarmInspectionId)
	if err != nil || farmInspectorJSON == nil {
		return fmt.Errorf("Farm Inspector with ID %s does not exist", farmInspector.FarmInspectionId)
	}

	// Update farm inspector
	updatedFarmInspectorJSON, err := json.Marshal(farmInspector)
	if err != nil {
		return fmt.Errorf("Failed to marshal updated farm inspector: %v", err)
	}

	err = ctx.GetStub().PutState(farmInspector.FarmInspectionId, updatedFarmInspectorJSON)
	if err != nil {
		return fmt.Errorf("Failed to update farm inspector: %v", err)
	}

	return nil
}

// UpdateHarvester updates harvester information
func (s *SmartContract) UpdateHarvester(ctx contractapi.TransactionContextInterface, harvester Harvester) error {
	// Check if harvester exists
	harvesterJSON, err := ctx.GetStub().GetState(harvester.HarvestId)
	if err != nil || harvesterJSON == nil {
		return fmt.Errorf("Harvester with ID %s does not exist", harvester.HarvestId)
	}

	// Update harvester
	updatedHarvesterJSON, err := json.Marshal(harvester)
	if err != nil {
		return fmt.Errorf("Failed to marshal updated harvester: %v", err)
	}

	err = ctx.GetStub().PutState(harvester.HarvestId, updatedHarvesterJSON)
	if err != nil {
		return fmt.Errorf("Failed to update harvester: %v", err)
	}

	return nil
}

// UpdateImporter updates importer information
func (s *SmartContract) UpdateImporter(ctx contractapi.TransactionContextInterface, importer Importer) error {
	// Check if importer exists
	importerJSON, err := ctx.GetStub().GetState(importer.ImporterId)
	if err != nil || importerJSON == nil {
		return fmt.Errorf("Importer with ID %s does not exist", importer.ImporterId)
	}

	// Update importer
	updatedImporterJSON, err := json.Marshal(importer)
	if err != nil {
		return fmt.Errorf("Failed to marshal updated importer: %v", err)
	}

	err = ctx.GetStub().PutState(importer.ImporterId, updatedImporterJSON)
	if err != nil {
		return fmt.Errorf("Failed to update importer: %v", err)
	}

	return nil
}

// UpdateExporter updates exporter information
func (s *SmartContract) UpdateExporter(ctx contractapi.TransactionContextInterface, exporter Exporter) error {
	// Check if exporter exists
	exporterJSON, err := ctx.GetStub().GetState(exporter.ExporterId)
	if err != nil || exporterJSON == nil {
		return fmt.Errorf("Exporter with ID %s does not exist", exporter.ExporterId)
	}

	// Update exporter
	updatedExporterJSON, err := json.Marshal(exporter)
	if err != nil {
		return fmt.Errorf("Failed to marshal updated exporter: %v", err)
	}

	err = ctx.GetStub().PutState(exporter.ExporterId, updatedExporterJSON)
	if err != nil {
		return fmt.Errorf("Failed to update exporter: %v", err)
	}

	return nil
}

// UpdateProcessor updates processor information
func (s *SmartContract) UpdateProcessor(ctx contractapi.TransactionContextInterface, processor Processor) error {
	// Check if processor exists
	processorJSON, err := ctx.GetStub().GetState(processor.ProcessorId)
	if err != nil || processorJSON == nil {
		return fmt.Errorf("Processor with ID %s does not exist", processor.ProcessorId)
	}

	// Update processor
	updatedProcessorJSON, err := json.Marshal(processor)
	if err != nil {
		return fmt.Errorf("Failed to marshal updated processor: %v", err)
	}

	err = ctx.GetStub().PutState(processor.ProcessorId, updatedProcessorJSON)
	if err != nil {
		return fmt.Errorf("Failed to update processor: %v", err)
	}

	return nil
}


// ViewBatch retrieves the batch details by batchId
func (s *SmartContract) ViewUser(ctx contractapi.TransactionContextInterface, userId string) (User, error) {
	// Fetch batch details
	userJSON, err := ctx.GetStub().GetState(userId)
	if err != nil || userJSON == nil {
		return User{}, fmt.Errorf("User with ID %s does not exist", userId)
	}

	// Unmarshal the data
	var batch User
	err = json.Unmarshal(userJSON, &batch)
	if err != nil {
		return User{}, fmt.Errorf("Failed to unmarshal batch data: %v", err)
	}

	return batch, nil
}



// ViewBatch retrieves the batch details by batchId
func (s *SmartContract) ViewBatch(ctx contractapi.TransactionContextInterface, batchId string) (Batch, error) {
	// Fetch batch details
	batchJSON, err := ctx.GetStub().GetState(batchId)
	if err != nil || batchJSON == nil {
		return Batch{}, fmt.Errorf("Batch with ID %s does not exist", batchId)
	}

	// Unmarshal the data
	var batch Batch
	err = json.Unmarshal(batchJSON, &batch)
	if err != nil {
		return Batch{}, fmt.Errorf("Failed to unmarshal batch data: %v", err)
	}

	return batch, nil
}

// ViewFarmInspector retrieves the farm inspector details by farmInspectionId
func (s *SmartContract) ViewFarmInspector(ctx contractapi.TransactionContextInterface, farmInspectionId string) (FarmInspector, error) {
	// Fetch farm inspector details
	farmInspectorJSON, err := ctx.GetStub().GetState(farmInspectionId)
	if err != nil || farmInspectorJSON == nil {
		return FarmInspector{}, fmt.Errorf("Farm inspector with ID %s does not exist", farmInspectionId)
	}

	// Unmarshal the data
	var farmInspector FarmInspector
	err = json.Unmarshal(farmInspectorJSON, &farmInspector)
	if err != nil {
		return FarmInspector{}, fmt.Errorf("Failed to unmarshal farm inspector data: %v", err)
	}

	return farmInspector, nil
}

// ViewHarvester retrieves the harvester details by harvestId
func (s *SmartContract) ViewHarvester(ctx contractapi.TransactionContextInterface, harvestId string) (Harvester, error) {
	// Fetch harvester details
	harvesterJSON, err := ctx.GetStub().GetState(harvestId)
	if err != nil || harvesterJSON == nil {
		return Harvester{}, fmt.Errorf("Harvester with ID %s does not exist", harvestId)
	}

	// Unmarshal the data
	var harvester Harvester
	err = json.Unmarshal(harvesterJSON, &harvester)
	if err != nil {
		return Harvester{}, fmt.Errorf("Failed to unmarshal harvester data: %v", err)
	}

	return harvester, nil
}

// ViewImporter retrieves the importer details by importerId
func (s *SmartContract) ViewImporter(ctx contractapi.TransactionContextInterface, importerId string) (Importer, error) {
	// Fetch importer details
	importerJSON, err := ctx.GetStub().GetState(importerId)
	if err != nil || importerJSON == nil {
		return Importer{}, fmt.Errorf("Importer with ID %s does not exist", importerId)
	}

	// Unmarshal the data
	var importer Importer
	err = json.Unmarshal(importerJSON, &importer)
	if err != nil {
		return Importer{}, fmt.Errorf("Failed to unmarshal importer data: %v", err)
	}

	return importer, nil
}

// ViewExporter retrieves the exporter details by exporterId
func (s *SmartContract) ViewExporter(ctx contractapi.TransactionContextInterface, exporterId string) (Exporter, error) {
	// Fetch exporter details
	exporterJSON, err := ctx.GetStub().GetState(exporterId)
	if err != nil || exporterJSON == nil {
		return Exporter{}, fmt.Errorf("Exporter with ID %s does not exist", exporterId)
	}

	// Unmarshal the data
	var exporter Exporter
	err = json.Unmarshal(exporterJSON, &exporter)
	if err != nil {
		return Exporter{}, fmt.Errorf("Failed to unmarshal exporter data: %v", err)
	}

	return exporter, nil
}

// ViewProcessor retrieves the processor details by processorId
func (s *SmartContract) ViewProcessor(ctx contractapi.TransactionContextInterface, processorId string) (Processor, error) {
	// Fetch processor details
	processorJSON, err := ctx.GetStub().GetState(processorId)
	if err != nil || processorJSON == nil {
		return Processor{}, fmt.Errorf("Processor with ID %s does not exist", processorId)
	}

	// Unmarshal the data
	var processor Processor
	err = json.Unmarshal(processorJSON, &processor)
	if err != nil {
		return Processor{}, fmt.Errorf("Failed to unmarshal processor data: %v", err)
	}

	return processor, nil
}



// GetAllBatches retrieves all batch records from the ledger
func (s *SmartContract) GetAllBatches(ctx contractapi.TransactionContextInterface) ([]Batch, error) {
	queryIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("Failed to get state by range: %v", err)
	}
	defer queryIterator.Close()

	var batches []Batch

	for queryIterator.HasNext() {
		queryResponse, err := queryIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("Failed to iterate over results: %v", err)
		}

		var batch Batch
		err = json.Unmarshal(queryResponse.Value, &batch)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal batch data: %v", err)
		}

		batches = append(batches, batch)
	}

	return batches, nil
}


// GetAllUseres retrieves all batch records from the ledger
func (s *SmartContract) GetAllUsers(ctx contractapi.TransactionContextInterface) ([]User, error) {
	queryIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("Failed to get state by range: %v", err)
	}
	defer queryIterator.Close()

	var users []User

	for queryIterator.HasNext() {
		queryResponse, err := queryIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("Failed to iterate over results: %v", err)
		}

		var batch User
		err = json.Unmarshal(queryResponse.Value, &batch)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal batch data: %v", err)
		}

		users = append(users, batch)
	}

	return users, nil
}


