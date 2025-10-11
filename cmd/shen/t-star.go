package main

import . "github.com/tiancaiamao/shen-go/kl"

var TStarMain = MakeNative(func(__e *ControlFlow) {
tmp13925 := MakeNative(func(__e *ControlFlow) {
V4488 := __e.Get(1)
_ = V4488
V4489 := __e.Get(2)
_ = V4489
tmp13926 := MakeNative(func(__e *ControlFlow) {
W4490 := __e.Get(1)
_ = W4490
tmp13927 := MakeNative(func(__e *ControlFlow) {
W4491 := __e.Get(1)
_ = W4491
tmp13928 := MakeNative(func(__e *ControlFlow) {
W4492 := __e.Get(1)
_ = W4492
tmp13929 := MakeNative(func(__e *ControlFlow) {
Z4493 := __e.Get(1)
_ = Z4493
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4494 := __e.Get(1)
_ = Z4494
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4495 := __e.Get(1)
_ = Z4495
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4496 := __e.Get(1)
_ = Z4496
tmp13930 := MakeNative(func(__e *ControlFlow) {
W4497 := __e.Get(1)
_ = W4497
tmp13931 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13931

tmp13932 := Call(__e, PrimFunc(symshen_4deref), W4490, Z4493)


tmp13933 := Call(__e, PrimFunc(symreceive), tmp13932)


tmp13934 := Call(__e, PrimFunc(symshen_4deref), W4491, Z4493)


tmp13935 := Call(__e, PrimFunc(symreceive), tmp13934)


tmp13936 := MakeNative(func(__e *ControlFlow) {
tmp13937 := Call(__e, PrimFunc(symshen_4deref), W4492, Z4493)


tmp13938 := Call(__e, PrimFunc(symreceive), tmp13937)


tmp13939 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symreturn), W4497, Z4493, Z4494, Z4495, Z4496)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4toplevel_1forms), tmp13938, W4497, Z4493, Z4494, Z4495, tmp13939)
return


}, 0)

tmp13940 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), tmp13933, tmp13935, W4497, Z4493, Z4494, Z4495, tmp13936)


__e.TailApply(PrimFunc(symshen_4gc), Z4493, tmp13940)
return


}, 1)

tmp13941 := Call(__e, PrimFunc(symshen_4newpv), Z4493)


__e.TailApply(tmp13930, tmp13941)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp13942 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp13943 := Call(__e, tmp13929, tmp13942)


tmp13944 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp13945 := Call(__e, PrimFunc(sym_8v), MakeNumber(0), tmp13944)


tmp13946 := Call(__e, PrimFunc(sym_8v), True, tmp13945)


tmp13947 := Call(__e, tmp13943, tmp13946)


tmp13948 := Call(__e, tmp13947, MakeNumber(0))


tmp13949 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

__e.TailApply(tmp13948, tmp13949)
return


}, 1)

tmp13950 := Call(__e, PrimFunc(symshen_4curry), V4488)


__e.TailApply(tmp13928, tmp13950)
return


}, 1)

tmp13951 := Call(__e, PrimFunc(symshen_4rectify_1type), V4489)


__e.TailApply(tmp13927, tmp13951)
return


}, 1)

tmp13952 := Call(__e, PrimFunc(symshen_4extract_1vars), V4489)


__e.TailApply(tmp13926, tmp13952)
return


}, 2)

tmp13953 := Call(__e, ns2_1set, symshen_4typecheck, tmp13925)


_ = tmp13953

tmp13954 := MakeNative(func(__e *ControlFlow) {
V4498 := __e.Get(1)
_ = V4498
V4499 := __e.Get(2)
_ = V4499
V4500 := __e.Get(3)
_ = V4500
V4501 := __e.Get(4)
_ = V4501
V4502 := __e.Get(5)
_ = V4502
V4503 := __e.Get(6)
_ = V4503
V4504 := __e.Get(7)
_ = V4504
tmp13955 := MakeNative(func(__e *ControlFlow) {
W4505 := __e.Get(1)
_ = W4505
tmp13973 := PrimEqual(W4505, False)

if True == tmp13973 {
tmp13971 := Call(__e, PrimFunc(symshen_4unlocked_2), V4502)


if True == tmp13971 {
tmp13956 := MakeNative(func(__e *ControlFlow) {
W4507 := __e.Get(1)
_ = W4507
tmp13968 := PrimIsPair(W4507)

if True == tmp13968 {
tmp13957 := MakeNative(func(__e *ControlFlow) {
W4508 := __e.Get(1)
_ = W4508
tmp13958 := MakeNative(func(__e *ControlFlow) {
W4509 := __e.Get(1)
_ = W4509
tmp13959 := MakeNative(func(__e *ControlFlow) {
W4510 := __e.Get(1)
_ = W4510
tmp13960 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13960

tmp13961 := Call(__e, PrimFunc(symshen_4deref), W4510, V4501)


tmp13962 := Call(__e, PrimFunc(symsubst), tmp13961, W4508, V4499)


tmp13963 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), W4509, tmp13962, V4500, V4501, V4502, V4503, V4504)


__e.TailApply(PrimFunc(symshen_4gc), V4501, tmp13963)
return


}, 1)

tmp13964 := Call(__e, PrimFunc(symshen_4newpv), V4501)


__e.TailApply(tmp13959, tmp13964)
return


}, 1)

tmp13965 := PrimTail(W4507)

__e.TailApply(tmp13958, tmp13965)
return


}, 1)

tmp13966 := PrimHead(W4507)

__e.TailApply(tmp13957, tmp13966)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13969 := Call(__e, PrimFunc(symshen_4lazyderef), V4498, V4501)


__e.TailApply(tmp13956, tmp13969)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4505)
return
}


}, 1)

tmp13981 := Call(__e, PrimFunc(symshen_4unlocked_2), V4502)


var ifres13974 Obj

if True == tmp13981 {
tmp13975 := MakeNative(func(__e *ControlFlow) {
W4506 := __e.Get(1)
_ = W4506
tmp13978 := PrimEqual(W4506, Nil)

if True == tmp13978 {
tmp13976 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13976

__e.TailApply(PrimFunc(symis_b), V4499, V4500, V4501, V4502, V4503, V4504)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13979 := Call(__e, PrimFunc(symshen_4lazyderef), V4498, V4501)


tmp13980 := Call(__e, tmp13975, tmp13979)


ifres13974 = tmp13980


} else {
ifres13974 = False


}

__e.TailApply(tmp13955, ifres13974)
return


}, 7)

tmp13982 := Call(__e, ns2_1set, symshen_4insert_1prolog_1variables, tmp13954)


_ = tmp13982

tmp13983 := MakeNative(func(__e *ControlFlow) {
V4511 := __e.Get(1)
_ = V4511
V4512 := __e.Get(2)
_ = V4512
V4513 := __e.Get(3)
_ = V4513
V4514 := __e.Get(4)
_ = V4514
V4515 := __e.Get(5)
_ = V4515
V4516 := __e.Get(6)
_ = V4516
tmp13984 := MakeNative(func(__e *ControlFlow) {
W4517 := __e.Get(1)
_ = W4517
tmp13985 := MakeNative(func(__e *ControlFlow) {
W4518 := __e.Get(1)
_ = W4518
tmp13998 := PrimEqual(W4518, False)

if True == tmp13998 {
tmp13986 := MakeNative(func(__e *ControlFlow) {
W4524 := __e.Get(1)
_ = W4524
tmp13988 := PrimEqual(W4524, False)

if True == tmp13988 {
__e.TailApply(PrimFunc(symshen_4unlock), V4514, W4517)
return
} else {
__e.Return(W4524)
return
}


}, 1)

tmp13996 := Call(__e, PrimFunc(symshen_4unlocked_2), V4514)


var ifres13989 Obj

if True == tmp13996 {
tmp13990 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13990

tmp13991 := PrimIntern(MakeString(":"))

tmp13992 := PrimCons(V4512, Nil)

tmp13993 := PrimCons(tmp13991, tmp13992)

tmp13994 := PrimCons(V4511, tmp13993)

tmp13995 := Call(__e, PrimFunc(symshen_4system_1S), tmp13994, Nil, V4513, V4514, W4517, V4516)


ifres13989 = tmp13995


} else {
ifres13989 = False


}

__e.TailApply(tmp13986, ifres13989)
return


} else {
__e.Return(W4518)
return
}


}, 1)

tmp14027 := Call(__e, PrimFunc(symshen_4unlocked_2), V4514)


var ifres13999 Obj

if True == tmp14027 {
tmp14000 := MakeNative(func(__e *ControlFlow) {
W4519 := __e.Get(1)
_ = W4519
tmp14024 := PrimIsPair(W4519)

if True == tmp14024 {
tmp14001 := MakeNative(func(__e *ControlFlow) {
W4520 := __e.Get(1)
_ = W4520
tmp14020 := PrimEqual(W4520, symdefine)

if True == tmp14020 {
tmp14002 := MakeNative(func(__e *ControlFlow) {
W4521 := __e.Get(1)
_ = W4521
tmp14016 := PrimIsPair(W4521)

if True == tmp14016 {
tmp14003 := MakeNative(func(__e *ControlFlow) {
W4522 := __e.Get(1)
_ = W4522
tmp14004 := MakeNative(func(__e *ControlFlow) {
W4523 := __e.Get(1)
_ = W4523
tmp14005 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14005

tmp14006 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp14007 := MakeNative(func(__e *ControlFlow) {
tmp14008 := MakeNative(func(__e *ControlFlow) {
tmp14009 := PrimValue(symshen_4_dspy_d)

tmp14010 := MakeNative(func(__e *ControlFlow) {
tmp14011 := PrimCons(W4522, W4523)

tmp14012 := PrimCons(symdefine, tmp14011)

__e.TailApply(PrimFunc(symshen_4t_d), tmp14012, V4512, V4513, V4514, W4517, V4516)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4signal_1def), tmp14009, W4522, V4513, V4514, W4517, tmp14010)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4513, V4514, W4517, tmp14008)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14006, V4513, V4514, W4517, tmp14007)
return


}, 1)

tmp14013 := PrimTail(W4521)

__e.TailApply(tmp14004, tmp14013)
return


}, 1)

tmp14014 := PrimHead(W4521)

__e.TailApply(tmp14003, tmp14014)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14017 := PrimTail(W4519)

tmp14018 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14017, V4513)


__e.TailApply(tmp14002, tmp14018)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14021 := PrimHead(W4519)

tmp14022 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14021, V4513)


__e.TailApply(tmp14001, tmp14022)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14025 := Call(__e, PrimFunc(symshen_4lazyderef), V4511, V4513)


tmp14026 := Call(__e, tmp14000, tmp14025)


ifres13999 = tmp14026


} else {
ifres13999 = False


}

__e.TailApply(tmp13985, ifres13999)
return


}, 1)

tmp14028 := PrimNumberAdd(V4515, MakeNumber(1))

__e.TailApply(tmp13984, tmp14028)
return


}, 6)

tmp14029 := Call(__e, ns2_1set, symshen_4toplevel_1forms, tmp13983)


_ = tmp14029

tmp14030 := MakeNative(func(__e *ControlFlow) {
V4525 := __e.Get(1)
_ = V4525
V4526 := __e.Get(2)
_ = V4526
V4527 := __e.Get(3)
_ = V4527
V4528 := __e.Get(4)
_ = V4528
V4529 := __e.Get(5)
_ = V4529
V4530 := __e.Get(6)
_ = V4530
tmp14031 := MakeNative(func(__e *ControlFlow) {
W4531 := __e.Get(1)
_ = W4531
tmp14048 := PrimEqual(W4531, False)

if True == tmp14048 {
tmp14046 := Call(__e, PrimFunc(symshen_4unlocked_2), V4528)


if True == tmp14046 {
tmp14032 := MakeNative(func(__e *ControlFlow) {
W4533 := __e.Get(1)
_ = W4533
tmp14043 := PrimEqual(W4533, True)

if True == tmp14043 {
tmp14033 := MakeNative(func(__e *ControlFlow) {
W4534 := __e.Get(1)
_ = W4534
tmp14034 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14034

tmp14035 := Call(__e, PrimFunc(symshen_4deref), V4526, V4527)


tmp14036 := Call(__e, PrimFunc(symshen_4app), tmp14035, MakeString(")\n"), symshen_4a)


tmp14037 := PrimStringConcat(MakeString("\ntypechecking (fn "), tmp14036)

tmp14038 := Call(__e, PrimFunc(symstoutput))


tmp14039 := Call(__e, PrimFunc(sympr), tmp14037, tmp14038)


tmp14040 := Call(__e, PrimFunc(symis), W4534, tmp14039, V4527, V4528, V4529, V4530)


__e.TailApply(PrimFunc(symshen_4gc), V4527, tmp14040)
return


}, 1)

tmp14041 := Call(__e, PrimFunc(symshen_4newpv), V4527)


__e.TailApply(tmp14033, tmp14041)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14044 := Call(__e, PrimFunc(symshen_4lazyderef), V4525, V4527)


__e.TailApply(tmp14032, tmp14044)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4531)
return
}


}, 1)

tmp14056 := Call(__e, PrimFunc(symshen_4unlocked_2), V4528)


var ifres14049 Obj

if True == tmp14056 {
tmp14050 := MakeNative(func(__e *ControlFlow) {
W4532 := __e.Get(1)
_ = W4532
tmp14053 := PrimEqual(W4532, False)

if True == tmp14053 {
tmp14051 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14051

__e.TailApply(PrimFunc(symthaw), V4530)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14054 := Call(__e, PrimFunc(symshen_4lazyderef), V4525, V4527)


tmp14055 := Call(__e, tmp14050, tmp14054)


ifres14049 = tmp14055


} else {
ifres14049 = False


}

__e.TailApply(tmp14031, ifres14049)
return


}, 6)

tmp14057 := Call(__e, ns2_1set, symshen_4signal_1def, tmp14030)


_ = tmp14057

tmp14058 := MakeNative(func(__e *ControlFlow) {
V4535 := __e.Get(1)
_ = V4535
tmp14059 := Call(__e, PrimFunc(symshen_4curry_1type), V4535)


__e.TailApply(PrimFunc(symshen_4demodulate), tmp14059)
return


}, 1)

tmp14060 := Call(__e, ns2_1set, symshen_4rectify_1type, tmp14058)


_ = tmp14060

tmp14061 := MakeNative(func(__e *ControlFlow) {
V4536 := __e.Get(1)
_ = V4536
tmp14062 := MakeNative(func(__e *ControlFlow) {
tmp14063 := MakeNative(func(__e *ControlFlow) {
W4537 := __e.Get(1)
_ = W4537
tmp14065 := PrimEqual(W4537, V4536)

if True == tmp14065 {
__e.Return(V4536)
return
} else {
__e.TailApply(PrimFunc(symshen_4demodulate), W4537)
return
}


}, 1)

tmp14066 := MakeNative(func(__e *ControlFlow) {
Z4538 := __e.Get(1)
_ = Z4538
__e.TailApply(PrimFunc(symshen_4demod), Z4538)
return
}, 1)

tmp14067 := Call(__e, PrimFunc(symshen_4walk), tmp14066, V4536)


__e.TailApply(tmp14063, tmp14067)
return


}, 0)

tmp14068 := MakeNative(func(__e *ControlFlow) {
Z4539 := __e.Get(1)
_ = Z4539
__e.Return(V4536)
return
}, 1)

__e.TailApply(try_1catch, tmp14062, tmp14068)
return


}, 1)

tmp14069 := Call(__e, ns2_1set, symshen_4demodulate, tmp14061)


_ = tmp14069

tmp14070 := MakeNative(func(__e *ControlFlow) {
V4540 := __e.Get(1)
_ = V4540
tmp14194 := PrimIsPair(V4540)

var ifres14167 Obj

if True == tmp14194 {
tmp14192 := PrimTail(V4540)

tmp14193 := PrimIsPair(tmp14192)

var ifres14169 Obj

if True == tmp14193 {
tmp14189 := PrimTail(V4540)

tmp14190 := PrimHead(tmp14189)

tmp14191 := PrimEqual(sym_1_1_6, tmp14190)

var ifres14171 Obj

if True == tmp14191 {
tmp14186 := PrimTail(V4540)

tmp14187 := PrimTail(tmp14186)

tmp14188 := PrimIsPair(tmp14187)

var ifres14173 Obj

if True == tmp14188 {
tmp14182 := PrimTail(V4540)

tmp14183 := PrimTail(tmp14182)

tmp14184 := PrimTail(tmp14183)

tmp14185 := PrimIsPair(tmp14184)

var ifres14175 Obj

if True == tmp14185 {
tmp14177 := PrimTail(V4540)

tmp14178 := PrimTail(tmp14177)

tmp14179 := PrimTail(tmp14178)

tmp14180 := PrimHead(tmp14179)

tmp14181 := PrimEqual(sym_1_1_6, tmp14180)

var ifres14176 Obj

if True == tmp14181 {
ifres14176 = True


} else {
ifres14176 = False


}

ifres14175 = ifres14176


} else {
ifres14175 = False


}

var ifres14174 Obj

if True == ifres14175 {
ifres14174 = True


} else {
ifres14174 = False


}

ifres14173 = ifres14174


} else {
ifres14173 = False


}

var ifres14172 Obj

if True == ifres14173 {
ifres14172 = True


} else {
ifres14172 = False


}

ifres14171 = ifres14172


} else {
ifres14171 = False


}

var ifres14170 Obj

if True == ifres14171 {
ifres14170 = True


} else {
ifres14170 = False


}

ifres14169 = ifres14170


} else {
ifres14169 = False


}

var ifres14168 Obj

if True == ifres14169 {
ifres14168 = True


} else {
ifres14168 = False


}

ifres14167 = ifres14168


} else {
ifres14167 = False


}

if True == ifres14167 {
tmp14071 := PrimHead(V4540)

tmp14072 := PrimTail(V4540)

tmp14073 := PrimTail(tmp14072)

tmp14074 := PrimCons(tmp14073, Nil)

tmp14075 := PrimCons(sym_1_1_6, tmp14074)

tmp14076 := PrimCons(tmp14071, tmp14075)

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14076)
return


} else {
tmp14165 := PrimIsPair(V4540)

var ifres14125 Obj

if True == tmp14165 {
tmp14163 := PrimHead(V4540)

tmp14164 := PrimIsPair(tmp14163)

var ifres14127 Obj

if True == tmp14164 {
tmp14160 := PrimHead(V4540)

tmp14161 := PrimHead(tmp14160)

tmp14162 := PrimEqual(symlist, tmp14161)

var ifres14129 Obj

if True == tmp14162 {
tmp14157 := PrimHead(V4540)

tmp14158 := PrimTail(tmp14157)

tmp14159 := PrimIsPair(tmp14158)

var ifres14131 Obj

if True == tmp14159 {
tmp14153 := PrimHead(V4540)

tmp14154 := PrimTail(tmp14153)

tmp14155 := PrimTail(tmp14154)

tmp14156 := PrimEqual(Nil, tmp14155)

var ifres14133 Obj

if True == tmp14156 {
tmp14151 := PrimTail(V4540)

tmp14152 := PrimIsPair(tmp14151)

var ifres14135 Obj

if True == tmp14152 {
tmp14148 := PrimTail(V4540)

tmp14149 := PrimHead(tmp14148)

tmp14150 := PrimEqual(sym_a_a_6, tmp14149)

var ifres14137 Obj

if True == tmp14150 {
tmp14145 := PrimTail(V4540)

tmp14146 := PrimTail(tmp14145)

tmp14147 := PrimIsPair(tmp14146)

var ifres14139 Obj

if True == tmp14147 {
tmp14141 := PrimTail(V4540)

tmp14142 := PrimTail(tmp14141)

tmp14143 := PrimTail(tmp14142)

tmp14144 := PrimEqual(Nil, tmp14143)

var ifres14140 Obj

if True == tmp14144 {
ifres14140 = True


} else {
ifres14140 = False


}

ifres14139 = ifres14140


} else {
ifres14139 = False


}

var ifres14138 Obj

if True == ifres14139 {
ifres14138 = True


} else {
ifres14138 = False


}

ifres14137 = ifres14138


} else {
ifres14137 = False


}

var ifres14136 Obj

if True == ifres14137 {
ifres14136 = True


} else {
ifres14136 = False


}

ifres14135 = ifres14136


} else {
ifres14135 = False


}

var ifres14134 Obj

if True == ifres14135 {
ifres14134 = True


} else {
ifres14134 = False


}

ifres14133 = ifres14134


} else {
ifres14133 = False


}

var ifres14132 Obj

if True == ifres14133 {
ifres14132 = True


} else {
ifres14132 = False


}

ifres14131 = ifres14132


} else {
ifres14131 = False


}

var ifres14130 Obj

if True == ifres14131 {
ifres14130 = True


} else {
ifres14130 = False


}

ifres14129 = ifres14130


} else {
ifres14129 = False


}

var ifres14128 Obj

if True == ifres14129 {
ifres14128 = True


} else {
ifres14128 = False


}

ifres14127 = ifres14128


} else {
ifres14127 = False


}

var ifres14126 Obj

if True == ifres14127 {
ifres14126 = True


} else {
ifres14126 = False


}

ifres14125 = ifres14126


} else {
ifres14125 = False


}

if True == ifres14125 {
tmp14077 := PrimHead(V4540)

tmp14078 := PrimHead(V4540)

tmp14079 := PrimTail(V4540)

tmp14080 := PrimTail(tmp14079)

tmp14081 := PrimCons(tmp14078, tmp14080)

tmp14082 := PrimCons(symstr, tmp14081)

tmp14083 := PrimCons(tmp14082, Nil)

tmp14084 := PrimCons(sym_1_1_6, tmp14083)

tmp14085 := PrimCons(tmp14077, tmp14084)

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14085)
return


} else {
tmp14123 := PrimIsPair(V4540)

var ifres14096 Obj

if True == tmp14123 {
tmp14121 := PrimTail(V4540)

tmp14122 := PrimIsPair(tmp14121)

var ifres14098 Obj

if True == tmp14122 {
tmp14118 := PrimTail(V4540)

tmp14119 := PrimHead(tmp14118)

tmp14120 := PrimEqual(sym_d, tmp14119)

var ifres14100 Obj

if True == tmp14120 {
tmp14115 := PrimTail(V4540)

tmp14116 := PrimTail(tmp14115)

tmp14117 := PrimIsPair(tmp14116)

var ifres14102 Obj

if True == tmp14117 {
tmp14111 := PrimTail(V4540)

tmp14112 := PrimTail(tmp14111)

tmp14113 := PrimTail(tmp14112)

tmp14114 := PrimIsPair(tmp14113)

var ifres14104 Obj

if True == tmp14114 {
tmp14106 := PrimTail(V4540)

tmp14107 := PrimTail(tmp14106)

tmp14108 := PrimTail(tmp14107)

tmp14109 := PrimHead(tmp14108)

tmp14110 := PrimEqual(sym_d, tmp14109)

var ifres14105 Obj

if True == tmp14110 {
ifres14105 = True


} else {
ifres14105 = False


}

ifres14104 = ifres14105


} else {
ifres14104 = False


}

var ifres14103 Obj

if True == ifres14104 {
ifres14103 = True


} else {
ifres14103 = False


}

ifres14102 = ifres14103


} else {
ifres14102 = False


}

var ifres14101 Obj

if True == ifres14102 {
ifres14101 = True


} else {
ifres14101 = False


}

ifres14100 = ifres14101


} else {
ifres14100 = False


}

var ifres14099 Obj

if True == ifres14100 {
ifres14099 = True


} else {
ifres14099 = False


}

ifres14098 = ifres14099


} else {
ifres14098 = False


}

var ifres14097 Obj

if True == ifres14098 {
ifres14097 = True


} else {
ifres14097 = False


}

ifres14096 = ifres14097


} else {
ifres14096 = False


}

if True == ifres14096 {
tmp14086 := PrimHead(V4540)

tmp14087 := PrimTail(V4540)

tmp14088 := PrimTail(tmp14087)

tmp14089 := PrimCons(tmp14088, Nil)

tmp14090 := PrimCons(sym_d, tmp14089)

tmp14091 := PrimCons(tmp14086, tmp14090)

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14091)
return


} else {
tmp14094 := PrimIsPair(V4540)

if True == tmp14094 {
tmp14092 := MakeNative(func(__e *ControlFlow) {
Z4541 := __e.Get(1)
_ = Z4541
__e.TailApply(PrimFunc(symshen_4curry_1type), Z4541)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp14092, V4540)
return


} else {
__e.Return(V4540)
return
}


}


}


}


}, 1)

tmp14195 := Call(__e, ns2_1set, symshen_4curry_1type, tmp14070)


_ = tmp14195

tmp14196 := MakeNative(func(__e *ControlFlow) {
V4542 := __e.Get(1)
_ = V4542
tmp14314 := PrimIsPair(V4542)

var ifres14306 Obj

if True == tmp14314 {
tmp14312 := PrimHead(V4542)

tmp14313 := PrimEqual(symdefine, tmp14312)

var ifres14308 Obj

if True == tmp14313 {
tmp14310 := PrimTail(V4542)

tmp14311 := PrimIsPair(tmp14310)

var ifres14309 Obj

if True == tmp14311 {
ifres14309 = True


} else {
ifres14309 = False


}

ifres14308 = ifres14309


} else {
ifres14308 = False


}

var ifres14307 Obj

if True == ifres14308 {
ifres14307 = True


} else {
ifres14307 = False


}

ifres14306 = ifres14307


} else {
ifres14306 = False


}

if True == ifres14306 {
__e.Return(V4542)
return
} else {
tmp14304 := PrimIsPair(V4542)

var ifres14285 Obj

if True == tmp14304 {
tmp14302 := PrimHead(V4542)

tmp14303 := PrimEqual(symtype, tmp14302)

var ifres14287 Obj

if True == tmp14303 {
tmp14300 := PrimTail(V4542)

tmp14301 := PrimIsPair(tmp14300)

var ifres14289 Obj

if True == tmp14301 {
tmp14297 := PrimTail(V4542)

tmp14298 := PrimTail(tmp14297)

tmp14299 := PrimIsPair(tmp14298)

var ifres14291 Obj

if True == tmp14299 {
tmp14293 := PrimTail(V4542)

tmp14294 := PrimTail(tmp14293)

tmp14295 := PrimTail(tmp14294)

tmp14296 := PrimEqual(Nil, tmp14295)

var ifres14292 Obj

if True == tmp14296 {
ifres14292 = True


} else {
ifres14292 = False


}

ifres14291 = ifres14292


} else {
ifres14291 = False


}

var ifres14290 Obj

if True == ifres14291 {
ifres14290 = True


} else {
ifres14290 = False


}

ifres14289 = ifres14290


} else {
ifres14289 = False


}

var ifres14288 Obj

if True == ifres14289 {
ifres14288 = True


} else {
ifres14288 = False


}

ifres14287 = ifres14288


} else {
ifres14287 = False


}

var ifres14286 Obj

if True == ifres14287 {
ifres14286 = True


} else {
ifres14286 = False


}

ifres14285 = ifres14286


} else {
ifres14285 = False


}

if True == ifres14285 {
tmp14197 := PrimTail(V4542)

tmp14198 := PrimHead(tmp14197)

tmp14199 := Call(__e, PrimFunc(symshen_4curry), tmp14198)


tmp14200 := PrimTail(V4542)

tmp14201 := PrimTail(tmp14200)

tmp14202 := PrimCons(tmp14199, tmp14201)

__e.Return(PrimCons(symtype, tmp14202))
return


} else {
tmp14283 := PrimIsPair(V4542)

var ifres14264 Obj

if True == tmp14283 {
tmp14281 := PrimHead(V4542)

tmp14282 := PrimEqual(syminput_7, tmp14281)

var ifres14266 Obj

if True == tmp14282 {
tmp14279 := PrimTail(V4542)

tmp14280 := PrimIsPair(tmp14279)

var ifres14268 Obj

if True == tmp14280 {
tmp14276 := PrimTail(V4542)

tmp14277 := PrimTail(tmp14276)

tmp14278 := PrimIsPair(tmp14277)

var ifres14270 Obj

if True == tmp14278 {
tmp14272 := PrimTail(V4542)

tmp14273 := PrimTail(tmp14272)

tmp14274 := PrimTail(tmp14273)

tmp14275 := PrimEqual(Nil, tmp14274)

var ifres14271 Obj

if True == tmp14275 {
ifres14271 = True


} else {
ifres14271 = False


}

ifres14270 = ifres14271


} else {
ifres14270 = False


}

var ifres14269 Obj

if True == ifres14270 {
ifres14269 = True


} else {
ifres14269 = False


}

ifres14268 = ifres14269


} else {
ifres14268 = False


}

var ifres14267 Obj

if True == ifres14268 {
ifres14267 = True


} else {
ifres14267 = False


}

ifres14266 = ifres14267


} else {
ifres14266 = False


}

var ifres14265 Obj

if True == ifres14266 {
ifres14265 = True


} else {
ifres14265 = False


}

ifres14264 = ifres14265


} else {
ifres14264 = False


}

if True == ifres14264 {
tmp14203 := PrimTail(V4542)

tmp14204 := PrimHead(tmp14203)

tmp14205 := PrimTail(V4542)

tmp14206 := PrimTail(tmp14205)

tmp14207 := PrimHead(tmp14206)

tmp14208 := Call(__e, PrimFunc(symshen_4curry), tmp14207)


tmp14209 := PrimCons(tmp14208, Nil)

tmp14210 := PrimCons(tmp14204, tmp14209)

__e.Return(PrimCons(syminput_7, tmp14210))
return


} else {
tmp14262 := PrimIsPair(V4542)

var ifres14258 Obj

if True == tmp14262 {
tmp14260 := PrimHead(V4542)

tmp14261 := Call(__e, PrimFunc(symshen_4special_2), tmp14260)


var ifres14259 Obj

if True == tmp14261 {
ifres14259 = True


} else {
ifres14259 = False


}

ifres14258 = ifres14259


} else {
ifres14258 = False


}

if True == ifres14258 {
tmp14211 := PrimHead(V4542)

tmp14212 := MakeNative(func(__e *ControlFlow) {
Z4543 := __e.Get(1)
_ = Z4543
__e.TailApply(PrimFunc(symshen_4curry), Z4543)
return
}, 1)

tmp14213 := PrimTail(V4542)

tmp14214 := Call(__e, PrimFunc(symmap), tmp14212, tmp14213)


__e.Return(PrimCons(tmp14211, tmp14214))
return


} else {
tmp14256 := PrimIsPair(V4542)

var ifres14252 Obj

if True == tmp14256 {
tmp14254 := PrimHead(V4542)

tmp14255 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp14254)


var ifres14253 Obj

if True == tmp14255 {
ifres14253 = True


} else {
ifres14253 = False


}

ifres14252 = ifres14253


} else {
ifres14252 = False


}

if True == ifres14252 {
__e.Return(V4542)
return
} else {
tmp14250 := PrimIsPair(V4542)

var ifres14241 Obj

if True == tmp14250 {
tmp14248 := PrimTail(V4542)

tmp14249 := PrimIsPair(tmp14248)

var ifres14243 Obj

if True == tmp14249 {
tmp14245 := PrimTail(V4542)

tmp14246 := PrimTail(tmp14245)

tmp14247 := PrimIsPair(tmp14246)

var ifres14244 Obj

if True == tmp14247 {
ifres14244 = True


} else {
ifres14244 = False


}

ifres14243 = ifres14244


} else {
ifres14243 = False


}

var ifres14242 Obj

if True == ifres14243 {
ifres14242 = True


} else {
ifres14242 = False


}

ifres14241 = ifres14242


} else {
ifres14241 = False


}

if True == ifres14241 {
tmp14215 := PrimHead(V4542)

tmp14216 := PrimTail(V4542)

tmp14217 := PrimHead(tmp14216)

tmp14218 := PrimCons(tmp14217, Nil)

tmp14219 := PrimCons(tmp14215, tmp14218)

tmp14220 := PrimTail(V4542)

tmp14221 := PrimTail(tmp14220)

tmp14222 := PrimCons(tmp14219, tmp14221)

__e.TailApply(PrimFunc(symshen_4curry), tmp14222)
return


} else {
tmp14239 := PrimIsPair(V4542)

var ifres14230 Obj

if True == tmp14239 {
tmp14237 := PrimTail(V4542)

tmp14238 := PrimIsPair(tmp14237)

var ifres14232 Obj

if True == tmp14238 {
tmp14234 := PrimTail(V4542)

tmp14235 := PrimTail(tmp14234)

tmp14236 := PrimEqual(Nil, tmp14235)

var ifres14233 Obj

if True == tmp14236 {
ifres14233 = True


} else {
ifres14233 = False


}

ifres14232 = ifres14233


} else {
ifres14232 = False


}

var ifres14231 Obj

if True == ifres14232 {
ifres14231 = True


} else {
ifres14231 = False


}

ifres14230 = ifres14231


} else {
ifres14230 = False


}

if True == ifres14230 {
tmp14223 := PrimHead(V4542)

tmp14224 := Call(__e, PrimFunc(symshen_4curry), tmp14223)


tmp14225 := PrimTail(V4542)

tmp14226 := PrimHead(tmp14225)

tmp14227 := Call(__e, PrimFunc(symshen_4curry), tmp14226)


tmp14228 := PrimCons(tmp14227, Nil)

__e.Return(PrimCons(tmp14224, tmp14228))
return


} else {
__e.Return(V4542)
return
}


}


}


}


}


}


}


}, 1)

tmp14315 := Call(__e, ns2_1set, symshen_4curry, tmp14196)


_ = tmp14315

tmp14316 := MakeNative(func(__e *ControlFlow) {
V4544 := __e.Get(1)
_ = V4544
tmp14317 := PrimValue(symshen_4_dspecial_d)

__e.TailApply(PrimFunc(symelement_2), V4544, tmp14317)
return


}, 1)

tmp14318 := Call(__e, ns2_1set, symshen_4special_2, tmp14316)


_ = tmp14318

tmp14319 := MakeNative(func(__e *ControlFlow) {
V4545 := __e.Get(1)
_ = V4545
tmp14320 := PrimValue(symshen_4_dextraspecial_d)

__e.TailApply(PrimFunc(symelement_2), V4545, tmp14320)
return


}, 1)

tmp14321 := Call(__e, ns2_1set, symshen_4extraspecial_2, tmp14319)


_ = tmp14321

tmp14322 := MakeNative(func(__e *ControlFlow) {
V4546 := __e.Get(1)
_ = V4546
V4547 := __e.Get(2)
_ = V4547
V4548 := __e.Get(3)
_ = V4548
V4549 := __e.Get(4)
_ = V4549
V4550 := __e.Get(5)
_ = V4550
V4551 := __e.Get(6)
_ = V4551
tmp14323 := MakeNative(func(__e *ControlFlow) {
W4552 := __e.Get(1)
_ = W4552
tmp14324 := MakeNative(func(__e *ControlFlow) {
W4553 := __e.Get(1)
_ = W4553
tmp14382 := PrimEqual(W4553, False)

if True == tmp14382 {
tmp14325 := MakeNative(func(__e *ControlFlow) {
W4554 := __e.Get(1)
_ = W4554
tmp14344 := PrimEqual(W4554, False)

if True == tmp14344 {
tmp14326 := MakeNative(func(__e *ControlFlow) {
W4562 := __e.Get(1)
_ = W4562
tmp14336 := PrimEqual(W4562, False)

if True == tmp14336 {
tmp14327 := MakeNative(func(__e *ControlFlow) {
W4563 := __e.Get(1)
_ = W4563
tmp14329 := PrimEqual(W4563, False)

if True == tmp14329 {
__e.TailApply(PrimFunc(symshen_4unlock), V4549, W4552)
return
} else {
__e.Return(W4563)
return
}


}, 1)

tmp14334 := Call(__e, PrimFunc(symshen_4unlocked_2), V4549)


var ifres14330 Obj

if True == tmp14334 {
tmp14331 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14331

tmp14332 := PrimValue(symshen_4_ddatatypes_d)

tmp14333 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), V4546, V4547, tmp14332, V4548, V4549, W4552, V4551)


ifres14330 = tmp14333


} else {
ifres14330 = False


}

__e.TailApply(tmp14327, ifres14330)
return


} else {
__e.Return(W4562)
return
}


}, 1)

tmp14342 := Call(__e, PrimFunc(symshen_4unlocked_2), V4549)


var ifres14337 Obj

if True == tmp14342 {
tmp14338 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14338

tmp14339 := PrimValue(symshen_4_dspy_d)

tmp14340 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4show), V4546, V4547, V4548, V4549, W4552, V4551)
return
}, 0)

tmp14341 := Call(__e, PrimFunc(symwhen), tmp14339, V4548, V4549, W4552, tmp14340)


ifres14337 = tmp14341


} else {
ifres14337 = False


}

__e.TailApply(tmp14326, ifres14337)
return


} else {
__e.Return(W4554)
return
}


}, 1)

tmp14380 := Call(__e, PrimFunc(symshen_4unlocked_2), V4549)


var ifres14345 Obj

if True == tmp14380 {
tmp14346 := MakeNative(func(__e *ControlFlow) {
W4555 := __e.Get(1)
_ = W4555
tmp14377 := PrimIsPair(W4555)

if True == tmp14377 {
tmp14347 := MakeNative(func(__e *ControlFlow) {
W4556 := __e.Get(1)
_ = W4556
tmp14348 := MakeNative(func(__e *ControlFlow) {
W4557 := __e.Get(1)
_ = W4557
tmp14372 := PrimIsPair(W4557)

if True == tmp14372 {
tmp14349 := MakeNative(func(__e *ControlFlow) {
W4558 := __e.Get(1)
_ = W4558
tmp14350 := MakeNative(func(__e *ControlFlow) {
W4559 := __e.Get(1)
_ = W4559
tmp14367 := PrimIsPair(W4559)

if True == tmp14367 {
tmp14351 := MakeNative(func(__e *ControlFlow) {
W4560 := __e.Get(1)
_ = W4560
tmp14352 := MakeNative(func(__e *ControlFlow) {
W4561 := __e.Get(1)
_ = W4561
tmp14362 := PrimEqual(W4561, Nil)

if True == tmp14362 {
tmp14353 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14353

tmp14354 := Call(__e, PrimFunc(symshen_4deref), W4558, V4548)


tmp14355 := PrimIntern(MakeString(":"))

tmp14356 := PrimEqual(tmp14354, tmp14355)

tmp14357 := MakeNative(func(__e *ControlFlow) {
tmp14358 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp14359 := MakeNative(func(__e *ControlFlow) {
tmp14360 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4556, W4560, V4547, V4548, V4549, W4552, V4551)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4548, V4549, W4552, tmp14360)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14358, V4548, V4549, W4552, tmp14359)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14356, V4548, V4549, W4552, tmp14357)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14363 := PrimTail(W4559)

tmp14364 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14363, V4548)


__e.TailApply(tmp14352, tmp14364)
return


}, 1)

tmp14365 := PrimHead(W4559)

__e.TailApply(tmp14351, tmp14365)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14368 := PrimTail(W4557)

tmp14369 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14368, V4548)


__e.TailApply(tmp14350, tmp14369)
return


}, 1)

tmp14370 := PrimHead(W4557)

__e.TailApply(tmp14349, tmp14370)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14373 := PrimTail(W4555)

tmp14374 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14373, V4548)


__e.TailApply(tmp14348, tmp14374)
return


}, 1)

tmp14375 := PrimHead(W4555)

__e.TailApply(tmp14347, tmp14375)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14378 := Call(__e, PrimFunc(symshen_4lazyderef), V4546, V4548)


tmp14379 := Call(__e, tmp14346, tmp14378)


ifres14345 = tmp14379


} else {
ifres14345 = False


}

__e.TailApply(tmp14325, ifres14345)
return


} else {
__e.Return(W4553)
return
}


}, 1)

tmp14387 := Call(__e, PrimFunc(symshen_4unlocked_2), V4549)


var ifres14383 Obj

if True == tmp14387 {
tmp14384 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14384

tmp14385 := Call(__e, PrimFunc(symshen_4maxinfexceeded_2))


tmp14386 := Call(__e, PrimFunc(symwhen), tmp14385, V4548, V4549, W4552, V4551)


ifres14383 = tmp14386


} else {
ifres14383 = False


}

__e.TailApply(tmp14324, ifres14383)
return


}, 1)

tmp14388 := PrimNumberAdd(V4550, MakeNumber(1))

__e.TailApply(tmp14323, tmp14388)
return


}, 6)

tmp14389 := Call(__e, ns2_1set, symshen_4system_1S, tmp14322)


_ = tmp14389

tmp14390 := MakeNative(func(__e *ControlFlow) {
V4570 := __e.Get(1)
_ = V4570
V4571 := __e.Get(2)
_ = V4571
V4572 := __e.Get(3)
_ = V4572
V4573 := __e.Get(4)
_ = V4573
V4574 := __e.Get(5)
_ = V4574
V4575 := __e.Get(6)
_ = V4575
tmp14391 := Call(__e, PrimFunc(symshen_4line))


_ = tmp14391

tmp14392 := Call(__e, PrimFunc(symshen_4deref), V4570, V4572)


tmp14393 := Call(__e, PrimFunc(symshen_4show_1p), tmp14392)


_ = tmp14393

tmp14394 := Call(__e, PrimFunc(symnl), MakeNumber(2))


_ = tmp14394

tmp14395 := Call(__e, PrimFunc(symshen_4deref), V4571, V4572)


tmp14396 := Call(__e, PrimFunc(symshen_4show_1assumptions), tmp14395, MakeNumber(1))


_ = tmp14396

tmp14397 := Call(__e, PrimFunc(symshen_4pause_1for_1user))


_ = tmp14397

__e.Return(False)
return


}, 6)

tmp14398 := Call(__e, ns2_1set, symshen_4show, tmp14390)


_ = tmp14398

tmp14399 := MakeNative(func(__e *ControlFlow) {
tmp14400 := MakeNative(func(__e *ControlFlow) {
W4576 := __e.Get(1)
_ = W4576
tmp14402 := PrimEqual(MakeNumber(1), W4576)

var ifres14401 Obj

if True == tmp14402 {
ifres14401 = MakeString("")


} else {
ifres14401 = MakeString("s")


}

tmp14403 := Call(__e, PrimFunc(symshen_4app), ifres14401, MakeString(" \n?- "), symshen_4a)


tmp14404 := PrimStringConcat(MakeString(" inference"), tmp14403)

tmp14405 := Call(__e, PrimFunc(symshen_4app), W4576, tmp14404, symshen_4a)


tmp14406 := PrimStringConcat(MakeString("____________________________________________________________ "), tmp14405)

tmp14407 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14406, tmp14407)
return


}, 1)

tmp14408 := Call(__e, PrimFunc(syminferences))


__e.TailApply(tmp14400, tmp14408)
return


}, 0)

tmp14409 := Call(__e, ns2_1set, symshen_4line, tmp14399)


_ = tmp14409

tmp14410 := MakeNative(func(__e *ControlFlow) {
V4577 := __e.Get(1)
_ = V4577
tmp14442 := PrimIsPair(V4577)

var ifres14421 Obj

if True == tmp14442 {
tmp14440 := PrimTail(V4577)

tmp14441 := PrimIsPair(tmp14440)

var ifres14423 Obj

if True == tmp14441 {
tmp14437 := PrimTail(V4577)

tmp14438 := PrimTail(tmp14437)

tmp14439 := PrimIsPair(tmp14438)

var ifres14425 Obj

if True == tmp14439 {
tmp14433 := PrimTail(V4577)

tmp14434 := PrimTail(tmp14433)

tmp14435 := PrimTail(tmp14434)

tmp14436 := PrimEqual(Nil, tmp14435)

var ifres14427 Obj

if True == tmp14436 {
tmp14429 := PrimTail(V4577)

tmp14430 := PrimHead(tmp14429)

tmp14431 := PrimIntern(MakeString(":"))

tmp14432 := PrimEqual(tmp14430, tmp14431)

var ifres14428 Obj

if True == tmp14432 {
ifres14428 = True


} else {
ifres14428 = False


}

ifres14427 = ifres14428


} else {
ifres14427 = False


}

var ifres14426 Obj

if True == ifres14427 {
ifres14426 = True


} else {
ifres14426 = False


}

ifres14425 = ifres14426


} else {
ifres14425 = False


}

var ifres14424 Obj

if True == ifres14425 {
ifres14424 = True


} else {
ifres14424 = False


}

ifres14423 = ifres14424


} else {
ifres14423 = False


}

var ifres14422 Obj

if True == ifres14423 {
ifres14422 = True


} else {
ifres14422 = False


}

ifres14421 = ifres14422


} else {
ifres14421 = False


}

if True == ifres14421 {
tmp14411 := PrimHead(V4577)

tmp14412 := Call(__e, PrimFunc(symshen_4prterm), tmp14411)


_ = tmp14412

tmp14413 := Call(__e, PrimFunc(symstoutput))


tmp14414 := Call(__e, PrimFunc(sympr), MakeString(" : "), tmp14413)


_ = tmp14414

tmp14415 := PrimTail(V4577)

tmp14416 := PrimTail(tmp14415)

tmp14417 := PrimHead(tmp14416)

tmp14418 := Call(__e, PrimFunc(symshen_4app), tmp14417, MakeString(""), symshen_4r)


tmp14419 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14418, tmp14419)
return


} else {
__e.TailApply(PrimFunc(symshen_4prterm), V4577)
return
}


}, 1)

tmp14443 := Call(__e, ns2_1set, symshen_4show_1p, tmp14410)


_ = tmp14443

tmp14444 := MakeNative(func(__e *ControlFlow) {
V4578 := __e.Get(1)
_ = V4578
tmp14487 := PrimIsPair(V4578)

var ifres14468 Obj

if True == tmp14487 {
tmp14485 := PrimHead(V4578)

tmp14486 := PrimEqual(symcons, tmp14485)

var ifres14470 Obj

if True == tmp14486 {
tmp14483 := PrimTail(V4578)

tmp14484 := PrimIsPair(tmp14483)

var ifres14472 Obj

if True == tmp14484 {
tmp14480 := PrimTail(V4578)

tmp14481 := PrimTail(tmp14480)

tmp14482 := PrimIsPair(tmp14481)

var ifres14474 Obj

if True == tmp14482 {
tmp14476 := PrimTail(V4578)

tmp14477 := PrimTail(tmp14476)

tmp14478 := PrimTail(tmp14477)

tmp14479 := PrimEqual(Nil, tmp14478)

var ifres14475 Obj

if True == tmp14479 {
ifres14475 = True


} else {
ifres14475 = False


}

ifres14474 = ifres14475


} else {
ifres14474 = False


}

var ifres14473 Obj

if True == ifres14474 {
ifres14473 = True


} else {
ifres14473 = False


}

ifres14472 = ifres14473


} else {
ifres14472 = False


}

var ifres14471 Obj

if True == ifres14472 {
ifres14471 = True


} else {
ifres14471 = False


}

ifres14470 = ifres14471


} else {
ifres14470 = False


}

var ifres14469 Obj

if True == ifres14470 {
ifres14469 = True


} else {
ifres14469 = False


}

ifres14468 = ifres14469


} else {
ifres14468 = False


}

if True == ifres14468 {
tmp14445 := Call(__e, PrimFunc(symstoutput))


tmp14446 := Call(__e, PrimFunc(sympr), MakeString("["), tmp14445)


_ = tmp14446

tmp14447 := PrimTail(V4578)

tmp14448 := PrimHead(tmp14447)

tmp14449 := Call(__e, PrimFunc(symshen_4prterm), tmp14448)


_ = tmp14449

tmp14450 := PrimTail(V4578)

tmp14451 := PrimTail(tmp14450)

tmp14452 := PrimHead(tmp14451)

tmp14453 := Call(__e, PrimFunc(symshen_4prtl), tmp14452)


_ = tmp14453

tmp14454 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("]"), tmp14454)
return


} else {
tmp14466 := PrimIsPair(V4578)

if True == tmp14466 {
tmp14455 := Call(__e, PrimFunc(symstoutput))


tmp14456 := Call(__e, PrimFunc(sympr), MakeString("("), tmp14455)


_ = tmp14456

tmp14457 := PrimHead(V4578)

tmp14458 := Call(__e, PrimFunc(symshen_4prterm), tmp14457)


_ = tmp14458

tmp14459 := MakeNative(func(__e *ControlFlow) {
Z4579 := __e.Get(1)
_ = Z4579
tmp14460 := Call(__e, PrimFunc(symstoutput))


tmp14461 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp14460)


_ = tmp14461

__e.TailApply(PrimFunc(symshen_4prterm), Z4579)
return


}, 1)

tmp14462 := PrimTail(V4578)

tmp14463 := Call(__e, PrimFunc(symmap), tmp14459, tmp14462)


_ = tmp14463

tmp14464 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(")"), tmp14464)
return


} else {
__e.TailApply(PrimFunc(symprint), V4578)
return
}


}


}, 1)

tmp14488 := Call(__e, ns2_1set, symshen_4prterm, tmp14444)


_ = tmp14488

tmp14489 := MakeNative(func(__e *ControlFlow) {
V4580 := __e.Get(1)
_ = V4580
tmp14522 := PrimEqual(Nil, V4580)

if True == tmp14522 {
__e.Return(MakeString(""))
return
} else {
tmp14520 := PrimIsPair(V4580)

var ifres14501 Obj

if True == tmp14520 {
tmp14518 := PrimHead(V4580)

tmp14519 := PrimEqual(symcons, tmp14518)

var ifres14503 Obj

if True == tmp14519 {
tmp14516 := PrimTail(V4580)

tmp14517 := PrimIsPair(tmp14516)

var ifres14505 Obj

if True == tmp14517 {
tmp14513 := PrimTail(V4580)

tmp14514 := PrimTail(tmp14513)

tmp14515 := PrimIsPair(tmp14514)

var ifres14507 Obj

if True == tmp14515 {
tmp14509 := PrimTail(V4580)

tmp14510 := PrimTail(tmp14509)

tmp14511 := PrimTail(tmp14510)

tmp14512 := PrimEqual(Nil, tmp14511)

var ifres14508 Obj

if True == tmp14512 {
ifres14508 = True


} else {
ifres14508 = False


}

ifres14507 = ifres14508


} else {
ifres14507 = False


}

var ifres14506 Obj

if True == ifres14507 {
ifres14506 = True


} else {
ifres14506 = False


}

ifres14505 = ifres14506


} else {
ifres14505 = False


}

var ifres14504 Obj

if True == ifres14505 {
ifres14504 = True


} else {
ifres14504 = False


}

ifres14503 = ifres14504


} else {
ifres14503 = False


}

var ifres14502 Obj

if True == ifres14503 {
ifres14502 = True


} else {
ifres14502 = False


}

ifres14501 = ifres14502


} else {
ifres14501 = False


}

if True == ifres14501 {
tmp14490 := Call(__e, PrimFunc(symstoutput))


tmp14491 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp14490)


_ = tmp14491

tmp14492 := PrimTail(V4580)

tmp14493 := PrimHead(tmp14492)

tmp14494 := Call(__e, PrimFunc(symshen_4prterm), tmp14493)


_ = tmp14494

tmp14495 := PrimTail(V4580)

tmp14496 := PrimTail(tmp14495)

tmp14497 := PrimHead(tmp14496)

__e.TailApply(PrimFunc(symshen_4prtl), tmp14497)
return


} else {
tmp14498 := Call(__e, PrimFunc(symstoutput))


tmp14499 := Call(__e, PrimFunc(sympr), MakeString(" | "), tmp14498)


_ = tmp14499

__e.TailApply(PrimFunc(symshen_4prterm), V4580)
return


}


}


}, 1)

tmp14523 := Call(__e, ns2_1set, symshen_4prtl, tmp14489)


_ = tmp14523

tmp14524 := MakeNative(func(__e *ControlFlow) {
V4587 := __e.Get(1)
_ = V4587
V4588 := __e.Get(2)
_ = V4588
tmp14537 := PrimEqual(Nil, V4587)

if True == tmp14537 {
tmp14525 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("\n> "), tmp14525)
return


} else {
tmp14535 := PrimIsPair(V4587)

if True == tmp14535 {
tmp14526 := Call(__e, PrimFunc(symshen_4app), V4588, MakeString(". "), symshen_4a)


tmp14527 := Call(__e, PrimFunc(symstoutput))


tmp14528 := Call(__e, PrimFunc(sympr), tmp14526, tmp14527)


_ = tmp14528

tmp14529 := PrimHead(V4587)

tmp14530 := Call(__e, PrimFunc(symshen_4show_1p), tmp14529)


_ = tmp14530

tmp14531 := Call(__e, PrimFunc(symnl), MakeNumber(1))


_ = tmp14531

tmp14532 := PrimTail(V4587)

tmp14533 := PrimNumberAdd(V4588, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4show_1assumptions), tmp14532, tmp14533)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.show-assumptions")))
return
}


}


}, 2)

tmp14538 := Call(__e, ns2_1set, symshen_4show_1assumptions, tmp14524)


_ = tmp14538

tmp14539 := MakeNative(func(__e *ControlFlow) {
tmp14540 := MakeNative(func(__e *ControlFlow) {
W4589 := __e.Get(1)
_ = W4589
tmp14542 := PrimEqual(W4589, MakeNumber(94))

if True == tmp14542 {
__e.Return(PrimSimpleError(MakeString("input aborted\n")))
return
} else {
__e.TailApply(PrimFunc(symnl), MakeNumber(1))
return
}


}, 1)

tmp14543 := Call(__e, PrimFunc(symstinput))


tmp14544 := PrimReadByte(tmp14543)

__e.TailApply(tmp14540, tmp14544)
return


}, 0)

tmp14545 := Call(__e, ns2_1set, symshen_4pause_1for_1user, tmp14539)


_ = tmp14545

tmp14546 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d))
return
}, 0)

tmp14547 := Call(__e, ns2_1set, symshen_4type_1theory_1enabled_2, tmp14546)


_ = tmp14547

tmp14548 := MakeNative(func(__e *ControlFlow) {
tmp14550 := Call(__e, PrimFunc(syminferences))


tmp14551 := PrimValue(symshen_4_dmaxinferences_d)

tmp14552 := PrimGreatThan(tmp14550, tmp14551)

if True == tmp14552 {
__e.Return(PrimSimpleError(MakeString("maximum inferences exceeded")))
return
} else {
__e.Return(False)
return
}


}, 0)

tmp14553 := Call(__e, ns2_1set, symshen_4maxinfexceeded_2, tmp14548)


_ = tmp14553

tmp14554 := MakeNative(func(__e *ControlFlow) {
V4590 := __e.Get(1)
_ = V4590
V4591 := __e.Get(2)
_ = V4591
V4592 := __e.Get(3)
_ = V4592
V4593 := __e.Get(4)
_ = V4593
V4594 := __e.Get(5)
_ = V4594
V4595 := __e.Get(6)
_ = V4595
V4596 := __e.Get(7)
_ = V4596
tmp14555 := MakeNative(func(__e *ControlFlow) {
W4597 := __e.Get(1)
_ = W4597
tmp14556 := MakeNative(func(__e *ControlFlow) {
W4598 := __e.Get(1)
_ = W4598
tmp15470 := PrimEqual(W4598, False)

if True == tmp15470 {
tmp14557 := MakeNative(func(__e *ControlFlow) {
W4599 := __e.Get(1)
_ = W4599
tmp15460 := PrimEqual(W4599, False)

if True == tmp15460 {
tmp14558 := MakeNative(func(__e *ControlFlow) {
W4600 := __e.Get(1)
_ = W4600
tmp15454 := PrimEqual(W4600, False)

if True == tmp15454 {
tmp14559 := MakeNative(func(__e *ControlFlow) {
W4601 := __e.Get(1)
_ = W4601
tmp15435 := PrimEqual(W4601, False)

if True == tmp15435 {
tmp14560 := MakeNative(func(__e *ControlFlow) {
W4605 := __e.Get(1)
_ = W4605
tmp15402 := PrimEqual(W4605, False)

if True == tmp15402 {
tmp14561 := MakeNative(func(__e *ControlFlow) {
W4611 := __e.Get(1)
_ = W4611
tmp15375 := PrimEqual(W4611, False)

if True == tmp15375 {
tmp14562 := MakeNative(func(__e *ControlFlow) {
W4617 := __e.Get(1)
_ = W4617
tmp15340 := PrimEqual(W4617, False)

if True == tmp15340 {
tmp14563 := MakeNative(func(__e *ControlFlow) {
W4624 := __e.Get(1)
_ = W4624
tmp15309 := PrimEqual(W4624, False)

if True == tmp15309 {
tmp14564 := MakeNative(func(__e *ControlFlow) {
W4631 := __e.Get(1)
_ = W4631
tmp15224 := PrimEqual(W4631, False)

if True == tmp15224 {
tmp14565 := MakeNative(func(__e *ControlFlow) {
W4652 := __e.Get(1)
_ = W4652
tmp15118 := PrimEqual(W4652, False)

if True == tmp15118 {
tmp14566 := MakeNative(func(__e *ControlFlow) {
W4680 := __e.Get(1)
_ = W4680
tmp15033 := PrimEqual(W4680, False)

if True == tmp15033 {
tmp14567 := MakeNative(func(__e *ControlFlow) {
W4701 := __e.Get(1)
_ = W4701
tmp14990 := PrimEqual(W4701, False)

if True == tmp14990 {
tmp14568 := MakeNative(func(__e *ControlFlow) {
W4711 := __e.Get(1)
_ = W4711
tmp14866 := PrimEqual(W4711, False)

if True == tmp14866 {
tmp14569 := MakeNative(func(__e *ControlFlow) {
W4741 := __e.Get(1)
_ = W4741
tmp14802 := PrimEqual(W4741, False)

if True == tmp14802 {
tmp14570 := MakeNative(func(__e *ControlFlow) {
W4754 := __e.Get(1)
_ = W4754
tmp14714 := PrimEqual(W4754, False)

if True == tmp14714 {
tmp14571 := MakeNative(func(__e *ControlFlow) {
W4775 := __e.Get(1)
_ = W4775
tmp14676 := PrimEqual(W4775, False)

if True == tmp14676 {
tmp14572 := MakeNative(func(__e *ControlFlow) {
W4783 := __e.Get(1)
_ = W4783
tmp14637 := PrimEqual(W4783, False)

if True == tmp14637 {
tmp14573 := MakeNative(func(__e *ControlFlow) {
W4791 := __e.Get(1)
_ = W4791
tmp14599 := PrimEqual(W4791, False)

if True == tmp14599 {
tmp14574 := MakeNative(func(__e *ControlFlow) {
W4799 := __e.Get(1)
_ = W4799
tmp14588 := PrimEqual(W4799, False)

if True == tmp14588 {
tmp14575 := MakeNative(func(__e *ControlFlow) {
W4801 := __e.Get(1)
_ = W4801
tmp14577 := PrimEqual(W4801, False)

if True == tmp14577 {
__e.TailApply(PrimFunc(symshen_4unlock), V4594, W4597)
return
} else {
__e.Return(W4801)
return
}


}, 1)

tmp14586 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres14578 Obj

if True == tmp14586 {
tmp14579 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14579

tmp14580 := PrimIntern(MakeString(":"))

tmp14581 := PrimCons(V4591, Nil)

tmp14582 := PrimCons(tmp14580, tmp14581)

tmp14583 := PrimCons(V4590, tmp14582)

tmp14584 := PrimValue(symshen_4_ddatatypes_d)

tmp14585 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), tmp14583, V4592, tmp14584, V4593, V4594, W4597, V4596)


ifres14578 = tmp14585


} else {
ifres14578 = False


}

__e.TailApply(tmp14575, ifres14578)
return


} else {
__e.Return(W4799)
return
}


}, 1)

tmp14597 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres14589 Obj

if True == tmp14597 {
tmp14590 := MakeNative(func(__e *ControlFlow) {
W4800 := __e.Get(1)
_ = W4800
tmp14591 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14591

tmp14592 := MakeNative(func(__e *ControlFlow) {
tmp14593 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), V4590, V4591, W4800, V4593, V4594, W4597, V4596)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4593, V4594, W4597, tmp14593)
return


}, 0)

tmp14594 := Call(__e, PrimFunc(symshen_4l_1rules), V4592, W4800, False, V4593, V4594, W4597, tmp14592)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp14594)
return


}, 1)

tmp14595 := Call(__e, PrimFunc(symshen_4newpv), V4593)


tmp14596 := Call(__e, tmp14590, tmp14595)


ifres14589 = tmp14596


} else {
ifres14589 = False


}

__e.TailApply(tmp14574, ifres14589)
return


} else {
__e.Return(W4791)
return
}


}, 1)

tmp14635 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres14600 Obj

if True == tmp14635 {
tmp14601 := MakeNative(func(__e *ControlFlow) {
W4792 := __e.Get(1)
_ = W4792
tmp14632 := PrimIsPair(W4792)

if True == tmp14632 {
tmp14602 := MakeNative(func(__e *ControlFlow) {
W4793 := __e.Get(1)
_ = W4793
tmp14628 := PrimEqual(W4793, symset)

if True == tmp14628 {
tmp14603 := MakeNative(func(__e *ControlFlow) {
W4794 := __e.Get(1)
_ = W4794
tmp14624 := PrimIsPair(W4794)

if True == tmp14624 {
tmp14604 := MakeNative(func(__e *ControlFlow) {
W4795 := __e.Get(1)
_ = W4795
tmp14605 := MakeNative(func(__e *ControlFlow) {
W4796 := __e.Get(1)
_ = W4796
tmp14619 := PrimIsPair(W4796)

if True == tmp14619 {
tmp14606 := MakeNative(func(__e *ControlFlow) {
W4797 := __e.Get(1)
_ = W4797
tmp14607 := MakeNative(func(__e *ControlFlow) {
W4798 := __e.Get(1)
_ = W4798
tmp14614 := PrimEqual(W4798, Nil)

if True == tmp14614 {
tmp14608 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14608

tmp14609 := MakeNative(func(__e *ControlFlow) {
tmp14610 := PrimCons(W4795, Nil)

tmp14611 := PrimCons(symvalue, tmp14610)

tmp14612 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4797, V4591, V4592, V4593, V4594, W4597, V4596)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp14611, V4591, V4592, V4593, V4594, W4597, tmp14612)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4795, symsymbol, V4592, V4593, V4594, W4597, tmp14609)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14615 := PrimTail(W4796)

tmp14616 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14615, V4593)


__e.TailApply(tmp14607, tmp14616)
return


}, 1)

tmp14617 := PrimHead(W4796)

__e.TailApply(tmp14606, tmp14617)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14620 := PrimTail(W4794)

tmp14621 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14620, V4593)


__e.TailApply(tmp14605, tmp14621)
return


}, 1)

tmp14622 := PrimHead(W4794)

__e.TailApply(tmp14604, tmp14622)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14625 := PrimTail(W4792)

tmp14626 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14625, V4593)


__e.TailApply(tmp14603, tmp14626)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14629 := PrimHead(W4792)

tmp14630 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14629, V4593)


__e.TailApply(tmp14602, tmp14630)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14633 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp14634 := Call(__e, tmp14601, tmp14633)


ifres14600 = tmp14634


} else {
ifres14600 = False


}

__e.TailApply(tmp14573, ifres14600)
return


} else {
__e.Return(W4783)
return
}


}, 1)

tmp14674 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres14638 Obj

if True == tmp14674 {
tmp14639 := MakeNative(func(__e *ControlFlow) {
W4784 := __e.Get(1)
_ = W4784
tmp14671 := PrimIsPair(W4784)

if True == tmp14671 {
tmp14640 := MakeNative(func(__e *ControlFlow) {
W4785 := __e.Get(1)
_ = W4785
tmp14667 := PrimEqual(W4785, syminput_7)

if True == tmp14667 {
tmp14641 := MakeNative(func(__e *ControlFlow) {
W4786 := __e.Get(1)
_ = W4786
tmp14663 := PrimIsPair(W4786)

if True == tmp14663 {
tmp14642 := MakeNative(func(__e *ControlFlow) {
W4787 := __e.Get(1)
_ = W4787
tmp14643 := MakeNative(func(__e *ControlFlow) {
W4788 := __e.Get(1)
_ = W4788
tmp14658 := PrimIsPair(W4788)

if True == tmp14658 {
tmp14644 := MakeNative(func(__e *ControlFlow) {
W4789 := __e.Get(1)
_ = W4789
tmp14645 := MakeNative(func(__e *ControlFlow) {
W4790 := __e.Get(1)
_ = W4790
tmp14653 := PrimEqual(W4790, Nil)

if True == tmp14653 {
tmp14646 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14646

tmp14647 := Call(__e, PrimFunc(symshen_4deref), W4787, V4593)


tmp14648 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp14647)


tmp14649 := MakeNative(func(__e *ControlFlow) {
tmp14650 := PrimCons(symin, Nil)

tmp14651 := PrimCons(symstream, tmp14650)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4789, tmp14651, V4592, V4593, V4594, W4597, V4596)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), V4591, tmp14648, V4593, V4594, W4597, tmp14649)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14654 := PrimTail(W4788)

tmp14655 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14654, V4593)


__e.TailApply(tmp14645, tmp14655)
return


}, 1)

tmp14656 := PrimHead(W4788)

__e.TailApply(tmp14644, tmp14656)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14659 := PrimTail(W4786)

tmp14660 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14659, V4593)


__e.TailApply(tmp14643, tmp14660)
return


}, 1)

tmp14661 := PrimHead(W4786)

__e.TailApply(tmp14642, tmp14661)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14664 := PrimTail(W4784)

tmp14665 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14664, V4593)


__e.TailApply(tmp14641, tmp14665)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14668 := PrimHead(W4784)

tmp14669 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14668, V4593)


__e.TailApply(tmp14640, tmp14669)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14672 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp14673 := Call(__e, tmp14639, tmp14672)


ifres14638 = tmp14673


} else {
ifres14638 = False


}

__e.TailApply(tmp14572, ifres14638)
return


} else {
__e.Return(W4775)
return
}


}, 1)

tmp14712 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres14677 Obj

if True == tmp14712 {
tmp14678 := MakeNative(func(__e *ControlFlow) {
W4776 := __e.Get(1)
_ = W4776
tmp14709 := PrimIsPair(W4776)

if True == tmp14709 {
tmp14679 := MakeNative(func(__e *ControlFlow) {
W4777 := __e.Get(1)
_ = W4777
tmp14705 := PrimEqual(W4777, symtype)

if True == tmp14705 {
tmp14680 := MakeNative(func(__e *ControlFlow) {
W4778 := __e.Get(1)
_ = W4778
tmp14701 := PrimIsPair(W4778)

if True == tmp14701 {
tmp14681 := MakeNative(func(__e *ControlFlow) {
W4779 := __e.Get(1)
_ = W4779
tmp14682 := MakeNative(func(__e *ControlFlow) {
W4780 := __e.Get(1)
_ = W4780
tmp14696 := PrimIsPair(W4780)

if True == tmp14696 {
tmp14683 := MakeNative(func(__e *ControlFlow) {
W4781 := __e.Get(1)
_ = W4781
tmp14684 := MakeNative(func(__e *ControlFlow) {
W4782 := __e.Get(1)
_ = W4782
tmp14691 := PrimEqual(W4782, Nil)

if True == tmp14691 {
tmp14685 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14685

tmp14686 := MakeNative(func(__e *ControlFlow) {
tmp14687 := Call(__e, PrimFunc(symshen_4deref), W4781, V4593)


tmp14688 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp14687)


tmp14689 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4779, V4591, V4592, V4593, V4594, W4597, V4596)
return
}, 0)

__e.TailApply(PrimFunc(symis_b), tmp14688, V4591, V4593, V4594, W4597, tmp14689)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4593, V4594, W4597, tmp14686)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14692 := PrimTail(W4780)

tmp14693 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14692, V4593)


__e.TailApply(tmp14684, tmp14693)
return


}, 1)

tmp14694 := PrimHead(W4780)

__e.TailApply(tmp14683, tmp14694)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14697 := PrimTail(W4778)

tmp14698 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14697, V4593)


__e.TailApply(tmp14682, tmp14698)
return


}, 1)

tmp14699 := PrimHead(W4778)

__e.TailApply(tmp14681, tmp14699)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14702 := PrimTail(W4776)

tmp14703 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14702, V4593)


__e.TailApply(tmp14680, tmp14703)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14706 := PrimHead(W4776)

tmp14707 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14706, V4593)


__e.TailApply(tmp14679, tmp14707)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14710 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp14711 := Call(__e, tmp14678, tmp14710)


ifres14677 = tmp14711


} else {
ifres14677 = False


}

__e.TailApply(tmp14571, ifres14677)
return


} else {
__e.Return(W4754)
return
}


}, 1)

tmp14800 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres14715 Obj

if True == tmp14800 {
tmp14716 := MakeNative(func(__e *ControlFlow) {
W4755 := __e.Get(1)
_ = W4755
tmp14797 := PrimIsPair(W4755)

if True == tmp14797 {
tmp14717 := MakeNative(func(__e *ControlFlow) {
W4756 := __e.Get(1)
_ = W4756
tmp14793 := PrimEqual(W4756, symopen)

if True == tmp14793 {
tmp14718 := MakeNative(func(__e *ControlFlow) {
W4757 := __e.Get(1)
_ = W4757
tmp14789 := PrimIsPair(W4757)

if True == tmp14789 {
tmp14719 := MakeNative(func(__e *ControlFlow) {
W4758 := __e.Get(1)
_ = W4758
tmp14720 := MakeNative(func(__e *ControlFlow) {
W4759 := __e.Get(1)
_ = W4759
tmp14784 := PrimIsPair(W4759)

if True == tmp14784 {
tmp14721 := MakeNative(func(__e *ControlFlow) {
W4760 := __e.Get(1)
_ = W4760
tmp14722 := MakeNative(func(__e *ControlFlow) {
W4761 := __e.Get(1)
_ = W4761
tmp14779 := PrimEqual(W4761, Nil)

if True == tmp14779 {
tmp14723 := MakeNative(func(__e *ControlFlow) {
W4762 := __e.Get(1)
_ = W4762
tmp14724 := MakeNative(func(__e *ControlFlow) {
W4763 := __e.Get(1)
_ = W4763
tmp14768 := PrimIsPair(W4762)

if True == tmp14768 {
tmp14725 := MakeNative(func(__e *ControlFlow) {
W4765 := __e.Get(1)
_ = W4765
tmp14726 := MakeNative(func(__e *ControlFlow) {
W4766 := __e.Get(1)
_ = W4766
tmp14730 := PrimEqual(W4765, symstream)

if True == tmp14730 {
__e.TailApply(PrimFunc(symthaw), W4766)
return
} else {
tmp14728 := Call(__e, PrimFunc(symshen_4pvar_2), W4765)


if True == tmp14728 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4765, symstream, V4593, W4766)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14731 := MakeNative(func(__e *ControlFlow) {
tmp14732 := MakeNative(func(__e *ControlFlow) {
W4767 := __e.Get(1)
_ = W4767
tmp14733 := MakeNative(func(__e *ControlFlow) {
W4768 := __e.Get(1)
_ = W4768
tmp14753 := PrimIsPair(W4767)

if True == tmp14753 {
tmp14734 := MakeNative(func(__e *ControlFlow) {
W4770 := __e.Get(1)
_ = W4770
tmp14735 := MakeNative(func(__e *ControlFlow) {
W4771 := __e.Get(1)
_ = W4771
tmp14736 := MakeNative(func(__e *ControlFlow) {
W4772 := __e.Get(1)
_ = W4772
tmp14740 := PrimEqual(W4771, Nil)

if True == tmp14740 {
__e.TailApply(PrimFunc(symthaw), W4772)
return
} else {
tmp14738 := Call(__e, PrimFunc(symshen_4pvar_2), W4771)


if True == tmp14738 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4771, Nil, V4593, W4772)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14741 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4768, W4770)
return
}, 0)

__e.TailApply(tmp14736, tmp14741)
return


}, 1)

tmp14742 := PrimTail(W4767)

tmp14743 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14742, V4593)


__e.TailApply(tmp14735, tmp14743)
return


}, 1)

tmp14744 := PrimHead(W4767)

__e.TailApply(tmp14734, tmp14744)
return


} else {
tmp14751 := Call(__e, PrimFunc(symshen_4pvar_2), W4767)


if True == tmp14751 {
tmp14745 := MakeNative(func(__e *ControlFlow) {
W4773 := __e.Get(1)
_ = W4773
tmp14746 := PrimCons(W4773, Nil)

tmp14747 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4768, W4773)
return
}, 0)

tmp14748 := Call(__e, PrimFunc(symshen_4bind_b), W4767, tmp14746, V4593, tmp14747)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp14748)
return


}, 1)

tmp14749 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp14745, tmp14749)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14754 := MakeNative(func(__e *ControlFlow) {
Z4769 := __e.Get(1)
_ = Z4769
__e.TailApply(W4763, Z4769)
return
}, 1)

__e.TailApply(tmp14733, tmp14754)
return


}, 1)

tmp14755 := PrimTail(W4762)

tmp14756 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14755, V4593)


__e.TailApply(tmp14732, tmp14756)
return


}, 0)

__e.TailApply(tmp14726, tmp14731)
return


}, 1)

tmp14757 := PrimHead(W4762)

tmp14758 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14757, V4593)


__e.TailApply(tmp14725, tmp14758)
return


} else {
tmp14766 := Call(__e, PrimFunc(symshen_4pvar_2), W4762)


if True == tmp14766 {
tmp14759 := MakeNative(func(__e *ControlFlow) {
W4774 := __e.Get(1)
_ = W4774
tmp14760 := PrimCons(W4774, Nil)

tmp14761 := PrimCons(symstream, tmp14760)

tmp14762 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4763, W4774)
return
}, 0)

tmp14763 := Call(__e, PrimFunc(symshen_4bind_b), W4762, tmp14761, V4593, tmp14762)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp14763)
return


}, 1)

tmp14764 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp14759, tmp14764)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14769 := MakeNative(func(__e *ControlFlow) {
Z4764 := __e.Get(1)
_ = Z4764
tmp14770 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14770

tmp14771 := MakeNative(func(__e *ControlFlow) {
tmp14772 := Call(__e, PrimFunc(symshen_4lazyderef), Z4764, V4593)


tmp14773 := PrimCons(symout, Nil)

tmp14774 := PrimCons(symin, tmp14773)

tmp14775 := Call(__e, PrimFunc(symelement_2), tmp14772, tmp14774)


tmp14776 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4758, symstring, V4592, V4593, V4594, W4597, V4596)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14775, V4593, V4594, W4597, tmp14776)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W4760, Z4764, V4593, V4594, W4597, tmp14771)
return


}, 1)

__e.TailApply(tmp14724, tmp14769)
return


}, 1)

tmp14777 := Call(__e, PrimFunc(symshen_4lazyderef), V4591, V4593)


__e.TailApply(tmp14723, tmp14777)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14780 := PrimTail(W4759)

tmp14781 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14780, V4593)


__e.TailApply(tmp14722, tmp14781)
return


}, 1)

tmp14782 := PrimHead(W4759)

__e.TailApply(tmp14721, tmp14782)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14785 := PrimTail(W4757)

tmp14786 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14785, V4593)


__e.TailApply(tmp14720, tmp14786)
return


}, 1)

tmp14787 := PrimHead(W4757)

__e.TailApply(tmp14719, tmp14787)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14790 := PrimTail(W4755)

tmp14791 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14790, V4593)


__e.TailApply(tmp14718, tmp14791)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14794 := PrimHead(W4755)

tmp14795 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14794, V4593)


__e.TailApply(tmp14717, tmp14795)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14798 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp14799 := Call(__e, tmp14716, tmp14798)


ifres14715 = tmp14799


} else {
ifres14715 = False


}

__e.TailApply(tmp14570, ifres14715)
return


} else {
__e.Return(W4741)
return
}


}, 1)

tmp14864 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres14803 Obj

if True == tmp14864 {
tmp14804 := MakeNative(func(__e *ControlFlow) {
W4742 := __e.Get(1)
_ = W4742
tmp14861 := PrimIsPair(W4742)

if True == tmp14861 {
tmp14805 := MakeNative(func(__e *ControlFlow) {
W4743 := __e.Get(1)
_ = W4743
tmp14857 := PrimEqual(W4743, symlet)

if True == tmp14857 {
tmp14806 := MakeNative(func(__e *ControlFlow) {
W4744 := __e.Get(1)
_ = W4744
tmp14853 := PrimIsPair(W4744)

if True == tmp14853 {
tmp14807 := MakeNative(func(__e *ControlFlow) {
W4745 := __e.Get(1)
_ = W4745
tmp14808 := MakeNative(func(__e *ControlFlow) {
W4746 := __e.Get(1)
_ = W4746
tmp14848 := PrimIsPair(W4746)

if True == tmp14848 {
tmp14809 := MakeNative(func(__e *ControlFlow) {
W4747 := __e.Get(1)
_ = W4747
tmp14810 := MakeNative(func(__e *ControlFlow) {
W4748 := __e.Get(1)
_ = W4748
tmp14843 := PrimIsPair(W4748)

if True == tmp14843 {
tmp14811 := MakeNative(func(__e *ControlFlow) {
W4749 := __e.Get(1)
_ = W4749
tmp14812 := MakeNative(func(__e *ControlFlow) {
W4750 := __e.Get(1)
_ = W4750
tmp14838 := PrimEqual(W4750, Nil)

if True == tmp14838 {
tmp14813 := MakeNative(func(__e *ControlFlow) {
W4751 := __e.Get(1)
_ = W4751
tmp14814 := MakeNative(func(__e *ControlFlow) {
W4752 := __e.Get(1)
_ = W4752
tmp14815 := MakeNative(func(__e *ControlFlow) {
W4753 := __e.Get(1)
_ = W4753
tmp14816 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14816

tmp14817 := MakeNative(func(__e *ControlFlow) {
tmp14818 := Call(__e, PrimFunc(symshen_4lazyderef), W4745, V4593)


tmp14819 := Call(__e, PrimFunc(symshen_4freshterm), tmp14818)


tmp14820 := MakeNative(func(__e *ControlFlow) {
tmp14821 := Call(__e, PrimFunc(symshen_4lazyderef), W4745, V4593)


tmp14822 := Call(__e, PrimFunc(symshen_4lazyderef), W4752, V4593)


tmp14823 := Call(__e, PrimFunc(symshen_4lazyderef), W4749, V4593)


tmp14824 := Call(__e, PrimFunc(symshen_4beta), tmp14821, tmp14822, tmp14823)


tmp14825 := MakeNative(func(__e *ControlFlow) {
tmp14826 := PrimIntern(MakeString(":"))

tmp14827 := PrimCons(W4753, Nil)

tmp14828 := PrimCons(tmp14826, tmp14827)

tmp14829 := PrimCons(W4752, tmp14828)

tmp14830 := PrimCons(tmp14829, V4592)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4751, V4591, tmp14830, V4593, V4594, W4597, V4596)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4751, tmp14824, V4593, V4594, W4597, tmp14825)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4752, tmp14819, V4593, V4594, W4597, tmp14820)
return


}, 0)

tmp14831 := Call(__e, PrimFunc(symshen_4system_1S_1h), W4747, W4753, V4592, V4593, V4594, W4597, tmp14817)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp14831)
return


}, 1)

tmp14832 := Call(__e, PrimFunc(symshen_4newpv), V4593)


tmp14833 := Call(__e, tmp14815, tmp14832)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp14833)
return


}, 1)

tmp14834 := Call(__e, PrimFunc(symshen_4newpv), V4593)


tmp14835 := Call(__e, tmp14814, tmp14834)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp14835)
return


}, 1)

tmp14836 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp14813, tmp14836)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14839 := PrimTail(W4748)

tmp14840 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14839, V4593)


__e.TailApply(tmp14812, tmp14840)
return


}, 1)

tmp14841 := PrimHead(W4748)

__e.TailApply(tmp14811, tmp14841)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14844 := PrimTail(W4746)

tmp14845 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14844, V4593)


__e.TailApply(tmp14810, tmp14845)
return


}, 1)

tmp14846 := PrimHead(W4746)

__e.TailApply(tmp14809, tmp14846)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14849 := PrimTail(W4744)

tmp14850 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14849, V4593)


__e.TailApply(tmp14808, tmp14850)
return


}, 1)

tmp14851 := PrimHead(W4744)

__e.TailApply(tmp14807, tmp14851)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14854 := PrimTail(W4742)

tmp14855 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14854, V4593)


__e.TailApply(tmp14806, tmp14855)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14858 := PrimHead(W4742)

tmp14859 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14858, V4593)


__e.TailApply(tmp14805, tmp14859)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14862 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp14863 := Call(__e, tmp14804, tmp14862)


ifres14803 = tmp14863


} else {
ifres14803 = False


}

__e.TailApply(tmp14569, ifres14803)
return


} else {
__e.Return(W4711)
return
}


}, 1)

tmp14988 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres14867 Obj

if True == tmp14988 {
tmp14868 := MakeNative(func(__e *ControlFlow) {
W4712 := __e.Get(1)
_ = W4712
tmp14985 := PrimIsPair(W4712)

if True == tmp14985 {
tmp14869 := MakeNative(func(__e *ControlFlow) {
W4713 := __e.Get(1)
_ = W4713
tmp14981 := PrimEqual(W4713, symlambda)

if True == tmp14981 {
tmp14870 := MakeNative(func(__e *ControlFlow) {
W4714 := __e.Get(1)
_ = W4714
tmp14977 := PrimIsPair(W4714)

if True == tmp14977 {
tmp14871 := MakeNative(func(__e *ControlFlow) {
W4715 := __e.Get(1)
_ = W4715
tmp14872 := MakeNative(func(__e *ControlFlow) {
W4716 := __e.Get(1)
_ = W4716
tmp14972 := PrimIsPair(W4716)

if True == tmp14972 {
tmp14873 := MakeNative(func(__e *ControlFlow) {
W4717 := __e.Get(1)
_ = W4717
tmp14874 := MakeNative(func(__e *ControlFlow) {
W4718 := __e.Get(1)
_ = W4718
tmp14967 := PrimEqual(W4718, Nil)

if True == tmp14967 {
tmp14875 := MakeNative(func(__e *ControlFlow) {
W4719 := __e.Get(1)
_ = W4719
tmp14876 := MakeNative(func(__e *ControlFlow) {
W4720 := __e.Get(1)
_ = W4720
tmp14943 := PrimIsPair(W4719)

if True == tmp14943 {
tmp14877 := MakeNative(func(__e *ControlFlow) {
W4725 := __e.Get(1)
_ = W4725
tmp14878 := MakeNative(func(__e *ControlFlow) {
W4726 := __e.Get(1)
_ = W4726
tmp14879 := MakeNative(func(__e *ControlFlow) {
W4727 := __e.Get(1)
_ = W4727
tmp14923 := PrimIsPair(W4726)

if True == tmp14923 {
tmp14880 := MakeNative(func(__e *ControlFlow) {
W4729 := __e.Get(1)
_ = W4729
tmp14881 := MakeNative(func(__e *ControlFlow) {
W4730 := __e.Get(1)
_ = W4730
tmp14885 := PrimEqual(W4729, sym_1_1_6)

if True == tmp14885 {
__e.TailApply(PrimFunc(symthaw), W4730)
return
} else {
tmp14883 := Call(__e, PrimFunc(symshen_4pvar_2), W4729)


if True == tmp14883 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4729, sym_1_1_6, V4593, W4730)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14886 := MakeNative(func(__e *ControlFlow) {
tmp14887 := MakeNative(func(__e *ControlFlow) {
W4731 := __e.Get(1)
_ = W4731
tmp14888 := MakeNative(func(__e *ControlFlow) {
W4732 := __e.Get(1)
_ = W4732
tmp14908 := PrimIsPair(W4731)

if True == tmp14908 {
tmp14889 := MakeNative(func(__e *ControlFlow) {
W4734 := __e.Get(1)
_ = W4734
tmp14890 := MakeNative(func(__e *ControlFlow) {
W4735 := __e.Get(1)
_ = W4735
tmp14891 := MakeNative(func(__e *ControlFlow) {
W4736 := __e.Get(1)
_ = W4736
tmp14895 := PrimEqual(W4735, Nil)

if True == tmp14895 {
__e.TailApply(PrimFunc(symthaw), W4736)
return
} else {
tmp14893 := Call(__e, PrimFunc(symshen_4pvar_2), W4735)


if True == tmp14893 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4735, Nil, V4593, W4736)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14896 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4732, W4734)
return
}, 0)

__e.TailApply(tmp14891, tmp14896)
return


}, 1)

tmp14897 := PrimTail(W4731)

tmp14898 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14897, V4593)


__e.TailApply(tmp14890, tmp14898)
return


}, 1)

tmp14899 := PrimHead(W4731)

__e.TailApply(tmp14889, tmp14899)
return


} else {
tmp14906 := Call(__e, PrimFunc(symshen_4pvar_2), W4731)


if True == tmp14906 {
tmp14900 := MakeNative(func(__e *ControlFlow) {
W4737 := __e.Get(1)
_ = W4737
tmp14901 := PrimCons(W4737, Nil)

tmp14902 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4732, W4737)
return
}, 0)

tmp14903 := Call(__e, PrimFunc(symshen_4bind_b), W4731, tmp14901, V4593, tmp14902)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp14903)
return


}, 1)

tmp14904 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp14900, tmp14904)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14909 := MakeNative(func(__e *ControlFlow) {
Z4733 := __e.Get(1)
_ = Z4733
__e.TailApply(W4727, Z4733)
return
}, 1)

__e.TailApply(tmp14888, tmp14909)
return


}, 1)

tmp14910 := PrimTail(W4726)

tmp14911 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14910, V4593)


__e.TailApply(tmp14887, tmp14911)
return


}, 0)

__e.TailApply(tmp14881, tmp14886)
return


}, 1)

tmp14912 := PrimHead(W4726)

tmp14913 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14912, V4593)


__e.TailApply(tmp14880, tmp14913)
return


} else {
tmp14921 := Call(__e, PrimFunc(symshen_4pvar_2), W4726)


if True == tmp14921 {
tmp14914 := MakeNative(func(__e *ControlFlow) {
W4738 := __e.Get(1)
_ = W4738
tmp14915 := PrimCons(W4738, Nil)

tmp14916 := PrimCons(sym_1_1_6, tmp14915)

tmp14917 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4727, W4738)
return
}, 0)

tmp14918 := Call(__e, PrimFunc(symshen_4bind_b), W4726, tmp14916, V4593, tmp14917)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp14918)
return


}, 1)

tmp14919 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp14914, tmp14919)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14924 := MakeNative(func(__e *ControlFlow) {
Z4728 := __e.Get(1)
_ = Z4728
tmp14925 := Call(__e, W4720, W4725)


__e.TailApply(tmp14925, Z4728)
return


}, 1)

__e.TailApply(tmp14879, tmp14924)
return


}, 1)

tmp14926 := PrimTail(W4719)

tmp14927 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14926, V4593)


__e.TailApply(tmp14878, tmp14927)
return


}, 1)

tmp14928 := PrimHead(W4719)

__e.TailApply(tmp14877, tmp14928)
return


} else {
tmp14941 := Call(__e, PrimFunc(symshen_4pvar_2), W4719)


if True == tmp14941 {
tmp14929 := MakeNative(func(__e *ControlFlow) {
W4739 := __e.Get(1)
_ = W4739
tmp14930 := MakeNative(func(__e *ControlFlow) {
W4740 := __e.Get(1)
_ = W4740
tmp14931 := PrimCons(W4740, Nil)

tmp14932 := PrimCons(sym_1_1_6, tmp14931)

tmp14933 := PrimCons(W4739, tmp14932)

tmp14934 := MakeNative(func(__e *ControlFlow) {
tmp14935 := Call(__e, W4720, W4739)


__e.TailApply(tmp14935, W4740)
return


}, 0)

tmp14936 := Call(__e, PrimFunc(symshen_4bind_b), W4719, tmp14933, V4593, tmp14934)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp14936)
return


}, 1)

tmp14937 := Call(__e, PrimFunc(symshen_4newpv), V4593)


tmp14938 := Call(__e, tmp14930, tmp14937)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp14938)
return


}, 1)

tmp14939 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp14929, tmp14939)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14944 := MakeNative(func(__e *ControlFlow) {
Z4721 := __e.Get(1)
_ = Z4721
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4722 := __e.Get(1)
_ = Z4722
tmp14945 := MakeNative(func(__e *ControlFlow) {
W4723 := __e.Get(1)
_ = W4723
tmp14946 := MakeNative(func(__e *ControlFlow) {
W4724 := __e.Get(1)
_ = W4724
tmp14947 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14947

tmp14948 := Call(__e, PrimFunc(symshen_4lazyderef), W4715, V4593)


tmp14949 := Call(__e, PrimFunc(symshen_4freshterm), tmp14948)


tmp14950 := MakeNative(func(__e *ControlFlow) {
tmp14951 := Call(__e, PrimFunc(symshen_4lazyderef), W4715, V4593)


tmp14952 := Call(__e, PrimFunc(symshen_4deref), W4724, V4593)


tmp14953 := Call(__e, PrimFunc(symshen_4deref), W4717, V4593)


tmp14954 := Call(__e, PrimFunc(symshen_4beta), tmp14951, tmp14952, tmp14953)


tmp14955 := MakeNative(func(__e *ControlFlow) {
tmp14956 := PrimIntern(MakeString(":"))

tmp14957 := PrimCons(Z4721, Nil)

tmp14958 := PrimCons(tmp14956, tmp14957)

tmp14959 := PrimCons(W4724, tmp14958)

tmp14960 := PrimCons(tmp14959, V4592)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4723, Z4722, tmp14960, V4593, V4594, W4597, V4596)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4723, tmp14954, V4593, V4594, W4597, tmp14955)
return


}, 0)

tmp14961 := Call(__e, PrimFunc(symbind), W4724, tmp14949, V4593, V4594, W4597, tmp14950)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp14961)
return


}, 1)

tmp14962 := Call(__e, PrimFunc(symshen_4newpv), V4593)


tmp14963 := Call(__e, tmp14946, tmp14962)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp14963)
return


}, 1)

tmp14964 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp14945, tmp14964)
return


}, 1))
return
}, 1)

__e.TailApply(tmp14876, tmp14944)
return


}, 1)

tmp14965 := Call(__e, PrimFunc(symshen_4lazyderef), V4591, V4593)


__e.TailApply(tmp14875, tmp14965)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14968 := PrimTail(W4716)

tmp14969 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14968, V4593)


__e.TailApply(tmp14874, tmp14969)
return


}, 1)

tmp14970 := PrimHead(W4716)

__e.TailApply(tmp14873, tmp14970)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14973 := PrimTail(W4714)

tmp14974 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14973, V4593)


__e.TailApply(tmp14872, tmp14974)
return


}, 1)

tmp14975 := PrimHead(W4714)

__e.TailApply(tmp14871, tmp14975)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14978 := PrimTail(W4712)

tmp14979 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14978, V4593)


__e.TailApply(tmp14870, tmp14979)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14982 := PrimHead(W4712)

tmp14983 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14982, V4593)


__e.TailApply(tmp14869, tmp14983)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14986 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp14987 := Call(__e, tmp14868, tmp14986)


ifres14867 = tmp14987


} else {
ifres14867 = False


}

__e.TailApply(tmp14568, ifres14867)
return


} else {
__e.Return(W4701)
return
}


}, 1)

tmp15031 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres14991 Obj

if True == tmp15031 {
tmp14992 := MakeNative(func(__e *ControlFlow) {
W4702 := __e.Get(1)
_ = W4702
tmp15028 := PrimIsPair(W4702)

if True == tmp15028 {
tmp14993 := MakeNative(func(__e *ControlFlow) {
W4703 := __e.Get(1)
_ = W4703
tmp15024 := PrimEqual(W4703, sym_8s)

if True == tmp15024 {
tmp14994 := MakeNative(func(__e *ControlFlow) {
W4704 := __e.Get(1)
_ = W4704
tmp15020 := PrimIsPair(W4704)

if True == tmp15020 {
tmp14995 := MakeNative(func(__e *ControlFlow) {
W4705 := __e.Get(1)
_ = W4705
tmp14996 := MakeNative(func(__e *ControlFlow) {
W4706 := __e.Get(1)
_ = W4706
tmp15015 := PrimIsPair(W4706)

if True == tmp15015 {
tmp14997 := MakeNative(func(__e *ControlFlow) {
W4707 := __e.Get(1)
_ = W4707
tmp14998 := MakeNative(func(__e *ControlFlow) {
W4708 := __e.Get(1)
_ = W4708
tmp15010 := PrimEqual(W4708, Nil)

if True == tmp15010 {
tmp14999 := MakeNative(func(__e *ControlFlow) {
W4709 := __e.Get(1)
_ = W4709
tmp15000 := MakeNative(func(__e *ControlFlow) {
W4710 := __e.Get(1)
_ = W4710
tmp15004 := PrimEqual(W4709, symstring)

if True == tmp15004 {
__e.TailApply(PrimFunc(symthaw), W4710)
return
} else {
tmp15002 := Call(__e, PrimFunc(symshen_4pvar_2), W4709)


if True == tmp15002 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4709, symstring, V4593, W4710)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15005 := MakeNative(func(__e *ControlFlow) {
tmp15006 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15006

tmp15007 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4707, symstring, V4592, V4593, V4594, W4597, V4596)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4705, symstring, V4592, V4593, V4594, W4597, tmp15007)
return


}, 0)

__e.TailApply(tmp15000, tmp15005)
return


}, 1)

tmp15008 := Call(__e, PrimFunc(symshen_4lazyderef), V4591, V4593)


__e.TailApply(tmp14999, tmp15008)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15011 := PrimTail(W4706)

tmp15012 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15011, V4593)


__e.TailApply(tmp14998, tmp15012)
return


}, 1)

tmp15013 := PrimHead(W4706)

__e.TailApply(tmp14997, tmp15013)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15016 := PrimTail(W4704)

tmp15017 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15016, V4593)


__e.TailApply(tmp14996, tmp15017)
return


}, 1)

tmp15018 := PrimHead(W4704)

__e.TailApply(tmp14995, tmp15018)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15021 := PrimTail(W4702)

tmp15022 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15021, V4593)


__e.TailApply(tmp14994, tmp15022)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15025 := PrimHead(W4702)

tmp15026 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15025, V4593)


__e.TailApply(tmp14993, tmp15026)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15029 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp15030 := Call(__e, tmp14992, tmp15029)


ifres14991 = tmp15030


} else {
ifres14991 = False


}

__e.TailApply(tmp14567, ifres14991)
return


} else {
__e.Return(W4680)
return
}


}, 1)

tmp15116 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres15034 Obj

if True == tmp15116 {
tmp15035 := MakeNative(func(__e *ControlFlow) {
W4681 := __e.Get(1)
_ = W4681
tmp15113 := PrimIsPair(W4681)

if True == tmp15113 {
tmp15036 := MakeNative(func(__e *ControlFlow) {
W4682 := __e.Get(1)
_ = W4682
tmp15109 := PrimEqual(W4682, sym_8v)

if True == tmp15109 {
tmp15037 := MakeNative(func(__e *ControlFlow) {
W4683 := __e.Get(1)
_ = W4683
tmp15105 := PrimIsPair(W4683)

if True == tmp15105 {
tmp15038 := MakeNative(func(__e *ControlFlow) {
W4684 := __e.Get(1)
_ = W4684
tmp15039 := MakeNative(func(__e *ControlFlow) {
W4685 := __e.Get(1)
_ = W4685
tmp15100 := PrimIsPair(W4685)

if True == tmp15100 {
tmp15040 := MakeNative(func(__e *ControlFlow) {
W4686 := __e.Get(1)
_ = W4686
tmp15041 := MakeNative(func(__e *ControlFlow) {
W4687 := __e.Get(1)
_ = W4687
tmp15095 := PrimEqual(W4687, Nil)

if True == tmp15095 {
tmp15042 := MakeNative(func(__e *ControlFlow) {
W4688 := __e.Get(1)
_ = W4688
tmp15043 := MakeNative(func(__e *ControlFlow) {
W4689 := __e.Get(1)
_ = W4689
tmp15087 := PrimIsPair(W4688)

if True == tmp15087 {
tmp15044 := MakeNative(func(__e *ControlFlow) {
W4691 := __e.Get(1)
_ = W4691
tmp15045 := MakeNative(func(__e *ControlFlow) {
W4692 := __e.Get(1)
_ = W4692
tmp15049 := PrimEqual(W4691, symvector)

if True == tmp15049 {
__e.TailApply(PrimFunc(symthaw), W4692)
return
} else {
tmp15047 := Call(__e, PrimFunc(symshen_4pvar_2), W4691)


if True == tmp15047 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4691, symvector, V4593, W4692)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15050 := MakeNative(func(__e *ControlFlow) {
tmp15051 := MakeNative(func(__e *ControlFlow) {
W4693 := __e.Get(1)
_ = W4693
tmp15052 := MakeNative(func(__e *ControlFlow) {
W4694 := __e.Get(1)
_ = W4694
tmp15072 := PrimIsPair(W4693)

if True == tmp15072 {
tmp15053 := MakeNative(func(__e *ControlFlow) {
W4696 := __e.Get(1)
_ = W4696
tmp15054 := MakeNative(func(__e *ControlFlow) {
W4697 := __e.Get(1)
_ = W4697
tmp15055 := MakeNative(func(__e *ControlFlow) {
W4698 := __e.Get(1)
_ = W4698
tmp15059 := PrimEqual(W4697, Nil)

if True == tmp15059 {
__e.TailApply(PrimFunc(symthaw), W4698)
return
} else {
tmp15057 := Call(__e, PrimFunc(symshen_4pvar_2), W4697)


if True == tmp15057 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4697, Nil, V4593, W4698)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15060 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4694, W4696)
return
}, 0)

__e.TailApply(tmp15055, tmp15060)
return


}, 1)

tmp15061 := PrimTail(W4693)

tmp15062 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15061, V4593)


__e.TailApply(tmp15054, tmp15062)
return


}, 1)

tmp15063 := PrimHead(W4693)

__e.TailApply(tmp15053, tmp15063)
return


} else {
tmp15070 := Call(__e, PrimFunc(symshen_4pvar_2), W4693)


if True == tmp15070 {
tmp15064 := MakeNative(func(__e *ControlFlow) {
W4699 := __e.Get(1)
_ = W4699
tmp15065 := PrimCons(W4699, Nil)

tmp15066 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4694, W4699)
return
}, 0)

tmp15067 := Call(__e, PrimFunc(symshen_4bind_b), W4693, tmp15065, V4593, tmp15066)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp15067)
return


}, 1)

tmp15068 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp15064, tmp15068)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15073 := MakeNative(func(__e *ControlFlow) {
Z4695 := __e.Get(1)
_ = Z4695
__e.TailApply(W4689, Z4695)
return
}, 1)

__e.TailApply(tmp15052, tmp15073)
return


}, 1)

tmp15074 := PrimTail(W4688)

tmp15075 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15074, V4593)


__e.TailApply(tmp15051, tmp15075)
return


}, 0)

__e.TailApply(tmp15045, tmp15050)
return


}, 1)

tmp15076 := PrimHead(W4688)

tmp15077 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15076, V4593)


__e.TailApply(tmp15044, tmp15077)
return


} else {
tmp15085 := Call(__e, PrimFunc(symshen_4pvar_2), W4688)


if True == tmp15085 {
tmp15078 := MakeNative(func(__e *ControlFlow) {
W4700 := __e.Get(1)
_ = W4700
tmp15079 := PrimCons(W4700, Nil)

tmp15080 := PrimCons(symvector, tmp15079)

tmp15081 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4689, W4700)
return
}, 0)

tmp15082 := Call(__e, PrimFunc(symshen_4bind_b), W4688, tmp15080, V4593, tmp15081)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp15082)
return


}, 1)

tmp15083 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp15078, tmp15083)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15088 := MakeNative(func(__e *ControlFlow) {
Z4690 := __e.Get(1)
_ = Z4690
tmp15089 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15089

tmp15090 := MakeNative(func(__e *ControlFlow) {
tmp15091 := PrimCons(Z4690, Nil)

tmp15092 := PrimCons(symvector, tmp15091)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4686, tmp15092, V4592, V4593, V4594, W4597, V4596)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4684, Z4690, V4592, V4593, V4594, W4597, tmp15090)
return


}, 1)

__e.TailApply(tmp15043, tmp15088)
return


}, 1)

tmp15093 := Call(__e, PrimFunc(symshen_4lazyderef), V4591, V4593)


__e.TailApply(tmp15042, tmp15093)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15096 := PrimTail(W4685)

tmp15097 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15096, V4593)


__e.TailApply(tmp15041, tmp15097)
return


}, 1)

tmp15098 := PrimHead(W4685)

__e.TailApply(tmp15040, tmp15098)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15101 := PrimTail(W4683)

tmp15102 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15101, V4593)


__e.TailApply(tmp15039, tmp15102)
return


}, 1)

tmp15103 := PrimHead(W4683)

__e.TailApply(tmp15038, tmp15103)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15106 := PrimTail(W4681)

tmp15107 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15106, V4593)


__e.TailApply(tmp15037, tmp15107)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15110 := PrimHead(W4681)

tmp15111 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15110, V4593)


__e.TailApply(tmp15036, tmp15111)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15114 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp15115 := Call(__e, tmp15035, tmp15114)


ifres15034 = tmp15115


} else {
ifres15034 = False


}

__e.TailApply(tmp14566, ifres15034)
return


} else {
__e.Return(W4652)
return
}


}, 1)

tmp15222 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres15119 Obj

if True == tmp15222 {
tmp15120 := MakeNative(func(__e *ControlFlow) {
W4653 := __e.Get(1)
_ = W4653
tmp15219 := PrimIsPair(W4653)

if True == tmp15219 {
tmp15121 := MakeNative(func(__e *ControlFlow) {
W4654 := __e.Get(1)
_ = W4654
tmp15215 := PrimEqual(W4654, sym_8p)

if True == tmp15215 {
tmp15122 := MakeNative(func(__e *ControlFlow) {
W4655 := __e.Get(1)
_ = W4655
tmp15211 := PrimIsPair(W4655)

if True == tmp15211 {
tmp15123 := MakeNative(func(__e *ControlFlow) {
W4656 := __e.Get(1)
_ = W4656
tmp15124 := MakeNative(func(__e *ControlFlow) {
W4657 := __e.Get(1)
_ = W4657
tmp15206 := PrimIsPair(W4657)

if True == tmp15206 {
tmp15125 := MakeNative(func(__e *ControlFlow) {
W4658 := __e.Get(1)
_ = W4658
tmp15126 := MakeNative(func(__e *ControlFlow) {
W4659 := __e.Get(1)
_ = W4659
tmp15201 := PrimEqual(W4659, Nil)

if True == tmp15201 {
tmp15127 := MakeNative(func(__e *ControlFlow) {
W4660 := __e.Get(1)
_ = W4660
tmp15128 := MakeNative(func(__e *ControlFlow) {
W4661 := __e.Get(1)
_ = W4661
tmp15195 := PrimIsPair(W4660)

if True == tmp15195 {
tmp15129 := MakeNative(func(__e *ControlFlow) {
W4664 := __e.Get(1)
_ = W4664
tmp15130 := MakeNative(func(__e *ControlFlow) {
W4665 := __e.Get(1)
_ = W4665
tmp15131 := MakeNative(func(__e *ControlFlow) {
W4666 := __e.Get(1)
_ = W4666
tmp15175 := PrimIsPair(W4665)

if True == tmp15175 {
tmp15132 := MakeNative(func(__e *ControlFlow) {
W4668 := __e.Get(1)
_ = W4668
tmp15133 := MakeNative(func(__e *ControlFlow) {
W4669 := __e.Get(1)
_ = W4669
tmp15137 := PrimEqual(W4668, sym_d)

if True == tmp15137 {
__e.TailApply(PrimFunc(symthaw), W4669)
return
} else {
tmp15135 := Call(__e, PrimFunc(symshen_4pvar_2), W4668)


if True == tmp15135 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4668, sym_d, V4593, W4669)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15138 := MakeNative(func(__e *ControlFlow) {
tmp15139 := MakeNative(func(__e *ControlFlow) {
W4670 := __e.Get(1)
_ = W4670
tmp15140 := MakeNative(func(__e *ControlFlow) {
W4671 := __e.Get(1)
_ = W4671
tmp15160 := PrimIsPair(W4670)

if True == tmp15160 {
tmp15141 := MakeNative(func(__e *ControlFlow) {
W4673 := __e.Get(1)
_ = W4673
tmp15142 := MakeNative(func(__e *ControlFlow) {
W4674 := __e.Get(1)
_ = W4674
tmp15143 := MakeNative(func(__e *ControlFlow) {
W4675 := __e.Get(1)
_ = W4675
tmp15147 := PrimEqual(W4674, Nil)

if True == tmp15147 {
__e.TailApply(PrimFunc(symthaw), W4675)
return
} else {
tmp15145 := Call(__e, PrimFunc(symshen_4pvar_2), W4674)


if True == tmp15145 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4674, Nil, V4593, W4675)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15148 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4671, W4673)
return
}, 0)

__e.TailApply(tmp15143, tmp15148)
return


}, 1)

tmp15149 := PrimTail(W4670)

tmp15150 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15149, V4593)


__e.TailApply(tmp15142, tmp15150)
return


}, 1)

tmp15151 := PrimHead(W4670)

__e.TailApply(tmp15141, tmp15151)
return


} else {
tmp15158 := Call(__e, PrimFunc(symshen_4pvar_2), W4670)


if True == tmp15158 {
tmp15152 := MakeNative(func(__e *ControlFlow) {
W4676 := __e.Get(1)
_ = W4676
tmp15153 := PrimCons(W4676, Nil)

tmp15154 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4671, W4676)
return
}, 0)

tmp15155 := Call(__e, PrimFunc(symshen_4bind_b), W4670, tmp15153, V4593, tmp15154)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp15155)
return


}, 1)

tmp15156 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp15152, tmp15156)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15161 := MakeNative(func(__e *ControlFlow) {
Z4672 := __e.Get(1)
_ = Z4672
__e.TailApply(W4666, Z4672)
return
}, 1)

__e.TailApply(tmp15140, tmp15161)
return


}, 1)

tmp15162 := PrimTail(W4665)

tmp15163 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15162, V4593)


__e.TailApply(tmp15139, tmp15163)
return


}, 0)

__e.TailApply(tmp15133, tmp15138)
return


}, 1)

tmp15164 := PrimHead(W4665)

tmp15165 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15164, V4593)


__e.TailApply(tmp15132, tmp15165)
return


} else {
tmp15173 := Call(__e, PrimFunc(symshen_4pvar_2), W4665)


if True == tmp15173 {
tmp15166 := MakeNative(func(__e *ControlFlow) {
W4677 := __e.Get(1)
_ = W4677
tmp15167 := PrimCons(W4677, Nil)

tmp15168 := PrimCons(sym_d, tmp15167)

tmp15169 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4666, W4677)
return
}, 0)

tmp15170 := Call(__e, PrimFunc(symshen_4bind_b), W4665, tmp15168, V4593, tmp15169)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp15170)
return


}, 1)

tmp15171 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp15166, tmp15171)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15176 := MakeNative(func(__e *ControlFlow) {
Z4667 := __e.Get(1)
_ = Z4667
tmp15177 := Call(__e, W4661, W4664)


__e.TailApply(tmp15177, Z4667)
return


}, 1)

__e.TailApply(tmp15131, tmp15176)
return


}, 1)

tmp15178 := PrimTail(W4660)

tmp15179 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15178, V4593)


__e.TailApply(tmp15130, tmp15179)
return


}, 1)

tmp15180 := PrimHead(W4660)

__e.TailApply(tmp15129, tmp15180)
return


} else {
tmp15193 := Call(__e, PrimFunc(symshen_4pvar_2), W4660)


if True == tmp15193 {
tmp15181 := MakeNative(func(__e *ControlFlow) {
W4678 := __e.Get(1)
_ = W4678
tmp15182 := MakeNative(func(__e *ControlFlow) {
W4679 := __e.Get(1)
_ = W4679
tmp15183 := PrimCons(W4679, Nil)

tmp15184 := PrimCons(sym_d, tmp15183)

tmp15185 := PrimCons(W4678, tmp15184)

tmp15186 := MakeNative(func(__e *ControlFlow) {
tmp15187 := Call(__e, W4661, W4678)


__e.TailApply(tmp15187, W4679)
return


}, 0)

tmp15188 := Call(__e, PrimFunc(symshen_4bind_b), W4660, tmp15185, V4593, tmp15186)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp15188)
return


}, 1)

tmp15189 := Call(__e, PrimFunc(symshen_4newpv), V4593)


tmp15190 := Call(__e, tmp15182, tmp15189)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp15190)
return


}, 1)

tmp15191 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp15181, tmp15191)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15196 := MakeNative(func(__e *ControlFlow) {
Z4662 := __e.Get(1)
_ = Z4662
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4663 := __e.Get(1)
_ = Z4663
tmp15197 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15197

tmp15198 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4658, Z4663, V4592, V4593, V4594, W4597, V4596)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4656, Z4662, V4592, V4593, V4594, W4597, tmp15198)
return


}, 1))
return
}, 1)

__e.TailApply(tmp15128, tmp15196)
return


}, 1)

tmp15199 := Call(__e, PrimFunc(symshen_4lazyderef), V4591, V4593)


__e.TailApply(tmp15127, tmp15199)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15202 := PrimTail(W4657)

tmp15203 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15202, V4593)


__e.TailApply(tmp15126, tmp15203)
return


}, 1)

tmp15204 := PrimHead(W4657)

__e.TailApply(tmp15125, tmp15204)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15207 := PrimTail(W4655)

tmp15208 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15207, V4593)


__e.TailApply(tmp15124, tmp15208)
return


}, 1)

tmp15209 := PrimHead(W4655)

__e.TailApply(tmp15123, tmp15209)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15212 := PrimTail(W4653)

tmp15213 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15212, V4593)


__e.TailApply(tmp15122, tmp15213)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15216 := PrimHead(W4653)

tmp15217 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15216, V4593)


__e.TailApply(tmp15121, tmp15217)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15220 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp15221 := Call(__e, tmp15120, tmp15220)


ifres15119 = tmp15221


} else {
ifres15119 = False


}

__e.TailApply(tmp14565, ifres15119)
return


} else {
__e.Return(W4631)
return
}


}, 1)

tmp15307 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres15225 Obj

if True == tmp15307 {
tmp15226 := MakeNative(func(__e *ControlFlow) {
W4632 := __e.Get(1)
_ = W4632
tmp15304 := PrimIsPair(W4632)

if True == tmp15304 {
tmp15227 := MakeNative(func(__e *ControlFlow) {
W4633 := __e.Get(1)
_ = W4633
tmp15300 := PrimEqual(W4633, symcons)

if True == tmp15300 {
tmp15228 := MakeNative(func(__e *ControlFlow) {
W4634 := __e.Get(1)
_ = W4634
tmp15296 := PrimIsPair(W4634)

if True == tmp15296 {
tmp15229 := MakeNative(func(__e *ControlFlow) {
W4635 := __e.Get(1)
_ = W4635
tmp15230 := MakeNative(func(__e *ControlFlow) {
W4636 := __e.Get(1)
_ = W4636
tmp15291 := PrimIsPair(W4636)

if True == tmp15291 {
tmp15231 := MakeNative(func(__e *ControlFlow) {
W4637 := __e.Get(1)
_ = W4637
tmp15232 := MakeNative(func(__e *ControlFlow) {
W4638 := __e.Get(1)
_ = W4638
tmp15286 := PrimEqual(W4638, Nil)

if True == tmp15286 {
tmp15233 := MakeNative(func(__e *ControlFlow) {
W4639 := __e.Get(1)
_ = W4639
tmp15234 := MakeNative(func(__e *ControlFlow) {
W4640 := __e.Get(1)
_ = W4640
tmp15278 := PrimIsPair(W4639)

if True == tmp15278 {
tmp15235 := MakeNative(func(__e *ControlFlow) {
W4642 := __e.Get(1)
_ = W4642
tmp15236 := MakeNative(func(__e *ControlFlow) {
W4643 := __e.Get(1)
_ = W4643
tmp15240 := PrimEqual(W4642, symlist)

if True == tmp15240 {
__e.TailApply(PrimFunc(symthaw), W4643)
return
} else {
tmp15238 := Call(__e, PrimFunc(symshen_4pvar_2), W4642)


if True == tmp15238 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4642, symlist, V4593, W4643)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15241 := MakeNative(func(__e *ControlFlow) {
tmp15242 := MakeNative(func(__e *ControlFlow) {
W4644 := __e.Get(1)
_ = W4644
tmp15243 := MakeNative(func(__e *ControlFlow) {
W4645 := __e.Get(1)
_ = W4645
tmp15263 := PrimIsPair(W4644)

if True == tmp15263 {
tmp15244 := MakeNative(func(__e *ControlFlow) {
W4647 := __e.Get(1)
_ = W4647
tmp15245 := MakeNative(func(__e *ControlFlow) {
W4648 := __e.Get(1)
_ = W4648
tmp15246 := MakeNative(func(__e *ControlFlow) {
W4649 := __e.Get(1)
_ = W4649
tmp15250 := PrimEqual(W4648, Nil)

if True == tmp15250 {
__e.TailApply(PrimFunc(symthaw), W4649)
return
} else {
tmp15248 := Call(__e, PrimFunc(symshen_4pvar_2), W4648)


if True == tmp15248 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4648, Nil, V4593, W4649)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15251 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4645, W4647)
return
}, 0)

__e.TailApply(tmp15246, tmp15251)
return


}, 1)

tmp15252 := PrimTail(W4644)

tmp15253 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15252, V4593)


__e.TailApply(tmp15245, tmp15253)
return


}, 1)

tmp15254 := PrimHead(W4644)

__e.TailApply(tmp15244, tmp15254)
return


} else {
tmp15261 := Call(__e, PrimFunc(symshen_4pvar_2), W4644)


if True == tmp15261 {
tmp15255 := MakeNative(func(__e *ControlFlow) {
W4650 := __e.Get(1)
_ = W4650
tmp15256 := PrimCons(W4650, Nil)

tmp15257 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4645, W4650)
return
}, 0)

tmp15258 := Call(__e, PrimFunc(symshen_4bind_b), W4644, tmp15256, V4593, tmp15257)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp15258)
return


}, 1)

tmp15259 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp15255, tmp15259)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15264 := MakeNative(func(__e *ControlFlow) {
Z4646 := __e.Get(1)
_ = Z4646
__e.TailApply(W4640, Z4646)
return
}, 1)

__e.TailApply(tmp15243, tmp15264)
return


}, 1)

tmp15265 := PrimTail(W4639)

tmp15266 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15265, V4593)


__e.TailApply(tmp15242, tmp15266)
return


}, 0)

__e.TailApply(tmp15236, tmp15241)
return


}, 1)

tmp15267 := PrimHead(W4639)

tmp15268 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15267, V4593)


__e.TailApply(tmp15235, tmp15268)
return


} else {
tmp15276 := Call(__e, PrimFunc(symshen_4pvar_2), W4639)


if True == tmp15276 {
tmp15269 := MakeNative(func(__e *ControlFlow) {
W4651 := __e.Get(1)
_ = W4651
tmp15270 := PrimCons(W4651, Nil)

tmp15271 := PrimCons(symlist, tmp15270)

tmp15272 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4640, W4651)
return
}, 0)

tmp15273 := Call(__e, PrimFunc(symshen_4bind_b), W4639, tmp15271, V4593, tmp15272)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp15273)
return


}, 1)

tmp15274 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp15269, tmp15274)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15279 := MakeNative(func(__e *ControlFlow) {
Z4641 := __e.Get(1)
_ = Z4641
tmp15280 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15280

tmp15281 := MakeNative(func(__e *ControlFlow) {
tmp15282 := PrimCons(Z4641, Nil)

tmp15283 := PrimCons(symlist, tmp15282)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4637, tmp15283, V4592, V4593, V4594, W4597, V4596)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4635, Z4641, V4592, V4593, V4594, W4597, tmp15281)
return


}, 1)

__e.TailApply(tmp15234, tmp15279)
return


}, 1)

tmp15284 := Call(__e, PrimFunc(symshen_4lazyderef), V4591, V4593)


__e.TailApply(tmp15233, tmp15284)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15287 := PrimTail(W4636)

tmp15288 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15287, V4593)


__e.TailApply(tmp15232, tmp15288)
return


}, 1)

tmp15289 := PrimHead(W4636)

__e.TailApply(tmp15231, tmp15289)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15292 := PrimTail(W4634)

tmp15293 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15292, V4593)


__e.TailApply(tmp15230, tmp15293)
return


}, 1)

tmp15294 := PrimHead(W4634)

__e.TailApply(tmp15229, tmp15294)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15297 := PrimTail(W4632)

tmp15298 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15297, V4593)


__e.TailApply(tmp15228, tmp15298)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15301 := PrimHead(W4632)

tmp15302 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15301, V4593)


__e.TailApply(tmp15227, tmp15302)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15305 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp15306 := Call(__e, tmp15226, tmp15305)


ifres15225 = tmp15306


} else {
ifres15225 = False


}

__e.TailApply(tmp14564, ifres15225)
return


} else {
__e.Return(W4624)
return
}


}, 1)

tmp15338 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres15310 Obj

if True == tmp15338 {
tmp15311 := MakeNative(func(__e *ControlFlow) {
W4625 := __e.Get(1)
_ = W4625
tmp15335 := PrimIsPair(W4625)

if True == tmp15335 {
tmp15312 := MakeNative(func(__e *ControlFlow) {
W4626 := __e.Get(1)
_ = W4626
tmp15313 := MakeNative(func(__e *ControlFlow) {
W4627 := __e.Get(1)
_ = W4627
tmp15330 := PrimIsPair(W4627)

if True == tmp15330 {
tmp15314 := MakeNative(func(__e *ControlFlow) {
W4628 := __e.Get(1)
_ = W4628
tmp15315 := MakeNative(func(__e *ControlFlow) {
W4629 := __e.Get(1)
_ = W4629
tmp15325 := PrimEqual(W4629, Nil)

if True == tmp15325 {
tmp15316 := MakeNative(func(__e *ControlFlow) {
W4630 := __e.Get(1)
_ = W4630
tmp15317 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15317

tmp15318 := PrimCons(V4591, Nil)

tmp15319 := PrimCons(sym_1_1_6, tmp15318)

tmp15320 := PrimCons(W4630, tmp15319)

tmp15321 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4628, W4630, V4592, V4593, V4594, W4597, V4596)
return
}, 0)

tmp15322 := Call(__e, PrimFunc(symshen_4system_1S_1h), W4626, tmp15320, V4592, V4593, V4594, W4597, tmp15321)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp15322)
return


}, 1)

tmp15323 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp15316, tmp15323)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15326 := PrimTail(W4627)

tmp15327 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15326, V4593)


__e.TailApply(tmp15315, tmp15327)
return


}, 1)

tmp15328 := PrimHead(W4627)

__e.TailApply(tmp15314, tmp15328)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15331 := PrimTail(W4625)

tmp15332 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15331, V4593)


__e.TailApply(tmp15313, tmp15332)
return


}, 1)

tmp15333 := PrimHead(W4625)

__e.TailApply(tmp15312, tmp15333)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15336 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp15337 := Call(__e, tmp15311, tmp15336)


ifres15310 = tmp15337


} else {
ifres15310 = False


}

__e.TailApply(tmp14563, ifres15310)
return


} else {
__e.Return(W4617)
return
}


}, 1)

tmp15373 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres15341 Obj

if True == tmp15373 {
tmp15342 := MakeNative(func(__e *ControlFlow) {
W4618 := __e.Get(1)
_ = W4618
tmp15370 := PrimIsPair(W4618)

if True == tmp15370 {
tmp15343 := MakeNative(func(__e *ControlFlow) {
W4619 := __e.Get(1)
_ = W4619
tmp15344 := MakeNative(func(__e *ControlFlow) {
W4620 := __e.Get(1)
_ = W4620
tmp15365 := PrimIsPair(W4620)

if True == tmp15365 {
tmp15345 := MakeNative(func(__e *ControlFlow) {
W4621 := __e.Get(1)
_ = W4621
tmp15346 := MakeNative(func(__e *ControlFlow) {
W4622 := __e.Get(1)
_ = W4622
tmp15360 := PrimEqual(W4622, Nil)

if True == tmp15360 {
tmp15347 := MakeNative(func(__e *ControlFlow) {
W4623 := __e.Get(1)
_ = W4623
tmp15348 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15348

tmp15349 := Call(__e, PrimFunc(symshen_4lazyderef), W4619, V4593)


tmp15350 := PrimIsPair(tmp15349)

tmp15351 := PrimNot(tmp15350)

tmp15352 := MakeNative(func(__e *ControlFlow) {
tmp15353 := PrimCons(V4591, Nil)

tmp15354 := PrimCons(sym_1_1_6, tmp15353)

tmp15355 := PrimCons(W4623, tmp15354)

tmp15356 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4621, W4623, V4592, V4593, V4594, W4597, V4596)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4lookupsig), W4619, tmp15355, V4593, V4594, W4597, tmp15356)
return


}, 0)

tmp15357 := Call(__e, PrimFunc(symwhen), tmp15351, V4593, V4594, W4597, tmp15352)


__e.TailApply(PrimFunc(symshen_4gc), V4593, tmp15357)
return


}, 1)

tmp15358 := Call(__e, PrimFunc(symshen_4newpv), V4593)


__e.TailApply(tmp15347, tmp15358)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15361 := PrimTail(W4620)

tmp15362 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15361, V4593)


__e.TailApply(tmp15346, tmp15362)
return


}, 1)

tmp15363 := PrimHead(W4620)

__e.TailApply(tmp15345, tmp15363)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15366 := PrimTail(W4618)

tmp15367 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15366, V4593)


__e.TailApply(tmp15344, tmp15367)
return


}, 1)

tmp15368 := PrimHead(W4618)

__e.TailApply(tmp15343, tmp15368)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15371 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp15372 := Call(__e, tmp15342, tmp15371)


ifres15341 = tmp15372


} else {
ifres15341 = False


}

__e.TailApply(tmp14562, ifres15341)
return


} else {
__e.Return(W4611)
return
}


}, 1)

tmp15400 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres15376 Obj

if True == tmp15400 {
tmp15377 := MakeNative(func(__e *ControlFlow) {
W4612 := __e.Get(1)
_ = W4612
tmp15397 := PrimIsPair(W4612)

if True == tmp15397 {
tmp15378 := MakeNative(func(__e *ControlFlow) {
W4613 := __e.Get(1)
_ = W4613
tmp15393 := PrimEqual(W4613, symfn)

if True == tmp15393 {
tmp15379 := MakeNative(func(__e *ControlFlow) {
W4614 := __e.Get(1)
_ = W4614
tmp15389 := PrimIsPair(W4614)

if True == tmp15389 {
tmp15380 := MakeNative(func(__e *ControlFlow) {
W4615 := __e.Get(1)
_ = W4615
tmp15381 := MakeNative(func(__e *ControlFlow) {
W4616 := __e.Get(1)
_ = W4616
tmp15384 := PrimEqual(W4616, Nil)

if True == tmp15384 {
tmp15382 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15382

__e.TailApply(PrimFunc(symshen_4lookupsig), W4615, V4591, V4593, V4594, W4597, V4596)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15385 := PrimTail(W4614)

tmp15386 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15385, V4593)


__e.TailApply(tmp15381, tmp15386)
return


}, 1)

tmp15387 := PrimHead(W4614)

__e.TailApply(tmp15380, tmp15387)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15390 := PrimTail(W4612)

tmp15391 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15390, V4593)


__e.TailApply(tmp15379, tmp15391)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15394 := PrimHead(W4612)

tmp15395 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15394, V4593)


__e.TailApply(tmp15378, tmp15395)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15398 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp15399 := Call(__e, tmp15377, tmp15398)


ifres15376 = tmp15399


} else {
ifres15376 = False


}

__e.TailApply(tmp14561, ifres15376)
return


} else {
__e.Return(W4605)
return
}


}, 1)

tmp15433 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres15403 Obj

if True == tmp15433 {
tmp15404 := MakeNative(func(__e *ControlFlow) {
W4606 := __e.Get(1)
_ = W4606
tmp15430 := PrimIsPair(W4606)

if True == tmp15430 {
tmp15405 := MakeNative(func(__e *ControlFlow) {
W4607 := __e.Get(1)
_ = W4607
tmp15426 := PrimEqual(W4607, symfn)

if True == tmp15426 {
tmp15406 := MakeNative(func(__e *ControlFlow) {
W4608 := __e.Get(1)
_ = W4608
tmp15422 := PrimIsPair(W4608)

if True == tmp15422 {
tmp15407 := MakeNative(func(__e *ControlFlow) {
W4609 := __e.Get(1)
_ = W4609
tmp15408 := MakeNative(func(__e *ControlFlow) {
W4610 := __e.Get(1)
_ = W4610
tmp15417 := PrimEqual(W4610, Nil)

if True == tmp15417 {
tmp15409 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15409

tmp15410 := Call(__e, PrimFunc(symshen_4deref), W4609, V4593)


tmp15411 := Call(__e, PrimFunc(symarity), tmp15410)


tmp15412 := PrimEqual(tmp15411, MakeNumber(0))

tmp15413 := MakeNative(func(__e *ControlFlow) {
tmp15414 := MakeNative(func(__e *ControlFlow) {
tmp15415 := PrimCons(W4609, Nil)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp15415, V4591, V4592, V4593, V4594, W4597, V4596)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4593, V4594, W4597, tmp15414)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15412, V4593, V4594, W4597, tmp15413)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15418 := PrimTail(W4608)

tmp15419 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15418, V4593)


__e.TailApply(tmp15408, tmp15419)
return


}, 1)

tmp15420 := PrimHead(W4608)

__e.TailApply(tmp15407, tmp15420)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15423 := PrimTail(W4606)

tmp15424 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15423, V4593)


__e.TailApply(tmp15406, tmp15424)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15427 := PrimHead(W4606)

tmp15428 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15427, V4593)


__e.TailApply(tmp15405, tmp15428)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15431 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp15432 := Call(__e, tmp15404, tmp15431)


ifres15403 = tmp15432


} else {
ifres15403 = False


}

__e.TailApply(tmp14560, ifres15403)
return


} else {
__e.Return(W4601)
return
}


}, 1)

tmp15452 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres15436 Obj

if True == tmp15452 {
tmp15437 := MakeNative(func(__e *ControlFlow) {
W4602 := __e.Get(1)
_ = W4602
tmp15449 := PrimIsPair(W4602)

if True == tmp15449 {
tmp15438 := MakeNative(func(__e *ControlFlow) {
W4603 := __e.Get(1)
_ = W4603
tmp15439 := MakeNative(func(__e *ControlFlow) {
W4604 := __e.Get(1)
_ = W4604
tmp15444 := PrimEqual(W4604, Nil)

if True == tmp15444 {
tmp15440 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15440

tmp15441 := PrimCons(V4591, Nil)

tmp15442 := PrimCons(sym_1_1_6, tmp15441)

__e.TailApply(PrimFunc(symshen_4lookupsig), W4603, tmp15442, V4593, V4594, W4597, V4596)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15445 := PrimTail(W4602)

tmp15446 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15445, V4593)


__e.TailApply(tmp15439, tmp15446)
return


}, 1)

tmp15447 := PrimHead(W4602)

__e.TailApply(tmp15438, tmp15447)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15450 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp15451 := Call(__e, tmp15437, tmp15450)


ifres15436 = tmp15451


} else {
ifres15436 = False


}

__e.TailApply(tmp14559, ifres15436)
return


} else {
__e.Return(W4600)
return
}


}, 1)

tmp15458 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres15455 Obj

if True == tmp15458 {
tmp15456 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15456

tmp15457 := Call(__e, PrimFunc(symshen_4by_1hypothesis), V4590, V4591, V4592, V4593, V4594, W4597, V4596)


ifres15455 = tmp15457


} else {
ifres15455 = False


}

__e.TailApply(tmp14558, ifres15455)
return


} else {
__e.Return(W4599)
return
}


}, 1)

tmp15468 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres15461 Obj

if True == tmp15468 {
tmp15462 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15462

tmp15463 := Call(__e, PrimFunc(symshen_4lazyderef), V4590, V4593)


tmp15464 := PrimIsPair(tmp15463)

tmp15465 := PrimNot(tmp15464)

tmp15466 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4primitive), V4590, V4591, V4593, V4594, W4597, V4596)
return
}, 0)

tmp15467 := Call(__e, PrimFunc(symwhen), tmp15465, V4593, V4594, W4597, tmp15466)


ifres15461 = tmp15467


} else {
ifres15461 = False


}

__e.TailApply(tmp14557, ifres15461)
return


} else {
__e.Return(W4598)
return
}


}, 1)

tmp15480 := Call(__e, PrimFunc(symshen_4unlocked_2), V4594)


var ifres15471 Obj

if True == tmp15480 {
tmp15472 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15472

tmp15473 := PrimValue(symshen_4_dspy_d)

tmp15474 := MakeNative(func(__e *ControlFlow) {
tmp15475 := PrimIntern(MakeString(":"))

tmp15476 := PrimCons(V4591, Nil)

tmp15477 := PrimCons(tmp15475, tmp15476)

tmp15478 := PrimCons(V4590, tmp15477)

__e.TailApply(PrimFunc(symshen_4show), tmp15478, V4592, V4593, V4594, W4597, V4596)
return


}, 0)

tmp15479 := Call(__e, PrimFunc(symwhen), tmp15473, V4593, V4594, W4597, tmp15474)


ifres15471 = tmp15479


} else {
ifres15471 = False


}

__e.TailApply(tmp14556, ifres15471)
return


}, 1)

tmp15481 := PrimNumberAdd(V4595, MakeNumber(1))

__e.TailApply(tmp14555, tmp15481)
return


}, 7)

tmp15482 := Call(__e, ns2_1set, symshen_4system_1S_1h, tmp14554)


_ = tmp15482

tmp15483 := MakeNative(func(__e *ControlFlow) {
V4802 := __e.Get(1)
_ = V4802
V4803 := __e.Get(2)
_ = V4803
V4804 := __e.Get(3)
_ = V4804
V4805 := __e.Get(4)
_ = V4805
V4806 := __e.Get(5)
_ = V4806
V4807 := __e.Get(6)
_ = V4807
tmp15484 := MakeNative(func(__e *ControlFlow) {
W4808 := __e.Get(1)
_ = W4808
tmp15592 := PrimEqual(W4808, False)

if True == tmp15592 {
tmp15485 := MakeNative(func(__e *ControlFlow) {
W4811 := __e.Get(1)
_ = W4811
tmp15576 := PrimEqual(W4811, False)

if True == tmp15576 {
tmp15486 := MakeNative(func(__e *ControlFlow) {
W4814 := __e.Get(1)
_ = W4814
tmp15560 := PrimEqual(W4814, False)

if True == tmp15560 {
tmp15487 := MakeNative(func(__e *ControlFlow) {
W4817 := __e.Get(1)
_ = W4817
tmp15544 := PrimEqual(W4817, False)

if True == tmp15544 {
tmp15542 := Call(__e, PrimFunc(symshen_4unlocked_2), V4805)


if True == tmp15542 {
tmp15488 := MakeNative(func(__e *ControlFlow) {
W4820 := __e.Get(1)
_ = W4820
tmp15539 := PrimEqual(W4820, Nil)

if True == tmp15539 {
tmp15489 := MakeNative(func(__e *ControlFlow) {
W4821 := __e.Get(1)
_ = W4821
tmp15490 := MakeNative(func(__e *ControlFlow) {
W4822 := __e.Get(1)
_ = W4822
tmp15534 := PrimIsPair(W4821)

if True == tmp15534 {
tmp15491 := MakeNative(func(__e *ControlFlow) {
W4824 := __e.Get(1)
_ = W4824
tmp15492 := MakeNative(func(__e *ControlFlow) {
W4825 := __e.Get(1)
_ = W4825
tmp15496 := PrimEqual(W4824, symlist)

if True == tmp15496 {
__e.TailApply(PrimFunc(symthaw), W4825)
return
} else {
tmp15494 := Call(__e, PrimFunc(symshen_4pvar_2), W4824)


if True == tmp15494 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4824, symlist, V4804, W4825)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15497 := MakeNative(func(__e *ControlFlow) {
tmp15498 := MakeNative(func(__e *ControlFlow) {
W4826 := __e.Get(1)
_ = W4826
tmp15499 := MakeNative(func(__e *ControlFlow) {
W4827 := __e.Get(1)
_ = W4827
tmp15519 := PrimIsPair(W4826)

if True == tmp15519 {
tmp15500 := MakeNative(func(__e *ControlFlow) {
W4829 := __e.Get(1)
_ = W4829
tmp15501 := MakeNative(func(__e *ControlFlow) {
W4830 := __e.Get(1)
_ = W4830
tmp15502 := MakeNative(func(__e *ControlFlow) {
W4831 := __e.Get(1)
_ = W4831
tmp15506 := PrimEqual(W4830, Nil)

if True == tmp15506 {
__e.TailApply(PrimFunc(symthaw), W4831)
return
} else {
tmp15504 := Call(__e, PrimFunc(symshen_4pvar_2), W4830)


if True == tmp15504 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4830, Nil, V4804, W4831)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15507 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4827, W4829)
return
}, 0)

__e.TailApply(tmp15502, tmp15507)
return


}, 1)

tmp15508 := PrimTail(W4826)

tmp15509 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15508, V4804)


__e.TailApply(tmp15501, tmp15509)
return


}, 1)

tmp15510 := PrimHead(W4826)

__e.TailApply(tmp15500, tmp15510)
return


} else {
tmp15517 := Call(__e, PrimFunc(symshen_4pvar_2), W4826)


if True == tmp15517 {
tmp15511 := MakeNative(func(__e *ControlFlow) {
W4832 := __e.Get(1)
_ = W4832
tmp15512 := PrimCons(W4832, Nil)

tmp15513 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4827, W4832)
return
}, 0)

tmp15514 := Call(__e, PrimFunc(symshen_4bind_b), W4826, tmp15512, V4804, tmp15513)


__e.TailApply(PrimFunc(symshen_4gc), V4804, tmp15514)
return


}, 1)

tmp15515 := Call(__e, PrimFunc(symshen_4newpv), V4804)


__e.TailApply(tmp15511, tmp15515)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15520 := MakeNative(func(__e *ControlFlow) {
Z4828 := __e.Get(1)
_ = Z4828
__e.TailApply(W4822, Z4828)
return
}, 1)

__e.TailApply(tmp15499, tmp15520)
return


}, 1)

tmp15521 := PrimTail(W4821)

tmp15522 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15521, V4804)


__e.TailApply(tmp15498, tmp15522)
return


}, 0)

__e.TailApply(tmp15492, tmp15497)
return


}, 1)

tmp15523 := PrimHead(W4821)

tmp15524 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15523, V4804)


__e.TailApply(tmp15491, tmp15524)
return


} else {
tmp15532 := Call(__e, PrimFunc(symshen_4pvar_2), W4821)


if True == tmp15532 {
tmp15525 := MakeNative(func(__e *ControlFlow) {
W4833 := __e.Get(1)
_ = W4833
tmp15526 := PrimCons(W4833, Nil)

tmp15527 := PrimCons(symlist, tmp15526)

tmp15528 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4822, W4833)
return
}, 0)

tmp15529 := Call(__e, PrimFunc(symshen_4bind_b), W4821, tmp15527, V4804, tmp15528)


__e.TailApply(PrimFunc(symshen_4gc), V4804, tmp15529)
return


}, 1)

tmp15530 := Call(__e, PrimFunc(symshen_4newpv), V4804)


__e.TailApply(tmp15525, tmp15530)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15535 := MakeNative(func(__e *ControlFlow) {
Z4823 := __e.Get(1)
_ = Z4823
tmp15536 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15536

__e.TailApply(PrimFunc(symthaw), V4807)
return


}, 1)

__e.TailApply(tmp15490, tmp15535)
return


}, 1)

tmp15537 := Call(__e, PrimFunc(symshen_4lazyderef), V4803, V4804)


__e.TailApply(tmp15489, tmp15537)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15540 := Call(__e, PrimFunc(symshen_4lazyderef), V4802, V4804)


__e.TailApply(tmp15488, tmp15540)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4817)
return
}


}, 1)

tmp15558 := Call(__e, PrimFunc(symshen_4unlocked_2), V4805)


var ifres15545 Obj

if True == tmp15558 {
tmp15546 := MakeNative(func(__e *ControlFlow) {
W4818 := __e.Get(1)
_ = W4818
tmp15547 := MakeNative(func(__e *ControlFlow) {
W4819 := __e.Get(1)
_ = W4819
tmp15551 := PrimEqual(W4818, symsymbol)

if True == tmp15551 {
__e.TailApply(PrimFunc(symthaw), W4819)
return
} else {
tmp15549 := Call(__e, PrimFunc(symshen_4pvar_2), W4818)


if True == tmp15549 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4818, symsymbol, V4804, W4819)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15552 := MakeNative(func(__e *ControlFlow) {
tmp15553 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15553

tmp15554 := Call(__e, PrimFunc(symshen_4lazyderef), V4802, V4804)


tmp15555 := PrimIsSymbol(tmp15554)

__e.TailApply(PrimFunc(symwhen), tmp15555, V4804, V4805, V4806, V4807)
return


}, 0)

__e.TailApply(tmp15547, tmp15552)
return


}, 1)

tmp15556 := Call(__e, PrimFunc(symshen_4lazyderef), V4803, V4804)


tmp15557 := Call(__e, tmp15546, tmp15556)


ifres15545 = tmp15557


} else {
ifres15545 = False


}

__e.TailApply(tmp15487, ifres15545)
return


} else {
__e.Return(W4814)
return
}


}, 1)

tmp15574 := Call(__e, PrimFunc(symshen_4unlocked_2), V4805)


var ifres15561 Obj

if True == tmp15574 {
tmp15562 := MakeNative(func(__e *ControlFlow) {
W4815 := __e.Get(1)
_ = W4815
tmp15563 := MakeNative(func(__e *ControlFlow) {
W4816 := __e.Get(1)
_ = W4816
tmp15567 := PrimEqual(W4815, symstring)

if True == tmp15567 {
__e.TailApply(PrimFunc(symthaw), W4816)
return
} else {
tmp15565 := Call(__e, PrimFunc(symshen_4pvar_2), W4815)


if True == tmp15565 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4815, symstring, V4804, W4816)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15568 := MakeNative(func(__e *ControlFlow) {
tmp15569 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15569

tmp15570 := Call(__e, PrimFunc(symshen_4lazyderef), V4802, V4804)


tmp15571 := PrimIsString(tmp15570)

__e.TailApply(PrimFunc(symwhen), tmp15571, V4804, V4805, V4806, V4807)
return


}, 0)

__e.TailApply(tmp15563, tmp15568)
return


}, 1)

tmp15572 := Call(__e, PrimFunc(symshen_4lazyderef), V4803, V4804)


tmp15573 := Call(__e, tmp15562, tmp15572)


ifres15561 = tmp15573


} else {
ifres15561 = False


}

__e.TailApply(tmp15486, ifres15561)
return


} else {
__e.Return(W4811)
return
}


}, 1)

tmp15590 := Call(__e, PrimFunc(symshen_4unlocked_2), V4805)


var ifres15577 Obj

if True == tmp15590 {
tmp15578 := MakeNative(func(__e *ControlFlow) {
W4812 := __e.Get(1)
_ = W4812
tmp15579 := MakeNative(func(__e *ControlFlow) {
W4813 := __e.Get(1)
_ = W4813
tmp15583 := PrimEqual(W4812, symboolean)

if True == tmp15583 {
__e.TailApply(PrimFunc(symthaw), W4813)
return
} else {
tmp15581 := Call(__e, PrimFunc(symshen_4pvar_2), W4812)


if True == tmp15581 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4812, symboolean, V4804, W4813)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15584 := MakeNative(func(__e *ControlFlow) {
tmp15585 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15585

tmp15586 := Call(__e, PrimFunc(symshen_4lazyderef), V4802, V4804)


tmp15587 := Call(__e, PrimFunc(symboolean_2), tmp15586)


__e.TailApply(PrimFunc(symwhen), tmp15587, V4804, V4805, V4806, V4807)
return


}, 0)

__e.TailApply(tmp15579, tmp15584)
return


}, 1)

tmp15588 := Call(__e, PrimFunc(symshen_4lazyderef), V4803, V4804)


tmp15589 := Call(__e, tmp15578, tmp15588)


ifres15577 = tmp15589


} else {
ifres15577 = False


}

__e.TailApply(tmp15485, ifres15577)
return


} else {
__e.Return(W4808)
return
}


}, 1)

tmp15606 := Call(__e, PrimFunc(symshen_4unlocked_2), V4805)


var ifres15593 Obj

if True == tmp15606 {
tmp15594 := MakeNative(func(__e *ControlFlow) {
W4809 := __e.Get(1)
_ = W4809
tmp15595 := MakeNative(func(__e *ControlFlow) {
W4810 := __e.Get(1)
_ = W4810
tmp15599 := PrimEqual(W4809, symnumber)

if True == tmp15599 {
__e.TailApply(PrimFunc(symthaw), W4810)
return
} else {
tmp15597 := Call(__e, PrimFunc(symshen_4pvar_2), W4809)


if True == tmp15597 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4809, symnumber, V4804, W4810)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15600 := MakeNative(func(__e *ControlFlow) {
tmp15601 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15601

tmp15602 := Call(__e, PrimFunc(symshen_4lazyderef), V4802, V4804)


tmp15603 := PrimIsNumber(tmp15602)

__e.TailApply(PrimFunc(symwhen), tmp15603, V4804, V4805, V4806, V4807)
return


}, 0)

__e.TailApply(tmp15595, tmp15600)
return


}, 1)

tmp15604 := Call(__e, PrimFunc(symshen_4lazyderef), V4803, V4804)


tmp15605 := Call(__e, tmp15594, tmp15604)


ifres15593 = tmp15605


} else {
ifres15593 = False


}

__e.TailApply(tmp15484, ifres15593)
return


}, 6)

tmp15607 := Call(__e, ns2_1set, symshen_4primitive, tmp15483)


_ = tmp15607

tmp15608 := MakeNative(func(__e *ControlFlow) {
V4834 := __e.Get(1)
_ = V4834
V4835 := __e.Get(2)
_ = V4835
V4836 := __e.Get(3)
_ = V4836
V4837 := __e.Get(4)
_ = V4837
V4838 := __e.Get(5)
_ = V4838
V4839 := __e.Get(6)
_ = V4839
V4840 := __e.Get(7)
_ = V4840
tmp15609 := MakeNative(func(__e *ControlFlow) {
W4841 := __e.Get(1)
_ = W4841
tmp15620 := PrimEqual(W4841, False)

if True == tmp15620 {
tmp15618 := Call(__e, PrimFunc(symshen_4unlocked_2), V4838)


if True == tmp15618 {
tmp15610 := MakeNative(func(__e *ControlFlow) {
W4850 := __e.Get(1)
_ = W4850
tmp15615 := PrimIsPair(W4850)

if True == tmp15615 {
tmp15611 := MakeNative(func(__e *ControlFlow) {
W4851 := __e.Get(1)
_ = W4851
tmp15612 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15612

__e.TailApply(PrimFunc(symshen_4by_1hypothesis), V4834, V4835, W4851, V4837, V4838, V4839, V4840)
return


}, 1)

tmp15613 := PrimTail(W4850)

__e.TailApply(tmp15611, tmp15613)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15616 := Call(__e, PrimFunc(symshen_4lazyderef), V4836, V4837)


__e.TailApply(tmp15610, tmp15616)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4841)
return
}


}, 1)

tmp15662 := Call(__e, PrimFunc(symshen_4unlocked_2), V4838)


var ifres15621 Obj

if True == tmp15662 {
tmp15622 := MakeNative(func(__e *ControlFlow) {
W4842 := __e.Get(1)
_ = W4842
tmp15659 := PrimIsPair(W4842)

if True == tmp15659 {
tmp15623 := MakeNative(func(__e *ControlFlow) {
W4843 := __e.Get(1)
_ = W4843
tmp15655 := PrimIsPair(W4843)

if True == tmp15655 {
tmp15624 := MakeNative(func(__e *ControlFlow) {
W4844 := __e.Get(1)
_ = W4844
tmp15625 := MakeNative(func(__e *ControlFlow) {
W4845 := __e.Get(1)
_ = W4845
tmp15650 := PrimIsPair(W4845)

if True == tmp15650 {
tmp15626 := MakeNative(func(__e *ControlFlow) {
W4846 := __e.Get(1)
_ = W4846
tmp15627 := MakeNative(func(__e *ControlFlow) {
W4847 := __e.Get(1)
_ = W4847
tmp15645 := PrimIsPair(W4847)

if True == tmp15645 {
tmp15628 := MakeNative(func(__e *ControlFlow) {
W4848 := __e.Get(1)
_ = W4848
tmp15629 := MakeNative(func(__e *ControlFlow) {
W4849 := __e.Get(1)
_ = W4849
tmp15640 := PrimEqual(W4849, Nil)

if True == tmp15640 {
tmp15630 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15630

tmp15631 := Call(__e, PrimFunc(symshen_4deref), W4846, V4837)


tmp15632 := PrimIntern(MakeString(":"))

tmp15633 := PrimEqual(tmp15631, tmp15632)

tmp15634 := MakeNative(func(__e *ControlFlow) {
tmp15635 := Call(__e, PrimFunc(symshen_4deref), V4834, V4837)


tmp15636 := Call(__e, PrimFunc(symshen_4deref), W4844, V4837)


tmp15637 := PrimEqual(tmp15635, tmp15636)

tmp15638 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis_b), V4835, W4848, V4837, V4838, V4839, V4840)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15637, V4837, V4838, V4839, tmp15638)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15633, V4837, V4838, V4839, tmp15634)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15641 := PrimTail(W4847)

tmp15642 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15641, V4837)


__e.TailApply(tmp15629, tmp15642)
return


}, 1)

tmp15643 := PrimHead(W4847)

__e.TailApply(tmp15628, tmp15643)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15646 := PrimTail(W4845)

tmp15647 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15646, V4837)


__e.TailApply(tmp15627, tmp15647)
return


}, 1)

tmp15648 := PrimHead(W4845)

__e.TailApply(tmp15626, tmp15648)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15651 := PrimTail(W4843)

tmp15652 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15651, V4837)


__e.TailApply(tmp15625, tmp15652)
return


}, 1)

tmp15653 := PrimHead(W4843)

__e.TailApply(tmp15624, tmp15653)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15656 := PrimHead(W4842)

tmp15657 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15656, V4837)


__e.TailApply(tmp15623, tmp15657)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15660 := Call(__e, PrimFunc(symshen_4lazyderef), V4836, V4837)


tmp15661 := Call(__e, tmp15622, tmp15660)


ifres15621 = tmp15661


} else {
ifres15621 = False


}

__e.TailApply(tmp15609, ifres15621)
return


}, 7)

tmp15663 := Call(__e, ns2_1set, symshen_4by_1hypothesis, tmp15608)


_ = tmp15663

tmp15664 := MakeNative(func(__e *ControlFlow) {
V4852 := __e.Get(1)
_ = V4852
V4853 := __e.Get(2)
_ = V4853
V4854 := __e.Get(3)
_ = V4854
V4855 := __e.Get(4)
_ = V4855
V4856 := __e.Get(5)
_ = V4856
V4857 := __e.Get(6)
_ = V4857
tmp15669 := Call(__e, PrimFunc(symshen_4unlocked_2), V4855)


if True == tmp15669 {
tmp15665 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15665

tmp15666 := PrimValue(symshen_4_dsigf_d)

tmp15667 := Call(__e, PrimFunc(symassoc), V4852, tmp15666)


__e.TailApply(PrimFunc(symshen_4sigf), tmp15667, V4853, V4854, V4855, V4856, V4857)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp15670 := Call(__e, ns2_1set, symshen_4lookupsig, tmp15664)


_ = tmp15670

tmp15671 := MakeNative(func(__e *ControlFlow) {
V4872 := __e.Get(1)
_ = V4872
V4873 := __e.Get(2)
_ = V4873
V4874 := __e.Get(3)
_ = V4874
V4875 := __e.Get(4)
_ = V4875
V4876 := __e.Get(5)
_ = V4876
V4877 := __e.Get(6)
_ = V4877
tmp15678 := PrimIsPair(V4872)

if True == tmp15678 {
tmp15672 := PrimTail(V4872)

tmp15673 := Call(__e, tmp15672, V4873)


tmp15674 := Call(__e, tmp15673, V4874)


tmp15675 := Call(__e, tmp15674, V4875)


tmp15676 := Call(__e, tmp15675, V4876)


__e.TailApply(tmp15676, V4877)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp15679 := Call(__e, ns2_1set, symshen_4sigf, tmp15671)


_ = tmp15679

tmp15680 := MakeNative(func(__e *ControlFlow) {
V4878 := __e.Get(1)
_ = V4878
tmp15681 := MakeNative(func(__e *ControlFlow) {
W4879 := __e.Get(1)
_ = W4879
tmp15682 := MakeNative(func(__e *ControlFlow) {
W4880 := __e.Get(1)
_ = W4880
tmp15683 := MakeNative(func(__e *ControlFlow) {
W4881 := __e.Get(1)
_ = W4881
tmp15684 := MakeNative(func(__e *ControlFlow) {
W4882 := __e.Get(1)
_ = W4882
__e.Return(W4882)
return
}, 1)

tmp15685 := PrimValue(symshen_4_dgensym_d)

tmp15686 := PrimNumberAdd(MakeNumber(1), tmp15685)

tmp15687 := PrimSet(symshen_4_dgensym_d, tmp15686)

tmp15688 := PrimVectorSet(W4881, MakeNumber(2), tmp15687)

__e.TailApply(tmp15684, tmp15688)
return


}, 1)

tmp15689 := PrimVectorSet(W4880, MakeNumber(1), V4878)

__e.TailApply(tmp15683, tmp15689)
return


}, 1)

tmp15690 := PrimVectorSet(W4879, MakeNumber(0), symshen_4print_1freshterm)

__e.TailApply(tmp15682, tmp15690)
return


}, 1)

tmp15691 := PrimAbsvector(MakeNumber(3))

__e.TailApply(tmp15681, tmp15691)
return


}, 1)

tmp15692 := Call(__e, ns2_1set, symshen_4freshterm, tmp15680)


_ = tmp15692

tmp15693 := MakeNative(func(__e *ControlFlow) {
V4883 := __e.Get(1)
_ = V4883
tmp15694 := PrimVectorGet(V4883, MakeNumber(1))

tmp15695 := PrimStr(tmp15694)

__e.Return(PrimStringConcat(MakeString("&&"), tmp15695))
return


}, 1)

tmp15696 := Call(__e, ns2_1set, symshen_4print_1freshterm, tmp15693)


_ = tmp15696

tmp15697 := MakeNative(func(__e *ControlFlow) {
V4884 := __e.Get(1)
_ = V4884
V4885 := __e.Get(2)
_ = V4885
V4886 := __e.Get(3)
_ = V4886
V4887 := __e.Get(4)
_ = V4887
V4888 := __e.Get(5)
_ = V4888
V4889 := __e.Get(6)
_ = V4889
V4890 := __e.Get(7)
_ = V4890
tmp15698 := MakeNative(func(__e *ControlFlow) {
W4891 := __e.Get(1)
_ = W4891
tmp15709 := PrimEqual(W4891, False)

if True == tmp15709 {
tmp15707 := Call(__e, PrimFunc(symshen_4unlocked_2), V4888)


if True == tmp15707 {
tmp15699 := MakeNative(func(__e *ControlFlow) {
W4895 := __e.Get(1)
_ = W4895
tmp15704 := PrimIsPair(W4895)

if True == tmp15704 {
tmp15700 := MakeNative(func(__e *ControlFlow) {
W4896 := __e.Get(1)
_ = W4896
tmp15701 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15701

__e.TailApply(PrimFunc(symshen_4search_1user_1datatypes), V4884, V4885, W4896, V4887, V4888, V4889, V4890)
return


}, 1)

tmp15702 := PrimTail(W4895)

__e.TailApply(tmp15700, tmp15702)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15705 := Call(__e, PrimFunc(symshen_4lazyderef), V4886, V4887)


__e.TailApply(tmp15699, tmp15705)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4891)
return
}


}, 1)

tmp15729 := Call(__e, PrimFunc(symshen_4unlocked_2), V4888)


var ifres15710 Obj

if True == tmp15729 {
tmp15711 := MakeNative(func(__e *ControlFlow) {
W4892 := __e.Get(1)
_ = W4892
tmp15726 := PrimIsPair(W4892)

if True == tmp15726 {
tmp15712 := MakeNative(func(__e *ControlFlow) {
W4893 := __e.Get(1)
_ = W4893
tmp15722 := PrimIsPair(W4893)

if True == tmp15722 {
tmp15713 := MakeNative(func(__e *ControlFlow) {
W4894 := __e.Get(1)
_ = W4894
tmp15714 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15714

tmp15715 := Call(__e, PrimFunc(symshen_4deref), W4894, V4887)


tmp15716 := Call(__e, PrimFunc(symshen_4deref), V4884, V4887)


tmp15717 := Call(__e, tmp15715, tmp15716)


tmp15718 := Call(__e, PrimFunc(symshen_4deref), V4885, V4887)


tmp15719 := Call(__e, tmp15717, tmp15718)


__e.TailApply(PrimFunc(symcall), tmp15719, V4887, V4888, V4889, V4890)
return


}, 1)

tmp15720 := PrimTail(W4893)

__e.TailApply(tmp15713, tmp15720)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15723 := PrimHead(W4892)

tmp15724 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15723, V4887)


__e.TailApply(tmp15712, tmp15724)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15727 := Call(__e, PrimFunc(symshen_4lazyderef), V4886, V4887)


tmp15728 := Call(__e, tmp15711, tmp15727)


ifres15710 = tmp15728


} else {
ifres15710 = False


}

__e.TailApply(tmp15698, ifres15710)
return


}, 7)

tmp15730 := Call(__e, ns2_1set, symshen_4search_1user_1datatypes, tmp15697)


_ = tmp15730

tmp15731 := MakeNative(func(__e *ControlFlow) {
V4897 := __e.Get(1)
_ = V4897
V4898 := __e.Get(2)
_ = V4898
V4899 := __e.Get(3)
_ = V4899
V4900 := __e.Get(4)
_ = V4900
V4901 := __e.Get(5)
_ = V4901
V4902 := __e.Get(6)
_ = V4902
V4903 := __e.Get(7)
_ = V4903
tmp15732 := MakeNative(func(__e *ControlFlow) {
W4904 := __e.Get(1)
_ = W4904
tmp15733 := MakeNative(func(__e *ControlFlow) {
W4905 := __e.Get(1)
_ = W4905
tmp16163 := PrimEqual(W4905, False)

if True == tmp16163 {
tmp15734 := MakeNative(func(__e *ControlFlow) {
W4908 := __e.Get(1)
_ = W4908
tmp16063 := PrimEqual(W4908, False)

if True == tmp16063 {
tmp15735 := MakeNative(func(__e *ControlFlow) {
W4928 := __e.Get(1)
_ = W4928
tmp15958 := PrimEqual(W4928, False)

if True == tmp15958 {
tmp15736 := MakeNative(func(__e *ControlFlow) {
W4950 := __e.Get(1)
_ = W4950
tmp15877 := PrimEqual(W4950, False)

if True == tmp15877 {
tmp15737 := MakeNative(func(__e *ControlFlow) {
W4966 := __e.Get(1)
_ = W4966
tmp15777 := PrimEqual(W4966, False)

if True == tmp15777 {
tmp15738 := MakeNative(func(__e *ControlFlow) {
W4986 := __e.Get(1)
_ = W4986
tmp15740 := PrimEqual(W4986, False)

if True == tmp15740 {
__e.TailApply(PrimFunc(symshen_4unlock), V4901, W4904)
return
} else {
__e.Return(W4986)
return
}


}, 1)

tmp15775 := Call(__e, PrimFunc(symshen_4unlocked_2), V4901)


var ifres15741 Obj

if True == tmp15775 {
tmp15742 := MakeNative(func(__e *ControlFlow) {
W4987 := __e.Get(1)
_ = W4987
tmp15772 := PrimIsPair(W4987)

if True == tmp15772 {
tmp15743 := MakeNative(func(__e *ControlFlow) {
W4988 := __e.Get(1)
_ = W4988
tmp15744 := MakeNative(func(__e *ControlFlow) {
W4989 := __e.Get(1)
_ = W4989
tmp15745 := MakeNative(func(__e *ControlFlow) {
W4990 := __e.Get(1)
_ = W4990
tmp15746 := MakeNative(func(__e *ControlFlow) {
W4991 := __e.Get(1)
_ = W4991
tmp15764 := PrimIsPair(W4990)

if True == tmp15764 {
tmp15747 := MakeNative(func(__e *ControlFlow) {
W4994 := __e.Get(1)
_ = W4994
tmp15748 := MakeNative(func(__e *ControlFlow) {
W4995 := __e.Get(1)
_ = W4995
tmp15749 := Call(__e, W4991, W4994)


__e.TailApply(tmp15749, W4995)
return


}, 1)

tmp15750 := PrimTail(W4990)

__e.TailApply(tmp15748, tmp15750)
return


}, 1)

tmp15751 := PrimHead(W4990)

__e.TailApply(tmp15747, tmp15751)
return


} else {
tmp15762 := Call(__e, PrimFunc(symshen_4pvar_2), W4990)


if True == tmp15762 {
tmp15752 := MakeNative(func(__e *ControlFlow) {
W4996 := __e.Get(1)
_ = W4996
tmp15753 := MakeNative(func(__e *ControlFlow) {
W4997 := __e.Get(1)
_ = W4997
tmp15754 := PrimCons(W4996, W4997)

tmp15755 := MakeNative(func(__e *ControlFlow) {
tmp15756 := Call(__e, W4991, W4996)


__e.TailApply(tmp15756, W4997)
return


}, 0)

tmp15757 := Call(__e, PrimFunc(symshen_4bind_b), W4990, tmp15754, V4900, tmp15755)


__e.TailApply(PrimFunc(symshen_4gc), V4900, tmp15757)
return


}, 1)

tmp15758 := Call(__e, PrimFunc(symshen_4newpv), V4900)


tmp15759 := Call(__e, tmp15753, tmp15758)


__e.TailApply(PrimFunc(symshen_4gc), V4900, tmp15759)
return


}, 1)

tmp15760 := Call(__e, PrimFunc(symshen_4newpv), V4900)


__e.TailApply(tmp15752, tmp15760)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15765 := MakeNative(func(__e *ControlFlow) {
Z4992 := __e.Get(1)
_ = Z4992
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4993 := __e.Get(1)
_ = Z4993
tmp15766 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15766

tmp15767 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4l_1rules), W4989, Z4993, V4899, V4900, V4901, W4904, V4903)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z4992, W4988, V4900, V4901, W4904, tmp15767)
return


}, 1))
return
}, 1)

__e.TailApply(tmp15746, tmp15765)
return


}, 1)

tmp15768 := Call(__e, PrimFunc(symshen_4lazyderef), V4898, V4900)


__e.TailApply(tmp15745, tmp15768)
return


}, 1)

tmp15769 := PrimTail(W4987)

__e.TailApply(tmp15744, tmp15769)
return


}, 1)

tmp15770 := PrimHead(W4987)

__e.TailApply(tmp15743, tmp15770)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15773 := Call(__e, PrimFunc(symshen_4lazyderef), V4897, V4900)


tmp15774 := Call(__e, tmp15742, tmp15773)


ifres15741 = tmp15774


} else {
ifres15741 = False


}

__e.TailApply(tmp15738, ifres15741)
return


} else {
__e.Return(W4966)
return
}


}, 1)

tmp15875 := Call(__e, PrimFunc(symshen_4unlocked_2), V4901)


var ifres15778 Obj

if True == tmp15875 {
tmp15779 := MakeNative(func(__e *ControlFlow) {
W4967 := __e.Get(1)
_ = W4967
tmp15872 := PrimIsPair(W4967)

if True == tmp15872 {
tmp15780 := MakeNative(func(__e *ControlFlow) {
W4968 := __e.Get(1)
_ = W4968
tmp15868 := PrimIsPair(W4968)

if True == tmp15868 {
tmp15781 := MakeNative(func(__e *ControlFlow) {
W4969 := __e.Get(1)
_ = W4969
tmp15864 := PrimIsPair(W4969)

if True == tmp15864 {
tmp15782 := MakeNative(func(__e *ControlFlow) {
W4970 := __e.Get(1)
_ = W4970
tmp15860 := PrimEqual(W4970, sym_8v)

if True == tmp15860 {
tmp15783 := MakeNative(func(__e *ControlFlow) {
W4971 := __e.Get(1)
_ = W4971
tmp15856 := PrimIsPair(W4971)

if True == tmp15856 {
tmp15784 := MakeNative(func(__e *ControlFlow) {
W4972 := __e.Get(1)
_ = W4972
tmp15785 := MakeNative(func(__e *ControlFlow) {
W4973 := __e.Get(1)
_ = W4973
tmp15851 := PrimIsPair(W4973)

if True == tmp15851 {
tmp15786 := MakeNative(func(__e *ControlFlow) {
W4974 := __e.Get(1)
_ = W4974
tmp15787 := MakeNative(func(__e *ControlFlow) {
W4975 := __e.Get(1)
_ = W4975
tmp15846 := PrimEqual(W4975, Nil)

if True == tmp15846 {
tmp15788 := MakeNative(func(__e *ControlFlow) {
W4976 := __e.Get(1)
_ = W4976
tmp15842 := PrimIsPair(W4976)

if True == tmp15842 {
tmp15789 := MakeNative(func(__e *ControlFlow) {
W4977 := __e.Get(1)
_ = W4977
tmp15790 := MakeNative(func(__e *ControlFlow) {
W4978 := __e.Get(1)
_ = W4978
tmp15837 := PrimIsPair(W4978)

if True == tmp15837 {
tmp15791 := MakeNative(func(__e *ControlFlow) {
W4979 := __e.Get(1)
_ = W4979
tmp15833 := PrimIsPair(W4979)

if True == tmp15833 {
tmp15792 := MakeNative(func(__e *ControlFlow) {
W4980 := __e.Get(1)
_ = W4980
tmp15829 := PrimEqual(W4980, symvector)

if True == tmp15829 {
tmp15793 := MakeNative(func(__e *ControlFlow) {
W4981 := __e.Get(1)
_ = W4981
tmp15825 := PrimIsPair(W4981)

if True == tmp15825 {
tmp15794 := MakeNative(func(__e *ControlFlow) {
W4982 := __e.Get(1)
_ = W4982
tmp15795 := MakeNative(func(__e *ControlFlow) {
W4983 := __e.Get(1)
_ = W4983
tmp15820 := PrimEqual(W4983, Nil)

if True == tmp15820 {
tmp15796 := MakeNative(func(__e *ControlFlow) {
W4984 := __e.Get(1)
_ = W4984
tmp15816 := PrimEqual(W4984, Nil)

if True == tmp15816 {
tmp15797 := MakeNative(func(__e *ControlFlow) {
W4985 := __e.Get(1)
_ = W4985
tmp15798 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15798

tmp15799 := Call(__e, PrimFunc(symshen_4deref), W4977, V4900)


tmp15800 := PrimIntern(MakeString(":"))

tmp15801 := PrimEqual(tmp15799, tmp15800)

tmp15802 := MakeNative(func(__e *ControlFlow) {
tmp15803 := MakeNative(func(__e *ControlFlow) {
tmp15804 := PrimCons(W4982, Nil)

tmp15805 := PrimCons(W4977, tmp15804)

tmp15806 := PrimCons(W4972, tmp15805)

tmp15807 := PrimCons(W4982, Nil)

tmp15808 := PrimCons(symvector, tmp15807)

tmp15809 := PrimCons(tmp15808, Nil)

tmp15810 := PrimCons(W4977, tmp15809)

tmp15811 := PrimCons(W4974, tmp15810)

tmp15812 := PrimCons(tmp15811, W4985)

tmp15813 := PrimCons(tmp15806, tmp15812)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp15813, V4898, True, V4900, V4901, W4904, V4903)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4900, V4901, W4904, tmp15803)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15801, V4900, V4901, W4904, tmp15802)
return


}, 1)

tmp15814 := PrimTail(W4967)

__e.TailApply(tmp15797, tmp15814)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15817 := PrimTail(W4978)

tmp15818 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15817, V4900)


__e.TailApply(tmp15796, tmp15818)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15821 := PrimTail(W4981)

tmp15822 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15821, V4900)


__e.TailApply(tmp15795, tmp15822)
return


}, 1)

tmp15823 := PrimHead(W4981)

__e.TailApply(tmp15794, tmp15823)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15826 := PrimTail(W4979)

tmp15827 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15826, V4900)


__e.TailApply(tmp15793, tmp15827)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15830 := PrimHead(W4979)

tmp15831 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15830, V4900)


__e.TailApply(tmp15792, tmp15831)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15834 := PrimHead(W4978)

tmp15835 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15834, V4900)


__e.TailApply(tmp15791, tmp15835)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15838 := PrimTail(W4976)

tmp15839 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15838, V4900)


__e.TailApply(tmp15790, tmp15839)
return


}, 1)

tmp15840 := PrimHead(W4976)

__e.TailApply(tmp15789, tmp15840)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15843 := PrimTail(W4968)

tmp15844 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15843, V4900)


__e.TailApply(tmp15788, tmp15844)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15847 := PrimTail(W4973)

tmp15848 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15847, V4900)


__e.TailApply(tmp15787, tmp15848)
return


}, 1)

tmp15849 := PrimHead(W4973)

__e.TailApply(tmp15786, tmp15849)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15852 := PrimTail(W4971)

tmp15853 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15852, V4900)


__e.TailApply(tmp15785, tmp15853)
return


}, 1)

tmp15854 := PrimHead(W4971)

__e.TailApply(tmp15784, tmp15854)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15857 := PrimTail(W4969)

tmp15858 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15857, V4900)


__e.TailApply(tmp15783, tmp15858)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15861 := PrimHead(W4969)

tmp15862 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15861, V4900)


__e.TailApply(tmp15782, tmp15862)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15865 := PrimHead(W4968)

tmp15866 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15865, V4900)


__e.TailApply(tmp15781, tmp15866)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15869 := PrimHead(W4967)

tmp15870 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15869, V4900)


__e.TailApply(tmp15780, tmp15870)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15873 := Call(__e, PrimFunc(symshen_4lazyderef), V4897, V4900)


tmp15874 := Call(__e, tmp15779, tmp15873)


ifres15778 = tmp15874


} else {
ifres15778 = False


}

__e.TailApply(tmp15737, ifres15778)
return


} else {
__e.Return(W4950)
return
}


}, 1)

tmp15956 := Call(__e, PrimFunc(symshen_4unlocked_2), V4901)


var ifres15878 Obj

if True == tmp15956 {
tmp15879 := MakeNative(func(__e *ControlFlow) {
W4951 := __e.Get(1)
_ = W4951
tmp15953 := PrimIsPair(W4951)

if True == tmp15953 {
tmp15880 := MakeNative(func(__e *ControlFlow) {
W4952 := __e.Get(1)
_ = W4952
tmp15949 := PrimIsPair(W4952)

if True == tmp15949 {
tmp15881 := MakeNative(func(__e *ControlFlow) {
W4953 := __e.Get(1)
_ = W4953
tmp15945 := PrimIsPair(W4953)

if True == tmp15945 {
tmp15882 := MakeNative(func(__e *ControlFlow) {
W4954 := __e.Get(1)
_ = W4954
tmp15941 := PrimEqual(W4954, sym_8s)

if True == tmp15941 {
tmp15883 := MakeNative(func(__e *ControlFlow) {
W4955 := __e.Get(1)
_ = W4955
tmp15937 := PrimIsPair(W4955)

if True == tmp15937 {
tmp15884 := MakeNative(func(__e *ControlFlow) {
W4956 := __e.Get(1)
_ = W4956
tmp15885 := MakeNative(func(__e *ControlFlow) {
W4957 := __e.Get(1)
_ = W4957
tmp15932 := PrimIsPair(W4957)

if True == tmp15932 {
tmp15886 := MakeNative(func(__e *ControlFlow) {
W4958 := __e.Get(1)
_ = W4958
tmp15887 := MakeNative(func(__e *ControlFlow) {
W4959 := __e.Get(1)
_ = W4959
tmp15927 := PrimEqual(W4959, Nil)

if True == tmp15927 {
tmp15888 := MakeNative(func(__e *ControlFlow) {
W4960 := __e.Get(1)
_ = W4960
tmp15923 := PrimIsPair(W4960)

if True == tmp15923 {
tmp15889 := MakeNative(func(__e *ControlFlow) {
W4961 := __e.Get(1)
_ = W4961
tmp15890 := MakeNative(func(__e *ControlFlow) {
W4962 := __e.Get(1)
_ = W4962
tmp15918 := PrimIsPair(W4962)

if True == tmp15918 {
tmp15891 := MakeNative(func(__e *ControlFlow) {
W4963 := __e.Get(1)
_ = W4963
tmp15914 := PrimEqual(W4963, symstring)

if True == tmp15914 {
tmp15892 := MakeNative(func(__e *ControlFlow) {
W4964 := __e.Get(1)
_ = W4964
tmp15910 := PrimEqual(W4964, Nil)

if True == tmp15910 {
tmp15893 := MakeNative(func(__e *ControlFlow) {
W4965 := __e.Get(1)
_ = W4965
tmp15894 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15894

tmp15895 := Call(__e, PrimFunc(symshen_4deref), W4961, V4900)


tmp15896 := PrimIntern(MakeString(":"))

tmp15897 := PrimEqual(tmp15895, tmp15896)

tmp15898 := MakeNative(func(__e *ControlFlow) {
tmp15899 := MakeNative(func(__e *ControlFlow) {
tmp15900 := PrimCons(symstring, Nil)

tmp15901 := PrimCons(W4961, tmp15900)

tmp15902 := PrimCons(W4956, tmp15901)

tmp15903 := PrimCons(symstring, Nil)

tmp15904 := PrimCons(W4961, tmp15903)

tmp15905 := PrimCons(W4958, tmp15904)

tmp15906 := PrimCons(tmp15905, W4965)

tmp15907 := PrimCons(tmp15902, tmp15906)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp15907, V4898, True, V4900, V4901, W4904, V4903)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4900, V4901, W4904, tmp15899)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15897, V4900, V4901, W4904, tmp15898)
return


}, 1)

tmp15908 := PrimTail(W4951)

__e.TailApply(tmp15893, tmp15908)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15911 := PrimTail(W4962)

tmp15912 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15911, V4900)


__e.TailApply(tmp15892, tmp15912)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15915 := PrimHead(W4962)

tmp15916 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15915, V4900)


__e.TailApply(tmp15891, tmp15916)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15919 := PrimTail(W4960)

tmp15920 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15919, V4900)


__e.TailApply(tmp15890, tmp15920)
return


}, 1)

tmp15921 := PrimHead(W4960)

__e.TailApply(tmp15889, tmp15921)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15924 := PrimTail(W4952)

tmp15925 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15924, V4900)


__e.TailApply(tmp15888, tmp15925)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15928 := PrimTail(W4957)

tmp15929 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15928, V4900)


__e.TailApply(tmp15887, tmp15929)
return


}, 1)

tmp15930 := PrimHead(W4957)

__e.TailApply(tmp15886, tmp15930)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15933 := PrimTail(W4955)

tmp15934 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15933, V4900)


__e.TailApply(tmp15885, tmp15934)
return


}, 1)

tmp15935 := PrimHead(W4955)

__e.TailApply(tmp15884, tmp15935)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15938 := PrimTail(W4953)

tmp15939 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15938, V4900)


__e.TailApply(tmp15883, tmp15939)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15942 := PrimHead(W4953)

tmp15943 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15942, V4900)


__e.TailApply(tmp15882, tmp15943)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15946 := PrimHead(W4952)

tmp15947 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15946, V4900)


__e.TailApply(tmp15881, tmp15947)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15950 := PrimHead(W4951)

tmp15951 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15950, V4900)


__e.TailApply(tmp15880, tmp15951)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15954 := Call(__e, PrimFunc(symshen_4lazyderef), V4897, V4900)


tmp15955 := Call(__e, tmp15879, tmp15954)


ifres15878 = tmp15955


} else {
ifres15878 = False


}

__e.TailApply(tmp15736, ifres15878)
return


} else {
__e.Return(W4928)
return
}


}, 1)

tmp16061 := Call(__e, PrimFunc(symshen_4unlocked_2), V4901)


var ifres15959 Obj

if True == tmp16061 {
tmp15960 := MakeNative(func(__e *ControlFlow) {
W4929 := __e.Get(1)
_ = W4929
tmp16058 := PrimIsPair(W4929)

if True == tmp16058 {
tmp15961 := MakeNative(func(__e *ControlFlow) {
W4930 := __e.Get(1)
_ = W4930
tmp16054 := PrimIsPair(W4930)

if True == tmp16054 {
tmp15962 := MakeNative(func(__e *ControlFlow) {
W4931 := __e.Get(1)
_ = W4931
tmp16050 := PrimIsPair(W4931)

if True == tmp16050 {
tmp15963 := MakeNative(func(__e *ControlFlow) {
W4932 := __e.Get(1)
_ = W4932
tmp16046 := PrimEqual(W4932, sym_8p)

if True == tmp16046 {
tmp15964 := MakeNative(func(__e *ControlFlow) {
W4933 := __e.Get(1)
_ = W4933
tmp16042 := PrimIsPair(W4933)

if True == tmp16042 {
tmp15965 := MakeNative(func(__e *ControlFlow) {
W4934 := __e.Get(1)
_ = W4934
tmp15966 := MakeNative(func(__e *ControlFlow) {
W4935 := __e.Get(1)
_ = W4935
tmp16037 := PrimIsPair(W4935)

if True == tmp16037 {
tmp15967 := MakeNative(func(__e *ControlFlow) {
W4936 := __e.Get(1)
_ = W4936
tmp15968 := MakeNative(func(__e *ControlFlow) {
W4937 := __e.Get(1)
_ = W4937
tmp16032 := PrimEqual(W4937, Nil)

if True == tmp16032 {
tmp15969 := MakeNative(func(__e *ControlFlow) {
W4938 := __e.Get(1)
_ = W4938
tmp16028 := PrimIsPair(W4938)

if True == tmp16028 {
tmp15970 := MakeNative(func(__e *ControlFlow) {
W4939 := __e.Get(1)
_ = W4939
tmp15971 := MakeNative(func(__e *ControlFlow) {
W4940 := __e.Get(1)
_ = W4940
tmp16023 := PrimIsPair(W4940)

if True == tmp16023 {
tmp15972 := MakeNative(func(__e *ControlFlow) {
W4941 := __e.Get(1)
_ = W4941
tmp16019 := PrimIsPair(W4941)

if True == tmp16019 {
tmp15973 := MakeNative(func(__e *ControlFlow) {
W4942 := __e.Get(1)
_ = W4942
tmp15974 := MakeNative(func(__e *ControlFlow) {
W4943 := __e.Get(1)
_ = W4943
tmp16014 := PrimIsPair(W4943)

if True == tmp16014 {
tmp15975 := MakeNative(func(__e *ControlFlow) {
W4944 := __e.Get(1)
_ = W4944
tmp16010 := PrimEqual(W4944, sym_d)

if True == tmp16010 {
tmp15976 := MakeNative(func(__e *ControlFlow) {
W4945 := __e.Get(1)
_ = W4945
tmp16006 := PrimIsPair(W4945)

if True == tmp16006 {
tmp15977 := MakeNative(func(__e *ControlFlow) {
W4946 := __e.Get(1)
_ = W4946
tmp15978 := MakeNative(func(__e *ControlFlow) {
W4947 := __e.Get(1)
_ = W4947
tmp16001 := PrimEqual(W4947, Nil)

if True == tmp16001 {
tmp15979 := MakeNative(func(__e *ControlFlow) {
W4948 := __e.Get(1)
_ = W4948
tmp15997 := PrimEqual(W4948, Nil)

if True == tmp15997 {
tmp15980 := MakeNative(func(__e *ControlFlow) {
W4949 := __e.Get(1)
_ = W4949
tmp15981 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15981

tmp15982 := Call(__e, PrimFunc(symshen_4deref), W4939, V4900)


tmp15983 := PrimIntern(MakeString(":"))

tmp15984 := PrimEqual(tmp15982, tmp15983)

tmp15985 := MakeNative(func(__e *ControlFlow) {
tmp15986 := MakeNative(func(__e *ControlFlow) {
tmp15987 := PrimCons(W4942, Nil)

tmp15988 := PrimCons(W4939, tmp15987)

tmp15989 := PrimCons(W4934, tmp15988)

tmp15990 := PrimCons(W4946, Nil)

tmp15991 := PrimCons(W4939, tmp15990)

tmp15992 := PrimCons(W4936, tmp15991)

tmp15993 := PrimCons(tmp15992, W4949)

tmp15994 := PrimCons(tmp15989, tmp15993)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp15994, V4898, True, V4900, V4901, W4904, V4903)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4900, V4901, W4904, tmp15986)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15984, V4900, V4901, W4904, tmp15985)
return


}, 1)

tmp15995 := PrimTail(W4929)

__e.TailApply(tmp15980, tmp15995)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15998 := PrimTail(W4940)

tmp15999 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15998, V4900)


__e.TailApply(tmp15979, tmp15999)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16002 := PrimTail(W4945)

tmp16003 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16002, V4900)


__e.TailApply(tmp15978, tmp16003)
return


}, 1)

tmp16004 := PrimHead(W4945)

__e.TailApply(tmp15977, tmp16004)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16007 := PrimTail(W4943)

tmp16008 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16007, V4900)


__e.TailApply(tmp15976, tmp16008)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16011 := PrimHead(W4943)

tmp16012 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16011, V4900)


__e.TailApply(tmp15975, tmp16012)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16015 := PrimTail(W4941)

tmp16016 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16015, V4900)


__e.TailApply(tmp15974, tmp16016)
return


}, 1)

tmp16017 := PrimHead(W4941)

__e.TailApply(tmp15973, tmp16017)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16020 := PrimHead(W4940)

tmp16021 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16020, V4900)


__e.TailApply(tmp15972, tmp16021)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16024 := PrimTail(W4938)

tmp16025 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16024, V4900)


__e.TailApply(tmp15971, tmp16025)
return


}, 1)

tmp16026 := PrimHead(W4938)

__e.TailApply(tmp15970, tmp16026)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16029 := PrimTail(W4930)

tmp16030 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16029, V4900)


__e.TailApply(tmp15969, tmp16030)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16033 := PrimTail(W4935)

tmp16034 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16033, V4900)


__e.TailApply(tmp15968, tmp16034)
return


}, 1)

tmp16035 := PrimHead(W4935)

__e.TailApply(tmp15967, tmp16035)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16038 := PrimTail(W4933)

tmp16039 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16038, V4900)


__e.TailApply(tmp15966, tmp16039)
return


}, 1)

tmp16040 := PrimHead(W4933)

__e.TailApply(tmp15965, tmp16040)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16043 := PrimTail(W4931)

tmp16044 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16043, V4900)


__e.TailApply(tmp15964, tmp16044)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16047 := PrimHead(W4931)

tmp16048 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16047, V4900)


__e.TailApply(tmp15963, tmp16048)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16051 := PrimHead(W4930)

tmp16052 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16051, V4900)


__e.TailApply(tmp15962, tmp16052)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16055 := PrimHead(W4929)

tmp16056 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16055, V4900)


__e.TailApply(tmp15961, tmp16056)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16059 := Call(__e, PrimFunc(symshen_4lazyderef), V4897, V4900)


tmp16060 := Call(__e, tmp15960, tmp16059)


ifres15959 = tmp16060


} else {
ifres15959 = False


}

__e.TailApply(tmp15735, ifres15959)
return


} else {
__e.Return(W4908)
return
}


}, 1)

tmp16161 := Call(__e, PrimFunc(symshen_4unlocked_2), V4901)


var ifres16064 Obj

if True == tmp16161 {
tmp16065 := MakeNative(func(__e *ControlFlow) {
W4909 := __e.Get(1)
_ = W4909
tmp16158 := PrimIsPair(W4909)

if True == tmp16158 {
tmp16066 := MakeNative(func(__e *ControlFlow) {
W4910 := __e.Get(1)
_ = W4910
tmp16154 := PrimIsPair(W4910)

if True == tmp16154 {
tmp16067 := MakeNative(func(__e *ControlFlow) {
W4911 := __e.Get(1)
_ = W4911
tmp16150 := PrimIsPair(W4911)

if True == tmp16150 {
tmp16068 := MakeNative(func(__e *ControlFlow) {
W4912 := __e.Get(1)
_ = W4912
tmp16146 := PrimEqual(W4912, symcons)

if True == tmp16146 {
tmp16069 := MakeNative(func(__e *ControlFlow) {
W4913 := __e.Get(1)
_ = W4913
tmp16142 := PrimIsPair(W4913)

if True == tmp16142 {
tmp16070 := MakeNative(func(__e *ControlFlow) {
W4914 := __e.Get(1)
_ = W4914
tmp16071 := MakeNative(func(__e *ControlFlow) {
W4915 := __e.Get(1)
_ = W4915
tmp16137 := PrimIsPair(W4915)

if True == tmp16137 {
tmp16072 := MakeNative(func(__e *ControlFlow) {
W4916 := __e.Get(1)
_ = W4916
tmp16073 := MakeNative(func(__e *ControlFlow) {
W4917 := __e.Get(1)
_ = W4917
tmp16132 := PrimEqual(W4917, Nil)

if True == tmp16132 {
tmp16074 := MakeNative(func(__e *ControlFlow) {
W4918 := __e.Get(1)
_ = W4918
tmp16128 := PrimIsPair(W4918)

if True == tmp16128 {
tmp16075 := MakeNative(func(__e *ControlFlow) {
W4919 := __e.Get(1)
_ = W4919
tmp16076 := MakeNative(func(__e *ControlFlow) {
W4920 := __e.Get(1)
_ = W4920
tmp16123 := PrimIsPair(W4920)

if True == tmp16123 {
tmp16077 := MakeNative(func(__e *ControlFlow) {
W4921 := __e.Get(1)
_ = W4921
tmp16119 := PrimIsPair(W4921)

if True == tmp16119 {
tmp16078 := MakeNative(func(__e *ControlFlow) {
W4922 := __e.Get(1)
_ = W4922
tmp16115 := PrimEqual(W4922, symlist)

if True == tmp16115 {
tmp16079 := MakeNative(func(__e *ControlFlow) {
W4923 := __e.Get(1)
_ = W4923
tmp16111 := PrimIsPair(W4923)

if True == tmp16111 {
tmp16080 := MakeNative(func(__e *ControlFlow) {
W4924 := __e.Get(1)
_ = W4924
tmp16081 := MakeNative(func(__e *ControlFlow) {
W4925 := __e.Get(1)
_ = W4925
tmp16106 := PrimEqual(W4925, Nil)

if True == tmp16106 {
tmp16082 := MakeNative(func(__e *ControlFlow) {
W4926 := __e.Get(1)
_ = W4926
tmp16102 := PrimEqual(W4926, Nil)

if True == tmp16102 {
tmp16083 := MakeNative(func(__e *ControlFlow) {
W4927 := __e.Get(1)
_ = W4927
tmp16084 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16084

tmp16085 := Call(__e, PrimFunc(symshen_4deref), W4919, V4900)


tmp16086 := PrimIntern(MakeString(":"))

tmp16087 := PrimEqual(tmp16085, tmp16086)

tmp16088 := MakeNative(func(__e *ControlFlow) {
tmp16089 := MakeNative(func(__e *ControlFlow) {
tmp16090 := PrimCons(W4924, Nil)

tmp16091 := PrimCons(W4919, tmp16090)

tmp16092 := PrimCons(W4914, tmp16091)

tmp16093 := PrimCons(W4924, Nil)

tmp16094 := PrimCons(symlist, tmp16093)

tmp16095 := PrimCons(tmp16094, Nil)

tmp16096 := PrimCons(W4919, tmp16095)

tmp16097 := PrimCons(W4916, tmp16096)

tmp16098 := PrimCons(tmp16097, W4927)

tmp16099 := PrimCons(tmp16092, tmp16098)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16099, V4898, True, V4900, V4901, W4904, V4903)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4900, V4901, W4904, tmp16089)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp16087, V4900, V4901, W4904, tmp16088)
return


}, 1)

tmp16100 := PrimTail(W4909)

__e.TailApply(tmp16083, tmp16100)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16103 := PrimTail(W4920)

tmp16104 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16103, V4900)


__e.TailApply(tmp16082, tmp16104)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16107 := PrimTail(W4923)

tmp16108 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16107, V4900)


__e.TailApply(tmp16081, tmp16108)
return


}, 1)

tmp16109 := PrimHead(W4923)

__e.TailApply(tmp16080, tmp16109)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16112 := PrimTail(W4921)

tmp16113 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16112, V4900)


__e.TailApply(tmp16079, tmp16113)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16116 := PrimHead(W4921)

tmp16117 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16116, V4900)


__e.TailApply(tmp16078, tmp16117)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16120 := PrimHead(W4920)

tmp16121 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16120, V4900)


__e.TailApply(tmp16077, tmp16121)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16124 := PrimTail(W4918)

tmp16125 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16124, V4900)


__e.TailApply(tmp16076, tmp16125)
return


}, 1)

tmp16126 := PrimHead(W4918)

__e.TailApply(tmp16075, tmp16126)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16129 := PrimTail(W4910)

tmp16130 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16129, V4900)


__e.TailApply(tmp16074, tmp16130)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16133 := PrimTail(W4915)

tmp16134 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16133, V4900)


__e.TailApply(tmp16073, tmp16134)
return


}, 1)

tmp16135 := PrimHead(W4915)

__e.TailApply(tmp16072, tmp16135)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16138 := PrimTail(W4913)

tmp16139 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16138, V4900)


__e.TailApply(tmp16071, tmp16139)
return


}, 1)

tmp16140 := PrimHead(W4913)

__e.TailApply(tmp16070, tmp16140)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16143 := PrimTail(W4911)

tmp16144 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16143, V4900)


__e.TailApply(tmp16069, tmp16144)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16147 := PrimHead(W4911)

tmp16148 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16147, V4900)


__e.TailApply(tmp16068, tmp16148)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16151 := PrimHead(W4910)

tmp16152 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16151, V4900)


__e.TailApply(tmp16067, tmp16152)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16155 := PrimHead(W4909)

tmp16156 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16155, V4900)


__e.TailApply(tmp16066, tmp16156)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16159 := Call(__e, PrimFunc(symshen_4lazyderef), V4897, V4900)


tmp16160 := Call(__e, tmp16065, tmp16159)


ifres16064 = tmp16160


} else {
ifres16064 = False


}

__e.TailApply(tmp15734, ifres16064)
return


} else {
__e.Return(W4905)
return
}


}, 1)

tmp16176 := Call(__e, PrimFunc(symshen_4unlocked_2), V4901)


var ifres16164 Obj

if True == tmp16176 {
tmp16165 := MakeNative(func(__e *ControlFlow) {
W4906 := __e.Get(1)
_ = W4906
tmp16173 := PrimEqual(W4906, Nil)

if True == tmp16173 {
tmp16166 := MakeNative(func(__e *ControlFlow) {
W4907 := __e.Get(1)
_ = W4907
tmp16170 := PrimEqual(W4907, True)

if True == tmp16170 {
tmp16167 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16167

tmp16168 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symbind), V4898, Nil, V4900, V4901, W4904, V4903)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4900, V4901, W4904, tmp16168)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16171 := Call(__e, PrimFunc(symshen_4lazyderef), V4899, V4900)


__e.TailApply(tmp16166, tmp16171)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16174 := Call(__e, PrimFunc(symshen_4lazyderef), V4897, V4900)


tmp16175 := Call(__e, tmp16165, tmp16174)


ifres16164 = tmp16175


} else {
ifres16164 = False


}

__e.TailApply(tmp15733, ifres16164)
return


}, 1)

tmp16177 := PrimNumberAdd(V4902, MakeNumber(1))

__e.TailApply(tmp15732, tmp16177)
return


}, 7)

tmp16178 := Call(__e, ns2_1set, symshen_4l_1rules, tmp15731)


_ = tmp16178

tmp16179 := MakeNative(func(__e *ControlFlow) {
V4998 := __e.Get(1)
_ = V4998
V4999 := __e.Get(2)
_ = V4999
V5000 := __e.Get(3)
_ = V5000
V5001 := __e.Get(4)
_ = V5001
V5002 := __e.Get(5)
_ = V5002
V5003 := __e.Get(6)
_ = V5003
tmp16180 := MakeNative(func(__e *ControlFlow) {
W5004 := __e.Get(1)
_ = W5004
tmp16181 := MakeNative(func(__e *ControlFlow) {
W5005 := __e.Get(1)
_ = W5005
tmp16183 := PrimEqual(W5005, False)

if True == tmp16183 {
__e.TailApply(PrimFunc(symshen_4unlock), V5001, W5004)
return
} else {
__e.Return(W5005)
return
}


}, 1)

tmp16231 := Call(__e, PrimFunc(symshen_4unlocked_2), V5001)


var ifres16184 Obj

if True == tmp16231 {
tmp16185 := MakeNative(func(__e *ControlFlow) {
W5006 := __e.Get(1)
_ = W5006
tmp16228 := PrimIsPair(W5006)

if True == tmp16228 {
tmp16186 := MakeNative(func(__e *ControlFlow) {
W5007 := __e.Get(1)
_ = W5007
tmp16224 := PrimEqual(W5007, symdefine)

if True == tmp16224 {
tmp16187 := MakeNative(func(__e *ControlFlow) {
W5008 := __e.Get(1)
_ = W5008
tmp16220 := PrimIsPair(W5008)

if True == tmp16220 {
tmp16188 := MakeNative(func(__e *ControlFlow) {
W5009 := __e.Get(1)
_ = W5009
tmp16189 := MakeNative(func(__e *ControlFlow) {
W5010 := __e.Get(1)
_ = W5010
tmp16190 := MakeNative(func(__e *ControlFlow) {
W5011 := __e.Get(1)
_ = W5011
tmp16191 := MakeNative(func(__e *ControlFlow) {
W5012 := __e.Get(1)
_ = W5012
tmp16192 := MakeNative(func(__e *ControlFlow) {
W5013 := __e.Get(1)
_ = W5013
tmp16193 := MakeNative(func(__e *ControlFlow) {
W5014 := __e.Get(1)
_ = W5014
tmp16194 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16194

tmp16195 := MakeNative(func(__e *ControlFlow) {
tmp16196 := PrimCons(W5009, W5010)

tmp16197 := Call(__e, PrimFunc(symshen_4sigxrules), tmp16196)


tmp16198 := MakeNative(func(__e *ControlFlow) {
tmp16199 := Call(__e, PrimFunc(symshen_4lazyderef), W5011, V5000)


tmp16200 := Call(__e, PrimFunc(symfst), tmp16199)


tmp16201 := MakeNative(func(__e *ControlFlow) {
tmp16202 := Call(__e, PrimFunc(symshen_4lazyderef), W5011, V5000)


tmp16203 := Call(__e, PrimFunc(symsnd), tmp16202)


tmp16204 := MakeNative(func(__e *ControlFlow) {
tmp16205 := Call(__e, PrimFunc(symshen_4deref), W5014, V5000)


tmp16206 := Call(__e, PrimFunc(symshen_4freshen_1sig), tmp16205)


tmp16207 := MakeNative(func(__e *ControlFlow) {
tmp16208 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis), W5014, V4999, V5000, V5001, W5004, V5003)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rules), W5009, W5012, W5013, MakeNumber(1), V5000, V5001, W5004, tmp16208)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5013, tmp16206, V5000, V5001, W5004, tmp16207)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5012, tmp16203, V5000, V5001, W5004, tmp16204)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5014, tmp16200, V5000, V5001, W5004, tmp16201)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5011, tmp16197, V5000, V5001, W5004, tmp16198)
return


}, 0)

tmp16209 := Call(__e, PrimFunc(symshen_4cut), V5000, V5001, W5004, tmp16195)


__e.TailApply(PrimFunc(symshen_4gc), V5000, tmp16209)
return


}, 1)

tmp16210 := Call(__e, PrimFunc(symshen_4newpv), V5000)


tmp16211 := Call(__e, tmp16193, tmp16210)


__e.TailApply(PrimFunc(symshen_4gc), V5000, tmp16211)
return


}, 1)

tmp16212 := Call(__e, PrimFunc(symshen_4newpv), V5000)


tmp16213 := Call(__e, tmp16192, tmp16212)


__e.TailApply(PrimFunc(symshen_4gc), V5000, tmp16213)
return


}, 1)

tmp16214 := Call(__e, PrimFunc(symshen_4newpv), V5000)


tmp16215 := Call(__e, tmp16191, tmp16214)


__e.TailApply(PrimFunc(symshen_4gc), V5000, tmp16215)
return


}, 1)

tmp16216 := Call(__e, PrimFunc(symshen_4newpv), V5000)


__e.TailApply(tmp16190, tmp16216)
return


}, 1)

tmp16217 := PrimTail(W5008)

__e.TailApply(tmp16189, tmp16217)
return


}, 1)

tmp16218 := PrimHead(W5008)

__e.TailApply(tmp16188, tmp16218)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16221 := PrimTail(W5006)

tmp16222 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16221, V5000)


__e.TailApply(tmp16187, tmp16222)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16225 := PrimHead(W5006)

tmp16226 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16225, V5000)


__e.TailApply(tmp16186, tmp16226)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16229 := Call(__e, PrimFunc(symshen_4lazyderef), V4998, V5000)


tmp16230 := Call(__e, tmp16185, tmp16229)


ifres16184 = tmp16230


} else {
ifres16184 = False


}

__e.TailApply(tmp16181, ifres16184)
return


}, 1)

tmp16232 := PrimNumberAdd(V5002, MakeNumber(1))

__e.TailApply(tmp16180, tmp16232)
return


}, 6)

tmp16233 := Call(__e, ns2_1set, symshen_4t_d, tmp16179)


_ = tmp16233

tmp16234 := MakeNative(func(__e *ControlFlow) {
V5015 := __e.Get(1)
_ = V5015
tmp16235 := MakeNative(func(__e *ControlFlow) {
Z5016 := __e.Get(1)
_ = Z5016
__e.TailApply(PrimFunc(symshen_4_5sig_drules_6), Z5016)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp16235, V5015)
return


}, 1)

tmp16236 := Call(__e, ns2_1set, symshen_4sigxrules, tmp16234)


_ = tmp16236

tmp16237 := MakeNative(func(__e *ControlFlow) {
V5017 := __e.Get(1)
_ = V5017
tmp16238 := MakeNative(func(__e *ControlFlow) {
W5018 := __e.Get(1)
_ = W5018
tmp16240 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5018)


if True == tmp16240 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W5018)
return
}


}, 1)

tmp16273 := PrimIsPair(V5017)

var ifres16241 Obj

if True == tmp16273 {
tmp16242 := MakeNative(func(__e *ControlFlow) {
W5019 := __e.Get(1)
_ = W5019
tmp16269 := Call(__e, PrimFunc(symshen_4hds_a_2), W5019, sym_i)


if True == tmp16269 {
tmp16243 := MakeNative(func(__e *ControlFlow) {
W5020 := __e.Get(1)
_ = W5020
tmp16244 := MakeNative(func(__e *ControlFlow) {
W5021 := __e.Get(1)
_ = W5021
tmp16265 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5021)


if True == tmp16265 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16245 := MakeNative(func(__e *ControlFlow) {
W5022 := __e.Get(1)
_ = W5022
tmp16246 := MakeNative(func(__e *ControlFlow) {
W5023 := __e.Get(1)
_ = W5023
tmp16261 := Call(__e, PrimFunc(symshen_4hds_a_2), W5023, sym_j)


if True == tmp16261 {
tmp16247 := MakeNative(func(__e *ControlFlow) {
W5024 := __e.Get(1)
_ = W5024
tmp16248 := MakeNative(func(__e *ControlFlow) {
W5025 := __e.Get(1)
_ = W5025
tmp16257 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5025)


if True == tmp16257 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16249 := MakeNative(func(__e *ControlFlow) {
W5026 := __e.Get(1)
_ = W5026
tmp16250 := MakeNative(func(__e *ControlFlow) {
W5027 := __e.Get(1)
_ = W5027
tmp16251 := MakeNative(func(__e *ControlFlow) {
W5028 := __e.Get(1)
_ = W5028
__e.TailApply(PrimFunc(sym_8p), W5028, W5026)
return
}, 1)

tmp16252 := Call(__e, PrimFunc(symshen_4rectify_1type), W5022)


tmp16253 := Call(__e, tmp16251, tmp16252)


__e.TailApply(PrimFunc(symshen_4comb), W5027, tmp16253)
return


}, 1)

tmp16254 := Call(__e, PrimFunc(symshen_4in_1_6), W5025)


__e.TailApply(tmp16250, tmp16254)
return


}, 1)

tmp16255 := Call(__e, PrimFunc(symshen_4_5_1out), W5025)


__e.TailApply(tmp16249, tmp16255)
return


}


}, 1)

tmp16258 := Call(__e, PrimFunc(symshen_4_5rules_d_6), W5024)


__e.TailApply(tmp16248, tmp16258)
return


}, 1)

tmp16259 := Call(__e, PrimFunc(symtail), W5023)


__e.TailApply(tmp16247, tmp16259)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16262 := Call(__e, PrimFunc(symshen_4in_1_6), W5021)


__e.TailApply(tmp16246, tmp16262)
return


}, 1)

tmp16263 := Call(__e, PrimFunc(symshen_4_5_1out), W5021)


__e.TailApply(tmp16245, tmp16263)
return


}


}, 1)

tmp16266 := Call(__e, PrimFunc(symshen_4_5signature_6), W5020)


__e.TailApply(tmp16244, tmp16266)
return


}, 1)

tmp16267 := Call(__e, PrimFunc(symtail), W5019)


__e.TailApply(tmp16243, tmp16267)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16270 := Call(__e, PrimFunc(symtail), V5017)


tmp16271 := Call(__e, tmp16242, tmp16270)


ifres16241 = tmp16271


} else {
tmp16272 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16241 = tmp16272


}

__e.TailApply(tmp16238, ifres16241)
return


}, 1)

tmp16274 := Call(__e, ns2_1set, symshen_4_5sig_drules_6, tmp16237)


_ = tmp16274

tmp16275 := MakeNative(func(__e *ControlFlow) {
V5029 := __e.Get(1)
_ = V5029
tmp16276 := MakeNative(func(__e *ControlFlow) {
W5030 := __e.Get(1)
_ = W5030
tmp16277 := MakeNative(func(__e *ControlFlow) {
W5031 := __e.Get(1)
_ = W5031
__e.TailApply(PrimFunc(symshen_4freshen_1type), W5031, V5029)
return
}, 1)

tmp16278 := MakeNative(func(__e *ControlFlow) {
Z5032 := __e.Get(1)
_ = Z5032
tmp16279 := Call(__e, PrimFunc(symconcat), sym_e, Z5032)


tmp16280 := Call(__e, PrimFunc(symshen_4freshterm), tmp16279)


__e.Return(PrimCons(Z5032, tmp16280))
return


}, 1)

tmp16281 := Call(__e, PrimFunc(symmap), tmp16278, W5030)


__e.TailApply(tmp16277, tmp16281)
return


}, 1)

tmp16282 := Call(__e, PrimFunc(symshen_4extract_1vars), V5029)


__e.TailApply(tmp16276, tmp16282)
return


}, 1)

tmp16283 := Call(__e, ns2_1set, symshen_4freshen_1sig, tmp16275)


_ = tmp16283

tmp16284 := MakeNative(func(__e *ControlFlow) {
V5033 := __e.Get(1)
_ = V5033
V5034 := __e.Get(2)
_ = V5034
tmp16298 := PrimEqual(Nil, V5033)

if True == tmp16298 {
__e.Return(V5034)
return
} else {
tmp16296 := PrimIsPair(V5033)

var ifres16292 Obj

if True == tmp16296 {
tmp16294 := PrimHead(V5033)

tmp16295 := PrimIsPair(tmp16294)

var ifres16293 Obj

if True == tmp16295 {
ifres16293 = True


} else {
ifres16293 = False


}

ifres16292 = ifres16293


} else {
ifres16292 = False


}

if True == ifres16292 {
tmp16285 := PrimTail(V5033)

tmp16286 := PrimHead(V5033)

tmp16287 := PrimTail(tmp16286)

tmp16288 := PrimHead(V5033)

tmp16289 := PrimHead(tmp16288)

tmp16290 := Call(__e, PrimFunc(symsubst), tmp16287, tmp16289, V5034)


__e.TailApply(PrimFunc(symshen_4freshen_1type), tmp16285, tmp16290)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.freshen-type")))
return
}


}


}, 2)

tmp16299 := Call(__e, ns2_1set, symshen_4freshen_1type, tmp16284)


_ = tmp16299

tmp16300 := MakeNative(func(__e *ControlFlow) {
V5035 := __e.Get(1)
_ = V5035
tmp16301 := MakeNative(func(__e *ControlFlow) {
W5036 := __e.Get(1)
_ = W5036
tmp16316 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5036)


if True == tmp16316 {
tmp16302 := MakeNative(func(__e *ControlFlow) {
W5043 := __e.Get(1)
_ = W5043
tmp16304 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5043)


if True == tmp16304 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W5043)
return
}


}, 1)

tmp16305 := MakeNative(func(__e *ControlFlow) {
W5044 := __e.Get(1)
_ = W5044
tmp16312 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5044)


if True == tmp16312 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16306 := MakeNative(func(__e *ControlFlow) {
W5045 := __e.Get(1)
_ = W5045
tmp16307 := MakeNative(func(__e *ControlFlow) {
W5046 := __e.Get(1)
_ = W5046
tmp16308 := PrimCons(W5045, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W5046, tmp16308)
return


}, 1)

tmp16309 := Call(__e, PrimFunc(symshen_4in_1_6), W5044)


__e.TailApply(tmp16307, tmp16309)
return


}, 1)

tmp16310 := Call(__e, PrimFunc(symshen_4_5_1out), W5044)


__e.TailApply(tmp16306, tmp16310)
return


}


}, 1)

tmp16313 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5035)


tmp16314 := Call(__e, tmp16305, tmp16313)


__e.TailApply(tmp16302, tmp16314)
return


} else {
__e.Return(W5036)
return
}


}, 1)

tmp16317 := MakeNative(func(__e *ControlFlow) {
W5037 := __e.Get(1)
_ = W5037
tmp16332 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5037)


if True == tmp16332 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16318 := MakeNative(func(__e *ControlFlow) {
W5038 := __e.Get(1)
_ = W5038
tmp16319 := MakeNative(func(__e *ControlFlow) {
W5039 := __e.Get(1)
_ = W5039
tmp16320 := MakeNative(func(__e *ControlFlow) {
W5040 := __e.Get(1)
_ = W5040
tmp16327 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5040)


if True == tmp16327 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16321 := MakeNative(func(__e *ControlFlow) {
W5041 := __e.Get(1)
_ = W5041
tmp16322 := MakeNative(func(__e *ControlFlow) {
W5042 := __e.Get(1)
_ = W5042
tmp16323 := PrimCons(W5038, W5041)

__e.TailApply(PrimFunc(symshen_4comb), W5042, tmp16323)
return


}, 1)

tmp16324 := Call(__e, PrimFunc(symshen_4in_1_6), W5040)


__e.TailApply(tmp16322, tmp16324)
return


}, 1)

tmp16325 := Call(__e, PrimFunc(symshen_4_5_1out), W5040)


__e.TailApply(tmp16321, tmp16325)
return


}


}, 1)

tmp16328 := Call(__e, PrimFunc(symshen_4_5rules_d_6), W5039)


__e.TailApply(tmp16320, tmp16328)
return


}, 1)

tmp16329 := Call(__e, PrimFunc(symshen_4in_1_6), W5037)


__e.TailApply(tmp16319, tmp16329)
return


}, 1)

tmp16330 := Call(__e, PrimFunc(symshen_4_5_1out), W5037)


__e.TailApply(tmp16318, tmp16330)
return


}


}, 1)

tmp16333 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5035)


tmp16334 := Call(__e, tmp16317, tmp16333)


__e.TailApply(tmp16301, tmp16334)
return


}, 1)

tmp16335 := Call(__e, ns2_1set, symshen_4_5rules_d_6, tmp16300)


_ = tmp16335

tmp16336 := MakeNative(func(__e *ControlFlow) {
V5047 := __e.Get(1)
_ = V5047
tmp16337 := MakeNative(func(__e *ControlFlow) {
W5048 := __e.Get(1)
_ = W5048
tmp16423 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5048)


if True == tmp16423 {
tmp16338 := MakeNative(func(__e *ControlFlow) {
W5058 := __e.Get(1)
_ = W5058
tmp16387 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5058)


if True == tmp16387 {
tmp16339 := MakeNative(func(__e *ControlFlow) {
W5068 := __e.Get(1)
_ = W5068
tmp16364 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5068)


if True == tmp16364 {
tmp16340 := MakeNative(func(__e *ControlFlow) {
W5075 := __e.Get(1)
_ = W5075
tmp16342 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5075)


if True == tmp16342 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W5075)
return
}


}, 1)

tmp16343 := MakeNative(func(__e *ControlFlow) {
W5076 := __e.Get(1)
_ = W5076
tmp16360 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5076)


if True == tmp16360 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16344 := MakeNative(func(__e *ControlFlow) {
W5077 := __e.Get(1)
_ = W5077
tmp16345 := MakeNative(func(__e *ControlFlow) {
W5078 := __e.Get(1)
_ = W5078
tmp16356 := Call(__e, PrimFunc(symshen_4hds_a_2), W5078, sym_1_6)


if True == tmp16356 {
tmp16346 := MakeNative(func(__e *ControlFlow) {
W5079 := __e.Get(1)
_ = W5079
tmp16353 := PrimIsPair(W5079)

if True == tmp16353 {
tmp16347 := MakeNative(func(__e *ControlFlow) {
W5080 := __e.Get(1)
_ = W5080
tmp16348 := MakeNative(func(__e *ControlFlow) {
W5081 := __e.Get(1)
_ = W5081
tmp16349 := Call(__e, PrimFunc(sym_8p), W5077, W5080)


__e.TailApply(PrimFunc(symshen_4comb), W5081, tmp16349)
return


}, 1)

tmp16350 := Call(__e, PrimFunc(symtail), W5079)


__e.TailApply(tmp16348, tmp16350)
return


}, 1)

tmp16351 := Call(__e, PrimFunc(symhead), W5079)


__e.TailApply(tmp16347, tmp16351)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16354 := Call(__e, PrimFunc(symtail), W5078)


__e.TailApply(tmp16346, tmp16354)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16357 := Call(__e, PrimFunc(symshen_4in_1_6), W5076)


__e.TailApply(tmp16345, tmp16357)
return


}, 1)

tmp16358 := Call(__e, PrimFunc(symshen_4_5_1out), W5076)


__e.TailApply(tmp16344, tmp16358)
return


}


}, 1)

tmp16361 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5047)


tmp16362 := Call(__e, tmp16343, tmp16361)


__e.TailApply(tmp16340, tmp16362)
return


} else {
__e.Return(W5068)
return
}


}, 1)

tmp16365 := MakeNative(func(__e *ControlFlow) {
W5069 := __e.Get(1)
_ = W5069
tmp16383 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5069)


if True == tmp16383 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16366 := MakeNative(func(__e *ControlFlow) {
W5070 := __e.Get(1)
_ = W5070
tmp16367 := MakeNative(func(__e *ControlFlow) {
W5071 := __e.Get(1)
_ = W5071
tmp16379 := Call(__e, PrimFunc(symshen_4hds_a_2), W5071, sym_5_1)


if True == tmp16379 {
tmp16368 := MakeNative(func(__e *ControlFlow) {
W5072 := __e.Get(1)
_ = W5072
tmp16376 := PrimIsPair(W5072)

if True == tmp16376 {
tmp16369 := MakeNative(func(__e *ControlFlow) {
W5073 := __e.Get(1)
_ = W5073
tmp16370 := MakeNative(func(__e *ControlFlow) {
W5074 := __e.Get(1)
_ = W5074
tmp16371 := Call(__e, PrimFunc(symshen_4correct), W5073)


tmp16372 := Call(__e, PrimFunc(sym_8p), W5070, tmp16371)


__e.TailApply(PrimFunc(symshen_4comb), W5074, tmp16372)
return


}, 1)

tmp16373 := Call(__e, PrimFunc(symtail), W5072)


__e.TailApply(tmp16370, tmp16373)
return


}, 1)

tmp16374 := Call(__e, PrimFunc(symhead), W5072)


__e.TailApply(tmp16369, tmp16374)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16377 := Call(__e, PrimFunc(symtail), W5071)


__e.TailApply(tmp16368, tmp16377)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16380 := Call(__e, PrimFunc(symshen_4in_1_6), W5069)


__e.TailApply(tmp16367, tmp16380)
return


}, 1)

tmp16381 := Call(__e, PrimFunc(symshen_4_5_1out), W5069)


__e.TailApply(tmp16366, tmp16381)
return


}


}, 1)

tmp16384 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5047)


tmp16385 := Call(__e, tmp16365, tmp16384)


__e.TailApply(tmp16339, tmp16385)
return


} else {
__e.Return(W5058)
return
}


}, 1)

tmp16388 := MakeNative(func(__e *ControlFlow) {
W5059 := __e.Get(1)
_ = W5059
tmp16419 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5059)


if True == tmp16419 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16389 := MakeNative(func(__e *ControlFlow) {
W5060 := __e.Get(1)
_ = W5060
tmp16390 := MakeNative(func(__e *ControlFlow) {
W5061 := __e.Get(1)
_ = W5061
tmp16415 := Call(__e, PrimFunc(symshen_4hds_a_2), W5061, sym_5_1)


if True == tmp16415 {
tmp16391 := MakeNative(func(__e *ControlFlow) {
W5062 := __e.Get(1)
_ = W5062
tmp16412 := PrimIsPair(W5062)

if True == tmp16412 {
tmp16392 := MakeNative(func(__e *ControlFlow) {
W5063 := __e.Get(1)
_ = W5063
tmp16393 := MakeNative(func(__e *ControlFlow) {
W5064 := __e.Get(1)
_ = W5064
tmp16408 := Call(__e, PrimFunc(symshen_4hds_a_2), W5064, symwhere)


if True == tmp16408 {
tmp16394 := MakeNative(func(__e *ControlFlow) {
W5065 := __e.Get(1)
_ = W5065
tmp16405 := PrimIsPair(W5065)

if True == tmp16405 {
tmp16395 := MakeNative(func(__e *ControlFlow) {
W5066 := __e.Get(1)
_ = W5066
tmp16396 := MakeNative(func(__e *ControlFlow) {
W5067 := __e.Get(1)
_ = W5067
tmp16397 := PrimCons(W5063, Nil)

tmp16398 := PrimCons(W5066, tmp16397)

tmp16399 := PrimCons(symwhere, tmp16398)

tmp16400 := Call(__e, PrimFunc(symshen_4correct), tmp16399)


tmp16401 := Call(__e, PrimFunc(sym_8p), W5060, tmp16400)


__e.TailApply(PrimFunc(symshen_4comb), W5067, tmp16401)
return


}, 1)

tmp16402 := Call(__e, PrimFunc(symtail), W5065)


__e.TailApply(tmp16396, tmp16402)
return


}, 1)

tmp16403 := Call(__e, PrimFunc(symhead), W5065)


__e.TailApply(tmp16395, tmp16403)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16406 := Call(__e, PrimFunc(symtail), W5064)


__e.TailApply(tmp16394, tmp16406)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16409 := Call(__e, PrimFunc(symtail), W5062)


__e.TailApply(tmp16393, tmp16409)
return


}, 1)

tmp16410 := Call(__e, PrimFunc(symhead), W5062)


__e.TailApply(tmp16392, tmp16410)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16413 := Call(__e, PrimFunc(symtail), W5061)


__e.TailApply(tmp16391, tmp16413)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16416 := Call(__e, PrimFunc(symshen_4in_1_6), W5059)


__e.TailApply(tmp16390, tmp16416)
return


}, 1)

tmp16417 := Call(__e, PrimFunc(symshen_4_5_1out), W5059)


__e.TailApply(tmp16389, tmp16417)
return


}


}, 1)

tmp16420 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5047)


tmp16421 := Call(__e, tmp16388, tmp16420)


__e.TailApply(tmp16338, tmp16421)
return


} else {
__e.Return(W5048)
return
}


}, 1)

tmp16424 := MakeNative(func(__e *ControlFlow) {
W5049 := __e.Get(1)
_ = W5049
tmp16454 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5049)


if True == tmp16454 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16425 := MakeNative(func(__e *ControlFlow) {
W5050 := __e.Get(1)
_ = W5050
tmp16426 := MakeNative(func(__e *ControlFlow) {
W5051 := __e.Get(1)
_ = W5051
tmp16450 := Call(__e, PrimFunc(symshen_4hds_a_2), W5051, sym_1_6)


if True == tmp16450 {
tmp16427 := MakeNative(func(__e *ControlFlow) {
W5052 := __e.Get(1)
_ = W5052
tmp16447 := PrimIsPair(W5052)

if True == tmp16447 {
tmp16428 := MakeNative(func(__e *ControlFlow) {
W5053 := __e.Get(1)
_ = W5053
tmp16429 := MakeNative(func(__e *ControlFlow) {
W5054 := __e.Get(1)
_ = W5054
tmp16443 := Call(__e, PrimFunc(symshen_4hds_a_2), W5054, symwhere)


if True == tmp16443 {
tmp16430 := MakeNative(func(__e *ControlFlow) {
W5055 := __e.Get(1)
_ = W5055
tmp16440 := PrimIsPair(W5055)

if True == tmp16440 {
tmp16431 := MakeNative(func(__e *ControlFlow) {
W5056 := __e.Get(1)
_ = W5056
tmp16432 := MakeNative(func(__e *ControlFlow) {
W5057 := __e.Get(1)
_ = W5057
tmp16433 := PrimCons(W5053, Nil)

tmp16434 := PrimCons(W5056, tmp16433)

tmp16435 := PrimCons(symwhere, tmp16434)

tmp16436 := Call(__e, PrimFunc(sym_8p), W5050, tmp16435)


__e.TailApply(PrimFunc(symshen_4comb), W5057, tmp16436)
return


}, 1)

tmp16437 := Call(__e, PrimFunc(symtail), W5055)


__e.TailApply(tmp16432, tmp16437)
return


}, 1)

tmp16438 := Call(__e, PrimFunc(symhead), W5055)


__e.TailApply(tmp16431, tmp16438)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16441 := Call(__e, PrimFunc(symtail), W5054)


__e.TailApply(tmp16430, tmp16441)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16444 := Call(__e, PrimFunc(symtail), W5052)


__e.TailApply(tmp16429, tmp16444)
return


}, 1)

tmp16445 := Call(__e, PrimFunc(symhead), W5052)


__e.TailApply(tmp16428, tmp16445)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16448 := Call(__e, PrimFunc(symtail), W5051)


__e.TailApply(tmp16427, tmp16448)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16451 := Call(__e, PrimFunc(symshen_4in_1_6), W5049)


__e.TailApply(tmp16426, tmp16451)
return


}, 1)

tmp16452 := Call(__e, PrimFunc(symshen_4_5_1out), W5049)


__e.TailApply(tmp16425, tmp16452)
return


}


}, 1)

tmp16455 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5047)


tmp16456 := Call(__e, tmp16424, tmp16455)


__e.TailApply(tmp16337, tmp16456)
return


}, 1)

tmp16457 := Call(__e, ns2_1set, symshen_4_5rule_d_6, tmp16336)


_ = tmp16457

tmp16458 := MakeNative(func(__e *ControlFlow) {
V5082 := __e.Get(1)
_ = V5082
tmp16606 := PrimIsPair(V5082)

var ifres16550 Obj

if True == tmp16606 {
tmp16604 := PrimHead(V5082)

tmp16605 := PrimEqual(symwhere, tmp16604)

var ifres16552 Obj

if True == tmp16605 {
tmp16602 := PrimTail(V5082)

tmp16603 := PrimIsPair(tmp16602)

var ifres16554 Obj

if True == tmp16603 {
tmp16599 := PrimTail(V5082)

tmp16600 := PrimTail(tmp16599)

tmp16601 := PrimIsPair(tmp16600)

var ifres16556 Obj

if True == tmp16601 {
tmp16595 := PrimTail(V5082)

tmp16596 := PrimTail(tmp16595)

tmp16597 := PrimHead(tmp16596)

tmp16598 := PrimIsPair(tmp16597)

var ifres16558 Obj

if True == tmp16598 {
tmp16590 := PrimTail(V5082)

tmp16591 := PrimTail(tmp16590)

tmp16592 := PrimHead(tmp16591)

tmp16593 := PrimHead(tmp16592)

tmp16594 := PrimEqual(symfail_1if, tmp16593)

var ifres16560 Obj

if True == tmp16594 {
tmp16585 := PrimTail(V5082)

tmp16586 := PrimTail(tmp16585)

tmp16587 := PrimHead(tmp16586)

tmp16588 := PrimTail(tmp16587)

tmp16589 := PrimIsPair(tmp16588)

var ifres16562 Obj

if True == tmp16589 {
tmp16579 := PrimTail(V5082)

tmp16580 := PrimTail(tmp16579)

tmp16581 := PrimHead(tmp16580)

tmp16582 := PrimTail(tmp16581)

tmp16583 := PrimTail(tmp16582)

tmp16584 := PrimIsPair(tmp16583)

var ifres16564 Obj

if True == tmp16584 {
tmp16572 := PrimTail(V5082)

tmp16573 := PrimTail(tmp16572)

tmp16574 := PrimHead(tmp16573)

tmp16575 := PrimTail(tmp16574)

tmp16576 := PrimTail(tmp16575)

tmp16577 := PrimTail(tmp16576)

tmp16578 := PrimEqual(Nil, tmp16577)

var ifres16566 Obj

if True == tmp16578 {
tmp16568 := PrimTail(V5082)

tmp16569 := PrimTail(tmp16568)

tmp16570 := PrimTail(tmp16569)

tmp16571 := PrimEqual(Nil, tmp16570)

var ifres16567 Obj

if True == tmp16571 {
ifres16567 = True


} else {
ifres16567 = False


}

ifres16566 = ifres16567


} else {
ifres16566 = False


}

var ifres16565 Obj

if True == ifres16566 {
ifres16565 = True


} else {
ifres16565 = False


}

ifres16564 = ifres16565


} else {
ifres16564 = False


}

var ifres16563 Obj

if True == ifres16564 {
ifres16563 = True


} else {
ifres16563 = False


}

ifres16562 = ifres16563


} else {
ifres16562 = False


}

var ifres16561 Obj

if True == ifres16562 {
ifres16561 = True


} else {
ifres16561 = False


}

ifres16560 = ifres16561


} else {
ifres16560 = False


}

var ifres16559 Obj

if True == ifres16560 {
ifres16559 = True


} else {
ifres16559 = False


}

ifres16558 = ifres16559


} else {
ifres16558 = False


}

var ifres16557 Obj

if True == ifres16558 {
ifres16557 = True


} else {
ifres16557 = False


}

ifres16556 = ifres16557


} else {
ifres16556 = False


}

var ifres16555 Obj

if True == ifres16556 {
ifres16555 = True


} else {
ifres16555 = False


}

ifres16554 = ifres16555


} else {
ifres16554 = False


}

var ifres16553 Obj

if True == ifres16554 {
ifres16553 = True


} else {
ifres16553 = False


}

ifres16552 = ifres16553


} else {
ifres16552 = False


}

var ifres16551 Obj

if True == ifres16552 {
ifres16551 = True


} else {
ifres16551 = False


}

ifres16550 = ifres16551


} else {
ifres16550 = False


}

if True == ifres16550 {
tmp16459 := PrimTail(V5082)

tmp16460 := PrimHead(tmp16459)

tmp16461 := PrimTail(V5082)

tmp16462 := PrimTail(tmp16461)

tmp16463 := PrimHead(tmp16462)

tmp16464 := PrimTail(tmp16463)

tmp16465 := PrimCons(tmp16464, Nil)

tmp16466 := PrimCons(symnot, tmp16465)

tmp16467 := PrimCons(tmp16466, Nil)

tmp16468 := PrimCons(tmp16460, tmp16467)

tmp16469 := PrimCons(symand, tmp16468)

tmp16470 := PrimTail(V5082)

tmp16471 := PrimTail(tmp16470)

tmp16472 := PrimHead(tmp16471)

tmp16473 := PrimTail(tmp16472)

tmp16474 := PrimTail(tmp16473)

tmp16475 := PrimCons(tmp16469, tmp16474)

__e.Return(PrimCons(symwhere, tmp16475))
return


} else {
tmp16548 := PrimIsPair(V5082)

var ifres16529 Obj

if True == tmp16548 {
tmp16546 := PrimHead(V5082)

tmp16547 := PrimEqual(symwhere, tmp16546)

var ifres16531 Obj

if True == tmp16547 {
tmp16544 := PrimTail(V5082)

tmp16545 := PrimIsPair(tmp16544)

var ifres16533 Obj

if True == tmp16545 {
tmp16541 := PrimTail(V5082)

tmp16542 := PrimTail(tmp16541)

tmp16543 := PrimIsPair(tmp16542)

var ifres16535 Obj

if True == tmp16543 {
tmp16537 := PrimTail(V5082)

tmp16538 := PrimTail(tmp16537)

tmp16539 := PrimTail(tmp16538)

tmp16540 := PrimEqual(Nil, tmp16539)

var ifres16536 Obj

if True == tmp16540 {
ifres16536 = True


} else {
ifres16536 = False


}

ifres16535 = ifres16536


} else {
ifres16535 = False


}

var ifres16534 Obj

if True == ifres16535 {
ifres16534 = True


} else {
ifres16534 = False


}

ifres16533 = ifres16534


} else {
ifres16533 = False


}

var ifres16532 Obj

if True == ifres16533 {
ifres16532 = True


} else {
ifres16532 = False


}

ifres16531 = ifres16532


} else {
ifres16531 = False


}

var ifres16530 Obj

if True == ifres16531 {
ifres16530 = True


} else {
ifres16530 = False


}

ifres16529 = ifres16530


} else {
ifres16529 = False


}

if True == ifres16529 {
tmp16476 := PrimTail(V5082)

tmp16477 := PrimHead(tmp16476)

tmp16478 := PrimTail(V5082)

tmp16479 := PrimTail(tmp16478)

tmp16480 := PrimHead(tmp16479)

tmp16481 := PrimCons(symfail, Nil)

tmp16482 := PrimCons(tmp16481, Nil)

tmp16483 := PrimCons(tmp16480, tmp16482)

tmp16484 := PrimCons(sym_a, tmp16483)

tmp16485 := PrimCons(tmp16484, Nil)

tmp16486 := PrimCons(symnot, tmp16485)

tmp16487 := PrimCons(tmp16486, Nil)

tmp16488 := PrimCons(tmp16477, tmp16487)

tmp16489 := PrimCons(symand, tmp16488)

tmp16490 := PrimTail(V5082)

tmp16491 := PrimTail(tmp16490)

tmp16492 := PrimCons(tmp16489, tmp16491)

__e.Return(PrimCons(symwhere, tmp16492))
return


} else {
tmp16527 := PrimIsPair(V5082)

var ifres16508 Obj

if True == tmp16527 {
tmp16525 := PrimHead(V5082)

tmp16526 := PrimEqual(symfail_1if, tmp16525)

var ifres16510 Obj

if True == tmp16526 {
tmp16523 := PrimTail(V5082)

tmp16524 := PrimIsPair(tmp16523)

var ifres16512 Obj

if True == tmp16524 {
tmp16520 := PrimTail(V5082)

tmp16521 := PrimTail(tmp16520)

tmp16522 := PrimIsPair(tmp16521)

var ifres16514 Obj

if True == tmp16522 {
tmp16516 := PrimTail(V5082)

tmp16517 := PrimTail(tmp16516)

tmp16518 := PrimTail(tmp16517)

tmp16519 := PrimEqual(Nil, tmp16518)

var ifres16515 Obj

if True == tmp16519 {
ifres16515 = True


} else {
ifres16515 = False


}

ifres16514 = ifres16515


} else {
ifres16514 = False


}

var ifres16513 Obj

if True == ifres16514 {
ifres16513 = True


} else {
ifres16513 = False


}

ifres16512 = ifres16513


} else {
ifres16512 = False


}

var ifres16511 Obj

if True == ifres16512 {
ifres16511 = True


} else {
ifres16511 = False


}

ifres16510 = ifres16511


} else {
ifres16510 = False


}

var ifres16509 Obj

if True == ifres16510 {
ifres16509 = True


} else {
ifres16509 = False


}

ifres16508 = ifres16509


} else {
ifres16508 = False


}

if True == ifres16508 {
tmp16493 := PrimTail(V5082)

tmp16494 := PrimCons(tmp16493, Nil)

tmp16495 := PrimCons(symnot, tmp16494)

tmp16496 := PrimTail(V5082)

tmp16497 := PrimTail(tmp16496)

tmp16498 := PrimCons(tmp16495, tmp16497)

__e.Return(PrimCons(symwhere, tmp16498))
return


} else {
tmp16499 := PrimCons(symfail, Nil)

tmp16500 := PrimCons(tmp16499, Nil)

tmp16501 := PrimCons(V5082, tmp16500)

tmp16502 := PrimCons(sym_a, tmp16501)

tmp16503 := PrimCons(tmp16502, Nil)

tmp16504 := PrimCons(symnot, tmp16503)

tmp16505 := PrimCons(V5082, Nil)

tmp16506 := PrimCons(tmp16504, tmp16505)

__e.Return(PrimCons(symwhere, tmp16506))
return


}


}


}


}, 1)

tmp16607 := Call(__e, ns2_1set, symshen_4correct, tmp16458)


_ = tmp16607

tmp16608 := MakeNative(func(__e *ControlFlow) {
V5083 := __e.Get(1)
_ = V5083
V5084 := __e.Get(2)
_ = V5084
V5085 := __e.Get(3)
_ = V5085
V5086 := __e.Get(4)
_ = V5086
V5087 := __e.Get(5)
_ = V5087
V5088 := __e.Get(6)
_ = V5088
V5089 := __e.Get(7)
_ = V5089
V5090 := __e.Get(8)
_ = V5090
tmp16609 := MakeNative(func(__e *ControlFlow) {
W5091 := __e.Get(1)
_ = W5091
tmp16610 := MakeNative(func(__e *ControlFlow) {
W5092 := __e.Get(1)
_ = W5092
tmp16640 := PrimEqual(W5092, False)

if True == tmp16640 {
tmp16611 := MakeNative(func(__e *ControlFlow) {
W5094 := __e.Get(1)
_ = W5094
tmp16613 := PrimEqual(W5094, False)

if True == tmp16613 {
__e.TailApply(PrimFunc(symshen_4unlock), V5088, W5091)
return
} else {
__e.Return(W5094)
return
}


}, 1)

tmp16638 := Call(__e, PrimFunc(symshen_4unlocked_2), V5088)


var ifres16614 Obj

if True == tmp16638 {
tmp16615 := MakeNative(func(__e *ControlFlow) {
W5095 := __e.Get(1)
_ = W5095
tmp16635 := PrimIsPair(W5095)

if True == tmp16635 {
tmp16616 := MakeNative(func(__e *ControlFlow) {
W5096 := __e.Get(1)
_ = W5096
tmp16617 := MakeNative(func(__e *ControlFlow) {
W5097 := __e.Get(1)
_ = W5097
tmp16618 := MakeNative(func(__e *ControlFlow) {
W5098 := __e.Get(1)
_ = W5098
tmp16619 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16619

tmp16620 := Call(__e, PrimFunc(symshen_4deref), W5096, V5087)


tmp16621 := Call(__e, PrimFunc(symshen_4freshen_1rule), tmp16620)


tmp16622 := MakeNative(func(__e *ControlFlow) {
tmp16623 := Call(__e, PrimFunc(symshen_4lazyderef), W5098, V5087)


tmp16624 := Call(__e, PrimFunc(symfst), tmp16623)


tmp16625 := Call(__e, PrimFunc(symshen_4lazyderef), W5098, V5087)


tmp16626 := Call(__e, PrimFunc(symsnd), tmp16625)


tmp16627 := MakeNative(func(__e *ControlFlow) {
tmp16628 := MakeNative(func(__e *ControlFlow) {
tmp16629 := PrimNumberAdd(V5086, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4t_d_1rules), V5083, W5097, V5085, tmp16629, V5087, V5088, W5091, V5090)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5087, V5088, W5091, tmp16628)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rule), V5083, V5086, tmp16624, tmp16626, V5085, V5087, V5088, W5091, tmp16627)
return


}, 0)

tmp16630 := Call(__e, PrimFunc(symbind), W5098, tmp16621, V5087, V5088, W5091, tmp16622)


__e.TailApply(PrimFunc(symshen_4gc), V5087, tmp16630)
return


}, 1)

tmp16631 := Call(__e, PrimFunc(symshen_4newpv), V5087)


__e.TailApply(tmp16618, tmp16631)
return


}, 1)

tmp16632 := PrimTail(W5095)

__e.TailApply(tmp16617, tmp16632)
return


}, 1)

tmp16633 := PrimHead(W5095)

__e.TailApply(tmp16616, tmp16633)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16636 := Call(__e, PrimFunc(symshen_4lazyderef), V5084, V5087)


tmp16637 := Call(__e, tmp16615, tmp16636)


ifres16614 = tmp16637


} else {
ifres16614 = False


}

__e.TailApply(tmp16611, ifres16614)
return


} else {
__e.Return(W5092)
return
}


}, 1)

tmp16648 := Call(__e, PrimFunc(symshen_4unlocked_2), V5088)


var ifres16641 Obj

if True == tmp16648 {
tmp16642 := MakeNative(func(__e *ControlFlow) {
W5093 := __e.Get(1)
_ = W5093
tmp16645 := PrimEqual(W5093, Nil)

if True == tmp16645 {
tmp16643 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16643

__e.TailApply(PrimFunc(symthaw), V5090)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16646 := Call(__e, PrimFunc(symshen_4lazyderef), V5084, V5087)


tmp16647 := Call(__e, tmp16642, tmp16646)


ifres16641 = tmp16647


} else {
ifres16641 = False


}

__e.TailApply(tmp16610, ifres16641)
return


}, 1)

tmp16649 := PrimNumberAdd(V5089, MakeNumber(1))

__e.TailApply(tmp16609, tmp16649)
return


}, 8)

tmp16650 := Call(__e, ns2_1set, symshen_4t_d_1rules, tmp16608)


_ = tmp16650

tmp16651 := MakeNative(func(__e *ControlFlow) {
V5099 := __e.Get(1)
_ = V5099
tmp16664 := Call(__e, PrimFunc(symtuple_2), V5099)


if True == tmp16664 {
tmp16652 := MakeNative(func(__e *ControlFlow) {
W5100 := __e.Get(1)
_ = W5100
tmp16653 := MakeNative(func(__e *ControlFlow) {
W5101 := __e.Get(1)
_ = W5101
tmp16654 := Call(__e, PrimFunc(symfst), V5099)


tmp16655 := Call(__e, PrimFunc(symshen_4freshen), W5101, tmp16654)


tmp16656 := Call(__e, PrimFunc(symsnd), V5099)


tmp16657 := Call(__e, PrimFunc(symshen_4freshen), W5101, tmp16656)


__e.TailApply(PrimFunc(sym_8p), tmp16655, tmp16657)
return


}, 1)

tmp16658 := MakeNative(func(__e *ControlFlow) {
Z5102 := __e.Get(1)
_ = Z5102
tmp16659 := Call(__e, PrimFunc(symshen_4freshterm), Z5102)


__e.Return(PrimCons(Z5102, tmp16659))
return


}, 1)

tmp16660 := Call(__e, PrimFunc(symmap), tmp16658, W5100)


__e.TailApply(tmp16653, tmp16660)
return


}, 1)

tmp16661 := Call(__e, PrimFunc(symfst), V5099)


tmp16662 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp16661)


__e.TailApply(tmp16652, tmp16662)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.freshen-rule")))
return
}


}, 1)

tmp16665 := Call(__e, ns2_1set, symshen_4freshen_1rule, tmp16651)


_ = tmp16665

tmp16666 := MakeNative(func(__e *ControlFlow) {
V5103 := __e.Get(1)
_ = V5103
V5104 := __e.Get(2)
_ = V5104
tmp16680 := PrimEqual(Nil, V5103)

if True == tmp16680 {
__e.Return(V5104)
return
} else {
tmp16678 := PrimIsPair(V5103)

var ifres16674 Obj

if True == tmp16678 {
tmp16676 := PrimHead(V5103)

tmp16677 := PrimIsPair(tmp16676)

var ifres16675 Obj

if True == tmp16677 {
ifres16675 = True


} else {
ifres16675 = False


}

ifres16674 = ifres16675


} else {
ifres16674 = False


}

if True == ifres16674 {
tmp16667 := PrimTail(V5103)

tmp16668 := PrimHead(V5103)

tmp16669 := PrimHead(tmp16668)

tmp16670 := PrimHead(V5103)

tmp16671 := PrimTail(tmp16670)

tmp16672 := Call(__e, PrimFunc(symshen_4beta), tmp16669, tmp16671, V5104)


__e.TailApply(PrimFunc(symshen_4freshen), tmp16667, tmp16672)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.freshen")))
return
}


}


}, 2)

tmp16681 := Call(__e, ns2_1set, symshen_4freshen, tmp16666)


_ = tmp16681

tmp16682 := MakeNative(func(__e *ControlFlow) {
V5105 := __e.Get(1)
_ = V5105
V5106 := __e.Get(2)
_ = V5106
V5107 := __e.Get(3)
_ = V5107
V5108 := __e.Get(4)
_ = V5108
V5109 := __e.Get(5)
_ = V5109
V5110 := __e.Get(6)
_ = V5110
V5111 := __e.Get(7)
_ = V5111
V5112 := __e.Get(8)
_ = V5112
V5113 := __e.Get(9)
_ = V5113
tmp16683 := MakeNative(func(__e *ControlFlow) {
W5114 := __e.Get(1)
_ = W5114
tmp16696 := PrimEqual(W5114, False)

if True == tmp16696 {
tmp16694 := Call(__e, PrimFunc(symshen_4unlocked_2), V5111)


if True == tmp16694 {
tmp16684 := MakeNative(func(__e *ControlFlow) {
W5115 := __e.Get(1)
_ = W5115
tmp16685 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16685

tmp16686 := Call(__e, PrimFunc(symshen_4app), V5105, MakeString("\n"), symshen_4a)


tmp16687 := PrimStringConcat(MakeString(" of "), tmp16686)

tmp16688 := Call(__e, PrimFunc(symshen_4app), V5106, tmp16687, symshen_4a)


tmp16689 := PrimStringConcat(MakeString("type error in rule "), tmp16688)

tmp16690 := PrimSimpleError(tmp16689)

tmp16691 := Call(__e, PrimFunc(symbind), W5115, tmp16690, V5110, V5111, V5112, V5113)


__e.TailApply(PrimFunc(symshen_4gc), V5110, tmp16691)
return


}, 1)

tmp16692 := Call(__e, PrimFunc(symshen_4newpv), V5110)


__e.TailApply(tmp16684, tmp16692)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5114)
return
}


}, 1)

tmp16700 := Call(__e, PrimFunc(symshen_4unlocked_2), V5111)


var ifres16697 Obj

if True == tmp16700 {
tmp16698 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16698

tmp16699 := Call(__e, PrimFunc(symshen_4t_d_1rule_1h), V5107, V5108, V5109, V5110, V5111, V5112, V5113)


ifres16697 = tmp16699


} else {
ifres16697 = False


}

__e.TailApply(tmp16683, ifres16697)
return


}, 9)

tmp16701 := Call(__e, ns2_1set, symshen_4t_d_1rule, tmp16682)


_ = tmp16701

tmp16702 := MakeNative(func(__e *ControlFlow) {
V5116 := __e.Get(1)
_ = V5116
V5117 := __e.Get(2)
_ = V5117
V5118 := __e.Get(3)
_ = V5118
V5119 := __e.Get(4)
_ = V5119
V5120 := __e.Get(5)
_ = V5120
V5121 := __e.Get(6)
_ = V5121
V5122 := __e.Get(7)
_ = V5122
tmp16703 := MakeNative(func(__e *ControlFlow) {
W5123 := __e.Get(1)
_ = W5123
tmp16704 := MakeNative(func(__e *ControlFlow) {
W5124 := __e.Get(1)
_ = W5124
tmp16727 := PrimEqual(W5124, False)

if True == tmp16727 {
tmp16705 := MakeNative(func(__e *ControlFlow) {
W5131 := __e.Get(1)
_ = W5131
tmp16707 := PrimEqual(W5131, False)

if True == tmp16707 {
__e.TailApply(PrimFunc(symshen_4unlock), V5120, W5123)
return
} else {
__e.Return(W5131)
return
}


}, 1)

tmp16725 := Call(__e, PrimFunc(symshen_4unlocked_2), V5120)


var ifres16708 Obj

if True == tmp16725 {
tmp16709 := MakeNative(func(__e *ControlFlow) {
W5132 := __e.Get(1)
_ = W5132
tmp16710 := MakeNative(func(__e *ControlFlow) {
W5133 := __e.Get(1)
_ = W5133
tmp16711 := MakeNative(func(__e *ControlFlow) {
W5134 := __e.Get(1)
_ = W5134
tmp16712 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16712

tmp16713 := Call(__e, PrimFunc(symshen_4freshterms), V5116)


tmp16714 := MakeNative(func(__e *ControlFlow) {
tmp16715 := MakeNative(func(__e *ControlFlow) {
tmp16716 := MakeNative(func(__e *ControlFlow) {
tmp16717 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5117, W5133, W5134, V5119, V5120, W5123, V5122)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4myassume), V5116, V5118, W5134, V5119, V5120, W5123, tmp16717)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5119, V5120, W5123, tmp16716)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1integrity), V5116, V5118, W5132, W5133, V5119, V5120, W5123, tmp16715)
return


}, 0)

tmp16718 := Call(__e, PrimFunc(symshen_4p_1hyps), tmp16713, W5132, V5119, V5120, W5123, tmp16714)


__e.TailApply(PrimFunc(symshen_4gc), V5119, tmp16718)
return


}, 1)

tmp16719 := Call(__e, PrimFunc(symshen_4newpv), V5119)


tmp16720 := Call(__e, tmp16711, tmp16719)


__e.TailApply(PrimFunc(symshen_4gc), V5119, tmp16720)
return


}, 1)

tmp16721 := Call(__e, PrimFunc(symshen_4newpv), V5119)


tmp16722 := Call(__e, tmp16710, tmp16721)


__e.TailApply(PrimFunc(symshen_4gc), V5119, tmp16722)
return


}, 1)

tmp16723 := Call(__e, PrimFunc(symshen_4newpv), V5119)


tmp16724 := Call(__e, tmp16709, tmp16723)


ifres16708 = tmp16724


} else {
ifres16708 = False


}

__e.TailApply(tmp16705, ifres16708)
return


} else {
__e.Return(W5124)
return
}


}, 1)

tmp16757 := Call(__e, PrimFunc(symshen_4unlocked_2), V5120)


var ifres16728 Obj

if True == tmp16757 {
tmp16729 := MakeNative(func(__e *ControlFlow) {
W5125 := __e.Get(1)
_ = W5125
tmp16754 := PrimEqual(W5125, Nil)

if True == tmp16754 {
tmp16730 := MakeNative(func(__e *ControlFlow) {
W5126 := __e.Get(1)
_ = W5126
tmp16751 := PrimIsPair(W5126)

if True == tmp16751 {
tmp16731 := MakeNative(func(__e *ControlFlow) {
W5127 := __e.Get(1)
_ = W5127
tmp16747 := PrimEqual(W5127, sym_1_1_6)

if True == tmp16747 {
tmp16732 := MakeNative(func(__e *ControlFlow) {
W5128 := __e.Get(1)
_ = W5128
tmp16743 := PrimIsPair(W5128)

if True == tmp16743 {
tmp16733 := MakeNative(func(__e *ControlFlow) {
W5129 := __e.Get(1)
_ = W5129
tmp16734 := MakeNative(func(__e *ControlFlow) {
W5130 := __e.Get(1)
_ = W5130
tmp16738 := PrimEqual(W5130, Nil)

if True == tmp16738 {
tmp16735 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16735

tmp16736 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5117, W5129, Nil, V5119, V5120, W5123, V5122)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5119, V5120, W5123, tmp16736)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16739 := PrimTail(W5128)

tmp16740 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16739, V5119)


__e.TailApply(tmp16734, tmp16740)
return


}, 1)

tmp16741 := PrimHead(W5128)

__e.TailApply(tmp16733, tmp16741)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16744 := PrimTail(W5126)

tmp16745 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16744, V5119)


__e.TailApply(tmp16732, tmp16745)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16748 := PrimHead(W5126)

tmp16749 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16748, V5119)


__e.TailApply(tmp16731, tmp16749)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16752 := Call(__e, PrimFunc(symshen_4lazyderef), V5118, V5119)


__e.TailApply(tmp16730, tmp16752)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16755 := Call(__e, PrimFunc(symshen_4lazyderef), V5116, V5119)


tmp16756 := Call(__e, tmp16729, tmp16755)


ifres16728 = tmp16756


} else {
ifres16728 = False


}

__e.TailApply(tmp16704, ifres16728)
return


}, 1)

tmp16758 := PrimNumberAdd(V5121, MakeNumber(1))

__e.TailApply(tmp16703, tmp16758)
return


}, 7)

tmp16759 := Call(__e, ns2_1set, symshen_4t_d_1rule_1h, tmp16702)


_ = tmp16759

tmp16760 := MakeNative(func(__e *ControlFlow) {
V5135 := __e.Get(1)
_ = V5135
V5136 := __e.Get(2)
_ = V5136
V5137 := __e.Get(3)
_ = V5137
V5138 := __e.Get(4)
_ = V5138
V5139 := __e.Get(5)
_ = V5139
V5140 := __e.Get(6)
_ = V5140
V5141 := __e.Get(7)
_ = V5141
tmp16761 := MakeNative(func(__e *ControlFlow) {
W5142 := __e.Get(1)
_ = W5142
tmp16914 := PrimEqual(W5142, False)

if True == tmp16914 {
tmp16912 := Call(__e, PrimFunc(symshen_4unlocked_2), V5139)


if True == tmp16912 {
tmp16762 := MakeNative(func(__e *ControlFlow) {
W5146 := __e.Get(1)
_ = W5146
tmp16909 := PrimIsPair(W5146)

if True == tmp16909 {
tmp16763 := MakeNative(func(__e *ControlFlow) {
W5147 := __e.Get(1)
_ = W5147
tmp16764 := MakeNative(func(__e *ControlFlow) {
W5148 := __e.Get(1)
_ = W5148
tmp16765 := MakeNative(func(__e *ControlFlow) {
W5149 := __e.Get(1)
_ = W5149
tmp16904 := PrimIsPair(W5149)

if True == tmp16904 {
tmp16766 := MakeNative(func(__e *ControlFlow) {
W5150 := __e.Get(1)
_ = W5150
tmp16767 := MakeNative(func(__e *ControlFlow) {
W5151 := __e.Get(1)
_ = W5151
tmp16899 := PrimIsPair(W5151)

if True == tmp16899 {
tmp16768 := MakeNative(func(__e *ControlFlow) {
W5152 := __e.Get(1)
_ = W5152
tmp16895 := PrimEqual(W5152, sym_1_1_6)

if True == tmp16895 {
tmp16769 := MakeNative(func(__e *ControlFlow) {
W5153 := __e.Get(1)
_ = W5153
tmp16891 := PrimIsPair(W5153)

if True == tmp16891 {
tmp16770 := MakeNative(func(__e *ControlFlow) {
W5154 := __e.Get(1)
_ = W5154
tmp16771 := MakeNative(func(__e *ControlFlow) {
W5155 := __e.Get(1)
_ = W5155
tmp16886 := PrimEqual(W5155, Nil)

if True == tmp16886 {
tmp16772 := MakeNative(func(__e *ControlFlow) {
W5156 := __e.Get(1)
_ = W5156
tmp16773 := MakeNative(func(__e *ControlFlow) {
W5157 := __e.Get(1)
_ = W5157
tmp16877 := PrimIsPair(W5156)

if True == tmp16877 {
tmp16774 := MakeNative(func(__e *ControlFlow) {
W5162 := __e.Get(1)
_ = W5162
tmp16775 := MakeNative(func(__e *ControlFlow) {
W5163 := __e.Get(1)
_ = W5163
tmp16845 := PrimIsPair(W5162)

if True == tmp16845 {
tmp16776 := MakeNative(func(__e *ControlFlow) {
W5168 := __e.Get(1)
_ = W5168
tmp16777 := MakeNative(func(__e *ControlFlow) {
W5169 := __e.Get(1)
_ = W5169
tmp16778 := MakeNative(func(__e *ControlFlow) {
W5170 := __e.Get(1)
_ = W5170
tmp16820 := PrimIsPair(W5169)

if True == tmp16820 {
tmp16779 := MakeNative(func(__e *ControlFlow) {
W5173 := __e.Get(1)
_ = W5173
tmp16780 := MakeNative(func(__e *ControlFlow) {
W5174 := __e.Get(1)
_ = W5174
tmp16781 := MakeNative(func(__e *ControlFlow) {
W5175 := __e.Get(1)
_ = W5175
tmp16801 := PrimIsPair(W5174)

if True == tmp16801 {
tmp16782 := MakeNative(func(__e *ControlFlow) {
W5177 := __e.Get(1)
_ = W5177
tmp16783 := MakeNative(func(__e *ControlFlow) {
W5178 := __e.Get(1)
_ = W5178
tmp16784 := MakeNative(func(__e *ControlFlow) {
W5179 := __e.Get(1)
_ = W5179
tmp16788 := PrimEqual(W5178, Nil)

if True == tmp16788 {
__e.TailApply(PrimFunc(symthaw), W5179)
return
} else {
tmp16786 := Call(__e, PrimFunc(symshen_4pvar_2), W5178)


if True == tmp16786 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5178, Nil, V5138, W5179)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp16789 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5175, W5177)
return
}, 0)

__e.TailApply(tmp16784, tmp16789)
return


}, 1)

tmp16790 := PrimTail(W5174)

tmp16791 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16790, V5138)


__e.TailApply(tmp16783, tmp16791)
return


}, 1)

tmp16792 := PrimHead(W5174)

__e.TailApply(tmp16782, tmp16792)
return


} else {
tmp16799 := Call(__e, PrimFunc(symshen_4pvar_2), W5174)


if True == tmp16799 {
tmp16793 := MakeNative(func(__e *ControlFlow) {
W5180 := __e.Get(1)
_ = W5180
tmp16794 := PrimCons(W5180, Nil)

tmp16795 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5175, W5180)
return
}, 0)

tmp16796 := Call(__e, PrimFunc(symshen_4bind_b), W5174, tmp16794, V5138, tmp16795)


__e.TailApply(PrimFunc(symshen_4gc), V5138, tmp16796)
return


}, 1)

tmp16797 := Call(__e, PrimFunc(symshen_4newpv), V5138)


__e.TailApply(tmp16793, tmp16797)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16802 := MakeNative(func(__e *ControlFlow) {
Z5176 := __e.Get(1)
_ = Z5176
tmp16803 := Call(__e, W5170, W5173)


__e.TailApply(tmp16803, Z5176)
return


}, 1)

__e.TailApply(tmp16781, tmp16802)
return


}, 1)

tmp16804 := PrimTail(W5169)

tmp16805 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16804, V5138)


__e.TailApply(tmp16780, tmp16805)
return


}, 1)

tmp16806 := PrimHead(W5169)

__e.TailApply(tmp16779, tmp16806)
return


} else {
tmp16818 := Call(__e, PrimFunc(symshen_4pvar_2), W5169)


if True == tmp16818 {
tmp16807 := MakeNative(func(__e *ControlFlow) {
W5181 := __e.Get(1)
_ = W5181
tmp16808 := MakeNative(func(__e *ControlFlow) {
W5182 := __e.Get(1)
_ = W5182
tmp16809 := PrimCons(W5182, Nil)

tmp16810 := PrimCons(W5181, tmp16809)

tmp16811 := MakeNative(func(__e *ControlFlow) {
tmp16812 := Call(__e, W5170, W5181)


__e.TailApply(tmp16812, W5182)
return


}, 0)

tmp16813 := Call(__e, PrimFunc(symshen_4bind_b), W5169, tmp16810, V5138, tmp16811)


__e.TailApply(PrimFunc(symshen_4gc), V5138, tmp16813)
return


}, 1)

tmp16814 := Call(__e, PrimFunc(symshen_4newpv), V5138)


tmp16815 := Call(__e, tmp16808, tmp16814)


__e.TailApply(PrimFunc(symshen_4gc), V5138, tmp16815)
return


}, 1)

tmp16816 := Call(__e, PrimFunc(symshen_4newpv), V5138)


__e.TailApply(tmp16807, tmp16816)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16821 := MakeNative(func(__e *ControlFlow) {
Z5171 := __e.Get(1)
_ = Z5171
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5172 := __e.Get(1)
_ = Z5172
tmp16822 := Call(__e, W5163, W5168)


tmp16823 := Call(__e, tmp16822, Z5171)


__e.TailApply(tmp16823, Z5172)
return


}, 1))
return
}, 1)

__e.TailApply(tmp16778, tmp16821)
return


}, 1)

tmp16824 := PrimTail(W5162)

tmp16825 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16824, V5138)


__e.TailApply(tmp16777, tmp16825)
return


}, 1)

tmp16826 := PrimHead(W5162)

__e.TailApply(tmp16776, tmp16826)
return


} else {
tmp16843 := Call(__e, PrimFunc(symshen_4pvar_2), W5162)


if True == tmp16843 {
tmp16827 := MakeNative(func(__e *ControlFlow) {
W5183 := __e.Get(1)
_ = W5183
tmp16828 := MakeNative(func(__e *ControlFlow) {
W5184 := __e.Get(1)
_ = W5184
tmp16829 := MakeNative(func(__e *ControlFlow) {
W5185 := __e.Get(1)
_ = W5185
tmp16830 := PrimCons(W5185, Nil)

tmp16831 := PrimCons(W5184, tmp16830)

tmp16832 := PrimCons(W5183, tmp16831)

tmp16833 := MakeNative(func(__e *ControlFlow) {
tmp16834 := Call(__e, W5163, W5183)


tmp16835 := Call(__e, tmp16834, W5184)


__e.TailApply(tmp16835, W5185)
return


}, 0)

tmp16836 := Call(__e, PrimFunc(symshen_4bind_b), W5162, tmp16832, V5138, tmp16833)


__e.TailApply(PrimFunc(symshen_4gc), V5138, tmp16836)
return


}, 1)

tmp16837 := Call(__e, PrimFunc(symshen_4newpv), V5138)


tmp16838 := Call(__e, tmp16829, tmp16837)


__e.TailApply(PrimFunc(symshen_4gc), V5138, tmp16838)
return


}, 1)

tmp16839 := Call(__e, PrimFunc(symshen_4newpv), V5138)


tmp16840 := Call(__e, tmp16828, tmp16839)


__e.TailApply(PrimFunc(symshen_4gc), V5138, tmp16840)
return


}, 1)

tmp16841 := Call(__e, PrimFunc(symshen_4newpv), V5138)


__e.TailApply(tmp16827, tmp16841)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16846 := MakeNative(func(__e *ControlFlow) {
Z5164 := __e.Get(1)
_ = Z5164
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5165 := __e.Get(1)
_ = Z5165
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5166 := __e.Get(1)
_ = Z5166
tmp16847 := MakeNative(func(__e *ControlFlow) {
W5167 := __e.Get(1)
_ = W5167
tmp16848 := Call(__e, W5157, Z5164)


tmp16849 := Call(__e, tmp16848, Z5165)


tmp16850 := Call(__e, tmp16849, Z5166)


__e.TailApply(tmp16850, W5167)
return


}, 1)

tmp16851 := PrimTail(W5156)

__e.TailApply(tmp16847, tmp16851)
return


}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp16775, tmp16846)
return


}, 1)

tmp16852 := PrimHead(W5156)

tmp16853 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16852, V5138)


__e.TailApply(tmp16774, tmp16853)
return


} else {
tmp16875 := Call(__e, PrimFunc(symshen_4pvar_2), W5156)


if True == tmp16875 {
tmp16854 := MakeNative(func(__e *ControlFlow) {
W5186 := __e.Get(1)
_ = W5186
tmp16855 := MakeNative(func(__e *ControlFlow) {
W5187 := __e.Get(1)
_ = W5187
tmp16856 := MakeNative(func(__e *ControlFlow) {
W5188 := __e.Get(1)
_ = W5188
tmp16857 := MakeNative(func(__e *ControlFlow) {
W5189 := __e.Get(1)
_ = W5189
tmp16858 := PrimCons(W5188, Nil)

tmp16859 := PrimCons(W5187, tmp16858)

tmp16860 := PrimCons(W5186, tmp16859)

tmp16861 := PrimCons(tmp16860, W5189)

tmp16862 := MakeNative(func(__e *ControlFlow) {
tmp16863 := Call(__e, W5157, W5186)


tmp16864 := Call(__e, tmp16863, W5187)


tmp16865 := Call(__e, tmp16864, W5188)


__e.TailApply(tmp16865, W5189)
return


}, 0)

tmp16866 := Call(__e, PrimFunc(symshen_4bind_b), W5156, tmp16861, V5138, tmp16862)


__e.TailApply(PrimFunc(symshen_4gc), V5138, tmp16866)
return


}, 1)

tmp16867 := Call(__e, PrimFunc(symshen_4newpv), V5138)


tmp16868 := Call(__e, tmp16857, tmp16867)


__e.TailApply(PrimFunc(symshen_4gc), V5138, tmp16868)
return


}, 1)

tmp16869 := Call(__e, PrimFunc(symshen_4newpv), V5138)


tmp16870 := Call(__e, tmp16856, tmp16869)


__e.TailApply(PrimFunc(symshen_4gc), V5138, tmp16870)
return


}, 1)

tmp16871 := Call(__e, PrimFunc(symshen_4newpv), V5138)


tmp16872 := Call(__e, tmp16855, tmp16871)


__e.TailApply(PrimFunc(symshen_4gc), V5138, tmp16872)
return


}, 1)

tmp16873 := Call(__e, PrimFunc(symshen_4newpv), V5138)


__e.TailApply(tmp16854, tmp16873)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16878 := MakeNative(func(__e *ControlFlow) {
Z5158 := __e.Get(1)
_ = Z5158
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5159 := __e.Get(1)
_ = Z5159
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5160 := __e.Get(1)
_ = Z5160
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5161 := __e.Get(1)
_ = Z5161
tmp16879 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16879

tmp16880 := MakeNative(func(__e *ControlFlow) {
tmp16881 := MakeNative(func(__e *ControlFlow) {
tmp16882 := PrimIntern(MakeString(":"))

tmp16883 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4myassume), W5148, W5154, Z5161, V5138, V5139, V5140, V5141)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5159, tmp16882, V5138, V5139, V5140, tmp16883)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W5147, Z5158, V5138, V5139, V5140, tmp16881)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W5150, Z5160, V5138, V5139, V5140, tmp16880)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp16773, tmp16878)
return


}, 1)

tmp16884 := Call(__e, PrimFunc(symshen_4lazyderef), V5137, V5138)


__e.TailApply(tmp16772, tmp16884)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16887 := PrimTail(W5153)

tmp16888 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16887, V5138)


__e.TailApply(tmp16771, tmp16888)
return


}, 1)

tmp16889 := PrimHead(W5153)

__e.TailApply(tmp16770, tmp16889)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16892 := PrimTail(W5151)

tmp16893 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16892, V5138)


__e.TailApply(tmp16769, tmp16893)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16896 := PrimHead(W5151)

tmp16897 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16896, V5138)


__e.TailApply(tmp16768, tmp16897)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16900 := PrimTail(W5149)

tmp16901 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16900, V5138)


__e.TailApply(tmp16767, tmp16901)
return


}, 1)

tmp16902 := PrimHead(W5149)

__e.TailApply(tmp16766, tmp16902)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16905 := Call(__e, PrimFunc(symshen_4lazyderef), V5136, V5138)


__e.TailApply(tmp16765, tmp16905)
return


}, 1)

tmp16906 := PrimTail(W5146)

__e.TailApply(tmp16764, tmp16906)
return


}, 1)

tmp16907 := PrimHead(W5146)

__e.TailApply(tmp16763, tmp16907)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16910 := Call(__e, PrimFunc(symshen_4lazyderef), V5135, V5138)


__e.TailApply(tmp16762, tmp16910)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5142)
return
}


}, 1)

tmp16930 := Call(__e, PrimFunc(symshen_4unlocked_2), V5139)


var ifres16915 Obj

if True == tmp16930 {
tmp16916 := MakeNative(func(__e *ControlFlow) {
W5143 := __e.Get(1)
_ = W5143
tmp16927 := PrimEqual(W5143, Nil)

if True == tmp16927 {
tmp16917 := MakeNative(func(__e *ControlFlow) {
W5144 := __e.Get(1)
_ = W5144
tmp16918 := MakeNative(func(__e *ControlFlow) {
W5145 := __e.Get(1)
_ = W5145
tmp16922 := PrimEqual(W5144, Nil)

if True == tmp16922 {
__e.TailApply(PrimFunc(symthaw), W5145)
return
} else {
tmp16920 := Call(__e, PrimFunc(symshen_4pvar_2), W5144)


if True == tmp16920 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5144, Nil, V5138, W5145)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp16923 := MakeNative(func(__e *ControlFlow) {
tmp16924 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16924

__e.TailApply(PrimFunc(symthaw), V5141)
return


}, 0)

__e.TailApply(tmp16918, tmp16923)
return


}, 1)

tmp16925 := Call(__e, PrimFunc(symshen_4lazyderef), V5137, V5138)


__e.TailApply(tmp16917, tmp16925)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16928 := Call(__e, PrimFunc(symshen_4lazyderef), V5135, V5138)


tmp16929 := Call(__e, tmp16916, tmp16928)


ifres16915 = tmp16929


} else {
ifres16915 = False


}

__e.TailApply(tmp16761, ifres16915)
return


}, 7)

tmp16931 := Call(__e, ns2_1set, symshen_4myassume, tmp16760)


_ = tmp16931

tmp16932 := MakeNative(func(__e *ControlFlow) {
V5192 := __e.Get(1)
_ = V5192
tmp16955 := PrimEqual(Nil, V5192)

if True == tmp16955 {
__e.Return(Nil)
return
} else {
tmp16953 := PrimIsPair(V5192)

var ifres16949 Obj

if True == tmp16953 {
tmp16951 := PrimHead(V5192)

tmp16952 := PrimIsPair(tmp16951)

var ifres16950 Obj

if True == tmp16952 {
ifres16950 = True


} else {
ifres16950 = False


}

ifres16949 = ifres16950


} else {
ifres16949 = False


}

if True == ifres16949 {
tmp16933 := PrimHead(V5192)

tmp16934 := PrimTail(V5192)

tmp16935 := Call(__e, PrimFunc(symappend), tmp16933, tmp16934)


__e.TailApply(PrimFunc(symshen_4freshterms), tmp16935)
return


} else {
tmp16947 := PrimIsPair(V5192)

var ifres16943 Obj

if True == tmp16947 {
tmp16945 := PrimHead(V5192)

tmp16946 := Call(__e, PrimFunc(symshen_4freshterm_2), tmp16945)


var ifres16944 Obj

if True == tmp16946 {
ifres16944 = True


} else {
ifres16944 = False


}

ifres16943 = ifres16944


} else {
ifres16943 = False


}

if True == ifres16943 {
tmp16936 := PrimHead(V5192)

tmp16937 := PrimTail(V5192)

tmp16938 := Call(__e, PrimFunc(symshen_4freshterms), tmp16937)


__e.TailApply(PrimFunc(symadjoin), tmp16936, tmp16938)
return


} else {
tmp16941 := PrimIsPair(V5192)

if True == tmp16941 {
tmp16939 := PrimTail(V5192)

__e.TailApply(PrimFunc(symshen_4freshterms), tmp16939)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.freshterms")))
return
}


}


}


}


}, 1)

tmp16956 := Call(__e, ns2_1set, symshen_4freshterms, tmp16932)


_ = tmp16956

tmp16957 := MakeNative(func(__e *ControlFlow) {
V5193 := __e.Get(1)
_ = V5193
V5194 := __e.Get(2)
_ = V5194
V5195 := __e.Get(3)
_ = V5195
V5196 := __e.Get(4)
_ = V5196
V5197 := __e.Get(5)
_ = V5197
V5198 := __e.Get(6)
_ = V5198
tmp16958 := MakeNative(func(__e *ControlFlow) {
W5199 := __e.Get(1)
_ = W5199
tmp17082 := PrimEqual(W5199, False)

if True == tmp17082 {
tmp17080 := Call(__e, PrimFunc(symshen_4unlocked_2), V5196)


if True == tmp17080 {
tmp16959 := MakeNative(func(__e *ControlFlow) {
W5203 := __e.Get(1)
_ = W5203
tmp17077 := PrimIsPair(W5203)

if True == tmp17077 {
tmp16960 := MakeNative(func(__e *ControlFlow) {
W5204 := __e.Get(1)
_ = W5204
tmp16961 := MakeNative(func(__e *ControlFlow) {
W5205 := __e.Get(1)
_ = W5205
tmp16962 := MakeNative(func(__e *ControlFlow) {
W5206 := __e.Get(1)
_ = W5206
tmp16963 := MakeNative(func(__e *ControlFlow) {
W5207 := __e.Get(1)
_ = W5207
tmp17067 := PrimIsPair(W5206)

if True == tmp17067 {
tmp16964 := MakeNative(func(__e *ControlFlow) {
W5212 := __e.Get(1)
_ = W5212
tmp16965 := MakeNative(func(__e *ControlFlow) {
W5213 := __e.Get(1)
_ = W5213
tmp17035 := PrimIsPair(W5212)

if True == tmp17035 {
tmp16966 := MakeNative(func(__e *ControlFlow) {
W5218 := __e.Get(1)
_ = W5218
tmp16967 := MakeNative(func(__e *ControlFlow) {
W5219 := __e.Get(1)
_ = W5219
tmp16968 := MakeNative(func(__e *ControlFlow) {
W5220 := __e.Get(1)
_ = W5220
tmp17010 := PrimIsPair(W5219)

if True == tmp17010 {
tmp16969 := MakeNative(func(__e *ControlFlow) {
W5223 := __e.Get(1)
_ = W5223
tmp16970 := MakeNative(func(__e *ControlFlow) {
W5224 := __e.Get(1)
_ = W5224
tmp16971 := MakeNative(func(__e *ControlFlow) {
W5225 := __e.Get(1)
_ = W5225
tmp16991 := PrimIsPair(W5224)

if True == tmp16991 {
tmp16972 := MakeNative(func(__e *ControlFlow) {
W5227 := __e.Get(1)
_ = W5227
tmp16973 := MakeNative(func(__e *ControlFlow) {
W5228 := __e.Get(1)
_ = W5228
tmp16974 := MakeNative(func(__e *ControlFlow) {
W5229 := __e.Get(1)
_ = W5229
tmp16978 := PrimEqual(W5228, Nil)

if True == tmp16978 {
__e.TailApply(PrimFunc(symthaw), W5229)
return
} else {
tmp16976 := Call(__e, PrimFunc(symshen_4pvar_2), W5228)


if True == tmp16976 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5228, Nil, V5195, W5229)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp16979 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5225, W5227)
return
}, 0)

__e.TailApply(tmp16974, tmp16979)
return


}, 1)

tmp16980 := PrimTail(W5224)

tmp16981 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16980, V5195)


__e.TailApply(tmp16973, tmp16981)
return


}, 1)

tmp16982 := PrimHead(W5224)

__e.TailApply(tmp16972, tmp16982)
return


} else {
tmp16989 := Call(__e, PrimFunc(symshen_4pvar_2), W5224)


if True == tmp16989 {
tmp16983 := MakeNative(func(__e *ControlFlow) {
W5230 := __e.Get(1)
_ = W5230
tmp16984 := PrimCons(W5230, Nil)

tmp16985 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5225, W5230)
return
}, 0)

tmp16986 := Call(__e, PrimFunc(symshen_4bind_b), W5224, tmp16984, V5195, tmp16985)


__e.TailApply(PrimFunc(symshen_4gc), V5195, tmp16986)
return


}, 1)

tmp16987 := Call(__e, PrimFunc(symshen_4newpv), V5195)


__e.TailApply(tmp16983, tmp16987)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16992 := MakeNative(func(__e *ControlFlow) {
Z5226 := __e.Get(1)
_ = Z5226
tmp16993 := Call(__e, W5220, W5223)


__e.TailApply(tmp16993, Z5226)
return


}, 1)

__e.TailApply(tmp16971, tmp16992)
return


}, 1)

tmp16994 := PrimTail(W5219)

tmp16995 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16994, V5195)


__e.TailApply(tmp16970, tmp16995)
return


}, 1)

tmp16996 := PrimHead(W5219)

__e.TailApply(tmp16969, tmp16996)
return


} else {
tmp17008 := Call(__e, PrimFunc(symshen_4pvar_2), W5219)


if True == tmp17008 {
tmp16997 := MakeNative(func(__e *ControlFlow) {
W5231 := __e.Get(1)
_ = W5231
tmp16998 := MakeNative(func(__e *ControlFlow) {
W5232 := __e.Get(1)
_ = W5232
tmp16999 := PrimCons(W5232, Nil)

tmp17000 := PrimCons(W5231, tmp16999)

tmp17001 := MakeNative(func(__e *ControlFlow) {
tmp17002 := Call(__e, W5220, W5231)


__e.TailApply(tmp17002, W5232)
return


}, 0)

tmp17003 := Call(__e, PrimFunc(symshen_4bind_b), W5219, tmp17000, V5195, tmp17001)


__e.TailApply(PrimFunc(symshen_4gc), V5195, tmp17003)
return


}, 1)

tmp17004 := Call(__e, PrimFunc(symshen_4newpv), V5195)


tmp17005 := Call(__e, tmp16998, tmp17004)


__e.TailApply(PrimFunc(symshen_4gc), V5195, tmp17005)
return


}, 1)

tmp17006 := Call(__e, PrimFunc(symshen_4newpv), V5195)


__e.TailApply(tmp16997, tmp17006)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17011 := MakeNative(func(__e *ControlFlow) {
Z5221 := __e.Get(1)
_ = Z5221
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5222 := __e.Get(1)
_ = Z5222
tmp17012 := Call(__e, W5213, W5218)


tmp17013 := Call(__e, tmp17012, Z5221)


__e.TailApply(tmp17013, Z5222)
return


}, 1))
return
}, 1)

__e.TailApply(tmp16968, tmp17011)
return


}, 1)

tmp17014 := PrimTail(W5212)

tmp17015 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17014, V5195)


__e.TailApply(tmp16967, tmp17015)
return


}, 1)

tmp17016 := PrimHead(W5212)

__e.TailApply(tmp16966, tmp17016)
return


} else {
tmp17033 := Call(__e, PrimFunc(symshen_4pvar_2), W5212)


if True == tmp17033 {
tmp17017 := MakeNative(func(__e *ControlFlow) {
W5233 := __e.Get(1)
_ = W5233
tmp17018 := MakeNative(func(__e *ControlFlow) {
W5234 := __e.Get(1)
_ = W5234
tmp17019 := MakeNative(func(__e *ControlFlow) {
W5235 := __e.Get(1)
_ = W5235
tmp17020 := PrimCons(W5235, Nil)

tmp17021 := PrimCons(W5234, tmp17020)

tmp17022 := PrimCons(W5233, tmp17021)

tmp17023 := MakeNative(func(__e *ControlFlow) {
tmp17024 := Call(__e, W5213, W5233)


tmp17025 := Call(__e, tmp17024, W5234)


__e.TailApply(tmp17025, W5235)
return


}, 0)

tmp17026 := Call(__e, PrimFunc(symshen_4bind_b), W5212, tmp17022, V5195, tmp17023)


__e.TailApply(PrimFunc(symshen_4gc), V5195, tmp17026)
return


}, 1)

tmp17027 := Call(__e, PrimFunc(symshen_4newpv), V5195)


tmp17028 := Call(__e, tmp17019, tmp17027)


__e.TailApply(PrimFunc(symshen_4gc), V5195, tmp17028)
return


}, 1)

tmp17029 := Call(__e, PrimFunc(symshen_4newpv), V5195)


tmp17030 := Call(__e, tmp17018, tmp17029)


__e.TailApply(PrimFunc(symshen_4gc), V5195, tmp17030)
return


}, 1)

tmp17031 := Call(__e, PrimFunc(symshen_4newpv), V5195)


__e.TailApply(tmp17017, tmp17031)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17036 := MakeNative(func(__e *ControlFlow) {
Z5214 := __e.Get(1)
_ = Z5214
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5215 := __e.Get(1)
_ = Z5215
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5216 := __e.Get(1)
_ = Z5216
tmp17037 := MakeNative(func(__e *ControlFlow) {
W5217 := __e.Get(1)
_ = W5217
tmp17038 := Call(__e, W5207, Z5214)


tmp17039 := Call(__e, tmp17038, Z5215)


tmp17040 := Call(__e, tmp17039, Z5216)


__e.TailApply(tmp17040, W5217)
return


}, 1)

tmp17041 := PrimTail(W5206)

__e.TailApply(tmp17037, tmp17041)
return


}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp16965, tmp17036)
return


}, 1)

tmp17042 := PrimHead(W5206)

tmp17043 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17042, V5195)


__e.TailApply(tmp16964, tmp17043)
return


} else {
tmp17065 := Call(__e, PrimFunc(symshen_4pvar_2), W5206)


if True == tmp17065 {
tmp17044 := MakeNative(func(__e *ControlFlow) {
W5236 := __e.Get(1)
_ = W5236
tmp17045 := MakeNative(func(__e *ControlFlow) {
W5237 := __e.Get(1)
_ = W5237
tmp17046 := MakeNative(func(__e *ControlFlow) {
W5238 := __e.Get(1)
_ = W5238
tmp17047 := MakeNative(func(__e *ControlFlow) {
W5239 := __e.Get(1)
_ = W5239
tmp17048 := PrimCons(W5238, Nil)

tmp17049 := PrimCons(W5237, tmp17048)

tmp17050 := PrimCons(W5236, tmp17049)

tmp17051 := PrimCons(tmp17050, W5239)

tmp17052 := MakeNative(func(__e *ControlFlow) {
tmp17053 := Call(__e, W5207, W5236)


tmp17054 := Call(__e, tmp17053, W5237)


tmp17055 := Call(__e, tmp17054, W5238)


__e.TailApply(tmp17055, W5239)
return


}, 0)

tmp17056 := Call(__e, PrimFunc(symshen_4bind_b), W5206, tmp17051, V5195, tmp17052)


__e.TailApply(PrimFunc(symshen_4gc), V5195, tmp17056)
return


}, 1)

tmp17057 := Call(__e, PrimFunc(symshen_4newpv), V5195)


tmp17058 := Call(__e, tmp17047, tmp17057)


__e.TailApply(PrimFunc(symshen_4gc), V5195, tmp17058)
return


}, 1)

tmp17059 := Call(__e, PrimFunc(symshen_4newpv), V5195)


tmp17060 := Call(__e, tmp17046, tmp17059)


__e.TailApply(PrimFunc(symshen_4gc), V5195, tmp17060)
return


}, 1)

tmp17061 := Call(__e, PrimFunc(symshen_4newpv), V5195)


tmp17062 := Call(__e, tmp17045, tmp17061)


__e.TailApply(PrimFunc(symshen_4gc), V5195, tmp17062)
return


}, 1)

tmp17063 := Call(__e, PrimFunc(symshen_4newpv), V5195)


__e.TailApply(tmp17044, tmp17063)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17068 := MakeNative(func(__e *ControlFlow) {
Z5208 := __e.Get(1)
_ = Z5208
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5209 := __e.Get(1)
_ = Z5209
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5210 := __e.Get(1)
_ = Z5210
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5211 := __e.Get(1)
_ = Z5211
tmp17069 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17069

tmp17070 := MakeNative(func(__e *ControlFlow) {
tmp17071 := PrimIntern(MakeString(":"))

tmp17072 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4p_1hyps), W5205, Z5211, V5195, V5196, V5197, V5198)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5209, tmp17071, V5195, V5196, V5197, tmp17072)
return


}, 0)

__e.TailApply(PrimFunc(symbind), Z5208, W5204, V5195, V5196, V5197, tmp17070)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp16963, tmp17068)
return


}, 1)

tmp17073 := Call(__e, PrimFunc(symshen_4lazyderef), V5194, V5195)


__e.TailApply(tmp16962, tmp17073)
return


}, 1)

tmp17074 := PrimTail(W5203)

__e.TailApply(tmp16961, tmp17074)
return


}, 1)

tmp17075 := PrimHead(W5203)

__e.TailApply(tmp16960, tmp17075)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17078 := Call(__e, PrimFunc(symshen_4lazyderef), V5193, V5195)


__e.TailApply(tmp16959, tmp17078)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5199)
return
}


}, 1)

tmp17098 := Call(__e, PrimFunc(symshen_4unlocked_2), V5196)


var ifres17083 Obj

if True == tmp17098 {
tmp17084 := MakeNative(func(__e *ControlFlow) {
W5200 := __e.Get(1)
_ = W5200
tmp17095 := PrimEqual(W5200, Nil)

if True == tmp17095 {
tmp17085 := MakeNative(func(__e *ControlFlow) {
W5201 := __e.Get(1)
_ = W5201
tmp17086 := MakeNative(func(__e *ControlFlow) {
W5202 := __e.Get(1)
_ = W5202
tmp17090 := PrimEqual(W5201, Nil)

if True == tmp17090 {
__e.TailApply(PrimFunc(symthaw), W5202)
return
} else {
tmp17088 := Call(__e, PrimFunc(symshen_4pvar_2), W5201)


if True == tmp17088 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5201, Nil, V5195, W5202)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp17091 := MakeNative(func(__e *ControlFlow) {
tmp17092 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17092

__e.TailApply(PrimFunc(symthaw), V5198)
return


}, 0)

__e.TailApply(tmp17086, tmp17091)
return


}, 1)

tmp17093 := Call(__e, PrimFunc(symshen_4lazyderef), V5194, V5195)


__e.TailApply(tmp17085, tmp17093)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17096 := Call(__e, PrimFunc(symshen_4lazyderef), V5193, V5195)


tmp17097 := Call(__e, tmp17084, tmp17096)


ifres17083 = tmp17097


} else {
ifres17083 = False


}

__e.TailApply(tmp16958, ifres17083)
return


}, 6)

tmp17099 := Call(__e, ns2_1set, symshen_4p_1hyps, tmp16957)


_ = tmp17099

tmp17100 := MakeNative(func(__e *ControlFlow) {
V5240 := __e.Get(1)
_ = V5240
V5241 := __e.Get(2)
_ = V5241
V5242 := __e.Get(3)
_ = V5242
V5243 := __e.Get(4)
_ = V5243
V5244 := __e.Get(5)
_ = V5244
V5245 := __e.Get(6)
_ = V5245
V5246 := __e.Get(7)
_ = V5246
tmp17101 := MakeNative(func(__e *ControlFlow) {
W5247 := __e.Get(1)
_ = W5247
tmp17102 := MakeNative(func(__e *ControlFlow) {
W5248 := __e.Get(1)
_ = W5248
tmp17112 := PrimEqual(W5248, False)

if True == tmp17112 {
tmp17103 := MakeNative(func(__e *ControlFlow) {
W5257 := __e.Get(1)
_ = W5257
tmp17105 := PrimEqual(W5257, False)

if True == tmp17105 {
__e.TailApply(PrimFunc(symshen_4unlock), V5244, W5247)
return
} else {
__e.Return(W5257)
return
}


}, 1)

tmp17110 := Call(__e, PrimFunc(symshen_4unlocked_2), V5244)


var ifres17106 Obj

if True == tmp17110 {
tmp17107 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17107

tmp17108 := Call(__e, PrimFunc(symshen_4curry), V5240)


tmp17109 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp17108, V5241, V5242, V5243, V5244, W5247, V5246)


ifres17106 = tmp17109


} else {
ifres17106 = False


}

__e.TailApply(tmp17103, ifres17106)
return


} else {
__e.Return(W5248)
return
}


}, 1)

tmp17157 := Call(__e, PrimFunc(symshen_4unlocked_2), V5244)


var ifres17113 Obj

if True == tmp17157 {
tmp17114 := MakeNative(func(__e *ControlFlow) {
W5249 := __e.Get(1)
_ = W5249
tmp17154 := PrimIsPair(W5249)

if True == tmp17154 {
tmp17115 := MakeNative(func(__e *ControlFlow) {
W5250 := __e.Get(1)
_ = W5250
tmp17150 := PrimEqual(W5250, symwhere)

if True == tmp17150 {
tmp17116 := MakeNative(func(__e *ControlFlow) {
W5251 := __e.Get(1)
_ = W5251
tmp17146 := PrimIsPair(W5251)

if True == tmp17146 {
tmp17117 := MakeNative(func(__e *ControlFlow) {
W5252 := __e.Get(1)
_ = W5252
tmp17118 := MakeNative(func(__e *ControlFlow) {
W5253 := __e.Get(1)
_ = W5253
tmp17141 := PrimIsPair(W5253)

if True == tmp17141 {
tmp17119 := MakeNative(func(__e *ControlFlow) {
W5254 := __e.Get(1)
_ = W5254
tmp17120 := MakeNative(func(__e *ControlFlow) {
W5255 := __e.Get(1)
_ = W5255
tmp17136 := PrimEqual(W5255, Nil)

if True == tmp17136 {
tmp17121 := MakeNative(func(__e *ControlFlow) {
W5256 := __e.Get(1)
_ = W5256
tmp17122 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17122

tmp17123 := MakeNative(func(__e *ControlFlow) {
tmp17124 := Call(__e, PrimFunc(symshen_4curry), W5252)


tmp17125 := MakeNative(func(__e *ControlFlow) {
tmp17126 := MakeNative(func(__e *ControlFlow) {
tmp17127 := MakeNative(func(__e *ControlFlow) {
tmp17128 := PrimIntern(MakeString(":"))

tmp17129 := PrimCons(symverified, Nil)

tmp17130 := PrimCons(tmp17128, tmp17129)

tmp17131 := PrimCons(W5256, tmp17130)

tmp17132 := PrimCons(tmp17131, V5242)

__e.TailApply(PrimFunc(symshen_4t_d_1correct), W5254, V5241, tmp17132, V5243, V5244, W5247, V5246)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5243, V5244, W5247, tmp17127)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5256, symboolean, V5242, V5243, V5244, W5247, tmp17126)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5256, tmp17124, V5243, V5244, W5247, tmp17125)
return


}, 0)

tmp17133 := Call(__e, PrimFunc(symshen_4cut), V5243, V5244, W5247, tmp17123)


__e.TailApply(PrimFunc(symshen_4gc), V5243, tmp17133)
return


}, 1)

tmp17134 := Call(__e, PrimFunc(symshen_4newpv), V5243)


__e.TailApply(tmp17121, tmp17134)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17137 := PrimTail(W5253)

tmp17138 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17137, V5243)


__e.TailApply(tmp17120, tmp17138)
return


}, 1)

tmp17139 := PrimHead(W5253)

__e.TailApply(tmp17119, tmp17139)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17142 := PrimTail(W5251)

tmp17143 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17142, V5243)


__e.TailApply(tmp17118, tmp17143)
return


}, 1)

tmp17144 := PrimHead(W5251)

__e.TailApply(tmp17117, tmp17144)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17147 := PrimTail(W5249)

tmp17148 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17147, V5243)


__e.TailApply(tmp17116, tmp17148)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17151 := PrimHead(W5249)

tmp17152 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17151, V5243)


__e.TailApply(tmp17115, tmp17152)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17155 := Call(__e, PrimFunc(symshen_4lazyderef), V5240, V5243)


tmp17156 := Call(__e, tmp17114, tmp17155)


ifres17113 = tmp17156


} else {
ifres17113 = False


}

__e.TailApply(tmp17102, ifres17113)
return


}, 1)

tmp17158 := PrimNumberAdd(V5245, MakeNumber(1))

__e.TailApply(tmp17101, tmp17158)
return


}, 7)

tmp17159 := Call(__e, ns2_1set, symshen_4t_d_1correct, tmp17100)


_ = tmp17159

tmp17160 := MakeNative(func(__e *ControlFlow) {
V5258 := __e.Get(1)
_ = V5258
V5259 := __e.Get(2)
_ = V5259
V5260 := __e.Get(3)
_ = V5260
V5261 := __e.Get(4)
_ = V5261
V5262 := __e.Get(5)
_ = V5262
V5263 := __e.Get(6)
_ = V5263
V5264 := __e.Get(7)
_ = V5264
V5265 := __e.Get(8)
_ = V5265
tmp17161 := MakeNative(func(__e *ControlFlow) {
W5266 := __e.Get(1)
_ = W5266
tmp17203 := PrimEqual(W5266, False)

if True == tmp17203 {
tmp17201 := Call(__e, PrimFunc(symshen_4unlocked_2), V5263)


if True == tmp17201 {
tmp17162 := MakeNative(func(__e *ControlFlow) {
W5268 := __e.Get(1)
_ = W5268
tmp17198 := PrimIsPair(W5268)

if True == tmp17198 {
tmp17163 := MakeNative(func(__e *ControlFlow) {
W5269 := __e.Get(1)
_ = W5269
tmp17164 := MakeNative(func(__e *ControlFlow) {
W5270 := __e.Get(1)
_ = W5270
tmp17165 := MakeNative(func(__e *ControlFlow) {
W5271 := __e.Get(1)
_ = W5271
tmp17193 := PrimIsPair(W5271)

if True == tmp17193 {
tmp17166 := MakeNative(func(__e *ControlFlow) {
W5272 := __e.Get(1)
_ = W5272
tmp17167 := MakeNative(func(__e *ControlFlow) {
W5273 := __e.Get(1)
_ = W5273
tmp17188 := PrimIsPair(W5273)

if True == tmp17188 {
tmp17168 := MakeNative(func(__e *ControlFlow) {
W5274 := __e.Get(1)
_ = W5274
tmp17184 := PrimEqual(W5274, sym_1_1_6)

if True == tmp17184 {
tmp17169 := MakeNative(func(__e *ControlFlow) {
W5275 := __e.Get(1)
_ = W5275
tmp17180 := PrimIsPair(W5275)

if True == tmp17180 {
tmp17170 := MakeNative(func(__e *ControlFlow) {
W5276 := __e.Get(1)
_ = W5276
tmp17171 := MakeNative(func(__e *ControlFlow) {
W5277 := __e.Get(1)
_ = W5277
tmp17175 := PrimEqual(W5277, Nil)

if True == tmp17175 {
tmp17172 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17172

tmp17173 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1integrity), W5270, W5276, V5260, V5261, V5262, V5263, V5264, V5265)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5269, W5272, V5260, V5262, V5263, V5264, tmp17173)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17176 := PrimTail(W5275)

tmp17177 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17176, V5262)


__e.TailApply(tmp17171, tmp17177)
return


}, 1)

tmp17178 := PrimHead(W5275)

__e.TailApply(tmp17170, tmp17178)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17181 := PrimTail(W5273)

tmp17182 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17181, V5262)


__e.TailApply(tmp17169, tmp17182)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17185 := PrimHead(W5273)

tmp17186 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17185, V5262)


__e.TailApply(tmp17168, tmp17186)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17189 := PrimTail(W5271)

tmp17190 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17189, V5262)


__e.TailApply(tmp17167, tmp17190)
return


}, 1)

tmp17191 := PrimHead(W5271)

__e.TailApply(tmp17166, tmp17191)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17194 := Call(__e, PrimFunc(symshen_4lazyderef), V5259, V5262)


__e.TailApply(tmp17165, tmp17194)
return


}, 1)

tmp17195 := PrimTail(W5268)

__e.TailApply(tmp17164, tmp17195)
return


}, 1)

tmp17196 := PrimHead(W5268)

__e.TailApply(tmp17163, tmp17196)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17199 := Call(__e, PrimFunc(symshen_4lazyderef), V5258, V5262)


__e.TailApply(tmp17162, tmp17199)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5266)
return
}


}, 1)

tmp17211 := Call(__e, PrimFunc(symshen_4unlocked_2), V5263)


var ifres17204 Obj

if True == tmp17211 {
tmp17205 := MakeNative(func(__e *ControlFlow) {
W5267 := __e.Get(1)
_ = W5267
tmp17208 := PrimEqual(W5267, Nil)

if True == tmp17208 {
tmp17206 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17206

__e.TailApply(PrimFunc(symis_b), V5259, V5261, V5262, V5263, V5264, V5265)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17209 := Call(__e, PrimFunc(symshen_4lazyderef), V5258, V5262)


tmp17210 := Call(__e, tmp17205, tmp17209)


ifres17204 = tmp17210


} else {
ifres17204 = False


}

__e.TailApply(tmp17161, ifres17204)
return


}, 8)

tmp17212 := Call(__e, ns2_1set, symshen_4t_d_1integrity, tmp17160)


_ = tmp17212

tmp17213 := MakeNative(func(__e *ControlFlow) {
V5278 := __e.Get(1)
_ = V5278
tmp17222 := PrimIsVector(V5278)

if True == tmp17222 {
tmp17219 := PrimIsString(V5278)

tmp17220 := PrimNot(tmp17219)

var ifres17215 Obj

if True == tmp17220 {
tmp17217 := PrimVectorGet(V5278, MakeNumber(0))

tmp17218 := PrimEqual(tmp17217, symshen_4print_1freshterm)

var ifres17216 Obj

if True == tmp17218 {
ifres17216 = True


} else {
ifres17216 = False


}

ifres17215 = ifres17216


} else {
ifres17215 = False


}

if True == ifres17215 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


}, 1)

__e.TailApply(ns2_1set, symshen_4freshterm_2, tmp17213)
return




}, 0)

