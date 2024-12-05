package vaahan_service

import "encoding/xml"

func ConvertXMLToJSON(xmlData string) (map[string]interface{}, error) {
	// Define a struct to match the structure of the XML response
	var vehicleDetails struct {
		XMLName              xml.Name `xml:"VehicleDetails"`
		StatusMessage        string   `xml:"stautsMessage"`
		RcRegnNo             string   `xml:"rc_regn_no"`
		RcRegnDt             string   `xml:"rc_regn_dt"`
		RcRegnUpto           string   `xml:"rc_regn_upto"`
		RcPurchaseDt         string   `xml:"rc_purchase_dt"`
		RcOwnerSr            string   `xml:"rc_owner_sr"`
		RcOwnerName          string   `xml:"rc_owner_name"`
		StateCd              string   `xml:"state_cd"`
		RtoCd                string   `xml:"rto_cd"`
		RcPresentAddress     string   `xml:"rc_present_address"`
		RcPermanentAddress   string   `xml:"rc_permanent_address"`
		RcVchCatg            string   `xml:"rc_vch_catg"`
		RcVhClassDesc        string   `xml:"rc_vh_class_desc"`
		RcChasiNo            string   `xml:"rc_chasi_no"`
		RcEngNo              string   `xml:"rc_eng_no"`
		RcMakerDesc          string   `xml:"rc_maker_desc"`
		RcMakerModel         string   `xml:"rc_maker_model"`
		RcBodyTypeDesc       string   `xml:"rc_body_type_desc"`
		RcFuelDesc           string   `xml:"rc_fuel_desc"`
		RcColor              string   `xml:"rc_color"`
		RcNormsDesc          string   `xml:"rc_norms_desc"`
		RcFitUpto            string   `xml:"rc_fit_upto"`
		RcTaxUpto            string   `xml:"rc_tax_upto"`
		RcFinancer           string   `xml:"rc_financer"`
		RcInsuranceComp      string   `xml:"rc_insurance_comp"`
		RcInsurancePolicyNo  string   `xml:"rc_insurance_policy_no"`
		RcInsuranceUpto      string   `xml:"rc_insurance_upto"`
		RcManuMonthYr        string   `xml:"rc_manu_month_yr"`
		RcUnldWt             string   `xml:"rc_unld_wt"`
		RcGvw                string   `xml:"rc_gvw"`
		RcNoCyl              string   `xml:"rc_no_cyl"`
		RcCubicCap           string   `xml:"rc_cubic_cap"`
		RcSeatCap            string   `xml:"rc_seat_cap"`
		RcSleeperCap         string   `xml:"rc_sleeper_cap"`
		RcStandCap           string   `xml:"rc_stand_cap"`
		RcWheelbase          string   `xml:"rc_wheelbase"`
		RcRegisteredAt       string   `xml:"rc_registered_at"`
		RcStatus             string   `xml:"rc_status"`
		RcStatusAsOn         string   `xml:"rc_status_as_on"`
		RcNcrbStatus         string   `xml:"rc_ncrb_status"`
		RcBlacklistStatus    string   `xml:"rc_blacklist_status"`
		RcNocDetails         string   `xml:"rc_noc_details"`
		RcVhType             string   `xml:"rc_vh_type"`
		RcVhClass            string   `xml:"rc_vh_class"`
		RcNocDt              string   `xml:"rc_noc_dt"`
		RcFuelCd             string   `xml:"rc_fuel_cd"`
		RcMakerCd            string   `xml:"rc_maker_cd"`
		RcModelCd            string   `xml:"rc_model_cd"`
		RcNormsCd            string   `xml:"rc_norms_cd"`
		RcSaleAmt            string   `xml:"rc_sale_amt"`
		RcOwnCatgDesc        string   `xml:"rc_own_catg_desc"`
		RcVchCatgDesc        string   `xml:"rc_vch_catg_desc"`
		RcOwnerCdDesc        string   `xml:"rc_owner_cd_desc"`
		RcDeemedOwnerDetails struct {
			AucValidUpto             string `xml:"auc_valid_upto"`
			AuthorizationCertificate string `xml:"authorization_certificate_number"`
			DealerCode               string `xml:"dealer_code"`
			DealerMail               string `xml:"dealer_mail"`
			DealerMobile             string `xml:"dealer_mobile"`
			DealerName               string `xml:"dealer_name"`
		} `xml:"rc_deemed_owner_details"`
		RcVehicleSurrenderedToDealer string `xml:"rc_vehicle_surrendered_to_dealer"`
		RcDealer                     struct {
			DealerContactNo string `xml:"dealer_contact_no"`
			DealerDistrict  string `xml:"dealer_district"`
			DealerPincode   string `xml:"dealer_pincode"`
		} `xml:"rc_dealer"`
		RcOwnerHistory struct {
			OffName   string `xml:"offName"`
			OwnerName string `xml:"owner_name"`
			OwnerSr   string `xml:"owner_sr"`
			StateCd   string `xml:"stateCd"`
		} `xml:"rc_owner_history"`
		RcCurrentAddDistrictCode string `xml:"rc_currentadd_districtcode"`
		RcNonUse                 string `xml:"rc_non_use"`
		RcPassengerTax           string `xml:"rc_passenger_tax"`
		RcGoodsTax               string `xml:"rc_goods_tax"`
		TempPermit               struct {
			RcPermitCode      string `xml:"rc_permit_code"`
			RcPermitIssueDt   string `xml:"rc_permit_issue_dt"`
			RcPermitNo        string `xml:"rc_permit_no"`
			RcPermitType      string `xml:"rc_permit_type"`
			RcPermitValidFrom string `xml:"rc_permit_valid_from"`
			RcPermitValidUpto string `xml:"rc_permit_valid_upto"`
		} `xml:"temp_permit"`
		RcNoOfAxle     string `xml:"rc_no_of_axle"`
		RcQrUrl        string `xml:"rc_qr_url"`
		RcAuthName     string `xml:"rc_auth_name"`
		RcAuthSign     string `xml:"rc_auth_sign"`
		RcApprovalDate string `xml:"rc_approval_date"`
		RcHp           string `xml:"rc_hp"`
		RcMandalDesc   string `xml:"rc_mandal_desc"`
		RcTalukCd      string `xml:"rc_taluk_cd"`
	}

	// Unmarshal the XML data into the struct
	if err := xml.Unmarshal([]byte(xmlData), &vehicleDetails); err != nil {
		return nil, err
	}

	// Convert the struct to a map for JSON response
	vehicleDetailsMap := map[string]interface{}{
		"statusMessage":                vehicleDetails.StatusMessage,
		"rcRegnNo":                     vehicleDetails.RcRegnNo,
		"rcRegnDt":                     vehicleDetails.RcRegnDt,
		"rcRegnUpto":                   vehicleDetails.RcRegnUpto,
		"rcPurchaseDt":                 vehicleDetails.RcPurchaseDt,
		"rcOwnerSr":                    vehicleDetails.RcOwnerSr,
		"rcOwnerName":                  vehicleDetails.RcOwnerName,
		"stateCd":                      vehicleDetails.StateCd,
		"rtoCd":                        vehicleDetails.RtoCd,
		"rcPresentAddress":             vehicleDetails.RcPresentAddress,
		"rcPermanentAddress":           vehicleDetails.RcPermanentAddress,
		"rcVchCatg":                    vehicleDetails.RcVchCatg,
		"rcVhClassDesc":                vehicleDetails.RcVhClassDesc,
		"rcChasiNo":                    vehicleDetails.RcChasiNo,
		"rcEngNo":                      vehicleDetails.RcEngNo,
		"rcMakerDesc":                  vehicleDetails.RcMakerDesc,
		"rcMakerModel":                 vehicleDetails.RcMakerModel,
		"rcBodyTypeDesc":               vehicleDetails.RcBodyTypeDesc,
		"rcFuelDesc":                   vehicleDetails.RcFuelDesc,
		"rcColor":                      vehicleDetails.RcColor,
		"rcNormsDesc":                  vehicleDetails.RcNormsDesc,
		"rcFitUpto":                    vehicleDetails.RcFitUpto,
		"rcTaxUpto":                    vehicleDetails.RcTaxUpto,
		"rcFinancer":                   vehicleDetails.RcFinancer,
		"rcInsuranceComp":              vehicleDetails.RcInsuranceComp,
		"rcInsurancePolicyNo":          vehicleDetails.RcInsurancePolicyNo,
		"rcInsuranceUpto":              vehicleDetails.RcInsuranceUpto,
		"rcManuMonthYr":                vehicleDetails.RcManuMonthYr,
		"rcUnldWt":                     vehicleDetails.RcUnldWt,
		"rcGvw":                        vehicleDetails.RcGvw,
		"rcNoCyl":                      vehicleDetails.RcNoCyl,
		"rcCubicCap":                   vehicleDetails.RcCubicCap,
		"rcSeatCap":                    vehicleDetails.RcSeatCap,
		"rcSleeperCap":                 vehicleDetails.RcSleeperCap,
		"rcStandCap":                   vehicleDetails.RcStandCap,
		"rcWheelbase":                  vehicleDetails.RcWheelbase,
		"rcRegisteredAt":               vehicleDetails.RcRegisteredAt,
		"rcStatus":                     vehicleDetails.RcStatus,
		"rcStatusAsOn":                 vehicleDetails.RcStatusAsOn,
		"rcNcrbStatus":                 vehicleDetails.RcNcrbStatus,
		"rcBlacklistStatus":            vehicleDetails.RcBlacklistStatus,
		"rcNocDetails":                 vehicleDetails.RcNocDetails,
		"rcVhType":                     vehicleDetails.RcVhType,
		"rcVhClass":                    vehicleDetails.RcVhClass,
		"rcNocDt":                      vehicleDetails.RcNocDt,
		"rcFuelCd":                     vehicleDetails.RcFuelCd,
		"rcMakerCd":                    vehicleDetails.RcMakerCd,
		"rcModelCd":                    vehicleDetails.RcModelCd,
		"rcNormsCd":                    vehicleDetails.RcNormsCd,
		"rcSaleAmt":                    vehicleDetails.RcSaleAmt,
		"rcOwnCatgDesc":                vehicleDetails.RcOwnCatgDesc,
		"rcVchCatgDesc":                vehicleDetails.RcVchCatgDesc,
		"rcOwnerCdDesc":                vehicleDetails.RcOwnerCdDesc,
		"rcDeemedOwnerDetails":         vehicleDetails.RcDeemedOwnerDetails,
		"rcVehicleSurrenderedToDealer": vehicleDetails.RcVehicleSurrenderedToDealer,
		"rcDealer":                     vehicleDetails.RcDealer,
		"rcOwnerHistory":               vehicleDetails.RcOwnerHistory,
		"rcCurrentAddDistrictCode":     vehicleDetails.RcCurrentAddDistrictCode,
		"rcNonUse":                     vehicleDetails.RcNonUse,
		"rcPassengerTax":               vehicleDetails.RcPassengerTax,
		"rcGoodsTax":                   vehicleDetails.RcGoodsTax,
		"tempPermit":                   vehicleDetails.TempPermit,
		"rcNoOfAxle":                   vehicleDetails.RcNoOfAxle,
		"rcQrUrl":                      vehicleDetails.RcQrUrl,
		"rcAuthName":                   vehicleDetails.RcAuthName,
		"rcAuthSign":                   vehicleDetails.RcAuthSign,
		"rcApprovalDate":               vehicleDetails.RcApprovalDate,
		"rcHp":                         vehicleDetails.RcHp,
		"rcMandalDesc":                 vehicleDetails.RcMandalDesc,
		"rcTalukCd":                    vehicleDetails.RcTalukCd,
	}

	return vehicleDetailsMap, nil
}
