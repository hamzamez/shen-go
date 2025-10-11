package main

import . "github.com/tiancaiamao/shen-go/kl"

var TypesMain = MakeNative(func(__e *ControlFlow) {
tmp18369 := MakeNative(func(__e *ControlFlow) {
V5540 := __e.Get(1)
_ = V5540
V5541 := __e.Get(2)
_ = V5541
tmp18370 := MakeNative(func(__e *ControlFlow) {
W5542 := __e.Get(1)
_ = W5542
tmp18371 := MakeNative(func(__e *ControlFlow) {
W5543 := __e.Get(1)
_ = W5543
tmp18372 := MakeNative(func(__e *ControlFlow) {
W5548 := __e.Get(1)
_ = W5548
tmp18373 := MakeNative(func(__e *ControlFlow) {
W5549 := __e.Get(1)
_ = W5549
__e.Return(V5540)
return
}, 1)

tmp18374 := PrimValue(symshen_4_dsigf_d)

tmp18375 := Call(__e, PrimFunc(symshen_4assoc_1_6), V5540, W5548, tmp18374)


tmp18376 := PrimSet(symshen_4_dsigf_d, tmp18375)

__e.TailApply(tmp18373, tmp18376)
return


}, 1)

tmp18377 := Call(__e, PrimFunc(symshen_4prolog_1abstraction), V5541)


tmp18378 := Call(__e, PrimFunc(symeval_1kl), tmp18377)


__e.TailApply(tmp18372, tmp18378)
return


}, 1)

tmp18379 := MakeNative(func(__e *ControlFlow) {
Z5544 := __e.Get(1)
_ = Z5544
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5545 := __e.Get(1)
_ = Z5545
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5546 := __e.Get(1)
_ = Z5546
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5547 := __e.Get(1)
_ = Z5547
tmp18380 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18380

tmp18381 := Call(__e, PrimFunc(symshen_4deref), V5540, Z5544)


tmp18382 := Call(__e, PrimFunc(symreceive), tmp18381)


tmp18383 := Call(__e, PrimFunc(symshen_4deref), W5542, Z5544)


tmp18384 := Call(__e, PrimFunc(symreceive), tmp18383)


__e.TailApply(PrimFunc(symshen_4variancy), tmp18382, tmp18384, Z5544, Z5545, Z5546, Z5547)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18385 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp18386 := Call(__e, tmp18379, tmp18385)


tmp18387 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp18388 := Call(__e, PrimFunc(sym_8v), MakeNumber(0), tmp18387)


tmp18389 := Call(__e, PrimFunc(sym_8v), True, tmp18388)


tmp18390 := Call(__e, tmp18386, tmp18389)


tmp18391 := Call(__e, tmp18390, MakeNumber(0))


tmp18392 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

tmp18393 := Call(__e, tmp18391, tmp18392)


__e.TailApply(tmp18371, tmp18393)
return


}, 1)

tmp18394 := Call(__e, PrimFunc(symshen_4rectify_1type), V5541)


__e.TailApply(tmp18370, tmp18394)
return


}, 2)

tmp18395 := Call(__e, ns2_1set, symdeclare, tmp18369)


_ = tmp18395

tmp18396 := MakeNative(func(__e *ControlFlow) {
V5550 := __e.Get(1)
_ = V5550
V5551 := __e.Get(2)
_ = V5551
V5552 := __e.Get(3)
_ = V5552
V5553 := __e.Get(4)
_ = V5553
V5554 := __e.Get(5)
_ = V5554
V5555 := __e.Get(6)
_ = V5555
tmp18403 := Call(__e, PrimFunc(symshen_4unlocked_2), V5553)


if True == tmp18403 {
tmp18397 := MakeNative(func(__e *ControlFlow) {
W5556 := __e.Get(1)
_ = W5556
tmp18398 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18398

tmp18399 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4variants_2), V5550, W5556, V5551, V5552, V5553, V5554, V5555)
return
}, 0)

tmp18400 := Call(__e, PrimFunc(symshen_4system_1S_1h), V5550, W5556, Nil, V5552, V5553, V5554, tmp18399)


__e.TailApply(PrimFunc(symshen_4gc), V5552, tmp18400)
return


}, 1)

tmp18401 := Call(__e, PrimFunc(symshen_4newpv), V5552)


__e.TailApply(tmp18397, tmp18401)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp18404 := Call(__e, ns2_1set, symshen_4variancy, tmp18396)


_ = tmp18404

tmp18405 := MakeNative(func(__e *ControlFlow) {
V5557 := __e.Get(1)
_ = V5557
V5558 := __e.Get(2)
_ = V5558
V5559 := __e.Get(3)
_ = V5559
V5560 := __e.Get(4)
_ = V5560
V5561 := __e.Get(5)
_ = V5561
V5562 := __e.Get(6)
_ = V5562
V5563 := __e.Get(7)
_ = V5563
tmp18406 := MakeNative(func(__e *ControlFlow) {
W5564 := __e.Get(1)
_ = W5564
tmp18407 := MakeNative(func(__e *ControlFlow) {
W5565 := __e.Get(1)
_ = W5565
tmp18431 := PrimEqual(W5565, False)

if True == tmp18431 {
tmp18408 := MakeNative(func(__e *ControlFlow) {
W5568 := __e.Get(1)
_ = W5568
tmp18425 := PrimEqual(W5568, False)

if True == tmp18425 {
tmp18409 := MakeNative(func(__e *ControlFlow) {
W5569 := __e.Get(1)
_ = W5569
tmp18411 := PrimEqual(W5569, False)

if True == tmp18411 {
__e.TailApply(PrimFunc(symshen_4unlock), V5561, W5564)
return
} else {
__e.Return(W5569)
return
}


}, 1)

tmp18423 := Call(__e, PrimFunc(symshen_4unlocked_2), V5561)


var ifres18412 Obj

if True == tmp18423 {
tmp18413 := MakeNative(func(__e *ControlFlow) {
W5570 := __e.Get(1)
_ = W5570
tmp18414 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18414

tmp18415 := Call(__e, PrimFunc(symshen_4deref), V5557, V5560)


tmp18416 := Call(__e, PrimFunc(symshen_4app), tmp18415, MakeString(" may create errors\n"), symshen_4a)


tmp18417 := PrimStringConcat(MakeString("warning: changing the type of "), tmp18416)

tmp18418 := Call(__e, PrimFunc(symstoutput))


tmp18419 := Call(__e, PrimFunc(sympr), tmp18417, tmp18418)


tmp18420 := Call(__e, PrimFunc(symis), W5570, tmp18419, V5560, V5561, W5564, V5563)


__e.TailApply(PrimFunc(symshen_4gc), V5560, tmp18420)
return


}, 1)

tmp18421 := Call(__e, PrimFunc(symshen_4newpv), V5560)


tmp18422 := Call(__e, tmp18413, tmp18421)


ifres18412 = tmp18422


} else {
ifres18412 = False


}

__e.TailApply(tmp18409, ifres18412)
return


} else {
__e.Return(W5568)
return
}


}, 1)

tmp18429 := Call(__e, PrimFunc(symshen_4unlocked_2), V5561)


var ifres18426 Obj

if True == tmp18429 {
tmp18427 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18427

tmp18428 := Call(__e, PrimFunc(symis_b), V5558, V5559, V5560, V5561, W5564, V5563)


ifres18426 = tmp18428


} else {
ifres18426 = False


}

__e.TailApply(tmp18408, ifres18426)
return


} else {
__e.Return(W5565)
return
}


}, 1)

tmp18443 := Call(__e, PrimFunc(symshen_4unlocked_2), V5561)


var ifres18432 Obj

if True == tmp18443 {
tmp18433 := MakeNative(func(__e *ControlFlow) {
W5566 := __e.Get(1)
_ = W5566
tmp18434 := MakeNative(func(__e *ControlFlow) {
W5567 := __e.Get(1)
_ = W5567
tmp18438 := PrimEqual(W5566, symsymbol)

if True == tmp18438 {
__e.TailApply(PrimFunc(symthaw), W5567)
return
} else {
tmp18436 := Call(__e, PrimFunc(symshen_4pvar_2), W5566)


if True == tmp18436 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5566, symsymbol, V5560, W5567)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp18439 := MakeNative(func(__e *ControlFlow) {
tmp18440 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18440

__e.TailApply(PrimFunc(symshen_4cut), V5560, V5561, W5564, V5563)
return


}, 0)

__e.TailApply(tmp18434, tmp18439)
return


}, 1)

tmp18441 := Call(__e, PrimFunc(symshen_4lazyderef), V5558, V5560)


tmp18442 := Call(__e, tmp18433, tmp18441)


ifres18432 = tmp18442


} else {
ifres18432 = False


}

__e.TailApply(tmp18407, ifres18432)
return


}, 1)

tmp18444 := PrimNumberAdd(V5562, MakeNumber(1))

__e.TailApply(tmp18406, tmp18444)
return


}, 7)

tmp18445 := Call(__e, ns2_1set, symshen_4variants_2, tmp18405)


_ = tmp18445

tmp18446 := MakeNative(func(__e *ControlFlow) {
V5571 := __e.Get(1)
_ = V5571
tmp18447 := MakeNative(func(__e *ControlFlow) {
W5572 := __e.Get(1)
_ = W5572
tmp18448 := MakeNative(func(__e *ControlFlow) {
W5573 := __e.Get(1)
_ = W5573
tmp18449 := MakeNative(func(__e *ControlFlow) {
W5574 := __e.Get(1)
_ = W5574
tmp18450 := MakeNative(func(__e *ControlFlow) {
W5575 := __e.Get(1)
_ = W5575
tmp18451 := MakeNative(func(__e *ControlFlow) {
W5576 := __e.Get(1)
_ = W5576
tmp18452 := MakeNative(func(__e *ControlFlow) {
W5577 := __e.Get(1)
_ = W5577
tmp18453 := Call(__e, PrimFunc(symshen_4rcons__form), V5571)


tmp18454 := PrimCons(W5575, Nil)

tmp18455 := PrimCons(W5574, tmp18454)

tmp18456 := PrimCons(W5573, tmp18455)

tmp18457 := PrimCons(W5572, tmp18456)

tmp18458 := PrimCons(tmp18453, tmp18457)

tmp18459 := PrimCons(W5576, tmp18458)

tmp18460 := PrimCons(symis_b, tmp18459)

tmp18461 := Call(__e, PrimFunc(symshen_4stpart), W5577, tmp18460, W5572)


tmp18462 := PrimCons(tmp18461, Nil)

tmp18463 := PrimCons(W5575, tmp18462)

tmp18464 := PrimCons(symlambda, tmp18463)

tmp18465 := PrimCons(tmp18464, Nil)

tmp18466 := PrimCons(W5574, tmp18465)

tmp18467 := PrimCons(symlambda, tmp18466)

tmp18468 := PrimCons(tmp18467, Nil)

tmp18469 := PrimCons(W5573, tmp18468)

tmp18470 := PrimCons(symlambda, tmp18469)

tmp18471 := PrimCons(tmp18470, Nil)

tmp18472 := PrimCons(W5572, tmp18471)

tmp18473 := PrimCons(symlambda, tmp18472)

tmp18474 := PrimCons(tmp18473, Nil)

tmp18475 := PrimCons(W5576, tmp18474)

__e.Return(PrimCons(symlambda, tmp18475))
return


}, 1)

tmp18476 := Call(__e, PrimFunc(symshen_4extract_1vars), V5571)


__e.TailApply(tmp18452, tmp18476)
return


}, 1)

tmp18477 := Call(__e, PrimFunc(symgensym), symV)


__e.TailApply(tmp18451, tmp18477)
return


}, 1)

tmp18478 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp18450, tmp18478)
return


}, 1)

tmp18479 := Call(__e, PrimFunc(symgensym), symKey)


__e.TailApply(tmp18449, tmp18479)
return


}, 1)

tmp18480 := Call(__e, PrimFunc(symgensym), symL)


__e.TailApply(tmp18448, tmp18480)
return


}, 1)

tmp18481 := Call(__e, PrimFunc(symgensym), symB)


__e.TailApply(tmp18447, tmp18481)
return


}, 1)

tmp18482 := Call(__e, ns2_1set, symshen_4prolog_1abstraction, tmp18446)


_ = tmp18482

tmp18483 := MakeNative(func(__e *ControlFlow) {
V5578 := __e.Get(1)
_ = V5578
__e.Return(V5578)
return
}, 1)

tmp18484 := Call(__e, ns2_1set, symshen_4demod, tmp18483)


_ = tmp18484

tmp18485 := PrimCons(symA, Nil)

tmp18486 := PrimCons(sym_1_1_6, tmp18485)

tmp18487 := Call(__e, PrimFunc(symdeclare), symabort, tmp18486)


_ = tmp18487

tmp18488 := PrimCons(symboolean, Nil)

tmp18489 := PrimCons(sym_1_1_6, tmp18488)

tmp18490 := PrimCons(symA, tmp18489)

tmp18491 := Call(__e, PrimFunc(symdeclare), symabsvector_2, tmp18490)


_ = tmp18491

tmp18492 := PrimCons(symA, Nil)

tmp18493 := PrimCons(symlist, tmp18492)

tmp18494 := PrimCons(symA, Nil)

tmp18495 := PrimCons(symlist, tmp18494)

tmp18496 := PrimCons(tmp18495, Nil)

tmp18497 := PrimCons(sym_1_1_6, tmp18496)

tmp18498 := PrimCons(tmp18493, tmp18497)

tmp18499 := PrimCons(tmp18498, Nil)

tmp18500 := PrimCons(sym_1_1_6, tmp18499)

tmp18501 := PrimCons(symA, tmp18500)

tmp18502 := Call(__e, PrimFunc(symdeclare), symadjoin, tmp18501)


_ = tmp18502

tmp18503 := PrimCons(symboolean, Nil)

tmp18504 := PrimCons(sym_1_1_6, tmp18503)

tmp18505 := PrimCons(symboolean, tmp18504)

tmp18506 := PrimCons(tmp18505, Nil)

tmp18507 := PrimCons(sym_1_1_6, tmp18506)

tmp18508 := PrimCons(symboolean, tmp18507)

tmp18509 := Call(__e, PrimFunc(symdeclare), symand, tmp18508)


_ = tmp18509

tmp18510 := PrimCons(symstring, Nil)

tmp18511 := PrimCons(sym_1_1_6, tmp18510)

tmp18512 := PrimCons(symsymbol, tmp18511)

tmp18513 := PrimCons(tmp18512, Nil)

tmp18514 := PrimCons(sym_1_1_6, tmp18513)

tmp18515 := PrimCons(symstring, tmp18514)

tmp18516 := PrimCons(tmp18515, Nil)

tmp18517 := PrimCons(sym_1_1_6, tmp18516)

tmp18518 := PrimCons(symA, tmp18517)

tmp18519 := Call(__e, PrimFunc(symdeclare), symshen_4app, tmp18518)


_ = tmp18519

tmp18520 := PrimCons(symA, Nil)

tmp18521 := PrimCons(symlist, tmp18520)

tmp18522 := PrimCons(symA, Nil)

tmp18523 := PrimCons(symlist, tmp18522)

tmp18524 := PrimCons(symA, Nil)

tmp18525 := PrimCons(symlist, tmp18524)

tmp18526 := PrimCons(tmp18525, Nil)

tmp18527 := PrimCons(sym_1_1_6, tmp18526)

tmp18528 := PrimCons(tmp18523, tmp18527)

tmp18529 := PrimCons(tmp18528, Nil)

tmp18530 := PrimCons(sym_1_1_6, tmp18529)

tmp18531 := PrimCons(tmp18521, tmp18530)

tmp18532 := Call(__e, PrimFunc(symdeclare), symappend, tmp18531)


_ = tmp18532

tmp18533 := PrimCons(symnumber, Nil)

tmp18534 := PrimCons(sym_1_1_6, tmp18533)

tmp18535 := PrimCons(symA, tmp18534)

tmp18536 := Call(__e, PrimFunc(symdeclare), symarity, tmp18535)


_ = tmp18536

tmp18537 := PrimCons(symA, Nil)

tmp18538 := PrimCons(symlist, tmp18537)

tmp18539 := PrimCons(tmp18538, Nil)

tmp18540 := PrimCons(symlist, tmp18539)

tmp18541 := PrimCons(symA, Nil)

tmp18542 := PrimCons(symlist, tmp18541)

tmp18543 := PrimCons(tmp18542, Nil)

tmp18544 := PrimCons(sym_1_1_6, tmp18543)

tmp18545 := PrimCons(tmp18540, tmp18544)

tmp18546 := PrimCons(tmp18545, Nil)

tmp18547 := PrimCons(sym_1_1_6, tmp18546)

tmp18548 := PrimCons(symA, tmp18547)

tmp18549 := Call(__e, PrimFunc(symdeclare), symassoc, tmp18548)


_ = tmp18549

tmp18550 := PrimCons(symboolean, Nil)

tmp18551 := PrimCons(sym_1_1_6, tmp18550)

tmp18552 := PrimCons(symA, tmp18551)

tmp18553 := Call(__e, PrimFunc(symdeclare), symatom_2, tmp18552)


_ = tmp18553

tmp18554 := PrimCons(symboolean, Nil)

tmp18555 := PrimCons(sym_1_1_6, tmp18554)

tmp18556 := PrimCons(symA, tmp18555)

tmp18557 := Call(__e, PrimFunc(symdeclare), symboolean_2, tmp18556)


_ = tmp18557

tmp18558 := PrimCons(symstring, Nil)

tmp18559 := PrimCons(sym_1_1_6, tmp18558)

tmp18560 := PrimCons(symstring, tmp18559)

tmp18561 := Call(__e, PrimFunc(symdeclare), symbootstrap, tmp18560)


_ = tmp18561

tmp18562 := PrimCons(symboolean, Nil)

tmp18563 := PrimCons(sym_1_1_6, tmp18562)

tmp18564 := PrimCons(symsymbol, tmp18563)

tmp18565 := Call(__e, PrimFunc(symdeclare), symbound_2, tmp18564)


_ = tmp18565

tmp18566 := PrimCons(symA, Nil)

tmp18567 := PrimCons(symlist, tmp18566)

tmp18568 := PrimCons(symboolean, Nil)

tmp18569 := PrimCons(sym_1_1_6, tmp18568)

tmp18570 := PrimCons(tmp18567, tmp18569)

tmp18571 := Call(__e, PrimFunc(symdeclare), symshen_4ccons_2, tmp18570)


_ = tmp18571

tmp18572 := PrimCons(symstring, Nil)

tmp18573 := PrimCons(sym_1_1_6, tmp18572)

tmp18574 := PrimCons(symstring, tmp18573)

tmp18575 := Call(__e, PrimFunc(symdeclare), symcd, tmp18574)


_ = tmp18575

tmp18576 := PrimCons(symA, Nil)

tmp18577 := PrimCons(symstream, tmp18576)

tmp18578 := PrimCons(symB, Nil)

tmp18579 := PrimCons(symlist, tmp18578)

tmp18580 := PrimCons(tmp18579, Nil)

tmp18581 := PrimCons(sym_1_1_6, tmp18580)

tmp18582 := PrimCons(tmp18577, tmp18581)

tmp18583 := Call(__e, PrimFunc(symdeclare), symclose, tmp18582)


_ = tmp18583

tmp18584 := PrimCons(symstring, Nil)

tmp18585 := PrimCons(sym_1_1_6, tmp18584)

tmp18586 := PrimCons(symstring, tmp18585)

tmp18587 := PrimCons(tmp18586, Nil)

tmp18588 := PrimCons(sym_1_1_6, tmp18587)

tmp18589 := PrimCons(symstring, tmp18588)

tmp18590 := Call(__e, PrimFunc(symdeclare), symcn, tmp18589)


_ = tmp18590

tmp18591 := PrimCons(symA, Nil)

tmp18592 := PrimCons(symlist, tmp18591)

tmp18593 := PrimCons(symA, Nil)

tmp18594 := PrimCons(symlist, tmp18593)

tmp18595 := PrimCons(symB, Nil)

tmp18596 := PrimCons(tmp18594, tmp18595)

tmp18597 := PrimCons(symstr, tmp18596)

tmp18598 := PrimCons(tmp18597, Nil)

tmp18599 := PrimCons(sym_1_1_6, tmp18598)

tmp18600 := PrimCons(tmp18592, tmp18599)

tmp18601 := PrimCons(symA, Nil)

tmp18602 := PrimCons(symlist, tmp18601)

tmp18603 := PrimCons(symB, Nil)

tmp18604 := PrimCons(sym_1_1_6, tmp18603)

tmp18605 := PrimCons(tmp18602, tmp18604)

tmp18606 := PrimCons(tmp18605, Nil)

tmp18607 := PrimCons(sym_1_1_6, tmp18606)

tmp18608 := PrimCons(tmp18600, tmp18607)

tmp18609 := Call(__e, PrimFunc(symdeclare), symcompile, tmp18608)


_ = tmp18609

tmp18610 := PrimCons(symboolean, Nil)

tmp18611 := PrimCons(sym_1_1_6, tmp18610)

tmp18612 := PrimCons(symA, tmp18611)

tmp18613 := Call(__e, PrimFunc(symdeclare), symcons_2, tmp18612)


_ = tmp18613

tmp18614 := PrimCons(symsymbol, Nil)

tmp18615 := PrimCons(symlist, tmp18614)

tmp18616 := PrimCons(tmp18615, Nil)

tmp18617 := PrimCons(sym_1_1_6, tmp18616)

tmp18618 := Call(__e, PrimFunc(symdeclare), symdatatypes, tmp18617)


_ = tmp18618

tmp18619 := PrimCons(symsymbol, Nil)

tmp18620 := PrimCons(sym_1_1_6, tmp18619)

tmp18621 := PrimCons(symsymbol, tmp18620)

tmp18622 := Call(__e, PrimFunc(symdeclare), symdestroy, tmp18621)


_ = tmp18622

tmp18623 := PrimCons(symA, Nil)

tmp18624 := PrimCons(symlist, tmp18623)

tmp18625 := PrimCons(symA, Nil)

tmp18626 := PrimCons(symlist, tmp18625)

tmp18627 := PrimCons(symA, Nil)

tmp18628 := PrimCons(symlist, tmp18627)

tmp18629 := PrimCons(tmp18628, Nil)

tmp18630 := PrimCons(sym_1_1_6, tmp18629)

tmp18631 := PrimCons(tmp18626, tmp18630)

tmp18632 := PrimCons(tmp18631, Nil)

tmp18633 := PrimCons(sym_1_1_6, tmp18632)

tmp18634 := PrimCons(tmp18624, tmp18633)

tmp18635 := Call(__e, PrimFunc(symdeclare), symdifference, tmp18634)


_ = tmp18635

tmp18636 := PrimCons(symB, Nil)

tmp18637 := PrimCons(sym_1_1_6, tmp18636)

tmp18638 := PrimCons(symB, tmp18637)

tmp18639 := PrimCons(tmp18638, Nil)

tmp18640 := PrimCons(sym_1_1_6, tmp18639)

tmp18641 := PrimCons(symA, tmp18640)

tmp18642 := Call(__e, PrimFunc(symdeclare), symdo, tmp18641)


_ = tmp18642

tmp18643 := PrimCons(symA, Nil)

tmp18644 := PrimCons(symlist, tmp18643)

tmp18645 := PrimCons(symA, Nil)

tmp18646 := PrimCons(symlist, tmp18645)

tmp18647 := PrimCons(symB, Nil)

tmp18648 := PrimCons(symlist, tmp18647)

tmp18649 := PrimCons(tmp18648, Nil)

tmp18650 := PrimCons(tmp18646, tmp18649)

tmp18651 := PrimCons(symstr, tmp18650)

tmp18652 := PrimCons(tmp18651, Nil)

tmp18653 := PrimCons(sym_1_1_6, tmp18652)

tmp18654 := PrimCons(tmp18644, tmp18653)

tmp18655 := Call(__e, PrimFunc(symdeclare), sym_5e_6, tmp18654)


_ = tmp18655

tmp18656 := PrimCons(symA, Nil)

tmp18657 := PrimCons(symlist, tmp18656)

tmp18658 := PrimCons(symB, Nil)

tmp18659 := PrimCons(symlist, tmp18658)

tmp18660 := PrimCons(symA, Nil)

tmp18661 := PrimCons(symlist, tmp18660)

tmp18662 := PrimCons(tmp18661, Nil)

tmp18663 := PrimCons(tmp18659, tmp18662)

tmp18664 := PrimCons(symstr, tmp18663)

tmp18665 := PrimCons(tmp18664, Nil)

tmp18666 := PrimCons(sym_1_1_6, tmp18665)

tmp18667 := PrimCons(tmp18657, tmp18666)

tmp18668 := Call(__e, PrimFunc(symdeclare), sym_5_b_6, tmp18667)


_ = tmp18668

tmp18669 := PrimCons(symA, Nil)

tmp18670 := PrimCons(symlist, tmp18669)

tmp18671 := PrimCons(symA, Nil)

tmp18672 := PrimCons(symlist, tmp18671)

tmp18673 := PrimCons(symB, Nil)

tmp18674 := PrimCons(symlist, tmp18673)

tmp18675 := PrimCons(tmp18674, Nil)

tmp18676 := PrimCons(tmp18672, tmp18675)

tmp18677 := PrimCons(symstr, tmp18676)

tmp18678 := PrimCons(tmp18677, Nil)

tmp18679 := PrimCons(sym_1_1_6, tmp18678)

tmp18680 := PrimCons(tmp18670, tmp18679)

tmp18681 := Call(__e, PrimFunc(symdeclare), sym_5end_6, tmp18680)


_ = tmp18681

tmp18682 := PrimCons(symA, Nil)

tmp18683 := PrimCons(symlist, tmp18682)

tmp18684 := PrimCons(symB, Nil)

tmp18685 := PrimCons(tmp18683, tmp18684)

tmp18686 := PrimCons(symstr, tmp18685)

tmp18687 := PrimCons(symboolean, Nil)

tmp18688 := PrimCons(sym_1_1_6, tmp18687)

tmp18689 := PrimCons(tmp18686, tmp18688)

tmp18690 := Call(__e, PrimFunc(symdeclare), symshen_4parse_1failure_2, tmp18689)


_ = tmp18690

tmp18691 := PrimCons(symA, Nil)

tmp18692 := PrimCons(symlist, tmp18691)

tmp18693 := PrimCons(symB, Nil)

tmp18694 := PrimCons(tmp18692, tmp18693)

tmp18695 := PrimCons(symstr, tmp18694)

tmp18696 := PrimCons(tmp18695, Nil)

tmp18697 := PrimCons(sym_1_1_6, tmp18696)

tmp18698 := Call(__e, PrimFunc(symdeclare), symshen_4parse_1failure, tmp18697)


_ = tmp18698

tmp18699 := PrimCons(symA, Nil)

tmp18700 := PrimCons(symlist, tmp18699)

tmp18701 := PrimCons(symB, Nil)

tmp18702 := PrimCons(tmp18700, tmp18701)

tmp18703 := PrimCons(symstr, tmp18702)

tmp18704 := PrimCons(symB, Nil)

tmp18705 := PrimCons(sym_1_1_6, tmp18704)

tmp18706 := PrimCons(tmp18703, tmp18705)

tmp18707 := Call(__e, PrimFunc(symdeclare), symshen_4_5_1out, tmp18706)


_ = tmp18707

tmp18708 := PrimCons(symA, Nil)

tmp18709 := PrimCons(symlist, tmp18708)

tmp18710 := PrimCons(symB, Nil)

tmp18711 := PrimCons(tmp18709, tmp18710)

tmp18712 := PrimCons(symstr, tmp18711)

tmp18713 := PrimCons(symA, Nil)

tmp18714 := PrimCons(symlist, tmp18713)

tmp18715 := PrimCons(tmp18714, Nil)

tmp18716 := PrimCons(sym_1_1_6, tmp18715)

tmp18717 := PrimCons(tmp18712, tmp18716)

tmp18718 := Call(__e, PrimFunc(symdeclare), symshen_4in_1_6, tmp18717)


_ = tmp18718

tmp18719 := PrimCons(symA, Nil)

tmp18720 := PrimCons(symlist, tmp18719)

tmp18721 := PrimCons(symA, Nil)

tmp18722 := PrimCons(symlist, tmp18721)

tmp18723 := PrimCons(symB, Nil)

tmp18724 := PrimCons(tmp18722, tmp18723)

tmp18725 := PrimCons(symstr, tmp18724)

tmp18726 := PrimCons(tmp18725, Nil)

tmp18727 := PrimCons(sym_1_1_6, tmp18726)

tmp18728 := PrimCons(symB, tmp18727)

tmp18729 := PrimCons(tmp18728, Nil)

tmp18730 := PrimCons(sym_1_1_6, tmp18729)

tmp18731 := PrimCons(tmp18720, tmp18730)

tmp18732 := Call(__e, PrimFunc(symdeclare), symshen_4comb, tmp18731)


_ = tmp18732

tmp18733 := PrimCons(symA, Nil)

tmp18734 := PrimCons(symlist, tmp18733)

tmp18735 := PrimCons(symboolean, Nil)

tmp18736 := PrimCons(sym_1_1_6, tmp18735)

tmp18737 := PrimCons(tmp18734, tmp18736)

tmp18738 := PrimCons(tmp18737, Nil)

tmp18739 := PrimCons(sym_1_1_6, tmp18738)

tmp18740 := PrimCons(symA, tmp18739)

tmp18741 := Call(__e, PrimFunc(symdeclare), symelement_2, tmp18740)


_ = tmp18741

tmp18742 := PrimCons(symboolean, Nil)

tmp18743 := PrimCons(sym_1_1_6, tmp18742)

tmp18744 := PrimCons(symA, tmp18743)

tmp18745 := Call(__e, PrimFunc(symdeclare), symempty_2, tmp18744)


_ = tmp18745

tmp18746 := PrimCons(symboolean, Nil)

tmp18747 := PrimCons(sym_1_1_6, tmp18746)

tmp18748 := PrimCons(symsymbol, tmp18747)

tmp18749 := Call(__e, PrimFunc(symdeclare), symenable_1type_1theory, tmp18748)


_ = tmp18749

tmp18750 := PrimCons(symsymbol, Nil)

tmp18751 := PrimCons(symlist, tmp18750)

tmp18752 := PrimCons(tmp18751, Nil)

tmp18753 := PrimCons(sym_1_1_6, tmp18752)

tmp18754 := PrimCons(symsymbol, tmp18753)

tmp18755 := Call(__e, PrimFunc(symdeclare), symexternal, tmp18754)


_ = tmp18755

tmp18756 := PrimCons(symstring, Nil)

tmp18757 := PrimCons(sym_1_1_6, tmp18756)

tmp18758 := PrimCons(symexception, tmp18757)

tmp18759 := Call(__e, PrimFunc(symdeclare), symerror_1to_1string, tmp18758)


_ = tmp18759

tmp18760 := PrimCons(symstring, Nil)

tmp18761 := PrimCons(symlist, tmp18760)

tmp18762 := PrimCons(tmp18761, Nil)

tmp18763 := PrimCons(sym_1_1_6, tmp18762)

tmp18764 := PrimCons(symA, tmp18763)

tmp18765 := Call(__e, PrimFunc(symdeclare), symexplode, tmp18764)


_ = tmp18765

tmp18766 := PrimCons(symsymbol, Nil)

tmp18767 := PrimCons(sym_1_1_6, tmp18766)

tmp18768 := PrimCons(symsymbol, tmp18767)

tmp18769 := Call(__e, PrimFunc(symdeclare), symfactorise, tmp18768)


_ = tmp18769

tmp18770 := PrimCons(symboolean, Nil)

tmp18771 := PrimCons(sym_1_1_6, tmp18770)

tmp18772 := Call(__e, PrimFunc(symdeclare), symfactorise_2, tmp18771)


_ = tmp18772

tmp18773 := PrimCons(symsymbol, Nil)

tmp18774 := PrimCons(sym_1_1_6, tmp18773)

tmp18775 := Call(__e, PrimFunc(symdeclare), symfail, tmp18774)


_ = tmp18775

tmp18776 := PrimCons(symA, Nil)

tmp18777 := PrimCons(sym_1_1_6, tmp18776)

tmp18778 := PrimCons(symA, tmp18777)

tmp18779 := PrimCons(symA, Nil)

tmp18780 := PrimCons(sym_1_1_6, tmp18779)

tmp18781 := PrimCons(symA, tmp18780)

tmp18782 := PrimCons(tmp18781, Nil)

tmp18783 := PrimCons(sym_1_1_6, tmp18782)

tmp18784 := PrimCons(tmp18778, tmp18783)

tmp18785 := Call(__e, PrimFunc(symdeclare), symfix, tmp18784)


_ = tmp18785

tmp18786 := PrimCons(symA, Nil)

tmp18787 := PrimCons(symlazy, tmp18786)

tmp18788 := PrimCons(tmp18787, Nil)

tmp18789 := PrimCons(sym_1_1_6, tmp18788)

tmp18790 := PrimCons(symA, tmp18789)

tmp18791 := Call(__e, PrimFunc(symdeclare), symfreeze, tmp18790)


_ = tmp18791

tmp18792 := PrimCons(symB, Nil)

tmp18793 := PrimCons(sym_d, tmp18792)

tmp18794 := PrimCons(symA, tmp18793)

tmp18795 := PrimCons(symA, Nil)

tmp18796 := PrimCons(sym_1_1_6, tmp18795)

tmp18797 := PrimCons(tmp18794, tmp18796)

tmp18798 := Call(__e, PrimFunc(symdeclare), symfst, tmp18797)


_ = tmp18798

tmp18799 := PrimCons(symsymbol, Nil)

tmp18800 := PrimCons(sym_1_1_6, tmp18799)

tmp18801 := PrimCons(symsymbol, tmp18800)

tmp18802 := Call(__e, PrimFunc(symdeclare), symgensym, tmp18801)


_ = tmp18802

tmp18803 := PrimCons(symA, Nil)

tmp18804 := PrimCons(symlist, tmp18803)

tmp18805 := PrimCons(symboolean, Nil)

tmp18806 := PrimCons(sym_1_1_6, tmp18805)

tmp18807 := PrimCons(symA, tmp18806)

tmp18808 := PrimCons(tmp18807, Nil)

tmp18809 := PrimCons(sym_1_1_6, tmp18808)

tmp18810 := PrimCons(tmp18804, tmp18809)

tmp18811 := Call(__e, PrimFunc(symdeclare), symshen_4hds_a_2, tmp18810)


_ = tmp18811

tmp18812 := PrimCons(symboolean, Nil)

tmp18813 := PrimCons(sym_1_1_6, tmp18812)

tmp18814 := PrimCons(symsymbol, tmp18813)

tmp18815 := Call(__e, PrimFunc(symdeclare), symhush, tmp18814)


_ = tmp18815

tmp18816 := PrimCons(symboolean, Nil)

tmp18817 := PrimCons(sym_1_1_6, tmp18816)

tmp18818 := Call(__e, PrimFunc(symdeclare), symhush_2, tmp18817)


_ = tmp18818

tmp18819 := PrimCons(symA, Nil)

tmp18820 := PrimCons(symvector, tmp18819)

tmp18821 := PrimCons(symA, Nil)

tmp18822 := PrimCons(sym_1_1_6, tmp18821)

tmp18823 := PrimCons(symnumber, tmp18822)

tmp18824 := PrimCons(tmp18823, Nil)

tmp18825 := PrimCons(sym_1_1_6, tmp18824)

tmp18826 := PrimCons(tmp18820, tmp18825)

tmp18827 := Call(__e, PrimFunc(symdeclare), sym_5_1vector, tmp18826)


_ = tmp18827

tmp18828 := PrimCons(symA, Nil)

tmp18829 := PrimCons(symvector, tmp18828)

tmp18830 := PrimCons(symA, Nil)

tmp18831 := PrimCons(symvector, tmp18830)

tmp18832 := PrimCons(tmp18831, Nil)

tmp18833 := PrimCons(sym_1_1_6, tmp18832)

tmp18834 := PrimCons(symA, tmp18833)

tmp18835 := PrimCons(tmp18834, Nil)

tmp18836 := PrimCons(sym_1_1_6, tmp18835)

tmp18837 := PrimCons(symnumber, tmp18836)

tmp18838 := PrimCons(tmp18837, Nil)

tmp18839 := PrimCons(sym_1_1_6, tmp18838)

tmp18840 := PrimCons(tmp18829, tmp18839)

tmp18841 := Call(__e, PrimFunc(symdeclare), symvector_1_6, tmp18840)


_ = tmp18841

tmp18842 := PrimCons(symA, Nil)

tmp18843 := PrimCons(symvector, tmp18842)

tmp18844 := PrimCons(tmp18843, Nil)

tmp18845 := PrimCons(sym_1_1_6, tmp18844)

tmp18846 := PrimCons(symnumber, tmp18845)

tmp18847 := Call(__e, PrimFunc(symdeclare), symvector, tmp18846)


_ = tmp18847

tmp18848 := PrimCons(symnumber, Nil)

tmp18849 := PrimCons(sym_1_1_6, tmp18848)

tmp18850 := PrimCons(symsymbol, tmp18849)

tmp18851 := Call(__e, PrimFunc(symdeclare), symget_1time, tmp18850)


_ = tmp18851

tmp18852 := PrimCons(symnumber, Nil)

tmp18853 := PrimCons(sym_1_1_6, tmp18852)

tmp18854 := PrimCons(symnumber, tmp18853)

tmp18855 := PrimCons(tmp18854, Nil)

tmp18856 := PrimCons(sym_1_1_6, tmp18855)

tmp18857 := PrimCons(symA, tmp18856)

tmp18858 := Call(__e, PrimFunc(symdeclare), symhash, tmp18857)


_ = tmp18858

tmp18859 := PrimCons(symA, Nil)

tmp18860 := PrimCons(symlist, tmp18859)

tmp18861 := PrimCons(symA, Nil)

tmp18862 := PrimCons(sym_1_1_6, tmp18861)

tmp18863 := PrimCons(tmp18860, tmp18862)

tmp18864 := Call(__e, PrimFunc(symdeclare), symhead, tmp18863)


_ = tmp18864

tmp18865 := PrimCons(symA, Nil)

tmp18866 := PrimCons(symvector, tmp18865)

tmp18867 := PrimCons(symA, Nil)

tmp18868 := PrimCons(sym_1_1_6, tmp18867)

tmp18869 := PrimCons(tmp18866, tmp18868)

tmp18870 := Call(__e, PrimFunc(symdeclare), symhdv, tmp18869)


_ = tmp18870

tmp18871 := PrimCons(symstring, Nil)

tmp18872 := PrimCons(sym_1_1_6, tmp18871)

tmp18873 := PrimCons(symstring, tmp18872)

tmp18874 := Call(__e, PrimFunc(symdeclare), symhdstr, tmp18873)


_ = tmp18874

tmp18875 := PrimCons(symA, Nil)

tmp18876 := PrimCons(sym_1_1_6, tmp18875)

tmp18877 := PrimCons(symA, tmp18876)

tmp18878 := PrimCons(tmp18877, Nil)

tmp18879 := PrimCons(sym_1_1_6, tmp18878)

tmp18880 := PrimCons(symA, tmp18879)

tmp18881 := PrimCons(tmp18880, Nil)

tmp18882 := PrimCons(sym_1_1_6, tmp18881)

tmp18883 := PrimCons(symboolean, tmp18882)

tmp18884 := Call(__e, PrimFunc(symdeclare), symif, tmp18883)


_ = tmp18884

tmp18885 := PrimCons(symsymbol, Nil)

tmp18886 := PrimCons(sym_1_1_6, tmp18885)

tmp18887 := PrimCons(symsymbol, tmp18886)

tmp18888 := Call(__e, PrimFunc(symdeclare), symin_1package, tmp18887)


_ = tmp18888

tmp18889 := PrimCons(symstring, Nil)

tmp18890 := PrimCons(sym_1_1_6, tmp18889)

tmp18891 := Call(__e, PrimFunc(symdeclare), symit, tmp18890)


_ = tmp18891

tmp18892 := PrimCons(symstring, Nil)

tmp18893 := PrimCons(sym_1_1_6, tmp18892)

tmp18894 := Call(__e, PrimFunc(symdeclare), symimplementation, tmp18893)


_ = tmp18894

tmp18895 := PrimCons(symsymbol, Nil)

tmp18896 := PrimCons(symlist, tmp18895)

tmp18897 := PrimCons(symsymbol, Nil)

tmp18898 := PrimCons(symlist, tmp18897)

tmp18899 := PrimCons(tmp18898, Nil)

tmp18900 := PrimCons(sym_1_1_6, tmp18899)

tmp18901 := PrimCons(tmp18896, tmp18900)

tmp18902 := Call(__e, PrimFunc(symdeclare), syminclude, tmp18901)


_ = tmp18902

tmp18903 := PrimCons(symsymbol, Nil)

tmp18904 := PrimCons(symlist, tmp18903)

tmp18905 := PrimCons(symsymbol, Nil)

tmp18906 := PrimCons(symlist, tmp18905)

tmp18907 := PrimCons(tmp18906, Nil)

tmp18908 := PrimCons(sym_1_1_6, tmp18907)

tmp18909 := PrimCons(tmp18904, tmp18908)

tmp18910 := Call(__e, PrimFunc(symdeclare), syminclude_1all_1but, tmp18909)


_ = tmp18910

tmp18911 := PrimCons(symsymbol, Nil)

tmp18912 := PrimCons(symlist, tmp18911)

tmp18913 := PrimCons(tmp18912, Nil)

tmp18914 := PrimCons(sym_1_1_6, tmp18913)

tmp18915 := Call(__e, PrimFunc(symdeclare), symincluded, tmp18914)


_ = tmp18915

tmp18916 := PrimCons(symnumber, Nil)

tmp18917 := PrimCons(sym_1_1_6, tmp18916)

tmp18918 := Call(__e, PrimFunc(symdeclare), syminferences, tmp18917)


_ = tmp18918

tmp18919 := PrimCons(symstring, Nil)

tmp18920 := PrimCons(sym_1_1_6, tmp18919)

tmp18921 := PrimCons(symstring, tmp18920)

tmp18922 := PrimCons(tmp18921, Nil)

tmp18923 := PrimCons(sym_1_1_6, tmp18922)

tmp18924 := PrimCons(symA, tmp18923)

tmp18925 := Call(__e, PrimFunc(symdeclare), symshen_4insert, tmp18924)


_ = tmp18925

tmp18926 := PrimCons(symboolean, Nil)

tmp18927 := PrimCons(sym_1_1_6, tmp18926)

tmp18928 := PrimCons(symA, tmp18927)

tmp18929 := Call(__e, PrimFunc(symdeclare), syminteger_2, tmp18928)


_ = tmp18929

tmp18930 := PrimCons(symsymbol, Nil)

tmp18931 := PrimCons(symlist, tmp18930)

tmp18932 := PrimCons(tmp18931, Nil)

tmp18933 := PrimCons(sym_1_1_6, tmp18932)

tmp18934 := PrimCons(symsymbol, tmp18933)

tmp18935 := Call(__e, PrimFunc(symdeclare), syminternal, tmp18934)


_ = tmp18935

tmp18936 := PrimCons(symA, Nil)

tmp18937 := PrimCons(symlist, tmp18936)

tmp18938 := PrimCons(symA, Nil)

tmp18939 := PrimCons(symlist, tmp18938)

tmp18940 := PrimCons(symA, Nil)

tmp18941 := PrimCons(symlist, tmp18940)

tmp18942 := PrimCons(tmp18941, Nil)

tmp18943 := PrimCons(sym_1_1_6, tmp18942)

tmp18944 := PrimCons(tmp18939, tmp18943)

tmp18945 := PrimCons(tmp18944, Nil)

tmp18946 := PrimCons(sym_1_1_6, tmp18945)

tmp18947 := PrimCons(tmp18937, tmp18946)

tmp18948 := Call(__e, PrimFunc(symdeclare), symintersection, tmp18947)


_ = tmp18948

tmp18949 := PrimCons(symstring, Nil)

tmp18950 := PrimCons(sym_1_1_6, tmp18949)

tmp18951 := Call(__e, PrimFunc(symdeclare), symlanguage, tmp18950)


_ = tmp18951

tmp18952 := PrimCons(symA, Nil)

tmp18953 := PrimCons(symlist, tmp18952)

tmp18954 := PrimCons(symnumber, Nil)

tmp18955 := PrimCons(sym_1_1_6, tmp18954)

tmp18956 := PrimCons(tmp18953, tmp18955)

tmp18957 := Call(__e, PrimFunc(symdeclare), symlength, tmp18956)


_ = tmp18957

tmp18958 := PrimCons(symA, Nil)

tmp18959 := PrimCons(symvector, tmp18958)

tmp18960 := PrimCons(symnumber, Nil)

tmp18961 := PrimCons(sym_1_1_6, tmp18960)

tmp18962 := PrimCons(tmp18959, tmp18961)

tmp18963 := Call(__e, PrimFunc(symdeclare), symlimit, tmp18962)


_ = tmp18963

tmp18964 := PrimCons(symin, Nil)

tmp18965 := PrimCons(symstream, tmp18964)

tmp18966 := PrimCons(symunit, Nil)

tmp18967 := PrimCons(symlist, tmp18966)

tmp18968 := PrimCons(tmp18967, Nil)

tmp18969 := PrimCons(sym_1_1_6, tmp18968)

tmp18970 := PrimCons(tmp18965, tmp18969)

tmp18971 := Call(__e, PrimFunc(symdeclare), symlineread, tmp18970)


_ = tmp18971

tmp18972 := PrimCons(symsymbol, Nil)

tmp18973 := PrimCons(sym_1_1_6, tmp18972)

tmp18974 := PrimCons(symstring, tmp18973)

tmp18975 := Call(__e, PrimFunc(symdeclare), symload, tmp18974)


_ = tmp18975

tmp18976 := PrimCons(symB, Nil)

tmp18977 := PrimCons(sym_1_1_6, tmp18976)

tmp18978 := PrimCons(symA, tmp18977)

tmp18979 := PrimCons(symA, Nil)

tmp18980 := PrimCons(symlist, tmp18979)

tmp18981 := PrimCons(symB, Nil)

tmp18982 := PrimCons(symlist, tmp18981)

tmp18983 := PrimCons(tmp18982, Nil)

tmp18984 := PrimCons(sym_1_1_6, tmp18983)

tmp18985 := PrimCons(tmp18980, tmp18984)

tmp18986 := PrimCons(tmp18985, Nil)

tmp18987 := PrimCons(sym_1_1_6, tmp18986)

tmp18988 := PrimCons(tmp18978, tmp18987)

tmp18989 := Call(__e, PrimFunc(symdeclare), symmap, tmp18988)


_ = tmp18989

tmp18990 := PrimCons(symB, Nil)

tmp18991 := PrimCons(symlist, tmp18990)

tmp18992 := PrimCons(tmp18991, Nil)

tmp18993 := PrimCons(sym_1_1_6, tmp18992)

tmp18994 := PrimCons(symA, tmp18993)

tmp18995 := PrimCons(symA, Nil)

tmp18996 := PrimCons(symlist, tmp18995)

tmp18997 := PrimCons(symB, Nil)

tmp18998 := PrimCons(symlist, tmp18997)

tmp18999 := PrimCons(tmp18998, Nil)

tmp19000 := PrimCons(sym_1_1_6, tmp18999)

tmp19001 := PrimCons(tmp18996, tmp19000)

tmp19002 := PrimCons(tmp19001, Nil)

tmp19003 := PrimCons(sym_1_1_6, tmp19002)

tmp19004 := PrimCons(tmp18994, tmp19003)

tmp19005 := Call(__e, PrimFunc(symdeclare), symmapcan, tmp19004)


_ = tmp19005

tmp19006 := PrimCons(symnumber, Nil)

tmp19007 := PrimCons(sym_1_1_6, tmp19006)

tmp19008 := PrimCons(symnumber, tmp19007)

tmp19009 := Call(__e, PrimFunc(symdeclare), symmaxinferences, tmp19008)


_ = tmp19009

tmp19010 := PrimCons(symstring, Nil)

tmp19011 := PrimCons(sym_1_1_6, tmp19010)

tmp19012 := PrimCons(symnumber, tmp19011)

tmp19013 := Call(__e, PrimFunc(symdeclare), symn_1_6string, tmp19012)


_ = tmp19013

tmp19014 := PrimCons(symnumber, Nil)

tmp19015 := PrimCons(sym_1_1_6, tmp19014)

tmp19016 := PrimCons(symnumber, tmp19015)

tmp19017 := Call(__e, PrimFunc(symdeclare), symnl, tmp19016)


_ = tmp19017

tmp19018 := PrimCons(symboolean, Nil)

tmp19019 := PrimCons(sym_1_1_6, tmp19018)

tmp19020 := PrimCons(symboolean, tmp19019)

tmp19021 := Call(__e, PrimFunc(symdeclare), symnot, tmp19020)


_ = tmp19021

tmp19022 := PrimCons(symA, Nil)

tmp19023 := PrimCons(symlist, tmp19022)

tmp19024 := PrimCons(symA, Nil)

tmp19025 := PrimCons(sym_1_1_6, tmp19024)

tmp19026 := PrimCons(tmp19023, tmp19025)

tmp19027 := PrimCons(tmp19026, Nil)

tmp19028 := PrimCons(sym_1_1_6, tmp19027)

tmp19029 := PrimCons(symnumber, tmp19028)

tmp19030 := Call(__e, PrimFunc(symdeclare), symnth, tmp19029)


_ = tmp19030

tmp19031 := PrimCons(symboolean, Nil)

tmp19032 := PrimCons(sym_1_1_6, tmp19031)

tmp19033 := PrimCons(symA, tmp19032)

tmp19034 := Call(__e, PrimFunc(symdeclare), symnumber_2, tmp19033)


_ = tmp19034

tmp19035 := PrimCons(symnumber, Nil)

tmp19036 := PrimCons(sym_1_1_6, tmp19035)

tmp19037 := PrimCons(symB, tmp19036)

tmp19038 := PrimCons(tmp19037, Nil)

tmp19039 := PrimCons(sym_1_1_6, tmp19038)

tmp19040 := PrimCons(symA, tmp19039)

tmp19041 := Call(__e, PrimFunc(symdeclare), symoccurrences, tmp19040)


_ = tmp19041

tmp19042 := PrimCons(symboolean, Nil)

tmp19043 := PrimCons(sym_1_1_6, tmp19042)

tmp19044 := PrimCons(symsymbol, tmp19043)

tmp19045 := Call(__e, PrimFunc(symdeclare), symoccurs_1check, tmp19044)


_ = tmp19045

tmp19046 := PrimCons(symboolean, Nil)

tmp19047 := PrimCons(sym_1_1_6, tmp19046)

tmp19048 := Call(__e, PrimFunc(symdeclare), symoccurs_2, tmp19047)


_ = tmp19048

tmp19049 := PrimCons(symboolean, Nil)

tmp19050 := PrimCons(sym_1_1_6, tmp19049)

tmp19051 := PrimCons(symsymbol, tmp19050)

tmp19052 := Call(__e, PrimFunc(symdeclare), symoptimise, tmp19051)


_ = tmp19052

tmp19053 := PrimCons(symboolean, Nil)

tmp19054 := PrimCons(sym_1_1_6, tmp19053)

tmp19055 := Call(__e, PrimFunc(symdeclare), symoptimise_2, tmp19054)


_ = tmp19055

tmp19056 := PrimCons(symboolean, Nil)

tmp19057 := PrimCons(sym_1_1_6, tmp19056)

tmp19058 := PrimCons(symboolean, tmp19057)

tmp19059 := PrimCons(tmp19058, Nil)

tmp19060 := PrimCons(sym_1_1_6, tmp19059)

tmp19061 := PrimCons(symboolean, tmp19060)

tmp19062 := Call(__e, PrimFunc(symdeclare), symor, tmp19061)


_ = tmp19062

tmp19063 := PrimCons(symstring, Nil)

tmp19064 := PrimCons(sym_1_1_6, tmp19063)

tmp19065 := Call(__e, PrimFunc(symdeclare), symos, tmp19064)


_ = tmp19065

tmp19066 := PrimCons(symboolean, Nil)

tmp19067 := PrimCons(sym_1_1_6, tmp19066)

tmp19068 := PrimCons(symsymbol, tmp19067)

tmp19069 := Call(__e, PrimFunc(symdeclare), sympackage_2, tmp19068)


_ = tmp19069

tmp19070 := PrimCons(symstring, Nil)

tmp19071 := PrimCons(sym_1_1_6, tmp19070)

tmp19072 := Call(__e, PrimFunc(symdeclare), symport, tmp19071)


_ = tmp19072

tmp19073 := PrimCons(symstring, Nil)

tmp19074 := PrimCons(sym_1_1_6, tmp19073)

tmp19075 := Call(__e, PrimFunc(symdeclare), symporters, tmp19074)


_ = tmp19075

tmp19076 := PrimCons(symstring, Nil)

tmp19077 := PrimCons(sym_1_1_6, tmp19076)

tmp19078 := PrimCons(symnumber, tmp19077)

tmp19079 := PrimCons(tmp19078, Nil)

tmp19080 := PrimCons(sym_1_1_6, tmp19079)

tmp19081 := PrimCons(symstring, tmp19080)

tmp19082 := Call(__e, PrimFunc(symdeclare), sympos, tmp19081)


_ = tmp19082

tmp19083 := PrimCons(symout, Nil)

tmp19084 := PrimCons(symstream, tmp19083)

tmp19085 := PrimCons(symstring, Nil)

tmp19086 := PrimCons(sym_1_1_6, tmp19085)

tmp19087 := PrimCons(tmp19084, tmp19086)

tmp19088 := PrimCons(tmp19087, Nil)

tmp19089 := PrimCons(sym_1_1_6, tmp19088)

tmp19090 := PrimCons(symstring, tmp19089)

tmp19091 := Call(__e, PrimFunc(symdeclare), sympr, tmp19090)


_ = tmp19091

tmp19092 := PrimCons(symA, Nil)

tmp19093 := PrimCons(sym_1_1_6, tmp19092)

tmp19094 := PrimCons(symA, tmp19093)

tmp19095 := Call(__e, PrimFunc(symdeclare), symprint, tmp19094)


_ = tmp19095

tmp19096 := PrimCons(symsymbol, Nil)

tmp19097 := PrimCons(sym_1_1_6, tmp19096)

tmp19098 := PrimCons(symsymbol, tmp19097)

tmp19099 := Call(__e, PrimFunc(symdeclare), symprofile, tmp19098)


_ = tmp19099

tmp19100 := PrimCons(symsymbol, Nil)

tmp19101 := PrimCons(symlist, tmp19100)

tmp19102 := PrimCons(symsymbol, Nil)

tmp19103 := PrimCons(symlist, tmp19102)

tmp19104 := PrimCons(tmp19103, Nil)

tmp19105 := PrimCons(sym_1_1_6, tmp19104)

tmp19106 := PrimCons(tmp19101, tmp19105)

tmp19107 := Call(__e, PrimFunc(symdeclare), sympreclude, tmp19106)


_ = tmp19107

tmp19108 := PrimCons(symstring, Nil)

tmp19109 := PrimCons(sym_1_1_6, tmp19108)

tmp19110 := PrimCons(symstring, tmp19109)

tmp19111 := Call(__e, PrimFunc(symdeclare), symshen_4proc_1nl, tmp19110)


_ = tmp19111

tmp19112 := PrimCons(symnumber, Nil)

tmp19113 := PrimCons(sym_d, tmp19112)

tmp19114 := PrimCons(symsymbol, tmp19113)

tmp19115 := PrimCons(tmp19114, Nil)

tmp19116 := PrimCons(sym_1_1_6, tmp19115)

tmp19117 := PrimCons(symsymbol, tmp19116)

tmp19118 := Call(__e, PrimFunc(symdeclare), symprofile_1results, tmp19117)


_ = tmp19118

tmp19119 := PrimCons(symA, Nil)

tmp19120 := PrimCons(sym_1_1_6, tmp19119)

tmp19121 := PrimCons(symA, tmp19120)

tmp19122 := Call(__e, PrimFunc(symdeclare), symprotect, tmp19121)


_ = tmp19122

tmp19123 := PrimCons(symsymbol, Nil)

tmp19124 := PrimCons(symlist, tmp19123)

tmp19125 := PrimCons(symsymbol, Nil)

tmp19126 := PrimCons(symlist, tmp19125)

tmp19127 := PrimCons(tmp19126, Nil)

tmp19128 := PrimCons(sym_1_1_6, tmp19127)

tmp19129 := PrimCons(tmp19124, tmp19128)

tmp19130 := Call(__e, PrimFunc(symdeclare), sympreclude_1all_1but, tmp19129)


_ = tmp19130

tmp19131 := PrimCons(symout, Nil)

tmp19132 := PrimCons(symstream, tmp19131)

tmp19133 := PrimCons(symstring, Nil)

tmp19134 := PrimCons(sym_1_1_6, tmp19133)

tmp19135 := PrimCons(tmp19132, tmp19134)

tmp19136 := PrimCons(tmp19135, Nil)

tmp19137 := PrimCons(sym_1_1_6, tmp19136)

tmp19138 := PrimCons(symstring, tmp19137)

tmp19139 := Call(__e, PrimFunc(symdeclare), symshen_4prhush, tmp19138)


_ = tmp19139

tmp19140 := PrimCons(symnumber, Nil)

tmp19141 := PrimCons(sym_1_1_6, tmp19140)

tmp19142 := PrimCons(symnumber, tmp19141)

tmp19143 := Call(__e, PrimFunc(symdeclare), symprolog_1memory, tmp19142)


_ = tmp19143

tmp19144 := PrimCons(symunit, Nil)

tmp19145 := PrimCons(symlist, tmp19144)

tmp19146 := PrimCons(tmp19145, Nil)

tmp19147 := PrimCons(sym_1_1_6, tmp19146)

tmp19148 := PrimCons(symsymbol, tmp19147)

tmp19149 := Call(__e, PrimFunc(symdeclare), symps, tmp19148)


_ = tmp19149

tmp19150 := PrimCons(symin, Nil)

tmp19151 := PrimCons(symstream, tmp19150)

tmp19152 := PrimCons(symunit, Nil)

tmp19153 := PrimCons(sym_1_1_6, tmp19152)

tmp19154 := PrimCons(tmp19151, tmp19153)

tmp19155 := Call(__e, PrimFunc(symdeclare), symread, tmp19154)


_ = tmp19155

tmp19156 := PrimCons(symin, Nil)

tmp19157 := PrimCons(symstream, tmp19156)

tmp19158 := PrimCons(symnumber, Nil)

tmp19159 := PrimCons(sym_1_1_6, tmp19158)

tmp19160 := PrimCons(tmp19157, tmp19159)

tmp19161 := Call(__e, PrimFunc(symdeclare), symread_1byte, tmp19160)


_ = tmp19161

tmp19162 := PrimCons(symnumber, Nil)

tmp19163 := PrimCons(symlist, tmp19162)

tmp19164 := PrimCons(tmp19163, Nil)

tmp19165 := PrimCons(sym_1_1_6, tmp19164)

tmp19166 := PrimCons(symstring, tmp19165)

tmp19167 := Call(__e, PrimFunc(symdeclare), symread_1file_1as_1bytelist, tmp19166)


_ = tmp19167

tmp19168 := PrimCons(symstring, Nil)

tmp19169 := PrimCons(sym_1_1_6, tmp19168)

tmp19170 := PrimCons(symstring, tmp19169)

tmp19171 := Call(__e, PrimFunc(symdeclare), symread_1file_1as_1string, tmp19170)


_ = tmp19171

tmp19172 := PrimCons(symunit, Nil)

tmp19173 := PrimCons(symlist, tmp19172)

tmp19174 := PrimCons(tmp19173, Nil)

tmp19175 := PrimCons(sym_1_1_6, tmp19174)

tmp19176 := PrimCons(symstring, tmp19175)

tmp19177 := Call(__e, PrimFunc(symdeclare), symread_1file, tmp19176)


_ = tmp19177

tmp19178 := PrimCons(symunit, Nil)

tmp19179 := PrimCons(symlist, tmp19178)

tmp19180 := PrimCons(tmp19179, Nil)

tmp19181 := PrimCons(sym_1_1_6, tmp19180)

tmp19182 := PrimCons(symstring, tmp19181)

tmp19183 := Call(__e, PrimFunc(symdeclare), symread_1from_1string, tmp19182)


_ = tmp19183

tmp19184 := PrimCons(symunit, Nil)

tmp19185 := PrimCons(symlist, tmp19184)

tmp19186 := PrimCons(tmp19185, Nil)

tmp19187 := PrimCons(sym_1_1_6, tmp19186)

tmp19188 := PrimCons(symstring, tmp19187)

tmp19189 := Call(__e, PrimFunc(symdeclare), symread_1from_1string_1unprocessed, tmp19188)


_ = tmp19189

tmp19190 := PrimCons(symstring, Nil)

tmp19191 := PrimCons(sym_1_1_6, tmp19190)

tmp19192 := Call(__e, PrimFunc(symdeclare), symrelease, tmp19191)


_ = tmp19192

tmp19193 := PrimCons(symA, Nil)

tmp19194 := PrimCons(symlist, tmp19193)

tmp19195 := PrimCons(symA, Nil)

tmp19196 := PrimCons(symlist, tmp19195)

tmp19197 := PrimCons(tmp19196, Nil)

tmp19198 := PrimCons(sym_1_1_6, tmp19197)

tmp19199 := PrimCons(tmp19194, tmp19198)

tmp19200 := PrimCons(tmp19199, Nil)

tmp19201 := PrimCons(sym_1_1_6, tmp19200)

tmp19202 := PrimCons(symA, tmp19201)

tmp19203 := Call(__e, PrimFunc(symdeclare), symremove, tmp19202)


_ = tmp19203

tmp19204 := PrimCons(symA, Nil)

tmp19205 := PrimCons(symlist, tmp19204)

tmp19206 := PrimCons(symA, Nil)

tmp19207 := PrimCons(symlist, tmp19206)

tmp19208 := PrimCons(tmp19207, Nil)

tmp19209 := PrimCons(sym_1_1_6, tmp19208)

tmp19210 := PrimCons(tmp19205, tmp19209)

tmp19211 := Call(__e, PrimFunc(symdeclare), symreverse, tmp19210)


_ = tmp19211

tmp19212 := PrimCons(symA, Nil)

tmp19213 := PrimCons(sym_1_1_6, tmp19212)

tmp19214 := PrimCons(symstring, tmp19213)

tmp19215 := Call(__e, PrimFunc(symdeclare), symsimple_1error, tmp19214)


_ = tmp19215

tmp19216 := PrimCons(symB, Nil)

tmp19217 := PrimCons(sym_d, tmp19216)

tmp19218 := PrimCons(symA, tmp19217)

tmp19219 := PrimCons(symB, Nil)

tmp19220 := PrimCons(sym_1_1_6, tmp19219)

tmp19221 := PrimCons(tmp19218, tmp19220)

tmp19222 := Call(__e, PrimFunc(symdeclare), symsnd, tmp19221)


_ = tmp19222

tmp19223 := PrimCons(symsymbol, Nil)

tmp19224 := PrimCons(sym_1_1_6, tmp19223)

tmp19225 := PrimCons(symnumber, tmp19224)

tmp19226 := PrimCons(tmp19225, Nil)

tmp19227 := PrimCons(sym_1_1_6, tmp19226)

tmp19228 := PrimCons(symsymbol, tmp19227)

tmp19229 := Call(__e, PrimFunc(symdeclare), symspecialise, tmp19228)


_ = tmp19229

tmp19230 := PrimCons(symboolean, Nil)

tmp19231 := PrimCons(sym_1_1_6, tmp19230)

tmp19232 := PrimCons(symsymbol, tmp19231)

tmp19233 := Call(__e, PrimFunc(symdeclare), symspy, tmp19232)


_ = tmp19233

tmp19234 := PrimCons(symboolean, Nil)

tmp19235 := PrimCons(sym_1_1_6, tmp19234)

tmp19236 := Call(__e, PrimFunc(symdeclare), symspy_2, tmp19235)


_ = tmp19236

tmp19237 := PrimCons(symboolean, Nil)

tmp19238 := PrimCons(sym_1_1_6, tmp19237)

tmp19239 := PrimCons(symsymbol, tmp19238)

tmp19240 := Call(__e, PrimFunc(symdeclare), symstep, tmp19239)


_ = tmp19240

tmp19241 := PrimCons(symboolean, Nil)

tmp19242 := PrimCons(sym_1_1_6, tmp19241)

tmp19243 := Call(__e, PrimFunc(symdeclare), symstep_2, tmp19242)


_ = tmp19243

tmp19244 := PrimCons(symin, Nil)

tmp19245 := PrimCons(symstream, tmp19244)

tmp19246 := PrimCons(tmp19245, Nil)

tmp19247 := PrimCons(sym_1_1_6, tmp19246)

tmp19248 := Call(__e, PrimFunc(symdeclare), symstinput, tmp19247)


_ = tmp19248

tmp19249 := PrimCons(symout, Nil)

tmp19250 := PrimCons(symstream, tmp19249)

tmp19251 := PrimCons(tmp19250, Nil)

tmp19252 := PrimCons(sym_1_1_6, tmp19251)

tmp19253 := Call(__e, PrimFunc(symdeclare), symstoutput, tmp19252)


_ = tmp19253

tmp19254 := PrimCons(symboolean, Nil)

tmp19255 := PrimCons(sym_1_1_6, tmp19254)

tmp19256 := PrimCons(symA, tmp19255)

tmp19257 := Call(__e, PrimFunc(symdeclare), symstring_2, tmp19256)


_ = tmp19257

tmp19258 := PrimCons(symstring, Nil)

tmp19259 := PrimCons(sym_1_1_6, tmp19258)

tmp19260 := PrimCons(symA, tmp19259)

tmp19261 := Call(__e, PrimFunc(symdeclare), symstr, tmp19260)


_ = tmp19261

tmp19262 := PrimCons(symnumber, Nil)

tmp19263 := PrimCons(sym_1_1_6, tmp19262)

tmp19264 := PrimCons(symstring, tmp19263)

tmp19265 := Call(__e, PrimFunc(symdeclare), symstring_1_6n, tmp19264)


_ = tmp19265

tmp19266 := PrimCons(symsymbol, Nil)

tmp19267 := PrimCons(sym_1_1_6, tmp19266)

tmp19268 := PrimCons(symstring, tmp19267)

tmp19269 := Call(__e, PrimFunc(symdeclare), symstring_1_6symbol, tmp19268)


_ = tmp19269

tmp19270 := PrimCons(symnumber, Nil)

tmp19271 := PrimCons(symlist, tmp19270)

tmp19272 := PrimCons(symnumber, Nil)

tmp19273 := PrimCons(sym_1_1_6, tmp19272)

tmp19274 := PrimCons(tmp19271, tmp19273)

tmp19275 := Call(__e, PrimFunc(symdeclare), symsum, tmp19274)


_ = tmp19275

tmp19276 := PrimCons(symboolean, Nil)

tmp19277 := PrimCons(sym_1_1_6, tmp19276)

tmp19278 := PrimCons(symA, tmp19277)

tmp19279 := Call(__e, PrimFunc(symdeclare), symsymbol_2, tmp19278)


_ = tmp19279

tmp19280 := PrimCons(symsymbol, Nil)

tmp19281 := PrimCons(sym_1_1_6, tmp19280)

tmp19282 := PrimCons(symsymbol, tmp19281)

tmp19283 := Call(__e, PrimFunc(symdeclare), symsystemf, tmp19282)


_ = tmp19283

tmp19284 := PrimCons(symboolean, Nil)

tmp19285 := PrimCons(sym_1_1_6, tmp19284)

tmp19286 := Call(__e, PrimFunc(symdeclare), symsystem_1S_2, tmp19285)


_ = tmp19286

tmp19287 := PrimCons(symA, Nil)

tmp19288 := PrimCons(symlist, tmp19287)

tmp19289 := PrimCons(symA, Nil)

tmp19290 := PrimCons(symlist, tmp19289)

tmp19291 := PrimCons(tmp19290, Nil)

tmp19292 := PrimCons(sym_1_1_6, tmp19291)

tmp19293 := PrimCons(tmp19288, tmp19292)

tmp19294 := Call(__e, PrimFunc(symdeclare), symtail, tmp19293)


_ = tmp19294

tmp19295 := PrimCons(symstring, Nil)

tmp19296 := PrimCons(sym_1_1_6, tmp19295)

tmp19297 := PrimCons(symstring, tmp19296)

tmp19298 := Call(__e, PrimFunc(symdeclare), symtlstr, tmp19297)


_ = tmp19298

tmp19299 := PrimCons(symA, Nil)

tmp19300 := PrimCons(symvector, tmp19299)

tmp19301 := PrimCons(symA, Nil)

tmp19302 := PrimCons(symvector, tmp19301)

tmp19303 := PrimCons(tmp19302, Nil)

tmp19304 := PrimCons(sym_1_1_6, tmp19303)

tmp19305 := PrimCons(tmp19300, tmp19304)

tmp19306 := Call(__e, PrimFunc(symdeclare), symtlv, tmp19305)


_ = tmp19306

tmp19307 := PrimCons(symboolean, Nil)

tmp19308 := PrimCons(sym_1_1_6, tmp19307)

tmp19309 := PrimCons(symsymbol, tmp19308)

tmp19310 := Call(__e, PrimFunc(symdeclare), symtc, tmp19309)


_ = tmp19310

tmp19311 := PrimCons(symboolean, Nil)

tmp19312 := PrimCons(sym_1_1_6, tmp19311)

tmp19313 := Call(__e, PrimFunc(symdeclare), symtc_2, tmp19312)


_ = tmp19313

tmp19314 := PrimCons(symA, Nil)

tmp19315 := PrimCons(symlazy, tmp19314)

tmp19316 := PrimCons(symA, Nil)

tmp19317 := PrimCons(sym_1_1_6, tmp19316)

tmp19318 := PrimCons(tmp19315, tmp19317)

tmp19319 := Call(__e, PrimFunc(symdeclare), symthaw, tmp19318)


_ = tmp19319

tmp19320 := PrimCons(symsymbol, Nil)

tmp19321 := PrimCons(sym_1_1_6, tmp19320)

tmp19322 := PrimCons(symsymbol, tmp19321)

tmp19323 := Call(__e, PrimFunc(symdeclare), symtrack, tmp19322)


_ = tmp19323

tmp19324 := PrimCons(symsymbol, Nil)

tmp19325 := PrimCons(symlist, tmp19324)

tmp19326 := PrimCons(tmp19325, Nil)

tmp19327 := PrimCons(sym_1_1_6, tmp19326)

tmp19328 := Call(__e, PrimFunc(symdeclare), symtracked, tmp19327)


_ = tmp19328

tmp19329 := PrimCons(symA, Nil)

tmp19330 := PrimCons(sym_1_1_6, tmp19329)

tmp19331 := PrimCons(symexception, tmp19330)

tmp19332 := PrimCons(symA, Nil)

tmp19333 := PrimCons(sym_1_1_6, tmp19332)

tmp19334 := PrimCons(tmp19331, tmp19333)

tmp19335 := PrimCons(tmp19334, Nil)

tmp19336 := PrimCons(sym_1_1_6, tmp19335)

tmp19337 := PrimCons(symA, tmp19336)

tmp19338 := Call(__e, PrimFunc(symdeclare), symtrap_1error, tmp19337)


_ = tmp19338

tmp19339 := PrimCons(symboolean, Nil)

tmp19340 := PrimCons(sym_1_1_6, tmp19339)

tmp19341 := PrimCons(symA, tmp19340)

tmp19342 := Call(__e, PrimFunc(symdeclare), symtuple_2, tmp19341)


_ = tmp19342

tmp19343 := PrimCons(symsymbol, Nil)

tmp19344 := PrimCons(sym_1_1_6, tmp19343)

tmp19345 := PrimCons(symsymbol, tmp19344)

tmp19346 := Call(__e, PrimFunc(symdeclare), symundefmacro, tmp19345)


_ = tmp19346

tmp19347 := PrimCons(symA, Nil)

tmp19348 := PrimCons(symlist, tmp19347)

tmp19349 := PrimCons(symA, Nil)

tmp19350 := PrimCons(symlist, tmp19349)

tmp19351 := PrimCons(symA, Nil)

tmp19352 := PrimCons(symlist, tmp19351)

tmp19353 := PrimCons(tmp19352, Nil)

tmp19354 := PrimCons(sym_1_1_6, tmp19353)

tmp19355 := PrimCons(tmp19350, tmp19354)

tmp19356 := PrimCons(tmp19355, Nil)

tmp19357 := PrimCons(sym_1_1_6, tmp19356)

tmp19358 := PrimCons(tmp19348, tmp19357)

tmp19359 := Call(__e, PrimFunc(symdeclare), symunion, tmp19358)


_ = tmp19359

tmp19360 := PrimCons(symsymbol, Nil)

tmp19361 := PrimCons(sym_1_1_6, tmp19360)

tmp19362 := PrimCons(symsymbol, tmp19361)

tmp19363 := Call(__e, PrimFunc(symdeclare), symunprofile, tmp19362)


_ = tmp19363

tmp19364 := PrimCons(symsymbol, Nil)

tmp19365 := PrimCons(sym_1_1_6, tmp19364)

tmp19366 := PrimCons(symsymbol, tmp19365)

tmp19367 := Call(__e, PrimFunc(symdeclare), symuntrack, tmp19366)


_ = tmp19367

tmp19368 := PrimCons(symsymbol, Nil)

tmp19369 := PrimCons(symlist, tmp19368)

tmp19370 := PrimCons(tmp19369, Nil)

tmp19371 := PrimCons(sym_1_1_6, tmp19370)

tmp19372 := Call(__e, PrimFunc(symdeclare), symuserdefs, tmp19371)


_ = tmp19372

tmp19373 := PrimCons(symboolean, Nil)

tmp19374 := PrimCons(sym_1_1_6, tmp19373)

tmp19375 := PrimCons(symA, tmp19374)

tmp19376 := Call(__e, PrimFunc(symdeclare), symvariable_2, tmp19375)


_ = tmp19376

tmp19377 := PrimCons(symboolean, Nil)

tmp19378 := PrimCons(sym_1_1_6, tmp19377)

tmp19379 := PrimCons(symA, tmp19378)

tmp19380 := Call(__e, PrimFunc(symdeclare), symvector_2, tmp19379)


_ = tmp19380

tmp19381 := PrimCons(symstring, Nil)

tmp19382 := PrimCons(sym_1_1_6, tmp19381)

tmp19383 := Call(__e, PrimFunc(symdeclare), symversion, tmp19382)


_ = tmp19383

tmp19384 := PrimCons(symA, Nil)

tmp19385 := PrimCons(sym_1_1_6, tmp19384)

tmp19386 := PrimCons(symA, tmp19385)

tmp19387 := PrimCons(tmp19386, Nil)

tmp19388 := PrimCons(sym_1_1_6, tmp19387)

tmp19389 := PrimCons(symstring, tmp19388)

tmp19390 := Call(__e, PrimFunc(symdeclare), symwrite_1to_1file, tmp19389)


_ = tmp19390

tmp19391 := PrimCons(symout, Nil)

tmp19392 := PrimCons(symstream, tmp19391)

tmp19393 := PrimCons(symnumber, Nil)

tmp19394 := PrimCons(sym_1_1_6, tmp19393)

tmp19395 := PrimCons(tmp19392, tmp19394)

tmp19396 := PrimCons(tmp19395, Nil)

tmp19397 := PrimCons(sym_1_1_6, tmp19396)

tmp19398 := PrimCons(symnumber, tmp19397)

tmp19399 := Call(__e, PrimFunc(symdeclare), symwrite_1byte, tmp19398)


_ = tmp19399

tmp19400 := PrimCons(symboolean, Nil)

tmp19401 := PrimCons(sym_1_1_6, tmp19400)

tmp19402 := PrimCons(symstring, tmp19401)

tmp19403 := Call(__e, PrimFunc(symdeclare), symy_1or_1n_2, tmp19402)


_ = tmp19403

tmp19404 := PrimCons(symboolean, Nil)

tmp19405 := PrimCons(sym_1_1_6, tmp19404)

tmp19406 := PrimCons(symnumber, tmp19405)

tmp19407 := PrimCons(tmp19406, Nil)

tmp19408 := PrimCons(sym_1_1_6, tmp19407)

tmp19409 := PrimCons(symnumber, tmp19408)

tmp19410 := Call(__e, PrimFunc(symdeclare), sym_6, tmp19409)


_ = tmp19410

tmp19411 := PrimCons(symboolean, Nil)

tmp19412 := PrimCons(sym_1_1_6, tmp19411)

tmp19413 := PrimCons(symnumber, tmp19412)

tmp19414 := PrimCons(tmp19413, Nil)

tmp19415 := PrimCons(sym_1_1_6, tmp19414)

tmp19416 := PrimCons(symnumber, tmp19415)

tmp19417 := Call(__e, PrimFunc(symdeclare), sym_5, tmp19416)


_ = tmp19417

tmp19418 := PrimCons(symboolean, Nil)

tmp19419 := PrimCons(sym_1_1_6, tmp19418)

tmp19420 := PrimCons(symnumber, tmp19419)

tmp19421 := PrimCons(tmp19420, Nil)

tmp19422 := PrimCons(sym_1_1_6, tmp19421)

tmp19423 := PrimCons(symnumber, tmp19422)

tmp19424 := Call(__e, PrimFunc(symdeclare), sym_6_a, tmp19423)


_ = tmp19424

tmp19425 := PrimCons(symboolean, Nil)

tmp19426 := PrimCons(sym_1_1_6, tmp19425)

tmp19427 := PrimCons(symnumber, tmp19426)

tmp19428 := PrimCons(tmp19427, Nil)

tmp19429 := PrimCons(sym_1_1_6, tmp19428)

tmp19430 := PrimCons(symnumber, tmp19429)

tmp19431 := Call(__e, PrimFunc(symdeclare), sym_5_a, tmp19430)


_ = tmp19431

tmp19432 := PrimCons(symboolean, Nil)

tmp19433 := PrimCons(sym_1_1_6, tmp19432)

tmp19434 := PrimCons(symA, tmp19433)

tmp19435 := PrimCons(tmp19434, Nil)

tmp19436 := PrimCons(sym_1_1_6, tmp19435)

tmp19437 := PrimCons(symA, tmp19436)

tmp19438 := Call(__e, PrimFunc(symdeclare), sym_a, tmp19437)


_ = tmp19438

tmp19439 := PrimCons(symnumber, Nil)

tmp19440 := PrimCons(sym_1_1_6, tmp19439)

tmp19441 := PrimCons(symnumber, tmp19440)

tmp19442 := PrimCons(tmp19441, Nil)

tmp19443 := PrimCons(sym_1_1_6, tmp19442)

tmp19444 := PrimCons(symnumber, tmp19443)

tmp19445 := Call(__e, PrimFunc(symdeclare), sym_7, tmp19444)


_ = tmp19445

tmp19446 := PrimCons(symnumber, Nil)

tmp19447 := PrimCons(sym_1_1_6, tmp19446)

tmp19448 := PrimCons(symnumber, tmp19447)

tmp19449 := PrimCons(tmp19448, Nil)

tmp19450 := PrimCons(sym_1_1_6, tmp19449)

tmp19451 := PrimCons(symnumber, tmp19450)

tmp19452 := Call(__e, PrimFunc(symdeclare), sym_c, tmp19451)


_ = tmp19452

tmp19453 := PrimCons(symnumber, Nil)

tmp19454 := PrimCons(sym_1_1_6, tmp19453)

tmp19455 := PrimCons(symnumber, tmp19454)

tmp19456 := PrimCons(tmp19455, Nil)

tmp19457 := PrimCons(sym_1_1_6, tmp19456)

tmp19458 := PrimCons(symnumber, tmp19457)

tmp19459 := Call(__e, PrimFunc(symdeclare), sym_1, tmp19458)


_ = tmp19459

tmp19460 := PrimCons(symnumber, Nil)

tmp19461 := PrimCons(sym_1_1_6, tmp19460)

tmp19462 := PrimCons(symnumber, tmp19461)

tmp19463 := PrimCons(tmp19462, Nil)

tmp19464 := PrimCons(sym_1_1_6, tmp19463)

tmp19465 := PrimCons(symnumber, tmp19464)

tmp19466 := Call(__e, PrimFunc(symdeclare), sym_d, tmp19465)


_ = tmp19466

tmp19467 := PrimCons(symboolean, Nil)

tmp19468 := PrimCons(sym_1_1_6, tmp19467)

tmp19469 := PrimCons(symB, tmp19468)

tmp19470 := PrimCons(tmp19469, Nil)

tmp19471 := PrimCons(sym_1_1_6, tmp19470)

tmp19472 := PrimCons(symA, tmp19471)

__e.TailApply(PrimFunc(symdeclare), sym_a_a, tmp19472)
return




}, 0)

var symshen_4profile_1help = MakeSymbol("shen.profile-help")
var symshen_4bytes_1_6string = MakeSymbol("shen.bytes->string")
var symprolog_1memory = MakeSymbol("prolog-memory")
var symshen_4find_1arity = MakeSymbol("shen.find-arity")
var symshen_4evaluate_1lineread = MakeSymbol("shen.evaluate-lineread")
var symshen_4string_1prefix_2 = MakeSymbol("shen.string-prefix?")
var symnull = MakeSymbol("null")
var symshen_4_dfactorise_2_d = MakeSymbol("shen.*factorise?*")
var symnth = MakeSymbol("nth")
var symport = MakeSymbol("port")
var symlineread = MakeSymbol("lineread")
var symshen_4_5str_6 = MakeSymbol("shen.<str>")
var symshen_4_5log10_6 = MakeSymbol("shen.<log10>")
var symreceive = MakeSymbol("receive")
var symshen_4copyfromvector = MakeSymbol("shen.copyfromvector")
var symshen_4write_1string = MakeSymbol("shen.write-string")
var symshen_4my_1read_1byte = MakeSymbol("shen.my-read-byte")
var sym_c = MakeSymbol("/")
var symloaded = MakeSymbol("loaded")
var symshen_4use_1history = MakeSymbol("shen.use-history")
var symshen_4process_1synonyms = MakeSymbol("shen.process-synonyms")
var symK = MakeSymbol("K")
var symsubst = MakeSymbol("subst")
var symshen_4_dsigf_d = MakeSymbol("shen.*sigf*")
var symFreeze = MakeSymbol("Freeze")
var symZ = MakeSymbol("Z")
var symshen_4_5number_6 = MakeSymbol("shen.<number>")
var symshen_4fn_1call = MakeSymbol("shen.fn-call")
var symshen_4_1m = MakeSymbol("shen.-m")
var symshen_4make_1prolog_1variable = MakeSymbol("shen.make-prolog-variable")
var symeval_1kl = MakeSymbol("eval-kl")
var symshen_4external_1symbols = MakeSymbol("shen.external-symbols")
var symconcat = MakeSymbol("concat")
var symshen_4_5datatype_1rule_6 = MakeSymbol("shen.<datatype-rule>")
var symshen_4t_d_1rules = MakeSymbol("shen.t*-rules")
var symshen_4correct = MakeSymbol("shen.correct")
var symversion = MakeSymbol("version")
var symshen_4_5pattern2_6 = MakeSymbol("shen.<pattern2>")
var symshen_4pivot_1on = MakeSymbol("shen.pivot-on")
var symshen_4add_1sexpr = MakeSymbol("shen.add-sexpr")
var symshen_4check_1eval_1and_1print = MakeSymbol("shen.check-eval-and-print")
var symshen_4write_1kl = MakeSymbol("shen.write-kl")
var symshen_4top = MakeSymbol("shen.top")
var symshen_4_5sng_6 = MakeSymbol("shen.<sng>")
var symshen_4arg_1_6str = MakeSymbol("shen.arg->str")
var symshen_4_5bar_6 = MakeSymbol("shen.<bar>")
var symshen_4process_1read_1byte = MakeSymbol("shen.process-read-byte")
var symshen_4_5syntax_6 = MakeSymbol("shen.<syntax>")
var symshen_4combine_1c_1code = MakeSymbol("shen.combine-c-code")
var symshen_4yacc_1syntax = MakeSymbol("shen.yacc-syntax")
var symempty_2 = MakeSymbol("empty?")
var symprint = MakeSymbol("print")
var symshen_4compute_1E = MakeSymbol("shen.compute-E")
var symshen_4_7m = MakeSymbol("shen.+m")
var symshen_4passive = MakeSymbol("shen.passive")
var symHd = MakeSymbol("Hd")
var symshen_4variants_2 = MakeSymbol("shen.variants?")
var symcn = MakeSymbol("cn")
var symshen_4_5body_6 = MakeSymbol("shen.<body>")
var symshen_4lazyderef = MakeSymbol("shen.lazyderef")
var symshen_4deref_1forked_1literals = MakeSymbol("shen.deref-forked-literals")
var symshen_4search_1user_1datatypes = MakeSymbol("shen.search-user-datatypes")
var sym_dstinput_d = MakeSymbol("*stinput*")
var symshen_4r = MakeSymbol("shen.r")
var sym_5e_6 = MakeSymbol("<e>")
var symshen_4choicepoint_2 = MakeSymbol("shen.choicepoint?")
var symshen_4type_1F = MakeSymbol("shen.type-F")
var symprotect = MakeSymbol("protect")
var symshen_4mkstr_1r = MakeSymbol("shen.mkstr-r")
var sym_a = MakeSymbol("=")
var symshen_4nothing_1doing_2 = MakeSymbol("shen.nothing-doing?")
var symshen_4update_1lambdatable = MakeSymbol("shen.update-lambdatable")
var symshen_4_5float_6 = MakeSymbol("shen.<float>")
var symshen_4atom_1case_1minus = MakeSymbol("shen.atom-case-minus")
var symshen_4freshen_1sig = MakeSymbol("shen.freshen-sig")
var sym_1 = MakeSymbol("-")
var symshen_4_5comment_6 = MakeSymbol("shen.<comment>")
var symshen_4_5sym_6 = MakeSymbol("shen.<sym>")
var symshen_4newname = MakeSymbol("shen.newname")
var symshen_4check_1byte = MakeSymbol("shen.check-byte")
var symshen_4mod = MakeSymbol("shen.mod")
var symshen_4mkstr = MakeSymbol("shen.mkstr")
var symshen_4fbound_2 = MakeSymbol("shen.fbound?")
var symSelect = MakeSymbol("Select")
var symabsvector_2 = MakeSymbol("absvector?")
var symunit = MakeSymbol("unit")
var symshen_4compile_1side_1conditions = MakeSymbol("shen.compile-side-conditions")
var symshen_4_5rule_d_6 = MakeSymbol("shen.<rule*>")
var symtc_2 = MakeSymbol("tc?")
var symshen_4analyse_1variable_2 = MakeSymbol("shen.analyse-variable?")
var symshen_4unpackage_emacroexpand = MakeSymbol("shen.unpackage&macroexpand")
var symshen_4_5control_6 = MakeSymbol("shen.<control>")
var symis = MakeSymbol("is")
var symshen_4dynamic_1default = MakeSymbol("shen.dynamic-default")
var symshen_4remember_1datatype = MakeSymbol("shen.remember-datatype")
var symshen_4optimise_1typing = MakeSymbol("shen.optimise-typing")
var symcd = MakeSymbol("cd")
var symcons = MakeSymbol("cons")
var symshen_4whitespace_2 = MakeSymbol("shen.whitespace?")
var symshen_4internal_1to_1shen_2 = MakeSymbol("shen.internal-to-shen?")
var symshen_4pui_1h = MakeSymbol("shen.pui-h")
var symshen_4compile_1search_1procedure = MakeSymbol("shen.compile-search-procedure")
var symshen_4_5yaccsig_6 = MakeSymbol("shen.<yaccsig>")
var symshen_4non_1terminal_2 = MakeSymbol("shen.non-terminal?")
var symdestroy = MakeSymbol("destroy")
var symResult = MakeSymbol("Result")
var symstream = MakeSymbol("stream")
var symshen_4stpart = MakeSymbol("shen.stpart")
var symshen_4sigxrules = MakeSymbol("shen.sigxrules")
var symS = MakeSymbol("S")
var symtlv = MakeSymbol("tlv")
var symshen_4abs = MakeSymbol("shen.abs")
var symshen_4_doptimise_d = MakeSymbol("shen.*optimise*")
var symshen_4kl_1body = MakeSymbol("shen.kl-body")
var symfactorise = MakeSymbol("factorise")
var symshen_4_5returns_6 = MakeSymbol("shen.<returns>")
var symshen_4vector_1dereference = MakeSymbol("shen.vector-dereference")
var symshen_4f = MakeSymbol("shen.f")
var symshen_4_5semicolon_6 = MakeSymbol("shen.<semicolon>")
var symsymbol = MakeSymbol("symbol")
var symshen_4prolog_1arity_1check = MakeSymbol("shen.prolog-arity-check")
var symshen_4by_1hypothesis = MakeSymbol("shen.by-hypothesis")
var symshen_4_5name_6 = MakeSymbol("shen.<name>")
var symshen_4string_1match = MakeSymbol("shen.string-match")
var symshen_4t_d_1rule_1h = MakeSymbol("shen.t*-rule-h")
var symshen_4t_d_1integrity = MakeSymbol("shen.t*-integrity")
var symshen_4_5colon_1equal_6 = MakeSymbol("shen.<colon-equal>")
var symshen_4a = MakeSymbol("shen.a")
var symshen_4skip = MakeSymbol("shen.skip")
var sym_e_e = MakeSymbol("&&")
var sym_b = MakeSymbol("!")
var symshen_4macroexpand_1h = MakeSymbol("shen.macroexpand-h")
var symeval = MakeSymbol("eval")
var symstinput = MakeSymbol("stinput")
var symos = MakeSymbol("os")
var symshen_4_5pattern_6 = MakeSymbol("shen.<pattern>")
var symshen_4constructor_2 = MakeSymbol("shen.constructor?")
var symshen_4alpha_1convert = MakeSymbol("shen.alpha-convert")
var symshen_4_5whitespaces_6 = MakeSymbol("shen.<whitespaces>")
var symtrack = MakeSymbol("track")
var symshen_4typecheck = MakeSymbol("shen.typecheck")
var symshen_4_5lowC_6 = MakeSymbol("shen.<lowC>")
var sym_7 = MakeSymbol("+")
var symin_1package = MakeSymbol("in-package")
var symlazy = MakeSymbol("lazy")
var symshen_4eval_1and_1print = MakeSymbol("shen.eval-and-print")
var symshen_4yacc_1_6shen = MakeSymbol("shen.yacc->shen")
var symshen_4call_1prolog = MakeSymbol("shen.call-prolog")
var symshen_4integer_1test_2 = MakeSymbol("shen.integer-test?")
var symshen_4ccons_2 = MakeSymbol("shen.ccons?")
var symlet = MakeSymbol("let")
var symW = MakeSymbol("W")
var symshen_4_dresidue_d = MakeSymbol("shen.*residue*")
var syminclude_1all_1but = MakeSymbol("include-all-but")
var symerror = MakeSymbol("error")
var symshen_4_5hterm_6 = MakeSymbol("shen.<hterm>")
var symshen_4unpackage = MakeSymbol("shen.unpackage")
var sym_dstoutput_d = MakeSymbol("*stoutput*")
var sympreclude_1all_1but = MakeSymbol("preclude-all-but")
var symshen_4klfile = MakeSymbol("shen.klfile")
var symshen_4rule_1_6clause = MakeSymbol("shen.rule->clause")
var symshen_4output_1track = MakeSymbol("shen.output-track")
var symshen_4t_d_1rule = MakeSymbol("shen.t*-rule")
var symshen_4objectcode = MakeSymbol("shen.objectcode")
var symshen_4_5strcontents_6 = MakeSymbol("shen.<strcontents>")
var symshen_4compute_1fraction = MakeSymbol("shen.compute-fraction")
var symforeign = MakeSymbol("foreign")
var symfix = MakeSymbol("fix")
var symY = MakeSymbol("Y")
var symshen_4printF = MakeSymbol("shen.printF")
var symshen_4free_1variable_1error_1message = MakeSymbol("shen.free-variable-error-message")
var symwrite_1byte = MakeSymbol("write-byte")
var symshen_4non_1application_2 = MakeSymbol("shen.non-application?")
var symmacroexpand = MakeSymbol("macroexpand")
var symshen_4_7vector_2 = MakeSymbol("shen.+vector?")
var symshen_4execute_1store_1arity = MakeSymbol("shen.execute-store-arity")
var symundefmacro = MakeSymbol("undefmacro")
var symcall = MakeSymbol("call")
var symdefcc = MakeSymbol("defcc")
var symshen_4write_1kl_1h = MakeSymbol("shen.write-kl-h")
var symshen_4lowercase_2 = MakeSymbol("shen.lowercase?")
var symshen_4f_1error = MakeSymbol("shen.f-error")
var symshen_4_5bterms_6 = MakeSymbol("shen.<bterms>")
var symshen_4head_1foundit_b = MakeSymbol("shen.head-foundit!")
var symshen_4use_1type_1info = MakeSymbol("shen.use-type-info")
var symout = MakeSymbol("out")
var symshen_4arity_1chk = MakeSymbol("shen.arity-chk")
var symshen_4compute_1fraction_1h = MakeSymbol("shen.compute-fraction-h")
var sym_dmacros_d = MakeSymbol("*macros*")
var symlist = MakeSymbol("list")
var symshen_4initialise__environment = MakeSymbol("shen.initialise_environment")
var symshen_4incinfs = MakeSymbol("shen.incinfs")
var symshen_4decrement_1ticket = MakeSymbol("shen.decrement-ticket")
var syminternal = MakeSymbol("internal")
var symshen_4s = MakeSymbol("shen.s")
var symshen_4_5pattern1_6 = MakeSymbol("shen.<pattern1>")
var symshen_4initialise_1lambda_1tables = MakeSymbol("shen.initialise-lambda-tables")
var symshen_4trim_1it = MakeSymbol("shen.trim-it")
var symshen_4hascut_2 = MakeSymbol("shen.hascut?")
var symshen_4deref = MakeSymbol("shen.deref")
var symshen_4lzy_a_b = MakeSymbol("shen.lzy=!")
var symstoutput = MakeSymbol("stoutput")
var symexplode = MakeSymbol("explode")
var sym_dlanguage_d = MakeSymbol("*language*")
var symclose = MakeSymbol("close")
var symshen_4predicate = MakeSymbol("shen.predicate")
var symshen_4_5sides_6 = MakeSymbol("shen.<sides>")
var symshen_4insert_1tracking_1code = MakeSymbol("shen.insert-tracking-code")
var symoutput = MakeSymbol("output")
var symshen_4process_1lambda = MakeSymbol("shen.process-lambda")
var symshen_4key_1in_1sequent_1calculus_2 = MakeSymbol("shen.key-in-sequent-calculus?")
var symshen_4keep_1looking = MakeSymbol("shen.keep-looking")
var symshen_4put_1profile = MakeSymbol("shen.put-profile")
var symshen_4map_1h = MakeSymbol("shen.map-h")
var symshen_4iter_1vector = MakeSymbol("shen.iter-vector")
var symshen_4rectify_1test = MakeSymbol("shen.rectify-test")
var symshen_4special_1case = MakeSymbol("shen.special-case")
var symor = MakeSymbol("or")
var symshen_4newpv = MakeSymbol("shen.newpv")
var symshen_4overbind = MakeSymbol("shen.overbind")
var symshen_4_5datatype_1rules_6 = MakeSymbol("shen.<datatype-rules>")
var symhdv = MakeSymbol("hdv")
var symshen_4hush = MakeSymbol("shen.hush")
var sym__ = MakeSymbol("_")
var symfunction = MakeSymbol("function")
var symload = MakeSymbol("load")
var sym_c_4 = MakeSymbol("/.")
var symshen_4type_1error = MakeSymbol("shen.type-error")
var symhead = MakeSymbol("head")
var symshen_4string_1_6bytes = MakeSymbol("shen.string->bytes")
var symstring_1_6symbol = MakeSymbol("string->symbol")
var symshen_4list_1_6str = MakeSymbol("shen.list->str")
var symshen_4vector_1_6str = MakeSymbol("shen.vector->str")
var symshen_4op1 = MakeSymbol("shen.op1")
var symshen_4_dit_d = MakeSymbol("shen.*it*")
var symshen_4_5alphanum_6 = MakeSymbol("shen.<alphanum>")
var sym_5_1vector = MakeSymbol("<-vector")
var symlanguage = MakeSymbol("language")
var sym_3 = MakeSymbol("$")
var sym_6 = MakeSymbol(">")
var symrun = MakeSymbol("run")
var symshen_4intern_1type = MakeSymbol("shen.intern-type")
var symFinish = MakeSymbol("Finish")
var symshen_4_5prems_6 = MakeSymbol("shen.<prems>")
var symshen_4fillvector = MakeSymbol("shen.fillvector")
var symmap = MakeSymbol("map")
var symshen_4freshterm = MakeSymbol("shen.freshterm")
var symremove = MakeSymbol("remove")
var symshen_4factor_1selectors = MakeSymbol("shen.factor-selectors")
var sympackage = MakeSymbol("package")
var symoccurs_1check = MakeSymbol("occurs-check")
var symshen_4_5datatype_6 = MakeSymbol("shen.<datatype>")
var symshen_4_5alpha_6 = MakeSymbol("shen.<alpha>")
var symshen_4_5hash_6 = MakeSymbol("shen.<hash>")
var symshen_4credits = MakeSymbol("shen.credits")
var symshen_4make_1uppercase = MakeSymbol("shen.make-uppercase")
var symshen_4_5sc_6 = MakeSymbol("shen.<sc>")
var symshen_4cons_1case_1plus = MakeSymbol("shen.cons-case-plus")
var symshen_4coll_1formulae = MakeSymbol("shen.coll-formulae")
var symshen_4line = MakeSymbol("shen.line")
var symintersection = MakeSymbol("intersection")
var symshen_4linearise = MakeSymbol("shen.linearise")
var symabsvector = MakeSymbol("absvector")
var symshen_4compile_1synonyms = MakeSymbol("shen.compile-synonyms")
var symshen_4deref_1terms = MakeSymbol("shen.deref-terms")
var symshen_4terpri_1or_1read_1char = MakeSymbol("shen.terpri-or-read-char")
var symshen_4p_1hyps = MakeSymbol("shen.p-hyps")
var symshen_4_5c_1rules_6 = MakeSymbol("shen.<c-rules>")
var symshen_4unassoc = MakeSymbol("shen.unassoc")
var symshen_4_5rules_6 = MakeSymbol("shen.<rules>")
var symshen_4_5patterns_6 = MakeSymbol("shen.<patterns>")
var sym_6_a = MakeSymbol(">=")
var symprolog_2 = MakeSymbol("prolog?")
var symunix = MakeSymbol("unix")
var symshen_4build_1lambda_1table = MakeSymbol("shen.build-lambda-table")
var symshen_4process_1time = MakeSymbol("shen.process-time")
var symoccurs_2 = MakeSymbol("occurs?")
var symshen_4record_1and_1evaluate = MakeSymbol("shen.record-and-evaluate")
var symif = MakeSymbol("if")
var symshen_4_ddatatypes_d = MakeSymbol("shen.*datatypes*")
var symshen_4_dstep_d = MakeSymbol("shen.*step*")
var symshen_4_5head_6 = MakeSymbol("shen.<head>")
var symshen_4spaces = MakeSymbol("shen.spaces")
var symshen_4autocomplete = MakeSymbol("shen.autocomplete")
var symput = MakeSymbol("put")
var symshen_4_5rrb_6 = MakeSymbol("shen.<rrb>")
var sympos = MakeSymbol("pos")
var symshen_4goto_1h = MakeSymbol("shen.goto-h")
var symshen_4show_1p = MakeSymbol("shen.show-p")
var symAction = MakeSymbol("Action")
var symshen_4_7string_2 = MakeSymbol("shen.+string?")
var symoptimise = MakeSymbol("optimise")
var symshen_4compile_1to_1kl = MakeSymbol("shen.compile-to-kl")
var symshen_4factor = MakeSymbol("shen.factor")
var symshen_4_5ass_6 = MakeSymbol("shen.<ass>")
var symshen_4dbl_1h_2 = MakeSymbol("shen.dbl-h?")
var symshen_4yacc_1semantics = MakeSymbol("shen.yacc-semantics")
var syminferences = MakeSymbol("inferences")
var sym_drelease_d = MakeSymbol("*release*")
var symshen_4_dlambdatable_d = MakeSymbol("shen.*lambdatable*")
var sym_5end_6 = MakeSymbol("<end>")
var symshen_4update_1assoc = MakeSymbol("shen.update-assoc")
var symshen_4terms = MakeSymbol("shen.terms")
var symshen_4_5single_6 = MakeSymbol("shen.<single>")
var symshen_4track_1function = MakeSymbol("shen.track-function")
var sym_j = MakeSymbol("}")
var symmaxinferences = MakeSymbol("maxinferences")
var symfindall = MakeSymbol("findall")
var symshen_4_5iscomma_6 = MakeSymbol("shen.<iscomma>")
var symappend = MakeSymbol("append")
var symdo = MakeSymbol("do")
var symV = MakeSymbol("V")
var symshen_4assoc_1_6 = MakeSymbol("shen.assoc->")
var symMessage = MakeSymbol("Message")
var symnl = MakeSymbol("nl")
var symshen_4proc_1nl = MakeSymbol("shen.proc-nl")
var symshen_4linearise_1h = MakeSymbol("shen.linearise-h")
var symshen_4_5lowE_6 = MakeSymbol("shen.<lowE>")
var symshen_4_5non_1terminal_2_6 = MakeSymbol("shen.<non-terminal?>")
var symshen_4gc = MakeSymbol("shen.gc")
var symshen_4profiled_2 = MakeSymbol("shen.profiled?")
var symRecord = MakeSymbol("Record")
var symvector = MakeSymbol("vector")
var symtail = MakeSymbol("tail")
var symshen_4special_2 = MakeSymbol("shen.special?")
var symshen_4process_1let = MakeSymbol("shen.process-let")
var symshen_4prolog_1vector_1size = MakeSymbol("shen.prolog-vector-size")
var symshen_4remove_1h = MakeSymbol("shen.remove-h")
var symcompile = MakeSymbol("compile")
var symshen_4_5alphanums_6 = MakeSymbol("shen.<alphanums>")
var sym_a_a_6 = MakeSymbol("==>")
var symStart = MakeSymbol("Start")
var symshen_4_1null_1 = MakeSymbol("shen.-null-")
var symvariable_2 = MakeSymbol("variable?")
var symshen_4modh = MakeSymbol("shen.modh")
var symbound_2 = MakeSymbol("bound?")
var symshen_4print_1vector_2 = MakeSymbol("shen.print-vector?")
var sym_5_b_6 = MakeSymbol("<!>")
var symprofile_1results = MakeSymbol("profile-results")
var symshen_4prolog_1parameters = MakeSymbol("shen.prolog-parameters")
var symshen_4passive_2 = MakeSymbol("shen.passive?")
var symshen_4remove_1pointer = MakeSymbol("shen.remove-pointer")
var symoccurrences = MakeSymbol("occurrences")
var symshen_4store_1arity = MakeSymbol("shen.store-arity")
var symnumber_2 = MakeSymbol("number?")
var symhush = MakeSymbol("hush")
var sym_5_1_1 = MakeSymbol("<--")
var symshen_4compile_1prolog = MakeSymbol("shen.compile-prolog")
var symshen_4profile_1func = MakeSymbol("shen.profile-func")
var symtracked = MakeSymbol("tracked")
var symshen_4monotype = MakeSymbol("shen.monotype")
var symshen_4_5shortnatter_6 = MakeSymbol("shen.<shortnatter>")
var symshen_4application_2 = MakeSymbol("shen.application?")
var symexception = MakeSymbol("exception")
var symshen_4goto = MakeSymbol("shen.goto")
var symshen_4demodulate = MakeSymbol("shen.demodulate")
var symshen_4non_1terminalcode = MakeSymbol("shen.non-terminalcode")
var symreverse = MakeSymbol("reverse")
var sym_dport_d = MakeSymbol("*port*")
var symshen_4mkstr_1l = MakeSymbol("shen.mkstr-l")
var symfn = MakeSymbol("fn")
var symshen_4beta = MakeSymbol("shen.beta")
var symshen_4op = MakeSymbol("shen.op")
var symshen_4simple_1curry = MakeSymbol("shen.simple-curry")
var sympreclude = MakeSymbol("preclude")
var symshen_4insert_1l = MakeSymbol("shen.insert-l")
var symshen_4print_1prolog_1vector = MakeSymbol("shen.print-prolog-vector")
var symerror_1to_1string = MakeSymbol("error-to-string")
var symshen_4mu_1h = MakeSymbol("shen.mu-h")
var symshen_4_5wildcard_6 = MakeSymbol("shen.<wildcard>")
var symshen_4pvar_2 = MakeSymbol("shen.pvar?")
var symshen_4_5dbl_6 = MakeSymbol("shen.<dbl>")
var symDelta = MakeSymbol("Delta")
var symshen_4shendef_1_6kldef = MakeSymbol("shen.shendef->kldef")
var symshen = MakeSymbol("shen")
var symshen_4package_1user_1input = MakeSymbol("shen.package-user-input")
var symshen_4maxinfexceeded_2 = MakeSymbol("shen.maxinfexceeded?")
var symParse = MakeSymbol("Parse")
var symps = MakeSymbol("ps")
var symshen_4hds_a_2 = MakeSymbol("shen.hds=?")
var syminput = MakeSymbol("input")
var sym_6_6 = MakeSymbol(">>")
var symshen_4_5formula_6 = MakeSymbol("shen.<formula>")
var symshen_4foundit_b = MakeSymbol("shen.foundit!")
var symporters = MakeSymbol("porters")
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var symis_b = MakeSymbol("is!")
var symshen_4partial = MakeSymbol("shen.partial")
var symshen_4create_1skeleton = MakeSymbol("shen.create-skeleton")
var symshen_4unlocked_2 = MakeSymbol("shen.unlocked?")
var symshen_4_5c_1rule_6 = MakeSymbol("shen.<c-rule>")
var symshen_4_dspecial_d = MakeSymbol("shen.*special*")
var symshen_4rcons__form = MakeSymbol("shen.rcons_form")
var symshen_4record_1internal = MakeSymbol("shen.record-internal")
var symshen_4_5iscolon_6 = MakeSymbol("shen.<iscolon>")
var symshen_4shendef_1_6kldef_1h = MakeSymbol("shen.shendef->kldef-h")
var symread_1from_1string = MakeSymbol("read-from-string")
var symshen_4_5longnatter_6 = MakeSymbol("shen.<longnatter>")
var symshen_4record_1external = MakeSymbol("shen.record-external")
var symdatatypes = MakeSymbol("datatypes")
var symshen_4pac_1h = MakeSymbol("shen.pac-h")
var symshen_4extract_1free_1vars = MakeSymbol("shen.extract-free-vars")
var symshen_4prterm = MakeSymbol("shen.prterm")
var symshen_4_dtracking_d = MakeSymbol("shen.*tracking*")
var symshen_4magless = MakeSymbol("shen.magless")
var symcons_2 = MakeSymbol("cons?")
var symread_1from_1string_1unprocessed = MakeSymbol("read-from-string-unprocessed")
var symshen_4_5atom_6 = MakeSymbol("shen.<atom>")
var symshen_4_dhistory_d = MakeSymbol("shen.*history*")
var symshen_4typename = MakeSymbol("shen.typename")
var symstring_2 = MakeSymbol("string?")
var symshen_4tlv_1help = MakeSymbol("shen.tlv-help")
var sympackage_2 = MakeSymbol("package?")
var symshen_4char_1stoutput_2 = MakeSymbol("shen.char-stoutput?")
var symshen_4_5stop_6 = MakeSymbol("shen.<stop>")
var symspy = MakeSymbol("spy")
var symverified = MakeSymbol("verified")
var symshen_4peek_1history = MakeSymbol("shen.peek-history")
var symshen_4process_1def = MakeSymbol("shen.process-def")
var symshen_4scan_1body = MakeSymbol("shen.scan-body")
var symshen_4process_1sexprs = MakeSymbol("shen.process-sexprs")
var symshen_4compute_1integer_1h = MakeSymbol("shen.compute-integer-h")
var symatom_2 = MakeSymbol("atom?")
var symshen_4compile_1premise = MakeSymbol("shen.compile-premise")
var symshen_4variancy = MakeSymbol("shen.variancy")
var symshen_4eos = MakeSymbol("shen.eos")
var symtl = MakeSymbol("tl")
var symshen_4macros = MakeSymbol("shen.macros")
var symshen_4_5bterm_6 = MakeSymbol("shen.<bterm>")
var symshen_4insert_1prolog_1variables = MakeSymbol("shen.insert-prolog-variables")
var symshen_4sigf = MakeSymbol("shen.sigf")
var symshen_4_5non_1terminal_1name_6 = MakeSymbol("shen.<non-terminal-name>")
var symshen_4_5packagenames_6 = MakeSymbol("shen.<packagenames>")
var symshen_4_8v_1help = MakeSymbol("shen.@v-help")
var sympr = MakeSymbol("pr")
var symshen_4_5equal_6 = MakeSymbol("shen.<equal>")
var symshen_4_5dbq_6 = MakeSymbol("shen.<dbq>")
var symshen_4bindv = MakeSymbol("shen.bindv")
var symshen_4_dpackage_d = MakeSymbol("shen.*package*")
var symshen_4rectify_1type = MakeSymbol("shen.rectify-type")
var symshen_4find_1arities = MakeSymbol("shen.find-arities")
var symshen_4process_1application = MakeSymbol("shen.process-application")
var symshen_4loading_2 = MakeSymbol("shen.loading?")
var symfile = MakeSymbol("file")
var symshen_4_5defprolog_6 = MakeSymbol("shen.<defprolog>")
var symshen_4prolog_1keyword_2 = MakeSymbol("shen.prolog-keyword?")
var symmapcan = MakeSymbol("mapcan")
var sym_dversion_d = MakeSymbol("*version*")
var symshen_4record_1it = MakeSymbol("shen.record-it")
var symshen_4assert_d = MakeSymbol("shen.assert*")
var symshen_4cons_1form_1no_1modes = MakeSymbol("shen.cons-form-no-modes")
var symshen_4l_1rules = MakeSymbol("shen.l-rules")
var sym_e = MakeSymbol("&")
var symshen_4insert = MakeSymbol("shen.insert")
var symwhere = MakeSymbol("where")
var symand = MakeSymbol("and")
var symshen_4process_1after_1type = MakeSymbol("shen.process-after-type")
var symtrap_1error = MakeSymbol("trap-error")
var symshen_4_5packagename_6 = MakeSymbol("shen.<packagename>")
var symKey = MakeSymbol("Key")
var symshen_4alphanums_2 = MakeSymbol("shen.alphanums?")
var symshen_4change_1pointer_1value = MakeSymbol("shen.change-pointer-value")
var symshen_4iter_1list = MakeSymbol("shen.iter-list")
var symshen_4parameters = MakeSymbol("shen.parameters")
var symshen_4i_1failed_b = MakeSymbol("shen.i-failed!")
var symshen_4compute_1integer = MakeSymbol("shen.compute-integer")
var symshen_4_dsystem_d = MakeSymbol("shen.*system*")
var symn_1_6string = MakeSymbol("n->string")
var symshen_4atom_1_6str = MakeSymbol("shen.atom->str")
var symshen_4lambda_1function = MakeSymbol("shen.lambda-function")
var symB = MakeSymbol("B")
var symshen_4compile_1side_1condition = MakeSymbol("shen.compile-side-condition")
var symshen_4_5sig_drules_6 = MakeSymbol("shen.<sig*rules>")
var symshen_4freshterm_2 = MakeSymbol("shen.freshterm?")
var symfail = MakeSymbol("fail")
var symuserdefs = MakeSymbol("userdefs")
var symshen_4insert_1info = MakeSymbol("shen.insert-info")
var symshen_4dynamic = MakeSymbol("shen.dynamic")
var symshen_4call_1dynamic = MakeSymbol("shen.call-dynamic")
var syminteger_2 = MakeSymbol("integer?")
var symshen_4_dmaxinferences_d = MakeSymbol("shen.*maxinferences*")
var symshen_4_dloading_2_d = MakeSymbol("shen.*loading?*")
var symshen_4unwind_1types = MakeSymbol("shen.unwind-types")
var symshen_4invoke = MakeSymbol("shen.invoke")
var symshen_4show_1datatypes = MakeSymbol("shen.show-datatypes")
var symshen_4tracked_2 = MakeSymbol("shen.tracked?")
var symshen_4explode_1h = MakeSymbol("shen.explode-h")
var symshen_4factorise_1code = MakeSymbol("shen.factorise-code")
var symshen_4constructor_1error = MakeSymbol("shen.constructor-error")
var symshen_4find_1free_1vars = MakeSymbol("shen.find-free-vars")
var symshen_4_5integer_6 = MakeSymbol("shen.<integer>")
var symshen_4_dalldatatypes_d = MakeSymbol("shen.*alldatatypes*")
var symshen_4_dcall_d = MakeSymbol("shen.*call*")
var symstep = MakeSymbol("step")
var symshen_4hashkey = MakeSymbol("shen.hashkey")
var symshen_4lambda_1entry = MakeSymbol("shen.lambda-entry")
var symshen_4rep_1X = MakeSymbol("shen.rep-X")
var symshen_4_5colon_6 = MakeSymbol("shen.<colon>")
var symshen_4remove_1datatypes = MakeSymbol("shen.remove-datatypes")
var symshen_4_5packagechar_6 = MakeSymbol("shen.<packagechar>")
var symRemainder = MakeSymbol("Remainder")
var symshen_4prhush = MakeSymbol("shen.prhush")
var symarity = MakeSymbol("arity")
var symshen_4string_1_6byte = MakeSymbol("shen.string->byte")
var symshen_4str_1_6bytes = MakeSymbol("shen.str->bytes")
var symaddress_1_6 = MakeSymbol("address->")
var symdefmacro = MakeSymbol("defmacro")
var symu_b = MakeSymbol("u!")
var symshen_4shen = MakeSymbol("shen.shen")
var symshen_4system_1S = MakeSymbol("shen.system-S")
var symshen_4posint_2 = MakeSymbol("shen.posint?")
var symshen_4remove_1indirection = MakeSymbol("shen.remove-indirection")
var symvalue = MakeSymbol("value")
var symdeclare = MakeSymbol("declare")
var symretract = MakeSymbol("retract")
var symshen_4continue = MakeSymbol("shen.continue")
var symshen_4ticket_1number = MakeSymbol("shen.ticket-number")
var symshen_4prolog_1track = MakeSymbol("shen.prolog-track")
var symopen = MakeSymbol("open")
var symasserta = MakeSymbol("asserta")
var symshen_4tuple_1up = MakeSymbol("shen.tuple-up")
var symshen_4assumetypes = MakeSymbol("shen.assumetypes")
var symshen_4_5return_6 = MakeSymbol("shen.<return>")
var symshen_4linearise_1clause = MakeSymbol("shen.linearise-clause")
var symshen_4wildcard_2 = MakeSymbol("shen.wildcard?")
var symshen_4_5yacc_6 = MakeSymbol("shen.<yacc>")
var symshen_4processed = MakeSymbol("shen.processed")
var symTl = MakeSymbol("Tl")
var symfactorise_2 = MakeSymbol("factorise?")
var symshen_4analyse_1symbol_2 = MakeSymbol("shen.analyse-symbol?")
var symshen_4prodbutzero = MakeSymbol("shen.prodbutzero")
var symshen_4initialise_1arity_1table = MakeSymbol("shen.initialise-arity-table")
var symreturn = MakeSymbol("return")
var symnumber = MakeSymbol("number")
var sym_1_1_6 = MakeSymbol("-->")
var symshen_4freeze_1literals = MakeSymbol("shen.freeze-literals")
var symthaw = MakeSymbol("thaw")
var symwrite_1to_1file = MakeSymbol("write-to-file")
var symshen_4_5fraction_6 = MakeSymbol("shen.<fraction>")
var symshen_4unpack_1foreign = MakeSymbol("shen.unpack-foreign")
var symshen_4cut = MakeSymbol("shen.cut")
var symshen_4_5prem_6 = MakeSymbol("shen.<prem>")
var symshen_4_5side_6 = MakeSymbol("shen.<side>")
var symshen_4let_b = MakeSymbol("shen.let!")
var symshen_4reverse_1help = MakeSymbol("shen.reverse-help")
var sym_dmaximum_1print_1sequence_1size_d = MakeSymbol("*maximum-print-sequence-size*")
var symshen_4_dnames_d = MakeSymbol("shen.*names*")
var symshen_4typename_1h = MakeSymbol("shen.typename-h")
var symshen_4recursive_1string_1match = MakeSymbol("shen.recursive-string-match")
var symshen_4load_1help = MakeSymbol("shen.load-help")
var symTm = MakeSymbol("Tm")
var symP = MakeSymbol("P")
var symshen_4intern_1in_1package = MakeSymbol("shen.intern-in-package")
var symshen_4find_1types = MakeSymbol("shen.find-types")
var symshen_4process_1applications = MakeSymbol("shen.process-applications")
var symshen_4factor_1selectors_1h = MakeSymbol("shen.factor-selectors-h")
var symshen_4char_1stinput_2 = MakeSymbol("shen.char-stinput?")
var symshen_4_5multiline_6 = MakeSymbol("shen.<multiline>")
var symshen_4_5times_6 = MakeSymbol("shen.<times>")
var symshen_4compile_1premises = MakeSymbol("shen.compile-premises")
var symshen_4shen_1_6kl = MakeSymbol("shen.shen->kl")
var symexternal = MakeSymbol("external")
var symlimit = MakeSymbol("limit")
var sym_dimplementation_d = MakeSymbol("*implementation*")
var symdefun = MakeSymbol("defun")
var symshen_4try_1parse = MakeSymbol("shen.try-parse")
var symshen_4_5clause_6 = MakeSymbol("shen.<clause>")
var symshen_4deref_1calls = MakeSymbol("shen.deref-calls")
var symshen_4_dtc_d = MakeSymbol("shen.*tc*")
var sym_4_4_4 = MakeSymbol("...")
var symshen_4factor_1recognisors = MakeSymbol("shen.factor-recognisors")
var symread_1file = MakeSymbol("read-file")
var symshen_4extraspecial_2 = MakeSymbol("shen.extraspecial?")
var symshen_4sng_1h_2 = MakeSymbol("shen.sng-h?")
var symshen_4c_1rule_1_6shen = MakeSymbol("shen.c-rule->shen")
var symshen_4monomorphic_2 = MakeSymbol("shen.monomorphic?")
var symshen_4this_1symbol_1is_1unbound = MakeSymbol("shen.this-symbol-is-unbound")
var sym_dos_d = MakeSymbol("*os*")
var symsystem_1S_2 = MakeSymbol("system-S?")
var symshen_4bad_1pivot_2 = MakeSymbol("shen.bad-pivot?")
var symshen_4_5comma_6 = MakeSymbol("shen.<comma>")
var symassertz = MakeSymbol("assertz")
var symshen_4prompt = MakeSymbol("shen.prompt")
var symshen_4typetable = MakeSymbol("shen.typetable")
var symtuple_2 = MakeSymbol("tuple?")
var syminput_7 = MakeSymbol("input+")
var symshen_4_5define_6 = MakeSymbol("shen.<define>")
var symshen_4nchars = MakeSymbol("shen.nchars")
var symshen_4read_1unit_1string = MakeSymbol("shen.read-unit-string")
var symbar_b = MakeSymbol("bar!")
var symget_1time = MakeSymbol("get-time")
var symshen_4update_1history = MakeSymbol("shen.update-history")
var symread_1file_1as_1string = MakeSymbol("read-file-as-string")
var symassoc = MakeSymbol("assoc")
var symshen_4write_1chars = MakeSymbol("shen.write-chars")
var symshen_4maxseq = MakeSymbol("shen.maxseq")
var symshen_4extract_1vars = MakeSymbol("shen.extract-vars")
var sym_5_1address = MakeSymbol("<-address")
var symbootstrap = MakeSymbol("bootstrap")
var symshen_4_5type_6 = MakeSymbol("shen.<type>")
var symshen_4digit_2 = MakeSymbol("shen.digit?")
var symgensym = MakeSymbol("gensym")
var symunion = MakeSymbol("union")
var symshen_4sysfunc_2 = MakeSymbol("shen.sysfunc?")
var symshen_4_5e_1number_6 = MakeSymbol("shen.<e-number>")
var symshen_4lch = MakeSymbol("shen.lch")
var symshen_4tame = MakeSymbol("shen.tame")
var symshen_4findall_1h = MakeSymbol("shen.findall-h")
var symupdate_1lambda_1table = MakeSymbol("update-lambda-table")
var symcond = MakeSymbol("cond")
var symshen_4_5whitespace_6 = MakeSymbol("shen.<whitespace>")
var symshen_4process_1cases = MakeSymbol("shen.process-cases")
var symTime = MakeSymbol("Time")
var symshen_4_5clauses_6 = MakeSymbol("shen.<clauses>")
var symshen_4t_d = MakeSymbol("shen.t*")
var symshen_4myassume = MakeSymbol("shen.myassume")
var symshen_4_5s_1exprs2_6 = MakeSymbol("shen.<s-exprs2>")
var symshen_4_5digits_6 = MakeSymbol("shen.<digits>")
var symshen_4lowercase_1symbol_2 = MakeSymbol("shen.lowercase-symbol?")
var symshen_4app = MakeSymbol("shen.app")
var symfst = MakeSymbol("fst")
var symshen_4list_2 = MakeSymbol("shen.list?")
var symshen_4unprotect = MakeSymbol("shen.unprotect")
var symGoTo = MakeSymbol("GoTo")
var symboolean_2 = MakeSymbol("boolean?")
var symshen_4insert_1h = MakeSymbol("shen.insert-h")
var symshen_4choicepoint_b = MakeSymbol("shen.choicepoint!")
var symshen_4print_1residue = MakeSymbol("shen.print-residue")
var symwhen = MakeSymbol("when")
var symshen_4work_1through = MakeSymbol("shen.work-through")
var symshen_4retract_1clause = MakeSymbol("shen.retract-clause")
var symshen_4horn_1clause_1procedure = MakeSymbol("shen.horn-clause-procedure")
var symshen_4shen_1call_2 = MakeSymbol("shen.shen-call?")
var symshen_4alpha_2 = MakeSymbol("shen.alpha?")
var sym_i = MakeSymbol("{")
var symshen_4_5simple_1pattern_6 = MakeSymbol("shen.<simple-pattern>")
var symshen_4read_1file_1as_1bytelist_1help = MakeSymbol("shen.read-file-as-bytelist-help")
var symshen_4_5rcurly_6 = MakeSymbol("shen.<rcurly>")
var symprofile = MakeSymbol("profile")
var symshen_4rules_1_6prolog = MakeSymbol("shen.rules->prolog")
var symshen_4pvar = MakeSymbol("shen.pvar")
var symshen_4parse_1failure_2 = MakeSymbol("shen.parse-failure?")
var symshen_4read_1loop = MakeSymbol("shen.read-loop")
var symshen_4_5lsb_6 = MakeSymbol("shen.<lsb>")
var symshen_4package_1symbols = MakeSymbol("shen.package-symbols")
var symX = MakeSymbol("X")
var symshen_4compile_1assumptions = MakeSymbol("shen.compile-assumptions")
var symshen_4input_1track = MakeSymbol("shen.input-track")
var symshen_4_doccurs_d = MakeSymbol("shen.*occurs*")
var symshen_4return_2 = MakeSymbol("shen.return?")
var symshen_4sng_2 = MakeSymbol("shen.sng?")
var sym_d = MakeSymbol("*")
var symsimple_1error = MakeSymbol("simple-error")
var symtime = MakeSymbol("time")
var symshen_4prolog_1fbody = MakeSymbol("shen.prolog-fbody")
var symshen_4get_1profile = MakeSymbol("shen.get-profile")
var symfail_1if = MakeSymbol("fail-if")
var symrelease = MakeSymbol("release")
var symtype = MakeSymbol("type")
var symshen_4comb = MakeSymbol("shen.comb")
var symin = MakeSymbol("in")
var symshen_4fn_1call_2 = MakeSymbol("shen.fn-call?")
var symmode = MakeSymbol("mode")
var symsave = MakeSymbol("save")
var sym_8v = MakeSymbol("@v")
var symshen_4str_1_6str = MakeSymbol("shen.str->str")
var symshen_4triple_1stack = MakeSymbol("shen.triple-stack")
var symread_1file_1as_1bytelist = MakeSymbol("read-file-as-bytelist")
var syminclude = MakeSymbol("include")
var symshen_4atom_1case_1plus = MakeSymbol("shen.atom-case-plus")
var symshen_4lr_1rule = MakeSymbol("shen.lr-rule")
var symshen_4type_1theory_1enabled_2 = MakeSymbol("shen.type-theory-enabled?")
var symshen_4_duserdefs_d = MakeSymbol("shen.*userdefs*")
var symshen_4shen_1_6kl_1h = MakeSymbol("shen.shen->kl-h")
var symincluded = MakeSymbol("included")
var sym_5 = MakeSymbol("<")
var symL = MakeSymbol("L")
var symshen_4_dsize_1prolog_1vector_d = MakeSymbol("shen.*size-prolog-vector*")
var symshen_4pause_1for_1user = MakeSymbol("shen.pause-for-user")
var symshen_4lookupsig = MakeSymbol("shen.lookupsig")
var symget = MakeSymbol("get")
var symenable_1type_1theory = MakeSymbol("enable-type-theory")
var symshen_4_dextraspecial_d = MakeSymbol("shen.*extraspecial*")
var symshen_4parse_1failure = MakeSymbol("shen.parse-failure")
var sym_5_1 = MakeSymbol("<-")
var symshen_4op_1test = MakeSymbol("shen.op-test")
var symtlstr = MakeSymbol("tlstr")
var symshen_4undefined_1f_2 = MakeSymbol("shen.undefined-f?")
var sym_dproperty_1vector_d = MakeSymbol("*property-vector*")
var symvector_2 = MakeSymbol("vector?")
var symshen_4fix_1help = MakeSymbol("shen.fix-help")
var symabort = MakeSymbol("abort")
var symshen_4misc_2 = MakeSymbol("shen.misc?")
var symshen_4internal_1to_1P_2 = MakeSymbol("shen.internal-to-P?")
var symA = MakeSymbol("A")
var symshen_4cons_1case_1minus = MakeSymbol("shen.cons-case-minus")
var symshen_4internal_1symbols = MakeSymbol("shen.internal-symbols")
var symshen_4_5constructor_6 = MakeSymbol("shen.<constructor>")
var symvar_2 = MakeSymbol("var?")
var symshen_4process_1_8s = MakeSymbol("shen.process-@s")
var symshen_4demode = MakeSymbol("shen.demode")
var symshen_4expand_1mode_1forms = MakeSymbol("shen.expand-mode-forms")
var symshen_4search = MakeSymbol("shen.search")
var symshen_4freshen_1rule = MakeSymbol("shen.freshen-rule")
var symvector_1_6 = MakeSymbol("vector->")
var symstring_1_6n = MakeSymbol("string->n")
var symshen_4lock = MakeSymbol("shen.lock")
var symshen_4unwind = MakeSymbol("shen.unwind")
var symshen_4vector_1parameter = MakeSymbol("shen.vector-parameter")
var symshen_4system_1S_1h = MakeSymbol("shen.system-S-h")
var symshen_4freshen_1type = MakeSymbol("shen.freshen-type")
var symspecialise = MakeSymbol("specialise")
var symshen_4dbl_2 = MakeSymbol("shen.dbl?")
var symboolean = MakeSymbol("boolean")
var symshen_4bottom = MakeSymbol("shen.bottom")
var symshen_4curry = MakeSymbol("shen.curry")
var symread = MakeSymbol("read")
var symshen_4_5singleline_6 = MakeSymbol("shen.<singleline>")
var symshen_4expt = MakeSymbol("shen.expt")
var symshen_4_dsynonyms_d = MakeSymbol("shen.*synonyms*")
var symshen_4optimise_1passive = MakeSymbol("shen.optimise-passive")
var symshen_4show = MakeSymbol("shen.show")
var symshen_4_5syntax_1item_6 = MakeSymbol("shen.<syntax-item>")
var symshen_4variablecode = MakeSymbol("shen.variablecode")
var symhdstr = MakeSymbol("hdstr")
var symoptimise_2 = MakeSymbol("optimise?")
var symshen_4cond_1form = MakeSymbol("shen.cond-form")
var symshen_4_5shortnatters_6 = MakeSymbol("shen.<shortnatters>")
var symset = MakeSymbol("set")
var symshen_4nextticket = MakeSymbol("shen.nextticket")
var symshen_4occurs_1check_2 = MakeSymbol("shen.occurs-check?")
var symshen_4compile_1assumption = MakeSymbol("shen.compile-assumption")
var symelement_2 = MakeSymbol("element?")
var symhush_2 = MakeSymbol("hush?")
var symshen_4_5lcurly_6 = MakeSymbol("shen.<lcurly>")
var symshen_4optimise_1passive_1h = MakeSymbol("shen.optimise-passive-h")
var symshen_4recursively_1print = MakeSymbol("shen.recursively-print")
var symshen_4toplevel_1forms = MakeSymbol("shen.toplevel-forms")
var symread_1byte = MakeSymbol("read-byte")
var symshen_4show_1assumptions = MakeSymbol("shen.show-assumptions")
var symshen_4colon_1equal_2 = MakeSymbol("shen.colon-equal?")
var symshen_4syntax_1item_2 = MakeSymbol("shen.syntax-item?")
var symshen_4process_1yacc_1semantics = MakeSymbol("shen.process-yacc-semantics")
var symdifference = MakeSymbol("difference")
var symshen_4funexstring = MakeSymbol("shen.funexstring")
var sym_1_6 = MakeSymbol("->")
var symshen_4op2 = MakeSymbol("shen.op2")
var symshen_4cons_1form = MakeSymbol("shen.cons-form")
var symshen_4_5numeral_6 = MakeSymbol("shen.<numeral>")
var symshen_4_5strc_6 = MakeSymbol("shen.<strc>")
var symshen_4overapplication_2 = MakeSymbol("shen.overapplication?")
var symunput = MakeSymbol("unput")
var sym_dhome_1directory_d = MakeSymbol("*home-directory*")
var symshen_4length_1h = MakeSymbol("shen.length-h")
var symshen_4_5rule_6 = MakeSymbol("shen.<rule>")
var symshen_4walk = MakeSymbol("shen.walk")
var symshen_4process_1assoc = MakeSymbol("shen.process-assoc")
var symshen_4synonyms_1h = MakeSymbol("shen.synonyms-h")
var symshen_4_5expr_6 = MakeSymbol("shen.<expr>")
var symshen_4uppercase_2 = MakeSymbol("shen.uppercase?")
var symshen_4record_1kl = MakeSymbol("shen.record-kl")
var symdefine = MakeSymbol("define")
var symshen_4_5signature_6 = MakeSymbol("shen.<signature>")
var symdatatype = MakeSymbol("datatype")
var symshen_4g = MakeSymbol("shen.g")
var symshen_4_5literal_6 = MakeSymbol("shen.<literal>")
var symshen_4function_1calls = MakeSymbol("shen.function-calls")
var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var symshen_4in_1_6 = MakeSymbol("shen.in->")
var symshen_4free_1variable_2 = MakeSymbol("shen.free-variable?")
var symstr = MakeSymbol("str")
var symdefprolog = MakeSymbol("defprolog")
var symshen_4prtl = MakeSymbol("shen.prtl")
var symshen_4_dgensym_d = MakeSymbol("shen.*gensym*")
var symunprofile = MakeSymbol("unprofile")
var symC = MakeSymbol("C")
var symshen_4signal_1def = MakeSymbol("shen.signal-def")
var symsymbol_2 = MakeSymbol("symbol?")
var symshen_4_dinfs_d = MakeSymbol("shen.*infs*")
var symfresh = MakeSymbol("fresh")
var symshen_4_5rsb_6 = MakeSymbol("shen.<rsb>")
var symshen_4_5plus_6 = MakeSymbol("shen.<plus>")
var symshen_4packaged_2 = MakeSymbol("shen.packaged?")
var symshen_4curry_1type = MakeSymbol("shen.curry-type")
var symshen_4received = MakeSymbol("shen.received")
var symshen_4rfas_1h = MakeSymbol("shen.rfas-h")
var symsystemf = MakeSymbol("systemf")
var symshen_4process_1datatype = MakeSymbol("shen.process-datatype")
var symshen_4pretty_1type = MakeSymbol("shen.pretty-type")
var symshen_4compile_1body = MakeSymbol("shen.compile-body")
var symshen_4primitive = MakeSymbol("shen.primitive")
var symsum = MakeSymbol("sum")
var symshen_4_5_1out = MakeSymbol("shen.<-out")
var symshen_4_5lrb_6 = MakeSymbol("shen.<lrb>")
var symshen_4zero_1place_2 = MakeSymbol("shen.zero-place?")
var symuntrack = MakeSymbol("untrack")
var symshen_4callrec = MakeSymbol("shen.callrec")
var symimplementation = MakeSymbol("implementation")
var symlambda = MakeSymbol("lambda")
var symshen_4compileBytes = MakeSymbol("shen.compileBytes")
var symshen_4internal_2 = MakeSymbol("shen.internal?")
var symintern = MakeSymbol("intern")
var symshen_4cons_1form_1with_1modes = MakeSymbol("shen.cons-form-with-modes")
var symshen_4tabulate_1passive = MakeSymbol("shen.tabulate-passive")
var symshen_4prolog_1abstraction = MakeSymbol("shen.prolog-abstraction")
var symadjoin = MakeSymbol("adjoin")
var symshen_4unlock = MakeSymbol("shen.unlock")
var symshen_4fits_2 = MakeSymbol("shen.fits?")
var symshen_4tuple = MakeSymbol("shen.tuple")
var symhd = MakeSymbol("hd")
var symstring = MakeSymbol("string")
var symshen_4_5formulae_6 = MakeSymbol("shen.<formulae>")
var symPrevious = MakeSymbol("Previous")
var symsnd = MakeSymbol("snd")
var symfreeze = MakeSymbol("freeze")
var symshen_4_dprolog_1memory_d = MakeSymbol("shen.*prolog-memory*")
var symspy_2 = MakeSymbol("spy?")
var symshen_4demod = MakeSymbol("shen.demod")
var symshen_4lchh = MakeSymbol("shen.lchh")
var symshen_4c_1rules_1_6shen = MakeSymbol("shen.c-rules->shen")
var symhash = MakeSymbol("hash")
var symshen_4multiples = MakeSymbol("shen.multiples")
var symshen_4t = MakeSymbol("shen.t")
var symit = MakeSymbol("it")
var symshen_4_dprofiled_d = MakeSymbol("shen.*profiled*")
var symshen_4print_1freshterm = MakeSymbol("shen.print-freshterm")
var symshen_4_5hterm2_6 = MakeSymbol("shen.<hterm2>")
var symshen_4locked_2 = MakeSymbol("shen.locked?")
var sym_8s = MakeSymbol("@s")
var symshen_4_dshen_1type_1theory_1enabled_2_d = MakeSymbol("shen.*shen-type-theory-enabled?*")
var symcases = MakeSymbol("cases")
var symshen_4lzy_a = MakeSymbol("shen.lzy=")
var symshen_4decons = MakeSymbol("shen.decons")
var symshen_4factor_1cn = MakeSymbol("shen.factor-cn")
var symshen_4_5s_1exprs_6 = MakeSymbol("shen.<s-exprs>")
var sym_5_a = MakeSymbol("<=")
var symshen_4prolog_1vector = MakeSymbol("shen.prolog-vector")
var symshen_4compile_1consequent = MakeSymbol("shen.compile-consequent")
var symstep_2 = MakeSymbol("step?")
var symshen_4read_1evaluate_1print = MakeSymbol("shen.read-evaluate-print")
var symshen_4wildcardcode = MakeSymbol("shen.wildcardcode")
var sym_dhush_d = MakeSymbol("*hush*")
var symshen_4out_1of_1bounds = MakeSymbol("shen.out-of-bounds")
var symshen_4choicepoint = MakeSymbol("shen.choicepoint")
var symfork = MakeSymbol("fork")
var symunspecialise = MakeSymbol("unspecialise")
var symshen_4loop = MakeSymbol("shen.loop")
var symshen_4goals = MakeSymbol("shen.goals")
var symshen_4freshterms = MakeSymbol("shen.freshterms")
var symlength = MakeSymbol("length")
var symshen_4_5s_1exprs1_6 = MakeSymbol("shen.<s-exprs1>")
var symshen_4partial_1application_d_2 = MakeSymbol("shen.partial-application*?")
var symshen_4compile_1head = MakeSymbol("shen.compile-head")
var symshen_4_5rules_d_6 = MakeSymbol("shen.<rules*>")
var symshen_4t_d_1correct = MakeSymbol("shen.t*-correct")
var symshen_4_5semantics_6 = MakeSymbol("shen.<semantics>")
var symshen_4conscode = MakeSymbol("shen.conscode")
var sym_8p = MakeSymbol("@p")
var sym_dporters_d = MakeSymbol("*porters*")
var symshen_4_5backslash_6 = MakeSymbol("shen.<backslash>")
var symshen_4_5digit_6 = MakeSymbol("shen.<digit>")
var symsynonyms = MakeSymbol("synonyms")
var symshen_4foreign_2 = MakeSymbol("shen.foreign?")
var symbind = MakeSymbol("bind")
var symshen_4record_1macro = MakeSymbol("shen.record-macro")
var symshen_4source = MakeSymbol("shen.source")
var symshen_4fn_1print = MakeSymbol("shen.fn-print")
var symshen_4_5notdbq_6 = MakeSymbol("shen.<notdbq>")
var symshen_4_5minus_6 = MakeSymbol("shen.<minus>")
var symshen_4_5hterm1_6 = MakeSymbol("shen.<hterm1>")
var symshen_4semicolon_2 = MakeSymbol("shen.semicolon?")
var symshen_4openlock = MakeSymbol("shen.openlock")
var symshen_4variable_1case = MakeSymbol("shen.variable-case")
var symnot = MakeSymbol("not")
var symmake_1string = MakeSymbol("make-string")
var symshen_4bind_b = MakeSymbol("shen.bind!")
var symshen_4_5double_6 = MakeSymbol("shen.<double>")
var symshen_4nvars = MakeSymbol("shen.nvars")
var symshen_4ok = MakeSymbol("shen.ok")
var symshen_4freshen = MakeSymbol("shen.freshen")
var symshen_4terminalcode = MakeSymbol("shen.terminalcode")
var sym_a_a = MakeSymbol("==")
var symtc = MakeSymbol("tc")
var symshen_4free_1var_1chk = MakeSymbol("shen.free-var-chk")
var symshen_4_dspy_d = MakeSymbol("shen.*spy*")
var symshen_4_5conc_6 = MakeSymbol("shen.<conc>")
var symshen_4body_1foundit_b = MakeSymbol("shen.body-foundit!")
var symshen_4compile_1premise_1h = MakeSymbol("shen.compile-premise-h")
