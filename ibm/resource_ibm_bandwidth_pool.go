package ibm

import (
	"fmt"
	"log"
	"strconv"

	"github.com/softlayer/softlayer-go/sl"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

func resourceIBMBandwidthPool() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMBandwidthPoolCreate,
		Read:     resourceIBMBandwidthPoolRead,
		Update:   resourceIBMBandwidthPoolUpdate,
		Delete:   resourceIBMBandwidthPoolDelete,
		Exists:   resourceIBMBandwidthPoolExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"locationgroupid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceIBMBandwidthPoolCreate(d *schema.ResourceData, meta interface{}) error {
	///create  session
	sess := meta.(ClientSession).SoftLayerSession()
	log.Println("ordering bandwidth Pool service...")
	///get the value of all the parameters
	name := d.Get("name").(string)
	locationGroupID := d.Get("locationgroupd").(int)
	///creat an object of Bandwidth Pool service
	service := services.GetNetworkBandwidthVersion1AllotmentService(sess)
	account := services.GetAccountService(sess)
	// get the account Id
	accountData, err := account.Mask("id").GetObject()
	if err != nil {
		return fmt.Errorf("Error retreiving Account Details: %s", err)
	}
	////pass the parameters to create bandwidth pool
	receipt1, err := service.CreateObject(&datatypes.Network_Bandwidth_Version1_Allotment{
		AccountId:                sl.Int(*accountData.Id),
		BandwidthAllotmentTypeId: sl.Int(2),
		LocationGroupId:          sl.Int(locationGroupID),
		Name:                     sl.String(name),
	})
	if err != nil {
		return fmt.Errorf("Error creating Bandwidth Pool: %s", err)
	}
	id := strconv.Itoa(*receipt1.Id)
	d.SetId(id)
	log.Println(d.Id())
	return resourceIBMBandwidthPoolRead(d, meta)
}

func resourceIBMBandwidthPoolRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	log.Println("reading BW Pool service...")
	bwID, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Not a valid BW pool ID, must be an integer: %s", err)
	}
	bw, err := services.GetNetworkBandwidthVersion1AllotmentService(sess).Id(bwID).GetObject()
	if err != nil {
		return fmt.Errorf("Error retrieving virtual guest: %s", err)
	}
	d.Set("name", *bw.Name)
	return nil
}

func resourceIBMBandwidthPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	log.Printf("[INFO] Updating BW Pool")
	bwID, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Not a valid BW pool ID, must be an integer: %s", err)
	}
	result, err := services.GetNetworkBandwidthVersion1AllotmentService(sess).Id(bwID).GetObject()
	if err != nil {
		return fmt.Errorf("Error retrieving virtual guest: %s", err)
	}
	isChanged := false

	if d.HasChange("name") {
		result.Name = sl.String(d.Get("name").(string))
		isChanged = true
	}
	if isChanged {
		_, err = services.GetNetworkBandwidthVersion1AllotmentService(sess).Id(bwID).EditObject(&result)
		if err != nil {
			return fmt.Errorf("Couldn't update virtual guest: %s", err)
		}
	}
	return resourceIBMBandwidthPoolRead(d, meta)
}

func resourceIBMBandwidthPoolDelete(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	log.Printf("[INFO] Deleting BW Pool")
	///pass the id to delete the resource.
	bwID, err := strconv.Atoi(d.Id())
	_, err = services.GetNetworkBandwidthVersion1AllotmentService(sess).
		Id(bwID).
		RequestVdrCancellation()
	if err != nil {
		log.Println("error destroying")
		log.Println(err)
	}
	d.SetId("")
	return nil
}

func resourceIBMBandwidthPoolExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess := meta.(ClientSession).SoftLayerSession()
	log.Println("Exists BW pool service...")
	//service := services.GetNetworkBandwidthVersion1AllotmentService(sess)
	///check if the resource exists with the given id.
	bwID, err := strconv.Atoi(d.Id())
	_, err = services.GetNetworkBandwidthVersion1AllotmentService(sess).
		Id(bwID).
		GetObject()
	//exists, err := service.GetObject(id)
	if err != nil {
		if apiErr, ok := err.(sl.Error); ok {
			if apiErr.StatusCode == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error retrieving BW pool info: %s", err)
	}
	return true, nil
}
