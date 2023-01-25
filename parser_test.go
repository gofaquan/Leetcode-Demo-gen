package parser

import (
	"reflect"
	"testing"
)

func Test_initParser(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *parser
	}{
		{
			"test nil string",
			args{s: ""},
			&parser{
				s:    "",
				data: make([]string, 1, 1),
			}},

		{
			"test  [ string",
			args{s: "["},
			&parser{
				s:    "",
				data: make([]string, 1, 1),
			}},

		{
			"test  ] string",
			args{s: "["},
			&parser{
				s:    "",
				data: make([]string, 1, 1),
			}},

		{
			"test  [] string",
			args{s: "[]"},
			&parser{
				s:    "",
				data: make([]string, 1, 1),
			}},

		{
			"test no space with comma string",
			args{s: "[1,2,3,4]"},
			&parser{
				s:    "1,2,3,4",
				data: []string{"1", "2", "3", "4"},
			}},

		{
			"test space with comma string",
			args{s: "[1 ,2,3 ,4]"},
			&parser{
				s:    "1,2,3,4",
				data: []string{"1", "2", "3", "4"},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initParser(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtoList(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test no space with comma string",
			args{s: "[1,2,3,3,4,5]"},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},

		{
			"test space with comma string",
			args{s: "[1, 2, 3, 3, 4, 5]"},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &ListNode{}
			if AtoList(tt.args.s, got); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AtoList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtoBinaryTree(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"test no space with comma string",
			args{s: "[7,3,15,null,null,9,20]"},
			&TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   20,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},

		{
			"test space with comma string",
			args{s: "[7, 3, 15, null, null, 9, 20]"},
			&TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   20,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &TreeNode{}
			if AtoBinaryTree(tt.args.s, got); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AtoBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
