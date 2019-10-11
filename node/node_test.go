package node

import (
	"reflect"
	"testing"
)

func TestNode_Key(t *testing.T) {
	t.Parallel()
	type fields struct {
		key  int
		next *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Key",
			fields: fields{
				key:  1,
				next: nil,
			},
			want: 1,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			n := &Node{
				key:  tc.fields.key,
				next: tc.fields.next,
			}
			if got := n.Key(); got != tc.want {
				t.Errorf("Node.Key() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestNode_Next(t *testing.T) {
	t.Parallel()
	type fields struct {
		key  int
		next *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		{
			name: "Nil Next",
			fields: fields{
				key:  1,
				next: nil,
			},
		},
		{
			name: "Non-nil Next",
			fields: fields{
				key:  1,
				next: &Node{},
			},
			want: &Node{},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			n := &Node{
				key:  tc.fields.key,
				next: tc.fields.next,
			}
			if got := n.Next(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Node.Next() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestNode_SetKey(t *testing.T) {
	t.Parallel()
	type fields struct {
		key  int
		next *Node
	}
	type args struct {
		newKey int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantKey int
	}{
		{
			name: "Set key",
			fields: fields{
				key:  1,
				next: nil,
			},
			args: args{
				newKey: 0,
			},
			wantKey: 0,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			n := &Node{
				key:  tc.fields.key,
				next: tc.fields.next,
			}
			n.SetKey(tc.args.newKey)
			if n.key != tc.wantKey {
				t.Errorf("Node.SetKey() = %v, want %v", n.key, tc.wantKey)
			}
		})
	}
}

func TestNode_SetNext(t *testing.T) {
	t.Parallel()
	type fields struct {
		key  int
		next *Node
	}
	type args struct {
		newNext *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Set next",
			fields: fields{
				key:  1,
				next: nil,
			},
			args: args{newNext: &Node{key: 0, next: nil}},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			n := &Node{
				key:  tc.fields.key,
				next: tc.fields.next,
			}
			n.SetNext(tc.args.newNext)
			if n.next != tc.args.newNext {
				t.Errorf("Node.SetNext() = %v, want %v", n.next, tc.args.newNext)
			}
		})
	}
}
