package app

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/berquerant/weaver-pokemon-type/app/GetEffectivityListByAttackQuery",
		Iface: reflect.TypeOf((*GetEffectivityListByAttackQuery)(nil)).Elem(),
		New:   func() any { return &getEffectivityListByAttackQuery{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return getEffectivityListByAttackQuery_local_stub{impl: impl.(GetEffectivityListByAttackQuery), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return getEffectivityListByAttackQuery_client_stub{stub: stub, getEffectivityListByAttackMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/berquerant/weaver-pokemon-type/app/GetEffectivityListByAttackQuery", Method: "GetEffectivityListByAttack"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return getEffectivityListByAttackQuery_server_stub{impl: impl.(GetEffectivityListByAttackQuery), addLoad: addLoad}
		},
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/berquerant/weaver-pokemon-type/app/GetEffectivityListByDefenseListQuery",
		Iface: reflect.TypeOf((*GetEffectivityListByDefenseListQuery)(nil)).Elem(),
		New:   func() any { return &getEffectivityListByDefenseListQuery{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return getEffectivityListByDefenseListQuery_local_stub{impl: impl.(GetEffectivityListByDefenseListQuery), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return getEffectivityListByDefenseListQuery_client_stub{stub: stub, getEffectivityListByDefenseListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/berquerant/weaver-pokemon-type/app/GetEffectivityListByDefenseListQuery", Method: "GetEffectivityListByDefenseList"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return getEffectivityListByDefenseListQuery_server_stub{impl: impl.(GetEffectivityListByDefenseListQuery), addLoad: addLoad}
		},
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/berquerant/weaver-pokemon-type/app/GetTypeByNameQuery",
		Iface: reflect.TypeOf((*GetTypeByNameQuery)(nil)).Elem(),
		New:   func() any { return &getTypeByNameQuery{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return getTypeByNameQuery_local_stub{impl: impl.(GetTypeByNameQuery), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return getTypeByNameQuery_client_stub{stub: stub, getTypeByNameMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/berquerant/weaver-pokemon-type/app/GetTypeByNameQuery", Method: "GetTypeByName"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return getTypeByNameQuery_server_stub{impl: impl.(GetTypeByNameQuery), addLoad: addLoad}
		},
	})
}

// Local stub implementations.

type getEffectivityListByAttackQuery_local_stub struct {
	impl   GetEffectivityListByAttackQuery
	tracer trace.Tracer
}

func (s getEffectivityListByAttackQuery_local_stub) GetEffectivityListByAttack(ctx context.Context, a0 PileIndex, a1 int) (r0 []PiledEffectivity, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "app.GetEffectivityListByAttackQuery.GetEffectivityListByAttack", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetEffectivityListByAttack(ctx, a0, a1)
}

type getEffectivityListByDefenseListQuery_local_stub struct {
	impl   GetEffectivityListByDefenseListQuery
	tracer trace.Tracer
}

func (s getEffectivityListByDefenseListQuery_local_stub) GetEffectivityListByDefenseList(ctx context.Context, a0 DefenseTypeIDList) (r0 []PiledEffectivity, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "app.GetEffectivityListByDefenseListQuery.GetEffectivityListByDefenseList", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetEffectivityListByDefenseList(ctx, a0)
}

type getTypeByNameQuery_local_stub struct {
	impl   GetTypeByNameQuery
	tracer trace.Tracer
}

func (s getTypeByNameQuery_local_stub) GetTypeByName(ctx context.Context, a0 string) (r0 *Type, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "app.GetTypeByNameQuery.GetTypeByName", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetTypeByName(ctx, a0)
}

// Client stub implementations.

type getEffectivityListByAttackQuery_client_stub struct {
	stub                              codegen.Stub
	getEffectivityListByAttackMetrics *codegen.MethodMetrics
}

func (s getEffectivityListByAttackQuery_client_stub) GetEffectivityListByAttack(ctx context.Context, a0 PileIndex, a1 int) (r0 []PiledEffectivity, err error) {
	// Update metrics.
	start := time.Now()
	s.getEffectivityListByAttackMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "app.GetEffectivityListByAttackQuery.GetEffectivityListByAttack", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.getEffectivityListByAttackMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.getEffectivityListByAttackMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int((int)(a0))
	enc.Int(a1)
	var shardKey uint64

	// Call the remote method.
	s.getEffectivityListByAttackMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		return
	}
	s.getEffectivityListByAttackMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_PiledEffectivity_0a8633c8(dec)
	err = dec.Error()
	return
}

type getEffectivityListByDefenseListQuery_client_stub struct {
	stub                                   codegen.Stub
	getEffectivityListByDefenseListMetrics *codegen.MethodMetrics
}

func (s getEffectivityListByDefenseListQuery_client_stub) GetEffectivityListByDefenseList(ctx context.Context, a0 DefenseTypeIDList) (r0 []PiledEffectivity, err error) {
	// Update metrics.
	start := time.Now()
	s.getEffectivityListByDefenseListMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "app.GetEffectivityListByDefenseListQuery.GetEffectivityListByDefenseList", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.getEffectivityListByDefenseListMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.getEffectivityListByDefenseListMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + (len(a0) * 8))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	serviceweaver_enc_slice_int_7c8c8866(enc, ([]int)(a0))
	var shardKey uint64

	// Call the remote method.
	s.getEffectivityListByDefenseListMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		return
	}
	s.getEffectivityListByDefenseListMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_PiledEffectivity_0a8633c8(dec)
	err = dec.Error()
	return
}

type getTypeByNameQuery_client_stub struct {
	stub                 codegen.Stub
	getTypeByNameMetrics *codegen.MethodMetrics
}

func (s getTypeByNameQuery_client_stub) GetTypeByName(ctx context.Context, a0 string) (r0 *Type, err error) {
	// Update metrics.
	start := time.Now()
	s.getTypeByNameMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "app.GetTypeByNameQuery.GetTypeByName", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.getTypeByNameMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.getTypeByNameMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.getTypeByNameMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		return
	}
	s.getTypeByNameMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_ptr_Type_298d5b80(dec)
	err = dec.Error()
	return
}

// Server stub implementations.

type getEffectivityListByAttackQuery_server_stub struct {
	impl    GetEffectivityListByAttackQuery
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s getEffectivityListByAttackQuery_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "GetEffectivityListByAttack":
		return s.getEffectivityListByAttack
	default:
		return nil
	}
}

func (s getEffectivityListByAttackQuery_server_stub) getEffectivityListByAttack(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 PileIndex
	*(*int)(&a0) = dec.Int()
	var a1 int
	a1 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetEffectivityListByAttack(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_PiledEffectivity_0a8633c8(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type getEffectivityListByDefenseListQuery_server_stub struct {
	impl    GetEffectivityListByDefenseListQuery
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s getEffectivityListByDefenseListQuery_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "GetEffectivityListByDefenseList":
		return s.getEffectivityListByDefenseList
	default:
		return nil
	}
}

func (s getEffectivityListByDefenseListQuery_server_stub) getEffectivityListByDefenseList(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 DefenseTypeIDList
	*(*[]int)(&a0) = serviceweaver_dec_slice_int_7c8c8866(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetEffectivityListByDefenseList(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_PiledEffectivity_0a8633c8(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type getTypeByNameQuery_server_stub struct {
	impl    GetTypeByNameQuery
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s getTypeByNameQuery_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "GetTypeByName":
		return s.getTypeByName
	default:
		return nil
	}
}

func (s getTypeByNameQuery_server_stub) getTypeByName(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetTypeByName(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_Type_298d5b80(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = &Effectivity{}

func (x *Effectivity) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Effectivity.WeaverMarshal: nil receiver"))
	}
	enc.Int(x.ID)
	(x.Attack).WeaverMarshal(enc)
	(x.Defense).WeaverMarshal(enc)
	enc.Float32(x.Multiplier)
}

func (x *Effectivity) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Effectivity.WeaverUnmarshal: nil receiver"))
	}
	x.ID = dec.Int()
	(&x.Attack).WeaverUnmarshal(dec)
	(&x.Defense).WeaverUnmarshal(dec)
	x.Multiplier = dec.Float32()
}

var _ codegen.AutoMarshal = &Type{}

func (x *Type) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Type.WeaverMarshal: nil receiver"))
	}
	enc.Int(x.ID)
	enc.String(x.Name)
}

func (x *Type) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Type.WeaverUnmarshal: nil receiver"))
	}
	x.ID = dec.Int()
	x.Name = dec.String()
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_Effectivity_12530f17(enc *codegen.Encoder, arg []Effectivity) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_Effectivity_12530f17(dec *codegen.Decoder) []Effectivity {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]Effectivity, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

func serviceweaver_enc_slice_PiledEffectivity_0a8633c8(enc *codegen.Encoder, arg []PiledEffectivity) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		serviceweaver_enc_slice_Effectivity_12530f17(enc, ([]Effectivity)(arg[i]))
	}
}

func serviceweaver_dec_slice_PiledEffectivity_0a8633c8(dec *codegen.Decoder) []PiledEffectivity {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]PiledEffectivity, n)
	for i := 0; i < n; i++ {
		*(*[]Effectivity)(&res[i]) = serviceweaver_dec_slice_Effectivity_12530f17(dec)
	}
	return res
}

func serviceweaver_enc_slice_int_7c8c8866(enc *codegen.Encoder, arg []int) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.Int(arg[i])
	}
}

func serviceweaver_dec_slice_int_7c8c8866(dec *codegen.Decoder) []int {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = dec.Int()
	}
	return res
}

func serviceweaver_enc_ptr_Type_298d5b80(enc *codegen.Encoder, arg *Type) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		(*arg).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_ptr_Type_298d5b80(dec *codegen.Decoder) *Type {
	if !dec.Bool() {
		return nil
	}
	var res Type
	(&res).WeaverUnmarshal(dec)
	return &res
}
