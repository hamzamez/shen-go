package main

import . "github.com/tiancaiamao/shen-go/kl"

var PrologMain = MakeNative(func(__e *ControlFlow) {
tmp9952 := MakeNative(func(__e *ControlFlow) {
V1212 := __e.Get(1)
_ = V1212
__e.TailApply(PrimFunc(symshen_4assert_d), V1212, symshen_4top)
return
}, 1)

tmp9953 := Call(__e, ns2_1set, symasserta, tmp9952)


_ = tmp9953

tmp9954 := MakeNative(func(__e *ControlFlow) {
V1213 := __e.Get(1)
_ = V1213
__e.TailApply(PrimFunc(symshen_4assert_d), V1213, symshen_4bottom)
return
}, 1)

tmp9955 := Call(__e, ns2_1set, symassertz, tmp9954)


_ = tmp9955

tmp9956 := MakeNative(func(__e *ControlFlow) {
V1214 := __e.Get(1)
_ = V1214
V1215 := __e.Get(2)
_ = V1215
tmp9990 := PrimIsPair(V1214)

var ifres9981 Obj

if True == tmp9990 {
tmp9988 := PrimTail(V1214)

tmp9989 := PrimIsPair(tmp9988)

var ifres9983 Obj

if True == tmp9989 {
tmp9985 := PrimTail(V1214)

tmp9986 := PrimHead(tmp9985)

tmp9987 := PrimEqual(sym_5_1_1, tmp9986)

var ifres9984 Obj

if True == tmp9987 {
ifres9984 = True


} else {
ifres9984 = False


}

ifres9983 = ifres9984


} else {
ifres9983 = False


}

var ifres9982 Obj

if True == ifres9983 {
ifres9982 = True


} else {
ifres9982 = False


}

ifres9981 = ifres9982


} else {
ifres9981 = False


}

if True == ifres9981 {
tmp9957 := MakeNative(func(__e *ControlFlow) {
W1216 := __e.Get(1)
_ = W1216
tmp9958 := MakeNative(func(__e *ControlFlow) {
W1217 := __e.Get(1)
_ = W1217
tmp9959 := MakeNative(func(__e *ControlFlow) {
W1218 := __e.Get(1)
_ = W1218
tmp9960 := MakeNative(func(__e *ControlFlow) {
W1219 := __e.Get(1)
_ = W1219
tmp9961 := MakeNative(func(__e *ControlFlow) {
W1220 := __e.Get(1)
_ = W1220
tmp9962 := MakeNative(func(__e *ControlFlow) {
W1221 := __e.Get(1)
_ = W1221
tmp9963 := MakeNative(func(__e *ControlFlow) {
W1222 := __e.Get(1)
_ = W1222
__e.Return(W1216)
return
}, 1)

tmp9964 := PrimTail(V1214)

tmp9965 := PrimTail(tmp9964)

tmp9966 := Call(__e, PrimFunc(symshen_4insert_1info), W1216, W1217, tmp9965, V1214, V1215)


__e.TailApply(tmp9963, tmp9966)
return


}, 1)

tmp9972 := PrimEqual(W1220, MakeNumber(-1))

var ifres9967 Obj

if True == tmp9972 {
tmp9968 := Call(__e, PrimFunc(symshen_4create_1skeleton), W1216, W1219)


tmp9969 := Call(__e, PrimFunc(symeval), tmp9968)


_ = tmp9969

tmp9970 := PrimValue(sym_dproperty_1vector_d)

tmp9971 := Call(__e, PrimFunc(symput), W1216, symshen_4dynamic, Nil, tmp9970)


ifres9967 = tmp9971


} else {
ifres9967 = symshen_4skip


}

__e.TailApply(tmp9962, ifres9967)
return


}, 1)

tmp9973 := Call(__e, PrimFunc(symarity), W1216)


__e.TailApply(tmp9961, tmp9973)
return


}, 1)

tmp9974 := Call(__e, PrimFunc(symshen_4parameters), W1218)


__e.TailApply(tmp9960, tmp9974)
return


}, 1)

tmp9975 := Call(__e, PrimFunc(symlength), W1217)


__e.TailApply(tmp9959, tmp9975)
return


}, 1)

tmp9976 := PrimHead(V1214)

tmp9977 := Call(__e, PrimFunc(symshen_4terms), tmp9976)


__e.TailApply(tmp9958, tmp9977)
return


}, 1)

tmp9978 := PrimHead(V1214)

tmp9979 := Call(__e, PrimFunc(symshen_4predicate), tmp9978)


__e.TailApply(tmp9957, tmp9979)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.assert*")))
return
}


}, 2)

tmp9991 := Call(__e, ns2_1set, symshen_4assert_d, tmp9956)


_ = tmp9991

tmp9992 := MakeNative(func(__e *ControlFlow) {
V1225 := __e.Get(1)
_ = V1225
tmp9994 := PrimIsPair(V1225)

if True == tmp9994 {
__e.Return(PrimHead(V1225))
return
} else {
__e.Return(V1225)
return
}


}, 1)

tmp9995 := Call(__e, ns2_1set, symshen_4predicate, tmp9992)


_ = tmp9995

tmp9996 := MakeNative(func(__e *ControlFlow) {
V1230 := __e.Get(1)
_ = V1230
tmp9998 := PrimIsPair(V1230)

if True == tmp9998 {
__e.Return(PrimTail(V1230))
return
} else {
__e.Return(Nil)
return
}


}, 1)

tmp9999 := Call(__e, ns2_1set, symshen_4terms, tmp9996)


_ = tmp9999

tmp10000 := MakeNative(func(__e *ControlFlow) {
V1231 := __e.Get(1)
_ = V1231
V1232 := __e.Get(2)
_ = V1232
tmp10001 := Call(__e, PrimFunc(symshen_4dynamic_1default), V1231, V1232)


tmp10002 := PrimCons(V1231, tmp10001)

__e.Return(PrimCons(symdefprolog, tmp10002))
return


}, 2)

tmp10003 := Call(__e, ns2_1set, symshen_4create_1skeleton, tmp10000)


_ = tmp10003

tmp10004 := MakeNative(func(__e *ControlFlow) {
V1233 := __e.Get(1)
_ = V1233
V1234 := __e.Get(2)
_ = V1234
tmp10005 := Call(__e, PrimFunc(symshen_4cons_1form), V1234)


tmp10006 := PrimCons(symshen_4dynamic, Nil)

tmp10007 := PrimCons(V1233, tmp10006)

tmp10008 := PrimCons(symget, tmp10007)

tmp10009 := PrimCons(tmp10008, Nil)

tmp10010 := PrimCons(tmp10005, tmp10009)

tmp10011 := PrimCons(symshen_4call_1dynamic, tmp10010)

tmp10012 := PrimIntern(MakeString(";"))

tmp10013 := PrimCons(tmp10012, Nil)

tmp10014 := PrimCons(tmp10011, tmp10013)

tmp10015 := PrimCons(sym_5_1_1, tmp10014)

__e.TailApply(PrimFunc(symappend), V1234, tmp10015)
return


}, 2)

tmp10016 := Call(__e, ns2_1set, symshen_4dynamic_1default, tmp10004)


_ = tmp10016

tmp10017 := MakeNative(func(__e *ControlFlow) {
V1235 := __e.Get(1)
_ = V1235
V1236 := __e.Get(2)
_ = V1236
V1237 := __e.Get(3)
_ = V1237
V1238 := __e.Get(4)
_ = V1238
V1239 := __e.Get(5)
_ = V1239
tmp10018 := MakeNative(func(__e *ControlFlow) {
W1240 := __e.Get(1)
_ = W1240
tmp10019 := MakeNative(func(__e *ControlFlow) {
W1241 := __e.Get(1)
_ = W1241
tmp10020 := MakeNative(func(__e *ControlFlow) {
W1242 := __e.Get(1)
_ = W1242
tmp10021 := MakeNative(func(__e *ControlFlow) {
W1243 := __e.Get(1)
_ = W1243
tmp10022 := MakeNative(func(__e *ControlFlow) {
W1244 := __e.Get(1)
_ = W1244
tmp10023 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V1235, symshen_4dynamic, W1244, tmp10023)
return


}, 1)

tmp10028 := PrimEqual(V1239, symshen_4top)

var ifres10024 Obj

if True == tmp10028 {
tmp10025 := PrimCons(W1242, W1243)

ifres10024 = tmp10025


} else {
tmp10026 := PrimCons(W1242, Nil)

tmp10027 := Call(__e, PrimFunc(symappend), W1243, tmp10026)


ifres10024 = tmp10027


}

__e.TailApply(tmp10022, ifres10024)
return


}, 1)

tmp10029 := PrimValue(sym_dproperty_1vector_d)

tmp10030 := Call(__e, PrimFunc(symget), V1235, symshen_4dynamic, tmp10029)


__e.TailApply(tmp10021, tmp10030)
return


}, 1)

tmp10031 := Call(__e, PrimFunc(symfn), W1240)


tmp10032 := PrimCons(W1240, V1238)

tmp10033 := PrimCons(tmp10031, tmp10032)

__e.TailApply(tmp10020, tmp10033)
return


}, 1)

tmp10034 := PrimCons(W1240, Nil)

tmp10035 := PrimCons(symdefprolog, tmp10034)

tmp10036 := PrimCons(sym_5_1_1, V1237)

tmp10037 := Call(__e, PrimFunc(symappend), V1236, tmp10036)


tmp10038 := Call(__e, PrimFunc(symappend), tmp10035, tmp10037)


tmp10039 := Call(__e, PrimFunc(symeval), tmp10038)


__e.TailApply(tmp10019, tmp10039)
return


}, 1)

tmp10040 := Call(__e, PrimFunc(symgensym), symshen_4g)


__e.TailApply(tmp10018, tmp10040)
return


}, 5)

tmp10041 := Call(__e, ns2_1set, symshen_4insert_1info, tmp10017)


_ = tmp10041

tmp10042 := MakeNative(func(__e *ControlFlow) {
tmp10043 := MakeNative(func(__e *ControlFlow) {
W1245 := __e.Get(1)
_ = W1245
tmp10044 := MakeNative(func(__e *ControlFlow) {
W1246 := __e.Get(1)
_ = W1246
__e.Return(W1246)
return
}, 1)

tmp10050 := Call(__e, PrimFunc(symempty_2), W1245)


var ifres10045 Obj

if True == tmp10050 {
tmp10046 := Call(__e, PrimFunc(symgensym), symshen_4g)


ifres10045 = tmp10046


} else {
tmp10047 := PrimTail(W1245)

tmp10048 := PrimSet(symshen_4_dnames_d, tmp10047)

_ = tmp10048

tmp10049 := PrimHead(W1245)

ifres10045 = tmp10049


}

__e.TailApply(tmp10044, ifres10045)
return


}, 1)

tmp10051 := PrimValue(symshen_4_dnames_d)

__e.TailApply(tmp10043, tmp10051)
return


}, 0)

tmp10052 := Call(__e, ns2_1set, symshen_4newname, tmp10042)


_ = tmp10052

tmp10053 := MakeNative(func(__e *ControlFlow) {
V1247 := __e.Get(1)
_ = V1247
V1248 := __e.Get(2)
_ = V1248
V1249 := __e.Get(3)
_ = V1249
V1250 := __e.Get(4)
_ = V1250
V1251 := __e.Get(5)
_ = V1251
V1252 := __e.Get(6)
_ = V1252
tmp10054 := MakeNative(func(__e *ControlFlow) {
W1253 := __e.Get(1)
_ = W1253
tmp10065 := PrimEqual(W1253, False)

if True == tmp10065 {
tmp10063 := Call(__e, PrimFunc(symshen_4unlocked_2), V1250)


if True == tmp10063 {
tmp10055 := MakeNative(func(__e *ControlFlow) {
W1257 := __e.Get(1)
_ = W1257
tmp10060 := PrimIsPair(W1257)

if True == tmp10060 {
tmp10056 := MakeNative(func(__e *ControlFlow) {
W1258 := __e.Get(1)
_ = W1258
tmp10057 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10057

__e.TailApply(PrimFunc(symshen_4call_1dynamic), V1247, W1258, V1249, V1250, V1251, V1252)
return


}, 1)

tmp10058 := PrimTail(W1257)

__e.TailApply(tmp10056, tmp10058)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp10061 := Call(__e, PrimFunc(symshen_4lazyderef), V1248, V1249)


__e.TailApply(tmp10055, tmp10061)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W1253)
return
}


}, 1)

tmp10080 := Call(__e, PrimFunc(symshen_4unlocked_2), V1250)


var ifres10066 Obj

if True == tmp10080 {
tmp10067 := MakeNative(func(__e *ControlFlow) {
W1254 := __e.Get(1)
_ = W1254
tmp10077 := PrimIsPair(W1254)

if True == tmp10077 {
tmp10068 := MakeNative(func(__e *ControlFlow) {
W1255 := __e.Get(1)
_ = W1255
tmp10073 := PrimIsPair(W1255)

if True == tmp10073 {
tmp10069 := MakeNative(func(__e *ControlFlow) {
W1256 := __e.Get(1)
_ = W1256
tmp10070 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10070

__e.TailApply(PrimFunc(symshen_4callrec), W1256, V1247, V1249, V1250, V1251, V1252)
return


}, 1)

tmp10071 := PrimHead(W1255)

__e.TailApply(tmp10069, tmp10071)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp10074 := PrimHead(W1254)

tmp10075 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10074, V1249)


__e.TailApply(tmp10068, tmp10075)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp10078 := Call(__e, PrimFunc(symshen_4lazyderef), V1248, V1249)


tmp10079 := Call(__e, tmp10067, tmp10078)


ifres10066 = tmp10079


} else {
ifres10066 = False


}

__e.TailApply(tmp10054, ifres10066)
return


}, 6)

tmp10081 := Call(__e, ns2_1set, symshen_4call_1dynamic, tmp10053)


_ = tmp10081

tmp10082 := MakeNative(func(__e *ControlFlow) {
V1259 := __e.Get(1)
_ = V1259
V1260 := __e.Get(2)
_ = V1260
V1261 := __e.Get(3)
_ = V1261
V1262 := __e.Get(4)
_ = V1262
V1263 := __e.Get(5)
_ = V1263
V1264 := __e.Get(6)
_ = V1264
tmp10092 := PrimEqual(Nil, V1260)

if True == tmp10092 {
tmp10083 := Call(__e, V1259, V1261)


tmp10084 := Call(__e, tmp10083, V1262)


tmp10085 := Call(__e, tmp10084, V1263)


__e.TailApply(tmp10085, V1264)
return


} else {
tmp10090 := PrimIsPair(V1260)

if True == tmp10090 {
tmp10086 := PrimHead(V1260)

tmp10087 := Call(__e, V1259, tmp10086)


tmp10088 := PrimTail(V1260)

__e.TailApply(PrimFunc(symshen_4callrec), tmp10087, tmp10088, V1261, V1262, V1263, V1264)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.callrec")))
return
}


}


}, 6)

tmp10093 := Call(__e, ns2_1set, symshen_4callrec, tmp10082)


_ = tmp10093

tmp10094 := MakeNative(func(__e *ControlFlow) {
V1265 := __e.Get(1)
_ = V1265
tmp10113 := PrimIsPair(V1265)

var ifres10104 Obj

if True == tmp10113 {
tmp10111 := PrimTail(V1265)

tmp10112 := PrimIsPair(tmp10111)

var ifres10106 Obj

if True == tmp10112 {
tmp10108 := PrimTail(V1265)

tmp10109 := PrimHead(tmp10108)

tmp10110 := PrimEqual(sym_5_1_1, tmp10109)

var ifres10107 Obj

if True == tmp10110 {
ifres10107 = True


} else {
ifres10107 = False


}

ifres10106 = ifres10107


} else {
ifres10106 = False


}

var ifres10105 Obj

if True == ifres10106 {
ifres10105 = True


} else {
ifres10105 = False


}

ifres10104 = ifres10105


} else {
ifres10104 = False


}

if True == ifres10104 {
tmp10095 := MakeNative(func(__e *ControlFlow) {
W1266 := __e.Get(1)
_ = W1266
tmp10096 := MakeNative(func(__e *ControlFlow) {
W1267 := __e.Get(1)
_ = W1267
tmp10097 := Call(__e, PrimFunc(symshen_4retract_1clause), V1265, W1267)


tmp10098 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), W1266, symshen_4dynamic, tmp10097, tmp10098)
return


}, 1)

tmp10099 := PrimValue(sym_dproperty_1vector_d)

tmp10100 := Call(__e, PrimFunc(symget), W1266, symshen_4dynamic, tmp10099)


__e.TailApply(tmp10096, tmp10100)
return


}, 1)

tmp10101 := PrimHead(V1265)

tmp10102 := Call(__e, PrimFunc(symshen_4predicate), tmp10101)


__e.TailApply(tmp10095, tmp10102)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function retract")))
return
}


}, 1)

tmp10114 := Call(__e, ns2_1set, symretract, tmp10094)


_ = tmp10114

tmp10115 := MakeNative(func(__e *ControlFlow) {
V1273 := __e.Get(1)
_ = V1273
V1274 := __e.Get(2)
_ = V1274
tmp10145 := PrimEqual(Nil, V1274)

if True == tmp10145 {
__e.Return(Nil)
return
} else {
tmp10143 := PrimIsPair(V1274)

var ifres10128 Obj

if True == tmp10143 {
tmp10141 := PrimHead(V1274)

tmp10142 := PrimIsPair(tmp10141)

var ifres10130 Obj

if True == tmp10142 {
tmp10138 := PrimHead(V1274)

tmp10139 := PrimTail(tmp10138)

tmp10140 := PrimIsPair(tmp10139)

var ifres10132 Obj

if True == tmp10140 {
tmp10134 := PrimHead(V1274)

tmp10135 := PrimTail(tmp10134)

tmp10136 := PrimTail(tmp10135)

tmp10137 := PrimEqual(V1273, tmp10136)

var ifres10133 Obj

if True == tmp10137 {
ifres10133 = True


} else {
ifres10133 = False


}

ifres10132 = ifres10133


} else {
ifres10132 = False


}

var ifres10131 Obj

if True == ifres10132 {
ifres10131 = True


} else {
ifres10131 = False


}

ifres10130 = ifres10131


} else {
ifres10130 = False


}

var ifres10129 Obj

if True == ifres10130 {
ifres10129 = True


} else {
ifres10129 = False


}

ifres10128 = ifres10129


} else {
ifres10128 = False


}

if True == ifres10128 {
tmp10116 := PrimHead(V1274)

tmp10117 := PrimTail(tmp10116)

tmp10118 := PrimHead(tmp10117)

tmp10119 := PrimValue(symshen_4_dnames_d)

tmp10120 := PrimCons(tmp10118, tmp10119)

tmp10121 := PrimSet(symshen_4_dnames_d, tmp10120)

_ = tmp10121

__e.Return(PrimTail(V1274))
return


} else {
tmp10126 := PrimIsPair(V1274)

if True == tmp10126 {
tmp10122 := PrimHead(V1274)

tmp10123 := PrimTail(V1274)

tmp10124 := Call(__e, PrimFunc(symshen_4retract_1clause), V1273, tmp10123)


__e.Return(PrimCons(tmp10122, tmp10124))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.retract-clause")))
return
}


}


}


}, 2)

tmp10146 := Call(__e, ns2_1set, symshen_4retract_1clause, tmp10115)


_ = tmp10146

tmp10147 := MakeNative(func(__e *ControlFlow) {
V1275 := __e.Get(1)
_ = V1275
V1276 := __e.Get(2)
_ = V1276
tmp10148 := MakeNative(func(__e *ControlFlow) {
Z1277 := __e.Get(1)
_ = Z1277
__e.TailApply(PrimFunc(symshen_4_5defprolog_6), Z1277)
return
}, 1)

tmp10149 := PrimCons(V1275, V1276)

__e.TailApply(PrimFunc(symcompile), tmp10148, tmp10149)
return


}, 2)

tmp10150 := Call(__e, ns2_1set, symshen_4compile_1prolog, tmp10147)


_ = tmp10150

tmp10151 := MakeNative(func(__e *ControlFlow) {
V1278 := __e.Get(1)
_ = V1278
tmp10152 := MakeNative(func(__e *ControlFlow) {
W1279 := __e.Get(1)
_ = W1279
tmp10154 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1279)


if True == tmp10154 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1279)
return
}


}, 1)

tmp10176 := PrimIsPair(V1278)

var ifres10155 Obj

if True == tmp10176 {
tmp10156 := MakeNative(func(__e *ControlFlow) {
W1280 := __e.Get(1)
_ = W1280
tmp10157 := MakeNative(func(__e *ControlFlow) {
W1281 := __e.Get(1)
_ = W1281
tmp10158 := MakeNative(func(__e *ControlFlow) {
W1282 := __e.Get(1)
_ = W1282
tmp10170 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1282)


if True == tmp10170 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10159 := MakeNative(func(__e *ControlFlow) {
W1283 := __e.Get(1)
_ = W1283
tmp10160 := MakeNative(func(__e *ControlFlow) {
W1284 := __e.Get(1)
_ = W1284
tmp10161 := MakeNative(func(__e *ControlFlow) {
W1285 := __e.Get(1)
_ = W1285
tmp10162 := MakeNative(func(__e *ControlFlow) {
W1286 := __e.Get(1)
_ = W1286
__e.TailApply(PrimFunc(symshen_4horn_1clause_1procedure), W1280, W1286)
return
}, 1)

tmp10163 := MakeNative(func(__e *ControlFlow) {
Z1287 := __e.Get(1)
_ = Z1287
__e.TailApply(PrimFunc(symshen_4linearise_1clause), Z1287)
return
}, 1)

tmp10164 := Call(__e, PrimFunc(symmap), tmp10163, W1283)


__e.TailApply(tmp10162, tmp10164)
return


}, 1)

tmp10165 := Call(__e, PrimFunc(symshen_4prolog_1arity_1check), W1280, W1283)


tmp10166 := Call(__e, tmp10161, tmp10165)


__e.TailApply(PrimFunc(symshen_4comb), W1284, tmp10166)
return


}, 1)

tmp10167 := Call(__e, PrimFunc(symshen_4in_1_6), W1282)


__e.TailApply(tmp10160, tmp10167)
return


}, 1)

tmp10168 := Call(__e, PrimFunc(symshen_4_5_1out), W1282)


__e.TailApply(tmp10159, tmp10168)
return


}


}, 1)

tmp10171 := Call(__e, PrimFunc(symshen_4_5clauses_6), W1281)


__e.TailApply(tmp10158, tmp10171)
return


}, 1)

tmp10172 := Call(__e, PrimFunc(symtail), V1278)


__e.TailApply(tmp10157, tmp10172)
return


}, 1)

tmp10173 := Call(__e, PrimFunc(symhead), V1278)


tmp10174 := Call(__e, tmp10156, tmp10173)


ifres10155 = tmp10174


} else {
tmp10175 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10155 = tmp10175


}

__e.TailApply(tmp10152, ifres10155)
return


}, 1)

tmp10177 := Call(__e, ns2_1set, symshen_4_5defprolog_6, tmp10151)


_ = tmp10177

tmp10178 := MakeNative(func(__e *ControlFlow) {
V1290 := __e.Get(1)
_ = V1290
V1291 := __e.Get(2)
_ = V1291
tmp10222 := PrimIsPair(V1291)

var ifres10203 Obj

if True == tmp10222 {
tmp10220 := PrimHead(V1291)

tmp10221 := PrimIsPair(tmp10220)

var ifres10205 Obj

if True == tmp10221 {
tmp10217 := PrimHead(V1291)

tmp10218 := PrimTail(tmp10217)

tmp10219 := PrimIsPair(tmp10218)

var ifres10207 Obj

if True == tmp10219 {
tmp10213 := PrimHead(V1291)

tmp10214 := PrimTail(tmp10213)

tmp10215 := PrimTail(tmp10214)

tmp10216 := PrimEqual(Nil, tmp10215)

var ifres10209 Obj

if True == tmp10216 {
tmp10211 := PrimTail(V1291)

tmp10212 := PrimEqual(Nil, tmp10211)

var ifres10210 Obj

if True == tmp10212 {
ifres10210 = True


} else {
ifres10210 = False


}

ifres10209 = ifres10210


} else {
ifres10209 = False


}

var ifres10208 Obj

if True == ifres10209 {
ifres10208 = True


} else {
ifres10208 = False


}

ifres10207 = ifres10208


} else {
ifres10207 = False


}

var ifres10206 Obj

if True == ifres10207 {
ifres10206 = True


} else {
ifres10206 = False


}

ifres10205 = ifres10206


} else {
ifres10205 = False


}

var ifres10204 Obj

if True == ifres10205 {
ifres10204 = True


} else {
ifres10204 = False


}

ifres10203 = ifres10204


} else {
ifres10203 = False


}

if True == ifres10203 {
tmp10179 := PrimHead(V1291)

tmp10180 := PrimHead(tmp10179)

__e.TailApply(PrimFunc(symlength), tmp10180)
return


} else {
tmp10201 := PrimIsPair(V1291)

var ifres10186 Obj

if True == tmp10201 {
tmp10199 := PrimHead(V1291)

tmp10200 := PrimIsPair(tmp10199)

var ifres10188 Obj

if True == tmp10200 {
tmp10196 := PrimHead(V1291)

tmp10197 := PrimTail(tmp10196)

tmp10198 := PrimIsPair(tmp10197)

var ifres10190 Obj

if True == tmp10198 {
tmp10192 := PrimHead(V1291)

tmp10193 := PrimTail(tmp10192)

tmp10194 := PrimTail(tmp10193)

tmp10195 := PrimEqual(Nil, tmp10194)

var ifres10191 Obj

if True == tmp10195 {
ifres10191 = True


} else {
ifres10191 = False


}

ifres10190 = ifres10191


} else {
ifres10190 = False


}

var ifres10189 Obj

if True == ifres10190 {
ifres10189 = True


} else {
ifres10189 = False


}

ifres10188 = ifres10189


} else {
ifres10188 = False


}

var ifres10187 Obj

if True == ifres10188 {
ifres10187 = True


} else {
ifres10187 = False


}

ifres10186 = ifres10187


} else {
ifres10186 = False


}

if True == ifres10186 {
tmp10181 := PrimHead(V1291)

tmp10182 := PrimHead(tmp10181)

tmp10183 := Call(__e, PrimFunc(symlength), tmp10182)


tmp10184 := PrimTail(V1291)

__e.TailApply(PrimFunc(symshen_4pac_1h), V1290, tmp10183, tmp10184)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.prolog-arity-check")))
return
}


}


}, 2)

tmp10223 := Call(__e, ns2_1set, symshen_4prolog_1arity_1check, tmp10178)


_ = tmp10223

tmp10224 := MakeNative(func(__e *ControlFlow) {
V1296 := __e.Get(1)
_ = V1296
V1297 := __e.Get(2)
_ = V1297
V1298 := __e.Get(3)
_ = V1298
tmp10240 := PrimEqual(Nil, V1298)

if True == tmp10240 {
__e.Return(V1297)
return
} else {
tmp10238 := PrimIsPair(V1298)

var ifres10234 Obj

if True == tmp10238 {
tmp10236 := PrimHead(V1298)

tmp10237 := PrimIsPair(tmp10236)

var ifres10235 Obj

if True == tmp10237 {
ifres10235 = True


} else {
ifres10235 = False


}

ifres10234 = ifres10235


} else {
ifres10234 = False


}

if True == ifres10234 {
tmp10229 := PrimHead(V1298)

tmp10230 := PrimHead(tmp10229)

tmp10231 := Call(__e, PrimFunc(symlength), tmp10230)


tmp10232 := PrimEqual(V1297, tmp10231)

if True == tmp10232 {
tmp10225 := PrimTail(V1298)

__e.TailApply(PrimFunc(symshen_4pac_1h), V1296, V1297, tmp10225)
return


} else {
tmp10226 := Call(__e, PrimFunc(symshen_4app), V1296, MakeString("\n"), symshen_4a)


tmp10227 := PrimStringConcat(MakeString("arity error in prolog procedure "), tmp10226)

__e.Return(PrimSimpleError(tmp10227))
return


}


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.pac-h")))
return
}


}


}, 3)

tmp10241 := Call(__e, ns2_1set, symshen_4pac_1h, tmp10224)


_ = tmp10241

tmp10242 := MakeNative(func(__e *ControlFlow) {
V1299 := __e.Get(1)
_ = V1299
tmp10243 := MakeNative(func(__e *ControlFlow) {
W1300 := __e.Get(1)
_ = W1300
tmp10262 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1300)


if True == tmp10262 {
tmp10244 := MakeNative(func(__e *ControlFlow) {
W1307 := __e.Get(1)
_ = W1307
tmp10246 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1307)


if True == tmp10246 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1307)
return
}


}, 1)

tmp10247 := MakeNative(func(__e *ControlFlow) {
W1308 := __e.Get(1)
_ = W1308
tmp10258 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1308)


if True == tmp10258 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10248 := MakeNative(func(__e *ControlFlow) {
W1309 := __e.Get(1)
_ = W1309
tmp10249 := MakeNative(func(__e *ControlFlow) {
W1310 := __e.Get(1)
_ = W1310
tmp10254 := Call(__e, PrimFunc(symempty_2), W1309)


var ifres10250 Obj

if True == tmp10254 {
ifres10250 = Nil


} else {
tmp10251 := Call(__e, PrimFunc(symshen_4app), W1309, MakeString("\n ..."), symshen_4r)


tmp10252 := PrimStringConcat(MakeString("Prolog syntax error here:\n "), tmp10251)

tmp10253 := PrimSimpleError(tmp10252)

ifres10250 = tmp10253


}

__e.TailApply(PrimFunc(symshen_4comb), W1310, ifres10250)
return


}, 1)

tmp10255 := Call(__e, PrimFunc(symshen_4in_1_6), W1308)


__e.TailApply(tmp10249, tmp10255)
return


}, 1)

tmp10256 := Call(__e, PrimFunc(symshen_4_5_1out), W1308)


__e.TailApply(tmp10248, tmp10256)
return


}


}, 1)

tmp10259 := Call(__e, PrimFunc(sym_5_b_6), V1299)


tmp10260 := Call(__e, tmp10247, tmp10259)


__e.TailApply(tmp10244, tmp10260)
return


} else {
__e.Return(W1300)
return
}


}, 1)

tmp10263 := MakeNative(func(__e *ControlFlow) {
W1301 := __e.Get(1)
_ = W1301
tmp10278 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1301)


if True == tmp10278 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10264 := MakeNative(func(__e *ControlFlow) {
W1302 := __e.Get(1)
_ = W1302
tmp10265 := MakeNative(func(__e *ControlFlow) {
W1303 := __e.Get(1)
_ = W1303
tmp10266 := MakeNative(func(__e *ControlFlow) {
W1304 := __e.Get(1)
_ = W1304
tmp10273 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1304)


if True == tmp10273 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10267 := MakeNative(func(__e *ControlFlow) {
W1305 := __e.Get(1)
_ = W1305
tmp10268 := MakeNative(func(__e *ControlFlow) {
W1306 := __e.Get(1)
_ = W1306
tmp10269 := PrimCons(W1302, W1305)

__e.TailApply(PrimFunc(symshen_4comb), W1306, tmp10269)
return


}, 1)

tmp10270 := Call(__e, PrimFunc(symshen_4in_1_6), W1304)


__e.TailApply(tmp10268, tmp10270)
return


}, 1)

tmp10271 := Call(__e, PrimFunc(symshen_4_5_1out), W1304)


__e.TailApply(tmp10267, tmp10271)
return


}


}, 1)

tmp10274 := Call(__e, PrimFunc(symshen_4_5clauses_6), W1303)


__e.TailApply(tmp10266, tmp10274)
return


}, 1)

tmp10275 := Call(__e, PrimFunc(symshen_4in_1_6), W1301)


__e.TailApply(tmp10265, tmp10275)
return


}, 1)

tmp10276 := Call(__e, PrimFunc(symshen_4_5_1out), W1301)


__e.TailApply(tmp10264, tmp10276)
return


}


}, 1)

tmp10279 := Call(__e, PrimFunc(symshen_4_5clause_6), V1299)


tmp10280 := Call(__e, tmp10263, tmp10279)


__e.TailApply(tmp10243, tmp10280)
return


}, 1)

tmp10281 := Call(__e, ns2_1set, symshen_4_5clauses_6, tmp10242)


_ = tmp10281

tmp10282 := MakeNative(func(__e *ControlFlow) {
V1311 := __e.Get(1)
_ = V1311
tmp10298 := PrimIsPair(V1311)

var ifres10289 Obj

if True == tmp10298 {
tmp10296 := PrimTail(V1311)

tmp10297 := PrimIsPair(tmp10296)

var ifres10291 Obj

if True == tmp10297 {
tmp10293 := PrimTail(V1311)

tmp10294 := PrimTail(tmp10293)

tmp10295 := PrimEqual(Nil, tmp10294)

var ifres10292 Obj

if True == tmp10295 {
ifres10292 = True


} else {
ifres10292 = False


}

ifres10291 = ifres10292


} else {
ifres10291 = False


}

var ifres10290 Obj

if True == ifres10291 {
ifres10290 = True


} else {
ifres10290 = False


}

ifres10289 = ifres10290


} else {
ifres10289 = False


}

if True == ifres10289 {
tmp10283 := PrimHead(V1311)

tmp10284 := PrimTail(V1311)

tmp10285 := PrimHead(tmp10284)

tmp10286 := Call(__e, PrimFunc(sym_8p), tmp10283, tmp10285)


tmp10287 := Call(__e, PrimFunc(symshen_4linearise), tmp10286)


__e.TailApply(PrimFunc(symshen_4lch), tmp10287)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.linearise-clause")))
return
}


}, 1)

tmp10299 := Call(__e, ns2_1set, symshen_4linearise_1clause, tmp10282)


_ = tmp10299

tmp10300 := MakeNative(func(__e *ControlFlow) {
V1312 := __e.Get(1)
_ = V1312
tmp10306 := Call(__e, PrimFunc(symtuple_2), V1312)


if True == tmp10306 {
tmp10301 := Call(__e, PrimFunc(symfst), V1312)


tmp10302 := Call(__e, PrimFunc(symsnd), V1312)


tmp10303 := Call(__e, PrimFunc(symshen_4lchh), tmp10302)


tmp10304 := PrimCons(tmp10303, Nil)

__e.Return(PrimCons(tmp10301, tmp10304))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.lch")))
return
}


}, 1)

tmp10307 := Call(__e, ns2_1set, symshen_4lch, tmp10300)


_ = tmp10307

tmp10308 := MakeNative(func(__e *ControlFlow) {
V1313 := __e.Get(1)
_ = V1313
tmp10371 := PrimIsPair(V1313)

var ifres10320 Obj

if True == tmp10371 {
tmp10369 := PrimHead(V1313)

tmp10370 := PrimEqual(symwhere, tmp10369)

var ifres10322 Obj

if True == tmp10370 {
tmp10367 := PrimTail(V1313)

tmp10368 := PrimIsPair(tmp10367)

var ifres10324 Obj

if True == tmp10368 {
tmp10364 := PrimTail(V1313)

tmp10365 := PrimHead(tmp10364)

tmp10366 := PrimIsPair(tmp10365)

var ifres10326 Obj

if True == tmp10366 {
tmp10360 := PrimTail(V1313)

tmp10361 := PrimHead(tmp10360)

tmp10362 := PrimHead(tmp10361)

tmp10363 := PrimEqual(sym_a, tmp10362)

var ifres10328 Obj

if True == tmp10363 {
tmp10356 := PrimTail(V1313)

tmp10357 := PrimHead(tmp10356)

tmp10358 := PrimTail(tmp10357)

tmp10359 := PrimIsPair(tmp10358)

var ifres10330 Obj

if True == tmp10359 {
tmp10351 := PrimTail(V1313)

tmp10352 := PrimHead(tmp10351)

tmp10353 := PrimTail(tmp10352)

tmp10354 := PrimTail(tmp10353)

tmp10355 := PrimIsPair(tmp10354)

var ifres10332 Obj

if True == tmp10355 {
tmp10345 := PrimTail(V1313)

tmp10346 := PrimHead(tmp10345)

tmp10347 := PrimTail(tmp10346)

tmp10348 := PrimTail(tmp10347)

tmp10349 := PrimTail(tmp10348)

tmp10350 := PrimEqual(Nil, tmp10349)

var ifres10334 Obj

if True == tmp10350 {
tmp10342 := PrimTail(V1313)

tmp10343 := PrimTail(tmp10342)

tmp10344 := PrimIsPair(tmp10343)

var ifres10336 Obj

if True == tmp10344 {
tmp10338 := PrimTail(V1313)

tmp10339 := PrimTail(tmp10338)

tmp10340 := PrimTail(tmp10339)

tmp10341 := PrimEqual(Nil, tmp10340)

var ifres10337 Obj

if True == tmp10341 {
ifres10337 = True


} else {
ifres10337 = False


}

ifres10336 = ifres10337


} else {
ifres10336 = False


}

var ifres10335 Obj

if True == ifres10336 {
ifres10335 = True


} else {
ifres10335 = False


}

ifres10334 = ifres10335


} else {
ifres10334 = False


}

var ifres10333 Obj

if True == ifres10334 {
ifres10333 = True


} else {
ifres10333 = False


}

ifres10332 = ifres10333


} else {
ifres10332 = False


}

var ifres10331 Obj

if True == ifres10332 {
ifres10331 = True


} else {
ifres10331 = False


}

ifres10330 = ifres10331


} else {
ifres10330 = False


}

var ifres10329 Obj

if True == ifres10330 {
ifres10329 = True


} else {
ifres10329 = False


}

ifres10328 = ifres10329


} else {
ifres10328 = False


}

var ifres10327 Obj

if True == ifres10328 {
ifres10327 = True


} else {
ifres10327 = False


}

ifres10326 = ifres10327


} else {
ifres10326 = False


}

var ifres10325 Obj

if True == ifres10326 {
ifres10325 = True


} else {
ifres10325 = False


}

ifres10324 = ifres10325


} else {
ifres10324 = False


}

var ifres10323 Obj

if True == ifres10324 {
ifres10323 = True


} else {
ifres10323 = False


}

ifres10322 = ifres10323


} else {
ifres10322 = False


}

var ifres10321 Obj

if True == ifres10322 {
ifres10321 = True


} else {
ifres10321 = False


}

ifres10320 = ifres10321


} else {
ifres10320 = False


}

if True == ifres10320 {
tmp10310 := PrimValue(symshen_4_doccurs_d)

var ifres10309 Obj

if True == tmp10310 {
ifres10309 = symis_b


} else {
ifres10309 = symis


}

tmp10311 := PrimTail(V1313)

tmp10312 := PrimHead(tmp10311)

tmp10313 := PrimTail(tmp10312)

tmp10314 := PrimCons(ifres10309, tmp10313)

tmp10315 := PrimTail(V1313)

tmp10316 := PrimTail(tmp10315)

tmp10317 := PrimHead(tmp10316)

tmp10318 := Call(__e, PrimFunc(symshen_4lchh), tmp10317)


__e.Return(PrimCons(tmp10314, tmp10318))
return


} else {
__e.Return(V1313)
return
}


}, 1)

tmp10372 := Call(__e, ns2_1set, symshen_4lchh, tmp10308)


_ = tmp10372

tmp10373 := MakeNative(func(__e *ControlFlow) {
V1314 := __e.Get(1)
_ = V1314
tmp10374 := MakeNative(func(__e *ControlFlow) {
W1315 := __e.Get(1)
_ = W1315
tmp10376 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1315)


if True == tmp10376 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1315)
return
}


}, 1)

tmp10377 := MakeNative(func(__e *ControlFlow) {
W1316 := __e.Get(1)
_ = W1316
tmp10403 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1316)


if True == tmp10403 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10378 := MakeNative(func(__e *ControlFlow) {
W1317 := __e.Get(1)
_ = W1317
tmp10379 := MakeNative(func(__e *ControlFlow) {
W1318 := __e.Get(1)
_ = W1318
tmp10399 := Call(__e, PrimFunc(symshen_4hds_a_2), W1318, sym_5_1_1)


if True == tmp10399 {
tmp10380 := MakeNative(func(__e *ControlFlow) {
W1319 := __e.Get(1)
_ = W1319
tmp10381 := MakeNative(func(__e *ControlFlow) {
W1320 := __e.Get(1)
_ = W1320
tmp10395 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1320)


if True == tmp10395 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10382 := MakeNative(func(__e *ControlFlow) {
W1321 := __e.Get(1)
_ = W1321
tmp10383 := MakeNative(func(__e *ControlFlow) {
W1322 := __e.Get(1)
_ = W1322
tmp10384 := MakeNative(func(__e *ControlFlow) {
W1323 := __e.Get(1)
_ = W1323
tmp10390 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1323)


if True == tmp10390 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10385 := MakeNative(func(__e *ControlFlow) {
W1324 := __e.Get(1)
_ = W1324
tmp10386 := PrimCons(W1321, Nil)

tmp10387 := PrimCons(W1317, tmp10386)

__e.TailApply(PrimFunc(symshen_4comb), W1324, tmp10387)
return


}, 1)

tmp10388 := Call(__e, PrimFunc(symshen_4in_1_6), W1323)


__e.TailApply(tmp10385, tmp10388)
return


}


}, 1)

tmp10391 := Call(__e, PrimFunc(symshen_4_5sc_6), W1322)


__e.TailApply(tmp10384, tmp10391)
return


}, 1)

tmp10392 := Call(__e, PrimFunc(symshen_4in_1_6), W1320)


__e.TailApply(tmp10383, tmp10392)
return


}, 1)

tmp10393 := Call(__e, PrimFunc(symshen_4_5_1out), W1320)


__e.TailApply(tmp10382, tmp10393)
return


}


}, 1)

tmp10396 := Call(__e, PrimFunc(symshen_4_5body_6), W1319)


__e.TailApply(tmp10381, tmp10396)
return


}, 1)

tmp10397 := Call(__e, PrimFunc(symtail), W1318)


__e.TailApply(tmp10380, tmp10397)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10400 := Call(__e, PrimFunc(symshen_4in_1_6), W1316)


__e.TailApply(tmp10379, tmp10400)
return


}, 1)

tmp10401 := Call(__e, PrimFunc(symshen_4_5_1out), W1316)


__e.TailApply(tmp10378, tmp10401)
return


}


}, 1)

tmp10404 := Call(__e, PrimFunc(symshen_4_5head_6), V1314)


tmp10405 := Call(__e, tmp10377, tmp10404)


__e.TailApply(tmp10374, tmp10405)
return


}, 1)

tmp10406 := Call(__e, ns2_1set, symshen_4_5clause_6, tmp10373)


_ = tmp10406

tmp10407 := MakeNative(func(__e *ControlFlow) {
V1325 := __e.Get(1)
_ = V1325
tmp10408 := MakeNative(func(__e *ControlFlow) {
W1326 := __e.Get(1)
_ = W1326
tmp10420 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1326)


if True == tmp10420 {
tmp10409 := MakeNative(func(__e *ControlFlow) {
W1333 := __e.Get(1)
_ = W1333
tmp10411 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1333)


if True == tmp10411 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1333)
return
}


}, 1)

tmp10412 := MakeNative(func(__e *ControlFlow) {
W1334 := __e.Get(1)
_ = W1334
tmp10416 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1334)


if True == tmp10416 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10413 := MakeNative(func(__e *ControlFlow) {
W1335 := __e.Get(1)
_ = W1335
__e.TailApply(PrimFunc(symshen_4comb), W1335, Nil)
return
}, 1)

tmp10414 := Call(__e, PrimFunc(symshen_4in_1_6), W1334)


__e.TailApply(tmp10413, tmp10414)
return


}


}, 1)

tmp10417 := Call(__e, PrimFunc(sym_5e_6), V1325)


tmp10418 := Call(__e, tmp10412, tmp10417)


__e.TailApply(tmp10409, tmp10418)
return


} else {
__e.Return(W1326)
return
}


}, 1)

tmp10421 := MakeNative(func(__e *ControlFlow) {
W1327 := __e.Get(1)
_ = W1327
tmp10436 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1327)


if True == tmp10436 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10422 := MakeNative(func(__e *ControlFlow) {
W1328 := __e.Get(1)
_ = W1328
tmp10423 := MakeNative(func(__e *ControlFlow) {
W1329 := __e.Get(1)
_ = W1329
tmp10424 := MakeNative(func(__e *ControlFlow) {
W1330 := __e.Get(1)
_ = W1330
tmp10431 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1330)


if True == tmp10431 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10425 := MakeNative(func(__e *ControlFlow) {
W1331 := __e.Get(1)
_ = W1331
tmp10426 := MakeNative(func(__e *ControlFlow) {
W1332 := __e.Get(1)
_ = W1332
tmp10427 := PrimCons(W1328, W1331)

__e.TailApply(PrimFunc(symshen_4comb), W1332, tmp10427)
return


}, 1)

tmp10428 := Call(__e, PrimFunc(symshen_4in_1_6), W1330)


__e.TailApply(tmp10426, tmp10428)
return


}, 1)

tmp10429 := Call(__e, PrimFunc(symshen_4_5_1out), W1330)


__e.TailApply(tmp10425, tmp10429)
return


}


}, 1)

tmp10432 := Call(__e, PrimFunc(symshen_4_5head_6), W1329)


__e.TailApply(tmp10424, tmp10432)
return


}, 1)

tmp10433 := Call(__e, PrimFunc(symshen_4in_1_6), W1327)


__e.TailApply(tmp10423, tmp10433)
return


}, 1)

tmp10434 := Call(__e, PrimFunc(symshen_4_5_1out), W1327)


__e.TailApply(tmp10422, tmp10434)
return


}


}, 1)

tmp10437 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1325)


tmp10438 := Call(__e, tmp10421, tmp10437)


__e.TailApply(tmp10408, tmp10438)
return


}, 1)

tmp10439 := Call(__e, ns2_1set, symshen_4_5head_6, tmp10407)


_ = tmp10439

tmp10440 := MakeNative(func(__e *ControlFlow) {
V1336 := __e.Get(1)
_ = V1336
tmp10441 := MakeNative(func(__e *ControlFlow) {
W1337 := __e.Get(1)
_ = W1337
tmp10629 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1337)


if True == tmp10629 {
tmp10442 := MakeNative(func(__e *ControlFlow) {
W1340 := __e.Get(1)
_ = W1340
tmp10616 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1340)


if True == tmp10616 {
tmp10443 := MakeNative(func(__e *ControlFlow) {
W1343 := __e.Get(1)
_ = W1343
tmp10577 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1343)


if True == tmp10577 {
tmp10444 := MakeNative(func(__e *ControlFlow) {
W1355 := __e.Get(1)
_ = W1355
tmp10547 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1355)


if True == tmp10547 {
tmp10445 := MakeNative(func(__e *ControlFlow) {
W1364 := __e.Get(1)
_ = W1364
tmp10517 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1364)


if True == tmp10517 {
tmp10446 := MakeNative(func(__e *ControlFlow) {
W1373 := __e.Get(1)
_ = W1373
tmp10483 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1373)


if True == tmp10483 {
tmp10447 := MakeNative(func(__e *ControlFlow) {
W1383 := __e.Get(1)
_ = W1383
tmp10449 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1383)


if True == tmp10449 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1383)
return
}


}, 1)

tmp10481 := Call(__e, PrimFunc(symshen_4ccons_2), V1336)


var ifres10450 Obj

if True == tmp10481 {
tmp10451 := MakeNative(func(__e *ControlFlow) {
W1384 := __e.Get(1)
_ = W1384
tmp10452 := MakeNative(func(__e *ControlFlow) {
W1385 := __e.Get(1)
_ = W1385
tmp10476 := Call(__e, PrimFunc(symshen_4hds_a_2), W1384, symmode)


if True == tmp10476 {
tmp10453 := MakeNative(func(__e *ControlFlow) {
W1386 := __e.Get(1)
_ = W1386
tmp10454 := MakeNative(func(__e *ControlFlow) {
W1387 := __e.Get(1)
_ = W1387
tmp10472 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1387)


if True == tmp10472 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10455 := MakeNative(func(__e *ControlFlow) {
W1388 := __e.Get(1)
_ = W1388
tmp10456 := MakeNative(func(__e *ControlFlow) {
W1389 := __e.Get(1)
_ = W1389
tmp10468 := Call(__e, PrimFunc(symshen_4hds_a_2), W1389, sym_1)


if True == tmp10468 {
tmp10457 := MakeNative(func(__e *ControlFlow) {
W1390 := __e.Get(1)
_ = W1390
tmp10458 := MakeNative(func(__e *ControlFlow) {
W1391 := __e.Get(1)
_ = W1391
tmp10464 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1391)


if True == tmp10464 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10459 := MakeNative(func(__e *ControlFlow) {
W1392 := __e.Get(1)
_ = W1392
tmp10460 := PrimCons(W1388, Nil)

tmp10461 := PrimCons(symshen_4_1m, tmp10460)

__e.TailApply(PrimFunc(symshen_4comb), W1385, tmp10461)
return


}, 1)

tmp10462 := Call(__e, PrimFunc(symshen_4in_1_6), W1391)


__e.TailApply(tmp10459, tmp10462)
return


}


}, 1)

tmp10465 := Call(__e, PrimFunc(sym_5end_6), W1390)


__e.TailApply(tmp10458, tmp10465)
return


}, 1)

tmp10466 := Call(__e, PrimFunc(symtail), W1389)


__e.TailApply(tmp10457, tmp10466)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10469 := Call(__e, PrimFunc(symshen_4in_1_6), W1387)


__e.TailApply(tmp10456, tmp10469)
return


}, 1)

tmp10470 := Call(__e, PrimFunc(symshen_4_5_1out), W1387)


__e.TailApply(tmp10455, tmp10470)
return


}


}, 1)

tmp10473 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1386)


__e.TailApply(tmp10454, tmp10473)
return


}, 1)

tmp10474 := Call(__e, PrimFunc(symtail), W1384)


__e.TailApply(tmp10453, tmp10474)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10477 := Call(__e, PrimFunc(symtail), V1336)


__e.TailApply(tmp10452, tmp10477)
return


}, 1)

tmp10478 := Call(__e, PrimFunc(symhead), V1336)


tmp10479 := Call(__e, tmp10451, tmp10478)


ifres10450 = tmp10479


} else {
tmp10480 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10450 = tmp10480


}

__e.TailApply(tmp10447, ifres10450)
return


} else {
__e.Return(W1373)
return
}


}, 1)

tmp10515 := Call(__e, PrimFunc(symshen_4ccons_2), V1336)


var ifres10484 Obj

if True == tmp10515 {
tmp10485 := MakeNative(func(__e *ControlFlow) {
W1374 := __e.Get(1)
_ = W1374
tmp10486 := MakeNative(func(__e *ControlFlow) {
W1375 := __e.Get(1)
_ = W1375
tmp10510 := Call(__e, PrimFunc(symshen_4hds_a_2), W1374, symmode)


if True == tmp10510 {
tmp10487 := MakeNative(func(__e *ControlFlow) {
W1376 := __e.Get(1)
_ = W1376
tmp10488 := MakeNative(func(__e *ControlFlow) {
W1377 := __e.Get(1)
_ = W1377
tmp10506 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1377)


if True == tmp10506 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10489 := MakeNative(func(__e *ControlFlow) {
W1378 := __e.Get(1)
_ = W1378
tmp10490 := MakeNative(func(__e *ControlFlow) {
W1379 := __e.Get(1)
_ = W1379
tmp10502 := Call(__e, PrimFunc(symshen_4hds_a_2), W1379, sym_7)


if True == tmp10502 {
tmp10491 := MakeNative(func(__e *ControlFlow) {
W1380 := __e.Get(1)
_ = W1380
tmp10492 := MakeNative(func(__e *ControlFlow) {
W1381 := __e.Get(1)
_ = W1381
tmp10498 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1381)


if True == tmp10498 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10493 := MakeNative(func(__e *ControlFlow) {
W1382 := __e.Get(1)
_ = W1382
tmp10494 := PrimCons(W1378, Nil)

tmp10495 := PrimCons(symshen_4_7m, tmp10494)

__e.TailApply(PrimFunc(symshen_4comb), W1375, tmp10495)
return


}, 1)

tmp10496 := Call(__e, PrimFunc(symshen_4in_1_6), W1381)


__e.TailApply(tmp10493, tmp10496)
return


}


}, 1)

tmp10499 := Call(__e, PrimFunc(sym_5end_6), W1380)


__e.TailApply(tmp10492, tmp10499)
return


}, 1)

tmp10500 := Call(__e, PrimFunc(symtail), W1379)


__e.TailApply(tmp10491, tmp10500)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10503 := Call(__e, PrimFunc(symshen_4in_1_6), W1377)


__e.TailApply(tmp10490, tmp10503)
return


}, 1)

tmp10504 := Call(__e, PrimFunc(symshen_4_5_1out), W1377)


__e.TailApply(tmp10489, tmp10504)
return


}


}, 1)

tmp10507 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1376)


__e.TailApply(tmp10488, tmp10507)
return


}, 1)

tmp10508 := Call(__e, PrimFunc(symtail), W1374)


__e.TailApply(tmp10487, tmp10508)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10511 := Call(__e, PrimFunc(symtail), V1336)


__e.TailApply(tmp10486, tmp10511)
return


}, 1)

tmp10512 := Call(__e, PrimFunc(symhead), V1336)


tmp10513 := Call(__e, tmp10485, tmp10512)


ifres10484 = tmp10513


} else {
tmp10514 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10484 = tmp10514


}

__e.TailApply(tmp10446, ifres10484)
return


} else {
__e.Return(W1364)
return
}


}, 1)

tmp10545 := Call(__e, PrimFunc(symshen_4ccons_2), V1336)


var ifres10518 Obj

if True == tmp10545 {
tmp10519 := MakeNative(func(__e *ControlFlow) {
W1365 := __e.Get(1)
_ = W1365
tmp10520 := MakeNative(func(__e *ControlFlow) {
W1366 := __e.Get(1)
_ = W1366
tmp10540 := Call(__e, PrimFunc(symshen_4hds_a_2), W1365, sym_1)


if True == tmp10540 {
tmp10521 := MakeNative(func(__e *ControlFlow) {
W1367 := __e.Get(1)
_ = W1367
tmp10522 := MakeNative(func(__e *ControlFlow) {
W1368 := __e.Get(1)
_ = W1368
tmp10536 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1368)


if True == tmp10536 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10523 := MakeNative(func(__e *ControlFlow) {
W1369 := __e.Get(1)
_ = W1369
tmp10524 := MakeNative(func(__e *ControlFlow) {
W1370 := __e.Get(1)
_ = W1370
tmp10525 := MakeNative(func(__e *ControlFlow) {
W1371 := __e.Get(1)
_ = W1371
tmp10531 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1371)


if True == tmp10531 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10526 := MakeNative(func(__e *ControlFlow) {
W1372 := __e.Get(1)
_ = W1372
tmp10527 := PrimCons(W1369, Nil)

tmp10528 := PrimCons(symshen_4_1m, tmp10527)

__e.TailApply(PrimFunc(symshen_4comb), W1366, tmp10528)
return


}, 1)

tmp10529 := Call(__e, PrimFunc(symshen_4in_1_6), W1371)


__e.TailApply(tmp10526, tmp10529)
return


}


}, 1)

tmp10532 := Call(__e, PrimFunc(sym_5end_6), W1370)


__e.TailApply(tmp10525, tmp10532)
return


}, 1)

tmp10533 := Call(__e, PrimFunc(symshen_4in_1_6), W1368)


__e.TailApply(tmp10524, tmp10533)
return


}, 1)

tmp10534 := Call(__e, PrimFunc(symshen_4_5_1out), W1368)


__e.TailApply(tmp10523, tmp10534)
return


}


}, 1)

tmp10537 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1367)


__e.TailApply(tmp10522, tmp10537)
return


}, 1)

tmp10538 := Call(__e, PrimFunc(symtail), W1365)


__e.TailApply(tmp10521, tmp10538)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10541 := Call(__e, PrimFunc(symtail), V1336)


__e.TailApply(tmp10520, tmp10541)
return


}, 1)

tmp10542 := Call(__e, PrimFunc(symhead), V1336)


tmp10543 := Call(__e, tmp10519, tmp10542)


ifres10518 = tmp10543


} else {
tmp10544 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10518 = tmp10544


}

__e.TailApply(tmp10445, ifres10518)
return


} else {
__e.Return(W1355)
return
}


}, 1)

tmp10575 := Call(__e, PrimFunc(symshen_4ccons_2), V1336)


var ifres10548 Obj

if True == tmp10575 {
tmp10549 := MakeNative(func(__e *ControlFlow) {
W1356 := __e.Get(1)
_ = W1356
tmp10550 := MakeNative(func(__e *ControlFlow) {
W1357 := __e.Get(1)
_ = W1357
tmp10570 := Call(__e, PrimFunc(symshen_4hds_a_2), W1356, sym_7)


if True == tmp10570 {
tmp10551 := MakeNative(func(__e *ControlFlow) {
W1358 := __e.Get(1)
_ = W1358
tmp10552 := MakeNative(func(__e *ControlFlow) {
W1359 := __e.Get(1)
_ = W1359
tmp10566 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1359)


if True == tmp10566 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10553 := MakeNative(func(__e *ControlFlow) {
W1360 := __e.Get(1)
_ = W1360
tmp10554 := MakeNative(func(__e *ControlFlow) {
W1361 := __e.Get(1)
_ = W1361
tmp10555 := MakeNative(func(__e *ControlFlow) {
W1362 := __e.Get(1)
_ = W1362
tmp10561 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1362)


if True == tmp10561 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10556 := MakeNative(func(__e *ControlFlow) {
W1363 := __e.Get(1)
_ = W1363
tmp10557 := PrimCons(W1360, Nil)

tmp10558 := PrimCons(symshen_4_7m, tmp10557)

__e.TailApply(PrimFunc(symshen_4comb), W1357, tmp10558)
return


}, 1)

tmp10559 := Call(__e, PrimFunc(symshen_4in_1_6), W1362)


__e.TailApply(tmp10556, tmp10559)
return


}


}, 1)

tmp10562 := Call(__e, PrimFunc(sym_5end_6), W1361)


__e.TailApply(tmp10555, tmp10562)
return


}, 1)

tmp10563 := Call(__e, PrimFunc(symshen_4in_1_6), W1359)


__e.TailApply(tmp10554, tmp10563)
return


}, 1)

tmp10564 := Call(__e, PrimFunc(symshen_4_5_1out), W1359)


__e.TailApply(tmp10553, tmp10564)
return


}


}, 1)

tmp10567 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1358)


__e.TailApply(tmp10552, tmp10567)
return


}, 1)

tmp10568 := Call(__e, PrimFunc(symtail), W1356)


__e.TailApply(tmp10551, tmp10568)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10571 := Call(__e, PrimFunc(symtail), V1336)


__e.TailApply(tmp10550, tmp10571)
return


}, 1)

tmp10572 := Call(__e, PrimFunc(symhead), V1336)


tmp10573 := Call(__e, tmp10549, tmp10572)


ifres10548 = tmp10573


} else {
tmp10574 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10548 = tmp10574


}

__e.TailApply(tmp10444, ifres10548)
return


} else {
__e.Return(W1343)
return
}


}, 1)

tmp10614 := Call(__e, PrimFunc(symshen_4ccons_2), V1336)


var ifres10578 Obj

if True == tmp10614 {
tmp10579 := MakeNative(func(__e *ControlFlow) {
W1344 := __e.Get(1)
_ = W1344
tmp10580 := MakeNative(func(__e *ControlFlow) {
W1345 := __e.Get(1)
_ = W1345
tmp10609 := Call(__e, PrimFunc(symshen_4hds_a_2), W1344, symcons)


if True == tmp10609 {
tmp10581 := MakeNative(func(__e *ControlFlow) {
W1346 := __e.Get(1)
_ = W1346
tmp10582 := MakeNative(func(__e *ControlFlow) {
W1347 := __e.Get(1)
_ = W1347
tmp10605 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1347)


if True == tmp10605 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10583 := MakeNative(func(__e *ControlFlow) {
W1348 := __e.Get(1)
_ = W1348
tmp10584 := MakeNative(func(__e *ControlFlow) {
W1349 := __e.Get(1)
_ = W1349
tmp10585 := MakeNative(func(__e *ControlFlow) {
W1350 := __e.Get(1)
_ = W1350
tmp10600 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1350)


if True == tmp10600 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10586 := MakeNative(func(__e *ControlFlow) {
W1351 := __e.Get(1)
_ = W1351
tmp10587 := MakeNative(func(__e *ControlFlow) {
W1352 := __e.Get(1)
_ = W1352
tmp10588 := MakeNative(func(__e *ControlFlow) {
W1353 := __e.Get(1)
_ = W1353
tmp10595 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1353)


if True == tmp10595 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10589 := MakeNative(func(__e *ControlFlow) {
W1354 := __e.Get(1)
_ = W1354
tmp10590 := PrimCons(W1351, Nil)

tmp10591 := PrimCons(W1348, tmp10590)

tmp10592 := PrimCons(symcons, tmp10591)

__e.TailApply(PrimFunc(symshen_4comb), W1345, tmp10592)
return


}, 1)

tmp10593 := Call(__e, PrimFunc(symshen_4in_1_6), W1353)


__e.TailApply(tmp10589, tmp10593)
return


}


}, 1)

tmp10596 := Call(__e, PrimFunc(sym_5end_6), W1352)


__e.TailApply(tmp10588, tmp10596)
return


}, 1)

tmp10597 := Call(__e, PrimFunc(symshen_4in_1_6), W1350)


__e.TailApply(tmp10587, tmp10597)
return


}, 1)

tmp10598 := Call(__e, PrimFunc(symshen_4_5_1out), W1350)


__e.TailApply(tmp10586, tmp10598)
return


}


}, 1)

tmp10601 := Call(__e, PrimFunc(symshen_4_5hterm2_6), W1349)


__e.TailApply(tmp10585, tmp10601)
return


}, 1)

tmp10602 := Call(__e, PrimFunc(symshen_4in_1_6), W1347)


__e.TailApply(tmp10584, tmp10602)
return


}, 1)

tmp10603 := Call(__e, PrimFunc(symshen_4_5_1out), W1347)


__e.TailApply(tmp10583, tmp10603)
return


}


}, 1)

tmp10606 := Call(__e, PrimFunc(symshen_4_5hterm1_6), W1346)


__e.TailApply(tmp10582, tmp10606)
return


}, 1)

tmp10607 := Call(__e, PrimFunc(symtail), W1344)


__e.TailApply(tmp10581, tmp10607)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10610 := Call(__e, PrimFunc(symtail), V1336)


__e.TailApply(tmp10580, tmp10610)
return


}, 1)

tmp10611 := Call(__e, PrimFunc(symhead), V1336)


tmp10612 := Call(__e, tmp10579, tmp10611)


ifres10578 = tmp10612


} else {
tmp10613 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10578 = tmp10613


}

__e.TailApply(tmp10443, ifres10578)
return


} else {
__e.Return(W1340)
return
}


}, 1)

tmp10627 := PrimIsPair(V1336)

var ifres10617 Obj

if True == tmp10627 {
tmp10618 := MakeNative(func(__e *ControlFlow) {
W1341 := __e.Get(1)
_ = W1341
tmp10619 := MakeNative(func(__e *ControlFlow) {
W1342 := __e.Get(1)
_ = W1342
tmp10621 := PrimIntern(MakeString(":"))

tmp10622 := PrimEqual(W1341, tmp10621)

if True == tmp10622 {
__e.TailApply(PrimFunc(symshen_4comb), W1342, W1341)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10623 := Call(__e, PrimFunc(symtail), V1336)


__e.TailApply(tmp10619, tmp10623)
return


}, 1)

tmp10624 := Call(__e, PrimFunc(symhead), V1336)


tmp10625 := Call(__e, tmp10618, tmp10624)


ifres10617 = tmp10625


} else {
tmp10626 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10617 = tmp10626


}

__e.TailApply(tmp10442, ifres10617)
return


} else {
__e.Return(W1337)
return
}


}, 1)

tmp10643 := PrimIsPair(V1336)

var ifres10630 Obj

if True == tmp10643 {
tmp10631 := MakeNative(func(__e *ControlFlow) {
W1338 := __e.Get(1)
_ = W1338
tmp10632 := MakeNative(func(__e *ControlFlow) {
W1339 := __e.Get(1)
_ = W1339
tmp10638 := Call(__e, PrimFunc(symatom_2), W1338)


var ifres10634 Obj

if True == tmp10638 {
tmp10636 := Call(__e, PrimFunc(symshen_4prolog_1keyword_2), W1338)


tmp10637 := PrimNot(tmp10636)

var ifres10635 Obj

if True == tmp10637 {
ifres10635 = True


} else {
ifres10635 = False


}

ifres10634 = ifres10635


} else {
ifres10634 = False


}

if True == ifres10634 {
__e.TailApply(PrimFunc(symshen_4comb), W1339, W1338)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10639 := Call(__e, PrimFunc(symtail), V1336)


__e.TailApply(tmp10632, tmp10639)
return


}, 1)

tmp10640 := Call(__e, PrimFunc(symhead), V1336)


tmp10641 := Call(__e, tmp10631, tmp10640)


ifres10630 = tmp10641


} else {
tmp10642 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10630 = tmp10642


}

__e.TailApply(tmp10441, ifres10630)
return


}, 1)

tmp10644 := Call(__e, ns2_1set, symshen_4_5hterm_6, tmp10440)


_ = tmp10644

tmp10645 := MakeNative(func(__e *ControlFlow) {
V1393 := __e.Get(1)
_ = V1393
tmp10646 := PrimIntern(MakeString(";"))

tmp10647 := PrimCons(sym_5_1_1, Nil)

tmp10648 := PrimCons(tmp10646, tmp10647)

__e.TailApply(PrimFunc(symelement_2), V1393, tmp10648)
return


}, 1)

tmp10649 := Call(__e, ns2_1set, symshen_4prolog_1keyword_2, tmp10645)


_ = tmp10649

tmp10650 := MakeNative(func(__e *ControlFlow) {
V1394 := __e.Get(1)
_ = V1394
tmp10663 := PrimIsSymbol(V1394)

if True == tmp10663 {
__e.Return(True)
return
} else {
tmp10661 := PrimIsString(V1394)

var ifres10652 Obj

if True == tmp10661 {
ifres10652 = True


} else {
tmp10660 := Call(__e, PrimFunc(symboolean_2), V1394)


var ifres10654 Obj

if True == tmp10660 {
ifres10654 = True


} else {
tmp10659 := PrimIsNumber(V1394)

var ifres10656 Obj

if True == tmp10659 {
ifres10656 = True


} else {
tmp10658 := Call(__e, PrimFunc(symempty_2), V1394)


var ifres10657 Obj

if True == tmp10658 {
ifres10657 = True


} else {
ifres10657 = False


}

ifres10656 = ifres10657


}

var ifres10655 Obj

if True == ifres10656 {
ifres10655 = True


} else {
ifres10655 = False


}

ifres10654 = ifres10655


}

var ifres10653 Obj

if True == ifres10654 {
ifres10653 = True


} else {
ifres10653 = False


}

ifres10652 = ifres10653


}

if True == ifres10652 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp10664 := Call(__e, ns2_1set, symatom_2, tmp10650)


_ = tmp10664

tmp10665 := MakeNative(func(__e *ControlFlow) {
V1395 := __e.Get(1)
_ = V1395
tmp10666 := MakeNative(func(__e *ControlFlow) {
W1396 := __e.Get(1)
_ = W1396
tmp10668 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1396)


if True == tmp10668 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1396)
return
}


}, 1)

tmp10669 := MakeNative(func(__e *ControlFlow) {
W1397 := __e.Get(1)
_ = W1397
tmp10675 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1397)


if True == tmp10675 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10670 := MakeNative(func(__e *ControlFlow) {
W1398 := __e.Get(1)
_ = W1398
tmp10671 := MakeNative(func(__e *ControlFlow) {
W1399 := __e.Get(1)
_ = W1399
__e.TailApply(PrimFunc(symshen_4comb), W1399, W1398)
return
}, 1)

tmp10672 := Call(__e, PrimFunc(symshen_4in_1_6), W1397)


__e.TailApply(tmp10671, tmp10672)
return


}, 1)

tmp10673 := Call(__e, PrimFunc(symshen_4_5_1out), W1397)


__e.TailApply(tmp10670, tmp10673)
return


}


}, 1)

tmp10676 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1395)


tmp10677 := Call(__e, tmp10669, tmp10676)


__e.TailApply(tmp10666, tmp10677)
return


}, 1)

tmp10678 := Call(__e, ns2_1set, symshen_4_5hterm1_6, tmp10665)


_ = tmp10678

tmp10679 := MakeNative(func(__e *ControlFlow) {
V1400 := __e.Get(1)
_ = V1400
tmp10680 := MakeNative(func(__e *ControlFlow) {
W1401 := __e.Get(1)
_ = W1401
tmp10682 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1401)


if True == tmp10682 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1401)
return
}


}, 1)

tmp10683 := MakeNative(func(__e *ControlFlow) {
W1402 := __e.Get(1)
_ = W1402
tmp10689 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1402)


if True == tmp10689 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10684 := MakeNative(func(__e *ControlFlow) {
W1403 := __e.Get(1)
_ = W1403
tmp10685 := MakeNative(func(__e *ControlFlow) {
W1404 := __e.Get(1)
_ = W1404
__e.TailApply(PrimFunc(symshen_4comb), W1404, W1403)
return
}, 1)

tmp10686 := Call(__e, PrimFunc(symshen_4in_1_6), W1402)


__e.TailApply(tmp10685, tmp10686)
return


}, 1)

tmp10687 := Call(__e, PrimFunc(symshen_4_5_1out), W1402)


__e.TailApply(tmp10684, tmp10687)
return


}


}, 1)

tmp10690 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1400)


tmp10691 := Call(__e, tmp10683, tmp10690)


__e.TailApply(tmp10680, tmp10691)
return


}, 1)

tmp10692 := Call(__e, ns2_1set, symshen_4_5hterm2_6, tmp10679)


_ = tmp10692

tmp10693 := MakeNative(func(__e *ControlFlow) {
V1405 := __e.Get(1)
_ = V1405
tmp10694 := MakeNative(func(__e *ControlFlow) {
W1406 := __e.Get(1)
_ = W1406
tmp10706 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1406)


if True == tmp10706 {
tmp10695 := MakeNative(func(__e *ControlFlow) {
W1413 := __e.Get(1)
_ = W1413
tmp10697 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1413)


if True == tmp10697 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1413)
return
}


}, 1)

tmp10698 := MakeNative(func(__e *ControlFlow) {
W1414 := __e.Get(1)
_ = W1414
tmp10702 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1414)


if True == tmp10702 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10699 := MakeNative(func(__e *ControlFlow) {
W1415 := __e.Get(1)
_ = W1415
__e.TailApply(PrimFunc(symshen_4comb), W1415, Nil)
return
}, 1)

tmp10700 := Call(__e, PrimFunc(symshen_4in_1_6), W1414)


__e.TailApply(tmp10699, tmp10700)
return


}


}, 1)

tmp10703 := Call(__e, PrimFunc(sym_5e_6), V1405)


tmp10704 := Call(__e, tmp10698, tmp10703)


__e.TailApply(tmp10695, tmp10704)
return


} else {
__e.Return(W1406)
return
}


}, 1)

tmp10707 := MakeNative(func(__e *ControlFlow) {
W1407 := __e.Get(1)
_ = W1407
tmp10722 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1407)


if True == tmp10722 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10708 := MakeNative(func(__e *ControlFlow) {
W1408 := __e.Get(1)
_ = W1408
tmp10709 := MakeNative(func(__e *ControlFlow) {
W1409 := __e.Get(1)
_ = W1409
tmp10710 := MakeNative(func(__e *ControlFlow) {
W1410 := __e.Get(1)
_ = W1410
tmp10717 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1410)


if True == tmp10717 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10711 := MakeNative(func(__e *ControlFlow) {
W1411 := __e.Get(1)
_ = W1411
tmp10712 := MakeNative(func(__e *ControlFlow) {
W1412 := __e.Get(1)
_ = W1412
tmp10713 := PrimCons(W1408, W1411)

__e.TailApply(PrimFunc(symshen_4comb), W1412, tmp10713)
return


}, 1)

tmp10714 := Call(__e, PrimFunc(symshen_4in_1_6), W1410)


__e.TailApply(tmp10712, tmp10714)
return


}, 1)

tmp10715 := Call(__e, PrimFunc(symshen_4_5_1out), W1410)


__e.TailApply(tmp10711, tmp10715)
return


}


}, 1)

tmp10718 := Call(__e, PrimFunc(symshen_4_5body_6), W1409)


__e.TailApply(tmp10710, tmp10718)
return


}, 1)

tmp10719 := Call(__e, PrimFunc(symshen_4in_1_6), W1407)


__e.TailApply(tmp10709, tmp10719)
return


}, 1)

tmp10720 := Call(__e, PrimFunc(symshen_4_5_1out), W1407)


__e.TailApply(tmp10708, tmp10720)
return


}


}, 1)

tmp10723 := Call(__e, PrimFunc(symshen_4_5literal_6), V1405)


tmp10724 := Call(__e, tmp10707, tmp10723)


__e.TailApply(tmp10694, tmp10724)
return


}, 1)

tmp10725 := Call(__e, ns2_1set, symshen_4_5body_6, tmp10693)


_ = tmp10725

tmp10726 := MakeNative(func(__e *ControlFlow) {
V1416 := __e.Get(1)
_ = V1416
tmp10727 := MakeNative(func(__e *ControlFlow) {
W1417 := __e.Get(1)
_ = W1417
tmp10754 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1417)


if True == tmp10754 {
tmp10728 := MakeNative(func(__e *ControlFlow) {
W1419 := __e.Get(1)
_ = W1419
tmp10730 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1419)


if True == tmp10730 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1419)
return
}


}, 1)

tmp10752 := Call(__e, PrimFunc(symshen_4ccons_2), V1416)


var ifres10731 Obj

if True == tmp10752 {
tmp10732 := MakeNative(func(__e *ControlFlow) {
W1420 := __e.Get(1)
_ = W1420
tmp10733 := MakeNative(func(__e *ControlFlow) {
W1421 := __e.Get(1)
_ = W1421
tmp10734 := MakeNative(func(__e *ControlFlow) {
W1422 := __e.Get(1)
_ = W1422
tmp10746 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1422)


if True == tmp10746 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10735 := MakeNative(func(__e *ControlFlow) {
W1423 := __e.Get(1)
_ = W1423
tmp10736 := MakeNative(func(__e *ControlFlow) {
W1424 := __e.Get(1)
_ = W1424
tmp10737 := MakeNative(func(__e *ControlFlow) {
W1425 := __e.Get(1)
_ = W1425
tmp10741 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1425)


if True == tmp10741 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10738 := MakeNative(func(__e *ControlFlow) {
W1426 := __e.Get(1)
_ = W1426
__e.TailApply(PrimFunc(symshen_4comb), W1421, W1423)
return
}, 1)

tmp10739 := Call(__e, PrimFunc(symshen_4in_1_6), W1425)


__e.TailApply(tmp10738, tmp10739)
return


}


}, 1)

tmp10742 := Call(__e, PrimFunc(sym_5end_6), W1424)


__e.TailApply(tmp10737, tmp10742)
return


}, 1)

tmp10743 := Call(__e, PrimFunc(symshen_4in_1_6), W1422)


__e.TailApply(tmp10736, tmp10743)
return


}, 1)

tmp10744 := Call(__e, PrimFunc(symshen_4_5_1out), W1422)


__e.TailApply(tmp10735, tmp10744)
return


}


}, 1)

tmp10747 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1420)


__e.TailApply(tmp10734, tmp10747)
return


}, 1)

tmp10748 := Call(__e, PrimFunc(symtail), V1416)


__e.TailApply(tmp10733, tmp10748)
return


}, 1)

tmp10749 := Call(__e, PrimFunc(symhead), V1416)


tmp10750 := Call(__e, tmp10732, tmp10749)


ifres10731 = tmp10750


} else {
tmp10751 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10731 = tmp10751


}

__e.TailApply(tmp10728, ifres10731)
return


} else {
__e.Return(W1417)
return
}


}, 1)

tmp10760 := Call(__e, PrimFunc(symshen_4hds_a_2), V1416, sym_b)


var ifres10755 Obj

if True == tmp10760 {
tmp10756 := MakeNative(func(__e *ControlFlow) {
W1418 := __e.Get(1)
_ = W1418
__e.TailApply(PrimFunc(symshen_4comb), W1418, sym_b)
return
}, 1)

tmp10757 := Call(__e, PrimFunc(symtail), V1416)


tmp10758 := Call(__e, tmp10756, tmp10757)


ifres10755 = tmp10758


} else {
tmp10759 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10755 = tmp10759


}

__e.TailApply(tmp10727, ifres10755)
return


}, 1)

tmp10761 := Call(__e, ns2_1set, symshen_4_5literal_6, tmp10726)


_ = tmp10761

tmp10762 := MakeNative(func(__e *ControlFlow) {
V1427 := __e.Get(1)
_ = V1427
tmp10763 := MakeNative(func(__e *ControlFlow) {
W1428 := __e.Get(1)
_ = W1428
tmp10775 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1428)


if True == tmp10775 {
tmp10764 := MakeNative(func(__e *ControlFlow) {
W1435 := __e.Get(1)
_ = W1435
tmp10766 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1435)


if True == tmp10766 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1435)
return
}


}, 1)

tmp10767 := MakeNative(func(__e *ControlFlow) {
W1436 := __e.Get(1)
_ = W1436
tmp10771 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1436)


if True == tmp10771 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10768 := MakeNative(func(__e *ControlFlow) {
W1437 := __e.Get(1)
_ = W1437
__e.TailApply(PrimFunc(symshen_4comb), W1437, Nil)
return
}, 1)

tmp10769 := Call(__e, PrimFunc(symshen_4in_1_6), W1436)


__e.TailApply(tmp10768, tmp10769)
return


}


}, 1)

tmp10772 := Call(__e, PrimFunc(sym_5e_6), V1427)


tmp10773 := Call(__e, tmp10767, tmp10772)


__e.TailApply(tmp10764, tmp10773)
return


} else {
__e.Return(W1428)
return
}


}, 1)

tmp10776 := MakeNative(func(__e *ControlFlow) {
W1429 := __e.Get(1)
_ = W1429
tmp10791 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1429)


if True == tmp10791 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10777 := MakeNative(func(__e *ControlFlow) {
W1430 := __e.Get(1)
_ = W1430
tmp10778 := MakeNative(func(__e *ControlFlow) {
W1431 := __e.Get(1)
_ = W1431
tmp10779 := MakeNative(func(__e *ControlFlow) {
W1432 := __e.Get(1)
_ = W1432
tmp10786 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1432)


if True == tmp10786 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10780 := MakeNative(func(__e *ControlFlow) {
W1433 := __e.Get(1)
_ = W1433
tmp10781 := MakeNative(func(__e *ControlFlow) {
W1434 := __e.Get(1)
_ = W1434
tmp10782 := PrimCons(W1430, W1433)

__e.TailApply(PrimFunc(symshen_4comb), W1434, tmp10782)
return


}, 1)

tmp10783 := Call(__e, PrimFunc(symshen_4in_1_6), W1432)


__e.TailApply(tmp10781, tmp10783)
return


}, 1)

tmp10784 := Call(__e, PrimFunc(symshen_4_5_1out), W1432)


__e.TailApply(tmp10780, tmp10784)
return


}


}, 1)

tmp10787 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1431)


__e.TailApply(tmp10779, tmp10787)
return


}, 1)

tmp10788 := Call(__e, PrimFunc(symshen_4in_1_6), W1429)


__e.TailApply(tmp10778, tmp10788)
return


}, 1)

tmp10789 := Call(__e, PrimFunc(symshen_4_5_1out), W1429)


__e.TailApply(tmp10777, tmp10789)
return


}


}, 1)

tmp10792 := Call(__e, PrimFunc(symshen_4_5bterm_6), V1427)


tmp10793 := Call(__e, tmp10776, tmp10792)


__e.TailApply(tmp10763, tmp10793)
return


}, 1)

tmp10794 := Call(__e, ns2_1set, symshen_4_5bterms_6, tmp10762)


_ = tmp10794

tmp10795 := MakeNative(func(__e *ControlFlow) {
V1438 := __e.Get(1)
_ = V1438
tmp10796 := MakeNative(func(__e *ControlFlow) {
W1439 := __e.Get(1)
_ = W1439
tmp10836 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1439)


if True == tmp10836 {
tmp10797 := MakeNative(func(__e *ControlFlow) {
W1443 := __e.Get(1)
_ = W1443
tmp10824 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1443)


if True == tmp10824 {
tmp10798 := MakeNative(func(__e *ControlFlow) {
W1446 := __e.Get(1)
_ = W1446
tmp10800 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1446)


if True == tmp10800 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1446)
return
}


}, 1)

tmp10822 := Call(__e, PrimFunc(symshen_4ccons_2), V1438)


var ifres10801 Obj

if True == tmp10822 {
tmp10802 := MakeNative(func(__e *ControlFlow) {
W1447 := __e.Get(1)
_ = W1447
tmp10803 := MakeNative(func(__e *ControlFlow) {
W1448 := __e.Get(1)
_ = W1448
tmp10804 := MakeNative(func(__e *ControlFlow) {
W1449 := __e.Get(1)
_ = W1449
tmp10816 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1449)


if True == tmp10816 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10805 := MakeNative(func(__e *ControlFlow) {
W1450 := __e.Get(1)
_ = W1450
tmp10806 := MakeNative(func(__e *ControlFlow) {
W1451 := __e.Get(1)
_ = W1451
tmp10807 := MakeNative(func(__e *ControlFlow) {
W1452 := __e.Get(1)
_ = W1452
tmp10811 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1452)


if True == tmp10811 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10808 := MakeNative(func(__e *ControlFlow) {
W1453 := __e.Get(1)
_ = W1453
__e.TailApply(PrimFunc(symshen_4comb), W1448, W1450)
return
}, 1)

tmp10809 := Call(__e, PrimFunc(symshen_4in_1_6), W1452)


__e.TailApply(tmp10808, tmp10809)
return


}


}, 1)

tmp10812 := Call(__e, PrimFunc(sym_5end_6), W1451)


__e.TailApply(tmp10807, tmp10812)
return


}, 1)

tmp10813 := Call(__e, PrimFunc(symshen_4in_1_6), W1449)


__e.TailApply(tmp10806, tmp10813)
return


}, 1)

tmp10814 := Call(__e, PrimFunc(symshen_4_5_1out), W1449)


__e.TailApply(tmp10805, tmp10814)
return


}


}, 1)

tmp10817 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1447)


__e.TailApply(tmp10804, tmp10817)
return


}, 1)

tmp10818 := Call(__e, PrimFunc(symtail), V1438)


__e.TailApply(tmp10803, tmp10818)
return


}, 1)

tmp10819 := Call(__e, PrimFunc(symhead), V1438)


tmp10820 := Call(__e, tmp10802, tmp10819)


ifres10801 = tmp10820


} else {
tmp10821 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10801 = tmp10821


}

__e.TailApply(tmp10798, ifres10801)
return


} else {
__e.Return(W1443)
return
}


}, 1)

tmp10834 := PrimIsPair(V1438)

var ifres10825 Obj

if True == tmp10834 {
tmp10826 := MakeNative(func(__e *ControlFlow) {
W1444 := __e.Get(1)
_ = W1444
tmp10827 := MakeNative(func(__e *ControlFlow) {
W1445 := __e.Get(1)
_ = W1445
tmp10829 := Call(__e, PrimFunc(symatom_2), W1444)


if True == tmp10829 {
__e.TailApply(PrimFunc(symshen_4comb), W1445, W1444)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10830 := Call(__e, PrimFunc(symtail), V1438)


__e.TailApply(tmp10827, tmp10830)
return


}, 1)

tmp10831 := Call(__e, PrimFunc(symhead), V1438)


tmp10832 := Call(__e, tmp10826, tmp10831)


ifres10825 = tmp10832


} else {
tmp10833 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10825 = tmp10833


}

__e.TailApply(tmp10797, ifres10825)
return


} else {
__e.Return(W1439)
return
}


}, 1)

tmp10837 := MakeNative(func(__e *ControlFlow) {
W1440 := __e.Get(1)
_ = W1440
tmp10843 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1440)


if True == tmp10843 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10838 := MakeNative(func(__e *ControlFlow) {
W1441 := __e.Get(1)
_ = W1441
tmp10839 := MakeNative(func(__e *ControlFlow) {
W1442 := __e.Get(1)
_ = W1442
__e.TailApply(PrimFunc(symshen_4comb), W1442, W1441)
return
}, 1)

tmp10840 := Call(__e, PrimFunc(symshen_4in_1_6), W1440)


__e.TailApply(tmp10839, tmp10840)
return


}, 1)

tmp10841 := Call(__e, PrimFunc(symshen_4_5_1out), W1440)


__e.TailApply(tmp10838, tmp10841)
return


}


}, 1)

tmp10844 := Call(__e, PrimFunc(symshen_4_5wildcard_6), V1438)


tmp10845 := Call(__e, tmp10837, tmp10844)


__e.TailApply(tmp10796, tmp10845)
return


}, 1)

tmp10846 := Call(__e, ns2_1set, symshen_4_5bterm_6, tmp10795)


_ = tmp10846

tmp10847 := MakeNative(func(__e *ControlFlow) {
V1454 := __e.Get(1)
_ = V1454
tmp10848 := MakeNative(func(__e *ControlFlow) {
W1455 := __e.Get(1)
_ = W1455
tmp10850 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1455)


if True == tmp10850 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1455)
return
}


}, 1)

tmp10861 := PrimIsPair(V1454)

var ifres10851 Obj

if True == tmp10861 {
tmp10852 := MakeNative(func(__e *ControlFlow) {
W1456 := __e.Get(1)
_ = W1456
tmp10853 := MakeNative(func(__e *ControlFlow) {
W1457 := __e.Get(1)
_ = W1457
tmp10856 := PrimEqual(W1456, sym__)

if True == tmp10856 {
tmp10854 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(PrimFunc(symshen_4comb), W1457, tmp10854)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10857 := Call(__e, PrimFunc(symtail), V1454)


__e.TailApply(tmp10853, tmp10857)
return


}, 1)

tmp10858 := Call(__e, PrimFunc(symhead), V1454)


tmp10859 := Call(__e, tmp10852, tmp10858)


ifres10851 = tmp10859


} else {
tmp10860 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10851 = tmp10860


}

__e.TailApply(tmp10848, ifres10851)
return


}, 1)

tmp10862 := Call(__e, ns2_1set, symshen_4_5wildcard_6, tmp10847)


_ = tmp10862

tmp10863 := MakeNative(func(__e *ControlFlow) {
V1458 := __e.Get(1)
_ = V1458
tmp10864 := MakeNative(func(__e *ControlFlow) {
W1459 := __e.Get(1)
_ = W1459
tmp10866 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1459)


if True == tmp10866 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1459)
return
}


}, 1)

tmp10876 := PrimIsPair(V1458)

var ifres10867 Obj

if True == tmp10876 {
tmp10868 := MakeNative(func(__e *ControlFlow) {
W1460 := __e.Get(1)
_ = W1460
tmp10869 := MakeNative(func(__e *ControlFlow) {
W1461 := __e.Get(1)
_ = W1461
tmp10871 := Call(__e, PrimFunc(symshen_4semicolon_2), W1460)


if True == tmp10871 {
__e.TailApply(PrimFunc(symshen_4comb), W1461, W1460)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10872 := Call(__e, PrimFunc(symtail), V1458)


__e.TailApply(tmp10869, tmp10872)
return


}, 1)

tmp10873 := Call(__e, PrimFunc(symhead), V1458)


tmp10874 := Call(__e, tmp10868, tmp10873)


ifres10867 = tmp10874


} else {
tmp10875 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10867 = tmp10875


}

__e.TailApply(tmp10864, ifres10867)
return


}, 1)

tmp10877 := Call(__e, ns2_1set, symshen_4_5sc_6, tmp10863)


_ = tmp10877

tmp10878 := MakeNative(func(__e *ControlFlow) {
V1462 := __e.Get(1)
_ = V1462
V1463 := __e.Get(2)
_ = V1463
tmp10879 := MakeNative(func(__e *ControlFlow) {
W1464 := __e.Get(1)
_ = W1464
tmp10880 := MakeNative(func(__e *ControlFlow) {
W1465 := __e.Get(1)
_ = W1465
tmp10881 := MakeNative(func(__e *ControlFlow) {
W1466 := __e.Get(1)
_ = W1466
tmp10882 := MakeNative(func(__e *ControlFlow) {
W1467 := __e.Get(1)
_ = W1467
tmp10883 := MakeNative(func(__e *ControlFlow) {
W1468 := __e.Get(1)
_ = W1468
tmp10884 := MakeNative(func(__e *ControlFlow) {
W1469 := __e.Get(1)
_ = W1469
tmp10885 := MakeNative(func(__e *ControlFlow) {
W1470 := __e.Get(1)
_ = W1470
tmp10886 := MakeNative(func(__e *ControlFlow) {
W1471 := __e.Get(1)
_ = W1471
tmp10887 := MakeNative(func(__e *ControlFlow) {
W1472 := __e.Get(1)
_ = W1472
__e.Return(W1472)
return
}, 1)

tmp10888 := PrimCons(sym_1_6, Nil)

tmp10889 := PrimCons(W1467, tmp10888)

tmp10890 := PrimCons(W1466, tmp10889)

tmp10891 := PrimCons(W1465, tmp10890)

tmp10892 := PrimCons(W1464, tmp10891)

tmp10893 := PrimCons(W1471, Nil)

tmp10894 := Call(__e, PrimFunc(symappend), tmp10892, tmp10893)


tmp10895 := Call(__e, PrimFunc(symappend), W1468, tmp10894)


tmp10896 := PrimCons(V1462, tmp10895)

tmp10897 := PrimCons(symdefine, tmp10896)

__e.TailApply(tmp10887, tmp10897)
return


}, 1)

var ifres10898 Obj

if True == W1469 {
tmp10899 := PrimCons(MakeNumber(1), Nil)

tmp10900 := PrimCons(W1466, tmp10899)

tmp10901 := PrimCons(sym_7, tmp10900)

tmp10902 := PrimCons(W1470, Nil)

tmp10903 := PrimCons(tmp10901, tmp10902)

tmp10904 := PrimCons(W1466, tmp10903)

tmp10905 := PrimCons(symlet, tmp10904)

ifres10898 = tmp10905


} else {
ifres10898 = W1470


}

__e.TailApply(tmp10886, ifres10898)
return


}, 1)

tmp10906 := Call(__e, PrimFunc(symshen_4prolog_1fbody), V1463, W1468, W1464, W1465, W1466, W1467, W1469)


__e.TailApply(tmp10885, tmp10906)
return


}, 1)

tmp10907 := Call(__e, PrimFunc(symshen_4hascut_2), V1463)


__e.TailApply(tmp10884, tmp10907)
return


}, 1)

tmp10908 := Call(__e, PrimFunc(symshen_4prolog_1parameters), V1463)


__e.TailApply(tmp10883, tmp10908)
return


}, 1)

tmp10909 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp10882, tmp10909)
return


}, 1)

tmp10910 := Call(__e, PrimFunc(symgensym), symK)


__e.TailApply(tmp10881, tmp10910)
return


}, 1)

tmp10911 := Call(__e, PrimFunc(symgensym), symL)


__e.TailApply(tmp10880, tmp10911)
return


}, 1)

tmp10912 := Call(__e, PrimFunc(symgensym), symB)


__e.TailApply(tmp10879, tmp10912)
return


}, 2)

tmp10913 := Call(__e, ns2_1set, symshen_4horn_1clause_1procedure, tmp10878)


_ = tmp10913

tmp10914 := MakeNative(func(__e *ControlFlow) {
V1475 := __e.Get(1)
_ = V1475
tmp10924 := PrimEqual(sym_b, V1475)

if True == tmp10924 {
__e.Return(True)
return
} else {
tmp10922 := PrimIsPair(V1475)

if True == tmp10922 {
tmp10919 := PrimHead(V1475)

tmp10920 := Call(__e, PrimFunc(symshen_4hascut_2), tmp10919)


if True == tmp10920 {
__e.Return(True)
return
} else {
tmp10916 := PrimTail(V1475)

tmp10917 := Call(__e, PrimFunc(symshen_4hascut_2), tmp10916)


if True == tmp10917 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


} else {
__e.Return(False)
return
}


}


}, 1)

tmp10925 := Call(__e, ns2_1set, symshen_4hascut_2, tmp10914)


_ = tmp10925

tmp10926 := MakeNative(func(__e *ControlFlow) {
V1480 := __e.Get(1)
_ = V1480
tmp10935 := PrimIsPair(V1480)

var ifres10931 Obj

if True == tmp10935 {
tmp10933 := PrimHead(V1480)

tmp10934 := PrimIsPair(tmp10933)

var ifres10932 Obj

if True == tmp10934 {
ifres10932 = True


} else {
ifres10932 = False


}

ifres10931 = ifres10932


} else {
ifres10931 = False


}

if True == ifres10931 {
tmp10927 := PrimHead(V1480)

tmp10928 := PrimHead(tmp10927)

tmp10929 := Call(__e, PrimFunc(symlength), tmp10928)


__e.TailApply(PrimFunc(symshen_4parameters), tmp10929)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.prolog-parameters")))
return
}


}, 1)

tmp10936 := Call(__e, ns2_1set, symshen_4prolog_1parameters, tmp10926)


_ = tmp10936

tmp10937 := MakeNative(func(__e *ControlFlow) {
V1501 := __e.Get(1)
_ = V1501
V1502 := __e.Get(2)
_ = V1502
V1503 := __e.Get(3)
_ = V1503
V1504 := __e.Get(4)
_ = V1504
V1505 := __e.Get(5)
_ = V1505
V1506 := __e.Get(6)
_ = V1506
V1507 := __e.Get(7)
_ = V1507
tmp11030 := PrimEqual(Nil, V1501)

var ifres11027 Obj

if True == tmp11030 {
tmp11029 := PrimEqual(True, V1507)

var ifres11028 Obj

if True == tmp11029 {
ifres11028 = True


} else {
ifres11028 = False


}

ifres11027 = ifres11028


} else {
ifres11027 = False


}

if True == ifres11027 {
tmp10938 := PrimCons(V1505, Nil)

tmp10939 := PrimCons(V1504, tmp10938)

__e.Return(PrimCons(symshen_4unlock, tmp10939))
return


} else {
tmp11025 := PrimIsPair(V1501)

var ifres11003 Obj

if True == tmp11025 {
tmp11023 := PrimHead(V1501)

tmp11024 := PrimIsPair(tmp11023)

var ifres11005 Obj

if True == tmp11024 {
tmp11020 := PrimHead(V1501)

tmp11021 := PrimTail(tmp11020)

tmp11022 := PrimIsPair(tmp11021)

var ifres11007 Obj

if True == tmp11022 {
tmp11016 := PrimHead(V1501)

tmp11017 := PrimTail(tmp11016)

tmp11018 := PrimTail(tmp11017)

tmp11019 := PrimEqual(Nil, tmp11018)

var ifres11009 Obj

if True == tmp11019 {
tmp11014 := PrimTail(V1501)

tmp11015 := PrimEqual(Nil, tmp11014)

var ifres11011 Obj

if True == tmp11015 {
tmp11013 := PrimEqual(False, V1507)

var ifres11012 Obj

if True == tmp11013 {
ifres11012 = True


} else {
ifres11012 = False


}

ifres11011 = ifres11012


} else {
ifres11011 = False


}

var ifres11010 Obj

if True == ifres11011 {
ifres11010 = True


} else {
ifres11010 = False


}

ifres11009 = ifres11010


} else {
ifres11009 = False


}

var ifres11008 Obj

if True == ifres11009 {
ifres11008 = True


} else {
ifres11008 = False


}

ifres11007 = ifres11008


} else {
ifres11007 = False


}

var ifres11006 Obj

if True == ifres11007 {
ifres11006 = True


} else {
ifres11006 = False


}

ifres11005 = ifres11006


} else {
ifres11005 = False


}

var ifres11004 Obj

if True == ifres11005 {
ifres11004 = True


} else {
ifres11004 = False


}

ifres11003 = ifres11004


} else {
ifres11003 = False


}

if True == ifres11003 {
tmp10940 := MakeNative(func(__e *ControlFlow) {
W1508 := __e.Get(1)
_ = W1508
tmp10941 := PrimCons(V1504, Nil)

tmp10942 := PrimCons(symshen_4unlocked_2, tmp10941)

tmp10943 := PrimHead(V1501)

tmp10944 := PrimHead(tmp10943)

tmp10945 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10944, V1502, V1503, W1508)


tmp10946 := PrimCons(False, Nil)

tmp10947 := PrimCons(tmp10945, tmp10946)

tmp10948 := PrimCons(tmp10942, tmp10947)

__e.Return(PrimCons(symif, tmp10948))
return


}, 1)

tmp10949 := PrimHead(V1501)

tmp10950 := PrimHead(tmp10949)

tmp10951 := PrimHead(V1501)

tmp10952 := PrimTail(tmp10951)

tmp10953 := PrimHead(tmp10952)

tmp10954 := Call(__e, PrimFunc(symshen_4continue), tmp10950, tmp10953, V1503, V1504, V1505, V1506)


__e.TailApply(tmp10940, tmp10954)
return


} else {
tmp11001 := PrimIsPair(V1501)

var ifres10986 Obj

if True == tmp11001 {
tmp10999 := PrimHead(V1501)

tmp11000 := PrimIsPair(tmp10999)

var ifres10988 Obj

if True == tmp11000 {
tmp10996 := PrimHead(V1501)

tmp10997 := PrimTail(tmp10996)

tmp10998 := PrimIsPair(tmp10997)

var ifres10990 Obj

if True == tmp10998 {
tmp10992 := PrimHead(V1501)

tmp10993 := PrimTail(tmp10992)

tmp10994 := PrimTail(tmp10993)

tmp10995 := PrimEqual(Nil, tmp10994)

var ifres10991 Obj

if True == tmp10995 {
ifres10991 = True


} else {
ifres10991 = False


}

ifres10990 = ifres10991


} else {
ifres10990 = False


}

var ifres10989 Obj

if True == ifres10990 {
ifres10989 = True


} else {
ifres10989 = False


}

ifres10988 = ifres10989


} else {
ifres10988 = False


}

var ifres10987 Obj

if True == ifres10988 {
ifres10987 = True


} else {
ifres10987 = False


}

ifres10986 = ifres10987


} else {
ifres10986 = False


}

if True == ifres10986 {
tmp10955 := MakeNative(func(__e *ControlFlow) {
W1509 := __e.Get(1)
_ = W1509
tmp10956 := MakeNative(func(__e *ControlFlow) {
W1510 := __e.Get(1)
_ = W1510
tmp10957 := PrimCons(V1504, Nil)

tmp10958 := PrimCons(symshen_4unlocked_2, tmp10957)

tmp10959 := PrimHead(V1501)

tmp10960 := PrimHead(tmp10959)

tmp10961 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10960, V1502, V1503, W1510)


tmp10962 := PrimCons(False, Nil)

tmp10963 := PrimCons(tmp10961, tmp10962)

tmp10964 := PrimCons(tmp10958, tmp10963)

tmp10965 := PrimCons(symif, tmp10964)

tmp10966 := PrimCons(False, Nil)

tmp10967 := PrimCons(W1509, tmp10966)

tmp10968 := PrimCons(sym_a, tmp10967)

tmp10969 := PrimTail(V1501)

tmp10970 := Call(__e, PrimFunc(symshen_4prolog_1fbody), tmp10969, V1502, V1503, V1504, V1505, V1506, V1507)


tmp10971 := PrimCons(W1509, Nil)

tmp10972 := PrimCons(tmp10970, tmp10971)

tmp10973 := PrimCons(tmp10968, tmp10972)

tmp10974 := PrimCons(symif, tmp10973)

tmp10975 := PrimCons(tmp10974, Nil)

tmp10976 := PrimCons(tmp10965, tmp10975)

tmp10977 := PrimCons(W1509, tmp10976)

__e.Return(PrimCons(symlet, tmp10977))
return


}, 1)

tmp10978 := PrimHead(V1501)

tmp10979 := PrimHead(tmp10978)

tmp10980 := PrimHead(V1501)

tmp10981 := PrimTail(tmp10980)

tmp10982 := PrimHead(tmp10981)

tmp10983 := Call(__e, PrimFunc(symshen_4continue), tmp10979, tmp10982, V1503, V1504, V1505, V1506)


__e.TailApply(tmp10956, tmp10983)
return


}, 1)

tmp10984 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp10955, tmp10984)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.prolog-fbody")))
return
}


}


}


}, 7)

tmp11031 := Call(__e, ns2_1set, symshen_4prolog_1fbody, tmp10937)


_ = tmp11031

tmp11032 := MakeNative(func(__e *ControlFlow) {
V1511 := __e.Get(1)
_ = V1511
V1512 := __e.Get(2)
_ = V1512
tmp11037 := Call(__e, PrimFunc(symshen_4locked_2), V1511)


var ifres11034 Obj

if True == tmp11037 {
tmp11036 := Call(__e, PrimFunc(symshen_4fits_2), V1512, V1511)


var ifres11035 Obj

if True == tmp11036 {
ifres11035 = True


} else {
ifres11035 = False


}

ifres11034 = ifres11035


} else {
ifres11034 = False


}

if True == ifres11034 {
__e.TailApply(PrimFunc(symshen_4openlock), V1511)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp11038 := Call(__e, ns2_1set, symshen_4unlock, tmp11032)


_ = tmp11038

tmp11039 := MakeNative(func(__e *ControlFlow) {
V1513 := __e.Get(1)
_ = V1513
tmp11040 := Call(__e, PrimFunc(symshen_4unlocked_2), V1513)


__e.Return(PrimNot(tmp11040))
return


}, 1)

tmp11041 := Call(__e, ns2_1set, symshen_4locked_2, tmp11039)


_ = tmp11041

tmp11042 := MakeNative(func(__e *ControlFlow) {
V1514 := __e.Get(1)
_ = V1514
__e.Return(PrimVectorGet(V1514, MakeNumber(1)))
return
}, 1)

tmp11043 := Call(__e, ns2_1set, symshen_4unlocked_2, tmp11042)


_ = tmp11043

tmp11044 := MakeNative(func(__e *ControlFlow) {
V1515 := __e.Get(1)
_ = V1515
tmp11045 := PrimVectorSet(V1515, MakeNumber(1), True)

_ = tmp11045

__e.Return(False)
return


}, 1)

tmp11046 := Call(__e, ns2_1set, symshen_4openlock, tmp11044)


_ = tmp11046

tmp11047 := MakeNative(func(__e *ControlFlow) {
V1516 := __e.Get(1)
_ = V1516
V1517 := __e.Get(2)
_ = V1517
tmp11048 := PrimVectorGet(V1517, MakeNumber(2))

__e.Return(PrimEqual(V1516, tmp11048))
return


}, 2)

tmp11049 := Call(__e, ns2_1set, symshen_4fits_2, tmp11047)


_ = tmp11049

tmp11050 := MakeNative(func(__e *ControlFlow) {
V1520 := __e.Get(1)
_ = V1520
V1521 := __e.Get(2)
_ = V1521
V1522 := __e.Get(3)
_ = V1522
V1523 := __e.Get(4)
_ = V1523
tmp11051 := MakeNative(func(__e *ControlFlow) {
W1524 := __e.Get(1)
_ = W1524
tmp11056 := PrimEqual(W1524, False)

var ifres11053 Obj

if True == tmp11056 {
tmp11055 := Call(__e, PrimFunc(symshen_4unlocked_2), V1521)


var ifres11054 Obj

if True == tmp11055 {
ifres11054 = True


} else {
ifres11054 = False


}

ifres11053 = ifres11054


} else {
ifres11053 = False


}

if True == ifres11053 {
__e.TailApply(PrimFunc(symshen_4lock), V1522, V1521)
return
} else {
__e.Return(W1524)
return
}


}, 1)

tmp11057 := Call(__e, PrimFunc(symthaw), V1523)


__e.TailApply(tmp11051, tmp11057)
return


}, 4)

tmp11058 := Call(__e, ns2_1set, symshen_4cut, tmp11050)


_ = tmp11058

tmp11059 := MakeNative(func(__e *ControlFlow) {
V1525 := __e.Get(1)
_ = V1525
V1526 := __e.Get(2)
_ = V1526
tmp11060 := MakeNative(func(__e *ControlFlow) {
W1527 := __e.Get(1)
_ = W1527
tmp11061 := MakeNative(func(__e *ControlFlow) {
W1528 := __e.Get(1)
_ = W1528
__e.Return(False)
return
}, 1)

tmp11062 := PrimVectorSet(V1526, MakeNumber(2), V1525)

__e.TailApply(tmp11061, tmp11062)
return


}, 1)

tmp11063 := PrimVectorSet(V1526, MakeNumber(1), False)

__e.TailApply(tmp11060, tmp11063)
return


}, 2)

tmp11064 := Call(__e, ns2_1set, symshen_4lock, tmp11059)


_ = tmp11064

tmp11065 := MakeNative(func(__e *ControlFlow) {
V1529 := __e.Get(1)
_ = V1529
V1530 := __e.Get(2)
_ = V1530
V1531 := __e.Get(3)
_ = V1531
V1532 := __e.Get(4)
_ = V1532
V1533 := __e.Get(5)
_ = V1533
V1534 := __e.Get(6)
_ = V1534
tmp11066 := MakeNative(func(__e *ControlFlow) {
W1535 := __e.Get(1)
_ = W1535
tmp11067 := MakeNative(func(__e *ControlFlow) {
W1536 := __e.Get(1)
_ = W1536
tmp11068 := MakeNative(func(__e *ControlFlow) {
W1537 := __e.Get(1)
_ = W1537
tmp11069 := MakeNative(func(__e *ControlFlow) {
W1538 := __e.Get(1)
_ = W1538
__e.TailApply(PrimFunc(symshen_4stpart), W1537, W1538, V1531)
return
}, 1)

tmp11070 := PrimCons(symshen_4incinfs, Nil)

tmp11071 := Call(__e, PrimFunc(symshen_4compile_1body), V1530, V1531, V1532, V1533, V1534)


tmp11072 := PrimCons(tmp11071, Nil)

tmp11073 := PrimCons(tmp11070, tmp11072)

tmp11074 := PrimCons(symdo, tmp11073)

__e.TailApply(tmp11069, tmp11074)
return


}, 1)

tmp11075 := Call(__e, PrimFunc(symdifference), W1536, W1535)


__e.TailApply(tmp11068, tmp11075)
return


}, 1)

tmp11076 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), V1530)


__e.TailApply(tmp11067, tmp11076)
return


}, 1)

tmp11077 := Call(__e, PrimFunc(symshen_4extract_1vars), V1529)


__e.TailApply(tmp11066, tmp11077)
return


}, 6)

tmp11078 := Call(__e, ns2_1set, symshen_4continue, tmp11065)


_ = tmp11078

tmp11079 := MakeNative(func(__e *ControlFlow) {
V1541 := __e.Get(1)
_ = V1541
tmp11114 := PrimIsPair(V1541)

var ifres11095 Obj

if True == tmp11114 {
tmp11112 := PrimHead(V1541)

tmp11113 := PrimEqual(symlambda, tmp11112)

var ifres11097 Obj

if True == tmp11113 {
tmp11110 := PrimTail(V1541)

tmp11111 := PrimIsPair(tmp11110)

var ifres11099 Obj

if True == tmp11111 {
tmp11107 := PrimTail(V1541)

tmp11108 := PrimTail(tmp11107)

tmp11109 := PrimIsPair(tmp11108)

var ifres11101 Obj

if True == tmp11109 {
tmp11103 := PrimTail(V1541)

tmp11104 := PrimTail(tmp11103)

tmp11105 := PrimTail(tmp11104)

tmp11106 := PrimEqual(Nil, tmp11105)

var ifres11102 Obj

if True == tmp11106 {
ifres11102 = True


} else {
ifres11102 = False


}

ifres11101 = ifres11102


} else {
ifres11101 = False


}

var ifres11100 Obj

if True == ifres11101 {
ifres11100 = True


} else {
ifres11100 = False


}

ifres11099 = ifres11100


} else {
ifres11099 = False


}

var ifres11098 Obj

if True == ifres11099 {
ifres11098 = True


} else {
ifres11098 = False


}

ifres11097 = ifres11098


} else {
ifres11097 = False


}

var ifres11096 Obj

if True == ifres11097 {
ifres11096 = True


} else {
ifres11096 = False


}

ifres11095 = ifres11096


} else {
ifres11095 = False


}

if True == ifres11095 {
tmp11080 := PrimTail(V1541)

tmp11081 := PrimHead(tmp11080)

tmp11082 := PrimTail(V1541)

tmp11083 := PrimTail(tmp11082)

tmp11084 := PrimHead(tmp11083)

tmp11085 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11084)


__e.TailApply(PrimFunc(symremove), tmp11081, tmp11085)
return


} else {
tmp11093 := PrimIsPair(V1541)

if True == tmp11093 {
tmp11086 := PrimHead(V1541)

tmp11087 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11086)


tmp11088 := PrimTail(V1541)

tmp11089 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11088)


__e.TailApply(PrimFunc(symunion), tmp11087, tmp11089)
return


} else {
tmp11091 := PrimIsVariable(V1541)

if True == tmp11091 {
__e.Return(PrimCons(V1541, Nil))
return
} else {
__e.Return(Nil)
return
}


}


}


}, 1)

tmp11115 := Call(__e, ns2_1set, symshen_4extract_1free_1vars, tmp11079)


_ = tmp11115

tmp11116 := MakeNative(func(__e *ControlFlow) {
V1558 := __e.Get(1)
_ = V1558
V1559 := __e.Get(2)
_ = V1559
V1560 := __e.Get(3)
_ = V1560
V1561 := __e.Get(4)
_ = V1561
V1562 := __e.Get(5)
_ = V1562
tmp11151 := PrimEqual(Nil, V1558)

if True == tmp11151 {
tmp11117 := PrimCons(V1562, Nil)

__e.Return(PrimCons(symthaw, tmp11117))
return


} else {
tmp11149 := PrimIsPair(V1558)

var ifres11145 Obj

if True == tmp11149 {
tmp11147 := PrimHead(V1558)

tmp11148 := PrimEqual(sym_b, tmp11147)

var ifres11146 Obj

if True == tmp11148 {
ifres11146 = True


} else {
ifres11146 = False


}

ifres11145 = ifres11146


} else {
ifres11145 = False


}

if True == ifres11145 {
tmp11118 := PrimCons(symshen_4cut, Nil)

tmp11119 := PrimTail(V1558)

tmp11120 := PrimCons(tmp11118, tmp11119)

__e.TailApply(PrimFunc(symshen_4compile_1body), tmp11120, V1559, V1560, V1561, V1562)
return


} else {
tmp11143 := PrimIsPair(V1558)

var ifres11139 Obj

if True == tmp11143 {
tmp11141 := PrimTail(V1558)

tmp11142 := PrimEqual(Nil, tmp11141)

var ifres11140 Obj

if True == tmp11142 {
ifres11140 = True


} else {
ifres11140 = False


}

ifres11139 = ifres11140


} else {
ifres11139 = False


}

if True == ifres11139 {
tmp11121 := PrimHead(V1558)

tmp11122 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11121, V1559)


tmp11123 := PrimCons(V1562, Nil)

tmp11124 := PrimCons(V1561, tmp11123)

tmp11125 := PrimCons(V1560, tmp11124)

tmp11126 := PrimCons(V1559, tmp11125)

__e.TailApply(PrimFunc(symappend), tmp11122, tmp11126)
return


} else {
tmp11137 := PrimIsPair(V1558)

if True == tmp11137 {
tmp11127 := MakeNative(func(__e *ControlFlow) {
W1563 := __e.Get(1)
_ = W1563
tmp11128 := PrimTail(V1558)

tmp11129 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11128, V1559, V1560, V1561, V1562)


tmp11130 := PrimCons(tmp11129, Nil)

tmp11131 := PrimCons(V1561, tmp11130)

tmp11132 := PrimCons(V1560, tmp11131)

tmp11133 := PrimCons(V1559, tmp11132)

__e.TailApply(PrimFunc(symappend), W1563, tmp11133)
return


}, 1)

tmp11134 := PrimHead(V1558)

tmp11135 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11134, V1559)


__e.TailApply(tmp11127, tmp11135)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-fbody")))
return
}


}


}


}


}, 5)

tmp11152 := Call(__e, ns2_1set, symshen_4compile_1body, tmp11116)


_ = tmp11152

tmp11153 := MakeNative(func(__e *ControlFlow) {
V1580 := __e.Get(1)
_ = V1580
V1581 := __e.Get(2)
_ = V1581
V1582 := __e.Get(3)
_ = V1582
V1583 := __e.Get(4)
_ = V1583
V1584 := __e.Get(5)
_ = V1584
tmp11177 := PrimEqual(Nil, V1580)

if True == tmp11177 {
__e.Return(V1584)
return
} else {
tmp11175 := PrimIsPair(V1580)

var ifres11171 Obj

if True == tmp11175 {
tmp11173 := PrimHead(V1580)

tmp11174 := PrimEqual(sym_b, tmp11173)

var ifres11172 Obj

if True == tmp11174 {
ifres11172 = True


} else {
ifres11172 = False


}

ifres11171 = ifres11172


} else {
ifres11171 = False


}

if True == ifres11171 {
tmp11154 := PrimCons(symshen_4cut, Nil)

tmp11155 := PrimTail(V1580)

tmp11156 := PrimCons(tmp11154, tmp11155)

__e.TailApply(PrimFunc(symshen_4freeze_1literals), tmp11156, V1581, V1582, V1583, V1584)
return


} else {
tmp11169 := PrimIsPair(V1580)

if True == tmp11169 {
tmp11157 := MakeNative(func(__e *ControlFlow) {
W1585 := __e.Get(1)
_ = W1585
tmp11158 := PrimTail(V1580)

tmp11159 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11158, V1581, V1582, V1583, V1584)


tmp11160 := PrimCons(tmp11159, Nil)

tmp11161 := PrimCons(V1583, tmp11160)

tmp11162 := PrimCons(V1582, tmp11161)

tmp11163 := PrimCons(V1581, tmp11162)

tmp11164 := Call(__e, PrimFunc(symappend), W1585, tmp11163)


tmp11165 := PrimCons(tmp11164, Nil)

__e.Return(PrimCons(symfreeze, tmp11165))
return


}, 1)

tmp11166 := PrimHead(V1580)

tmp11167 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11166, V1581)


__e.TailApply(tmp11157, tmp11167)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.freeze-literals")))
return
}


}


}


}, 5)

tmp11178 := Call(__e, ns2_1set, symshen_4freeze_1literals, tmp11153)


_ = tmp11178

tmp11179 := MakeNative(func(__e *ControlFlow) {
V1590 := __e.Get(1)
_ = V1590
V1591 := __e.Get(2)
_ = V1591
tmp11194 := PrimIsPair(V1590)

var ifres11190 Obj

if True == tmp11194 {
tmp11192 := PrimHead(V1590)

tmp11193 := PrimEqual(symfork, tmp11192)

var ifres11191 Obj

if True == tmp11193 {
ifres11191 = True


} else {
ifres11191 = False


}

ifres11190 = ifres11191


} else {
ifres11190 = False


}

if True == ifres11190 {
tmp11180 := PrimTail(V1590)

tmp11181 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11180, V1591)


tmp11182 := PrimCons(tmp11181, Nil)

__e.Return(PrimCons(symfork, tmp11182))
return


} else {
tmp11188 := PrimIsPair(V1590)

if True == tmp11188 {
tmp11183 := PrimHead(V1590)

tmp11184 := MakeNative(func(__e *ControlFlow) {
Z1592 := __e.Get(1)
_ = Z1592
__e.TailApply(PrimFunc(symshen_4function_1calls), Z1592, V1591)
return
}, 1)

tmp11185 := PrimTail(V1590)

tmp11186 := Call(__e, PrimFunc(symmap), tmp11184, tmp11185)


__e.Return(PrimCons(tmp11183, tmp11186))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.deref-calls")))
return
}


}


}, 2)

tmp11195 := Call(__e, ns2_1set, symshen_4deref_1calls, tmp11179)


_ = tmp11195

tmp11196 := MakeNative(func(__e *ControlFlow) {
V1599 := __e.Get(1)
_ = V1599
V1600 := __e.Get(2)
_ = V1600
tmp11206 := PrimEqual(Nil, V1599)

if True == tmp11206 {
__e.Return(Nil)
return
} else {
tmp11204 := PrimIsPair(V1599)

if True == tmp11204 {
tmp11197 := PrimHead(V1599)

tmp11198 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11197, V1600)


tmp11199 := PrimTail(V1599)

tmp11200 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11199, V1600)


tmp11201 := PrimCons(tmp11200, Nil)

tmp11202 := PrimCons(tmp11198, tmp11201)

__e.Return(PrimCons(symcons, tmp11202))
return


} else {
__e.Return(PrimSimpleError(MakeString("fork requires a list of literals\n")))
return
}


}


}, 2)

tmp11207 := Call(__e, ns2_1set, symshen_4deref_1forked_1literals, tmp11196)


_ = tmp11207

tmp11208 := MakeNative(func(__e *ControlFlow) {
V1603 := __e.Get(1)
_ = V1603
V1604 := __e.Get(2)
_ = V1604
tmp11240 := PrimIsPair(V1603)

var ifres11221 Obj

if True == tmp11240 {
tmp11238 := PrimHead(V1603)

tmp11239 := PrimEqual(symcons, tmp11238)

var ifres11223 Obj

if True == tmp11239 {
tmp11236 := PrimTail(V1603)

tmp11237 := PrimIsPair(tmp11236)

var ifres11225 Obj

if True == tmp11237 {
tmp11233 := PrimTail(V1603)

tmp11234 := PrimTail(tmp11233)

tmp11235 := PrimIsPair(tmp11234)

var ifres11227 Obj

if True == tmp11235 {
tmp11229 := PrimTail(V1603)

tmp11230 := PrimTail(tmp11229)

tmp11231 := PrimTail(tmp11230)

tmp11232 := PrimEqual(Nil, tmp11231)

var ifres11228 Obj

if True == tmp11232 {
ifres11228 = True


} else {
ifres11228 = False


}

ifres11227 = ifres11228


} else {
ifres11227 = False


}

var ifres11226 Obj

if True == ifres11227 {
ifres11226 = True


} else {
ifres11226 = False


}

ifres11225 = ifres11226


} else {
ifres11225 = False


}

var ifres11224 Obj

if True == ifres11225 {
ifres11224 = True


} else {
ifres11224 = False


}

ifres11223 = ifres11224


} else {
ifres11223 = False


}

var ifres11222 Obj

if True == ifres11223 {
ifres11222 = True


} else {
ifres11222 = False


}

ifres11221 = ifres11222


} else {
ifres11221 = False


}

if True == ifres11221 {
tmp11209 := PrimTail(V1603)

tmp11210 := PrimHead(tmp11209)

tmp11211 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11210, V1604)


tmp11212 := PrimTail(V1603)

tmp11213 := PrimTail(tmp11212)

tmp11214 := PrimHead(tmp11213)

tmp11215 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11214, V1604)


tmp11216 := PrimCons(tmp11215, Nil)

tmp11217 := PrimCons(tmp11211, tmp11216)

__e.Return(PrimCons(symcons, tmp11217))
return


} else {
tmp11219 := PrimIsPair(V1603)

if True == tmp11219 {
__e.TailApply(PrimFunc(symshen_4deref_1terms), V1603, V1604, Nil)
return
} else {
__e.Return(V1603)
return
}


}


}, 2)

tmp11241 := Call(__e, ns2_1set, symshen_4function_1calls, tmp11208)


_ = tmp11241

tmp11242 := MakeNative(func(__e *ControlFlow) {
V1613 := __e.Get(1)
_ = V1613
V1614 := __e.Get(2)
_ = V1614
V1615 := __e.Get(3)
_ = V1615
tmp11336 := PrimIsPair(V1613)

var ifres11323 Obj

if True == tmp11336 {
tmp11334 := PrimHead(V1613)

tmp11335 := PrimEqual(MakeNumber(0), tmp11334)

var ifres11325 Obj

if True == tmp11335 {
tmp11332 := PrimTail(V1613)

tmp11333 := PrimIsPair(tmp11332)

var ifres11327 Obj

if True == tmp11333 {
tmp11329 := PrimTail(V1613)

tmp11330 := PrimTail(tmp11329)

tmp11331 := PrimEqual(Nil, tmp11330)

var ifres11328 Obj

if True == tmp11331 {
ifres11328 = True


} else {
ifres11328 = False


}

ifres11327 = ifres11328


} else {
ifres11327 = False


}

var ifres11326 Obj

if True == ifres11327 {
ifres11326 = True


} else {
ifres11326 = False


}

ifres11325 = ifres11326


} else {
ifres11325 = False


}

var ifres11324 Obj

if True == ifres11325 {
ifres11324 = True


} else {
ifres11324 = False


}

ifres11323 = ifres11324


} else {
ifres11323 = False


}

if True == ifres11323 {
tmp11249 := PrimTail(V1613)

tmp11250 := PrimHead(tmp11249)

tmp11251 := PrimIsVariable(tmp11250)

if True == tmp11251 {
tmp11243 := PrimTail(V1613)

__e.Return(PrimHead(tmp11243))
return


} else {
tmp11244 := PrimTail(V1613)

tmp11245 := PrimHead(tmp11244)

tmp11246 := Call(__e, PrimFunc(symshen_4app), tmp11245, MakeString("\n"), symshen_4s)


tmp11247 := PrimStringConcat(MakeString("attempt to optimise a non-variable "), tmp11246)

__e.Return(PrimSimpleError(tmp11247))
return


}


} else {
tmp11321 := PrimIsPair(V1613)

var ifres11308 Obj

if True == tmp11321 {
tmp11319 := PrimHead(V1613)

tmp11320 := PrimEqual(MakeNumber(1), tmp11319)

var ifres11310 Obj

if True == tmp11320 {
tmp11317 := PrimTail(V1613)

tmp11318 := PrimIsPair(tmp11317)

var ifres11312 Obj

if True == tmp11318 {
tmp11314 := PrimTail(V1613)

tmp11315 := PrimTail(tmp11314)

tmp11316 := PrimEqual(Nil, tmp11315)

var ifres11313 Obj

if True == tmp11316 {
ifres11313 = True


} else {
ifres11313 = False


}

ifres11312 = ifres11313


} else {
ifres11312 = False


}

var ifres11311 Obj

if True == ifres11312 {
ifres11311 = True


} else {
ifres11311 = False


}

ifres11310 = ifres11311


} else {
ifres11310 = False


}

var ifres11309 Obj

if True == ifres11310 {
ifres11309 = True


} else {
ifres11309 = False


}

ifres11308 = ifres11309


} else {
ifres11308 = False


}

if True == ifres11308 {
tmp11261 := PrimTail(V1613)

tmp11262 := PrimHead(tmp11261)

tmp11263 := PrimIsVariable(tmp11262)

if True == tmp11263 {
tmp11252 := PrimTail(V1613)

tmp11253 := PrimHead(tmp11252)

tmp11254 := PrimCons(V1614, Nil)

tmp11255 := PrimCons(tmp11253, tmp11254)

__e.Return(PrimCons(symshen_4lazyderef, tmp11255))
return


} else {
tmp11256 := PrimTail(V1613)

tmp11257 := PrimHead(tmp11256)

tmp11258 := Call(__e, PrimFunc(symshen_4app), tmp11257, MakeString("\n"), symshen_4s)


tmp11259 := PrimStringConcat(MakeString("attempt to optimise a non-variable "), tmp11258)

__e.Return(PrimSimpleError(tmp11259))
return


}


} else {
tmp11305 := Call(__e, PrimFunc(symelement_2), V1613, V1615)


tmp11306 := PrimNot(tmp11305)

var ifres11302 Obj

if True == tmp11306 {
tmp11304 := PrimIsVariable(V1613)

var ifres11303 Obj

if True == tmp11304 {
ifres11303 = True


} else {
ifres11303 = False


}

ifres11302 = ifres11303


} else {
ifres11302 = False


}

if True == ifres11302 {
tmp11264 := PrimCons(V1614, Nil)

tmp11265 := PrimCons(V1613, tmp11264)

__e.Return(PrimCons(symshen_4deref, tmp11265))
return


} else {
tmp11300 := PrimIsPair(V1613)

var ifres11281 Obj

if True == tmp11300 {
tmp11298 := PrimHead(V1613)

tmp11299 := PrimEqual(symlambda, tmp11298)

var ifres11283 Obj

if True == tmp11299 {
tmp11296 := PrimTail(V1613)

tmp11297 := PrimIsPair(tmp11296)

var ifres11285 Obj

if True == tmp11297 {
tmp11293 := PrimTail(V1613)

tmp11294 := PrimTail(tmp11293)

tmp11295 := PrimIsPair(tmp11294)

var ifres11287 Obj

if True == tmp11295 {
tmp11289 := PrimTail(V1613)

tmp11290 := PrimTail(tmp11289)

tmp11291 := PrimTail(tmp11290)

tmp11292 := PrimEqual(Nil, tmp11291)

var ifres11288 Obj

if True == tmp11292 {
ifres11288 = True


} else {
ifres11288 = False


}

ifres11287 = ifres11288


} else {
ifres11287 = False


}

var ifres11286 Obj

if True == ifres11287 {
ifres11286 = True


} else {
ifres11286 = False


}

ifres11285 = ifres11286


} else {
ifres11285 = False


}

var ifres11284 Obj

if True == ifres11285 {
ifres11284 = True


} else {
ifres11284 = False


}

ifres11283 = ifres11284


} else {
ifres11283 = False


}

var ifres11282 Obj

if True == ifres11283 {
ifres11282 = True


} else {
ifres11282 = False


}

ifres11281 = ifres11282


} else {
ifres11281 = False


}

if True == ifres11281 {
tmp11266 := PrimTail(V1613)

tmp11267 := PrimHead(tmp11266)

tmp11268 := PrimTail(V1613)

tmp11269 := PrimTail(tmp11268)

tmp11270 := PrimHead(tmp11269)

tmp11271 := PrimTail(V1613)

tmp11272 := PrimHead(tmp11271)

tmp11273 := PrimCons(tmp11272, V1615)

tmp11274 := Call(__e, PrimFunc(symshen_4deref_1terms), tmp11270, V1614, tmp11273)


tmp11275 := PrimCons(tmp11274, Nil)

tmp11276 := PrimCons(tmp11267, tmp11275)

__e.Return(PrimCons(symlambda, tmp11276))
return


} else {
tmp11279 := PrimIsPair(V1613)

if True == tmp11279 {
tmp11277 := MakeNative(func(__e *ControlFlow) {
Z1616 := __e.Get(1)
_ = Z1616
__e.TailApply(PrimFunc(symshen_4deref_1terms), Z1616, V1614, V1615)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11277, V1613)
return


} else {
__e.Return(V1613)
return
}


}


}


}


}


}, 3)

tmp11337 := Call(__e, ns2_1set, symshen_4deref_1terms, tmp11242)


_ = tmp11337

tmp11338 := MakeNative(func(__e *ControlFlow) {
V1634 := __e.Get(1)
_ = V1634
V1635 := __e.Get(2)
_ = V1635
V1636 := __e.Get(3)
_ = V1636
V1637 := __e.Get(4)
_ = V1637
V1638 := __e.Get(5)
_ = V1638
tmp11514 := PrimEqual(Nil, V1635)

var ifres11511 Obj

if True == tmp11514 {
tmp11513 := PrimEqual(Nil, V1636)

var ifres11512 Obj

if True == tmp11513 {
ifres11512 = True


} else {
ifres11512 = False


}

ifres11511 = ifres11512


} else {
ifres11511 = False


}

if True == ifres11511 {
__e.Return(V1638)
return
} else {
tmp11509 := PrimIsPair(V1635)

var ifres11489 Obj

if True == tmp11509 {
tmp11507 := PrimHead(V1635)

tmp11508 := PrimIsPair(tmp11507)

var ifres11491 Obj

if True == tmp11508 {
tmp11504 := PrimHead(V1635)

tmp11505 := PrimHead(tmp11504)

tmp11506 := PrimEqual(symshen_4_7m, tmp11505)

var ifres11493 Obj

if True == tmp11506 {
tmp11501 := PrimHead(V1635)

tmp11502 := PrimTail(tmp11501)

tmp11503 := PrimIsPair(tmp11502)

var ifres11495 Obj

if True == tmp11503 {
tmp11497 := PrimHead(V1635)

tmp11498 := PrimTail(tmp11497)

tmp11499 := PrimTail(tmp11498)

tmp11500 := PrimEqual(Nil, tmp11499)

var ifres11496 Obj

if True == tmp11500 {
ifres11496 = True


} else {
ifres11496 = False


}

ifres11495 = ifres11496


} else {
ifres11495 = False


}

var ifres11494 Obj

if True == ifres11495 {
ifres11494 = True


} else {
ifres11494 = False


}

ifres11493 = ifres11494


} else {
ifres11493 = False


}

var ifres11492 Obj

if True == ifres11493 {
ifres11492 = True


} else {
ifres11492 = False


}

ifres11491 = ifres11492


} else {
ifres11491 = False


}

var ifres11490 Obj

if True == ifres11491 {
ifres11490 = True


} else {
ifres11490 = False


}

ifres11489 = ifres11490


} else {
ifres11489 = False


}

if True == ifres11489 {
tmp11339 := PrimHead(V1635)

tmp11340 := PrimTail(tmp11339)

tmp11341 := PrimHead(tmp11340)

tmp11342 := PrimTail(V1635)

tmp11343 := PrimCons(V1634, tmp11342)

tmp11344 := PrimCons(tmp11341, tmp11343)

tmp11345 := PrimCons(symshen_4_7m, tmp11344)

__e.TailApply(PrimFunc(symshen_4compile_1head), V1634, tmp11345, V1636, V1637, V1638)
return


} else {
tmp11487 := PrimIsPair(V1635)

var ifres11467 Obj

if True == tmp11487 {
tmp11485 := PrimHead(V1635)

tmp11486 := PrimIsPair(tmp11485)

var ifres11469 Obj

if True == tmp11486 {
tmp11482 := PrimHead(V1635)

tmp11483 := PrimHead(tmp11482)

tmp11484 := PrimEqual(symshen_4_1m, tmp11483)

var ifres11471 Obj

if True == tmp11484 {
tmp11479 := PrimHead(V1635)

tmp11480 := PrimTail(tmp11479)

tmp11481 := PrimIsPair(tmp11480)

var ifres11473 Obj

if True == tmp11481 {
tmp11475 := PrimHead(V1635)

tmp11476 := PrimTail(tmp11475)

tmp11477 := PrimTail(tmp11476)

tmp11478 := PrimEqual(Nil, tmp11477)

var ifres11474 Obj

if True == tmp11478 {
ifres11474 = True


} else {
ifres11474 = False


}

ifres11473 = ifres11474


} else {
ifres11473 = False


}

var ifres11472 Obj

if True == ifres11473 {
ifres11472 = True


} else {
ifres11472 = False


}

ifres11471 = ifres11472


} else {
ifres11471 = False


}

var ifres11470 Obj

if True == ifres11471 {
ifres11470 = True


} else {
ifres11470 = False


}

ifres11469 = ifres11470


} else {
ifres11469 = False


}

var ifres11468 Obj

if True == ifres11469 {
ifres11468 = True


} else {
ifres11468 = False


}

ifres11467 = ifres11468


} else {
ifres11467 = False


}

if True == ifres11467 {
tmp11346 := PrimHead(V1635)

tmp11347 := PrimTail(tmp11346)

tmp11348 := PrimHead(tmp11347)

tmp11349 := PrimTail(V1635)

tmp11350 := PrimCons(V1634, tmp11349)

tmp11351 := PrimCons(tmp11348, tmp11350)

tmp11352 := PrimCons(symshen_4_1m, tmp11351)

__e.TailApply(PrimFunc(symshen_4compile_1head), V1634, tmp11352, V1636, V1637, V1638)
return


} else {
tmp11465 := PrimIsPair(V1635)

var ifres11461 Obj

if True == tmp11465 {
tmp11463 := PrimHead(V1635)

tmp11464 := PrimEqual(symshen_4_1m, tmp11463)

var ifres11462 Obj

if True == tmp11464 {
ifres11462 = True


} else {
ifres11462 = False


}

ifres11461 = ifres11462


} else {
ifres11461 = False


}

if True == ifres11461 {
tmp11353 := PrimTail(V1635)

__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11353, V1636, V1637, V1638)
return


} else {
tmp11459 := PrimIsPair(V1635)

var ifres11455 Obj

if True == tmp11459 {
tmp11457 := PrimHead(V1635)

tmp11458 := PrimEqual(symshen_4_7m, tmp11457)

var ifres11456 Obj

if True == tmp11458 {
ifres11456 = True


} else {
ifres11456 = False


}

ifres11455 = ifres11456


} else {
ifres11455 = False


}

if True == ifres11455 {
tmp11354 := PrimTail(V1635)

__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11354, V1636, V1637, V1638)
return


} else {
tmp11453 := PrimIsPair(V1635)

var ifres11446 Obj

if True == tmp11453 {
tmp11452 := PrimIsPair(V1636)

var ifres11448 Obj

if True == tmp11452 {
tmp11450 := PrimHead(V1635)

tmp11451 := Call(__e, PrimFunc(symshen_4wildcard_2), tmp11450)


var ifres11449 Obj

if True == tmp11451 {
ifres11449 = True


} else {
ifres11449 = False


}

ifres11448 = ifres11449


} else {
ifres11448 = False


}

var ifres11447 Obj

if True == ifres11448 {
ifres11447 = True


} else {
ifres11447 = False


}

ifres11446 = ifres11447


} else {
ifres11446 = False


}

if True == ifres11446 {
tmp11355 := PrimTail(V1635)

tmp11356 := PrimTail(V1636)

__e.TailApply(PrimFunc(symshen_4compile_1head), V1634, tmp11355, tmp11356, V1637, V1638)
return


} else {
tmp11444 := PrimIsPair(V1635)

var ifres11440 Obj

if True == tmp11444 {
tmp11442 := PrimHead(V1635)

tmp11443 := PrimIsVariable(tmp11442)

var ifres11441 Obj

if True == tmp11443 {
ifres11441 = True


} else {
ifres11441 = False


}

ifres11440 = ifres11441


} else {
ifres11440 = False


}

if True == ifres11440 {
__e.TailApply(PrimFunc(symshen_4variable_1case), V1634, V1635, V1636, V1637, V1638)
return
} else {
tmp11438 := PrimEqual(symshen_4_1m, V1634)

var ifres11431 Obj

if True == tmp11438 {
tmp11437 := PrimIsPair(V1635)

var ifres11433 Obj

if True == tmp11437 {
tmp11435 := PrimHead(V1635)

tmp11436 := Call(__e, PrimFunc(symatom_2), tmp11435)


var ifres11434 Obj

if True == tmp11436 {
ifres11434 = True


} else {
ifres11434 = False


}

ifres11433 = ifres11434


} else {
ifres11433 = False


}

var ifres11432 Obj

if True == ifres11433 {
ifres11432 = True


} else {
ifres11432 = False


}

ifres11431 = ifres11432


} else {
ifres11431 = False


}

if True == ifres11431 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1minus), V1635, V1636, V1637, V1638)
return
} else {
tmp11429 := PrimEqual(symshen_4_1m, V1634)

var ifres11399 Obj

if True == tmp11429 {
tmp11428 := PrimIsPair(V1635)

var ifres11401 Obj

if True == tmp11428 {
tmp11426 := PrimHead(V1635)

tmp11427 := PrimIsPair(tmp11426)

var ifres11403 Obj

if True == tmp11427 {
tmp11423 := PrimHead(V1635)

tmp11424 := PrimHead(tmp11423)

tmp11425 := PrimEqual(symcons, tmp11424)

var ifres11405 Obj

if True == tmp11425 {
tmp11420 := PrimHead(V1635)

tmp11421 := PrimTail(tmp11420)

tmp11422 := PrimIsPair(tmp11421)

var ifres11407 Obj

if True == tmp11422 {
tmp11416 := PrimHead(V1635)

tmp11417 := PrimTail(tmp11416)

tmp11418 := PrimTail(tmp11417)

tmp11419 := PrimIsPair(tmp11418)

var ifres11409 Obj

if True == tmp11419 {
tmp11411 := PrimHead(V1635)

tmp11412 := PrimTail(tmp11411)

tmp11413 := PrimTail(tmp11412)

tmp11414 := PrimTail(tmp11413)

tmp11415 := PrimEqual(Nil, tmp11414)

var ifres11410 Obj

if True == tmp11415 {
ifres11410 = True


} else {
ifres11410 = False


}

ifres11409 = ifres11410


} else {
ifres11409 = False


}

var ifres11408 Obj

if True == ifres11409 {
ifres11408 = True


} else {
ifres11408 = False


}

ifres11407 = ifres11408


} else {
ifres11407 = False


}

var ifres11406 Obj

if True == ifres11407 {
ifres11406 = True


} else {
ifres11406 = False


}

ifres11405 = ifres11406


} else {
ifres11405 = False


}

var ifres11404 Obj

if True == ifres11405 {
ifres11404 = True


} else {
ifres11404 = False


}

ifres11403 = ifres11404


} else {
ifres11403 = False


}

var ifres11402 Obj

if True == ifres11403 {
ifres11402 = True


} else {
ifres11402 = False


}

ifres11401 = ifres11402


} else {
ifres11401 = False


}

var ifres11400 Obj

if True == ifres11401 {
ifres11400 = True


} else {
ifres11400 = False


}

ifres11399 = ifres11400


} else {
ifres11399 = False


}

if True == ifres11399 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1minus), V1635, V1636, V1637, V1638)
return
} else {
tmp11397 := PrimEqual(symshen_4_7m, V1634)

var ifres11390 Obj

if True == tmp11397 {
tmp11396 := PrimIsPair(V1635)

var ifres11392 Obj

if True == tmp11396 {
tmp11394 := PrimHead(V1635)

tmp11395 := Call(__e, PrimFunc(symatom_2), tmp11394)


var ifres11393 Obj

if True == tmp11395 {
ifres11393 = True


} else {
ifres11393 = False


}

ifres11392 = ifres11393


} else {
ifres11392 = False


}

var ifres11391 Obj

if True == ifres11392 {
ifres11391 = True


} else {
ifres11391 = False


}

ifres11390 = ifres11391


} else {
ifres11390 = False


}

if True == ifres11390 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1plus), V1635, V1636, V1637, V1638)
return
} else {
tmp11388 := PrimEqual(symshen_4_7m, V1634)

var ifres11358 Obj

if True == tmp11388 {
tmp11387 := PrimIsPair(V1635)

var ifres11360 Obj

if True == tmp11387 {
tmp11385 := PrimHead(V1635)

tmp11386 := PrimIsPair(tmp11385)

var ifres11362 Obj

if True == tmp11386 {
tmp11382 := PrimHead(V1635)

tmp11383 := PrimHead(tmp11382)

tmp11384 := PrimEqual(symcons, tmp11383)

var ifres11364 Obj

if True == tmp11384 {
tmp11379 := PrimHead(V1635)

tmp11380 := PrimTail(tmp11379)

tmp11381 := PrimIsPair(tmp11380)

var ifres11366 Obj

if True == tmp11381 {
tmp11375 := PrimHead(V1635)

tmp11376 := PrimTail(tmp11375)

tmp11377 := PrimTail(tmp11376)

tmp11378 := PrimIsPair(tmp11377)

var ifres11368 Obj

if True == tmp11378 {
tmp11370 := PrimHead(V1635)

tmp11371 := PrimTail(tmp11370)

tmp11372 := PrimTail(tmp11371)

tmp11373 := PrimTail(tmp11372)

tmp11374 := PrimEqual(Nil, tmp11373)

var ifres11369 Obj

if True == tmp11374 {
ifres11369 = True


} else {
ifres11369 = False


}

ifres11368 = ifres11369


} else {
ifres11368 = False


}

var ifres11367 Obj

if True == ifres11368 {
ifres11367 = True


} else {
ifres11367 = False


}

ifres11366 = ifres11367


} else {
ifres11366 = False


}

var ifres11365 Obj

if True == ifres11366 {
ifres11365 = True


} else {
ifres11365 = False


}

ifres11364 = ifres11365


} else {
ifres11364 = False


}

var ifres11363 Obj

if True == ifres11364 {
ifres11363 = True


} else {
ifres11363 = False


}

ifres11362 = ifres11363


} else {
ifres11362 = False


}

var ifres11361 Obj

if True == ifres11362 {
ifres11361 = True


} else {
ifres11361 = False


}

ifres11360 = ifres11361


} else {
ifres11360 = False


}

var ifres11359 Obj

if True == ifres11360 {
ifres11359 = True


} else {
ifres11359 = False


}

ifres11358 = ifres11359


} else {
ifres11358 = False


}

if True == ifres11358 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1plus), V1635, V1636, V1637, V1638)
return
} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-head")))
return
}


}


}


}


}


}


}


}


}


}


}


}, 5)

tmp11515 := Call(__e, ns2_1set, symshen_4compile_1head, tmp11338)


_ = tmp11515

tmp11516 := MakeNative(func(__e *ControlFlow) {
V1649 := __e.Get(1)
_ = V1649
V1650 := __e.Get(2)
_ = V1650
V1651 := __e.Get(3)
_ = V1651
V1652 := __e.Get(4)
_ = V1652
V1653 := __e.Get(5)
_ = V1653
tmp11537 := PrimIsPair(V1650)

var ifres11534 Obj

if True == tmp11537 {
tmp11536 := PrimIsPair(V1651)

var ifres11535 Obj

if True == tmp11536 {
ifres11535 = True


} else {
ifres11535 = False


}

ifres11534 = ifres11535


} else {
ifres11534 = False


}

if True == ifres11534 {
tmp11531 := PrimHead(V1651)

tmp11532 := PrimIsVariable(tmp11531)

if True == tmp11532 {
tmp11517 := PrimTail(V1650)

tmp11518 := PrimTail(V1651)

tmp11519 := PrimHead(V1651)

tmp11520 := PrimHead(V1650)

tmp11521 := Call(__e, PrimFunc(symsubst), tmp11519, tmp11520, V1653)


__e.TailApply(PrimFunc(symshen_4compile_1head), V1649, tmp11517, tmp11518, V1652, tmp11521)
return


} else {
tmp11522 := PrimHead(V1650)

tmp11523 := PrimHead(V1651)

tmp11524 := PrimTail(V1650)

tmp11525 := PrimTail(V1651)

tmp11526 := Call(__e, PrimFunc(symshen_4compile_1head), V1649, tmp11524, tmp11525, V1652, V1653)


tmp11527 := PrimCons(tmp11526, Nil)

tmp11528 := PrimCons(tmp11523, tmp11527)

tmp11529 := PrimCons(tmp11522, tmp11528)

__e.Return(PrimCons(symlet, tmp11529))
return


}


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.variable-case")))
return
}


}, 5)

tmp11538 := Call(__e, ns2_1set, symshen_4variable_1case, tmp11516)


_ = tmp11538

tmp11539 := MakeNative(func(__e *ControlFlow) {
V1662 := __e.Get(1)
_ = V1662
V1663 := __e.Get(2)
_ = V1663
V1664 := __e.Get(3)
_ = V1664
V1665 := __e.Get(4)
_ = V1665
tmp11564 := PrimIsPair(V1662)

var ifres11561 Obj

if True == tmp11564 {
tmp11563 := PrimIsPair(V1663)

var ifres11562 Obj

if True == tmp11563 {
ifres11562 = True


} else {
ifres11562 = False


}

ifres11561 = ifres11562


} else {
ifres11561 = False


}

if True == ifres11561 {
tmp11540 := MakeNative(func(__e *ControlFlow) {
W1666 := __e.Get(1)
_ = W1666
tmp11541 := PrimHead(V1663)

tmp11542 := PrimCons(V1664, Nil)

tmp11543 := PrimCons(tmp11541, tmp11542)

tmp11544 := PrimCons(symshen_4lazyderef, tmp11543)

tmp11545 := PrimHead(V1662)

tmp11546 := PrimCons(tmp11545, Nil)

tmp11547 := PrimCons(W1666, tmp11546)

tmp11548 := PrimCons(sym_a, tmp11547)

tmp11549 := PrimTail(V1662)

tmp11550 := PrimTail(V1663)

tmp11551 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11549, tmp11550, V1664, V1665)


tmp11552 := PrimCons(False, Nil)

tmp11553 := PrimCons(tmp11551, tmp11552)

tmp11554 := PrimCons(tmp11548, tmp11553)

tmp11555 := PrimCons(symif, tmp11554)

tmp11556 := PrimCons(tmp11555, Nil)

tmp11557 := PrimCons(tmp11544, tmp11556)

tmp11558 := PrimCons(W1666, tmp11557)

__e.Return(PrimCons(symlet, tmp11558))
return


}, 1)

tmp11559 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11540, tmp11559)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.atom-case-minus")))
return
}


}, 4)

tmp11565 := Call(__e, ns2_1set, symshen_4atom_1case_1minus, tmp11539)


_ = tmp11565

tmp11566 := MakeNative(func(__e *ControlFlow) {
V1675 := __e.Get(1)
_ = V1675
V1676 := __e.Get(2)
_ = V1676
V1677 := __e.Get(3)
_ = V1677
V1678 := __e.Get(4)
_ = V1678
tmp11631 := PrimIsPair(V1675)

var ifres11601 Obj

if True == tmp11631 {
tmp11629 := PrimHead(V1675)

tmp11630 := PrimIsPair(tmp11629)

var ifres11603 Obj

if True == tmp11630 {
tmp11626 := PrimHead(V1675)

tmp11627 := PrimHead(tmp11626)

tmp11628 := PrimEqual(symcons, tmp11627)

var ifres11605 Obj

if True == tmp11628 {
tmp11623 := PrimHead(V1675)

tmp11624 := PrimTail(tmp11623)

tmp11625 := PrimIsPair(tmp11624)

var ifres11607 Obj

if True == tmp11625 {
tmp11619 := PrimHead(V1675)

tmp11620 := PrimTail(tmp11619)

tmp11621 := PrimTail(tmp11620)

tmp11622 := PrimIsPair(tmp11621)

var ifres11609 Obj

if True == tmp11622 {
tmp11614 := PrimHead(V1675)

tmp11615 := PrimTail(tmp11614)

tmp11616 := PrimTail(tmp11615)

tmp11617 := PrimTail(tmp11616)

tmp11618 := PrimEqual(Nil, tmp11617)

var ifres11611 Obj

if True == tmp11618 {
tmp11613 := PrimIsPair(V1676)

var ifres11612 Obj

if True == tmp11613 {
ifres11612 = True


} else {
ifres11612 = False


}

ifres11611 = ifres11612


} else {
ifres11611 = False


}

var ifres11610 Obj

if True == ifres11611 {
ifres11610 = True


} else {
ifres11610 = False


}

ifres11609 = ifres11610


} else {
ifres11609 = False


}

var ifres11608 Obj

if True == ifres11609 {
ifres11608 = True


} else {
ifres11608 = False


}

ifres11607 = ifres11608


} else {
ifres11607 = False


}

var ifres11606 Obj

if True == ifres11607 {
ifres11606 = True


} else {
ifres11606 = False


}

ifres11605 = ifres11606


} else {
ifres11605 = False


}

var ifres11604 Obj

if True == ifres11605 {
ifres11604 = True


} else {
ifres11604 = False


}

ifres11603 = ifres11604


} else {
ifres11603 = False


}

var ifres11602 Obj

if True == ifres11603 {
ifres11602 = True


} else {
ifres11602 = False


}

ifres11601 = ifres11602


} else {
ifres11601 = False


}

if True == ifres11601 {
tmp11567 := MakeNative(func(__e *ControlFlow) {
W1679 := __e.Get(1)
_ = W1679
tmp11568 := PrimHead(V1676)

tmp11569 := PrimCons(V1677, Nil)

tmp11570 := PrimCons(tmp11568, tmp11569)

tmp11571 := PrimCons(symshen_4lazyderef, tmp11570)

tmp11572 := PrimCons(W1679, Nil)

tmp11573 := PrimCons(symcons_2, tmp11572)

tmp11574 := PrimHead(V1675)

tmp11575 := PrimTail(tmp11574)

tmp11576 := PrimHead(tmp11575)

tmp11577 := PrimHead(V1675)

tmp11578 := PrimTail(tmp11577)

tmp11579 := PrimTail(tmp11578)

tmp11580 := PrimHead(tmp11579)

tmp11581 := PrimTail(V1675)

tmp11582 := PrimCons(tmp11580, tmp11581)

tmp11583 := PrimCons(tmp11576, tmp11582)

tmp11584 := PrimCons(W1679, Nil)

tmp11585 := PrimCons(symhd, tmp11584)

tmp11586 := PrimCons(W1679, Nil)

tmp11587 := PrimCons(symtl, tmp11586)

tmp11588 := PrimTail(V1676)

tmp11589 := PrimCons(tmp11587, tmp11588)

tmp11590 := PrimCons(tmp11585, tmp11589)

tmp11591 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11583, tmp11590, V1677, V1678)


tmp11592 := PrimCons(False, Nil)

tmp11593 := PrimCons(tmp11591, tmp11592)

tmp11594 := PrimCons(tmp11573, tmp11593)

tmp11595 := PrimCons(symif, tmp11594)

tmp11596 := PrimCons(tmp11595, Nil)

tmp11597 := PrimCons(tmp11571, tmp11596)

tmp11598 := PrimCons(W1679, tmp11597)

__e.Return(PrimCons(symlet, tmp11598))
return


}, 1)

tmp11599 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11567, tmp11599)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.cons-case-minus")))
return
}


}, 4)

tmp11632 := Call(__e, ns2_1set, symshen_4cons_1case_1minus, tmp11566)


_ = tmp11632

tmp11633 := MakeNative(func(__e *ControlFlow) {
V1688 := __e.Get(1)
_ = V1688
V1689 := __e.Get(2)
_ = V1689
V1690 := __e.Get(3)
_ = V1690
V1691 := __e.Get(4)
_ = V1691
tmp11679 := PrimIsPair(V1688)

var ifres11676 Obj

if True == tmp11679 {
tmp11678 := PrimIsPair(V1689)

var ifres11677 Obj

if True == tmp11678 {
ifres11677 = True


} else {
ifres11677 = False


}

ifres11676 = ifres11677


} else {
ifres11676 = False


}

if True == ifres11676 {
tmp11634 := MakeNative(func(__e *ControlFlow) {
W1692 := __e.Get(1)
_ = W1692
tmp11635 := MakeNative(func(__e *ControlFlow) {
W1693 := __e.Get(1)
_ = W1693
tmp11636 := PrimHead(V1689)

tmp11637 := PrimCons(V1690, Nil)

tmp11638 := PrimCons(tmp11636, tmp11637)

tmp11639 := PrimCons(symshen_4lazyderef, tmp11638)

tmp11640 := PrimTail(V1688)

tmp11641 := PrimTail(V1689)

tmp11642 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11640, tmp11641, V1690, V1691)


tmp11643 := PrimCons(tmp11642, Nil)

tmp11644 := PrimCons(symfreeze, tmp11643)

tmp11645 := PrimHead(V1688)

tmp11646 := PrimCons(tmp11645, Nil)

tmp11647 := PrimCons(W1692, tmp11646)

tmp11648 := PrimCons(sym_a, tmp11647)

tmp11649 := PrimCons(W1693, Nil)

tmp11650 := PrimCons(symthaw, tmp11649)

tmp11651 := PrimCons(W1692, Nil)

tmp11652 := PrimCons(symshen_4pvar_2, tmp11651)

tmp11653 := PrimHead(V1688)

tmp11654 := Call(__e, PrimFunc(symshen_4demode), tmp11653)


tmp11655 := PrimCons(W1693, Nil)

tmp11656 := PrimCons(V1690, tmp11655)

tmp11657 := PrimCons(tmp11654, tmp11656)

tmp11658 := PrimCons(W1692, tmp11657)

tmp11659 := PrimCons(symshen_4bind_b, tmp11658)

tmp11660 := PrimCons(False, Nil)

tmp11661 := PrimCons(tmp11659, tmp11660)

tmp11662 := PrimCons(tmp11652, tmp11661)

tmp11663 := PrimCons(symif, tmp11662)

tmp11664 := PrimCons(tmp11663, Nil)

tmp11665 := PrimCons(tmp11650, tmp11664)

tmp11666 := PrimCons(tmp11648, tmp11665)

tmp11667 := PrimCons(symif, tmp11666)

tmp11668 := PrimCons(tmp11667, Nil)

tmp11669 := PrimCons(tmp11644, tmp11668)

tmp11670 := PrimCons(W1693, tmp11669)

tmp11671 := PrimCons(tmp11639, tmp11670)

tmp11672 := PrimCons(W1692, tmp11671)

__e.Return(PrimCons(symlet, tmp11672))
return


}, 1)

tmp11673 := Call(__e, PrimFunc(symgensym), symGoTo)


__e.TailApply(tmp11635, tmp11673)
return


}, 1)

tmp11674 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11634, tmp11674)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.atom-case-plus")))
return
}


}, 4)

tmp11680 := Call(__e, ns2_1set, symshen_4atom_1case_1plus, tmp11633)


_ = tmp11680

tmp11681 := MakeNative(func(__e *ControlFlow) {
V1702 := __e.Get(1)
_ = V1702
V1703 := __e.Get(2)
_ = V1703
V1704 := __e.Get(3)
_ = V1704
V1705 := __e.Get(4)
_ = V1705
tmp11777 := PrimIsPair(V1702)

var ifres11747 Obj

if True == tmp11777 {
tmp11775 := PrimHead(V1702)

tmp11776 := PrimIsPair(tmp11775)

var ifres11749 Obj

if True == tmp11776 {
tmp11772 := PrimHead(V1702)

tmp11773 := PrimHead(tmp11772)

tmp11774 := PrimEqual(symcons, tmp11773)

var ifres11751 Obj

if True == tmp11774 {
tmp11769 := PrimHead(V1702)

tmp11770 := PrimTail(tmp11769)

tmp11771 := PrimIsPair(tmp11770)

var ifres11753 Obj

if True == tmp11771 {
tmp11765 := PrimHead(V1702)

tmp11766 := PrimTail(tmp11765)

tmp11767 := PrimTail(tmp11766)

tmp11768 := PrimIsPair(tmp11767)

var ifres11755 Obj

if True == tmp11768 {
tmp11760 := PrimHead(V1702)

tmp11761 := PrimTail(tmp11760)

tmp11762 := PrimTail(tmp11761)

tmp11763 := PrimTail(tmp11762)

tmp11764 := PrimEqual(Nil, tmp11763)

var ifres11757 Obj

if True == tmp11764 {
tmp11759 := PrimIsPair(V1703)

var ifres11758 Obj

if True == tmp11759 {
ifres11758 = True


} else {
ifres11758 = False


}

ifres11757 = ifres11758


} else {
ifres11757 = False


}

var ifres11756 Obj

if True == ifres11757 {
ifres11756 = True


} else {
ifres11756 = False


}

ifres11755 = ifres11756


} else {
ifres11755 = False


}

var ifres11754 Obj

if True == ifres11755 {
ifres11754 = True


} else {
ifres11754 = False


}

ifres11753 = ifres11754


} else {
ifres11753 = False


}

var ifres11752 Obj

if True == ifres11753 {
ifres11752 = True


} else {
ifres11752 = False


}

ifres11751 = ifres11752


} else {
ifres11751 = False


}

var ifres11750 Obj

if True == ifres11751 {
ifres11750 = True


} else {
ifres11750 = False


}

ifres11749 = ifres11750


} else {
ifres11749 = False


}

var ifres11748 Obj

if True == ifres11749 {
ifres11748 = True


} else {
ifres11748 = False


}

ifres11747 = ifres11748


} else {
ifres11747 = False


}

if True == ifres11747 {
tmp11682 := MakeNative(func(__e *ControlFlow) {
W1706 := __e.Get(1)
_ = W1706
tmp11683 := MakeNative(func(__e *ControlFlow) {
W1707 := __e.Get(1)
_ = W1707
tmp11684 := MakeNative(func(__e *ControlFlow) {
W1708 := __e.Get(1)
_ = W1708
tmp11685 := MakeNative(func(__e *ControlFlow) {
W1709 := __e.Get(1)
_ = W1709
tmp11686 := MakeNative(func(__e *ControlFlow) {
W1710 := __e.Get(1)
_ = W1710
tmp11687 := PrimHead(V1703)

tmp11688 := PrimCons(V1704, Nil)

tmp11689 := PrimCons(tmp11687, tmp11688)

tmp11690 := PrimCons(symshen_4lazyderef, tmp11689)

tmp11691 := PrimTail(V1702)

tmp11692 := PrimTail(V1703)

tmp11693 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11691, tmp11692, V1704, V1705)


tmp11694 := Call(__e, PrimFunc(symshen_4goto), W1708, tmp11693)


tmp11695 := PrimCons(W1706, Nil)

tmp11696 := PrimCons(symcons_2, tmp11695)

tmp11697 := PrimHead(V1702)

tmp11698 := PrimTail(tmp11697)

tmp11699 := PrimCons(W1706, Nil)

tmp11700 := PrimCons(symhd, tmp11699)

tmp11701 := PrimCons(W1706, Nil)

tmp11702 := PrimCons(symtl, tmp11701)

tmp11703 := PrimCons(tmp11702, Nil)

tmp11704 := PrimCons(tmp11700, tmp11703)

tmp11705 := Call(__e, PrimFunc(symshen_4invoke), W1707, W1708)


tmp11706 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11698, tmp11704, V1704, tmp11705)


tmp11707 := PrimCons(W1706, Nil)

tmp11708 := PrimCons(symshen_4pvar_2, tmp11707)

tmp11709 := Call(__e, PrimFunc(symshen_4demode), W1709)


tmp11710 := Call(__e, PrimFunc(symshen_4invoke), W1707, W1708)


tmp11711 := PrimCons(tmp11710, Nil)

tmp11712 := PrimCons(symfreeze, tmp11711)

tmp11713 := PrimCons(tmp11712, Nil)

tmp11714 := PrimCons(V1704, tmp11713)

tmp11715 := PrimCons(tmp11709, tmp11714)

tmp11716 := PrimCons(W1706, tmp11715)

tmp11717 := PrimCons(symshen_4bind_b, tmp11716)

tmp11718 := Call(__e, PrimFunc(symshen_4stpart), W1710, tmp11717, V1704)


tmp11719 := PrimCons(False, Nil)

tmp11720 := PrimCons(tmp11718, tmp11719)

tmp11721 := PrimCons(tmp11708, tmp11720)

tmp11722 := PrimCons(symif, tmp11721)

tmp11723 := PrimCons(tmp11722, Nil)

tmp11724 := PrimCons(tmp11706, tmp11723)

tmp11725 := PrimCons(tmp11696, tmp11724)

tmp11726 := PrimCons(symif, tmp11725)

tmp11727 := PrimCons(tmp11726, Nil)

tmp11728 := PrimCons(tmp11694, tmp11727)

tmp11729 := PrimCons(W1707, tmp11728)

tmp11730 := PrimCons(tmp11690, tmp11729)

tmp11731 := PrimCons(W1706, tmp11730)

__e.Return(PrimCons(symlet, tmp11731))
return


}, 1)

tmp11732 := Call(__e, PrimFunc(symshen_4extract_1vars), W1709)


__e.TailApply(tmp11686, tmp11732)
return


}, 1)

tmp11733 := PrimHead(V1702)

tmp11734 := Call(__e, PrimFunc(symshen_4tame), tmp11733)


__e.TailApply(tmp11685, tmp11734)
return


}, 1)

tmp11735 := PrimHead(V1702)

tmp11736 := PrimTail(tmp11735)

tmp11737 := PrimHead(tmp11736)

tmp11738 := PrimHead(V1702)

tmp11739 := PrimTail(tmp11738)

tmp11740 := PrimTail(tmp11739)

tmp11741 := PrimHead(tmp11740)

tmp11742 := PrimCons(tmp11737, tmp11741)

tmp11743 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp11742)


__e.TailApply(tmp11684, tmp11743)
return


}, 1)

tmp11744 := Call(__e, PrimFunc(symgensym), symGoTo)


__e.TailApply(tmp11683, tmp11744)
return


}, 1)

tmp11745 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11682, tmp11745)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.cons-case-plus")))
return
}


}, 4)

tmp11778 := Call(__e, ns2_1set, symshen_4cons_1case_1plus, tmp11681)


_ = tmp11778

tmp11779 := MakeNative(func(__e *ControlFlow) {
V1711 := __e.Get(1)
_ = V1711
tmp11816 := PrimIsPair(V1711)

var ifres11803 Obj

if True == tmp11816 {
tmp11814 := PrimHead(V1711)

tmp11815 := PrimEqual(symshen_4_7m, tmp11814)

var ifres11805 Obj

if True == tmp11815 {
tmp11812 := PrimTail(V1711)

tmp11813 := PrimIsPair(tmp11812)

var ifres11807 Obj

if True == tmp11813 {
tmp11809 := PrimTail(V1711)

tmp11810 := PrimTail(tmp11809)

tmp11811 := PrimEqual(Nil, tmp11810)

var ifres11808 Obj

if True == tmp11811 {
ifres11808 = True


} else {
ifres11808 = False


}

ifres11807 = ifres11808


} else {
ifres11807 = False


}

var ifres11806 Obj

if True == ifres11807 {
ifres11806 = True


} else {
ifres11806 = False


}

ifres11805 = ifres11806


} else {
ifres11805 = False


}

var ifres11804 Obj

if True == ifres11805 {
ifres11804 = True


} else {
ifres11804 = False


}

ifres11803 = ifres11804


} else {
ifres11803 = False


}

if True == ifres11803 {
tmp11780 := PrimTail(V1711)

tmp11781 := PrimHead(tmp11780)

__e.TailApply(PrimFunc(symshen_4demode), tmp11781)
return


} else {
tmp11801 := PrimIsPair(V1711)

var ifres11788 Obj

if True == tmp11801 {
tmp11799 := PrimHead(V1711)

tmp11800 := PrimEqual(symshen_4_1m, tmp11799)

var ifres11790 Obj

if True == tmp11800 {
tmp11797 := PrimTail(V1711)

tmp11798 := PrimIsPair(tmp11797)

var ifres11792 Obj

if True == tmp11798 {
tmp11794 := PrimTail(V1711)

tmp11795 := PrimTail(tmp11794)

tmp11796 := PrimEqual(Nil, tmp11795)

var ifres11793 Obj

if True == tmp11796 {
ifres11793 = True


} else {
ifres11793 = False


}

ifres11792 = ifres11793


} else {
ifres11792 = False


}

var ifres11791 Obj

if True == ifres11792 {
ifres11791 = True


} else {
ifres11791 = False


}

ifres11790 = ifres11791


} else {
ifres11790 = False


}

var ifres11789 Obj

if True == ifres11790 {
ifres11789 = True


} else {
ifres11789 = False


}

ifres11788 = ifres11789


} else {
ifres11788 = False


}

if True == ifres11788 {
tmp11782 := PrimTail(V1711)

tmp11783 := PrimHead(tmp11782)

__e.TailApply(PrimFunc(symshen_4demode), tmp11783)
return


} else {
tmp11786 := PrimIsPair(V1711)

if True == tmp11786 {
tmp11784 := MakeNative(func(__e *ControlFlow) {
Z1712 := __e.Get(1)
_ = Z1712
__e.TailApply(PrimFunc(symshen_4demode), Z1712)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11784, V1711)
return


} else {
__e.Return(V1711)
return
}


}


}


}, 1)

tmp11817 := Call(__e, ns2_1set, symshen_4demode, tmp11779)


_ = tmp11817

tmp11818 := MakeNative(func(__e *ControlFlow) {
V1713 := __e.Get(1)
_ = V1713
tmp11823 := Call(__e, PrimFunc(symshen_4wildcard_2), V1713)


if True == tmp11823 {
__e.TailApply(PrimFunc(symgensym), symY)
return
} else {
tmp11821 := PrimIsPair(V1713)

if True == tmp11821 {
tmp11819 := MakeNative(func(__e *ControlFlow) {
Z1714 := __e.Get(1)
_ = Z1714
__e.TailApply(PrimFunc(symshen_4tame), Z1714)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11819, V1713)
return


} else {
__e.Return(V1713)
return
}


}


}, 1)

tmp11824 := Call(__e, ns2_1set, symshen_4tame, tmp11818)


_ = tmp11824

tmp11825 := MakeNative(func(__e *ControlFlow) {
V1715 := __e.Get(1)
_ = V1715
V1716 := __e.Get(2)
_ = V1716
tmp11828 := PrimEqual(Nil, V1715)

if True == tmp11828 {
tmp11826 := PrimCons(V1716, Nil)

__e.Return(PrimCons(symfreeze, tmp11826))
return


} else {
__e.TailApply(PrimFunc(symshen_4goto_1h), V1715, V1716)
return
}


}, 2)

tmp11829 := Call(__e, ns2_1set, symshen_4goto, tmp11825)


_ = tmp11829

tmp11830 := MakeNative(func(__e *ControlFlow) {
V1717 := __e.Get(1)
_ = V1717
V1718 := __e.Get(2)
_ = V1718
tmp11839 := PrimEqual(Nil, V1717)

if True == tmp11839 {
__e.Return(V1718)
return
} else {
tmp11837 := PrimIsPair(V1717)

if True == tmp11837 {
tmp11831 := PrimHead(V1717)

tmp11832 := PrimTail(V1717)

tmp11833 := Call(__e, PrimFunc(symshen_4goto_1h), tmp11832, V1718)


tmp11834 := PrimCons(tmp11833, Nil)

tmp11835 := PrimCons(tmp11831, tmp11834)

__e.Return(PrimCons(symlambda, tmp11835))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.goto-h")))
return
}


}


}, 2)

tmp11840 := Call(__e, ns2_1set, symshen_4goto_1h, tmp11830)


_ = tmp11840

tmp11841 := MakeNative(func(__e *ControlFlow) {
V1719 := __e.Get(1)
_ = V1719
V1720 := __e.Get(2)
_ = V1720
tmp11844 := PrimEqual(Nil, V1720)

if True == tmp11844 {
tmp11842 := PrimCons(V1719, Nil)

__e.Return(PrimCons(symthaw, tmp11842))
return


} else {
__e.Return(PrimCons(V1719, V1720))
return
}


}, 2)

tmp11845 := Call(__e, ns2_1set, symshen_4invoke, tmp11841)


_ = tmp11845

tmp11846 := MakeNative(func(__e *ControlFlow) {
V1721 := __e.Get(1)
_ = V1721
__e.Return(PrimEqual(V1721, sym__))
return
}, 1)

tmp11847 := Call(__e, ns2_1set, symshen_4wildcard_2, tmp11846)


_ = tmp11847

tmp11848 := MakeNative(func(__e *ControlFlow) {
V1722 := __e.Get(1)
_ = V1722
tmp11849 := MakeNative(func(__e *ControlFlow) {
tmp11854 := PrimIsVector(V1722)

if True == tmp11854 {
tmp11851 := PrimVectorGet(V1722, MakeNumber(0))

tmp11852 := PrimEqual(tmp11851, symshen_4pvar)

if True == tmp11852 {
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


}, 0)

tmp11855 := MakeNative(func(__e *ControlFlow) {
Z1723 := __e.Get(1)
_ = Z1723
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp11849, tmp11855)
return


}, 1)

tmp11856 := Call(__e, ns2_1set, symshen_4pvar_2, tmp11848)


_ = tmp11856

tmp11857 := MakeNative(func(__e *ControlFlow) {
V1724 := __e.Get(1)
_ = V1724
V1725 := __e.Get(2)
_ = V1725
tmp11864 := Call(__e, PrimFunc(symshen_4pvar_2), V1724)


if True == tmp11864 {
tmp11858 := MakeNative(func(__e *ControlFlow) {
W1726 := __e.Get(1)
_ = W1726
tmp11860 := PrimEqual(W1726, symshen_4_1null_1)

if True == tmp11860 {
__e.Return(V1724)
return
} else {
__e.TailApply(PrimFunc(symshen_4lazyderef), W1726, V1725)
return
}


}, 1)

tmp11861 := PrimVectorGet(V1724, MakeNumber(1))

tmp11862 := PrimVectorGet(V1725, tmp11861)

__e.TailApply(tmp11858, tmp11862)
return


} else {
__e.Return(V1724)
return
}


}, 2)

tmp11865 := Call(__e, ns2_1set, symshen_4lazyderef, tmp11857)


_ = tmp11865

tmp11866 := MakeNative(func(__e *ControlFlow) {
V1727 := __e.Get(1)
_ = V1727
V1728 := __e.Get(2)
_ = V1728
tmp11879 := PrimIsPair(V1727)

if True == tmp11879 {
tmp11867 := PrimHead(V1727)

tmp11868 := Call(__e, PrimFunc(symshen_4deref), tmp11867, V1728)


tmp11869 := PrimTail(V1727)

tmp11870 := Call(__e, PrimFunc(symshen_4deref), tmp11869, V1728)


__e.Return(PrimCons(tmp11868, tmp11870))
return


} else {
tmp11877 := Call(__e, PrimFunc(symshen_4pvar_2), V1727)


if True == tmp11877 {
tmp11871 := MakeNative(func(__e *ControlFlow) {
W1729 := __e.Get(1)
_ = W1729
tmp11873 := PrimEqual(W1729, symshen_4_1null_1)

if True == tmp11873 {
__e.Return(V1727)
return
} else {
__e.TailApply(PrimFunc(symshen_4deref), W1729, V1728)
return
}


}, 1)

tmp11874 := PrimVectorGet(V1727, MakeNumber(1))

tmp11875 := PrimVectorGet(V1728, tmp11874)

__e.TailApply(tmp11871, tmp11875)
return


} else {
__e.Return(V1727)
return
}


}


}, 2)

tmp11880 := Call(__e, ns2_1set, symshen_4deref, tmp11866)


_ = tmp11880

tmp11881 := MakeNative(func(__e *ControlFlow) {
V1730 := __e.Get(1)
_ = V1730
V1731 := __e.Get(2)
_ = V1731
V1732 := __e.Get(3)
_ = V1732
V1733 := __e.Get(4)
_ = V1733
tmp11882 := MakeNative(func(__e *ControlFlow) {
W1734 := __e.Get(1)
_ = W1734
tmp11883 := MakeNative(func(__e *ControlFlow) {
W1735 := __e.Get(1)
_ = W1735
tmp11885 := PrimEqual(W1735, False)

if True == tmp11885 {
__e.TailApply(PrimFunc(symshen_4unwind), V1730, V1732, W1735)
return
} else {
__e.Return(W1735)
return
}


}, 1)

tmp11886 := Call(__e, PrimFunc(symthaw), V1733)


__e.TailApply(tmp11883, tmp11886)
return


}, 1)

tmp11887 := Call(__e, PrimFunc(symshen_4bindv), V1730, V1731, V1732)


__e.TailApply(tmp11882, tmp11887)
return


}, 4)

tmp11888 := Call(__e, ns2_1set, symshen_4bind_b, tmp11881)


_ = tmp11888

tmp11889 := MakeNative(func(__e *ControlFlow) {
V1736 := __e.Get(1)
_ = V1736
V1737 := __e.Get(2)
_ = V1737
V1738 := __e.Get(3)
_ = V1738
tmp11890 := PrimVectorGet(V1736, MakeNumber(1))

__e.Return(PrimVectorSet(V1738, tmp11890, V1737))
return


}, 3)

tmp11891 := Call(__e, ns2_1set, symshen_4bindv, tmp11889)


_ = tmp11891

tmp11892 := MakeNative(func(__e *ControlFlow) {
V1739 := __e.Get(1)
_ = V1739
V1740 := __e.Get(2)
_ = V1740
V1741 := __e.Get(3)
_ = V1741
tmp11893 := PrimVectorGet(V1739, MakeNumber(1))

tmp11894 := PrimVectorSet(V1740, tmp11893, symshen_4_1null_1)

_ = tmp11894

__e.Return(V1741)
return


}, 3)

tmp11895 := Call(__e, ns2_1set, symshen_4unwind, tmp11892)


_ = tmp11895

tmp11896 := MakeNative(func(__e *ControlFlow) {
V1750 := __e.Get(1)
_ = V1750
V1751 := __e.Get(2)
_ = V1751
V1752 := __e.Get(3)
_ = V1752
tmp11911 := PrimEqual(Nil, V1750)

if True == tmp11911 {
__e.Return(V1751)
return
} else {
tmp11909 := PrimIsPair(V1750)

if True == tmp11909 {
tmp11897 := PrimHead(V1750)

tmp11898 := PrimCons(V1752, Nil)

tmp11899 := PrimCons(symshen_4newpv, tmp11898)

tmp11900 := PrimTail(V1750)

tmp11901 := Call(__e, PrimFunc(symshen_4stpart), tmp11900, V1751, V1752)


tmp11902 := PrimCons(tmp11901, Nil)

tmp11903 := PrimCons(V1752, tmp11902)

tmp11904 := PrimCons(symshen_4gc, tmp11903)

tmp11905 := PrimCons(tmp11904, Nil)

tmp11906 := PrimCons(tmp11899, tmp11905)

tmp11907 := PrimCons(tmp11897, tmp11906)

__e.Return(PrimCons(symlet, tmp11907))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.stpart")))
return
}


}


}, 3)

tmp11912 := Call(__e, ns2_1set, symshen_4stpart, tmp11896)


_ = tmp11912

tmp11913 := MakeNative(func(__e *ControlFlow) {
V1753 := __e.Get(1)
_ = V1753
V1754 := __e.Get(2)
_ = V1754
tmp11918 := PrimEqual(V1754, False)

if True == tmp11918 {
tmp11914 := MakeNative(func(__e *ControlFlow) {
W1755 := __e.Get(1)
_ = W1755
tmp11915 := Call(__e, PrimFunc(symshen_4decrement_1ticket), W1755, V1753)


_ = tmp11915

__e.Return(V1754)
return


}, 1)

tmp11916 := Call(__e, PrimFunc(symshen_4ticket_1number), V1753)


__e.TailApply(tmp11914, tmp11916)
return


} else {
__e.Return(V1754)
return
}


}, 2)

tmp11919 := Call(__e, ns2_1set, symshen_4gc, tmp11913)


_ = tmp11919

tmp11920 := MakeNative(func(__e *ControlFlow) {
V1756 := __e.Get(1)
_ = V1756
V1757 := __e.Get(2)
_ = V1757
tmp11921 := PrimNumberSubtract(V1756, MakeNumber(1))

__e.Return(PrimVectorSet(V1757, MakeNumber(1), tmp11921))
return


}, 2)

tmp11922 := Call(__e, ns2_1set, symshen_4decrement_1ticket, tmp11920)


_ = tmp11922

tmp11923 := MakeNative(func(__e *ControlFlow) {
V1758 := __e.Get(1)
_ = V1758
tmp11924 := MakeNative(func(__e *ControlFlow) {
W1759 := __e.Get(1)
_ = W1759
tmp11925 := MakeNative(func(__e *ControlFlow) {
W1760 := __e.Get(1)
_ = W1760
tmp11926 := MakeNative(func(__e *ControlFlow) {
W1761 := __e.Get(1)
_ = W1761
__e.Return(W1760)
return
}, 1)

tmp11927 := Call(__e, PrimFunc(symshen_4nextticket), V1758, W1759)


__e.TailApply(tmp11926, tmp11927)
return


}, 1)

tmp11928 := Call(__e, PrimFunc(symshen_4make_1prolog_1variable), W1759)


__e.TailApply(tmp11925, tmp11928)
return


}, 1)

tmp11929 := Call(__e, PrimFunc(symshen_4ticket_1number), V1758)


__e.TailApply(tmp11924, tmp11929)
return


}, 1)

tmp11930 := Call(__e, ns2_1set, symshen_4newpv, tmp11923)


_ = tmp11930

tmp11931 := MakeNative(func(__e *ControlFlow) {
V1762 := __e.Get(1)
_ = V1762
__e.Return(PrimVectorGet(V1762, MakeNumber(1)))
return
}, 1)

tmp11932 := Call(__e, ns2_1set, symshen_4ticket_1number, tmp11931)


_ = tmp11932

tmp11933 := MakeNative(func(__e *ControlFlow) {
V1763 := __e.Get(1)
_ = V1763
V1764 := __e.Get(2)
_ = V1764
tmp11934 := MakeNative(func(__e *ControlFlow) {
W1765 := __e.Get(1)
_ = W1765
tmp11935 := PrimNumberAdd(V1764, MakeNumber(1))

__e.Return(PrimVectorSet(W1765, MakeNumber(1), tmp11935))
return


}, 1)

tmp11936 := PrimVectorSet(V1763, V1764, symshen_4_1null_1)

__e.TailApply(tmp11934, tmp11936)
return


}, 2)

tmp11937 := Call(__e, ns2_1set, symshen_4nextticket, tmp11933)


_ = tmp11937

tmp11938 := MakeNative(func(__e *ControlFlow) {
V1766 := __e.Get(1)
_ = V1766
tmp11939 := PrimAbsvector(MakeNumber(2))

tmp11940 := PrimVectorSet(tmp11939, MakeNumber(0), symshen_4pvar)

__e.Return(PrimVectorSet(tmp11940, MakeNumber(1), V1766))
return


}, 1)

tmp11941 := Call(__e, ns2_1set, symshen_4make_1prolog_1variable, tmp11938)


_ = tmp11941

tmp11942 := MakeNative(func(__e *ControlFlow) {
V1767 := __e.Get(1)
_ = V1767
tmp11943 := PrimVectorGet(V1767, MakeNumber(1))

tmp11944 := Call(__e, PrimFunc(symshen_4app), tmp11943, MakeString(""), symshen_4a)


__e.Return(PrimStringConcat(MakeString("Var"), tmp11944))
return


}, 1)

tmp11945 := Call(__e, ns2_1set, symshen_4pvar, tmp11942)


_ = tmp11945

tmp11946 := MakeNative(func(__e *ControlFlow) {
tmp11947 := PrimValue(symshen_4_dinfs_d)

tmp11948 := PrimNumberAdd(MakeNumber(1), tmp11947)

__e.Return(PrimSet(symshen_4_dinfs_d, tmp11948))
return


}, 0)

tmp11949 := Call(__e, ns2_1set, symshen_4incinfs, tmp11946)


_ = tmp11949

tmp11950 := MakeNative(func(__e *ControlFlow) {
V1768 := __e.Get(1)
_ = V1768
tmp11957 := PrimIsInteger(V1768)

var ifres11954 Obj

if True == tmp11957 {
tmp11956 := PrimGreatThan(V1768, MakeNumber(0))

var ifres11955 Obj

if True == tmp11956 {
ifres11955 = True


} else {
ifres11955 = False


}

ifres11954 = ifres11955


} else {
ifres11954 = False


}

if True == ifres11954 {
__e.Return(PrimSet(symshen_4_dsize_1prolog_1vector_d, V1768))
return
} else {
tmp11951 := Call(__e, PrimFunc(symshen_4app), V1768, MakeString(""), symshen_4a)


tmp11952 := PrimStringConcat(MakeString("prolog vector size: size should be a positive integer; not "), tmp11951)

__e.Return(PrimSimpleError(tmp11952))
return


}


}, 1)

tmp11958 := Call(__e, ns2_1set, symshen_4prolog_1vector_1size, tmp11950)


_ = tmp11958

tmp11959 := MakeNative(func(__e *ControlFlow) {
V1780 := __e.Get(1)
_ = V1780
V1781 := __e.Get(2)
_ = V1781
V1782 := __e.Get(3)
_ = V1782
V1783 := __e.Get(4)
_ = V1783
tmp11989 := PrimEqual(V1780, V1781)

if True == tmp11989 {
__e.TailApply(PrimFunc(symthaw), V1783)
return
} else {
tmp11987 := Call(__e, PrimFunc(symshen_4pvar_2), V1780)


var ifres11982 Obj

if True == tmp11987 {
tmp11984 := Call(__e, PrimFunc(symshen_4deref), V1781, V1782)


tmp11985 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1780, tmp11984)


tmp11986 := PrimNot(tmp11985)

var ifres11983 Obj

if True == tmp11986 {
ifres11983 = True


} else {
ifres11983 = False


}

ifres11982 = ifres11983


} else {
ifres11982 = False


}

if True == ifres11982 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1780, V1781, V1782, V1783)
return
} else {
tmp11980 := Call(__e, PrimFunc(symshen_4pvar_2), V1781)


var ifres11975 Obj

if True == tmp11980 {
tmp11977 := Call(__e, PrimFunc(symshen_4deref), V1780, V1782)


tmp11978 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1781, tmp11977)


tmp11979 := PrimNot(tmp11978)

var ifres11976 Obj

if True == tmp11979 {
ifres11976 = True


} else {
ifres11976 = False


}

ifres11975 = ifres11976


} else {
ifres11975 = False


}

if True == ifres11975 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1781, V1780, V1782, V1783)
return
} else {
tmp11973 := PrimIsPair(V1780)

var ifres11970 Obj

if True == tmp11973 {
tmp11972 := PrimIsPair(V1781)

var ifres11971 Obj

if True == tmp11972 {
ifres11971 = True


} else {
ifres11971 = False


}

ifres11970 = ifres11971


} else {
ifres11970 = False


}

if True == ifres11970 {
tmp11960 := PrimHead(V1780)

tmp11961 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11960, V1782)


tmp11962 := PrimHead(V1781)

tmp11963 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11962, V1782)


tmp11964 := MakeNative(func(__e *ControlFlow) {
tmp11965 := PrimTail(V1780)

tmp11966 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11965, V1782)


tmp11967 := PrimTail(V1781)

tmp11968 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11967, V1782)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp11966, tmp11968, V1782, V1783)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp11961, tmp11963, V1782, tmp11964)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp11990 := Call(__e, ns2_1set, symshen_4lzy_a_b, tmp11959)


_ = tmp11990

tmp11991 := MakeNative(func(__e *ControlFlow) {
V1795 := __e.Get(1)
_ = V1795
V1796 := __e.Get(2)
_ = V1796
V1797 := __e.Get(3)
_ = V1797
V1798 := __e.Get(4)
_ = V1798
tmp12011 := PrimEqual(V1795, V1796)

if True == tmp12011 {
__e.TailApply(PrimFunc(symthaw), V1798)
return
} else {
tmp12009 := Call(__e, PrimFunc(symshen_4pvar_2), V1795)


if True == tmp12009 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1795, V1796, V1797, V1798)
return
} else {
tmp12007 := Call(__e, PrimFunc(symshen_4pvar_2), V1796)


if True == tmp12007 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1796, V1795, V1797, V1798)
return
} else {
tmp12005 := PrimIsPair(V1795)

var ifres12002 Obj

if True == tmp12005 {
tmp12004 := PrimIsPair(V1796)

var ifres12003 Obj

if True == tmp12004 {
ifres12003 = True


} else {
ifres12003 = False


}

ifres12002 = ifres12003


} else {
ifres12002 = False


}

if True == ifres12002 {
tmp11992 := PrimHead(V1795)

tmp11993 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11992, V1797)


tmp11994 := PrimHead(V1796)

tmp11995 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11994, V1797)


tmp11996 := MakeNative(func(__e *ControlFlow) {
tmp11997 := PrimTail(V1795)

tmp11998 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11997, V1797)


tmp11999 := PrimTail(V1796)

tmp12000 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11999, V1797)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp11998, tmp12000, V1797, V1798)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4lzy_a), tmp11993, tmp11995, V1797, tmp11996)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp12012 := Call(__e, ns2_1set, symshen_4lzy_a, tmp11991)


_ = tmp12012

tmp12013 := MakeNative(func(__e *ControlFlow) {
V1804 := __e.Get(1)
_ = V1804
V1805 := __e.Get(2)
_ = V1805
tmp12023 := PrimEqual(V1804, V1805)

if True == tmp12023 {
__e.Return(True)
return
} else {
tmp12021 := PrimIsPair(V1805)

if True == tmp12021 {
tmp12018 := PrimHead(V1805)

tmp12019 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1804, tmp12018)


if True == tmp12019 {
__e.Return(True)
return
} else {
tmp12015 := PrimTail(V1805)

tmp12016 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1804, tmp12015)


if True == tmp12016 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


} else {
__e.Return(False)
return
}


}


}, 2)

tmp12024 := Call(__e, ns2_1set, symshen_4occurs_1check_2, tmp12013)


_ = tmp12024

tmp12025 := MakeNative(func(__e *ControlFlow) {
V1806 := __e.Get(1)
_ = V1806
V1807 := __e.Get(2)
_ = V1807
V1808 := __e.Get(3)
_ = V1808
V1809 := __e.Get(4)
_ = V1809
V1810 := __e.Get(5)
_ = V1810
tmp12026 := Call(__e, V1806, V1807)


tmp12027 := Call(__e, tmp12026, V1808)


tmp12028 := Call(__e, tmp12027, V1809)


__e.TailApply(tmp12028, V1810)
return


}, 5)

tmp12029 := Call(__e, ns2_1set, symcall, tmp12025)


_ = tmp12029

tmp12030 := MakeNative(func(__e *ControlFlow) {
V1817 := __e.Get(1)
_ = V1817
V1818 := __e.Get(2)
_ = V1818
V1819 := __e.Get(3)
_ = V1819
V1820 := __e.Get(4)
_ = V1820
V1821 := __e.Get(5)
_ = V1821
__e.TailApply(PrimFunc(symshen_4deref), V1817, V1818)
return
}, 5)

tmp12031 := Call(__e, ns2_1set, symreturn, tmp12030)


_ = tmp12031

tmp12032 := MakeNative(func(__e *ControlFlow) {
V1828 := __e.Get(1)
_ = V1828
V1829 := __e.Get(2)
_ = V1829
V1830 := __e.Get(3)
_ = V1830
V1831 := __e.Get(4)
_ = V1831
V1832 := __e.Get(5)
_ = V1832
if True == V1828 {
__e.TailApply(PrimFunc(symthaw), V1832)
return
} else {
__e.Return(False)
return
}
}, 5)

tmp12034 := Call(__e, ns2_1set, symwhen, tmp12032)


_ = tmp12034

tmp12035 := MakeNative(func(__e *ControlFlow) {
V1833 := __e.Get(1)
_ = V1833
V1834 := __e.Get(2)
_ = V1834
V1835 := __e.Get(3)
_ = V1835
V1836 := __e.Get(4)
_ = V1836
V1837 := __e.Get(5)
_ = V1837
V1838 := __e.Get(6)
_ = V1838
tmp12036 := Call(__e, PrimFunc(symshen_4lazyderef), V1833, V1835)


tmp12037 := Call(__e, PrimFunc(symshen_4lazyderef), V1834, V1835)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12036, tmp12037, V1835, V1838)
return


}, 6)

tmp12038 := Call(__e, ns2_1set, symis, tmp12035)


_ = tmp12038

tmp12039 := MakeNative(func(__e *ControlFlow) {
V1839 := __e.Get(1)
_ = V1839
V1840 := __e.Get(2)
_ = V1840
V1841 := __e.Get(3)
_ = V1841
V1842 := __e.Get(4)
_ = V1842
V1843 := __e.Get(5)
_ = V1843
V1844 := __e.Get(6)
_ = V1844
tmp12040 := Call(__e, PrimFunc(symshen_4lazyderef), V1839, V1841)


tmp12041 := Call(__e, PrimFunc(symshen_4lazyderef), V1840, V1841)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp12040, tmp12041, V1841, V1844)
return


}, 6)

tmp12042 := Call(__e, ns2_1set, symis_b, tmp12039)


_ = tmp12042

tmp12043 := MakeNative(func(__e *ControlFlow) {
V1849 := __e.Get(1)
_ = V1849
V1850 := __e.Get(2)
_ = V1850
V1851 := __e.Get(3)
_ = V1851
V1852 := __e.Get(4)
_ = V1852
V1853 := __e.Get(5)
_ = V1853
V1854 := __e.Get(6)
_ = V1854
__e.TailApply(PrimFunc(symshen_4bind_b), V1849, V1850, V1851, V1854)
return
}, 6)

tmp12044 := Call(__e, ns2_1set, symbind, tmp12043)


_ = tmp12044

tmp12045 := MakeNative(func(__e *ControlFlow) {
V1855 := __e.Get(1)
_ = V1855
V1856 := __e.Get(2)
_ = V1856
V1857 := __e.Get(3)
_ = V1857
V1858 := __e.Get(4)
_ = V1858
V1859 := __e.Get(5)
_ = V1859
tmp12047 := Call(__e, PrimFunc(symshen_4lazyderef), V1855, V1856)


tmp12048 := Call(__e, PrimFunc(symshen_4pvar_2), tmp12047)


if True == tmp12048 {
__e.TailApply(PrimFunc(symthaw), V1859)
return
} else {
__e.Return(False)
return
}


}, 5)

tmp12049 := Call(__e, ns2_1set, symvar_2, tmp12045)


_ = tmp12049

tmp12050 := MakeNative(func(__e *ControlFlow) {
V1862 := __e.Get(1)
_ = V1862
__e.Return(MakeString("|prolog vector|"))
return
}, 1)

tmp12051 := Call(__e, ns2_1set, symshen_4print_1prolog_1vector, tmp12050)


_ = tmp12051

tmp12052 := MakeNative(func(__e *ControlFlow) {
V1881 := __e.Get(1)
_ = V1881
V1882 := __e.Get(2)
_ = V1882
V1883 := __e.Get(3)
_ = V1883
V1884 := __e.Get(4)
_ = V1884
V1885 := __e.Get(5)
_ = V1885
tmp12065 := PrimEqual(Nil, V1881)

if True == tmp12065 {
__e.Return(False)
return
} else {
tmp12063 := PrimIsPair(V1881)

if True == tmp12063 {
tmp12053 := MakeNative(func(__e *ControlFlow) {
W1886 := __e.Get(1)
_ = W1886
tmp12056 := PrimEqual(W1886, False)

if True == tmp12056 {
tmp12054 := PrimTail(V1881)

__e.TailApply(PrimFunc(symfork), tmp12054, V1882, V1883, V1884, V1885)
return


} else {
__e.Return(W1886)
return
}


}, 1)

tmp12057 := PrimHead(V1881)

tmp12058 := Call(__e, tmp12057, V1882)


tmp12059 := Call(__e, tmp12058, V1883)


tmp12060 := Call(__e, tmp12059, V1884)


tmp12061 := Call(__e, tmp12060, V1885)


__e.TailApply(tmp12053, tmp12061)
return


} else {
__e.Return(PrimSimpleError(MakeString("fork expects a list of literals\n")))
return
}


}


}, 5)

tmp12066 := Call(__e, ns2_1set, symfork, tmp12052)


_ = tmp12066

tmp12067 := MakeNative(func(__e *ControlFlow) {
V1887 := __e.Get(1)
_ = V1887
V1888 := __e.Get(2)
_ = V1888
V1889 := __e.Get(3)
_ = V1889
V1890 := __e.Get(4)
_ = V1890
V1891 := __e.Get(5)
_ = V1891
V1892 := __e.Get(6)
_ = V1892
V1893 := __e.Get(7)
_ = V1893
tmp12074 := Call(__e, PrimFunc(symshen_4unlocked_2), V1891)


if True == tmp12074 {
tmp12068 := MakeNative(func(__e *ControlFlow) {
W1894 := __e.Get(1)
_ = W1894
tmp12069 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12069

tmp12070 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4findall_1h), V1887, V1888, V1889, W1894, V1890, V1891, V1892, V1893)
return
}, 0)

tmp12071 := Call(__e, PrimFunc(symis), W1894, Nil, V1890, V1891, V1892, tmp12070)


__e.TailApply(PrimFunc(symshen_4gc), V1890, tmp12071)
return


}, 1)

tmp12072 := Call(__e, PrimFunc(symshen_4newpv), V1890)


__e.TailApply(tmp12068, tmp12072)
return


} else {
__e.Return(False)
return
}


}, 7)

tmp12075 := Call(__e, ns2_1set, symfindall, tmp12067)


_ = tmp12075

tmp12076 := MakeNative(func(__e *ControlFlow) {
V1895 := __e.Get(1)
_ = V1895
V1896 := __e.Get(2)
_ = V1896
V1897 := __e.Get(3)
_ = V1897
V1898 := __e.Get(4)
_ = V1898
V1899 := __e.Get(5)
_ = V1899
V1900 := __e.Get(6)
_ = V1900
V1901 := __e.Get(7)
_ = V1901
V1902 := __e.Get(8)
_ = V1902
tmp12077 := MakeNative(func(__e *ControlFlow) {
W1903 := __e.Get(1)
_ = W1903
tmp12082 := PrimEqual(W1903, False)

if True == tmp12082 {
tmp12080 := Call(__e, PrimFunc(symshen_4unlocked_2), V1900)


if True == tmp12080 {
tmp12078 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12078

__e.TailApply(PrimFunc(symis_b), V1897, V1898, V1899, V1900, V1901, V1902)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W1903)
return
}


}, 1)

tmp12087 := Call(__e, PrimFunc(symshen_4unlocked_2), V1900)


var ifres12083 Obj

if True == tmp12087 {
tmp12084 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12084

tmp12085 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4overbind), V1895, V1898, V1899, V1900, V1901, V1902)
return
}, 0)

tmp12086 := Call(__e, PrimFunc(symcall), V1896, V1899, V1900, V1901, tmp12085)


ifres12083 = tmp12086


} else {
ifres12083 = False


}

__e.TailApply(tmp12077, ifres12083)
return


}, 8)

tmp12088 := Call(__e, ns2_1set, symshen_4findall_1h, tmp12076)


_ = tmp12088

tmp12089 := MakeNative(func(__e *ControlFlow) {
V1910 := __e.Get(1)
_ = V1910
V1911 := __e.Get(2)
_ = V1911
V1912 := __e.Get(3)
_ = V1912
V1913 := __e.Get(4)
_ = V1913
V1914 := __e.Get(5)
_ = V1914
V1915 := __e.Get(6)
_ = V1915
tmp12090 := Call(__e, PrimFunc(symshen_4deref), V1910, V1912)


tmp12091 := Call(__e, PrimFunc(symshen_4lazyderef), V1911, V1912)


tmp12092 := PrimCons(tmp12090, tmp12091)

tmp12093 := Call(__e, PrimFunc(symshen_4bindv), V1911, tmp12092, V1912)


_ = tmp12093

__e.Return(False)
return


}, 6)

tmp12094 := Call(__e, ns2_1set, symshen_4overbind, tmp12089)


_ = tmp12094

tmp12095 := MakeNative(func(__e *ControlFlow) {
V1918 := __e.Get(1)
_ = V1918
tmp12099 := PrimEqual(sym_7, V1918)

if True == tmp12099 {
__e.Return(PrimSet(symshen_4_doccurs_d, True))
return
} else {
tmp12097 := PrimEqual(sym_1, V1918)

if True == tmp12097 {
__e.Return(PrimSet(symshen_4_doccurs_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("occurs-check expects a + or a -.\n")))
return
}


}


}, 1)

__e.TailApply(ns2_1set, symoccurs_1check, tmp12095)
return




}, 0)

