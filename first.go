package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type MyEvent struct {
	CivId string
}

func getCivByName(civilization int) map[string]string {

	if civilization > 0 && civilization < 33 {

		f, err := os.Open("/tmp/civilizations.csv")
		if err != nil {
			log.Fatal("Error leyendo archivo", err)
		}
		r := csv.NewReader(bufio.NewReader(f))
		i := 0
		m := make(map[string]string)

		for {
			record, err := r.Read()
			// Stop at EOF.
			if err == io.EOF {
				break
			}

			if i == (civilization) {
				m["name"] = record[0]
				m["specialty"] = record[2]
				m["uniqueUnit"] = record[3]
				break
			}

			i += 1
		}

		return m

	}

	log.Println("Valor no aceptado")
	m := make(map[string]string)
	m["error"] = "parameter not supported"
	return m

}

func example(event MyEvent) (map[string]string, error) {
	// NOTE: you need to store your AWS credentials in ~/.aws/credentials
	log.Print("Event:", event)

	// 1) Define your bucket and item names
	bucket := "wizeline-go-bootcamp-canodez"
	item := "civilizations.csv"

	// 2) Create an AWS session
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// 3) Create a new AWS S3 downloader
	downloader := s3manager.NewDownloader(sess)

	// 4) Download the item from the bucket. If an error occurs, log it and exit. Otherwise, notify the user that the download succeeded.
	file, err := os.Create("/tmp/" + item)
	if err != nil {
		fmt.Println("Error creando archivo", err)
	}

	defer file.Close()

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

	intVar, err := strconv.Atoi(event.CivId)
	if err != nil {
		println("Hubo error convirtiendo a string: ", err)
	}

	log.Println(intVar)

	result := getCivByName(intVar)

	log.Println("CivName: ", result["name"])
	log.Println("Specialty: ", result["specialty"])
	log.Println("UniqueUnit: ", result["uniqueUnit"])

	return result, nil

}

func main() {

	lambda.Start(example)

}
