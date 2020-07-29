package domains

import (
	"context"
	"testing"
)

// t.Error(f) はテスト失敗としてログを出すが、以降の処理も実行される
// t.Fatal(f) はテスト失敗としてログを出し、以降の処理は実行されない
func Test_todoRepository(t *testing.T) {
	ctx := context.Background()
	repo := NewTodoRepository()

	// Create
	todoA, err := repo.Create(ctx, &Todo{
		Text: "test A",
	})
	if err != nil {
		t.Fatal(err)
	}
	if v := todoA.Done; v {
		t.Errorf("unexpected: %#v", v)
	}
	if v := todoA.DoneAt; !v.IsZero() {
		t.Errorf("unexpected: %#v", v)
	}
	if v := todoA.CreatedAt; v.IsZero() {
		t.Errorf("unexpected: %#v", v)
	}

	// Update
	todoANew, err := repo.Update(ctx, &Todo{
		ID:   todoA.ID,
		Text: "test A!",
		Done: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if v := todoANew.Done; !v {
		t.Errorf("unexpected: %#v", v)
	}
	if v := todoANew.DoneAt; v.IsZero() {
		t.Errorf("unexpected: %#v", v)
	}

	// Get
	todoA2, err := repo.Get(ctx, todoA.ID)
	if err != nil {
		t.Fatal(err)
	}
	if v1, v2 := todoA.ID, todoA2.ID; v1 != v2 {
		t.Errorf("unexpected: %#v, %#v", v1, v2)
	}

	todoB, err := repo.Create(ctx, &Todo{
		Text: "test B",
	})
	if err != nil {
		t.Fatal(err)
	}

	// GetAll
	list, err := repo.GetAll(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if v := len(list); v != 2 {
		t.Fatalf("unexpected: %#v", v)
	}
	if v := list[0]; v.ID != todoB.ID {
		t.Errorf("unexpected: %#v", v)
	}
	if v := list[1]; v.ID != todoA.ID {
		t.Errorf("unexpected: %#v", v)
	}
}
