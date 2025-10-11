package main

import . "github.com/tiancaiamao/shen-go/kl"

var TopLevelMain = MakeNative(func(__e *ControlFlow) {
tmp8422 := MakeNative(func(__e *ControlFlow) {
tmp8423 := Call(__e, PrimFunc(symshen_4credits))


_ = tmp8423

__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp8424 := Call(__e, ns2_1set, symshen_4shen, tmp8422)


_ = tmp8424

tmp8425 := MakeNative(func(__e *ControlFlow) {
tmp8426 := Call(__e, PrimFunc(symshen_4initialise__environment))


_ = tmp8426

tmp8427 := Call(__e, PrimFunc(symshen_4prompt))


_ = tmp8427

tmp8428 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4read_1evaluate_1print))
return
}, 0)

tmp8429 := MakeNative(func(__e *ControlFlow) {
Z5301 := __e.Get(1)
_ = Z5301
tmp8430 := PrimErrorToString(Z5301)

tmp8431 := Call(__e, PrimFunc(symstoutput))


tmp8432 := Call(__e, PrimFunc(sympr), tmp8430, tmp8431)


_ = tmp8432

__e.TailApply(PrimFunc(symnl), MakeNumber(0))
return


}, 1)

tmp8433 := Call(__e, try_1catch, tmp8428, tmp8429)


_ = tmp8433

__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp8434 := Call(__e, ns2_1set, symshen_4loop, tmp8425)


_ = tmp8434

tmp8435 := MakeNative(func(__e *ControlFlow) {
tmp8436 := Call(__e, PrimFunc(symstoutput))


tmp8437 := Call(__e, PrimFunc(sympr), MakeString("\nShen, www.shenlanguage.org, copyright (C) 2010-2024, Mark Tarver\n"), tmp8436)


_ = tmp8437

tmp8438 := PrimValue(sym_dversion_d)

tmp8439 := PrimValue(sym_dlanguage_d)

tmp8440 := PrimValue(sym_dimplementation_d)

tmp8441 := PrimValue(sym_drelease_d)

tmp8442 := Call(__e, PrimFunc(symshen_4app), tmp8441, MakeString("\n"), symshen_4a)


tmp8443 := PrimStringConcat(MakeString(" "), tmp8442)

tmp8444 := Call(__e, PrimFunc(symshen_4app), tmp8440, tmp8443, symshen_4a)


tmp8445 := PrimStringConcat(MakeString(", platform: "), tmp8444)

tmp8446 := Call(__e, PrimFunc(symshen_4app), tmp8439, tmp8445, symshen_4a)


tmp8447 := PrimStringConcat(MakeString(", language: "), tmp8446)

tmp8448 := Call(__e, PrimFunc(symshen_4app), tmp8438, tmp8447, symshen_4a)


tmp8449 := PrimStringConcat(MakeString("version: S"), tmp8448)

tmp8450 := Call(__e, PrimFunc(symstoutput))


tmp8451 := Call(__e, PrimFunc(sympr), tmp8449, tmp8450)


_ = tmp8451

tmp8452 := PrimValue(sym_dport_d)

tmp8453 := PrimValue(sym_dporters_d)

tmp8454 := Call(__e, PrimFunc(symshen_4app), tmp8453, MakeString("\n\n"), symshen_4a)


tmp8455 := PrimStringConcat(MakeString(", ported by "), tmp8454)

tmp8456 := Call(__e, PrimFunc(symshen_4app), tmp8452, tmp8455, symshen_4a)


tmp8457 := PrimStringConcat(MakeString("port "), tmp8456)

tmp8458 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8457, tmp8458)
return


}, 0)

tmp8459 := Call(__e, ns2_1set, symshen_4credits, tmp8435)


_ = tmp8459

tmp8460 := MakeNative(func(__e *ControlFlow) {
tmp8461 := PrimSet(symshen_4_dcall_d, MakeNumber(0))

_ = tmp8461

__e.Return(PrimSet(symshen_4_dinfs_d, MakeNumber(0)))
return


}, 0)

tmp8462 := Call(__e, ns2_1set, symshen_4initialise__environment, tmp8460)


_ = tmp8462

tmp8463 := MakeNative(func(__e *ControlFlow) {
tmp8475 := PrimValue(symshen_4_dtc_d)

if True == tmp8475 {
tmp8464 := PrimValue(symshen_4_dhistory_d)

tmp8465 := Call(__e, PrimFunc(symlength), tmp8464)


tmp8466 := Call(__e, PrimFunc(symshen_4app), tmp8465, MakeString("+) "), symshen_4a)


tmp8467 := PrimStringConcat(MakeString("\n("), tmp8466)

tmp8468 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8467, tmp8468)
return


} else {
tmp8469 := PrimValue(symshen_4_dhistory_d)

tmp8470 := Call(__e, PrimFunc(symlength), tmp8469)


tmp8471 := Call(__e, PrimFunc(symshen_4app), tmp8470, MakeString("-) "), symshen_4a)


tmp8472 := PrimStringConcat(MakeString("\n("), tmp8471)

tmp8473 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8472, tmp8473)
return


}


}, 0)

tmp8476 := Call(__e, ns2_1set, symshen_4prompt, tmp8463)


_ = tmp8476

tmp8477 := MakeNative(func(__e *ControlFlow) {
tmp8478 := MakeNative(func(__e *ControlFlow) {
W5302 := __e.Get(1)
_ = W5302
tmp8479 := MakeNative(func(__e *ControlFlow) {
W5303 := __e.Get(1)
_ = W5303
tmp8480 := MakeNative(func(__e *ControlFlow) {
W5304 := __e.Get(1)
_ = W5304
tmp8481 := PrimValue(symshen_4_dtc_d)

__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5303, W5304, tmp8481)
return


}, 1)

tmp8482 := Call(__e, PrimFunc(symshen_4update_1history))


__e.TailApply(tmp8480, tmp8482)
return


}, 1)

tmp8483 := Call(__e, PrimFunc(symstinput))


tmp8484 := Call(__e, PrimFunc(symlineread), tmp8483)


tmp8485 := Call(__e, PrimFunc(symshen_4package_1user_1input), W5302, tmp8484)


__e.TailApply(tmp8479, tmp8485)
return


}, 1)

tmp8486 := PrimValue(symshen_4_dpackage_d)

__e.TailApply(tmp8478, tmp8486)
return


}, 0)

tmp8487 := Call(__e, ns2_1set, symshen_4read_1evaluate_1print, tmp8477)


_ = tmp8487

tmp8488 := MakeNative(func(__e *ControlFlow) {
V5305 := __e.Get(1)
_ = V5305
V5306 := __e.Get(2)
_ = V5306
tmp8495 := PrimEqual(symnull, V5305)

if True == tmp8495 {
__e.Return(V5306)
return
} else {
tmp8489 := MakeNative(func(__e *ControlFlow) {
W5307 := __e.Get(1)
_ = W5307
tmp8490 := MakeNative(func(__e *ControlFlow) {
W5308 := __e.Get(1)
_ = W5308
tmp8491 := MakeNative(func(__e *ControlFlow) {
Z5309 := __e.Get(1)
_ = Z5309
__e.TailApply(PrimFunc(symshen_4pui_1h), W5307, W5308, Z5309)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8491, V5306)
return


}, 1)

tmp8492 := Call(__e, PrimFunc(symexternal), V5305)


__e.TailApply(tmp8490, tmp8492)
return


}, 1)

tmp8493 := PrimStr(V5305)

__e.TailApply(tmp8489, tmp8493)
return


}


}, 2)

tmp8496 := Call(__e, ns2_1set, symshen_4package_1user_1input, tmp8488)


_ = tmp8496

tmp8497 := MakeNative(func(__e *ControlFlow) {
V5314 := __e.Get(1)
_ = V5314
V5315 := __e.Get(2)
_ = V5315
V5316 := __e.Get(3)
_ = V5316
tmp8538 := PrimIsPair(V5316)

var ifres8525 Obj

if True == tmp8538 {
tmp8536 := PrimHead(V5316)

tmp8537 := PrimEqual(symfn, tmp8536)

var ifres8527 Obj

if True == tmp8537 {
tmp8534 := PrimTail(V5316)

tmp8535 := PrimIsPair(tmp8534)

var ifres8529 Obj

if True == tmp8535 {
tmp8531 := PrimTail(V5316)

tmp8532 := PrimTail(tmp8531)

tmp8533 := PrimEqual(Nil, tmp8532)

var ifres8530 Obj

if True == tmp8533 {
ifres8530 = True


} else {
ifres8530 = False


}

ifres8529 = ifres8530


} else {
ifres8529 = False


}

var ifres8528 Obj

if True == ifres8529 {
ifres8528 = True


} else {
ifres8528 = False


}

ifres8527 = ifres8528


} else {
ifres8527 = False


}

var ifres8526 Obj

if True == ifres8527 {
ifres8526 = True


} else {
ifres8526 = False


}

ifres8525 = ifres8526


} else {
ifres8525 = False


}

if True == ifres8525 {
tmp8503 := PrimTail(V5316)

tmp8504 := PrimHead(tmp8503)

tmp8505 := Call(__e, PrimFunc(symshen_4internal_2), tmp8504, V5314, V5315)


if True == tmp8505 {
tmp8498 := PrimTail(V5316)

tmp8499 := PrimHead(tmp8498)

tmp8500 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5314, tmp8499)


tmp8501 := PrimCons(tmp8500, Nil)

__e.Return(PrimCons(symfn, tmp8501))
return


} else {
__e.Return(V5316)
return
}


} else {
tmp8523 := PrimIsPair(V5316)

if True == tmp8523 {
tmp8520 := PrimHead(V5316)

tmp8521 := Call(__e, PrimFunc(symshen_4internal_2), tmp8520, V5314, V5315)


if True == tmp8521 {
tmp8506 := PrimHead(V5316)

tmp8507 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5314, tmp8506)


tmp8508 := MakeNative(func(__e *ControlFlow) {
Z5317 := __e.Get(1)
_ = Z5317
__e.TailApply(PrimFunc(symshen_4pui_1h), V5314, V5315, Z5317)
return
}, 1)

tmp8509 := PrimTail(V5316)

tmp8510 := Call(__e, PrimFunc(symmap), tmp8508, tmp8509)


__e.Return(PrimCons(tmp8507, tmp8510))
return


} else {
tmp8517 := PrimHead(V5316)

tmp8518 := PrimIsPair(tmp8517)

if True == tmp8518 {
tmp8511 := MakeNative(func(__e *ControlFlow) {
Z5318 := __e.Get(1)
_ = Z5318
__e.TailApply(PrimFunc(symshen_4pui_1h), V5314, V5315, Z5318)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8511, V5316)
return


} else {
tmp8512 := PrimHead(V5316)

tmp8513 := MakeNative(func(__e *ControlFlow) {
Z5319 := __e.Get(1)
_ = Z5319
__e.TailApply(PrimFunc(symshen_4pui_1h), V5314, V5315, Z5319)
return
}, 1)

tmp8514 := PrimTail(V5316)

tmp8515 := Call(__e, PrimFunc(symmap), tmp8513, tmp8514)


__e.Return(PrimCons(tmp8512, tmp8515))
return


}


}


} else {
__e.Return(V5316)
return
}


}


}, 3)

tmp8539 := Call(__e, ns2_1set, symshen_4pui_1h, tmp8497)


_ = tmp8539

tmp8540 := MakeNative(func(__e *ControlFlow) {
tmp8541 := Call(__e, PrimFunc(symit))


tmp8542 := Call(__e, PrimFunc(symshen_4trim_1it), tmp8541)


tmp8543 := PrimValue(symshen_4_dhistory_d)

tmp8544 := PrimCons(tmp8542, tmp8543)

__e.Return(PrimSet(symshen_4_dhistory_d, tmp8544))
return


}, 0)

tmp8545 := Call(__e, ns2_1set, symshen_4update_1history, tmp8540)


_ = tmp8545

tmp8546 := MakeNative(func(__e *ControlFlow) {
V5320 := __e.Get(1)
_ = V5320
tmp8554 := Call(__e, PrimFunc(symshen_4_7string_2), V5320)


var ifres8549 Obj

if True == tmp8554 {
tmp8551 := Call(__e, PrimFunc(symhdstr), V5320)


tmp8552 := PrimStringToNumber(tmp8551)

tmp8553 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8552)


var ifres8550 Obj

if True == tmp8553 {
ifres8550 = True


} else {
ifres8550 = False


}

ifres8549 = ifres8550


} else {
ifres8549 = False


}

if True == ifres8549 {
tmp8547 := PrimTailString(V5320)

__e.TailApply(PrimFunc(symshen_4trim_1it), tmp8547)
return


} else {
__e.Return(V5320)
return
}


}, 1)

tmp8555 := Call(__e, ns2_1set, symshen_4trim_1it, tmp8546)


_ = tmp8555

tmp8556 := MakeNative(func(__e *ControlFlow) {
V5339 := __e.Get(1)
_ = V5339
V5340 := __e.Get(2)
_ = V5340
V5341 := __e.Get(3)
_ = V5341
tmp8686 := PrimIsPair(V5339)

var ifres8655 Obj

if True == tmp8686 {
tmp8684 := PrimTail(V5339)

tmp8685 := PrimEqual(Nil, tmp8684)

var ifres8657 Obj

if True == tmp8685 {
tmp8683 := PrimIsPair(V5340)

var ifres8659 Obj

if True == tmp8683 {
tmp8681 := PrimHead(V5340)

tmp8682 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8681)


var ifres8661 Obj

if True == tmp8682 {
tmp8678 := PrimHead(V5340)

tmp8679 := Call(__e, PrimFunc(symhdstr), tmp8678)


tmp8680 := PrimEqual(MakeString("!"), tmp8679)

var ifres8663 Obj

if True == tmp8680 {
tmp8675 := PrimHead(V5340)

tmp8676 := PrimTailString(tmp8675)

tmp8677 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8676)


var ifres8665 Obj

if True == tmp8677 {
tmp8671 := PrimHead(V5340)

tmp8672 := PrimTailString(tmp8671)

tmp8673 := Call(__e, PrimFunc(symhdstr), tmp8672)


tmp8674 := PrimEqual(MakeString("!"), tmp8673)

var ifres8667 Obj

if True == tmp8674 {
tmp8669 := PrimTail(V5340)

tmp8670 := PrimIsPair(tmp8669)

var ifres8668 Obj

if True == tmp8670 {
ifres8668 = True


} else {
ifres8668 = False


}

ifres8667 = ifres8668


} else {
ifres8667 = False


}

var ifres8666 Obj

if True == ifres8667 {
ifres8666 = True


} else {
ifres8666 = False


}

ifres8665 = ifres8666


} else {
ifres8665 = False


}

var ifres8664 Obj

if True == ifres8665 {
ifres8664 = True


} else {
ifres8664 = False


}

ifres8663 = ifres8664


} else {
ifres8663 = False


}

var ifres8662 Obj

if True == ifres8663 {
ifres8662 = True


} else {
ifres8662 = False


}

ifres8661 = ifres8662


} else {
ifres8661 = False


}

var ifres8660 Obj

if True == ifres8661 {
ifres8660 = True


} else {
ifres8660 = False


}

ifres8659 = ifres8660


} else {
ifres8659 = False


}

var ifres8658 Obj

if True == ifres8659 {
ifres8658 = True


} else {
ifres8658 = False


}

ifres8657 = ifres8658


} else {
ifres8657 = False


}

var ifres8656 Obj

if True == ifres8657 {
ifres8656 = True


} else {
ifres8656 = False


}

ifres8655 = ifres8656


} else {
ifres8655 = False


}

if True == ifres8655 {
tmp8557 := MakeNative(func(__e *ControlFlow) {
W5342 := __e.Get(1)
_ = W5342
tmp8558 := MakeNative(func(__e *ControlFlow) {
W5343 := __e.Get(1)
_ = W5343
tmp8559 := MakeNative(func(__e *ControlFlow) {
W5344 := __e.Get(1)
_ = W5344
__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5342, W5343, V5341)
return
}, 1)

tmp8560 := PrimTail(V5340)

tmp8561 := PrimHead(tmp8560)

tmp8562 := Call(__e, PrimFunc(symshen_4app), tmp8561, MakeString("\n"), symshen_4a)


tmp8563 := Call(__e, PrimFunc(symstoutput))


tmp8564 := Call(__e, PrimFunc(sympr), tmp8562, tmp8563)


__e.TailApply(tmp8559, tmp8564)
return


}, 1)

tmp8565 := PrimTail(V5340)

tmp8566 := PrimHead(tmp8565)

tmp8567 := PrimTail(V5340)

tmp8568 := PrimCons(tmp8566, tmp8567)

tmp8569 := PrimSet(symshen_4_dhistory_d, tmp8568)

__e.TailApply(tmp8558, tmp8569)
return


}, 1)

tmp8570 := PrimTail(V5340)

tmp8571 := PrimHead(tmp8570)

tmp8572 := Call(__e, PrimFunc(symread_1from_1string), tmp8571)


__e.TailApply(tmp8557, tmp8572)
return


} else {
tmp8653 := PrimIsPair(V5339)

var ifres8637 Obj

if True == tmp8653 {
tmp8651 := PrimTail(V5339)

tmp8652 := PrimEqual(Nil, tmp8651)

var ifres8639 Obj

if True == tmp8652 {
tmp8650 := PrimIsPair(V5340)

var ifres8641 Obj

if True == tmp8650 {
tmp8648 := PrimHead(V5340)

tmp8649 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8648)


var ifres8643 Obj

if True == tmp8649 {
tmp8645 := PrimHead(V5340)

tmp8646 := Call(__e, PrimFunc(symhdstr), tmp8645)


tmp8647 := PrimEqual(MakeString("!"), tmp8646)

var ifres8644 Obj

if True == tmp8647 {
ifres8644 = True


} else {
ifres8644 = False


}

ifres8643 = ifres8644


} else {
ifres8643 = False


}

var ifres8642 Obj

if True == ifres8643 {
ifres8642 = True


} else {
ifres8642 = False


}

ifres8641 = ifres8642


} else {
ifres8641 = False


}

var ifres8640 Obj

if True == ifres8641 {
ifres8640 = True


} else {
ifres8640 = False


}

ifres8639 = ifres8640


} else {
ifres8639 = False


}

var ifres8638 Obj

if True == ifres8639 {
ifres8638 = True


} else {
ifres8638 = False


}

ifres8637 = ifres8638


} else {
ifres8637 = False


}

if True == ifres8637 {
tmp8573 := MakeNative(func(__e *ControlFlow) {
W5345 := __e.Get(1)
_ = W5345
tmp8574 := MakeNative(func(__e *ControlFlow) {
W5346 := __e.Get(1)
_ = W5346
tmp8575 := MakeNative(func(__e *ControlFlow) {
W5347 := __e.Get(1)
_ = W5347
tmp8576 := MakeNative(func(__e *ControlFlow) {
W5348 := __e.Get(1)
_ = W5348
tmp8577 := MakeNative(func(__e *ControlFlow) {
W5349 := __e.Get(1)
_ = W5349
__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5348, W5349, V5341)
return
}, 1)

tmp8578 := PrimTail(V5340)

tmp8579 := PrimCons(W5346, tmp8578)

tmp8580 := PrimSet(symshen_4_dhistory_d, tmp8579)

__e.TailApply(tmp8577, tmp8580)
return


}, 1)

tmp8581 := Call(__e, PrimFunc(symread_1from_1string), W5346)


__e.TailApply(tmp8576, tmp8581)
return


}, 1)

tmp8582 := Call(__e, PrimFunc(symshen_4app), W5346, MakeString("\n"), symshen_4a)


tmp8583 := Call(__e, PrimFunc(symstoutput))


tmp8584 := Call(__e, PrimFunc(sympr), tmp8582, tmp8583)


__e.TailApply(tmp8575, tmp8584)
return


}, 1)

tmp8585 := PrimHead(V5340)

tmp8586 := PrimTailString(tmp8585)

tmp8587 := PrimTail(V5340)

tmp8588 := Call(__e, PrimFunc(symshen_4use_1history), W5345, tmp8586, tmp8587)


__e.TailApply(tmp8574, tmp8588)
return


}, 1)

tmp8594 := PrimHead(V5340)

tmp8595 := PrimTailString(tmp8594)

tmp8596 := PrimEqual(tmp8595, MakeString(""))

var ifres8589 Obj

if True == tmp8596 {
ifres8589 = Nil


} else {
tmp8590 := PrimHead(V5340)

tmp8591 := PrimTailString(tmp8590)

tmp8592 := Call(__e, PrimFunc(symread_1from_1string), tmp8591)


tmp8593 := PrimHead(tmp8592)

ifres8589 = tmp8593


}

__e.TailApply(tmp8573, ifres8589)
return


} else {
tmp8635 := PrimIsPair(V5339)

var ifres8619 Obj

if True == tmp8635 {
tmp8633 := PrimTail(V5339)

tmp8634 := PrimEqual(Nil, tmp8633)

var ifres8621 Obj

if True == tmp8634 {
tmp8632 := PrimIsPair(V5340)

var ifres8623 Obj

if True == tmp8632 {
tmp8630 := PrimHead(V5340)

tmp8631 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8630)


var ifres8625 Obj

if True == tmp8631 {
tmp8627 := PrimHead(V5340)

tmp8628 := Call(__e, PrimFunc(symhdstr), tmp8627)


tmp8629 := PrimEqual(MakeString("%"), tmp8628)

var ifres8626 Obj

if True == tmp8629 {
ifres8626 = True


} else {
ifres8626 = False


}

ifres8625 = ifres8626


} else {
ifres8625 = False


}

var ifres8624 Obj

if True == ifres8625 {
ifres8624 = True


} else {
ifres8624 = False


}

ifres8623 = ifres8624


} else {
ifres8623 = False


}

var ifres8622 Obj

if True == ifres8623 {
ifres8622 = True


} else {
ifres8622 = False


}

ifres8621 = ifres8622


} else {
ifres8621 = False


}

var ifres8620 Obj

if True == ifres8621 {
ifres8620 = True


} else {
ifres8620 = False


}

ifres8619 = ifres8620


} else {
ifres8619 = False


}

if True == ifres8619 {
tmp8597 := MakeNative(func(__e *ControlFlow) {
W5350 := __e.Get(1)
_ = W5350
tmp8598 := MakeNative(func(__e *ControlFlow) {
W5351 := __e.Get(1)
_ = W5351
tmp8599 := MakeNative(func(__e *ControlFlow) {
W5352 := __e.Get(1)
_ = W5352
__e.TailApply(PrimFunc(symabort))
return
}, 1)

tmp8600 := PrimTail(V5340)

tmp8601 := PrimSet(symshen_4_dhistory_d, tmp8600)

__e.TailApply(tmp8599, tmp8601)
return


}, 1)

tmp8602 := PrimHead(V5340)

tmp8603 := PrimTailString(tmp8602)

tmp8604 := PrimTail(V5340)

tmp8605 := Call(__e, PrimFunc(symshen_4peek_1history), W5350, tmp8603, tmp8604)


__e.TailApply(tmp8598, tmp8605)
return


}, 1)

tmp8611 := PrimHead(V5340)

tmp8612 := PrimTailString(tmp8611)

tmp8613 := PrimEqual(tmp8612, MakeString(""))

var ifres8606 Obj

if True == tmp8613 {
ifres8606 = Nil


} else {
tmp8607 := PrimHead(V5340)

tmp8608 := PrimTailString(tmp8607)

tmp8609 := Call(__e, PrimFunc(symread_1from_1string), tmp8608)


tmp8610 := PrimHead(tmp8609)

ifres8606 = tmp8610


}

__e.TailApply(tmp8597, ifres8606)
return


} else {
tmp8617 := PrimEqual(True, V5341)

if True == tmp8617 {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V5339)
return
} else {
tmp8615 := PrimEqual(False, V5341)

if True == tmp8615 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V5339)
return
} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.evaluate-lineread")))
return
}


}


}


}


}


}, 3)

tmp8687 := Call(__e, ns2_1set, symshen_4evaluate_1lineread, tmp8556)


_ = tmp8687

tmp8688 := MakeNative(func(__e *ControlFlow) {
V5353 := __e.Get(1)
_ = V5353
V5354 := __e.Get(2)
_ = V5354
V5355 := __e.Get(3)
_ = V5355
tmp8694 := PrimIsInteger(V5353)

if True == tmp8694 {
tmp8689 := PrimNumberAdd(MakeNumber(1), V5353)

tmp8690 := Call(__e, PrimFunc(symreverse), V5355)


__e.TailApply(PrimFunc(symnth), tmp8689, tmp8690)
return


} else {
tmp8692 := PrimIsSymbol(V5353)

if True == tmp8692 {
__e.TailApply(PrimFunc(symshen_4string_1match), V5354, V5355)
return
} else {
__e.Return(PrimSimpleError(MakeString("! expects a number or a symbol\n")))
return
}


}


}, 3)

tmp8695 := Call(__e, ns2_1set, symshen_4use_1history, tmp8688)


_ = tmp8695

tmp8696 := MakeNative(func(__e *ControlFlow) {
V5356 := __e.Get(1)
_ = V5356
V5357 := __e.Get(2)
_ = V5357
V5358 := __e.Get(3)
_ = V5358
tmp8710 := PrimIsInteger(V5356)

if True == tmp8710 {
tmp8697 := PrimNumberAdd(MakeNumber(1), V5356)

tmp8698 := Call(__e, PrimFunc(symreverse), V5358)


tmp8699 := Call(__e, PrimFunc(symnth), tmp8697, tmp8698)


tmp8700 := Call(__e, PrimFunc(symshen_4app), tmp8699, MakeString(""), symshen_4a)


tmp8701 := PrimStringConcat(MakeString("\n"), tmp8700)

tmp8702 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8701, tmp8702)
return


} else {
tmp8708 := PrimEqual(V5357, MakeString(""))

var ifres8705 Obj

if True == tmp8708 {
ifres8705 = True


} else {
tmp8707 := PrimIsSymbol(V5356)

var ifres8706 Obj

if True == tmp8707 {
ifres8706 = True


} else {
ifres8706 = False


}

ifres8705 = ifres8706


}

if True == ifres8705 {
tmp8703 := Call(__e, PrimFunc(symreverse), V5358)


__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), MakeNumber(0), V5357, tmp8703)
return


} else {
__e.Return(PrimSimpleError(MakeString("% expects a number or a symbol\n")))
return
}


}


}, 3)

tmp8711 := Call(__e, ns2_1set, symshen_4peek_1history, tmp8696)


_ = tmp8711

tmp8712 := MakeNative(func(__e *ControlFlow) {
V5368 := __e.Get(1)
_ = V5368
V5369 := __e.Get(2)
_ = V5369
tmp8723 := PrimEqual(Nil, V5369)

if True == tmp8723 {
__e.Return(PrimSimpleError(MakeString("\ninput not found")))
return
} else {
tmp8721 := PrimIsPair(V5369)

var ifres8717 Obj

if True == tmp8721 {
tmp8719 := PrimHead(V5369)

tmp8720 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5368, tmp8719)


var ifres8718 Obj

if True == tmp8720 {
ifres8718 = True


} else {
ifres8718 = False


}

ifres8717 = ifres8718


} else {
ifres8717 = False


}

if True == ifres8717 {
__e.Return(PrimHead(V5369))
return
} else {
tmp8715 := PrimIsPair(V5369)

if True == tmp8715 {
tmp8713 := PrimTail(V5369)

__e.TailApply(PrimFunc(symshen_4string_1match), V5368, tmp8713)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.string-match")))
return
}


}


}


}, 2)

tmp8724 := Call(__e, ns2_1set, symshen_4string_1match, tmp8712)


_ = tmp8724

tmp8725 := MakeNative(func(__e *ControlFlow) {
V5377 := __e.Get(1)
_ = V5377
V5378 := __e.Get(2)
_ = V5378
tmp8762 := PrimEqual(MakeString(""), V5377)

if True == tmp8762 {
__e.Return(True)
return
} else {
tmp8760 := Call(__e, PrimFunc(symshen_4_7string_2), V5377)


var ifres8755 Obj

if True == tmp8760 {
tmp8757 := Call(__e, PrimFunc(symhdstr), V5377)


tmp8758 := PrimStringToNumber(tmp8757)

tmp8759 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8758)


var ifres8756 Obj

if True == tmp8759 {
ifres8756 = True


} else {
ifres8756 = False


}

ifres8755 = ifres8756


} else {
ifres8755 = False


}

if True == ifres8755 {
tmp8726 := PrimTailString(V5377)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp8726, V5378)
return


} else {
tmp8753 := Call(__e, PrimFunc(symshen_4_7string_2), V5378)


var ifres8748 Obj

if True == tmp8753 {
tmp8750 := Call(__e, PrimFunc(symhdstr), V5378)


tmp8751 := PrimStringToNumber(tmp8750)

tmp8752 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8751)


var ifres8749 Obj

if True == tmp8752 {
ifres8749 = True


} else {
ifres8749 = False


}

ifres8748 = ifres8749


} else {
ifres8748 = False


}

if True == ifres8748 {
tmp8727 := PrimTailString(V5378)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5377, tmp8727)
return


} else {
tmp8746 := Call(__e, PrimFunc(symshen_4_7string_2), V5378)


var ifres8742 Obj

if True == tmp8746 {
tmp8744 := Call(__e, PrimFunc(symhdstr), V5378)


tmp8745 := PrimEqual(MakeString("("), tmp8744)

var ifres8743 Obj

if True == tmp8745 {
ifres8743 = True


} else {
ifres8743 = False


}

ifres8742 = ifres8743


} else {
ifres8742 = False


}

if True == ifres8742 {
tmp8728 := PrimTailString(V5378)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5377, tmp8728)
return


} else {
tmp8740 := Call(__e, PrimFunc(symshen_4_7string_2), V5377)


var ifres8732 Obj

if True == tmp8740 {
tmp8739 := Call(__e, PrimFunc(symshen_4_7string_2), V5378)


var ifres8734 Obj

if True == tmp8739 {
tmp8736 := Call(__e, PrimFunc(symhdstr), V5377)


tmp8737 := Call(__e, PrimFunc(symhdstr), V5378)


tmp8738 := PrimEqual(tmp8736, tmp8737)

var ifres8735 Obj

if True == tmp8738 {
ifres8735 = True


} else {
ifres8735 = False


}

ifres8734 = ifres8735


} else {
ifres8734 = False


}

var ifres8733 Obj

if True == ifres8734 {
ifres8733 = True


} else {
ifres8733 = False


}

ifres8732 = ifres8733


} else {
ifres8732 = False


}

if True == ifres8732 {
tmp8729 := PrimTailString(V5377)

tmp8730 := PrimTailString(V5378)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp8729, tmp8730)
return


} else {
__e.Return(False)
return
}


}


}


}


}


}, 2)

tmp8763 := Call(__e, ns2_1set, symshen_4string_1prefix_2, tmp8725)


_ = tmp8763

tmp8764 := MakeNative(func(__e *ControlFlow) {
V5389 := __e.Get(1)
_ = V5389
V5390 := __e.Get(2)
_ = V5390
V5391 := __e.Get(3)
_ = V5391
tmp8779 := PrimEqual(Nil, V5391)

if True == tmp8779 {
__e.Return(symshen_4skip)
return
} else {
tmp8777 := PrimIsPair(V5391)

if True == tmp8777 {
tmp8772 := PrimHead(V5391)

tmp8773 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5390, tmp8772)


var ifres8765 Obj

if True == tmp8773 {
tmp8766 := PrimHead(V5391)

tmp8767 := Call(__e, PrimFunc(symshen_4app), tmp8766, MakeString("\n"), symshen_4a)


tmp8768 := PrimStringConcat(MakeString(". "), tmp8767)

tmp8769 := Call(__e, PrimFunc(symshen_4app), V5389, tmp8768, symshen_4a)


tmp8770 := Call(__e, PrimFunc(symstoutput))


tmp8771 := Call(__e, PrimFunc(sympr), tmp8769, tmp8770)


ifres8765 = tmp8771


} else {
ifres8765 = symshen_4skip


}

_ = ifres8765

tmp8774 := PrimNumberAdd(V5389, MakeNumber(1))

tmp8775 := PrimTail(V5391)

__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), tmp8774, V5390, tmp8775)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.recursive-string-match")))
return
}


}


}, 3)

__e.TailApply(ns2_1set, symshen_4recursive_1string_1match, tmp8764)
return




}, 0)

