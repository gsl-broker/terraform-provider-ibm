package ibm

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMBandwidthPool_Basic(t *testing.T) {
	var bwPool datatypes.Network_Bandwidth_Version1_Allotment

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckIBMBandwidthPooDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMBandwidthPool_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMBandwidthPoolExists("ibm_bandwidth_pool.BWPool_first", &bwPool),
					resource.TestCheckResourceAttr(
						"ibm_bandwidth_pool.BWPool_first", "accountId", "1521909"),
					resource.TestCheckResourceAttr(
						"ibm_bandwidth_pool.BWPool_first", "bandwidthAllotmentTypeId", "2"),
					resource.TestCheckResourceAttr(
						"ibm_bandwidth_pool.BWPool_first", "name", "Checkdelete1"),
					resource.TestCheckResourceAttr(
						"ibm_bandwidth_pool.BWPool_first", "locationGroupId", "1"),
				),
			},
		},
	})
}

// func testAccCheckIBMBandwidthPooDestroy(s *terraform.State) error {
// 	service := services.GetNetworkBandwidthVersion1AllotmentService(testAccProvider.Meta().(ClientSession).SoftLayerSession())

// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "ibm_bandwidth_pool" {
// 			continue
// 		}
// 		bwID, err := strconv.Atoi(rs.Primary.ID)
// 		if err != nil {
// 			return err
// 		}
// 		// Try to find the domain
// 		response, err := service.Id(bwID).GetObject()
// 		fmt.Print(response)
// 		if err == nil {
// 			return fmt.Errorf("BW Pool with id %d still exists", bwID)
// 		}
// 	}

// 	return nil
// }

func testAccCheckIBMBandwidthPoolExists(n string, bwPool *datatypes.Network_Bandwidth_Version1_Allotment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}
		bwID, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}
		service := services.GetNetworkBandwidthVersion1AllotmentService(testAccProvider.Meta().(ClientSession).SoftLayerSession())

		foundId, err := service.Id(bwID).GetObject()
		if err != nil {
			return err
		}
		resourceId := strconv.Itoa(*foundId.Id)
		if err != nil {
			return err
		}
		if resourceId != rs.Primary.ID {
			return errors.New("Record not found")
		}
		return nil
	}
}

const testAccCheckIBMBandwidthPool_basic = `
resource "ibm_bandwidth_pool" "BWPool_first" {
	name="Checkdelete1"
	locationGroupId=1
	bandwidthAllotmentTypeId=2
	accountId=123456
	}`
