package common

// TODO: Add description regarding the ETSI documentation
// TODO: Add validation functions that follow the ETSI documentation

type ProblemDetails struct {
	Type          Uri            `json:"type,omitempty"`          // A URI reference according to IETF RFC 3986 that identifies the problem type.
	Title         string         `json:"title,omitempty"`         // A short, human-readable summary of the problem type. It SHOULD NOT change from occurrence to occurrence of the problem.
	Status        int32          `json:"status,omitempty"`        // The HTTP status code for this occurrence of the problem.
	Detail        string         `json:"detail,omitempty"`        // A human-readable explanation specific to this occurrence of the problem.
	Instance      Uri            `json:"instance,omitempty"`      // A URI reference that identifies the specific occurrence of the problem.
	Cause         string         `json:"cause,omitempty"`         // A machine-readable application error cause specific to this occurrence of the problem This IE should be present and provide application- related error information, if available.
	InvalidParams []InvalidParam `json:"invalidParams,omitempty"` // Description of invalid parameters, for a request rejected due to invalid parameters.
}

// Validate validates this problem details.
func (p *ProblemDetails) Validate() error {
	if err := p.Type.Validate(); err != nil {
		return err
	}
	if err := p.Instance.Validate(); err != nil {
		return err
	}
	for _, invalidParam := range p.InvalidParams {
		if err := invalidParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Link struct {
	Href Uri `json:"href"` // It contains the URI of the linked resource.
}

// Validate validates this link.
func (l *Link) Validate() error {
	if err := l.Href.Validate(); err != nil {
		return err
	}
	return nil
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

func (i InvalidParam) Validate() error {
	return nil
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
