package common

//Employees is
type Employees struct {
	EmployeeID      string `json:"CustomerID"`
	LastName        string `json:"LastName"`
	FirstName       string `json:"FirstName"`
	Title           string `json:"Title"`
	TitleOfCourtesy string `json:"TitleOfCourtesy"`
	BirthDate       string `json:"BirthDate"`
	HireDate        string `json:"HireDate"`
	Address         string `json:"Address"`
	City            string `json:"City"`
	Region          string `json:"Region"`
	PostalCode      string `json:"PostalCode"`
	Country         string `json:"Country"`
	HomePhone       string `json:"HomePhone"`
	Extension       string `json:"Extension"`
	Photo           string `json:"Photo"`
	Notes           string `json:"Notes"`
	ReportsTo       string `json:"ReportsTo"`
	ProvinceName    string `json:"ProvinceName"`
}
