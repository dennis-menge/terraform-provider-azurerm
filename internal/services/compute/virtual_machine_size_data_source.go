package compute

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

func dataSourceVirtualMachineSize() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Read: dataSourceVirtualMachineRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"number_of_cores": {
				Type:     pluginsdk.TypeInt,
				Computed: true,
			},
			"os_disk_size_in_mb": {
				Type:     pluginsdk.TypeInt,
				Computed: true,
			},
			"resource_disk_size_in_mb": {
				Type:     pluginsdk.TypeInt,
				Computed: true,
			},
			"memory_in_mb": {
				Type:     pluginsdk.TypeInt,
				Computed: true,
			},
			"max_data_disk_count": {
				Type:     pluginsdk.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceVirtualMachineSizeRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.VMSizeClient

	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	vmSizeName := d.Get("name").(string)
	location := d.Get("location").(string)

	resp, err := client.List(ctx, location)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return fmt.Errorf("%s was not found", location)
		}

		return fmt.Errorf("making Read request on %s: %+v", location, err)
	}

	for _, vmSize := range *resp.Value {
		if *vmSize.Name == vmSizeName {

			d.Set("name", vmSize.Name)
			if err != nil {
				return err
			}

			d.Set("number_of_cores", vmSize.Name)
			if err != nil {
				return err
			}

			d.Set("os_disk_size_in_mb", vmSize.Name)
			if err != nil {
				return err
			}

			d.Set("resource_disk_size_in_mb", vmSize.Name)
			if err != nil {
				return err
			}

			d.Set("memory_in_mb", vmSize.Name)
			if err != nil {
				return err
			}

			d.Set("max_data_disk_count", vmSize.Name)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return nil
}
