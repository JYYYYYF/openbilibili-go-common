package special

import (
	"context"
	"flag"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"go-common/app/interface/main/app-card/model/card/operate"
	"go-common/app/interface/main/app-feed/conf"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	d *Dao
)

func init() {
	dir, _ := filepath.Abs("../../cmd/app-feed-test.toml")
	flag.Set("conf", dir)
	conf.Init()
	d = New(conf.Conf)
}

func TestNew(t *testing.T) {
	type args struct {
		c *conf.Config
	}
	tests := []struct {
		name  string
		args  args
		wantD *Dao
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotD := New(tt.args.c); !reflect.DeepEqual(gotD, tt.wantD) {
				t.Errorf("New() = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}

func TestDao_Card(t *testing.T) {
	type args struct {
		c   context.Context
		now time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantScm map[int64]*operate.Special
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Convey(tt.name, t, func() {
			gotScm, err := d.Card(tt.args.c, tt.args.now)
			So(gotScm, ShouldEqual, tt.wantScm)
			So(err, ShouldEqual, tt.wantErr)
		})
	}
}

func TestDao_Close(t *testing.T) {
	tests := []struct {
		name string
		d    *Dao
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Close()
		})
	}
}
