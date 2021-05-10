package goatapp

import (
	"testing"
)

func TestMockupapp(t *testing.T) {
	var (
		err  error
		mapp *MockupApp
	)
	t.Parallel()
	if mapp, err = NewMockupApp(Params{
		Arguments: []string{"arg1", "arg2"},
		Version:   NewVersion(2, 1, 1, "-dev"),
	}); err != nil {
		t.Error(err)
	}
	if len(mapp.Arguments()) != 2 {
		t.Errorf("Expected 2 arguments")
	}
	if mapp.DependencyProvider() == nil {
		t.Errorf("Expected dependency provider")
	}
	if mapp.Filespaces().CWD() == nil {
		t.Errorf("Expected CWD filespaces")
	}
	if mapp.Filespaces().Home() == nil {
		t.Errorf("Expected Home filespaces")
	}
	if mapp.Filespaces().Root() == nil {
		t.Errorf("Expected Root filespaces")
	}
	if mapp.Filespaces().Tmp() == nil {
		t.Errorf("Expected Tmp filespaces")
	}
	if mapp.IOContext() == nil {
		t.Errorf("Expected IOContext")
	}
	if mapp.Name() == "" {
		t.Errorf("Expected Name")
	}
	if mapp.Scopes().App() == nil {
		t.Errorf("Expected App scope")
	}
	if mapp.Scopes().Arguments() == nil {
		t.Errorf("Expected Arguments scope")
	}
	if mapp.Scopes().Config() == nil {
		t.Errorf("Expected Config scope")
	}
	if mapp.Scopes().Filespace() == nil {
		t.Errorf("Expected Filespace scope")
	}
	if mapp.Terminal() == nil {
		t.Errorf("Expected Terminal scope")
	}
	if mapp.Version().Major() != 2 {
		t.Errorf("Expected version major equals to 2")
	}
	if len(mapp.HealthCheckerNames()) != 0 {
		t.Errorf("Expected empty helth checker list")
	}
}
