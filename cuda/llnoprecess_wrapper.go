package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"sync"
	"unsafe"
)

// CUDA handle for llnoprecess kernel
var llnoprecess_code cu.Function

// Stores the arguments for llnoprecess kernel invocation
type llnoprecess_args_t struct {
	arg_tx unsafe.Pointer
	arg_ty unsafe.Pointer
	arg_tz unsafe.Pointer
	arg_mx unsafe.Pointer
	arg_my unsafe.Pointer
	arg_mz unsafe.Pointer
	arg_hx unsafe.Pointer
	arg_hy unsafe.Pointer
	arg_hz unsafe.Pointer
	arg_N  int
	argptr [10]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for llnoprecess kernel invocation
var llnoprecess_args llnoprecess_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	llnoprecess_args.argptr[0] = unsafe.Pointer(&llnoprecess_args.arg_tx)
	llnoprecess_args.argptr[1] = unsafe.Pointer(&llnoprecess_args.arg_ty)
	llnoprecess_args.argptr[2] = unsafe.Pointer(&llnoprecess_args.arg_tz)
	llnoprecess_args.argptr[3] = unsafe.Pointer(&llnoprecess_args.arg_mx)
	llnoprecess_args.argptr[4] = unsafe.Pointer(&llnoprecess_args.arg_my)
	llnoprecess_args.argptr[5] = unsafe.Pointer(&llnoprecess_args.arg_mz)
	llnoprecess_args.argptr[6] = unsafe.Pointer(&llnoprecess_args.arg_hx)
	llnoprecess_args.argptr[7] = unsafe.Pointer(&llnoprecess_args.arg_hy)
	llnoprecess_args.argptr[8] = unsafe.Pointer(&llnoprecess_args.arg_hz)
	llnoprecess_args.argptr[9] = unsafe.Pointer(&llnoprecess_args.arg_N)
}

// Wrapper for llnoprecess CUDA kernel, asynchronous.
func k_llnoprecess_async(tx unsafe.Pointer, ty unsafe.Pointer, tz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, hx unsafe.Pointer, hy unsafe.Pointer, hz unsafe.Pointer, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
	}

	llnoprecess_args.Lock()
	defer llnoprecess_args.Unlock()

	if llnoprecess_code == 0 {
		llnoprecess_code = fatbinLoad(llnoprecess_map, "llnoprecess")
	}

	llnoprecess_args.arg_tx = tx
	llnoprecess_args.arg_ty = ty
	llnoprecess_args.arg_tz = tz
	llnoprecess_args.arg_mx = mx
	llnoprecess_args.arg_my = my
	llnoprecess_args.arg_mz = mz
	llnoprecess_args.arg_hx = hx
	llnoprecess_args.arg_hy = hy
	llnoprecess_args.arg_hz = hz
	llnoprecess_args.arg_N = N

	args := llnoprecess_args.argptr[:]
	cu.LaunchKernel(llnoprecess_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
	}
}

// maps compute capability on PTX code for llnoprecess kernel.
var llnoprecess_map = map[int]string{0: "",
	20: llnoprecess_ptx_20,
	30: llnoprecess_ptx_30,
	35: llnoprecess_ptx_35}

// llnoprecess PTX code for various compute capabilities.
const (
	llnoprecess_ptx_20 = `
.version 3.2
.target sm_20
.address_size 64


.visible .entry llnoprecess(
	.param .u64 llnoprecess_param_0,
	.param .u64 llnoprecess_param_1,
	.param .u64 llnoprecess_param_2,
	.param .u64 llnoprecess_param_3,
	.param .u64 llnoprecess_param_4,
	.param .u64 llnoprecess_param_5,
	.param .u64 llnoprecess_param_6,
	.param .u64 llnoprecess_param_7,
	.param .u64 llnoprecess_param_8,
	.param .u32 llnoprecess_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<28>;
	.reg .s64 	%rd<29>;


	ld.param.u64 	%rd10, [llnoprecess_param_0];
	ld.param.u64 	%rd11, [llnoprecess_param_1];
	ld.param.u64 	%rd12, [llnoprecess_param_2];
	ld.param.u64 	%rd13, [llnoprecess_param_3];
	ld.param.u64 	%rd14, [llnoprecess_param_4];
	ld.param.u64 	%rd15, [llnoprecess_param_5];
	ld.param.u64 	%rd16, [llnoprecess_param_6];
	ld.param.u64 	%rd17, [llnoprecess_param_7];
	ld.param.u64 	%rd18, [llnoprecess_param_8];
	ld.param.u32 	%r2, [llnoprecess_param_9];
	cvta.to.global.u64 	%rd1, %rd12;
	cvta.to.global.u64 	%rd2, %rd11;
	cvta.to.global.u64 	%rd3, %rd10;
	cvta.to.global.u64 	%rd4, %rd18;
	cvta.to.global.u64 	%rd5, %rd17;
	cvta.to.global.u64 	%rd6, %rd16;
	cvta.to.global.u64 	%rd7, %rd15;
	cvta.to.global.u64 	%rd8, %rd14;
	cvta.to.global.u64 	%rd9, %rd13;
	.loc 1 10 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 1 11 1
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	mul.wide.s32 	%rd19, %r1, 4;
	add.s64 	%rd20, %rd9, %rd19;
	add.s64 	%rd21, %rd8, %rd19;
	add.s64 	%rd22, %rd7, %rd19;
	add.s64 	%rd23, %rd6, %rd19;
	add.s64 	%rd24, %rd5, %rd19;
	add.s64 	%rd25, %rd4, %rd19;
	.loc 1 14 1
	ld.global.f32 	%f1, [%rd25];
	.loc 1 13 1
	ld.global.f32 	%f2, [%rd21];
	.loc 1 16 1
	mul.f32 	%f3, %f2, %f1;
	.loc 1 14 1
	ld.global.f32 	%f4, [%rd24];
	.loc 1 13 1
	ld.global.f32 	%f5, [%rd22];
	.loc 1 16 1
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	.loc 1 14 1
	ld.global.f32 	%f8, [%rd23];
	.loc 1 16 1
	mul.f32 	%f9, %f5, %f8;
	.loc 1 13 1
	ld.global.f32 	%f10, [%rd20];
	.loc 1 16 1
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	.loc 1 17 1
	mul.f32 	%f16, %f2, %f15;
	mul.f32 	%f17, %f5, %f12;
	sub.f32 	%f18, %f16, %f17;
	mul.f32 	%f19, %f5, %f7;
	mul.f32 	%f20, %f10, %f15;
	sub.f32 	%f21, %f19, %f20;
	mul.f32 	%f22, %f10, %f12;
	mul.f32 	%f23, %f2, %f7;
	sub.f32 	%f24, %f22, %f23;
	neg.f32 	%f25, %f18;
	neg.f32 	%f26, %f21;
	neg.f32 	%f27, %f24;
	add.s64 	%rd26, %rd3, %rd19;
	.loc 1 19 1
	st.global.f32 	[%rd26], %f25;
	add.s64 	%rd27, %rd2, %rd19;
	.loc 1 20 1
	st.global.f32 	[%rd27], %f26;
	add.s64 	%rd28, %rd1, %rd19;
	.loc 1 21 1
	st.global.f32 	[%rd28], %f27;

BB0_2:
	.loc 1 23 2
	ret;
}


`
	llnoprecess_ptx_30 = `
.version 3.2
.target sm_30
.address_size 64


.visible .entry llnoprecess(
	.param .u64 llnoprecess_param_0,
	.param .u64 llnoprecess_param_1,
	.param .u64 llnoprecess_param_2,
	.param .u64 llnoprecess_param_3,
	.param .u64 llnoprecess_param_4,
	.param .u64 llnoprecess_param_5,
	.param .u64 llnoprecess_param_6,
	.param .u64 llnoprecess_param_7,
	.param .u64 llnoprecess_param_8,
	.param .u32 llnoprecess_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<28>;
	.reg .s64 	%rd<29>;


	ld.param.u64 	%rd10, [llnoprecess_param_0];
	ld.param.u64 	%rd11, [llnoprecess_param_1];
	ld.param.u64 	%rd12, [llnoprecess_param_2];
	ld.param.u64 	%rd13, [llnoprecess_param_3];
	ld.param.u64 	%rd14, [llnoprecess_param_4];
	ld.param.u64 	%rd15, [llnoprecess_param_5];
	ld.param.u64 	%rd16, [llnoprecess_param_6];
	ld.param.u64 	%rd17, [llnoprecess_param_7];
	ld.param.u64 	%rd18, [llnoprecess_param_8];
	ld.param.u32 	%r2, [llnoprecess_param_9];
	cvta.to.global.u64 	%rd1, %rd12;
	cvta.to.global.u64 	%rd2, %rd11;
	cvta.to.global.u64 	%rd3, %rd10;
	cvta.to.global.u64 	%rd4, %rd18;
	cvta.to.global.u64 	%rd5, %rd17;
	cvta.to.global.u64 	%rd6, %rd16;
	cvta.to.global.u64 	%rd7, %rd15;
	cvta.to.global.u64 	%rd8, %rd14;
	cvta.to.global.u64 	%rd9, %rd13;
	.loc 1 10 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 1 11 1
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	mul.wide.s32 	%rd19, %r1, 4;
	add.s64 	%rd20, %rd9, %rd19;
	add.s64 	%rd21, %rd8, %rd19;
	add.s64 	%rd22, %rd7, %rd19;
	add.s64 	%rd23, %rd6, %rd19;
	add.s64 	%rd24, %rd5, %rd19;
	add.s64 	%rd25, %rd4, %rd19;
	.loc 1 14 1
	ld.global.f32 	%f1, [%rd25];
	.loc 1 13 1
	ld.global.f32 	%f2, [%rd21];
	.loc 1 16 1
	mul.f32 	%f3, %f2, %f1;
	.loc 1 14 1
	ld.global.f32 	%f4, [%rd24];
	.loc 1 13 1
	ld.global.f32 	%f5, [%rd22];
	.loc 1 16 1
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	.loc 1 14 1
	ld.global.f32 	%f8, [%rd23];
	.loc 1 16 1
	mul.f32 	%f9, %f5, %f8;
	.loc 1 13 1
	ld.global.f32 	%f10, [%rd20];
	.loc 1 16 1
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	.loc 1 17 1
	mul.f32 	%f16, %f2, %f15;
	mul.f32 	%f17, %f5, %f12;
	sub.f32 	%f18, %f16, %f17;
	mul.f32 	%f19, %f5, %f7;
	mul.f32 	%f20, %f10, %f15;
	sub.f32 	%f21, %f19, %f20;
	mul.f32 	%f22, %f10, %f12;
	mul.f32 	%f23, %f2, %f7;
	sub.f32 	%f24, %f22, %f23;
	neg.f32 	%f25, %f18;
	neg.f32 	%f26, %f21;
	neg.f32 	%f27, %f24;
	add.s64 	%rd26, %rd3, %rd19;
	.loc 1 19 1
	st.global.f32 	[%rd26], %f25;
	add.s64 	%rd27, %rd2, %rd19;
	.loc 1 20 1
	st.global.f32 	[%rd27], %f26;
	add.s64 	%rd28, %rd1, %rd19;
	.loc 1 21 1
	st.global.f32 	[%rd28], %f27;

BB0_2:
	.loc 1 23 2
	ret;
}


`
	llnoprecess_ptx_35 = `
.version 3.2
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry llnoprecess(
	.param .u64 llnoprecess_param_0,
	.param .u64 llnoprecess_param_1,
	.param .u64 llnoprecess_param_2,
	.param .u64 llnoprecess_param_3,
	.param .u64 llnoprecess_param_4,
	.param .u64 llnoprecess_param_5,
	.param .u64 llnoprecess_param_6,
	.param .u64 llnoprecess_param_7,
	.param .u64 llnoprecess_param_8,
	.param .u32 llnoprecess_param_9
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<28>;
	.reg .s64 	%rd<29>;


	ld.param.u64 	%rd10, [llnoprecess_param_0];
	ld.param.u64 	%rd11, [llnoprecess_param_1];
	ld.param.u64 	%rd12, [llnoprecess_param_2];
	ld.param.u64 	%rd13, [llnoprecess_param_3];
	ld.param.u64 	%rd14, [llnoprecess_param_4];
	ld.param.u64 	%rd15, [llnoprecess_param_5];
	ld.param.u64 	%rd16, [llnoprecess_param_6];
	ld.param.u64 	%rd17, [llnoprecess_param_7];
	ld.param.u64 	%rd18, [llnoprecess_param_8];
	ld.param.u32 	%r2, [llnoprecess_param_9];
	cvta.to.global.u64 	%rd1, %rd12;
	cvta.to.global.u64 	%rd2, %rd11;
	cvta.to.global.u64 	%rd3, %rd10;
	cvta.to.global.u64 	%rd4, %rd18;
	cvta.to.global.u64 	%rd5, %rd17;
	cvta.to.global.u64 	%rd6, %rd16;
	cvta.to.global.u64 	%rd7, %rd15;
	cvta.to.global.u64 	%rd8, %rd14;
	cvta.to.global.u64 	%rd9, %rd13;
	.loc 1 10 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 1 11 1
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB2_2;

	mul.wide.s32 	%rd19, %r1, 4;
	add.s64 	%rd20, %rd9, %rd19;
	add.s64 	%rd21, %rd8, %rd19;
	add.s64 	%rd22, %rd7, %rd19;
	add.s64 	%rd23, %rd6, %rd19;
	add.s64 	%rd24, %rd5, %rd19;
	add.s64 	%rd25, %rd4, %rd19;
	.loc 1 14 1
	ld.global.nc.f32 	%f1, [%rd25];
	.loc 1 13 1
	ld.global.nc.f32 	%f2, [%rd21];
	.loc 1 16 1
	mul.f32 	%f3, %f2, %f1;
	.loc 1 14 1
	ld.global.nc.f32 	%f4, [%rd24];
	.loc 1 13 1
	ld.global.nc.f32 	%f5, [%rd22];
	.loc 1 16 1
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	.loc 1 14 1
	ld.global.nc.f32 	%f8, [%rd23];
	.loc 1 16 1
	mul.f32 	%f9, %f5, %f8;
	.loc 1 13 1
	ld.global.nc.f32 	%f10, [%rd20];
	.loc 1 16 1
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	.loc 1 17 1
	mul.f32 	%f16, %f2, %f15;
	mul.f32 	%f17, %f5, %f12;
	sub.f32 	%f18, %f16, %f17;
	mul.f32 	%f19, %f5, %f7;
	mul.f32 	%f20, %f10, %f15;
	sub.f32 	%f21, %f19, %f20;
	mul.f32 	%f22, %f10, %f12;
	mul.f32 	%f23, %f2, %f7;
	sub.f32 	%f24, %f22, %f23;
	neg.f32 	%f25, %f18;
	neg.f32 	%f26, %f21;
	neg.f32 	%f27, %f24;
	add.s64 	%rd26, %rd3, %rd19;
	.loc 1 19 1
	st.global.f32 	[%rd26], %f25;
	add.s64 	%rd27, %rd2, %rd19;
	.loc 1 20 1
	st.global.f32 	[%rd27], %f26;
	add.s64 	%rd28, %rd1, %rd19;
	.loc 1 21 1
	st.global.f32 	[%rd28], %f27;

BB2_2:
	.loc 1 23 2
	ret;
}


`
)