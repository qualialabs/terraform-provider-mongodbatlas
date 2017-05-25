package mongodbatlas

import (
    "log"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
    return &schema.Provider{

      Schema: map[string]*schema.Schema{
        "atlas_api_key": &schema.Schema{
          Type:        schema.TypeString,
          Optional:    true,
          DefaultFunc: schema.EnvDefaultFunc("ATLAS_API_KEY", nil),
        },
        "atlas_username": &schema.Schema{
          Type:        schema.TypeString,
          Optional:    true,
          DefaultFunc: schema.EnvDefaultFunc("ATLAS_USERNAME", nil),
        },
        "atlas_group_id": &schema.Schema{
          Type:        schema.TypeString,
          Optional:    true,
          DefaultFunc: schema.EnvDefaultFunc("ATLAS_GROUP_ID", nil),
        },

      },


      ResourcesMap: map[string]*schema.Resource{
          "mongodbatlas_mongodbuser": resourceMongodbUser(),
      },

      ConfigureFunc: providerConfigure,

    }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

  config := Config{
    ATLAS_API_KEY: d.Get("atlas_api_key").(string),
    ATLAS_USERNAME: d.Get("atlas_username").(string),
    ATLAS_GROUP_ID: d.Get("atlas_group_id").(string),    
  }

  log.Println("[INFO] Initializing Atlas client")
  return config.Client()
}