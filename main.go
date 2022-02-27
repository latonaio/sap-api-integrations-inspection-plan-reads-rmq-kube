package main

import (
	sap_api_caller "sap-api-integrations-inspection-plan-reads-rmq-kube/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-inspection-plan-reads-rmq-kube/SAP_API_Input_Reader"
	"sap-api-integrations-inspection-plan-reads-rmq-kube/config"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client"
	"golang.org/x/xerrors"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	rmq, err := rabbitmq.NewRabbitmqClient(conf.RMQ.URL(), conf.RMQ.QueueFrom(), conf.RMQ.QueueTo())
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Close()

	caller := sap_api_caller.NewSAPAPICaller(
		conf.SAP.BaseURL(),
		conf.RMQ.QueueTo(),
		rmq,
		l,
	)

	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()

	for msg := range iter {
		err = callInspectionPlan(caller, msg)
		if err != nil {
			msg.Fail()
			l.Error(err)
			continue
		}
		msg.Success()
	}
}

func callInspectionPlan(caller *sap_api_caller.SAPAPICaller, msg rabbitmq.RabbitmqMessage) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = xerrors.Errorf("error occurred: %w", e)
			return
		}
	}()
	inspectionPlanGroup, inspectionPlan, plant, material, billOfOperationsDesc, inspectionSpecification := extractData(msg.Data())
	accepter := getAccepter(msg.Data())
	caller.AsyncGetInspectionPlan(inspectionPlanGroup, inspectionPlan, plant, material, billOfOperationsDesc, inspectionSpecification, accepter)
	return nil
}

func extractData(data map[string]interface{}) (inspectionPlanGroup, inspectionPlan, plant, material, billOfOperationsDesc, inspectionSpecification string) {
	sdc := sap_api_input_reader.ConvertToSDC(data)
	inspectionPlanGroup = sdc.InspectionPlan.InspectionPlanGroup
	inspectionPlan = sdc.InspectionPlan.InspectionPlan
	plant = sdc.InspectionPlan.Plant
	material = sdc.InspectionPlan.MaterialAssignment.Material
	billOfOperationsDesc = sdc.InspectionPlan.BillOfOperationsDesc
	inspectionSpecification = sdc.InspectionPlan.Operation.InspectionSpecification
	return
}

func getAccepter(data map[string]interface{}) []string {
	sdc := sap_api_input_reader.ConvertToSDC(data)
	accepter := sdc.Accepter
	if len(sdc.Accepter) == 0 {
		accepter = []string{"All"}
	}

	if accepter[0] == "All" {
		accepter = []string{
			"CustomerMaterial",
		}
	}
	return accepter
}

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Inspection_Plan_Operation_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "MaterialAssignment", "Operation",
			"BillOfOperationsDesc", "InspectionSpecification",
		}
	}

	caller.AsyncGetInspectionPlan(
		inoutSDC.InspectionPlan.InspectionPlanGroup,
		inoutSDC.InspectionPlan.InspectionPlan,
		inoutSDC.InspectionPlan.Plant,
		inoutSDC.InspectionPlan.MaterialAssignment.Material,
		inoutSDC.InspectionPlan.BillOfOperationsDesc,
		inoutSDC.InspectionPlan.Operation.InspectionSpecification,
		accepter,
	)
}