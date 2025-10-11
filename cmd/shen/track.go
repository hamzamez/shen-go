package main

import . "github.com/tiancaiamao/shen-go/kl"

var TrackMain = MakeNative(func(__e *ControlFlow) {
tmp13480 := MakeNative(func(__e *ControlFlow) {
V5423 := __e.Get(1)
_ = V5423
tmp13481 := Call(__e, PrimFunc(symshen_4app), V5423, MakeString(";\n"), symshen_4a)


tmp13482 := PrimStringConcat(MakeString("partial function "), tmp13481)

tmp13483 := Call(__e, PrimFunc(symstoutput))


tmp13484 := Call(__e, PrimFunc(sympr), tmp13482, tmp13483)


_ = tmp13484

tmp13493 := Call(__e, PrimFunc(symshen_4tracked_2), V5423)


tmp13494 := PrimNot(tmp13493)

var ifres13488 Obj

if True == tmp13494 {
tmp13490 := Call(__e, PrimFunc(symshen_4app), V5423, MakeString("? "), symshen_4a)


tmp13491 := PrimStringConcat(MakeString("track "), tmp13490)

tmp13492 := Call(__e, PrimFunc(symy_1or_1n_2), tmp13491)


var ifres13489 Obj

if True == tmp13492 {
ifres13489 = True


} else {
ifres13489 = False


}

ifres13488 = ifres13489


} else {
ifres13488 = False


}

var ifres13485 Obj

if True == ifres13488 {
tmp13486 := Call(__e, PrimFunc(symps), V5423)


tmp13487 := Call(__e, PrimFunc(symshen_4track_1function), tmp13486)


ifres13485 = tmp13487


} else {
ifres13485 = symshen_4ok


}

_ = ifres13485

__e.Return(PrimSimpleError(MakeString("aborted")))
return


}, 1)

tmp13495 := Call(__e, ns2_1set, symshen_4f_1error, tmp13480)


_ = tmp13495

tmp13496 := MakeNative(func(__e *ControlFlow) {
V5424 := __e.Get(1)
_ = V5424
tmp13497 := PrimValue(symshen_4_dtracking_d)

__e.TailApply(PrimFunc(symelement_2), V5424, tmp13497)
return


}, 1)

tmp13498 := Call(__e, ns2_1set, symshen_4tracked_2, tmp13496)


_ = tmp13498

tmp13499 := MakeNative(func(__e *ControlFlow) {
V5425 := __e.Get(1)
_ = V5425
tmp13500 := MakeNative(func(__e *ControlFlow) {
W5426 := __e.Get(1)
_ = W5426
__e.TailApply(PrimFunc(symshen_4track_1function), W5426)
return
}, 1)

tmp13501 := Call(__e, PrimFunc(symps), V5425)


__e.TailApply(tmp13500, tmp13501)
return


}, 1)

tmp13502 := Call(__e, ns2_1set, symtrack, tmp13499)


_ = tmp13502

tmp13503 := MakeNative(func(__e *ControlFlow) {
V5429 := __e.Get(1)
_ = V5429
tmp13560 := PrimIsPair(V5429)

var ifres13534 Obj

if True == tmp13560 {
tmp13558 := PrimHead(V5429)

tmp13559 := PrimEqual(symdefun, tmp13558)

var ifres13536 Obj

if True == tmp13559 {
tmp13556 := PrimTail(V5429)

tmp13557 := PrimIsPair(tmp13556)

var ifres13538 Obj

if True == tmp13557 {
tmp13553 := PrimTail(V5429)

tmp13554 := PrimTail(tmp13553)

tmp13555 := PrimIsPair(tmp13554)

var ifres13540 Obj

if True == tmp13555 {
tmp13549 := PrimTail(V5429)

tmp13550 := PrimTail(tmp13549)

tmp13551 := PrimTail(tmp13550)

tmp13552 := PrimIsPair(tmp13551)

var ifres13542 Obj

if True == tmp13552 {
tmp13544 := PrimTail(V5429)

tmp13545 := PrimTail(tmp13544)

tmp13546 := PrimTail(tmp13545)

tmp13547 := PrimTail(tmp13546)

tmp13548 := PrimEqual(Nil, tmp13547)

var ifres13543 Obj

if True == tmp13548 {
ifres13543 = True


} else {
ifres13543 = False


}

ifres13542 = ifres13543


} else {
ifres13542 = False


}

var ifres13541 Obj

if True == ifres13542 {
ifres13541 = True


} else {
ifres13541 = False


}

ifres13540 = ifres13541


} else {
ifres13540 = False


}

var ifres13539 Obj

if True == ifres13540 {
ifres13539 = True


} else {
ifres13539 = False


}

ifres13538 = ifres13539


} else {
ifres13538 = False


}

var ifres13537 Obj

if True == ifres13538 {
ifres13537 = True


} else {
ifres13537 = False


}

ifres13536 = ifres13537


} else {
ifres13536 = False


}

var ifres13535 Obj

if True == ifres13536 {
ifres13535 = True


} else {
ifres13535 = False


}

ifres13534 = ifres13535


} else {
ifres13534 = False


}

if True == ifres13534 {
tmp13504 := MakeNative(func(__e *ControlFlow) {
W5430 := __e.Get(1)
_ = W5430
tmp13505 := MakeNative(func(__e *ControlFlow) {
W5431 := __e.Get(1)
_ = W5431
tmp13506 := MakeNative(func(__e *ControlFlow) {
W5432 := __e.Get(1)
_ = W5432
tmp13507 := PrimTail(V5429)

__e.Return(PrimHead(tmp13507))
return


}, 1)

tmp13508 := PrimTail(V5429)

tmp13509 := PrimHead(tmp13508)

tmp13510 := PrimValue(symshen_4_dtracking_d)

tmp13511 := Call(__e, PrimFunc(symadjoin), tmp13509, tmp13510)


tmp13512 := PrimSet(symshen_4_dtracking_d, tmp13511)

__e.TailApply(tmp13506, tmp13512)
return


}, 1)

tmp13513 := Call(__e, PrimFunc(symeval_1kl), W5430)


__e.TailApply(tmp13505, tmp13513)
return


}, 1)

tmp13514 := PrimTail(V5429)

tmp13515 := PrimHead(tmp13514)

tmp13516 := PrimTail(V5429)

tmp13517 := PrimTail(tmp13516)

tmp13518 := PrimHead(tmp13517)

tmp13519 := PrimTail(V5429)

tmp13520 := PrimHead(tmp13519)

tmp13521 := PrimTail(V5429)

tmp13522 := PrimTail(tmp13521)

tmp13523 := PrimHead(tmp13522)

tmp13524 := PrimTail(V5429)

tmp13525 := PrimTail(tmp13524)

tmp13526 := PrimTail(tmp13525)

tmp13527 := PrimHead(tmp13526)

tmp13528 := Call(__e, PrimFunc(symshen_4insert_1tracking_1code), tmp13520, tmp13523, tmp13527)


tmp13529 := PrimCons(tmp13528, Nil)

tmp13530 := PrimCons(tmp13518, tmp13529)

tmp13531 := PrimCons(tmp13515, tmp13530)

tmp13532 := PrimCons(symdefun, tmp13531)

__e.TailApply(tmp13504, tmp13532)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.track-function")))
return
}


}, 1)

tmp13561 := Call(__e, ns2_1set, symshen_4track_1function, tmp13503)


_ = tmp13561

tmp13562 := MakeNative(func(__e *ControlFlow) {
V5433 := __e.Get(1)
_ = V5433
V5434 := __e.Get(2)
_ = V5434
V5435 := __e.Get(3)
_ = V5435
tmp13563 := PrimCons(symshen_4_dcall_d, Nil)

tmp13564 := PrimCons(symvalue, tmp13563)

tmp13565 := PrimCons(MakeNumber(1), Nil)

tmp13566 := PrimCons(tmp13564, tmp13565)

tmp13567 := PrimCons(sym_7, tmp13566)

tmp13568 := PrimCons(tmp13567, Nil)

tmp13569 := PrimCons(symshen_4_dcall_d, tmp13568)

tmp13570 := PrimCons(symset, tmp13569)

tmp13571 := PrimCons(symshen_4_dcall_d, Nil)

tmp13572 := PrimCons(symvalue, tmp13571)

tmp13573 := Call(__e, PrimFunc(symshen_4prolog_1track), V5435, V5434)


tmp13574 := Call(__e, PrimFunc(symshen_4cons_1form), tmp13573)


tmp13575 := PrimCons(tmp13574, Nil)

tmp13576 := PrimCons(V5433, tmp13575)

tmp13577 := PrimCons(tmp13572, tmp13576)

tmp13578 := PrimCons(symshen_4input_1track, tmp13577)

tmp13579 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

tmp13580 := PrimCons(symshen_4_dcall_d, Nil)

tmp13581 := PrimCons(symvalue, tmp13580)

tmp13582 := PrimCons(symResult, Nil)

tmp13583 := PrimCons(V5433, tmp13582)

tmp13584 := PrimCons(tmp13581, tmp13583)

tmp13585 := PrimCons(symshen_4output_1track, tmp13584)

tmp13586 := PrimCons(symshen_4_dcall_d, Nil)

tmp13587 := PrimCons(symvalue, tmp13586)

tmp13588 := PrimCons(MakeNumber(1), Nil)

tmp13589 := PrimCons(tmp13587, tmp13588)

tmp13590 := PrimCons(sym_1, tmp13589)

tmp13591 := PrimCons(tmp13590, Nil)

tmp13592 := PrimCons(symshen_4_dcall_d, tmp13591)

tmp13593 := PrimCons(symset, tmp13592)

tmp13594 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

tmp13595 := PrimCons(symResult, Nil)

tmp13596 := PrimCons(tmp13594, tmp13595)

tmp13597 := PrimCons(symdo, tmp13596)

tmp13598 := PrimCons(tmp13597, Nil)

tmp13599 := PrimCons(tmp13593, tmp13598)

tmp13600 := PrimCons(symdo, tmp13599)

tmp13601 := PrimCons(tmp13600, Nil)

tmp13602 := PrimCons(tmp13585, tmp13601)

tmp13603 := PrimCons(symdo, tmp13602)

tmp13604 := PrimCons(tmp13603, Nil)

tmp13605 := PrimCons(V5435, tmp13604)

tmp13606 := PrimCons(symResult, tmp13605)

tmp13607 := PrimCons(symlet, tmp13606)

tmp13608 := PrimCons(tmp13607, Nil)

tmp13609 := PrimCons(tmp13579, tmp13608)

tmp13610 := PrimCons(symdo, tmp13609)

tmp13611 := PrimCons(tmp13610, Nil)

tmp13612 := PrimCons(tmp13578, tmp13611)

tmp13613 := PrimCons(symdo, tmp13612)

tmp13614 := PrimCons(tmp13613, Nil)

tmp13615 := PrimCons(tmp13570, tmp13614)

__e.Return(PrimCons(symdo, tmp13615))
return


}, 3)

tmp13616 := Call(__e, ns2_1set, symshen_4insert_1tracking_1code, tmp13562)


_ = tmp13616

tmp13617 := MakeNative(func(__e *ControlFlow) {
V5436 := __e.Get(1)
_ = V5436
V5437 := __e.Get(2)
_ = V5437
tmp13620 := Call(__e, PrimFunc(symoccurrences), symshen_4incinfs, V5436)


tmp13621 := PrimEqual(tmp13620, MakeNumber(0))

if True == tmp13621 {
__e.Return(V5437)
return
} else {
tmp13618 := Call(__e, PrimFunc(symshen_4vector_1parameter), V5437)


__e.TailApply(PrimFunc(symshen_4vector_1dereference), V5437, tmp13618)
return


}


}, 2)

tmp13622 := Call(__e, ns2_1set, symshen_4prolog_1track, tmp13617)


_ = tmp13622

tmp13623 := MakeNative(func(__e *ControlFlow) {
V5440 := __e.Get(1)
_ = V5440
tmp13652 := PrimEqual(Nil, V5440)

if True == tmp13652 {
__e.Return(Nil)
return
} else {
tmp13650 := PrimIsPair(V5440)

var ifres13628 Obj

if True == tmp13650 {
tmp13648 := PrimTail(V5440)

tmp13649 := PrimIsPair(tmp13648)

var ifres13630 Obj

if True == tmp13649 {
tmp13645 := PrimTail(V5440)

tmp13646 := PrimTail(tmp13645)

tmp13647 := PrimIsPair(tmp13646)

var ifres13632 Obj

if True == tmp13647 {
tmp13641 := PrimTail(V5440)

tmp13642 := PrimTail(tmp13641)

tmp13643 := PrimTail(tmp13642)

tmp13644 := PrimIsPair(tmp13643)

var ifres13634 Obj

if True == tmp13644 {
tmp13636 := PrimTail(V5440)

tmp13637 := PrimTail(tmp13636)

tmp13638 := PrimTail(tmp13637)

tmp13639 := PrimTail(tmp13638)

tmp13640 := PrimEqual(Nil, tmp13639)

var ifres13635 Obj

if True == tmp13640 {
ifres13635 = True


} else {
ifres13635 = False


}

ifres13634 = ifres13635


} else {
ifres13634 = False


}

var ifres13633 Obj

if True == ifres13634 {
ifres13633 = True


} else {
ifres13633 = False


}

ifres13632 = ifres13633


} else {
ifres13632 = False


}

var ifres13631 Obj

if True == ifres13632 {
ifres13631 = True


} else {
ifres13631 = False


}

ifres13630 = ifres13631


} else {
ifres13630 = False


}

var ifres13629 Obj

if True == ifres13630 {
ifres13629 = True


} else {
ifres13629 = False


}

ifres13628 = ifres13629


} else {
ifres13628 = False


}

if True == ifres13628 {
__e.Return(PrimHead(V5440))
return
} else {
tmp13626 := PrimIsPair(V5440)

if True == tmp13626 {
tmp13624 := PrimTail(V5440)

__e.TailApply(PrimFunc(symshen_4vector_1parameter), tmp13624)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.vector-parameter")))
return
}


}


}


}, 1)

tmp13653 := Call(__e, ns2_1set, symshen_4vector_1parameter, tmp13623)


_ = tmp13653

tmp13654 := MakeNative(func(__e *ControlFlow) {
V5443 := __e.Get(1)
_ = V5443
V5444 := __e.Get(2)
_ = V5444
tmp13688 := PrimEqual(Nil, V5444)

if True == tmp13688 {
__e.Return(V5443)
return
} else {
tmp13686 := PrimIsPair(V5443)

var ifres13664 Obj

if True == tmp13686 {
tmp13684 := PrimTail(V5443)

tmp13685 := PrimIsPair(tmp13684)

var ifres13666 Obj

if True == tmp13685 {
tmp13681 := PrimTail(V5443)

tmp13682 := PrimTail(tmp13681)

tmp13683 := PrimIsPair(tmp13682)

var ifres13668 Obj

if True == tmp13683 {
tmp13677 := PrimTail(V5443)

tmp13678 := PrimTail(tmp13677)

tmp13679 := PrimTail(tmp13678)

tmp13680 := PrimIsPair(tmp13679)

var ifres13670 Obj

if True == tmp13680 {
tmp13672 := PrimTail(V5443)

tmp13673 := PrimTail(tmp13672)

tmp13674 := PrimTail(tmp13673)

tmp13675 := PrimTail(tmp13674)

tmp13676 := PrimEqual(Nil, tmp13675)

var ifres13671 Obj

if True == tmp13676 {
ifres13671 = True


} else {
ifres13671 = False


}

ifres13670 = ifres13671


} else {
ifres13670 = False


}

var ifres13669 Obj

if True == ifres13670 {
ifres13669 = True


} else {
ifres13669 = False


}

ifres13668 = ifres13669


} else {
ifres13668 = False


}

var ifres13667 Obj

if True == ifres13668 {
ifres13667 = True


} else {
ifres13667 = False


}

ifres13666 = ifres13667


} else {
ifres13666 = False


}

var ifres13665 Obj

if True == ifres13666 {
ifres13665 = True


} else {
ifres13665 = False


}

ifres13664 = ifres13665


} else {
ifres13664 = False


}

if True == ifres13664 {
__e.Return(V5443)
return
} else {
tmp13662 := PrimIsPair(V5443)

if True == tmp13662 {
tmp13655 := PrimHead(V5443)

tmp13656 := PrimCons(V5444, Nil)

tmp13657 := PrimCons(tmp13655, tmp13656)

tmp13658 := PrimCons(symshen_4deref, tmp13657)

tmp13659 := PrimTail(V5443)

tmp13660 := Call(__e, PrimFunc(symshen_4vector_1dereference), tmp13659, V5444)


__e.Return(PrimCons(tmp13658, tmp13660))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.vector-dereference")))
return
}


}


}


}, 2)

tmp13689 := Call(__e, ns2_1set, symshen_4vector_1dereference, tmp13654)


_ = tmp13689

tmp13690 := MakeNative(func(__e *ControlFlow) {
V5447 := __e.Get(1)
_ = V5447
tmp13694 := PrimEqual(sym_7, V5447)

if True == tmp13694 {
__e.Return(PrimSet(symshen_4_dstep_d, True))
return
} else {
tmp13692 := PrimEqual(sym_1, V5447)

if True == tmp13692 {
__e.Return(PrimSet(symshen_4_dstep_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("step expects a + or a -.\n")))
return
}


}


}, 1)

tmp13695 := Call(__e, ns2_1set, symstep, tmp13690)


_ = tmp13695

tmp13696 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dstep_d))
return
}, 0)

tmp13697 := Call(__e, ns2_1set, symstep_2, tmp13696)


_ = tmp13697

tmp13698 := MakeNative(func(__e *ControlFlow) {
V5450 := __e.Get(1)
_ = V5450
tmp13702 := PrimEqual(sym_7, V5450)

if True == tmp13702 {
__e.Return(PrimSet(symshen_4_dspy_d, True))
return
} else {
tmp13700 := PrimEqual(sym_1, V5450)

if True == tmp13700 {
__e.Return(PrimSet(symshen_4_dspy_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("spy expects a + or a -.\n")))
return
}


}


}, 1)

tmp13703 := Call(__e, ns2_1set, symspy, tmp13698)


_ = tmp13703

tmp13704 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dspy_d))
return
}, 0)

tmp13705 := Call(__e, ns2_1set, symspy_2, tmp13704)


_ = tmp13705

tmp13706 := MakeNative(func(__e *ControlFlow) {
tmp13710 := PrimValue(symshen_4_dstep_d)

if True == tmp13710 {
tmp13707 := PrimValue(sym_dstinput_d)

tmp13708 := PrimReadByte(tmp13707)

__e.TailApply(PrimFunc(symshen_4check_1byte), tmp13708)
return


} else {
__e.TailApply(PrimFunc(symnl), MakeNumber(1))
return
}


}, 0)

tmp13711 := Call(__e, ns2_1set, symshen_4terpri_1or_1read_1char, tmp13706)


_ = tmp13711

tmp13712 := MakeNative(func(__e *ControlFlow) {
V5453 := __e.Get(1)
_ = V5453
tmp13714 := PrimEqual(MakeNumber(94), V5453)

if True == tmp13714 {
__e.Return(PrimSimpleError(MakeString("aborted")))
return
} else {
__e.Return(True)
return
}


}, 1)

tmp13715 := Call(__e, ns2_1set, symshen_4check_1byte, tmp13712)


_ = tmp13715

tmp13716 := MakeNative(func(__e *ControlFlow) {
V5454 := __e.Get(1)
_ = V5454
V5455 := __e.Get(2)
_ = V5455
V5456 := __e.Get(3)
_ = V5456
tmp13717 := Call(__e, PrimFunc(symshen_4spaces), V5454)


tmp13718 := Call(__e, PrimFunc(symshen_4spaces), V5454)


tmp13719 := Call(__e, PrimFunc(symshen_4app), tmp13718, MakeString(""), symshen_4a)


tmp13720 := PrimStringConcat(MakeString(" \n"), tmp13719)

tmp13721 := Call(__e, PrimFunc(symshen_4app), V5455, tmp13720, symshen_4a)


tmp13722 := PrimStringConcat(MakeString("> Inputs to "), tmp13721)

tmp13723 := Call(__e, PrimFunc(symshen_4app), V5454, tmp13722, symshen_4a)


tmp13724 := PrimStringConcat(MakeString("<"), tmp13723)

tmp13725 := Call(__e, PrimFunc(symshen_4app), tmp13717, tmp13724, symshen_4a)


tmp13726 := PrimStringConcat(MakeString("\n"), tmp13725)

tmp13727 := Call(__e, PrimFunc(symstoutput))


tmp13728 := Call(__e, PrimFunc(sympr), tmp13726, tmp13727)


_ = tmp13728

__e.TailApply(PrimFunc(symshen_4recursively_1print), V5456)
return


}, 3)

tmp13729 := Call(__e, ns2_1set, symshen_4input_1track, tmp13716)


_ = tmp13729

tmp13730 := MakeNative(func(__e *ControlFlow) {
V5459 := __e.Get(1)
_ = V5459
tmp13740 := PrimEqual(Nil, V5459)

if True == tmp13740 {
tmp13731 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(" ==>"), tmp13731)
return


} else {
tmp13738 := PrimIsPair(V5459)

if True == tmp13738 {
tmp13732 := PrimHead(V5459)

tmp13733 := Call(__e, PrimFunc(symprint), tmp13732)


_ = tmp13733

tmp13734 := Call(__e, PrimFunc(symstoutput))


tmp13735 := Call(__e, PrimFunc(sympr), MakeString(", "), tmp13734)


_ = tmp13735

tmp13736 := PrimTail(V5459)

__e.TailApply(PrimFunc(symshen_4recursively_1print), tmp13736)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.recursively-print")))
return
}


}


}, 1)

tmp13741 := Call(__e, ns2_1set, symshen_4recursively_1print, tmp13730)


_ = tmp13741

tmp13742 := MakeNative(func(__e *ControlFlow) {
V5460 := __e.Get(1)
_ = V5460
tmp13746 := PrimEqual(MakeNumber(0), V5460)

if True == tmp13746 {
__e.Return(MakeString(""))
return
} else {
tmp13743 := PrimNumberSubtract(V5460, MakeNumber(1))

tmp13744 := Call(__e, PrimFunc(symshen_4spaces), tmp13743)


__e.Return(PrimStringConcat(MakeString(" "), tmp13744))
return


}


}, 1)

tmp13747 := Call(__e, ns2_1set, symshen_4spaces, tmp13742)


_ = tmp13747

tmp13748 := MakeNative(func(__e *ControlFlow) {
V5461 := __e.Get(1)
_ = V5461
V5462 := __e.Get(2)
_ = V5462
V5463 := __e.Get(3)
_ = V5463
tmp13749 := Call(__e, PrimFunc(symshen_4spaces), V5461)


tmp13750 := Call(__e, PrimFunc(symshen_4spaces), V5461)


tmp13751 := Call(__e, PrimFunc(symshen_4app), V5463, MakeString(""), symshen_4s)


tmp13752 := PrimStringConcat(MakeString("==> "), tmp13751)

tmp13753 := Call(__e, PrimFunc(symshen_4app), tmp13750, tmp13752, symshen_4a)


tmp13754 := PrimStringConcat(MakeString(" \n"), tmp13753)

tmp13755 := Call(__e, PrimFunc(symshen_4app), V5462, tmp13754, symshen_4a)


tmp13756 := PrimStringConcat(MakeString("> Output from "), tmp13755)

tmp13757 := Call(__e, PrimFunc(symshen_4app), V5461, tmp13756, symshen_4a)


tmp13758 := PrimStringConcat(MakeString("<"), tmp13757)

tmp13759 := Call(__e, PrimFunc(symshen_4app), tmp13749, tmp13758, symshen_4a)


tmp13760 := PrimStringConcat(MakeString("\n"), tmp13759)

tmp13761 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp13760, tmp13761)
return


}, 3)

tmp13762 := Call(__e, ns2_1set, symshen_4output_1track, tmp13748)


_ = tmp13762

tmp13763 := MakeNative(func(__e *ControlFlow) {
V5464 := __e.Get(1)
_ = V5464
tmp13764 := PrimValue(symshen_4_dtracking_d)

tmp13765 := Call(__e, PrimFunc(symremove), V5464, tmp13764)


tmp13766 := PrimSet(symshen_4_dtracking_d, tmp13765)

_ = tmp13766

tmp13767 := MakeNative(func(__e *ControlFlow) {
tmp13768 := Call(__e, PrimFunc(symps), V5464)


__e.TailApply(PrimFunc(symeval), tmp13768)
return


}, 0)

tmp13769 := MakeNative(func(__e *ControlFlow) {
Z5465 := __e.Get(1)
_ = Z5465
__e.Return(V5464)
return
}, 1)

tmp13770 := Call(__e, try_1catch, tmp13767, tmp13769)


_ = tmp13770

__e.Return(V5464)
return


}, 1)

tmp13771 := Call(__e, ns2_1set, symuntrack, tmp13763)


_ = tmp13771

tmp13772 := MakeNative(func(__e *ControlFlow) {
V5466 := __e.Get(1)
_ = V5466
V5467 := __e.Get(2)
_ = V5467
__e.TailApply(PrimFunc(symshen_4remove_1h), V5466, V5467, Nil)
return
}, 2)

tmp13773 := Call(__e, ns2_1set, symremove, tmp13772)


_ = tmp13773

tmp13774 := MakeNative(func(__e *ControlFlow) {
V5477 := __e.Get(1)
_ = V5477
V5478 := __e.Get(2)
_ = V5478
V5479 := __e.Get(3)
_ = V5479
tmp13789 := PrimEqual(Nil, V5478)

if True == tmp13789 {
__e.TailApply(PrimFunc(symreverse), V5479)
return
} else {
tmp13787 := PrimIsPair(V5478)

var ifres13783 Obj

if True == tmp13787 {
tmp13785 := PrimHead(V5478)

tmp13786 := PrimEqual(V5477, tmp13785)

var ifres13784 Obj

if True == tmp13786 {
ifres13784 = True


} else {
ifres13784 = False


}

ifres13783 = ifres13784


} else {
ifres13783 = False


}

if True == ifres13783 {
tmp13775 := PrimHead(V5478)

tmp13776 := PrimTail(V5478)

__e.TailApply(PrimFunc(symshen_4remove_1h), tmp13775, tmp13776, V5479)
return


} else {
tmp13781 := PrimIsPair(V5478)

if True == tmp13781 {
tmp13777 := PrimTail(V5478)

tmp13778 := PrimHead(V5478)

tmp13779 := PrimCons(tmp13778, V5479)

__e.TailApply(PrimFunc(symshen_4remove_1h), V5477, tmp13777, tmp13779)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-h")))
return
}


}


}


}, 3)

tmp13790 := Call(__e, ns2_1set, symshen_4remove_1h, tmp13774)


_ = tmp13790

tmp13791 := MakeNative(func(__e *ControlFlow) {
V5480 := __e.Get(1)
_ = V5480
tmp13792 := PrimValue(symshen_4_dprofiled_d)

tmp13793 := PrimCons(V5480, tmp13792)

tmp13794 := PrimSet(symshen_4_dprofiled_d, tmp13793)

_ = tmp13794

tmp13795 := Call(__e, PrimFunc(symps), V5480)


__e.TailApply(PrimFunc(symshen_4profile_1help), tmp13795)
return


}, 1)

tmp13796 := Call(__e, ns2_1set, symprofile, tmp13791)


_ = tmp13796

tmp13797 := MakeNative(func(__e *ControlFlow) {
V5483 := __e.Get(1)
_ = V5483
tmp13867 := PrimIsPair(V5483)

var ifres13841 Obj

if True == tmp13867 {
tmp13865 := PrimHead(V5483)

tmp13866 := PrimEqual(symdefun, tmp13865)

var ifres13843 Obj

if True == tmp13866 {
tmp13863 := PrimTail(V5483)

tmp13864 := PrimIsPair(tmp13863)

var ifres13845 Obj

if True == tmp13864 {
tmp13860 := PrimTail(V5483)

tmp13861 := PrimTail(tmp13860)

tmp13862 := PrimIsPair(tmp13861)

var ifres13847 Obj

if True == tmp13862 {
tmp13856 := PrimTail(V5483)

tmp13857 := PrimTail(tmp13856)

tmp13858 := PrimTail(tmp13857)

tmp13859 := PrimIsPair(tmp13858)

var ifres13849 Obj

if True == tmp13859 {
tmp13851 := PrimTail(V5483)

tmp13852 := PrimTail(tmp13851)

tmp13853 := PrimTail(tmp13852)

tmp13854 := PrimTail(tmp13853)

tmp13855 := PrimEqual(Nil, tmp13854)

var ifres13850 Obj

if True == tmp13855 {
ifres13850 = True


} else {
ifres13850 = False


}

ifres13849 = ifres13850


} else {
ifres13849 = False


}

var ifres13848 Obj

if True == ifres13849 {
ifres13848 = True


} else {
ifres13848 = False


}

ifres13847 = ifres13848


} else {
ifres13847 = False


}

var ifres13846 Obj

if True == ifres13847 {
ifres13846 = True


} else {
ifres13846 = False


}

ifres13845 = ifres13846


} else {
ifres13845 = False


}

var ifres13844 Obj

if True == ifres13845 {
ifres13844 = True


} else {
ifres13844 = False


}

ifres13843 = ifres13844


} else {
ifres13843 = False


}

var ifres13842 Obj

if True == ifres13843 {
ifres13842 = True


} else {
ifres13842 = False


}

ifres13841 = ifres13842


} else {
ifres13841 = False


}

if True == ifres13841 {
tmp13798 := MakeNative(func(__e *ControlFlow) {
W5484 := __e.Get(1)
_ = W5484
tmp13799 := MakeNative(func(__e *ControlFlow) {
W5485 := __e.Get(1)
_ = W5485
tmp13800 := MakeNative(func(__e *ControlFlow) {
W5486 := __e.Get(1)
_ = W5486
tmp13801 := MakeNative(func(__e *ControlFlow) {
W5487 := __e.Get(1)
_ = W5487
tmp13802 := MakeNative(func(__e *ControlFlow) {
W5488 := __e.Get(1)
_ = W5488
tmp13803 := PrimTail(V5483)

__e.Return(PrimHead(tmp13803))
return


}, 1)

tmp13804 := Call(__e, PrimFunc(symeval_1kl), W5486)


__e.TailApply(tmp13802, tmp13804)
return


}, 1)

tmp13805 := Call(__e, PrimFunc(symeval_1kl), W5485)


__e.TailApply(tmp13801, tmp13805)
return


}, 1)

tmp13806 := PrimTail(V5483)

tmp13807 := PrimTail(tmp13806)

tmp13808 := PrimHead(tmp13807)

tmp13809 := PrimTail(V5483)

tmp13810 := PrimHead(tmp13809)

tmp13811 := PrimTail(V5483)

tmp13812 := PrimTail(tmp13811)

tmp13813 := PrimTail(tmp13812)

tmp13814 := PrimHead(tmp13813)

tmp13815 := Call(__e, PrimFunc(symsubst), W5484, tmp13810, tmp13814)


tmp13816 := PrimCons(tmp13815, Nil)

tmp13817 := PrimCons(tmp13808, tmp13816)

tmp13818 := PrimCons(W5484, tmp13817)

tmp13819 := PrimCons(symdefun, tmp13818)

__e.TailApply(tmp13800, tmp13819)
return


}, 1)

tmp13820 := PrimTail(V5483)

tmp13821 := PrimHead(tmp13820)

tmp13822 := PrimTail(V5483)

tmp13823 := PrimTail(tmp13822)

tmp13824 := PrimHead(tmp13823)

tmp13825 := PrimTail(V5483)

tmp13826 := PrimHead(tmp13825)

tmp13827 := PrimTail(V5483)

tmp13828 := PrimTail(tmp13827)

tmp13829 := PrimHead(tmp13828)

tmp13830 := PrimTail(V5483)

tmp13831 := PrimTail(tmp13830)

tmp13832 := PrimHead(tmp13831)

tmp13833 := PrimCons(W5484, tmp13832)

tmp13834 := Call(__e, PrimFunc(symshen_4profile_1func), tmp13826, tmp13829, tmp13833)


tmp13835 := PrimCons(tmp13834, Nil)

tmp13836 := PrimCons(tmp13824, tmp13835)

tmp13837 := PrimCons(tmp13821, tmp13836)

tmp13838 := PrimCons(symdefun, tmp13837)

__e.TailApply(tmp13799, tmp13838)
return


}, 1)

tmp13839 := Call(__e, PrimFunc(symgensym), symshen_4f)


__e.TailApply(tmp13798, tmp13839)
return


} else {
__e.Return(PrimSimpleError(MakeString("Cannot profile.\n")))
return
}


}, 1)

tmp13868 := Call(__e, ns2_1set, symshen_4profile_1help, tmp13797)


_ = tmp13868

tmp13869 := MakeNative(func(__e *ControlFlow) {
V5489 := __e.Get(1)
_ = V5489
tmp13870 := PrimValue(symshen_4_dprofiled_d)

tmp13871 := Call(__e, PrimFunc(symremove), V5489, tmp13870)


tmp13872 := PrimSet(symshen_4_dprofiled_d, tmp13871)

_ = tmp13872

tmp13873 := MakeNative(func(__e *ControlFlow) {
tmp13874 := Call(__e, PrimFunc(symps), V5489)


__e.TailApply(PrimFunc(symeval), tmp13874)
return


}, 0)

tmp13875 := MakeNative(func(__e *ControlFlow) {
Z5490 := __e.Get(1)
_ = Z5490
__e.Return(V5489)
return
}, 1)

__e.TailApply(try_1catch, tmp13873, tmp13875)
return


}, 1)

tmp13876 := Call(__e, ns2_1set, symunprofile, tmp13869)


_ = tmp13876

tmp13877 := MakeNative(func(__e *ControlFlow) {
V5491 := __e.Get(1)
_ = V5491
tmp13878 := PrimValue(symshen_4_dprofiled_d)

__e.TailApply(PrimFunc(symelement_2), V5491, tmp13878)
return


}, 1)

tmp13879 := Call(__e, ns2_1set, symshen_4profiled_2, tmp13877)


_ = tmp13879

tmp13880 := MakeNative(func(__e *ControlFlow) {
V5492 := __e.Get(1)
_ = V5492
V5493 := __e.Get(2)
_ = V5493
V5494 := __e.Get(3)
_ = V5494
tmp13881 := PrimCons(symrun, Nil)

tmp13882 := PrimCons(symget_1time, tmp13881)

tmp13883 := PrimCons(symrun, Nil)

tmp13884 := PrimCons(symget_1time, tmp13883)

tmp13885 := PrimCons(symStart, Nil)

tmp13886 := PrimCons(tmp13884, tmp13885)

tmp13887 := PrimCons(sym_1, tmp13886)

tmp13888 := PrimCons(V5492, Nil)

tmp13889 := PrimCons(symshen_4get_1profile, tmp13888)

tmp13890 := PrimCons(symFinish, Nil)

tmp13891 := PrimCons(tmp13889, tmp13890)

tmp13892 := PrimCons(sym_7, tmp13891)

tmp13893 := PrimCons(tmp13892, Nil)

tmp13894 := PrimCons(V5492, tmp13893)

tmp13895 := PrimCons(symshen_4put_1profile, tmp13894)

tmp13896 := PrimCons(symResult, Nil)

tmp13897 := PrimCons(tmp13895, tmp13896)

tmp13898 := PrimCons(symRecord, tmp13897)

tmp13899 := PrimCons(symlet, tmp13898)

tmp13900 := PrimCons(tmp13899, Nil)

tmp13901 := PrimCons(tmp13887, tmp13900)

tmp13902 := PrimCons(symFinish, tmp13901)

tmp13903 := PrimCons(symlet, tmp13902)

tmp13904 := PrimCons(tmp13903, Nil)

tmp13905 := PrimCons(V5494, tmp13904)

tmp13906 := PrimCons(symResult, tmp13905)

tmp13907 := PrimCons(symlet, tmp13906)

tmp13908 := PrimCons(tmp13907, Nil)

tmp13909 := PrimCons(tmp13882, tmp13908)

tmp13910 := PrimCons(symStart, tmp13909)

__e.Return(PrimCons(symlet, tmp13910))
return


}, 3)

tmp13911 := Call(__e, ns2_1set, symshen_4profile_1func, tmp13880)


_ = tmp13911

tmp13912 := MakeNative(func(__e *ControlFlow) {
V5495 := __e.Get(1)
_ = V5495
tmp13913 := MakeNative(func(__e *ControlFlow) {
W5496 := __e.Get(1)
_ = W5496
tmp13914 := MakeNative(func(__e *ControlFlow) {
W5497 := __e.Get(1)
_ = W5497
__e.TailApply(PrimFunc(sym_8p), V5495, W5496)
return
}, 1)

tmp13915 := Call(__e, PrimFunc(symshen_4put_1profile), V5495, MakeNumber(0))


__e.TailApply(tmp13914, tmp13915)
return


}, 1)

tmp13916 := Call(__e, PrimFunc(symshen_4get_1profile), V5495)


__e.TailApply(tmp13913, tmp13916)
return


}, 1)

tmp13917 := Call(__e, ns2_1set, symprofile_1results, tmp13912)


_ = tmp13917

tmp13918 := MakeNative(func(__e *ControlFlow) {
V5498 := __e.Get(1)
_ = V5498
tmp13919 := MakeNative(func(__e *ControlFlow) {
tmp13920 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V5498, symprofile, tmp13920)
return


}, 0)

tmp13921 := MakeNative(func(__e *ControlFlow) {
Z5499 := __e.Get(1)
_ = Z5499
__e.Return(MakeNumber(0))
return
}, 1)

__e.TailApply(try_1catch, tmp13919, tmp13921)
return


}, 1)

tmp13922 := Call(__e, ns2_1set, symshen_4get_1profile, tmp13918)


_ = tmp13922

tmp13923 := MakeNative(func(__e *ControlFlow) {
V5500 := __e.Get(1)
_ = V5500
V5501 := __e.Get(2)
_ = V5501
tmp13924 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V5500, symprofile, V5501, tmp13924)
return


}, 2)

__e.TailApply(ns2_1set, symshen_4put_1profile, tmp13923)
return




}, 0)

