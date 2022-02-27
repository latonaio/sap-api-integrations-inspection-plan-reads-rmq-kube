package responses

type MaterialAssignment struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			Material                      string      `json:"Material"`
			Plant                         string      `json:"Plant"`
			InspectionPlanGroup           string      `json:"InspectionPlanGroup"`
			InspectionPlan                string      `json:"InspectionPlan"`
			InspPlanMatlAssignment        string      `json:"InspPlanMatlAssignment"`
			InspPlanMatlAssgmtIntVersion  string      `json:"InspPlanMatlAssgmtIntVersion"`
			InspectionPlanInternalVersion string      `json:"InspectionPlanInternalVersion"`
			ValidityStartDate             string      `json:"ValidityStartDate"`
			ValidityEndDate               string      `json:"ValidityEndDate"`
			ChangeNumber                  string      `json:"ChangeNumber"`
			CreationDate                  string      `json:"CreationDate"`
			LastChangeDate                string      `json:"LastChangeDate"`
			IsDeleted                     bool        `json:"IsDeleted"`
			Supplier                      string      `json:"Supplier"`
			Customer                      string      `json:"Customer"`
			MultipleSpecificationObject   string      `json:"MultipleSpecificationObject"`
			MultipleSpecificationObjType  string      `json:"MultipleSpecificationObjType"`
			BOOSearchText                 string      `json:"BOOSearchText"`
			ChangedDateTime               string      `json:"ChangedDateTime"`
		} `json:"results"`
	} `json:"d"`
}
