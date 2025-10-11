package main

import . "github.com/tiancaiamao/shen-go/kl"

var DeclarationsMain = MakeNative(func(__e *ControlFlow) {
tmp7580 := PrimSet(symshen_4_dhistory_d, Nil)

_ = tmp7580

tmp7581 := PrimSet(symshen_4_dtc_d, False)

_ = tmp7581

tmp7582 := Call(__e, PrimFunc(symvector), MakeNumber(20000))


tmp7583 := PrimSet(sym_dproperty_1vector_d, tmp7582)

_ = tmp7583

tmp7584 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4macros), X)
return
}, 1)

tmp7585 := PrimCons(symshen_4macros, tmp7584)

tmp7586 := PrimCons(tmp7585, Nil)

tmp7587 := PrimSet(sym_dmacros_d, tmp7586)

_ = tmp7587

tmp7588 := PrimSet(symshen_4_dgensym_d, MakeNumber(0))

_ = tmp7588

tmp7589 := PrimSet(symshen_4_dtracking_d, Nil)

_ = tmp7589

tmp7590 := PrimSet(symshen_4_dprofiled_d, Nil)

_ = tmp7590

tmp7591 := PrimSet(sym_dhome_1directory_d, MakeString(""))

_ = tmp7591

tmp7592 := PrimCons(symtype, Nil)

tmp7593 := PrimCons(syminput_7, tmp7592)

tmp7594 := PrimCons(symopen, tmp7593)

tmp7595 := PrimCons(symset, tmp7594)

tmp7596 := PrimCons(symwhere, tmp7595)

tmp7597 := PrimCons(symlet, tmp7596)

tmp7598 := PrimCons(symlambda, tmp7597)

tmp7599 := PrimCons(symcons, tmp7598)

tmp7600 := PrimCons(sym_8v, tmp7599)

tmp7601 := PrimCons(sym_8s, tmp7600)

tmp7602 := PrimCons(sym_8p, tmp7601)

tmp7603 := PrimSet(symshen_4_dspecial_d, tmp7602)

_ = tmp7603

tmp7604 := PrimSet(symshen_4_dextraspecial_d, Nil)

_ = tmp7604

tmp7605 := PrimSet(symshen_4_dspy_d, False)

_ = tmp7605

tmp7606 := PrimSet(symshen_4_ddatatypes_d, Nil)

_ = tmp7606

tmp7607 := PrimSet(symshen_4_dalldatatypes_d, Nil)

_ = tmp7607

tmp7608 := PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True)

_ = tmp7608

tmp7609 := PrimSet(symshen_4_dpackage_d, symnull)

_ = tmp7609

tmp7610 := PrimSet(symshen_4_dsynonyms_d, Nil)

_ = tmp7610

tmp7611 := PrimSet(symshen_4_dsystem_d, Nil)

_ = tmp7611

tmp7612 := PrimSet(symshen_4_dsigf_d, Nil)

_ = tmp7612

tmp7613 := PrimSet(symshen_4_doccurs_d, True)

_ = tmp7613

tmp7614 := PrimSet(symshen_4_dfactorise_2_d, False)

_ = tmp7614

tmp7615 := PrimSet(symshen_4_dmaxinferences_d, MakeNumber(1000000))

_ = tmp7615

tmp7616 := PrimSet(sym_dmaximum_1print_1sequence_1size_d, MakeNumber(20))

_ = tmp7616

tmp7617 := PrimSet(symshen_4_dcall_d, MakeNumber(0))

_ = tmp7617

tmp7618 := PrimSet(symshen_4_dinfs_d, MakeNumber(0))

_ = tmp7618

tmp7619 := PrimSet(sym_dhush_d, False)

_ = tmp7619

tmp7620 := PrimSet(symshen_4_doptimise_d, False)

_ = tmp7620

tmp7621 := PrimSet(sym_dversion_d, MakeString("39.2"))

_ = tmp7621

tmp7622 := PrimSet(symshen_4_dnames_d, Nil)

_ = tmp7622

tmp7623 := PrimSet(symshen_4_dstep_d, False)

_ = tmp7623

tmp7624 := PrimSet(symshen_4_dit_d, MakeString(""))

_ = tmp7624

tmp7625 := PrimSet(symshen_4_dresidue_d, Nil)

_ = tmp7625

tmp7626 := PrimSet(symshen_4_dprolog_1memory_d, MakeNumber(1000))

_ = tmp7626

tmp7627 := PrimSet(symshen_4_dloading_2_d, False)

_ = tmp7627

tmp7628 := PrimSet(symshen_4_duserdefs_d, Nil)

_ = tmp7628

tmp7629 := MakeNative(func(__e *ControlFlow) {
tmp7630 := MakeNative(func(__e *ControlFlow) {
Z5780 := __e.Get(1)
_ = Z5780
__e.TailApply(PrimFunc(symshen_4typename), Z5780)
return
}, 1)

tmp7631 := PrimValue(symshen_4_dalldatatypes_d)

__e.TailApply(PrimFunc(symmap), tmp7630, tmp7631)
return


}, 0)

tmp7632 := Call(__e, ns2_1set, symdatatypes, tmp7629)


_ = tmp7632

tmp7633 := MakeNative(func(__e *ControlFlow) {
tmp7634 := MakeNative(func(__e *ControlFlow) {
Z5781 := __e.Get(1)
_ = Z5781
__e.TailApply(PrimFunc(symshen_4typename), Z5781)
return
}, 1)

tmp7635 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(PrimFunc(symmap), tmp7634, tmp7635)
return


}, 0)

tmp7636 := Call(__e, ns2_1set, symincluded, tmp7633)


_ = tmp7636

tmp7637 := MakeNative(func(__e *ControlFlow) {
V5784 := __e.Get(1)
_ = V5784
tmp7642 := PrimIsPair(V5784)

if True == tmp7642 {
tmp7638 := PrimHead(V5784)

tmp7639 := PrimStr(tmp7638)

tmp7640 := Call(__e, PrimFunc(symshen_4typename_1h), tmp7639)


__e.Return(PrimIntern(tmp7640))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.typename")))
return
}


}, 1)

tmp7643 := Call(__e, ns2_1set, symshen_4typename, tmp7637)


_ = tmp7643

tmp7644 := MakeNative(func(__e *ControlFlow) {
V5785 := __e.Get(1)
_ = V5785
tmp7651 := PrimEqual(MakeString("#type"), V5785)

if True == tmp7651 {
__e.Return(MakeString(""))
return
} else {
tmp7649 := Call(__e, PrimFunc(symshen_4_7string_2), V5785)


if True == tmp7649 {
tmp7645 := Call(__e, PrimFunc(symhdstr), V5785)


tmp7646 := PrimTailString(V5785)

tmp7647 := Call(__e, PrimFunc(symshen_4typename_1h), tmp7646)


__e.Return(PrimStringConcat(tmp7645, tmp7647))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.typename-h")))
return
}


}


}, 1)

tmp7652 := Call(__e, ns2_1set, symshen_4typename_1h, tmp7644)


_ = tmp7652

tmp7653 := MakeNative(func(__e *ControlFlow) {
V5786 := __e.Get(1)
_ = V5786
tmp7657 := PrimLessThan(V5786, MakeNumber(0))

if True == tmp7657 {
__e.Return(PrimValue(symshen_4_dprolog_1memory_d))
return
} else {
tmp7655 := PrimIsInteger(V5786)

if True == tmp7655 {
__e.Return(PrimSet(symshen_4_dprolog_1memory_d, V5786))
return
} else {
__e.Return(PrimSimpleError(MakeString("prolog memory expects an integer value\n")))
return
}


}


}, 1)

tmp7658 := Call(__e, ns2_1set, symprolog_1memory, tmp7653)


_ = tmp7658

tmp7659 := MakeNative(func(__e *ControlFlow) {
V5789 := __e.Get(1)
_ = V5789
tmp7675 := PrimEqual(Nil, V5789)

if True == tmp7675 {
__e.Return(Nil)
return
} else {
tmp7673 := PrimIsPair(V5789)

var ifres7669 Obj

if True == tmp7673 {
tmp7671 := PrimTail(V5789)

tmp7672 := PrimIsPair(tmp7671)

var ifres7670 Obj

if True == tmp7672 {
ifres7670 = True


} else {
ifres7670 = False


}

ifres7669 = ifres7670


} else {
ifres7669 = False


}

if True == ifres7669 {
tmp7660 := MakeNative(func(__e *ControlFlow) {
W5790 := __e.Get(1)
_ = W5790
tmp7661 := PrimTail(V5789)

tmp7662 := PrimTail(tmp7661)

__e.TailApply(PrimFunc(symshen_4initialise_1arity_1table), tmp7662)
return


}, 1)

tmp7663 := PrimHead(V5789)

tmp7664 := PrimTail(V5789)

tmp7665 := PrimHead(tmp7664)

tmp7666 := PrimValue(sym_dproperty_1vector_d)

tmp7667 := Call(__e, PrimFunc(symput), tmp7663, symarity, tmp7665, tmp7666)


__e.TailApply(tmp7660, tmp7667)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.initialise-arity-table")))
return
}


}


}, 1)

tmp7676 := Call(__e, ns2_1set, symshen_4initialise_1lambda_1tables, tmp7659)


_ = tmp7676

tmp7677 := MakeNative(func(__e *ControlFlow) {
V5791 := __e.Get(1)
_ = V5791
tmp7678 := MakeNative(func(__e *ControlFlow) {
tmp7679 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V5791, symarity, tmp7679)
return


}, 0)

tmp7680 := MakeNative(func(__e *ControlFlow) {
Z5792 := __e.Get(1)
_ = Z5792
__e.Return(MakeNumber(-1))
return
}, 1)

__e.TailApply(try_1catch, tmp7678, tmp7680)
return


}, 1)

tmp7681 := Call(__e, ns2_1set, symarity, tmp7677)


_ = tmp7681

tmp7682 := MakeNative(func(__e *ControlFlow) {
V5795 := __e.Get(1)
_ = V5795
tmp7698 := PrimEqual(Nil, V5795)

if True == tmp7698 {
__e.Return(Nil)
return
} else {
tmp7696 := PrimIsPair(V5795)

var ifres7692 Obj

if True == tmp7696 {
tmp7694 := PrimTail(V5795)

tmp7695 := PrimIsPair(tmp7694)

var ifres7693 Obj

if True == tmp7695 {
ifres7693 = True


} else {
ifres7693 = False


}

ifres7692 = ifres7693


} else {
ifres7692 = False


}

if True == ifres7692 {
tmp7683 := MakeNative(func(__e *ControlFlow) {
W5796 := __e.Get(1)
_ = W5796
tmp7684 := PrimTail(V5795)

tmp7685 := PrimTail(tmp7684)

__e.TailApply(PrimFunc(symshen_4initialise_1arity_1table), tmp7685)
return


}, 1)

tmp7686 := PrimHead(V5795)

tmp7687 := PrimTail(V5795)

tmp7688 := PrimHead(tmp7687)

tmp7689 := PrimValue(sym_dproperty_1vector_d)

tmp7690 := Call(__e, PrimFunc(symput), tmp7686, symarity, tmp7688, tmp7689)


__e.TailApply(tmp7683, tmp7690)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.initialise-arity-table")))
return
}


}


}, 1)

tmp7699 := Call(__e, ns2_1set, symshen_4initialise_1arity_1table, tmp7682)


_ = tmp7699

tmp7700 := PrimCons(MakeNumber(2), Nil)

tmp7701 := PrimCons(sym_8s, tmp7700)

tmp7702 := PrimCons(MakeNumber(2), tmp7701)

tmp7703 := PrimCons(sym_8v, tmp7702)

tmp7704 := PrimCons(MakeNumber(2), tmp7703)

tmp7705 := PrimCons(sym_8p, tmp7704)

tmp7706 := PrimCons(MakeNumber(1), tmp7705)

tmp7707 := PrimCons(sym_5_b_6, tmp7706)

tmp7708 := PrimCons(MakeNumber(1), tmp7707)

tmp7709 := PrimCons(sym_5end_6, tmp7708)

tmp7710 := PrimCons(MakeNumber(1), tmp7709)

tmp7711 := PrimCons(sym_5e_6, tmp7710)

tmp7712 := PrimCons(MakeNumber(2), tmp7711)

tmp7713 := PrimCons(sym_a_a, tmp7712)

tmp7714 := PrimCons(MakeNumber(2), tmp7713)

tmp7715 := PrimCons(sym_1, tmp7714)

tmp7716 := PrimCons(MakeNumber(2), tmp7715)

tmp7717 := PrimCons(sym_c, tmp7716)

tmp7718 := PrimCons(MakeNumber(2), tmp7717)

tmp7719 := PrimCons(sym_d, tmp7718)

tmp7720 := PrimCons(MakeNumber(2), tmp7719)

tmp7721 := PrimCons(sym_7, tmp7720)

tmp7722 := PrimCons(MakeNumber(1), tmp7721)

tmp7723 := PrimCons(symy_1or_1n_2, tmp7722)

tmp7724 := PrimCons(MakeNumber(2), tmp7723)

tmp7725 := PrimCons(symwrite_1to_1file, tmp7724)

tmp7726 := PrimCons(MakeNumber(2), tmp7725)

tmp7727 := PrimCons(symwrite_1byte, tmp7726)

tmp7728 := PrimCons(MakeNumber(5), tmp7727)

tmp7729 := PrimCons(symwhen, tmp7728)

tmp7730 := PrimCons(MakeNumber(0), tmp7729)

tmp7731 := PrimCons(symversion, tmp7730)

tmp7732 := PrimCons(MakeNumber(5), tmp7731)

tmp7733 := PrimCons(symvar_2, tmp7732)

tmp7734 := PrimCons(MakeNumber(1), tmp7733)

tmp7735 := PrimCons(symvariable_2, tmp7734)

tmp7736 := PrimCons(MakeNumber(1), tmp7735)

tmp7737 := PrimCons(symvalue, tmp7736)

tmp7738 := PrimCons(MakeNumber(3), tmp7737)

tmp7739 := PrimCons(symvector_1_6, tmp7738)

tmp7740 := PrimCons(MakeNumber(1), tmp7739)

tmp7741 := PrimCons(symvector_2, tmp7740)

tmp7742 := PrimCons(MakeNumber(1), tmp7741)

tmp7743 := PrimCons(symvector, tmp7742)

tmp7744 := PrimCons(MakeNumber(0), tmp7743)

tmp7745 := PrimCons(symuserdefs, tmp7744)

tmp7746 := PrimCons(MakeNumber(2), tmp7745)

tmp7747 := PrimCons(symupdate_1lambda_1table, tmp7746)

tmp7748 := PrimCons(MakeNumber(1), tmp7747)

tmp7749 := PrimCons(symundefmacro, tmp7748)

tmp7750 := PrimCons(MakeNumber(1), tmp7749)

tmp7751 := PrimCons(symuntrack, tmp7750)

tmp7752 := PrimCons(MakeNumber(2), tmp7751)

tmp7753 := PrimCons(symunion, tmp7752)

tmp7754 := PrimCons(MakeNumber(1), tmp7753)

tmp7755 := PrimCons(symunprofile, tmp7754)

tmp7756 := PrimCons(MakeNumber(3), tmp7755)

tmp7757 := PrimCons(symunput, tmp7756)

tmp7758 := PrimCons(MakeNumber(1), tmp7757)

tmp7759 := PrimCons(symundefmacro, tmp7758)

tmp7760 := PrimCons(MakeNumber(5), tmp7759)

tmp7761 := PrimCons(symreturn, tmp7760)

tmp7762 := PrimCons(MakeNumber(2), tmp7761)

tmp7763 := PrimCons(symtype, tmp7762)

tmp7764 := PrimCons(MakeNumber(1), tmp7763)

tmp7765 := PrimCons(symtuple_2, tmp7764)

tmp7766 := PrimCons(MakeNumber(2), tmp7765)

tmp7767 := PrimCons(symtrap_1error, tmp7766)

tmp7768 := PrimCons(MakeNumber(0), tmp7767)

tmp7769 := PrimCons(symtracked, tmp7768)

tmp7770 := PrimCons(MakeNumber(1), tmp7769)

tmp7771 := PrimCons(symtrack, tmp7770)

tmp7772 := PrimCons(MakeNumber(1), tmp7771)

tmp7773 := PrimCons(symtlstr, tmp7772)

tmp7774 := PrimCons(MakeNumber(1), tmp7773)

tmp7775 := PrimCons(symthaw, tmp7774)

tmp7776 := PrimCons(MakeNumber(0), tmp7775)

tmp7777 := PrimCons(symtc_2, tmp7776)

tmp7778 := PrimCons(MakeNumber(1), tmp7777)

tmp7779 := PrimCons(symtc, tmp7778)

tmp7780 := PrimCons(MakeNumber(1), tmp7779)

tmp7781 := PrimCons(symtl, tmp7780)

tmp7782 := PrimCons(MakeNumber(1), tmp7781)

tmp7783 := PrimCons(symtail, tmp7782)

tmp7784 := PrimCons(MakeNumber(1), tmp7783)

tmp7785 := PrimCons(symsystemf, tmp7784)

tmp7786 := PrimCons(MakeNumber(1), tmp7785)

tmp7787 := PrimCons(symsymbol_2, tmp7786)

tmp7788 := PrimCons(MakeNumber(1), tmp7787)

tmp7789 := PrimCons(symsum, tmp7788)

tmp7790 := PrimCons(MakeNumber(3), tmp7789)

tmp7791 := PrimCons(symsubst, tmp7790)

tmp7792 := PrimCons(MakeNumber(1), tmp7791)

tmp7793 := PrimCons(symstring_2, tmp7792)

tmp7794 := PrimCons(MakeNumber(1), tmp7793)

tmp7795 := PrimCons(symstring_1_6symbol, tmp7794)

tmp7796 := PrimCons(MakeNumber(1), tmp7795)

tmp7797 := PrimCons(symstring_1_6n, tmp7796)

tmp7798 := PrimCons(MakeNumber(1), tmp7797)

tmp7799 := PrimCons(symstr, tmp7798)

tmp7800 := PrimCons(MakeNumber(0), tmp7799)

tmp7801 := PrimCons(symstoutput, tmp7800)

tmp7802 := PrimCons(MakeNumber(0), tmp7801)

tmp7803 := PrimCons(symstinput, tmp7802)

tmp7804 := PrimCons(MakeNumber(0), tmp7803)

tmp7805 := PrimCons(symstep_2, tmp7804)

tmp7806 := PrimCons(MakeNumber(1), tmp7805)

tmp7807 := PrimCons(symstep, tmp7806)

tmp7808 := PrimCons(MakeNumber(0), tmp7807)

tmp7809 := PrimCons(symspy_2, tmp7808)

tmp7810 := PrimCons(MakeNumber(1), tmp7809)

tmp7811 := PrimCons(symspy, tmp7810)

tmp7812 := PrimCons(MakeNumber(2), tmp7811)

tmp7813 := PrimCons(symspecialise, tmp7812)

tmp7814 := PrimCons(MakeNumber(1), tmp7813)

tmp7815 := PrimCons(symsnd, tmp7814)

tmp7816 := PrimCons(MakeNumber(1), tmp7815)

tmp7817 := PrimCons(symsimple_1error, tmp7816)

tmp7818 := PrimCons(MakeNumber(2), tmp7817)

tmp7819 := PrimCons(symset, tmp7818)

tmp7820 := PrimCons(MakeNumber(1), tmp7819)

tmp7821 := PrimCons(symreverse, tmp7820)

tmp7822 := PrimCons(MakeNumber(2), tmp7821)

tmp7823 := PrimCons(symremove, tmp7822)

tmp7824 := PrimCons(MakeNumber(0), tmp7823)

tmp7825 := PrimCons(symrelease, tmp7824)

tmp7826 := PrimCons(MakeNumber(1), tmp7825)

tmp7827 := PrimCons(symreceive, tmp7826)

tmp7828 := PrimCons(MakeNumber(1), tmp7827)

tmp7829 := PrimCons(symshen_4read_1unit_1string, tmp7828)

tmp7830 := PrimCons(MakeNumber(1), tmp7829)

tmp7831 := PrimCons(symread_1from_1string_1unprocessed, tmp7830)

tmp7832 := PrimCons(MakeNumber(1), tmp7831)

tmp7833 := PrimCons(symread_1from_1string, tmp7832)

tmp7834 := PrimCons(MakeNumber(1), tmp7833)

tmp7835 := PrimCons(symread_1byte, tmp7834)

tmp7836 := PrimCons(MakeNumber(1), tmp7835)

tmp7837 := PrimCons(symread, tmp7836)

tmp7838 := PrimCons(MakeNumber(1), tmp7837)

tmp7839 := PrimCons(symread_1file, tmp7838)

tmp7840 := PrimCons(MakeNumber(1), tmp7839)

tmp7841 := PrimCons(symread_1file_1as_1bytelist, tmp7840)

tmp7842 := PrimCons(MakeNumber(1), tmp7841)

tmp7843 := PrimCons(symread_1file_1as_1string, tmp7842)

tmp7844 := PrimCons(MakeNumber(4), tmp7843)

tmp7845 := PrimCons(symput, tmp7844)

tmp7846 := PrimCons(MakeNumber(1), tmp7845)

tmp7847 := PrimCons(symprotect, tmp7846)

tmp7848 := PrimCons(MakeNumber(1), tmp7847)

tmp7849 := PrimCons(sympreclude_1all_1but, tmp7848)

tmp7850 := PrimCons(MakeNumber(1), tmp7849)

tmp7851 := PrimCons(sympreclude, tmp7850)

tmp7852 := PrimCons(MakeNumber(1), tmp7851)

tmp7853 := PrimCons(symps, tmp7852)

tmp7854 := PrimCons(MakeNumber(2), tmp7853)

tmp7855 := PrimCons(sympr, tmp7854)

tmp7856 := PrimCons(MakeNumber(1), tmp7855)

tmp7857 := PrimCons(symprofile_1results, tmp7856)

tmp7858 := PrimCons(MakeNumber(1), tmp7857)

tmp7859 := PrimCons(symprolog_1memory, tmp7858)

tmp7860 := PrimCons(MakeNumber(1), tmp7859)

tmp7861 := PrimCons(symshen_4printF, tmp7860)

tmp7862 := PrimCons(MakeNumber(1), tmp7861)

tmp7863 := PrimCons(symshen_4print_1freshterm, tmp7862)

tmp7864 := PrimCons(MakeNumber(1), tmp7863)

tmp7865 := PrimCons(symshen_4print_1prolog_1vector, tmp7864)

tmp7866 := PrimCons(MakeNumber(1), tmp7865)

tmp7867 := PrimCons(symprofile, tmp7866)

tmp7868 := PrimCons(MakeNumber(1), tmp7867)

tmp7869 := PrimCons(symprint, tmp7868)

tmp7870 := PrimCons(MakeNumber(1), tmp7869)

tmp7871 := PrimCons(sympreclude_1all_1but, tmp7870)

tmp7872 := PrimCons(MakeNumber(2), tmp7871)

tmp7873 := PrimCons(sympos, tmp7872)

tmp7874 := PrimCons(MakeNumber(0), tmp7873)

tmp7875 := PrimCons(symporters, tmp7874)

tmp7876 := PrimCons(MakeNumber(0), tmp7875)

tmp7877 := PrimCons(symport, tmp7876)

tmp7878 := PrimCons(MakeNumber(1), tmp7877)

tmp7879 := PrimCons(sympackage_2, tmp7878)

tmp7880 := PrimCons(MakeNumber(3), tmp7879)

tmp7881 := PrimCons(sympackage, tmp7880)

tmp7882 := PrimCons(MakeNumber(0), tmp7881)

tmp7883 := PrimCons(symos, tmp7882)

tmp7884 := PrimCons(MakeNumber(2), tmp7883)

tmp7885 := PrimCons(symor, tmp7884)

tmp7886 := PrimCons(MakeNumber(0), tmp7885)

tmp7887 := PrimCons(symoptimise_2, tmp7886)

tmp7888 := PrimCons(MakeNumber(1), tmp7887)

tmp7889 := PrimCons(symoptimise, tmp7888)

tmp7890 := PrimCons(MakeNumber(2), tmp7889)

tmp7891 := PrimCons(symopen, tmp7890)

tmp7892 := PrimCons(MakeNumber(1), tmp7891)

tmp7893 := PrimCons(symoccurs_1check, tmp7892)

tmp7894 := PrimCons(MakeNumber(0), tmp7893)

tmp7895 := PrimCons(symoccurs_2, tmp7894)

tmp7896 := PrimCons(MakeNumber(2), tmp7895)

tmp7897 := PrimCons(symoccurrences, tmp7896)

tmp7898 := PrimCons(MakeNumber(1), tmp7897)

tmp7899 := PrimCons(symoccurs_1check, tmp7898)

tmp7900 := PrimCons(MakeNumber(1), tmp7899)

tmp7901 := PrimCons(symnumber_2, tmp7900)

tmp7902 := PrimCons(MakeNumber(1), tmp7901)

tmp7903 := PrimCons(symn_1_6string, tmp7902)

tmp7904 := PrimCons(MakeNumber(2), tmp7903)

tmp7905 := PrimCons(symnth, tmp7904)

tmp7906 := PrimCons(MakeNumber(1), tmp7905)

tmp7907 := PrimCons(symnot, tmp7906)

tmp7908 := PrimCons(MakeNumber(1), tmp7907)

tmp7909 := PrimCons(symnl, tmp7908)

tmp7910 := PrimCons(MakeNumber(1), tmp7909)

tmp7911 := PrimCons(symmaxinferences, tmp7910)

tmp7912 := PrimCons(MakeNumber(2), tmp7911)

tmp7913 := PrimCons(symmapcan, tmp7912)

tmp7914 := PrimCons(MakeNumber(2), tmp7913)

tmp7915 := PrimCons(symmap, tmp7914)

tmp7916 := PrimCons(MakeNumber(1), tmp7915)

tmp7917 := PrimCons(symmacroexpand, tmp7916)

tmp7918 := PrimCons(MakeNumber(1), tmp7917)

tmp7919 := PrimCons(symvector, tmp7918)

tmp7920 := PrimCons(MakeNumber(2), tmp7919)

tmp7921 := PrimCons(sym_5_a, tmp7920)

tmp7922 := PrimCons(MakeNumber(2), tmp7921)

tmp7923 := PrimCons(sym_5, tmp7922)

tmp7924 := PrimCons(MakeNumber(1), tmp7923)

tmp7925 := PrimCons(symload, tmp7924)

tmp7926 := PrimCons(MakeNumber(1), tmp7925)

tmp7927 := PrimCons(symlist, tmp7926)

tmp7928 := PrimCons(MakeNumber(1), tmp7927)

tmp7929 := PrimCons(symlineread, tmp7928)

tmp7930 := PrimCons(MakeNumber(1), tmp7929)

tmp7931 := PrimCons(symlimit, tmp7930)

tmp7932 := PrimCons(MakeNumber(1), tmp7931)

tmp7933 := PrimCons(symlength, tmp7932)

tmp7934 := PrimCons(MakeNumber(0), tmp7933)

tmp7935 := PrimCons(symlanguage, tmp7934)

tmp7936 := PrimCons(MakeNumber(6), tmp7935)

tmp7937 := PrimCons(symis_b, tmp7936)

tmp7938 := PrimCons(MakeNumber(6), tmp7937)

tmp7939 := PrimCons(symis, tmp7938)

tmp7940 := PrimCons(MakeNumber(0), tmp7939)

tmp7941 := PrimCons(symit, tmp7940)

tmp7942 := PrimCons(MakeNumber(1), tmp7941)

tmp7943 := PrimCons(syminternal, tmp7942)

tmp7944 := PrimCons(MakeNumber(2), tmp7943)

tmp7945 := PrimCons(symintersection, tmp7944)

tmp7946 := PrimCons(MakeNumber(1), tmp7945)

tmp7947 := PrimCons(syminclude_1all_1but, tmp7946)

tmp7948 := PrimCons(MakeNumber(0), tmp7947)

tmp7949 := PrimCons(symimplementation, tmp7948)

tmp7950 := PrimCons(MakeNumber(2), tmp7949)

tmp7951 := PrimCons(syminput_7, tmp7950)

tmp7952 := PrimCons(MakeNumber(1), tmp7951)

tmp7953 := PrimCons(syminput, tmp7952)

tmp7954 := PrimCons(MakeNumber(0), tmp7953)

tmp7955 := PrimCons(syminferences, tmp7954)

tmp7956 := PrimCons(MakeNumber(1), tmp7955)

tmp7957 := PrimCons(symintern, tmp7956)

tmp7958 := PrimCons(MakeNumber(1), tmp7957)

tmp7959 := PrimCons(syminternal, tmp7958)

tmp7960 := PrimCons(MakeNumber(1), tmp7959)

tmp7961 := PrimCons(syminteger_2, tmp7960)

tmp7962 := PrimCons(MakeNumber(1), tmp7961)

tmp7963 := PrimCons(symin_1package, tmp7962)

tmp7964 := PrimCons(MakeNumber(0), tmp7963)

tmp7965 := PrimCons(symincluded, tmp7964)

tmp7966 := PrimCons(MakeNumber(1), tmp7965)

tmp7967 := PrimCons(syminclude, tmp7966)

tmp7968 := PrimCons(MakeNumber(3), tmp7967)

tmp7969 := PrimCons(symif, tmp7968)

tmp7970 := PrimCons(MakeNumber(1), tmp7969)

tmp7971 := PrimCons(symhush, tmp7970)

tmp7972 := PrimCons(MakeNumber(0), tmp7971)

tmp7973 := PrimCons(symhush_2, tmp7972)

tmp7974 := PrimCons(MakeNumber(1), tmp7973)

tmp7975 := PrimCons(symhead, tmp7974)

tmp7976 := PrimCons(MakeNumber(1), tmp7975)

tmp7977 := PrimCons(symhdstr, tmp7976)

tmp7978 := PrimCons(MakeNumber(1), tmp7977)

tmp7979 := PrimCons(symhdv, tmp7978)

tmp7980 := PrimCons(MakeNumber(1), tmp7979)

tmp7981 := PrimCons(symhd, tmp7980)

tmp7982 := PrimCons(MakeNumber(2), tmp7981)

tmp7983 := PrimCons(symhash, tmp7982)

tmp7984 := PrimCons(MakeNumber(2), tmp7983)

tmp7985 := PrimCons(sym_a, tmp7984)

tmp7986 := PrimCons(MakeNumber(2), tmp7985)

tmp7987 := PrimCons(sym_6_a, tmp7986)

tmp7988 := PrimCons(MakeNumber(2), tmp7987)

tmp7989 := PrimCons(sym_6, tmp7988)

tmp7990 := PrimCons(MakeNumber(2), tmp7989)

tmp7991 := PrimCons(sym_5_1vector, tmp7990)

tmp7992 := PrimCons(MakeNumber(2), tmp7991)

tmp7993 := PrimCons(sym_5_1address, tmp7992)

tmp7994 := PrimCons(MakeNumber(3), tmp7993)

tmp7995 := PrimCons(symaddress_1_6, tmp7994)

tmp7996 := PrimCons(MakeNumber(1), tmp7995)

tmp7997 := PrimCons(symget_1time, tmp7996)

tmp7998 := PrimCons(MakeNumber(3), tmp7997)

tmp7999 := PrimCons(symget, tmp7998)

tmp8000 := PrimCons(MakeNumber(1), tmp7999)

tmp8001 := PrimCons(symgensym, tmp8000)

tmp8002 := PrimCons(MakeNumber(1), tmp8001)

tmp8003 := PrimCons(symfunction, tmp8002)

tmp8004 := PrimCons(MakeNumber(1), tmp8003)

tmp8005 := PrimCons(symfn, tmp8004)

tmp8006 := PrimCons(MakeNumber(1), tmp8005)

tmp8007 := PrimCons(symfst, tmp8006)

tmp8008 := PrimCons(MakeNumber(0), tmp8007)

tmp8009 := PrimCons(symfresh, tmp8008)

tmp8010 := PrimCons(MakeNumber(1), tmp8009)

tmp8011 := PrimCons(symfreeze, tmp8010)

tmp8012 := PrimCons(MakeNumber(5), tmp8011)

tmp8013 := PrimCons(symfork, tmp8012)

tmp8014 := PrimCons(MakeNumber(1), tmp8013)

tmp8015 := PrimCons(symforeign, tmp8014)

tmp8016 := PrimCons(MakeNumber(7), tmp8015)

tmp8017 := PrimCons(symfindall, tmp8016)

tmp8018 := PrimCons(MakeNumber(2), tmp8017)

tmp8019 := PrimCons(symfix, tmp8018)

tmp8020 := PrimCons(MakeNumber(0), tmp8019)

tmp8021 := PrimCons(symfail, tmp8020)

tmp8022 := PrimCons(MakeNumber(2), tmp8021)

tmp8023 := PrimCons(symfail_1if, tmp8022)

tmp8024 := PrimCons(MakeNumber(0), tmp8023)

tmp8025 := PrimCons(symfactorise_2, tmp8024)

tmp8026 := PrimCons(MakeNumber(1), tmp8025)

tmp8027 := PrimCons(symfactorise, tmp8026)

tmp8028 := PrimCons(MakeNumber(1), tmp8027)

tmp8029 := PrimCons(symexternal, tmp8028)

tmp8030 := PrimCons(MakeNumber(1), tmp8029)

tmp8031 := PrimCons(symexplode, tmp8030)

tmp8032 := PrimCons(MakeNumber(1), tmp8031)

tmp8033 := PrimCons(symeval_1kl, tmp8032)

tmp8034 := PrimCons(MakeNumber(1), tmp8033)

tmp8035 := PrimCons(symeval, tmp8034)

tmp8036 := PrimCons(MakeNumber(1), tmp8035)

tmp8037 := PrimCons(symerror_1to_1string, tmp8036)

tmp8038 := PrimCons(MakeNumber(1), tmp8037)

tmp8039 := PrimCons(symexternal, tmp8038)

tmp8040 := PrimCons(MakeNumber(1), tmp8039)

tmp8041 := PrimCons(symenable_1type_1theory, tmp8040)

tmp8042 := PrimCons(MakeNumber(1), tmp8041)

tmp8043 := PrimCons(symempty_2, tmp8042)

tmp8044 := PrimCons(MakeNumber(2), tmp8043)

tmp8045 := PrimCons(symelement_2, tmp8044)

tmp8046 := PrimCons(MakeNumber(2), tmp8045)

tmp8047 := PrimCons(symdo, tmp8046)

tmp8048 := PrimCons(MakeNumber(2), tmp8047)

tmp8049 := PrimCons(symdifference, tmp8048)

tmp8050 := PrimCons(MakeNumber(1), tmp8049)

tmp8051 := PrimCons(symdestroy, tmp8050)

tmp8052 := PrimCons(MakeNumber(2), tmp8051)

tmp8053 := PrimCons(symdeclare, tmp8052)

tmp8054 := PrimCons(MakeNumber(0), tmp8053)

tmp8055 := PrimCons(symdatatypes, tmp8054)

tmp8056 := PrimCons(MakeNumber(1), tmp8055)

tmp8057 := PrimCons(symclose, tmp8056)

tmp8058 := PrimCons(MakeNumber(2), tmp8057)

tmp8059 := PrimCons(symcn, tmp8058)

tmp8060 := PrimCons(MakeNumber(1), tmp8059)

tmp8061 := PrimCons(symcons_2, tmp8060)

tmp8062 := PrimCons(MakeNumber(2), tmp8061)

tmp8063 := PrimCons(symcons, tmp8062)

tmp8064 := PrimCons(MakeNumber(2), tmp8063)

tmp8065 := PrimCons(symconcat, tmp8064)

tmp8066 := PrimCons(MakeNumber(2), tmp8065)

tmp8067 := PrimCons(symcompile, tmp8066)

tmp8068 := PrimCons(MakeNumber(1), tmp8067)

tmp8069 := PrimCons(symcd, tmp8068)

tmp8070 := PrimCons(MakeNumber(5), tmp8069)

tmp8071 := PrimCons(symcall, tmp8070)

tmp8072 := PrimCons(MakeNumber(6), tmp8071)

tmp8073 := PrimCons(symbind, tmp8072)

tmp8074 := PrimCons(MakeNumber(1), tmp8073)

tmp8075 := PrimCons(symbound_2, tmp8074)

tmp8076 := PrimCons(MakeNumber(1), tmp8075)

tmp8077 := PrimCons(symbootstrap, tmp8076)

tmp8078 := PrimCons(MakeNumber(1), tmp8077)

tmp8079 := PrimCons(symboolean_2, tmp8078)

tmp8080 := PrimCons(MakeNumber(1), tmp8079)

tmp8081 := PrimCons(symatom_2, tmp8080)

tmp8082 := PrimCons(MakeNumber(2), tmp8081)

tmp8083 := PrimCons(symassoc, tmp8082)

tmp8084 := PrimCons(MakeNumber(1), tmp8083)

tmp8085 := PrimCons(symarity, tmp8084)

tmp8086 := PrimCons(MakeNumber(2), tmp8085)

tmp8087 := PrimCons(symappend, tmp8086)

tmp8088 := PrimCons(MakeNumber(2), tmp8087)

tmp8089 := PrimCons(symand, tmp8088)

tmp8090 := PrimCons(MakeNumber(2), tmp8089)

tmp8091 := PrimCons(symadjoin, tmp8090)

tmp8092 := PrimCons(MakeNumber(3), tmp8091)

tmp8093 := PrimCons(symaddress_1_6, tmp8092)

tmp8094 := PrimCons(MakeNumber(1), tmp8093)

tmp8095 := PrimCons(symabsvector, tmp8094)

tmp8096 := PrimCons(MakeNumber(1), tmp8095)

tmp8097 := PrimCons(symabsvector_2, tmp8096)

tmp8098 := PrimCons(MakeNumber(0), tmp8097)

tmp8099 := PrimCons(symabort, tmp8098)

tmp8100 := Call(__e, PrimFunc(symshen_4initialise_1arity_1table), tmp8099)


_ = tmp8100

tmp8101 := MakeNative(func(__e *ControlFlow) {
V5797 := __e.Get(1)
_ = V5797
tmp8102 := MakeNative(func(__e *ControlFlow) {
W5798 := __e.Get(1)
_ = W5798
tmp8103 := MakeNative(func(__e *ControlFlow) {
W5799 := __e.Get(1)
_ = W5799
__e.Return(V5797)
return
}, 1)

tmp8104 := Call(__e, PrimFunc(symadjoin), V5797, W5798)


tmp8105 := PrimValue(sym_dproperty_1vector_d)

tmp8106 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp8104, tmp8105)


__e.TailApply(tmp8103, tmp8106)
return


}, 1)

tmp8107 := PrimValue(sym_dproperty_1vector_d)

tmp8108 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp8107)


__e.TailApply(tmp8102, tmp8108)
return


}, 1)

tmp8109 := Call(__e, ns2_1set, symsystemf, tmp8101)


_ = tmp8109

tmp8110 := MakeNative(func(__e *ControlFlow) {
V5800 := __e.Get(1)
_ = V5800
V5801 := __e.Get(2)
_ = V5801
tmp8112 := Call(__e, PrimFunc(symelement_2), V5800, V5801)


if True == tmp8112 {
__e.Return(V5801)
return
} else {
__e.Return(PrimCons(V5800, V5801))
return
}


}, 2)

tmp8113 := Call(__e, ns2_1set, symadjoin, tmp8110)


_ = tmp8113

tmp8114 := PrimIntern(MakeString(":"))

tmp8115 := PrimIntern(MakeString(";"))

tmp8116 := PrimIntern(MakeString(":="))

tmp8117 := PrimIntern(MakeString(","))

tmp8118 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp8119 := PrimIntern(MakeString("bar!"))

tmp8120 := PrimCons(symabort, Nil)

tmp8121 := PrimCons(symabsvector, tmp8120)

tmp8122 := PrimCons(symabsvector_2, tmp8121)

tmp8123 := PrimCons(symaddress_1_6, tmp8122)

tmp8124 := PrimCons(sym_5_1address, tmp8123)

tmp8125 := PrimCons(symadjoin, tmp8124)

tmp8126 := PrimCons(symand, tmp8125)

tmp8127 := PrimCons(symappend, tmp8126)

tmp8128 := PrimCons(symarity, tmp8127)

tmp8129 := PrimCons(symassoc, tmp8128)

tmp8130 := PrimCons(symassertz, tmp8129)

tmp8131 := PrimCons(symasserta, tmp8130)

tmp8132 := PrimCons(symatom_2, tmp8131)

tmp8133 := PrimCons(symstep_2, tmp8132)

tmp8134 := PrimCons(symspy_2, tmp8133)

tmp8135 := PrimCons(tmp8119, tmp8134)

tmp8136 := PrimCons(symbootstrap, tmp8135)

tmp8137 := PrimCons(symboolean, tmp8136)

tmp8138 := PrimCons(symboolean_2, tmp8137)

tmp8139 := PrimCons(symbound_2, tmp8138)

tmp8140 := PrimCons(symbind, tmp8139)

tmp8141 := PrimCons(symclose, tmp8140)

tmp8142 := PrimCons(symcall, tmp8141)

tmp8143 := PrimCons(symcases, tmp8142)

tmp8144 := PrimCons(symcd, tmp8143)

tmp8145 := PrimCons(symcompile, tmp8144)

tmp8146 := PrimCons(symconcat, tmp8145)

tmp8147 := PrimCons(symcond, tmp8146)

tmp8148 := PrimCons(symcons, tmp8147)

tmp8149 := PrimCons(symcons_2, tmp8148)

tmp8150 := PrimCons(symcn, tmp8149)

tmp8151 := PrimCons(symdatatypes, tmp8150)

tmp8152 := PrimCons(symdatatype, tmp8151)

tmp8153 := PrimCons(symdeclare, tmp8152)

tmp8154 := PrimCons(symdefprolog, tmp8153)

tmp8155 := PrimCons(symdefcc, tmp8154)

tmp8156 := PrimCons(symdefmacro, tmp8155)

tmp8157 := PrimCons(symdefine, tmp8156)

tmp8158 := PrimCons(symdefun, tmp8157)

tmp8159 := PrimCons(symdestroy, tmp8158)

tmp8160 := PrimCons(symdifference, tmp8159)

tmp8161 := PrimCons(symdo, tmp8160)

tmp8162 := PrimCons(symelement_2, tmp8161)

tmp8163 := PrimCons(symempty_2, tmp8162)

tmp8164 := PrimCons(symerror, tmp8163)

tmp8165 := PrimCons(symerror_1to_1string, tmp8164)

tmp8166 := PrimCons(symeval, tmp8165)

tmp8167 := PrimCons(symeval_1kl, tmp8166)

tmp8168 := PrimCons(symexception, tmp8167)

tmp8169 := PrimCons(symexternal, tmp8168)

tmp8170 := PrimCons(symexplode, tmp8169)

tmp8171 := PrimCons(symenable_1type_1theory, tmp8170)

tmp8172 := PrimCons(False, tmp8171)

tmp8173 := PrimCons(symfindall, tmp8172)

tmp8174 := PrimCons(symfactorise_2, tmp8173)

tmp8175 := PrimCons(symfactorise, tmp8174)

tmp8176 := PrimCons(symfail_1if, tmp8175)

tmp8177 := PrimCons(symfail, tmp8176)

tmp8178 := PrimCons(symfile, tmp8177)

tmp8179 := PrimCons(symfix, tmp8178)

tmp8180 := PrimCons(symforeign, tmp8179)

tmp8181 := PrimCons(symfork, tmp8180)

tmp8182 := PrimCons(symfresh, tmp8181)

tmp8183 := PrimCons(symfreeze, tmp8182)

tmp8184 := PrimCons(symfst, tmp8183)

tmp8185 := PrimCons(symfunction, tmp8184)

tmp8186 := PrimCons(symfn, tmp8185)

tmp8187 := PrimCons(symgensym, tmp8186)

tmp8188 := PrimCons(symget_1time, tmp8187)

tmp8189 := PrimCons(symget, tmp8188)

tmp8190 := PrimCons(symhash, tmp8189)

tmp8191 := PrimCons(symhdstr, tmp8190)

tmp8192 := PrimCons(symhdv, tmp8191)

tmp8193 := PrimCons(symhd, tmp8192)

tmp8194 := PrimCons(symhead, tmp8193)

tmp8195 := PrimCons(symhush_2, tmp8194)

tmp8196 := PrimCons(symhush_2, tmp8195)

tmp8197 := PrimCons(symif, tmp8196)

tmp8198 := PrimCons(symimplementation, tmp8197)

tmp8199 := PrimCons(syminternal, tmp8198)

tmp8200 := PrimCons(symin_1package, tmp8199)

tmp8201 := PrimCons(symin, tmp8200)

tmp8202 := PrimCons(symis_b, tmp8201)

tmp8203 := PrimCons(symis, tmp8202)

tmp8204 := PrimCons(symit, tmp8203)

tmp8205 := PrimCons(syminclude_1all_1but, tmp8204)

tmp8206 := PrimCons(syminclude, tmp8205)

tmp8207 := PrimCons(symincluded, tmp8206)

tmp8208 := PrimCons(syminput_7, tmp8207)

tmp8209 := PrimCons(syminput, tmp8208)

tmp8210 := PrimCons(syminteger_2, tmp8209)

tmp8211 := PrimCons(symintern, tmp8210)

tmp8212 := PrimCons(syminferences, tmp8211)

tmp8213 := PrimCons(symintersection, tmp8212)

tmp8214 := PrimCons(symis, tmp8213)

tmp8215 := PrimCons(symlanguage, tmp8214)

tmp8216 := PrimCons(symlambda, tmp8215)

tmp8217 := PrimCons(symlazy, tmp8216)

tmp8218 := PrimCons(symlet, tmp8217)

tmp8219 := PrimCons(symlength, tmp8218)

tmp8220 := PrimCons(symlimit, tmp8219)

tmp8221 := PrimCons(symlineread, tmp8220)

tmp8222 := PrimCons(symlist, tmp8221)

tmp8223 := PrimCons(symloaded, tmp8222)

tmp8224 := PrimCons(symload, tmp8223)

tmp8225 := PrimCons(symmake_1string, tmp8224)

tmp8226 := PrimCons(symmap, tmp8225)

tmp8227 := PrimCons(symmapcan, tmp8226)

tmp8228 := PrimCons(symmaxinferences, tmp8227)

tmp8229 := PrimCons(symmacroexpand, tmp8228)

tmp8230 := PrimCons(symmode, tmp8229)

tmp8231 := PrimCons(symnl, tmp8230)

tmp8232 := PrimCons(symnot, tmp8231)

tmp8233 := PrimCons(symnth, tmp8232)

tmp8234 := PrimCons(symnull, tmp8233)

tmp8235 := PrimCons(symnumber, tmp8234)

tmp8236 := PrimCons(symnumber_2, tmp8235)

tmp8237 := PrimCons(symn_1_6string, tmp8236)

tmp8238 := PrimCons(symoccurs_2, tmp8237)

tmp8239 := PrimCons(symoccurs_1check, tmp8238)

tmp8240 := PrimCons(symoccurrences, tmp8239)

tmp8241 := PrimCons(symopen, tmp8240)

tmp8242 := PrimCons(symoptimise_2, tmp8241)

tmp8243 := PrimCons(symoptimise, tmp8242)

tmp8244 := PrimCons(symor, tmp8243)

tmp8245 := PrimCons(symos, tmp8244)

tmp8246 := PrimCons(symout, tmp8245)

tmp8247 := PrimCons(symoutput, tmp8246)

tmp8248 := PrimCons(sympackage, tmp8247)

tmp8249 := PrimCons(symport, tmp8248)

tmp8250 := PrimCons(symporters, tmp8249)

tmp8251 := PrimCons(sympos, tmp8250)

tmp8252 := PrimCons(sympr, tmp8251)

tmp8253 := PrimCons(symprint, tmp8252)

tmp8254 := PrimCons(symprolog_1memory, tmp8253)

tmp8255 := PrimCons(symprofile, tmp8254)

tmp8256 := PrimCons(symprofile_1results, tmp8255)

tmp8257 := PrimCons(symprotect, tmp8256)

tmp8258 := PrimCons(symprolog_2, tmp8257)

tmp8259 := PrimCons(symps, tmp8258)

tmp8260 := PrimCons(sympreclude_1all_1but, tmp8259)

tmp8261 := PrimCons(sympreclude, tmp8260)

tmp8262 := PrimCons(symput, tmp8261)

tmp8263 := PrimCons(sympackage_2, tmp8262)

tmp8264 := PrimCons(symread_1from_1string_1unprocessed, tmp8263)

tmp8265 := PrimCons(symread_1from_1string, tmp8264)

tmp8266 := PrimCons(symread_1byte, tmp8265)

tmp8267 := PrimCons(symread_1file_1as_1string, tmp8266)

tmp8268 := PrimCons(symread_1file_1as_1bytelist, tmp8267)

tmp8269 := PrimCons(symread_1file, tmp8268)

tmp8270 := PrimCons(symreceive, tmp8269)

tmp8271 := PrimCons(symread, tmp8270)

tmp8272 := PrimCons(symrelease, tmp8271)

tmp8273 := PrimCons(symremove, tmp8272)

tmp8274 := PrimCons(symretract, tmp8273)

tmp8275 := PrimCons(symreverse, tmp8274)

tmp8276 := PrimCons(symrun, tmp8275)

tmp8277 := PrimCons(symstr, tmp8276)

tmp8278 := PrimCons(symsave, tmp8277)

tmp8279 := PrimCons(symset, tmp8278)

tmp8280 := PrimCons(symsimple_1error, tmp8279)

tmp8281 := PrimCons(symsnd, tmp8280)

tmp8282 := PrimCons(symspecialise, tmp8281)

tmp8283 := PrimCons(symspy, tmp8282)

tmp8284 := PrimCons(symstep, tmp8283)

tmp8285 := PrimCons(symstoutput, tmp8284)

tmp8286 := PrimCons(symstinput, tmp8285)

tmp8287 := PrimCons(symstring, tmp8286)

tmp8288 := PrimCons(symstream, tmp8287)

tmp8289 := PrimCons(symstring_1_6n, tmp8288)

tmp8290 := PrimCons(symstring_2, tmp8289)

tmp8291 := PrimCons(symsubst, tmp8290)

tmp8292 := PrimCons(symsum, tmp8291)

tmp8293 := PrimCons(symstring_1_6symbol, tmp8292)

tmp8294 := PrimCons(symsymbol_2, tmp8293)

tmp8295 := PrimCons(symsymbol, tmp8294)

tmp8296 := PrimCons(symsystem_1S_2, tmp8295)

tmp8297 := PrimCons(symsynonyms, tmp8296)

tmp8298 := PrimCons(symsystemf, tmp8297)

tmp8299 := PrimCons(symtail, tmp8298)

tmp8300 := PrimCons(symtlv, tmp8299)

tmp8301 := PrimCons(symtlstr, tmp8300)

tmp8302 := PrimCons(symtl, tmp8301)

tmp8303 := PrimCons(symtc, tmp8302)

tmp8304 := PrimCons(symtc_2, tmp8303)

tmp8305 := PrimCons(symthaw, tmp8304)

tmp8306 := PrimCons(symtime, tmp8305)

tmp8307 := PrimCons(symtrack, tmp8306)

tmp8308 := PrimCons(symtracked, tmp8307)

tmp8309 := PrimCons(symtrap_1error, tmp8308)

tmp8310 := PrimCons(True, tmp8309)

tmp8311 := PrimCons(symtuple_2, tmp8310)

tmp8312 := PrimCons(symtype, tmp8311)

tmp8313 := PrimCons(symreturn, tmp8312)

tmp8314 := PrimCons(symundefmacro, tmp8313)

tmp8315 := PrimCons(symunprofile, tmp8314)

tmp8316 := PrimCons(symunput, tmp8315)

tmp8317 := PrimCons(symunion, tmp8316)

tmp8318 := PrimCons(symunix, tmp8317)

tmp8319 := PrimCons(symunit, tmp8318)

tmp8320 := PrimCons(symuntrack, tmp8319)

tmp8321 := PrimCons(symunspecialise, tmp8320)

tmp8322 := PrimCons(symupdate_1lambda_1table, tmp8321)

tmp8323 := PrimCons(symu_b, tmp8322)

tmp8324 := PrimCons(symuserdefs, tmp8323)

tmp8325 := PrimCons(symvector_2, tmp8324)

tmp8326 := PrimCons(symvector, tmp8325)

tmp8327 := PrimCons(sym_5_1vector, tmp8326)

tmp8328 := PrimCons(symvector_1_6, tmp8327)

tmp8329 := PrimCons(symvalue, tmp8328)

tmp8330 := PrimCons(symvar_2, tmp8329)

tmp8331 := PrimCons(symvariable_2, tmp8330)

tmp8332 := PrimCons(symverified, tmp8331)

tmp8333 := PrimCons(symversion, tmp8332)

tmp8334 := PrimCons(symwhen, tmp8333)

tmp8335 := PrimCons(symwhere, tmp8334)

tmp8336 := PrimCons(symwrite_1byte, tmp8335)

tmp8337 := PrimCons(symwrite_1to_1file, tmp8336)

tmp8338 := PrimCons(symy_1or_1n_2, tmp8337)

tmp8339 := PrimCons(tmp8118, tmp8338)

tmp8340 := PrimCons(sym_6_6, tmp8339)

tmp8341 := PrimCons(sym_5, tmp8340)

tmp8342 := PrimCons(sym_5_a, tmp8341)

tmp8343 := PrimCons(sym_7, tmp8342)

tmp8344 := PrimCons(sym_d, tmp8343)

tmp8345 := PrimCons(sym_c, tmp8344)

tmp8346 := PrimCons(sym_1, tmp8345)

tmp8347 := PrimCons(sym_3, tmp8346)

tmp8348 := PrimCons(sym_5end_6, tmp8347)

tmp8349 := PrimCons(sym_5_b_6, tmp8348)

tmp8350 := PrimCons(sym_c_4, tmp8349)

tmp8351 := PrimCons(sym_a_a_6, tmp8350)

tmp8352 := PrimCons(sym_6, tmp8351)

tmp8353 := PrimCons(sym_6_a, tmp8352)

tmp8354 := PrimCons(sym_a, tmp8353)

tmp8355 := PrimCons(sym_a_a, tmp8354)

tmp8356 := PrimCons(sym_5e_6, tmp8355)

tmp8357 := PrimCons(sym_1_6, tmp8356)

tmp8358 := PrimCons(sym_5_1, tmp8357)

tmp8359 := PrimCons(sym_dhush_d, tmp8358)

tmp8360 := PrimCons(sym_dporters_d, tmp8359)

tmp8361 := PrimCons(sym_dport_d, tmp8360)

tmp8362 := PrimCons(sym_8s, tmp8361)

tmp8363 := PrimCons(sym_8p, tmp8362)

tmp8364 := PrimCons(sym_8v, tmp8363)

tmp8365 := PrimCons(sym_dproperty_1vector_d, tmp8364)

tmp8366 := PrimCons(sym_drelease_d, tmp8365)

tmp8367 := PrimCons(sym_dos_d, tmp8366)

tmp8368 := PrimCons(sym_dmacros_d, tmp8367)

tmp8369 := PrimCons(sym_dmaximum_1print_1sequence_1size_d, tmp8368)

tmp8370 := PrimCons(sym_dversion_d, tmp8369)

tmp8371 := PrimCons(sym_dhome_1directory_d, tmp8370)

tmp8372 := PrimCons(sym_dstoutput_d, tmp8371)

tmp8373 := PrimCons(sym_dstinput_d, tmp8372)

tmp8374 := PrimCons(sym_dimplementation_d, tmp8373)

tmp8375 := PrimCons(sym_dlanguage_d, tmp8374)

tmp8376 := PrimCons(sym__, tmp8375)

tmp8377 := PrimCons(tmp8117, tmp8376)

tmp8378 := PrimCons(tmp8116, tmp8377)

tmp8379 := PrimCons(tmp8115, tmp8378)

tmp8380 := PrimCons(tmp8114, tmp8379)

tmp8381 := PrimCons(sym_e_e, tmp8380)

tmp8382 := PrimCons(sym_5_1_1, tmp8381)

tmp8383 := PrimCons(sym_1_1_6, tmp8382)

tmp8384 := PrimCons(sym_i, tmp8383)

tmp8385 := PrimCons(sym_j, tmp8384)

tmp8386 := PrimCons(sym_b, tmp8385)

tmp8387 := PrimValue(sym_dproperty_1vector_d)

tmp8388 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp8386, tmp8387)


_ = tmp8388

tmp8389 := MakeNative(func(__e *ControlFlow) {
V5802 := __e.Get(1)
_ = V5802
tmp8390 := MakeNative(func(__e *ControlFlow) {
W5803 := __e.Get(1)
_ = W5803
tmp8398 := PrimEqual(W5803, MakeNumber(-1))

var ifres8395 Obj

if True == tmp8398 {
ifres8395 = True


} else {
tmp8397 := PrimEqual(W5803, MakeNumber(0))

var ifres8396 Obj

if True == tmp8397 {
ifres8396 = True


} else {
ifres8396 = False


}

ifres8395 = ifres8396


}

if True == ifres8395 {
__e.Return(Nil)
return
} else {
tmp8391 := PrimCons(V5802, Nil)

tmp8392 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp8391, W5803)


tmp8393 := Call(__e, PrimFunc(symeval_1kl), tmp8392)


__e.Return(PrimCons(V5802, tmp8393))
return


}


}, 1)

tmp8399 := Call(__e, PrimFunc(symarity), V5802)


__e.TailApply(tmp8390, tmp8399)
return


}, 1)

tmp8400 := Call(__e, ns2_1set, symshen_4lambda_1entry, tmp8389)


_ = tmp8400

tmp8401 := MakeNative(func(__e *ControlFlow) {
V5804 := __e.Get(1)
_ = V5804
tmp8402 := MakeNative(func(__e *ControlFlow) {
W5805 := __e.Get(1)
_ = W5805
tmp8403 := MakeNative(func(__e *ControlFlow) {
Z5807 := __e.Get(1)
_ = Z5807
__e.TailApply(PrimFunc(symshen_4tuple), Z5807)
return
}, 1)

tmp8404 := PrimCons(symshen_4tuple, tmp8403)

tmp8405 := MakeNative(func(__e *ControlFlow) {
Z5808 := __e.Get(1)
_ = Z5808
__e.TailApply(PrimFunc(symshen_4pvar), Z5808)
return
}, 1)

tmp8406 := PrimCons(symshen_4pvar, tmp8405)

tmp8407 := MakeNative(func(__e *ControlFlow) {
Z5809 := __e.Get(1)
_ = Z5809
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Z5809)
return
}, 1)

tmp8408 := PrimCons(symshen_4print_1prolog_1vector, tmp8407)

tmp8409 := MakeNative(func(__e *ControlFlow) {
Z5810 := __e.Get(1)
_ = Z5810
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Z5810)
return
}, 1)

tmp8410 := PrimCons(symshen_4print_1freshterm, tmp8409)

tmp8411 := MakeNative(func(__e *ControlFlow) {
Z5811 := __e.Get(1)
_ = Z5811
__e.TailApply(PrimFunc(symshen_4printF), Z5811)
return
}, 1)

tmp8412 := PrimCons(symshen_4printF, tmp8411)

tmp8413 := PrimCons(tmp8412, W5805)

tmp8414 := PrimCons(tmp8410, tmp8413)

tmp8415 := PrimCons(tmp8408, tmp8414)

tmp8416 := PrimCons(tmp8406, tmp8415)

tmp8417 := PrimCons(tmp8404, tmp8416)

__e.Return(PrimSet(symshen_4_dlambdatable_d, tmp8417))
return


}, 1)

tmp8418 := MakeNative(func(__e *ControlFlow) {
Z5806 := __e.Get(1)
_ = Z5806
__e.TailApply(PrimFunc(symshen_4lambda_1entry), Z5806)
return
}, 1)

tmp8419 := Call(__e, PrimFunc(symmap), tmp8418, V5804)


__e.TailApply(tmp8402, tmp8419)
return


}, 1)

tmp8420 := Call(__e, ns2_1set, symshen_4build_1lambda_1table, tmp8401)


_ = tmp8420

tmp8421 := Call(__e, PrimFunc(symexternal), symshen)


__e.TailApply(PrimFunc(symshen_4build_1lambda_1table), tmp8421)
return




}, 0)

