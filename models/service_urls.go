package models

type ServiceUrlConfig struct {
    AuthServiceURL string
    HrmServiceURL string
    LogsServiceURL string
	PmServiceUrl string
	ScmServiceURL string
	AccountingServiceURL string
	ManufacturingServiceURL string
	NotificationsServiceURL string
	SalesServiceURL string
	DocumentsServiceURL string
	CrmServiceURL string

	SaveLogsUrl string
	CheckUserExistsUrl string
}