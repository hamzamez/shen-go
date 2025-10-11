package main

import . "github.com/tiancaiamao/shen-go/kl"

var SequentMain = MakeNative(func(__e *ControlFlow) {
tmp12100 := MakeNative(func(__e *ControlFlow) {
V3020 := __e.Get(1)
_ = V3020
tmp12101 := MakeNative(func(__e *ControlFlow) {
W3021 := __e.Get(1)
_ = W3021
tmp12103 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3021)


if True == tmp12103 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3021)
return
}


}, 1)

tmp12123 := PrimIsPair(V3020)

var ifres12104 Obj

if True == tmp12123 {
tmp12105 := MakeNative(func(__e *ControlFlow) {
W3022 := __e.Get(1)
_ = W3022
tmp12106 := MakeNative(func(__e *ControlFlow) {
W3023 := __e.Get(1)
_ = W3023
tmp12107 := MakeNative(func(__e *ControlFlow) {
W3024 := __e.Get(1)
_ = W3024
tmp12117 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3024)


if True == tmp12117 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12108 := MakeNative(func(__e *ControlFlow) {
W3025 := __e.Get(1)
_ = W3025
tmp12109 := MakeNative(func(__e *ControlFlow) {
W3026 := __e.Get(1)
_ = W3026
tmp12110 := MakeNative(func(__e *ControlFlow) {
W3027 := __e.Get(1)
_ = W3027
tmp12111 := Call(__e, PrimFunc(symfn), W3022)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), W3022, tmp12111)
return


}, 1)

tmp12112 := Call(__e, PrimFunc(symshen_4rules_1_6prolog), W3022, W3025)


tmp12113 := Call(__e, tmp12110, tmp12112)


__e.TailApply(PrimFunc(symshen_4comb), W3026, tmp12113)
return


}, 1)

tmp12114 := Call(__e, PrimFunc(symshen_4in_1_6), W3024)


__e.TailApply(tmp12109, tmp12114)
return


}, 1)

tmp12115 := Call(__e, PrimFunc(symshen_4_5_1out), W3024)


__e.TailApply(tmp12108, tmp12115)
return


}


}, 1)

tmp12118 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), W3023)


__e.TailApply(tmp12107, tmp12118)
return


}, 1)

tmp12119 := Call(__e, PrimFunc(symtail), V3020)


__e.TailApply(tmp12106, tmp12119)
return


}, 1)

tmp12120 := Call(__e, PrimFunc(symhead), V3020)


tmp12121 := Call(__e, tmp12105, tmp12120)


ifres12104 = tmp12121


} else {
tmp12122 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12104 = tmp12122


}

__e.TailApply(tmp12101, ifres12104)
return


}, 1)

tmp12124 := Call(__e, ns2_1set, symshen_4_5datatype_6, tmp12100)


_ = tmp12124

tmp12125 := MakeNative(func(__e *ControlFlow) {
V3028 := __e.Get(1)
_ = V3028
tmp12126 := MakeNative(func(__e *ControlFlow) {
W3029 := __e.Get(1)
_ = W3029
tmp12145 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3029)


if True == tmp12145 {
tmp12127 := MakeNative(func(__e *ControlFlow) {
W3036 := __e.Get(1)
_ = W3036
tmp12129 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3036)


if True == tmp12129 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3036)
return
}


}, 1)

tmp12130 := MakeNative(func(__e *ControlFlow) {
W3037 := __e.Get(1)
_ = W3037
tmp12141 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3037)


if True == tmp12141 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12131 := MakeNative(func(__e *ControlFlow) {
W3038 := __e.Get(1)
_ = W3038
tmp12132 := MakeNative(func(__e *ControlFlow) {
W3039 := __e.Get(1)
_ = W3039
tmp12137 := Call(__e, PrimFunc(symempty_2), W3038)


var ifres12133 Obj

if True == tmp12137 {
ifres12133 = Nil


} else {
tmp12134 := Call(__e, PrimFunc(symshen_4app), W3038, MakeString("\n ..."), symshen_4r)


tmp12135 := PrimStringConcat(MakeString("datatype syntax error here:\n "), tmp12134)

tmp12136 := PrimSimpleError(tmp12135)

ifres12133 = tmp12136


}

__e.TailApply(PrimFunc(symshen_4comb), W3039, ifres12133)
return


}, 1)

tmp12138 := Call(__e, PrimFunc(symshen_4in_1_6), W3037)


__e.TailApply(tmp12132, tmp12138)
return


}, 1)

tmp12139 := Call(__e, PrimFunc(symshen_4_5_1out), W3037)


__e.TailApply(tmp12131, tmp12139)
return


}


}, 1)

tmp12142 := Call(__e, PrimFunc(sym_5_b_6), V3028)


tmp12143 := Call(__e, tmp12130, tmp12142)


__e.TailApply(tmp12127, tmp12143)
return


} else {
__e.Return(W3029)
return
}


}, 1)

tmp12146 := MakeNative(func(__e *ControlFlow) {
W3030 := __e.Get(1)
_ = W3030
tmp12161 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3030)


if True == tmp12161 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12147 := MakeNative(func(__e *ControlFlow) {
W3031 := __e.Get(1)
_ = W3031
tmp12148 := MakeNative(func(__e *ControlFlow) {
W3032 := __e.Get(1)
_ = W3032
tmp12149 := MakeNative(func(__e *ControlFlow) {
W3033 := __e.Get(1)
_ = W3033
tmp12156 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3033)


if True == tmp12156 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12150 := MakeNative(func(__e *ControlFlow) {
W3034 := __e.Get(1)
_ = W3034
tmp12151 := MakeNative(func(__e *ControlFlow) {
W3035 := __e.Get(1)
_ = W3035
tmp12152 := Call(__e, PrimFunc(symappend), W3031, W3034)


__e.TailApply(PrimFunc(symshen_4comb), W3035, tmp12152)
return


}, 1)

tmp12153 := Call(__e, PrimFunc(symshen_4in_1_6), W3033)


__e.TailApply(tmp12151, tmp12153)
return


}, 1)

tmp12154 := Call(__e, PrimFunc(symshen_4_5_1out), W3033)


__e.TailApply(tmp12150, tmp12154)
return


}


}, 1)

tmp12157 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), W3032)


__e.TailApply(tmp12149, tmp12157)
return


}, 1)

tmp12158 := Call(__e, PrimFunc(symshen_4in_1_6), W3030)


__e.TailApply(tmp12148, tmp12158)
return


}, 1)

tmp12159 := Call(__e, PrimFunc(symshen_4_5_1out), W3030)


__e.TailApply(tmp12147, tmp12159)
return


}


}, 1)

tmp12162 := Call(__e, PrimFunc(symshen_4_5datatype_1rule_6), V3028)


tmp12163 := Call(__e, tmp12146, tmp12162)


__e.TailApply(tmp12126, tmp12163)
return


}, 1)

tmp12164 := Call(__e, ns2_1set, symshen_4_5datatype_1rules_6, tmp12125)


_ = tmp12164

tmp12165 := MakeNative(func(__e *ControlFlow) {
V3040 := __e.Get(1)
_ = V3040
tmp12166 := MakeNative(func(__e *ControlFlow) {
W3041 := __e.Get(1)
_ = W3041
tmp12180 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3041)


if True == tmp12180 {
tmp12167 := MakeNative(func(__e *ControlFlow) {
W3045 := __e.Get(1)
_ = W3045
tmp12169 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3045)


if True == tmp12169 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3045)
return
}


}, 1)

tmp12170 := MakeNative(func(__e *ControlFlow) {
W3046 := __e.Get(1)
_ = W3046
tmp12176 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3046)


if True == tmp12176 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12171 := MakeNative(func(__e *ControlFlow) {
W3047 := __e.Get(1)
_ = W3047
tmp12172 := MakeNative(func(__e *ControlFlow) {
W3048 := __e.Get(1)
_ = W3048
__e.TailApply(PrimFunc(symshen_4comb), W3048, W3047)
return
}, 1)

tmp12173 := Call(__e, PrimFunc(symshen_4in_1_6), W3046)


__e.TailApply(tmp12172, tmp12173)
return


}, 1)

tmp12174 := Call(__e, PrimFunc(symshen_4_5_1out), W3046)


__e.TailApply(tmp12171, tmp12174)
return


}


}, 1)

tmp12177 := Call(__e, PrimFunc(symshen_4_5double_6), V3040)


tmp12178 := Call(__e, tmp12170, tmp12177)


__e.TailApply(tmp12167, tmp12178)
return


} else {
__e.Return(W3041)
return
}


}, 1)

tmp12181 := MakeNative(func(__e *ControlFlow) {
W3042 := __e.Get(1)
_ = W3042
tmp12187 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3042)


if True == tmp12187 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12182 := MakeNative(func(__e *ControlFlow) {
W3043 := __e.Get(1)
_ = W3043
tmp12183 := MakeNative(func(__e *ControlFlow) {
W3044 := __e.Get(1)
_ = W3044
__e.TailApply(PrimFunc(symshen_4comb), W3044, W3043)
return
}, 1)

tmp12184 := Call(__e, PrimFunc(symshen_4in_1_6), W3042)


__e.TailApply(tmp12183, tmp12184)
return


}, 1)

tmp12185 := Call(__e, PrimFunc(symshen_4_5_1out), W3042)


__e.TailApply(tmp12182, tmp12185)
return


}


}, 1)

tmp12188 := Call(__e, PrimFunc(symshen_4_5single_6), V3040)


tmp12189 := Call(__e, tmp12181, tmp12188)


__e.TailApply(tmp12166, tmp12189)
return


}, 1)

tmp12190 := Call(__e, ns2_1set, symshen_4_5datatype_1rule_6, tmp12165)


_ = tmp12190

tmp12191 := MakeNative(func(__e *ControlFlow) {
V3049 := __e.Get(1)
_ = V3049
tmp12192 := MakeNative(func(__e *ControlFlow) {
W3050 := __e.Get(1)
_ = W3050
tmp12194 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3050)


if True == tmp12194 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3050)
return
}


}, 1)

tmp12195 := MakeNative(func(__e *ControlFlow) {
W3051 := __e.Get(1)
_ = W3051
tmp12233 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3051)


if True == tmp12233 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12196 := MakeNative(func(__e *ControlFlow) {
W3052 := __e.Get(1)
_ = W3052
tmp12197 := MakeNative(func(__e *ControlFlow) {
W3053 := __e.Get(1)
_ = W3053
tmp12198 := MakeNative(func(__e *ControlFlow) {
W3054 := __e.Get(1)
_ = W3054
tmp12228 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3054)


if True == tmp12228 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12199 := MakeNative(func(__e *ControlFlow) {
W3055 := __e.Get(1)
_ = W3055
tmp12200 := MakeNative(func(__e *ControlFlow) {
W3056 := __e.Get(1)
_ = W3056
tmp12201 := MakeNative(func(__e *ControlFlow) {
W3057 := __e.Get(1)
_ = W3057
tmp12223 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3057)


if True == tmp12223 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12202 := MakeNative(func(__e *ControlFlow) {
W3058 := __e.Get(1)
_ = W3058
tmp12203 := MakeNative(func(__e *ControlFlow) {
W3059 := __e.Get(1)
_ = W3059
tmp12219 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3059)


if True == tmp12219 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12204 := MakeNative(func(__e *ControlFlow) {
W3060 := __e.Get(1)
_ = W3060
tmp12205 := MakeNative(func(__e *ControlFlow) {
W3061 := __e.Get(1)
_ = W3061
tmp12206 := MakeNative(func(__e *ControlFlow) {
W3062 := __e.Get(1)
_ = W3062
tmp12214 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3062)


if True == tmp12214 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12207 := MakeNative(func(__e *ControlFlow) {
W3063 := __e.Get(1)
_ = W3063
tmp12208 := PrimCons(W3060, Nil)

tmp12209 := PrimCons(W3055, tmp12208)

tmp12210 := PrimCons(W3052, tmp12209)

tmp12211 := PrimCons(tmp12210, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W3063, tmp12211)
return


}, 1)

tmp12212 := Call(__e, PrimFunc(symshen_4in_1_6), W3062)


__e.TailApply(tmp12207, tmp12212)
return


}


}, 1)

tmp12215 := Call(__e, PrimFunc(symshen_4_5sc_6), W3061)


__e.TailApply(tmp12206, tmp12215)
return


}, 1)

tmp12216 := Call(__e, PrimFunc(symshen_4in_1_6), W3059)


__e.TailApply(tmp12205, tmp12216)
return


}, 1)

tmp12217 := Call(__e, PrimFunc(symshen_4_5_1out), W3059)


__e.TailApply(tmp12204, tmp12217)
return


}


}, 1)

tmp12220 := Call(__e, PrimFunc(symshen_4_5conc_6), W3058)


__e.TailApply(tmp12203, tmp12220)
return


}, 1)

tmp12221 := Call(__e, PrimFunc(symshen_4in_1_6), W3057)


__e.TailApply(tmp12202, tmp12221)
return


}


}, 1)

tmp12224 := Call(__e, PrimFunc(symshen_4_5sng_6), W3056)


__e.TailApply(tmp12201, tmp12224)
return


}, 1)

tmp12225 := Call(__e, PrimFunc(symshen_4in_1_6), W3054)


__e.TailApply(tmp12200, tmp12225)
return


}, 1)

tmp12226 := Call(__e, PrimFunc(symshen_4_5_1out), W3054)


__e.TailApply(tmp12199, tmp12226)
return


}


}, 1)

tmp12229 := Call(__e, PrimFunc(symshen_4_5prems_6), W3053)


__e.TailApply(tmp12198, tmp12229)
return


}, 1)

tmp12230 := Call(__e, PrimFunc(symshen_4in_1_6), W3051)


__e.TailApply(tmp12197, tmp12230)
return


}, 1)

tmp12231 := Call(__e, PrimFunc(symshen_4_5_1out), W3051)


__e.TailApply(tmp12196, tmp12231)
return


}


}, 1)

tmp12234 := Call(__e, PrimFunc(symshen_4_5sides_6), V3049)


tmp12235 := Call(__e, tmp12195, tmp12234)


__e.TailApply(tmp12192, tmp12235)
return


}, 1)

tmp12236 := Call(__e, ns2_1set, symshen_4_5single_6, tmp12191)


_ = tmp12236

tmp12237 := MakeNative(func(__e *ControlFlow) {
V3064 := __e.Get(1)
_ = V3064
tmp12238 := MakeNative(func(__e *ControlFlow) {
W3065 := __e.Get(1)
_ = W3065
tmp12240 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3065)


if True == tmp12240 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3065)
return
}


}, 1)

tmp12241 := MakeNative(func(__e *ControlFlow) {
W3066 := __e.Get(1)
_ = W3066
tmp12278 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3066)


if True == tmp12278 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12242 := MakeNative(func(__e *ControlFlow) {
W3067 := __e.Get(1)
_ = W3067
tmp12243 := MakeNative(func(__e *ControlFlow) {
W3068 := __e.Get(1)
_ = W3068
tmp12244 := MakeNative(func(__e *ControlFlow) {
W3069 := __e.Get(1)
_ = W3069
tmp12273 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3069)


if True == tmp12273 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12245 := MakeNative(func(__e *ControlFlow) {
W3070 := __e.Get(1)
_ = W3070
tmp12246 := MakeNative(func(__e *ControlFlow) {
W3071 := __e.Get(1)
_ = W3071
tmp12247 := MakeNative(func(__e *ControlFlow) {
W3072 := __e.Get(1)
_ = W3072
tmp12268 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3072)


if True == tmp12268 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12248 := MakeNative(func(__e *ControlFlow) {
W3073 := __e.Get(1)
_ = W3073
tmp12249 := MakeNative(func(__e *ControlFlow) {
W3074 := __e.Get(1)
_ = W3074
tmp12264 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3074)


if True == tmp12264 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12250 := MakeNative(func(__e *ControlFlow) {
W3075 := __e.Get(1)
_ = W3075
tmp12251 := MakeNative(func(__e *ControlFlow) {
W3076 := __e.Get(1)
_ = W3076
tmp12252 := MakeNative(func(__e *ControlFlow) {
W3077 := __e.Get(1)
_ = W3077
tmp12259 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3077)


if True == tmp12259 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12253 := MakeNative(func(__e *ControlFlow) {
W3078 := __e.Get(1)
_ = W3078
tmp12254 := PrimCons(W3075, Nil)

tmp12255 := PrimCons(Nil, tmp12254)

tmp12256 := Call(__e, PrimFunc(symshen_4lr_1rule), W3067, W3070, tmp12255)


__e.TailApply(PrimFunc(symshen_4comb), W3078, tmp12256)
return


}, 1)

tmp12257 := Call(__e, PrimFunc(symshen_4in_1_6), W3077)


__e.TailApply(tmp12253, tmp12257)
return


}


}, 1)

tmp12260 := Call(__e, PrimFunc(symshen_4_5sc_6), W3076)


__e.TailApply(tmp12252, tmp12260)
return


}, 1)

tmp12261 := Call(__e, PrimFunc(symshen_4in_1_6), W3074)


__e.TailApply(tmp12251, tmp12261)
return


}, 1)

tmp12262 := Call(__e, PrimFunc(symshen_4_5_1out), W3074)


__e.TailApply(tmp12250, tmp12262)
return


}


}, 1)

tmp12265 := Call(__e, PrimFunc(symshen_4_5formula_6), W3073)


__e.TailApply(tmp12249, tmp12265)
return


}, 1)

tmp12266 := Call(__e, PrimFunc(symshen_4in_1_6), W3072)


__e.TailApply(tmp12248, tmp12266)
return


}


}, 1)

tmp12269 := Call(__e, PrimFunc(symshen_4_5dbl_6), W3071)


__e.TailApply(tmp12247, tmp12269)
return


}, 1)

tmp12270 := Call(__e, PrimFunc(symshen_4in_1_6), W3069)


__e.TailApply(tmp12246, tmp12270)
return


}, 1)

tmp12271 := Call(__e, PrimFunc(symshen_4_5_1out), W3069)


__e.TailApply(tmp12245, tmp12271)
return


}


}, 1)

tmp12274 := Call(__e, PrimFunc(symshen_4_5formulae_6), W3068)


__e.TailApply(tmp12244, tmp12274)
return


}, 1)

tmp12275 := Call(__e, PrimFunc(symshen_4in_1_6), W3066)


__e.TailApply(tmp12243, tmp12275)
return


}, 1)

tmp12276 := Call(__e, PrimFunc(symshen_4_5_1out), W3066)


__e.TailApply(tmp12242, tmp12276)
return


}


}, 1)

tmp12279 := Call(__e, PrimFunc(symshen_4_5sides_6), V3064)


tmp12280 := Call(__e, tmp12241, tmp12279)


__e.TailApply(tmp12238, tmp12280)
return


}, 1)

tmp12281 := Call(__e, ns2_1set, symshen_4_5double_6, tmp12237)


_ = tmp12281

tmp12282 := MakeNative(func(__e *ControlFlow) {
V3079 := __e.Get(1)
_ = V3079
tmp12283 := MakeNative(func(__e *ControlFlow) {
W3080 := __e.Get(1)
_ = W3080
tmp12306 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3080)


if True == tmp12306 {
tmp12284 := MakeNative(func(__e *ControlFlow) {
W3089 := __e.Get(1)
_ = W3089
tmp12286 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3089)


if True == tmp12286 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3089)
return
}


}, 1)

tmp12287 := MakeNative(func(__e *ControlFlow) {
W3090 := __e.Get(1)
_ = W3090
tmp12302 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3090)


if True == tmp12302 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12288 := MakeNative(func(__e *ControlFlow) {
W3091 := __e.Get(1)
_ = W3091
tmp12289 := MakeNative(func(__e *ControlFlow) {
W3092 := __e.Get(1)
_ = W3092
tmp12290 := MakeNative(func(__e *ControlFlow) {
W3093 := __e.Get(1)
_ = W3093
tmp12297 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3093)


if True == tmp12297 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12291 := MakeNative(func(__e *ControlFlow) {
W3094 := __e.Get(1)
_ = W3094
tmp12292 := PrimCons(W3091, Nil)

tmp12293 := PrimCons(Nil, tmp12292)

tmp12294 := PrimCons(tmp12293, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W3094, tmp12294)
return


}, 1)

tmp12295 := Call(__e, PrimFunc(symshen_4in_1_6), W3093)


__e.TailApply(tmp12291, tmp12295)
return


}


}, 1)

tmp12298 := Call(__e, PrimFunc(symshen_4_5sc_6), W3092)


__e.TailApply(tmp12290, tmp12298)
return


}, 1)

tmp12299 := Call(__e, PrimFunc(symshen_4in_1_6), W3090)


__e.TailApply(tmp12289, tmp12299)
return


}, 1)

tmp12300 := Call(__e, PrimFunc(symshen_4_5_1out), W3090)


__e.TailApply(tmp12288, tmp12300)
return


}


}, 1)

tmp12303 := Call(__e, PrimFunc(symshen_4_5formula_6), V3079)


tmp12304 := Call(__e, tmp12287, tmp12303)


__e.TailApply(tmp12284, tmp12304)
return


} else {
__e.Return(W3080)
return
}


}, 1)

tmp12307 := MakeNative(func(__e *ControlFlow) {
W3081 := __e.Get(1)
_ = W3081
tmp12330 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3081)


if True == tmp12330 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12308 := MakeNative(func(__e *ControlFlow) {
W3082 := __e.Get(1)
_ = W3082
tmp12309 := MakeNative(func(__e *ControlFlow) {
W3083 := __e.Get(1)
_ = W3083
tmp12310 := MakeNative(func(__e *ControlFlow) {
W3084 := __e.Get(1)
_ = W3084
tmp12325 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3084)


if True == tmp12325 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12311 := MakeNative(func(__e *ControlFlow) {
W3085 := __e.Get(1)
_ = W3085
tmp12312 := MakeNative(func(__e *ControlFlow) {
W3086 := __e.Get(1)
_ = W3086
tmp12321 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3086)


if True == tmp12321 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12313 := MakeNative(func(__e *ControlFlow) {
W3087 := __e.Get(1)
_ = W3087
tmp12314 := MakeNative(func(__e *ControlFlow) {
W3088 := __e.Get(1)
_ = W3088
tmp12315 := PrimCons(W3082, Nil)

tmp12316 := PrimCons(Nil, tmp12315)

tmp12317 := PrimCons(tmp12316, W3087)

__e.TailApply(PrimFunc(symshen_4comb), W3088, tmp12317)
return


}, 1)

tmp12318 := Call(__e, PrimFunc(symshen_4in_1_6), W3086)


__e.TailApply(tmp12314, tmp12318)
return


}, 1)

tmp12319 := Call(__e, PrimFunc(symshen_4_5_1out), W3086)


__e.TailApply(tmp12313, tmp12319)
return


}


}, 1)

tmp12322 := Call(__e, PrimFunc(symshen_4_5formulae_6), W3085)


__e.TailApply(tmp12312, tmp12322)
return


}, 1)

tmp12323 := Call(__e, PrimFunc(symshen_4in_1_6), W3084)


__e.TailApply(tmp12311, tmp12323)
return


}


}, 1)

tmp12326 := Call(__e, PrimFunc(symshen_4_5sc_6), W3083)


__e.TailApply(tmp12310, tmp12326)
return


}, 1)

tmp12327 := Call(__e, PrimFunc(symshen_4in_1_6), W3081)


__e.TailApply(tmp12309, tmp12327)
return


}, 1)

tmp12328 := Call(__e, PrimFunc(symshen_4_5_1out), W3081)


__e.TailApply(tmp12308, tmp12328)
return


}


}, 1)

tmp12331 := Call(__e, PrimFunc(symshen_4_5formula_6), V3079)


tmp12332 := Call(__e, tmp12307, tmp12331)


__e.TailApply(tmp12283, tmp12332)
return


}, 1)

tmp12333 := Call(__e, ns2_1set, symshen_4_5formulae_6, tmp12282)


_ = tmp12333

tmp12334 := MakeNative(func(__e *ControlFlow) {
V3095 := __e.Get(1)
_ = V3095
tmp12335 := MakeNative(func(__e *ControlFlow) {
W3096 := __e.Get(1)
_ = W3096
tmp12351 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3096)


if True == tmp12351 {
tmp12336 := MakeNative(func(__e *ControlFlow) {
W3104 := __e.Get(1)
_ = W3104
tmp12338 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3104)


if True == tmp12338 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3104)
return
}


}, 1)

tmp12339 := MakeNative(func(__e *ControlFlow) {
W3105 := __e.Get(1)
_ = W3105
tmp12347 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3105)


if True == tmp12347 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12340 := MakeNative(func(__e *ControlFlow) {
W3106 := __e.Get(1)
_ = W3106
tmp12341 := MakeNative(func(__e *ControlFlow) {
W3107 := __e.Get(1)
_ = W3107
tmp12342 := PrimCons(W3106, Nil)

tmp12343 := PrimCons(Nil, tmp12342)

__e.TailApply(PrimFunc(symshen_4comb), W3107, tmp12343)
return


}, 1)

tmp12344 := Call(__e, PrimFunc(symshen_4in_1_6), W3105)


__e.TailApply(tmp12341, tmp12344)
return


}, 1)

tmp12345 := Call(__e, PrimFunc(symshen_4_5_1out), W3105)


__e.TailApply(tmp12340, tmp12345)
return


}


}, 1)

tmp12348 := Call(__e, PrimFunc(symshen_4_5formula_6), V3095)


tmp12349 := Call(__e, tmp12339, tmp12348)


__e.TailApply(tmp12336, tmp12349)
return


} else {
__e.Return(W3096)
return
}


}, 1)

tmp12352 := MakeNative(func(__e *ControlFlow) {
W3097 := __e.Get(1)
_ = W3097
tmp12372 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3097)


if True == tmp12372 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12353 := MakeNative(func(__e *ControlFlow) {
W3098 := __e.Get(1)
_ = W3098
tmp12354 := MakeNative(func(__e *ControlFlow) {
W3099 := __e.Get(1)
_ = W3099
tmp12368 := Call(__e, PrimFunc(symshen_4hds_a_2), W3099, sym_6_6)


if True == tmp12368 {
tmp12355 := MakeNative(func(__e *ControlFlow) {
W3100 := __e.Get(1)
_ = W3100
tmp12356 := MakeNative(func(__e *ControlFlow) {
W3101 := __e.Get(1)
_ = W3101
tmp12364 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3101)


if True == tmp12364 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12357 := MakeNative(func(__e *ControlFlow) {
W3102 := __e.Get(1)
_ = W3102
tmp12358 := MakeNative(func(__e *ControlFlow) {
W3103 := __e.Get(1)
_ = W3103
tmp12359 := PrimCons(W3102, Nil)

tmp12360 := PrimCons(W3098, tmp12359)

__e.TailApply(PrimFunc(symshen_4comb), W3103, tmp12360)
return


}, 1)

tmp12361 := Call(__e, PrimFunc(symshen_4in_1_6), W3101)


__e.TailApply(tmp12358, tmp12361)
return


}, 1)

tmp12362 := Call(__e, PrimFunc(symshen_4_5_1out), W3101)


__e.TailApply(tmp12357, tmp12362)
return


}


}, 1)

tmp12365 := Call(__e, PrimFunc(symshen_4_5formula_6), W3100)


__e.TailApply(tmp12356, tmp12365)
return


}, 1)

tmp12366 := Call(__e, PrimFunc(symtail), W3099)


__e.TailApply(tmp12355, tmp12366)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12369 := Call(__e, PrimFunc(symshen_4in_1_6), W3097)


__e.TailApply(tmp12354, tmp12369)
return


}, 1)

tmp12370 := Call(__e, PrimFunc(symshen_4_5_1out), W3097)


__e.TailApply(tmp12353, tmp12370)
return


}


}, 1)

tmp12373 := Call(__e, PrimFunc(symshen_4_5ass_6), V3095)


tmp12374 := Call(__e, tmp12352, tmp12373)


__e.TailApply(tmp12335, tmp12374)
return


}, 1)

tmp12375 := Call(__e, ns2_1set, symshen_4_5conc_6, tmp12334)


_ = tmp12375

tmp12376 := MakeNative(func(__e *ControlFlow) {
V3108 := __e.Get(1)
_ = V3108
tmp12377 := MakeNative(func(__e *ControlFlow) {
W3109 := __e.Get(1)
_ = W3109
tmp12389 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3109)


if True == tmp12389 {
tmp12378 := MakeNative(func(__e *ControlFlow) {
W3118 := __e.Get(1)
_ = W3118
tmp12380 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3118)


if True == tmp12380 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3118)
return
}


}, 1)

tmp12381 := MakeNative(func(__e *ControlFlow) {
W3119 := __e.Get(1)
_ = W3119
tmp12385 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3119)


if True == tmp12385 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12382 := MakeNative(func(__e *ControlFlow) {
W3120 := __e.Get(1)
_ = W3120
__e.TailApply(PrimFunc(symshen_4comb), W3120, Nil)
return
}, 1)

tmp12383 := Call(__e, PrimFunc(symshen_4in_1_6), W3119)


__e.TailApply(tmp12382, tmp12383)
return


}


}, 1)

tmp12386 := Call(__e, PrimFunc(sym_5e_6), V3108)


tmp12387 := Call(__e, tmp12381, tmp12386)


__e.TailApply(tmp12378, tmp12387)
return


} else {
__e.Return(W3109)
return
}


}, 1)

tmp12390 := MakeNative(func(__e *ControlFlow) {
W3110 := __e.Get(1)
_ = W3110
tmp12411 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3110)


if True == tmp12411 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12391 := MakeNative(func(__e *ControlFlow) {
W3111 := __e.Get(1)
_ = W3111
tmp12392 := MakeNative(func(__e *ControlFlow) {
W3112 := __e.Get(1)
_ = W3112
tmp12393 := MakeNative(func(__e *ControlFlow) {
W3113 := __e.Get(1)
_ = W3113
tmp12406 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3113)


if True == tmp12406 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12394 := MakeNative(func(__e *ControlFlow) {
W3114 := __e.Get(1)
_ = W3114
tmp12395 := MakeNative(func(__e *ControlFlow) {
W3115 := __e.Get(1)
_ = W3115
tmp12402 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3115)


if True == tmp12402 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12396 := MakeNative(func(__e *ControlFlow) {
W3116 := __e.Get(1)
_ = W3116
tmp12397 := MakeNative(func(__e *ControlFlow) {
W3117 := __e.Get(1)
_ = W3117
tmp12398 := PrimCons(W3111, W3116)

__e.TailApply(PrimFunc(symshen_4comb), W3117, tmp12398)
return


}, 1)

tmp12399 := Call(__e, PrimFunc(symshen_4in_1_6), W3115)


__e.TailApply(tmp12397, tmp12399)
return


}, 1)

tmp12400 := Call(__e, PrimFunc(symshen_4_5_1out), W3115)


__e.TailApply(tmp12396, tmp12400)
return


}


}, 1)

tmp12403 := Call(__e, PrimFunc(symshen_4_5prems_6), W3114)


__e.TailApply(tmp12395, tmp12403)
return


}, 1)

tmp12404 := Call(__e, PrimFunc(symshen_4in_1_6), W3113)


__e.TailApply(tmp12394, tmp12404)
return


}


}, 1)

tmp12407 := Call(__e, PrimFunc(symshen_4_5sc_6), W3112)


__e.TailApply(tmp12393, tmp12407)
return


}, 1)

tmp12408 := Call(__e, PrimFunc(symshen_4in_1_6), W3110)


__e.TailApply(tmp12392, tmp12408)
return


}, 1)

tmp12409 := Call(__e, PrimFunc(symshen_4_5_1out), W3110)


__e.TailApply(tmp12391, tmp12409)
return


}


}, 1)

tmp12412 := Call(__e, PrimFunc(symshen_4_5prem_6), V3108)


tmp12413 := Call(__e, tmp12390, tmp12412)


__e.TailApply(tmp12377, tmp12413)
return


}, 1)

tmp12414 := Call(__e, ns2_1set, symshen_4_5prems_6, tmp12376)


_ = tmp12414

tmp12415 := MakeNative(func(__e *ControlFlow) {
V3121 := __e.Get(1)
_ = V3121
tmp12416 := MakeNative(func(__e *ControlFlow) {
W3122 := __e.Get(1)
_ = W3122
tmp12458 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3122)


if True == tmp12458 {
tmp12417 := MakeNative(func(__e *ControlFlow) {
W3124 := __e.Get(1)
_ = W3124
tmp12433 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3124)


if True == tmp12433 {
tmp12418 := MakeNative(func(__e *ControlFlow) {
W3132 := __e.Get(1)
_ = W3132
tmp12420 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3132)


if True == tmp12420 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3132)
return
}


}, 1)

tmp12421 := MakeNative(func(__e *ControlFlow) {
W3133 := __e.Get(1)
_ = W3133
tmp12429 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3133)


if True == tmp12429 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12422 := MakeNative(func(__e *ControlFlow) {
W3134 := __e.Get(1)
_ = W3134
tmp12423 := MakeNative(func(__e *ControlFlow) {
W3135 := __e.Get(1)
_ = W3135
tmp12424 := PrimCons(W3134, Nil)

tmp12425 := PrimCons(Nil, tmp12424)

__e.TailApply(PrimFunc(symshen_4comb), W3135, tmp12425)
return


}, 1)

tmp12426 := Call(__e, PrimFunc(symshen_4in_1_6), W3133)


__e.TailApply(tmp12423, tmp12426)
return


}, 1)

tmp12427 := Call(__e, PrimFunc(symshen_4_5_1out), W3133)


__e.TailApply(tmp12422, tmp12427)
return


}


}, 1)

tmp12430 := Call(__e, PrimFunc(symshen_4_5formula_6), V3121)


tmp12431 := Call(__e, tmp12421, tmp12430)


__e.TailApply(tmp12418, tmp12431)
return


} else {
__e.Return(W3124)
return
}


}, 1)

tmp12434 := MakeNative(func(__e *ControlFlow) {
W3125 := __e.Get(1)
_ = W3125
tmp12454 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3125)


if True == tmp12454 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12435 := MakeNative(func(__e *ControlFlow) {
W3126 := __e.Get(1)
_ = W3126
tmp12436 := MakeNative(func(__e *ControlFlow) {
W3127 := __e.Get(1)
_ = W3127
tmp12450 := Call(__e, PrimFunc(symshen_4hds_a_2), W3127, sym_6_6)


if True == tmp12450 {
tmp12437 := MakeNative(func(__e *ControlFlow) {
W3128 := __e.Get(1)
_ = W3128
tmp12438 := MakeNative(func(__e *ControlFlow) {
W3129 := __e.Get(1)
_ = W3129
tmp12446 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3129)


if True == tmp12446 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12439 := MakeNative(func(__e *ControlFlow) {
W3130 := __e.Get(1)
_ = W3130
tmp12440 := MakeNative(func(__e *ControlFlow) {
W3131 := __e.Get(1)
_ = W3131
tmp12441 := PrimCons(W3130, Nil)

tmp12442 := PrimCons(W3126, tmp12441)

__e.TailApply(PrimFunc(symshen_4comb), W3131, tmp12442)
return


}, 1)

tmp12443 := Call(__e, PrimFunc(symshen_4in_1_6), W3129)


__e.TailApply(tmp12440, tmp12443)
return


}, 1)

tmp12444 := Call(__e, PrimFunc(symshen_4_5_1out), W3129)


__e.TailApply(tmp12439, tmp12444)
return


}


}, 1)

tmp12447 := Call(__e, PrimFunc(symshen_4_5formula_6), W3128)


__e.TailApply(tmp12438, tmp12447)
return


}, 1)

tmp12448 := Call(__e, PrimFunc(symtail), W3127)


__e.TailApply(tmp12437, tmp12448)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12451 := Call(__e, PrimFunc(symshen_4in_1_6), W3125)


__e.TailApply(tmp12436, tmp12451)
return


}, 1)

tmp12452 := Call(__e, PrimFunc(symshen_4_5_1out), W3125)


__e.TailApply(tmp12435, tmp12452)
return


}


}, 1)

tmp12455 := Call(__e, PrimFunc(symshen_4_5ass_6), V3121)


tmp12456 := Call(__e, tmp12434, tmp12455)


__e.TailApply(tmp12417, tmp12456)
return


} else {
__e.Return(W3122)
return
}


}, 1)

tmp12464 := Call(__e, PrimFunc(symshen_4hds_a_2), V3121, sym_b)


var ifres12459 Obj

if True == tmp12464 {
tmp12460 := MakeNative(func(__e *ControlFlow) {
W3123 := __e.Get(1)
_ = W3123
__e.TailApply(PrimFunc(symshen_4comb), W3123, sym_b)
return
}, 1)

tmp12461 := Call(__e, PrimFunc(symtail), V3121)


tmp12462 := Call(__e, tmp12460, tmp12461)


ifres12459 = tmp12462


} else {
tmp12463 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12459 = tmp12463


}

__e.TailApply(tmp12416, ifres12459)
return


}, 1)

tmp12465 := Call(__e, ns2_1set, symshen_4_5prem_6, tmp12415)


_ = tmp12465

tmp12466 := MakeNative(func(__e *ControlFlow) {
V3136 := __e.Get(1)
_ = V3136
tmp12467 := MakeNative(func(__e *ControlFlow) {
W3137 := __e.Get(1)
_ = W3137
tmp12492 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3137)


if True == tmp12492 {
tmp12468 := MakeNative(func(__e *ControlFlow) {
W3146 := __e.Get(1)
_ = W3146
tmp12480 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3146)


if True == tmp12480 {
tmp12469 := MakeNative(func(__e *ControlFlow) {
W3150 := __e.Get(1)
_ = W3150
tmp12471 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3150)


if True == tmp12471 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3150)
return
}


}, 1)

tmp12472 := MakeNative(func(__e *ControlFlow) {
W3151 := __e.Get(1)
_ = W3151
tmp12476 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3151)


if True == tmp12476 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12473 := MakeNative(func(__e *ControlFlow) {
W3152 := __e.Get(1)
_ = W3152
__e.TailApply(PrimFunc(symshen_4comb), W3152, Nil)
return
}, 1)

tmp12474 := Call(__e, PrimFunc(symshen_4in_1_6), W3151)


__e.TailApply(tmp12473, tmp12474)
return


}


}, 1)

tmp12477 := Call(__e, PrimFunc(sym_5e_6), V3136)


tmp12478 := Call(__e, tmp12472, tmp12477)


__e.TailApply(tmp12469, tmp12478)
return


} else {
__e.Return(W3146)
return
}


}, 1)

tmp12481 := MakeNative(func(__e *ControlFlow) {
W3147 := __e.Get(1)
_ = W3147
tmp12488 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3147)


if True == tmp12488 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12482 := MakeNative(func(__e *ControlFlow) {
W3148 := __e.Get(1)
_ = W3148
tmp12483 := MakeNative(func(__e *ControlFlow) {
W3149 := __e.Get(1)
_ = W3149
tmp12484 := PrimCons(W3148, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W3149, tmp12484)
return


}, 1)

tmp12485 := Call(__e, PrimFunc(symshen_4in_1_6), W3147)


__e.TailApply(tmp12483, tmp12485)
return


}, 1)

tmp12486 := Call(__e, PrimFunc(symshen_4_5_1out), W3147)


__e.TailApply(tmp12482, tmp12486)
return


}


}, 1)

tmp12489 := Call(__e, PrimFunc(symshen_4_5formula_6), V3136)


tmp12490 := Call(__e, tmp12481, tmp12489)


__e.TailApply(tmp12468, tmp12490)
return


} else {
__e.Return(W3137)
return
}


}, 1)

tmp12493 := MakeNative(func(__e *ControlFlow) {
W3138 := __e.Get(1)
_ = W3138
tmp12514 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3138)


if True == tmp12514 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12494 := MakeNative(func(__e *ControlFlow) {
W3139 := __e.Get(1)
_ = W3139
tmp12495 := MakeNative(func(__e *ControlFlow) {
W3140 := __e.Get(1)
_ = W3140
tmp12496 := MakeNative(func(__e *ControlFlow) {
W3141 := __e.Get(1)
_ = W3141
tmp12509 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3141)


if True == tmp12509 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12497 := MakeNative(func(__e *ControlFlow) {
W3142 := __e.Get(1)
_ = W3142
tmp12498 := MakeNative(func(__e *ControlFlow) {
W3143 := __e.Get(1)
_ = W3143
tmp12505 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3143)


if True == tmp12505 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12499 := MakeNative(func(__e *ControlFlow) {
W3144 := __e.Get(1)
_ = W3144
tmp12500 := MakeNative(func(__e *ControlFlow) {
W3145 := __e.Get(1)
_ = W3145
tmp12501 := PrimCons(W3139, W3144)

__e.TailApply(PrimFunc(symshen_4comb), W3145, tmp12501)
return


}, 1)

tmp12502 := Call(__e, PrimFunc(symshen_4in_1_6), W3143)


__e.TailApply(tmp12500, tmp12502)
return


}, 1)

tmp12503 := Call(__e, PrimFunc(symshen_4_5_1out), W3143)


__e.TailApply(tmp12499, tmp12503)
return


}


}, 1)

tmp12506 := Call(__e, PrimFunc(symshen_4_5ass_6), W3142)


__e.TailApply(tmp12498, tmp12506)
return


}, 1)

tmp12507 := Call(__e, PrimFunc(symshen_4in_1_6), W3141)


__e.TailApply(tmp12497, tmp12507)
return


}


}, 1)

tmp12510 := Call(__e, PrimFunc(symshen_4_5iscomma_6), W3140)


__e.TailApply(tmp12496, tmp12510)
return


}, 1)

tmp12511 := Call(__e, PrimFunc(symshen_4in_1_6), W3138)


__e.TailApply(tmp12495, tmp12511)
return


}, 1)

tmp12512 := Call(__e, PrimFunc(symshen_4_5_1out), W3138)


__e.TailApply(tmp12494, tmp12512)
return


}


}, 1)

tmp12515 := Call(__e, PrimFunc(symshen_4_5formula_6), V3136)


tmp12516 := Call(__e, tmp12493, tmp12515)


__e.TailApply(tmp12467, tmp12516)
return


}, 1)

tmp12517 := Call(__e, ns2_1set, symshen_4_5ass_6, tmp12466)


_ = tmp12517

tmp12518 := MakeNative(func(__e *ControlFlow) {
V3153 := __e.Get(1)
_ = V3153
tmp12519 := MakeNative(func(__e *ControlFlow) {
W3154 := __e.Get(1)
_ = W3154
tmp12521 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3154)


if True == tmp12521 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3154)
return
}


}, 1)

tmp12532 := PrimIsPair(V3153)

var ifres12522 Obj

if True == tmp12532 {
tmp12523 := MakeNative(func(__e *ControlFlow) {
W3155 := __e.Get(1)
_ = W3155
tmp12524 := MakeNative(func(__e *ControlFlow) {
W3156 := __e.Get(1)
_ = W3156
tmp12526 := PrimIntern(MakeString(","))

tmp12527 := PrimEqual(W3155, tmp12526)

if True == tmp12527 {
__e.TailApply(PrimFunc(symshen_4comb), W3156, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12528 := Call(__e, PrimFunc(symtail), V3153)


__e.TailApply(tmp12524, tmp12528)
return


}, 1)

tmp12529 := Call(__e, PrimFunc(symhead), V3153)


tmp12530 := Call(__e, tmp12523, tmp12529)


ifres12522 = tmp12530


} else {
tmp12531 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12522 = tmp12531


}

__e.TailApply(tmp12519, ifres12522)
return


}, 1)

tmp12533 := Call(__e, ns2_1set, symshen_4_5iscomma_6, tmp12518)


_ = tmp12533

tmp12534 := MakeNative(func(__e *ControlFlow) {
V3157 := __e.Get(1)
_ = V3157
tmp12535 := MakeNative(func(__e *ControlFlow) {
W3158 := __e.Get(1)
_ = W3158
tmp12549 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3158)


if True == tmp12549 {
tmp12536 := MakeNative(func(__e *ControlFlow) {
W3167 := __e.Get(1)
_ = W3167
tmp12538 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3167)


if True == tmp12538 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3167)
return
}


}, 1)

tmp12539 := MakeNative(func(__e *ControlFlow) {
W3168 := __e.Get(1)
_ = W3168
tmp12545 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3168)


if True == tmp12545 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12540 := MakeNative(func(__e *ControlFlow) {
W3169 := __e.Get(1)
_ = W3169
tmp12541 := MakeNative(func(__e *ControlFlow) {
W3170 := __e.Get(1)
_ = W3170
__e.TailApply(PrimFunc(symshen_4comb), W3170, W3169)
return
}, 1)

tmp12542 := Call(__e, PrimFunc(symshen_4in_1_6), W3168)


__e.TailApply(tmp12541, tmp12542)
return


}, 1)

tmp12543 := Call(__e, PrimFunc(symshen_4_5_1out), W3168)


__e.TailApply(tmp12540, tmp12543)
return


}


}, 1)

tmp12546 := Call(__e, PrimFunc(symshen_4_5expr_6), V3157)


tmp12547 := Call(__e, tmp12539, tmp12546)


__e.TailApply(tmp12536, tmp12547)
return


} else {
__e.Return(W3158)
return
}


}, 1)

tmp12550 := MakeNative(func(__e *ControlFlow) {
W3159 := __e.Get(1)
_ = W3159
tmp12576 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3159)


if True == tmp12576 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12551 := MakeNative(func(__e *ControlFlow) {
W3160 := __e.Get(1)
_ = W3160
tmp12552 := MakeNative(func(__e *ControlFlow) {
W3161 := __e.Get(1)
_ = W3161
tmp12553 := MakeNative(func(__e *ControlFlow) {
W3162 := __e.Get(1)
_ = W3162
tmp12571 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3162)


if True == tmp12571 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12554 := MakeNative(func(__e *ControlFlow) {
W3163 := __e.Get(1)
_ = W3163
tmp12555 := MakeNative(func(__e *ControlFlow) {
W3164 := __e.Get(1)
_ = W3164
tmp12567 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3164)


if True == tmp12567 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12556 := MakeNative(func(__e *ControlFlow) {
W3165 := __e.Get(1)
_ = W3165
tmp12557 := MakeNative(func(__e *ControlFlow) {
W3166 := __e.Get(1)
_ = W3166
tmp12558 := Call(__e, PrimFunc(symshen_4curry), W3160)


tmp12559 := PrimIntern(MakeString(":"))

tmp12560 := Call(__e, PrimFunc(symshen_4rectify_1type), W3165)


tmp12561 := PrimCons(tmp12560, Nil)

tmp12562 := PrimCons(tmp12559, tmp12561)

tmp12563 := PrimCons(tmp12558, tmp12562)

__e.TailApply(PrimFunc(symshen_4comb), W3166, tmp12563)
return


}, 1)

tmp12564 := Call(__e, PrimFunc(symshen_4in_1_6), W3164)


__e.TailApply(tmp12557, tmp12564)
return


}, 1)

tmp12565 := Call(__e, PrimFunc(symshen_4_5_1out), W3164)


__e.TailApply(tmp12556, tmp12565)
return


}


}, 1)

tmp12568 := Call(__e, PrimFunc(symshen_4_5type_6), W3163)


__e.TailApply(tmp12555, tmp12568)
return


}, 1)

tmp12569 := Call(__e, PrimFunc(symshen_4in_1_6), W3162)


__e.TailApply(tmp12554, tmp12569)
return


}


}, 1)

tmp12572 := Call(__e, PrimFunc(symshen_4_5iscolon_6), W3161)


__e.TailApply(tmp12553, tmp12572)
return


}, 1)

tmp12573 := Call(__e, PrimFunc(symshen_4in_1_6), W3159)


__e.TailApply(tmp12552, tmp12573)
return


}, 1)

tmp12574 := Call(__e, PrimFunc(symshen_4_5_1out), W3159)


__e.TailApply(tmp12551, tmp12574)
return


}


}, 1)

tmp12577 := Call(__e, PrimFunc(symshen_4_5expr_6), V3157)


tmp12578 := Call(__e, tmp12550, tmp12577)


__e.TailApply(tmp12535, tmp12578)
return


}, 1)

tmp12579 := Call(__e, ns2_1set, symshen_4_5formula_6, tmp12534)


_ = tmp12579

tmp12580 := MakeNative(func(__e *ControlFlow) {
V3171 := __e.Get(1)
_ = V3171
tmp12581 := MakeNative(func(__e *ControlFlow) {
W3172 := __e.Get(1)
_ = W3172
tmp12583 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3172)


if True == tmp12583 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3172)
return
}


}, 1)

tmp12594 := PrimIsPair(V3171)

var ifres12584 Obj

if True == tmp12594 {
tmp12585 := MakeNative(func(__e *ControlFlow) {
W3173 := __e.Get(1)
_ = W3173
tmp12586 := MakeNative(func(__e *ControlFlow) {
W3174 := __e.Get(1)
_ = W3174
tmp12588 := PrimIntern(MakeString(":"))

tmp12589 := PrimEqual(W3173, tmp12588)

if True == tmp12589 {
__e.TailApply(PrimFunc(symshen_4comb), W3174, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12590 := Call(__e, PrimFunc(symtail), V3171)


__e.TailApply(tmp12586, tmp12590)
return


}, 1)

tmp12591 := Call(__e, PrimFunc(symhead), V3171)


tmp12592 := Call(__e, tmp12585, tmp12591)


ifres12584 = tmp12592


} else {
tmp12593 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12584 = tmp12593


}

__e.TailApply(tmp12581, ifres12584)
return


}, 1)

tmp12595 := Call(__e, ns2_1set, symshen_4_5iscolon_6, tmp12580)


_ = tmp12595

tmp12596 := MakeNative(func(__e *ControlFlow) {
V3175 := __e.Get(1)
_ = V3175
tmp12597 := MakeNative(func(__e *ControlFlow) {
W3176 := __e.Get(1)
_ = W3176
tmp12609 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3176)


if True == tmp12609 {
tmp12598 := MakeNative(func(__e *ControlFlow) {
W3183 := __e.Get(1)
_ = W3183
tmp12600 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3183)


if True == tmp12600 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3183)
return
}


}, 1)

tmp12601 := MakeNative(func(__e *ControlFlow) {
W3184 := __e.Get(1)
_ = W3184
tmp12605 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3184)


if True == tmp12605 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12602 := MakeNative(func(__e *ControlFlow) {
W3185 := __e.Get(1)
_ = W3185
__e.TailApply(PrimFunc(symshen_4comb), W3185, Nil)
return
}, 1)

tmp12603 := Call(__e, PrimFunc(symshen_4in_1_6), W3184)


__e.TailApply(tmp12602, tmp12603)
return


}


}, 1)

tmp12606 := Call(__e, PrimFunc(sym_5e_6), V3175)


tmp12607 := Call(__e, tmp12601, tmp12606)


__e.TailApply(tmp12598, tmp12607)
return


} else {
__e.Return(W3176)
return
}


}, 1)

tmp12610 := MakeNative(func(__e *ControlFlow) {
W3177 := __e.Get(1)
_ = W3177
tmp12625 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3177)


if True == tmp12625 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12611 := MakeNative(func(__e *ControlFlow) {
W3178 := __e.Get(1)
_ = W3178
tmp12612 := MakeNative(func(__e *ControlFlow) {
W3179 := __e.Get(1)
_ = W3179
tmp12613 := MakeNative(func(__e *ControlFlow) {
W3180 := __e.Get(1)
_ = W3180
tmp12620 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3180)


if True == tmp12620 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12614 := MakeNative(func(__e *ControlFlow) {
W3181 := __e.Get(1)
_ = W3181
tmp12615 := MakeNative(func(__e *ControlFlow) {
W3182 := __e.Get(1)
_ = W3182
tmp12616 := PrimCons(W3178, W3181)

__e.TailApply(PrimFunc(symshen_4comb), W3182, tmp12616)
return


}, 1)

tmp12617 := Call(__e, PrimFunc(symshen_4in_1_6), W3180)


__e.TailApply(tmp12615, tmp12617)
return


}, 1)

tmp12618 := Call(__e, PrimFunc(symshen_4_5_1out), W3180)


__e.TailApply(tmp12614, tmp12618)
return


}


}, 1)

tmp12621 := Call(__e, PrimFunc(symshen_4_5sides_6), W3179)


__e.TailApply(tmp12613, tmp12621)
return


}, 1)

tmp12622 := Call(__e, PrimFunc(symshen_4in_1_6), W3177)


__e.TailApply(tmp12612, tmp12622)
return


}, 1)

tmp12623 := Call(__e, PrimFunc(symshen_4_5_1out), W3177)


__e.TailApply(tmp12611, tmp12623)
return


}


}, 1)

tmp12626 := Call(__e, PrimFunc(symshen_4_5side_6), V3175)


tmp12627 := Call(__e, tmp12610, tmp12626)


__e.TailApply(tmp12597, tmp12627)
return


}, 1)

tmp12628 := Call(__e, ns2_1set, symshen_4_5sides_6, tmp12596)


_ = tmp12628

tmp12629 := MakeNative(func(__e *ControlFlow) {
V3186 := __e.Get(1)
_ = V3186
tmp12630 := MakeNative(func(__e *ControlFlow) {
W3187 := __e.Get(1)
_ = W3187
tmp12680 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3187)


if True == tmp12680 {
tmp12631 := MakeNative(func(__e *ControlFlow) {
W3191 := __e.Get(1)
_ = W3191
tmp12657 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3191)


if True == tmp12657 {
tmp12632 := MakeNative(func(__e *ControlFlow) {
W3197 := __e.Get(1)
_ = W3197
tmp12634 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3197)


if True == tmp12634 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3197)
return
}


}, 1)

tmp12655 := Call(__e, PrimFunc(symshen_4hds_a_2), V3186, symshen_4let_b)


var ifres12635 Obj

if True == tmp12655 {
tmp12636 := MakeNative(func(__e *ControlFlow) {
W3198 := __e.Get(1)
_ = W3198
tmp12651 := PrimIsPair(W3198)

if True == tmp12651 {
tmp12637 := MakeNative(func(__e *ControlFlow) {
W3199 := __e.Get(1)
_ = W3199
tmp12638 := MakeNative(func(__e *ControlFlow) {
W3200 := __e.Get(1)
_ = W3200
tmp12647 := PrimIsPair(W3200)

if True == tmp12647 {
tmp12639 := MakeNative(func(__e *ControlFlow) {
W3201 := __e.Get(1)
_ = W3201
tmp12640 := MakeNative(func(__e *ControlFlow) {
W3202 := __e.Get(1)
_ = W3202
tmp12641 := PrimCons(W3201, Nil)

tmp12642 := PrimCons(W3199, tmp12641)

tmp12643 := PrimCons(symshen_4let_b, tmp12642)

__e.TailApply(PrimFunc(symshen_4comb), W3202, tmp12643)
return


}, 1)

tmp12644 := Call(__e, PrimFunc(symtail), W3200)


__e.TailApply(tmp12640, tmp12644)
return


}, 1)

tmp12645 := Call(__e, PrimFunc(symhead), W3200)


__e.TailApply(tmp12639, tmp12645)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12648 := Call(__e, PrimFunc(symtail), W3198)


__e.TailApply(tmp12638, tmp12648)
return


}, 1)

tmp12649 := Call(__e, PrimFunc(symhead), W3198)


__e.TailApply(tmp12637, tmp12649)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12652 := Call(__e, PrimFunc(symtail), V3186)


tmp12653 := Call(__e, tmp12636, tmp12652)


ifres12635 = tmp12653


} else {
tmp12654 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12635 = tmp12654


}

__e.TailApply(tmp12632, ifres12635)
return


} else {
__e.Return(W3191)
return
}


}, 1)

tmp12678 := Call(__e, PrimFunc(symshen_4hds_a_2), V3186, symlet)


var ifres12658 Obj

if True == tmp12678 {
tmp12659 := MakeNative(func(__e *ControlFlow) {
W3192 := __e.Get(1)
_ = W3192
tmp12674 := PrimIsPair(W3192)

if True == tmp12674 {
tmp12660 := MakeNative(func(__e *ControlFlow) {
W3193 := __e.Get(1)
_ = W3193
tmp12661 := MakeNative(func(__e *ControlFlow) {
W3194 := __e.Get(1)
_ = W3194
tmp12670 := PrimIsPair(W3194)

if True == tmp12670 {
tmp12662 := MakeNative(func(__e *ControlFlow) {
W3195 := __e.Get(1)
_ = W3195
tmp12663 := MakeNative(func(__e *ControlFlow) {
W3196 := __e.Get(1)
_ = W3196
tmp12664 := PrimCons(W3195, Nil)

tmp12665 := PrimCons(W3193, tmp12664)

tmp12666 := PrimCons(symlet, tmp12665)

__e.TailApply(PrimFunc(symshen_4comb), W3196, tmp12666)
return


}, 1)

tmp12667 := Call(__e, PrimFunc(symtail), W3194)


__e.TailApply(tmp12663, tmp12667)
return


}, 1)

tmp12668 := Call(__e, PrimFunc(symhead), W3194)


__e.TailApply(tmp12662, tmp12668)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12671 := Call(__e, PrimFunc(symtail), W3192)


__e.TailApply(tmp12661, tmp12671)
return


}, 1)

tmp12672 := Call(__e, PrimFunc(symhead), W3192)


__e.TailApply(tmp12660, tmp12672)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12675 := Call(__e, PrimFunc(symtail), V3186)


tmp12676 := Call(__e, tmp12659, tmp12675)


ifres12658 = tmp12676


} else {
tmp12677 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12658 = tmp12677


}

__e.TailApply(tmp12631, ifres12658)
return


} else {
__e.Return(W3187)
return
}


}, 1)

tmp12694 := Call(__e, PrimFunc(symshen_4hds_a_2), V3186, symif)


var ifres12681 Obj

if True == tmp12694 {
tmp12682 := MakeNative(func(__e *ControlFlow) {
W3188 := __e.Get(1)
_ = W3188
tmp12690 := PrimIsPair(W3188)

if True == tmp12690 {
tmp12683 := MakeNative(func(__e *ControlFlow) {
W3189 := __e.Get(1)
_ = W3189
tmp12684 := MakeNative(func(__e *ControlFlow) {
W3190 := __e.Get(1)
_ = W3190
tmp12685 := PrimCons(W3189, Nil)

tmp12686 := PrimCons(symif, tmp12685)

__e.TailApply(PrimFunc(symshen_4comb), W3190, tmp12686)
return


}, 1)

tmp12687 := Call(__e, PrimFunc(symtail), W3188)


__e.TailApply(tmp12684, tmp12687)
return


}, 1)

tmp12688 := Call(__e, PrimFunc(symhead), W3188)


__e.TailApply(tmp12683, tmp12688)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12691 := Call(__e, PrimFunc(symtail), V3186)


tmp12692 := Call(__e, tmp12682, tmp12691)


ifres12681 = tmp12692


} else {
tmp12693 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12681 = tmp12693


}

__e.TailApply(tmp12630, ifres12681)
return


}, 1)

tmp12695 := Call(__e, ns2_1set, symshen_4_5side_6, tmp12629)


_ = tmp12695

tmp12696 := MakeNative(func(__e *ControlFlow) {
V3209 := __e.Get(1)
_ = V3209
V3210 := __e.Get(2)
_ = V3210
V3211 := __e.Get(3)
_ = V3211
tmp12731 := PrimIsPair(V3211)

var ifres12718 Obj

if True == tmp12731 {
tmp12729 := PrimHead(V3211)

tmp12730 := PrimEqual(Nil, tmp12729)

var ifres12720 Obj

if True == tmp12730 {
tmp12727 := PrimTail(V3211)

tmp12728 := PrimIsPair(tmp12727)

var ifres12722 Obj

if True == tmp12728 {
tmp12724 := PrimTail(V3211)

tmp12725 := PrimTail(tmp12724)

tmp12726 := PrimEqual(Nil, tmp12725)

var ifres12723 Obj

if True == tmp12726 {
ifres12723 = True


} else {
ifres12723 = False


}

ifres12722 = ifres12723


} else {
ifres12722 = False


}

var ifres12721 Obj

if True == ifres12722 {
ifres12721 = True


} else {
ifres12721 = False


}

ifres12720 = ifres12721


} else {
ifres12720 = False


}

var ifres12719 Obj

if True == ifres12720 {
ifres12719 = True


} else {
ifres12719 = False


}

ifres12718 = ifres12719


} else {
ifres12718 = False


}

if True == ifres12718 {
tmp12697 := MakeNative(func(__e *ControlFlow) {
W3212 := __e.Get(1)
_ = W3212
tmp12698 := MakeNative(func(__e *ControlFlow) {
W3213 := __e.Get(1)
_ = W3213
tmp12699 := MakeNative(func(__e *ControlFlow) {
W3214 := __e.Get(1)
_ = W3214
tmp12700 := MakeNative(func(__e *ControlFlow) {
W3215 := __e.Get(1)
_ = W3215
tmp12701 := MakeNative(func(__e *ControlFlow) {
W3216 := __e.Get(1)
_ = W3216
tmp12702 := PrimCons(W3215, Nil)

__e.Return(PrimCons(W3216, tmp12702))
return


}, 1)

tmp12703 := PrimCons(V3211, Nil)

tmp12704 := PrimCons(V3210, tmp12703)

tmp12705 := PrimCons(V3209, tmp12704)

__e.TailApply(tmp12701, tmp12705)
return


}, 1)

tmp12706 := PrimCons(W3214, Nil)

tmp12707 := PrimCons(W3213, Nil)

tmp12708 := PrimCons(tmp12706, tmp12707)

tmp12709 := PrimCons(V3209, tmp12708)

__e.TailApply(tmp12700, tmp12709)
return


}, 1)

tmp12710 := Call(__e, PrimFunc(symshen_4coll_1formulae), V3210)


tmp12711 := PrimCons(W3212, Nil)

tmp12712 := PrimCons(tmp12710, tmp12711)

__e.TailApply(tmp12699, tmp12712)
return


}, 1)

tmp12713 := PrimTail(V3211)

tmp12714 := PrimCons(W3212, Nil)

tmp12715 := PrimCons(tmp12713, tmp12714)

__e.TailApply(tmp12698, tmp12715)
return


}, 1)

tmp12716 := Call(__e, PrimFunc(symgensym), symP)


__e.TailApply(tmp12697, tmp12716)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.lr-rule")))
return
}


}, 3)

tmp12732 := Call(__e, ns2_1set, symshen_4lr_1rule, tmp12696)


_ = tmp12732

tmp12733 := MakeNative(func(__e *ControlFlow) {
V3219 := __e.Get(1)
_ = V3219
tmp12762 := PrimEqual(Nil, V3219)

if True == tmp12762 {
__e.Return(Nil)
return
} else {
tmp12760 := PrimIsPair(V3219)

var ifres12740 Obj

if True == tmp12760 {
tmp12758 := PrimHead(V3219)

tmp12759 := PrimIsPair(tmp12758)

var ifres12742 Obj

if True == tmp12759 {
tmp12755 := PrimHead(V3219)

tmp12756 := PrimHead(tmp12755)

tmp12757 := PrimEqual(Nil, tmp12756)

var ifres12744 Obj

if True == tmp12757 {
tmp12752 := PrimHead(V3219)

tmp12753 := PrimTail(tmp12752)

tmp12754 := PrimIsPair(tmp12753)

var ifres12746 Obj

if True == tmp12754 {
tmp12748 := PrimHead(V3219)

tmp12749 := PrimTail(tmp12748)

tmp12750 := PrimTail(tmp12749)

tmp12751 := PrimEqual(Nil, tmp12750)

var ifres12747 Obj

if True == tmp12751 {
ifres12747 = True


} else {
ifres12747 = False


}

ifres12746 = ifres12747


} else {
ifres12746 = False


}

var ifres12745 Obj

if True == ifres12746 {
ifres12745 = True


} else {
ifres12745 = False


}

ifres12744 = ifres12745


} else {
ifres12744 = False


}

var ifres12743 Obj

if True == ifres12744 {
ifres12743 = True


} else {
ifres12743 = False


}

ifres12742 = ifres12743


} else {
ifres12742 = False


}

var ifres12741 Obj

if True == ifres12742 {
ifres12741 = True


} else {
ifres12741 = False


}

ifres12740 = ifres12741


} else {
ifres12740 = False


}

if True == ifres12740 {
tmp12734 := PrimHead(V3219)

tmp12735 := PrimTail(tmp12734)

tmp12736 := PrimHead(tmp12735)

tmp12737 := PrimTail(V3219)

tmp12738 := Call(__e, PrimFunc(symshen_4coll_1formulae), tmp12737)


__e.Return(PrimCons(tmp12736, tmp12738))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.coll-formulae")))
return
}


}


}, 1)

tmp12763 := Call(__e, ns2_1set, symshen_4coll_1formulae, tmp12733)


_ = tmp12763

tmp12764 := MakeNative(func(__e *ControlFlow) {
V3220 := __e.Get(1)
_ = V3220
tmp12765 := MakeNative(func(__e *ControlFlow) {
W3221 := __e.Get(1)
_ = W3221
tmp12767 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3221)


if True == tmp12767 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3221)
return
}


}, 1)

tmp12779 := PrimIsPair(V3220)

var ifres12768 Obj

if True == tmp12779 {
tmp12769 := MakeNative(func(__e *ControlFlow) {
W3222 := __e.Get(1)
_ = W3222
tmp12770 := MakeNative(func(__e *ControlFlow) {
W3223 := __e.Get(1)
_ = W3223
tmp12773 := Call(__e, PrimFunc(symshen_4key_1in_1sequent_1calculus_2), W3222)


tmp12774 := PrimNot(tmp12773)

if True == tmp12774 {
tmp12771 := Call(__e, PrimFunc(symmacroexpand), W3222)


__e.TailApply(PrimFunc(symshen_4comb), W3223, tmp12771)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12775 := Call(__e, PrimFunc(symtail), V3220)


__e.TailApply(tmp12770, tmp12775)
return


}, 1)

tmp12776 := Call(__e, PrimFunc(symhead), V3220)


tmp12777 := Call(__e, tmp12769, tmp12776)


ifres12768 = tmp12777


} else {
tmp12778 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12768 = tmp12778


}

__e.TailApply(tmp12765, ifres12768)
return


}, 1)

tmp12780 := Call(__e, ns2_1set, symshen_4_5expr_6, tmp12764)


_ = tmp12780

tmp12781 := MakeNative(func(__e *ControlFlow) {
V3224 := __e.Get(1)
_ = V3224
tmp12788 := PrimIntern(MakeString(";"))

tmp12789 := PrimIntern(MakeString(","))

tmp12790 := PrimIntern(MakeString(":"))

tmp12791 := PrimCons(sym_5_1_1, Nil)

tmp12792 := PrimCons(tmp12790, tmp12791)

tmp12793 := PrimCons(tmp12789, tmp12792)

tmp12794 := PrimCons(tmp12788, tmp12793)

tmp12795 := PrimCons(sym_6_6, tmp12794)

tmp12796 := Call(__e, PrimFunc(symelement_2), V3224, tmp12795)


if True == tmp12796 {
__e.Return(True)
return
} else {
tmp12786 := Call(__e, PrimFunc(symshen_4sng_2), V3224)


var ifres12783 Obj

if True == tmp12786 {
ifres12783 = True


} else {
tmp12785 := Call(__e, PrimFunc(symshen_4dbl_2), V3224)


var ifres12784 Obj

if True == tmp12785 {
ifres12784 = True


} else {
ifres12784 = False


}

ifres12783 = ifres12784


}

if True == ifres12783 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp12797 := Call(__e, ns2_1set, symshen_4key_1in_1sequent_1calculus_2, tmp12781)


_ = tmp12797

tmp12798 := MakeNative(func(__e *ControlFlow) {
V3225 := __e.Get(1)
_ = V3225
tmp12799 := MakeNative(func(__e *ControlFlow) {
W3226 := __e.Get(1)
_ = W3226
tmp12801 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3226)


if True == tmp12801 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3226)
return
}


}, 1)

tmp12802 := MakeNative(func(__e *ControlFlow) {
W3227 := __e.Get(1)
_ = W3227
tmp12808 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3227)


if True == tmp12808 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12803 := MakeNative(func(__e *ControlFlow) {
W3228 := __e.Get(1)
_ = W3228
tmp12804 := MakeNative(func(__e *ControlFlow) {
W3229 := __e.Get(1)
_ = W3229
__e.TailApply(PrimFunc(symshen_4comb), W3229, W3228)
return
}, 1)

tmp12805 := Call(__e, PrimFunc(symshen_4in_1_6), W3227)


__e.TailApply(tmp12804, tmp12805)
return


}, 1)

tmp12806 := Call(__e, PrimFunc(symshen_4_5_1out), W3227)


__e.TailApply(tmp12803, tmp12806)
return


}


}, 1)

tmp12809 := Call(__e, PrimFunc(symshen_4_5expr_6), V3225)


tmp12810 := Call(__e, tmp12802, tmp12809)


__e.TailApply(tmp12799, tmp12810)
return


}, 1)

tmp12811 := Call(__e, ns2_1set, symshen_4_5type_6, tmp12798)


_ = tmp12811

tmp12812 := MakeNative(func(__e *ControlFlow) {
V3230 := __e.Get(1)
_ = V3230
tmp12813 := MakeNative(func(__e *ControlFlow) {
W3231 := __e.Get(1)
_ = W3231
tmp12815 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3231)


if True == tmp12815 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3231)
return
}


}, 1)

tmp12825 := PrimIsPair(V3230)

var ifres12816 Obj

if True == tmp12825 {
tmp12817 := MakeNative(func(__e *ControlFlow) {
W3232 := __e.Get(1)
_ = W3232
tmp12818 := MakeNative(func(__e *ControlFlow) {
W3233 := __e.Get(1)
_ = W3233
tmp12820 := Call(__e, PrimFunc(symshen_4dbl_2), W3232)


if True == tmp12820 {
__e.TailApply(PrimFunc(symshen_4comb), W3233, W3232)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12821 := Call(__e, PrimFunc(symtail), V3230)


__e.TailApply(tmp12818, tmp12821)
return


}, 1)

tmp12822 := Call(__e, PrimFunc(symhead), V3230)


tmp12823 := Call(__e, tmp12817, tmp12822)


ifres12816 = tmp12823


} else {
tmp12824 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12816 = tmp12824


}

__e.TailApply(tmp12813, ifres12816)
return


}, 1)

tmp12826 := Call(__e, ns2_1set, symshen_4_5dbl_6, tmp12812)


_ = tmp12826

tmp12827 := MakeNative(func(__e *ControlFlow) {
V3234 := __e.Get(1)
_ = V3234
tmp12828 := MakeNative(func(__e *ControlFlow) {
W3235 := __e.Get(1)
_ = W3235
tmp12830 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3235)


if True == tmp12830 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3235)
return
}


}, 1)

tmp12840 := PrimIsPair(V3234)

var ifres12831 Obj

if True == tmp12840 {
tmp12832 := MakeNative(func(__e *ControlFlow) {
W3236 := __e.Get(1)
_ = W3236
tmp12833 := MakeNative(func(__e *ControlFlow) {
W3237 := __e.Get(1)
_ = W3237
tmp12835 := Call(__e, PrimFunc(symshen_4sng_2), W3236)


if True == tmp12835 {
__e.TailApply(PrimFunc(symshen_4comb), W3237, W3236)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12836 := Call(__e, PrimFunc(symtail), V3234)


__e.TailApply(tmp12833, tmp12836)
return


}, 1)

tmp12837 := Call(__e, PrimFunc(symhead), V3234)


tmp12838 := Call(__e, tmp12832, tmp12837)


ifres12831 = tmp12838


} else {
tmp12839 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12831 = tmp12839


}

__e.TailApply(tmp12828, ifres12831)
return


}, 1)

tmp12841 := Call(__e, ns2_1set, symshen_4_5sng_6, tmp12827)


_ = tmp12841

tmp12842 := MakeNative(func(__e *ControlFlow) {
V3238 := __e.Get(1)
_ = V3238
tmp12847 := PrimIsSymbol(V3238)

if True == tmp12847 {
tmp12844 := PrimStr(V3238)

tmp12845 := Call(__e, PrimFunc(symshen_4sng_1h_2), tmp12844)


if True == tmp12845 {
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

tmp12848 := Call(__e, ns2_1set, symshen_4sng_2, tmp12842)


_ = tmp12848

tmp12849 := MakeNative(func(__e *ControlFlow) {
V3241 := __e.Get(1)
_ = V3241
tmp12858 := PrimEqual(MakeString("___"), V3241)

if True == tmp12858 {
__e.Return(True)
return
} else {
tmp12856 := Call(__e, PrimFunc(symshen_4_7string_2), V3241)


var ifres12852 Obj

if True == tmp12856 {
tmp12854 := Call(__e, PrimFunc(symhdstr), V3241)


tmp12855 := PrimEqual(MakeString("_"), tmp12854)

var ifres12853 Obj

if True == tmp12855 {
ifres12853 = True


} else {
ifres12853 = False


}

ifres12852 = ifres12853


} else {
ifres12852 = False


}

if True == ifres12852 {
tmp12850 := PrimTailString(V3241)

__e.TailApply(PrimFunc(symshen_4sng_1h_2), tmp12850)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp12859 := Call(__e, ns2_1set, symshen_4sng_1h_2, tmp12849)


_ = tmp12859

tmp12860 := MakeNative(func(__e *ControlFlow) {
V3242 := __e.Get(1)
_ = V3242
tmp12865 := PrimIsSymbol(V3242)

if True == tmp12865 {
tmp12862 := PrimStr(V3242)

tmp12863 := Call(__e, PrimFunc(symshen_4dbl_1h_2), tmp12862)


if True == tmp12863 {
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

tmp12866 := Call(__e, ns2_1set, symshen_4dbl_2, tmp12860)


_ = tmp12866

tmp12867 := MakeNative(func(__e *ControlFlow) {
V3245 := __e.Get(1)
_ = V3245
tmp12876 := PrimEqual(MakeString("==="), V3245)

if True == tmp12876 {
__e.Return(True)
return
} else {
tmp12874 := Call(__e, PrimFunc(symshen_4_7string_2), V3245)


var ifres12870 Obj

if True == tmp12874 {
tmp12872 := Call(__e, PrimFunc(symhdstr), V3245)


tmp12873 := PrimEqual(MakeString("="), tmp12872)

var ifres12871 Obj

if True == tmp12873 {
ifres12871 = True


} else {
ifres12871 = False


}

ifres12870 = ifres12871


} else {
ifres12870 = False


}

if True == ifres12870 {
tmp12868 := PrimTailString(V3245)

__e.TailApply(PrimFunc(symshen_4dbl_1h_2), tmp12868)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp12877 := Call(__e, ns2_1set, symshen_4dbl_1h_2, tmp12867)


_ = tmp12877

tmp12878 := MakeNative(func(__e *ControlFlow) {
V3246 := __e.Get(1)
_ = V3246
V3247 := __e.Get(2)
_ = V3247
tmp12879 := PrimValue(symshen_4_ddatatypes_d)

tmp12880 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3246, V3247, tmp12879)


tmp12881 := PrimSet(symshen_4_ddatatypes_d, tmp12880)

_ = tmp12881

tmp12882 := PrimValue(symshen_4_dalldatatypes_d)

tmp12883 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3246, V3247, tmp12882)


tmp12884 := PrimSet(symshen_4_dalldatatypes_d, tmp12883)

_ = tmp12884

__e.Return(V3246)
return


}, 2)

tmp12885 := Call(__e, ns2_1set, symshen_4remember_1datatype, tmp12878)


_ = tmp12885

tmp12886 := MakeNative(func(__e *ControlFlow) {
V3248 := __e.Get(1)
_ = V3248
V3249 := __e.Get(2)
_ = V3249
tmp12887 := MakeNative(func(__e *ControlFlow) {
W3250 := __e.Get(1)
_ = W3250
tmp12888 := PrimCons(V3248, W3250)

tmp12889 := PrimCons(symdefprolog, tmp12888)

__e.TailApply(PrimFunc(symeval), tmp12889)
return


}, 1)

tmp12890 := MakeNative(func(__e *ControlFlow) {
Z3251 := __e.Get(1)
_ = Z3251
__e.TailApply(PrimFunc(symshen_4rule_1_6clause), Z3251)
return
}, 1)

tmp12891 := Call(__e, PrimFunc(symmapcan), tmp12890, V3249)


__e.TailApply(tmp12887, tmp12891)
return


}, 2)

tmp12892 := Call(__e, ns2_1set, symshen_4rules_1_6prolog, tmp12886)


_ = tmp12892

tmp12893 := MakeNative(func(__e *ControlFlow) {
V3254 := __e.Get(1)
_ = V3254
tmp12970 := PrimIsPair(V3254)

var ifres12934 Obj

if True == tmp12970 {
tmp12968 := PrimTail(V3254)

tmp12969 := PrimIsPair(tmp12968)

var ifres12936 Obj

if True == tmp12969 {
tmp12965 := PrimTail(V3254)

tmp12966 := PrimTail(tmp12965)

tmp12967 := PrimIsPair(tmp12966)

var ifres12938 Obj

if True == tmp12967 {
tmp12961 := PrimTail(V3254)

tmp12962 := PrimTail(tmp12961)

tmp12963 := PrimHead(tmp12962)

tmp12964 := PrimIsPair(tmp12963)

var ifres12940 Obj

if True == tmp12964 {
tmp12956 := PrimTail(V3254)

tmp12957 := PrimTail(tmp12956)

tmp12958 := PrimHead(tmp12957)

tmp12959 := PrimTail(tmp12958)

tmp12960 := PrimIsPair(tmp12959)

var ifres12942 Obj

if True == tmp12960 {
tmp12950 := PrimTail(V3254)

tmp12951 := PrimTail(tmp12950)

tmp12952 := PrimHead(tmp12951)

tmp12953 := PrimTail(tmp12952)

tmp12954 := PrimTail(tmp12953)

tmp12955 := PrimEqual(Nil, tmp12954)

var ifres12944 Obj

if True == tmp12955 {
tmp12946 := PrimTail(V3254)

tmp12947 := PrimTail(tmp12946)

tmp12948 := PrimTail(tmp12947)

tmp12949 := PrimEqual(Nil, tmp12948)

var ifres12945 Obj

if True == tmp12949 {
ifres12945 = True


} else {
ifres12945 = False


}

ifres12944 = ifres12945


} else {
ifres12944 = False


}

var ifres12943 Obj

if True == ifres12944 {
ifres12943 = True


} else {
ifres12943 = False


}

ifres12942 = ifres12943


} else {
ifres12942 = False


}

var ifres12941 Obj

if True == ifres12942 {
ifres12941 = True


} else {
ifres12941 = False


}

ifres12940 = ifres12941


} else {
ifres12940 = False


}

var ifres12939 Obj

if True == ifres12940 {
ifres12939 = True


} else {
ifres12939 = False


}

ifres12938 = ifres12939


} else {
ifres12938 = False


}

var ifres12937 Obj

if True == ifres12938 {
ifres12937 = True


} else {
ifres12937 = False


}

ifres12936 = ifres12937


} else {
ifres12936 = False


}

var ifres12935 Obj

if True == ifres12936 {
ifres12935 = True


} else {
ifres12935 = False


}

ifres12934 = ifres12935


} else {
ifres12934 = False


}

if True == ifres12934 {
tmp12894 := MakeNative(func(__e *ControlFlow) {
W3255 := __e.Get(1)
_ = W3255
tmp12895 := MakeNative(func(__e *ControlFlow) {
W3256 := __e.Get(1)
_ = W3256
tmp12896 := MakeNative(func(__e *ControlFlow) {
W3257 := __e.Get(1)
_ = W3257
tmp12897 := MakeNative(func(__e *ControlFlow) {
W3258 := __e.Get(1)
_ = W3258
tmp12898 := MakeNative(func(__e *ControlFlow) {
W3259 := __e.Get(1)
_ = W3259
tmp12899 := PrimCons(sym_5_1_1, Nil)

tmp12900 := PrimIntern(MakeString(";"))

tmp12901 := PrimCons(tmp12900, Nil)

tmp12902 := Call(__e, PrimFunc(symappend), W3259, tmp12901)


tmp12903 := Call(__e, PrimFunc(symappend), tmp12899, tmp12902)


__e.TailApply(PrimFunc(symappend), W3258, tmp12903)
return


}, 1)

tmp12904 := PrimTail(V3254)

tmp12905 := PrimTail(tmp12904)

tmp12906 := PrimHead(tmp12905)

tmp12907 := PrimHead(tmp12906)

tmp12908 := PrimHead(V3254)

tmp12909 := PrimTail(V3254)

tmp12910 := PrimHead(tmp12909)

tmp12911 := Call(__e, PrimFunc(symshen_4goals), W3255, tmp12907, tmp12908, tmp12910, W3256, W3257)


__e.TailApply(tmp12898, tmp12911)
return


}, 1)

tmp12912 := PrimTail(V3254)

tmp12913 := PrimTail(tmp12912)

tmp12914 := PrimHead(tmp12913)

tmp12915 := PrimTail(tmp12914)

tmp12916 := PrimHead(tmp12915)

tmp12917 := Call(__e, PrimFunc(symshen_4compile_1consequent), tmp12916, W3256)


__e.TailApply(tmp12897, tmp12917)
return


}, 1)

tmp12918 := PrimTail(V3254)

tmp12919 := PrimTail(tmp12918)

tmp12920 := PrimHead(tmp12919)

tmp12921 := PrimTail(tmp12920)

tmp12922 := PrimHead(tmp12921)

tmp12923 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp12922)


__e.TailApply(tmp12896, tmp12923)
return


}, 1)

tmp12924 := PrimTail(V3254)

tmp12925 := PrimTail(tmp12924)

tmp12926 := PrimHead(tmp12925)

tmp12927 := PrimHead(tmp12926)

tmp12928 := Call(__e, PrimFunc(symlength), tmp12927)


tmp12929 := Call(__e, PrimFunc(symshen_4nvars), tmp12928)


tmp12930 := PrimCons(symDelta, Nil)

tmp12931 := Call(__e, PrimFunc(symappend), tmp12929, tmp12930)


__e.TailApply(tmp12895, tmp12931)
return


}, 1)

tmp12932 := Call(__e, PrimFunc(symshen_4extract_1vars), V3254)


__e.TailApply(tmp12894, tmp12932)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.rule->clause")))
return
}


}, 1)

tmp12971 := Call(__e, ns2_1set, symshen_4rule_1_6clause, tmp12893)


_ = tmp12971

tmp12972 := MakeNative(func(__e *ControlFlow) {
V3266 := __e.Get(1)
_ = V3266
V3267 := __e.Get(2)
_ = V3267
tmp12977 := PrimIsPair(V3267)

if True == tmp12977 {
tmp12973 := Call(__e, PrimFunc(symshen_4optimise_1typing), V3266)


tmp12974 := PrimHead(V3267)

tmp12975 := PrimCons(tmp12974, Nil)

__e.Return(PrimCons(tmp12973, tmp12975))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-consequent")))
return
}


}, 2)

tmp12978 := Call(__e, ns2_1set, symshen_4compile_1consequent, tmp12972)


_ = tmp12978

tmp12979 := MakeNative(func(__e *ControlFlow) {
V3268 := __e.Get(1)
_ = V3268
tmp12984 := PrimEqual(MakeNumber(0), V3268)

if True == tmp12984 {
__e.Return(Nil)
return
} else {
tmp12980 := Call(__e, PrimFunc(symgensym), symV)


tmp12981 := PrimNumberSubtract(V3268, MakeNumber(1))

tmp12982 := Call(__e, PrimFunc(symshen_4nvars), tmp12981)


__e.Return(PrimCons(tmp12980, tmp12982))
return


}


}, 1)

tmp12985 := Call(__e, ns2_1set, symshen_4nvars, tmp12979)


_ = tmp12985

tmp12986 := MakeNative(func(__e *ControlFlow) {
V3269 := __e.Get(1)
_ = V3269
tmp13026 := PrimIsPair(V3269)

var ifres13005 Obj

if True == tmp13026 {
tmp13024 := PrimTail(V3269)

tmp13025 := PrimIsPair(tmp13024)

var ifres13007 Obj

if True == tmp13025 {
tmp13021 := PrimTail(V3269)

tmp13022 := PrimTail(tmp13021)

tmp13023 := PrimIsPair(tmp13022)

var ifres13009 Obj

if True == tmp13023 {
tmp13017 := PrimTail(V3269)

tmp13018 := PrimTail(tmp13017)

tmp13019 := PrimTail(tmp13018)

tmp13020 := PrimEqual(Nil, tmp13019)

var ifres13011 Obj

if True == tmp13020 {
tmp13013 := PrimTail(V3269)

tmp13014 := PrimHead(tmp13013)

tmp13015 := PrimIntern(MakeString(":"))

tmp13016 := PrimEqual(tmp13014, tmp13015)

var ifres13012 Obj

if True == tmp13016 {
ifres13012 = True


} else {
ifres13012 = False


}

ifres13011 = ifres13012


} else {
ifres13011 = False


}

var ifres13010 Obj

if True == ifres13011 {
ifres13010 = True


} else {
ifres13010 = False


}

ifres13009 = ifres13010


} else {
ifres13009 = False


}

var ifres13008 Obj

if True == ifres13009 {
ifres13008 = True


} else {
ifres13008 = False


}

ifres13007 = ifres13008


} else {
ifres13007 = False


}

var ifres13006 Obj

if True == ifres13007 {
ifres13006 = True


} else {
ifres13006 = False


}

ifres13005 = ifres13006


} else {
ifres13005 = False


}

if True == ifres13005 {
tmp12987 := MakeNative(func(__e *ControlFlow) {
W3270 := __e.Get(1)
_ = W3270
__e.TailApply(PrimFunc(symshen_4cons_1form_1with_1modes), W3270)
return
}, 1)

tmp12988 := PrimHead(V3269)

tmp12989 := PrimTail(V3269)

tmp12990 := PrimHead(tmp12989)

tmp12991 := PrimTail(V3269)

tmp12992 := PrimTail(tmp12991)

tmp12993 := PrimCons(sym_7, tmp12992)

tmp12994 := PrimCons(tmp12993, Nil)

tmp12995 := PrimCons(tmp12990, tmp12994)

tmp12996 := PrimCons(tmp12988, tmp12995)

tmp12997 := PrimCons(tmp12996, Nil)

tmp12998 := PrimCons(sym_1, tmp12997)

tmp12999 := Call(__e, PrimFunc(symshen_4expand_1mode_1forms), tmp12998)


__e.TailApply(tmp12987, tmp12999)
return


} else {
tmp13000 := MakeNative(func(__e *ControlFlow) {
W3271 := __e.Get(1)
_ = W3271
__e.TailApply(PrimFunc(symshen_4cons_1form_1with_1modes), W3271)
return
}, 1)

tmp13001 := PrimCons(V3269, Nil)

tmp13002 := PrimCons(sym_7, tmp13001)

tmp13003 := Call(__e, PrimFunc(symshen_4expand_1mode_1forms), tmp13002)


__e.TailApply(tmp13000, tmp13003)
return


}


}, 1)

tmp13027 := Call(__e, ns2_1set, symshen_4optimise_1typing, tmp12986)


_ = tmp13027

tmp13028 := MakeNative(func(__e *ControlFlow) {
V3272 := __e.Get(1)
_ = V3272
tmp13071 := PrimIsPair(V3272)

var ifres13058 Obj

if True == tmp13071 {
tmp13069 := PrimHead(V3272)

tmp13070 := PrimEqual(sym_7, tmp13069)

var ifres13060 Obj

if True == tmp13070 {
tmp13067 := PrimTail(V3272)

tmp13068 := PrimIsPair(tmp13067)

var ifres13062 Obj

if True == tmp13068 {
tmp13064 := PrimTail(V3272)

tmp13065 := PrimTail(tmp13064)

tmp13066 := PrimEqual(Nil, tmp13065)

var ifres13063 Obj

if True == tmp13066 {
ifres13063 = True


} else {
ifres13063 = False


}

ifres13062 = ifres13063


} else {
ifres13062 = False


}

var ifres13061 Obj

if True == ifres13062 {
ifres13061 = True


} else {
ifres13061 = False


}

ifres13060 = ifres13061


} else {
ifres13060 = False


}

var ifres13059 Obj

if True == ifres13060 {
ifres13059 = True


} else {
ifres13059 = False


}

ifres13058 = ifres13059


} else {
ifres13058 = False


}

if True == ifres13058 {
tmp13029 := PrimTail(V3272)

tmp13030 := PrimHead(tmp13029)

tmp13031 := Call(__e, PrimFunc(symshen_4expand_1mode_1forms), tmp13030)


tmp13032 := PrimCons(sym_7, Nil)

tmp13033 := PrimCons(tmp13031, tmp13032)

__e.Return(PrimCons(symmode, tmp13033))
return


} else {
tmp13056 := PrimIsPair(V3272)

var ifres13043 Obj

if True == tmp13056 {
tmp13054 := PrimHead(V3272)

tmp13055 := PrimEqual(sym_1, tmp13054)

var ifres13045 Obj

if True == tmp13055 {
tmp13052 := PrimTail(V3272)

tmp13053 := PrimIsPair(tmp13052)

var ifres13047 Obj

if True == tmp13053 {
tmp13049 := PrimTail(V3272)

tmp13050 := PrimTail(tmp13049)

tmp13051 := PrimEqual(Nil, tmp13050)

var ifres13048 Obj

if True == tmp13051 {
ifres13048 = True


} else {
ifres13048 = False


}

ifres13047 = ifres13048


} else {
ifres13047 = False


}

var ifres13046 Obj

if True == ifres13047 {
ifres13046 = True


} else {
ifres13046 = False


}

ifres13045 = ifres13046


} else {
ifres13045 = False


}

var ifres13044 Obj

if True == ifres13045 {
ifres13044 = True


} else {
ifres13044 = False


}

ifres13043 = ifres13044


} else {
ifres13043 = False


}

if True == ifres13043 {
tmp13034 := PrimTail(V3272)

tmp13035 := PrimHead(tmp13034)

tmp13036 := Call(__e, PrimFunc(symshen_4expand_1mode_1forms), tmp13035)


tmp13037 := PrimCons(sym_1, Nil)

tmp13038 := PrimCons(tmp13036, tmp13037)

__e.Return(PrimCons(symmode, tmp13038))
return


} else {
tmp13041 := PrimIsPair(V3272)

if True == tmp13041 {
tmp13039 := MakeNative(func(__e *ControlFlow) {
Z3273 := __e.Get(1)
_ = Z3273
__e.TailApply(PrimFunc(symshen_4expand_1mode_1forms), Z3273)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13039, V3272)
return


} else {
__e.Return(V3272)
return
}


}


}


}, 1)

tmp13072 := Call(__e, ns2_1set, symshen_4expand_1mode_1forms, tmp13028)


_ = tmp13072

tmp13073 := MakeNative(func(__e *ControlFlow) {
V3274 := __e.Get(1)
_ = V3274
tmp13125 := PrimIsPair(V3274)

var ifres13106 Obj

if True == tmp13125 {
tmp13123 := PrimHead(V3274)

tmp13124 := PrimEqual(symmode, tmp13123)

var ifres13108 Obj

if True == tmp13124 {
tmp13121 := PrimTail(V3274)

tmp13122 := PrimIsPair(tmp13121)

var ifres13110 Obj

if True == tmp13122 {
tmp13118 := PrimTail(V3274)

tmp13119 := PrimTail(tmp13118)

tmp13120 := PrimIsPair(tmp13119)

var ifres13112 Obj

if True == tmp13120 {
tmp13114 := PrimTail(V3274)

tmp13115 := PrimTail(tmp13114)

tmp13116 := PrimTail(tmp13115)

tmp13117 := PrimEqual(Nil, tmp13116)

var ifres13113 Obj

if True == tmp13117 {
ifres13113 = True


} else {
ifres13113 = False


}

ifres13112 = ifres13113


} else {
ifres13112 = False


}

var ifres13111 Obj

if True == ifres13112 {
ifres13111 = True


} else {
ifres13111 = False


}

ifres13110 = ifres13111


} else {
ifres13110 = False


}

var ifres13109 Obj

if True == ifres13110 {
ifres13109 = True


} else {
ifres13109 = False


}

ifres13108 = ifres13109


} else {
ifres13108 = False


}

var ifres13107 Obj

if True == ifres13108 {
ifres13107 = True


} else {
ifres13107 = False


}

ifres13106 = ifres13107


} else {
ifres13106 = False


}

if True == ifres13106 {
tmp13074 := PrimTail(V3274)

tmp13075 := PrimTail(tmp13074)

tmp13076 := PrimHead(tmp13075)

tmp13077 := PrimTail(V3274)

tmp13078 := PrimHead(tmp13077)

tmp13079 := Call(__e, PrimFunc(symshen_4cons_1form_1with_1modes), tmp13078)


tmp13080 := PrimCons(tmp13079, Nil)

__e.Return(PrimCons(tmp13076, tmp13080))
return


} else {
tmp13104 := PrimIsPair(V3274)

var ifres13091 Obj

if True == tmp13104 {
tmp13102 := PrimHead(V3274)

tmp13103 := PrimEqual(symbar_b, tmp13102)

var ifres13093 Obj

if True == tmp13103 {
tmp13100 := PrimTail(V3274)

tmp13101 := PrimIsPair(tmp13100)

var ifres13095 Obj

if True == tmp13101 {
tmp13097 := PrimTail(V3274)

tmp13098 := PrimTail(tmp13097)

tmp13099 := PrimEqual(Nil, tmp13098)

var ifres13096 Obj

if True == tmp13099 {
ifres13096 = True


} else {
ifres13096 = False


}

ifres13095 = ifres13096


} else {
ifres13095 = False


}

var ifres13094 Obj

if True == ifres13095 {
ifres13094 = True


} else {
ifres13094 = False


}

ifres13093 = ifres13094


} else {
ifres13093 = False


}

var ifres13092 Obj

if True == ifres13093 {
ifres13092 = True


} else {
ifres13092 = False


}

ifres13091 = ifres13092


} else {
ifres13091 = False


}

if True == ifres13091 {
tmp13081 := PrimTail(V3274)

__e.Return(PrimHead(tmp13081))
return


} else {
tmp13089 := PrimIsPair(V3274)

if True == tmp13089 {
tmp13082 := PrimHead(V3274)

tmp13083 := Call(__e, PrimFunc(symshen_4cons_1form_1with_1modes), tmp13082)


tmp13084 := PrimTail(V3274)

tmp13085 := Call(__e, PrimFunc(symshen_4cons_1form_1with_1modes), tmp13084)


tmp13086 := PrimCons(tmp13085, Nil)

tmp13087 := PrimCons(tmp13083, tmp13086)

__e.Return(PrimCons(symcons, tmp13087))
return


} else {
__e.Return(V3274)
return
}


}


}


}, 1)

tmp13126 := Call(__e, ns2_1set, symshen_4cons_1form_1with_1modes, tmp13073)


_ = tmp13126

tmp13127 := MakeNative(func(__e *ControlFlow) {
V3275 := __e.Get(1)
_ = V3275
V3276 := __e.Get(2)
_ = V3276
V3277 := __e.Get(3)
_ = V3277
V3278 := __e.Get(4)
_ = V3278
V3279 := __e.Get(5)
_ = V3279
V3280 := __e.Get(6)
_ = V3280
tmp13128 := MakeNative(func(__e *ControlFlow) {
W3281 := __e.Get(1)
_ = W3281
tmp13129 := MakeNative(func(__e *ControlFlow) {
W3282 := __e.Get(1)
_ = W3282
tmp13130 := MakeNative(func(__e *ControlFlow) {
W3283 := __e.Get(1)
_ = W3283
tmp13131 := Call(__e, PrimFunc(symappend), W3282, W3283)


__e.TailApply(PrimFunc(symappend), W3281, tmp13131)
return


}, 1)

tmp13132 := Call(__e, PrimFunc(symshen_4compile_1premises), V3278, V3279)


__e.TailApply(tmp13130, tmp13132)
return


}, 1)

tmp13133 := Call(__e, PrimFunc(symshen_4compile_1side_1conditions), V3277)


__e.TailApply(tmp13129, tmp13133)
return


}, 1)

tmp13134 := Call(__e, PrimFunc(symshen_4compile_1assumptions), V3276, V3275, V3279, V3280)


__e.TailApply(tmp13128, tmp13134)
return


}, 6)

tmp13135 := Call(__e, ns2_1set, symshen_4goals, tmp13127)


_ = tmp13135

tmp13136 := MakeNative(func(__e *ControlFlow) {
V3298 := __e.Get(1)
_ = V3298
V3299 := __e.Get(2)
_ = V3299
V3300 := __e.Get(3)
_ = V3300
V3301 := __e.Get(4)
_ = V3301
tmp13159 := PrimEqual(Nil, V3298)

if True == tmp13159 {
__e.Return(Nil)
return
} else {
tmp13157 := PrimIsPair(V3298)

var ifres13150 Obj

if True == tmp13157 {
tmp13156 := PrimIsPair(V3300)

var ifres13152 Obj

if True == tmp13156 {
tmp13154 := PrimTail(V3300)

tmp13155 := PrimIsPair(tmp13154)

var ifres13153 Obj

if True == tmp13155 {
ifres13153 = True


} else {
ifres13153 = False


}

ifres13152 = ifres13153


} else {
ifres13152 = False


}

var ifres13151 Obj

if True == ifres13152 {
ifres13151 = True


} else {
ifres13151 = False


}

ifres13150 = ifres13151


} else {
ifres13150 = False


}

if True == ifres13150 {
tmp13137 := MakeNative(func(__e *ControlFlow) {
W3302 := __e.Get(1)
_ = W3302
tmp13138 := PrimHead(V3298)

tmp13139 := PrimHead(V3300)

tmp13140 := PrimTail(V3300)

tmp13141 := PrimHead(tmp13140)

tmp13142 := Call(__e, PrimFunc(symshen_4compile_1assumption), tmp13138, tmp13139, tmp13141, V3299, V3301)


tmp13143 := PrimTail(V3298)

tmp13144 := PrimTail(V3300)

tmp13145 := Call(__e, PrimFunc(symshen_4compile_1assumptions), tmp13143, V3299, tmp13144, W3302)


__e.Return(PrimCons(tmp13142, tmp13145))
return


}, 1)

tmp13146 := PrimHead(V3298)

tmp13147 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp13146)


tmp13148 := Call(__e, PrimFunc(symappend), tmp13147, V3301)


__e.TailApply(tmp13137, tmp13148)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-assumptions")))
return
}


}


}, 4)

tmp13160 := Call(__e, ns2_1set, symshen_4compile_1assumptions, tmp13136)


_ = tmp13160

tmp13161 := MakeNative(func(__e *ControlFlow) {
V3303 := __e.Get(1)
_ = V3303
V3304 := __e.Get(2)
_ = V3304
V3305 := __e.Get(3)
_ = V3305
V3306 := __e.Get(4)
_ = V3306
V3307 := __e.Get(5)
_ = V3307
tmp13162 := MakeNative(func(__e *ControlFlow) {
W3308 := __e.Get(1)
_ = W3308
tmp13163 := MakeNative(func(__e *ControlFlow) {
W3309 := __e.Get(1)
_ = W3309
tmp13164 := PrimCons(V3305, V3306)

tmp13165 := PrimCons(Nil, tmp13164)

tmp13166 := PrimCons(V3304, tmp13165)

__e.Return(PrimCons(W3308, tmp13166))
return


}, 1)

tmp13167 := Call(__e, PrimFunc(symshen_4compile_1search_1procedure), W3308, V3303, V3304, V3305, V3306, V3307)


__e.TailApply(tmp13163, tmp13167)
return


}, 1)

tmp13168 := Call(__e, PrimFunc(symgensym), symshen_4search)


__e.TailApply(tmp13162, tmp13168)
return


}, 5)

tmp13169 := Call(__e, ns2_1set, symshen_4compile_1assumption, tmp13161)


_ = tmp13169

tmp13170 := MakeNative(func(__e *ControlFlow) {
V3310 := __e.Get(1)
_ = V3310
V3311 := __e.Get(2)
_ = V3311
V3312 := __e.Get(3)
_ = V3312
V3313 := __e.Get(4)
_ = V3313
V3314 := __e.Get(5)
_ = V3314
V3315 := __e.Get(6)
_ = V3315
tmp13171 := MakeNative(func(__e *ControlFlow) {
W3316 := __e.Get(1)
_ = W3316
tmp13172 := MakeNative(func(__e *ControlFlow) {
W3317 := __e.Get(1)
_ = W3317
tmp13173 := MakeNative(func(__e *ControlFlow) {
W3318 := __e.Get(1)
_ = W3318
tmp13174 := Call(__e, PrimFunc(symappend), W3317, W3318)


tmp13175 := PrimCons(V3310, tmp13174)

tmp13176 := PrimCons(symdefprolog, tmp13175)

__e.TailApply(PrimFunc(symeval), tmp13176)
return


}, 1)

tmp13177 := Call(__e, PrimFunc(symshen_4keep_1looking), V3310, V3312, W3316, V3313, V3314)


__e.TailApply(tmp13173, tmp13177)
return


}, 1)

tmp13178 := Call(__e, PrimFunc(symshen_4foundit_b), V3311, V3312, W3316, V3313, V3314, V3315)


__e.TailApply(tmp13172, tmp13178)
return


}, 1)

tmp13179 := Call(__e, PrimFunc(symgensym), symPrevious)


__e.TailApply(tmp13171, tmp13179)
return


}, 6)

tmp13180 := Call(__e, ns2_1set, symshen_4compile_1search_1procedure, tmp13170)


_ = tmp13180

tmp13181 := MakeNative(func(__e *ControlFlow) {
V3319 := __e.Get(1)
_ = V3319
V3320 := __e.Get(2)
_ = V3320
V3321 := __e.Get(3)
_ = V3321
V3322 := __e.Get(4)
_ = V3322
V3323 := __e.Get(5)
_ = V3323
V3324 := __e.Get(6)
_ = V3324
tmp13182 := MakeNative(func(__e *ControlFlow) {
W3325 := __e.Get(1)
_ = W3325
tmp13183 := MakeNative(func(__e *ControlFlow) {
W3326 := __e.Get(1)
_ = W3326
tmp13184 := MakeNative(func(__e *ControlFlow) {
W3327 := __e.Get(1)
_ = W3327
tmp13185 := MakeNative(func(__e *ControlFlow) {
W3328 := __e.Get(1)
_ = W3328
tmp13186 := PrimCons(sym_5_1_1, Nil)

tmp13187 := PrimIntern(MakeString(";"))

tmp13188 := PrimCons(tmp13187, Nil)

tmp13189 := Call(__e, PrimFunc(symappend), W3328, tmp13188)


tmp13190 := Call(__e, PrimFunc(symappend), tmp13186, tmp13189)


__e.TailApply(PrimFunc(symappend), W3327, tmp13190)
return


}, 1)

tmp13191 := Call(__e, PrimFunc(symshen_4body_1foundit_b), V3320, V3321, V3322, W3326)


__e.TailApply(tmp13185, tmp13191)
return


}, 1)

tmp13192 := Call(__e, PrimFunc(symshen_4head_1foundit_b), V3319, V3320, V3321, V3322, V3323, W3326)


__e.TailApply(tmp13184, tmp13192)
return


}, 1)

tmp13193 := Call(__e, PrimFunc(symshen_4tabulate_1passive), W3325)


__e.TailApply(tmp13183, tmp13193)
return


}, 1)

tmp13194 := Call(__e, PrimFunc(symshen_4passive), V3319, V3324)


__e.TailApply(tmp13182, tmp13194)
return


}, 6)

tmp13195 := Call(__e, ns2_1set, symshen_4foundit_b, tmp13181)


_ = tmp13195

tmp13196 := MakeNative(func(__e *ControlFlow) {
V3329 := __e.Get(1)
_ = V3329
V3330 := __e.Get(2)
_ = V3330
V3331 := __e.Get(3)
_ = V3331
V3332 := __e.Get(4)
_ = V3332
V3333 := __e.Get(5)
_ = V3333
tmp13197 := MakeNative(func(__e *ControlFlow) {
W3334 := __e.Get(1)
_ = W3334
tmp13198 := MakeNative(func(__e *ControlFlow) {
W3335 := __e.Get(1)
_ = W3335
tmp13199 := MakeNative(func(__e *ControlFlow) {
W3336 := __e.Get(1)
_ = W3336
tmp13200 := PrimCons(sym_5_1_1, Nil)

tmp13201 := PrimIntern(MakeString(";"))

tmp13202 := PrimCons(tmp13201, Nil)

tmp13203 := Call(__e, PrimFunc(symappend), W3336, tmp13202)


tmp13204 := Call(__e, PrimFunc(symappend), tmp13200, tmp13203)


__e.TailApply(PrimFunc(symappend), W3335, tmp13204)
return


}, 1)

tmp13205 := PrimCons(V3331, Nil)

tmp13206 := PrimCons(W3334, tmp13205)

tmp13207 := PrimCons(symcons, tmp13206)

tmp13208 := PrimCons(V3332, V3333)

tmp13209 := PrimCons(tmp13207, tmp13208)

tmp13210 := PrimCons(V3330, tmp13209)

tmp13211 := PrimCons(V3329, tmp13210)

tmp13212 := PrimCons(tmp13211, Nil)

__e.TailApply(tmp13199, tmp13212)
return


}, 1)

tmp13213 := PrimCons(V3330, Nil)

tmp13214 := PrimCons(W3334, tmp13213)

tmp13215 := PrimCons(symcons, tmp13214)

tmp13216 := PrimCons(tmp13215, Nil)

tmp13217 := PrimCons(sym_1, tmp13216)

tmp13218 := PrimCons(V3332, V3333)

tmp13219 := PrimCons(V3331, tmp13218)

tmp13220 := PrimCons(tmp13217, tmp13219)

__e.TailApply(tmp13198, tmp13220)
return


}, 1)

tmp13221 := Call(__e, PrimFunc(symgensym), symV)


__e.TailApply(tmp13197, tmp13221)
return


}, 5)

tmp13222 := Call(__e, ns2_1set, symshen_4keep_1looking, tmp13196)


_ = tmp13222

tmp13223 := MakeNative(func(__e *ControlFlow) {
V3341 := __e.Get(1)
_ = V3341
V3342 := __e.Get(2)
_ = V3342
tmp13231 := PrimIsPair(V3341)

if True == tmp13231 {
tmp13224 := PrimHead(V3341)

tmp13225 := Call(__e, PrimFunc(symshen_4passive), tmp13224, V3342)


tmp13226 := PrimTail(V3341)

tmp13227 := Call(__e, PrimFunc(symshen_4passive), tmp13226, V3342)


__e.TailApply(PrimFunc(symunion), tmp13225, tmp13227)
return


} else {
tmp13229 := Call(__e, PrimFunc(symshen_4passive_2), V3341, V3342)


if True == tmp13229 {
__e.Return(PrimCons(V3341, Nil))
return
} else {
__e.Return(Nil)
return
}


}


}, 2)

tmp13232 := Call(__e, ns2_1set, symshen_4passive, tmp13223)


_ = tmp13232

tmp13233 := MakeNative(func(__e *ControlFlow) {
V3343 := __e.Get(1)
_ = V3343
V3344 := __e.Get(2)
_ = V3344
tmp13237 := Call(__e, PrimFunc(symelement_2), V3343, V3344)


tmp13238 := PrimNot(tmp13237)

if True == tmp13238 {
tmp13235 := PrimIsVariable(V3343)

if True == tmp13235 {
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


}, 2)

tmp13239 := Call(__e, ns2_1set, symshen_4passive_2, tmp13233)


_ = tmp13239

tmp13240 := MakeNative(func(__e *ControlFlow) {
V3345 := __e.Get(1)
_ = V3345
tmp13241 := MakeNative(func(__e *ControlFlow) {
Z3346 := __e.Get(1)
_ = Z3346
tmp13242 := Call(__e, PrimFunc(symgensym), symV)


__e.Return(PrimCons(Z3346, tmp13242))
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp13241, V3345)
return


}, 1)

tmp13243 := Call(__e, ns2_1set, symshen_4tabulate_1passive, tmp13240)


_ = tmp13243

tmp13244 := MakeNative(func(__e *ControlFlow) {
V3347 := __e.Get(1)
_ = V3347
V3348 := __e.Get(2)
_ = V3348
V3349 := __e.Get(3)
_ = V3349
V3350 := __e.Get(4)
_ = V3350
V3351 := __e.Get(5)
_ = V3351
V3352 := __e.Get(6)
_ = V3352
tmp13245 := MakeNative(func(__e *ControlFlow) {
W3353 := __e.Get(1)
_ = W3353
tmp13246 := Call(__e, PrimFunc(symshen_4optimise_1typing), V3347)


tmp13247 := PrimCons(V3348, Nil)

tmp13248 := PrimCons(tmp13246, tmp13247)

tmp13249 := PrimCons(symcons, tmp13248)

tmp13250 := PrimCons(tmp13249, Nil)

tmp13251 := PrimCons(sym_1, tmp13250)

tmp13252 := PrimCons(V3350, W3353)

tmp13253 := PrimCons(V3349, tmp13252)

__e.Return(PrimCons(tmp13251, tmp13253))
return


}, 1)

tmp13254 := Call(__e, PrimFunc(symshen_4optimise_1passive), V3351, V3352)


__e.TailApply(tmp13245, tmp13254)
return


}, 6)

tmp13255 := Call(__e, ns2_1set, symshen_4head_1foundit_b, tmp13244)


_ = tmp13255

tmp13256 := MakeNative(func(__e *ControlFlow) {
V3354 := __e.Get(1)
_ = V3354
V3355 := __e.Get(2)
_ = V3355
tmp13257 := MakeNative(func(__e *ControlFlow) {
Z3356 := __e.Get(1)
_ = Z3356
__e.TailApply(PrimFunc(symshen_4optimise_1passive_1h), Z3356, V3355)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13257, V3354)
return


}, 2)

tmp13258 := Call(__e, ns2_1set, symshen_4optimise_1passive, tmp13256)


_ = tmp13258

tmp13259 := MakeNative(func(__e *ControlFlow) {
V3357 := __e.Get(1)
_ = V3357
V3358 := __e.Get(2)
_ = V3358
tmp13260 := MakeNative(func(__e *ControlFlow) {
W3359 := __e.Get(1)
_ = W3359
tmp13262 := Call(__e, PrimFunc(symempty_2), W3359)


if True == tmp13262 {
__e.Return(V3357)
return
} else {
__e.Return(PrimTail(W3359))
return
}


}, 1)

tmp13263 := Call(__e, PrimFunc(symassoc), V3357, V3358)


__e.TailApply(tmp13260, tmp13263)
return


}, 2)

tmp13264 := Call(__e, ns2_1set, symshen_4optimise_1passive_1h, tmp13259)


_ = tmp13264

tmp13265 := MakeNative(func(__e *ControlFlow) {
V3368 := __e.Get(1)
_ = V3368
V3369 := __e.Get(2)
_ = V3369
V3370 := __e.Get(3)
_ = V3370
V3371 := __e.Get(4)
_ = V3371
tmp13292 := PrimEqual(Nil, V3371)

if True == tmp13292 {
tmp13266 := PrimCons(V3369, Nil)

tmp13267 := PrimCons(MakeNumber(1), tmp13266)

tmp13268 := PrimCons(V3368, Nil)

tmp13269 := PrimCons(MakeNumber(1), tmp13268)

tmp13270 := PrimCons(tmp13269, Nil)

tmp13271 := PrimCons(tmp13267, tmp13270)

tmp13272 := PrimCons(symappend, tmp13271)

tmp13273 := PrimCons(tmp13272, Nil)

tmp13274 := PrimCons(V3370, tmp13273)

tmp13275 := PrimCons(symbind, tmp13274)

__e.Return(PrimCons(tmp13275, Nil))
return


} else {
tmp13290 := PrimIsPair(V3371)

var ifres13286 Obj

if True == tmp13290 {
tmp13288 := PrimHead(V3371)

tmp13289 := PrimIsPair(tmp13288)

var ifres13287 Obj

if True == tmp13289 {
ifres13287 = True


} else {
ifres13287 = False


}

ifres13286 = ifres13287


} else {
ifres13286 = False


}

if True == ifres13286 {
tmp13276 := PrimHead(V3371)

tmp13277 := PrimTail(tmp13276)

tmp13278 := PrimHead(V3371)

tmp13279 := PrimHead(tmp13278)

tmp13280 := PrimCons(tmp13279, Nil)

tmp13281 := PrimCons(tmp13277, tmp13280)

tmp13282 := PrimCons(symbind, tmp13281)

tmp13283 := PrimTail(V3371)

tmp13284 := Call(__e, PrimFunc(symshen_4body_1foundit_b), V3368, V3369, V3370, tmp13283)


__e.Return(PrimCons(tmp13282, tmp13284))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.body-foundit!")))
return
}


}


}, 4)

tmp13293 := Call(__e, ns2_1set, symshen_4body_1foundit_b, tmp13265)


_ = tmp13293

tmp13294 := MakeNative(func(__e *ControlFlow) {
V3372 := __e.Get(1)
_ = V3372
tmp13295 := MakeNative(func(__e *ControlFlow) {
Z3373 := __e.Get(1)
_ = Z3373
__e.TailApply(PrimFunc(symshen_4compile_1side_1condition), Z3373)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13295, V3372)
return


}, 1)

tmp13296 := Call(__e, ns2_1set, symshen_4compile_1side_1conditions, tmp13294)


_ = tmp13296

tmp13297 := MakeNative(func(__e *ControlFlow) {
V3376 := __e.Get(1)
_ = V3376
tmp13357 := PrimIsPair(V3376)

var ifres13338 Obj

if True == tmp13357 {
tmp13355 := PrimHead(V3376)

tmp13356 := PrimEqual(symlet, tmp13355)

var ifres13340 Obj

if True == tmp13356 {
tmp13353 := PrimTail(V3376)

tmp13354 := PrimIsPair(tmp13353)

var ifres13342 Obj

if True == tmp13354 {
tmp13350 := PrimTail(V3376)

tmp13351 := PrimTail(tmp13350)

tmp13352 := PrimIsPair(tmp13351)

var ifres13344 Obj

if True == tmp13352 {
tmp13346 := PrimTail(V3376)

tmp13347 := PrimTail(tmp13346)

tmp13348 := PrimTail(tmp13347)

tmp13349 := PrimEqual(Nil, tmp13348)

var ifres13345 Obj

if True == tmp13349 {
ifres13345 = True


} else {
ifres13345 = False


}

ifres13344 = ifres13345


} else {
ifres13344 = False


}

var ifres13343 Obj

if True == ifres13344 {
ifres13343 = True


} else {
ifres13343 = False


}

ifres13342 = ifres13343


} else {
ifres13342 = False


}

var ifres13341 Obj

if True == ifres13342 {
ifres13341 = True


} else {
ifres13341 = False


}

ifres13340 = ifres13341


} else {
ifres13340 = False


}

var ifres13339 Obj

if True == ifres13340 {
ifres13339 = True


} else {
ifres13339 = False


}

ifres13338 = ifres13339


} else {
ifres13338 = False


}

if True == ifres13338 {
tmp13298 := PrimTail(V3376)

__e.Return(PrimCons(symis, tmp13298))
return


} else {
tmp13336 := PrimIsPair(V3376)

var ifres13317 Obj

if True == tmp13336 {
tmp13334 := PrimHead(V3376)

tmp13335 := PrimEqual(symshen_4let_b, tmp13334)

var ifres13319 Obj

if True == tmp13335 {
tmp13332 := PrimTail(V3376)

tmp13333 := PrimIsPair(tmp13332)

var ifres13321 Obj

if True == tmp13333 {
tmp13329 := PrimTail(V3376)

tmp13330 := PrimTail(tmp13329)

tmp13331 := PrimIsPair(tmp13330)

var ifres13323 Obj

if True == tmp13331 {
tmp13325 := PrimTail(V3376)

tmp13326 := PrimTail(tmp13325)

tmp13327 := PrimTail(tmp13326)

tmp13328 := PrimEqual(Nil, tmp13327)

var ifres13324 Obj

if True == tmp13328 {
ifres13324 = True


} else {
ifres13324 = False


}

ifres13323 = ifres13324


} else {
ifres13323 = False


}

var ifres13322 Obj

if True == ifres13323 {
ifres13322 = True


} else {
ifres13322 = False


}

ifres13321 = ifres13322


} else {
ifres13321 = False


}

var ifres13320 Obj

if True == ifres13321 {
ifres13320 = True


} else {
ifres13320 = False


}

ifres13319 = ifres13320


} else {
ifres13319 = False


}

var ifres13318 Obj

if True == ifres13319 {
ifres13318 = True


} else {
ifres13318 = False


}

ifres13317 = ifres13318


} else {
ifres13317 = False


}

if True == ifres13317 {
tmp13299 := PrimTail(V3376)

__e.Return(PrimCons(symis_b, tmp13299))
return


} else {
tmp13315 := PrimIsPair(V3376)

var ifres13302 Obj

if True == tmp13315 {
tmp13313 := PrimHead(V3376)

tmp13314 := PrimEqual(symif, tmp13313)

var ifres13304 Obj

if True == tmp13314 {
tmp13311 := PrimTail(V3376)

tmp13312 := PrimIsPair(tmp13311)

var ifres13306 Obj

if True == tmp13312 {
tmp13308 := PrimTail(V3376)

tmp13309 := PrimTail(tmp13308)

tmp13310 := PrimEqual(Nil, tmp13309)

var ifres13307 Obj

if True == tmp13310 {
ifres13307 = True


} else {
ifres13307 = False


}

ifres13306 = ifres13307


} else {
ifres13306 = False


}

var ifres13305 Obj

if True == ifres13306 {
ifres13305 = True


} else {
ifres13305 = False


}

ifres13304 = ifres13305


} else {
ifres13304 = False


}

var ifres13303 Obj

if True == ifres13304 {
ifres13303 = True


} else {
ifres13303 = False


}

ifres13302 = ifres13303


} else {
ifres13302 = False


}

if True == ifres13302 {
tmp13300 := PrimTail(V3376)

__e.Return(PrimCons(symwhen, tmp13300))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-side-condition")))
return
}


}


}


}, 1)

tmp13358 := Call(__e, ns2_1set, symshen_4compile_1side_1condition, tmp13297)


_ = tmp13358

tmp13359 := MakeNative(func(__e *ControlFlow) {
V3377 := __e.Get(1)
_ = V3377
V3378 := __e.Get(2)
_ = V3378
tmp13360 := MakeNative(func(__e *ControlFlow) {
W3379 := __e.Get(1)
_ = W3379
tmp13361 := MakeNative(func(__e *ControlFlow) {
Z3380 := __e.Get(1)
_ = Z3380
__e.TailApply(PrimFunc(symshen_4compile_1premise), Z3380, W3379)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13361, V3377)
return


}, 1)

tmp13362 := Call(__e, PrimFunc(symreverse), V3378)


tmp13363 := PrimHead(tmp13362)

__e.TailApply(tmp13360, tmp13363)
return


}, 2)

tmp13364 := Call(__e, ns2_1set, symshen_4compile_1premises, tmp13359)


_ = tmp13364

tmp13365 := MakeNative(func(__e *ControlFlow) {
V3387 := __e.Get(1)
_ = V3387
V3388 := __e.Get(2)
_ = V3388
tmp13382 := PrimEqual(sym_b, V3387)

if True == tmp13382 {
__e.Return(sym_b)
return
} else {
tmp13380 := PrimIsPair(V3387)

var ifres13371 Obj

if True == tmp13380 {
tmp13378 := PrimTail(V3387)

tmp13379 := PrimIsPair(tmp13378)

var ifres13373 Obj

if True == tmp13379 {
tmp13375 := PrimTail(V3387)

tmp13376 := PrimTail(tmp13375)

tmp13377 := PrimEqual(Nil, tmp13376)

var ifres13374 Obj

if True == tmp13377 {
ifres13374 = True


} else {
ifres13374 = False


}

ifres13373 = ifres13374


} else {
ifres13373 = False


}

var ifres13372 Obj

if True == ifres13373 {
ifres13372 = True


} else {
ifres13372 = False


}

ifres13371 = ifres13372


} else {
ifres13371 = False


}

if True == ifres13371 {
tmp13366 := PrimHead(V3387)

tmp13367 := Call(__e, PrimFunc(symreverse), tmp13366)


tmp13368 := PrimTail(V3387)

tmp13369 := PrimHead(tmp13368)

__e.TailApply(PrimFunc(symshen_4compile_1premise_1h), tmp13367, tmp13369, V3388)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.premise")))
return
}


}


}, 2)

tmp13383 := Call(__e, ns2_1set, symshen_4compile_1premise, tmp13365)


_ = tmp13383

tmp13384 := MakeNative(func(__e *ControlFlow) {
V3395 := __e.Get(1)
_ = V3395
V3396 := __e.Get(2)
_ = V3396
V3397 := __e.Get(3)
_ = V3397
tmp13397 := PrimEqual(Nil, V3395)

if True == tmp13397 {
tmp13385 := Call(__e, PrimFunc(symshen_4cons_1form_1no_1modes), V3396)


tmp13386 := PrimCons(V3397, Nil)

tmp13387 := PrimCons(tmp13385, tmp13386)

__e.Return(PrimCons(symshen_4system_1S, tmp13387))
return


} else {
tmp13395 := PrimIsPair(V3395)

if True == tmp13395 {
tmp13388 := PrimTail(V3395)

tmp13389 := PrimHead(V3395)

tmp13390 := Call(__e, PrimFunc(symshen_4cons_1form_1no_1modes), tmp13389)


tmp13391 := PrimCons(V3397, Nil)

tmp13392 := PrimCons(tmp13390, tmp13391)

tmp13393 := PrimCons(symcons, tmp13392)

__e.TailApply(PrimFunc(symshen_4compile_1premise_1h), tmp13388, V3396, tmp13393)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-premise-h")))
return
}


}


}, 3)

tmp13398 := Call(__e, ns2_1set, symshen_4compile_1premise_1h, tmp13384)


_ = tmp13398

tmp13399 := MakeNative(func(__e *ControlFlow) {
V3398 := __e.Get(1)
_ = V3398
tmp13423 := PrimIsPair(V3398)

var ifres13410 Obj

if True == tmp13423 {
tmp13421 := PrimHead(V3398)

tmp13422 := PrimEqual(symbar_b, tmp13421)

var ifres13412 Obj

if True == tmp13422 {
tmp13419 := PrimTail(V3398)

tmp13420 := PrimIsPair(tmp13419)

var ifres13414 Obj

if True == tmp13420 {
tmp13416 := PrimTail(V3398)

tmp13417 := PrimTail(tmp13416)

tmp13418 := PrimEqual(Nil, tmp13417)

var ifres13415 Obj

if True == tmp13418 {
ifres13415 = True


} else {
ifres13415 = False


}

ifres13414 = ifres13415


} else {
ifres13414 = False


}

var ifres13413 Obj

if True == ifres13414 {
ifres13413 = True


} else {
ifres13413 = False


}

ifres13412 = ifres13413


} else {
ifres13412 = False


}

var ifres13411 Obj

if True == ifres13412 {
ifres13411 = True


} else {
ifres13411 = False


}

ifres13410 = ifres13411


} else {
ifres13410 = False


}

if True == ifres13410 {
tmp13400 := PrimTail(V3398)

__e.Return(PrimHead(tmp13400))
return


} else {
tmp13408 := PrimIsPair(V3398)

if True == tmp13408 {
tmp13401 := PrimHead(V3398)

tmp13402 := Call(__e, PrimFunc(symshen_4cons_1form_1no_1modes), tmp13401)


tmp13403 := PrimTail(V3398)

tmp13404 := Call(__e, PrimFunc(symshen_4cons_1form_1no_1modes), tmp13403)


tmp13405 := PrimCons(tmp13404, Nil)

tmp13406 := PrimCons(tmp13402, tmp13405)

__e.Return(PrimCons(symcons, tmp13406))
return


} else {
__e.Return(V3398)
return
}


}


}, 1)

tmp13424 := Call(__e, ns2_1set, symshen_4cons_1form_1no_1modes, tmp13399)


_ = tmp13424

tmp13425 := MakeNative(func(__e *ControlFlow) {
V3399 := __e.Get(1)
_ = V3399
tmp13426 := MakeNative(func(__e *ControlFlow) {
W3400 := __e.Get(1)
_ = W3400
tmp13427 := MakeNative(func(__e *ControlFlow) {
W3402 := __e.Get(1)
_ = W3402
tmp13428 := MakeNative(func(__e *ControlFlow) {
W3403 := __e.Get(1)
_ = W3403
tmp13429 := MakeNative(func(__e *ControlFlow) {
W3404 := __e.Get(1)
_ = W3404
__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3404)
return
}, 1)

tmp13430 := PrimSet(symshen_4_ddatatypes_d, W3403)

__e.TailApply(tmp13429, tmp13430)
return


}, 1)

tmp13431 := Call(__e, PrimFunc(symshen_4remove_1datatypes), W3400, W3402)


__e.TailApply(tmp13428, tmp13431)
return


}, 1)

tmp13432 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(tmp13427, tmp13432)
return


}, 1)

tmp13433 := MakeNative(func(__e *ControlFlow) {
Z3401 := __e.Get(1)
_ = Z3401
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3401)
return
}, 1)

tmp13434 := Call(__e, PrimFunc(symmap), tmp13433, V3399)


__e.TailApply(tmp13426, tmp13434)
return


}, 1)

tmp13435 := Call(__e, ns2_1set, sympreclude, tmp13425)


_ = tmp13435

tmp13436 := MakeNative(func(__e *ControlFlow) {
V3409 := __e.Get(1)
_ = V3409
V3410 := __e.Get(2)
_ = V3410
tmp13443 := PrimEqual(Nil, V3409)

if True == tmp13443 {
__e.Return(V3410)
return
} else {
tmp13441 := PrimIsPair(V3409)

if True == tmp13441 {
tmp13437 := PrimTail(V3409)

tmp13438 := PrimHead(V3409)

tmp13439 := Call(__e, PrimFunc(symshen_4unassoc), tmp13438, V3410)


__e.TailApply(PrimFunc(symshen_4remove_1datatypes), tmp13437, tmp13439)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-datatypes")))
return
}


}


}, 2)

tmp13444 := Call(__e, ns2_1set, symshen_4remove_1datatypes, tmp13436)


_ = tmp13444

tmp13445 := MakeNative(func(__e *ControlFlow) {
V3411 := __e.Get(1)
_ = V3411
tmp13446 := MakeNative(func(__e *ControlFlow) {
Z3412 := __e.Get(1)
_ = Z3412
__e.Return(PrimHead(Z3412))
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13446, V3411)
return


}, 1)

tmp13447 := Call(__e, ns2_1set, symshen_4show_1datatypes, tmp13445)


_ = tmp13447

tmp13448 := MakeNative(func(__e *ControlFlow) {
V3413 := __e.Get(1)
_ = V3413
tmp13449 := MakeNative(func(__e *ControlFlow) {
W3414 := __e.Get(1)
_ = W3414
tmp13450 := MakeNative(func(__e *ControlFlow) {
W3416 := __e.Get(1)
_ = W3416
tmp13451 := MakeNative(func(__e *ControlFlow) {
W3418 := __e.Get(1)
_ = W3418
__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3418)
return
}, 1)

tmp13452 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(tmp13451, tmp13452)
return


}, 1)

tmp13453 := MakeNative(func(__e *ControlFlow) {
Z3417 := __e.Get(1)
_ = Z3417
tmp13454 := Call(__e, PrimFunc(symfn), Z3417)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3417, tmp13454)
return


}, 1)

tmp13455 := Call(__e, PrimFunc(symmap), tmp13453, W3414)


__e.TailApply(tmp13450, tmp13455)
return


}, 1)

tmp13456 := MakeNative(func(__e *ControlFlow) {
Z3415 := __e.Get(1)
_ = Z3415
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3415)
return
}, 1)

tmp13457 := Call(__e, PrimFunc(symmap), tmp13456, V3413)


__e.TailApply(tmp13449, tmp13457)
return


}, 1)

tmp13458 := Call(__e, ns2_1set, syminclude, tmp13448)


_ = tmp13458

tmp13459 := MakeNative(func(__e *ControlFlow) {
V3419 := __e.Get(1)
_ = V3419
tmp13460 := MakeNative(func(__e *ControlFlow) {
W3420 := __e.Get(1)
_ = W3420
tmp13461 := MakeNative(func(__e *ControlFlow) {
W3421 := __e.Get(1)
_ = W3421
tmp13462 := MakeNative(func(__e *ControlFlow) {
W3423 := __e.Get(1)
_ = W3423
tmp13463 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(PrimFunc(symshen_4show_1datatypes), tmp13463)
return


}, 1)

tmp13464 := MakeNative(func(__e *ControlFlow) {
Z3424 := __e.Get(1)
_ = Z3424
tmp13465 := Call(__e, PrimFunc(symfn), Z3424)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3424, tmp13465)
return


}, 1)

tmp13466 := Call(__e, PrimFunc(symmap), tmp13464, W3421)


__e.TailApply(tmp13462, tmp13466)
return


}, 1)

tmp13467 := MakeNative(func(__e *ControlFlow) {
Z3422 := __e.Get(1)
_ = Z3422
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3422)
return
}, 1)

tmp13468 := Call(__e, PrimFunc(symmap), tmp13467, V3419)


__e.TailApply(tmp13461, tmp13468)
return


}, 1)

tmp13469 := PrimSet(symshen_4_ddatatypes_d, Nil)

__e.TailApply(tmp13460, tmp13469)
return


}, 1)

tmp13470 := Call(__e, ns2_1set, sympreclude_1all_1but, tmp13459)


_ = tmp13470

tmp13471 := MakeNative(func(__e *ControlFlow) {
V3425 := __e.Get(1)
_ = V3425
tmp13472 := MakeNative(func(__e *ControlFlow) {
W3426 := __e.Get(1)
_ = W3426
tmp13473 := MakeNative(func(__e *ControlFlow) {
W3428 := __e.Get(1)
_ = W3428
tmp13474 := MakeNative(func(__e *ControlFlow) {
W3429 := __e.Get(1)
_ = W3429
__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3429)
return
}, 1)

tmp13475 := Call(__e, PrimFunc(symshen_4remove_1datatypes), W3426, W3428)


tmp13476 := PrimSet(symshen_4_ddatatypes_d, tmp13475)

__e.TailApply(tmp13474, tmp13476)
return


}, 1)

tmp13477 := PrimValue(symshen_4_dalldatatypes_d)

__e.TailApply(tmp13473, tmp13477)
return


}, 1)

tmp13478 := MakeNative(func(__e *ControlFlow) {
Z3427 := __e.Get(1)
_ = Z3427
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3427)
return
}, 1)

tmp13479 := Call(__e, PrimFunc(symmap), tmp13478, V3425)


__e.TailApply(tmp13472, tmp13479)
return


}, 1)

__e.TailApply(ns2_1set, syminclude_1all_1but, tmp13471)
return




}, 0)

