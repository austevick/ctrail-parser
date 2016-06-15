package main
import (
    "flag"
    "os"
    "github.com/austevick/ctrail-parser/parser"
)

func main(){
    currentDirectory, _ := os.Getwd()
    // Directory to scan logs for. Looks for *.json.gz files. Defaults to current directory
    directory := flag.String("logs", currentDirectory, "Directory for logs")
    // JMESPath style query
    queryExpression := flag.String("query", "Records[*]", "JMESPath query expression. Defaults to all records")
    flag.Parse()
    parser.Execute(*directory,*queryExpression)
}
