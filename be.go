package main                                                       
                                                                   
import (                                                           
        "encoding/json"                                            
        "fmt"                                                      
        "strconv"                                                  
                                                                   
        "github.com/hyperledger/fabric-contract-api-go/contractapi"
)                                                              
                                                               
                                                               
// SmartContract provides functions for managing a car         
type SmartContract struct {                                    
        contractapi.Contract                                   
}                                                              
// SimpleAsset implements a simple chaincode to manage an asset
type Product struct {                                              
        Brand   string `json:"brand"`                              
        Price  int `json:"price"`                                  
        Count int `json:"count"`                                   
}                                                                  
// QueryResult structure used for handling result of query         
type QueryResult struct {                                      
        Key    string `json:"Key"`                             
        Record *Product                                        
}                                                              
                                                               
// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface)
        // Set up any variables or assets here by calling stub.PutState()      
        products := []Product{                                                 
                Product{Brand: "Samsung TV", Price: 250, Count: 20},           
                Product{Brand: "Apple TV", Price: 250, Count: 30},             
                Product{Brand: "Xiaomi Mi Phone", Price: 150, Count: 50},      
                Product{Brand: "Toshiba Laptop", Price: 200, Count: 40},       
                Product{Brand: "Huawei Watch", Price: 150, Count: 60},         
        }                                                                      
                                                                               
        for i, product := range products {                                     
                productAsBytes, _ := json.Marshal(product)                     
                err := ctx.GetStub().PutState("PRODUCT"+strconv.Itoa(i), product
                                                                                
                if err != nil {                                                 
                        return fmt.Errorf("Failed to put to world state. %s", er
                }                                                               
        }                                                                       
                                                                                
        return nil                                                              
}
// QueryAllProducts gets all products from the world state                      
func (s *SmartContract) QueryAllProducts(ctx contractapi.TransactionContextInterface)
        startKey := "PRODUCT0"                                                  
        endKey := "PRODUCT99"                                                   
                                                                                
        resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey) 
        if err != nil {                                                         
                return nil, err                                                 
        }                                                                       
                                                                                
        defer resultsIterator.Close()                                           
        results := []QueryResult{}                                              
        for resultsIterator.HasNext() {                                         
                queryResponse, err := resultsIterator.Next()                    
                if err != nil {                                                 
                        return nil, err                                         
                }                                                               
                                                                                
                product := new(Product)                                         
                _ = json.Unmarshal(queryResponse.Value, product)                
                                                                                
                queryResult := QueryResult{Key: queryResponse.Key, Record: product}
                results = append(results, queryResult)
        }                                                                       
                                                                                
        return results, nil                                                     
}                                                                               
// CreateProduct adds a new product to the world state with given details       
func (s *SmartContract) CreateProduct(ctx contractapi.TransactionContextInterfac
        product := Product{                                                     
                Brand:   brand,                                                 
                Price:  price,                                                  
                Count: count,                                                   
        }                                                                       
                                                                                
        productAsBytes, _ := json.Marshal(product)                              
        return ctx.GetStub().PutState(productNumber, productAsBytes)            
}
// ChangeProductPrice updates the price of the new product with given details   
func (s *SmartContract) ChangeProductPrice(ctx contractapi.TransactionContextInt
        productAsBytes, err := ctx.GetStub().GetState(productNumber)            
        product := new(Product)                                                 
        _ = json.Unmarshal(productAsBytes, product)                             
                                                                                
        if err != nil {                                                         
                return err                                                      
        }                                                                       
                                                                                
        product.Price = newPrice                                                
        productAsBytesToPut, _ := json.Marshal(product)                         
        return ctx.GetStub().PutState(productNumber,                            
productAsBytesToPut)                                                            
}
// main function starts up the chaincode in the container during instantiate    
func main() {                                                                   
        chaincode, err := contractapi.NewChaincode(new(SmartContract))          
                                                                                
        if err != nil {                                                         
                fmt.Printf("Error create be chaincode: %s", err.Error())        
                return                                                          
        }                                                                       
                                                                                
        if err := chaincode.Start(); err != nil {                               
                fmt.Printf("Error starting be chaincode: %s", err.Error())      
        }                                                                       
}
