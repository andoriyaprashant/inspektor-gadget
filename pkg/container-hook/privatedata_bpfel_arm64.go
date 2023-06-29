// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package containerhook

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadPrivatedata returns the embedded CollectionSpec for privatedata.
func loadPrivatedata() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_PrivatedataBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load privatedata: %w", err)
	}

	return spec, err
}

// loadPrivatedataObjects loads privatedata and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*privatedataObjects
//	*privatedataPrograms
//	*privatedataMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadPrivatedataObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadPrivatedata()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// privatedataSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type privatedataSpecs struct {
	privatedataProgramSpecs
	privatedataMapSpecs
}

// privatedataSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type privatedataProgramSpecs struct {
	IgFgetX   *ebpf.ProgramSpec `ebpf:"ig_fget_x"`
	IgScmSndE *ebpf.ProgramSpec `ebpf:"ig_scm_snd_e"`
}

// privatedataMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type privatedataMapSpecs struct {
	IgPrivateData *ebpf.MapSpec `ebpf:"ig_private_data"`
}

// privatedataObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadPrivatedataObjects or ebpf.CollectionSpec.LoadAndAssign.
type privatedataObjects struct {
	privatedataPrograms
	privatedataMaps
}

func (o *privatedataObjects) Close() error {
	return _PrivatedataClose(
		&o.privatedataPrograms,
		&o.privatedataMaps,
	)
}

// privatedataMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadPrivatedataObjects or ebpf.CollectionSpec.LoadAndAssign.
type privatedataMaps struct {
	IgPrivateData *ebpf.Map `ebpf:"ig_private_data"`
}

func (m *privatedataMaps) Close() error {
	return _PrivatedataClose(
		m.IgPrivateData,
	)
}

// privatedataPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadPrivatedataObjects or ebpf.CollectionSpec.LoadAndAssign.
type privatedataPrograms struct {
	IgFgetX   *ebpf.Program `ebpf:"ig_fget_x"`
	IgScmSndE *ebpf.Program `ebpf:"ig_scm_snd_e"`
}

func (p *privatedataPrograms) Close() error {
	return _PrivatedataClose(
		p.IgFgetX,
		p.IgScmSndE,
	)
}

func _PrivatedataClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed privatedata_bpfel_arm64.o
var _PrivatedataBytes []byte
