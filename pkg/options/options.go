package options

import (
	"github.com/corneliusweig/ketall/pkg/printer"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericclioptions/printers"
)

type KetallOptions struct {
	CfgFile         string
	GenericCliFlags *genericclioptions.ConfigFlags
	PrintFlags      KAPrintFlags
	Verbs           []v1.Verbs
}

func NewCmdOptions() *KetallOptions {
	return &KetallOptions{
		GenericCliFlags: genericclioptions.NewConfigFlags(),
		PrintFlags:      KAPrintFlags{genericclioptions.NewPrintFlags("")},
	}
}

type KAPrintFlags struct {
	*genericclioptions.PrintFlags
}

func (f *KAPrintFlags) ToPrinter() (printers.ResourcePrinter, error) {
	if f.OutputFormat == nil || *f.OutputFormat == "" || *f.OutputFormat == "short" {
		return printer.BasicTablePrinter{}, nil
	}
	return f.PrintFlags.ToPrinter()
}