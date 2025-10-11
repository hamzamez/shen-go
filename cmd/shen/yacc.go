package main

import . "github.com/tiancaiamao/shen-go/kl"

var YaccMain = MakeNative(func(__e *ControlFlow) {
tmp17223 := MakeNative(func(__e *ControlFlow) {
V107 := __e.Get(1)
_ = V107
V108 := __e.Get(2)
_ = V108
tmp17224 := MakeNative(func(__e *ControlFlow) {
W109 := __e.Get(1)
_ = W109
tmp17233 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W109)


if True == tmp17233 {
__e.Return(PrimSimpleError(MakeString("parse failure\n")))
return
} else {
tmp17230 := Call(__e, PrimFunc(symshen_4in_1_6), W109)


tmp17231 := PrimIsPair(tmp17230)

if True == tmp17231 {
tmp17225 := Call(__e, PrimFunc(symshen_4in_1_6), W109)


tmp17226 := PrimHead(tmp17225)

tmp17227 := Call(__e, PrimFunc(symshen_4app), tmp17226, MakeString(" ..."), symshen_4s)


tmp17228 := PrimStringConcat(MakeString("syntax error here: "), tmp17227)

__e.Return(PrimSimpleError(tmp17228))
return


} else {
__e.TailApply(PrimFunc(symshen_4_5_1out), W109)
return
}


}


}, 1)

tmp17234 := Call(__e, V107, V108)


__e.TailApply(tmp17224, tmp17234)
return


}, 2)

tmp17235 := Call(__e, ns2_1set, symcompile, tmp17223)


_ = tmp17235

tmp17236 := MakeNative(func(__e *ControlFlow) {
V110 := __e.Get(1)
_ = V110
tmp17237 := Call(__e, PrimFunc(symfail))


__e.Return(PrimEqual(V110, tmp17237))
return


}, 1)

tmp17238 := Call(__e, ns2_1set, symshen_4parse_1failure_2, tmp17236)


_ = tmp17238

tmp17239 := MakeNative(func(__e *ControlFlow) {
V113 := __e.Get(1)
_ = V113
tmp17252 := PrimIsPair(V113)

var ifres17243 Obj

if True == tmp17252 {
tmp17250 := PrimTail(V113)

tmp17251 := PrimIsPair(tmp17250)

var ifres17245 Obj

if True == tmp17251 {
tmp17247 := PrimTail(V113)

tmp17248 := PrimTail(tmp17247)

tmp17249 := PrimEqual(Nil, tmp17248)

var ifres17246 Obj

if True == tmp17249 {
ifres17246 = True


} else {
ifres17246 = False


}

ifres17245 = ifres17246


} else {
ifres17245 = False


}

var ifres17244 Obj

if True == ifres17245 {
ifres17244 = True


} else {
ifres17244 = False


}

ifres17243 = ifres17244


} else {
ifres17243 = False


}

if True == ifres17243 {
tmp17240 := PrimTail(V113)

__e.Return(PrimHead(tmp17240))
return


} else {
tmp17241 := Call(__e, PrimFunc(symshen_4app), V113, MakeString(" is not a YACC stream\n"), symshen_4s)


__e.Return(PrimSimpleError(tmp17241))
return


}


}, 1)

tmp17253 := Call(__e, ns2_1set, symshen_4objectcode, tmp17239)


_ = tmp17253

tmp17254 := MakeNative(func(__e *ControlFlow) {
V114 := __e.Get(1)
_ = V114
tmp17255 := MakeNative(func(__e *ControlFlow) {
Z115 := __e.Get(1)
_ = Z115
__e.TailApply(PrimFunc(symshen_4_5yacc_6), Z115)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp17255, V114)
return


}, 1)

tmp17256 := Call(__e, ns2_1set, symshen_4yacc_1_6shen, tmp17254)


_ = tmp17256

tmp17257 := MakeNative(func(__e *ControlFlow) {
V116 := __e.Get(1)
_ = V116
tmp17258 := MakeNative(func(__e *ControlFlow) {
W117 := __e.Get(1)
_ = W117
tmp17260 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W117)


if True == tmp17260 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W117)
return
}


}, 1)

tmp17296 := PrimIsPair(V116)

var ifres17261 Obj

if True == tmp17296 {
tmp17262 := MakeNative(func(__e *ControlFlow) {
W118 := __e.Get(1)
_ = W118
tmp17263 := MakeNative(func(__e *ControlFlow) {
W119 := __e.Get(1)
_ = W119
tmp17264 := MakeNative(func(__e *ControlFlow) {
W120 := __e.Get(1)
_ = W120
tmp17290 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W120)


if True == tmp17290 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17265 := MakeNative(func(__e *ControlFlow) {
W121 := __e.Get(1)
_ = W121
tmp17266 := MakeNative(func(__e *ControlFlow) {
W122 := __e.Get(1)
_ = W122
tmp17267 := MakeNative(func(__e *ControlFlow) {
W123 := __e.Get(1)
_ = W123
tmp17285 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W123)


if True == tmp17285 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17268 := MakeNative(func(__e *ControlFlow) {
W124 := __e.Get(1)
_ = W124
tmp17269 := MakeNative(func(__e *ControlFlow) {
W125 := __e.Get(1)
_ = W125
tmp17270 := MakeNative(func(__e *ControlFlow) {
W126 := __e.Get(1)
_ = W126
tmp17271 := MakeNative(func(__e *ControlFlow) {
W127 := __e.Get(1)
_ = W127
__e.Return(W127)
return
}, 1)

tmp17272 := PrimCons(W118, Nil)

tmp17273 := PrimCons(symdefine, tmp17272)

tmp17274 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), W121, W126, W124)


tmp17275 := PrimCons(tmp17274, Nil)

tmp17276 := PrimCons(sym_1_6, tmp17275)

tmp17277 := PrimCons(W126, tmp17276)

tmp17278 := Call(__e, PrimFunc(symappend), W121, tmp17277)


tmp17279 := Call(__e, PrimFunc(symappend), tmp17273, tmp17278)


__e.TailApply(tmp17271, tmp17279)
return


}, 1)

tmp17280 := Call(__e, PrimFunc(symgensym), symS)


tmp17281 := Call(__e, tmp17270, tmp17280)


__e.TailApply(PrimFunc(symshen_4comb), W125, tmp17281)
return


}, 1)

tmp17282 := Call(__e, PrimFunc(symshen_4in_1_6), W123)


__e.TailApply(tmp17269, tmp17282)
return


}, 1)

tmp17283 := Call(__e, PrimFunc(symshen_4_5_1out), W123)


__e.TailApply(tmp17268, tmp17283)
return


}


}, 1)

tmp17286 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), W122)


__e.TailApply(tmp17267, tmp17286)
return


}, 1)

tmp17287 := Call(__e, PrimFunc(symshen_4in_1_6), W120)


__e.TailApply(tmp17266, tmp17287)
return


}, 1)

tmp17288 := Call(__e, PrimFunc(symshen_4_5_1out), W120)


__e.TailApply(tmp17265, tmp17288)
return


}


}, 1)

tmp17291 := Call(__e, PrimFunc(symshen_4_5yaccsig_6), W119)


__e.TailApply(tmp17264, tmp17291)
return


}, 1)

tmp17292 := Call(__e, PrimFunc(symtail), V116)


__e.TailApply(tmp17263, tmp17292)
return


}, 1)

tmp17293 := Call(__e, PrimFunc(symhead), V116)


tmp17294 := Call(__e, tmp17262, tmp17293)


ifres17261 = tmp17294


} else {
tmp17295 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17261 = tmp17295


}

__e.TailApply(tmp17258, ifres17261)
return


}, 1)

tmp17297 := Call(__e, ns2_1set, symshen_4_5yacc_6, tmp17257)


_ = tmp17297

tmp17298 := MakeNative(func(__e *ControlFlow) {
V128 := __e.Get(1)
_ = V128
tmp17299 := MakeNative(func(__e *ControlFlow) {
W129 := __e.Get(1)
_ = W129
tmp17311 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W129)


if True == tmp17311 {
tmp17300 := MakeNative(func(__e *ControlFlow) {
W144 := __e.Get(1)
_ = W144
tmp17302 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W144)


if True == tmp17302 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W144)
return
}


}, 1)

tmp17303 := MakeNative(func(__e *ControlFlow) {
W145 := __e.Get(1)
_ = W145
tmp17307 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W145)


if True == tmp17307 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17304 := MakeNative(func(__e *ControlFlow) {
W146 := __e.Get(1)
_ = W146
__e.TailApply(PrimFunc(symshen_4comb), W146, Nil)
return
}, 1)

tmp17305 := Call(__e, PrimFunc(symshen_4in_1_6), W145)


__e.TailApply(tmp17304, tmp17305)
return


}


}, 1)

tmp17308 := Call(__e, PrimFunc(sym_5e_6), V128)


tmp17309 := Call(__e, tmp17303, tmp17308)


__e.TailApply(tmp17300, tmp17309)
return


} else {
__e.Return(W129)
return
}


}, 1)

tmp17374 := PrimIsPair(V128)

var ifres17312 Obj

if True == tmp17374 {
tmp17313 := MakeNative(func(__e *ControlFlow) {
W130 := __e.Get(1)
_ = W130
tmp17314 := MakeNative(func(__e *ControlFlow) {
W131 := __e.Get(1)
_ = W131
tmp17369 := Call(__e, PrimFunc(symshen_4ccons_2), W131)


if True == tmp17369 {
tmp17315 := MakeNative(func(__e *ControlFlow) {
W132 := __e.Get(1)
_ = W132
tmp17316 := MakeNative(func(__e *ControlFlow) {
W133 := __e.Get(1)
_ = W133
tmp17365 := Call(__e, PrimFunc(symshen_4hds_a_2), W132, symlist)


if True == tmp17365 {
tmp17317 := MakeNative(func(__e *ControlFlow) {
W134 := __e.Get(1)
_ = W134
tmp17362 := PrimIsPair(W134)

if True == tmp17362 {
tmp17318 := MakeNative(func(__e *ControlFlow) {
W135 := __e.Get(1)
_ = W135
tmp17319 := MakeNative(func(__e *ControlFlow) {
W136 := __e.Get(1)
_ = W136
tmp17320 := MakeNative(func(__e *ControlFlow) {
W137 := __e.Get(1)
_ = W137
tmp17357 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W137)


if True == tmp17357 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17321 := MakeNative(func(__e *ControlFlow) {
W138 := __e.Get(1)
_ = W138
tmp17354 := Call(__e, PrimFunc(symshen_4hds_a_2), W133, sym_a_a_6)


if True == tmp17354 {
tmp17322 := MakeNative(func(__e *ControlFlow) {
W139 := __e.Get(1)
_ = W139
tmp17351 := PrimIsPair(W139)

if True == tmp17351 {
tmp17323 := MakeNative(func(__e *ControlFlow) {
W140 := __e.Get(1)
_ = W140
tmp17324 := MakeNative(func(__e *ControlFlow) {
W141 := __e.Get(1)
_ = W141
tmp17347 := PrimIsPair(W141)

if True == tmp17347 {
tmp17325 := MakeNative(func(__e *ControlFlow) {
W142 := __e.Get(1)
_ = W142
tmp17326 := MakeNative(func(__e *ControlFlow) {
W143 := __e.Get(1)
_ = W143
tmp17343 := PrimEqual(sym_i, W130)

var ifres17340 Obj

if True == tmp17343 {
tmp17342 := PrimEqual(sym_j, W142)

var ifres17341 Obj

if True == tmp17342 {
ifres17341 = True


} else {
ifres17341 = False


}

ifres17340 = ifres17341


} else {
ifres17340 = False


}

if True == ifres17340 {
tmp17327 := PrimCons(W135, Nil)

tmp17328 := PrimCons(symlist, tmp17327)

tmp17329 := PrimCons(W135, Nil)

tmp17330 := PrimCons(symlist, tmp17329)

tmp17331 := PrimCons(W140, Nil)

tmp17332 := PrimCons(tmp17330, tmp17331)

tmp17333 := PrimCons(symstr, tmp17332)

tmp17334 := PrimCons(sym_j, Nil)

tmp17335 := PrimCons(tmp17333, tmp17334)

tmp17336 := PrimCons(sym_1_1_6, tmp17335)

tmp17337 := PrimCons(tmp17328, tmp17336)

tmp17338 := PrimCons(sym_i, tmp17337)

__e.TailApply(PrimFunc(symshen_4comb), W143, tmp17338)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17344 := Call(__e, PrimFunc(symtail), W141)


__e.TailApply(tmp17326, tmp17344)
return


}, 1)

tmp17345 := Call(__e, PrimFunc(symhead), W141)


__e.TailApply(tmp17325, tmp17345)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17348 := Call(__e, PrimFunc(symtail), W139)


__e.TailApply(tmp17324, tmp17348)
return


}, 1)

tmp17349 := Call(__e, PrimFunc(symhead), W139)


__e.TailApply(tmp17323, tmp17349)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17352 := Call(__e, PrimFunc(symtail), W133)


__e.TailApply(tmp17322, tmp17352)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17355 := Call(__e, PrimFunc(symshen_4in_1_6), W137)


__e.TailApply(tmp17321, tmp17355)
return


}


}, 1)

tmp17358 := Call(__e, PrimFunc(sym_5end_6), W136)


__e.TailApply(tmp17320, tmp17358)
return


}, 1)

tmp17359 := Call(__e, PrimFunc(symtail), W134)


__e.TailApply(tmp17319, tmp17359)
return


}, 1)

tmp17360 := Call(__e, PrimFunc(symhead), W134)


__e.TailApply(tmp17318, tmp17360)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17363 := Call(__e, PrimFunc(symtail), W132)


__e.TailApply(tmp17317, tmp17363)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17366 := Call(__e, PrimFunc(symtail), W131)


__e.TailApply(tmp17316, tmp17366)
return


}, 1)

tmp17367 := Call(__e, PrimFunc(symhead), W131)


__e.TailApply(tmp17315, tmp17367)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17370 := Call(__e, PrimFunc(symtail), V128)


__e.TailApply(tmp17314, tmp17370)
return


}, 1)

tmp17371 := Call(__e, PrimFunc(symhead), V128)


tmp17372 := Call(__e, tmp17313, tmp17371)


ifres17312 = tmp17372


} else {
tmp17373 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17312 = tmp17373


}

__e.TailApply(tmp17299, ifres17312)
return


}, 1)

tmp17375 := Call(__e, ns2_1set, symshen_4_5yaccsig_6, tmp17298)


_ = tmp17375

tmp17376 := MakeNative(func(__e *ControlFlow) {
V147 := __e.Get(1)
_ = V147
tmp17377 := MakeNative(func(__e *ControlFlow) {
W148 := __e.Get(1)
_ = W148
tmp17396 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W148)


if True == tmp17396 {
tmp17378 := MakeNative(func(__e *ControlFlow) {
W155 := __e.Get(1)
_ = W155
tmp17380 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W155)


if True == tmp17380 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W155)
return
}


}, 1)

tmp17381 := MakeNative(func(__e *ControlFlow) {
W156 := __e.Get(1)
_ = W156
tmp17392 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W156)


if True == tmp17392 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17382 := MakeNative(func(__e *ControlFlow) {
W157 := __e.Get(1)
_ = W157
tmp17383 := MakeNative(func(__e *ControlFlow) {
W158 := __e.Get(1)
_ = W158
tmp17388 := Call(__e, PrimFunc(symempty_2), W157)


var ifres17384 Obj

if True == tmp17388 {
ifres17384 = Nil


} else {
tmp17385 := Call(__e, PrimFunc(symshen_4app), W157, MakeString("\n ..."), symshen_4r)


tmp17386 := PrimStringConcat(MakeString("YACC syntax error here:\n "), tmp17385)

tmp17387 := PrimSimpleError(tmp17386)

ifres17384 = tmp17387


}

__e.TailApply(PrimFunc(symshen_4comb), W158, ifres17384)
return


}, 1)

tmp17389 := Call(__e, PrimFunc(symshen_4in_1_6), W156)


__e.TailApply(tmp17383, tmp17389)
return


}, 1)

tmp17390 := Call(__e, PrimFunc(symshen_4_5_1out), W156)


__e.TailApply(tmp17382, tmp17390)
return


}


}, 1)

tmp17393 := Call(__e, PrimFunc(sym_5_b_6), V147)


tmp17394 := Call(__e, tmp17381, tmp17393)


__e.TailApply(tmp17378, tmp17394)
return


} else {
__e.Return(W148)
return
}


}, 1)

tmp17397 := MakeNative(func(__e *ControlFlow) {
W149 := __e.Get(1)
_ = W149
tmp17412 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W149)


if True == tmp17412 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17398 := MakeNative(func(__e *ControlFlow) {
W150 := __e.Get(1)
_ = W150
tmp17399 := MakeNative(func(__e *ControlFlow) {
W151 := __e.Get(1)
_ = W151
tmp17400 := MakeNative(func(__e *ControlFlow) {
W152 := __e.Get(1)
_ = W152
tmp17407 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W152)


if True == tmp17407 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17401 := MakeNative(func(__e *ControlFlow) {
W153 := __e.Get(1)
_ = W153
tmp17402 := MakeNative(func(__e *ControlFlow) {
W154 := __e.Get(1)
_ = W154
tmp17403 := PrimCons(W150, W153)

__e.TailApply(PrimFunc(symshen_4comb), W154, tmp17403)
return


}, 1)

tmp17404 := Call(__e, PrimFunc(symshen_4in_1_6), W152)


__e.TailApply(tmp17402, tmp17404)
return


}, 1)

tmp17405 := Call(__e, PrimFunc(symshen_4_5_1out), W152)


__e.TailApply(tmp17401, tmp17405)
return


}


}, 1)

tmp17408 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), W151)


__e.TailApply(tmp17400, tmp17408)
return


}, 1)

tmp17409 := Call(__e, PrimFunc(symshen_4in_1_6), W149)


__e.TailApply(tmp17399, tmp17409)
return


}, 1)

tmp17410 := Call(__e, PrimFunc(symshen_4_5_1out), W149)


__e.TailApply(tmp17398, tmp17410)
return


}


}, 1)

tmp17413 := Call(__e, PrimFunc(symshen_4_5c_1rule_6), V147)


tmp17414 := Call(__e, tmp17397, tmp17413)


__e.TailApply(tmp17377, tmp17414)
return


}, 1)

tmp17415 := Call(__e, ns2_1set, symshen_4_5c_1rules_6, tmp17376)


_ = tmp17415

tmp17416 := MakeNative(func(__e *ControlFlow) {
V159 := __e.Get(1)
_ = V159
tmp17417 := MakeNative(func(__e *ControlFlow) {
W160 := __e.Get(1)
_ = W160
tmp17440 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W160)


if True == tmp17440 {
tmp17418 := MakeNative(func(__e *ControlFlow) {
W169 := __e.Get(1)
_ = W169
tmp17420 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W169)


if True == tmp17420 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W169)
return
}


}, 1)

tmp17421 := MakeNative(func(__e *ControlFlow) {
W170 := __e.Get(1)
_ = W170
tmp17436 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W170)


if True == tmp17436 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17422 := MakeNative(func(__e *ControlFlow) {
W171 := __e.Get(1)
_ = W171
tmp17423 := MakeNative(func(__e *ControlFlow) {
W172 := __e.Get(1)
_ = W172
tmp17424 := MakeNative(func(__e *ControlFlow) {
W173 := __e.Get(1)
_ = W173
tmp17431 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W173)


if True == tmp17431 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17425 := MakeNative(func(__e *ControlFlow) {
W174 := __e.Get(1)
_ = W174
tmp17426 := Call(__e, PrimFunc(symshen_4autocomplete), W171)


tmp17427 := PrimCons(tmp17426, Nil)

tmp17428 := PrimCons(W171, tmp17427)

__e.TailApply(PrimFunc(symshen_4comb), W174, tmp17428)
return


}, 1)

tmp17429 := Call(__e, PrimFunc(symshen_4in_1_6), W173)


__e.TailApply(tmp17425, tmp17429)
return


}


}, 1)

tmp17432 := Call(__e, PrimFunc(symshen_4_5sc_6), W172)


__e.TailApply(tmp17424, tmp17432)
return


}, 1)

tmp17433 := Call(__e, PrimFunc(symshen_4in_1_6), W170)


__e.TailApply(tmp17423, tmp17433)
return


}, 1)

tmp17434 := Call(__e, PrimFunc(symshen_4_5_1out), W170)


__e.TailApply(tmp17422, tmp17434)
return


}


}, 1)

tmp17437 := Call(__e, PrimFunc(symshen_4_5syntax_6), V159)


tmp17438 := Call(__e, tmp17421, tmp17437)


__e.TailApply(tmp17418, tmp17438)
return


} else {
__e.Return(W160)
return
}


}, 1)

tmp17441 := MakeNative(func(__e *ControlFlow) {
W161 := __e.Get(1)
_ = W161
tmp17463 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W161)


if True == tmp17463 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17442 := MakeNative(func(__e *ControlFlow) {
W162 := __e.Get(1)
_ = W162
tmp17443 := MakeNative(func(__e *ControlFlow) {
W163 := __e.Get(1)
_ = W163
tmp17444 := MakeNative(func(__e *ControlFlow) {
W164 := __e.Get(1)
_ = W164
tmp17458 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W164)


if True == tmp17458 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17445 := MakeNative(func(__e *ControlFlow) {
W165 := __e.Get(1)
_ = W165
tmp17446 := MakeNative(func(__e *ControlFlow) {
W166 := __e.Get(1)
_ = W166
tmp17447 := MakeNative(func(__e *ControlFlow) {
W167 := __e.Get(1)
_ = W167
tmp17453 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W167)


if True == tmp17453 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17448 := MakeNative(func(__e *ControlFlow) {
W168 := __e.Get(1)
_ = W168
tmp17449 := PrimCons(W165, Nil)

tmp17450 := PrimCons(W162, tmp17449)

__e.TailApply(PrimFunc(symshen_4comb), W168, tmp17450)
return


}, 1)

tmp17451 := Call(__e, PrimFunc(symshen_4in_1_6), W167)


__e.TailApply(tmp17448, tmp17451)
return


}


}, 1)

tmp17454 := Call(__e, PrimFunc(symshen_4_5sc_6), W166)


__e.TailApply(tmp17447, tmp17454)
return


}, 1)

tmp17455 := Call(__e, PrimFunc(symshen_4in_1_6), W164)


__e.TailApply(tmp17446, tmp17455)
return


}, 1)

tmp17456 := Call(__e, PrimFunc(symshen_4_5_1out), W164)


__e.TailApply(tmp17445, tmp17456)
return


}


}, 1)

tmp17459 := Call(__e, PrimFunc(symshen_4_5semantics_6), W163)


__e.TailApply(tmp17444, tmp17459)
return


}, 1)

tmp17460 := Call(__e, PrimFunc(symshen_4in_1_6), W161)


__e.TailApply(tmp17443, tmp17460)
return


}, 1)

tmp17461 := Call(__e, PrimFunc(symshen_4_5_1out), W161)


__e.TailApply(tmp17442, tmp17461)
return


}


}, 1)

tmp17464 := Call(__e, PrimFunc(symshen_4_5syntax_6), V159)


tmp17465 := Call(__e, tmp17441, tmp17464)


__e.TailApply(tmp17417, tmp17465)
return


}, 1)

tmp17466 := Call(__e, ns2_1set, symshen_4_5c_1rule_6, tmp17416)


_ = tmp17466

tmp17467 := MakeNative(func(__e *ControlFlow) {
V175 := __e.Get(1)
_ = V175
tmp17496 := PrimIsPair(V175)

var ifres17488 Obj

if True == tmp17496 {
tmp17494 := PrimTail(V175)

tmp17495 := PrimEqual(Nil, tmp17494)

var ifres17490 Obj

if True == tmp17495 {
tmp17492 := PrimHead(V175)

tmp17493 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17492)


var ifres17491 Obj

if True == tmp17493 {
ifres17491 = True


} else {
ifres17491 = False


}

ifres17490 = ifres17491


} else {
ifres17490 = False


}

var ifres17489 Obj

if True == ifres17490 {
ifres17489 = True


} else {
ifres17489 = False


}

ifres17488 = ifres17489


} else {
ifres17488 = False


}

if True == ifres17488 {
__e.Return(PrimHead(V175))
return
} else {
tmp17486 := PrimIsPair(V175)

var ifres17482 Obj

if True == tmp17486 {
tmp17484 := PrimHead(V175)

tmp17485 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17484)


var ifres17483 Obj

if True == tmp17485 {
ifres17483 = True


} else {
ifres17483 = False


}

ifres17482 = ifres17483


} else {
ifres17482 = False


}

if True == ifres17482 {
tmp17468 := PrimHead(V175)

tmp17469 := PrimTail(V175)

tmp17470 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17469)


tmp17471 := PrimCons(tmp17470, Nil)

tmp17472 := PrimCons(tmp17468, tmp17471)

__e.Return(PrimCons(symappend, tmp17472))
return


} else {
tmp17480 := PrimIsPair(V175)

if True == tmp17480 {
tmp17473 := PrimHead(V175)

tmp17474 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17473)


tmp17475 := PrimTail(V175)

tmp17476 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17475)


tmp17477 := PrimCons(tmp17476, Nil)

tmp17478 := PrimCons(tmp17474, tmp17477)

__e.Return(PrimCons(symcons, tmp17478))
return


} else {
__e.Return(V175)
return
}


}


}


}, 1)

tmp17497 := Call(__e, ns2_1set, symshen_4autocomplete, tmp17467)


_ = tmp17497

tmp17498 := MakeNative(func(__e *ControlFlow) {
V176 := __e.Get(1)
_ = V176
tmp17505 := PrimIsSymbol(V176)

if True == tmp17505 {
tmp17500 := MakeNative(func(__e *ControlFlow) {
W177 := __e.Get(1)
_ = W177
tmp17501 := MakeNative(func(__e *ControlFlow) {
Z178 := __e.Get(1)
_ = Z178
__e.TailApply(PrimFunc(symshen_4_5non_1terminal_2_6), Z178)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp17501, W177)
return


}, 1)

tmp17502 := Call(__e, PrimFunc(symexplode), V176)


tmp17503 := Call(__e, tmp17500, tmp17502)


if True == tmp17503 {
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

tmp17506 := Call(__e, ns2_1set, symshen_4non_1terminal_2, tmp17498)


_ = tmp17506

tmp17507 := MakeNative(func(__e *ControlFlow) {
V179 := __e.Get(1)
_ = V179
tmp17508 := MakeNative(func(__e *ControlFlow) {
W180 := __e.Get(1)
_ = W180
tmp17530 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W180)


if True == tmp17530 {
tmp17509 := MakeNative(func(__e *ControlFlow) {
W185 := __e.Get(1)
_ = W185
tmp17521 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W185)


if True == tmp17521 {
tmp17510 := MakeNative(func(__e *ControlFlow) {
W188 := __e.Get(1)
_ = W188
tmp17512 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W188)


if True == tmp17512 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W188)
return
}


}, 1)

tmp17513 := MakeNative(func(__e *ControlFlow) {
W189 := __e.Get(1)
_ = W189
tmp17517 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W189)


if True == tmp17517 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17514 := MakeNative(func(__e *ControlFlow) {
W190 := __e.Get(1)
_ = W190
__e.TailApply(PrimFunc(symshen_4comb), W190, False)
return
}, 1)

tmp17515 := Call(__e, PrimFunc(symshen_4in_1_6), W189)


__e.TailApply(tmp17514, tmp17515)
return


}


}, 1)

tmp17518 := Call(__e, PrimFunc(sym_5_b_6), V179)


tmp17519 := Call(__e, tmp17513, tmp17518)


__e.TailApply(tmp17510, tmp17519)
return


} else {
__e.Return(W185)
return
}


}, 1)

tmp17522 := MakeNative(func(__e *ControlFlow) {
W186 := __e.Get(1)
_ = W186
tmp17526 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W186)


if True == tmp17526 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17523 := MakeNative(func(__e *ControlFlow) {
W187 := __e.Get(1)
_ = W187
__e.TailApply(PrimFunc(symshen_4comb), W187, True)
return
}, 1)

tmp17524 := Call(__e, PrimFunc(symshen_4in_1_6), W186)


__e.TailApply(tmp17523, tmp17524)
return


}


}, 1)

tmp17527 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), V179)


tmp17528 := Call(__e, tmp17522, tmp17527)


__e.TailApply(tmp17509, tmp17528)
return


} else {
__e.Return(W180)
return
}


}, 1)

tmp17531 := MakeNative(func(__e *ControlFlow) {
W181 := __e.Get(1)
_ = W181
tmp17541 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W181)


if True == tmp17541 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17532 := MakeNative(func(__e *ControlFlow) {
W182 := __e.Get(1)
_ = W182
tmp17533 := MakeNative(func(__e *ControlFlow) {
W183 := __e.Get(1)
_ = W183
tmp17537 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W183)


if True == tmp17537 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17534 := MakeNative(func(__e *ControlFlow) {
W184 := __e.Get(1)
_ = W184
__e.TailApply(PrimFunc(symshen_4comb), W184, True)
return
}, 1)

tmp17535 := Call(__e, PrimFunc(symshen_4in_1_6), W183)


__e.TailApply(tmp17534, tmp17535)
return


}


}, 1)

tmp17538 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), W182)


__e.TailApply(tmp17533, tmp17538)
return


}, 1)

tmp17539 := Call(__e, PrimFunc(symshen_4in_1_6), W181)


__e.TailApply(tmp17532, tmp17539)
return


}


}, 1)

tmp17542 := Call(__e, PrimFunc(symshen_4_5packagenames_6), V179)


tmp17543 := Call(__e, tmp17531, tmp17542)


__e.TailApply(tmp17508, tmp17543)
return


}, 1)

tmp17544 := Call(__e, ns2_1set, symshen_4_5non_1terminal_2_6, tmp17507)


_ = tmp17544

tmp17545 := MakeNative(func(__e *ControlFlow) {
V191 := __e.Get(1)
_ = V191
tmp17546 := MakeNative(func(__e *ControlFlow) {
W192 := __e.Get(1)
_ = W192
tmp17562 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W192)


if True == tmp17562 {
tmp17547 := MakeNative(func(__e *ControlFlow) {
W198 := __e.Get(1)
_ = W198
tmp17549 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W198)


if True == tmp17549 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W198)
return
}


}, 1)

tmp17550 := MakeNative(func(__e *ControlFlow) {
W199 := __e.Get(1)
_ = W199
tmp17558 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W199)


if True == tmp17558 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17551 := MakeNative(func(__e *ControlFlow) {
W200 := __e.Get(1)
_ = W200
tmp17555 := Call(__e, PrimFunc(symshen_4hds_a_2), W200, MakeString("."))


if True == tmp17555 {
tmp17552 := MakeNative(func(__e *ControlFlow) {
W201 := __e.Get(1)
_ = W201
__e.TailApply(PrimFunc(symshen_4comb), W201, symshen_4skip)
return
}, 1)

tmp17553 := Call(__e, PrimFunc(symtail), W200)


__e.TailApply(tmp17552, tmp17553)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17556 := Call(__e, PrimFunc(symshen_4in_1_6), W199)


__e.TailApply(tmp17551, tmp17556)
return


}


}, 1)

tmp17559 := Call(__e, PrimFunc(symshen_4_5packagename_6), V191)


tmp17560 := Call(__e, tmp17550, tmp17559)


__e.TailApply(tmp17547, tmp17560)
return


} else {
__e.Return(W192)
return
}


}, 1)

tmp17563 := MakeNative(func(__e *ControlFlow) {
W193 := __e.Get(1)
_ = W193
tmp17577 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W193)


if True == tmp17577 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17564 := MakeNative(func(__e *ControlFlow) {
W194 := __e.Get(1)
_ = W194
tmp17574 := Call(__e, PrimFunc(symshen_4hds_a_2), W194, MakeString("."))


if True == tmp17574 {
tmp17565 := MakeNative(func(__e *ControlFlow) {
W195 := __e.Get(1)
_ = W195
tmp17566 := MakeNative(func(__e *ControlFlow) {
W196 := __e.Get(1)
_ = W196
tmp17570 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W196)


if True == tmp17570 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17567 := MakeNative(func(__e *ControlFlow) {
W197 := __e.Get(1)
_ = W197
__e.TailApply(PrimFunc(symshen_4comb), W197, symshen_4skip)
return
}, 1)

tmp17568 := Call(__e, PrimFunc(symshen_4in_1_6), W196)


__e.TailApply(tmp17567, tmp17568)
return


}


}, 1)

tmp17571 := Call(__e, PrimFunc(symshen_4_5packagenames_6), W195)


__e.TailApply(tmp17566, tmp17571)
return


}, 1)

tmp17572 := Call(__e, PrimFunc(symtail), W194)


__e.TailApply(tmp17565, tmp17572)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17575 := Call(__e, PrimFunc(symshen_4in_1_6), W193)


__e.TailApply(tmp17564, tmp17575)
return


}


}, 1)

tmp17578 := Call(__e, PrimFunc(symshen_4_5packagename_6), V191)


tmp17579 := Call(__e, tmp17563, tmp17578)


__e.TailApply(tmp17546, tmp17579)
return


}, 1)

tmp17580 := Call(__e, ns2_1set, symshen_4_5packagenames_6, tmp17545)


_ = tmp17580

tmp17581 := MakeNative(func(__e *ControlFlow) {
V202 := __e.Get(1)
_ = V202
tmp17582 := MakeNative(func(__e *ControlFlow) {
W203 := __e.Get(1)
_ = W203
tmp17594 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W203)


if True == tmp17594 {
tmp17583 := MakeNative(func(__e *ControlFlow) {
W208 := __e.Get(1)
_ = W208
tmp17585 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W208)


if True == tmp17585 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W208)
return
}


}, 1)

tmp17586 := MakeNative(func(__e *ControlFlow) {
W209 := __e.Get(1)
_ = W209
tmp17590 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W209)


if True == tmp17590 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17587 := MakeNative(func(__e *ControlFlow) {
W210 := __e.Get(1)
_ = W210
__e.TailApply(PrimFunc(symshen_4comb), W210, symshen_4skip)
return
}, 1)

tmp17588 := Call(__e, PrimFunc(symshen_4in_1_6), W209)


__e.TailApply(tmp17587, tmp17588)
return


}


}, 1)

tmp17591 := Call(__e, PrimFunc(sym_5e_6), V202)


tmp17592 := Call(__e, tmp17586, tmp17591)


__e.TailApply(tmp17583, tmp17592)
return


} else {
__e.Return(W203)
return
}


}, 1)

tmp17595 := MakeNative(func(__e *ControlFlow) {
W204 := __e.Get(1)
_ = W204
tmp17605 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W204)


if True == tmp17605 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17596 := MakeNative(func(__e *ControlFlow) {
W205 := __e.Get(1)
_ = W205
tmp17597 := MakeNative(func(__e *ControlFlow) {
W206 := __e.Get(1)
_ = W206
tmp17601 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W206)


if True == tmp17601 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17598 := MakeNative(func(__e *ControlFlow) {
W207 := __e.Get(1)
_ = W207
__e.TailApply(PrimFunc(symshen_4comb), W207, symshen_4skip)
return
}, 1)

tmp17599 := Call(__e, PrimFunc(symshen_4in_1_6), W206)


__e.TailApply(tmp17598, tmp17599)
return


}


}, 1)

tmp17602 := Call(__e, PrimFunc(symshen_4_5packagename_6), W205)


__e.TailApply(tmp17597, tmp17602)
return


}, 1)

tmp17603 := Call(__e, PrimFunc(symshen_4in_1_6), W204)


__e.TailApply(tmp17596, tmp17603)
return


}


}, 1)

tmp17606 := Call(__e, PrimFunc(symshen_4_5packagechar_6), V202)


tmp17607 := Call(__e, tmp17595, tmp17606)


__e.TailApply(tmp17582, tmp17607)
return


}, 1)

tmp17608 := Call(__e, ns2_1set, symshen_4_5packagename_6, tmp17581)


_ = tmp17608

tmp17609 := MakeNative(func(__e *ControlFlow) {
V211 := __e.Get(1)
_ = V211
tmp17610 := MakeNative(func(__e *ControlFlow) {
W212 := __e.Get(1)
_ = W212
tmp17612 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W212)


if True == tmp17612 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W212)
return
}


}, 1)

tmp17623 := PrimIsPair(V211)

var ifres17613 Obj

if True == tmp17623 {
tmp17614 := MakeNative(func(__e *ControlFlow) {
W213 := __e.Get(1)
_ = W213
tmp17615 := MakeNative(func(__e *ControlFlow) {
W214 := __e.Get(1)
_ = W214
tmp17617 := PrimEqual(W213, MakeString("."))

tmp17618 := PrimNot(tmp17617)

if True == tmp17618 {
__e.TailApply(PrimFunc(symshen_4comb), W214, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17619 := Call(__e, PrimFunc(symtail), V211)


__e.TailApply(tmp17615, tmp17619)
return


}, 1)

tmp17620 := Call(__e, PrimFunc(symhead), V211)


tmp17621 := Call(__e, tmp17614, tmp17620)


ifres17613 = tmp17621


} else {
tmp17622 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17613 = tmp17622


}

__e.TailApply(tmp17610, ifres17613)
return


}, 1)

tmp17624 := Call(__e, ns2_1set, symshen_4_5packagechar_6, tmp17609)


_ = tmp17624

tmp17625 := MakeNative(func(__e *ControlFlow) {
V215 := __e.Get(1)
_ = V215
tmp17626 := MakeNative(func(__e *ControlFlow) {
W216 := __e.Get(1)
_ = W216
tmp17628 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W216)


if True == tmp17628 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W216)
return
}


}, 1)

tmp17651 := Call(__e, PrimFunc(symshen_4hds_a_2), V215, MakeString("<"))


var ifres17629 Obj

if True == tmp17651 {
tmp17630 := MakeNative(func(__e *ControlFlow) {
W217 := __e.Get(1)
_ = W217
tmp17631 := MakeNative(func(__e *ControlFlow) {
W218 := __e.Get(1)
_ = W218
tmp17646 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W218)


if True == tmp17646 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17632 := MakeNative(func(__e *ControlFlow) {
W219 := __e.Get(1)
_ = W219
tmp17633 := MakeNative(func(__e *ControlFlow) {
W220 := __e.Get(1)
_ = W220
tmp17635 := MakeNative(func(__e *ControlFlow) {
W221 := __e.Get(1)
_ = W221
tmp17640 := PrimIsPair(W221)

if True == tmp17640 {
tmp17637 := PrimHead(W221)

tmp17638 := PrimEqual(tmp17637, MakeString(">"))

if True == tmp17638 {
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

tmp17641 := Call(__e, PrimFunc(symreverse), W219)


tmp17642 := Call(__e, tmp17635, tmp17641)


if True == tmp17642 {
__e.TailApply(PrimFunc(symshen_4comb), W220, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17643 := Call(__e, PrimFunc(symshen_4in_1_6), W218)


__e.TailApply(tmp17633, tmp17643)
return


}, 1)

tmp17644 := Call(__e, PrimFunc(symshen_4_5_1out), W218)


__e.TailApply(tmp17632, tmp17644)
return


}


}, 1)

tmp17647 := Call(__e, PrimFunc(sym_5_b_6), W217)


__e.TailApply(tmp17631, tmp17647)
return


}, 1)

tmp17648 := Call(__e, PrimFunc(symtail), V215)


tmp17649 := Call(__e, tmp17630, tmp17648)


ifres17629 = tmp17649


} else {
tmp17650 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17629 = tmp17650


}

__e.TailApply(tmp17626, ifres17629)
return


}, 1)

tmp17652 := Call(__e, ns2_1set, symshen_4_5non_1terminal_1name_6, tmp17625)


_ = tmp17652

tmp17653 := MakeNative(func(__e *ControlFlow) {
V222 := __e.Get(1)
_ = V222
tmp17654 := PrimIntern(MakeString(";"))

__e.Return(PrimEqual(V222, tmp17654))
return


}, 1)

tmp17655 := Call(__e, ns2_1set, symshen_4semicolon_2, tmp17653)


_ = tmp17655

tmp17656 := MakeNative(func(__e *ControlFlow) {
V223 := __e.Get(1)
_ = V223
tmp17657 := MakeNative(func(__e *ControlFlow) {
W224 := __e.Get(1)
_ = W224
tmp17659 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W224)


if True == tmp17659 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W224)
return
}


}, 1)

tmp17669 := PrimIsPair(V223)

var ifres17660 Obj

if True == tmp17669 {
tmp17661 := MakeNative(func(__e *ControlFlow) {
W225 := __e.Get(1)
_ = W225
tmp17662 := MakeNative(func(__e *ControlFlow) {
W226 := __e.Get(1)
_ = W226
tmp17664 := Call(__e, PrimFunc(symshen_4colon_1equal_2), W225)


if True == tmp17664 {
__e.TailApply(PrimFunc(symshen_4comb), W226, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17665 := Call(__e, PrimFunc(symtail), V223)


__e.TailApply(tmp17662, tmp17665)
return


}, 1)

tmp17666 := Call(__e, PrimFunc(symhead), V223)


tmp17667 := Call(__e, tmp17661, tmp17666)


ifres17660 = tmp17667


} else {
tmp17668 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17660 = tmp17668


}

__e.TailApply(tmp17657, ifres17660)
return


}, 1)

tmp17670 := Call(__e, ns2_1set, symshen_4_5colon_1equal_6, tmp17656)


_ = tmp17670

tmp17671 := MakeNative(func(__e *ControlFlow) {
V227 := __e.Get(1)
_ = V227
tmp17672 := PrimIntern(MakeString(":="))

__e.Return(PrimEqual(tmp17672, V227))
return


}, 1)

tmp17673 := Call(__e, ns2_1set, symshen_4colon_1equal_2, tmp17671)


_ = tmp17673

tmp17674 := MakeNative(func(__e *ControlFlow) {
V228 := __e.Get(1)
_ = V228
tmp17675 := MakeNative(func(__e *ControlFlow) {
W229 := __e.Get(1)
_ = W229
tmp17690 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W229)


if True == tmp17690 {
tmp17676 := MakeNative(func(__e *ControlFlow) {
W236 := __e.Get(1)
_ = W236
tmp17678 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W236)


if True == tmp17678 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W236)
return
}


}, 1)

tmp17679 := MakeNative(func(__e *ControlFlow) {
W237 := __e.Get(1)
_ = W237
tmp17686 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W237)


if True == tmp17686 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17680 := MakeNative(func(__e *ControlFlow) {
W238 := __e.Get(1)
_ = W238
tmp17681 := MakeNative(func(__e *ControlFlow) {
W239 := __e.Get(1)
_ = W239
tmp17682 := PrimCons(W238, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W239, tmp17682)
return


}, 1)

tmp17683 := Call(__e, PrimFunc(symshen_4in_1_6), W237)


__e.TailApply(tmp17681, tmp17683)
return


}, 1)

tmp17684 := Call(__e, PrimFunc(symshen_4_5_1out), W237)


__e.TailApply(tmp17680, tmp17684)
return


}


}, 1)

tmp17687 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V228)


tmp17688 := Call(__e, tmp17679, tmp17687)


__e.TailApply(tmp17676, tmp17688)
return


} else {
__e.Return(W229)
return
}


}, 1)

tmp17691 := MakeNative(func(__e *ControlFlow) {
W230 := __e.Get(1)
_ = W230
tmp17706 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W230)


if True == tmp17706 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17692 := MakeNative(func(__e *ControlFlow) {
W231 := __e.Get(1)
_ = W231
tmp17693 := MakeNative(func(__e *ControlFlow) {
W232 := __e.Get(1)
_ = W232
tmp17694 := MakeNative(func(__e *ControlFlow) {
W233 := __e.Get(1)
_ = W233
tmp17701 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W233)


if True == tmp17701 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17695 := MakeNative(func(__e *ControlFlow) {
W234 := __e.Get(1)
_ = W234
tmp17696 := MakeNative(func(__e *ControlFlow) {
W235 := __e.Get(1)
_ = W235
tmp17697 := PrimCons(W231, W234)

__e.TailApply(PrimFunc(symshen_4comb), W235, tmp17697)
return


}, 1)

tmp17698 := Call(__e, PrimFunc(symshen_4in_1_6), W233)


__e.TailApply(tmp17696, tmp17698)
return


}, 1)

tmp17699 := Call(__e, PrimFunc(symshen_4_5_1out), W233)


__e.TailApply(tmp17695, tmp17699)
return


}


}, 1)

tmp17702 := Call(__e, PrimFunc(symshen_4_5syntax_6), W232)


__e.TailApply(tmp17694, tmp17702)
return


}, 1)

tmp17703 := Call(__e, PrimFunc(symshen_4in_1_6), W230)


__e.TailApply(tmp17693, tmp17703)
return


}, 1)

tmp17704 := Call(__e, PrimFunc(symshen_4_5_1out), W230)


__e.TailApply(tmp17692, tmp17704)
return


}


}, 1)

tmp17707 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V228)


tmp17708 := Call(__e, tmp17691, tmp17707)


__e.TailApply(tmp17675, tmp17708)
return


}, 1)

tmp17709 := Call(__e, ns2_1set, symshen_4_5syntax_6, tmp17674)


_ = tmp17709

tmp17710 := MakeNative(func(__e *ControlFlow) {
V240 := __e.Get(1)
_ = V240
tmp17711 := MakeNative(func(__e *ControlFlow) {
W241 := __e.Get(1)
_ = W241
tmp17713 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W241)


if True == tmp17713 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W241)
return
}


}, 1)

tmp17723 := PrimIsPair(V240)

var ifres17714 Obj

if True == tmp17723 {
tmp17715 := MakeNative(func(__e *ControlFlow) {
W242 := __e.Get(1)
_ = W242
tmp17716 := MakeNative(func(__e *ControlFlow) {
W243 := __e.Get(1)
_ = W243
tmp17718 := Call(__e, PrimFunc(symshen_4syntax_1item_2), W242)


if True == tmp17718 {
__e.TailApply(PrimFunc(symshen_4comb), W243, W242)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17719 := Call(__e, PrimFunc(symtail), V240)


__e.TailApply(tmp17716, tmp17719)
return


}, 1)

tmp17720 := Call(__e, PrimFunc(symhead), V240)


tmp17721 := Call(__e, tmp17715, tmp17720)


ifres17714 = tmp17721


} else {
tmp17722 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17714 = tmp17722


}

__e.TailApply(tmp17711, ifres17714)
return


}, 1)

tmp17724 := Call(__e, ns2_1set, symshen_4_5syntax_1item_6, tmp17710)


_ = tmp17724

tmp17725 := MakeNative(func(__e *ControlFlow) {
V246 := __e.Get(1)
_ = V246
tmp17761 := Call(__e, PrimFunc(symshen_4colon_1equal_2), V246)


if True == tmp17761 {
__e.Return(False)
return
} else {
tmp17759 := Call(__e, PrimFunc(symshen_4semicolon_2), V246)


if True == tmp17759 {
__e.Return(False)
return
} else {
tmp17757 := Call(__e, PrimFunc(symatom_2), V246)


if True == tmp17757 {
__e.Return(True)
return
} else {
tmp17755 := PrimIsPair(V246)

var ifres17736 Obj

if True == tmp17755 {
tmp17753 := PrimHead(V246)

tmp17754 := PrimEqual(symcons, tmp17753)

var ifres17738 Obj

if True == tmp17754 {
tmp17751 := PrimTail(V246)

tmp17752 := PrimIsPair(tmp17751)

var ifres17740 Obj

if True == tmp17752 {
tmp17748 := PrimTail(V246)

tmp17749 := PrimTail(tmp17748)

tmp17750 := PrimIsPair(tmp17749)

var ifres17742 Obj

if True == tmp17750 {
tmp17744 := PrimTail(V246)

tmp17745 := PrimTail(tmp17744)

tmp17746 := PrimTail(tmp17745)

tmp17747 := PrimEqual(Nil, tmp17746)

var ifres17743 Obj

if True == tmp17747 {
ifres17743 = True


} else {
ifres17743 = False


}

ifres17742 = ifres17743


} else {
ifres17742 = False


}

var ifres17741 Obj

if True == ifres17742 {
ifres17741 = True


} else {
ifres17741 = False


}

ifres17740 = ifres17741


} else {
ifres17740 = False


}

var ifres17739 Obj

if True == ifres17740 {
ifres17739 = True


} else {
ifres17739 = False


}

ifres17738 = ifres17739


} else {
ifres17738 = False


}

var ifres17737 Obj

if True == ifres17738 {
ifres17737 = True


} else {
ifres17737 = False


}

ifres17736 = ifres17737


} else {
ifres17736 = False


}

if True == ifres17736 {
tmp17732 := PrimTail(V246)

tmp17733 := PrimHead(tmp17732)

tmp17734 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp17733)


if True == tmp17734 {
tmp17727 := PrimTail(V246)

tmp17728 := PrimTail(tmp17727)

tmp17729 := PrimHead(tmp17728)

tmp17730 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp17729)


if True == tmp17730 {
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


} else {
__e.Return(False)
return
}


}


}


}


}, 1)

tmp17762 := Call(__e, ns2_1set, symshen_4syntax_1item_2, tmp17725)


_ = tmp17762

tmp17763 := MakeNative(func(__e *ControlFlow) {
V247 := __e.Get(1)
_ = V247
tmp17764 := MakeNative(func(__e *ControlFlow) {
W248 := __e.Get(1)
_ = W248
tmp17785 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W248)


if True == tmp17785 {
tmp17765 := MakeNative(func(__e *ControlFlow) {
W256 := __e.Get(1)
_ = W256
tmp17767 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W256)


if True == tmp17767 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W256)
return
}


}, 1)

tmp17768 := MakeNative(func(__e *ControlFlow) {
W257 := __e.Get(1)
_ = W257
tmp17781 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W257)


if True == tmp17781 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17769 := MakeNative(func(__e *ControlFlow) {
W258 := __e.Get(1)
_ = W258
tmp17778 := PrimIsPair(W258)

if True == tmp17778 {
tmp17770 := MakeNative(func(__e *ControlFlow) {
W259 := __e.Get(1)
_ = W259
tmp17771 := MakeNative(func(__e *ControlFlow) {
W260 := __e.Get(1)
_ = W260
tmp17773 := Call(__e, PrimFunc(symshen_4semicolon_2), W259)


tmp17774 := PrimNot(tmp17773)

if True == tmp17774 {
__e.TailApply(PrimFunc(symshen_4comb), W260, W259)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17775 := Call(__e, PrimFunc(symtail), W258)


__e.TailApply(tmp17771, tmp17775)
return


}, 1)

tmp17776 := Call(__e, PrimFunc(symhead), W258)


__e.TailApply(tmp17770, tmp17776)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17779 := Call(__e, PrimFunc(symshen_4in_1_6), W257)


__e.TailApply(tmp17769, tmp17779)
return


}


}, 1)

tmp17782 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V247)


tmp17783 := Call(__e, tmp17768, tmp17782)


__e.TailApply(tmp17765, tmp17783)
return


} else {
__e.Return(W248)
return
}


}, 1)

tmp17786 := MakeNative(func(__e *ControlFlow) {
W249 := __e.Get(1)
_ = W249
tmp17812 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W249)


if True == tmp17812 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17787 := MakeNative(func(__e *ControlFlow) {
W250 := __e.Get(1)
_ = W250
tmp17809 := PrimIsPair(W250)

if True == tmp17809 {
tmp17788 := MakeNative(func(__e *ControlFlow) {
W251 := __e.Get(1)
_ = W251
tmp17789 := MakeNative(func(__e *ControlFlow) {
W252 := __e.Get(1)
_ = W252
tmp17805 := Call(__e, PrimFunc(symshen_4hds_a_2), W252, symwhere)


if True == tmp17805 {
tmp17790 := MakeNative(func(__e *ControlFlow) {
W253 := __e.Get(1)
_ = W253
tmp17802 := PrimIsPair(W253)

if True == tmp17802 {
tmp17791 := MakeNative(func(__e *ControlFlow) {
W254 := __e.Get(1)
_ = W254
tmp17792 := MakeNative(func(__e *ControlFlow) {
W255 := __e.Get(1)
_ = W255
tmp17797 := Call(__e, PrimFunc(symshen_4semicolon_2), W251)


tmp17798 := PrimNot(tmp17797)

if True == tmp17798 {
tmp17793 := PrimCons(W251, Nil)

tmp17794 := PrimCons(W254, tmp17793)

tmp17795 := PrimCons(symwhere, tmp17794)

__e.TailApply(PrimFunc(symshen_4comb), W255, tmp17795)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17799 := Call(__e, PrimFunc(symtail), W253)


__e.TailApply(tmp17792, tmp17799)
return


}, 1)

tmp17800 := Call(__e, PrimFunc(symhead), W253)


__e.TailApply(tmp17791, tmp17800)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17803 := Call(__e, PrimFunc(symtail), W252)


__e.TailApply(tmp17790, tmp17803)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17806 := Call(__e, PrimFunc(symtail), W250)


__e.TailApply(tmp17789, tmp17806)
return


}, 1)

tmp17807 := Call(__e, PrimFunc(symhead), W250)


__e.TailApply(tmp17788, tmp17807)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17810 := Call(__e, PrimFunc(symshen_4in_1_6), W249)


__e.TailApply(tmp17787, tmp17810)
return


}


}, 1)

tmp17813 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V247)


tmp17814 := Call(__e, tmp17786, tmp17813)


__e.TailApply(tmp17764, tmp17814)
return


}, 1)

tmp17815 := Call(__e, ns2_1set, symshen_4_5semantics_6, tmp17763)


_ = tmp17815

tmp17816 := MakeNative(func(__e *ControlFlow) {
V269 := __e.Get(1)
_ = V269
V270 := __e.Get(2)
_ = V270
V271 := __e.Get(3)
_ = V271
tmp17824 := PrimEqual(Nil, V271)

if True == tmp17824 {
__e.Return(PrimCons(symshen_4parse_1failure, Nil))
return
} else {
tmp17822 := PrimIsPair(V271)

if True == tmp17822 {
tmp17817 := PrimHead(V271)

tmp17818 := Call(__e, PrimFunc(symshen_4c_1rule_1_6shen), V269, tmp17817, V270)


tmp17819 := PrimTail(V271)

tmp17820 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), V269, V270, tmp17819)


__e.TailApply(PrimFunc(symshen_4combine_1c_1code), tmp17818, tmp17820)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.c-rules->shen\n")))
return
}


}


}, 3)

tmp17825 := Call(__e, ns2_1set, symshen_4c_1rules_1_6shen, tmp17816)


_ = tmp17825

tmp17826 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symfail))
return
}, 0)

tmp17827 := Call(__e, ns2_1set, symshen_4parse_1failure, tmp17826)


_ = tmp17827

tmp17828 := MakeNative(func(__e *ControlFlow) {
V272 := __e.Get(1)
_ = V272
V273 := __e.Get(2)
_ = V273
tmp17829 := PrimCons(symResult, Nil)

tmp17830 := PrimCons(symshen_4parse_1failure_2, tmp17829)

tmp17831 := PrimCons(symResult, Nil)

tmp17832 := PrimCons(V273, tmp17831)

tmp17833 := PrimCons(tmp17830, tmp17832)

tmp17834 := PrimCons(symif, tmp17833)

tmp17835 := PrimCons(tmp17834, Nil)

tmp17836 := PrimCons(V272, tmp17835)

tmp17837 := PrimCons(symResult, tmp17836)

__e.Return(PrimCons(symlet, tmp17837))
return


}, 2)

tmp17838 := Call(__e, ns2_1set, symshen_4combine_1c_1code, tmp17828)


_ = tmp17838

tmp17839 := MakeNative(func(__e *ControlFlow) {
V280 := __e.Get(1)
_ = V280
V281 := __e.Get(2)
_ = V281
V282 := __e.Get(3)
_ = V282
tmp17853 := PrimIsPair(V281)

var ifres17844 Obj

if True == tmp17853 {
tmp17851 := PrimTail(V281)

tmp17852 := PrimIsPair(tmp17851)

var ifres17846 Obj

if True == tmp17852 {
tmp17848 := PrimTail(V281)

tmp17849 := PrimTail(tmp17848)

tmp17850 := PrimEqual(Nil, tmp17849)

var ifres17847 Obj

if True == tmp17850 {
ifres17847 = True


} else {
ifres17847 = False


}

ifres17846 = ifres17847


} else {
ifres17846 = False


}

var ifres17845 Obj

if True == ifres17846 {
ifres17845 = True


} else {
ifres17845 = False


}

ifres17844 = ifres17845


} else {
ifres17844 = False


}

if True == ifres17844 {
tmp17840 := PrimHead(V281)

tmp17841 := PrimTail(V281)

tmp17842 := PrimHead(tmp17841)

__e.TailApply(PrimFunc(symshen_4yacc_1syntax), V280, V282, tmp17840, tmp17842)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.c-rule->shen\n")))
return
}


}, 3)

tmp17854 := Call(__e, ns2_1set, symshen_4c_1rule_1_6shen, tmp17839)


_ = tmp17854

tmp17855 := MakeNative(func(__e *ControlFlow) {
V291 := __e.Get(1)
_ = V291
V292 := __e.Get(2)
_ = V292
V293 := __e.Get(3)
_ = V293
V294 := __e.Get(4)
_ = V294
tmp17919 := PrimEqual(Nil, V293)

var ifres17897 Obj

if True == tmp17919 {
tmp17918 := PrimIsPair(V294)

var ifres17899 Obj

if True == tmp17918 {
tmp17916 := PrimHead(V294)

tmp17917 := PrimEqual(symwhere, tmp17916)

var ifres17901 Obj

if True == tmp17917 {
tmp17914 := PrimTail(V294)

tmp17915 := PrimIsPair(tmp17914)

var ifres17903 Obj

if True == tmp17915 {
tmp17911 := PrimTail(V294)

tmp17912 := PrimTail(tmp17911)

tmp17913 := PrimIsPair(tmp17912)

var ifres17905 Obj

if True == tmp17913 {
tmp17907 := PrimTail(V294)

tmp17908 := PrimTail(tmp17907)

tmp17909 := PrimTail(tmp17908)

tmp17910 := PrimEqual(Nil, tmp17909)

var ifres17906 Obj

if True == tmp17910 {
ifres17906 = True


} else {
ifres17906 = False


}

ifres17905 = ifres17906


} else {
ifres17905 = False


}

var ifres17904 Obj

if True == ifres17905 {
ifres17904 = True


} else {
ifres17904 = False


}

ifres17903 = ifres17904


} else {
ifres17903 = False


}

var ifres17902 Obj

if True == ifres17903 {
ifres17902 = True


} else {
ifres17902 = False


}

ifres17901 = ifres17902


} else {
ifres17901 = False


}

var ifres17900 Obj

if True == ifres17901 {
ifres17900 = True


} else {
ifres17900 = False


}

ifres17899 = ifres17900


} else {
ifres17899 = False


}

var ifres17898 Obj

if True == ifres17899 {
ifres17898 = True


} else {
ifres17898 = False


}

ifres17897 = ifres17898


} else {
ifres17897 = False


}

if True == ifres17897 {
tmp17856 := PrimTail(V294)

tmp17857 := PrimHead(tmp17856)

tmp17858 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), tmp17857)


tmp17859 := PrimTail(V294)

tmp17860 := PrimTail(tmp17859)

tmp17861 := PrimHead(tmp17860)

tmp17862 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V291, V292, Nil, tmp17861)


tmp17863 := PrimCons(symshen_4parse_1failure, Nil)

tmp17864 := PrimCons(tmp17863, Nil)

tmp17865 := PrimCons(tmp17862, tmp17864)

tmp17866 := PrimCons(tmp17858, tmp17865)

__e.Return(PrimCons(symif, tmp17866))
return


} else {
tmp17895 := PrimEqual(Nil, V293)

if True == tmp17895 {
__e.TailApply(PrimFunc(symshen_4yacc_1semantics), V291, V292, V294)
return
} else {
tmp17893 := PrimIsPair(V293)

if True == tmp17893 {
tmp17890 := PrimHead(V293)

tmp17891 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17890)


if True == tmp17891 {
tmp17867 := PrimHead(V293)

tmp17868 := PrimTail(V293)

__e.TailApply(PrimFunc(symshen_4non_1terminalcode), V291, V292, tmp17867, tmp17868, V294)
return


} else {
tmp17887 := PrimHead(V293)

tmp17888 := PrimIsVariable(tmp17887)

if True == tmp17888 {
tmp17869 := PrimHead(V293)

tmp17870 := PrimTail(V293)

__e.TailApply(PrimFunc(symshen_4variablecode), V291, V292, tmp17869, tmp17870, V294)
return


} else {
tmp17884 := PrimHead(V293)

tmp17885 := PrimEqual(sym__, tmp17884)

if True == tmp17885 {
tmp17871 := PrimHead(V293)

tmp17872 := PrimTail(V293)

__e.TailApply(PrimFunc(symshen_4wildcardcode), V291, V292, tmp17871, tmp17872, V294)
return


} else {
tmp17881 := PrimHead(V293)

tmp17882 := Call(__e, PrimFunc(symatom_2), tmp17881)


if True == tmp17882 {
tmp17873 := PrimHead(V293)

tmp17874 := PrimTail(V293)

__e.TailApply(PrimFunc(symshen_4terminalcode), V291, V292, tmp17873, tmp17874, V294)
return


} else {
tmp17878 := PrimHead(V293)

tmp17879 := PrimIsPair(tmp17878)

if True == tmp17879 {
tmp17875 := PrimHead(V293)

tmp17876 := PrimTail(V293)

__e.TailApply(PrimFunc(symshen_4conscode), V291, V292, tmp17875, tmp17876, V294)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n")))
return
}


}


}


}


}


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n")))
return
}


}


}


}, 4)

tmp17920 := Call(__e, ns2_1set, symshen_4yacc_1syntax, tmp17855)


_ = tmp17920

tmp17921 := MakeNative(func(__e *ControlFlow) {
V295 := __e.Get(1)
_ = V295
V296 := __e.Get(2)
_ = V296
V297 := __e.Get(3)
_ = V297
V298 := __e.Get(4)
_ = V298
V299 := __e.Get(5)
_ = V299
tmp17922 := MakeNative(func(__e *ControlFlow) {
W300 := __e.Get(1)
_ = W300
tmp17923 := MakeNative(func(__e *ControlFlow) {
W301 := __e.Get(1)
_ = W301
tmp17924 := MakeNative(func(__e *ControlFlow) {
W302 := __e.Get(1)
_ = W302
tmp17925 := PrimCons(V296, Nil)

tmp17926 := PrimCons(V297, tmp17925)

tmp17927 := PrimCons(W300, Nil)

tmp17928 := PrimCons(symshen_4parse_1failure_2, tmp17927)

tmp17929 := PrimCons(symshen_4parse_1failure, Nil)

tmp17930 := MakeNative(func(__e *ControlFlow) {
W303 := __e.Get(1)
_ = W303
tmp17940 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V297, V299)


var ifres17937 Obj

if True == tmp17940 {
ifres17937 = True


} else {
tmp17939 := Call(__e, PrimFunc(symshen_4occurs_1check_2), W301, V299)


var ifres17938 Obj

if True == tmp17939 {
ifres17938 = True


} else {
ifres17938 = False


}

ifres17937 = ifres17938


}

if True == ifres17937 {
tmp17931 := PrimCons(W300, Nil)

tmp17932 := PrimCons(symshen_4_5_1out, tmp17931)

tmp17933 := PrimCons(W303, Nil)

tmp17934 := PrimCons(tmp17932, tmp17933)

tmp17935 := PrimCons(W301, tmp17934)

__e.Return(PrimCons(symlet, tmp17935))
return


} else {
__e.Return(W303)
return
}


}, 1)

tmp17941 := PrimCons(W300, Nil)

tmp17942 := PrimCons(symshen_4in_1_6, tmp17941)

tmp17943 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V295, W302, V298, V299)


tmp17944 := PrimCons(tmp17943, Nil)

tmp17945 := PrimCons(tmp17942, tmp17944)

tmp17946 := PrimCons(W302, tmp17945)

tmp17947 := PrimCons(symlet, tmp17946)

tmp17948 := Call(__e, tmp17930, tmp17947)


tmp17949 := PrimCons(tmp17948, Nil)

tmp17950 := PrimCons(tmp17929, tmp17949)

tmp17951 := PrimCons(tmp17928, tmp17950)

tmp17952 := PrimCons(symif, tmp17951)

tmp17953 := PrimCons(tmp17952, Nil)

tmp17954 := PrimCons(tmp17926, tmp17953)

tmp17955 := PrimCons(W300, tmp17954)

__e.Return(PrimCons(symlet, tmp17955))
return


}, 1)

tmp17956 := Call(__e, PrimFunc(symconcat), symRemainder, V297)


__e.TailApply(tmp17924, tmp17956)
return


}, 1)

tmp17957 := Call(__e, PrimFunc(symconcat), symAction, V297)


__e.TailApply(tmp17923, tmp17957)
return


}, 1)

tmp17958 := Call(__e, PrimFunc(symconcat), symParse, V297)


__e.TailApply(tmp17922, tmp17958)
return


}, 5)

tmp17959 := Call(__e, ns2_1set, symshen_4non_1terminalcode, tmp17921)


_ = tmp17959

tmp17960 := MakeNative(func(__e *ControlFlow) {
V304 := __e.Get(1)
_ = V304
V305 := __e.Get(2)
_ = V305
V306 := __e.Get(3)
_ = V306
V307 := __e.Get(4)
_ = V307
V308 := __e.Get(5)
_ = V308
tmp17961 := MakeNative(func(__e *ControlFlow) {
W309 := __e.Get(1)
_ = W309
tmp17962 := PrimCons(V305, Nil)

tmp17963 := PrimCons(symcons_2, tmp17962)

tmp17964 := MakeNative(func(__e *ControlFlow) {
W310 := __e.Get(1)
_ = W310
tmp17971 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V306, V308)


if True == tmp17971 {
tmp17965 := PrimCons(V305, Nil)

tmp17966 := PrimCons(symhead, tmp17965)

tmp17967 := PrimCons(W310, Nil)

tmp17968 := PrimCons(tmp17966, tmp17967)

tmp17969 := PrimCons(V306, tmp17968)

__e.Return(PrimCons(symlet, tmp17969))
return


} else {
__e.Return(W310)
return
}


}, 1)

tmp17972 := PrimCons(V305, Nil)

tmp17973 := PrimCons(symtail, tmp17972)

tmp17974 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V304, W309, V307, V308)


tmp17975 := PrimCons(tmp17974, Nil)

tmp17976 := PrimCons(tmp17973, tmp17975)

tmp17977 := PrimCons(W309, tmp17976)

tmp17978 := PrimCons(symlet, tmp17977)

tmp17979 := Call(__e, tmp17964, tmp17978)


tmp17980 := PrimCons(symshen_4parse_1failure, Nil)

tmp17981 := PrimCons(tmp17980, Nil)

tmp17982 := PrimCons(tmp17979, tmp17981)

tmp17983 := PrimCons(tmp17963, tmp17982)

__e.Return(PrimCons(symif, tmp17983))
return


}, 1)

tmp17984 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp17961, tmp17984)
return


}, 5)

tmp17985 := Call(__e, ns2_1set, symshen_4variablecode, tmp17960)


_ = tmp17985

tmp17986 := MakeNative(func(__e *ControlFlow) {
V311 := __e.Get(1)
_ = V311
V312 := __e.Get(2)
_ = V312
V313 := __e.Get(3)
_ = V313
V314 := __e.Get(4)
_ = V314
V315 := __e.Get(5)
_ = V315
tmp17987 := MakeNative(func(__e *ControlFlow) {
W316 := __e.Get(1)
_ = W316
tmp17988 := PrimCons(V312, Nil)

tmp17989 := PrimCons(symcons_2, tmp17988)

tmp17990 := PrimCons(V312, Nil)

tmp17991 := PrimCons(symtail, tmp17990)

tmp17992 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V311, W316, V314, V315)


tmp17993 := PrimCons(tmp17992, Nil)

tmp17994 := PrimCons(tmp17991, tmp17993)

tmp17995 := PrimCons(W316, tmp17994)

tmp17996 := PrimCons(symlet, tmp17995)

tmp17997 := PrimCons(symshen_4parse_1failure, Nil)

tmp17998 := PrimCons(tmp17997, Nil)

tmp17999 := PrimCons(tmp17996, tmp17998)

tmp18000 := PrimCons(tmp17989, tmp17999)

__e.Return(PrimCons(symif, tmp18000))
return


}, 1)

tmp18001 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp17987, tmp18001)
return


}, 5)

tmp18002 := Call(__e, ns2_1set, symshen_4wildcardcode, tmp17986)


_ = tmp18002

tmp18003 := MakeNative(func(__e *ControlFlow) {
V317 := __e.Get(1)
_ = V317
V318 := __e.Get(2)
_ = V318
V319 := __e.Get(3)
_ = V319
V320 := __e.Get(4)
_ = V320
V321 := __e.Get(5)
_ = V321
tmp18004 := MakeNative(func(__e *ControlFlow) {
W322 := __e.Get(1)
_ = W322
tmp18005 := PrimCons(V319, Nil)

tmp18006 := PrimCons(V318, tmp18005)

tmp18007 := PrimCons(symshen_4hds_a_2, tmp18006)

tmp18008 := PrimCons(V318, Nil)

tmp18009 := PrimCons(symtail, tmp18008)

tmp18010 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V317, W322, V320, V321)


tmp18011 := PrimCons(tmp18010, Nil)

tmp18012 := PrimCons(tmp18009, tmp18011)

tmp18013 := PrimCons(W322, tmp18012)

tmp18014 := PrimCons(symlet, tmp18013)

tmp18015 := PrimCons(symshen_4parse_1failure, Nil)

tmp18016 := PrimCons(tmp18015, Nil)

tmp18017 := PrimCons(tmp18014, tmp18016)

tmp18018 := PrimCons(tmp18007, tmp18017)

__e.Return(PrimCons(symif, tmp18018))
return


}, 1)

tmp18019 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp18004, tmp18019)
return


}, 5)

tmp18020 := Call(__e, ns2_1set, symshen_4terminalcode, tmp18003)


_ = tmp18020

tmp18021 := MakeNative(func(__e *ControlFlow) {
V330 := __e.Get(1)
_ = V330
V331 := __e.Get(2)
_ = V331
tmp18027 := PrimIsPair(V330)

var ifres18023 Obj

if True == tmp18027 {
tmp18025 := PrimHead(V330)

tmp18026 := PrimEqual(tmp18025, V331)

var ifres18024 Obj

if True == tmp18026 {
ifres18024 = True


} else {
ifres18024 = False


}

ifres18023 = ifres18024


} else {
ifres18023 = False


}

if True == ifres18023 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp18028 := Call(__e, ns2_1set, symshen_4hds_a_2, tmp18021)


_ = tmp18028

tmp18029 := MakeNative(func(__e *ControlFlow) {
V332 := __e.Get(1)
_ = V332
V333 := __e.Get(2)
_ = V333
V334 := __e.Get(3)
_ = V334
V335 := __e.Get(4)
_ = V335
V336 := __e.Get(5)
_ = V336
tmp18030 := MakeNative(func(__e *ControlFlow) {
W337 := __e.Get(1)
_ = W337
tmp18031 := MakeNative(func(__e *ControlFlow) {
W338 := __e.Get(1)
_ = W338
tmp18032 := MakeNative(func(__e *ControlFlow) {
W339 := __e.Get(1)
_ = W339
tmp18033 := PrimCons(V333, Nil)

tmp18034 := PrimCons(symshen_4ccons_2, tmp18033)

tmp18035 := PrimCons(V333, Nil)

tmp18036 := PrimCons(symhead, tmp18035)

tmp18037 := PrimCons(V333, Nil)

tmp18038 := PrimCons(symtail, tmp18037)

tmp18039 := Call(__e, PrimFunc(symshen_4decons), V334)


tmp18040 := PrimCons(sym_5end_6, Nil)

tmp18041 := Call(__e, PrimFunc(symappend), tmp18039, tmp18040)


tmp18042 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V332, W339, V335, V336)


tmp18043 := PrimCons(tmp18042, Nil)

tmp18044 := PrimCons(symshen_4processed, tmp18043)

tmp18045 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V332, W338, tmp18041, tmp18044)


tmp18046 := PrimCons(tmp18045, Nil)

tmp18047 := PrimCons(tmp18038, tmp18046)

tmp18048 := PrimCons(W339, tmp18047)

tmp18049 := PrimCons(tmp18036, tmp18048)

tmp18050 := PrimCons(W338, tmp18049)

tmp18051 := PrimCons(symlet, tmp18050)

tmp18052 := PrimCons(symshen_4parse_1failure, Nil)

tmp18053 := PrimCons(tmp18052, Nil)

tmp18054 := PrimCons(tmp18051, tmp18053)

tmp18055 := PrimCons(tmp18034, tmp18054)

__e.Return(PrimCons(symif, tmp18055))
return


}, 1)

tmp18056 := Call(__e, PrimFunc(symgensym), symTl)


__e.TailApply(tmp18032, tmp18056)
return


}, 1)

tmp18057 := Call(__e, PrimFunc(symgensym), symHd)


__e.TailApply(tmp18031, tmp18057)
return


}, 1)

tmp18058 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp18030, tmp18058)
return


}, 5)

tmp18059 := Call(__e, ns2_1set, symshen_4conscode, tmp18029)


_ = tmp18059

tmp18060 := MakeNative(func(__e *ControlFlow) {
V348 := __e.Get(1)
_ = V348
tmp18066 := PrimIsPair(V348)

var ifres18062 Obj

if True == tmp18066 {
tmp18064 := PrimHead(V348)

tmp18065 := PrimIsPair(tmp18064)

var ifres18063 Obj

if True == tmp18065 {
ifres18063 = True


} else {
ifres18063 = False


}

ifres18062 = ifres18063


} else {
ifres18062 = False


}

if True == ifres18062 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp18067 := Call(__e, ns2_1set, symshen_4ccons_2, tmp18060)


_ = tmp18067

tmp18068 := MakeNative(func(__e *ControlFlow) {
V349 := __e.Get(1)
_ = V349
tmp18095 := PrimIsPair(V349)

var ifres18076 Obj

if True == tmp18095 {
tmp18093 := PrimHead(V349)

tmp18094 := PrimEqual(symcons, tmp18093)

var ifres18078 Obj

if True == tmp18094 {
tmp18091 := PrimTail(V349)

tmp18092 := PrimIsPair(tmp18091)

var ifres18080 Obj

if True == tmp18092 {
tmp18088 := PrimTail(V349)

tmp18089 := PrimTail(tmp18088)

tmp18090 := PrimIsPair(tmp18089)

var ifres18082 Obj

if True == tmp18090 {
tmp18084 := PrimTail(V349)

tmp18085 := PrimTail(tmp18084)

tmp18086 := PrimTail(tmp18085)

tmp18087 := PrimEqual(Nil, tmp18086)

var ifres18083 Obj

if True == tmp18087 {
ifres18083 = True


} else {
ifres18083 = False


}

ifres18082 = ifres18083


} else {
ifres18082 = False


}

var ifres18081 Obj

if True == ifres18082 {
ifres18081 = True


} else {
ifres18081 = False


}

ifres18080 = ifres18081


} else {
ifres18080 = False


}

var ifres18079 Obj

if True == ifres18080 {
ifres18079 = True


} else {
ifres18079 = False


}

ifres18078 = ifres18079


} else {
ifres18078 = False


}

var ifres18077 Obj

if True == ifres18078 {
ifres18077 = True


} else {
ifres18077 = False


}

ifres18076 = ifres18077


} else {
ifres18076 = False


}

if True == ifres18076 {
tmp18069 := PrimTail(V349)

tmp18070 := PrimHead(tmp18069)

tmp18071 := PrimTail(V349)

tmp18072 := PrimTail(tmp18071)

tmp18073 := PrimHead(tmp18072)

tmp18074 := Call(__e, PrimFunc(symshen_4decons), tmp18073)


__e.Return(PrimCons(tmp18070, tmp18074))
return


} else {
__e.Return(V349)
return
}


}, 1)

tmp18096 := Call(__e, ns2_1set, symshen_4decons, tmp18068)


_ = tmp18096

tmp18097 := MakeNative(func(__e *ControlFlow) {
V350 := __e.Get(1)
_ = V350
V351 := __e.Get(2)
_ = V351
tmp18098 := PrimCons(V351, Nil)

__e.Return(PrimCons(V350, tmp18098))
return


}, 2)

tmp18099 := Call(__e, ns2_1set, symshen_4comb, tmp18097)


_ = tmp18099

tmp18100 := MakeNative(func(__e *ControlFlow) {
V356 := __e.Get(1)
_ = V356
V357 := __e.Get(2)
_ = V357
V358 := __e.Get(3)
_ = V358
tmp18122 := PrimIsPair(V358)

var ifres18109 Obj

if True == tmp18122 {
tmp18120 := PrimHead(V358)

tmp18121 := PrimEqual(symshen_4processed, tmp18120)

var ifres18111 Obj

if True == tmp18121 {
tmp18118 := PrimTail(V358)

tmp18119 := PrimIsPair(tmp18118)

var ifres18113 Obj

if True == tmp18119 {
tmp18115 := PrimTail(V358)

tmp18116 := PrimTail(tmp18115)

tmp18117 := PrimEqual(Nil, tmp18116)

var ifres18114 Obj

if True == tmp18117 {
ifres18114 = True


} else {
ifres18114 = False


}

ifres18113 = ifres18114


} else {
ifres18113 = False


}

var ifres18112 Obj

if True == ifres18113 {
ifres18112 = True


} else {
ifres18112 = False


}

ifres18111 = ifres18112


} else {
ifres18111 = False


}

var ifres18110 Obj

if True == ifres18111 {
ifres18110 = True


} else {
ifres18110 = False


}

ifres18109 = ifres18110


} else {
ifres18109 = False


}

if True == ifres18109 {
tmp18101 := PrimTail(V358)

__e.Return(PrimHead(tmp18101))
return


} else {
tmp18102 := MakeNative(func(__e *ControlFlow) {
W359 := __e.Get(1)
_ = W359
tmp18103 := MakeNative(func(__e *ControlFlow) {
W360 := __e.Get(1)
_ = W360
tmp18104 := PrimCons(W360, Nil)

tmp18105 := PrimCons(V357, tmp18104)

__e.Return(PrimCons(symshen_4comb, tmp18105))
return


}, 1)

tmp18106 := Call(__e, PrimFunc(symshen_4use_1type_1info), V356, W359)


__e.TailApply(tmp18103, tmp18106)
return


}, 1)

tmp18107 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), V358)


__e.TailApply(tmp18102, tmp18107)
return


}


}, 3)

tmp18123 := Call(__e, ns2_1set, symshen_4yacc_1semantics, tmp18100)


_ = tmp18123

tmp18124 := MakeNative(func(__e *ControlFlow) {
V364 := __e.Get(1)
_ = V364
V365 := __e.Get(2)
_ = V365
tmp18312 := PrimIsPair(V364)

var ifres18133 Obj

if True == tmp18312 {
tmp18310 := PrimHead(V364)

tmp18311 := PrimEqual(sym_i, tmp18310)

var ifres18135 Obj

if True == tmp18311 {
tmp18308 := PrimTail(V364)

tmp18309 := PrimIsPair(tmp18308)

var ifres18137 Obj

if True == tmp18309 {
tmp18305 := PrimTail(V364)

tmp18306 := PrimHead(tmp18305)

tmp18307 := PrimIsPair(tmp18306)

var ifres18139 Obj

if True == tmp18307 {
tmp18301 := PrimTail(V364)

tmp18302 := PrimHead(tmp18301)

tmp18303 := PrimHead(tmp18302)

tmp18304 := PrimEqual(symlist, tmp18303)

var ifres18141 Obj

if True == tmp18304 {
tmp18297 := PrimTail(V364)

tmp18298 := PrimHead(tmp18297)

tmp18299 := PrimTail(tmp18298)

tmp18300 := PrimIsPair(tmp18299)

var ifres18143 Obj

if True == tmp18300 {
tmp18292 := PrimTail(V364)

tmp18293 := PrimHead(tmp18292)

tmp18294 := PrimTail(tmp18293)

tmp18295 := PrimTail(tmp18294)

tmp18296 := PrimEqual(Nil, tmp18295)

var ifres18145 Obj

if True == tmp18296 {
tmp18289 := PrimTail(V364)

tmp18290 := PrimTail(tmp18289)

tmp18291 := PrimIsPair(tmp18290)

var ifres18147 Obj

if True == tmp18291 {
tmp18285 := PrimTail(V364)

tmp18286 := PrimTail(tmp18285)

tmp18287 := PrimHead(tmp18286)

tmp18288 := PrimEqual(sym_1_1_6, tmp18287)

var ifres18149 Obj

if True == tmp18288 {
tmp18281 := PrimTail(V364)

tmp18282 := PrimTail(tmp18281)

tmp18283 := PrimTail(tmp18282)

tmp18284 := PrimIsPair(tmp18283)

var ifres18151 Obj

if True == tmp18284 {
tmp18276 := PrimTail(V364)

tmp18277 := PrimTail(tmp18276)

tmp18278 := PrimTail(tmp18277)

tmp18279 := PrimHead(tmp18278)

tmp18280 := PrimIsPair(tmp18279)

var ifres18153 Obj

if True == tmp18280 {
tmp18270 := PrimTail(V364)

tmp18271 := PrimTail(tmp18270)

tmp18272 := PrimTail(tmp18271)

tmp18273 := PrimHead(tmp18272)

tmp18274 := PrimHead(tmp18273)

tmp18275 := PrimEqual(symstr, tmp18274)

var ifres18155 Obj

if True == tmp18275 {
tmp18264 := PrimTail(V364)

tmp18265 := PrimTail(tmp18264)

tmp18266 := PrimTail(tmp18265)

tmp18267 := PrimHead(tmp18266)

tmp18268 := PrimTail(tmp18267)

tmp18269 := PrimIsPair(tmp18268)

var ifres18157 Obj

if True == tmp18269 {
tmp18257 := PrimTail(V364)

tmp18258 := PrimTail(tmp18257)

tmp18259 := PrimTail(tmp18258)

tmp18260 := PrimHead(tmp18259)

tmp18261 := PrimTail(tmp18260)

tmp18262 := PrimHead(tmp18261)

tmp18263 := PrimIsPair(tmp18262)

var ifres18159 Obj

if True == tmp18263 {
tmp18249 := PrimTail(V364)

tmp18250 := PrimTail(tmp18249)

tmp18251 := PrimTail(tmp18250)

tmp18252 := PrimHead(tmp18251)

tmp18253 := PrimTail(tmp18252)

tmp18254 := PrimHead(tmp18253)

tmp18255 := PrimHead(tmp18254)

tmp18256 := PrimEqual(symlist, tmp18255)

var ifres18161 Obj

if True == tmp18256 {
tmp18241 := PrimTail(V364)

tmp18242 := PrimTail(tmp18241)

tmp18243 := PrimTail(tmp18242)

tmp18244 := PrimHead(tmp18243)

tmp18245 := PrimTail(tmp18244)

tmp18246 := PrimHead(tmp18245)

tmp18247 := PrimTail(tmp18246)

tmp18248 := PrimIsPair(tmp18247)

var ifres18163 Obj

if True == tmp18248 {
tmp18232 := PrimTail(V364)

tmp18233 := PrimTail(tmp18232)

tmp18234 := PrimTail(tmp18233)

tmp18235 := PrimHead(tmp18234)

tmp18236 := PrimTail(tmp18235)

tmp18237 := PrimHead(tmp18236)

tmp18238 := PrimTail(tmp18237)

tmp18239 := PrimTail(tmp18238)

tmp18240 := PrimEqual(Nil, tmp18239)

var ifres18165 Obj

if True == tmp18240 {
tmp18225 := PrimTail(V364)

tmp18226 := PrimTail(tmp18225)

tmp18227 := PrimTail(tmp18226)

tmp18228 := PrimHead(tmp18227)

tmp18229 := PrimTail(tmp18228)

tmp18230 := PrimTail(tmp18229)

tmp18231 := PrimIsPair(tmp18230)

var ifres18167 Obj

if True == tmp18231 {
tmp18217 := PrimTail(V364)

tmp18218 := PrimTail(tmp18217)

tmp18219 := PrimTail(tmp18218)

tmp18220 := PrimHead(tmp18219)

tmp18221 := PrimTail(tmp18220)

tmp18222 := PrimTail(tmp18221)

tmp18223 := PrimTail(tmp18222)

tmp18224 := PrimEqual(Nil, tmp18223)

var ifres18169 Obj

if True == tmp18224 {
tmp18212 := PrimTail(V364)

tmp18213 := PrimTail(tmp18212)

tmp18214 := PrimTail(tmp18213)

tmp18215 := PrimTail(tmp18214)

tmp18216 := PrimIsPair(tmp18215)

var ifres18171 Obj

if True == tmp18216 {
tmp18206 := PrimTail(V364)

tmp18207 := PrimTail(tmp18206)

tmp18208 := PrimTail(tmp18207)

tmp18209 := PrimTail(tmp18208)

tmp18210 := PrimHead(tmp18209)

tmp18211 := PrimEqual(sym_j, tmp18210)

var ifres18173 Obj

if True == tmp18211 {
tmp18200 := PrimTail(V364)

tmp18201 := PrimTail(tmp18200)

tmp18202 := PrimTail(tmp18201)

tmp18203 := PrimTail(tmp18202)

tmp18204 := PrimTail(tmp18203)

tmp18205 := PrimEqual(Nil, tmp18204)

var ifres18175 Obj

if True == tmp18205 {
tmp18187 := PrimTail(V364)

tmp18188 := PrimHead(tmp18187)

tmp18189 := PrimTail(tmp18188)

tmp18190 := PrimHead(tmp18189)

tmp18191 := PrimTail(V364)

tmp18192 := PrimTail(tmp18191)

tmp18193 := PrimTail(tmp18192)

tmp18194 := PrimHead(tmp18193)

tmp18195 := PrimTail(tmp18194)

tmp18196 := PrimHead(tmp18195)

tmp18197 := PrimTail(tmp18196)

tmp18198 := PrimHead(tmp18197)

tmp18199 := PrimEqual(tmp18190, tmp18198)

var ifres18177 Obj

if True == tmp18199 {
tmp18179 := PrimTail(V364)

tmp18180 := PrimTail(tmp18179)

tmp18181 := PrimTail(tmp18180)

tmp18182 := PrimHead(tmp18181)

tmp18183 := PrimTail(tmp18182)

tmp18184 := PrimTail(tmp18183)

tmp18185 := PrimHead(tmp18184)

tmp18186 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18185)


var ifres18178 Obj

if True == tmp18186 {
ifres18178 = True


} else {
ifres18178 = False


}

ifres18177 = ifres18178


} else {
ifres18177 = False


}

var ifres18176 Obj

if True == ifres18177 {
ifres18176 = True


} else {
ifres18176 = False


}

ifres18175 = ifres18176


} else {
ifres18175 = False


}

var ifres18174 Obj

if True == ifres18175 {
ifres18174 = True


} else {
ifres18174 = False


}

ifres18173 = ifres18174


} else {
ifres18173 = False


}

var ifres18172 Obj

if True == ifres18173 {
ifres18172 = True


} else {
ifres18172 = False


}

ifres18171 = ifres18172


} else {
ifres18171 = False


}

var ifres18170 Obj

if True == ifres18171 {
ifres18170 = True


} else {
ifres18170 = False


}

ifres18169 = ifres18170


} else {
ifres18169 = False


}

var ifres18168 Obj

if True == ifres18169 {
ifres18168 = True


} else {
ifres18168 = False


}

ifres18167 = ifres18168


} else {
ifres18167 = False


}

var ifres18166 Obj

if True == ifres18167 {
ifres18166 = True


} else {
ifres18166 = False


}

ifres18165 = ifres18166


} else {
ifres18165 = False


}

var ifres18164 Obj

if True == ifres18165 {
ifres18164 = True


} else {
ifres18164 = False


}

ifres18163 = ifres18164


} else {
ifres18163 = False


}

var ifres18162 Obj

if True == ifres18163 {
ifres18162 = True


} else {
ifres18162 = False


}

ifres18161 = ifres18162


} else {
ifres18161 = False


}

var ifres18160 Obj

if True == ifres18161 {
ifres18160 = True


} else {
ifres18160 = False


}

ifres18159 = ifres18160


} else {
ifres18159 = False


}

var ifres18158 Obj

if True == ifres18159 {
ifres18158 = True


} else {
ifres18158 = False


}

ifres18157 = ifres18158


} else {
ifres18157 = False


}

var ifres18156 Obj

if True == ifres18157 {
ifres18156 = True


} else {
ifres18156 = False


}

ifres18155 = ifres18156


} else {
ifres18155 = False


}

var ifres18154 Obj

if True == ifres18155 {
ifres18154 = True


} else {
ifres18154 = False


}

ifres18153 = ifres18154


} else {
ifres18153 = False


}

var ifres18152 Obj

if True == ifres18153 {
ifres18152 = True


} else {
ifres18152 = False


}

ifres18151 = ifres18152


} else {
ifres18151 = False


}

var ifres18150 Obj

if True == ifres18151 {
ifres18150 = True


} else {
ifres18150 = False


}

ifres18149 = ifres18150


} else {
ifres18149 = False


}

var ifres18148 Obj

if True == ifres18149 {
ifres18148 = True


} else {
ifres18148 = False


}

ifres18147 = ifres18148


} else {
ifres18147 = False


}

var ifres18146 Obj

if True == ifres18147 {
ifres18146 = True


} else {
ifres18146 = False


}

ifres18145 = ifres18146


} else {
ifres18145 = False


}

var ifres18144 Obj

if True == ifres18145 {
ifres18144 = True


} else {
ifres18144 = False


}

ifres18143 = ifres18144


} else {
ifres18143 = False


}

var ifres18142 Obj

if True == ifres18143 {
ifres18142 = True


} else {
ifres18142 = False


}

ifres18141 = ifres18142


} else {
ifres18141 = False


}

var ifres18140 Obj

if True == ifres18141 {
ifres18140 = True


} else {
ifres18140 = False


}

ifres18139 = ifres18140


} else {
ifres18139 = False


}

var ifres18138 Obj

if True == ifres18139 {
ifres18138 = True


} else {
ifres18138 = False


}

ifres18137 = ifres18138


} else {
ifres18137 = False


}

var ifres18136 Obj

if True == ifres18137 {
ifres18136 = True


} else {
ifres18136 = False


}

ifres18135 = ifres18136


} else {
ifres18135 = False


}

var ifres18134 Obj

if True == ifres18135 {
ifres18134 = True


} else {
ifres18134 = False


}

ifres18133 = ifres18134


} else {
ifres18133 = False


}

if True == ifres18133 {
tmp18125 := PrimTail(V364)

tmp18126 := PrimTail(tmp18125)

tmp18127 := PrimTail(tmp18126)

tmp18128 := PrimHead(tmp18127)

tmp18129 := PrimTail(tmp18128)

tmp18130 := PrimTail(tmp18129)

tmp18131 := PrimCons(V365, tmp18130)

__e.Return(PrimCons(symtype, tmp18131))
return


} else {
__e.Return(V365)
return
}


}, 2)

tmp18313 := Call(__e, ns2_1set, symshen_4use_1type_1info, tmp18124)


_ = tmp18313

tmp18314 := MakeNative(func(__e *ControlFlow) {
V368 := __e.Get(1)
_ = V368
tmp18324 := PrimIsVariable(V368)

if True == tmp18324 {
__e.Return(False)
return
} else {
tmp18322 := PrimIsPair(V368)

if True == tmp18322 {
tmp18319 := PrimHead(V368)

tmp18320 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18319)


if True == tmp18320 {
tmp18316 := PrimTail(V368)

tmp18317 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18316)


if True == tmp18317 {
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


} else {
__e.Return(True)
return
}


}


}, 1)

tmp18325 := Call(__e, ns2_1set, symshen_4monomorphic_2, tmp18314)


_ = tmp18325

tmp18326 := MakeNative(func(__e *ControlFlow) {
V369 := __e.Get(1)
_ = V369
tmp18352 := PrimIsPair(V369)

var ifres18334 Obj

if True == tmp18352 {
tmp18350 := PrimHead(V369)

tmp18351 := PrimEqual(symprotect, tmp18350)

var ifres18336 Obj

if True == tmp18351 {
tmp18348 := PrimTail(V369)

tmp18349 := PrimIsPair(tmp18348)

var ifres18338 Obj

if True == tmp18349 {
tmp18345 := PrimTail(V369)

tmp18346 := PrimTail(tmp18345)

tmp18347 := PrimEqual(Nil, tmp18346)

var ifres18340 Obj

if True == tmp18347 {
tmp18342 := PrimTail(V369)

tmp18343 := PrimHead(tmp18342)

tmp18344 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18343)


var ifres18341 Obj

if True == tmp18344 {
ifres18341 = True


} else {
ifres18341 = False


}

ifres18340 = ifres18341


} else {
ifres18340 = False


}

var ifres18339 Obj

if True == ifres18340 {
ifres18339 = True


} else {
ifres18339 = False


}

ifres18338 = ifres18339


} else {
ifres18338 = False


}

var ifres18337 Obj

if True == ifres18338 {
ifres18337 = True


} else {
ifres18337 = False


}

ifres18336 = ifres18337


} else {
ifres18336 = False


}

var ifres18335 Obj

if True == ifres18336 {
ifres18335 = True


} else {
ifres18335 = False


}

ifres18334 = ifres18335


} else {
ifres18334 = False


}

if True == ifres18334 {
tmp18327 := PrimTail(V369)

__e.Return(PrimHead(tmp18327))
return


} else {
tmp18332 := PrimIsPair(V369)

if True == tmp18332 {
tmp18328 := MakeNative(func(__e *ControlFlow) {
Z370 := __e.Get(1)
_ = Z370
__e.TailApply(PrimFunc(symshen_4process_1yacc_1semantics), Z370)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp18328, V369)
return


} else {
tmp18330 := Call(__e, PrimFunc(symshen_4non_1terminal_2), V369)


if True == tmp18330 {
__e.TailApply(PrimFunc(symconcat), symAction, V369)
return
} else {
__e.Return(V369)
return
}


}


}


}, 1)

tmp18353 := Call(__e, ns2_1set, symshen_4process_1yacc_1semantics, tmp18326)


_ = tmp18353

tmp18354 := MakeNative(func(__e *ControlFlow) {
V371 := __e.Get(1)
_ = V371
tmp18355 := PrimTail(V371)

__e.Return(PrimHead(tmp18355))
return


}, 1)

tmp18356 := Call(__e, ns2_1set, symshen_4_5_1out, tmp18354)


_ = tmp18356

tmp18357 := MakeNative(func(__e *ControlFlow) {
V372 := __e.Get(1)
_ = V372
__e.Return(PrimHead(V372))
return
}, 1)

tmp18358 := Call(__e, ns2_1set, symshen_4in_1_6, tmp18357)


_ = tmp18358

tmp18359 := MakeNative(func(__e *ControlFlow) {
V373 := __e.Get(1)
_ = V373
tmp18360 := PrimCons(V373, Nil)

__e.Return(PrimCons(Nil, tmp18360))
return


}, 1)

tmp18361 := Call(__e, ns2_1set, sym_5_b_6, tmp18359)


_ = tmp18361

tmp18362 := MakeNative(func(__e *ControlFlow) {
V374 := __e.Get(1)
_ = V374
tmp18363 := PrimCons(Nil, Nil)

__e.Return(PrimCons(V374, tmp18363))
return


}, 1)

tmp18364 := Call(__e, ns2_1set, sym_5e_6, tmp18362)


_ = tmp18364

tmp18365 := MakeNative(func(__e *ControlFlow) {
V377 := __e.Get(1)
_ = V377
tmp18368 := PrimEqual(Nil, V377)

if True == tmp18368 {
tmp18366 := PrimCons(Nil, Nil)

__e.Return(PrimCons(Nil, tmp18366))
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

__e.TailApply(ns2_1set, sym_5end_6, tmp18365)
return




}, 0)

