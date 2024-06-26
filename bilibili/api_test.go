package bilibili

import (
	"testing"
)

func TestGetAllGuard(t *testing.T) {
	guardUser, err := GetAllGuard("628537")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", guardUser)
}

func TestGetDynamicDetail(t *testing.T) {
	cfg := NewCookieConfig("config.json")
	detail, err := GetDynamicDetail(cfg, "851252197280710664")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", detail)
}
