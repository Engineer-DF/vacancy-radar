package vacancy

type Vacancy struct {
	ID                    string
	Name                  string
	AreaID                string
	AreaName              string
	Salary                int
	SalaryFrequency       string
	SalaryMode            string
	PublishedAt           string
	CreatedAt             string
	AlternateURL          string
	Employer              string
	ProfessionalRolesID   string
	ProfessionalRolesName string
	Experience            string
}

// Вполне вероятно что почти половину фильтров придется удалить, но кого это волнует :D

type Filter struct {
	AgeRestriction             []DictionaryItem `json:"age_restriction"`
	ApplicantCommentAccessType []DictionaryItem `json:"applicant_comment_access_type"`
	ApplicantCommentsOrder     []DictionaryItem `json:"applicant_comments_order"`
	ApplicantNegotiationStatus []DictionaryItem `json:"applicant_negotiation_status"`
	BusinessTripReadiness      []DictionaryItem `json:"business_trip_readiness"`
	CivilLawContracts          []DictionaryItem `json:"civil_law_contracts"`
	Currency                   []CurrencyItem   `json:"currency"`
	DriverLicenseTypes         []struct {
		ID string
	} `json:"driver_license_types"`
	EducationLevel                 []DictionaryItem `json:"education_level"`
	EmployerActiveVacanciesOrder   []DictionaryItem `json:"employer_active_vacancies_order"`
	EmployerArchivedVacanciesOrder []DictionaryItem `json:"employer_archived_vacancies_order"`
	EmployerHiddenVacanciesOrder   []DictionaryItem `json:"employer_hidden_vacancies_order"`
	EmployerRelation               []DictionaryItem `json:"employer_relation"`
	EmployerType                   []DictionaryItem `json:"employer_type"`
	Employment                     []DictionaryItem `json:"employment"`
	EmploymentForm                 []struct {
		DictionaryItem
		Duration string
	} `json:"employment_form"`
	Experience          []DictionaryItem `json:"experience"`
	FlyInFlyOutDuration []DictionaryItem `json:"fly_in_fly_out_duration"`
	Gender              []DictionaryItem `json:"gender"`
	InclusivenessTypes  []struct {
		DictionaryItem
		Legacy bool
	} `json:"inclusiveness_types"`
	JobSearchStatusesApplicant   []DictionaryItem `json:"job_search_statuses_applicant"`
	JobSearchStatusesEmployer    []DictionaryItem `json:"job_search_statuses_employer"`
	LanguageLevel                []DictionaryItem `json:"language_level"`
	LinkedSocials                []DictionaryItem `json:"linked_socials"`
	MessagingStatus              []DictionaryItem `json:"messaging_status"`
	NegotiationsOrder            []DictionaryItem `json:"negotiations_order"`
	NegotiationsParticipantType  []DictionaryItem `json:"negotiations_participant_type"`
	NegotiationsState            []DictionaryItem `json:"negotiations_state"`
	PhoneCallStatus              []DictionaryItem `json:"phone_call_status"`
	PreferredContactType         []DictionaryItem `json:"preferred_contact_type"`
	RelocationType               []DictionaryItem `json:"relocation_type"`
	ResumeAccessType             []DictionaryItem `json:"resume_access_type"`
	ResumeContactsSiteType       []DictionaryItem `json:"resume_contacts_site_type"`
	ResumeEmploymentForm         []DictionaryItem `json:"resume_employment_form"`
	ResumeModerationNote         []DictionaryItem `json:"resume_moderation_note"`
	ResumeSearchExperiencePeriod []DictionaryItem `json:"resume_search_experience_period"`
	ResumeSearchFields           []DictionaryItem `json:"resume_search_fields"`
	ResumeSearchLabel            []DictionaryItem `json:"resume_search_label"`
	ResumeSearchLogic            []DictionaryItem `json:"resume_search_logic"`
	ResumeSearchOrder            []DictionaryItem `json:"resume_search_order"`
	ResumeSearchRelocation       []DictionaryItem `json:"resume_search_relocation"`
	ResumeStatus                 []DictionaryItem `json:"resume_status"`
	ResumeWorkFormat             []DictionaryItem `json:"resume_work_format"`
	SalaryRangeFrequency         []DictionaryItem `json:"salary_range_frequency"`
	SalaryRangeMode              []DictionaryItem `json:"salary_range_mode"`
	Schedule                     []struct {
		DictionaryItem
		UID string
	} `json:"schedule"`
	SetkaAccessType             []DictionaryItem `json:"setka_access_type"`
	TravelTime                  []DictionaryItem `json:"travel_time"`
	VacancyBillingType          []DictionaryItem `json:"vacancy_billing_type"`
	VacancyCluster              []DictionaryItem `json:"vacancy_cluster"`
	VacancyLabel                []DictionaryItem `json:"vacancy_label"`
	VacancyNotProlongedReason   []DictionaryItem `json:"vacancy_not_prolonged_reason"`
	VacancyRelation             []DictionaryItem `json:"vacancy_relation"`
	VacancySearchEmploymentForm []DictionaryItem `json:"vacancy_search_employment_form"`
	VacancySearchFields         []DictionaryItem `json:"vacancy_search_fields"`
	VacancySearchOrder          []DictionaryItem `json:"vacancy_search_order"`
	VacancyType                 []DictionaryItem `json:"vacancy_type"`
	WorkFormat                  []DictionaryItem `json:"work_format"`
	WorkScheduleByDays          []DictionaryItem `json:"work_schedule_by_days"`
	WorkingDays                 []DictionaryItem `json:"working_days"`
	WorkingHours                []DictionaryItem `json:"working_hours"`
	WorkingTimeIntervals        []DictionaryItem `json:"working_time_intervals"`
	WorkingTimeModes            []DictionaryItem `json:"working_time_modes"`
}

type DictionaryItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CurrencyItem struct {
	Abbr    string  `json:"abbr"`
	Code    string  `json:"code"`
	Default bool    `json:"default"`
	InUse   bool    `json:"in_use"`
	Name    string  `json:"name"`
	Rate    float64 `json:"rate"`
}
