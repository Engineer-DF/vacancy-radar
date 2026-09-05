package vacancy

import "time"

// Согласно разметке HH, служебные мета-поля начинаются с @.

type InitialState struct {
	VacancyView  Vacancy `json:"vacancyView"` // корень JSON содержит объект vacancyView
	CanonicalURL string  `json:"canonicalUrl"`
}

type Vacancy struct {
	VacancyID        int          `json:"vacancyId"`        // Уникальный ID вакансии
	Name             string       `json:"name"`             // Название вакансии
	Description      string       `json:"description"`      // Описание в формате HTML
	PublicationDate  time.Time    `json:"publicationDate"`  // Дата и время публикации
	ValidThroughTime time.Time    `json:"validThroughTime"` // Срок действия публикации
	WorkExperience   string       `json:"workExperience"`   // Опыт работы (например, "between1And3")
	EmploymentForm   string       `json:"employmentForm"`   // Занятость (например, "FULL")
	WorkFormats      []string     `json:"workFormats"`      // Формат работы (например, ["ON_SITE"])
	Compensation     Compensation `json:"compensation"`     // Зарплата
	Company          Company      `json:"company"`          // Работодатель
	Area             Area         `json:"area"`             // Регион
	Address          *Address     `json:"address"`          // Физический адрес (может быть null)
	KeySkills        KeySkills    `json:"keySkills"`        // Ключевые навыки
}

type Compensation struct {
	From         int    `json:"from"`         // Минимальная планка зарплаты
	To           int    `json:"to,omitempty"` // Максимальная планка (если указана)
	CurrencyCode string `json:"currencyCode"` // Код валюты (например, "RUR")
	Gross        bool   `json:"gross"`        // true - до вычета налогов, false - на руки
	Mode         string `json:"mode"`         // Период выплат (например, "MONTH")
}

type Logo struct {
	Type string `json:"@type"` // Тип размера (ORIGINAL, small, medium, vacancyPage)
	URL  string `json:"@url"`  // Путь к изображению
}

type Company struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`                 // Официальное название
	VisibleName          string `json:"visibleName"`          // Отображаемое название
	CompanySiteURL       string `json:"companySiteUrl"`       // Сайт компании
	AccreditedITEmployer bool   `json:"accreditedITEmployer"` // Аккредитованная ИТ-компания
	Logos                struct {
		Logo []Logo `json:"logo"`
	} `json:"logos"`
}

type Area struct {
	ID         int    `json:"@id"`
	Name       string `json:"name"`       // Город (например, "Краснодар")
	RegionName string `json:"regionName"` // Область/Край (например, "Краснодарский край")
}

type Address struct {
	City        string `json:"city"`        // Город
	Street      string `json:"street"`      // Улица
	Building    string `json:"building"`    // Дом/Корпус
	DisplayName string `json:"displayName"` // Полный адрес строкой
}

type KeySkills struct {
	KeySkill []string // Массив навыков, например ["Golang", "PostgreSQL"]
}
