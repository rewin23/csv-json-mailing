package main
 
import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "os"
    "bytes"
    "io/ioutil"
    "net/http"
  
    _"github.com/rewin23/csv-json-mailing/models"
)

type Contestant struct {
        FirstName    string  `json:"first_name"`
        LastName     string  `json:"last_name"`
        Genre        string  `json:"genre"`
        Phone        string  `json:"phone"` 
        OptIn        bool    `json:"optin"`
        Email        string  `json:"email"`
        Coupon       string  `json:"coupon"`
    }


 

func main() {
    csvFile, err := os.Open("./data.csv")
    if err != nil {
        fmt.Println(err)
    }
    defer csvFile.Close()
 
    reader := csv.NewReader(csvFile)
    reader.FieldsPerRecord = -1
 
    csvData, err := reader.ReadAll()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
 
    //var regex_ruta = regexp.MustCompile("^/([A-Z])$")
    var contestant Contestant
    var contestants []Contestant

    progress := 1

    url := "https://utils.lobulo.dev/mailing/nivea/covid"
 
    for _, each := range csvData {
        contestant.FirstName = each[2]
        contestant.LastName  = each[3]
        contestant.Genre     = "notspecified"
        contestant.Phone     = each[5]
        contestant.OptIn     = true
        contestant.Email     = each[6]
        contestant.Coupon    = each[7]
       
        contestants = append(contestants, contestant)

        jsonData, err := json.Marshal(contestant)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        //fmt.Println(string(jsonData))

        resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            print(err)
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            print("error")
            print(err)
        }
        progress = progress + 1
        fmt.Printf("%v ", progress)
        fmt.Println(string(body))

    }
 
    // Convert to JSON
    jsonData, err := json.Marshal(contestants)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
 
    //fmt.Println(string(jsonData))
 
    jsonFile, err := os.Create("./data.json")
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()
 
    jsonFile.Write(jsonData)
    jsonFile.Close()
}


