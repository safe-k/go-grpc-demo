package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"google.golang.org/grpc"

	"github.com/go-chi/chi"
	"github.com/safe-k/go-grpc-demo/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewCalculatorClient(conn)

	r := chi.NewRouter()
	r.Get("/add/{a}/{b}", callClient(client.Add))
	r.Get("/multiply/{a}/{b}", callClient(client.Multiply))

	panic(http.ListenAndServe(":8080", r))
}

type clientMethod func(ctx context.Context, in *proto.Request, opts ...grpc.CallOption) (*proto.Response, error)

func callClient(f clientMethod) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a, err := strconv.Atoi(chi.URLParam(r, "a"))
		if err != nil {
			panic(err)
		}

		b, err := strconv.Atoi(chi.URLParam(r, "b"))
		if err != nil {
			panic(err)
		}

		req := &proto.Request{A: int64(a), B: int64(b)}
		res, err := f(r.Context(), req)
		if err != nil {
			panic(err)
		}

		marshal, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		_, err = w.Write(marshal)
		if err != nil {
			panic(err)
		}
	}
}
