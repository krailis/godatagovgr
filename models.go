package godatagovgr

import (
	"encoding/json"
	"time"
)

// AcademicProfessorData struct representation.
// See more: https://www.data.gov.gr/datasets/minedu_dep/
type AcademicProfessorData struct {
	Year                int    `json:"year,omitempty"`
	Institution         string `json:"institution,omitempty"`
	FullProfessors      int    `json:"full_professors,omitempty"`
	AssociateProfessors int    `json:"associate_professors,omitempty"`
	AssistantProfessors int    `json:"assistant_professors,omitempty"`
	Lecturers           int    `json:"lecturers,omitempty"`
	PracticeProfessors  int    `json:"practice_professors,omitempty"`
	PracticeLecturers   int    `json:"practice_lecturers,omitempty"`
}

// AuditorData struct representation.
// See more: https://www.data.gov.gr/datasets/elte_auditors/
type AuditorData struct {
	NumberData
	Firms int `json:"firms,omitempty"`
}

// CadastrePlotData struct representation.
// This struct applies to multiple series of data:
//
// - NATURA Plots.
// See more: https://www.data.gov.gr/datasets/cadastre_natura_plot/
//
// - Plots.
// See more: https://www.data.gov.gr/datasets/cadastre_plot/
type CadastrePlotData struct {
	AverageArea      float64    `json:"avg_area,omitempty"`
	Date             *time.Time `json:"date,omitempty"`
	LocalAuthorityID string     `json:"local_authority_id,omitempty"`
	PlotNumber       int        `json:"plot_number,omitempty"`
}

// CasinoTicketData struct representation.
// See more: https://www.data.gov.gr/datasets/eeep_casino_tickets/
type CasinoTicketData struct {
	Year    int `json:"year,omitempty"`
	Tickets int `json:"tickets,omitempty"`
}

// CrimeStatData struct representation.
// See more: https://data.gov.gr/datasets/mcp_crime
type CrimeStatData struct {
	Year              int    `json:"year,omitempty"`
	Crime             string `json:"crime,omitempty"`
	Committed         int    `json:"committed,omitempty"`
	Attempted         int    `json:"attempted,omitempty"`
	Solved            int    `json:"solved,omitempty"`
	DomesticCriminals int    `json:"domestic_criminals,omitempty"`
	ForeignCriminals  int    `json:"foreign_criminals,omitempty"`
}

// ElectricityConsumptionData struct representation.
// See more: https://www.data.gov.gr/datasets/electricity_consumption/
type ElectricityConsumptionData struct {
	Area      string     `json:"area,omitempty"`
	Date      *time.Time `json:"date,omitempty"`
	EnergyMWH float64    `json:"energy_mwh,omitempty"`
}

// EnergyBalanceData struct representation.
// See more: https://www.data.gov.gr/datasets/admie_dailyenergybalanceanalysis/
type EnergyBalanceData struct {
	EnergyMWH  int        `json:"energy_mwh,omitempty"`
	Percentage float32    `json:"percentage,omitempty"`
	Date       *time.Time `json:"date,omitempty"`
	Fuel       string     `json:"fuel,omitempty"`
}

// EudoxusApplicationData struct representation.
// See more: https://www.data.gov.gr/datasets/grnet_eudoxus/
type EudoxusApplicationData struct {
	Year                  int    `json:"year,omitempty"`
	Period                string `json:"period,omitempty"`
	Institution           string `json:"institution,omitempty"`
	Department            string `json:"department,omitempty"`
	StudentWithStatements int    `json:"studentwithstatements,omitempty"`
	StudentWithDeliveries int    `json:"studentwithdeliveries,omitempty"`
}

// FinancialCrimeData struct representation.
// See more: https://www.data.gov.gr/datasets/mcp_financial_crimes/
type FinancialCrimeData struct {
	Year  int    `json:"year,omitempty"`
	Crime string `json:"crime,omitempty"`
	Count int    `json:"count,omitempty"`
}

// FoodInspectionData struct representation.
// See more: https://www.data.gov.gr/datasets/efet_inspections/
type FoodInspectionData struct {
	Year                   int     `json:"year,omitempty"`
	Inspections            int     `json:"inspections,omitempty"`
	Violations             int     `json:"violations,omitempty"`
	ViolatingOrganizations float64 `json:"violating_organizations,omitempty"`
	Penalties              float64 `json:"penalties,omitempty"`
}

// ForestFireData struct representation.
// See more: https://www.data.gov.gr/datasets/mcp_forest_fires/
type ForestFireData struct {
	Address                  string  `json:"address,omitempty"`
	AgriculturalAreaBurned   float64 `json:"agricultural_area_burned,omitempty"`
	AirplanesCL215           int     `json:"airplanes_cl215,omitempty"`
	AirplanesCL415           int     `json:"airplanes_cl415,omitempty"`
	AirplanesGRU             int     `json:"airplanes_gru,omitempty"`
	AirplanesPZL             int     `json:"airplanes_pzl,omitempty"`
	Army                     int     `json:"army,omitempty"`
	CropResidueAreaBurned    float64 `json:"crop_residue_area_burned,omitempty"`
	DumpingGroundAreaBurned  float64 `json:"dumping_ground_area_burned,omitempty"`
	EndTime                  string  `json:"end_time,omitempty"`
	Firefighters             int     `json:"firefighters,omitempty"`
	FireStation              string  `json:"fire_station,omitempty"`
	FireTrucks               int     `json:"fire_trucks,omitempty"`
	ForestAreaBurned         float64 `json:"forest_area_burned,omitempty"`
	Forestry                 *string `json:"forestry,omitempty"`
	GroveAreaBurned          float64 `json:"grove_area_burned,omitempty"`
	Helicopters              int     `json:"helicopters,omitempty"`
	LocalAuthoritiesVehicles int     `json:"local_authorities_vehicles,omitempty"`
	Location                 *string `json:"location,omitempty"`
	LowVegetationAreaBurned  float64 `json:"low_vegetation_area_burned,omitempty"`
	Machinery                int     `json:"machinery,omitempty"`
	Municipality             string  `json:"municipality,omitempty"`
	OtherFirefighters        int     `json:"other_firefighters,omitempty"`
	Prefecture               string  `json:"prefecture,omitempty"`
	StartTime                string  `json:"start_time,omitempty"`
	SwampAreaBurned          float64 `json:"swamp_area_burned,omitempty"`
	Volunteers               int     `json:"volunteers,omitempty"`
	WaterTankTrucks          int     `json:"water_tank_trucks,omitempty"`
	WildlandCrew             int     `json:"wildland_crew,omitempty"`
	WoodlandAreaBurned       float64 `json:"woodland_area_burned,omitempty"`
}

// HellenicCoastGuardIncidentData struct representation.
// See more: https://www.data.gov.gr/datasets/hcg_incidents/
type HellenicCoastGuardIncidentData struct {
	Year          int    `json:"year,omitempty"`
	Incident      string `json:"incident,omitempty"`
	Domestic      int    `json:"domestic,omitempty"`
	International int    `json:"international,omitempty"`
}

// InternetTrafficData struct representation.
// See more: https://data.gov.gr/datasets/internet_traffic/
type InternetTrafficData struct {
	Date       *time.Time `json:"date,omitempty"`
	AverageIn  int        `json:"avg_in,omitempty"`
	AverageOut int        `json:"avg_out,omitempty"`
	MaxIn      int        `json:"max_in,omitempty"`
	MaxOut     int        `json:"max_out,omitempty"`
}

// InternshipData struct representation.
// See more: https://www.data.gov.gr/datasets/grnet_atlas/
type InternshipData struct {
	Year                int    `json:"year,omitempty"`
	Institution         string `json:"institution,omitempty"`
	PrivateSector       int    `json:"private_sector,omitempty"`
	PublicSector        int    `json:"public_sector,omitempty"`
	NonGovOrganizations int    `json:"ngo,omitempty"`
}

// LandRegistryBase struct representation.
// It serves as the base of data structs related to Land Registry.
type LandRegistryBase struct {
	Date        string `json:"date,omitempty"`
	OtaFullName string `json:"ota_full_name,omitempty"`
	OtaID       int    `json:"ota_id,omitempty"`
	OtaName     string `json:"ota_name,omitempty"`
	OtaNameEn   string `json:"ota_name_en,omitempty"`
}

// LandConfiscationData struct representation.
// See more: https://www.data.gov.gr/datasets/ktm_confs/
type LandConfiscationData struct {
	LandRegistryBase
	Confiscations int `json:"confiscations,omitempty"`
}

// LandHorizontalAreaData struct representation.
// See more: https://www.data.gov.gr/datasets/ktm_harea/
type LandHorizontalAreaData struct {
	LandRegistryBase
	Area float64 `json:"area,omitempty"`
}

// LandHorizontalPlotData struct representation.
// See more: https://www.data.gov.gr/datasets/ktm_hplots/
type LandHorizontalPlotData struct {
	LandRegistryBase
	Plots int `json:"plots,omitempty"`
}

// LandLienData struct representation.
// See more: https://www.data.gov.gr/datasets/ktm_liens/
type LandLienData struct {
	LandRegistryBase
	Liens int `json:"liens,omitempty"`
}

// LandOwnerData struct representation.
// See more: https://www.data.gov.gr/datasets/ktm_owners/
type LandOwnerData struct {
	LandRegistryBase
	Owners int `json:"owners,omitempty"`
}

// LandPlotData struct representation.
// See more: https://www.data.gov.gr/datasets/ktm_plots/
type LandPlotData struct {
	LandRegistryBase
	Plots int `json:"plots,omitempty"`
}

// LandRegistryStatusData struct representation.
// See more: https://www.data.gov.gr/datasets/ktm_status/
type LandRegistryStatusData struct {
	LandRegistryBase
	Status string `json:"status,omitempty"`
}

// LandTransactionData struct representation.
// See more: https://www.data.gov.gr/datasets/ktm_transactions/
type LandTransactionData struct {
	LandRegistryBase
	Transactions int `json:"transactions,omitempty"`
}

// NumberData struct representation.
// This struct applies to multiple series of data:
//
// - Accountants.
// See more: https://www.data.gov.gr/datasets/oee_accountants/
//
// - Doctors.
// See more: https://www.data.gov.gr/datasets/minhealth_doctors/
//
// - Dendists.
// See more: https://www.data.gov.gr/datasets/minhealth_dentists/
//
// - Energy Inspector numbers.
// See more: https://www.data.gov.gr/datasets/minenv_inspectors/
//
// - Lawyer numbers.
// See more: https://www.data.gov.gr/datasets/minjust_lawyers/
//
// - Law firm numbers.
// See more: https://www.data.gov.gr/datasets/minjust_law_firms/
//
// - Pharmacy numbers.
// See more: https://data.gov.gr/datasets/minhealth_pharmacies/
//
// - Pharmacist numbers.
// See more: https://data.gov.gr/datasets/minhealth_pharmacists/
//
// - Realtors.
// See more: https://www.data.gov.gr/datasets/mindev_realtors/
//
// - Tourist Agencies.
// See more: https://www.data.gov.gr/datasets/mintour_agencies/
type NumberData struct {
	Year     int    `json:"year,omitempty"`
	Quarter  string `json:"quarter,omitempty"`
	Active   int    `json:"active,omitempty"`
	Entrants int    `json:"entrants,omitempty"`
	Exits    int    `json:"exits,omitempty"`
}

// PowerSystemLoadData struct representation.
// See more: https://www.data.gov.gr/datasets/admie_realtimescadasystemload/
type PowerSystemLoadData struct {
	Date      *time.Time `json:"date,omitempty"`
	EnergyMWH int        `json:"energy_mwh,omitempty"`
}

// ResProductionData struct representation.
// See more: https://www.data.gov.gr/datasets/admie_realtimescadares/
type ResProductionData struct {
	Date      *time.Time `json:"date,omitempty"`
	EnergyMWH int        `json:"energy_mwh,omitempty"`
}

// RidershipData (OASA) struct representation.
// See more: https://www.data.gov.gr/datasets/oasa_ridership/
type RidershipData struct {
	Validations     int        `json:"dv_validations,omitempty"`
	Agency          string     `json:"dv_agency,omitempty"`
	PlatenumStation string     `json:"dv_platenum_station,omitempty"`
	RoutesPerHour   *int       `json:"routes_per_hour,omitempty"`
	Date            *time.Time `json:"load_dt,omitempty"`
	DateHour        string     `json:"date_hour,omitempty"`
}

// RoadTrafficData struct representation.
// See more: https://www.data.gov.gr/datasets/road_traffic_attica/
type RoadTrafficData struct {
	DeviveID       string  `json:"year,omitempty"`
	CountedCars    int     `json:"countedcars,omitempty"`
	AppProcessTIme string  `json:"appprocesstime,omitempty"`
	RoadName       string  `json:"road_name,omitempty"`
	RoadInfo       string  `json:"road_info,omitempty"`
	AverageSpeed   float64 `json:"average_speed,omitempty"`
}

// SailingTrafficData struct representation.
// See more: https://www.data.gov.gr/datasets/minedu_students_school/
type SailingTrafficData struct {
	ArrivalPort       string `json:"arrivalport,omitempty"`
	ArrivalPortName   string `json:"arrivalportname,omitempty"`
	Date              string `json:"date,omitempty"`
	DeparturePort     string `json:"departureport,omitempty"`
	DeparturePortName string `json:"departureportname,omitempty"`
	PassengerCount    int    `json:"passengercount,omitempty"`
	RouteCode         string `json:"routecode,omitempty"`
	RouteCodeName     string `json:"routecodename,omitempty"`
	VehicleCount      int    `json:"vehiclecount,omitempty"`
}

// StudentSchoolData struct representation.
// See more: https://www.data.gov.gr/datasets/minedu_students_school/
type StudentSchoolData struct {
	Year                    int    `json:"year,omitempty"`
	Jurisdiction            string `json:"jurisdiction,omitempty"`
	District                string `json:"district,omitempty"`
	SchoolClass             string `json:"school_class,omitempty"`
	SchoolType              string `json:"school_type,omitempty"`
	SchoolName              string `json:"school_name,omitempty"`
	RegisteredStudentsBoys  int    `json:"registered_students_boys,omitempty"`
	RegisteredStudentsGirls int    `json:"registered_students_girls,omitempty"`
}

// TelecomIndicatorData struct definition.
// See more: https://www.data.gov.gr/datasets/eett_telecom_indicators/
type TelecomIndicatorData struct {
	Year      int     `json:"year,omitempty"`
	Category  string  `json:"category,omitempty"`
	Indicator string  `json:"indicator,omitempty"`
	Value     float64 `json:"value,omitempty"`
}

// TrafficAccidentData struct representation.
// See more: https://data.gov.gr/datasets/mcp_traffic_accidents
type TrafficAccidentData struct {
	Year             int    `json:"year,omitempty"`
	Jurisdiction     string `json:"jurisdiction,omitempty"`
	DeadlyAccidents  int    `json:"deadly_accidents,omitempty"`
	SeriousAccidents int    `json:"serious_accidents,omitempty"`
	OtherAccidents   int    `json:"other_accidents,omitempty"`
	Deaths           int    `json:"deaths,omitempty"`
	SeriouslyInjured int    `json:"seriously_injured,omitempty"`
	OtherInjured     int    `json:"other_injured,omitempty"`
}

// TrafficViolationData struct representation.
// See more: https://www.data.gov.gr/datasets/mcp_traffic_violations/
type TrafficViolationData struct {
	Year      int    `json:"year,omitempty"`
	Violation string `json:"violation,omitempty"`
	Count     int    `json:"count,omitempty"`
}

// UnemploymentData struct representation.
// See more: https://www.data.gov.gr/datasets/oaed_unemployment/
type UnemploymentData struct {
	Date       *time.Time `json:"asofdate,omitempty"`
	Unemployed int        `json:"unemployed,omitempty"`
	Benefits   int        `json:"benefits,omitempty"`
}

// UrbanIncidentData struct representation.
// See more: https://www.data.gov.gr/datasets/mcp_urban_incidents/
type UrbanIncidentData struct {
	AccidentType        *string `json:"accident_type,omitempty"`
	Burns               int     `json:"burns,omitempty"`
	Damages             int     `json:"damages,omitempty"`
	Deaths              int     `json:"deaths,omitempty"`
	EndTime             *string `json:"end_time,omitempty"`
	Firefighters        int     `json:"fire_fighters,omitempty"`
	FireStation         *string `json:"fire_station,omitempty"`
	FireTrucks          int     `json:"fire_trucks,omitempty"`
	FireVessels         int     `json:"fire_vessels,omitempty"`
	Incident            *string `json:"incident,omitempty"`
	IncidentDetail      *string `json:"incident_detail,omitempty"`
	Injuries            int     `json:"injuries,omitempty"`
	LocationDescription *string `json:"location_description,omitempty"`
	Municipality        *string `json:"municipality,omitempty"`
	PeopleInvolved      int     `json:"people_involved,omitempty"`
	Prefecture          *string `json:"prefecture,omitempty"`
	StartTime           *string `json:"start_time,omitempty"`
	Village             *string `json:"village	,omitempty"`
}

// VaccinationData struct representation.
// See more: https://data.gov.gr/datasets/mdg_emvolio/
type VaccinationData struct {
	Area                 string `json:"area,omitempty"`
	AreaID               int    `json:"areaid,omitempty"`
	DailyDoseFirst       int    `json:"dailydose1,omitempty"`
	DailyDoseSecond      int    `json:"dailydose2,omitempty"`
	DailyDoseThird       int    `json:"dailydose3,omitempty"`
	DayDiff              int    `json:"daydiff,omitempty"`
	DayTotal             int    `json:"daytotal,omitempty"`
	ReferenceDate        string `json:"referencedate,omitempty"`
	TotalDistinctPersons int    `json:"totaldistinctpersons,omitempty"`
	TotalDoseFirst       int    `json:"totaldose1,omitempty"`
	TotalDoseSecond      int    `json:"totaldose2,omitempty"`
	TotalDoseThird       int    `json:"totaldose3,omitempty"`
	TotalVaccinations    int    `json:"totalvaccinations,omitempty"`
}

// VotersPerAgeData struct representation.
// See more: https://www.data.gov.gr/datasets/minint_election_age/
type VotersPerAgeData struct {
	Year              int    `json:"year,omitempty"`
	ElectionType      string `json:"election_type,omitempty"`
	ElectoralDistrict string `json:"electoral_district,omitempty"`
	AgeGroup          string `json:"age_group,omitempty"`
	Count             int    `json:"count,omitempty"`
}

// VotersPerMunicipalityData struct representation.
// See more: https://www.data.gov.gr/datasets/minint_election_distribution/
type VotersPerMunicipalityData struct {
	Year              int    `json:"year,omitempty"`
	ElectionType      string `json:"election_type,omitempty"`
	ElectoralDistrict string `json:"electoral_district,omitempty"`
	Municipality      string `json:"municipality,omitempty"`
	VotersMale        int    `json:"voters_male,omitempty"`
	VotersFemale      int    `json:"voters_female,omitempty"`
}

// Stringable interface definition.
type Stringable interface {
	String() string
}

// structToJSONString returns a struct formatted into a pretty JSON string.
func structToJSONString(t interface{}) string {
	json, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		return ""
	}
	return string(json)
}

// String returns a pretty JSON string representation.
func (i *AcademicProfessorData) String() string          { return structToJSONString(i) }
func (i *AuditorData) String() string                    { return structToJSONString(i) }
func (i *CadastrePlotData) String() string               { return structToJSONString(i) }
func (i *CasinoTicketData) String() string               { return structToJSONString(i) }
func (i *CrimeStatData) String() string                  { return structToJSONString(i) }
func (i *ElectricityConsumptionData) String() string     { return structToJSONString(i) }
func (i *EnergyBalanceData) String() string              { return structToJSONString(i) }
func (i *EudoxusApplicationData) String() string         { return structToJSONString(i) }
func (i *FinancialCrimeData) String() string             { return structToJSONString(i) }
func (i *FoodInspectionData) String() string             { return structToJSONString(i) }
func (i *ForestFireData) String() string                 { return structToJSONString(i) }
func (i *HellenicCoastGuardIncidentData) String() string { return structToJSONString(i) }
func (i *InternetTrafficData) String() string            { return structToJSONString(i) }
func (i *InternshipData) String() string                 { return structToJSONString(i) }
func (i *LandConfiscationData) String() string           { return structToJSONString(i) }
func (i *LandHorizontalAreaData) String() string         { return structToJSONString(i) }
func (i *LandHorizontalPlotData) String() string         { return structToJSONString(i) }
func (i *LandLienData) String() string                   { return structToJSONString(i) }
func (i *LandOwnerData) String() string                  { return structToJSONString(i) }
func (i *LandPlotData) String() string                   { return structToJSONString(i) }
func (i *LandRegistryStatusData) String() string         { return structToJSONString(i) }
func (i *LandTransactionData) String() string            { return structToJSONString(i) }
func (i *NumberData) String() string                     { return structToJSONString(i) }
func (i *PowerSystemLoadData) String() string            { return structToJSONString(i) }
func (i *ResProductionData) String() string              { return structToJSONString(i) }
func (i *RidershipData) String() string                  { return structToJSONString(i) }
func (i *RoadTrafficData) String() string                { return structToJSONString(i) }
func (i *SailingTrafficData) String() string             { return structToJSONString(i) }
func (i *StudentSchoolData) String() string              { return structToJSONString(i) }
func (i *TelecomIndicatorData) String() string           { return structToJSONString(i) }
func (i *TrafficAccidentData) String() string            { return structToJSONString(i) }
func (i *TrafficViolationData) String() string           { return structToJSONString(i) }
func (i *UnemploymentData) String() string               { return structToJSONString(i) }
func (i *UrbanIncidentData) String() string              { return structToJSONString(i) }
func (i *VaccinationData) String() string                { return structToJSONString(i) }
func (i *VotersPerAgeData) String() string               { return structToJSONString(i) }
func (i *VotersPerMunicipalityData) String() string      { return structToJSONString(i) }
