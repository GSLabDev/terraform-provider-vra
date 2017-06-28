package vra

type blueprintlistJsonstruct struct {
	Content []struct {
		_type       string `json:"@type"`
		CatalogItem struct {
			Callbacks          interface{} `json:"callbacks"`
			CatalogItemTypeRef struct {
				ID    string `json:"id"`
				Label string `json:"label"`
			} `json:"catalogItemTypeRef"`
			DateCreated string `json:"dateCreated"`
			Description string `json:"description"`
			Forms       struct {
				CatalogRequestInfoHidden bool `json:"catalogRequestInfoHidden"`
				ItemDetails              struct {
					FormID string `json:"formId"`
					Type   string `json:"type"`
				} `json:"itemDetails"`
				RequestDetails struct {
					FormID string `json:"formId"`
					Type   string `json:"type"`
				} `json:"requestDetails"`
				RequestFormScale  string `json:"requestFormScale"`
				RequestSubmission struct {
					FormID string `json:"formId"`
					Type   string `json:"type"`
				} `json:"requestSubmission"`
			} `json:"forms"`
			IconID          string `json:"iconId"`
			ID              string `json:"id"`
			IsNoteworthy    bool   `json:"isNoteworthy"`
			LastUpdatedDate string `json:"lastUpdatedDate"`
			Name            string `json:"name"`
			Organization    struct {
				SubtenantLabel interface{} `json:"subtenantLabel"`
				SubtenantRef   interface{} `json:"subtenantRef"`
				TenantLabel    string      `json:"tenantLabel"`
				TenantRef      string      `json:"tenantRef"`
			} `json:"organization"`
			ProviderBinding struct {
				BindingID   string `json:"bindingId"`
				ProviderRef struct {
					ID    string `json:"id"`
					Label string `json:"label"`
				} `json:"providerRef"`
			} `json:"providerBinding"`
			Quota       int  `json:"quota"`
			Requestable bool `json:"requestable"`
			ServiceRef  struct {
				ID    string `json:"id"`
				Label string `json:"label"`
			} `json:"serviceRef"`
			Status     string `json:"status"`
			StatusName string `json:"statusName"`
			Version    int    `json:"version"`
		} `json:"catalogItem"`
		EntitledOrganizations []struct {
			SubtenantLabel string `json:"subtenantLabel"`
			SubtenantRef   string `json:"subtenantRef"`
			TenantLabel    string `json:"tenantLabel"`
			TenantRef      string `json:"tenantRef"`
		} `json:"entitledOrganizations"`
	} `json:"content"`
	Links    []interface{} `json:"links"`
	Metadata struct {
		Number        int `json:"number"`
		Offset        int `json:"offset"`
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
	} `json:"metadata"`
}

type templateJsonstruct struct {
	BusinessGroupID string      `json:"businessGroupId"`
	CatalogItemID   string      `json:"catalogItemId"`
	Data            interface{} `json:"data"`
	Description     interface{} `json:"description"`
	Reasons         interface{} `json:"reasons"`
	RequestedFor    string      `json:"requestedFor"`
	Type            string      `json:"type"`
}

type requestJsonstruct struct {
	_type                      string `json:"@type"`
	ApprovalStatus             string `json:"approvalStatus"`
	CatalogItemProviderBinding struct {
		BindingID   string `json:"bindingId"`
		ProviderRef struct {
			ID    string `json:"id"`
			Label string `json:"label"`
		} `json:"providerRef"`
	} `json:"catalogItemProviderBinding"`
	CatalogItemRef struct {
		ID    string `json:"id"`
		Label string `json:"label"`
	} `json:"catalogItemRef"`
	Components      interface{} `json:"components"`
	DateApproved    interface{} `json:"dateApproved"`
	DateCompleted   interface{} `json:"dateCompleted"`
	DateCreated     string      `json:"dateCreated"`
	DateSubmitted   string      `json:"dateSubmitted"`
	Description     interface{} `json:"description"`
	ExecutionStatus string      `json:"executionStatus"`
	IconID          string      `json:"iconId"`
	ID              string      `json:"id"`
	LastUpdated     string      `json:"lastUpdated"`
	Organization    struct {
		SubtenantLabel interface{} `json:"subtenantLabel"`
		SubtenantRef   string      `json:"subtenantRef"`
		TenantLabel    interface{} `json:"tenantLabel"`
		TenantRef      string      `json:"tenantRef"`
	} `json:"organization"`
	Phase             string      `json:"phase"`
	PostApprovalID    interface{} `json:"postApprovalId"`
	PreApprovalID     interface{} `json:"preApprovalId"`
	Quote             struct{}    `json:"quote"`
	Reasons           interface{} `json:"reasons"`
	RequestCompletion interface{} `json:"requestCompletion"`
	RequestData       struct {
		Entries []struct {
			Key   string `json:"key"`
			Value struct {
				ClassID     string      `json:"classId"`
				ComponentID interface{} `json:"componentId"`
				ID          string      `json:"id"`
				Label       string      `json:"label"`
				Type        string      `json:"type"`
			} `json:"value"`
		} `json:"entries"`
	} `json:"requestData"`
	RequestNumber            interface{} `json:"requestNumber"`
	RequestedBy              string      `json:"requestedBy"`
	RequestedFor             string      `json:"requestedFor"`
	RequestedItemDescription string      `json:"requestedItemDescription"`
	RequestedItemName        string      `json:"requestedItemName"`
	RequestorEntitlementID   string      `json:"requestorEntitlementId"`
	RetriesRemaining         int         `json:"retriesRemaining"`
	State                    string      `json:"state"`
	StateName                interface{} `json:"stateName"`
	Version                  int         `json:"version"`
	WaitingStatus            string      `json:"waitingStatus"`
}
