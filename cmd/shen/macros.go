package main

import . "github.com/tiancaiamao/shen-go/kl"

var MacrosMain = MakeNative(func(__e *ControlFlow) {
tmp8780 := MakeNative(func(__e *ControlFlow) {
V5843 := __e.Get(1)
_ = V5843
tmp8781 := MakeNative(func(__e *ControlFlow) {
W5844 := __e.Get(1)
_ = W5844
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V5843, W5844, W5844)
return
}, 1)

tmp8782 := MakeNative(func(__e *ControlFlow) {
Z5845 := __e.Get(1)
_ = Z5845
__e.Return(PrimTail(Z5845))
return
}, 1)

tmp8783 := PrimValue(sym_dmacros_d)

tmp8784 := Call(__e, PrimFunc(symmap), tmp8782, tmp8783)


__e.TailApply(tmp8781, tmp8784)
return


}, 1)

tmp8785 := Call(__e, ns2_1set, symmacroexpand, tmp8780)


_ = tmp8785

tmp8786 := MakeNative(func(__e *ControlFlow) {
V5854 := __e.Get(1)
_ = V5854
V5855 := __e.Get(2)
_ = V5855
V5856 := __e.Get(3)
_ = V5856
tmp8796 := PrimEqual(Nil, V5855)

if True == tmp8796 {
__e.Return(V5854)
return
} else {
tmp8794 := PrimIsPair(V5855)

if True == tmp8794 {
tmp8787 := MakeNative(func(__e *ControlFlow) {
W5857 := __e.Get(1)
_ = W5857
tmp8790 := PrimEqual(V5854, W5857)

if True == tmp8790 {
tmp8788 := PrimTail(V5855)

__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V5854, tmp8788, V5856)
return


} else {
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), W5857, V5856, V5856)
return
}


}, 1)

tmp8791 := PrimHead(V5855)

tmp8792 := Call(__e, PrimFunc(symshen_4walk), tmp8791, V5854)


__e.TailApply(tmp8787, tmp8792)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.macroexpand-h")))
return
}


}


}, 3)

tmp8797 := Call(__e, ns2_1set, symshen_4macroexpand_1h, tmp8786)


_ = tmp8797

tmp8798 := MakeNative(func(__e *ControlFlow) {
V5858 := __e.Get(1)
_ = V5858
V5859 := __e.Get(2)
_ = V5859
tmp8802 := PrimIsPair(V5859)

if True == tmp8802 {
tmp8799 := MakeNative(func(__e *ControlFlow) {
Z5860 := __e.Get(1)
_ = Z5860
__e.TailApply(PrimFunc(symshen_4walk), V5858, Z5860)
return
}, 1)

tmp8800 := Call(__e, PrimFunc(symmap), tmp8799, V5859)


__e.TailApply(V5858, tmp8800)
return


} else {
__e.TailApply(V5858, V5859)
return
}


}, 2)

tmp8803 := Call(__e, ns2_1set, symshen_4walk, tmp8798)


_ = tmp8803

tmp8804 := MakeNative(func(__e *ControlFlow) {
V5861 := __e.Get(1)
_ = V5861
tmp8805 := MakeNative(func(__e *ControlFlow) {
GoTo5862 := __e.Get(1)
_ = GoTo5862
tmp9057 := PrimIsPair(V5861)

if True == tmp9057 {
tmp8806 := MakeNative(func(__e *ControlFlow) {
Select5863 := __e.Get(1)
_ = Select5863
tmp8807 := MakeNative(func(__e *ControlFlow) {
Select5864 := __e.Get(1)
_ = Select5864
tmp9053 := PrimEqual(symdefmacro, Select5863)

var ifres9050 Obj

if True == tmp9053 {
tmp9052 := PrimIsPair(Select5864)

var ifres9051 Obj

if True == tmp9052 {
ifres9051 = True


} else {
ifres9051 = False


}

ifres9050 = ifres9051


} else {
ifres9050 = False


}

if True == ifres9050 {
tmp8808 := PrimHead(Select5864)

tmp8809 := PrimTail(Select5864)

__e.TailApply(PrimFunc(symshen_4process_1def), tmp8808, tmp8809)
return


} else {
tmp9048 := PrimEqual(symdefcc, Select5863)

if True == tmp9048 {
__e.TailApply(PrimFunc(symshen_4yacc_1_6shen), Select5864)
return
} else {
tmp9046 := PrimEqual(symu_b, Select5863)

var ifres9039 Obj

if True == tmp9046 {
tmp9045 := PrimIsPair(Select5864)

var ifres9041 Obj

if True == tmp9045 {
tmp9043 := PrimTail(Select5864)

tmp9044 := PrimEqual(Nil, tmp9043)

var ifres9042 Obj

if True == tmp9044 {
ifres9042 = True


} else {
ifres9042 = False


}

ifres9041 = ifres9042


} else {
ifres9041 = False


}

var ifres9040 Obj

if True == ifres9041 {
ifres9040 = True


} else {
ifres9040 = False


}

ifres9039 = ifres9040


} else {
ifres9039 = False


}

if True == ifres9039 {
tmp8810 := PrimHead(Select5864)

tmp8811 := Call(__e, PrimFunc(symshen_4make_1uppercase), tmp8810)


tmp8812 := PrimCons(tmp8811, Nil)

__e.Return(PrimCons(symprotect, tmp8812))
return


} else {
tmp9037 := PrimEqual(symerror, Select5863)

var ifres9034 Obj

if True == tmp9037 {
tmp9036 := PrimIsPair(Select5864)

var ifres9035 Obj

if True == tmp9036 {
ifres9035 = True


} else {
ifres9035 = False


}

ifres9034 = ifres9035


} else {
ifres9034 = False


}

if True == ifres9034 {
tmp8813 := PrimHead(Select5864)

tmp8814 := PrimTail(Select5864)

tmp8815 := Call(__e, PrimFunc(symshen_4mkstr), tmp8813, tmp8814)


tmp8816 := PrimCons(tmp8815, Nil)

__e.Return(PrimCons(symsimple_1error, tmp8816))
return


} else {
tmp9032 := PrimEqual(symoutput, Select5863)

var ifres9029 Obj

if True == tmp9032 {
tmp9031 := PrimIsPair(Select5864)

var ifres9030 Obj

if True == tmp9031 {
ifres9030 = True


} else {
ifres9030 = False


}

ifres9029 = ifres9030


} else {
ifres9029 = False


}

if True == ifres9029 {
tmp8817 := PrimHead(Select5864)

tmp8818 := PrimTail(Select5864)

tmp8819 := Call(__e, PrimFunc(symshen_4mkstr), tmp8817, tmp8818)


tmp8820 := PrimCons(symstoutput, Nil)

tmp8821 := PrimCons(tmp8820, Nil)

tmp8822 := PrimCons(tmp8819, tmp8821)

__e.Return(PrimCons(sympr, tmp8822))
return


} else {
tmp9027 := PrimEqual(sympr, Select5863)

var ifres9020 Obj

if True == tmp9027 {
tmp9026 := PrimIsPair(Select5864)

var ifres9022 Obj

if True == tmp9026 {
tmp9024 := PrimTail(Select5864)

tmp9025 := PrimEqual(Nil, tmp9024)

var ifres9023 Obj

if True == tmp9025 {
ifres9023 = True


} else {
ifres9023 = False


}

ifres9022 = ifres9023


} else {
ifres9022 = False


}

var ifres9021 Obj

if True == ifres9022 {
ifres9021 = True


} else {
ifres9021 = False


}

ifres9020 = ifres9021


} else {
ifres9020 = False


}

if True == ifres9020 {
tmp8823 := PrimHead(Select5864)

tmp8824 := PrimCons(symstoutput, Nil)

tmp8825 := PrimCons(tmp8824, Nil)

tmp8826 := PrimCons(tmp8823, tmp8825)

__e.Return(PrimCons(sympr, tmp8826))
return


} else {
tmp9018 := PrimEqual(symmake_1string, Select5863)

var ifres9015 Obj

if True == tmp9018 {
tmp9017 := PrimIsPair(Select5864)

var ifres9016 Obj

if True == tmp9017 {
ifres9016 = True


} else {
ifres9016 = False


}

ifres9015 = ifres9016


} else {
ifres9015 = False


}

if True == ifres9015 {
tmp8827 := PrimHead(Select5864)

tmp8828 := PrimTail(Select5864)

__e.TailApply(PrimFunc(symshen_4mkstr), tmp8827, tmp8828)
return


} else {
tmp9013 := PrimEqual(symlineread, Select5863)

var ifres9010 Obj

if True == tmp9013 {
tmp9012 := PrimEqual(Nil, Select5864)

var ifres9011 Obj

if True == tmp9012 {
ifres9011 = True


} else {
ifres9011 = False


}

ifres9010 = ifres9011


} else {
ifres9010 = False


}

if True == ifres9010 {
tmp8829 := PrimCons(symstinput, Nil)

tmp8830 := PrimCons(tmp8829, Nil)

__e.Return(PrimCons(symlineread, tmp8830))
return


} else {
tmp9008 := PrimEqual(syminput, Select5863)

var ifres9005 Obj

if True == tmp9008 {
tmp9007 := PrimEqual(Nil, Select5864)

var ifres9006 Obj

if True == tmp9007 {
ifres9006 = True


} else {
ifres9006 = False


}

ifres9005 = ifres9006


} else {
ifres9005 = False


}

if True == ifres9005 {
tmp8831 := PrimCons(symstinput, Nil)

tmp8832 := PrimCons(tmp8831, Nil)

__e.Return(PrimCons(syminput, tmp8832))
return


} else {
tmp9003 := PrimEqual(symread, Select5863)

var ifres9000 Obj

if True == tmp9003 {
tmp9002 := PrimEqual(Nil, Select5864)

var ifres9001 Obj

if True == tmp9002 {
ifres9001 = True


} else {
ifres9001 = False


}

ifres9000 = ifres9001


} else {
ifres9000 = False


}

if True == ifres9000 {
tmp8833 := PrimCons(symstinput, Nil)

tmp8834 := PrimCons(tmp8833, Nil)

__e.Return(PrimCons(symread, tmp8834))
return


} else {
tmp8998 := PrimEqual(syminput_7, Select5863)

var ifres8991 Obj

if True == tmp8998 {
tmp8997 := PrimIsPair(Select5864)

var ifres8993 Obj

if True == tmp8997 {
tmp8995 := PrimTail(Select5864)

tmp8996 := PrimEqual(Nil, tmp8995)

var ifres8994 Obj

if True == tmp8996 {
ifres8994 = True


} else {
ifres8994 = False


}

ifres8993 = ifres8994


} else {
ifres8993 = False


}

var ifres8992 Obj

if True == ifres8993 {
ifres8992 = True


} else {
ifres8992 = False


}

ifres8991 = ifres8992


} else {
ifres8991 = False


}

if True == ifres8991 {
tmp8835 := PrimHead(Select5864)

tmp8836 := PrimCons(symstinput, Nil)

tmp8837 := PrimCons(tmp8836, Nil)

tmp8838 := PrimCons(tmp8835, tmp8837)

__e.Return(PrimCons(syminput_7, tmp8838))
return


} else {
tmp8989 := PrimEqual(symread_1byte, Select5863)

var ifres8986 Obj

if True == tmp8989 {
tmp8988 := PrimEqual(Nil, Select5864)

var ifres8987 Obj

if True == tmp8988 {
ifres8987 = True


} else {
ifres8987 = False


}

ifres8986 = ifres8987


} else {
ifres8986 = False


}

if True == ifres8986 {
__e.TailApply(PrimFunc(symshen_4process_1read_1byte))
return
} else {
tmp8984 := PrimEqual(symprolog_2, Select5863)

if True == tmp8984 {
__e.TailApply(PrimFunc(symshen_4call_1prolog), Select5864)
return
} else {
tmp8982 := PrimEqual(symdefprolog, Select5863)

var ifres8979 Obj

if True == tmp8982 {
tmp8981 := PrimIsPair(Select5864)

var ifres8980 Obj

if True == tmp8981 {
ifres8980 = True


} else {
ifres8980 = False


}

ifres8979 = ifres8980


} else {
ifres8979 = False


}

if True == ifres8979 {
tmp8839 := PrimHead(Select5864)

tmp8840 := PrimTail(Select5864)

__e.TailApply(PrimFunc(symshen_4compile_1prolog), tmp8839, tmp8840)
return


} else {
tmp8977 := PrimEqual(symdatatype, Select5863)

var ifres8974 Obj

if True == tmp8977 {
tmp8976 := PrimIsPair(Select5864)

var ifres8975 Obj

if True == tmp8976 {
ifres8975 = True


} else {
ifres8975 = False


}

ifres8974 = ifres8975


} else {
ifres8974 = False


}

if True == ifres8974 {
tmp8841 := PrimHead(Select5864)

tmp8842 := PrimTail(Select5864)

__e.TailApply(PrimFunc(symshen_4process_1datatype), tmp8841, tmp8842)
return


} else {
tmp8972 := PrimEqual(sym_8s, Select5863)

if True == tmp8972 {
__e.TailApply(PrimFunc(symshen_4process_1_8s), V5861)
return
} else {
tmp8970 := PrimEqual(symsynonyms, Select5863)

if True == tmp8970 {
__e.TailApply(PrimFunc(symshen_4process_1synonyms), Select5864)
return
} else {
tmp8968 := PrimEqual(symnl, Select5863)

var ifres8965 Obj

if True == tmp8968 {
tmp8967 := PrimEqual(Nil, Select5864)

var ifres8966 Obj

if True == tmp8967 {
ifres8966 = True


} else {
ifres8966 = False


}

ifres8965 = ifres8966


} else {
ifres8965 = False


}

if True == ifres8965 {
tmp8843 := PrimCons(MakeNumber(1), Nil)

__e.Return(PrimCons(symnl, tmp8843))
return


} else {
tmp8963 := PrimEqual(symlet, Select5863)

if True == tmp8963 {
__e.TailApply(PrimFunc(symshen_4process_1let), V5861)
return
} else {
tmp8961 := PrimEqual(sym_c_4, Select5863)

if True == tmp8961 {
__e.TailApply(PrimFunc(symshen_4process_1lambda), V5861)
return
} else {
tmp8959 := PrimEqual(symcases, Select5863)

if True == tmp8959 {
__e.TailApply(PrimFunc(symshen_4process_1cases), V5861)
return
} else {
tmp8957 := PrimEqual(symtime, Select5863)

var ifres8950 Obj

if True == tmp8957 {
tmp8956 := PrimIsPair(Select5864)

var ifres8952 Obj

if True == tmp8956 {
tmp8954 := PrimTail(Select5864)

tmp8955 := PrimEqual(Nil, tmp8954)

var ifres8953 Obj

if True == tmp8955 {
ifres8953 = True


} else {
ifres8953 = False


}

ifres8952 = ifres8953


} else {
ifres8952 = False


}

var ifres8951 Obj

if True == ifres8952 {
ifres8951 = True


} else {
ifres8951 = False


}

ifres8950 = ifres8951


} else {
ifres8950 = False


}

if True == ifres8950 {
tmp8844 := PrimHead(Select5864)

__e.TailApply(PrimFunc(symshen_4process_1time), tmp8844)
return


} else {
tmp8948 := PrimEqual(symput, Select5863)

var ifres8930 Obj

if True == tmp8948 {
tmp8947 := PrimIsPair(Select5864)

var ifres8932 Obj

if True == tmp8947 {
tmp8945 := PrimTail(Select5864)

tmp8946 := PrimIsPair(tmp8945)

var ifres8934 Obj

if True == tmp8946 {
tmp8942 := PrimTail(Select5864)

tmp8943 := PrimTail(tmp8942)

tmp8944 := PrimIsPair(tmp8943)

var ifres8936 Obj

if True == tmp8944 {
tmp8938 := PrimTail(Select5864)

tmp8939 := PrimTail(tmp8938)

tmp8940 := PrimTail(tmp8939)

tmp8941 := PrimEqual(Nil, tmp8940)

var ifres8937 Obj

if True == tmp8941 {
ifres8937 = True


} else {
ifres8937 = False


}

ifres8936 = ifres8937


} else {
ifres8936 = False


}

var ifres8935 Obj

if True == ifres8936 {
ifres8935 = True


} else {
ifres8935 = False


}

ifres8934 = ifres8935


} else {
ifres8934 = False


}

var ifres8933 Obj

if True == ifres8934 {
ifres8933 = True


} else {
ifres8933 = False


}

ifres8932 = ifres8933


} else {
ifres8932 = False


}

var ifres8931 Obj

if True == ifres8932 {
ifres8931 = True


} else {
ifres8931 = False


}

ifres8930 = ifres8931


} else {
ifres8930 = False


}

if True == ifres8930 {
tmp8845 := PrimHead(Select5864)

tmp8846 := PrimTail(Select5864)

tmp8847 := PrimHead(tmp8846)

tmp8848 := PrimTail(Select5864)

tmp8849 := PrimTail(tmp8848)

tmp8850 := PrimHead(tmp8849)

tmp8851 := PrimCons(sym_dproperty_1vector_d, Nil)

tmp8852 := PrimCons(symvalue, tmp8851)

tmp8853 := PrimCons(tmp8852, Nil)

tmp8854 := PrimCons(tmp8850, tmp8853)

tmp8855 := PrimCons(tmp8847, tmp8854)

tmp8856 := PrimCons(tmp8845, tmp8855)

__e.Return(PrimCons(symput, tmp8856))
return


} else {
tmp8928 := PrimEqual(symget, Select5863)

var ifres8916 Obj

if True == tmp8928 {
tmp8927 := PrimIsPair(Select5864)

var ifres8918 Obj

if True == tmp8927 {
tmp8925 := PrimTail(Select5864)

tmp8926 := PrimIsPair(tmp8925)

var ifres8920 Obj

if True == tmp8926 {
tmp8922 := PrimTail(Select5864)

tmp8923 := PrimTail(tmp8922)

tmp8924 := PrimEqual(Nil, tmp8923)

var ifres8921 Obj

if True == tmp8924 {
ifres8921 = True


} else {
ifres8921 = False


}

ifres8920 = ifres8921


} else {
ifres8920 = False


}

var ifres8919 Obj

if True == ifres8920 {
ifres8919 = True


} else {
ifres8919 = False


}

ifres8918 = ifres8919


} else {
ifres8918 = False


}

var ifres8917 Obj

if True == ifres8918 {
ifres8917 = True


} else {
ifres8917 = False


}

ifres8916 = ifres8917


} else {
ifres8916 = False


}

if True == ifres8916 {
tmp8857 := PrimHead(Select5864)

tmp8858 := PrimTail(Select5864)

tmp8859 := PrimHead(tmp8858)

tmp8860 := PrimCons(sym_dproperty_1vector_d, Nil)

tmp8861 := PrimCons(symvalue, tmp8860)

tmp8862 := PrimCons(tmp8861, Nil)

tmp8863 := PrimCons(tmp8859, tmp8862)

tmp8864 := PrimCons(tmp8857, tmp8863)

__e.Return(PrimCons(symget, tmp8864))
return


} else {
tmp8914 := PrimEqual(symunput, Select5863)

var ifres8902 Obj

if True == tmp8914 {
tmp8913 := PrimIsPair(Select5864)

var ifres8904 Obj

if True == tmp8913 {
tmp8911 := PrimTail(Select5864)

tmp8912 := PrimIsPair(tmp8911)

var ifres8906 Obj

if True == tmp8912 {
tmp8908 := PrimTail(Select5864)

tmp8909 := PrimTail(tmp8908)

tmp8910 := PrimEqual(Nil, tmp8909)

var ifres8907 Obj

if True == tmp8910 {
ifres8907 = True


} else {
ifres8907 = False


}

ifres8906 = ifres8907


} else {
ifres8906 = False


}

var ifres8905 Obj

if True == ifres8906 {
ifres8905 = True


} else {
ifres8905 = False


}

ifres8904 = ifres8905


} else {
ifres8904 = False


}

var ifres8903 Obj

if True == ifres8904 {
ifres8903 = True


} else {
ifres8903 = False


}

ifres8902 = ifres8903


} else {
ifres8902 = False


}

if True == ifres8902 {
tmp8865 := PrimHead(Select5864)

tmp8866 := PrimTail(Select5864)

tmp8867 := PrimHead(tmp8866)

tmp8868 := PrimCons(sym_dproperty_1vector_d, Nil)

tmp8869 := PrimCons(symvalue, tmp8868)

tmp8870 := PrimCons(tmp8869, Nil)

tmp8871 := PrimCons(tmp8867, tmp8870)

tmp8872 := PrimCons(tmp8865, tmp8871)

__e.Return(PrimCons(symunput, tmp8872))
return


} else {
tmp8900 := PrimIsPair(Select5864)

var ifres8880 Obj

if True == tmp8900 {
tmp8898 := PrimTail(Select5864)

tmp8899 := PrimIsPair(tmp8898)

var ifres8882 Obj

if True == tmp8899 {
tmp8895 := PrimTail(Select5864)

tmp8896 := PrimTail(tmp8895)

tmp8897 := PrimIsPair(tmp8896)

var ifres8884 Obj

if True == tmp8897 {
tmp8886 := PrimCons(symdo, Nil)

tmp8887 := PrimCons(sym_d, tmp8886)

tmp8888 := PrimCons(sym_7, tmp8887)

tmp8889 := PrimCons(symor, tmp8888)

tmp8890 := PrimCons(symand, tmp8889)

tmp8891 := PrimCons(symappend, tmp8890)

tmp8892 := PrimCons(sym_8v, tmp8891)

tmp8893 := PrimCons(sym_8p, tmp8892)

tmp8894 := Call(__e, PrimFunc(symelement_2), Select5863, tmp8893)


var ifres8885 Obj

if True == tmp8894 {
ifres8885 = True


} else {
ifres8885 = False


}

ifres8884 = ifres8885


} else {
ifres8884 = False


}

var ifres8883 Obj

if True == ifres8884 {
ifres8883 = True


} else {
ifres8883 = False


}

ifres8882 = ifres8883


} else {
ifres8882 = False


}

var ifres8881 Obj

if True == ifres8882 {
ifres8881 = True


} else {
ifres8881 = False


}

ifres8880 = ifres8881


} else {
ifres8880 = False


}

if True == ifres8880 {
tmp8873 := PrimHead(Select5864)

tmp8874 := PrimTail(Select5864)

tmp8875 := PrimCons(Select5863, tmp8874)

tmp8876 := Call(__e, PrimFunc(symshen_4process_1assoc), tmp8875)


tmp8877 := PrimCons(tmp8876, Nil)

tmp8878 := PrimCons(tmp8873, tmp8877)

__e.Return(PrimCons(Select5863, tmp8878))
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5862)
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


}


}


}


}


}, 1)

tmp9054 := PrimTail(V5861)

__e.TailApply(tmp8807, tmp9054)
return


}, 1)

tmp9055 := PrimHead(V5861)

__e.TailApply(tmp8806, tmp9055)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5862)
return
}


}, 1)

tmp9058 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5861)
return
}, 0)

__e.TailApply(tmp8805, tmp9058)
return


}, 1)

tmp9059 := Call(__e, ns2_1set, symshen_4macros, tmp8804)


_ = tmp9059

tmp9060 := MakeNative(func(__e *ControlFlow) {
V5865 := __e.Get(1)
_ = V5865
V5866 := __e.Get(2)
_ = V5866
tmp9061 := MakeNative(func(__e *ControlFlow) {
W5867 := __e.Get(1)
_ = W5867
tmp9062 := MakeNative(func(__e *ControlFlow) {
W5868 := __e.Get(1)
_ = W5868
tmp9063 := MakeNative(func(__e *ControlFlow) {
W5869 := __e.Get(1)
_ = W5869
__e.Return(V5865)
return
}, 1)

tmp9064 := Call(__e, PrimFunc(symfn), V5865)


tmp9065 := Call(__e, PrimFunc(symshen_4record_1macro), V5865, tmp9064)


__e.TailApply(tmp9063, tmp9065)
return


}, 1)

tmp9066 := Call(__e, PrimFunc(symappend), V5866, W5867)


tmp9067 := PrimCons(V5865, tmp9066)

tmp9068 := PrimCons(symdefine, tmp9067)

tmp9069 := Call(__e, PrimFunc(symeval), tmp9068)


__e.TailApply(tmp9062, tmp9069)
return


}, 1)

tmp9070 := PrimCons(symX, Nil)

tmp9071 := PrimCons(sym_1_6, tmp9070)

tmp9072 := PrimCons(symX, tmp9071)

__e.TailApply(tmp9061, tmp9072)
return


}, 2)

tmp9073 := Call(__e, ns2_1set, symshen_4process_1def, tmp9060)


_ = tmp9073

tmp9074 := MakeNative(func(__e *ControlFlow) {
V5870 := __e.Get(1)
_ = V5870
tmp9114 := PrimIsPair(V5870)

var ifres9088 Obj

if True == tmp9114 {
tmp9112 := PrimHead(V5870)

tmp9113 := PrimEqual(symlet, tmp9112)

var ifres9090 Obj

if True == tmp9113 {
tmp9110 := PrimTail(V5870)

tmp9111 := PrimIsPair(tmp9110)

var ifres9092 Obj

if True == tmp9111 {
tmp9107 := PrimTail(V5870)

tmp9108 := PrimTail(tmp9107)

tmp9109 := PrimIsPair(tmp9108)

var ifres9094 Obj

if True == tmp9109 {
tmp9103 := PrimTail(V5870)

tmp9104 := PrimTail(tmp9103)

tmp9105 := PrimTail(tmp9104)

tmp9106 := PrimIsPair(tmp9105)

var ifres9096 Obj

if True == tmp9106 {
tmp9098 := PrimTail(V5870)

tmp9099 := PrimTail(tmp9098)

tmp9100 := PrimTail(tmp9099)

tmp9101 := PrimTail(tmp9100)

tmp9102 := PrimIsPair(tmp9101)

var ifres9097 Obj

if True == tmp9102 {
ifres9097 = True


} else {
ifres9097 = False


}

ifres9096 = ifres9097


} else {
ifres9096 = False


}

var ifres9095 Obj

if True == ifres9096 {
ifres9095 = True


} else {
ifres9095 = False


}

ifres9094 = ifres9095


} else {
ifres9094 = False


}

var ifres9093 Obj

if True == ifres9094 {
ifres9093 = True


} else {
ifres9093 = False


}

ifres9092 = ifres9093


} else {
ifres9092 = False


}

var ifres9091 Obj

if True == ifres9092 {
ifres9091 = True


} else {
ifres9091 = False


}

ifres9090 = ifres9091


} else {
ifres9090 = False


}

var ifres9089 Obj

if True == ifres9090 {
ifres9089 = True


} else {
ifres9089 = False


}

ifres9088 = ifres9089


} else {
ifres9088 = False


}

if True == ifres9088 {
tmp9075 := PrimTail(V5870)

tmp9076 := PrimHead(tmp9075)

tmp9077 := PrimTail(V5870)

tmp9078 := PrimTail(tmp9077)

tmp9079 := PrimHead(tmp9078)

tmp9080 := PrimTail(V5870)

tmp9081 := PrimTail(tmp9080)

tmp9082 := PrimTail(tmp9081)

tmp9083 := PrimCons(symlet, tmp9082)

tmp9084 := PrimCons(tmp9083, Nil)

tmp9085 := PrimCons(tmp9079, tmp9084)

tmp9086 := PrimCons(tmp9076, tmp9085)

__e.Return(PrimCons(symlet, tmp9086))
return


} else {
__e.Return(V5870)
return
}


}, 1)

tmp9115 := Call(__e, ns2_1set, symshen_4process_1let, tmp9074)


_ = tmp9115

tmp9116 := MakeNative(func(__e *ControlFlow) {
V5871 := __e.Get(1)
_ = V5871
tmp9117 := MakeNative(func(__e *ControlFlow) {
GoTo5873 := __e.Get(1)
_ = GoTo5873
tmp9152 := PrimIsPair(V5871)

if True == tmp9152 {
tmp9118 := MakeNative(func(__e *ControlFlow) {
Select5880 := __e.Get(1)
_ = Select5880
tmp9148 := PrimHead(V5871)

tmp9149 := PrimEqual(sym_8s, tmp9148)

if True == tmp9149 {
tmp9146 := PrimIsPair(Select5880)

if True == tmp9146 {
tmp9119 := MakeNative(func(__e *ControlFlow) {
Select5878 := __e.Get(1)
_ = Select5878
tmp9120 := MakeNative(func(__e *ControlFlow) {
Select5879 := __e.Get(1)
_ = Select5879
tmp9142 := PrimIsPair(Select5879)

if True == tmp9142 {
tmp9121 := MakeNative(func(__e *ControlFlow) {
Select5877 := __e.Get(1)
_ = Select5877
tmp9139 := PrimIsPair(Select5877)

if True == tmp9139 {
tmp9122 := PrimCons(sym_8s, Select5879)

tmp9123 := Call(__e, PrimFunc(symshen_4process_1_8s), tmp9122)


tmp9124 := PrimCons(tmp9123, Nil)

tmp9125 := PrimCons(Select5878, tmp9124)

__e.Return(PrimCons(sym_8s, tmp9125))
return


} else {
tmp9137 := PrimEqual(Nil, Select5877)

var ifres9134 Obj

if True == tmp9137 {
tmp9136 := PrimIsString(Select5878)

var ifres9135 Obj

if True == tmp9136 {
ifres9135 = True


} else {
ifres9135 = False


}

ifres9134 = ifres9135


} else {
ifres9134 = False


}

if True == ifres9134 {
tmp9126 := MakeNative(func(__e *ControlFlow) {
W5872 := __e.Get(1)
_ = W5872
tmp9130 := Call(__e, PrimFunc(symlength), W5872)


tmp9131 := PrimGreatThan(tmp9130, MakeNumber(1))

if True == tmp9131 {
tmp9127 := Call(__e, PrimFunc(symappend), W5872, Select5879)


tmp9128 := PrimCons(sym_8s, tmp9127)

__e.TailApply(PrimFunc(symshen_4process_1_8s), tmp9128)
return


} else {
__e.Return(V5871)
return
}


}, 1)

tmp9132 := Call(__e, PrimFunc(symexplode), Select5878)


__e.TailApply(tmp9126, tmp9132)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5873)
return
}


}


}, 1)

tmp9140 := PrimTail(Select5879)

__e.TailApply(tmp9121, tmp9140)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5873)
return
}


}, 1)

tmp9143 := PrimTail(Select5880)

__e.TailApply(tmp9120, tmp9143)
return


}, 1)

tmp9144 := PrimHead(Select5880)

__e.TailApply(tmp9119, tmp9144)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5873)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5873)
return
}


}, 1)

tmp9150 := PrimTail(V5871)

__e.TailApply(tmp9118, tmp9150)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5873)
return
}


}, 1)

tmp9153 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5871)
return
}, 0)

__e.TailApply(tmp9117, tmp9153)
return


}, 1)

tmp9154 := Call(__e, ns2_1set, symshen_4process_1_8s, tmp9116)


_ = tmp9154

tmp9155 := MakeNative(func(__e *ControlFlow) {
V5881 := __e.Get(1)
_ = V5881
V5882 := __e.Get(2)
_ = V5882
tmp9156 := MakeNative(func(__e *ControlFlow) {
W5883 := __e.Get(1)
_ = W5883
tmp9157 := MakeNative(func(__e *ControlFlow) {
W5884 := __e.Get(1)
_ = W5884
__e.Return(W5883)
return
}, 1)

tmp9158 := MakeNative(func(__e *ControlFlow) {
Z5885 := __e.Get(1)
_ = Z5885
__e.TailApply(PrimFunc(symshen_4_5datatype_6), Z5885)
return
}, 1)

tmp9159 := PrimCons(W5883, V5882)

tmp9160 := Call(__e, PrimFunc(symcompile), tmp9158, tmp9159)


__e.TailApply(tmp9157, tmp9160)
return


}, 1)

tmp9161 := Call(__e, PrimFunc(symshen_4intern_1type), V5881)


__e.TailApply(tmp9156, tmp9161)
return


}, 2)

tmp9162 := Call(__e, ns2_1set, symshen_4process_1datatype, tmp9155)


_ = tmp9162

tmp9163 := MakeNative(func(__e *ControlFlow) {
V5886 := __e.Get(1)
_ = V5886
tmp9164 := PrimStr(V5886)

tmp9165 := PrimStringConcat(tmp9164, MakeString("#type"))

__e.Return(PrimIntern(tmp9165))
return


}, 1)

tmp9166 := Call(__e, ns2_1set, symshen_4intern_1type, tmp9163)


_ = tmp9166

tmp9167 := MakeNative(func(__e *ControlFlow) {
V5887 := __e.Get(1)
_ = V5887
tmp9168 := PrimValue(symshen_4_dsynonyms_d)

tmp9169 := Call(__e, PrimFunc(symappend), V5887, tmp9168)


tmp9170 := PrimSet(symshen_4_dsynonyms_d, tmp9169)

__e.TailApply(PrimFunc(symshen_4synonyms_1h), tmp9170)
return


}, 1)

tmp9171 := Call(__e, ns2_1set, symshen_4process_1synonyms, tmp9167)


_ = tmp9171

tmp9172 := MakeNative(func(__e *ControlFlow) {
V5888 := __e.Get(1)
_ = V5888
tmp9173 := MakeNative(func(__e *ControlFlow) {
W5889 := __e.Get(1)
_ = W5889
tmp9174 := MakeNative(func(__e *ControlFlow) {
W5891 := __e.Get(1)
_ = W5891
__e.Return(symsynonyms)
return
}, 1)

tmp9175 := Call(__e, PrimFunc(symshen_4compile_1synonyms), W5889)


tmp9176 := PrimCons(symshen_4demod, tmp9175)

tmp9177 := PrimCons(symdefine, tmp9176)

tmp9178 := Call(__e, PrimFunc(symeval), tmp9177)


__e.TailApply(tmp9174, tmp9178)
return


}, 1)

tmp9179 := MakeNative(func(__e *ControlFlow) {
Z5890 := __e.Get(1)
_ = Z5890
__e.TailApply(PrimFunc(symshen_4curry_1type), Z5890)
return
}, 1)

tmp9180 := Call(__e, PrimFunc(symmap), tmp9179, V5888)


__e.TailApply(tmp9173, tmp9180)
return


}, 1)

tmp9181 := Call(__e, ns2_1set, symshen_4synonyms_1h, tmp9172)


_ = tmp9181

tmp9182 := MakeNative(func(__e *ControlFlow) {
V5894 := __e.Get(1)
_ = V5894
tmp9204 := PrimEqual(Nil, V5894)

if True == tmp9204 {
tmp9183 := MakeNative(func(__e *ControlFlow) {
W5895 := __e.Get(1)
_ = W5895
tmp9184 := PrimCons(W5895, Nil)

tmp9185 := PrimCons(sym_1_6, tmp9184)

__e.Return(PrimCons(W5895, tmp9185))
return


}, 1)

tmp9186 := Call(__e, PrimFunc(symgensym), symX)


__e.TailApply(tmp9183, tmp9186)
return


} else {
tmp9202 := PrimIsPair(V5894)

var ifres9198 Obj

if True == tmp9202 {
tmp9200 := PrimTail(V5894)

tmp9201 := PrimIsPair(tmp9200)

var ifres9199 Obj

if True == tmp9201 {
ifres9199 = True


} else {
ifres9199 = False


}

ifres9198 = ifres9199


} else {
ifres9198 = False


}

if True == ifres9198 {
tmp9187 := PrimHead(V5894)

tmp9188 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9187)


tmp9189 := PrimTail(V5894)

tmp9190 := PrimHead(tmp9189)

tmp9191 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9190)


tmp9192 := PrimTail(V5894)

tmp9193 := PrimTail(tmp9192)

tmp9194 := Call(__e, PrimFunc(symshen_4compile_1synonyms), tmp9193)


tmp9195 := PrimCons(tmp9191, tmp9194)

tmp9196 := PrimCons(sym_1_6, tmp9195)

__e.Return(PrimCons(tmp9188, tmp9196))
return


} else {
__e.Return(PrimSimpleError(MakeString("synonyms requires an even number of arguments\n")))
return
}


}


}, 1)

tmp9205 := Call(__e, ns2_1set, symshen_4compile_1synonyms, tmp9182)


_ = tmp9205

tmp9206 := MakeNative(func(__e *ControlFlow) {
V5896 := __e.Get(1)
_ = V5896
tmp9207 := MakeNative(func(__e *ControlFlow) {
GoTo5897 := __e.Get(1)
_ = GoTo5897
tmp9235 := PrimIsPair(V5896)

if True == tmp9235 {
tmp9208 := MakeNative(func(__e *ControlFlow) {
Select5904 := __e.Get(1)
_ = Select5904
tmp9231 := PrimHead(V5896)

tmp9232 := PrimEqual(sym_c_4, tmp9231)

if True == tmp9232 {
tmp9229 := PrimIsPair(Select5904)

if True == tmp9229 {
tmp9209 := MakeNative(func(__e *ControlFlow) {
Select5902 := __e.Get(1)
_ = Select5902
tmp9210 := MakeNative(func(__e *ControlFlow) {
Select5903 := __e.Get(1)
_ = Select5903
tmp9225 := PrimIsPair(Select5903)

if True == tmp9225 {
tmp9211 := MakeNative(func(__e *ControlFlow) {
Select5901 := __e.Get(1)
_ = Select5901
tmp9222 := PrimIsPair(Select5901)

if True == tmp9222 {
tmp9212 := PrimCons(sym_c_4, Select5903)

tmp9213 := Call(__e, PrimFunc(symshen_4process_1lambda), tmp9212)


tmp9214 := PrimCons(tmp9213, Nil)

tmp9215 := PrimCons(Select5902, tmp9214)

__e.Return(PrimCons(symlambda, tmp9215))
return


} else {
tmp9220 := PrimEqual(Nil, Select5901)

if True == tmp9220 {
tmp9218 := PrimIsVariable(Select5902)

if True == tmp9218 {
__e.Return(PrimCons(symlambda, Select5904))
return
} else {
tmp9216 := Call(__e, PrimFunc(symshen_4app), Select5902, MakeString(" is not a variable\n"), symshen_4s)


__e.Return(PrimSimpleError(tmp9216))
return


}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5897)
return
}


}


}, 1)

tmp9223 := PrimTail(Select5903)

__e.TailApply(tmp9211, tmp9223)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5897)
return
}


}, 1)

tmp9226 := PrimTail(Select5904)

__e.TailApply(tmp9210, tmp9226)
return


}, 1)

tmp9227 := PrimHead(Select5904)

__e.TailApply(tmp9209, tmp9227)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5897)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5897)
return
}


}, 1)

tmp9233 := PrimTail(V5896)

__e.TailApply(tmp9208, tmp9233)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5897)
return
}


}, 1)

tmp9236 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5896)
return
}, 0)

__e.TailApply(tmp9207, tmp9236)
return


}, 1)

tmp9237 := Call(__e, ns2_1set, symshen_4process_1lambda, tmp9206)


_ = tmp9237

tmp9238 := MakeNative(func(__e *ControlFlow) {
V5907 := __e.Get(1)
_ = V5907
tmp9239 := MakeNative(func(__e *ControlFlow) {
GoTo5908 := __e.Get(1)
_ = GoTo5908
tmp9279 := PrimIsPair(V5907)

if True == tmp9279 {
tmp9240 := MakeNative(func(__e *ControlFlow) {
Select5916 := __e.Get(1)
_ = Select5916
tmp9275 := PrimHead(V5907)

tmp9276 := PrimEqual(symcases, tmp9275)

if True == tmp9276 {
tmp9273 := PrimIsPair(Select5916)

if True == tmp9273 {
tmp9241 := MakeNative(func(__e *ControlFlow) {
Select5914 := __e.Get(1)
_ = Select5914
tmp9242 := MakeNative(func(__e *ControlFlow) {
Select5915 := __e.Get(1)
_ = Select5915
tmp9269 := PrimEqual(True, Select5914)

var ifres9266 Obj

if True == tmp9269 {
tmp9268 := PrimIsPair(Select5915)

var ifres9267 Obj

if True == tmp9268 {
ifres9267 = True


} else {
ifres9267 = False


}

ifres9266 = ifres9267


} else {
ifres9266 = False


}

if True == ifres9266 {
__e.Return(PrimHead(Select5915))
return
} else {
tmp9243 := MakeNative(func(__e *ControlFlow) {
GoTo5911 := __e.Get(1)
_ = GoTo5911
tmp9261 := PrimIsPair(Select5915)

if True == tmp9261 {
tmp9244 := MakeNative(func(__e *ControlFlow) {
Select5912 := __e.Get(1)
_ = Select5912
tmp9245 := MakeNative(func(__e *ControlFlow) {
Select5913 := __e.Get(1)
_ = Select5913
tmp9257 := PrimEqual(Nil, Select5913)

if True == tmp9257 {
tmp9246 := PrimCons(MakeString("error: cases exhausted"), Nil)

tmp9247 := PrimCons(symsimple_1error, tmp9246)

tmp9248 := PrimCons(tmp9247, Nil)

tmp9249 := PrimCons(Select5912, tmp9248)

tmp9250 := PrimCons(Select5914, tmp9249)

__e.Return(PrimCons(symif, tmp9250))
return


} else {
tmp9251 := PrimCons(symcases, Select5913)

tmp9252 := Call(__e, PrimFunc(symshen_4process_1cases), tmp9251)


tmp9253 := PrimCons(tmp9252, Nil)

tmp9254 := PrimCons(Select5912, tmp9253)

tmp9255 := PrimCons(Select5914, tmp9254)

__e.Return(PrimCons(symif, tmp9255))
return


}


}, 1)

tmp9258 := PrimTail(Select5915)

__e.TailApply(tmp9245, tmp9258)
return


}, 1)

tmp9259 := PrimHead(Select5915)

__e.TailApply(tmp9244, tmp9259)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5911)
return
}


}, 1)

tmp9262 := MakeNative(func(__e *ControlFlow) {
tmp9264 := PrimEqual(Nil, Select5915)

if True == tmp9264 {
__e.Return(PrimSimpleError(MakeString("error: odd number of case elements\n")))
return
} else {
__e.TailApply(PrimFunc(symthaw), GoTo5908)
return
}


}, 0)

__e.TailApply(tmp9243, tmp9262)
return


}


}, 1)

tmp9270 := PrimTail(Select5916)

__e.TailApply(tmp9242, tmp9270)
return


}, 1)

tmp9271 := PrimHead(Select5916)

__e.TailApply(tmp9241, tmp9271)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5908)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5908)
return
}


}, 1)

tmp9277 := PrimTail(V5907)

__e.TailApply(tmp9240, tmp9277)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5908)
return
}


}, 1)

tmp9280 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5907)
return
}, 0)

__e.TailApply(tmp9239, tmp9280)
return


}, 1)

tmp9281 := Call(__e, ns2_1set, symshen_4process_1cases, tmp9238)


_ = tmp9281

tmp9282 := MakeNative(func(__e *ControlFlow) {
V5917 := __e.Get(1)
_ = V5917
tmp9283 := PrimCons(symrun, Nil)

tmp9284 := PrimCons(symget_1time, tmp9283)

tmp9285 := PrimCons(symrun, Nil)

tmp9286 := PrimCons(symget_1time, tmp9285)

tmp9287 := PrimCons(symStart, Nil)

tmp9288 := PrimCons(symFinish, tmp9287)

tmp9289 := PrimCons(sym_1, tmp9288)

tmp9290 := PrimCons(symTime, Nil)

tmp9291 := PrimCons(symstr, tmp9290)

tmp9292 := PrimCons(MakeString(" secs\n"), Nil)

tmp9293 := PrimCons(tmp9291, tmp9292)

tmp9294 := PrimCons(symcn, tmp9293)

tmp9295 := PrimCons(tmp9294, Nil)

tmp9296 := PrimCons(MakeString("\nrun time: "), tmp9295)

tmp9297 := PrimCons(symcn, tmp9296)

tmp9298 := PrimCons(symstoutput, Nil)

tmp9299 := PrimCons(tmp9298, Nil)

tmp9300 := PrimCons(tmp9297, tmp9299)

tmp9301 := PrimCons(sympr, tmp9300)

tmp9302 := PrimCons(symResult, Nil)

tmp9303 := PrimCons(tmp9301, tmp9302)

tmp9304 := PrimCons(symMessage, tmp9303)

tmp9305 := PrimCons(tmp9289, tmp9304)

tmp9306 := PrimCons(symTime, tmp9305)

tmp9307 := PrimCons(tmp9286, tmp9306)

tmp9308 := PrimCons(symFinish, tmp9307)

tmp9309 := PrimCons(V5917, tmp9308)

tmp9310 := PrimCons(symResult, tmp9309)

tmp9311 := PrimCons(tmp9284, tmp9310)

tmp9312 := PrimCons(symStart, tmp9311)

__e.Return(PrimCons(symlet, tmp9312))
return


}, 1)

tmp9313 := Call(__e, ns2_1set, symshen_4process_1time, tmp9282)


_ = tmp9313

tmp9314 := MakeNative(func(__e *ControlFlow) {
V5918 := __e.Get(1)
_ = V5918
tmp9340 := PrimIsPair(V5918)

var ifres9325 Obj

if True == tmp9340 {
tmp9338 := PrimTail(V5918)

tmp9339 := PrimIsPair(tmp9338)

var ifres9327 Obj

if True == tmp9339 {
tmp9335 := PrimTail(V5918)

tmp9336 := PrimTail(tmp9335)

tmp9337 := PrimIsPair(tmp9336)

var ifres9329 Obj

if True == tmp9337 {
tmp9331 := PrimTail(V5918)

tmp9332 := PrimTail(tmp9331)

tmp9333 := PrimTail(tmp9332)

tmp9334 := PrimIsPair(tmp9333)

var ifres9330 Obj

if True == tmp9334 {
ifres9330 = True


} else {
ifres9330 = False


}

ifres9329 = ifres9330


} else {
ifres9329 = False


}

var ifres9328 Obj

if True == ifres9329 {
ifres9328 = True


} else {
ifres9328 = False


}

ifres9327 = ifres9328


} else {
ifres9327 = False


}

var ifres9326 Obj

if True == ifres9327 {
ifres9326 = True


} else {
ifres9326 = False


}

ifres9325 = ifres9326


} else {
ifres9325 = False


}

if True == ifres9325 {
tmp9315 := PrimHead(V5918)

tmp9316 := PrimTail(V5918)

tmp9317 := PrimHead(tmp9316)

tmp9318 := PrimHead(V5918)

tmp9319 := PrimTail(V5918)

tmp9320 := PrimTail(tmp9319)

tmp9321 := PrimCons(tmp9318, tmp9320)

tmp9322 := PrimCons(tmp9321, Nil)

tmp9323 := PrimCons(tmp9317, tmp9322)

__e.Return(PrimCons(tmp9315, tmp9323))
return


} else {
__e.Return(V5918)
return
}


}, 1)

tmp9341 := Call(__e, ns2_1set, symshen_4process_1assoc, tmp9314)


_ = tmp9341

tmp9342 := MakeNative(func(__e *ControlFlow) {
V5919 := __e.Get(1)
_ = V5919
tmp9343 := PrimStr(V5919)

tmp9344 := Call(__e, PrimFunc(symshen_4mu_1h), tmp9343)


__e.Return(PrimIntern(tmp9344))
return


}, 1)

tmp9345 := Call(__e, ns2_1set, symshen_4make_1uppercase, tmp9342)


_ = tmp9345

tmp9346 := MakeNative(func(__e *ControlFlow) {
V5920 := __e.Get(1)
_ = V5920
tmp9365 := PrimEqual(MakeString(""), V5920)

if True == tmp9365 {
__e.Return(MakeString(""))
return
} else {
tmp9363 := Call(__e, PrimFunc(symshen_4_7string_2), V5920)


if True == tmp9363 {
tmp9347 := MakeNative(func(__e *ControlFlow) {
W5921 := __e.Get(1)
_ = W5921
tmp9348 := MakeNative(func(__e *ControlFlow) {
W5922 := __e.Get(1)
_ = W5922
tmp9349 := MakeNative(func(__e *ControlFlow) {
W5923 := __e.Get(1)
_ = W5923
tmp9350 := PrimTailString(V5920)

tmp9351 := Call(__e, PrimFunc(symshen_4mu_1h), tmp9350)


__e.TailApply(PrimFunc(sym_8s), W5923, tmp9351)
return


}, 1)

tmp9358 := PrimGreatEqual(W5921, MakeNumber(97))

var ifres9355 Obj

if True == tmp9358 {
tmp9357 := PrimLessEqual(W5921, MakeNumber(122))

var ifres9356 Obj

if True == tmp9357 {
ifres9356 = True


} else {
ifres9356 = False


}

ifres9355 = ifres9356


} else {
ifres9355 = False


}

var ifres9352 Obj

if True == ifres9355 {
tmp9353 := PrimNumberToString(W5922)

ifres9352 = tmp9353


} else {
tmp9354 := Call(__e, PrimFunc(symhdstr), V5920)


ifres9352 = tmp9354


}

__e.TailApply(tmp9349, ifres9352)
return


}, 1)

tmp9359 := PrimNumberSubtract(W5921, MakeNumber(32))

__e.TailApply(tmp9348, tmp9359)
return


}, 1)

tmp9360 := Call(__e, PrimFunc(symhdstr), V5920)


tmp9361 := PrimStringToNumber(tmp9360)

__e.TailApply(tmp9347, tmp9361)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.mu-h")))
return
}


}


}, 1)

tmp9366 := Call(__e, ns2_1set, symshen_4mu_1h, tmp9346)


_ = tmp9366

tmp9367 := MakeNative(func(__e *ControlFlow) {
V5924 := __e.Get(1)
_ = V5924
V5925 := __e.Get(2)
_ = V5925
tmp9368 := PrimValue(sym_dmacros_d)

tmp9369 := Call(__e, PrimFunc(symshen_4update_1assoc), V5924, V5925, tmp9368)


__e.Return(PrimSet(sym_dmacros_d, tmp9369))
return


}, 2)

tmp9370 := Call(__e, ns2_1set, symshen_4record_1macro, tmp9367)


_ = tmp9370

tmp9371 := MakeNative(func(__e *ControlFlow) {
V5935 := __e.Get(1)
_ = V5935
V5936 := __e.Get(2)
_ = V5936
V5937 := __e.Get(3)
_ = V5937
tmp9391 := PrimEqual(Nil, V5937)

if True == tmp9391 {
tmp9372 := PrimCons(V5935, V5936)

__e.Return(PrimCons(tmp9372, Nil))
return


} else {
tmp9373 := MakeNative(func(__e *ControlFlow) {
GoTo5938 := __e.Get(1)
_ = GoTo5938
tmp9388 := PrimIsPair(V5937)

if True == tmp9388 {
tmp9374 := MakeNative(func(__e *ControlFlow) {
Select5939 := __e.Get(1)
_ = Select5939
tmp9375 := MakeNative(func(__e *ControlFlow) {
Select5940 := __e.Get(1)
_ = Select5940
tmp9384 := PrimIsPair(Select5939)

var ifres9380 Obj

if True == tmp9384 {
tmp9382 := PrimHead(Select5939)

tmp9383 := PrimEqual(V5935, tmp9382)

var ifres9381 Obj

if True == tmp9383 {
ifres9381 = True


} else {
ifres9381 = False


}

ifres9380 = ifres9381


} else {
ifres9380 = False


}

if True == ifres9380 {
tmp9376 := PrimHead(Select5939)

tmp9377 := PrimCons(tmp9376, V5936)

__e.Return(PrimCons(tmp9377, Select5940))
return


} else {
tmp9378 := Call(__e, PrimFunc(symshen_4update_1assoc), V5935, V5936, Select5940)


__e.Return(PrimCons(Select5939, tmp9378))
return


}


}, 1)

tmp9385 := PrimTail(V5937)

__e.TailApply(tmp9375, tmp9385)
return


}, 1)

tmp9386 := PrimHead(V5937)

__e.TailApply(tmp9374, tmp9386)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5938)
return
}


}, 1)

tmp9389 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.update-assoc")))
return
}, 0)

__e.TailApply(tmp9373, tmp9389)
return


}


}, 3)

tmp9392 := Call(__e, ns2_1set, symshen_4update_1assoc, tmp9371)


_ = tmp9392

tmp9393 := MakeNative(func(__e *ControlFlow) {
tmp9401 := Call(__e, PrimFunc(symstinput))


tmp9402 := Call(__e, PrimFunc(symshen_4char_1stinput_2), tmp9401)


if True == tmp9402 {
tmp9394 := PrimCons(symstinput, Nil)

tmp9395 := PrimCons(tmp9394, Nil)

tmp9396 := PrimCons(symshen_4read_1unit_1string, tmp9395)

tmp9397 := PrimCons(tmp9396, Nil)

__e.Return(PrimCons(symstring_1_6n, tmp9397))
return


} else {
tmp9398 := PrimCons(symstinput, Nil)

tmp9399 := PrimCons(tmp9398, Nil)

__e.Return(PrimCons(symread_1byte, tmp9399))
return


}


}, 0)

tmp9403 := Call(__e, ns2_1set, symshen_4process_1read_1byte, tmp9393)


_ = tmp9403

tmp9404 := MakeNative(func(__e *ControlFlow) {
V5941 := __e.Get(1)
_ = V5941
tmp9405 := MakeNative(func(__e *ControlFlow) {
W5942 := __e.Get(1)
_ = W5942
tmp9406 := MakeNative(func(__e *ControlFlow) {
W5943 := __e.Get(1)
_ = W5943
tmp9407 := MakeNative(func(__e *ControlFlow) {
W5944 := __e.Get(1)
_ = W5944
tmp9408 := MakeNative(func(__e *ControlFlow) {
W5945 := __e.Get(1)
_ = W5945
tmp9409 := MakeNative(func(__e *ControlFlow) {
W5946 := __e.Get(1)
_ = W5946
tmp9410 := MakeNative(func(__e *ControlFlow) {
W5948 := __e.Get(1)
_ = W5948
tmp9411 := MakeNative(func(__e *ControlFlow) {
W5949 := __e.Get(1)
_ = W5949
tmp9412 := MakeNative(func(__e *ControlFlow) {
W5950 := __e.Get(1)
_ = W5950
tmp9413 := MakeNative(func(__e *ControlFlow) {
W5951 := __e.Get(1)
_ = W5951
tmp9414 := MakeNative(func(__e *ControlFlow) {
W5952 := __e.Get(1)
_ = W5952
tmp9415 := MakeNative(func(__e *ControlFlow) {
W5953 := __e.Get(1)
_ = W5953
tmp9416 := PrimCons(W5945, Nil)

tmp9417 := PrimCons(W5944, tmp9416)

tmp9418 := PrimCons(W5943, tmp9417)

tmp9419 := PrimCons(W5942, tmp9418)

__e.Return(PrimCons(W5953, tmp9419))
return


}, 1)

tmp9420 := Call(__e, PrimFunc(symshen_4continue), W5948, W5946, W5949, W5950, W5951, W5952)


tmp9421 := PrimCons(tmp9420, Nil)

tmp9422 := PrimCons(W5952, tmp9421)

tmp9423 := PrimCons(symlambda, tmp9422)

tmp9424 := PrimCons(tmp9423, Nil)

tmp9425 := PrimCons(W5951, tmp9424)

tmp9426 := PrimCons(symlambda, tmp9425)

tmp9427 := PrimCons(tmp9426, Nil)

tmp9428 := PrimCons(W5950, tmp9427)

tmp9429 := PrimCons(symlambda, tmp9428)

tmp9430 := PrimCons(tmp9429, Nil)

tmp9431 := PrimCons(W5949, tmp9430)

tmp9432 := PrimCons(symlambda, tmp9431)

__e.TailApply(tmp9415, tmp9432)
return


}, 1)

tmp9433 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp9414, tmp9433)
return


}, 1)

tmp9434 := Call(__e, PrimFunc(symgensym), symK)


__e.TailApply(tmp9413, tmp9434)
return


}, 1)

tmp9435 := Call(__e, PrimFunc(symgensym), symL)


__e.TailApply(tmp9412, tmp9435)
return


}, 1)

tmp9436 := Call(__e, PrimFunc(symgensym), symV)


__e.TailApply(tmp9411, tmp9436)
return


}, 1)

tmp9437 := Call(__e, PrimFunc(symshen_4received), V5941)


__e.TailApply(tmp9410, tmp9437)
return


}, 1)

tmp9438 := MakeNative(func(__e *ControlFlow) {
Z5947 := __e.Get(1)
_ = Z5947
__e.TailApply(PrimFunc(symshen_4_5body_6), Z5947)
return
}, 1)

tmp9439 := Call(__e, PrimFunc(symcompile), tmp9438, V5941)


__e.TailApply(tmp9409, tmp9439)
return


}, 1)

tmp9440 := PrimCons(True, Nil)

tmp9441 := PrimCons(symfreeze, tmp9440)

__e.TailApply(tmp9408, tmp9441)
return


}, 1)

__e.TailApply(tmp9407, MakeNumber(0))
return


}, 1)

tmp9442 := PrimCons(MakeNumber(0), Nil)

tmp9443 := PrimCons(symvector, tmp9442)

tmp9444 := PrimCons(tmp9443, Nil)

tmp9445 := PrimCons(MakeNumber(0), tmp9444)

tmp9446 := PrimCons(True, tmp9445)

tmp9447 := PrimCons(sym_8v, tmp9446)

__e.TailApply(tmp9406, tmp9447)
return


}, 1)

tmp9448 := PrimCons(symshen_4prolog_1vector, Nil)

__e.TailApply(tmp9405, tmp9448)
return


}, 1)

tmp9449 := Call(__e, ns2_1set, symshen_4call_1prolog, tmp9404)


_ = tmp9449

tmp9450 := MakeNative(func(__e *ControlFlow) {
V5956 := __e.Get(1)
_ = V5956
tmp9451 := MakeNative(func(__e *ControlFlow) {
GoTo5957 := __e.Get(1)
_ = GoTo5957
tmp9468 := PrimIsPair(V5956)

if True == tmp9468 {
tmp9452 := MakeNative(func(__e *ControlFlow) {
Select5958 := __e.Get(1)
_ = Select5958
tmp9453 := MakeNative(func(__e *ControlFlow) {
Select5959 := __e.Get(1)
_ = Select5959
tmp9464 := PrimEqual(symreceive, Select5958)

var ifres9457 Obj

if True == tmp9464 {
tmp9463 := PrimIsPair(Select5959)

var ifres9459 Obj

if True == tmp9463 {
tmp9461 := PrimTail(Select5959)

tmp9462 := PrimEqual(Nil, tmp9461)

var ifres9460 Obj

if True == tmp9462 {
ifres9460 = True


} else {
ifres9460 = False


}

ifres9459 = ifres9460


} else {
ifres9459 = False


}

var ifres9458 Obj

if True == ifres9459 {
ifres9458 = True


} else {
ifres9458 = False


}

ifres9457 = ifres9458


} else {
ifres9457 = False


}

if True == ifres9457 {
__e.Return(Select5959)
return
} else {
tmp9454 := Call(__e, PrimFunc(symshen_4received), Select5958)


tmp9455 := Call(__e, PrimFunc(symshen_4received), Select5959)


__e.TailApply(PrimFunc(symunion), tmp9454, tmp9455)
return


}


}, 1)

tmp9465 := PrimTail(V5956)

__e.TailApply(tmp9453, tmp9465)
return


}, 1)

tmp9466 := PrimHead(V5956)

__e.TailApply(tmp9452, tmp9466)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5957)
return
}


}, 1)

tmp9469 := MakeNative(func(__e *ControlFlow) {
__e.Return(Nil)
return
}, 0)

__e.TailApply(tmp9451, tmp9469)
return


}, 1)

tmp9470 := Call(__e, ns2_1set, symshen_4received, tmp9450)


_ = tmp9470

tmp9471 := MakeNative(func(__e *ControlFlow) {
tmp9472 := MakeNative(func(__e *ControlFlow) {
W5960 := __e.Get(1)
_ = W5960
tmp9473 := MakeNative(func(__e *ControlFlow) {
W5961 := __e.Get(1)
_ = W5961
tmp9474 := MakeNative(func(__e *ControlFlow) {
W5962 := __e.Get(1)
_ = W5962
__e.Return(W5962)
return
}, 1)

tmp9475 := PrimVectorSet(W5960, MakeNumber(1), MakeNumber(2))

__e.TailApply(tmp9474, tmp9475)
return


}, 1)

tmp9476 := PrimVectorSet(W5960, MakeNumber(0), symshen_4print_1prolog_1vector)

__e.TailApply(tmp9473, tmp9476)
return


}, 1)

tmp9477 := PrimValue(symshen_4_dprolog_1memory_d)

tmp9478 := PrimAbsvector(tmp9477)

__e.TailApply(tmp9472, tmp9478)
return


}, 0)

tmp9479 := Call(__e, ns2_1set, symshen_4prolog_1vector, tmp9471)


_ = tmp9479

tmp9480 := MakeNative(func(__e *ControlFlow) {
V5963 := __e.Get(1)
_ = V5963
__e.Return(V5963)
return
}, 1)

tmp9481 := Call(__e, ns2_1set, symreceive, tmp9480)


_ = tmp9481

tmp9482 := MakeNative(func(__e *ControlFlow) {
V5964 := __e.Get(1)
_ = V5964
tmp9490 := PrimIsPair(V5964)

if True == tmp9490 {
tmp9483 := PrimHead(V5964)

tmp9484 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9483)


tmp9485 := PrimTail(V5964)

tmp9486 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9485)


tmp9487 := PrimCons(tmp9486, Nil)

tmp9488 := PrimCons(tmp9484, tmp9487)

__e.Return(PrimCons(symcons, tmp9488))
return


} else {
__e.Return(V5964)
return
}


}, 1)

tmp9491 := Call(__e, ns2_1set, symshen_4rcons__form, tmp9482)


_ = tmp9491

tmp9492 := MakeNative(func(__e *ControlFlow) {
V5965 := __e.Get(1)
_ = V5965
tmp9499 := PrimIsPair(V5965)

if True == tmp9499 {
tmp9493 := PrimHead(V5965)

tmp9494 := PrimTail(V5965)

tmp9495 := Call(__e, PrimFunc(symshen_4tuple_1up), tmp9494)


tmp9496 := PrimCons(tmp9495, Nil)

tmp9497 := PrimCons(tmp9493, tmp9496)

__e.Return(PrimCons(sym_8p, tmp9497))
return


} else {
__e.Return(V5965)
return
}


}, 1)

tmp9500 := Call(__e, ns2_1set, symshen_4tuple_1up, tmp9492)


_ = tmp9500

tmp9501 := MakeNative(func(__e *ControlFlow) {
V5966 := __e.Get(1)
_ = V5966
tmp9502 := PrimValue(sym_dmacros_d)

tmp9503 := Call(__e, PrimFunc(symassoc), V5966, tmp9502)


tmp9504 := PrimValue(sym_dmacros_d)

tmp9505 := Call(__e, PrimFunc(symremove), tmp9503, tmp9504)


tmp9506 := PrimSet(sym_dmacros_d, tmp9505)

_ = tmp9506

__e.Return(V5966)
return


}, 1)

__e.TailApply(ns2_1set, symundefmacro, tmp9501)
return




}, 0)

