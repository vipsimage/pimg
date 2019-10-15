package pimg

import (
	"image"
	"reflect"
	"testing"

	"github.com/vipsimage/vips"
)

func Test_watermarkPosition(t *testing.T) {
	type args struct {
		direction vips.CompassDirection
		base      rect
		wm        rect
	}
	tests := []struct {
		name      string
		args      args
		wantPoint image.Point
	}{
		{
			name: "DirectionEast",
			args: args{
				direction: vips.DirectionEast,
				base: rect{
					width:  2300,
					height: 1200,
				},
				wm: rect{
					width:  120,
					height: 80,
				},
			},
			wantPoint: image.Point{
				X: 2300 - 120,
				Y: (1200 - 80) / 2,
			},
		},
		{
			name: "DirectionEast-1",
			args: args{
				direction: vips.DirectionEast,
				base: rect{
					width:  3300,
					height: 3200,
				},
				wm: rect{
					width:  220,
					height: 80,
				},
			},
			wantPoint: image.Point{
				X: 3300 - 220,
				Y: (3200 - 80) / 2,
			},
		},
		{
			name: "DirectionSouth",
			args: args{
				direction: vips.DirectionSouth,
				base: rect{
					width:  2300,
					height: 1200,
				},
				wm: rect{
					width:  120,
					height: 80,
				},
			},
			wantPoint: image.Point{
				X: (2300 - 120) / 2,
				Y: 1200 - 80,
			},
		},
		{
			name: "DirectionWest",
			args: args{
				direction: vips.DirectionWest,
				base: rect{
					width:  2300,
					height: 1200,
				},
				wm: rect{
					width:  120,
					height: 80,
				},
			},
			wantPoint: image.Point{
				X: 0,
				Y: (1200 - 80) / 2,
			},
		},
		{
			name: "DirectionNorth",
			args: args{
				direction: vips.DirectionNorth,
				base: rect{
					width:  2300,
					height: 1200,
				},
				wm: rect{
					width:  120,
					height: 80,
				},
			},
			wantPoint: image.Point{
				X: (2300 - 120) / 2,
				Y: 0,
			},
		},
		{
			name: "DirectionCentre",
			args: args{
				direction: vips.DirectionCentre,
				base: rect{
					width:  2300,
					height: 1200,
				},
				wm: rect{
					width:  120,
					height: 80,
				},
			},
			wantPoint: image.Point{
				X: (2300 - 120) / 2,
				Y: (1200 - 80) / 2,
			},
		},
		{
			name: "DirectionNorthEast",
			args: args{
				direction: vips.DirectionNorthEast,
				base: rect{
					width:  2300,
					height: 1200,
				},
				wm: rect{
					width:  120,
					height: 80,
				},
			},
			wantPoint: image.Point{
				X: 2300 - 120,
				Y: 0,
			},
		},
		{
			name: "DirectionSouthEast",
			args: args{
				direction: vips.DirectionSouthEast,
				base: rect{
					width:  2300,
					height: 1200,
				},
				wm: rect{
					width:  120,
					height: 80,
				},
			},
			wantPoint: image.Point{
				X: 2300 - 120,
				Y: 1200 - 80,
			},
		},
		{
			name: "DirectionSouthWest",
			args: args{
				direction: vips.DirectionSouthWest,
				base: rect{
					width:  2300,
					height: 1200,
				},
				wm: rect{
					width:  120,
					height: 80,
				},
			},
			wantPoint: image.Point{
				X: 0,
				Y: 1200 - 80,
			},
		},
		{
			name: "DirectionNorthWest",
			args: args{
				direction: vips.DirectionNorthWest,
				base: rect{
					width:  2300,
					height: 1200,
				},
				wm: rect{
					width:  120,
					height: 80,
				},
			},
			wantPoint: image.Point{
				X: 0,
				Y: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPoint := watermarkPosition(tt.args.direction, tt.args.base, tt.args.wm); !reflect.DeepEqual(gotPoint, tt.wantPoint) {
				t.Errorf("watermarkPosition() = %v, want %v", gotPoint, tt.wantPoint)
			}
		})
	}
}
