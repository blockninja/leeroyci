package config

import (
	"testing"
)

func TestValidate(t *testing.T) {
	c := Config{}

	if c.Validate() == nil {
		t.Error("Config is valid, should be invalid - Secret")
	}

	c.Secret = "foo"

	if c.Validate() == nil {
		t.Error("Config is valid, should be invalid - BuildLogPath")
	}

	c.BuildLogPath = "foo"

	if c.Validate() == nil {
		t.Error("Config is valid, should be invalid - URL")
	}

	c.URL = "http://foo"

	if c.Validate() != nil {
		t.Error("Config is invalid, should be invalid - http")
	}

	c.URL = "https://foo"

	if c.Validate() == nil {
		t.Error("Config is valid, should be invalid - Cert")
	}

	c.Cert = "foo"

	if c.Validate() == nil {
		t.Error("Config is valid, should be invalid - Key")
	}

	c.Key = "foo"

	if c.Validate() != nil {
		t.Error("Config is invalid, should be invalid")
	}
}

func TestValidateRepository(t *testing.T) {
	r1 := Repository{
		URL: "foo",
	}
	r2 := Repository{}

	e1 := validateRepository(&r1)
	e2 := validateRepository(&r2)

	if e1 != nil {
		t.Error("Repository is invalid, should be valid", e1)
	}

	if e2 == nil {
		t.Error("Repository is valid, should be invalid", e2)
	}
}

func TestValidateCommand(t *testing.T) {
	c1 := Command{
		Name:    "foo",
		Execute: "bar",
	}

	c2 := Command{
		Execute: "bar",
	}

	c3 := Command{
		Name: "foo",
	}

	e1 := validateCommand(&c1)
	e2 := validateCommand(&c2)
	e3 := validateCommand(&c3)

	if e1 != nil {
		t.Error("c1 is invalid, should be valid", e1)
	}

	if e2 == nil {
		t.Error("c2 is valid, should be invalid", e2)
	}

	if e3 == nil {
		t.Error("c3 is valid, should be invalid", e3)
	}
}

func TestValidateNotify(t *testing.T) {
	n1 := Notify{
		Service: "foo",
	}

	n2 := Notify{}

	e1 := validateNotify(&n1)
	e2 := validateNotify(&n2)

	if e1 != nil {
		t.Error("n1 is invalid, should be valid", e1)
	}

	if e2 == nil {
		t.Error("n2 is valid, should be invalid", e2)
	}
}

func TestValidateNotifySlack(t *testing.T) {
	n := Notify{
		Service:   "slack",
		Arguments: make(map[string]string),
	}

	n.Arguments["endpoint"] = "foo"

	err := validateNotify(&n)

	if err == nil {
		t.Error("Error should not be nil, no channel")
	}

	n.Arguments = make(map[string]string)
	n.Arguments["channel"] = "foo"

	err = validateNotify(&n)

	if err == nil {
		t.Error("Error should not be nil, no endpoint")
	}

	n.Arguments["endpoint"] = "bar"

	err = validateNotify(&n)

	if err != nil {
		t.Error("Slack configuration is invalid, should be valid")
	}
}

func TestValidateNotifyHipChat(t *testing.T) {
	n := Notify{
		Service:   "hipchat",
		Arguments: make(map[string]string),
	}

	n.Arguments["channel"] = "foo"

	err := validateNotify(&n)

	if err == nil {
		t.Error("Error should not be nil, no key")
	}

	n.Arguments = make(map[string]string)
	n.Arguments["key"] = "foo"

	err = validateNotify(&n)

	if err == nil {
		t.Error("Error should not be nil, no channel")
	}

	n.Arguments["channel"] = "bar"

	err = validateNotify(&n)

	if err != nil {
		t.Error("HipChat configuration is invalid, should be valid")
	}
}

func TestValidateNotifyCampfire(t *testing.T) {
	n := Notify{
		Service:   "campfire",
		Arguments: make(map[string]string),
	}

	n.Arguments["id"] = "foo"
	n.Arguments["room"] = "bar"

	err := validateNotify(&n)

	if err == nil {
		t.Error("Error should not be nil, no key")
	}

	n.Arguments = make(map[string]string)
	n.Arguments["id"] = "foo"
	n.Arguments["key"] = "bar"

	err = validateNotify(&n)

	if err == nil {
		t.Error("Error should not be nil, no room")
	}

	n.Arguments = make(map[string]string)
	n.Arguments["room"] = "foo"
	n.Arguments["key"] = "bar"

	err = validateNotify(&n)

	if err == nil {
		t.Error("Error should not be nil, no id")
	}

	n.Arguments["id"] = "bar"

	err = validateNotify(&n)

	if err != nil {
		t.Error("Campfire configuration is invalid, should be valid")
	}
}
