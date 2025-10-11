package main

import . "github.com/tiancaiamao/shen-go/kl"

var LoadMain = MakeNative(func(__e *ControlFlow) {
tmp9507 := MakeNative(func(__e *ControlFlow) {
V874 := __e.Get(1)
_ = V874
tmp9508 := MakeNative(func(__e *ControlFlow) {
W875 := __e.Get(1)
_ = W875
tmp9509 := MakeNative(func(__e *ControlFlow) {
W876 := __e.Get(1)
_ = W876
tmp9510 := MakeNative(func(__e *ControlFlow) {
W882 := __e.Get(1)
_ = W882
__e.Return(symloaded)
return
}, 1)

var ifres9511 Obj

if True == W875 {
tmp9512 := Call(__e, PrimFunc(syminferences))


tmp9513 := Call(__e, PrimFunc(symshen_4app), tmp9512, MakeString(" inferences\n"), symshen_4a)


tmp9514 := PrimStringConcat(MakeString("\ntypechecked in "), tmp9513)

tmp9515 := Call(__e, PrimFunc(symstoutput))


tmp9516 := Call(__e, PrimFunc(sympr), tmp9514, tmp9515)


ifres9511 = tmp9516


} else {
ifres9511 = symshen_4skip


}

__e.TailApply(tmp9510, ifres9511)
return


}, 1)

tmp9517 := MakeNative(func(__e *ControlFlow) {
W877 := __e.Get(1)
_ = W877
tmp9518 := MakeNative(func(__e *ControlFlow) {
W878 := __e.Get(1)
_ = W878
tmp9519 := MakeNative(func(__e *ControlFlow) {
W879 := __e.Get(1)
_ = W879
tmp9520 := MakeNative(func(__e *ControlFlow) {
W880 := __e.Get(1)
_ = W880
tmp9521 := MakeNative(func(__e *ControlFlow) {
W881 := __e.Get(1)
_ = W881
__e.Return(W878)
return
}, 1)

tmp9522 := PrimStr(W880)

tmp9523 := PrimStringConcat(tmp9522, MakeString(" secs\n"))

tmp9524 := PrimStringConcat(MakeString("\nrun time: "), tmp9523)

tmp9525 := Call(__e, PrimFunc(symstoutput))


tmp9526 := Call(__e, PrimFunc(sympr), tmp9524, tmp9525)


__e.TailApply(tmp9521, tmp9526)
return


}, 1)

tmp9527 := PrimNumberSubtract(W879, W877)

__e.TailApply(tmp9520, tmp9527)
return


}, 1)

tmp9528 := PrimGetTime(symrun)

__e.TailApply(tmp9519, tmp9528)
return


}, 1)

tmp9529 := Call(__e, PrimFunc(symread_1file), V874)


tmp9530 := Call(__e, PrimFunc(symshen_4load_1help), W875, tmp9529)


__e.TailApply(tmp9518, tmp9530)
return


}, 1)

tmp9531 := PrimGetTime(symrun)

tmp9532 := Call(__e, tmp9517, tmp9531)


__e.TailApply(tmp9509, tmp9532)
return


}, 1)

tmp9533 := PrimValue(symshen_4_dtc_d)

__e.TailApply(tmp9508, tmp9533)
return


}, 1)

tmp9534 := Call(__e, ns2_1set, symload, tmp9507)


_ = tmp9534

tmp9535 := MakeNative(func(__e *ControlFlow) {
V885 := __e.Get(1)
_ = V885
V886 := __e.Get(2)
_ = V886
tmp9537 := PrimEqual(False, V885)

if True == tmp9537 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V886)
return
} else {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V886)
return
}


}, 2)

tmp9538 := Call(__e, ns2_1set, symshen_4load_1help, tmp9535)


_ = tmp9538

tmp9539 := MakeNative(func(__e *ControlFlow) {
V887 := __e.Get(1)
_ = V887
tmp9540 := MakeNative(func(__e *ControlFlow) {
Z888 := __e.Get(1)
_ = Z888
tmp9541 := Call(__e, PrimFunc(symshen_4shen_1_6kl), Z888)


tmp9542 := Call(__e, PrimFunc(symeval_1kl), tmp9541)


tmp9543 := Call(__e, PrimFunc(symshen_4app), tmp9542, MakeString("\n"), symshen_4s)


tmp9544 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp9543, tmp9544)
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp9540, V887)
return


}, 1)

tmp9545 := Call(__e, ns2_1set, symshen_4eval_1and_1print, tmp9539)


_ = tmp9545

tmp9546 := MakeNative(func(__e *ControlFlow) {
V889 := __e.Get(1)
_ = V889
tmp9547 := MakeNative(func(__e *ControlFlow) {
W890 := __e.Get(1)
_ = W890
tmp9548 := MakeNative(func(__e *ControlFlow) {
W892 := __e.Get(1)
_ = W892
tmp9549 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4work_1through), V889)
return
}, 0)

tmp9550 := MakeNative(func(__e *ControlFlow) {
Z894 := __e.Get(1)
_ = Z894
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z894, W890)
return
}, 1)

__e.TailApply(try_1catch, tmp9549, tmp9550)
return


}, 1)

tmp9551 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4assumetypes), W890)
return
}, 0)

tmp9552 := MakeNative(func(__e *ControlFlow) {
Z893 := __e.Get(1)
_ = Z893
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z893, W890)
return
}, 1)

tmp9553 := Call(__e, try_1catch, tmp9551, tmp9552)


__e.TailApply(tmp9548, tmp9553)
return


}, 1)

tmp9554 := MakeNative(func(__e *ControlFlow) {
Z891 := __e.Get(1)
_ = Z891
__e.TailApply(PrimFunc(symshen_4typetable), Z891)
return
}, 1)

tmp9555 := Call(__e, PrimFunc(symmapcan), tmp9554, V889)


__e.TailApply(tmp9547, tmp9555)
return


}, 1)

tmp9556 := Call(__e, ns2_1set, symshen_4check_1eval_1and_1print, tmp9546)


_ = tmp9556

tmp9557 := MakeNative(func(__e *ControlFlow) {
V899 := __e.Get(1)
_ = V899
tmp9602 := PrimIsPair(V899)

var ifres9583 Obj

if True == tmp9602 {
tmp9600 := PrimHead(V899)

tmp9601 := PrimEqual(symdefine, tmp9600)

var ifres9585 Obj

if True == tmp9601 {
tmp9598 := PrimTail(V899)

tmp9599 := PrimIsPair(tmp9598)

var ifres9587 Obj

if True == tmp9599 {
tmp9595 := PrimTail(V899)

tmp9596 := PrimTail(tmp9595)

tmp9597 := PrimIsPair(tmp9596)

var ifres9589 Obj

if True == tmp9597 {
tmp9591 := PrimTail(V899)

tmp9592 := PrimTail(tmp9591)

tmp9593 := PrimHead(tmp9592)

tmp9594 := PrimEqual(sym_i, tmp9593)

var ifres9590 Obj

if True == tmp9594 {
ifres9590 = True


} else {
ifres9590 = False


}

ifres9589 = ifres9590


} else {
ifres9589 = False


}

var ifres9588 Obj

if True == ifres9589 {
ifres9588 = True


} else {
ifres9588 = False


}

ifres9587 = ifres9588


} else {
ifres9587 = False


}

var ifres9586 Obj

if True == ifres9587 {
ifres9586 = True


} else {
ifres9586 = False


}

ifres9585 = ifres9586


} else {
ifres9585 = False


}

var ifres9584 Obj

if True == ifres9585 {
ifres9584 = True


} else {
ifres9584 = False


}

ifres9583 = ifres9584


} else {
ifres9583 = False


}

if True == ifres9583 {
tmp9558 := PrimTail(V899)

tmp9559 := PrimHead(tmp9558)

tmp9560 := PrimTail(V899)

tmp9561 := PrimHead(tmp9560)

tmp9562 := PrimTail(V899)

tmp9563 := PrimTail(tmp9562)

tmp9564 := PrimTail(tmp9563)

tmp9565 := Call(__e, PrimFunc(symshen_4type_1F), tmp9561, tmp9564)


tmp9566 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp9565)


tmp9567 := PrimCons(tmp9566, Nil)

__e.Return(PrimCons(tmp9559, tmp9567))
return


} else {
tmp9581 := PrimIsPair(V899)

var ifres9573 Obj

if True == tmp9581 {
tmp9579 := PrimHead(V899)

tmp9580 := PrimEqual(symdefine, tmp9579)

var ifres9575 Obj

if True == tmp9580 {
tmp9577 := PrimTail(V899)

tmp9578 := PrimIsPair(tmp9577)

var ifres9576 Obj

if True == tmp9578 {
ifres9576 = True


} else {
ifres9576 = False


}

ifres9575 = ifres9576


} else {
ifres9575 = False


}

var ifres9574 Obj

if True == ifres9575 {
ifres9574 = True


} else {
ifres9574 = False


}

ifres9573 = ifres9574


} else {
ifres9573 = False


}

if True == ifres9573 {
tmp9568 := PrimTail(V899)

tmp9569 := PrimHead(tmp9568)

tmp9570 := Call(__e, PrimFunc(symshen_4app), tmp9569, MakeString("\n"), symshen_4a)


tmp9571 := PrimStringConcat(MakeString("missing { in "), tmp9570)

__e.Return(PrimSimpleError(tmp9571))
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp9603 := Call(__e, ns2_1set, symshen_4typetable, tmp9557)


_ = tmp9603

tmp9604 := MakeNative(func(__e *ControlFlow) {
V906 := __e.Get(1)
_ = V906
V907 := __e.Get(2)
_ = V907
tmp9617 := PrimIsPair(V907)

var ifres9613 Obj

if True == tmp9617 {
tmp9615 := PrimHead(V907)

tmp9616 := PrimEqual(sym_j, tmp9615)

var ifres9614 Obj

if True == tmp9616 {
ifres9614 = True


} else {
ifres9614 = False


}

ifres9613 = ifres9614


} else {
ifres9613 = False


}

if True == ifres9613 {
__e.Return(Nil)
return
} else {
tmp9611 := PrimIsPair(V907)

if True == tmp9611 {
tmp9605 := PrimHead(V907)

tmp9606 := PrimTail(V907)

tmp9607 := Call(__e, PrimFunc(symshen_4type_1F), V906, tmp9606)


__e.Return(PrimCons(tmp9605, tmp9607))
return


} else {
tmp9608 := Call(__e, PrimFunc(symshen_4app), V906, MakeString("\n"), symshen_4a)


tmp9609 := PrimStringConcat(MakeString("missing } in "), tmp9608)

__e.Return(PrimSimpleError(tmp9609))
return


}


}


}, 2)

tmp9618 := Call(__e, ns2_1set, symshen_4type_1F, tmp9604)


_ = tmp9618

tmp9619 := MakeNative(func(__e *ControlFlow) {
V910 := __e.Get(1)
_ = V910
tmp9633 := PrimEqual(Nil, V910)

if True == tmp9633 {
__e.Return(Nil)
return
} else {
tmp9631 := PrimIsPair(V910)

var ifres9627 Obj

if True == tmp9631 {
tmp9629 := PrimTail(V910)

tmp9630 := PrimIsPair(tmp9629)

var ifres9628 Obj

if True == tmp9630 {
ifres9628 = True


} else {
ifres9628 = False


}

ifres9627 = ifres9628


} else {
ifres9627 = False


}

if True == ifres9627 {
tmp9620 := PrimHead(V910)

tmp9621 := PrimTail(V910)

tmp9622 := PrimHead(tmp9621)

tmp9623 := Call(__e, PrimFunc(symdeclare), tmp9620, tmp9622)


_ = tmp9623

tmp9624 := PrimTail(V910)

tmp9625 := PrimTail(tmp9624)

__e.TailApply(PrimFunc(symshen_4assumetypes), tmp9625)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.assumetype")))
return
}


}


}, 1)

tmp9634 := Call(__e, ns2_1set, symshen_4assumetypes, tmp9619)


_ = tmp9634

tmp9635 := MakeNative(func(__e *ControlFlow) {
V915 := __e.Get(1)
_ = V915
V916 := __e.Get(2)
_ = V916
tmp9646 := PrimIsPair(V916)

var ifres9642 Obj

if True == tmp9646 {
tmp9644 := PrimTail(V916)

tmp9645 := PrimIsPair(tmp9644)

var ifres9643 Obj

if True == tmp9645 {
ifres9643 = True


} else {
ifres9643 = False


}

ifres9642 = ifres9643


} else {
ifres9642 = False


}

if True == ifres9642 {
tmp9636 := PrimHead(V916)

tmp9637 := Call(__e, PrimFunc(symdestroy), tmp9636)


_ = tmp9637

tmp9638 := PrimTail(V916)

tmp9639 := PrimTail(tmp9638)

__e.TailApply(PrimFunc(symshen_4unwind_1types), V915, tmp9639)
return


} else {
tmp9640 := PrimErrorToString(V915)

__e.Return(PrimSimpleError(tmp9640))
return


}


}, 2)

tmp9647 := Call(__e, ns2_1set, symshen_4unwind_1types, tmp9635)


_ = tmp9647

tmp9648 := MakeNative(func(__e *ControlFlow) {
V919 := __e.Get(1)
_ = V919
tmp9697 := PrimEqual(Nil, V919)

if True == tmp9697 {
__e.Return(Nil)
return
} else {
tmp9695 := PrimIsPair(V919)

var ifres9680 Obj

if True == tmp9695 {
tmp9693 := PrimTail(V919)

tmp9694 := PrimIsPair(tmp9693)

var ifres9682 Obj

if True == tmp9694 {
tmp9690 := PrimTail(V919)

tmp9691 := PrimTail(tmp9690)

tmp9692 := PrimIsPair(tmp9691)

var ifres9684 Obj

if True == tmp9692 {
tmp9686 := PrimTail(V919)

tmp9687 := PrimHead(tmp9686)

tmp9688 := PrimIntern(MakeString(":"))

tmp9689 := PrimEqual(tmp9687, tmp9688)

var ifres9685 Obj

if True == tmp9689 {
ifres9685 = True


} else {
ifres9685 = False


}

ifres9684 = ifres9685


} else {
ifres9684 = False


}

var ifres9683 Obj

if True == ifres9684 {
ifres9683 = True


} else {
ifres9683 = False


}

ifres9682 = ifres9683


} else {
ifres9682 = False


}

var ifres9681 Obj

if True == ifres9682 {
ifres9681 = True


} else {
ifres9681 = False


}

ifres9680 = ifres9681


} else {
ifres9680 = False


}

if True == ifres9680 {
tmp9649 := MakeNative(func(__e *ControlFlow) {
W920 := __e.Get(1)
_ = W920
tmp9665 := PrimEqual(W920, False)

if True == tmp9665 {
__e.TailApply(PrimFunc(symshen_4type_1error))
return
} else {
tmp9650 := MakeNative(func(__e *ControlFlow) {
W921 := __e.Get(1)
_ = W921
tmp9651 := MakeNative(func(__e *ControlFlow) {
W922 := __e.Get(1)
_ = W922
tmp9652 := PrimTail(V919)

tmp9653 := PrimTail(tmp9652)

tmp9654 := PrimTail(tmp9653)

__e.TailApply(PrimFunc(symshen_4work_1through), tmp9654)
return


}, 1)

tmp9655 := Call(__e, PrimFunc(symshen_4pretty_1type), W920)


tmp9656 := Call(__e, PrimFunc(symshen_4app), tmp9655, MakeString("\n"), symshen_4r)


tmp9657 := PrimStringConcat(MakeString(" : "), tmp9656)

tmp9658 := Call(__e, PrimFunc(symshen_4app), W921, tmp9657, symshen_4s)


tmp9659 := Call(__e, PrimFunc(symstoutput))


tmp9660 := Call(__e, PrimFunc(sympr), tmp9658, tmp9659)


__e.TailApply(tmp9651, tmp9660)
return


}, 1)

tmp9661 := PrimHead(V919)

tmp9662 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp9661)


tmp9663 := Call(__e, PrimFunc(symeval_1kl), tmp9662)


__e.TailApply(tmp9650, tmp9663)
return


}


}, 1)

tmp9666 := PrimHead(V919)

tmp9667 := PrimTail(V919)

tmp9668 := PrimTail(tmp9667)

tmp9669 := PrimHead(tmp9668)

tmp9670 := Call(__e, PrimFunc(symshen_4typecheck), tmp9666, tmp9669)


__e.TailApply(tmp9649, tmp9670)
return


} else {
tmp9678 := PrimIsPair(V919)

if True == tmp9678 {
tmp9671 := PrimHead(V919)

tmp9672 := PrimIntern(MakeString(":"))

tmp9673 := PrimTail(V919)

tmp9674 := PrimCons(symA, tmp9673)

tmp9675 := PrimCons(tmp9672, tmp9674)

tmp9676 := PrimCons(tmp9671, tmp9675)

__e.TailApply(PrimFunc(symshen_4work_1through), tmp9676)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.work-through")))
return
}


}


}


}, 1)

tmp9698 := Call(__e, ns2_1set, symshen_4work_1through, tmp9648)


_ = tmp9698

tmp9699 := MakeNative(func(__e *ControlFlow) {
V924 := __e.Get(1)
_ = V924
tmp9841 := PrimIsPair(V924)

var ifres9715 Obj

if True == tmp9841 {
tmp9839 := PrimHead(V924)

tmp9840 := PrimIsPair(tmp9839)

var ifres9717 Obj

if True == tmp9840 {
tmp9836 := PrimHead(V924)

tmp9837 := PrimHead(tmp9836)

tmp9838 := PrimEqual(symlist, tmp9837)

var ifres9719 Obj

if True == tmp9838 {
tmp9833 := PrimHead(V924)

tmp9834 := PrimTail(tmp9833)

tmp9835 := PrimIsPair(tmp9834)

var ifres9721 Obj

if True == tmp9835 {
tmp9829 := PrimHead(V924)

tmp9830 := PrimTail(tmp9829)

tmp9831 := PrimTail(tmp9830)

tmp9832 := PrimEqual(Nil, tmp9831)

var ifres9723 Obj

if True == tmp9832 {
tmp9827 := PrimTail(V924)

tmp9828 := PrimIsPair(tmp9827)

var ifres9725 Obj

if True == tmp9828 {
tmp9824 := PrimTail(V924)

tmp9825 := PrimHead(tmp9824)

tmp9826 := PrimEqual(sym_1_1_6, tmp9825)

var ifres9727 Obj

if True == tmp9826 {
tmp9821 := PrimTail(V924)

tmp9822 := PrimTail(tmp9821)

tmp9823 := PrimIsPair(tmp9822)

var ifres9729 Obj

if True == tmp9823 {
tmp9817 := PrimTail(V924)

tmp9818 := PrimTail(tmp9817)

tmp9819 := PrimHead(tmp9818)

tmp9820 := PrimIsPair(tmp9819)

var ifres9731 Obj

if True == tmp9820 {
tmp9812 := PrimTail(V924)

tmp9813 := PrimTail(tmp9812)

tmp9814 := PrimHead(tmp9813)

tmp9815 := PrimHead(tmp9814)

tmp9816 := PrimEqual(symstr, tmp9815)

var ifres9733 Obj

if True == tmp9816 {
tmp9807 := PrimTail(V924)

tmp9808 := PrimTail(tmp9807)

tmp9809 := PrimHead(tmp9808)

tmp9810 := PrimTail(tmp9809)

tmp9811 := PrimIsPair(tmp9810)

var ifres9735 Obj

if True == tmp9811 {
tmp9801 := PrimTail(V924)

tmp9802 := PrimTail(tmp9801)

tmp9803 := PrimHead(tmp9802)

tmp9804 := PrimTail(tmp9803)

tmp9805 := PrimHead(tmp9804)

tmp9806 := PrimIsPair(tmp9805)

var ifres9737 Obj

if True == tmp9806 {
tmp9794 := PrimTail(V924)

tmp9795 := PrimTail(tmp9794)

tmp9796 := PrimHead(tmp9795)

tmp9797 := PrimTail(tmp9796)

tmp9798 := PrimHead(tmp9797)

tmp9799 := PrimHead(tmp9798)

tmp9800 := PrimEqual(symlist, tmp9799)

var ifres9739 Obj

if True == tmp9800 {
tmp9787 := PrimTail(V924)

tmp9788 := PrimTail(tmp9787)

tmp9789 := PrimHead(tmp9788)

tmp9790 := PrimTail(tmp9789)

tmp9791 := PrimHead(tmp9790)

tmp9792 := PrimTail(tmp9791)

tmp9793 := PrimIsPair(tmp9792)

var ifres9741 Obj

if True == tmp9793 {
tmp9779 := PrimTail(V924)

tmp9780 := PrimTail(tmp9779)

tmp9781 := PrimHead(tmp9780)

tmp9782 := PrimTail(tmp9781)

tmp9783 := PrimHead(tmp9782)

tmp9784 := PrimTail(tmp9783)

tmp9785 := PrimTail(tmp9784)

tmp9786 := PrimEqual(Nil, tmp9785)

var ifres9743 Obj

if True == tmp9786 {
tmp9773 := PrimTail(V924)

tmp9774 := PrimTail(tmp9773)

tmp9775 := PrimHead(tmp9774)

tmp9776 := PrimTail(tmp9775)

tmp9777 := PrimTail(tmp9776)

tmp9778 := PrimIsPair(tmp9777)

var ifres9745 Obj

if True == tmp9778 {
tmp9766 := PrimTail(V924)

tmp9767 := PrimTail(tmp9766)

tmp9768 := PrimHead(tmp9767)

tmp9769 := PrimTail(tmp9768)

tmp9770 := PrimTail(tmp9769)

tmp9771 := PrimTail(tmp9770)

tmp9772 := PrimEqual(Nil, tmp9771)

var ifres9747 Obj

if True == tmp9772 {
tmp9762 := PrimTail(V924)

tmp9763 := PrimTail(tmp9762)

tmp9764 := PrimTail(tmp9763)

tmp9765 := PrimEqual(Nil, tmp9764)

var ifres9749 Obj

if True == tmp9765 {
tmp9751 := PrimHead(V924)

tmp9752 := PrimTail(tmp9751)

tmp9753 := PrimHead(tmp9752)

tmp9754 := PrimTail(V924)

tmp9755 := PrimTail(tmp9754)

tmp9756 := PrimHead(tmp9755)

tmp9757 := PrimTail(tmp9756)

tmp9758 := PrimHead(tmp9757)

tmp9759 := PrimTail(tmp9758)

tmp9760 := PrimHead(tmp9759)

tmp9761 := PrimEqual(tmp9753, tmp9760)

var ifres9750 Obj

if True == tmp9761 {
ifres9750 = True


} else {
ifres9750 = False


}

ifres9749 = ifres9750


} else {
ifres9749 = False


}

var ifres9748 Obj

if True == ifres9749 {
ifres9748 = True


} else {
ifres9748 = False


}

ifres9747 = ifres9748


} else {
ifres9747 = False


}

var ifres9746 Obj

if True == ifres9747 {
ifres9746 = True


} else {
ifres9746 = False


}

ifres9745 = ifres9746


} else {
ifres9745 = False


}

var ifres9744 Obj

if True == ifres9745 {
ifres9744 = True


} else {
ifres9744 = False


}

ifres9743 = ifres9744


} else {
ifres9743 = False


}

var ifres9742 Obj

if True == ifres9743 {
ifres9742 = True


} else {
ifres9742 = False


}

ifres9741 = ifres9742


} else {
ifres9741 = False


}

var ifres9740 Obj

if True == ifres9741 {
ifres9740 = True


} else {
ifres9740 = False


}

ifres9739 = ifres9740


} else {
ifres9739 = False


}

var ifres9738 Obj

if True == ifres9739 {
ifres9738 = True


} else {
ifres9738 = False


}

ifres9737 = ifres9738


} else {
ifres9737 = False


}

var ifres9736 Obj

if True == ifres9737 {
ifres9736 = True


} else {
ifres9736 = False


}

ifres9735 = ifres9736


} else {
ifres9735 = False


}

var ifres9734 Obj

if True == ifres9735 {
ifres9734 = True


} else {
ifres9734 = False


}

ifres9733 = ifres9734


} else {
ifres9733 = False


}

var ifres9732 Obj

if True == ifres9733 {
ifres9732 = True


} else {
ifres9732 = False


}

ifres9731 = ifres9732


} else {
ifres9731 = False


}

var ifres9730 Obj

if True == ifres9731 {
ifres9730 = True


} else {
ifres9730 = False


}

ifres9729 = ifres9730


} else {
ifres9729 = False


}

var ifres9728 Obj

if True == ifres9729 {
ifres9728 = True


} else {
ifres9728 = False


}

ifres9727 = ifres9728


} else {
ifres9727 = False


}

var ifres9726 Obj

if True == ifres9727 {
ifres9726 = True


} else {
ifres9726 = False


}

ifres9725 = ifres9726


} else {
ifres9725 = False


}

var ifres9724 Obj

if True == ifres9725 {
ifres9724 = True


} else {
ifres9724 = False


}

ifres9723 = ifres9724


} else {
ifres9723 = False


}

var ifres9722 Obj

if True == ifres9723 {
ifres9722 = True


} else {
ifres9722 = False


}

ifres9721 = ifres9722


} else {
ifres9721 = False


}

var ifres9720 Obj

if True == ifres9721 {
ifres9720 = True


} else {
ifres9720 = False


}

ifres9719 = ifres9720


} else {
ifres9719 = False


}

var ifres9718 Obj

if True == ifres9719 {
ifres9718 = True


} else {
ifres9718 = False


}

ifres9717 = ifres9718


} else {
ifres9717 = False


}

var ifres9716 Obj

if True == ifres9717 {
ifres9716 = True


} else {
ifres9716 = False


}

ifres9715 = ifres9716


} else {
ifres9715 = False


}

if True == ifres9715 {
tmp9700 := PrimTail(V924)

tmp9701 := PrimTail(tmp9700)

tmp9702 := PrimHead(tmp9701)

tmp9703 := PrimTail(tmp9702)

tmp9704 := PrimHead(tmp9703)

tmp9705 := PrimTail(V924)

tmp9706 := PrimTail(tmp9705)

tmp9707 := PrimHead(tmp9706)

tmp9708 := PrimTail(tmp9707)

tmp9709 := PrimTail(tmp9708)

tmp9710 := PrimCons(sym_a_a_6, tmp9709)

__e.Return(PrimCons(tmp9704, tmp9710))
return


} else {
tmp9713 := PrimIsPair(V924)

if True == tmp9713 {
tmp9711 := MakeNative(func(__e *ControlFlow) {
Z925 := __e.Get(1)
_ = Z925
__e.TailApply(PrimFunc(symshen_4pretty_1type), Z925)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp9711, V924)
return


} else {
__e.Return(V924)
return
}


}


}, 1)

tmp9842 := Call(__e, ns2_1set, symshen_4pretty_1type, tmp9699)


_ = tmp9842

tmp9843 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimSimpleError(MakeString("type error\n")))
return
}, 0)

tmp9844 := Call(__e, ns2_1set, symshen_4type_1error, tmp9843)


_ = tmp9844

tmp9845 := MakeNative(func(__e *ControlFlow) {
V926 := __e.Get(1)
_ = V926
tmp9846 := MakeNative(func(__e *ControlFlow) {
W927 := __e.Get(1)
_ = W927
tmp9847 := MakeNative(func(__e *ControlFlow) {
W928 := __e.Get(1)
_ = W928
tmp9848 := MakeNative(func(__e *ControlFlow) {
W929 := __e.Get(1)
_ = W929
tmp9849 := MakeNative(func(__e *ControlFlow) {
W930 := __e.Get(1)
_ = W930
tmp9850 := MakeNative(func(__e *ControlFlow) {
W932 := __e.Get(1)
_ = W932
__e.Return(W927)
return
}, 1)

tmp9851 := Call(__e, PrimFunc(symshen_4write_1kl), W930, W929)


__e.TailApply(tmp9850, tmp9851)
return


}, 1)

tmp9852 := MakeNative(func(__e *ControlFlow) {
Z931 := __e.Get(1)
_ = Z931
tmp9853 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), Z931)


__e.TailApply(PrimFunc(symshen_4partial), tmp9853)
return


}, 1)

tmp9854 := Call(__e, PrimFunc(symmap), tmp9852, W928)


__e.TailApply(tmp9849, tmp9854)
return


}, 1)

tmp9855 := PrimOpenStream(W927, symout)

__e.TailApply(tmp9848, tmp9855)
return


}, 1)

tmp9856 := Call(__e, PrimFunc(symread_1file), V926)


__e.TailApply(tmp9847, tmp9856)
return


}, 1)

tmp9857 := Call(__e, PrimFunc(symshen_4klfile), V926)


__e.TailApply(tmp9846, tmp9857)
return


}, 1)

tmp9858 := Call(__e, ns2_1set, symbootstrap, tmp9845)


_ = tmp9858

tmp9859 := MakeNative(func(__e *ControlFlow) {
V933 := __e.Get(1)
_ = V933
tmp9882 := PrimIsPair(V933)

var ifres9869 Obj

if True == tmp9882 {
tmp9880 := PrimHead(V933)

tmp9881 := PrimEqual(symshen_4f_1error, tmp9880)

var ifres9871 Obj

if True == tmp9881 {
tmp9878 := PrimTail(V933)

tmp9879 := PrimIsPair(tmp9878)

var ifres9873 Obj

if True == tmp9879 {
tmp9875 := PrimTail(V933)

tmp9876 := PrimTail(tmp9875)

tmp9877 := PrimEqual(Nil, tmp9876)

var ifres9874 Obj

if True == tmp9877 {
ifres9874 = True


} else {
ifres9874 = False


}

ifres9873 = ifres9874


} else {
ifres9873 = False


}

var ifres9872 Obj

if True == ifres9873 {
ifres9872 = True


} else {
ifres9872 = False


}

ifres9871 = ifres9872


} else {
ifres9871 = False


}

var ifres9870 Obj

if True == ifres9871 {
ifres9870 = True


} else {
ifres9870 = False


}

ifres9869 = ifres9870


} else {
ifres9869 = False


}

if True == ifres9869 {
tmp9860 := PrimTail(V933)

tmp9861 := PrimHead(tmp9860)

tmp9862 := PrimStr(tmp9861)

tmp9863 := PrimStringConcat(MakeString("partial function "), tmp9862)

tmp9864 := PrimCons(tmp9863, Nil)

__e.Return(PrimCons(symsimple_1error, tmp9864))
return


} else {
tmp9867 := PrimIsPair(V933)

if True == tmp9867 {
tmp9865 := MakeNative(func(__e *ControlFlow) {
Z934 := __e.Get(1)
_ = Z934
__e.TailApply(PrimFunc(symshen_4partial), Z934)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp9865, V933)
return


} else {
__e.Return(V933)
return
}


}


}, 1)

tmp9883 := Call(__e, ns2_1set, symshen_4partial, tmp9859)


_ = tmp9883

tmp9884 := MakeNative(func(__e *ControlFlow) {
V937 := __e.Get(1)
_ = V937
V938 := __e.Get(2)
_ = V938
tmp9898 := PrimEqual(Nil, V937)

if True == tmp9898 {
__e.Return(PrimCloseStream(V938))
return
} else {
tmp9896 := PrimIsPair(V937)

var ifres9892 Obj

if True == tmp9896 {
tmp9894 := PrimHead(V937)

tmp9895 := PrimIsPair(tmp9894)

var ifres9893 Obj

if True == tmp9895 {
ifres9893 = True


} else {
ifres9893 = False


}

ifres9892 = ifres9893


} else {
ifres9892 = False


}

if True == ifres9892 {
tmp9885 := PrimTail(V937)

tmp9886 := PrimHead(V937)

tmp9887 := Call(__e, PrimFunc(symshen_4write_1kl_1h), tmp9886, V938)


_ = tmp9887

__e.TailApply(PrimFunc(symshen_4write_1kl), tmp9885, V938)
return


} else {
tmp9890 := PrimIsPair(V937)

if True == tmp9890 {
tmp9888 := PrimTail(V937)

__e.TailApply(PrimFunc(symshen_4write_1kl), tmp9888, V938)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.write-kl")))
return
}


}


}


}, 2)

tmp9899 := Call(__e, ns2_1set, symshen_4write_1kl, tmp9884)


_ = tmp9899

tmp9900 := MakeNative(func(__e *ControlFlow) {
V941 := __e.Get(1)
_ = V941
V942 := __e.Get(2)
_ = V942
tmp9940 := PrimIsPair(V941)

var ifres9903 Obj

if True == tmp9940 {
tmp9938 := PrimHead(V941)

tmp9939 := PrimEqual(symdefun, tmp9938)

var ifres9905 Obj

if True == tmp9939 {
tmp9936 := PrimTail(V941)

tmp9937 := PrimIsPair(tmp9936)

var ifres9907 Obj

if True == tmp9937 {
tmp9933 := PrimTail(V941)

tmp9934 := PrimHead(tmp9933)

tmp9935 := PrimEqual(symfail, tmp9934)

var ifres9909 Obj

if True == tmp9935 {
tmp9930 := PrimTail(V941)

tmp9931 := PrimTail(tmp9930)

tmp9932 := PrimIsPair(tmp9931)

var ifres9911 Obj

if True == tmp9932 {
tmp9926 := PrimTail(V941)

tmp9927 := PrimTail(tmp9926)

tmp9928 := PrimHead(tmp9927)

tmp9929 := PrimEqual(Nil, tmp9928)

var ifres9913 Obj

if True == tmp9929 {
tmp9922 := PrimTail(V941)

tmp9923 := PrimTail(tmp9922)

tmp9924 := PrimTail(tmp9923)

tmp9925 := PrimIsPair(tmp9924)

var ifres9915 Obj

if True == tmp9925 {
tmp9917 := PrimTail(V941)

tmp9918 := PrimTail(tmp9917)

tmp9919 := PrimTail(tmp9918)

tmp9920 := PrimTail(tmp9919)

tmp9921 := PrimEqual(Nil, tmp9920)

var ifres9916 Obj

if True == tmp9921 {
ifres9916 = True


} else {
ifres9916 = False


}

ifres9915 = ifres9916


} else {
ifres9915 = False


}

var ifres9914 Obj

if True == ifres9915 {
ifres9914 = True


} else {
ifres9914 = False


}

ifres9913 = ifres9914


} else {
ifres9913 = False


}

var ifres9912 Obj

if True == ifres9913 {
ifres9912 = True


} else {
ifres9912 = False


}

ifres9911 = ifres9912


} else {
ifres9911 = False


}

var ifres9910 Obj

if True == ifres9911 {
ifres9910 = True


} else {
ifres9910 = False


}

ifres9909 = ifres9910


} else {
ifres9909 = False


}

var ifres9908 Obj

if True == ifres9909 {
ifres9908 = True


} else {
ifres9908 = False


}

ifres9907 = ifres9908


} else {
ifres9907 = False


}

var ifres9906 Obj

if True == ifres9907 {
ifres9906 = True


} else {
ifres9906 = False


}

ifres9905 = ifres9906


} else {
ifres9905 = False


}

var ifres9904 Obj

if True == ifres9905 {
ifres9904 = True


} else {
ifres9904 = False


}

ifres9903 = ifres9904


} else {
ifres9903 = False


}

if True == ifres9903 {
__e.TailApply(PrimFunc(sympr), MakeString("(defun fail () shen.fail!)"), V942)
return
} else {
tmp9901 := Call(__e, PrimFunc(symshen_4app), V941, MakeString("\n\n"), symshen_4r)


__e.TailApply(PrimFunc(sympr), tmp9901, V942)
return


}


}, 2)

tmp9941 := Call(__e, ns2_1set, symshen_4write_1kl_1h, tmp9900)


_ = tmp9941

tmp9942 := MakeNative(func(__e *ControlFlow) {
V943 := __e.Get(1)
_ = V943
tmp9951 := PrimEqual(MakeString(""), V943)

if True == tmp9951 {
__e.Return(MakeString(".kl"))
return
} else {
tmp9949 := PrimEqual(MakeString(".shen"), V943)

if True == tmp9949 {
__e.Return(MakeString(".kl"))
return
} else {
tmp9947 := Call(__e, PrimFunc(symshen_4_7string_2), V943)


if True == tmp9947 {
tmp9943 := Call(__e, PrimFunc(symhdstr), V943)


tmp9944 := PrimTailString(V943)

tmp9945 := Call(__e, PrimFunc(symshen_4klfile), tmp9944)


__e.TailApply(PrimFunc(sym_8s), tmp9943, tmp9945)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.klfile")))
return
}


}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4klfile, tmp9942)
return




}, 0)

