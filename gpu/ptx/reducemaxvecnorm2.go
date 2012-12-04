package ptx

//This file is auto-generated. Editing is futile.

func init() { Code["reducemaxvecnorm2"] = REDUCEMAXVECNORM2 }

const REDUCEMAXVECNORM2 = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_000010e5_00000000-9_reducemaxvecnorm2.cpp3.i"
	.file	2 "/home/arne/src/code.google.com/p/nimble-cube/gpu/ptx/reducemaxvecnorm2.cu"
	.file	3 "/usr/local/cuda-5.0/nvvm/ci_include.h"
	.file	4 "/usr/local/cuda/bin/../include/sm_11_atomic_functions.h"
// __cuda_local_var_35270_32_non_const_sdata has been demoted

.visible .entry reducemaxvecnorm2(
	.param .u64 reducemaxvecnorm2_param_0,
	.param .u64 reducemaxvecnorm2_param_1,
	.param .u64 reducemaxvecnorm2_param_2,
	.param .u64 reducemaxvecnorm2_param_3,
	.param .f32 reducemaxvecnorm2_param_4,
	.param .u32 reducemaxvecnorm2_param_5
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<42>;
	.reg .f32 	%f<33>;
	.reg .s64 	%rd<19>;
	// demoted variable
	.shared .align 4 .b8 __cuda_local_var_35270_32_non_const_sdata[2048];

	ld.param.u64 	%rd6, [reducemaxvecnorm2_param_0];
	ld.param.u64 	%rd7, [reducemaxvecnorm2_param_1];
	ld.param.u64 	%rd8, [reducemaxvecnorm2_param_2];
	ld.param.u64 	%rd9, [reducemaxvecnorm2_param_3];
	ld.param.f32 	%f32, [reducemaxvecnorm2_param_4];
	ld.param.u32 	%r9, [reducemaxvecnorm2_param_5];
	cvta.to.global.u64 	%rd1, %rd9;
	cvta.to.global.u64 	%rd2, %rd8;
	cvta.to.global.u64 	%rd3, %rd7;
	cvta.to.global.u64 	%rd4, %rd6;
	.loc 2 10 1
	mov.u32 	%r41, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r40, %r41, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r41, %r11;
	.loc 2 10 1
	setp.ge.s32 	%p1, %r40, %r9;
	@%p1 bra 	BB0_2;

BB0_1:
	.loc 2 10 1
	mul.wide.s32 	%rd10, %r40, 4;
	add.s64 	%rd11, %rd4, %rd10;
	ld.global.f32 	%f5, [%rd11];
	add.s64 	%rd12, %rd3, %rd10;
	ld.global.f32 	%f6, [%rd12];
	mul.f32 	%f7, %f6, %f6;
	fma.rn.f32 	%f8, %f5, %f5, %f7;
	add.s64 	%rd13, %rd2, %rd10;
	ld.global.f32 	%f9, [%rd13];
	fma.rn.f32 	%f10, %f9, %f9, %f8;
	.loc 3 435 5
	max.f32 	%f32, %f32, %f10;
	.loc 2 10 1
	add.s32 	%r40, %r40, %r4;
	.loc 2 10 1
	setp.lt.s32 	%p2, %r40, %r9;
	@%p2 bra 	BB0_1;

BB0_2:
	.loc 2 10 1
	mul.wide.s32 	%rd14, %r2, 4;
	mov.u64 	%rd15, __cuda_local_var_35270_32_non_const_sdata;
	add.s64 	%rd5, %rd15, %rd14;
	st.shared.f32 	[%rd5], %f32;
	bar.sync 	0;
	.loc 2 10 1
	setp.lt.u32 	%p3, %r41, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	.loc 2 10 1
	mov.u32 	%r7, %r41;
	shr.u32 	%r41, %r7, 1;
	.loc 2 10 1
	setp.ge.u32 	%p4, %r2, %r41;
	@%p4 bra 	BB0_5;

	.loc 2 10 1
	ld.shared.f32 	%f11, [%rd5];
	add.s32 	%r17, %r41, %r2;
	mul.wide.u32 	%rd16, %r17, 4;
	add.s64 	%rd18, %rd15, %rd16;
	ld.shared.f32 	%f12, [%rd18];
	.loc 3 435 5
	max.f32 	%f13, %f11, %f12;
	.loc 2 10 1
	st.shared.f32 	[%rd5], %f13;

BB0_5:
	.loc 2 10 1
	bar.sync 	0;
	.loc 2 10 1
	setp.gt.u32 	%p5, %r7, 131;
	@%p5 bra 	BB0_3;

BB0_6:
	.loc 2 10 1
	setp.gt.s32 	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	.loc 2 10 1
	ld.volatile.shared.f32 	%f14, [%rd5];
	ld.volatile.shared.f32 	%f15, [%rd5+128];
	.loc 3 435 5
	max.f32 	%f16, %f14, %f15;
	.loc 2 10 1
	st.volatile.shared.f32 	[%rd5], %f16;
	ld.volatile.shared.f32 	%f17, [%rd5+64];
	ld.volatile.shared.f32 	%f18, [%rd5];
	.loc 3 435 5
	max.f32 	%f19, %f18, %f17;
	.loc 2 10 1
	st.volatile.shared.f32 	[%rd5], %f19;
	ld.volatile.shared.f32 	%f20, [%rd5+32];
	ld.volatile.shared.f32 	%f21, [%rd5];
	.loc 3 435 5
	max.f32 	%f22, %f21, %f20;
	.loc 2 10 1
	st.volatile.shared.f32 	[%rd5], %f22;
	ld.volatile.shared.f32 	%f23, [%rd5+16];
	ld.volatile.shared.f32 	%f24, [%rd5];
	.loc 3 435 5
	max.f32 	%f25, %f24, %f23;
	.loc 2 10 1
	st.volatile.shared.f32 	[%rd5], %f25;
	ld.volatile.shared.f32 	%f26, [%rd5+8];
	ld.volatile.shared.f32 	%f27, [%rd5];
	.loc 3 435 5
	max.f32 	%f28, %f27, %f26;
	.loc 2 10 1
	st.volatile.shared.f32 	[%rd5], %f28;
	ld.volatile.shared.f32 	%f29, [%rd5+4];
	ld.volatile.shared.f32 	%f30, [%rd5];
	.loc 3 435 5
	max.f32 	%f31, %f30, %f29;
	.loc 2 10 1
	st.volatile.shared.f32 	[%rd5], %f31;

BB0_8:
	.loc 2 10 1
	setp.ne.s32 	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	.loc 2 10 1
	ld.shared.u32 	%r38, [__cuda_local_var_35270_32_non_const_sdata];
	.loc 3 1881 5
	atom.global.max.s32 	%r39, [%rd1], %r38;

BB0_10:
	.loc 2 11 2
	ret;
}


`
