package main

import (
	"flag"
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
    "encoding/csv"
	"gopkg.in/yaml.v2"
)

var (
	conf    Config
	cfgFlag = flag.String("config", "config.yaml", "path of the config file")
	outputFlag = flag.String("output", "", "csv file to export, non if left blank")
	resultParams = [][]string {
		{"title"},
		{"time_taken", "Time taken for tests:"},
		{"completed", "Complete requests:"},
		{"failed", "Failed requests:"},
		{"requests_per_second", "Requests per second:"},
	}
)

// Config contains the settings of the ST server getting testes and all the tests
type Config struct {
	Host        string `yaml:"host"`
	Version     string `yaml:"version"`
	Requests    int64  `yaml:"requests"`
	Concurrency int64  `yaml:"concurrency"`
	APIEndpoint string
	Tests       []Test `yaml:"tests"`
}

// Test is a test loaded from config/yaml and getting executed by AB
type Test struct {
	Name        string `yaml:"name"`
	Requests    int64  `yaml:"requests"`
	Concurrency int64  `yaml:"concurrency"`
	Endpoint    string `yaml:"endpoint"`
	RequestType string `yaml:"type"`
	File        string `yaml:"file"`
}

func main() {
	log.Print("Starting AB tests")

	flag.Parse()
	cfg := *cfgFlag
	output := *outputFlag

	var err error
	conf, err = getConfig(cfg)
	if err != nil {
		log.Fatal("config read error: ", err)
		return
	}

	results := runABTest()

	if len(output) > 0 {
		writeCSV(output, results)
	}

	log.Print("Done!")
}

func runABTest() [][]string {
	results := [][]string{}
	for i := 0; i < len(conf.Tests); i++ {
		//log.Print("Starting test " + conf.Tests[i].Name)
		
		testCommand := testToAb(conf.APIEndpoint, conf.Tests[i]);
		cmd := exec.Command("ab", testCommand...)
		stdout, err := cmd.Output()

		if err != nil {
			println(err.Error())
			return nil
		}

		result := getResultFromOutput(stdout)
		result = append([]string{conf.Tests[i].Name}, result...)

		results = append(results, result)
		printResults(conf.Tests[i].Name, result)
	}

	return results
}

func testToAb(apiEndpoint string, test Test) []string {	
	params := []string{
		"-n", fmt.Sprintf("%v", test.Requests),
		"-k",
		"-c", fmt.Sprintf("%v", test.Concurrency),
	}

	if test.RequestType == "POST" || test.RequestType == "PUT"  {		
		method := ""
		if test.RequestType == "POST"{
			method = "-p"
		} else if test.RequestType == "PUT"{
			method = "-u"
		}

		params = append(params, method)
		params = append(params, test.File)
		params = append(params, "-T")
		params = append(params, "content-type-:application/json")
	}

	endpoint := strings.Join([]string{apiEndpoint, strings.TrimPrefix(test.Endpoint, "/")}, "/")
	params = append(params, endpoint)

	return params
}

func getResultFromOutput(output []byte) []string {
	split := strings.Split(string(output), "\n")
	results := []string{}

	for i := 0; i < len(split); i++ {
		for j := 1; j < len(resultParams); j++ {
			if strings.Contains(split[i], resultParams[j][1]) {
				sub := strings.Trim(split[i], resultParams[j][1])
				results = append(results, strings.Split(sub, " ")[0])
			}
		}
	}

	return results
}

func printResults(test string, results []string){
	log.Printf("----------------------------------\n")
	indent := 25
	for i := 0; i < len(results); i++ {
		log.Printf("%s%s%s", resultParams[i][0], createIndent(indent - len(resultParams[i][0])), results[i])
	}
}

func createIndent(length int) string {
	indent := ""
	for i := 0; i < length; i++ {
		indent = fmt.Sprintf("%s ", indent)
	}

	return indent
}

func writeCSV(path string, results [][]string){
	file, err := os.Create(path)
	if err != nil {
        log.Fatal("Cannot create file", err)
    }

    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

	//add header to results
	header := []string{}
	for i := 0; i < len(resultParams); i++{
		header = append(header, resultParams[i][0])
	}
	results = append([][]string {header}, results...)

	// write
    for _, value := range results {
        err := writer.Write(value)        
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
    }

	log.Printf("CSV output written to %s", path)
}

func getConfig(cfgFile string) (Config, error) {
	config := Config{}

	content, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return config, err
	}

	config.Host = strings.TrimSuffix(config.Host, "/")
	config.APIEndpoint = strings.Join([]string{config.Host, config.Version}, "/")
	checkConfig(config)
	return config, nil
}

func checkConfig(config Config) {
	for i := 0; i < len(config.Tests); i++ {
		test := &config.Tests[i]		

		// fatal on error
		if(test.Name == ""){
			log.Fatalf("test at index %v missing name", i);
		}
		if(test.RequestType != "GET" && test.RequestType != "POST" && test.RequestType != "PUT"){
			log.Fatalf("%s test type '%s' is not support, supported types are GET, POST or PUT", test.Name, test.RequestType);
		}
		if((test.RequestType == "POST" || test.RequestType == "PUT") && len(test.File) == 0){
			log.Fatalf("%s test type is '%s' but no file given", test.Name, test.RequestType);
		}
		if(test.File != ""){
			_, err := ioutil.ReadFile(test.File)
			if err != nil {
				log.Fatalf("%s unable to read or find file '%s'", test.Name, test.File);
			}
		}

		// set default params
		if(test.Requests == 0){
			test.Requests = config.Requests
		}

		if(test.Concurrency == 0){
			test.Concurrency = config.Concurrency
		}
	}
}
