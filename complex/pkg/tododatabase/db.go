package tododatabase

import (
	memdb "github.com/hashicorp/go-memdb"
	"github.com/thedevelopnik/go-kit-example/complex/pb"
)

func CreateDB() *memdb.MemDB {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"todo": &memdb.TableSchema{
				Name: "todo",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UintFieldIndex{Field: "Id"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	// seed database
	todo1 := &pb.Todo{
		Id:   1,
		Name: "go to the grocery store",
	}
	todo2 := &pb.Todo{
		Id:   2,
		Name: "vacuum",
	}
	todo3 := &pb.Todo{
		Id:   3,
		Name: "learn go",
	}

	txn1 := db.Txn(true)
	if err := txn1.Insert("todo", todo1); err != nil {
		panic(err)
	}
	txn1.Commit()

	txn2 := db.Txn(true)
	if err := txn2.Insert("todo", todo2); err != nil {
		panic(err)
	}
	txn2.Commit()

	txn3 := db.Txn(true)
	if err := txn3.Insert("todo", todo3); err != nil {
		panic(err)
	}
	txn3.Commit()

	return db
}
