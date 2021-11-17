package BinaryTree

import (
	"testing"
)

func TestBinaryTree_Delete(t1 *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name    string
		bst     *BinaryTree
		args    args
		wantErr bool
	}{
		{
			name: "success_delete",
			bst:  NewEmptyBinaryTree(),
			args: args{key: 18},
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t *testing.T) {
			err := tt.bst.Insert(0, "BENZ")
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			err = tt.bst.Insert(tt.args.key, "BMW")
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			err = tt.bst.Insert(67, "Ford")
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			err = tt.bst.Insert(45, "Aston")
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			err = tt.bst.Insert(3, "Doge")
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			tt.bst.Delete(tt.args.key)
			find, err := tt.bst.Find(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			if find != "" {
				t.Errorf("got %q, wanted %q", find, "")
			}
		})
	}
}

func TestBinaryTree_Find(t1 *testing.T) {
	type args struct {
		key int
		value string
	}
	tests := []struct {
		name    string
		bst  *BinaryTree
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success_find",
			bst: NewEmptyBinaryTree(),
			args: args{key: 15, value: "BMW"},
			want: "BMW",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			err := tt.bst.Insert(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := tt.bst.Find(tt.args.key)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Insert(t1 *testing.T) {
	type args struct {
		key  int
		data string
	}
	tests := []struct {
		name    string
		bst     *BinaryTree
		args    args
		wantErr bool
	}{
		{
			name: "success_insert",
			bst:  NewEmptyBinaryTree(),
			args: args{key: 7, data: "BENZ"},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t *testing.T) {
			if err := tt.bst.Insert(tt.args.key, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}

			find, err := tt.bst.Find(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}

			if find != tt.args.data {
				t.Errorf("got %q, wanted %q", find, tt.args.data)
			}
		})
	}
}
