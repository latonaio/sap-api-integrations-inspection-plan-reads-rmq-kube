package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			InspectionPlanGroup            string `json:"InspectionPlanGroup"`
			InspectionPlan                 string `json:"InspectionPlan"`
			InspectionPlanInternalVersion  string `json:"InspectionPlanInternalVersion"`
			IsDeleted                      bool   `json:"IsDeleted"`
			BillOfOperationsDesc           string `json:"BillOfOperationsDesc"`
			Plant                          string `json:"Plant"`
			BillOfOperationsUsage          string `json:"BillOfOperationsUsage"`
			BillOfOperationsStatus         string `json:"BillOfOperationsStatus"`
			ResponsiblePlannerGroup        string `json:"ResponsiblePlannerGroup"`
			MinimumLotSizeQuantity         string `json:"MinimumLotSizeQuantity"`
			MaximumLotSizeQuantity         string `json:"MaximumLotSizeQuantity"`
			BillOfOperationsUnit           string `json:"BillOfOperationsUnit"`
			ReplacedBillOfOperations       string `json:"ReplacedBillOfOperations"`
			IsMarkedForDeletion            bool   `json:"IsMarkedForDeletion"`
			InspPlanHasMultipleSpec        string `json:"InspPlanHasMultipleSpec"`
			InspSubsetFieldCombination     string `json:"InspSubsetFieldCombination"`
			InspectionPartialLotAssignment string `json:"InspectionPartialLotAssignment"`
			SmplDrawingProcedure           string `json:"SmplDrawingProcedure"`
			SmplDrawingProcedureVersion    string `json:"SmplDrawingProcedureVersion"`
			InspectionLotDynamicLevel      string `json:"InspectionLotDynamicLevel"`
			InspLotDynamicRule             string `json:"InspLotDynamicRule"`
			InspExternalNumberingOfValues  string `json:"InspExternalNumberingOfValues"`
			CreationDate                   string `json:"CreationDate"`
			LastChangeDate                 string `json:"LastChangeDate"`
			ChangeNumber                   string `json:"ChangeNumber"`
			ValidityStartDate              string `json:"ValidityStartDate"`
			ValidityEndDate                string `json:"ValidityEndDate"`
			ChangedDateTime                string `json:"ChangedDateTime"`
		} `json:"results"`
	} `json:"d"`
}
