package main

import (
	aws2 "github.com/aws/aws-sdk-go/aws"
	s32 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/galaco/aggregatier/cmd/common/aws"
	"github.com/galaco/aggregatier/cmd/common/aws/s3"
	"github.com/galaco/aggregatier/cmd/smite/common"
	"github.com/galaco/aggregatier/server/config"
	"github.com/galaco/aggregatier/server/models"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

const s3UrlPrefix = "https://s3."
const s3UrlSuffix = ".amazonaws.com/"
const s3BucketName = "<BUCKET-NAME>" // e.g mycdn.galaco.me
const s3BucketRegion = "<REGION>" //e.g. eu-west-1
const s3SmiteGodsBasePath = "smite/gods/" // path to icon storage
const s3BucketUrl = s3UrlPrefix + s3BucketRegion + s3UrlSuffix + s3BucketName

func main() {
	configuration, err := config.NewConfig("config.json")
	if err != nil {
		panic(err)
	}

	api := common.NewApi(configuration.Smite.DevId, configuration.Smite.AuthKey)

	session, err := api.CreateSession()
	if err != nil {
		panic(err)
	}
	heroes, err := api.GetGods(session.Id)
	if err != nil {
		panic(err)
	}

	models.InitDB(configuration.Database.Username +
		":" +
		configuration.Database.Password +
		"@tcp(" +
		configuration.Database.Host +
		")/" +
		configuration.Database.Name)

	// Initialize S3
	awsSession := aws.NewSession(s3BucketRegion)
	s3Provider := s3.NewS3(awsSession)

	existingThumbnails := map[string]bool{}

	i := 0
	err = s3Provider.ListObjectsPages(&s32.ListObjectsInput{
		Bucket: aws2.String(s3BucketName),
	}, func(p *s32.ListObjectsOutput, last bool) (shouldContinue bool) {
		i++
		for _, obj := range p.Contents {
			dir, file := filepath.Split(*obj.Key)
			if dir != s3SmiteGodsBasePath {
				continue
			}
			existingThumbnails[file] = true
		}
		return true
	})
	if err != nil {
		panic(err)
	}
	for _, hero := range heroes {
		// upload to s3
		_,file := filepath.Split(hero.IconUrl)
		if _,ok := existingThumbnails[file]; !ok {
			// icon doesnt exist
			resp, err := http.Get(hero.IconUrl)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			iconData, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			// Put icon into s3
			if _,err = s3.PutObject(s3Provider, iconData, s3BucketName, s3SmiteGodsBasePath + file); err != nil {
				panic(err)
			}
		}

		// Add hero to database
		if err := models.InsertHero(hero.Id, hero.Name, 1, s3BucketUrl + "/" + s3SmiteGodsBasePath + file); err != nil {
			panic(err)
		}
		log.Printf("Added god: %s\n", hero.Name)
	}
}
