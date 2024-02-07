package RiotApi

import (
	"testing"
)

func TestApi_setRegion(t *testing.T) {
	type fields struct {
		Server string
		region string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "br1", fields: fields{
			region: "americas",
			Server: "br1",
		}},
		{name: "eun1", fields: fields{
			region: "europe",
			Server: "eun1",
		}},
		{name: "euw1", fields: fields{
			region: "europe",
			Server: "euw1",
		}},
		{name: "jp1", fields: fields{
			region: "asia",
			Server: "jp1",
		}},
		{name: "kr", fields: fields{
			region: "asia",
			Server: "kr",
		}},
		{name: "la1", fields: fields{
			region: "americas",
			Server: "la1",
		}},
		{name: "la2", fields: fields{
			region: "americas",
			Server: "la2",
		}},
		{name: "na1", fields: fields{
			region: "americas",
			Server: "na1",
		}},
		{name: "oc1", fields: fields{
			region: "sea",
			Server: "oc1",
		}},
		{name: "ph2", fields: fields{
			region: "sea",
			Server: "ph2",
		}},
		{name: "ru", fields: fields{
			region: "europe",
			Server: "ru",
		}},
		{name: "sg2", fields: fields{
			region: "sea",
			Server: "sg2",
		}},
		{name: "th2", fields: fields{
			region: "sea",
			Server: "th2",
		}},
		{name: "tr1", fields: fields{
			region: "europe",
			Server: "tr1",
		}},
		{name: "tw2", fields: fields{
			region: "sea",
			Server: "tw2",
		}},
		{name: "vn2", fields: fields{
			region: "sea",
			Server: "vn2",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Api{
				LoL: LoL{
					Server: tt.fields.Server,
				},
			}
			a.region = getRegion(a.Server)

			if a.region != tt.fields.region {
				t.Errorf("got %q, wanted %q", a.region, tt.fields.region)
			}
		})
	}
}
