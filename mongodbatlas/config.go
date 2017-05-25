package mongodbatlas

import (
  "log"
  "github.com/qualialabs/mongodbatlas-go-api/mongodbatlas"
)

type Config struct {
  ATLAS_API_KEY string
  ATLAS_USERNAME string
  ATLAS_GROUP_ID string
}

func (c *Config) Client() (*mongodbatlas.Client, error) {
  client := mongodbatlas.NewClient(c.ATLAS_API_KEY, c.ATLAS_USERNAME, c.ATLAS_GROUP_ID)

  log.Printf("[INFO] Atlas Client configured ")

  return client, nil
}