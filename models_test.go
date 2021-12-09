package godatagovgr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStringerOmitEmpty tests the strings produced for empty structs.
func TestStringerOmitEmpty(t *testing.T) {
	objects := []Stringable{
		&NumberData{},
		&AcademicProfessorData{},
		&AuditorData{},
		&CasinoTicketData{},
		&CrimeStatData{},
		&EudoxusApplicationData{},
		&FinancialCrimeData{},
		&HellenicCoastGuardIncidentData{},
		&InternshipData{},
		&InternetTrafficData{},
		&StudentSchoolData{},
		&TrafficAccidentData{},
		&TrafficViolationData{},
		&VaccinationData{},
	}

	for _, object := range objects {
		assert.Equal(t, "{}", object.String())
	}

}
