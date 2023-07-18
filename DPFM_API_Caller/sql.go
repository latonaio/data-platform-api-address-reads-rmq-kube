package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-address-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-address-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var address *dpfm_api_output_formatter.Address
	for _, fn := range accepter {
		switch fn {
		case "Address":
			func() {
				address = c.Address(mtx, input, output, errs, log)
			}()
		case "Addresses":
			func() {
				address = c.Addresses(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Address:         address,
	}

	return data
}

func (c *DPFMAPICaller) Address(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	where := fmt.Sprintf("WHERE address.AddressID = %d ", input.Address.AddressID)

	if input.Address.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND address.IsMarkedForDeletion = %v", where, *input.Address.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_address_address_data AS address
		` + where + ` ORDER BY address.IsMarkedForDeletion ASC, address.AddressID DESC;`,
	)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return &((*data)[0])
}

func (c *DPFMAPICaller) Addresses(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {

	where := "WHERE 1 = 1"
	
	if input.Address.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND address.IsMarkedForDeletion = %v", where, *input.Address.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_address_address_data AS address
		` + where + ` ORDER BY address.IsMarkedForDeletion ASC, address.AddressID DESC;`,
	)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return &((*data)[0])
}
