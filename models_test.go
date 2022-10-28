package godatagovgr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringer(t *testing.T) {
	// Test for number data.
	numberData := &NumberData{
		Year:     2020,
		Quarter:  "Q1",
		Active:   73768,
		Entrants: 8768,
		Exits:    897,
	}
	assert.Equal(t, `{
	"year": 2020,
	"quarter": "Q1",
	"active": 73768,
	"entrants": 8768,
	"exits": 897
}`,
		numberData.String(),
	)
	// Test for voters per age data.
	votersPerAge := &VotersPerAgeData{
		Year:              2020,
		AgeGroup:          "18-23",
		ElectionType:      "Βουλευτικές",
		ElectoralDistrict: "Καρδίτσας",
		Count:             35000,
	}
	assert.Equal(t, `{
	"year": 2020,
	"election_type": "Βουλευτικές",
	"electoral_district": "Καρδίτσας",
	"age_group": "18-23",
	"count": 35000
}`,
		votersPerAge.String(),
	)
}

func TestStringerOmitEmpty(t *testing.T) {
	objects := []Stringable{
		&AcademicProfessorData{},
		&AuditorData{},
		&CadastrePlotData{},
		&CasinoTicketData{},
		&CrimeStatData{},
		&ElectricityConsumptionData{},
		&EnergyBalanceData{},
		&EudoxusApplicationData{},
		&FinancialCrimeData{},
		&FoodInspectionData{},
		&ForestFireData{},
		&HellenicCoastGuardIncidentData{},
		&InternshipData{},
		&InternetTrafficData{},
		&LandConfiscationData{},
		&LandHorizontalAreaData{},
		&LandHorizontalPlotData{},
		&LandLienData{},
		&LandOwnerData{},
		&LandPlotData{},
		&LandRegistryStatusData{},
		&LandTransactionData{},
		&NumberData{},
		&PowerSystemLoadData{},
		&ResProductionData{},
		&RidershipData{},
		&RoadTrafficData{},
		&SailingTrafficData{},
		&StudentSchoolData{},
		&TelecomIndicatorData{},
		&TrafficAccidentData{},
		&TrafficViolationData{},
		&UnemploymentData{},
		&UrbanIncidentData{},
		&VaccinationData{},
		&VotersPerAgeData{},
		&VotersPerMunicipalityData{},
	}

	for _, object := range objects {
		assert.Equal(t, "{}", object.String())
	}

}
