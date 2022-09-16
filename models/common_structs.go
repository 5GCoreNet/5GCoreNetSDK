package models

// TODO: Add description regarding the ETSI documentation
// TODO: Add validation functions that follow the ETSI documentation

type ProblemDetails struct {
	Type          Uri            `json:"type,omitempty"`
	Title         string         `json:"title,omitempty"`
	Status        int32          `json:"status,omitempty"`
	Detail        string         `json:"detail,omitempty"`
	Instance      Uri            `json:"instance,omitempty"`
	Cause         string         `json:"cause,omitempty"`
	InvalidParams []InvalidParam `json:"invalidParams,omitempty"`
}

type Link struct {
	Href Uri `json:"href,omitempty"`
}

type PatchItem struct {
	Op    PatchOperation `json:"op,omitempty"`
	Path  string         `json:"path,omitempty"`
	From  string         `json:"from,omitempty"`
	Value interface{}    `json:"value,omitempty"`
}

type LinksValueSchema struct {
	Links []Link `json:"links,omitempty"`
	Link  Link   `json:"link,omitempty"`
}

type SelfLink struct {
	Self Link `json:"self,omitempty"`
}

type InvalidParam struct {
	Param  string `json:"param,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type LinkRM Link

type ChangeItem struct {
	Op        ChangeType  `json:"op,omitempty"`
	Path      string      `json:"path,omitempty"`
	From      string      `json:"from,omitempty"`
	OrigValue interface{} `json:"origValue,omitempty"`
	NewValue  interface{} `json:"newValue,omitempty"`
}

type NotifyItem struct {
	ResourceId Uri          `json:"resourceId,omitempty"`
	Changes    []ChangeItem `json:"changes,omitempty"`
}

type ComplexQuery struct {
	Cnf Cnf `json:"cnf,omitempty"`
	Dnf Dnf `json:"dnf,omitempty"`
}

type Cnf struct {
	CnfUnits []CnfUnit `json:"cnfUnits,omitempty"`
}

type Dnf struct {
	DnfUnits []DnfUnit `json:"dnfUnits,omitempty"`
}

type CnfUnit struct {
	CnfUnit []Atom `json:"cnfUnit,omitempty"`
}

type DnfUnit struct {
	DnfUnit []Atom `json:"dnfUnit,omitempty"`
}

type Atom struct {
	Attr     string      `json:"attr,omitempty"`
	Value    interface{} `json:"value,omitempty"`
	Negative bool        `json:"negative,omitempty"`
}
