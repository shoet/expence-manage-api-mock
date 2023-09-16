package handler

import "net/http"

type ListExpenses struct{}

func (l *ListExpenses) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hoge"))
}
