package common

import "fmt"

type ProblemDetails struct {
	Type          *Uri            `json:"type,omitempty"`          // A URI reference according to IETF RFC 3986 that identifies the problem type.
	Title         *string         `json:"title,omitempty"`         // A short, human-readable summary of the problem type. It SHOULD NOT change from occurrence to occurrence of the problem.
	Status        *int32          `json:"status,omitempty"`        // The HTTP status code for this occurrence of the problem.
	Detail        *string         `json:"detail,omitempty"`        // A human-readable explanation specific to this occurrence of the problem.
	Instance      *Uri            `json:"instance,omitempty"`      // A URI reference that identifies the specific occurrence of the problem.
	Cause         *string         `json:"cause,omitempty"`         // A machine-readable application error cause specific to this occurrence of the problem This IE should be present and provide application- related error information, if available.
	InvalidParams []*InvalidParam `json:"invalidParams,omitempty"` // Description of invalid parameters, for a request rejected due to invalid parameters.
	// NOTE 1: See IETF RFC 7807 for detailed information and guidance for each attribute, and 3GPP TS 29.501 for guidelines on error handling support by 5GC SBI APIs.
	//NOTE 2: Additional attributes may be defined per API.
}

// Validate validates this problem details.
func (p *ProblemDetails) Validate() error {
	if p.Type != nil {
		if err := p.Type.Validate(); err != nil {
			return err
		}
	}
	if p.Instance != nil {
		if err := p.Instance.Validate(); err != nil {
			return err
		}
	}
	for _, invalidParam := range p.InvalidParams {
		if err := invalidParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Link struct {
	Href *Uri `json:"href"` // It contains the URI of the linked resource.
}

// Validate validates this link.
func (l *Link) Validate() error {
	if err := l.Href.Validate(); err != nil {
		return err
	}
	return nil
}

type PatchItem struct {
	Op    *PatchOperation `json:"op"`              //This IE indicates the patch operation as defined in IETF RFC 6902 to be performed on resource.
	Path  *string         `json:"path"`            // This IE contains a JSON pointer value (as defined in IETF RFC 6901) that references a location of a resource on which the patch operation shall be performed.
	From  *string         `json:"from,omitempty"`  // This IE indicates the path of the source JSON element (according to JSON Pointer syntax) being moved or copied to the location indicated by the "path" attribute. It shall be present if the patch operation is "move" or "copy".
	Value *interface{}    `json:"value,omitempty"` // This IE indicates a new value for the resource specified in the path attribute. It shall be present if the patch operation is "add", "replace" or "test". The null value shall be allowed.
}

// Validate validates this patch item.
func (p *PatchItem) Validate() error {
	if err := p.Op.Validate(); err != nil {
		return err
	}
	if p.Path == nil {
		return fmt.Errorf("path is required")
	}
	if *p.Op == PatchOperationMove || *p.Op == PatchOperationCopy {
		if p.From == nil {
			return fmt.Errorf("from must be present if op is move or copy")
		}
	} else {
		if p.From != nil {
			return fmt.Errorf("from must not be present if op is not move or copy")
		}
		if p.Value == nil {
			return fmt.Errorf("value must be present if op is not move or copy")
		}
	}
	return nil
}

// LinksValueSchema is a list of mutually exclusive alternatives Links or Link.
type LinksValueSchema struct {
	Links []*Link `json:"links,omitempty"` // Array of links.
	Link  *Link   `json:"link,omitempty"`  // Link.
}

// Validate validates this links value schema.
func (l *LinksValueSchema) Validate() error {
	if l.Links != nil && l.Link != nil {
		return fmt.Errorf("links and link must not be both present")
	}
	if l.Links == nil && l.Link == nil {
		return fmt.Errorf("links or link must be present")
	}
	if l.Links != nil {
		for _, link := range l.Links {
			if err := link.Validate(); err != nil {
				return err
			}
		}
	}
	if l.Link != nil {
		if err := l.Link.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SelfLink struct {
	Self *Link `json:"self"` // It contains the URI of the linked resource.
}

// Validate validates this self link.
func (s *SelfLink) Validate() error {
	if err := s.Self.Validate(); err != nil {
		return err
	}
	return nil
}

type InvalidParam struct {
	Param  *string `json:"param"`            // Attribute's name encoded as a JSON Pointer.
	Reason *string `json:"reason,omitempty"` // A human-readable reason, e.g. "must be a positive integer".
}

// Validate validates this invalid param.
func (i *InvalidParam) Validate() error {
	if i.Param == nil {
		return fmt.Errorf("param is required")
	}
	return nil
}

// LinkRM is defined in the same way as the Link data type, but with the OpenAPI "nullable: true" property.
type LinkRM Link

// Validate validates this link rm.
func (l *LinkRM) Validate() error {
	if l == nil {
		return nil
	}
	return (*Link)(l).Validate()
}

type ChangeItem struct {
	Op        *ChangeType  `json:"op"`                  // This IE indicates the change type which happens to the resource.
	Path      *string      `json:"path"`                // This IE contains a JSON pointer value (as defined in IETF RFC 6901) that references a location of an attribute on which the change has been applied.
	From      *string      `json:"from,omitempty"`      // This IE indicates the path of the source JSON element (according to JSON Pointer syntax) being moved or copied to the location indicated by the "path" attribute. It shall be present if the "op" attribute is of value "MOVE".
	OrigValue *interface{} `json:"origValue,omitempty"` // This IE indicates the original value of the attribute specified in the path attribute. This attribute only applies when the "op" attribute is of value "REMOVE", "REPLACE" or "MOVE" Based on the use case, this attribute may be included.
	NewValue  *interface{} `json:"newValue,omitempty"`  // This IE indicates a new value of the attribute specified in the path attribute. It shall be present if the "op" attribute is of value "ADD", "REPLACE". The null value shall be allowed.
}

// Validate validates this change item.
func (c *ChangeItem) Validate() error {
	if err := c.Op.Validate(); err != nil {
		return err
	}
	if c.Path == nil {
		return fmt.Errorf("path is required")
	}
	if *c.Op == ChangeTypeMove {
		if c.From == nil {
			return fmt.Errorf("from must be present if op is move")
		}
	} else {
		if c.From != nil {
			return fmt.Errorf("from must not be present if op is not move")
		}
	}
	if *c.Op == ChangeTypeAdd || *c.Op == ChangeTypeReplace {
		if c.NewValue == nil {
			return fmt.Errorf("newValue must be present if op is add or replace")
		}
	} else {
		if c.NewValue != nil {
			return fmt.Errorf("newValue must not be present if op is not add or replace")
		}
	}
	if *c.Op == ChangeTypeRemove || *c.Op == ChangeTypeReplace || *c.Op == ChangeTypeMove {
		if c.OrigValue == nil {
			return fmt.Errorf("origValue must be present if op is remove, replace or move")
		}
	} else {
		if c.OrigValue != nil {
			return fmt.Errorf("origValue must not be present if op is not remove, replace or move")
		}
	}
	return nil
}

type NotifyItem struct {
	ResourceId *Uri          `json:"resourceId"` // This IE contains the URI of the resource which has been changed.
	Changes    []*ChangeItem `json:"changes"`    // This IE contains the changes which have been applied on the resource identified by the resourceId attribute.
}

// Validate validates this notify item.
func (n *NotifyItem) Validate() error {
	if err := n.ResourceId.Validate(); err != nil {
		return err
	}
	if len(n.Changes) == 0 {
		return fmt.Errorf("changes must not be empty")
	}
	for _, change := range n.Changes {
		if err := change.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// ComplexQuery data type is either a conjunctive normal form or a disjunctive normal form. The attribute names
// "cnfUnits" and "dnfUnits" serve as discriminator.
type ComplexQuery struct {
	Cnf *Cnf `json:"cnf,omitempty"` // A conjunctive normal form.
	Dnf *Dnf `json:"dnf,omitempty"` // A disjunctive normal form.
}

// Validate validates this complex query.
func (c *ComplexQuery) Validate() error {
	if c.Cnf != nil && c.Dnf != nil {
		return fmt.Errorf("cnf and dnf must not be both present")
	}
	if c.Cnf == nil && c.Dnf == nil {
		return fmt.Errorf("cnf or dnf must be present")
	}
	if c.Cnf != nil {
		if err := c.Cnf.Validate(); err != nil {
			return err
		}
	}
	if c.Dnf != nil {
		if err := c.Dnf.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Cnf struct {
	CnfUnits []*CnfUnit `json:"cnfUnits"` // During the processing of cnfUnits attribute, all the members in the array shall be interpreted as logically concatenated with logical "AND".
}

// Validate validates this cnf.
func (c *Cnf) Validate() error {
	if len(c.CnfUnits) == 0 {
		return fmt.Errorf("cnfUnits must not be empty")
	}
	for _, cnfUnit := range c.CnfUnits {
		if err := cnfUnit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Dnf struct {
	DnfUnits []*DnfUnit `json:"dnfUnits"` // During the processing of dnfUnits attribute, all the members in the array shall be interpreted as logically concatenated with logical "OR"
}

// Validate validates this dnf.
func (d *Dnf) Validate() error {
	if len(d.DnfUnits) == 0 {
		return fmt.Errorf("dnfUnits must not be empty")
	}
	for _, dnfUnit := range d.DnfUnits {
		if err := dnfUnit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CnfUnit struct {
	CnfUnit []*Atom `json:"cnfUnit"` // During the processing of cnfUnit attribute, all the members in the array shall be interpreted as logically concatenated with logical "OR"
}

// Validate validates this cnf unit.
func (c *CnfUnit) Validate() error {
	if len(c.CnfUnit) == 0 {
		return fmt.Errorf("cnfUnit must not be empty")
	}
	for _, atom := range c.CnfUnit {
		if err := atom.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DnfUnit struct {
	DnfUnit []*Atom `json:"dnfUnit"` // During the processing of dnfUnit attribute, all the members in the array shall be interpreted as logically concatenated with logical "AND"
}

// Validate validates this dnf unit.
func (d *DnfUnit) Validate() error {
	if len(d.DnfUnit) == 0 {
		return fmt.Errorf("dnfUnit must not be empty")
	}
	for _, atom := range d.DnfUnit {
		if err := atom.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Atom struct {
	Attr     *string      `json:"attr"`               // This attribute contains the name of a defined query parameter.
	Value    *interface{} `json:"value"`              // This attribute contains the value of the query parameter as indicated by attr attribute.
	Negative *bool        `json:"negative,omitempty"` // This attribute indicates whether the negative condition applies for the query condition.
}

// Validate validates this atom.
func (a *Atom) Validate() error {
	if a.Attr == nil {
		return fmt.Errorf("attr must be present")
	}
	if a.Value == nil {
		return fmt.Errorf("value must be present")
	}
	return nil
}
