package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	aws2 "github.com/Projects/coin-purchase/internal/aws"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"os"
)

const (
	ExchangesCoinsMapFileName = "exchanges-coin-map.json"
)

func PutJsonDataToBucket(objectFileName string, data []byte) error {
	newSession, err := session.NewSession(&aws.Config{})

	if err != nil {
		return fmt.Errorf("could not initialize new aws session: %v", err)
	}

	s3Client := s3.New(newSession)
	bucket := os.Getenv("CRYPTO_DATA_BUCKET")

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(objectFileName),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("application/json"),
	})

	return nil
}

func GetExchangeCoinsMap() (ExchangesCoinsMap, error) {
	object, err := aws2.GetS3Object(os.Getenv("CRYPTO_DATA_BUCKET"), ExchangesCoinsMapFileName)

	if err != nil {
		return nil, err
	}

	defer object.Body.Close()
	body, err := ioutil.ReadAll(object.Body)

	if err != nil {
		return nil, err
	}

	var exchangeCoinMap ExchangesCoinsMap
	err = json.Unmarshal(body, &exchangeCoinMap)

	if err != nil {
		return nil, err
	}

	return exchangeCoinMap, nil
}
