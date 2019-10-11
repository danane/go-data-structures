package nodelist

import (
	"testing"

	"github.com/danane/go-data-structures/node"
)

func TestNodeList_RemoveElement(t *testing.T) {
	t.Parallel()
	type fields struct {
		head   *node.Node
		length int
	}
	type args struct {
		e int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantLen int
	}{
		{
			name: "Head removal",
			args: args{e: 1},
			fields: fields{
				head:   node.NewNode(1, nil),
				length: 1,
			},
			wantLen: 0,
		},
		{
			name: "Middle removal",
			args: args{e: 2},
			fields: fields{
				head:   node.NewNode(1, node.NewNode(2, node.NewNode(3, nil))),
				length: 3,
			},
			wantLen: 2,
		},
		{
			name: "Tail removal",
			args: args{e: 2},
			fields: fields{
				head:   node.NewNode(1, node.NewNode(2, nil)),
				length: 2,
			},
			wantLen: 1,
		},
		{
			name: "Empty list",
			args: args{e: 1},
			fields: fields{
				head: nil,
			},
			wantLen: 0,
			wantErr: true,
		},
		{
			name: "Key not found",
			args: args{e: 2},
			fields: fields{
				head:   node.NewNode(1, nil),
				length: 1,
			},
			wantLen: 1,
			wantErr: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			nl := &NodeList{
				head:   tc.fields.head,
				length: tc.fields.length,
			}
			if err := nl.Remove(tc.args.e); (err != nil) != tc.wantErr {
				t.Errorf("NodeList.RemoveElement() error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantLen != nl.length {
				t.Errorf("NodeList length = %d, wantLen %d", nl.length, tc.wantLen)
			}
		})
	}
}

func TestNodeList_InsertElement(t *testing.T) {
	type fields struct {
		head   *node.Node
		length int
	}
	type args struct {
		e int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantLen int
	}{
		{
			name: "Empty list",
			fields: fields{
				head: nil,
			},
			args: args{
				e: 1,
			},
			wantLen: 1,
		},
		{
			name: "Non-empty list",
			fields: fields{
				head:   node.NewNode(1, nil),
				length: 1,
			},
			args: args{
				e: 2,
			},
			wantLen: 2,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			nl := &NodeList{
				head:   tc.fields.head,
				length: tc.fields.length,
			}
			nl.Insert(tc.args.e)
			if tc.wantLen != nl.length {
				t.Errorf("NodeList length = %d, wantLen %d", nl.length, tc.wantLen)
			}
		})
	}
}

func TestNodeList_Length(t *testing.T) {
	type fields struct {
		head   *node.Node
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Empty list",
			fields: fields{
				head: nil,
			},
			want: 0,
		},
		{
			name: "Non-empty list",
			fields: fields{
				head:   node.NewNode(1, nil),
				length: 1,
			},
			want: 1,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			n := &NodeList{
				head:   tc.fields.head,
				length: tc.fields.length,
			}
			if got := n.Length(); got != tc.want {
				t.Errorf("NodeList.Length() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestNodeList_String(t *testing.T) {
	type fields struct {
		head   *node.Node
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Empty list",
			fields: fields{
				head: nil,
			},
			want: `<>`,
		},
		{
			name: "One element list",
			fields: fields{
				head:   node.NewNode(1, nil),
				length: 1,
			},
			want: `<1>`,
		},
		{
			name: "More than one element list",
			fields: fields{
				head:   node.NewNode(1, node.NewNode(2, node.NewNode(3, nil))),
				length: 3,
			},
			want: `<1 -> 2 -> 3>`,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			nl := &NodeList{
				head:   tc.fields.head,
				length: tc.fields.length,
			}
			if got := nl.String(); got != tc.want {
				t.Errorf("NodeList.String() = %v, want %v", got, tc.want)
			}
		})
	}
}
