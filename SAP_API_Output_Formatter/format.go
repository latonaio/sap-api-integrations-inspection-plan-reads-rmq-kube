package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-inspection-plan-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			InspectionPlanGroup:            data.InspectionPlanGroup,
			InspectionPlan:                 data.InspectionPlan,
			InspectionPlanInternalVersion:  data.InspectionPlanInternalVersion,
			IsDeleted:                      data.IsDeleted,
			BillOfOperationsDesc:           data.BillOfOperationsDesc,
			Plant:                          data.Plant,
			BillOfOperationsUsage:          data.BillOfOperationsUsage,
			BillOfOperationsStatus:         data.BillOfOperationsStatus,
			ResponsiblePlannerGroup:        data.ResponsiblePlannerGroup,
			MinimumLotSizeQuantity:         data.MinimumLotSizeQuantity,
			MaximumLotSizeQuantity:         data.MaximumLotSizeQuantity,
			BillOfOperationsUnit:           data.BillOfOperationsUnit,
			ReplacedBillOfOperations:       data.ReplacedBillOfOperations,
			IsMarkedForDeletion:            data.IsMarkedForDeletion,
			InspPlanHasMultipleSpec:        data.InspPlanHasMultipleSpec,
			InspSubsetFieldCombination:     data.InspSubsetFieldCombination,
			InspectionPartialLotAssignment: data.InspectionPartialLotAssignment,
			SmplDrawingProcedure:           data.SmplDrawingProcedure,
			SmplDrawingProcedureVersion:    data.SmplDrawingProcedureVersion,
			InspectionLotDynamicLevel:      data.InspectionLotDynamicLevel,
			InspLotDynamicRule:             data.InspLotDynamicRule,
			InspExternalNumberingOfValues:  data.InspExternalNumberingOfValues,
			CreationDate:                   data.CreationDate,
			LastChangeDate:                 data.LastChangeDate,
			ChangeNumber:                   data.ChangeNumber,
			ValidityStartDate:              data.ValidityStartDate,
			ValidityEndDate:                data.ValidityEndDate,
			ChangedDateTime:                data.ChangedDateTime,
		})
	}

	return header, nil
}

func ConvertToMaterialAssignment(raw []byte, l *logger.Logger) ([]MaterialAssignment, error) {
	pm := &responses.MaterialAssignment{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to MaterialAssignment. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	materialAssignment := make([]MaterialAssignment, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		materialAssignment = append(materialAssignment, MaterialAssignment{
			Material:                      data.Material,
			Plant:                         data.Plant,
			InspectionPlanGroup:           data.InspectionPlanGroup,
			InspectionPlan:                data.InspectionPlan,
			InspPlanMatlAssignment:        data.InspPlanMatlAssignment,
			InspPlanMatlAssgmtIntVersion:  data.InspPlanMatlAssgmtIntVersion,
			InspectionPlanInternalVersion: data.InspectionPlanInternalVersion,
			ValidityStartDate:             data.ValidityStartDate,
			ValidityEndDate:               data.ValidityEndDate,
			ChangeNumber:                  data.ChangeNumber,
			CreationDate:                  data.CreationDate,
			LastChangeDate:                data.LastChangeDate,
			IsDeleted:                     data.IsDeleted,
			Supplier:                      data.Supplier,
			Customer:                      data.Customer,
			MultipleSpecificationObject:   data.MultipleSpecificationObject,
			MultipleSpecificationObjType:  data.MultipleSpecificationObjType,
			BOOSearchText:                 data.BOOSearchText,
			ChangedDateTime:               data.ChangedDateTime,
		})
	}

	return materialAssignment, nil
}

func ConvertToOperation(raw []byte, l *logger.Logger) ([]Operation, error) {
	pm := &responses.Operation{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Operation. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	operation := make([]Operation, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		operation = append(operation, Operation{
			InspectionPlanGroup:            data.InspectionPlanGroup,
			BOOOperationInternalID:         data.BOOOperationInternalID,
			BOOCharacteristic:              data.BOOCharacteristic,
			BOOCharacteristicVersion:       data.BOOCharacteristicVersion,
			BOOOpInternalVersionCounter:    data.BOOOpInternalVersionCounter,
			InspectionPlanInternalVersion:  data.InspectionPlanInternalVersion,
			InspectionPlan:                 data.InspectionPlan,
			ValidityStartDate:              data.ValidityStartDate,
			ValidityEndDate:                data.ValidityEndDate,
			ChangeNumber:                   data.ChangeNumber,
			IsDeleted:                      data.IsDeleted,
			BOOOperationPRTInternalID:      data.BOOOperationPRTInternalID,
			InspectionMethod:               data.InspectionMethod,
			InspectionMethodVersion:        data.InspectionMethodVersion,
			InspectionMethodPlant:          data.InspectionMethodPlant,
			InspSpecImportanceCode:         data.InspSpecImportanceCode,
			InspectorQualification:         data.InspectorQualification,
			InspectionSpecification:        data.InspectionSpecification,
			InspectionSpecificationVersion: data.InspectionSpecificationVersion,
			InspectionSpecificationPlant:   data.InspectionSpecificationPlant,
			BOOCharcHasInspSpecReference:   data.BOOCharcHasInspSpecReference,
			ProdnRsceToolItemNumber:        data.ProdnRsceToolItemNumber,
			InspSpecControlIndicators:      data.InspSpecControlIndicators,
			InspSpecIsQuantitative:         data.InspSpecIsQuantitative,
			InspSpecIsMeasuredValueRqd:     data.InspSpecIsMeasuredValueRqd,
			InspSpecIsSelectedSetRequired:  data.InspSpecIsSelectedSetRequired,
			InspSpecIsUpperLimitRequired:   data.InspSpecIsUpperLimitRequired,
			InspSpecIsLowerLimitRequired:   data.InspSpecIsLowerLimitRequired,
			InspSpecIsTargetValueInLimit:   data.InspSpecIsTargetValueInLimit,
			InspectionScope:                data.InspectionScope,
			InspSpecIsLongTermInspection:   data.InspSpecIsLongTermInspection,
			InspSpecRecordingType:          data.InspSpecRecordingType,
			InspResultIsDocumentationRqd:   data.InspResultIsDocumentationRqd,
			InspSpecCharcCategory:          data.InspSpecCharcCategory,
			InspSpecIsSampleQtyAdditive:    data.InspSpecIsSampleQtyAdditive,
			InspSpecIsDestructive:          data.InspSpecIsDestructive,
			InspSpecResultCalculation:      data.InspSpecResultCalculation,
			InspSpecIsSamplingProcedRqd:    data.InspSpecIsSamplingProcedRqd,
			InspSpecIsScrapRelevant:        data.InspSpecIsScrapRelevant,
			InspSpecHasFixedCtrlIndicators: data.InspSpecHasFixedCtrlIndicators,
			InspSpecIsTestEquipmentRqd:     data.InspSpecIsTestEquipmentRqd,
			InspSpecIsDefectRecordingRqd:   data.InspSpecIsDefectRecordingRqd,
			InspSpecIsDefectsRecgAutomatic: data.InspSpecIsDefectsRecgAutomatic,
			InspSpecIsChgDocRequired:       data.InspSpecIsChgDocRequired,
			InspSpecIsControlChartUsed:     data.InspSpecIsControlChartUsed,
			InspSpecPrintControl:           data.InspSpecPrintControl,
			InspSpecFirstUpperSpecLimit:    data.InspSpecFirstUpperSpecLimit,
			InspSpecHasFirstUpperSpecLimit: data.InspSpecHasFirstUpperSpecLimit,
			InspSpecFirstLowerSpecLimit:    data.InspSpecFirstLowerSpecLimit,
			InspSpecHasFirstLowerSpecLimit: data.InspSpecHasFirstLowerSpecLimit,
			InspSpecSecondUpperSpecLimit:   data.InspSpecSecondUpperSpecLimit,
			InspSpecHasSecondUprSpecLimit:  data.InspSpecHasSecondUprSpecLimit,
			InspSpecSecondLowerSpecLimit:   data.InspSpecSecondLowerSpecLimit,
			InspSpecHasSecondLowrSpecLimit: data.InspSpecHasSecondLowrSpecLimit,
			InspSpecInputProcedure:         data.InspSpecInputProcedure,
			InspSpecHasFormula:             data.InspSpecHasFormula,
			InspSpecFormula1:               data.InspSpecFormula1,
			InspSpecFormula2:               data.InspSpecFormula2,
			InspSpecNumberOfClasses:        data.InspSpecNumberOfClasses,
			InspSpecClassWidthQty:          data.InspSpecClassWidthQty,
			InspSpecHasClassWidth:          data.InspSpecHasClassWidth,
			InspSpecClassMidpointQty:       data.InspSpecClassMidpointQty,
			InspSpecHasClassMidpoint:       data.InspSpecHasClassMidpoint,
			InspToleranceSpecification:     data.InspToleranceSpecification,
			InspSpecDecimalPlaces:          data.InspSpecDecimalPlaces,
			InspectionSpecificationUnit:    data.InspectionSpecificationUnit,
			InspSpecTargetValue:            data.InspSpecTargetValue,
			InspSpecHasTargetValue:         data.InspSpecHasTargetValue,
			InspSpecUpperLimit:             data.InspSpecUpperLimit,
			InspSpecLowerLimit:             data.InspSpecLowerLimit,
			InspSpecHasLowerLimit:          data.InspSpecHasLowerLimit,
			InspSpecHasUpperLimit:          data.InspSpecHasUpperLimit,
			InspSpecDefectCodeGrpRejection: data.InspSpecDefectCodeGrpRejection,
			InspSpecDefectCodeRejection:    data.InspSpecDefectCodeRejection,
			InspSpecDefectCodeGrpRjcnUpper: data.InspSpecDefectCodeGrpRjcnUpper,
			InspSpecDefectCodeRjcnUpper:    data.InspSpecDefectCodeRjcnUpper,
			InspSpecDefectCodeGrpRjcnLower: data.InspSpecDefectCodeGrpRjcnLower,
			InspSpecDefectCodeRjcnLower:    data.InspSpecDefectCodeRjcnLower,
			SelectedCodeSet:                data.SelectedCodeSet,
			SelectedCodeSetPlant:           data.SelectedCodeSetPlant,
			InspSpecAdditionalCatalog2:     data.InspSpecAdditionalCatalog2,
			InspSpecAdditionalSelectedSet2: data.InspSpecAdditionalSelectedSet2,
			InspSpecAdditionalCodeGroup2:   data.InspSpecAdditionalCodeGroup2,
			InspSpecAddlSeldCodeSetPlant2:  data.InspSpecAddlSeldCodeSetPlant2,
			InspSpecAdditionalCatalog3:     data.InspSpecAdditionalCatalog3,
			InspSpecAdditionalSelectedSet3: data.InspSpecAdditionalSelectedSet3,
			InspSpecAdditionalCodeGroup3:   data.InspSpecAdditionalCodeGroup3,
			InspSpecAddlSeldCodeSetPlant3:  data.InspSpecAddlSeldCodeSetPlant3,
			InspSpecAdditionalCatalog4:     data.InspSpecAdditionalCatalog4,
			InspSpecAdditionalSelectedSet4: data.InspSpecAdditionalSelectedSet4,
			InspSpecAdditionalCodeGroup4:   data.InspSpecAdditionalCodeGroup4,
			InspSpecAddlSeldCodeSetPlant4:  data.InspSpecAddlSeldCodeSetPlant4,
			InspSpecAdditionalCatalog5:     data.InspSpecAdditionalCatalog5,
			InspSpecAdditionalSelectedSet5: data.InspSpecAdditionalSelectedSet5,
			InspSpecAdditionalCodeGroup5:   data.InspSpecAdditionalCodeGroup5,
			InspSpecAddlSeldCodeSetPlant5:  data.InspSpecAddlSeldCodeSetPlant5,
			SamplingProcedure:              data.SamplingProcedure,
			InspCharacteristicSampleUnit:   data.InspCharacteristicSampleUnit,
			BOOCharcSampleQuantity:         data.BOOCharcSampleQuantity,
			InspSpecInformationField1:      data.InspSpecInformationField1,
			InspSpecInformationField2:      data.InspSpecInformationField2,
			InspSpecInformationField3:      data.InspSpecInformationField3,
			InspectionSpecificationText:    data.InspectionSpecificationText,
			CreationDate:                   data.CreationDate,
			LastChangeDate:                 data.LastChangeDate,
			BillOfOperationsVersion:        data.BillOfOperationsVersion,
			ChangedDateTime:                data.ChangedDateTime,
		})
	}

	return operation, nil
}
