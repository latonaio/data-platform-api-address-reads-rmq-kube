package dpfm_api_output_formatter

import (
	"data-platform-api-address-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-address-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToAddress(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]Address, error) {
	defer rows.Close()
	addresses := make([]Address, 0, len(sdc.Address))

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Address{}

		err := rows.Scan(
			&pm.AddressID,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.PostalCode,
			&pm.LocalSubRegion,
			&pm.LocalRegion,
			&pm.Country,
			&pm.GlobalRegion,
			&pm.TimeZone,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.CityName,
			&pm.Room,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &addresses, err
		}

		data := pm

		addresses = append(addresses, Address{
			AddressID:				data.AddressID,
			ValidityStartDate:		data.ValidityStartDate,
			ValidityEndDate:		data.ValidityEndDate,
			PostalCode:				data.PostalCode,
			LocalSubRegion:			data.LocalSubRegion,
			LocalRegion:			data.LocalRegion,
			Country:				data.Country,
			GlobalRegion:			data.GlobalRegion,
			TimeZone:				data.TimeZone,
			District:				data.District,
			StreetName:				data.StreetName,
			CityName:				data.CityName,
			Building:				data.Building,
			Floor:					data.Floor,
			Room:					data.Room,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &addresses, nil
	}

	return &addresses, nil
}
