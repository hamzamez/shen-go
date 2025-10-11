package main

import . "github.com/tiancaiamao/shen-go/kl"

var SysMain = MakeNative(func(__e *ControlFlow) {
tmp1146 := MakeNative(func(__e *ControlFlow) {
V3487 := __e.Get(1)
_ = V3487
__e.TailApply(V3487)
return
}, 1)

tmp1147 := Call(__e, ns2_1set, symthaw, tmp1146)


_ = tmp1147

tmp1148 := MakeNative(func(__e *ControlFlow) {
V3488 := __e.Get(1)
_ = V3488
tmp1149 := Call(__e, PrimFunc(symmacroexpand), V3488)


tmp1150 := Call(__e, PrimFunc(symshen_4find_1types), V3488)


tmp1151 := Call(__e, PrimFunc(symshen_4process_1applications), tmp1149, tmp1150)


tmp1152 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp1151)


__e.TailApply(PrimFunc(symeval_1kl), tmp1152)
return


}, 1)

tmp1153 := Call(__e, ns2_1set, symeval, tmp1148)


_ = tmp1153

tmp1154 := MakeNative(func(__e *ControlFlow) {
V3489 := __e.Get(1)
_ = V3489
tmp1161 := PrimEqual(symnull, V3489)

if True == tmp1161 {
__e.Return(Nil)
return
} else {
tmp1155 := MakeNative(func(__e *ControlFlow) {
tmp1156 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V3489, symshen_4external_1symbols, tmp1156)
return


}, 0)

tmp1157 := MakeNative(func(__e *ControlFlow) {
Z3490 := __e.Get(1)
_ = Z3490
tmp1158 := Call(__e, PrimFunc(symshen_4app), V3489, MakeString(" does not exist.\n;"), symshen_4a)


tmp1159 := PrimStringConcat(MakeString("package "), tmp1158)

__e.Return(PrimSimpleError(tmp1159))
return


}, 1)

__e.TailApply(try_1catch, tmp1155, tmp1157)
return


}


}, 1)

tmp1162 := Call(__e, ns2_1set, symexternal, tmp1154)


_ = tmp1162

tmp1163 := MakeNative(func(__e *ControlFlow) {
V3491 := __e.Get(1)
_ = V3491
tmp1170 := PrimEqual(symnull, V3491)

if True == tmp1170 {
__e.Return(Nil)
return
} else {
tmp1164 := MakeNative(func(__e *ControlFlow) {
tmp1165 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V3491, symshen_4internal_1symbols, tmp1165)
return


}, 0)

tmp1166 := MakeNative(func(__e *ControlFlow) {
Z3492 := __e.Get(1)
_ = Z3492
tmp1167 := Call(__e, PrimFunc(symshen_4app), V3491, MakeString(" does not exist.\n;"), symshen_4a)


tmp1168 := PrimStringConcat(MakeString("package "), tmp1167)

__e.Return(PrimSimpleError(tmp1168))
return


}, 1)

__e.TailApply(try_1catch, tmp1164, tmp1166)
return


}


}, 1)

tmp1171 := Call(__e, ns2_1set, syminternal, tmp1163)


_ = tmp1171

tmp1172 := MakeNative(func(__e *ControlFlow) {
V3493 := __e.Get(1)
_ = V3493
V3494 := __e.Get(2)
_ = V3494
tmp1174 := Call(__e, V3493, V3494)


if True == tmp1174 {
__e.TailApply(PrimFunc(symfail))
return
} else {
__e.Return(V3494)
return
}


}, 2)

tmp1175 := Call(__e, ns2_1set, symfail_1if, tmp1172)


_ = tmp1175

tmp1176 := MakeNative(func(__e *ControlFlow) {
V3495 := __e.Get(1)
_ = V3495
V3496 := __e.Get(2)
_ = V3496
__e.Return(PrimStringConcat(V3495, V3496))
return
}, 2)

tmp1177 := Call(__e, ns2_1set, sym_8s, tmp1176)


_ = tmp1177

tmp1178 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dtc_d))
return
}, 0)

tmp1179 := Call(__e, ns2_1set, symtc_2, tmp1178)


_ = tmp1179

tmp1180 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_doccurs_d))
return
}, 0)

tmp1181 := Call(__e, ns2_1set, symoccurs_2, tmp1180)


_ = tmp1181

tmp1182 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dfactorise_2_d))
return
}, 0)

tmp1183 := Call(__e, ns2_1set, symfactorise_2, tmp1182)


_ = tmp1183

tmp1184 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dtracking_d))
return
}, 0)

tmp1185 := Call(__e, ns2_1set, symtracked, tmp1184)


_ = tmp1185

tmp1186 := MakeNative(func(__e *ControlFlow) {
V3497 := __e.Get(1)
_ = V3497
tmp1187 := MakeNative(func(__e *ControlFlow) {
tmp1188 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V3497, symshen_4source, tmp1188)
return


}, 0)

tmp1189 := MakeNative(func(__e *ControlFlow) {
Z3498 := __e.Get(1)
_ = Z3498
tmp1190 := Call(__e, PrimFunc(symshen_4app), V3497, MakeString(" not found.\n"), symshen_4a)


__e.Return(PrimSimpleError(tmp1190))
return


}, 1)

__e.TailApply(try_1catch, tmp1187, tmp1189)
return


}, 1)

tmp1191 := Call(__e, ns2_1set, symps, tmp1186)


_ = tmp1191

tmp1192 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dstinput_d))
return
}, 0)

tmp1193 := Call(__e, ns2_1set, symstinput, tmp1192)


_ = tmp1193

tmp1194 := MakeNative(func(__e *ControlFlow) {
V3499 := __e.Get(1)
_ = V3499
tmp1195 := MakeNative(func(__e *ControlFlow) {
W3500 := __e.Get(1)
_ = W3500
tmp1196 := MakeNative(func(__e *ControlFlow) {
W3501 := __e.Get(1)
_ = W3501
tmp1197 := MakeNative(func(__e *ControlFlow) {
W3502 := __e.Get(1)
_ = W3502
__e.Return(W3502)
return
}, 1)

tmp1201 := PrimEqual(V3499, MakeNumber(0))

var ifres1198 Obj

if True == tmp1201 {
ifres1198 = W3501


} else {
tmp1199 := Call(__e, PrimFunc(symfail))


tmp1200 := Call(__e, PrimFunc(symshen_4fillvector), W3501, MakeNumber(1), V3499, tmp1199)


ifres1198 = tmp1200


}

__e.TailApply(tmp1197, ifres1198)
return


}, 1)

tmp1202 := PrimVectorSet(W3500, MakeNumber(0), V3499)

__e.TailApply(tmp1196, tmp1202)
return


}, 1)

tmp1203 := PrimNumberAdd(V3499, MakeNumber(1))

tmp1204 := PrimAbsvector(tmp1203)

__e.TailApply(tmp1195, tmp1204)
return


}, 1)

tmp1205 := Call(__e, ns2_1set, symvector, tmp1194)


_ = tmp1205

tmp1206 := MakeNative(func(__e *ControlFlow) {
V3504 := __e.Get(1)
_ = V3504
V3505 := __e.Get(2)
_ = V3505
V3506 := __e.Get(3)
_ = V3506
V3507 := __e.Get(4)
_ = V3507
tmp1210 := PrimEqual(V3505, V3506)

if True == tmp1210 {
__e.Return(PrimVectorSet(V3504, V3506, V3507))
return
} else {
tmp1207 := PrimVectorSet(V3504, V3505, V3507)

tmp1208 := PrimNumberAdd(MakeNumber(1), V3505)

__e.TailApply(PrimFunc(symshen_4fillvector), tmp1207, tmp1208, V3506, V3507)
return


}


}, 4)

tmp1211 := Call(__e, ns2_1set, symshen_4fillvector, tmp1206)


_ = tmp1211

tmp1212 := MakeNative(func(__e *ControlFlow) {
V3508 := __e.Get(1)
_ = V3508
tmp1219 := PrimIsVector(V3508)

if True == tmp1219 {
tmp1214 := MakeNative(func(__e *ControlFlow) {
tmp1215 := PrimVectorGet(V3508, MakeNumber(0))

__e.Return(PrimGreatEqual(tmp1215, MakeNumber(0)))
return


}, 0)

tmp1216 := MakeNative(func(__e *ControlFlow) {
Z3509 := __e.Get(1)
_ = Z3509
__e.Return(False)
return
}, 1)

tmp1217 := Call(__e, try_1catch, tmp1214, tmp1216)


if True == tmp1217 {
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

tmp1220 := Call(__e, ns2_1set, symvector_2, tmp1212)


_ = tmp1220

tmp1221 := MakeNative(func(__e *ControlFlow) {
V3510 := __e.Get(1)
_ = V3510
V3511 := __e.Get(2)
_ = V3511
V3512 := __e.Get(3)
_ = V3512
tmp1223 := PrimEqual(V3511, MakeNumber(0))

if True == tmp1223 {
__e.Return(PrimSimpleError(MakeString("cannot access 0th element of a vector\n")))
return
} else {
__e.Return(PrimVectorSet(V3510, V3511, V3512))
return
}


}, 3)

tmp1224 := Call(__e, ns2_1set, symvector_1_6, tmp1221)


_ = tmp1224

tmp1225 := MakeNative(func(__e *ControlFlow) {
V3513 := __e.Get(1)
_ = V3513
V3514 := __e.Get(2)
_ = V3514
tmp1232 := PrimEqual(V3514, MakeNumber(0))

if True == tmp1232 {
__e.Return(PrimSimpleError(MakeString("cannot access 0th element of a vector\n")))
return
} else {
tmp1226 := MakeNative(func(__e *ControlFlow) {
W3515 := __e.Get(1)
_ = W3515
tmp1228 := Call(__e, PrimFunc(symfail))


tmp1229 := PrimEqual(W3515, tmp1228)

if True == tmp1229 {
__e.Return(PrimSimpleError(MakeString("vector element not found\n")))
return
} else {
__e.Return(W3515)
return
}


}, 1)

tmp1230 := PrimVectorGet(V3513, V3514)

__e.TailApply(tmp1226, tmp1230)
return


}


}, 2)

tmp1233 := Call(__e, ns2_1set, sym_5_1vector, tmp1225)


_ = tmp1233

tmp1234 := MakeNative(func(__e *ControlFlow) {
V3516 := __e.Get(1)
_ = V3516
tmp1238 := PrimIsInteger(V3516)

if True == tmp1238 {
tmp1236 := PrimGreatEqual(V3516, MakeNumber(0))

if True == tmp1236 {
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

tmp1239 := Call(__e, ns2_1set, symshen_4posint_2, tmp1234)


_ = tmp1239

tmp1240 := MakeNative(func(__e *ControlFlow) {
V3517 := __e.Get(1)
_ = V3517
__e.Return(PrimVectorGet(V3517, MakeNumber(0)))
return
}, 1)

tmp1241 := Call(__e, ns2_1set, symlimit, tmp1240)


_ = tmp1241

tmp1242 := MakeNative(func(__e *ControlFlow) {
V3518 := __e.Get(1)
_ = V3518
tmp1273 := Call(__e, PrimFunc(symboolean_2), V3518)


var ifres1258 Obj

if True == tmp1273 {
ifres1258 = True


} else {
tmp1272 := PrimIsNumber(V3518)

var ifres1260 Obj

if True == tmp1272 {
ifres1260 = True


} else {
tmp1271 := PrimIsString(V3518)

var ifres1262 Obj

if True == tmp1271 {
ifres1262 = True


} else {
tmp1270 := PrimIsPair(V3518)

var ifres1264 Obj

if True == tmp1270 {
ifres1264 = True


} else {
tmp1269 := Call(__e, PrimFunc(symempty_2), V3518)


var ifres1266 Obj

if True == tmp1269 {
ifres1266 = True


} else {
tmp1268 := Call(__e, PrimFunc(symvector_2), V3518)


var ifres1267 Obj

if True == tmp1268 {
ifres1267 = True


} else {
ifres1267 = False


}

ifres1266 = ifres1267


}

var ifres1265 Obj

if True == ifres1266 {
ifres1265 = True


} else {
ifres1265 = False


}

ifres1264 = ifres1265


}

var ifres1263 Obj

if True == ifres1264 {
ifres1263 = True


} else {
ifres1263 = False


}

ifres1262 = ifres1263


}

var ifres1261 Obj

if True == ifres1262 {
ifres1261 = True


} else {
ifres1261 = False


}

ifres1260 = ifres1261


}

var ifres1259 Obj

if True == ifres1260 {
ifres1259 = True


} else {
ifres1259 = False


}

ifres1258 = ifres1259


}

if True == ifres1258 {
__e.Return(False)
return
} else {
tmp1248 := PrimIntern(MakeString(":"))

tmp1249 := PrimIntern(MakeString(";"))

tmp1250 := PrimIntern(MakeString(","))

tmp1251 := PrimCons(tmp1250, Nil)

tmp1252 := PrimCons(tmp1249, tmp1251)

tmp1253 := PrimCons(tmp1248, tmp1252)

tmp1254 := PrimCons(sym_j, tmp1253)

tmp1255 := PrimCons(sym_i, tmp1254)

tmp1256 := Call(__e, PrimFunc(symelement_2), V3518, tmp1255)


if True == tmp1256 {
__e.Return(True)
return
} else {
tmp1243 := MakeNative(func(__e *ControlFlow) {
tmp1244 := MakeNative(func(__e *ControlFlow) {
W3519 := __e.Get(1)
_ = W3519
__e.TailApply(PrimFunc(symshen_4analyse_1symbol_2), W3519)
return
}, 1)

tmp1245 := PrimStr(V3518)

__e.TailApply(tmp1244, tmp1245)
return


}, 0)

tmp1246 := MakeNative(func(__e *ControlFlow) {
Z3520 := __e.Get(1)
_ = Z3520
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1243, tmp1246)
return


}


}


}, 1)

tmp1274 := Call(__e, ns2_1set, symsymbol_2, tmp1242)


_ = tmp1274

tmp1275 := MakeNative(func(__e *ControlFlow) {
V3523 := __e.Get(1)
_ = V3523
tmp1284 := Call(__e, PrimFunc(symshen_4_7string_2), V3523)


if True == tmp1284 {
tmp1280 := Call(__e, PrimFunc(symhdstr), V3523)


tmp1281 := PrimStringToNumber(tmp1280)

tmp1282 := Call(__e, PrimFunc(symshen_4alpha_2), tmp1281)


if True == tmp1282 {
tmp1277 := PrimTailString(V3523)

tmp1278 := Call(__e, PrimFunc(symshen_4alphanums_2), tmp1277)


if True == tmp1278 {
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
__e.Return(PrimSimpleError(MakeString("implementation error in shen.analyse-symbol?")))
return
}


}, 1)

tmp1285 := Call(__e, ns2_1set, symshen_4analyse_1symbol_2, tmp1275)


_ = tmp1285

tmp1286 := MakeNative(func(__e *ControlFlow) {
V3526 := __e.Get(1)
_ = V3526
tmp1301 := PrimEqual(MakeString(""), V3526)

if True == tmp1301 {
__e.Return(True)
return
} else {
tmp1299 := Call(__e, PrimFunc(symshen_4_7string_2), V3526)


if True == tmp1299 {
tmp1287 := MakeNative(func(__e *ControlFlow) {
W3527 := __e.Get(1)
_ = W3527
tmp1295 := Call(__e, PrimFunc(symshen_4alpha_2), W3527)


var ifres1292 Obj

if True == tmp1295 {
ifres1292 = True


} else {
tmp1294 := Call(__e, PrimFunc(symshen_4digit_2), W3527)


var ifres1293 Obj

if True == tmp1294 {
ifres1293 = True


} else {
ifres1293 = False


}

ifres1292 = ifres1293


}

if True == ifres1292 {
tmp1289 := PrimTailString(V3526)

tmp1290 := Call(__e, PrimFunc(symshen_4alphanums_2), tmp1289)


if True == tmp1290 {
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

tmp1296 := Call(__e, PrimFunc(symhdstr), V3526)


tmp1297 := PrimStringToNumber(tmp1296)

__e.TailApply(tmp1287, tmp1297)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.alphanums?")))
return
}


}


}, 1)

tmp1302 := Call(__e, ns2_1set, symshen_4alphanums_2, tmp1286)


_ = tmp1302

tmp1303 := MakeNative(func(__e *ControlFlow) {
V3528 := __e.Get(1)
_ = V3528
tmp1315 := Call(__e, PrimFunc(symboolean_2), V3528)


var ifres1309 Obj

if True == tmp1315 {
ifres1309 = True


} else {
tmp1314 := PrimIsNumber(V3528)

var ifres1311 Obj

if True == tmp1314 {
ifres1311 = True


} else {
tmp1313 := PrimIsString(V3528)

var ifres1312 Obj

if True == tmp1313 {
ifres1312 = True


} else {
ifres1312 = False


}

ifres1311 = ifres1312


}

var ifres1310 Obj

if True == ifres1311 {
ifres1310 = True


} else {
ifres1310 = False


}

ifres1309 = ifres1310


}

if True == ifres1309 {
__e.Return(False)
return
} else {
tmp1304 := MakeNative(func(__e *ControlFlow) {
tmp1305 := MakeNative(func(__e *ControlFlow) {
W3529 := __e.Get(1)
_ = W3529
__e.TailApply(PrimFunc(symshen_4analyse_1variable_2), W3529)
return
}, 1)

tmp1306 := PrimStr(V3528)

__e.TailApply(tmp1305, tmp1306)
return


}, 0)

tmp1307 := MakeNative(func(__e *ControlFlow) {
Z3530 := __e.Get(1)
_ = Z3530
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1304, tmp1307)
return


}


}, 1)

tmp1316 := Call(__e, ns2_1set, symvariable_2, tmp1303)


_ = tmp1316

tmp1317 := MakeNative(func(__e *ControlFlow) {
V3533 := __e.Get(1)
_ = V3533
tmp1326 := Call(__e, PrimFunc(symshen_4_7string_2), V3533)


if True == tmp1326 {
tmp1322 := Call(__e, PrimFunc(symhdstr), V3533)


tmp1323 := PrimStringToNumber(tmp1322)

tmp1324 := Call(__e, PrimFunc(symshen_4uppercase_2), tmp1323)


if True == tmp1324 {
tmp1319 := PrimTailString(V3533)

tmp1320 := Call(__e, PrimFunc(symshen_4alphanums_2), tmp1319)


if True == tmp1320 {
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
__e.Return(PrimSimpleError(MakeString("implementation error in shen.analyse-variable?")))
return
}


}, 1)

tmp1327 := Call(__e, ns2_1set, symshen_4analyse_1variable_2, tmp1317)


_ = tmp1327

tmp1328 := MakeNative(func(__e *ControlFlow) {
V3534 := __e.Get(1)
_ = V3534
tmp1329 := PrimValue(symshen_4_dgensym_d)

tmp1330 := PrimNumberAdd(MakeNumber(1), tmp1329)

tmp1331 := PrimSet(symshen_4_dgensym_d, tmp1330)

__e.TailApply(PrimFunc(symconcat), V3534, tmp1331)
return


}, 1)

tmp1332 := Call(__e, ns2_1set, symgensym, tmp1328)


_ = tmp1332

tmp1333 := MakeNative(func(__e *ControlFlow) {
V3535 := __e.Get(1)
_ = V3535
V3536 := __e.Get(2)
_ = V3536
tmp1334 := PrimStr(V3535)

tmp1335 := PrimStr(V3536)

tmp1336 := PrimStringConcat(tmp1334, tmp1335)

__e.Return(PrimIntern(tmp1336))
return


}, 2)

tmp1337 := Call(__e, ns2_1set, symconcat, tmp1333)


_ = tmp1337

tmp1338 := MakeNative(func(__e *ControlFlow) {
V3537 := __e.Get(1)
_ = V3537
V3538 := __e.Get(2)
_ = V3538
tmp1339 := MakeNative(func(__e *ControlFlow) {
W3539 := __e.Get(1)
_ = W3539
tmp1340 := MakeNative(func(__e *ControlFlow) {
W3540 := __e.Get(1)
_ = W3540
tmp1341 := MakeNative(func(__e *ControlFlow) {
W3541 := __e.Get(1)
_ = W3541
tmp1342 := MakeNative(func(__e *ControlFlow) {
W3542 := __e.Get(1)
_ = W3542
__e.Return(W3539)
return
}, 1)

tmp1343 := PrimVectorSet(W3539, MakeNumber(2), V3538)

__e.TailApply(tmp1342, tmp1343)
return


}, 1)

tmp1344 := PrimVectorSet(W3539, MakeNumber(1), V3537)

__e.TailApply(tmp1341, tmp1344)
return


}, 1)

tmp1345 := PrimVectorSet(W3539, MakeNumber(0), symshen_4tuple)

__e.TailApply(tmp1340, tmp1345)
return


}, 1)

tmp1346 := PrimAbsvector(MakeNumber(3))

__e.TailApply(tmp1339, tmp1346)
return


}, 2)

tmp1347 := Call(__e, ns2_1set, sym_8p, tmp1338)


_ = tmp1347

tmp1348 := MakeNative(func(__e *ControlFlow) {
V3543 := __e.Get(1)
_ = V3543
__e.Return(PrimVectorGet(V3543, MakeNumber(1)))
return
}, 1)

tmp1349 := Call(__e, ns2_1set, symfst, tmp1348)


_ = tmp1349

tmp1350 := MakeNative(func(__e *ControlFlow) {
V3544 := __e.Get(1)
_ = V3544
__e.Return(PrimVectorGet(V3544, MakeNumber(2)))
return
}, 1)

tmp1351 := Call(__e, ns2_1set, symsnd, tmp1350)


_ = tmp1351

tmp1352 := MakeNative(func(__e *ControlFlow) {
V3545 := __e.Get(1)
_ = V3545
tmp1353 := MakeNative(func(__e *ControlFlow) {
tmp1358 := PrimIsVector(V3545)

if True == tmp1358 {
tmp1355 := PrimVectorGet(V3545, MakeNumber(0))

tmp1356 := PrimEqual(symshen_4tuple, tmp1355)

if True == tmp1356 {
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

tmp1359 := MakeNative(func(__e *ControlFlow) {
Z3546 := __e.Get(1)
_ = Z3546
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1353, tmp1359)
return


}, 1)

tmp1360 := Call(__e, ns2_1set, symtuple_2, tmp1352)


_ = tmp1360

tmp1361 := MakeNative(func(__e *ControlFlow) {
V3551 := __e.Get(1)
_ = V3551
V3552 := __e.Get(2)
_ = V3552
tmp1368 := PrimEqual(Nil, V3551)

if True == tmp1368 {
__e.Return(V3552)
return
} else {
tmp1366 := PrimIsPair(V3551)

if True == tmp1366 {
tmp1362 := PrimHead(V3551)

tmp1363 := PrimTail(V3551)

tmp1364 := Call(__e, PrimFunc(symappend), tmp1363, V3552)


__e.Return(PrimCons(tmp1362, tmp1364))
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to append a non-list")))
return
}


}


}, 2)

tmp1369 := Call(__e, ns2_1set, symappend, tmp1361)


_ = tmp1369

tmp1370 := MakeNative(func(__e *ControlFlow) {
V3553 := __e.Get(1)
_ = V3553
V3554 := __e.Get(2)
_ = V3554
tmp1371 := MakeNative(func(__e *ControlFlow) {
W3555 := __e.Get(1)
_ = W3555
tmp1372 := MakeNative(func(__e *ControlFlow) {
W3556 := __e.Get(1)
_ = W3556
tmp1373 := MakeNative(func(__e *ControlFlow) {
W3557 := __e.Get(1)
_ = W3557
tmp1375 := PrimEqual(W3555, MakeNumber(0))

if True == tmp1375 {
__e.Return(W3557)
return
} else {
__e.TailApply(PrimFunc(symshen_4_8v_1help), V3554, MakeNumber(1), W3555, W3557)
return
}


}, 1)

tmp1376 := Call(__e, PrimFunc(symvector_1_6), W3556, MakeNumber(1), V3553)


__e.TailApply(tmp1373, tmp1376)
return


}, 1)

tmp1377 := PrimNumberAdd(W3555, MakeNumber(1))

tmp1378 := Call(__e, PrimFunc(symvector), tmp1377)


__e.TailApply(tmp1372, tmp1378)
return


}, 1)

tmp1379 := Call(__e, PrimFunc(symlimit), V3554)


__e.TailApply(tmp1371, tmp1379)
return


}, 2)

tmp1380 := Call(__e, ns2_1set, sym_8v, tmp1370)


_ = tmp1380

tmp1381 := MakeNative(func(__e *ControlFlow) {
V3559 := __e.Get(1)
_ = V3559
V3560 := __e.Get(2)
_ = V3560
V3561 := __e.Get(3)
_ = V3561
V3562 := __e.Get(4)
_ = V3562
tmp1387 := PrimEqual(V3560, V3561)

if True == tmp1387 {
tmp1382 := PrimNumberAdd(V3561, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4copyfromvector), V3559, V3562, V3561, tmp1382)
return


} else {
tmp1383 := PrimNumberAdd(V3560, MakeNumber(1))

tmp1384 := PrimNumberAdd(V3560, MakeNumber(1))

tmp1385 := Call(__e, PrimFunc(symshen_4copyfromvector), V3559, V3562, V3560, tmp1384)


__e.TailApply(PrimFunc(symshen_4_8v_1help), V3559, tmp1383, V3561, tmp1385)
return


}


}, 4)

tmp1388 := Call(__e, ns2_1set, symshen_4_8v_1help, tmp1381)


_ = tmp1388

tmp1389 := MakeNative(func(__e *ControlFlow) {
V3563 := __e.Get(1)
_ = V3563
V3564 := __e.Get(2)
_ = V3564
V3565 := __e.Get(3)
_ = V3565
V3566 := __e.Get(4)
_ = V3566
tmp1390 := MakeNative(func(__e *ControlFlow) {
tmp1391 := Call(__e, PrimFunc(sym_5_1vector), V3563, V3565)


__e.TailApply(PrimFunc(symvector_1_6), V3564, V3566, tmp1391)
return


}, 0)

tmp1392 := MakeNative(func(__e *ControlFlow) {
Z3567 := __e.Get(1)
_ = Z3567
__e.Return(V3564)
return
}, 1)

__e.TailApply(try_1catch, tmp1390, tmp1392)
return


}, 4)

tmp1393 := Call(__e, ns2_1set, symshen_4copyfromvector, tmp1389)


_ = tmp1393

tmp1394 := MakeNative(func(__e *ControlFlow) {
V3568 := __e.Get(1)
_ = V3568
tmp1395 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3568, MakeNumber(1))
return
}, 0)

tmp1396 := MakeNative(func(__e *ControlFlow) {
Z3569 := __e.Get(1)
_ = Z3569
__e.Return(PrimSimpleError(MakeString("hdv needs a non-empty vector as an argument\n")))
return
}, 1)

__e.TailApply(try_1catch, tmp1395, tmp1396)
return


}, 1)

tmp1397 := Call(__e, ns2_1set, symhdv, tmp1394)


_ = tmp1397

tmp1398 := MakeNative(func(__e *ControlFlow) {
V3570 := __e.Get(1)
_ = V3570
tmp1399 := MakeNative(func(__e *ControlFlow) {
W3571 := __e.Get(1)
_ = W3571
tmp1408 := PrimEqual(W3571, MakeNumber(0))

if True == tmp1408 {
__e.Return(PrimSimpleError(MakeString("cannot take the tail of the empty vector\n")))
return
} else {
tmp1406 := PrimEqual(W3571, MakeNumber(1))

if True == tmp1406 {
__e.TailApply(PrimFunc(symvector), MakeNumber(0))
return
} else {
tmp1400 := MakeNative(func(__e *ControlFlow) {
W3572 := __e.Get(1)
_ = W3572
tmp1401 := PrimNumberSubtract(W3571, MakeNumber(1))

tmp1402 := Call(__e, PrimFunc(symvector), tmp1401)


__e.TailApply(PrimFunc(symshen_4tlv_1help), V3570, MakeNumber(2), W3571, tmp1402)
return


}, 1)

tmp1403 := PrimNumberSubtract(W3571, MakeNumber(1))

tmp1404 := Call(__e, PrimFunc(symvector), tmp1403)


__e.TailApply(tmp1400, tmp1404)
return


}


}


}, 1)

tmp1409 := Call(__e, PrimFunc(symlimit), V3570)


__e.TailApply(tmp1399, tmp1409)
return


}, 1)

tmp1410 := Call(__e, ns2_1set, symtlv, tmp1398)


_ = tmp1410

tmp1411 := MakeNative(func(__e *ControlFlow) {
V3574 := __e.Get(1)
_ = V3574
V3575 := __e.Get(2)
_ = V3575
V3576 := __e.Get(3)
_ = V3576
V3577 := __e.Get(4)
_ = V3577
tmp1417 := PrimEqual(V3575, V3576)

if True == tmp1417 {
tmp1412 := PrimNumberSubtract(V3576, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4copyfromvector), V3574, V3577, V3576, tmp1412)
return


} else {
tmp1413 := PrimNumberAdd(V3575, MakeNumber(1))

tmp1414 := PrimNumberSubtract(V3575, MakeNumber(1))

tmp1415 := Call(__e, PrimFunc(symshen_4copyfromvector), V3574, V3577, V3575, tmp1414)


__e.TailApply(PrimFunc(symshen_4tlv_1help), V3574, tmp1413, V3576, tmp1415)
return


}


}, 4)

tmp1418 := Call(__e, ns2_1set, symshen_4tlv_1help, tmp1411)


_ = tmp1418

tmp1419 := MakeNative(func(__e *ControlFlow) {
V3589 := __e.Get(1)
_ = V3589
V3590 := __e.Get(2)
_ = V3590
tmp1435 := PrimEqual(Nil, V3590)

if True == tmp1435 {
__e.Return(Nil)
return
} else {
tmp1433 := PrimIsPair(V3590)

var ifres1424 Obj

if True == tmp1433 {
tmp1431 := PrimHead(V3590)

tmp1432 := PrimIsPair(tmp1431)

var ifres1426 Obj

if True == tmp1432 {
tmp1428 := PrimHead(V3590)

tmp1429 := PrimHead(tmp1428)

tmp1430 := PrimEqual(V3589, tmp1429)

var ifres1427 Obj

if True == tmp1430 {
ifres1427 = True


} else {
ifres1427 = False


}

ifres1426 = ifres1427


} else {
ifres1426 = False


}

var ifres1425 Obj

if True == ifres1426 {
ifres1425 = True


} else {
ifres1425 = False


}

ifres1424 = ifres1425


} else {
ifres1424 = False


}

if True == ifres1424 {
__e.Return(PrimHead(V3590))
return
} else {
tmp1422 := PrimIsPair(V3590)

if True == tmp1422 {
tmp1420 := PrimTail(V3590)

__e.TailApply(PrimFunc(symassoc), V3589, tmp1420)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to search a non-list with assoc\n")))
return
}


}


}


}, 2)

tmp1436 := Call(__e, ns2_1set, symassoc, tmp1419)


_ = tmp1436

tmp1437 := MakeNative(func(__e *ControlFlow) {
V3593 := __e.Get(1)
_ = V3593
tmp1441 := PrimEqual(True, V3593)

if True == tmp1441 {
__e.Return(True)
return
} else {
tmp1439 := PrimEqual(False, V3593)

if True == tmp1439 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp1442 := Call(__e, ns2_1set, symboolean_2, tmp1437)


_ = tmp1442

tmp1443 := MakeNative(func(__e *ControlFlow) {
V3594 := __e.Get(1)
_ = V3594
tmp1448 := PrimEqual(MakeNumber(0), V3594)

if True == tmp1448 {
__e.Return(MakeNumber(0))
return
} else {
tmp1444 := Call(__e, PrimFunc(symstoutput))


tmp1445 := Call(__e, PrimFunc(sympr), MakeString("\n"), tmp1444)


_ = tmp1445

tmp1446 := PrimNumberSubtract(V3594, MakeNumber(1))

__e.TailApply(PrimFunc(symnl), tmp1446)
return


}


}, 1)

tmp1449 := Call(__e, ns2_1set, symnl, tmp1443)


_ = tmp1449

tmp1450 := MakeNative(func(__e *ControlFlow) {
V3601 := __e.Get(1)
_ = V3601
V3602 := __e.Get(2)
_ = V3602
tmp1461 := PrimEqual(Nil, V3601)

if True == tmp1461 {
__e.Return(Nil)
return
} else {
tmp1459 := PrimIsPair(V3601)

if True == tmp1459 {
tmp1456 := PrimHead(V3601)

tmp1457 := Call(__e, PrimFunc(symelement_2), tmp1456, V3602)


if True == tmp1457 {
tmp1451 := PrimTail(V3601)

__e.TailApply(PrimFunc(symdifference), tmp1451, V3602)
return


} else {
tmp1452 := PrimHead(V3601)

tmp1453 := PrimTail(V3601)

tmp1454 := Call(__e, PrimFunc(symdifference), tmp1453, V3602)


__e.Return(PrimCons(tmp1452, tmp1454))
return


}


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find the difference with a non-list\n")))
return
}


}


}, 2)

tmp1462 := Call(__e, ns2_1set, symdifference, tmp1450)


_ = tmp1462

tmp1463 := MakeNative(func(__e *ControlFlow) {
V3603 := __e.Get(1)
_ = V3603
V3604 := __e.Get(2)
_ = V3604
__e.Return(V3604)
return
}, 2)

tmp1464 := Call(__e, ns2_1set, symdo, tmp1463)


_ = tmp1464

tmp1465 := MakeNative(func(__e *ControlFlow) {
V3616 := __e.Get(1)
_ = V3616
V3617 := __e.Get(2)
_ = V3617
tmp1476 := PrimEqual(Nil, V3617)

if True == tmp1476 {
__e.Return(False)
return
} else {
tmp1474 := PrimIsPair(V3617)

var ifres1470 Obj

if True == tmp1474 {
tmp1472 := PrimHead(V3617)

tmp1473 := PrimEqual(V3616, tmp1472)

var ifres1471 Obj

if True == tmp1473 {
ifres1471 = True


} else {
ifres1471 = False


}

ifres1470 = ifres1471


} else {
ifres1470 = False


}

if True == ifres1470 {
__e.Return(True)
return
} else {
tmp1468 := PrimIsPair(V3617)

if True == tmp1468 {
tmp1466 := PrimTail(V3617)

__e.TailApply(PrimFunc(symelement_2), V3616, tmp1466)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find an element in a non-list\n")))
return
}


}


}


}, 2)

tmp1477 := Call(__e, ns2_1set, symelement_2, tmp1465)


_ = tmp1477

tmp1478 := MakeNative(func(__e *ControlFlow) {
V3620 := __e.Get(1)
_ = V3620
tmp1480 := PrimEqual(Nil, V3620)

if True == tmp1480 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp1481 := Call(__e, ns2_1set, symempty_2, tmp1478)


_ = tmp1481

tmp1482 := MakeNative(func(__e *ControlFlow) {
V3621 := __e.Get(1)
_ = V3621
V3622 := __e.Get(2)
_ = V3622
tmp1483 := Call(__e, V3621, V3622)


__e.TailApply(PrimFunc(symshen_4fix_1help), V3621, V3622, tmp1483)
return


}, 2)

tmp1484 := Call(__e, ns2_1set, symfix, tmp1482)


_ = tmp1484

tmp1485 := MakeNative(func(__e *ControlFlow) {
V3628 := __e.Get(1)
_ = V3628
V3629 := __e.Get(2)
_ = V3629
V3630 := __e.Get(3)
_ = V3630
tmp1488 := PrimEqual(V3629, V3630)

if True == tmp1488 {
__e.Return(V3630)
return
} else {
tmp1486 := Call(__e, V3628, V3630)


__e.TailApply(PrimFunc(symshen_4fix_1help), V3628, V3630, tmp1486)
return


}


}, 3)

tmp1489 := Call(__e, ns2_1set, symshen_4fix_1help, tmp1485)


_ = tmp1489

tmp1490 := MakeNative(func(__e *ControlFlow) {
V3631 := __e.Get(1)
_ = V3631
V3632 := __e.Get(2)
_ = V3632
V3633 := __e.Get(3)
_ = V3633
V3634 := __e.Get(4)
_ = V3634
tmp1491 := MakeNative(func(__e *ControlFlow) {
W3635 := __e.Get(1)
_ = W3635
tmp1492 := MakeNative(func(__e *ControlFlow) {
W3636 := __e.Get(1)
_ = W3636
tmp1493 := MakeNative(func(__e *ControlFlow) {
W3638 := __e.Get(1)
_ = W3638
__e.Return(V3633)
return
}, 1)

tmp1494 := Call(__e, PrimFunc(symshen_4change_1pointer_1value), V3631, V3632, V3633, W3636)


tmp1495 := Call(__e, PrimFunc(symvector_1_6), V3634, W3635, tmp1494)


__e.TailApply(tmp1493, tmp1495)
return


}, 1)

tmp1496 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3634, W3635)
return
}, 0)

tmp1497 := MakeNative(func(__e *ControlFlow) {
Z3637 := __e.Get(1)
_ = Z3637
__e.Return(Nil)
return
}, 1)

tmp1498 := Call(__e, try_1catch, tmp1496, tmp1497)


__e.TailApply(tmp1492, tmp1498)
return


}, 1)

tmp1499 := Call(__e, PrimFunc(symlimit), V3634)


tmp1500 := Call(__e, PrimFunc(symhash), V3631, tmp1499)


__e.TailApply(tmp1491, tmp1500)
return


}, 4)

tmp1501 := Call(__e, ns2_1set, symput, tmp1490)


_ = tmp1501

tmp1502 := MakeNative(func(__e *ControlFlow) {
V3639 := __e.Get(1)
_ = V3639
V3640 := __e.Get(2)
_ = V3640
V3641 := __e.Get(3)
_ = V3641
tmp1503 := MakeNative(func(__e *ControlFlow) {
W3642 := __e.Get(1)
_ = W3642
tmp1504 := MakeNative(func(__e *ControlFlow) {
W3643 := __e.Get(1)
_ = W3643
tmp1505 := MakeNative(func(__e *ControlFlow) {
W3645 := __e.Get(1)
_ = W3645
__e.Return(V3639)
return
}, 1)

tmp1506 := Call(__e, PrimFunc(symshen_4remove_1pointer), V3639, V3640, W3643)


tmp1507 := Call(__e, PrimFunc(symvector_1_6), V3641, W3642, tmp1506)


__e.TailApply(tmp1505, tmp1507)
return


}, 1)

tmp1508 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3641, W3642)
return
}, 0)

tmp1509 := MakeNative(func(__e *ControlFlow) {
Z3644 := __e.Get(1)
_ = Z3644
__e.Return(Nil)
return
}, 1)

tmp1510 := Call(__e, try_1catch, tmp1508, tmp1509)


__e.TailApply(tmp1504, tmp1510)
return


}, 1)

tmp1511 := Call(__e, PrimFunc(symlimit), V3641)


tmp1512 := Call(__e, PrimFunc(symhash), V3639, tmp1511)


__e.TailApply(tmp1503, tmp1512)
return


}, 3)

tmp1513 := Call(__e, ns2_1set, symunput, tmp1502)


_ = tmp1513

tmp1514 := MakeNative(func(__e *ControlFlow) {
V3656 := __e.Get(1)
_ = V3656
V3657 := __e.Get(2)
_ = V3657
V3658 := __e.Get(3)
_ = V3658
tmp1558 := PrimEqual(Nil, V3658)

if True == tmp1558 {
__e.Return(Nil)
return
} else {
tmp1556 := PrimIsPair(V3658)

var ifres1521 Obj

if True == tmp1556 {
tmp1554 := PrimHead(V3658)

tmp1555 := PrimIsPair(tmp1554)

var ifres1523 Obj

if True == tmp1555 {
tmp1551 := PrimHead(V3658)

tmp1552 := PrimHead(tmp1551)

tmp1553 := PrimIsPair(tmp1552)

var ifres1525 Obj

if True == tmp1553 {
tmp1547 := PrimHead(V3658)

tmp1548 := PrimHead(tmp1547)

tmp1549 := PrimTail(tmp1548)

tmp1550 := PrimIsPair(tmp1549)

var ifres1527 Obj

if True == tmp1550 {
tmp1542 := PrimHead(V3658)

tmp1543 := PrimHead(tmp1542)

tmp1544 := PrimTail(tmp1543)

tmp1545 := PrimTail(tmp1544)

tmp1546 := PrimEqual(Nil, tmp1545)

var ifres1529 Obj

if True == tmp1546 {
tmp1537 := PrimHead(V3658)

tmp1538 := PrimHead(tmp1537)

tmp1539 := PrimTail(tmp1538)

tmp1540 := PrimHead(tmp1539)

tmp1541 := PrimEqual(V3657, tmp1540)

var ifres1531 Obj

if True == tmp1541 {
tmp1533 := PrimHead(V3658)

tmp1534 := PrimHead(tmp1533)

tmp1535 := PrimHead(tmp1534)

tmp1536 := PrimEqual(V3656, tmp1535)

var ifres1532 Obj

if True == tmp1536 {
ifres1532 = True


} else {
ifres1532 = False


}

ifres1531 = ifres1532


} else {
ifres1531 = False


}

var ifres1530 Obj

if True == ifres1531 {
ifres1530 = True


} else {
ifres1530 = False


}

ifres1529 = ifres1530


} else {
ifres1529 = False


}

var ifres1528 Obj

if True == ifres1529 {
ifres1528 = True


} else {
ifres1528 = False


}

ifres1527 = ifres1528


} else {
ifres1527 = False


}

var ifres1526 Obj

if True == ifres1527 {
ifres1526 = True


} else {
ifres1526 = False


}

ifres1525 = ifres1526


} else {
ifres1525 = False


}

var ifres1524 Obj

if True == ifres1525 {
ifres1524 = True


} else {
ifres1524 = False


}

ifres1523 = ifres1524


} else {
ifres1523 = False


}

var ifres1522 Obj

if True == ifres1523 {
ifres1522 = True


} else {
ifres1522 = False


}

ifres1521 = ifres1522


} else {
ifres1521 = False


}

if True == ifres1521 {
__e.Return(PrimTail(V3658))
return
} else {
tmp1519 := PrimIsPair(V3658)

if True == tmp1519 {
tmp1515 := PrimHead(V3658)

tmp1516 := PrimTail(V3658)

tmp1517 := Call(__e, PrimFunc(symshen_4remove_1pointer), V3656, V3657, tmp1516)


__e.Return(PrimCons(tmp1515, tmp1517))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-pointer")))
return
}


}


}


}, 3)

tmp1559 := Call(__e, ns2_1set, symshen_4remove_1pointer, tmp1514)


_ = tmp1559

tmp1560 := MakeNative(func(__e *ControlFlow) {
V3671 := __e.Get(1)
_ = V3671
V3672 := __e.Get(2)
_ = V3672
V3673 := __e.Get(3)
_ = V3673
V3674 := __e.Get(4)
_ = V3674
tmp1611 := PrimEqual(Nil, V3674)

if True == tmp1611 {
tmp1561 := PrimCons(V3672, Nil)

tmp1562 := PrimCons(V3671, tmp1561)

tmp1563 := PrimCons(tmp1562, V3673)

__e.Return(PrimCons(tmp1563, Nil))
return


} else {
tmp1609 := PrimIsPair(V3674)

var ifres1574 Obj

if True == tmp1609 {
tmp1607 := PrimHead(V3674)

tmp1608 := PrimIsPair(tmp1607)

var ifres1576 Obj

if True == tmp1608 {
tmp1604 := PrimHead(V3674)

tmp1605 := PrimHead(tmp1604)

tmp1606 := PrimIsPair(tmp1605)

var ifres1578 Obj

if True == tmp1606 {
tmp1600 := PrimHead(V3674)

tmp1601 := PrimHead(tmp1600)

tmp1602 := PrimTail(tmp1601)

tmp1603 := PrimIsPair(tmp1602)

var ifres1580 Obj

if True == tmp1603 {
tmp1595 := PrimHead(V3674)

tmp1596 := PrimHead(tmp1595)

tmp1597 := PrimTail(tmp1596)

tmp1598 := PrimTail(tmp1597)

tmp1599 := PrimEqual(Nil, tmp1598)

var ifres1582 Obj

if True == tmp1599 {
tmp1590 := PrimHead(V3674)

tmp1591 := PrimHead(tmp1590)

tmp1592 := PrimTail(tmp1591)

tmp1593 := PrimHead(tmp1592)

tmp1594 := PrimEqual(V3672, tmp1593)

var ifres1584 Obj

if True == tmp1594 {
tmp1586 := PrimHead(V3674)

tmp1587 := PrimHead(tmp1586)

tmp1588 := PrimHead(tmp1587)

tmp1589 := PrimEqual(V3671, tmp1588)

var ifres1585 Obj

if True == tmp1589 {
ifres1585 = True


} else {
ifres1585 = False


}

ifres1584 = ifres1585


} else {
ifres1584 = False


}

var ifres1583 Obj

if True == ifres1584 {
ifres1583 = True


} else {
ifres1583 = False


}

ifres1582 = ifres1583


} else {
ifres1582 = False


}

var ifres1581 Obj

if True == ifres1582 {
ifres1581 = True


} else {
ifres1581 = False


}

ifres1580 = ifres1581


} else {
ifres1580 = False


}

var ifres1579 Obj

if True == ifres1580 {
ifres1579 = True


} else {
ifres1579 = False


}

ifres1578 = ifres1579


} else {
ifres1578 = False


}

var ifres1577 Obj

if True == ifres1578 {
ifres1577 = True


} else {
ifres1577 = False


}

ifres1576 = ifres1577


} else {
ifres1576 = False


}

var ifres1575 Obj

if True == ifres1576 {
ifres1575 = True


} else {
ifres1575 = False


}

ifres1574 = ifres1575


} else {
ifres1574 = False


}

if True == ifres1574 {
tmp1564 := PrimHead(V3674)

tmp1565 := PrimHead(tmp1564)

tmp1566 := PrimCons(tmp1565, V3673)

tmp1567 := PrimTail(V3674)

__e.Return(PrimCons(tmp1566, tmp1567))
return


} else {
tmp1572 := PrimIsPair(V3674)

if True == tmp1572 {
tmp1568 := PrimHead(V3674)

tmp1569 := PrimTail(V3674)

tmp1570 := Call(__e, PrimFunc(symshen_4change_1pointer_1value), V3671, V3672, V3673, tmp1569)


__e.Return(PrimCons(tmp1568, tmp1570))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.change-pointer-value")))
return
}


}


}


}, 4)

tmp1612 := Call(__e, ns2_1set, symshen_4change_1pointer_1value, tmp1560)


_ = tmp1612

tmp1613 := MakeNative(func(__e *ControlFlow) {
V3675 := __e.Get(1)
_ = V3675
V3676 := __e.Get(2)
_ = V3676
V3677 := __e.Get(3)
_ = V3677
tmp1614 := MakeNative(func(__e *ControlFlow) {
W3678 := __e.Get(1)
_ = W3678
tmp1615 := MakeNative(func(__e *ControlFlow) {
W3679 := __e.Get(1)
_ = W3679
tmp1616 := MakeNative(func(__e *ControlFlow) {
W3681 := __e.Get(1)
_ = W3681
tmp1622 := Call(__e, PrimFunc(symempty_2), W3681)


if True == tmp1622 {
tmp1617 := Call(__e, PrimFunc(symshen_4app), V3675, MakeString("\n"), symshen_4s)


tmp1618 := PrimStringConcat(MakeString(" not found for "), tmp1617)

tmp1619 := Call(__e, PrimFunc(symshen_4app), V3676, tmp1618, symshen_4s)


tmp1620 := PrimStringConcat(MakeString("attribute "), tmp1619)

__e.Return(PrimSimpleError(tmp1620))
return


} else {
__e.Return(PrimTail(W3681))
return
}


}, 1)

tmp1623 := PrimCons(V3676, Nil)

tmp1624 := PrimCons(V3675, tmp1623)

tmp1625 := Call(__e, PrimFunc(symassoc), tmp1624, W3679)


__e.TailApply(tmp1616, tmp1625)
return


}, 1)

tmp1626 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3677, W3678)
return
}, 0)

tmp1627 := MakeNative(func(__e *ControlFlow) {
Z3680 := __e.Get(1)
_ = Z3680
tmp1628 := Call(__e, PrimFunc(symshen_4app), V3676, MakeString("\n"), symshen_4s)


tmp1629 := PrimStringConcat(MakeString(" has no attributes: "), tmp1628)

tmp1630 := Call(__e, PrimFunc(symshen_4app), V3675, tmp1629, symshen_4a)


__e.Return(PrimSimpleError(tmp1630))
return


}, 1)

tmp1631 := Call(__e, try_1catch, tmp1626, tmp1627)


__e.TailApply(tmp1615, tmp1631)
return


}, 1)

tmp1632 := Call(__e, PrimFunc(symlimit), V3677)


tmp1633 := Call(__e, PrimFunc(symhash), V3675, tmp1632)


__e.TailApply(tmp1614, tmp1633)
return


}, 3)

tmp1634 := Call(__e, ns2_1set, symget, tmp1613)


_ = tmp1634

tmp1635 := MakeNative(func(__e *ControlFlow) {
V3682 := __e.Get(1)
_ = V3682
V3683 := __e.Get(2)
_ = V3683
tmp1636 := MakeNative(func(__e *ControlFlow) {
W3684 := __e.Get(1)
_ = W3684
tmp1638 := PrimEqual(W3684, MakeNumber(0))

if True == tmp1638 {
__e.Return(MakeNumber(1))
return
} else {
__e.Return(W3684)
return
}


}, 1)

tmp1639 := Call(__e, PrimFunc(symshen_4hashkey), V3682)


tmp1640 := Call(__e, PrimFunc(symshen_4mod), tmp1639, V3683)


__e.TailApply(tmp1636, tmp1640)
return


}, 2)

tmp1641 := Call(__e, ns2_1set, symhash, tmp1635)


_ = tmp1641

tmp1642 := MakeNative(func(__e *ControlFlow) {
V3685 := __e.Get(1)
_ = V3685
tmp1643 := MakeNative(func(__e *ControlFlow) {
W3686 := __e.Get(1)
_ = W3686
__e.TailApply(PrimFunc(symshen_4prodbutzero), W3686, MakeNumber(1))
return
}, 1)

tmp1644 := MakeNative(func(__e *ControlFlow) {
Z3687 := __e.Get(1)
_ = Z3687
__e.Return(PrimStringToNumber(Z3687))
return
}, 1)

tmp1645 := Call(__e, PrimFunc(symexplode), V3685)


tmp1646 := Call(__e, PrimFunc(symmap), tmp1644, tmp1645)


__e.TailApply(tmp1643, tmp1646)
return


}, 1)

tmp1647 := Call(__e, ns2_1set, symshen_4hashkey, tmp1642)


_ = tmp1647

tmp1648 := MakeNative(func(__e *ControlFlow) {
V3688 := __e.Get(1)
_ = V3688
V3689 := __e.Get(2)
_ = V3689
tmp1667 := PrimEqual(Nil, V3688)

if True == tmp1667 {
__e.Return(V3689)
return
} else {
tmp1665 := PrimIsPair(V3688)

var ifres1661 Obj

if True == tmp1665 {
tmp1663 := PrimHead(V3688)

tmp1664 := PrimEqual(MakeNumber(0), tmp1663)

var ifres1662 Obj

if True == tmp1664 {
ifres1662 = True


} else {
ifres1662 = False


}

ifres1661 = ifres1662


} else {
ifres1661 = False


}

if True == ifres1661 {
tmp1649 := PrimTail(V3688)

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1649, V3689)
return


} else {
tmp1659 := PrimIsPair(V3688)

if True == tmp1659 {
tmp1657 := PrimGreatThan(V3689, MakeNumber(10000000000))

if True == tmp1657 {
tmp1650 := PrimTail(V3688)

tmp1651 := PrimHead(V3688)

tmp1652 := PrimNumberAdd(V3689, tmp1651)

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1650, tmp1652)
return


} else {
tmp1653 := PrimTail(V3688)

tmp1654 := PrimHead(V3688)

tmp1655 := PrimNumberMultiply(V3689, tmp1654)

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1653, tmp1655)
return


}


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.prodbutzero")))
return
}


}


}


}, 2)

tmp1668 := Call(__e, ns2_1set, symshen_4prodbutzero, tmp1648)


_ = tmp1668

tmp1669 := MakeNative(func(__e *ControlFlow) {
V3690 := __e.Get(1)
_ = V3690
V3691 := __e.Get(2)
_ = V3691
tmp1670 := PrimCons(V3691, Nil)

tmp1671 := Call(__e, PrimFunc(symshen_4multiples), V3690, tmp1670)


__e.TailApply(PrimFunc(symshen_4modh), V3690, tmp1671)
return


}, 2)

tmp1672 := Call(__e, ns2_1set, symshen_4mod, tmp1669)


_ = tmp1672

tmp1673 := MakeNative(func(__e *ControlFlow) {
V3696 := __e.Get(1)
_ = V3696
V3697 := __e.Get(2)
_ = V3697
tmp1684 := PrimIsPair(V3697)

var ifres1680 Obj

if True == tmp1684 {
tmp1682 := PrimHead(V3697)

tmp1683 := PrimGreatThan(tmp1682, V3696)

var ifres1681 Obj

if True == tmp1683 {
ifres1681 = True


} else {
ifres1681 = False


}

ifres1680 = ifres1681


} else {
ifres1680 = False


}

if True == ifres1680 {
__e.Return(PrimTail(V3697))
return
} else {
tmp1678 := PrimIsPair(V3697)

if True == tmp1678 {
tmp1674 := PrimHead(V3697)

tmp1675 := PrimNumberMultiply(MakeNumber(2), tmp1674)

tmp1676 := PrimCons(tmp1675, V3697)

__e.TailApply(PrimFunc(symshen_4multiples), V3696, tmp1676)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.multiples")))
return
}


}


}, 2)

tmp1685 := Call(__e, ns2_1set, symshen_4multiples, tmp1673)


_ = tmp1685

tmp1686 := MakeNative(func(__e *ControlFlow) {
V3704 := __e.Get(1)
_ = V3704
V3705 := __e.Get(2)
_ = V3705
tmp1704 := PrimEqual(MakeNumber(0), V3704)

if True == tmp1704 {
__e.Return(MakeNumber(0))
return
} else {
tmp1702 := PrimEqual(Nil, V3705)

if True == tmp1702 {
__e.Return(V3704)
return
} else {
tmp1700 := PrimIsPair(V3705)

var ifres1696 Obj

if True == tmp1700 {
tmp1698 := PrimHead(V3705)

tmp1699 := PrimGreatThan(tmp1698, V3704)

var ifres1697 Obj

if True == tmp1699 {
ifres1697 = True


} else {
ifres1697 = False


}

ifres1696 = ifres1697


} else {
ifres1696 = False


}

if True == ifres1696 {
tmp1689 := PrimTail(V3705)

tmp1690 := Call(__e, PrimFunc(symempty_2), tmp1689)


if True == tmp1690 {
__e.Return(V3704)
return
} else {
tmp1687 := PrimTail(V3705)

__e.TailApply(PrimFunc(symshen_4modh), V3704, tmp1687)
return


}


} else {
tmp1694 := PrimIsPair(V3705)

if True == tmp1694 {
tmp1691 := PrimHead(V3705)

tmp1692 := PrimNumberSubtract(V3704, tmp1691)

__e.TailApply(PrimFunc(symshen_4modh), tmp1692, V3705)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.modh")))
return
}


}


}


}


}, 2)

tmp1705 := Call(__e, ns2_1set, symshen_4modh, tmp1686)


_ = tmp1705

tmp1706 := MakeNative(func(__e *ControlFlow) {
V3708 := __e.Get(1)
_ = V3708
tmp1713 := PrimEqual(Nil, V3708)

if True == tmp1713 {
__e.Return(MakeNumber(0))
return
} else {
tmp1711 := PrimIsPair(V3708)

if True == tmp1711 {
tmp1707 := PrimHead(V3708)

tmp1708 := PrimTail(V3708)

tmp1709 := Call(__e, PrimFunc(symsum), tmp1708)


__e.Return(PrimNumberAdd(tmp1707, tmp1709))
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to sum a non-list\n")))
return
}


}


}, 1)

tmp1714 := Call(__e, ns2_1set, symsum, tmp1706)


_ = tmp1714

tmp1715 := MakeNative(func(__e *ControlFlow) {
V3713 := __e.Get(1)
_ = V3713
tmp1717 := PrimIsPair(V3713)

if True == tmp1717 {
__e.Return(PrimHead(V3713))
return
} else {
__e.Return(PrimSimpleError(MakeString("head expects a non-empty list\n")))
return
}


}, 1)

tmp1718 := Call(__e, ns2_1set, symhead, tmp1715)


_ = tmp1718

tmp1719 := MakeNative(func(__e *ControlFlow) {
V3718 := __e.Get(1)
_ = V3718
tmp1721 := PrimIsPair(V3718)

if True == tmp1721 {
__e.Return(PrimTail(V3718))
return
} else {
__e.Return(PrimSimpleError(MakeString("tail expects a non-empty list\n")))
return
}


}, 1)

tmp1722 := Call(__e, ns2_1set, symtail, tmp1719)


_ = tmp1722

tmp1723 := MakeNative(func(__e *ControlFlow) {
V3719 := __e.Get(1)
_ = V3719
__e.Return(PrimPos(V3719, MakeNumber(0)))
return
}, 1)

tmp1724 := Call(__e, ns2_1set, symhdstr, tmp1723)


_ = tmp1724

tmp1725 := MakeNative(func(__e *ControlFlow) {
V3726 := __e.Get(1)
_ = V3726
V3727 := __e.Get(2)
_ = V3727
tmp1736 := PrimEqual(Nil, V3726)

if True == tmp1736 {
__e.Return(Nil)
return
} else {
tmp1734 := PrimIsPair(V3726)

if True == tmp1734 {
tmp1731 := PrimHead(V3726)

tmp1732 := Call(__e, PrimFunc(symelement_2), tmp1731, V3727)


if True == tmp1732 {
tmp1726 := PrimHead(V3726)

tmp1727 := PrimTail(V3726)

tmp1728 := Call(__e, PrimFunc(symintersection), tmp1727, V3727)


__e.Return(PrimCons(tmp1726, tmp1728))
return


} else {
tmp1729 := PrimTail(V3726)

__e.TailApply(PrimFunc(symintersection), tmp1729, V3727)
return


}


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find the intersection with a non-list\n")))
return
}


}


}, 2)

tmp1737 := Call(__e, ns2_1set, symintersection, tmp1725)


_ = tmp1737

tmp1738 := MakeNative(func(__e *ControlFlow) {
V3728 := __e.Get(1)
_ = V3728
__e.TailApply(PrimFunc(symshen_4reverse_1help), V3728, Nil)
return
}, 1)

tmp1739 := Call(__e, ns2_1set, symreverse, tmp1738)


_ = tmp1739

tmp1740 := MakeNative(func(__e *ControlFlow) {
V3733 := __e.Get(1)
_ = V3733
V3734 := __e.Get(2)
_ = V3734
tmp1747 := PrimEqual(Nil, V3733)

if True == tmp1747 {
__e.Return(V3734)
return
} else {
tmp1745 := PrimIsPair(V3733)

if True == tmp1745 {
tmp1741 := PrimTail(V3733)

tmp1742 := PrimHead(V3733)

tmp1743 := PrimCons(tmp1742, V3734)

__e.TailApply(PrimFunc(symshen_4reverse_1help), tmp1741, tmp1743)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to reverse a non-list\n")))
return
}


}


}, 2)

tmp1748 := Call(__e, ns2_1set, symshen_4reverse_1help, tmp1740)


_ = tmp1748

tmp1749 := MakeNative(func(__e *ControlFlow) {
V3739 := __e.Get(1)
_ = V3739
V3740 := __e.Get(2)
_ = V3740
tmp1760 := PrimEqual(Nil, V3739)

if True == tmp1760 {
__e.Return(V3740)
return
} else {
tmp1758 := PrimIsPair(V3739)

if True == tmp1758 {
tmp1755 := PrimHead(V3739)

tmp1756 := Call(__e, PrimFunc(symelement_2), tmp1755, V3740)


if True == tmp1756 {
tmp1750 := PrimTail(V3739)

__e.TailApply(PrimFunc(symunion), tmp1750, V3740)
return


} else {
tmp1751 := PrimHead(V3739)

tmp1752 := PrimTail(V3739)

tmp1753 := Call(__e, PrimFunc(symunion), tmp1752, V3740)


__e.Return(PrimCons(tmp1751, tmp1753))
return


}


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find the union with a non-list\n")))
return
}


}


}, 2)

tmp1761 := Call(__e, ns2_1set, symunion, tmp1749)


_ = tmp1761

tmp1762 := MakeNative(func(__e *ControlFlow) {
V3741 := __e.Get(1)
_ = V3741
tmp1763 := MakeNative(func(__e *ControlFlow) {
W3742 := __e.Get(1)
_ = W3742
tmp1764 := MakeNative(func(__e *ControlFlow) {
W3743 := __e.Get(1)
_ = W3743
tmp1765 := MakeNative(func(__e *ControlFlow) {
W3744 := __e.Get(1)
_ = W3744
tmp1771 := PrimEqual(MakeString("y"), W3744)

if True == tmp1771 {
__e.Return(True)
return
} else {
tmp1769 := PrimEqual(MakeString("n"), W3744)

if True == tmp1769 {
__e.Return(False)
return
} else {
tmp1766 := Call(__e, PrimFunc(symstoutput))


tmp1767 := Call(__e, PrimFunc(sympr), MakeString("please answer y or n\n"), tmp1766)


_ = tmp1767

__e.TailApply(PrimFunc(symy_1or_1n_2), V3741)
return


}


}


}, 1)

tmp1772 := Call(__e, PrimFunc(symstinput))


tmp1773 := Call(__e, PrimFunc(symread), tmp1772)


tmp1774 := Call(__e, PrimFunc(symshen_4app), tmp1773, MakeString(""), symshen_4s)


__e.TailApply(tmp1765, tmp1774)
return


}, 1)

tmp1775 := Call(__e, PrimFunc(symstoutput))


tmp1776 := Call(__e, PrimFunc(sympr), MakeString(" (y/n) "), tmp1775)


__e.TailApply(tmp1764, tmp1776)
return


}, 1)

tmp1777 := Call(__e, PrimFunc(symshen_4proc_1nl), V3741)


tmp1778 := Call(__e, PrimFunc(symstoutput))


tmp1779 := Call(__e, PrimFunc(sympr), tmp1777, tmp1778)


__e.TailApply(tmp1763, tmp1779)
return


}, 1)

tmp1780 := Call(__e, ns2_1set, symy_1or_1n_2, tmp1762)


_ = tmp1780

tmp1781 := MakeNative(func(__e *ControlFlow) {
V3745 := __e.Get(1)
_ = V3745
if True == V3745 {
__e.Return(False)
return
} else {
__e.Return(True)
return
}
}, 1)

tmp1783 := Call(__e, ns2_1set, symnot, tmp1781)


_ = tmp1783

tmp1784 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimSimpleError(MakeString("")))
return
}, 0)

tmp1785 := Call(__e, ns2_1set, symabort, tmp1784)


_ = tmp1785

tmp1786 := MakeNative(func(__e *ControlFlow) {
V3751 := __e.Get(1)
_ = V3751
V3752 := __e.Get(2)
_ = V3752
V3753 := __e.Get(3)
_ = V3753
tmp1794 := PrimEqual(V3752, V3753)

if True == tmp1794 {
__e.Return(V3751)
return
} else {
tmp1792 := PrimIsPair(V3753)

if True == tmp1792 {
tmp1787 := PrimHead(V3753)

tmp1788 := Call(__e, PrimFunc(symsubst), V3751, V3752, tmp1787)


tmp1789 := PrimTail(V3753)

tmp1790 := Call(__e, PrimFunc(symsubst), V3751, V3752, tmp1789)


__e.Return(PrimCons(tmp1788, tmp1790))
return


} else {
__e.Return(V3753)
return
}


}


}, 3)

tmp1795 := Call(__e, ns2_1set, symsubst, tmp1786)


_ = tmp1795

tmp1796 := MakeNative(func(__e *ControlFlow) {
V3754 := __e.Get(1)
_ = V3754
tmp1797 := Call(__e, PrimFunc(symshen_4app), V3754, MakeString(""), symshen_4a)


__e.TailApply(PrimFunc(symshen_4explode_1h), tmp1797)
return


}, 1)

tmp1798 := Call(__e, ns2_1set, symexplode, tmp1796)


_ = tmp1798

tmp1799 := MakeNative(func(__e *ControlFlow) {
V3757 := __e.Get(1)
_ = V3757
tmp1806 := PrimEqual(MakeString(""), V3757)

if True == tmp1806 {
__e.Return(Nil)
return
} else {
tmp1804 := Call(__e, PrimFunc(symshen_4_7string_2), V3757)


if True == tmp1804 {
tmp1800 := Call(__e, PrimFunc(symhdstr), V3757)


tmp1801 := PrimTailString(V3757)

tmp1802 := Call(__e, PrimFunc(symshen_4explode_1h), tmp1801)


__e.Return(PrimCons(tmp1800, tmp1802))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in explode-h")))
return
}


}


}, 1)

tmp1807 := Call(__e, ns2_1set, symshen_4explode_1h, tmp1799)


_ = tmp1807

tmp1808 := MakeNative(func(__e *ControlFlow) {
V3758 := __e.Get(1)
_ = V3758
tmp1811 := PrimEqual(V3758, MakeString(""))

var ifres1809 Obj

if True == tmp1811 {
ifres1809 = MakeString("")


} else {
tmp1810 := Call(__e, PrimFunc(symshen_4app), V3758, MakeString("/"), symshen_4a)


ifres1809 = tmp1810


}

__e.Return(PrimSet(sym_dhome_1directory_d, ifres1809))
return


}, 1)

tmp1812 := Call(__e, ns2_1set, symcd, tmp1808)


_ = tmp1812

tmp1813 := MakeNative(func(__e *ControlFlow) {
V3759 := __e.Get(1)
_ = V3759
V3760 := __e.Get(2)
_ = V3760
__e.TailApply(PrimFunc(symshen_4map_1h), V3759, V3760, Nil)
return
}, 2)

tmp1814 := Call(__e, ns2_1set, symmap, tmp1813)


_ = tmp1814

tmp1815 := MakeNative(func(__e *ControlFlow) {
V3761 := __e.Get(1)
_ = V3761
V3762 := __e.Get(2)
_ = V3762
V3763 := __e.Get(3)
_ = V3763
tmp1823 := PrimEqual(Nil, V3762)

if True == tmp1823 {
__e.TailApply(PrimFunc(symreverse), V3763)
return
} else {
tmp1821 := PrimIsPair(V3762)

if True == tmp1821 {
tmp1816 := PrimTail(V3762)

tmp1817 := PrimHead(V3762)

tmp1818 := Call(__e, V3761, tmp1817)


tmp1819 := PrimCons(tmp1818, V3763)

__e.TailApply(PrimFunc(symshen_4map_1h), V3761, tmp1816, tmp1819)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.map-h")))
return
}


}


}, 3)

tmp1824 := Call(__e, ns2_1set, symshen_4map_1h, tmp1815)


_ = tmp1824

tmp1825 := MakeNative(func(__e *ControlFlow) {
V3764 := __e.Get(1)
_ = V3764
__e.TailApply(PrimFunc(symshen_4length_1h), V3764, MakeNumber(0))
return
}, 1)

tmp1826 := Call(__e, ns2_1set, symlength, tmp1825)


_ = tmp1826

tmp1827 := MakeNative(func(__e *ControlFlow) {
V3769 := __e.Get(1)
_ = V3769
V3770 := __e.Get(2)
_ = V3770
tmp1831 := PrimEqual(Nil, V3769)

if True == tmp1831 {
__e.Return(V3770)
return
} else {
tmp1828 := PrimTail(V3769)

tmp1829 := PrimNumberAdd(V3770, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4length_1h), tmp1828, tmp1829)
return


}


}, 2)

tmp1832 := Call(__e, ns2_1set, symshen_4length_1h, tmp1827)


_ = tmp1832

tmp1833 := MakeNative(func(__e *ControlFlow) {
V3776 := __e.Get(1)
_ = V3776
V3777 := __e.Get(2)
_ = V3777
tmp1841 := PrimEqual(V3776, V3777)

if True == tmp1841 {
__e.Return(MakeNumber(1))
return
} else {
tmp1839 := PrimIsPair(V3777)

if True == tmp1839 {
tmp1834 := PrimHead(V3777)

tmp1835 := Call(__e, PrimFunc(symoccurrences), V3776, tmp1834)


tmp1836 := PrimTail(V3777)

tmp1837 := Call(__e, PrimFunc(symoccurrences), V3776, tmp1836)


__e.Return(PrimNumberAdd(tmp1835, tmp1837))
return


} else {
__e.Return(MakeNumber(0))
return
}


}


}, 2)

tmp1842 := Call(__e, ns2_1set, symoccurrences, tmp1833)


_ = tmp1842

tmp1843 := MakeNative(func(__e *ControlFlow) {
V3782 := __e.Get(1)
_ = V3782
V3783 := __e.Get(2)
_ = V3783
tmp1856 := PrimEqual(MakeNumber(1), V3782)

var ifres1853 Obj

if True == tmp1856 {
tmp1855 := PrimIsPair(V3783)

var ifres1854 Obj

if True == tmp1855 {
ifres1854 = True


} else {
ifres1854 = False


}

ifres1853 = ifres1854


} else {
ifres1853 = False


}

if True == ifres1853 {
__e.Return(PrimHead(V3783))
return
} else {
tmp1851 := PrimIsPair(V3783)

if True == tmp1851 {
tmp1844 := PrimNumberSubtract(V3782, MakeNumber(1))

tmp1845 := PrimTail(V3783)

__e.TailApply(PrimFunc(symnth), tmp1844, tmp1845)
return


} else {
tmp1846 := Call(__e, PrimFunc(symshen_4app), V3783, MakeString("\n"), symshen_4a)


tmp1847 := PrimStringConcat(MakeString(", "), tmp1846)

tmp1848 := Call(__e, PrimFunc(symshen_4app), V3782, tmp1847, symshen_4a)


tmp1849 := PrimStringConcat(MakeString("nth applied to "), tmp1848)

__e.Return(PrimSimpleError(tmp1849))
return


}


}


}, 2)

tmp1857 := Call(__e, ns2_1set, symnth, tmp1843)


_ = tmp1857

tmp1858 := MakeNative(func(__e *ControlFlow) {
V3784 := __e.Get(1)
_ = V3784
tmp1865 := PrimIsNumber(V3784)

if True == tmp1865 {
tmp1860 := MakeNative(func(__e *ControlFlow) {
W3785 := __e.Get(1)
_ = W3785
tmp1861 := Call(__e, PrimFunc(symshen_4magless), W3785, MakeNumber(1))


__e.TailApply(PrimFunc(symshen_4integer_1test_2), W3785, tmp1861)
return


}, 1)

tmp1862 := Call(__e, PrimFunc(symshen_4abs), V3784)


tmp1863 := Call(__e, tmp1860, tmp1862)


if True == tmp1863 {
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

tmp1866 := Call(__e, ns2_1set, syminteger_2, tmp1858)


_ = tmp1866

tmp1867 := MakeNative(func(__e *ControlFlow) {
V3786 := __e.Get(1)
_ = V3786
tmp1869 := PrimGreatThan(V3786, MakeNumber(0))

if True == tmp1869 {
__e.Return(V3786)
return
} else {
__e.Return(PrimNumberSubtract(MakeNumber(0), V3786))
return
}


}, 1)

tmp1870 := Call(__e, ns2_1set, symshen_4abs, tmp1867)


_ = tmp1870

tmp1871 := MakeNative(func(__e *ControlFlow) {
V3787 := __e.Get(1)
_ = V3787
V3788 := __e.Get(2)
_ = V3788
tmp1872 := MakeNative(func(__e *ControlFlow) {
W3789 := __e.Get(1)
_ = W3789
tmp1874 := PrimGreatThan(W3789, V3787)

if True == tmp1874 {
__e.Return(V3788)
return
} else {
__e.TailApply(PrimFunc(symshen_4magless), V3787, W3789)
return
}


}, 1)

tmp1875 := PrimNumberMultiply(V3788, MakeNumber(2))

__e.TailApply(tmp1872, tmp1875)
return


}, 2)

tmp1876 := Call(__e, ns2_1set, symshen_4magless, tmp1871)


_ = tmp1876

tmp1877 := MakeNative(func(__e *ControlFlow) {
V3793 := __e.Get(1)
_ = V3793
V3794 := __e.Get(2)
_ = V3794
tmp1885 := PrimEqual(MakeNumber(0), V3793)

if True == tmp1885 {
__e.Return(True)
return
} else {
tmp1883 := PrimGreatThan(MakeNumber(1), V3793)

if True == tmp1883 {
__e.Return(False)
return
} else {
tmp1878 := MakeNative(func(__e *ControlFlow) {
W3795 := __e.Get(1)
_ = W3795
tmp1880 := PrimGreatThan(MakeNumber(0), W3795)

if True == tmp1880 {
__e.Return(PrimIsInteger(V3793))
return
} else {
__e.TailApply(PrimFunc(symshen_4integer_1test_2), W3795, V3794)
return
}


}, 1)

tmp1881 := PrimNumberSubtract(V3793, V3794)

__e.TailApply(tmp1878, tmp1881)
return


}


}


}, 2)

tmp1886 := Call(__e, ns2_1set, symshen_4integer_1test_2, tmp1877)


_ = tmp1886

tmp1887 := MakeNative(func(__e *ControlFlow) {
V3802 := __e.Get(1)
_ = V3802
V3803 := __e.Get(2)
_ = V3803
tmp1895 := PrimEqual(Nil, V3803)

if True == tmp1895 {
__e.Return(Nil)
return
} else {
tmp1893 := PrimIsPair(V3803)

if True == tmp1893 {
tmp1888 := PrimHead(V3803)

tmp1889 := Call(__e, V3802, tmp1888)


tmp1890 := PrimTail(V3803)

tmp1891 := Call(__e, PrimFunc(symmapcan), V3802, tmp1890)


__e.TailApply(PrimFunc(symappend), tmp1889, tmp1891)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to mapcan over a non-list\n")))
return
}


}


}, 2)

tmp1896 := Call(__e, ns2_1set, symmapcan, tmp1887)


_ = tmp1896

tmp1897 := MakeNative(func(__e *ControlFlow) {
V3809 := __e.Get(1)
_ = V3809
V3810 := __e.Get(2)
_ = V3810
tmp1899 := PrimEqual(V3809, V3810)

if True == tmp1899 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp1900 := Call(__e, ns2_1set, sym_a_a, tmp1897)


_ = tmp1900

tmp1901 := MakeNative(func(__e *ControlFlow) {
V3811 := __e.Get(1)
_ = V3811
tmp1911 := PrimIsSymbol(V3811)

if True == tmp1911 {
tmp1903 := MakeNative(func(__e *ControlFlow) {
W3812 := __e.Get(1)
_ = W3812
tmp1905 := PrimEqual(W3812, symshen_4this_1symbol_1is_1unbound)

if True == tmp1905 {
__e.Return(False)
return
} else {
__e.Return(True)
return
}


}, 1)

tmp1906 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(V3811))
return
}, 0)

tmp1907 := MakeNative(func(__e *ControlFlow) {
Z3813 := __e.Get(1)
_ = Z3813
__e.Return(symshen_4this_1symbol_1is_1unbound)
return
}, 1)

tmp1908 := Call(__e, try_1catch, tmp1906, tmp1907)


tmp1909 := Call(__e, tmp1903, tmp1908)


if True == tmp1909 {
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

tmp1912 := Call(__e, ns2_1set, symbound_2, tmp1901)


_ = tmp1912

tmp1913 := MakeNative(func(__e *ControlFlow) {
V3814 := __e.Get(1)
_ = V3814
tmp1919 := PrimEqual(MakeString(""), V3814)

if True == tmp1919 {
__e.Return(Nil)
return
} else {
tmp1914 := PrimPos(V3814, MakeNumber(0))

tmp1915 := PrimStringToNumber(tmp1914)

tmp1916 := PrimTailString(V3814)

tmp1917 := Call(__e, PrimFunc(symshen_4string_1_6bytes), tmp1916)


__e.Return(PrimCons(tmp1915, tmp1917))
return


}


}, 1)

tmp1920 := Call(__e, ns2_1set, symshen_4string_1_6bytes, tmp1913)


_ = tmp1920

tmp1921 := MakeNative(func(__e *ControlFlow) {
V3815 := __e.Get(1)
_ = V3815
tmp1925 := PrimLessThan(V3815, MakeNumber(0))

if True == tmp1925 {
__e.Return(PrimValue(symshen_4_dmaxinferences_d))
return
} else {
tmp1923 := PrimIsInteger(V3815)

if True == tmp1923 {
__e.Return(PrimSet(symshen_4_dmaxinferences_d, V3815))
return
} else {
__e.Return(PrimSimpleError(MakeString("maxinferences expects an integer value\n")))
return
}


}


}, 1)

tmp1926 := Call(__e, ns2_1set, symmaxinferences, tmp1921)


_ = tmp1926

tmp1927 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dinfs_d))
return
}, 0)

tmp1928 := Call(__e, ns2_1set, syminferences, tmp1927)


_ = tmp1928

tmp1929 := MakeNative(func(__e *ControlFlow) {
V3816 := __e.Get(1)
_ = V3816
__e.Return(V3816)
return
}, 1)

tmp1930 := Call(__e, ns2_1set, symprotect, tmp1929)


_ = tmp1930

tmp1931 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dstoutput_d))
return
}, 0)

tmp1932 := Call(__e, ns2_1set, symstoutput, tmp1931)


_ = tmp1932

tmp1933 := MakeNative(func(__e *ControlFlow) {
V3817 := __e.Get(1)
_ = V3817
tmp1934 := MakeNative(func(__e *ControlFlow) {
W3818 := __e.Get(1)
_ = W3818
tmp1938 := PrimIsSymbol(W3818)

if True == tmp1938 {
__e.Return(W3818)
return
} else {
tmp1935 := Call(__e, PrimFunc(symshen_4app), V3817, MakeString(" to a symbol"), symshen_4s)


tmp1936 := PrimStringConcat(MakeString("cannot intern "), tmp1935)

__e.Return(PrimSimpleError(tmp1936))
return


}


}, 1)

tmp1939 := PrimIntern(V3817)

__e.TailApply(tmp1934, tmp1939)
return


}, 1)

tmp1940 := Call(__e, ns2_1set, symstring_1_6symbol, tmp1933)


_ = tmp1940

tmp1941 := MakeNative(func(__e *ControlFlow) {
V3821 := __e.Get(1)
_ = V3821
tmp1945 := PrimEqual(sym_7, V3821)

if True == tmp1945 {
__e.Return(PrimSet(symshen_4_doptimise_d, True))
return
} else {
tmp1943 := PrimEqual(sym_1, V3821)

if True == tmp1943 {
__e.Return(PrimSet(symshen_4_doptimise_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("optimise expects a + or a -.\n")))
return
}


}


}, 1)

tmp1946 := Call(__e, ns2_1set, symoptimise, tmp1941)


_ = tmp1946

tmp1947 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dos_d))
return
}, 0)

tmp1948 := Call(__e, ns2_1set, symos, tmp1947)


_ = tmp1948

tmp1949 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dlanguage_d))
return
}, 0)

tmp1950 := Call(__e, ns2_1set, symlanguage, tmp1949)


_ = tmp1950

tmp1951 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dversion_d))
return
}, 0)

tmp1952 := Call(__e, ns2_1set, symversion, tmp1951)


_ = tmp1952

tmp1953 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dport_d))
return
}, 0)

tmp1954 := Call(__e, ns2_1set, symport, tmp1953)


_ = tmp1954

tmp1955 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dporters_d))
return
}, 0)

tmp1956 := Call(__e, ns2_1set, symporters, tmp1955)


_ = tmp1956

tmp1957 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dimplementation_d))
return
}, 0)

tmp1958 := Call(__e, ns2_1set, symimplementation, tmp1957)


_ = tmp1958

tmp1959 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_drelease_d))
return
}, 0)

tmp1960 := Call(__e, ns2_1set, symrelease, tmp1959)


_ = tmp1960

tmp1961 := MakeNative(func(__e *ControlFlow) {
V3822 := __e.Get(1)
_ = V3822
tmp1966 := PrimEqual(symnull, V3822)

if True == tmp1966 {
__e.Return(True)
return
} else {
tmp1962 := MakeNative(func(__e *ControlFlow) {
tmp1963 := Call(__e, PrimFunc(symexternal), V3822)


_ = tmp1963

__e.Return(True)
return


}, 0)

tmp1964 := MakeNative(func(__e *ControlFlow) {
Z3823 := __e.Get(1)
_ = Z3823
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1962, tmp1964)
return


}


}, 1)

tmp1967 := Call(__e, ns2_1set, sympackage_2, tmp1961)


_ = tmp1967

tmp1968 := MakeNative(func(__e *ControlFlow) {
__e.Return(sym_4_4_4)
return
}, 0)

tmp1969 := Call(__e, ns2_1set, symfail, tmp1968)


_ = tmp1969

tmp1970 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_duserdefs_d))
return
}, 0)

tmp1971 := Call(__e, ns2_1set, symuserdefs, tmp1970)


_ = tmp1971

tmp1972 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_doptimise_d))
return
}, 0)

tmp1973 := Call(__e, ns2_1set, symoptimise_2, tmp1972)


_ = tmp1973

tmp1974 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dhush_d))
return
}, 0)

tmp1975 := Call(__e, ns2_1set, symhush_2, tmp1974)


_ = tmp1975

tmp1976 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d))
return
}, 0)

tmp1977 := Call(__e, ns2_1set, symsystem_1S_2, tmp1976)


_ = tmp1977

tmp1978 := MakeNative(func(__e *ControlFlow) {
V3826 := __e.Get(1)
_ = V3826
tmp1982 := PrimEqual(sym_7, V3826)

if True == tmp1982 {
__e.Return(PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True))
return
} else {
tmp1980 := PrimEqual(sym_1, V3826)

if True == tmp1980 {
__e.Return(PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("enable-type-theory expects a + or a -\n")))
return
}


}


}, 1)

tmp1983 := Call(__e, ns2_1set, symenable_1type_1theory, tmp1978)


_ = tmp1983

tmp1984 := MakeNative(func(__e *ControlFlow) {
V3829 := __e.Get(1)
_ = V3829
tmp1988 := PrimEqual(sym_7, V3829)

if True == tmp1988 {
__e.Return(PrimSet(sym_dhush_d, True))
return
} else {
tmp1986 := PrimEqual(sym_1, V3829)

if True == tmp1986 {
__e.Return(PrimSet(sym_dhush_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("hush expects a + or a -\n")))
return
}


}


}, 1)

tmp1989 := Call(__e, ns2_1set, symshen_4hush, tmp1984)


_ = tmp1989

tmp1990 := MakeNative(func(__e *ControlFlow) {
V3832 := __e.Get(1)
_ = V3832
tmp1994 := PrimEqual(sym_7, V3832)

if True == tmp1994 {
__e.Return(PrimSet(symshen_4_dtc_d, True))
return
} else {
tmp1992 := PrimEqual(sym_1, V3832)

if True == tmp1992 {
__e.Return(PrimSet(symshen_4_dtc_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("tc expects a + or -")))
return
}


}


}, 1)

tmp1995 := Call(__e, ns2_1set, symtc, tmp1990)


_ = tmp1995

tmp1996 := MakeNative(func(__e *ControlFlow) {
V3833 := __e.Get(1)
_ = V3833
tmp1997 := PrimValue(symshen_4_dsigf_d)

tmp1998 := Call(__e, PrimFunc(symshen_4unassoc), V3833, tmp1997)


tmp1999 := PrimSet(symshen_4_dsigf_d, tmp1998)

_ = tmp1999

__e.Return(V3833)
return


}, 1)

tmp2000 := Call(__e, ns2_1set, symdestroy, tmp1996)


_ = tmp2000

tmp2001 := MakeNative(func(__e *ControlFlow) {
V3843 := __e.Get(1)
_ = V3843
V3844 := __e.Get(2)
_ = V3844
tmp2019 := PrimEqual(Nil, V3844)

if True == tmp2019 {
__e.Return(Nil)
return
} else {
tmp2017 := PrimIsPair(V3844)

var ifres2008 Obj

if True == tmp2017 {
tmp2015 := PrimHead(V3844)

tmp2016 := PrimIsPair(tmp2015)

var ifres2010 Obj

if True == tmp2016 {
tmp2012 := PrimHead(V3844)

tmp2013 := PrimHead(tmp2012)

tmp2014 := PrimEqual(V3843, tmp2013)

var ifres2011 Obj

if True == tmp2014 {
ifres2011 = True


} else {
ifres2011 = False


}

ifres2010 = ifres2011


} else {
ifres2010 = False


}

var ifres2009 Obj

if True == ifres2010 {
ifres2009 = True


} else {
ifres2009 = False


}

ifres2008 = ifres2009


} else {
ifres2008 = False


}

if True == ifres2008 {
__e.Return(PrimTail(V3844))
return
} else {
tmp2006 := PrimIsPair(V3844)

if True == tmp2006 {
tmp2002 := PrimHead(V3844)

tmp2003 := PrimTail(V3844)

tmp2004 := Call(__e, PrimFunc(symshen_4unassoc), V3843, tmp2003)


__e.Return(PrimCons(tmp2002, tmp2004))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.unassoc")))
return
}


}


}


}, 2)

tmp2020 := Call(__e, ns2_1set, symshen_4unassoc, tmp2001)


_ = tmp2020

tmp2021 := MakeNative(func(__e *ControlFlow) {
V3845 := __e.Get(1)
_ = V3845
tmp2025 := Call(__e, PrimFunc(sympackage_2), V3845)


if True == tmp2025 {
__e.Return(PrimSet(symshen_4_dpackage_d, V3845))
return
} else {
tmp2022 := Call(__e, PrimFunc(symshen_4app), V3845, MakeString(" does not exist\n"), symshen_4a)


tmp2023 := PrimStringConcat(MakeString("package "), tmp2022)

__e.Return(PrimSimpleError(tmp2023))
return


}


}, 1)

tmp2026 := Call(__e, ns2_1set, symin_1package, tmp2021)


_ = tmp2026

tmp2027 := MakeNative(func(__e *ControlFlow) {
V3846 := __e.Get(1)
_ = V3846
V3847 := __e.Get(2)
_ = V3847
tmp2028 := MakeNative(func(__e *ControlFlow) {
W3848 := __e.Get(1)
_ = W3848
tmp2029 := MakeNative(func(__e *ControlFlow) {
W3849 := __e.Get(1)
_ = W3849
tmp2030 := MakeNative(func(__e *ControlFlow) {
W3850 := __e.Get(1)
_ = W3850
tmp2031 := MakeNative(func(__e *ControlFlow) {
W3851 := __e.Get(1)
_ = W3851
__e.Return(V3847)
return
}, 1)

tmp2032 := PrimCloseStream(W3848)

__e.TailApply(tmp2031, tmp2032)
return


}, 1)

tmp2033 := Call(__e, PrimFunc(sympr), W3849, W3848)


__e.TailApply(tmp2030, tmp2033)
return


}, 1)

tmp2036 := PrimIsString(V3847)

var ifres2034 Obj

if True == tmp2036 {
ifres2034 = V3847


} else {
tmp2035 := Call(__e, PrimFunc(symshen_4app), V3847, MakeString(""), symshen_4s)


ifres2034 = tmp2035


}

__e.TailApply(tmp2029, ifres2034)
return


}, 1)

tmp2037 := PrimOpenStream(V3846, symout)

__e.TailApply(tmp2028, tmp2037)
return


}, 2)

tmp2038 := Call(__e, ns2_1set, symwrite_1to_1file, tmp2027)


_ = tmp2038

tmp2039 := MakeNative(func(__e *ControlFlow) {
tmp2040 := Call(__e, PrimFunc(symgensym), symshen_4t)


__e.TailApply(PrimFunc(symshen_4freshterm), tmp2040)
return


}, 0)

tmp2041 := Call(__e, ns2_1set, symfresh, tmp2039)


_ = tmp2041

tmp2042 := MakeNative(func(__e *ControlFlow) {
V3852 := __e.Get(1)
_ = V3852
V3853 := __e.Get(2)
_ = V3853
tmp2043 := MakeNative(func(__e *ControlFlow) {
W3854 := __e.Get(1)
_ = W3854
tmp2044 := MakeNative(func(__e *ControlFlow) {
W3855 := __e.Get(1)
_ = W3855
tmp2045 := MakeNative(func(__e *ControlFlow) {
W3856 := __e.Get(1)
_ = W3856
__e.Return(V3852)
return
}, 1)

tmp2046 := PrimValue(symshen_4_dlambdatable_d)

tmp2047 := PrimCons(W3855, tmp2046)

tmp2048 := PrimSet(symshen_4_dlambdatable_d, tmp2047)

__e.TailApply(tmp2045, tmp2048)
return


}, 1)

tmp2049 := Call(__e, PrimFunc(symshen_4lambda_1entry), V3852)


__e.TailApply(tmp2044, tmp2049)
return


}, 1)

tmp2050 := PrimValue(sym_dproperty_1vector_d)

tmp2051 := Call(__e, PrimFunc(symput), V3852, symarity, V3853, tmp2050)


__e.TailApply(tmp2043, tmp2051)
return


}, 2)

tmp2052 := Call(__e, ns2_1set, symupdate_1lambda_1table, tmp2042)


_ = tmp2052

tmp2053 := MakeNative(func(__e *ControlFlow) {
V3859 := __e.Get(1)
_ = V3859
V3860 := __e.Get(2)
_ = V3860
tmp2077 := PrimEqual(MakeNumber(0), V3860)

if True == tmp2077 {
tmp2054 := PrimValue(symshen_4_dspecial_d)

tmp2055 := Call(__e, PrimFunc(symremove), V3859, tmp2054)


tmp2056 := PrimSet(symshen_4_dspecial_d, tmp2055)

_ = tmp2056

tmp2057 := PrimValue(symshen_4_dextraspecial_d)

tmp2058 := Call(__e, PrimFunc(symremove), V3859, tmp2057)


tmp2059 := PrimSet(symshen_4_dextraspecial_d, tmp2058)

_ = tmp2059

__e.Return(V3859)
return


} else {
tmp2075 := PrimEqual(MakeNumber(1), V3860)

if True == tmp2075 {
tmp2060 := PrimValue(symshen_4_dspecial_d)

tmp2061 := Call(__e, PrimFunc(symadjoin), V3859, tmp2060)


tmp2062 := PrimSet(symshen_4_dspecial_d, tmp2061)

_ = tmp2062

tmp2063 := PrimValue(symshen_4_dextraspecial_d)

tmp2064 := Call(__e, PrimFunc(symremove), V3859, tmp2063)


tmp2065 := PrimSet(symshen_4_dextraspecial_d, tmp2064)

_ = tmp2065

__e.Return(V3859)
return


} else {
tmp2073 := PrimEqual(MakeNumber(2), V3860)

if True == tmp2073 {
tmp2066 := PrimValue(symshen_4_dspecial_d)

tmp2067 := Call(__e, PrimFunc(symremove), V3859, tmp2066)


tmp2068 := PrimSet(symshen_4_dspecial_d, tmp2067)

_ = tmp2068

tmp2069 := PrimValue(symshen_4_dextraspecial_d)

tmp2070 := Call(__e, PrimFunc(symadjoin), V3859, tmp2069)


tmp2071 := PrimSet(symshen_4_dextraspecial_d, tmp2070)

_ = tmp2071

__e.Return(V3859)
return


} else {
__e.Return(PrimSimpleError(MakeString("specialise requires values of 0, 1 or 2\n")))
return
}


}


}


}, 2)

__e.TailApply(ns2_1set, symspecialise, tmp2053)
return




}, 0)

