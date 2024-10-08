package database

import (
	"fmt"
	"testing"
)

func TestDefaultPagination(t *testing.T) {

	orderQuery := PagingBy(0, 0, "", "")

	if orderQuery == "" {
		t.Error("err query is empty")
		t.Fail()
	}

	if orderQuery != "created_at desc" {
		t.Error("err query was wrong")
		t.Fail()
	}

}

func TestPaginationPageAndLimitPassed(t *testing.T) {
	order := "created_at desc"

	orderQuery := PagingBy(1, 30, "", "")

	expected := fmt.Sprintf("ORDER BY %s LIMIT %d OFFSET %d", order, 30, 0)

	if orderQuery != expected {
		t.Error("err query was wrong:", orderQuery, ", expected: ", expected)
		t.Fail()
	}

	orderQuery = PagingBy(3, 30, "", "")

	expected = fmt.Sprintf("ORDER BY %s LIMIT %d OFFSET %d", order, 30, 60)

	if orderQuery != expected {
		t.Error("err query was wrong:", orderQuery, ", expected: ", expected)
		t.Fail()
	}
}

func TestPaginationPageAndLimitPassedWrong(t *testing.T) {
	expected := "created_at desc"
	orderQuery := PagingBy(-1, 30, "", "")

	if orderQuery != expected {
		t.Error("err query was wrong:", orderQuery, ", expected: ", expected)
		t.Fail()
	}

	orderQuery = PagingBy(1, -30, "", "")

	if orderQuery != expected {
		t.Error("err query was wrong:", orderQuery, ", expected: ", expected)
		t.Fail()
	}
}

func TestPaginationOrder(t *testing.T) {
	orderQuery := PagingBy(0, 0, "asc", "")

	if orderQuery != "created_at asc" {
		t.Error("err query was wrong, expected: ", orderQuery)
		t.Fail()
	}

	orderQuery = PagingBy(0, 0, "", "updated_at")

	if orderQuery != "updated_at desc" {
		t.Error("err query was wrong, expected: ", orderQuery)
		t.Fail()
	}
}
