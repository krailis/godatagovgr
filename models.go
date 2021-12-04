package godatagovgr

// Number Data struct definition.
//
// This struct applies to multiple series of data:
// - Accountants.
//   See more: https://www.data.gov.gr/datasets/oee_accountants/
// - Energy Inspector numbers.
//   See more: https://www.data.gov.gr/datasets/minenv_inspectors/
// - Lawyer numbers.
//   See more: https://www.data.gov.gr/datasets/minjust_lawyers/
// - Law firm numbers.
//   See more: https://www.data.gov.gr/datasets/minjust_law_firms/
// - Pharmacy numbers.
//   See more: https://data.gov.gr/datasets/minhealth_pharmacies/
// - Pharmacist numbers.
//   See more: https://data.gov.gr/datasets/minhealth_pharmacists/
// - Realtors.
//   See more: https://www.data.gov.gr/datasets/mindev_realtors/
// - Tourist Agencies.
//   See more: https://www.data.gov.gr/datasets/mintour_agencies/
type NumberData struct {
	Year     int    `json:"year,omitempty"`
	Quarter  string `json:"quarter,omitempty"`
	Active   int    `json:"active,omitempty"`
	Entrants int    `json:"entrants,omitempty"`
	Exits    int    `json:"exits,omitempty"`
}

// Auditors.
// See more: https://www.data.gov.gr/datasets/elte_auditors/
type AuditorData struct {
	NumberData
	Firms int `json:"firms,omitempty"`
}

// Casino Tickets.
// See more: https://www.data.gov.gr/datasets/eeep_casino_tickets/
type CasinoTicketData struct {
	Year    int `json:"year,omitempty"`
	Tickets int `json:"tickets,omitempty"`
}

// Crime Statistics.
// See more: https://data.gov.gr/api/v1/query/mcp_crime
type CrimeStatData struct {
	Year              int    `json:"year,omitempty"`
	Crime             string `json:"crime,omitempty"`
	Committed         int    `json:"committed,omitempty"`
	Attempted         int    `json:"attempted,omitempty"`
	Solved            int    `json:"solved,omitempty"`
	DomesticCriminals int    `json:"domestic_criminals,omitempty"`
	ForeignCriminals  int    `json:"foreign_criminals,omitempty"`
}

// Academic Professor (DEP) data.
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

// Eudoxus Applications & Deliveries.
// See more: https://www.data.gov.gr/datasets/grnet_eudoxus/
type EudoxusApplicationData struct {
	Year                  int    `json:"year,omitempty"`
	Period                string `json:"period,omitempty"`
	Institution           string `json:"institution,omitempty"`
	Department            string `json:"department,omitempty"`
	StudentWithStatements int    `json:"studentwithstatements,omitempty"`
	StudentWithDeliveries int    `json:"studentwithdeliveries,omitempty"`
}

// Financial Crime Data.
// See more: https://www.data.gov.gr/datasets/mcp_financial_crimes/
type FinancialCrimeData struct {
	Year  int    `json:"year,omitempty"`
	Crime string `json:"crime,omitempty"`
	Count int    `json:"count,omitempty"`
}

// Hellenic Coast Guard Incidents.
// See more: https://www.data.gov.gr/datasets/hcg_incidents/
type HellenicCoastGuardIncidents struct {
	Year          int    `json:"year,omitempty"`
	Incident      string `json:"incident,omitempty"`
	Domestic      int    `json:"domestic,omitempty"`
	International int    `json:"international,omitempty"`
}

// Internet Traffic.
// See more: https://data.gov.gr/datasets/internet_traffic/
type InternetTraffic struct {
	Date       string `json:"date"`
	AverageIn  int    `json:"avg_in"`
	AverageOut int    `json:"avg_out"`
	MaxIn      int    `json:"max_in"`
	MaxOut     int    `json:"max_out"`
}

// Internet Traffic Query Params.
type InternetTrafficQueryParams struct {
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
}

// Internship Data.
// See more: https://www.data.gov.gr/datasets/grnet_atlas/
type InternshipData struct {
	Year                int    `json:"year,omitempty"`
	Institution         string `json:"institution,omitempty"`
	PrivateSector       int    `json:"private_sector,omitempty"`
	PublicSector        int    `json:"public_sector,omitempty"`
	NonGovOrganizations int    `json:"ngo,omitempty"`
}

// Student School Data.
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

// Traffic Accidents.
// See more: https://data.gov.gr/api/v1/query/mcp_traffic_accidents
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

// Traffic Violations.
// See more: https://www.data.gov.gr/datasets/mcp_traffic_violations/
type TrafficViolationData struct {
	Year      int    `json:"year,omitempty"`
	Violation string `json:"violation,omitempty"`
	Count     int    `json:"count,omitempty"`
}

// Traffic Violation Query Params.
type TrafficViolationQueryParams struct {
}

// Vaccinations.
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

// Vaccination Query Parameters struct definition.
type VaccinationQueryParams struct {
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
	Area     string `json:"area,omitempty"`
}
