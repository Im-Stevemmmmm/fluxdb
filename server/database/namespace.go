package database

import (
	"fmt"
)

// NewNamespace creates a new namespace and saves it.
func NewNamespace(name string) (*Namespace, *CreateNamespaceError) {
	path, err := NewRelativePath("ns_data")
	if err != nil {
		return nil, NewCreateNamespaceError(nil, err)
	}

	file, err := path.OpenFile()
	if err != nil {
		return nil, NewCreateNamespaceError(nil, err)
	}

	// Write data header
	if _, err := file.Write(nil); err != nil {
		return nil, NewCreateNamespaceError(nil, err)
	}

	return &Namespace{
		Name:      name,
		IndexRefs: make([]*interface{}, 0),
	}, nil
}

// NewCreateNamespaceError creates a new namespace error.
func NewCreateNamespaceError(ns *Namespace, err error) *CreateNamespaceError {
	return &CreateNamespaceError{
		Namespace: ns,
		Err:       err,
		LogData: NewLogData(
			fmt.Sprintf("Error creating namespace %s", ns.Name),
		),
	}
}

func (e *CreateNamespaceError) Error() string {
	return fmt.Sprintf("error creating namespace: Name=%s Time=%s", e.Namespace.Name, e.Timestamp)
}

// Namespace is a collection of indexes.
type Namespace struct {
	Name      string
	IndexRefs []*interface{}
}

// CreateNamespaceError is a namespace creation error.
type CreateNamespaceError struct {
	Namespace *Namespace
	Err       error
	*LogData
}
