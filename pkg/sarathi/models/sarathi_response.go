package models

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    sarathiResponse, err := UnmarshalSarathiResponse(bytes)
//    bytes, err = sarathiResponse.Marshal()

import "encoding/json"

func (r *SarathiResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SarathiResponse struct {
	Response []ResponseElement `json:"response"`
	Error    string            `json:"error"`
	Code     string            `json:"code"`
	Message  string            `json:"message"`
}

type ResponseElement struct {
	Response       ResponseResponse `json:"response"`
	ResponseStatus string           `json:"responseStatus"`
}

type ResponseResponse struct {
	Dldetobj []Dldetobj `json:"dldetobj"`
}

type Dldetobj struct {
	TransReqObj      interface{} `json:"transReqObj"`
	DLTest           interface{} `json:"dlTest"`
	DLEndorsementObj interface{} `json:"dlEndorsementObj"`
	Objections       interface{} `json:"objections"`
	PsvBadgeissued   interface{} `json:"psvBadgeissued"`
	DBLOC            string      `json:"dbLoc"`
	DLHzHillObj      interface{} `json:"dlHzHillObj"`
	Errorcd          int64       `json:"errorcd"`
	DLCovHistObj     interface{} `json:"dlCovHistObj"`
	Dlcovs           []Dlcov     `json:"dlcovs"`
	BioObj           BioObj      `json:"bioObj"`
	IDIdpIssuedpPOJO interface{} `json:"idIdpIssuedpPojo"`
	Dlobj            Dlobj       `json:"dlobj"`
	Erormsg          interface{} `json:"erormsg"`
	BioImgObj        interface{} `json:"bioImgObj"`
	DLHistObj        interface{} `json:"dlHistObj"`
}

type BioObj struct {
	BioPermSDName            interface{} `json:"bioPermSdName"`
	BioDigest                interface{} `json:"bioDigest"`
	BioStayperiodPresentAddr interface{} `json:"bioStayperiodPresentAddr"`
	BioTempPin               string      `json:"bioTempPin"`
	BioBioidSearch           interface{} `json:"bioBioidSearch"`
	BioTempAdd2              string      `json:"bioTempAdd2"`
	BioTempAdd1              string      `json:"bioTempAdd1"`
	BioTokenID               interface{} `json:"bioTokenId"`
	BioEndorsedt             interface{} `json:"bioEndorsedt"`
	BioTempAdd3              string      `json:"bioTempAdd3"`
	BioSwdLname              interface{} `json:"bioSwdLname"`
	BioRecGenesis            string      `json:"bioRecGenesis"`
	BioPermVillTownCD        interface{} `json:"bioPermVillTownCd"`
	BioIdentityMark1         interface{} `json:"bioIdentityMark1"`
	BioIdentityMark2         interface{} `json:"bioIdentityMark2"`
	BioUserID                string      `json:"bioUserId"`
	BioNatName               string      `json:"bioNatName"`
	BioQmQualcd              string      `json:"bioQmQualcd"`
	BioNprNo                 interface{} `json:"bioNprNo"`
	BioAadhaarNo             interface{} `json:"bioAadhaarNo"`
	BioGender                string      `json:"bioGender"`
	BioPermDistCD            interface{} `json:"bioPermDistCd"`
	BioTempDistCD            interface{} `json:"bioTempDistCd"`
	BioPhoneNo               interface{} `json:"bioPhoneNo"`
	BioTempVillTownCD        interface{} `json:"bioTempVillTownCd"`
	BioEndorsetime           interface{} `json:"bioEndorsetime"`
	BioFirstName             string      `json:"bioFirstName"`
	BioMobileNo              string      `json:"bioMobileNo"`
	BioPermVillTownName      interface{} `json:"bioPermVillTownName"`
	BioOrganDonor            string      `json:"bioOrganDonor"`
	BioTempSDName            string      `json:"bioTempSdName"`
	BioBloodGroupname        string      `json:"bioBloodGroupname"`
	BioTempLocal             string      `json:"bioTempLocal"`
	Dob                      interface{} `json:"dob"`
	BioPermLOCType           string      `json:"bioPermLocType"`
	BioPermSdcode            string      `json:"bioPermSdcode"`
	FullAddress              interface{} `json:"fullAddress"`
	BioDob                   string      `json:"bioDob"`
	BioApplno                int64       `json:"bioApplno"`
	BioAadhaarName           interface{} `json:"bioAadhaarName"`
	BioBioID                 string      `json:"bioBioId"`
	BioEmailID               string      `json:"bioEmailId"`
	BioPoliceStncd           interface{} `json:"bioPoliceStncd"`
	BioPermPin               string      `json:"bioPermPin"`
	BioTempLOCType           string      `json:"bioTempLocType"`
	BIPhoto                  interface{} `json:"biPhoto"`
	BioBirthplace            interface{} `json:"bioBirthplace"`
	BioMiddleName            interface{} `json:"bioMiddleName"`
	BioPermDistName          interface{} `json:"bioPermDistName"`
	BioPermAdd3              interface{} `json:"bioPermAdd3"`
	BioTempVillTownName      interface{} `json:"bioTempVillTownName"`
	BioAltMobileNo           interface{} `json:"bioAltMobileNo"`
	BioPermAdd1              string      `json:"bioPermAdd1"`
	BioLastName              string      `json:"bioLastName"`
	BioPermAdd2              string      `json:"bioPermAdd2"`
	BioFullName              string      `json:"bioFullName"`
	BioGenderDesc            string      `json:"bioGenderDesc"`
	BioPerDetAadhaar         interface{} `json:"bioPerDetAadhaar"`
	BioSwdFname              string      `json:"bioSwdFname"`
	BioTempSdcode            string      `json:"bioTempSdcode"`
	BioQmQualdesc            interface{} `json:"bioQmQualdesc"`
	BioEndorsementNo         interface{} `json:"bioEndorsementNo"`
	BioApplicantCatg         string      `json:"bioApplicantCatg"`
	BioSwdFullName           string      `json:"bioSwdFullName"`
	BioTempDistName          interface{} `json:"bioTempDistName"`
	BioCitiZen               string      `json:"bioCitiZen"`
	BioSwdMname              interface{} `json:"bioSwdMname"`
	BioBloodGroup            string      `json:"bioBloodGroup"`
	BioPermLocal             string      `json:"bioPermLocal"`
	AadharAuthenticated      bool        `json:"aadharAuthenticated"`
	BioDlno                  string      `json:"bioDlno"`
	Pht                      interface{} `json:"pht"`
	BioDependentRelation     string      `json:"bioDependentRelation"`
}

type Dlcov struct {
	DLBacklogCovEndtime interface{} `json:"dlBacklogCovEndtime"`
	DcApplno            int64       `json:"dcApplno"`
	BadgeIssuedAuth     interface{} `json:"badgeIssuedAuth"`
	DcCovStatus         string      `json:"dcCovStatus"`
	DcInvcrgNo          interface{} `json:"dcInvcrgNo"`
	Olacd               string      `json:"olacd"`
	Endouserid          int64       `json:"endouserid"`
	Covabbrv            string      `json:"covabbrv"`
	DcIssuedt           string      `json:"dcIssuedt"`
	OlaName             string      `json:"olaName"`
	DbcImvName          interface{} `json:"dbcImvName"`
	Vecatg              string      `json:"vecatg"`
	DcReflicNo          interface{} `json:"dcReflicNo"`
	BadgeNo             interface{} `json:"badgeNo"`
	VeBadgeIssue        interface{} `json:"veBadgeIssue"`
	VeShortdesc         string      `json:"veShortdesc"`
	DcEndorsedt         string      `json:"dcEndorsedt"`
	DcEndorsetime       string      `json:"dcEndorsetime"`
	DcAuthDt            interface{} `json:"dcAuthDt"`
	CovIssueAuthCode    string      `json:"covIssueAuthCode"`
	DcTokenID           int64       `json:"dcTokenId"`
	BadgeIssuedt        interface{} `json:"badgeIssuedt"`
	DcCovcd             int64       `json:"dcCovcd"`
	DcReflicType        string      `json:"dcReflicType"`
	CovIssuedt          interface{} `json:"covIssuedt"`
	DbcImvDesig         interface{} `json:"dbcImvDesig"`
	DcEndorseNo         interface{} `json:"dcEndorseNo"`
	DcLicno             string      `json:"dcLicno"`
	DcAuthNo            interface{} `json:"dcAuthNo"`
	Covdesc             string      `json:"covdesc"`
	DcInvrgdesc         interface{} `json:"dcInvrgdesc"`
	DLTestdate          interface{} `json:"dlTestdate"`
}

type Dlobj struct {
	EnforceFromDate     interface{} `json:"enforceFromDate"`
	DLLicno             string      `json:"dlLicno"`
	DLDispatchStatus    string      `json:"dlDispatchStatus"`
	OmRtoFullname       string      `json:"omRtoFullname"`
	DLTokenID           int64       `json:"dlTokenId"`
	DLStateCode         interface{} `json:"dlStateCode"`
	EnforceRemark       string      `json:"enforceRemark"`
	DLIncChallanNo      string      `json:"dlIncChallanNo"`
	DLIssueauth         string      `json:"dlIssueauth"`
	StateName           string      `json:"stateName"`
	DLBioID             interface{} `json:"dlBioId"`
	DLAuthNo            interface{} `json:"dlAuthNo"`
	DLNTValdtoDate      interface{} `json:"dlNtValdtoDate"`
	DLHzValdtoDt        interface{} `json:"dlHzValdtoDt"`
	DLLatestTrcode      int64       `json:"dlLatestTrcode"`
	OmOfficeTownname    string      `json:"omOfficeTownname"`
	DLInvcrgNo          interface{} `json:"dlInvcrgNo"`
	DLAuthDt            interface{} `json:"dlAuthDt"`
	DLIssuedt           string      `json:"dlIssuedt"`
	DLHlValdtoDate      interface{} `json:"dlHlValdtoDate"`
	DLEndorseno         interface{} `json:"dlEndorseno"`
	Bioid               string      `json:"bioid"`
	Olacode             string      `json:"olacode"`
	DLIncSourceType     string      `json:"dlIncSourceType"`
	DLOldLicno          string      `json:"dlOldLicno"`
	DLHzValdfrDt        interface{} `json:"dlHzValdfrDt"`
	DLIntermediateStage string      `json:"dlIntermediateStage"`
	DLIssueDate         interface{} `json:"dlIssueDate"`
	EnforceEndDate      interface{} `json:"enforceEndDate"`
	DLSeqno             interface{} `json:"dlSeqno"`
	DLTrValdtoDt        interface{} `json:"dlTrValdtoDt"`
	DLRemarks           interface{} `json:"dlRemarks"`
	DLRecGenesis        string      `json:"dlRecGenesis"`
	DLNTValdfrDt        string      `json:"dlNtValdfrDt"`
	DLEndorsetime       string      `json:"dlEndorsetime"`
	DLHlValdtoDt        interface{} `json:"dlHlValdtoDt"`
	DLHzValdtoDate      interface{} `json:"dlHzValdtoDate"`
	DLEndorseAuth       string      `json:"dlEndorseAuth"`
	DLStatus            string      `json:"dlStatus"`
	OlaName             string      `json:"olaName"`
	Statecd             string      `json:"statecd"`
	DLUsid              int64       `json:"dlUsid"`
	DLEndorsedt         string      `json:"dlEndorsedt"`
	DLAuthIssauth       interface{} `json:"dlAuthIssauth"`
	DLIssuedesig        interface{} `json:"dlIssuedesig"`
	OmRtoShortname      string      `json:"omRtoShortname"`
	DLHlValdfrDt        interface{} `json:"dlHlValdfrDt"`
	DLTrValdtoDate      interface{} `json:"dlTrValdtoDate"`
	DLAuthCov           interface{} `json:"dlAuthCov"`
	DLPrintDate         string      `json:"dlPrintDate"`
	DLPrintStatus       string      `json:"dlPrintStatus"`
	DLApplno            int64       `json:"dlApplno"`
	DLDigest            interface{} `json:"dlDigest"`
	DLIncRtoAction      string      `json:"dlIncRtoAction"`
	DLRtoCode           string      `json:"dlRtoCode"`
	DLTrValdfrDt        interface{} `json:"dlTrValdfrDt"`
	DlolaCode           interface{} `json:"dlolaCode"`
	DLNTValdtoDt        string      `json:"dlNtValdtoDt"`
}
