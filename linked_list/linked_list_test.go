package linked_list

import "testing"

func TestAppend(t *testing.T) {
	cases := []struct {
		desc             string
		list             *linkedList
		newNode          *node
		expectedHeadData int
		expectedTailData int
	}{
		{
			desc:             "blank list",
			list:             &linkedList{},
			newNode:          &node{Data: 3},
			expectedHeadData: 3,
			expectedTailData: 3,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {

			l := tt.list
			l.Append(tt.newNode)

			if l.Head.Data != tt.expectedHeadData {
				t.Fatalf("expected l.Head.Data == %d, got: %d", tt.expectedHeadData, l.Head.Data)
			}
			if l.Tail.Data != tt.expectedTailData {
				t.Fatalf("expected l.Hail.Data == %d, got: %d", tt.expectedTailData, l.Tail.Data)
			}
		})
	}

}
