// Code generated by protoc-gen-go-http. DO NOT EDIT.

package v1

import (
	context "context"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	mux "github.com/gorilla/mux"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(http.Request)
var _ = new(context.Context)
var _ = binding.MapProto
var _ = mux.NewRouter

const _ = http1.SupportPackageIsVersion1

type BlogHandler interface {
	CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleReply, error)

	DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleReply, error)

	GetArticle(context.Context, *GetArticleRequest) (*GetArticleReply, error)

	ListArticle(context.Context, *ListArticleRequest) (*ListArticleReply, error)

	UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleReply, error)
}

func NewBlogHandler(srv BlogHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	r.HandleFunc("/v1/article", func(w http.ResponseWriter, r *http.Request) {
		var in CreateArticleRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateArticle(ctx, req.(*CreateArticleRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*CreateArticleReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	r.HandleFunc("/v1/article/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in UpdateArticleRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.BindVars(mux.Vars(r), &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateArticle(ctx, req.(*UpdateArticleRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*UpdateArticleReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("PUT")

	r.HandleFunc("/v1/article/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in DeleteArticleRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.BindVars(mux.Vars(r), &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteArticle(ctx, req.(*DeleteArticleRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*DeleteArticleReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("DELETE")

	r.HandleFunc("/v1/article/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in GetArticleRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.BindVars(mux.Vars(r), &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArticle(ctx, req.(*GetArticleRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*GetArticleReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/v1/article/", func(w http.ResponseWriter, r *http.Request) {
		var in ListArticleRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListArticle(ctx, req.(*ListArticleRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ListArticleReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	return r
}

type BlogHTTPClient interface {
	CreateArticle(ctx context.Context, req *CreateArticleRequest, opts ...http1.CallOption) (rsp *CreateArticleReply, err error)

	DeleteArticle(ctx context.Context, req *DeleteArticleRequest, opts ...http1.CallOption) (rsp *DeleteArticleReply, err error)

	GetArticle(ctx context.Context, req *GetArticleRequest, opts ...http1.CallOption) (rsp *GetArticleReply, err error)

	ListArticle(ctx context.Context, req *ListArticleRequest, opts ...http1.CallOption) (rsp *ListArticleReply, err error)

	UpdateArticle(ctx context.Context, req *UpdateArticleRequest, opts ...http1.CallOption) (rsp *UpdateArticleReply, err error)
}

type BlogHTTPClientImpl struct {
	cc *http1.Client
}

func NewBlogHTTPClient(client *http1.Client) BlogHTTPClient {
	return &BlogHTTPClientImpl{client}
}

func (c *BlogHTTPClientImpl) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...http1.CallOption) (out *CreateArticleReply, err error) {
	path := binding.EncodePath("POST", "/v1/article", in)
	out = &CreateArticleReply{}

	err = c.cc.Invoke(ctx, path, in, &out, http1.Method("POST"), http1.PathPattern("/v1/article"))

	return
}

func (c *BlogHTTPClientImpl) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...http1.CallOption) (out *DeleteArticleReply, err error) {
	path := binding.EncodePath("DELETE", "/v1/article/{id}", in)
	out = &DeleteArticleReply{}

	err = c.cc.Invoke(ctx, path, nil, &out, http1.Method("DELETE"), http1.PathPattern("/v1/article/{id}"))

	return
}

func (c *BlogHTTPClientImpl) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...http1.CallOption) (out *GetArticleReply, err error) {
	path := binding.EncodePath("GET", "/v1/article/{id}", in)
	out = &GetArticleReply{}

	err = c.cc.Invoke(ctx, path, nil, &out, http1.Method("GET"), http1.PathPattern("/v1/article/{id}"))

	return
}

func (c *BlogHTTPClientImpl) ListArticle(ctx context.Context, in *ListArticleRequest, opts ...http1.CallOption) (out *ListArticleReply, err error) {
	path := binding.EncodePath("GET", "/v1/article/", in)
	out = &ListArticleReply{}

	err = c.cc.Invoke(ctx, path, nil, &out, http1.Method("GET"), http1.PathPattern("/v1/article/"))

	return
}

func (c *BlogHTTPClientImpl) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...http1.CallOption) (out *UpdateArticleReply, err error) {
	path := binding.EncodePath("PUT", "/v1/article/{id}", in)
	out = &UpdateArticleReply{}

	err = c.cc.Invoke(ctx, path, in, &out, http1.Method("PUT"), http1.PathPattern("/v1/article/{id}"))

	return
}
