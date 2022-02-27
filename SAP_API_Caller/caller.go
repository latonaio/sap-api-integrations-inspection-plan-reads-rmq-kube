package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-inspection-plan-reads-rmq-kube/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL      string
	apiKey       string
	outputQueues []string
	outputter    RMQOutputter
	log          *logger.Logger
}

func NewSAPAPICaller(baseUrl string, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:      baseUrl,
		apiKey:       GetApiKey(),
		outputQueues: outputQueueTo,
		outputter:    outputter,
		log:          l,
	}
}

func (c *SAPAPICaller) AsyncGetInspectionPlan(inspectionPlanGroup, inspectionPlan, plant, material, billOfOperationsDesc, inspectionSpecification string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(inspectionPlanGroup, inspectionPlan)
				wg.Done()
			}()
		case "MaterialAssignment":
			func() {
				c.MaterialAssignment(plant, material)
				wg.Done()
			}()
		case "Operation":
			func() {
				c.Operation(inspectionPlanGroup, inspectionPlan)
				wg.Done()
			}()
		case "BillOfOperationsDesc":
			func() {
				c.BillOfOperationsDesc(plant, billOfOperationsDesc)
				wg.Done()
			}()
		case "InspectionSpecification":
			func() {
				c.InspectionSpecification(plant, inspectionSpecification)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(inspectionPlanGroup, inspectionPlan string) {
	data, err := c.callInspectionPlanSrvAPIRequirementHeader("A_InspectionPlan", inspectionPlanGroup, inspectionPlan)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "InspectionPlanHeader"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callInspectionPlanSrvAPIRequirementHeader(api, inspectionPlanGroup, inspectionPlan string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_INSPECTIONPLAN_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, inspectionPlanGroup, inspectionPlan)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) MaterialAssignment(plant, material string) {
	data, err := c.callInspectionPlanSrvAPIRequirementMaterialAssignment("A_InspPlanMaterialAssgmt", plant, material)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "InspectionPlanMaterialAssignment"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callInspectionPlanSrvAPIRequirementMaterialAssignment(api, plant, material string) ([]sap_api_output_formatter.MaterialAssignment, error) {
	url := strings.Join([]string{c.baseURL, "API_INSPECTIONPLAN_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithMaterialAssignment(req, plant, material)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToMaterialAssignment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Operation(inspectionPlanGroup, inspectionPlan string) {
	data, err := c.callInspectionPlanSrvAPIRequirementOperation("A_InspPlanOpCharacteristic", inspectionPlanGroup, inspectionPlan)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "InspectionPlanOperation"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callInspectionPlanSrvAPIRequirementOperation(api, inspectionPlanGroup, inspectionPlan string) ([]sap_api_output_formatter.Operation, error) {
	url := strings.Join([]string{c.baseURL, "API_INSPECTIONPLAN_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithOperation(req, inspectionPlanGroup, inspectionPlan)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToOperation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) BillOfOperationsDesc(plant, billOfOperationsDesc string) {
	data, err := c.callInspectionPlanSrvAPIRequirementBillOfOperationsDesc("A_InspectionPlan", plant, billOfOperationsDesc)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "InspectionPlanHeader"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callInspectionPlanSrvAPIRequirementBillOfOperationsDesc(api, plant, billOfOperationsDesc string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_INSPECTIONPLAN_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithBillOfOperationsDesc(req, plant, billOfOperationsDesc)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) InspectionSpecification(plant, inspectionSpecification string) {
	data, err := c.callInspectionPlanSrvAPIRequirementInspectionSpecification("A_InspPlanOpCharacteristic", plant, inspectionSpecification)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "InspectionPlanOperation"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callInspectionPlanSrvAPIRequirementInspectionSpecification(api, plant, inspectionSpecification string) ([]sap_api_output_formatter.Operation, error) {
	url := strings.Join([]string{c.baseURL, "API_INSPECTIONPLAN_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithInspectionSpecification(req, plant, inspectionSpecification)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToOperation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, inspectionPlanGroup, inspectionPlan string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InspectionPlanGroup eq '%s' and InspectionPlan eq '%s'", inspectionPlanGroup, inspectionPlan))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithMaterialAssignment(req *http.Request, plant, material string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Plant eq '%s' and Material eq '%s'", plant, material))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithOperation(req *http.Request, inspectionPlanGroup, inspectionPlan string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InspectionPlanGroup eq '%s' and InspectionPlan eq '%s'", inspectionPlanGroup, inspectionPlan))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithBillOfOperationsDesc(req *http.Request, plant, billOfOperationsDesc string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Plant eq '%s' and substringof('%s', BillOfOperationsDesc)", plant, billOfOperationsDesc))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithInspectionSpecification(req *http.Request, plant, inspectionSpecification string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InspectionSpecificationPlant eq '%s' and substringof('%s', InspectionSpecification)", plant, inspectionSpecification))
	req.URL.RawQuery = params.Encode()
}
