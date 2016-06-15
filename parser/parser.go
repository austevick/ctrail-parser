package parser

import (
    "fmt"
    "path"
    "path/filepath"
    "compress/gzip"
    "io/ioutil"
    "os"
    "log"
    "encoding/json"
    "github.com/jmespath/go-jmespath"
)
func Execute(dir,query string){
    files, _ := filepath.Glob(path.Join(dir,"*.json.gz")) //Get all *.json.gz files in directory
    if len(files) == 0 { //If it doesn't find any files, check to see if there are directories and recursivly loop through directories
        err := filepath.Walk(dir, func(fullpath string, f os.FileInfo, err error) error {
            if f.IsDir() { //If it is a directory, get al *.json.gz files
                filesInCurrentDir,_ := filepath.Glob(path.Join(fullpath,"*.json.gz"))
                files = append(files, filesInCurrentDir...) //Append to list of files
            }
            return nil
        })
        if err != nil {
            log.Fatal("Bad Walk ",err)
        }
    }
    if len(files) == 0 {
        log.Fatal("No *.json.gz files found in ",dir)
    }
    events := make(map[string]interface{}) //empty map of interfaces to hold all records
    for _,file := range files { //Loop over all the *.json.gz files in the directory
        cloudtrailResponses := extract(file) //extract the file from the gzip archive
        var data map[string]interface{}
        json.Unmarshal(cloudtrailResponses, &data) //Unmarshal json into empty map of interfaces
        if events["Records"] == nil {
            events["Records"] = data["Records"].([]interface{}) //Create "Records" key in map if not already present
        } else {
            //Append to events["Records"] if events["Records"] is already present
            events["Records"] = append(events["Records"].([]interface{}),data["Records"].([]interface{})...)
        }
    }
    result, err := jmespath.Search(query, events) //Apply the search query to the list of CloudTrail events
	if err != nil {
		log.Fatal("Error executing expression: ", err)
	}
	toJSON, err := json.MarshalIndent(result, "", "  ") //Convert our map back into json but pretty-print the json (i.e. indent it...)
	if err != nil {
		log.Fatal("Error serializing result to JSON: ", err)
	}
	fmt.Println(string(toJSON)) //Print the filtered json events
}

func extract(f string) []byte{
    file, err := os.Open(f)//OPen the gz file
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close() //Close the gz file once this function exits
    gzipReader, err := gzip.NewReader(file) //Create a new gzip reader using the opened gzip file
    if err != nil {
        log.Fatal(err)
    }
    defer gzipReader.Close() //close the gzip reader after this function closes
    cloudtrailResponses, err := ioutil.ReadAll(gzipReader) //Read all data from the gzip reader and put it into cloudtrailResponses as a list of bytes
    if err != nil {
            log.Fatal("ReadAll: %v", err)
        }
    return cloudtrailResponses
}
