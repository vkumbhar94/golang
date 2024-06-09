package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/vkumbhar94/golang/internal/util"
)

func main() {
	client := util.MinIOS3Client()
	// fmt.Println("processing file.parquet")
	// file(client)
	// fmt.Println("processing file-nil.parquet")
	// fileNil(client)
	fmt.Println("processing file-multitype.parquet")
	fileMultiType(client)
}

func fileMultiType(client *s3.Client) {
	params := &s3.SelectObjectContentInput{
		Bucket:         aws.String("localdev"),
		Key:            aws.String("upload-file-multitype.parquet"),
		ExpressionType: types.ExpressionTypeSql,
		// Expression:     aws.String("SELECT * FROM S3Object"),
		Expression: aws.String("SELECT max(Id_Int64) FROM S3Object"),
		// Expression: aws.String("SELECT FirstName_String FROM S3Object"),
		// Expression: aws.String("SELECT * FROM S3Object where Id_Int64 = 123"),
		// Expression: aws.String("SELECT * FROM S3Object where Age_Int64 > 25"),
		// Expression: aws.String("SELECT * FROM S3Object where FirstName_String like '%i%'"),
		// Expression: aws.String("SELECT * FROM S3Object where LastName_String like '%o%'"),
		// Expression: aws.String("SELECT * FROM S3Object where Age_Int64 > 25 and LastName_String like '%o%'"),
		InputSerialization: &types.InputSerialization{
			// CSV: &types.CSVInput{
			// 	FileHeaderInfo: types.FileHeaderInfoUse,
			// },
			Parquet: &types.ParquetInput{},
		},
		OutputSerialization: &types.OutputSerialization{
			// CSV: &types.CSVOutput{},
			JSON: &types.JSONOutput{},
		},
	}
	content, err := client.SelectObjectContent(context.TODO(), params)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%#v\n", content.ResultMetadata)

	for v := range content.GetStream().Events() {
		fmt.Printf("%T\n", v)
		switch v1 := v.(type) {
		case *types.SelectObjectContentEventStreamMemberProgress:
			fmt.Println("Progress", *v1.Value.Details.BytesProcessed)
		case *types.SelectObjectContentEventStreamMemberRecords:
			b := bytes.NewReader(v1.Value.Payload)
			br := bufio.NewReader(b)
			for {
				// isPrefix is set to true to indicate that there is more data to read
				// but the buffer size is not large enough to fit the next line, so
				// the reader should read more data into the buffer.
				line, _, err := br.ReadLine()
				if err != nil {
					break
				}
				fmt.Println(string(line))
			}
		case *types.SelectObjectContentEventStreamMemberStats:
			fmt.Println("Stats BytesProcessed", *v1.Value.Details.BytesProcessed)
			fmt.Println("Stats BytesReturned", *v1.Value.Details.BytesReturned)
			fmt.Println("Stats BytesScanned", *v1.Value.Details.BytesScanned)
		case *types.SelectObjectContentEventStreamMemberEnd:
			fmt.Println("End of stream")
		default:
			fmt.Println("Unknown type")

		}
	}
}

func fileNil(client *s3.Client) {
	params := &s3.SelectObjectContentInput{
		Bucket:         aws.String("localdev"),
		Key:            aws.String("file-nil.parquet"),
		ExpressionType: types.ExpressionTypeSql,
		Expression:     aws.String("SELECT FirstName, LastName FROM S3Object"),
		InputSerialization: &types.InputSerialization{
			// CSV: &types.CSVInput{
			// 	FileHeaderInfo: types.FileHeaderInfoUse,
			// },
			Parquet: &types.ParquetInput{},
		},
		OutputSerialization: &types.OutputSerialization{
			// CSV: &types.CSVOutput{},
			JSON: &types.JSONOutput{},
		},
	}
	content, err := client.SelectObjectContent(context.TODO(), params)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%#v\n", content.ResultMetadata)

	for v := range content.GetStream().Events() {
		fmt.Printf("%T\n", v)
		switch v1 := v.(type) {
		case *types.SelectObjectContentEventStreamMemberProgress:
			fmt.Println("Progress", *v1.Value.Details.BytesProcessed)
		case *types.SelectObjectContentEventStreamMemberRecords:
			fmt.Println("Records\n", string(v1.Value.Payload))
			b := bytes.NewReader(v1.Value.Payload)
			br := bufio.NewReader(b)
			for {
				// isPrefix is set to true to indicate that there is more data to read
				// but the buffer size is not large enough to fit the next line, so
				// the reader should read more data into the buffer.
				line, _, err := br.ReadLine()
				if err != nil {
					break
				}
				fmt.Println(string(line))
			}
		case *types.SelectObjectContentEventStreamMemberStats:
			fmt.Println("Stats BytesProcessed", *v1.Value.Details.BytesProcessed)
			fmt.Println("Stats BytesReturned", *v1.Value.Details.BytesReturned)
			fmt.Println("Stats BytesScanned", *v1.Value.Details.BytesScanned)
		case *types.SelectObjectContentEventStreamMemberEnd:
			fmt.Println("End of stream")
		default:
			fmt.Println("Unknown type")

		}
	}
}

func file(client *s3.Client) {
	params := &s3.SelectObjectContentInput{
		Bucket:         aws.String("localdev"),
		Key:            aws.String("file.parquet"),
		ExpressionType: types.ExpressionTypeSql,
		Expression:     aws.String("SELECT FirstName FROM S3Object"),
		InputSerialization: &types.InputSerialization{
			// CSV: &types.CSVInput{
			// 	FileHeaderInfo: types.FileHeaderInfoUse,
			// },
			Parquet: &types.ParquetInput{},
		},
		OutputSerialization: &types.OutputSerialization{
			// CSV: &types.CSVOutput{},
			JSON: &types.JSONOutput{},
		},
	}
	content, err := client.SelectObjectContent(context.TODO(), params)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%#v\n", content.ResultMetadata)

	for v := range content.GetStream().Events() {
		fmt.Printf("%T\n", v)
		switch v1 := v.(type) {
		case *types.SelectObjectContentEventStreamMemberProgress:
			fmt.Println("Progress", *v1.Value.Details.BytesProcessed)
		case *types.SelectObjectContentEventStreamMemberRecords:
			fmt.Println("Records\n", string(v1.Value.Payload))
		case *types.SelectObjectContentEventStreamMemberStats:
			fmt.Println("Stats BytesProcessed", *v1.Value.Details.BytesProcessed)
			fmt.Println("Stats BytesReturned", *v1.Value.Details.BytesReturned)
			fmt.Println("Stats BytesScanned", *v1.Value.Details.BytesScanned)
		case *types.SelectObjectContentEventStreamMemberEnd:
			fmt.Println("End of stream")
		default:
			fmt.Println("Unknown type")

		}
	}
}
