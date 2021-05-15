package helpc

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
)

// RunHeader show help header
func RunHeader(a app.App, ctx app.IOContext) (err error) {
	var (
		deps struct {
			Info commservices.Info `dependency:"CommonInfo"`
		}
		io = ctx.IO()
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	info := deps.Info
	// header
	io.Out().Printf("%s %s\n", a.Name(), a.Version().String())
	if info.License().Company() != "" {
		license := info.License()
		io.Out().Printf("Develop by @%s", license.Company())
		if !license.Start().IsZero() {
			io.Out().Printf(" %v", license.Start().Year())
			if !license.End().IsZero() {
				io.Out().Printf("-%v", license.End().Year())
			}
		}
		if license.Type() != "" {
			// type is like: "all rights reserved"
			io.Out().Printf(" %s", license.Type())
		}
		io.Out().Printf("\n")
	}
	if info.PoweredBy() != "" {
		io.Out().Printf("%s\n", info.PoweredBy())
	}
	if info.Description() != "" {
		io.Out().Printf("\n%s\n\n", info.Description())
	}
	return nil
}
