package mongodbatlas

import (
    // "log"
    // "fmt"
    // "time"
    "github.com/hashicorp/terraform/helper/schema"
    // "github.com/qualialabs/mongodbatlas-go-api/mongodbatlas"
)

func resourceMongodbUser() *schema.Resource {
    return &schema.Resource{
        Create: resourceMongodbUserCreate,
        Read:   resourceMongodbUserRead,
        Update: resourceMongodbUserUpdate,
        Delete: resourceMongodbUserDelete,

        Schema: map[string]*schema.Schema{
            "db_name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "db_user": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "db_password": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


// #######################
func resourceMongodbUserCreate(d *schema.ResourceData, meta interface{}) error {

    // client := meta.(*mongodbatlas.Client)
    // // create user
    // db_name := d.Get("db_name").(string)
    // db_user := d.Get("db_user").(string)
    // db_password := d.Get("db_password").(string)
    // log.Println("[DEBUG] Creating new user" + db_user +" on mongodb " + db_name )

    // mongodbuser := &mongodbatlas.MongodbUser{
    //   Username: db_user,
    //   Password: db_password,
    //   DatabaseName: db_name,
    // }
    // err := client.CreateMongodbUser(mongodbuser)

    // if err != nil {
    //   return fmt.Errorf("Failed to create mongodb user: %s", err.Error())
    // }

    // d.SetId(db_name)
    // log.Printf("[INFO] record ID: %s", d.Id())
    
    return nil
}

func resourceMongodbUserRead(d *schema.ResourceData, meta interface{}) error {

  return nil
}

func resourceMongodbUserUpdate(d *schema.ResourceData, meta interface{}) error {
    // client := meta.(*mongodbatlas.Client)

    // db_name := d.Get("db_name").(string)
    // db_user := d.Get("db_user").(string)
    // db_password := d.Get("db_password").(string)
    // log.Println("[DEBUG] Updating user" + db_user +" on mongodb " + db_name )

    // mongodbuser := &mongodbatlas.MongodbUser{
    //   Username: db_user,
    //   Password: db_password,
    //   DatabaseName: db_name,
    // }
    // err := client.UpdateMongodbUser(mongodbuser)

    // if err != nil {
    //   return fmt.Errorf("Failed to update mongodb user: %s", err.Error())
    // }

    // d.SetId(db_name)
    // log.Printf("[INFO] record ID: %s", d.Id())
    
    return nil
}

func resourceMongodbUserDelete(d *schema.ResourceData, meta interface{}) error {
    // client := meta.(*mongodbatlas.Client)

    // db_name := d.Get("db_name").(string)
    // db_user := d.Get("db_user").(string)
    // db_password := d.Get("db_password").(string)
    // log.Println("[DEBUG] Deleting user" + db_user +" on mongodb " + db_name )

    // mongodbuser := &mongodbatlas.MongodbUser{
    //   Username: db_user,
    //   Password: db_password,
    //   DatabaseName: db_name,
    // }
    // err := client.DeleteMongodbUser(mongodbuser)

    // if err != nil {
    //   return fmt.Errorf("Failed to delete mongodb user: %s", err.Error())
    // }

    return nil

}