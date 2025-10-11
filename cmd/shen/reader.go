package main

import . "github.com/tiancaiamao/shen-go/kl"

var ReaderMain = MakeNative(func(__e *ControlFlow) {
tmp4727 := MakeNative(func(__e *ControlFlow) {
V2162 := __e.Get(1)
_ = V2162
tmp4728 := MakeNative(func(__e *ControlFlow) {
W2163 := __e.Get(1)
_ = W2163
tmp4729 := MakeNative(func(__e *ControlFlow) {
W2164 := __e.Get(1)
_ = W2164
tmp4730 := MakeNative(func(__e *ControlFlow) {
W2166 := __e.Get(1)
_ = W2166
__e.Return(W2166)
return
}, 1)

tmp4731 := Call(__e, PrimFunc(symshen_4process_1sexprs), W2164)


__e.TailApply(tmp4730, tmp4731)
return


}, 1)

tmp4732 := MakeNative(func(__e *ControlFlow) {
Z2165 := __e.Get(1)
_ = Z2165
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2165)
return
}, 1)

tmp4733 := Call(__e, PrimFunc(symshen_4compileBytes), tmp4732, W2163)


__e.TailApply(tmp4729, tmp4733)
return


}, 1)

tmp4734 := PrimReadFileAsByteList(V2162)

__e.TailApply(tmp4728, tmp4734)
return


}, 1)

tmp4735 := Call(__e, ns2_1set, symread_1file, tmp4727)


_ = tmp4735

tmp4736 := MakeNative(func(__e *ControlFlow) {
V2167 := __e.Get(1)
_ = V2167
V2168 := __e.Get(2)
_ = V2168
tmp4737 := MakeNative(func(__e *ControlFlow) {
W2169 := __e.Get(1)
_ = W2169
tmp4743 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2169)


if True == tmp4743 {
__e.Return(PrimSimpleError(MakeString("parse failure\n")))
return
} else {
tmp4740 := Call(__e, PrimFunc(symshen_4in_1_6), W2169)


tmp4741 := PrimIsPair(tmp4740)

if True == tmp4741 {
tmp4738 := Call(__e, PrimFunc(symshen_4in_1_6), W2169)


__e.TailApply(PrimFunc(symshen_4print_1residue), tmp4738)
return


} else {
__e.TailApply(PrimFunc(symshen_4_5_1out), W2169)
return
}


}


}, 1)

tmp4744 := Call(__e, V2167, V2168)


__e.TailApply(tmp4737, tmp4744)
return


}, 2)

tmp4745 := Call(__e, ns2_1set, symshen_4compileBytes, tmp4736)


_ = tmp4745

tmp4746 := MakeNative(func(__e *ControlFlow) {
V2170 := __e.Get(1)
_ = V2170
tmp4747 := MakeNative(func(__e *ControlFlow) {
W2171 := __e.Get(1)
_ = W2171
__e.TailApply(PrimFunc(symshen_4nchars), MakeNumber(50), V2170)
return
}, 1)

tmp4748 := Call(__e, PrimFunc(symstoutput))


tmp4749 := Call(__e, PrimFunc(sympr), MakeString("syntax error here:\n"), tmp4748)


__e.TailApply(tmp4747, tmp4749)
return


}, 1)

tmp4750 := Call(__e, ns2_1set, symshen_4print_1residue, tmp4746)


_ = tmp4750

tmp4751 := MakeNative(func(__e *ControlFlow) {
V2176 := __e.Get(1)
_ = V2176
V2177 := __e.Get(2)
_ = V2177
tmp4767 := PrimEqual(MakeNumber(0), V2176)

if True == tmp4767 {
tmp4752 := Call(__e, PrimFunc(symstoutput))


tmp4753 := Call(__e, PrimFunc(sympr), MakeString(" ..."), tmp4752)


_ = tmp4753

__e.TailApply(PrimFunc(symabort))
return


} else {
tmp4765 := PrimEqual(Nil, V2177)

if True == tmp4765 {
tmp4754 := Call(__e, PrimFunc(symstoutput))


tmp4755 := Call(__e, PrimFunc(sympr), MakeString(" ..."), tmp4754)


_ = tmp4755

__e.TailApply(PrimFunc(symabort))
return


} else {
tmp4763 := PrimIsPair(V2177)

if True == tmp4763 {
tmp4756 := PrimHead(V2177)

tmp4757 := PrimNumberToString(tmp4756)

tmp4758 := Call(__e, PrimFunc(symstoutput))


tmp4759 := Call(__e, PrimFunc(sympr), tmp4757, tmp4758)


_ = tmp4759

tmp4760 := PrimNumberSubtract(V2176, MakeNumber(1))

tmp4761 := PrimTail(V2177)

__e.TailApply(PrimFunc(symshen_4nchars), tmp4760, tmp4761)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.nchars")))
return
}


}


}


}, 2)

tmp4768 := Call(__e, ns2_1set, symshen_4nchars, tmp4751)


_ = tmp4768

tmp4769 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dit_d))
return
}, 0)

tmp4770 := Call(__e, ns2_1set, symit, tmp4769)


_ = tmp4770

tmp4771 := MakeNative(func(__e *ControlFlow) {
V2178 := __e.Get(1)
_ = V2178
tmp4772 := MakeNative(func(__e *ControlFlow) {
W2179 := __e.Get(1)
_ = W2179
tmp4773 := MakeNative(func(__e *ControlFlow) {
W2180 := __e.Get(1)
_ = W2180
tmp4774 := MakeNative(func(__e *ControlFlow) {
W2181 := __e.Get(1)
_ = W2181
tmp4775 := MakeNative(func(__e *ControlFlow) {
W2182 := __e.Get(1)
_ = W2182
__e.TailApply(PrimFunc(symreverse), W2181)
return
}, 1)

tmp4776 := PrimCloseStream(W2179)

__e.TailApply(tmp4775, tmp4776)
return


}, 1)

tmp4777 := Call(__e, PrimFunc(symshen_4read_1file_1as_1bytelist_1help), W2179, W2180, Nil)


__e.TailApply(tmp4774, tmp4777)
return


}, 1)

tmp4778 := PrimReadByte(W2179)

__e.TailApply(tmp4773, tmp4778)
return


}, 1)

tmp4779 := PrimOpenStream(V2178, symin)

__e.TailApply(tmp4772, tmp4779)
return


}, 1)

tmp4780 := Call(__e, ns2_1set, symread_1file_1as_1bytelist, tmp4771)


_ = tmp4780

tmp4781 := MakeNative(func(__e *ControlFlow) {
V2183 := __e.Get(1)
_ = V2183
V2184 := __e.Get(2)
_ = V2184
V2185 := __e.Get(3)
_ = V2185
tmp4785 := PrimEqual(MakeNumber(-1), V2184)

if True == tmp4785 {
__e.Return(V2185)
return
} else {
tmp4782 := PrimReadByte(V2183)

tmp4783 := PrimCons(V2184, V2185)

__e.TailApply(PrimFunc(symshen_4read_1file_1as_1bytelist_1help), V2183, tmp4782, tmp4783)
return


}


}, 3)

tmp4786 := Call(__e, ns2_1set, symshen_4read_1file_1as_1bytelist_1help, tmp4781)


_ = tmp4786

tmp4787 := MakeNative(func(__e *ControlFlow) {
V2186 := __e.Get(1)
_ = V2186
tmp4788 := MakeNative(func(__e *ControlFlow) {
W2187 := __e.Get(1)
_ = W2187
tmp4789 := PrimReadByte(W2187)

__e.TailApply(PrimFunc(symshen_4rfas_1h), W2187, tmp4789, MakeString(""))
return


}, 1)

tmp4790 := PrimOpenStream(V2186, symin)

__e.TailApply(tmp4788, tmp4790)
return


}, 1)

tmp4791 := Call(__e, ns2_1set, symread_1file_1as_1string, tmp4787)


_ = tmp4791

tmp4792 := MakeNative(func(__e *ControlFlow) {
V2188 := __e.Get(1)
_ = V2188
V2189 := __e.Get(2)
_ = V2189
V2190 := __e.Get(3)
_ = V2190
tmp4798 := PrimEqual(MakeNumber(-1), V2189)

if True == tmp4798 {
tmp4793 := PrimCloseStream(V2188)

_ = tmp4793

__e.Return(V2190)
return


} else {
tmp4794 := PrimReadByte(V2188)

tmp4795 := PrimNumberToString(V2189)

tmp4796 := PrimStringConcat(V2190, tmp4795)

__e.TailApply(PrimFunc(symshen_4rfas_1h), V2188, tmp4794, tmp4796)
return


}


}, 3)

tmp4799 := Call(__e, ns2_1set, symshen_4rfas_1h, tmp4792)


_ = tmp4799

tmp4800 := MakeNative(func(__e *ControlFlow) {
V2191 := __e.Get(1)
_ = V2191
tmp4801 := Call(__e, PrimFunc(symread), V2191)


__e.TailApply(PrimFunc(symeval_1kl), tmp4801)
return


}, 1)

tmp4802 := Call(__e, ns2_1set, syminput, tmp4800)


_ = tmp4802

tmp4803 := MakeNative(func(__e *ControlFlow) {
V2192 := __e.Get(1)
_ = V2192
V2193 := __e.Get(2)
_ = V2193
tmp4804 := MakeNative(func(__e *ControlFlow) {
W2194 := __e.Get(1)
_ = W2194
tmp4805 := MakeNative(func(__e *ControlFlow) {
W2195 := __e.Get(1)
_ = W2195
tmp4811 := Call(__e, PrimFunc(symshen_4rectify_1type), V2192)


tmp4812 := Call(__e, PrimFunc(symshen_4typecheck), W2195, tmp4811)


tmp4813 := PrimEqual(False, tmp4812)

if True == tmp4813 {
tmp4806 := Call(__e, PrimFunc(symshen_4app), V2192, MakeString("\n"), symshen_4r)


tmp4807 := PrimStringConcat(MakeString(" is not of type "), tmp4806)

tmp4808 := Call(__e, PrimFunc(symshen_4app), W2195, tmp4807, symshen_4r)


tmp4809 := PrimStringConcat(MakeString("type error: "), tmp4808)

__e.Return(PrimSimpleError(tmp4809))
return


} else {
__e.TailApply(PrimFunc(symeval_1kl), W2195)
return
}


}, 1)

tmp4814 := Call(__e, PrimFunc(symread), V2193)


__e.TailApply(tmp4805, tmp4814)
return


}, 1)

tmp4815 := Call(__e, PrimFunc(symshen_4monotype), V2192)


__e.TailApply(tmp4804, tmp4815)
return


}, 2)

tmp4816 := Call(__e, ns2_1set, syminput_7, tmp4803)


_ = tmp4816

tmp4817 := MakeNative(func(__e *ControlFlow) {
V2196 := __e.Get(1)
_ = V2196
tmp4824 := PrimIsPair(V2196)

if True == tmp4824 {
tmp4818 := MakeNative(func(__e *ControlFlow) {
Z2197 := __e.Get(1)
_ = Z2197
__e.TailApply(PrimFunc(symshen_4monotype), Z2197)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4818, V2196)
return


} else {
tmp4822 := PrimIsVariable(V2196)

if True == tmp4822 {
tmp4819 := Call(__e, PrimFunc(symshen_4app), V2196, MakeString("\n"), symshen_4a)


tmp4820 := PrimStringConcat(MakeString("input+ expects a monotype: not "), tmp4819)

__e.Return(PrimSimpleError(tmp4820))
return


} else {
__e.Return(V2196)
return
}


}


}, 1)

tmp4825 := Call(__e, ns2_1set, symshen_4monotype, tmp4817)


_ = tmp4825

tmp4826 := MakeNative(func(__e *ControlFlow) {
V2198 := __e.Get(1)
_ = V2198
tmp4827 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2198)


tmp4828 := MakeNative(func(__e *ControlFlow) {
Z2199 := __e.Get(1)
_ = Z2199
__e.TailApply(PrimFunc(symshen_4return_2), Z2199)
return
}, 1)

__e.TailApply(PrimFunc(symshen_4read_1loop), V2198, tmp4827, Nil, tmp4828)
return


}, 1)

tmp4829 := Call(__e, ns2_1set, symlineread, tmp4826)


_ = tmp4829

tmp4830 := MakeNative(func(__e *ControlFlow) {
V2200 := __e.Get(1)
_ = V2200
tmp4831 := MakeNative(func(__e *ControlFlow) {
W2201 := __e.Get(1)
_ = W2201
tmp4832 := MakeNative(func(__e *ControlFlow) {
W2202 := __e.Get(1)
_ = W2202
tmp4833 := MakeNative(func(__e *ControlFlow) {
W2204 := __e.Get(1)
_ = W2204
__e.Return(W2204)
return
}, 1)

tmp4834 := Call(__e, PrimFunc(symshen_4process_1sexprs), W2202)


__e.TailApply(tmp4833, tmp4834)
return


}, 1)

tmp4835 := MakeNative(func(__e *ControlFlow) {
Z2203 := __e.Get(1)
_ = Z2203
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2203)
return
}, 1)

tmp4836 := Call(__e, PrimFunc(symcompile), tmp4835, W2201)


__e.TailApply(tmp4832, tmp4836)
return


}, 1)

tmp4837 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2200)


__e.TailApply(tmp4831, tmp4837)
return


}, 1)

tmp4838 := Call(__e, ns2_1set, symread_1from_1string, tmp4830)


_ = tmp4838

tmp4839 := MakeNative(func(__e *ControlFlow) {
V2205 := __e.Get(1)
_ = V2205
tmp4840 := MakeNative(func(__e *ControlFlow) {
W2206 := __e.Get(1)
_ = W2206
tmp4841 := MakeNative(func(__e *ControlFlow) {
W2207 := __e.Get(1)
_ = W2207
__e.Return(W2207)
return
}, 1)

tmp4842 := MakeNative(func(__e *ControlFlow) {
Z2208 := __e.Get(1)
_ = Z2208
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2208)
return
}, 1)

tmp4843 := Call(__e, PrimFunc(symcompile), tmp4842, W2206)


__e.TailApply(tmp4841, tmp4843)
return


}, 1)

tmp4844 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2205)


__e.TailApply(tmp4840, tmp4844)
return


}, 1)

tmp4845 := Call(__e, ns2_1set, symread_1from_1string_1unprocessed, tmp4839)


_ = tmp4845

tmp4846 := MakeNative(func(__e *ControlFlow) {
V2209 := __e.Get(1)
_ = V2209
tmp4854 := PrimEqual(MakeString(""), V2209)

if True == tmp4854 {
__e.Return(Nil)
return
} else {
tmp4852 := Call(__e, PrimFunc(symshen_4_7string_2), V2209)


if True == tmp4852 {
tmp4847 := Call(__e, PrimFunc(symhdstr), V2209)


tmp4848 := PrimStringToNumber(tmp4847)

tmp4849 := PrimTailString(V2209)

tmp4850 := Call(__e, PrimFunc(symshen_4str_1_6bytes), tmp4849)


__e.Return(PrimCons(tmp4848, tmp4850))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.str->bytes")))
return
}


}


}, 1)

tmp4855 := Call(__e, ns2_1set, symshen_4str_1_6bytes, tmp4846)


_ = tmp4855

tmp4856 := MakeNative(func(__e *ControlFlow) {
V2210 := __e.Get(1)
_ = V2210
tmp4857 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2210)


tmp4858 := MakeNative(func(__e *ControlFlow) {
Z2211 := __e.Get(1)
_ = Z2211
__e.TailApply(PrimFunc(symshen_4whitespace_2), Z2211)
return
}, 1)

tmp4859 := Call(__e, PrimFunc(symshen_4read_1loop), V2210, tmp4857, Nil, tmp4858)


__e.Return(PrimHead(tmp4859))
return


}, 1)

tmp4860 := Call(__e, ns2_1set, symread, tmp4856)


_ = tmp4860

tmp4861 := MakeNative(func(__e *ControlFlow) {
V2212 := __e.Get(1)
_ = V2212
tmp4864 := Call(__e, PrimFunc(symshen_4char_1stinput_2), V2212)


if True == tmp4864 {
tmp4862 := Call(__e, PrimFunc(symshen_4read_1unit_1string), V2212)


__e.Return(PrimStringToNumber(tmp4862))
return


} else {
__e.Return(PrimReadByte(V2212))
return
}


}, 1)

tmp4865 := Call(__e, ns2_1set, symshen_4my_1read_1byte, tmp4861)


_ = tmp4865

tmp4866 := MakeNative(func(__e *ControlFlow) {
V2217 := __e.Get(1)
_ = V2217
V2218 := __e.Get(2)
_ = V2218
V2219 := __e.Get(3)
_ = V2219
V2220 := __e.Get(4)
_ = V2220
tmp4889 := PrimEqual(MakeNumber(94), V2218)

if True == tmp4889 {
__e.Return(PrimSimpleError(MakeString("read aborted")))
return
} else {
tmp4887 := PrimEqual(MakeNumber(-1), V2218)

if True == tmp4887 {
tmp4869 := Call(__e, PrimFunc(symempty_2), V2219)


if True == tmp4869 {
__e.Return(PrimSimpleError(MakeString("error: empty stream")))
return
} else {
tmp4867 := MakeNative(func(__e *ControlFlow) {
Z2221 := __e.Get(1)
_ = Z2221
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2221)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp4867, V2219)
return


}


} else {
tmp4885 := PrimEqual(MakeNumber(0), V2218)

if True == tmp4885 {
tmp4870 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2217)


__e.TailApply(PrimFunc(symshen_4read_1loop), V2217, tmp4870, V2219, V2220)
return


} else {
tmp4883 := Call(__e, V2220, V2218)


if True == tmp4883 {
tmp4871 := MakeNative(func(__e *ControlFlow) {
W2222 := __e.Get(1)
_ = W2222
tmp4877 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), W2222)


if True == tmp4877 {
tmp4872 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2217)


tmp4873 := PrimCons(V2218, Nil)

tmp4874 := Call(__e, PrimFunc(symappend), V2219, tmp4873)


__e.TailApply(PrimFunc(symshen_4read_1loop), V2217, tmp4872, tmp4874, V2220)
return


} else {
tmp4875 := Call(__e, PrimFunc(symshen_4record_1it), V2219)


_ = tmp4875

__e.Return(W2222)
return


}


}, 1)

tmp4878 := Call(__e, PrimFunc(symshen_4try_1parse), V2219)


__e.TailApply(tmp4871, tmp4878)
return


} else {
tmp4879 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2217)


tmp4880 := PrimCons(V2218, Nil)

tmp4881 := Call(__e, PrimFunc(symappend), V2219, tmp4880)


__e.TailApply(PrimFunc(symshen_4read_1loop), V2217, tmp4879, tmp4881, V2220)
return


}


}


}


}


}, 4)

tmp4890 := Call(__e, ns2_1set, symshen_4read_1loop, tmp4866)


_ = tmp4890

tmp4891 := MakeNative(func(__e *ControlFlow) {
V2223 := __e.Get(1)
_ = V2223
tmp4892 := MakeNative(func(__e *ControlFlow) {
W2224 := __e.Get(1)
_ = W2224
tmp4894 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), W2224)


if True == tmp4894 {
__e.Return(symshen_4i_1failed_b)
return
} else {
__e.TailApply(PrimFunc(symshen_4process_1sexprs), W2224)
return
}


}, 1)

tmp4895 := MakeNative(func(__e *ControlFlow) {
tmp4896 := MakeNative(func(__e *ControlFlow) {
Z2225 := __e.Get(1)
_ = Z2225
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2225)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp4896, V2223)
return


}, 0)

tmp4897 := MakeNative(func(__e *ControlFlow) {
Z2226 := __e.Get(1)
_ = Z2226
__e.Return(symshen_4i_1failed_b)
return
}, 1)

tmp4898 := Call(__e, try_1catch, tmp4895, tmp4897)


__e.TailApply(tmp4892, tmp4898)
return


}, 1)

tmp4899 := Call(__e, ns2_1set, symshen_4try_1parse, tmp4891)


_ = tmp4899

tmp4900 := MakeNative(func(__e *ControlFlow) {
V2229 := __e.Get(1)
_ = V2229
tmp4904 := PrimEqual(symshen_4i_1failed_b, V2229)

if True == tmp4904 {
__e.Return(True)
return
} else {
tmp4902 := PrimEqual(Nil, V2229)

if True == tmp4902 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp4905 := Call(__e, ns2_1set, symshen_4nothing_1doing_2, tmp4900)


_ = tmp4905

tmp4906 := MakeNative(func(__e *ControlFlow) {
V2230 := __e.Get(1)
_ = V2230
tmp4907 := Call(__e, PrimFunc(symshen_4bytes_1_6string), V2230)


__e.Return(PrimSet(symshen_4_dit_d, tmp4907))
return


}, 1)

tmp4908 := Call(__e, ns2_1set, symshen_4record_1it, tmp4906)


_ = tmp4908

tmp4909 := MakeNative(func(__e *ControlFlow) {
V2231 := __e.Get(1)
_ = V2231
tmp4917 := PrimEqual(Nil, V2231)

if True == tmp4917 {
__e.Return(MakeString(""))
return
} else {
tmp4915 := PrimIsPair(V2231)

if True == tmp4915 {
tmp4910 := PrimHead(V2231)

tmp4911 := PrimNumberToString(tmp4910)

tmp4912 := PrimTail(V2231)

tmp4913 := Call(__e, PrimFunc(symshen_4bytes_1_6string), tmp4912)


__e.Return(PrimStringConcat(tmp4911, tmp4913))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.bytes->string")))
return
}


}


}, 1)

tmp4918 := Call(__e, ns2_1set, symshen_4bytes_1_6string, tmp4909)


_ = tmp4918

tmp4919 := MakeNative(func(__e *ControlFlow) {
V2232 := __e.Get(1)
_ = V2232
tmp4920 := MakeNative(func(__e *ControlFlow) {
W2233 := __e.Get(1)
_ = W2233
tmp4921 := MakeNative(func(__e *ControlFlow) {
W2234 := __e.Get(1)
_ = W2234
tmp4922 := MakeNative(func(__e *ControlFlow) {
W2235 := __e.Get(1)
_ = W2235
tmp4923 := MakeNative(func(__e *ControlFlow) {
Z2236 := __e.Get(1)
_ = Z2236
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2236, W2235)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4923, W2233)
return


}, 1)

tmp4924 := Call(__e, PrimFunc(symshen_4find_1types), W2233)


__e.TailApply(tmp4922, tmp4924)
return


}, 1)

tmp4925 := Call(__e, PrimFunc(symshen_4find_1arities), W2233)


__e.TailApply(tmp4921, tmp4925)
return


}, 1)

tmp4926 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), V2232)


__e.TailApply(tmp4920, tmp4926)
return


}, 1)

tmp4927 := Call(__e, ns2_1set, symshen_4process_1sexprs, tmp4919)


_ = tmp4927

tmp4928 := MakeNative(func(__e *ControlFlow) {
V2237 := __e.Get(1)
_ = V2237
tmp4950 := PrimIsPair(V2237)

var ifres4941 Obj

if True == tmp4950 {
tmp4948 := PrimTail(V2237)

tmp4949 := PrimIsPair(tmp4948)

var ifres4943 Obj

if True == tmp4949 {
tmp4945 := PrimHead(V2237)

tmp4946 := PrimIntern(MakeString(":"))

tmp4947 := PrimEqual(tmp4945, tmp4946)

var ifres4944 Obj

if True == tmp4947 {
ifres4944 = True


} else {
ifres4944 = False


}

ifres4943 = ifres4944


} else {
ifres4943 = False


}

var ifres4942 Obj

if True == ifres4943 {
ifres4942 = True


} else {
ifres4942 = False


}

ifres4941 = ifres4942


} else {
ifres4941 = False


}

if True == ifres4941 {
tmp4929 := PrimTail(V2237)

tmp4930 := PrimHead(tmp4929)

tmp4931 := PrimTail(V2237)

tmp4932 := PrimTail(tmp4931)

tmp4933 := Call(__e, PrimFunc(symshen_4find_1types), tmp4932)


__e.Return(PrimCons(tmp4930, tmp4933))
return


} else {
tmp4939 := PrimIsPair(V2237)

if True == tmp4939 {
tmp4934 := PrimHead(V2237)

tmp4935 := Call(__e, PrimFunc(symshen_4find_1types), tmp4934)


tmp4936 := PrimTail(V2237)

tmp4937 := Call(__e, PrimFunc(symshen_4find_1types), tmp4936)


__e.TailApply(PrimFunc(symappend), tmp4935, tmp4937)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp4951 := Call(__e, ns2_1set, symshen_4find_1types, tmp4928)


_ = tmp4951

tmp4952 := MakeNative(func(__e *ControlFlow) {
V2240 := __e.Get(1)
_ = V2240
tmp5001 := PrimIsPair(V2240)

var ifres4982 Obj

if True == tmp5001 {
tmp4999 := PrimHead(V2240)

tmp5000 := PrimEqual(symdefine, tmp4999)

var ifres4984 Obj

if True == tmp5000 {
tmp4997 := PrimTail(V2240)

tmp4998 := PrimIsPair(tmp4997)

var ifres4986 Obj

if True == tmp4998 {
tmp4994 := PrimTail(V2240)

tmp4995 := PrimTail(tmp4994)

tmp4996 := PrimIsPair(tmp4995)

var ifres4988 Obj

if True == tmp4996 {
tmp4990 := PrimTail(V2240)

tmp4991 := PrimTail(tmp4990)

tmp4992 := PrimHead(tmp4991)

tmp4993 := PrimEqual(sym_i, tmp4992)

var ifres4989 Obj

if True == tmp4993 {
ifres4989 = True


} else {
ifres4989 = False


}

ifres4988 = ifres4989


} else {
ifres4988 = False


}

var ifres4987 Obj

if True == ifres4988 {
ifres4987 = True


} else {
ifres4987 = False


}

ifres4986 = ifres4987


} else {
ifres4986 = False


}

var ifres4985 Obj

if True == ifres4986 {
ifres4985 = True


} else {
ifres4985 = False


}

ifres4984 = ifres4985


} else {
ifres4984 = False


}

var ifres4983 Obj

if True == ifres4984 {
ifres4983 = True


} else {
ifres4983 = False


}

ifres4982 = ifres4983


} else {
ifres4982 = False


}

if True == ifres4982 {
tmp4953 := PrimTail(V2240)

tmp4954 := PrimHead(tmp4953)

tmp4955 := PrimTail(V2240)

tmp4956 := PrimHead(tmp4955)

tmp4957 := PrimTail(V2240)

tmp4958 := PrimTail(tmp4957)

tmp4959 := PrimTail(tmp4958)

tmp4960 := Call(__e, PrimFunc(symshen_4find_1arity), tmp4956, MakeNumber(1), tmp4959)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp4954, tmp4960)
return


} else {
tmp4980 := PrimIsPair(V2240)

var ifres4972 Obj

if True == tmp4980 {
tmp4978 := PrimHead(V2240)

tmp4979 := PrimEqual(symdefine, tmp4978)

var ifres4974 Obj

if True == tmp4979 {
tmp4976 := PrimTail(V2240)

tmp4977 := PrimIsPair(tmp4976)

var ifres4975 Obj

if True == tmp4977 {
ifres4975 = True


} else {
ifres4975 = False


}

ifres4974 = ifres4975


} else {
ifres4974 = False


}

var ifres4973 Obj

if True == ifres4974 {
ifres4973 = True


} else {
ifres4973 = False


}

ifres4972 = ifres4973


} else {
ifres4972 = False


}

if True == ifres4972 {
tmp4961 := PrimTail(V2240)

tmp4962 := PrimHead(tmp4961)

tmp4963 := PrimTail(V2240)

tmp4964 := PrimHead(tmp4963)

tmp4965 := PrimTail(V2240)

tmp4966 := PrimTail(tmp4965)

tmp4967 := Call(__e, PrimFunc(symshen_4find_1arity), tmp4964, MakeNumber(0), tmp4966)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp4962, tmp4967)
return


} else {
tmp4970 := PrimIsPair(V2240)

if True == tmp4970 {
tmp4968 := MakeNative(func(__e *ControlFlow) {
Z2241 := __e.Get(1)
_ = Z2241
__e.TailApply(PrimFunc(symshen_4find_1arities), Z2241)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4968, V2240)
return


} else {
__e.Return(symshen_4skip)
return
}


}


}


}, 1)

tmp5002 := Call(__e, ns2_1set, symshen_4find_1arities, tmp4952)


_ = tmp5002

tmp5003 := MakeNative(func(__e *ControlFlow) {
V2242 := __e.Get(1)
_ = V2242
V2243 := __e.Get(2)
_ = V2243
tmp5004 := MakeNative(func(__e *ControlFlow) {
W2244 := __e.Get(1)
_ = W2244
tmp5015 := PrimEqual(W2244, MakeNumber(-1))

if True == tmp5015 {
__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2242, V2243)
return
} else {
tmp5013 := PrimEqual(W2244, V2243)

if True == tmp5013 {
__e.Return(symshen_4skip)
return
} else {
tmp5011 := Call(__e, PrimFunc(symshen_4sysfunc_2), V2242)


if True == tmp5011 {
tmp5005 := Call(__e, PrimFunc(symshen_4app), V2242, MakeString(" is a system function\n"), symshen_4a)


__e.Return(PrimSimpleError(tmp5005))
return


} else {
tmp5006 := Call(__e, PrimFunc(symshen_4app), V2242, MakeString(" may cause errors\n"), symshen_4a)


tmp5007 := PrimStringConcat(MakeString("changing the arity of "), tmp5006)

tmp5008 := Call(__e, PrimFunc(symstoutput))


tmp5009 := Call(__e, PrimFunc(sympr), tmp5007, tmp5008)


_ = tmp5009

__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2242, V2243)
return


}


}


}


}, 1)

tmp5016 := Call(__e, PrimFunc(symarity), V2242)


__e.TailApply(tmp5004, tmp5016)
return


}, 2)

tmp5017 := Call(__e, ns2_1set, symshen_4store_1arity, tmp5003)


_ = tmp5017

tmp5018 := MakeNative(func(__e *ControlFlow) {
V2245 := __e.Get(1)
_ = V2245
V2246 := __e.Get(2)
_ = V2246
tmp5023 := PrimEqual(MakeNumber(0), V2246)

if True == tmp5023 {
tmp5019 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V2245, symarity, MakeNumber(0), tmp5019)
return


} else {
tmp5020 := PrimValue(sym_dproperty_1vector_d)

tmp5021 := Call(__e, PrimFunc(symput), V2245, symarity, V2246, tmp5020)


_ = tmp5021

__e.TailApply(PrimFunc(symshen_4update_1lambdatable), V2245, V2246)
return


}


}, 2)

tmp5024 := Call(__e, ns2_1set, symshen_4execute_1store_1arity, tmp5018)


_ = tmp5024

tmp5025 := MakeNative(func(__e *ControlFlow) {
V2247 := __e.Get(1)
_ = V2247
V2248 := __e.Get(2)
_ = V2248
tmp5026 := MakeNative(func(__e *ControlFlow) {
W2249 := __e.Get(1)
_ = W2249
tmp5027 := MakeNative(func(__e *ControlFlow) {
W2250 := __e.Get(1)
_ = W2250
tmp5028 := MakeNative(func(__e *ControlFlow) {
W2251 := __e.Get(1)
_ = W2251
tmp5029 := MakeNative(func(__e *ControlFlow) {
W2252 := __e.Get(1)
_ = W2252
__e.Return(W2252)
return
}, 1)

tmp5030 := PrimSet(symshen_4_dlambdatable_d, W2251)

__e.TailApply(tmp5029, tmp5030)
return


}, 1)

tmp5031 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2247, W2250, W2249)


__e.TailApply(tmp5028, tmp5031)
return


}, 1)

tmp5032 := PrimCons(V2247, Nil)

tmp5033 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp5032, V2248)


tmp5034 := Call(__e, PrimFunc(symeval_1kl), tmp5033)


__e.TailApply(tmp5027, tmp5034)
return


}, 1)

tmp5035 := PrimValue(symshen_4_dlambdatable_d)

__e.TailApply(tmp5026, tmp5035)
return


}, 2)

tmp5036 := Call(__e, ns2_1set, symshen_4update_1lambdatable, tmp5025)


_ = tmp5036

tmp5037 := MakeNative(func(__e *ControlFlow) {
V2255 := __e.Get(1)
_ = V2255
V2256 := __e.Get(2)
_ = V2256
tmp5055 := PrimEqual(MakeNumber(0), V2256)

if True == tmp5055 {
__e.Return(symshen_4skip)
return
} else {
tmp5053 := PrimEqual(MakeNumber(1), V2256)

if True == tmp5053 {
tmp5038 := MakeNative(func(__e *ControlFlow) {
W2257 := __e.Get(1)
_ = W2257
tmp5039 := PrimCons(W2257, Nil)

tmp5040 := Call(__e, PrimFunc(symappend), V2255, tmp5039)


tmp5041 := PrimCons(tmp5040, Nil)

tmp5042 := PrimCons(W2257, tmp5041)

__e.Return(PrimCons(symlambda, tmp5042))
return


}, 1)

tmp5043 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(tmp5038, tmp5043)
return


} else {
tmp5044 := MakeNative(func(__e *ControlFlow) {
W2258 := __e.Get(1)
_ = W2258
tmp5045 := PrimCons(W2258, Nil)

tmp5046 := Call(__e, PrimFunc(symappend), V2255, tmp5045)


tmp5047 := PrimNumberSubtract(V2256, MakeNumber(1))

tmp5048 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp5046, tmp5047)


tmp5049 := PrimCons(tmp5048, Nil)

tmp5050 := PrimCons(W2258, tmp5049)

__e.Return(PrimCons(symlambda, tmp5050))
return


}, 1)

tmp5051 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(tmp5044, tmp5051)
return


}


}


}, 2)

tmp5056 := Call(__e, ns2_1set, symshen_4lambda_1function, tmp5037)


_ = tmp5056

tmp5057 := MakeNative(func(__e *ControlFlow) {
V2268 := __e.Get(1)
_ = V2268
V2269 := __e.Get(2)
_ = V2269
V2270 := __e.Get(3)
_ = V2270
tmp5080 := PrimEqual(Nil, V2270)

if True == tmp5080 {
tmp5058 := PrimCons(V2268, V2269)

__e.Return(PrimCons(tmp5058, Nil))
return


} else {
tmp5078 := PrimIsPair(V2270)

var ifres5069 Obj

if True == tmp5078 {
tmp5076 := PrimHead(V2270)

tmp5077 := PrimIsPair(tmp5076)

var ifres5071 Obj

if True == tmp5077 {
tmp5073 := PrimHead(V2270)

tmp5074 := PrimHead(tmp5073)

tmp5075 := PrimEqual(V2268, tmp5074)

var ifres5072 Obj

if True == tmp5075 {
ifres5072 = True


} else {
ifres5072 = False


}

ifres5071 = ifres5072


} else {
ifres5071 = False


}

var ifres5070 Obj

if True == ifres5071 {
ifres5070 = True


} else {
ifres5070 = False


}

ifres5069 = ifres5070


} else {
ifres5069 = False


}

if True == ifres5069 {
tmp5059 := PrimHead(V2270)

tmp5060 := PrimHead(tmp5059)

tmp5061 := PrimCons(tmp5060, V2269)

tmp5062 := PrimTail(V2270)

__e.Return(PrimCons(tmp5061, tmp5062))
return


} else {
tmp5067 := PrimIsPair(V2270)

if True == tmp5067 {
tmp5063 := PrimHead(V2270)

tmp5064 := PrimTail(V2270)

tmp5065 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2268, V2269, tmp5064)


__e.Return(PrimCons(tmp5063, tmp5065))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.assoc->")))
return
}


}


}


}, 3)

tmp5081 := Call(__e, ns2_1set, symshen_4assoc_1_6, tmp5057)


_ = tmp5081

tmp5082 := MakeNative(func(__e *ControlFlow) {
V2285 := __e.Get(1)
_ = V2285
V2286 := __e.Get(2)
_ = V2286
V2287 := __e.Get(3)
_ = V2287
tmp5129 := PrimEqual(MakeNumber(0), V2286)

var ifres5122 Obj

if True == tmp5129 {
tmp5128 := PrimIsPair(V2287)

var ifres5124 Obj

if True == tmp5128 {
tmp5126 := PrimHead(V2287)

tmp5127 := PrimEqual(tmp5126, sym_1_6)

var ifres5125 Obj

if True == tmp5127 {
ifres5125 = True


} else {
ifres5125 = False


}

ifres5124 = ifres5125


} else {
ifres5124 = False


}

var ifres5123 Obj

if True == ifres5124 {
ifres5123 = True


} else {
ifres5123 = False


}

ifres5122 = ifres5123


} else {
ifres5122 = False


}

if True == ifres5122 {
__e.Return(MakeNumber(0))
return
} else {
tmp5120 := PrimEqual(MakeNumber(0), V2286)

var ifres5113 Obj

if True == tmp5120 {
tmp5119 := PrimIsPair(V2287)

var ifres5115 Obj

if True == tmp5119 {
tmp5117 := PrimHead(V2287)

tmp5118 := PrimEqual(tmp5117, sym_5_1)

var ifres5116 Obj

if True == tmp5118 {
ifres5116 = True


} else {
ifres5116 = False


}

ifres5115 = ifres5116


} else {
ifres5115 = False


}

var ifres5114 Obj

if True == ifres5115 {
ifres5114 = True


} else {
ifres5114 = False


}

ifres5113 = ifres5114


} else {
ifres5113 = False


}

if True == ifres5113 {
__e.Return(MakeNumber(0))
return
} else {
tmp5111 := PrimEqual(MakeNumber(0), V2286)

var ifres5108 Obj

if True == tmp5111 {
tmp5110 := PrimIsPair(V2287)

var ifres5109 Obj

if True == tmp5110 {
ifres5109 = True


} else {
ifres5109 = False


}

ifres5108 = ifres5109


} else {
ifres5108 = False


}

if True == ifres5108 {
tmp5083 := PrimTail(V2287)

tmp5084 := Call(__e, PrimFunc(symshen_4find_1arity), V2285, MakeNumber(0), tmp5083)


__e.Return(PrimNumberAdd(MakeNumber(1), tmp5084))
return


} else {
tmp5106 := PrimEqual(MakeNumber(1), V2286)

var ifres5099 Obj

if True == tmp5106 {
tmp5105 := PrimIsPair(V2287)

var ifres5101 Obj

if True == tmp5105 {
tmp5103 := PrimHead(V2287)

tmp5104 := PrimEqual(sym_j, tmp5103)

var ifres5102 Obj

if True == tmp5104 {
ifres5102 = True


} else {
ifres5102 = False


}

ifres5101 = ifres5102


} else {
ifres5101 = False


}

var ifres5100 Obj

if True == ifres5101 {
ifres5100 = True


} else {
ifres5100 = False


}

ifres5099 = ifres5100


} else {
ifres5099 = False


}

if True == ifres5099 {
tmp5085 := PrimTail(V2287)

__e.TailApply(PrimFunc(symshen_4find_1arity), V2285, MakeNumber(0), tmp5085)
return


} else {
tmp5097 := PrimEqual(MakeNumber(1), V2286)

var ifres5094 Obj

if True == tmp5097 {
tmp5096 := PrimIsPair(V2287)

var ifres5095 Obj

if True == tmp5096 {
ifres5095 = True


} else {
ifres5095 = False


}

ifres5094 = ifres5095


} else {
ifres5094 = False


}

if True == ifres5094 {
tmp5086 := PrimTail(V2287)

__e.TailApply(PrimFunc(symshen_4find_1arity), V2285, MakeNumber(1), tmp5086)
return


} else {
tmp5092 := PrimEqual(MakeNumber(1), V2286)

if True == tmp5092 {
tmp5087 := Call(__e, PrimFunc(symshen_4app), V2285, MakeString(" definition: missing }\n"), symshen_4a)


tmp5088 := PrimStringConcat(MakeString("syntax error in "), tmp5087)

__e.Return(PrimSimpleError(tmp5088))
return


} else {
tmp5089 := Call(__e, PrimFunc(symshen_4app), V2285, MakeString(" definition: missing -> or <-\n"), symshen_4a)


tmp5090 := PrimStringConcat(MakeString("syntax error in "), tmp5089)

__e.Return(PrimSimpleError(tmp5090))
return


}


}


}


}


}


}


}, 3)

tmp5130 := Call(__e, ns2_1set, symshen_4find_1arity, tmp5082)


_ = tmp5130

tmp5131 := MakeNative(func(__e *ControlFlow) {
V2288 := __e.Get(1)
_ = V2288
tmp5132 := MakeNative(func(__e *ControlFlow) {
W2289 := __e.Get(1)
_ = W2289
tmp5377 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2289)


if True == tmp5377 {
tmp5133 := MakeNative(func(__e *ControlFlow) {
W2300 := __e.Get(1)
_ = W2300
tmp5345 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2300)


if True == tmp5345 {
tmp5134 := MakeNative(func(__e *ControlFlow) {
W2311 := __e.Get(1)
_ = W2311
tmp5327 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2311)


if True == tmp5327 {
tmp5135 := MakeNative(func(__e *ControlFlow) {
W2317 := __e.Get(1)
_ = W2317
tmp5309 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2317)


if True == tmp5309 {
tmp5136 := MakeNative(func(__e *ControlFlow) {
W2323 := __e.Get(1)
_ = W2323
tmp5291 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2323)


if True == tmp5291 {
tmp5137 := MakeNative(func(__e *ControlFlow) {
W2329 := __e.Get(1)
_ = W2329
tmp5272 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2329)


if True == tmp5272 {
tmp5138 := MakeNative(func(__e *ControlFlow) {
W2335 := __e.Get(1)
_ = W2335
tmp5247 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2335)


if True == tmp5247 {
tmp5139 := MakeNative(func(__e *ControlFlow) {
W2343 := __e.Get(1)
_ = W2343
tmp5228 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2343)


if True == tmp5228 {
tmp5140 := MakeNative(func(__e *ControlFlow) {
W2349 := __e.Get(1)
_ = W2349
tmp5209 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2349)


if True == tmp5209 {
tmp5141 := MakeNative(func(__e *ControlFlow) {
W2355 := __e.Get(1)
_ = W2355
tmp5192 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2355)


if True == tmp5192 {
tmp5142 := MakeNative(func(__e *ControlFlow) {
W2361 := __e.Get(1)
_ = W2361
tmp5172 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2361)


if True == tmp5172 {
tmp5143 := MakeNative(func(__e *ControlFlow) {
W2368 := __e.Get(1)
_ = W2368
tmp5155 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2368)


if True == tmp5155 {
tmp5144 := MakeNative(func(__e *ControlFlow) {
W2374 := __e.Get(1)
_ = W2374
tmp5146 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2374)


if True == tmp5146 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2374)
return
}


}, 1)

tmp5147 := MakeNative(func(__e *ControlFlow) {
W2375 := __e.Get(1)
_ = W2375
tmp5151 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2375)


if True == tmp5151 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5148 := MakeNative(func(__e *ControlFlow) {
W2376 := __e.Get(1)
_ = W2376
__e.TailApply(PrimFunc(symshen_4comb), W2376, Nil)
return
}, 1)

tmp5149 := Call(__e, PrimFunc(symshen_4in_1_6), W2375)


__e.TailApply(tmp5148, tmp5149)
return


}


}, 1)

tmp5152 := Call(__e, PrimFunc(sym_5e_6), V2288)


tmp5153 := Call(__e, tmp5147, tmp5152)


__e.TailApply(tmp5144, tmp5153)
return


} else {
__e.Return(W2368)
return
}


}, 1)

tmp5156 := MakeNative(func(__e *ControlFlow) {
W2369 := __e.Get(1)
_ = W2369
tmp5168 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2369)


if True == tmp5168 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5157 := MakeNative(func(__e *ControlFlow) {
W2370 := __e.Get(1)
_ = W2370
tmp5158 := MakeNative(func(__e *ControlFlow) {
W2371 := __e.Get(1)
_ = W2371
tmp5164 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2371)


if True == tmp5164 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5159 := MakeNative(func(__e *ControlFlow) {
W2372 := __e.Get(1)
_ = W2372
tmp5160 := MakeNative(func(__e *ControlFlow) {
W2373 := __e.Get(1)
_ = W2373
__e.TailApply(PrimFunc(symshen_4comb), W2373, W2372)
return
}, 1)

tmp5161 := Call(__e, PrimFunc(symshen_4in_1_6), W2371)


__e.TailApply(tmp5160, tmp5161)
return


}, 1)

tmp5162 := Call(__e, PrimFunc(symshen_4_5_1out), W2371)


__e.TailApply(tmp5159, tmp5162)
return


}


}, 1)

tmp5165 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2370)


__e.TailApply(tmp5158, tmp5165)
return


}, 1)

tmp5166 := Call(__e, PrimFunc(symshen_4in_1_6), W2369)


__e.TailApply(tmp5157, tmp5166)
return


}


}, 1)

tmp5169 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), V2288)


tmp5170 := Call(__e, tmp5156, tmp5169)


__e.TailApply(tmp5143, tmp5170)
return


} else {
__e.Return(W2361)
return
}


}, 1)

tmp5173 := MakeNative(func(__e *ControlFlow) {
W2362 := __e.Get(1)
_ = W2362
tmp5188 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2362)


if True == tmp5188 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5174 := MakeNative(func(__e *ControlFlow) {
W2363 := __e.Get(1)
_ = W2363
tmp5175 := MakeNative(func(__e *ControlFlow) {
W2364 := __e.Get(1)
_ = W2364
tmp5176 := MakeNative(func(__e *ControlFlow) {
W2365 := __e.Get(1)
_ = W2365
tmp5183 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2365)


if True == tmp5183 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5177 := MakeNative(func(__e *ControlFlow) {
W2366 := __e.Get(1)
_ = W2366
tmp5178 := MakeNative(func(__e *ControlFlow) {
W2367 := __e.Get(1)
_ = W2367
tmp5179 := PrimCons(W2363, W2366)

__e.TailApply(PrimFunc(symshen_4comb), W2367, tmp5179)
return


}, 1)

tmp5180 := Call(__e, PrimFunc(symshen_4in_1_6), W2365)


__e.TailApply(tmp5178, tmp5180)
return


}, 1)

tmp5181 := Call(__e, PrimFunc(symshen_4_5_1out), W2365)


__e.TailApply(tmp5177, tmp5181)
return


}


}, 1)

tmp5184 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2364)


__e.TailApply(tmp5176, tmp5184)
return


}, 1)

tmp5185 := Call(__e, PrimFunc(symshen_4in_1_6), W2362)


__e.TailApply(tmp5175, tmp5185)
return


}, 1)

tmp5186 := Call(__e, PrimFunc(symshen_4_5_1out), W2362)


__e.TailApply(tmp5174, tmp5186)
return


}


}, 1)

tmp5189 := Call(__e, PrimFunc(symshen_4_5atom_6), V2288)


tmp5190 := Call(__e, tmp5173, tmp5189)


__e.TailApply(tmp5142, tmp5190)
return


} else {
__e.Return(W2355)
return
}


}, 1)

tmp5193 := MakeNative(func(__e *ControlFlow) {
W2356 := __e.Get(1)
_ = W2356
tmp5205 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2356)


if True == tmp5205 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5194 := MakeNative(func(__e *ControlFlow) {
W2357 := __e.Get(1)
_ = W2357
tmp5195 := MakeNative(func(__e *ControlFlow) {
W2358 := __e.Get(1)
_ = W2358
tmp5201 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2358)


if True == tmp5201 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5196 := MakeNative(func(__e *ControlFlow) {
W2359 := __e.Get(1)
_ = W2359
tmp5197 := MakeNative(func(__e *ControlFlow) {
W2360 := __e.Get(1)
_ = W2360
__e.TailApply(PrimFunc(symshen_4comb), W2360, W2359)
return
}, 1)

tmp5198 := Call(__e, PrimFunc(symshen_4in_1_6), W2358)


__e.TailApply(tmp5197, tmp5198)
return


}, 1)

tmp5199 := Call(__e, PrimFunc(symshen_4_5_1out), W2358)


__e.TailApply(tmp5196, tmp5199)
return


}


}, 1)

tmp5202 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2357)


__e.TailApply(tmp5195, tmp5202)
return


}, 1)

tmp5203 := Call(__e, PrimFunc(symshen_4in_1_6), W2356)


__e.TailApply(tmp5194, tmp5203)
return


}


}, 1)

tmp5206 := Call(__e, PrimFunc(symshen_4_5comment_6), V2288)


tmp5207 := Call(__e, tmp5193, tmp5206)


__e.TailApply(tmp5141, tmp5207)
return


} else {
__e.Return(W2349)
return
}


}, 1)

tmp5210 := MakeNative(func(__e *ControlFlow) {
W2350 := __e.Get(1)
_ = W2350
tmp5224 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2350)


if True == tmp5224 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5211 := MakeNative(func(__e *ControlFlow) {
W2351 := __e.Get(1)
_ = W2351
tmp5212 := MakeNative(func(__e *ControlFlow) {
W2352 := __e.Get(1)
_ = W2352
tmp5220 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2352)


if True == tmp5220 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5213 := MakeNative(func(__e *ControlFlow) {
W2353 := __e.Get(1)
_ = W2353
tmp5214 := MakeNative(func(__e *ControlFlow) {
W2354 := __e.Get(1)
_ = W2354
tmp5215 := PrimIntern(MakeString(","))

tmp5216 := PrimCons(tmp5215, W2353)

__e.TailApply(PrimFunc(symshen_4comb), W2354, tmp5216)
return


}, 1)

tmp5217 := Call(__e, PrimFunc(symshen_4in_1_6), W2352)


__e.TailApply(tmp5214, tmp5217)
return


}, 1)

tmp5218 := Call(__e, PrimFunc(symshen_4_5_1out), W2352)


__e.TailApply(tmp5213, tmp5218)
return


}


}, 1)

tmp5221 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2351)


__e.TailApply(tmp5212, tmp5221)
return


}, 1)

tmp5222 := Call(__e, PrimFunc(symshen_4in_1_6), W2350)


__e.TailApply(tmp5211, tmp5222)
return


}


}, 1)

tmp5225 := Call(__e, PrimFunc(symshen_4_5comma_6), V2288)


tmp5226 := Call(__e, tmp5210, tmp5225)


__e.TailApply(tmp5140, tmp5226)
return


} else {
__e.Return(W2343)
return
}


}, 1)

tmp5229 := MakeNative(func(__e *ControlFlow) {
W2344 := __e.Get(1)
_ = W2344
tmp5243 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2344)


if True == tmp5243 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5230 := MakeNative(func(__e *ControlFlow) {
W2345 := __e.Get(1)
_ = W2345
tmp5231 := MakeNative(func(__e *ControlFlow) {
W2346 := __e.Get(1)
_ = W2346
tmp5239 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2346)


if True == tmp5239 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5232 := MakeNative(func(__e *ControlFlow) {
W2347 := __e.Get(1)
_ = W2347
tmp5233 := MakeNative(func(__e *ControlFlow) {
W2348 := __e.Get(1)
_ = W2348
tmp5234 := PrimIntern(MakeString(":"))

tmp5235 := PrimCons(tmp5234, W2347)

__e.TailApply(PrimFunc(symshen_4comb), W2348, tmp5235)
return


}, 1)

tmp5236 := Call(__e, PrimFunc(symshen_4in_1_6), W2346)


__e.TailApply(tmp5233, tmp5236)
return


}, 1)

tmp5237 := Call(__e, PrimFunc(symshen_4_5_1out), W2346)


__e.TailApply(tmp5232, tmp5237)
return


}


}, 1)

tmp5240 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2345)


__e.TailApply(tmp5231, tmp5240)
return


}, 1)

tmp5241 := Call(__e, PrimFunc(symshen_4in_1_6), W2344)


__e.TailApply(tmp5230, tmp5241)
return


}


}, 1)

tmp5244 := Call(__e, PrimFunc(symshen_4_5colon_6), V2288)


tmp5245 := Call(__e, tmp5229, tmp5244)


__e.TailApply(tmp5139, tmp5245)
return


} else {
__e.Return(W2335)
return
}


}, 1)

tmp5248 := MakeNative(func(__e *ControlFlow) {
W2336 := __e.Get(1)
_ = W2336
tmp5268 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2336)


if True == tmp5268 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5249 := MakeNative(func(__e *ControlFlow) {
W2337 := __e.Get(1)
_ = W2337
tmp5250 := MakeNative(func(__e *ControlFlow) {
W2338 := __e.Get(1)
_ = W2338
tmp5264 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2338)


if True == tmp5264 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5251 := MakeNative(func(__e *ControlFlow) {
W2339 := __e.Get(1)
_ = W2339
tmp5252 := MakeNative(func(__e *ControlFlow) {
W2340 := __e.Get(1)
_ = W2340
tmp5260 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2340)


if True == tmp5260 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5253 := MakeNative(func(__e *ControlFlow) {
W2341 := __e.Get(1)
_ = W2341
tmp5254 := MakeNative(func(__e *ControlFlow) {
W2342 := __e.Get(1)
_ = W2342
tmp5255 := PrimIntern(MakeString(":="))

tmp5256 := PrimCons(tmp5255, W2341)

__e.TailApply(PrimFunc(symshen_4comb), W2342, tmp5256)
return


}, 1)

tmp5257 := Call(__e, PrimFunc(symshen_4in_1_6), W2340)


__e.TailApply(tmp5254, tmp5257)
return


}, 1)

tmp5258 := Call(__e, PrimFunc(symshen_4_5_1out), W2340)


__e.TailApply(tmp5253, tmp5258)
return


}


}, 1)

tmp5261 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2339)


__e.TailApply(tmp5252, tmp5261)
return


}, 1)

tmp5262 := Call(__e, PrimFunc(symshen_4in_1_6), W2338)


__e.TailApply(tmp5251, tmp5262)
return


}


}, 1)

tmp5265 := Call(__e, PrimFunc(symshen_4_5equal_6), W2337)


__e.TailApply(tmp5250, tmp5265)
return


}, 1)

tmp5266 := Call(__e, PrimFunc(symshen_4in_1_6), W2336)


__e.TailApply(tmp5249, tmp5266)
return


}


}, 1)

tmp5269 := Call(__e, PrimFunc(symshen_4_5colon_6), V2288)


tmp5270 := Call(__e, tmp5248, tmp5269)


__e.TailApply(tmp5138, tmp5270)
return


} else {
__e.Return(W2329)
return
}


}, 1)

tmp5273 := MakeNative(func(__e *ControlFlow) {
W2330 := __e.Get(1)
_ = W2330
tmp5287 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2330)


if True == tmp5287 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5274 := MakeNative(func(__e *ControlFlow) {
W2331 := __e.Get(1)
_ = W2331
tmp5275 := MakeNative(func(__e *ControlFlow) {
W2332 := __e.Get(1)
_ = W2332
tmp5283 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2332)


if True == tmp5283 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5276 := MakeNative(func(__e *ControlFlow) {
W2333 := __e.Get(1)
_ = W2333
tmp5277 := MakeNative(func(__e *ControlFlow) {
W2334 := __e.Get(1)
_ = W2334
tmp5278 := PrimIntern(MakeString(";"))

tmp5279 := PrimCons(tmp5278, W2333)

__e.TailApply(PrimFunc(symshen_4comb), W2334, tmp5279)
return


}, 1)

tmp5280 := Call(__e, PrimFunc(symshen_4in_1_6), W2332)


__e.TailApply(tmp5277, tmp5280)
return


}, 1)

tmp5281 := Call(__e, PrimFunc(symshen_4_5_1out), W2332)


__e.TailApply(tmp5276, tmp5281)
return


}


}, 1)

tmp5284 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2331)


__e.TailApply(tmp5275, tmp5284)
return


}, 1)

tmp5285 := Call(__e, PrimFunc(symshen_4in_1_6), W2330)


__e.TailApply(tmp5274, tmp5285)
return


}


}, 1)

tmp5288 := Call(__e, PrimFunc(symshen_4_5semicolon_6), V2288)


tmp5289 := Call(__e, tmp5273, tmp5288)


__e.TailApply(tmp5137, tmp5289)
return


} else {
__e.Return(W2323)
return
}


}, 1)

tmp5292 := MakeNative(func(__e *ControlFlow) {
W2324 := __e.Get(1)
_ = W2324
tmp5305 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2324)


if True == tmp5305 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5293 := MakeNative(func(__e *ControlFlow) {
W2325 := __e.Get(1)
_ = W2325
tmp5294 := MakeNative(func(__e *ControlFlow) {
W2326 := __e.Get(1)
_ = W2326
tmp5301 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2326)


if True == tmp5301 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5295 := MakeNative(func(__e *ControlFlow) {
W2327 := __e.Get(1)
_ = W2327
tmp5296 := MakeNative(func(__e *ControlFlow) {
W2328 := __e.Get(1)
_ = W2328
tmp5297 := PrimCons(symbar_b, W2327)

__e.TailApply(PrimFunc(symshen_4comb), W2328, tmp5297)
return


}, 1)

tmp5298 := Call(__e, PrimFunc(symshen_4in_1_6), W2326)


__e.TailApply(tmp5296, tmp5298)
return


}, 1)

tmp5299 := Call(__e, PrimFunc(symshen_4_5_1out), W2326)


__e.TailApply(tmp5295, tmp5299)
return


}


}, 1)

tmp5302 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2325)


__e.TailApply(tmp5294, tmp5302)
return


}, 1)

tmp5303 := Call(__e, PrimFunc(symshen_4in_1_6), W2324)


__e.TailApply(tmp5293, tmp5303)
return


}


}, 1)

tmp5306 := Call(__e, PrimFunc(symshen_4_5bar_6), V2288)


tmp5307 := Call(__e, tmp5292, tmp5306)


__e.TailApply(tmp5136, tmp5307)
return


} else {
__e.Return(W2317)
return
}


}, 1)

tmp5310 := MakeNative(func(__e *ControlFlow) {
W2318 := __e.Get(1)
_ = W2318
tmp5323 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2318)


if True == tmp5323 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5311 := MakeNative(func(__e *ControlFlow) {
W2319 := __e.Get(1)
_ = W2319
tmp5312 := MakeNative(func(__e *ControlFlow) {
W2320 := __e.Get(1)
_ = W2320
tmp5319 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2320)


if True == tmp5319 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5313 := MakeNative(func(__e *ControlFlow) {
W2321 := __e.Get(1)
_ = W2321
tmp5314 := MakeNative(func(__e *ControlFlow) {
W2322 := __e.Get(1)
_ = W2322
tmp5315 := PrimCons(sym_j, W2321)

__e.TailApply(PrimFunc(symshen_4comb), W2322, tmp5315)
return


}, 1)

tmp5316 := Call(__e, PrimFunc(symshen_4in_1_6), W2320)


__e.TailApply(tmp5314, tmp5316)
return


}, 1)

tmp5317 := Call(__e, PrimFunc(symshen_4_5_1out), W2320)


__e.TailApply(tmp5313, tmp5317)
return


}


}, 1)

tmp5320 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2319)


__e.TailApply(tmp5312, tmp5320)
return


}, 1)

tmp5321 := Call(__e, PrimFunc(symshen_4in_1_6), W2318)


__e.TailApply(tmp5311, tmp5321)
return


}


}, 1)

tmp5324 := Call(__e, PrimFunc(symshen_4_5rcurly_6), V2288)


tmp5325 := Call(__e, tmp5310, tmp5324)


__e.TailApply(tmp5135, tmp5325)
return


} else {
__e.Return(W2311)
return
}


}, 1)

tmp5328 := MakeNative(func(__e *ControlFlow) {
W2312 := __e.Get(1)
_ = W2312
tmp5341 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2312)


if True == tmp5341 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5329 := MakeNative(func(__e *ControlFlow) {
W2313 := __e.Get(1)
_ = W2313
tmp5330 := MakeNative(func(__e *ControlFlow) {
W2314 := __e.Get(1)
_ = W2314
tmp5337 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2314)


if True == tmp5337 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5331 := MakeNative(func(__e *ControlFlow) {
W2315 := __e.Get(1)
_ = W2315
tmp5332 := MakeNative(func(__e *ControlFlow) {
W2316 := __e.Get(1)
_ = W2316
tmp5333 := PrimCons(sym_i, W2315)

__e.TailApply(PrimFunc(symshen_4comb), W2316, tmp5333)
return


}, 1)

tmp5334 := Call(__e, PrimFunc(symshen_4in_1_6), W2314)


__e.TailApply(tmp5332, tmp5334)
return


}, 1)

tmp5335 := Call(__e, PrimFunc(symshen_4_5_1out), W2314)


__e.TailApply(tmp5331, tmp5335)
return


}


}, 1)

tmp5338 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2313)


__e.TailApply(tmp5330, tmp5338)
return


}, 1)

tmp5339 := Call(__e, PrimFunc(symshen_4in_1_6), W2312)


__e.TailApply(tmp5329, tmp5339)
return


}


}, 1)

tmp5342 := Call(__e, PrimFunc(symshen_4_5lcurly_6), V2288)


tmp5343 := Call(__e, tmp5328, tmp5342)


__e.TailApply(tmp5134, tmp5343)
return


} else {
__e.Return(W2300)
return
}


}, 1)

tmp5346 := MakeNative(func(__e *ControlFlow) {
W2301 := __e.Get(1)
_ = W2301
tmp5373 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2301)


if True == tmp5373 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5347 := MakeNative(func(__e *ControlFlow) {
W2302 := __e.Get(1)
_ = W2302
tmp5348 := MakeNative(func(__e *ControlFlow) {
W2303 := __e.Get(1)
_ = W2303
tmp5369 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2303)


if True == tmp5369 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5349 := MakeNative(func(__e *ControlFlow) {
W2304 := __e.Get(1)
_ = W2304
tmp5350 := MakeNative(func(__e *ControlFlow) {
W2305 := __e.Get(1)
_ = W2305
tmp5351 := MakeNative(func(__e *ControlFlow) {
W2306 := __e.Get(1)
_ = W2306
tmp5364 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2306)


if True == tmp5364 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5352 := MakeNative(func(__e *ControlFlow) {
W2307 := __e.Get(1)
_ = W2307
tmp5353 := MakeNative(func(__e *ControlFlow) {
W2308 := __e.Get(1)
_ = W2308
tmp5360 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2308)


if True == tmp5360 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5354 := MakeNative(func(__e *ControlFlow) {
W2309 := __e.Get(1)
_ = W2309
tmp5355 := MakeNative(func(__e *ControlFlow) {
W2310 := __e.Get(1)
_ = W2310
tmp5356 := Call(__e, PrimFunc(symshen_4add_1sexpr), W2304, W2309)


__e.TailApply(PrimFunc(symshen_4comb), W2310, tmp5356)
return


}, 1)

tmp5357 := Call(__e, PrimFunc(symshen_4in_1_6), W2308)


__e.TailApply(tmp5355, tmp5357)
return


}, 1)

tmp5358 := Call(__e, PrimFunc(symshen_4_5_1out), W2308)


__e.TailApply(tmp5354, tmp5358)
return


}


}, 1)

tmp5361 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), W2307)


__e.TailApply(tmp5353, tmp5361)
return


}, 1)

tmp5362 := Call(__e, PrimFunc(symshen_4in_1_6), W2306)


__e.TailApply(tmp5352, tmp5362)
return


}


}, 1)

tmp5365 := Call(__e, PrimFunc(symshen_4_5rrb_6), W2305)


__e.TailApply(tmp5351, tmp5365)
return


}, 1)

tmp5366 := Call(__e, PrimFunc(symshen_4in_1_6), W2303)


__e.TailApply(tmp5350, tmp5366)
return


}, 1)

tmp5367 := Call(__e, PrimFunc(symshen_4_5_1out), W2303)


__e.TailApply(tmp5349, tmp5367)
return


}


}, 1)

tmp5370 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), W2302)


__e.TailApply(tmp5348, tmp5370)
return


}, 1)

tmp5371 := Call(__e, PrimFunc(symshen_4in_1_6), W2301)


__e.TailApply(tmp5347, tmp5371)
return


}


}, 1)

tmp5374 := Call(__e, PrimFunc(symshen_4_5lrb_6), V2288)


tmp5375 := Call(__e, tmp5346, tmp5374)


__e.TailApply(tmp5133, tmp5375)
return


} else {
__e.Return(W2289)
return
}


}, 1)

tmp5378 := MakeNative(func(__e *ControlFlow) {
W2290 := __e.Get(1)
_ = W2290
tmp5406 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2290)


if True == tmp5406 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5379 := MakeNative(func(__e *ControlFlow) {
W2291 := __e.Get(1)
_ = W2291
tmp5380 := MakeNative(func(__e *ControlFlow) {
W2292 := __e.Get(1)
_ = W2292
tmp5402 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2292)


if True == tmp5402 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5381 := MakeNative(func(__e *ControlFlow) {
W2293 := __e.Get(1)
_ = W2293
tmp5382 := MakeNative(func(__e *ControlFlow) {
W2294 := __e.Get(1)
_ = W2294
tmp5383 := MakeNative(func(__e *ControlFlow) {
W2295 := __e.Get(1)
_ = W2295
tmp5397 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2295)


if True == tmp5397 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5384 := MakeNative(func(__e *ControlFlow) {
W2296 := __e.Get(1)
_ = W2296
tmp5385 := MakeNative(func(__e *ControlFlow) {
W2297 := __e.Get(1)
_ = W2297
tmp5393 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2297)


if True == tmp5393 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5386 := MakeNative(func(__e *ControlFlow) {
W2298 := __e.Get(1)
_ = W2298
tmp5387 := MakeNative(func(__e *ControlFlow) {
W2299 := __e.Get(1)
_ = W2299
tmp5388 := Call(__e, PrimFunc(symshen_4cons_1form), W2293)


tmp5389 := PrimCons(tmp5388, W2298)

__e.TailApply(PrimFunc(symshen_4comb), W2299, tmp5389)
return


}, 1)

tmp5390 := Call(__e, PrimFunc(symshen_4in_1_6), W2297)


__e.TailApply(tmp5387, tmp5390)
return


}, 1)

tmp5391 := Call(__e, PrimFunc(symshen_4_5_1out), W2297)


__e.TailApply(tmp5386, tmp5391)
return


}


}, 1)

tmp5394 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), W2296)


__e.TailApply(tmp5385, tmp5394)
return


}, 1)

tmp5395 := Call(__e, PrimFunc(symshen_4in_1_6), W2295)


__e.TailApply(tmp5384, tmp5395)
return


}


}, 1)

tmp5398 := Call(__e, PrimFunc(symshen_4_5rsb_6), W2294)


__e.TailApply(tmp5383, tmp5398)
return


}, 1)

tmp5399 := Call(__e, PrimFunc(symshen_4in_1_6), W2292)


__e.TailApply(tmp5382, tmp5399)
return


}, 1)

tmp5400 := Call(__e, PrimFunc(symshen_4_5_1out), W2292)


__e.TailApply(tmp5381, tmp5400)
return


}


}, 1)

tmp5403 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), W2291)


__e.TailApply(tmp5380, tmp5403)
return


}, 1)

tmp5404 := Call(__e, PrimFunc(symshen_4in_1_6), W2290)


__e.TailApply(tmp5379, tmp5404)
return


}


}, 1)

tmp5407 := Call(__e, PrimFunc(symshen_4_5lsb_6), V2288)


tmp5408 := Call(__e, tmp5378, tmp5407)


__e.TailApply(tmp5132, tmp5408)
return


}, 1)

tmp5409 := Call(__e, ns2_1set, symshen_4_5s_1exprs_6, tmp5131)


_ = tmp5409

tmp5410 := MakeNative(func(__e *ControlFlow) {
V2377 := __e.Get(1)
_ = V2377
V2378 := __e.Get(2)
_ = V2378
tmp5428 := PrimIsPair(V2377)

var ifres5415 Obj

if True == tmp5428 {
tmp5426 := PrimHead(V2377)

tmp5427 := PrimEqual(sym_3, tmp5426)

var ifres5417 Obj

if True == tmp5427 {
tmp5424 := PrimTail(V2377)

tmp5425 := PrimIsPair(tmp5424)

var ifres5419 Obj

if True == tmp5425 {
tmp5421 := PrimTail(V2377)

tmp5422 := PrimTail(tmp5421)

tmp5423 := PrimEqual(Nil, tmp5422)

var ifres5420 Obj

if True == tmp5423 {
ifres5420 = True


} else {
ifres5420 = False


}

ifres5419 = ifres5420


} else {
ifres5419 = False


}

var ifres5418 Obj

if True == ifres5419 {
ifres5418 = True


} else {
ifres5418 = False


}

ifres5417 = ifres5418


} else {
ifres5417 = False


}

var ifres5416 Obj

if True == ifres5417 {
ifres5416 = True


} else {
ifres5416 = False


}

ifres5415 = ifres5416


} else {
ifres5415 = False


}

if True == ifres5415 {
tmp5411 := PrimTail(V2377)

tmp5412 := PrimHead(tmp5411)

tmp5413 := Call(__e, PrimFunc(symexplode), tmp5412)


__e.TailApply(PrimFunc(symappend), tmp5413, V2378)
return


} else {
__e.Return(PrimCons(V2377, V2378))
return
}


}, 2)

tmp5429 := Call(__e, ns2_1set, symshen_4add_1sexpr, tmp5410)


_ = tmp5429

tmp5430 := MakeNative(func(__e *ControlFlow) {
V2379 := __e.Get(1)
_ = V2379
tmp5431 := MakeNative(func(__e *ControlFlow) {
W2380 := __e.Get(1)
_ = W2380
tmp5433 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2380)


if True == tmp5433 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2380)
return
}


}, 1)

tmp5439 := Call(__e, PrimFunc(symshen_4hds_a_2), V2379, MakeNumber(91))


var ifres5434 Obj

if True == tmp5439 {
tmp5435 := MakeNative(func(__e *ControlFlow) {
W2381 := __e.Get(1)
_ = W2381
__e.TailApply(PrimFunc(symshen_4comb), W2381, symshen_4skip)
return
}, 1)

tmp5436 := Call(__e, PrimFunc(symtail), V2379)


tmp5437 := Call(__e, tmp5435, tmp5436)


ifres5434 = tmp5437


} else {
tmp5438 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5434 = tmp5438


}

__e.TailApply(tmp5431, ifres5434)
return


}, 1)

tmp5440 := Call(__e, ns2_1set, symshen_4_5lsb_6, tmp5430)


_ = tmp5440

tmp5441 := MakeNative(func(__e *ControlFlow) {
V2382 := __e.Get(1)
_ = V2382
tmp5442 := MakeNative(func(__e *ControlFlow) {
W2383 := __e.Get(1)
_ = W2383
tmp5444 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2383)


if True == tmp5444 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2383)
return
}


}, 1)

tmp5450 := Call(__e, PrimFunc(symshen_4hds_a_2), V2382, MakeNumber(93))


var ifres5445 Obj

if True == tmp5450 {
tmp5446 := MakeNative(func(__e *ControlFlow) {
W2384 := __e.Get(1)
_ = W2384
__e.TailApply(PrimFunc(symshen_4comb), W2384, symshen_4skip)
return
}, 1)

tmp5447 := Call(__e, PrimFunc(symtail), V2382)


tmp5448 := Call(__e, tmp5446, tmp5447)


ifres5445 = tmp5448


} else {
tmp5449 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5445 = tmp5449


}

__e.TailApply(tmp5442, ifres5445)
return


}, 1)

tmp5451 := Call(__e, ns2_1set, symshen_4_5rsb_6, tmp5441)


_ = tmp5451

tmp5452 := MakeNative(func(__e *ControlFlow) {
V2385 := __e.Get(1)
_ = V2385
tmp5453 := MakeNative(func(__e *ControlFlow) {
W2386 := __e.Get(1)
_ = W2386
tmp5455 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2386)


if True == tmp5455 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2386)
return
}


}, 1)

tmp5456 := MakeNative(func(__e *ControlFlow) {
W2387 := __e.Get(1)
_ = W2387
tmp5462 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2387)


if True == tmp5462 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5457 := MakeNative(func(__e *ControlFlow) {
W2388 := __e.Get(1)
_ = W2388
tmp5458 := MakeNative(func(__e *ControlFlow) {
W2389 := __e.Get(1)
_ = W2389
__e.TailApply(PrimFunc(symshen_4comb), W2389, W2388)
return
}, 1)

tmp5459 := Call(__e, PrimFunc(symshen_4in_1_6), W2387)


__e.TailApply(tmp5458, tmp5459)
return


}, 1)

tmp5460 := Call(__e, PrimFunc(symshen_4_5_1out), W2387)


__e.TailApply(tmp5457, tmp5460)
return


}


}, 1)

tmp5463 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2385)


tmp5464 := Call(__e, tmp5456, tmp5463)


__e.TailApply(tmp5453, tmp5464)
return


}, 1)

tmp5465 := Call(__e, ns2_1set, symshen_4_5s_1exprs1_6, tmp5452)


_ = tmp5465

tmp5466 := MakeNative(func(__e *ControlFlow) {
V2390 := __e.Get(1)
_ = V2390
tmp5467 := MakeNative(func(__e *ControlFlow) {
W2391 := __e.Get(1)
_ = W2391
tmp5469 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2391)


if True == tmp5469 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2391)
return
}


}, 1)

tmp5470 := MakeNative(func(__e *ControlFlow) {
W2392 := __e.Get(1)
_ = W2392
tmp5476 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2392)


if True == tmp5476 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5471 := MakeNative(func(__e *ControlFlow) {
W2393 := __e.Get(1)
_ = W2393
tmp5472 := MakeNative(func(__e *ControlFlow) {
W2394 := __e.Get(1)
_ = W2394
__e.TailApply(PrimFunc(symshen_4comb), W2394, W2393)
return
}, 1)

tmp5473 := Call(__e, PrimFunc(symshen_4in_1_6), W2392)


__e.TailApply(tmp5472, tmp5473)
return


}, 1)

tmp5474 := Call(__e, PrimFunc(symshen_4_5_1out), W2392)


__e.TailApply(tmp5471, tmp5474)
return


}


}, 1)

tmp5477 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2390)


tmp5478 := Call(__e, tmp5470, tmp5477)


__e.TailApply(tmp5467, tmp5478)
return


}, 1)

tmp5479 := Call(__e, ns2_1set, symshen_4_5s_1exprs2_6, tmp5466)


_ = tmp5479

tmp5480 := MakeNative(func(__e *ControlFlow) {
V2396 := __e.Get(1)
_ = V2396
tmp5537 := PrimEqual(Nil, V2396)

if True == tmp5537 {
__e.Return(Nil)
return
} else {
tmp5535 := PrimIsPair(V2396)

var ifres5515 Obj

if True == tmp5535 {
tmp5533 := PrimTail(V2396)

tmp5534 := PrimIsPair(tmp5533)

var ifres5517 Obj

if True == tmp5534 {
tmp5530 := PrimTail(V2396)

tmp5531 := PrimTail(tmp5530)

tmp5532 := PrimIsPair(tmp5531)

var ifres5519 Obj

if True == tmp5532 {
tmp5526 := PrimTail(V2396)

tmp5527 := PrimTail(tmp5526)

tmp5528 := PrimTail(tmp5527)

tmp5529 := PrimEqual(Nil, tmp5528)

var ifres5521 Obj

if True == tmp5529 {
tmp5523 := PrimTail(V2396)

tmp5524 := PrimHead(tmp5523)

tmp5525 := PrimEqual(tmp5524, symbar_b)

var ifres5522 Obj

if True == tmp5525 {
ifres5522 = True


} else {
ifres5522 = False


}

ifres5521 = ifres5522


} else {
ifres5521 = False


}

var ifres5520 Obj

if True == ifres5521 {
ifres5520 = True


} else {
ifres5520 = False


}

ifres5519 = ifres5520


} else {
ifres5519 = False


}

var ifres5518 Obj

if True == ifres5519 {
ifres5518 = True


} else {
ifres5518 = False


}

ifres5517 = ifres5518


} else {
ifres5517 = False


}

var ifres5516 Obj

if True == ifres5517 {
ifres5516 = True


} else {
ifres5516 = False


}

ifres5515 = ifres5516


} else {
ifres5515 = False


}

if True == ifres5515 {
tmp5481 := PrimHead(V2396)

tmp5482 := PrimTail(V2396)

tmp5483 := PrimTail(tmp5482)

tmp5484 := PrimCons(tmp5481, tmp5483)

__e.Return(PrimCons(symcons, tmp5484))
return


} else {
tmp5513 := PrimIsPair(V2396)

var ifres5493 Obj

if True == tmp5513 {
tmp5511 := PrimTail(V2396)

tmp5512 := PrimIsPair(tmp5511)

var ifres5495 Obj

if True == tmp5512 {
tmp5508 := PrimTail(V2396)

tmp5509 := PrimTail(tmp5508)

tmp5510 := PrimIsPair(tmp5509)

var ifres5497 Obj

if True == tmp5510 {
tmp5504 := PrimTail(V2396)

tmp5505 := PrimTail(tmp5504)

tmp5506 := PrimTail(tmp5505)

tmp5507 := PrimIsPair(tmp5506)

var ifres5499 Obj

if True == tmp5507 {
tmp5501 := PrimTail(V2396)

tmp5502 := PrimHead(tmp5501)

tmp5503 := PrimEqual(tmp5502, symbar_b)

var ifres5500 Obj

if True == tmp5503 {
ifres5500 = True


} else {
ifres5500 = False


}

ifres5499 = ifres5500


} else {
ifres5499 = False


}

var ifres5498 Obj

if True == ifres5499 {
ifres5498 = True


} else {
ifres5498 = False


}

ifres5497 = ifres5498


} else {
ifres5497 = False


}

var ifres5496 Obj

if True == ifres5497 {
ifres5496 = True


} else {
ifres5496 = False


}

ifres5495 = ifres5496


} else {
ifres5495 = False


}

var ifres5494 Obj

if True == ifres5495 {
ifres5494 = True


} else {
ifres5494 = False


}

ifres5493 = ifres5494


} else {
ifres5493 = False


}

if True == ifres5493 {
__e.Return(PrimSimpleError(MakeString("misapplication of |\n")))
return
} else {
tmp5491 := PrimIsPair(V2396)

if True == tmp5491 {
tmp5485 := PrimHead(V2396)

tmp5486 := PrimTail(V2396)

tmp5487 := Call(__e, PrimFunc(symshen_4cons_1form), tmp5486)


tmp5488 := PrimCons(tmp5487, Nil)

tmp5489 := PrimCons(tmp5485, tmp5488)

__e.Return(PrimCons(symcons, tmp5489))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.cons-form")))
return
}


}


}


}


}, 1)

tmp5538 := Call(__e, ns2_1set, symshen_4cons_1form, tmp5480)


_ = tmp5538

tmp5539 := MakeNative(func(__e *ControlFlow) {
V2397 := __e.Get(1)
_ = V2397
tmp5540 := MakeNative(func(__e *ControlFlow) {
W2398 := __e.Get(1)
_ = W2398
tmp5542 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2398)


if True == tmp5542 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2398)
return
}


}, 1)

tmp5548 := Call(__e, PrimFunc(symshen_4hds_a_2), V2397, MakeNumber(40))


var ifres5543 Obj

if True == tmp5548 {
tmp5544 := MakeNative(func(__e *ControlFlow) {
W2399 := __e.Get(1)
_ = W2399
__e.TailApply(PrimFunc(symshen_4comb), W2399, symshen_4skip)
return
}, 1)

tmp5545 := Call(__e, PrimFunc(symtail), V2397)


tmp5546 := Call(__e, tmp5544, tmp5545)


ifres5543 = tmp5546


} else {
tmp5547 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5543 = tmp5547


}

__e.TailApply(tmp5540, ifres5543)
return


}, 1)

tmp5549 := Call(__e, ns2_1set, symshen_4_5lrb_6, tmp5539)


_ = tmp5549

tmp5550 := MakeNative(func(__e *ControlFlow) {
V2400 := __e.Get(1)
_ = V2400
tmp5551 := MakeNative(func(__e *ControlFlow) {
W2401 := __e.Get(1)
_ = W2401
tmp5553 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2401)


if True == tmp5553 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2401)
return
}


}, 1)

tmp5559 := Call(__e, PrimFunc(symshen_4hds_a_2), V2400, MakeNumber(41))


var ifres5554 Obj

if True == tmp5559 {
tmp5555 := MakeNative(func(__e *ControlFlow) {
W2402 := __e.Get(1)
_ = W2402
__e.TailApply(PrimFunc(symshen_4comb), W2402, symshen_4skip)
return
}, 1)

tmp5556 := Call(__e, PrimFunc(symtail), V2400)


tmp5557 := Call(__e, tmp5555, tmp5556)


ifres5554 = tmp5557


} else {
tmp5558 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5554 = tmp5558


}

__e.TailApply(tmp5551, ifres5554)
return


}, 1)

tmp5560 := Call(__e, ns2_1set, symshen_4_5rrb_6, tmp5550)


_ = tmp5560

tmp5561 := MakeNative(func(__e *ControlFlow) {
V2403 := __e.Get(1)
_ = V2403
tmp5562 := MakeNative(func(__e *ControlFlow) {
W2404 := __e.Get(1)
_ = W2404
tmp5564 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2404)


if True == tmp5564 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2404)
return
}


}, 1)

tmp5570 := Call(__e, PrimFunc(symshen_4hds_a_2), V2403, MakeNumber(123))


var ifres5565 Obj

if True == tmp5570 {
tmp5566 := MakeNative(func(__e *ControlFlow) {
W2405 := __e.Get(1)
_ = W2405
__e.TailApply(PrimFunc(symshen_4comb), W2405, symshen_4skip)
return
}, 1)

tmp5567 := Call(__e, PrimFunc(symtail), V2403)


tmp5568 := Call(__e, tmp5566, tmp5567)


ifres5565 = tmp5568


} else {
tmp5569 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5565 = tmp5569


}

__e.TailApply(tmp5562, ifres5565)
return


}, 1)

tmp5571 := Call(__e, ns2_1set, symshen_4_5lcurly_6, tmp5561)


_ = tmp5571

tmp5572 := MakeNative(func(__e *ControlFlow) {
V2406 := __e.Get(1)
_ = V2406
tmp5573 := MakeNative(func(__e *ControlFlow) {
W2407 := __e.Get(1)
_ = W2407
tmp5575 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2407)


if True == tmp5575 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2407)
return
}


}, 1)

tmp5581 := Call(__e, PrimFunc(symshen_4hds_a_2), V2406, MakeNumber(125))


var ifres5576 Obj

if True == tmp5581 {
tmp5577 := MakeNative(func(__e *ControlFlow) {
W2408 := __e.Get(1)
_ = W2408
__e.TailApply(PrimFunc(symshen_4comb), W2408, symshen_4skip)
return
}, 1)

tmp5578 := Call(__e, PrimFunc(symtail), V2406)


tmp5579 := Call(__e, tmp5577, tmp5578)


ifres5576 = tmp5579


} else {
tmp5580 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5576 = tmp5580


}

__e.TailApply(tmp5573, ifres5576)
return


}, 1)

tmp5582 := Call(__e, ns2_1set, symshen_4_5rcurly_6, tmp5572)


_ = tmp5582

tmp5583 := MakeNative(func(__e *ControlFlow) {
V2409 := __e.Get(1)
_ = V2409
tmp5584 := MakeNative(func(__e *ControlFlow) {
W2410 := __e.Get(1)
_ = W2410
tmp5586 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2410)


if True == tmp5586 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2410)
return
}


}, 1)

tmp5592 := Call(__e, PrimFunc(symshen_4hds_a_2), V2409, MakeNumber(124))


var ifres5587 Obj

if True == tmp5592 {
tmp5588 := MakeNative(func(__e *ControlFlow) {
W2411 := __e.Get(1)
_ = W2411
__e.TailApply(PrimFunc(symshen_4comb), W2411, symshen_4skip)
return
}, 1)

tmp5589 := Call(__e, PrimFunc(symtail), V2409)


tmp5590 := Call(__e, tmp5588, tmp5589)


ifres5587 = tmp5590


} else {
tmp5591 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5587 = tmp5591


}

__e.TailApply(tmp5584, ifres5587)
return


}, 1)

tmp5593 := Call(__e, ns2_1set, symshen_4_5bar_6, tmp5583)


_ = tmp5593

tmp5594 := MakeNative(func(__e *ControlFlow) {
V2412 := __e.Get(1)
_ = V2412
tmp5595 := MakeNative(func(__e *ControlFlow) {
W2413 := __e.Get(1)
_ = W2413
tmp5597 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2413)


if True == tmp5597 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2413)
return
}


}, 1)

tmp5603 := Call(__e, PrimFunc(symshen_4hds_a_2), V2412, MakeNumber(59))


var ifres5598 Obj

if True == tmp5603 {
tmp5599 := MakeNative(func(__e *ControlFlow) {
W2414 := __e.Get(1)
_ = W2414
__e.TailApply(PrimFunc(symshen_4comb), W2414, symshen_4skip)
return
}, 1)

tmp5600 := Call(__e, PrimFunc(symtail), V2412)


tmp5601 := Call(__e, tmp5599, tmp5600)


ifres5598 = tmp5601


} else {
tmp5602 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5598 = tmp5602


}

__e.TailApply(tmp5595, ifres5598)
return


}, 1)

tmp5604 := Call(__e, ns2_1set, symshen_4_5semicolon_6, tmp5594)


_ = tmp5604

tmp5605 := MakeNative(func(__e *ControlFlow) {
V2415 := __e.Get(1)
_ = V2415
tmp5606 := MakeNative(func(__e *ControlFlow) {
W2416 := __e.Get(1)
_ = W2416
tmp5608 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2416)


if True == tmp5608 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2416)
return
}


}, 1)

tmp5614 := Call(__e, PrimFunc(symshen_4hds_a_2), V2415, MakeNumber(58))


var ifres5609 Obj

if True == tmp5614 {
tmp5610 := MakeNative(func(__e *ControlFlow) {
W2417 := __e.Get(1)
_ = W2417
__e.TailApply(PrimFunc(symshen_4comb), W2417, symshen_4skip)
return
}, 1)

tmp5611 := Call(__e, PrimFunc(symtail), V2415)


tmp5612 := Call(__e, tmp5610, tmp5611)


ifres5609 = tmp5612


} else {
tmp5613 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5609 = tmp5613


}

__e.TailApply(tmp5606, ifres5609)
return


}, 1)

tmp5615 := Call(__e, ns2_1set, symshen_4_5colon_6, tmp5605)


_ = tmp5615

tmp5616 := MakeNative(func(__e *ControlFlow) {
V2418 := __e.Get(1)
_ = V2418
tmp5617 := MakeNative(func(__e *ControlFlow) {
W2419 := __e.Get(1)
_ = W2419
tmp5619 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2419)


if True == tmp5619 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2419)
return
}


}, 1)

tmp5625 := Call(__e, PrimFunc(symshen_4hds_a_2), V2418, MakeNumber(44))


var ifres5620 Obj

if True == tmp5625 {
tmp5621 := MakeNative(func(__e *ControlFlow) {
W2420 := __e.Get(1)
_ = W2420
__e.TailApply(PrimFunc(symshen_4comb), W2420, symshen_4skip)
return
}, 1)

tmp5622 := Call(__e, PrimFunc(symtail), V2418)


tmp5623 := Call(__e, tmp5621, tmp5622)


ifres5620 = tmp5623


} else {
tmp5624 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5620 = tmp5624


}

__e.TailApply(tmp5617, ifres5620)
return


}, 1)

tmp5626 := Call(__e, ns2_1set, symshen_4_5comma_6, tmp5616)


_ = tmp5626

tmp5627 := MakeNative(func(__e *ControlFlow) {
V2421 := __e.Get(1)
_ = V2421
tmp5628 := MakeNative(func(__e *ControlFlow) {
W2422 := __e.Get(1)
_ = W2422
tmp5630 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2422)


if True == tmp5630 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2422)
return
}


}, 1)

tmp5636 := Call(__e, PrimFunc(symshen_4hds_a_2), V2421, MakeNumber(61))


var ifres5631 Obj

if True == tmp5636 {
tmp5632 := MakeNative(func(__e *ControlFlow) {
W2423 := __e.Get(1)
_ = W2423
__e.TailApply(PrimFunc(symshen_4comb), W2423, symshen_4skip)
return
}, 1)

tmp5633 := Call(__e, PrimFunc(symtail), V2421)


tmp5634 := Call(__e, tmp5632, tmp5633)


ifres5631 = tmp5634


} else {
tmp5635 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5631 = tmp5635


}

__e.TailApply(tmp5628, ifres5631)
return


}, 1)

tmp5637 := Call(__e, ns2_1set, symshen_4_5equal_6, tmp5627)


_ = tmp5637

tmp5638 := MakeNative(func(__e *ControlFlow) {
V2424 := __e.Get(1)
_ = V2424
tmp5639 := MakeNative(func(__e *ControlFlow) {
W2425 := __e.Get(1)
_ = W2425
tmp5651 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2425)


if True == tmp5651 {
tmp5640 := MakeNative(func(__e *ControlFlow) {
W2428 := __e.Get(1)
_ = W2428
tmp5642 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2428)


if True == tmp5642 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2428)
return
}


}, 1)

tmp5643 := MakeNative(func(__e *ControlFlow) {
W2429 := __e.Get(1)
_ = W2429
tmp5647 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2429)


if True == tmp5647 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5644 := MakeNative(func(__e *ControlFlow) {
W2430 := __e.Get(1)
_ = W2430
__e.TailApply(PrimFunc(symshen_4comb), W2430, symshen_4skip)
return
}, 1)

tmp5645 := Call(__e, PrimFunc(symshen_4in_1_6), W2429)


__e.TailApply(tmp5644, tmp5645)
return


}


}, 1)

tmp5648 := Call(__e, PrimFunc(symshen_4_5multiline_6), V2424)


tmp5649 := Call(__e, tmp5643, tmp5648)


__e.TailApply(tmp5640, tmp5649)
return


} else {
__e.Return(W2425)
return
}


}, 1)

tmp5652 := MakeNative(func(__e *ControlFlow) {
W2426 := __e.Get(1)
_ = W2426
tmp5656 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2426)


if True == tmp5656 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5653 := MakeNative(func(__e *ControlFlow) {
W2427 := __e.Get(1)
_ = W2427
__e.TailApply(PrimFunc(symshen_4comb), W2427, symshen_4skip)
return
}, 1)

tmp5654 := Call(__e, PrimFunc(symshen_4in_1_6), W2426)


__e.TailApply(tmp5653, tmp5654)
return


}


}, 1)

tmp5657 := Call(__e, PrimFunc(symshen_4_5singleline_6), V2424)


tmp5658 := Call(__e, tmp5652, tmp5657)


__e.TailApply(tmp5639, tmp5658)
return


}, 1)

tmp5659 := Call(__e, ns2_1set, symshen_4_5comment_6, tmp5638)


_ = tmp5659

tmp5660 := MakeNative(func(__e *ControlFlow) {
V2431 := __e.Get(1)
_ = V2431
tmp5661 := MakeNative(func(__e *ControlFlow) {
W2432 := __e.Get(1)
_ = W2432
tmp5663 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2432)


if True == tmp5663 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2432)
return
}


}, 1)

tmp5664 := MakeNative(func(__e *ControlFlow) {
W2433 := __e.Get(1)
_ = W2433
tmp5686 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2433)


if True == tmp5686 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5665 := MakeNative(func(__e *ControlFlow) {
W2434 := __e.Get(1)
_ = W2434
tmp5666 := MakeNative(func(__e *ControlFlow) {
W2435 := __e.Get(1)
_ = W2435
tmp5682 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2435)


if True == tmp5682 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5667 := MakeNative(func(__e *ControlFlow) {
W2436 := __e.Get(1)
_ = W2436
tmp5668 := MakeNative(func(__e *ControlFlow) {
W2437 := __e.Get(1)
_ = W2437
tmp5678 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2437)


if True == tmp5678 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5669 := MakeNative(func(__e *ControlFlow) {
W2438 := __e.Get(1)
_ = W2438
tmp5670 := MakeNative(func(__e *ControlFlow) {
W2439 := __e.Get(1)
_ = W2439
tmp5674 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2439)


if True == tmp5674 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5671 := MakeNative(func(__e *ControlFlow) {
W2440 := __e.Get(1)
_ = W2440
__e.TailApply(PrimFunc(symshen_4comb), W2440, symshen_4skip)
return
}, 1)

tmp5672 := Call(__e, PrimFunc(symshen_4in_1_6), W2439)


__e.TailApply(tmp5671, tmp5672)
return


}


}, 1)

tmp5675 := Call(__e, PrimFunc(symshen_4_5returns_6), W2438)


__e.TailApply(tmp5670, tmp5675)
return


}, 1)

tmp5676 := Call(__e, PrimFunc(symshen_4in_1_6), W2437)


__e.TailApply(tmp5669, tmp5676)
return


}


}, 1)

tmp5679 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), W2436)


__e.TailApply(tmp5668, tmp5679)
return


}, 1)

tmp5680 := Call(__e, PrimFunc(symshen_4in_1_6), W2435)


__e.TailApply(tmp5667, tmp5680)
return


}


}, 1)

tmp5683 := Call(__e, PrimFunc(symshen_4_5backslash_6), W2434)


__e.TailApply(tmp5666, tmp5683)
return


}, 1)

tmp5684 := Call(__e, PrimFunc(symshen_4in_1_6), W2433)


__e.TailApply(tmp5665, tmp5684)
return


}


}, 1)

tmp5687 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2431)


tmp5688 := Call(__e, tmp5664, tmp5687)


__e.TailApply(tmp5661, tmp5688)
return


}, 1)

tmp5689 := Call(__e, ns2_1set, symshen_4_5singleline_6, tmp5660)


_ = tmp5689

tmp5690 := MakeNative(func(__e *ControlFlow) {
V2441 := __e.Get(1)
_ = V2441
tmp5691 := MakeNative(func(__e *ControlFlow) {
W2442 := __e.Get(1)
_ = W2442
tmp5693 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2442)


if True == tmp5693 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2442)
return
}


}, 1)

tmp5699 := Call(__e, PrimFunc(symshen_4hds_a_2), V2441, MakeNumber(92))


var ifres5694 Obj

if True == tmp5699 {
tmp5695 := MakeNative(func(__e *ControlFlow) {
W2443 := __e.Get(1)
_ = W2443
__e.TailApply(PrimFunc(symshen_4comb), W2443, symshen_4skip)
return
}, 1)

tmp5696 := Call(__e, PrimFunc(symtail), V2441)


tmp5697 := Call(__e, tmp5695, tmp5696)


ifres5694 = tmp5697


} else {
tmp5698 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5694 = tmp5698


}

__e.TailApply(tmp5691, ifres5694)
return


}, 1)

tmp5700 := Call(__e, ns2_1set, symshen_4_5backslash_6, tmp5690)


_ = tmp5700

tmp5701 := MakeNative(func(__e *ControlFlow) {
V2444 := __e.Get(1)
_ = V2444
tmp5702 := MakeNative(func(__e *ControlFlow) {
W2445 := __e.Get(1)
_ = W2445
tmp5714 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2445)


if True == tmp5714 {
tmp5703 := MakeNative(func(__e *ControlFlow) {
W2450 := __e.Get(1)
_ = W2450
tmp5705 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2450)


if True == tmp5705 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2450)
return
}


}, 1)

tmp5706 := MakeNative(func(__e *ControlFlow) {
W2451 := __e.Get(1)
_ = W2451
tmp5710 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2451)


if True == tmp5710 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5707 := MakeNative(func(__e *ControlFlow) {
W2452 := __e.Get(1)
_ = W2452
__e.TailApply(PrimFunc(symshen_4comb), W2452, symshen_4skip)
return
}, 1)

tmp5708 := Call(__e, PrimFunc(symshen_4in_1_6), W2451)


__e.TailApply(tmp5707, tmp5708)
return


}


}, 1)

tmp5711 := Call(__e, PrimFunc(sym_5e_6), V2444)


tmp5712 := Call(__e, tmp5706, tmp5711)


__e.TailApply(tmp5703, tmp5712)
return


} else {
__e.Return(W2445)
return
}


}, 1)

tmp5715 := MakeNative(func(__e *ControlFlow) {
W2446 := __e.Get(1)
_ = W2446
tmp5725 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2446)


if True == tmp5725 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5716 := MakeNative(func(__e *ControlFlow) {
W2447 := __e.Get(1)
_ = W2447
tmp5717 := MakeNative(func(__e *ControlFlow) {
W2448 := __e.Get(1)
_ = W2448
tmp5721 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2448)


if True == tmp5721 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5718 := MakeNative(func(__e *ControlFlow) {
W2449 := __e.Get(1)
_ = W2449
__e.TailApply(PrimFunc(symshen_4comb), W2449, symshen_4skip)
return
}, 1)

tmp5719 := Call(__e, PrimFunc(symshen_4in_1_6), W2448)


__e.TailApply(tmp5718, tmp5719)
return


}


}, 1)

tmp5722 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), W2447)


__e.TailApply(tmp5717, tmp5722)
return


}, 1)

tmp5723 := Call(__e, PrimFunc(symshen_4in_1_6), W2446)


__e.TailApply(tmp5716, tmp5723)
return


}


}, 1)

tmp5726 := Call(__e, PrimFunc(symshen_4_5shortnatter_6), V2444)


tmp5727 := Call(__e, tmp5715, tmp5726)


__e.TailApply(tmp5702, tmp5727)
return


}, 1)

tmp5728 := Call(__e, ns2_1set, symshen_4_5shortnatters_6, tmp5701)


_ = tmp5728

tmp5729 := MakeNative(func(__e *ControlFlow) {
V2453 := __e.Get(1)
_ = V2453
tmp5730 := MakeNative(func(__e *ControlFlow) {
W2454 := __e.Get(1)
_ = W2454
tmp5732 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2454)


if True == tmp5732 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2454)
return
}


}, 1)

tmp5743 := PrimIsPair(V2453)

var ifres5733 Obj

if True == tmp5743 {
tmp5734 := MakeNative(func(__e *ControlFlow) {
W2455 := __e.Get(1)
_ = W2455
tmp5735 := MakeNative(func(__e *ControlFlow) {
W2456 := __e.Get(1)
_ = W2456
tmp5737 := Call(__e, PrimFunc(symshen_4return_2), W2455)


tmp5738 := PrimNot(tmp5737)

if True == tmp5738 {
__e.TailApply(PrimFunc(symshen_4comb), W2456, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp5739 := Call(__e, PrimFunc(symtail), V2453)


__e.TailApply(tmp5735, tmp5739)
return


}, 1)

tmp5740 := Call(__e, PrimFunc(symhead), V2453)


tmp5741 := Call(__e, tmp5734, tmp5740)


ifres5733 = tmp5741


} else {
tmp5742 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5733 = tmp5742


}

__e.TailApply(tmp5730, ifres5733)
return


}, 1)

tmp5744 := Call(__e, ns2_1set, symshen_4_5shortnatter_6, tmp5729)


_ = tmp5744

tmp5745 := MakeNative(func(__e *ControlFlow) {
V2457 := __e.Get(1)
_ = V2457
tmp5746 := MakeNative(func(__e *ControlFlow) {
W2458 := __e.Get(1)
_ = W2458
tmp5758 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2458)


if True == tmp5758 {
tmp5747 := MakeNative(func(__e *ControlFlow) {
W2463 := __e.Get(1)
_ = W2463
tmp5749 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2463)


if True == tmp5749 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2463)
return
}


}, 1)

tmp5750 := MakeNative(func(__e *ControlFlow) {
W2464 := __e.Get(1)
_ = W2464
tmp5754 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2464)


if True == tmp5754 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5751 := MakeNative(func(__e *ControlFlow) {
W2465 := __e.Get(1)
_ = W2465
__e.TailApply(PrimFunc(symshen_4comb), W2465, symshen_4skip)
return
}, 1)

tmp5752 := Call(__e, PrimFunc(symshen_4in_1_6), W2464)


__e.TailApply(tmp5751, tmp5752)
return


}


}, 1)

tmp5755 := Call(__e, PrimFunc(symshen_4_5return_6), V2457)


tmp5756 := Call(__e, tmp5750, tmp5755)


__e.TailApply(tmp5747, tmp5756)
return


} else {
__e.Return(W2458)
return
}


}, 1)

tmp5759 := MakeNative(func(__e *ControlFlow) {
W2459 := __e.Get(1)
_ = W2459
tmp5769 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2459)


if True == tmp5769 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5760 := MakeNative(func(__e *ControlFlow) {
W2460 := __e.Get(1)
_ = W2460
tmp5761 := MakeNative(func(__e *ControlFlow) {
W2461 := __e.Get(1)
_ = W2461
tmp5765 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2461)


if True == tmp5765 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5762 := MakeNative(func(__e *ControlFlow) {
W2462 := __e.Get(1)
_ = W2462
__e.TailApply(PrimFunc(symshen_4comb), W2462, symshen_4skip)
return
}, 1)

tmp5763 := Call(__e, PrimFunc(symshen_4in_1_6), W2461)


__e.TailApply(tmp5762, tmp5763)
return


}


}, 1)

tmp5766 := Call(__e, PrimFunc(symshen_4_5returns_6), W2460)


__e.TailApply(tmp5761, tmp5766)
return


}, 1)

tmp5767 := Call(__e, PrimFunc(symshen_4in_1_6), W2459)


__e.TailApply(tmp5760, tmp5767)
return


}


}, 1)

tmp5770 := Call(__e, PrimFunc(symshen_4_5return_6), V2457)


tmp5771 := Call(__e, tmp5759, tmp5770)


__e.TailApply(tmp5746, tmp5771)
return


}, 1)

tmp5772 := Call(__e, ns2_1set, symshen_4_5returns_6, tmp5745)


_ = tmp5772

tmp5773 := MakeNative(func(__e *ControlFlow) {
V2466 := __e.Get(1)
_ = V2466
tmp5774 := MakeNative(func(__e *ControlFlow) {
W2467 := __e.Get(1)
_ = W2467
tmp5776 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2467)


if True == tmp5776 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2467)
return
}


}, 1)

tmp5786 := PrimIsPair(V2466)

var ifres5777 Obj

if True == tmp5786 {
tmp5778 := MakeNative(func(__e *ControlFlow) {
W2468 := __e.Get(1)
_ = W2468
tmp5779 := MakeNative(func(__e *ControlFlow) {
W2469 := __e.Get(1)
_ = W2469
tmp5781 := Call(__e, PrimFunc(symshen_4return_2), W2468)


if True == tmp5781 {
__e.TailApply(PrimFunc(symshen_4comb), W2469, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp5782 := Call(__e, PrimFunc(symtail), V2466)


__e.TailApply(tmp5779, tmp5782)
return


}, 1)

tmp5783 := Call(__e, PrimFunc(symhead), V2466)


tmp5784 := Call(__e, tmp5778, tmp5783)


ifres5777 = tmp5784


} else {
tmp5785 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5777 = tmp5785


}

__e.TailApply(tmp5774, ifres5777)
return


}, 1)

tmp5787 := Call(__e, ns2_1set, symshen_4_5return_6, tmp5773)


_ = tmp5787

tmp5788 := MakeNative(func(__e *ControlFlow) {
V2470 := __e.Get(1)
_ = V2470
tmp5789 := PrimCons(MakeNumber(13), Nil)

tmp5790 := PrimCons(MakeNumber(10), tmp5789)

tmp5791 := PrimCons(MakeNumber(9), tmp5790)

__e.TailApply(PrimFunc(symelement_2), V2470, tmp5791)
return


}, 1)

tmp5792 := Call(__e, ns2_1set, symshen_4return_2, tmp5788)


_ = tmp5792

tmp5793 := MakeNative(func(__e *ControlFlow) {
V2471 := __e.Get(1)
_ = V2471
tmp5794 := MakeNative(func(__e *ControlFlow) {
W2472 := __e.Get(1)
_ = W2472
tmp5796 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2472)


if True == tmp5796 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2472)
return
}


}, 1)

tmp5797 := MakeNative(func(__e *ControlFlow) {
W2473 := __e.Get(1)
_ = W2473
tmp5813 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2473)


if True == tmp5813 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5798 := MakeNative(func(__e *ControlFlow) {
W2474 := __e.Get(1)
_ = W2474
tmp5799 := MakeNative(func(__e *ControlFlow) {
W2475 := __e.Get(1)
_ = W2475
tmp5809 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2475)


if True == tmp5809 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5800 := MakeNative(func(__e *ControlFlow) {
W2476 := __e.Get(1)
_ = W2476
tmp5801 := MakeNative(func(__e *ControlFlow) {
W2477 := __e.Get(1)
_ = W2477
tmp5805 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2477)


if True == tmp5805 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5802 := MakeNative(func(__e *ControlFlow) {
W2478 := __e.Get(1)
_ = W2478
__e.TailApply(PrimFunc(symshen_4comb), W2478, symshen_4skip)
return
}, 1)

tmp5803 := Call(__e, PrimFunc(symshen_4in_1_6), W2477)


__e.TailApply(tmp5802, tmp5803)
return


}


}, 1)

tmp5806 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2476)


__e.TailApply(tmp5801, tmp5806)
return


}, 1)

tmp5807 := Call(__e, PrimFunc(symshen_4in_1_6), W2475)


__e.TailApply(tmp5800, tmp5807)
return


}


}, 1)

tmp5810 := Call(__e, PrimFunc(symshen_4_5times_6), W2474)


__e.TailApply(tmp5799, tmp5810)
return


}, 1)

tmp5811 := Call(__e, PrimFunc(symshen_4in_1_6), W2473)


__e.TailApply(tmp5798, tmp5811)
return


}


}, 1)

tmp5814 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2471)


tmp5815 := Call(__e, tmp5797, tmp5814)


__e.TailApply(tmp5794, tmp5815)
return


}, 1)

tmp5816 := Call(__e, ns2_1set, symshen_4_5multiline_6, tmp5793)


_ = tmp5816

tmp5817 := MakeNative(func(__e *ControlFlow) {
V2479 := __e.Get(1)
_ = V2479
tmp5818 := MakeNative(func(__e *ControlFlow) {
W2480 := __e.Get(1)
_ = W2480
tmp5820 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2480)


if True == tmp5820 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2480)
return
}


}, 1)

tmp5826 := Call(__e, PrimFunc(symshen_4hds_a_2), V2479, MakeNumber(42))


var ifres5821 Obj

if True == tmp5826 {
tmp5822 := MakeNative(func(__e *ControlFlow) {
W2481 := __e.Get(1)
_ = W2481
__e.TailApply(PrimFunc(symshen_4comb), W2481, symshen_4skip)
return
}, 1)

tmp5823 := Call(__e, PrimFunc(symtail), V2479)


tmp5824 := Call(__e, tmp5822, tmp5823)


ifres5821 = tmp5824


} else {
tmp5825 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5821 = tmp5825


}

__e.TailApply(tmp5818, ifres5821)
return


}, 1)

tmp5827 := Call(__e, ns2_1set, symshen_4_5times_6, tmp5817)


_ = tmp5827

tmp5828 := MakeNative(func(__e *ControlFlow) {
V2482 := __e.Get(1)
_ = V2482
tmp5829 := MakeNative(func(__e *ControlFlow) {
W2483 := __e.Get(1)
_ = W2483
tmp5862 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2483)


if True == tmp5862 {
tmp5830 := MakeNative(func(__e *ControlFlow) {
W2488 := __e.Get(1)
_ = W2488
tmp5847 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2488)


if True == tmp5847 {
tmp5831 := MakeNative(func(__e *ControlFlow) {
W2493 := __e.Get(1)
_ = W2493
tmp5833 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2493)


if True == tmp5833 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2493)
return
}


}, 1)

tmp5845 := PrimIsPair(V2482)

var ifres5834 Obj

if True == tmp5845 {
tmp5835 := MakeNative(func(__e *ControlFlow) {
W2494 := __e.Get(1)
_ = W2494
tmp5836 := MakeNative(func(__e *ControlFlow) {
W2495 := __e.Get(1)
_ = W2495
tmp5840 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2495)


if True == tmp5840 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5837 := MakeNative(func(__e *ControlFlow) {
W2496 := __e.Get(1)
_ = W2496
__e.TailApply(PrimFunc(symshen_4comb), W2496, symshen_4skip)
return
}, 1)

tmp5838 := Call(__e, PrimFunc(symshen_4in_1_6), W2495)


__e.TailApply(tmp5837, tmp5838)
return


}


}, 1)

tmp5841 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2494)


__e.TailApply(tmp5836, tmp5841)
return


}, 1)

tmp5842 := Call(__e, PrimFunc(symtail), V2482)


tmp5843 := Call(__e, tmp5835, tmp5842)


ifres5834 = tmp5843


} else {
tmp5844 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5834 = tmp5844


}

__e.TailApply(tmp5831, ifres5834)
return


} else {
__e.Return(W2488)
return
}


}, 1)

tmp5848 := MakeNative(func(__e *ControlFlow) {
W2489 := __e.Get(1)
_ = W2489
tmp5858 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2489)


if True == tmp5858 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5849 := MakeNative(func(__e *ControlFlow) {
W2490 := __e.Get(1)
_ = W2490
tmp5850 := MakeNative(func(__e *ControlFlow) {
W2491 := __e.Get(1)
_ = W2491
tmp5854 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2491)


if True == tmp5854 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5851 := MakeNative(func(__e *ControlFlow) {
W2492 := __e.Get(1)
_ = W2492
__e.TailApply(PrimFunc(symshen_4comb), W2492, symshen_4skip)
return
}, 1)

tmp5852 := Call(__e, PrimFunc(symshen_4in_1_6), W2491)


__e.TailApply(tmp5851, tmp5852)
return


}


}, 1)

tmp5855 := Call(__e, PrimFunc(symshen_4_5backslash_6), W2490)


__e.TailApply(tmp5850, tmp5855)
return


}, 1)

tmp5856 := Call(__e, PrimFunc(symshen_4in_1_6), W2489)


__e.TailApply(tmp5849, tmp5856)
return


}


}, 1)

tmp5859 := Call(__e, PrimFunc(symshen_4_5times_6), V2482)


tmp5860 := Call(__e, tmp5848, tmp5859)


__e.TailApply(tmp5830, tmp5860)
return


} else {
__e.Return(W2483)
return
}


}, 1)

tmp5863 := MakeNative(func(__e *ControlFlow) {
W2484 := __e.Get(1)
_ = W2484
tmp5873 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2484)


if True == tmp5873 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5864 := MakeNative(func(__e *ControlFlow) {
W2485 := __e.Get(1)
_ = W2485
tmp5865 := MakeNative(func(__e *ControlFlow) {
W2486 := __e.Get(1)
_ = W2486
tmp5869 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2486)


if True == tmp5869 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5866 := MakeNative(func(__e *ControlFlow) {
W2487 := __e.Get(1)
_ = W2487
__e.TailApply(PrimFunc(symshen_4comb), W2487, symshen_4skip)
return
}, 1)

tmp5867 := Call(__e, PrimFunc(symshen_4in_1_6), W2486)


__e.TailApply(tmp5866, tmp5867)
return


}


}, 1)

tmp5870 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2485)


__e.TailApply(tmp5865, tmp5870)
return


}, 1)

tmp5871 := Call(__e, PrimFunc(symshen_4in_1_6), W2484)


__e.TailApply(tmp5864, tmp5871)
return


}


}, 1)

tmp5874 := Call(__e, PrimFunc(symshen_4_5comment_6), V2482)


tmp5875 := Call(__e, tmp5863, tmp5874)


__e.TailApply(tmp5829, tmp5875)
return


}, 1)

tmp5876 := Call(__e, ns2_1set, symshen_4_5longnatter_6, tmp5828)


_ = tmp5876

tmp5877 := MakeNative(func(__e *ControlFlow) {
V2497 := __e.Get(1)
_ = V2497
tmp5878 := MakeNative(func(__e *ControlFlow) {
W2498 := __e.Get(1)
_ = W2498
tmp5909 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2498)


if True == tmp5909 {
tmp5879 := MakeNative(func(__e *ControlFlow) {
W2502 := __e.Get(1)
_ = W2502
tmp5898 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2502)


if True == tmp5898 {
tmp5880 := MakeNative(func(__e *ControlFlow) {
W2506 := __e.Get(1)
_ = W2506
tmp5882 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2506)


if True == tmp5882 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2506)
return
}


}, 1)

tmp5883 := MakeNative(func(__e *ControlFlow) {
W2507 := __e.Get(1)
_ = W2507
tmp5894 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2507)


if True == tmp5894 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5884 := MakeNative(func(__e *ControlFlow) {
W2508 := __e.Get(1)
_ = W2508
tmp5885 := MakeNative(func(__e *ControlFlow) {
W2509 := __e.Get(1)
_ = W2509
tmp5890 := PrimEqual(W2508, MakeString("<>"))

var ifres5886 Obj

if True == tmp5890 {
tmp5887 := PrimCons(MakeNumber(0), Nil)

tmp5888 := PrimCons(symvector, tmp5887)

ifres5886 = tmp5888


} else {
tmp5889 := PrimIntern(W2508)

ifres5886 = tmp5889


}

__e.TailApply(PrimFunc(symshen_4comb), W2509, ifres5886)
return


}, 1)

tmp5891 := Call(__e, PrimFunc(symshen_4in_1_6), W2507)


__e.TailApply(tmp5885, tmp5891)
return


}, 1)

tmp5892 := Call(__e, PrimFunc(symshen_4_5_1out), W2507)


__e.TailApply(tmp5884, tmp5892)
return


}


}, 1)

tmp5895 := Call(__e, PrimFunc(symshen_4_5sym_6), V2497)


tmp5896 := Call(__e, tmp5883, tmp5895)


__e.TailApply(tmp5880, tmp5896)
return


} else {
__e.Return(W2502)
return
}


}, 1)

tmp5899 := MakeNative(func(__e *ControlFlow) {
W2503 := __e.Get(1)
_ = W2503
tmp5905 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2503)


if True == tmp5905 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5900 := MakeNative(func(__e *ControlFlow) {
W2504 := __e.Get(1)
_ = W2504
tmp5901 := MakeNative(func(__e *ControlFlow) {
W2505 := __e.Get(1)
_ = W2505
__e.TailApply(PrimFunc(symshen_4comb), W2505, W2504)
return
}, 1)

tmp5902 := Call(__e, PrimFunc(symshen_4in_1_6), W2503)


__e.TailApply(tmp5901, tmp5902)
return


}, 1)

tmp5903 := Call(__e, PrimFunc(symshen_4_5_1out), W2503)


__e.TailApply(tmp5900, tmp5903)
return


}


}, 1)

tmp5906 := Call(__e, PrimFunc(symshen_4_5number_6), V2497)


tmp5907 := Call(__e, tmp5899, tmp5906)


__e.TailApply(tmp5879, tmp5907)
return


} else {
__e.Return(W2498)
return
}


}, 1)

tmp5910 := MakeNative(func(__e *ControlFlow) {
W2499 := __e.Get(1)
_ = W2499
tmp5916 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2499)


if True == tmp5916 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5911 := MakeNative(func(__e *ControlFlow) {
W2500 := __e.Get(1)
_ = W2500
tmp5912 := MakeNative(func(__e *ControlFlow) {
W2501 := __e.Get(1)
_ = W2501
__e.TailApply(PrimFunc(symshen_4comb), W2501, W2500)
return
}, 1)

tmp5913 := Call(__e, PrimFunc(symshen_4in_1_6), W2499)


__e.TailApply(tmp5912, tmp5913)
return


}, 1)

tmp5914 := Call(__e, PrimFunc(symshen_4_5_1out), W2499)


__e.TailApply(tmp5911, tmp5914)
return


}


}, 1)

tmp5917 := Call(__e, PrimFunc(symshen_4_5str_6), V2497)


tmp5918 := Call(__e, tmp5910, tmp5917)


__e.TailApply(tmp5878, tmp5918)
return


}, 1)

tmp5919 := Call(__e, ns2_1set, symshen_4_5atom_6, tmp5877)


_ = tmp5919

tmp5920 := MakeNative(func(__e *ControlFlow) {
V2510 := __e.Get(1)
_ = V2510
tmp5921 := MakeNative(func(__e *ControlFlow) {
W2511 := __e.Get(1)
_ = W2511
tmp5923 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2511)


if True == tmp5923 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2511)
return
}


}, 1)

tmp5924 := MakeNative(func(__e *ControlFlow) {
W2512 := __e.Get(1)
_ = W2512
tmp5939 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2512)


if True == tmp5939 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5925 := MakeNative(func(__e *ControlFlow) {
W2513 := __e.Get(1)
_ = W2513
tmp5926 := MakeNative(func(__e *ControlFlow) {
W2514 := __e.Get(1)
_ = W2514
tmp5927 := MakeNative(func(__e *ControlFlow) {
W2515 := __e.Get(1)
_ = W2515
tmp5934 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2515)


if True == tmp5934 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5928 := MakeNative(func(__e *ControlFlow) {
W2516 := __e.Get(1)
_ = W2516
tmp5929 := MakeNative(func(__e *ControlFlow) {
W2517 := __e.Get(1)
_ = W2517
tmp5930 := PrimStringConcat(W2513, W2516)

__e.TailApply(PrimFunc(symshen_4comb), W2517, tmp5930)
return


}, 1)

tmp5931 := Call(__e, PrimFunc(symshen_4in_1_6), W2515)


__e.TailApply(tmp5929, tmp5931)
return


}, 1)

tmp5932 := Call(__e, PrimFunc(symshen_4_5_1out), W2515)


__e.TailApply(tmp5928, tmp5932)
return


}


}, 1)

tmp5935 := Call(__e, PrimFunc(symshen_4_5alphanums_6), W2514)


__e.TailApply(tmp5927, tmp5935)
return


}, 1)

tmp5936 := Call(__e, PrimFunc(symshen_4in_1_6), W2512)


__e.TailApply(tmp5926, tmp5936)
return


}, 1)

tmp5937 := Call(__e, PrimFunc(symshen_4_5_1out), W2512)


__e.TailApply(tmp5925, tmp5937)
return


}


}, 1)

tmp5940 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2510)


tmp5941 := Call(__e, tmp5924, tmp5940)


__e.TailApply(tmp5921, tmp5941)
return


}, 1)

tmp5942 := Call(__e, ns2_1set, symshen_4_5sym_6, tmp5920)


_ = tmp5942

tmp5943 := MakeNative(func(__e *ControlFlow) {
V2518 := __e.Get(1)
_ = V2518
tmp5944 := MakeNative(func(__e *ControlFlow) {
W2519 := __e.Get(1)
_ = W2519
tmp5946 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2519)


if True == tmp5946 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2519)
return
}


}, 1)

tmp5957 := PrimIsPair(V2518)

var ifres5947 Obj

if True == tmp5957 {
tmp5948 := MakeNative(func(__e *ControlFlow) {
W2520 := __e.Get(1)
_ = W2520
tmp5949 := MakeNative(func(__e *ControlFlow) {
W2521 := __e.Get(1)
_ = W2521
tmp5952 := Call(__e, PrimFunc(symshen_4alpha_2), W2520)


if True == tmp5952 {
tmp5950 := PrimNumberToString(W2520)

__e.TailApply(PrimFunc(symshen_4comb), W2521, tmp5950)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp5953 := Call(__e, PrimFunc(symtail), V2518)


__e.TailApply(tmp5949, tmp5953)
return


}, 1)

tmp5954 := Call(__e, PrimFunc(symhead), V2518)


tmp5955 := Call(__e, tmp5948, tmp5954)


ifres5947 = tmp5955


} else {
tmp5956 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5947 = tmp5956


}

__e.TailApply(tmp5944, ifres5947)
return


}, 1)

tmp5958 := Call(__e, ns2_1set, symshen_4_5alpha_6, tmp5943)


_ = tmp5958

tmp5959 := MakeNative(func(__e *ControlFlow) {
V2522 := __e.Get(1)
_ = V2522
tmp5966 := Call(__e, PrimFunc(symshen_4lowercase_2), V2522)


if True == tmp5966 {
__e.Return(True)
return
} else {
tmp5964 := Call(__e, PrimFunc(symshen_4uppercase_2), V2522)


var ifres5961 Obj

if True == tmp5964 {
ifres5961 = True


} else {
tmp5963 := Call(__e, PrimFunc(symshen_4misc_2), V2522)


var ifres5962 Obj

if True == tmp5963 {
ifres5962 = True


} else {
ifres5962 = False


}

ifres5961 = ifres5962


}

if True == ifres5961 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp5967 := Call(__e, ns2_1set, symshen_4alpha_2, tmp5959)


_ = tmp5967

tmp5968 := MakeNative(func(__e *ControlFlow) {
V2523 := __e.Get(1)
_ = V2523
tmp5972 := PrimGreatEqual(V2523, MakeNumber(97))

if True == tmp5972 {
tmp5970 := PrimLessEqual(V2523, MakeNumber(122))

if True == tmp5970 {
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

tmp5973 := Call(__e, ns2_1set, symshen_4lowercase_2, tmp5968)


_ = tmp5973

tmp5974 := MakeNative(func(__e *ControlFlow) {
V2524 := __e.Get(1)
_ = V2524
tmp5978 := PrimGreatEqual(V2524, MakeNumber(65))

if True == tmp5978 {
tmp5976 := PrimLessEqual(V2524, MakeNumber(90))

if True == tmp5976 {
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

tmp5979 := Call(__e, ns2_1set, symshen_4uppercase_2, tmp5974)


_ = tmp5979

tmp5980 := MakeNative(func(__e *ControlFlow) {
V2525 := __e.Get(1)
_ = V2525
tmp5981 := PrimCons(MakeNumber(96), Nil)

tmp5982 := PrimCons(MakeNumber(35), tmp5981)

tmp5983 := PrimCons(MakeNumber(39), tmp5982)

tmp5984 := PrimCons(MakeNumber(37), tmp5983)

tmp5985 := PrimCons(MakeNumber(38), tmp5984)

tmp5986 := PrimCons(MakeNumber(60), tmp5985)

tmp5987 := PrimCons(MakeNumber(62), tmp5986)

tmp5988 := PrimCons(MakeNumber(46), tmp5987)

tmp5989 := PrimCons(MakeNumber(126), tmp5988)

tmp5990 := PrimCons(MakeNumber(64), tmp5989)

tmp5991 := PrimCons(MakeNumber(33), tmp5990)

tmp5992 := PrimCons(MakeNumber(36), tmp5991)

tmp5993 := PrimCons(MakeNumber(63), tmp5992)

tmp5994 := PrimCons(MakeNumber(95), tmp5993)

tmp5995 := PrimCons(MakeNumber(43), tmp5994)

tmp5996 := PrimCons(MakeNumber(47), tmp5995)

tmp5997 := PrimCons(MakeNumber(42), tmp5996)

tmp5998 := PrimCons(MakeNumber(45), tmp5997)

tmp5999 := PrimCons(MakeNumber(61), tmp5998)

__e.TailApply(PrimFunc(symelement_2), V2525, tmp5999)
return


}, 1)

tmp6000 := Call(__e, ns2_1set, symshen_4misc_2, tmp5980)


_ = tmp6000

tmp6001 := MakeNative(func(__e *ControlFlow) {
V2526 := __e.Get(1)
_ = V2526
tmp6002 := MakeNative(func(__e *ControlFlow) {
W2527 := __e.Get(1)
_ = W2527
tmp6014 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2527)


if True == tmp6014 {
tmp6003 := MakeNative(func(__e *ControlFlow) {
W2534 := __e.Get(1)
_ = W2534
tmp6005 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2534)


if True == tmp6005 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2534)
return
}


}, 1)

tmp6006 := MakeNative(func(__e *ControlFlow) {
W2535 := __e.Get(1)
_ = W2535
tmp6010 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2535)


if True == tmp6010 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6007 := MakeNative(func(__e *ControlFlow) {
W2536 := __e.Get(1)
_ = W2536
__e.TailApply(PrimFunc(symshen_4comb), W2536, MakeString(""))
return
}, 1)

tmp6008 := Call(__e, PrimFunc(symshen_4in_1_6), W2535)


__e.TailApply(tmp6007, tmp6008)
return


}


}, 1)

tmp6011 := Call(__e, PrimFunc(sym_5e_6), V2526)


tmp6012 := Call(__e, tmp6006, tmp6011)


__e.TailApply(tmp6003, tmp6012)
return


} else {
__e.Return(W2527)
return
}


}, 1)

tmp6015 := MakeNative(func(__e *ControlFlow) {
W2528 := __e.Get(1)
_ = W2528
tmp6030 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2528)


if True == tmp6030 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6016 := MakeNative(func(__e *ControlFlow) {
W2529 := __e.Get(1)
_ = W2529
tmp6017 := MakeNative(func(__e *ControlFlow) {
W2530 := __e.Get(1)
_ = W2530
tmp6018 := MakeNative(func(__e *ControlFlow) {
W2531 := __e.Get(1)
_ = W2531
tmp6025 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2531)


if True == tmp6025 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6019 := MakeNative(func(__e *ControlFlow) {
W2532 := __e.Get(1)
_ = W2532
tmp6020 := MakeNative(func(__e *ControlFlow) {
W2533 := __e.Get(1)
_ = W2533
tmp6021 := PrimStringConcat(W2529, W2532)

__e.TailApply(PrimFunc(symshen_4comb), W2533, tmp6021)
return


}, 1)

tmp6022 := Call(__e, PrimFunc(symshen_4in_1_6), W2531)


__e.TailApply(tmp6020, tmp6022)
return


}, 1)

tmp6023 := Call(__e, PrimFunc(symshen_4_5_1out), W2531)


__e.TailApply(tmp6019, tmp6023)
return


}


}, 1)

tmp6026 := Call(__e, PrimFunc(symshen_4_5alphanums_6), W2530)


__e.TailApply(tmp6018, tmp6026)
return


}, 1)

tmp6027 := Call(__e, PrimFunc(symshen_4in_1_6), W2528)


__e.TailApply(tmp6017, tmp6027)
return


}, 1)

tmp6028 := Call(__e, PrimFunc(symshen_4_5_1out), W2528)


__e.TailApply(tmp6016, tmp6028)
return


}


}, 1)

tmp6031 := Call(__e, PrimFunc(symshen_4_5alphanum_6), V2526)


tmp6032 := Call(__e, tmp6015, tmp6031)


__e.TailApply(tmp6002, tmp6032)
return


}, 1)

tmp6033 := Call(__e, ns2_1set, symshen_4_5alphanums_6, tmp6001)


_ = tmp6033

tmp6034 := MakeNative(func(__e *ControlFlow) {
V2537 := __e.Get(1)
_ = V2537
tmp6035 := MakeNative(func(__e *ControlFlow) {
W2538 := __e.Get(1)
_ = W2538
tmp6049 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2538)


if True == tmp6049 {
tmp6036 := MakeNative(func(__e *ControlFlow) {
W2542 := __e.Get(1)
_ = W2542
tmp6038 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2542)


if True == tmp6038 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2542)
return
}


}, 1)

tmp6039 := MakeNative(func(__e *ControlFlow) {
W2543 := __e.Get(1)
_ = W2543
tmp6045 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2543)


if True == tmp6045 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6040 := MakeNative(func(__e *ControlFlow) {
W2544 := __e.Get(1)
_ = W2544
tmp6041 := MakeNative(func(__e *ControlFlow) {
W2545 := __e.Get(1)
_ = W2545
__e.TailApply(PrimFunc(symshen_4comb), W2545, W2544)
return
}, 1)

tmp6042 := Call(__e, PrimFunc(symshen_4in_1_6), W2543)


__e.TailApply(tmp6041, tmp6042)
return


}, 1)

tmp6043 := Call(__e, PrimFunc(symshen_4_5_1out), W2543)


__e.TailApply(tmp6040, tmp6043)
return


}


}, 1)

tmp6046 := Call(__e, PrimFunc(symshen_4_5numeral_6), V2537)


tmp6047 := Call(__e, tmp6039, tmp6046)


__e.TailApply(tmp6036, tmp6047)
return


} else {
__e.Return(W2538)
return
}


}, 1)

tmp6050 := MakeNative(func(__e *ControlFlow) {
W2539 := __e.Get(1)
_ = W2539
tmp6056 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2539)


if True == tmp6056 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6051 := MakeNative(func(__e *ControlFlow) {
W2540 := __e.Get(1)
_ = W2540
tmp6052 := MakeNative(func(__e *ControlFlow) {
W2541 := __e.Get(1)
_ = W2541
__e.TailApply(PrimFunc(symshen_4comb), W2541, W2540)
return
}, 1)

tmp6053 := Call(__e, PrimFunc(symshen_4in_1_6), W2539)


__e.TailApply(tmp6052, tmp6053)
return


}, 1)

tmp6054 := Call(__e, PrimFunc(symshen_4_5_1out), W2539)


__e.TailApply(tmp6051, tmp6054)
return


}


}, 1)

tmp6057 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2537)


tmp6058 := Call(__e, tmp6050, tmp6057)


__e.TailApply(tmp6035, tmp6058)
return


}, 1)

tmp6059 := Call(__e, ns2_1set, symshen_4_5alphanum_6, tmp6034)


_ = tmp6059

tmp6060 := MakeNative(func(__e *ControlFlow) {
V2546 := __e.Get(1)
_ = V2546
tmp6061 := MakeNative(func(__e *ControlFlow) {
W2547 := __e.Get(1)
_ = W2547
tmp6063 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2547)


if True == tmp6063 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2547)
return
}


}, 1)

tmp6074 := PrimIsPair(V2546)

var ifres6064 Obj

if True == tmp6074 {
tmp6065 := MakeNative(func(__e *ControlFlow) {
W2548 := __e.Get(1)
_ = W2548
tmp6066 := MakeNative(func(__e *ControlFlow) {
W2549 := __e.Get(1)
_ = W2549
tmp6069 := Call(__e, PrimFunc(symshen_4digit_2), W2548)


if True == tmp6069 {
tmp6067 := PrimNumberToString(W2548)

__e.TailApply(PrimFunc(symshen_4comb), W2549, tmp6067)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6070 := Call(__e, PrimFunc(symtail), V2546)


__e.TailApply(tmp6066, tmp6070)
return


}, 1)

tmp6071 := Call(__e, PrimFunc(symhead), V2546)


tmp6072 := Call(__e, tmp6065, tmp6071)


ifres6064 = tmp6072


} else {
tmp6073 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6064 = tmp6073


}

__e.TailApply(tmp6061, ifres6064)
return


}, 1)

tmp6075 := Call(__e, ns2_1set, symshen_4_5numeral_6, tmp6060)


_ = tmp6075

tmp6076 := MakeNative(func(__e *ControlFlow) {
V2550 := __e.Get(1)
_ = V2550
tmp6080 := PrimGreatEqual(V2550, MakeNumber(48))

if True == tmp6080 {
tmp6078 := PrimLessEqual(V2550, MakeNumber(57))

if True == tmp6078 {
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

tmp6081 := Call(__e, ns2_1set, symshen_4digit_2, tmp6076)


_ = tmp6081

tmp6082 := MakeNative(func(__e *ControlFlow) {
V2551 := __e.Get(1)
_ = V2551
tmp6083 := MakeNative(func(__e *ControlFlow) {
W2552 := __e.Get(1)
_ = W2552
tmp6085 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2552)


if True == tmp6085 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2552)
return
}


}, 1)

tmp6086 := MakeNative(func(__e *ControlFlow) {
W2553 := __e.Get(1)
_ = W2553
tmp6104 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2553)


if True == tmp6104 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6087 := MakeNative(func(__e *ControlFlow) {
W2554 := __e.Get(1)
_ = W2554
tmp6088 := MakeNative(func(__e *ControlFlow) {
W2555 := __e.Get(1)
_ = W2555
tmp6100 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2555)


if True == tmp6100 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6089 := MakeNative(func(__e *ControlFlow) {
W2556 := __e.Get(1)
_ = W2556
tmp6090 := MakeNative(func(__e *ControlFlow) {
W2557 := __e.Get(1)
_ = W2557
tmp6091 := MakeNative(func(__e *ControlFlow) {
W2558 := __e.Get(1)
_ = W2558
tmp6095 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2558)


if True == tmp6095 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6092 := MakeNative(func(__e *ControlFlow) {
W2559 := __e.Get(1)
_ = W2559
__e.TailApply(PrimFunc(symshen_4comb), W2559, W2556)
return
}, 1)

tmp6093 := Call(__e, PrimFunc(symshen_4in_1_6), W2558)


__e.TailApply(tmp6092, tmp6093)
return


}


}, 1)

tmp6096 := Call(__e, PrimFunc(symshen_4_5dbq_6), W2557)


__e.TailApply(tmp6091, tmp6096)
return


}, 1)

tmp6097 := Call(__e, PrimFunc(symshen_4in_1_6), W2555)


__e.TailApply(tmp6090, tmp6097)
return


}, 1)

tmp6098 := Call(__e, PrimFunc(symshen_4_5_1out), W2555)


__e.TailApply(tmp6089, tmp6098)
return


}


}, 1)

tmp6101 := Call(__e, PrimFunc(symshen_4_5strcontents_6), W2554)


__e.TailApply(tmp6088, tmp6101)
return


}, 1)

tmp6102 := Call(__e, PrimFunc(symshen_4in_1_6), W2553)


__e.TailApply(tmp6087, tmp6102)
return


}


}, 1)

tmp6105 := Call(__e, PrimFunc(symshen_4_5dbq_6), V2551)


tmp6106 := Call(__e, tmp6086, tmp6105)


__e.TailApply(tmp6083, tmp6106)
return


}, 1)

tmp6107 := Call(__e, ns2_1set, symshen_4_5str_6, tmp6082)


_ = tmp6107

tmp6108 := MakeNative(func(__e *ControlFlow) {
V2560 := __e.Get(1)
_ = V2560
tmp6109 := MakeNative(func(__e *ControlFlow) {
W2561 := __e.Get(1)
_ = W2561
tmp6111 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2561)


if True == tmp6111 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2561)
return
}


}, 1)

tmp6117 := Call(__e, PrimFunc(symshen_4hds_a_2), V2560, MakeNumber(34))


var ifres6112 Obj

if True == tmp6117 {
tmp6113 := MakeNative(func(__e *ControlFlow) {
W2562 := __e.Get(1)
_ = W2562
__e.TailApply(PrimFunc(symshen_4comb), W2562, symshen_4skip)
return
}, 1)

tmp6114 := Call(__e, PrimFunc(symtail), V2560)


tmp6115 := Call(__e, tmp6113, tmp6114)


ifres6112 = tmp6115


} else {
tmp6116 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6112 = tmp6116


}

__e.TailApply(tmp6109, ifres6112)
return


}, 1)

tmp6118 := Call(__e, ns2_1set, symshen_4_5dbq_6, tmp6108)


_ = tmp6118

tmp6119 := MakeNative(func(__e *ControlFlow) {
V2563 := __e.Get(1)
_ = V2563
tmp6120 := MakeNative(func(__e *ControlFlow) {
W2564 := __e.Get(1)
_ = W2564
tmp6132 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2564)


if True == tmp6132 {
tmp6121 := MakeNative(func(__e *ControlFlow) {
W2571 := __e.Get(1)
_ = W2571
tmp6123 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2571)


if True == tmp6123 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2571)
return
}


}, 1)

tmp6124 := MakeNative(func(__e *ControlFlow) {
W2572 := __e.Get(1)
_ = W2572
tmp6128 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2572)


if True == tmp6128 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6125 := MakeNative(func(__e *ControlFlow) {
W2573 := __e.Get(1)
_ = W2573
__e.TailApply(PrimFunc(symshen_4comb), W2573, MakeString(""))
return
}, 1)

tmp6126 := Call(__e, PrimFunc(symshen_4in_1_6), W2572)


__e.TailApply(tmp6125, tmp6126)
return


}


}, 1)

tmp6129 := Call(__e, PrimFunc(sym_5e_6), V2563)


tmp6130 := Call(__e, tmp6124, tmp6129)


__e.TailApply(tmp6121, tmp6130)
return


} else {
__e.Return(W2564)
return
}


}, 1)

tmp6133 := MakeNative(func(__e *ControlFlow) {
W2565 := __e.Get(1)
_ = W2565
tmp6148 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2565)


if True == tmp6148 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6134 := MakeNative(func(__e *ControlFlow) {
W2566 := __e.Get(1)
_ = W2566
tmp6135 := MakeNative(func(__e *ControlFlow) {
W2567 := __e.Get(1)
_ = W2567
tmp6136 := MakeNative(func(__e *ControlFlow) {
W2568 := __e.Get(1)
_ = W2568
tmp6143 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2568)


if True == tmp6143 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6137 := MakeNative(func(__e *ControlFlow) {
W2569 := __e.Get(1)
_ = W2569
tmp6138 := MakeNative(func(__e *ControlFlow) {
W2570 := __e.Get(1)
_ = W2570
tmp6139 := PrimStringConcat(W2566, W2569)

__e.TailApply(PrimFunc(symshen_4comb), W2570, tmp6139)
return


}, 1)

tmp6140 := Call(__e, PrimFunc(symshen_4in_1_6), W2568)


__e.TailApply(tmp6138, tmp6140)
return


}, 1)

tmp6141 := Call(__e, PrimFunc(symshen_4_5_1out), W2568)


__e.TailApply(tmp6137, tmp6141)
return


}


}, 1)

tmp6144 := Call(__e, PrimFunc(symshen_4_5strcontents_6), W2567)


__e.TailApply(tmp6136, tmp6144)
return


}, 1)

tmp6145 := Call(__e, PrimFunc(symshen_4in_1_6), W2565)


__e.TailApply(tmp6135, tmp6145)
return


}, 1)

tmp6146 := Call(__e, PrimFunc(symshen_4_5_1out), W2565)


__e.TailApply(tmp6134, tmp6146)
return


}


}, 1)

tmp6149 := Call(__e, PrimFunc(symshen_4_5strc_6), V2563)


tmp6150 := Call(__e, tmp6133, tmp6149)


__e.TailApply(tmp6120, tmp6150)
return


}, 1)

tmp6151 := Call(__e, ns2_1set, symshen_4_5strcontents_6, tmp6119)


_ = tmp6151

tmp6152 := MakeNative(func(__e *ControlFlow) {
V2574 := __e.Get(1)
_ = V2574
tmp6153 := MakeNative(func(__e *ControlFlow) {
W2575 := __e.Get(1)
_ = W2575
tmp6167 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2575)


if True == tmp6167 {
tmp6154 := MakeNative(func(__e *ControlFlow) {
W2579 := __e.Get(1)
_ = W2579
tmp6156 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2579)


if True == tmp6156 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2579)
return
}


}, 1)

tmp6157 := MakeNative(func(__e *ControlFlow) {
W2580 := __e.Get(1)
_ = W2580
tmp6163 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2580)


if True == tmp6163 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6158 := MakeNative(func(__e *ControlFlow) {
W2581 := __e.Get(1)
_ = W2581
tmp6159 := MakeNative(func(__e *ControlFlow) {
W2582 := __e.Get(1)
_ = W2582
__e.TailApply(PrimFunc(symshen_4comb), W2582, W2581)
return
}, 1)

tmp6160 := Call(__e, PrimFunc(symshen_4in_1_6), W2580)


__e.TailApply(tmp6159, tmp6160)
return


}, 1)

tmp6161 := Call(__e, PrimFunc(symshen_4_5_1out), W2580)


__e.TailApply(tmp6158, tmp6161)
return


}


}, 1)

tmp6164 := Call(__e, PrimFunc(symshen_4_5notdbq_6), V2574)


tmp6165 := Call(__e, tmp6157, tmp6164)


__e.TailApply(tmp6154, tmp6165)
return


} else {
__e.Return(W2575)
return
}


}, 1)

tmp6168 := MakeNative(func(__e *ControlFlow) {
W2576 := __e.Get(1)
_ = W2576
tmp6174 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2576)


if True == tmp6174 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6169 := MakeNative(func(__e *ControlFlow) {
W2577 := __e.Get(1)
_ = W2577
tmp6170 := MakeNative(func(__e *ControlFlow) {
W2578 := __e.Get(1)
_ = W2578
__e.TailApply(PrimFunc(symshen_4comb), W2578, W2577)
return
}, 1)

tmp6171 := Call(__e, PrimFunc(symshen_4in_1_6), W2576)


__e.TailApply(tmp6170, tmp6171)
return


}, 1)

tmp6172 := Call(__e, PrimFunc(symshen_4_5_1out), W2576)


__e.TailApply(tmp6169, tmp6172)
return


}


}, 1)

tmp6175 := Call(__e, PrimFunc(symshen_4_5control_6), V2574)


tmp6176 := Call(__e, tmp6168, tmp6175)


__e.TailApply(tmp6153, tmp6176)
return


}, 1)

tmp6177 := Call(__e, ns2_1set, symshen_4_5strc_6, tmp6152)


_ = tmp6177

tmp6178 := MakeNative(func(__e *ControlFlow) {
V2583 := __e.Get(1)
_ = V2583
tmp6179 := MakeNative(func(__e *ControlFlow) {
W2584 := __e.Get(1)
_ = W2584
tmp6181 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2584)


if True == tmp6181 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2584)
return
}


}, 1)

tmp6182 := MakeNative(func(__e *ControlFlow) {
W2585 := __e.Get(1)
_ = W2585
tmp6207 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2585)


if True == tmp6207 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6183 := MakeNative(func(__e *ControlFlow) {
W2586 := __e.Get(1)
_ = W2586
tmp6184 := MakeNative(func(__e *ControlFlow) {
W2587 := __e.Get(1)
_ = W2587
tmp6203 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2587)


if True == tmp6203 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6185 := MakeNative(func(__e *ControlFlow) {
W2588 := __e.Get(1)
_ = W2588
tmp6186 := MakeNative(func(__e *ControlFlow) {
W2589 := __e.Get(1)
_ = W2589
tmp6199 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2589)


if True == tmp6199 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6187 := MakeNative(func(__e *ControlFlow) {
W2590 := __e.Get(1)
_ = W2590
tmp6188 := MakeNative(func(__e *ControlFlow) {
W2591 := __e.Get(1)
_ = W2591
tmp6189 := MakeNative(func(__e *ControlFlow) {
W2592 := __e.Get(1)
_ = W2592
tmp6194 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2592)


if True == tmp6194 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6190 := MakeNative(func(__e *ControlFlow) {
W2593 := __e.Get(1)
_ = W2593
tmp6191 := PrimNumberToString(W2590)

__e.TailApply(PrimFunc(symshen_4comb), W2593, tmp6191)
return


}, 1)

tmp6192 := Call(__e, PrimFunc(symshen_4in_1_6), W2592)


__e.TailApply(tmp6190, tmp6192)
return


}


}, 1)

tmp6195 := Call(__e, PrimFunc(symshen_4_5semicolon_6), W2591)


__e.TailApply(tmp6189, tmp6195)
return


}, 1)

tmp6196 := Call(__e, PrimFunc(symshen_4in_1_6), W2589)


__e.TailApply(tmp6188, tmp6196)
return


}, 1)

tmp6197 := Call(__e, PrimFunc(symshen_4_5_1out), W2589)


__e.TailApply(tmp6187, tmp6197)
return


}


}, 1)

tmp6200 := Call(__e, PrimFunc(symshen_4_5integer_6), W2588)


__e.TailApply(tmp6186, tmp6200)
return


}, 1)

tmp6201 := Call(__e, PrimFunc(symshen_4in_1_6), W2587)


__e.TailApply(tmp6185, tmp6201)
return


}


}, 1)

tmp6204 := Call(__e, PrimFunc(symshen_4_5hash_6), W2586)


__e.TailApply(tmp6184, tmp6204)
return


}, 1)

tmp6205 := Call(__e, PrimFunc(symshen_4in_1_6), W2585)


__e.TailApply(tmp6183, tmp6205)
return


}


}, 1)

tmp6208 := Call(__e, PrimFunc(symshen_4_5lowC_6), V2583)


tmp6209 := Call(__e, tmp6182, tmp6208)


__e.TailApply(tmp6179, tmp6209)
return


}, 1)

tmp6210 := Call(__e, ns2_1set, symshen_4_5control_6, tmp6178)


_ = tmp6210

tmp6211 := MakeNative(func(__e *ControlFlow) {
V2594 := __e.Get(1)
_ = V2594
tmp6212 := MakeNative(func(__e *ControlFlow) {
W2595 := __e.Get(1)
_ = W2595
tmp6214 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2595)


if True == tmp6214 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2595)
return
}


}, 1)

tmp6226 := PrimIsPair(V2594)

var ifres6215 Obj

if True == tmp6226 {
tmp6216 := MakeNative(func(__e *ControlFlow) {
W2596 := __e.Get(1)
_ = W2596
tmp6217 := MakeNative(func(__e *ControlFlow) {
W2597 := __e.Get(1)
_ = W2597
tmp6220 := PrimEqual(W2596, MakeNumber(34))

tmp6221 := PrimNot(tmp6220)

if True == tmp6221 {
tmp6218 := PrimNumberToString(W2596)

__e.TailApply(PrimFunc(symshen_4comb), W2597, tmp6218)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6222 := Call(__e, PrimFunc(symtail), V2594)


__e.TailApply(tmp6217, tmp6222)
return


}, 1)

tmp6223 := Call(__e, PrimFunc(symhead), V2594)


tmp6224 := Call(__e, tmp6216, tmp6223)


ifres6215 = tmp6224


} else {
tmp6225 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6215 = tmp6225


}

__e.TailApply(tmp6212, ifres6215)
return


}, 1)

tmp6227 := Call(__e, ns2_1set, symshen_4_5notdbq_6, tmp6211)


_ = tmp6227

tmp6228 := MakeNative(func(__e *ControlFlow) {
V2598 := __e.Get(1)
_ = V2598
tmp6229 := MakeNative(func(__e *ControlFlow) {
W2599 := __e.Get(1)
_ = W2599
tmp6231 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2599)


if True == tmp6231 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2599)
return
}


}, 1)

tmp6237 := Call(__e, PrimFunc(symshen_4hds_a_2), V2598, MakeNumber(99))


var ifres6232 Obj

if True == tmp6237 {
tmp6233 := MakeNative(func(__e *ControlFlow) {
W2600 := __e.Get(1)
_ = W2600
__e.TailApply(PrimFunc(symshen_4comb), W2600, symshen_4skip)
return
}, 1)

tmp6234 := Call(__e, PrimFunc(symtail), V2598)


tmp6235 := Call(__e, tmp6233, tmp6234)


ifres6232 = tmp6235


} else {
tmp6236 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6232 = tmp6236


}

__e.TailApply(tmp6229, ifres6232)
return


}, 1)

tmp6238 := Call(__e, ns2_1set, symshen_4_5lowC_6, tmp6228)


_ = tmp6238

tmp6239 := MakeNative(func(__e *ControlFlow) {
V2601 := __e.Get(1)
_ = V2601
tmp6240 := MakeNative(func(__e *ControlFlow) {
W2602 := __e.Get(1)
_ = W2602
tmp6242 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2602)


if True == tmp6242 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2602)
return
}


}, 1)

tmp6248 := Call(__e, PrimFunc(symshen_4hds_a_2), V2601, MakeNumber(35))


var ifres6243 Obj

if True == tmp6248 {
tmp6244 := MakeNative(func(__e *ControlFlow) {
W2603 := __e.Get(1)
_ = W2603
__e.TailApply(PrimFunc(symshen_4comb), W2603, symshen_4skip)
return
}, 1)

tmp6245 := Call(__e, PrimFunc(symtail), V2601)


tmp6246 := Call(__e, tmp6244, tmp6245)


ifres6243 = tmp6246


} else {
tmp6247 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6243 = tmp6247


}

__e.TailApply(tmp6240, ifres6243)
return


}, 1)

tmp6249 := Call(__e, ns2_1set, symshen_4_5hash_6, tmp6239)


_ = tmp6249

tmp6250 := MakeNative(func(__e *ControlFlow) {
V2604 := __e.Get(1)
_ = V2604
tmp6251 := MakeNative(func(__e *ControlFlow) {
W2605 := __e.Get(1)
_ = W2605
tmp6307 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2605)


if True == tmp6307 {
tmp6252 := MakeNative(func(__e *ControlFlow) {
W2611 := __e.Get(1)
_ = W2611
tmp6290 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2611)


if True == tmp6290 {
tmp6253 := MakeNative(func(__e *ControlFlow) {
W2617 := __e.Get(1)
_ = W2617
tmp6279 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2617)


if True == tmp6279 {
tmp6254 := MakeNative(func(__e *ControlFlow) {
W2621 := __e.Get(1)
_ = W2621
tmp6268 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2621)


if True == tmp6268 {
tmp6255 := MakeNative(func(__e *ControlFlow) {
W2625 := __e.Get(1)
_ = W2625
tmp6257 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2625)


if True == tmp6257 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2625)
return
}


}, 1)

tmp6258 := MakeNative(func(__e *ControlFlow) {
W2626 := __e.Get(1)
_ = W2626
tmp6264 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2626)


if True == tmp6264 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6259 := MakeNative(func(__e *ControlFlow) {
W2627 := __e.Get(1)
_ = W2627
tmp6260 := MakeNative(func(__e *ControlFlow) {
W2628 := __e.Get(1)
_ = W2628
__e.TailApply(PrimFunc(symshen_4comb), W2628, W2627)
return
}, 1)

tmp6261 := Call(__e, PrimFunc(symshen_4in_1_6), W2626)


__e.TailApply(tmp6260, tmp6261)
return


}, 1)

tmp6262 := Call(__e, PrimFunc(symshen_4_5_1out), W2626)


__e.TailApply(tmp6259, tmp6262)
return


}


}, 1)

tmp6265 := Call(__e, PrimFunc(symshen_4_5integer_6), V2604)


tmp6266 := Call(__e, tmp6258, tmp6265)


__e.TailApply(tmp6255, tmp6266)
return


} else {
__e.Return(W2621)
return
}


}, 1)

tmp6269 := MakeNative(func(__e *ControlFlow) {
W2622 := __e.Get(1)
_ = W2622
tmp6275 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2622)


if True == tmp6275 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6270 := MakeNative(func(__e *ControlFlow) {
W2623 := __e.Get(1)
_ = W2623
tmp6271 := MakeNative(func(__e *ControlFlow) {
W2624 := __e.Get(1)
_ = W2624
__e.TailApply(PrimFunc(symshen_4comb), W2624, W2623)
return
}, 1)

tmp6272 := Call(__e, PrimFunc(symshen_4in_1_6), W2622)


__e.TailApply(tmp6271, tmp6272)
return


}, 1)

tmp6273 := Call(__e, PrimFunc(symshen_4_5_1out), W2622)


__e.TailApply(tmp6270, tmp6273)
return


}


}, 1)

tmp6276 := Call(__e, PrimFunc(symshen_4_5float_6), V2604)


tmp6277 := Call(__e, tmp6269, tmp6276)


__e.TailApply(tmp6254, tmp6277)
return


} else {
__e.Return(W2617)
return
}


}, 1)

tmp6280 := MakeNative(func(__e *ControlFlow) {
W2618 := __e.Get(1)
_ = W2618
tmp6286 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2618)


if True == tmp6286 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6281 := MakeNative(func(__e *ControlFlow) {
W2619 := __e.Get(1)
_ = W2619
tmp6282 := MakeNative(func(__e *ControlFlow) {
W2620 := __e.Get(1)
_ = W2620
__e.TailApply(PrimFunc(symshen_4comb), W2620, W2619)
return
}, 1)

tmp6283 := Call(__e, PrimFunc(symshen_4in_1_6), W2618)


__e.TailApply(tmp6282, tmp6283)
return


}, 1)

tmp6284 := Call(__e, PrimFunc(symshen_4_5_1out), W2618)


__e.TailApply(tmp6281, tmp6284)
return


}


}, 1)

tmp6287 := Call(__e, PrimFunc(symshen_4_5e_1number_6), V2604)


tmp6288 := Call(__e, tmp6280, tmp6287)


__e.TailApply(tmp6253, tmp6288)
return


} else {
__e.Return(W2611)
return
}


}, 1)

tmp6291 := MakeNative(func(__e *ControlFlow) {
W2612 := __e.Get(1)
_ = W2612
tmp6303 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2612)


if True == tmp6303 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6292 := MakeNative(func(__e *ControlFlow) {
W2613 := __e.Get(1)
_ = W2613
tmp6293 := MakeNative(func(__e *ControlFlow) {
W2614 := __e.Get(1)
_ = W2614
tmp6299 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2614)


if True == tmp6299 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6294 := MakeNative(func(__e *ControlFlow) {
W2615 := __e.Get(1)
_ = W2615
tmp6295 := MakeNative(func(__e *ControlFlow) {
W2616 := __e.Get(1)
_ = W2616
__e.TailApply(PrimFunc(symshen_4comb), W2616, W2615)
return
}, 1)

tmp6296 := Call(__e, PrimFunc(symshen_4in_1_6), W2614)


__e.TailApply(tmp6295, tmp6296)
return


}, 1)

tmp6297 := Call(__e, PrimFunc(symshen_4_5_1out), W2614)


__e.TailApply(tmp6294, tmp6297)
return


}


}, 1)

tmp6300 := Call(__e, PrimFunc(symshen_4_5number_6), W2613)


__e.TailApply(tmp6293, tmp6300)
return


}, 1)

tmp6301 := Call(__e, PrimFunc(symshen_4in_1_6), W2612)


__e.TailApply(tmp6292, tmp6301)
return


}


}, 1)

tmp6304 := Call(__e, PrimFunc(symshen_4_5plus_6), V2604)


tmp6305 := Call(__e, tmp6291, tmp6304)


__e.TailApply(tmp6252, tmp6305)
return


} else {
__e.Return(W2605)
return
}


}, 1)

tmp6308 := MakeNative(func(__e *ControlFlow) {
W2606 := __e.Get(1)
_ = W2606
tmp6321 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2606)


if True == tmp6321 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6309 := MakeNative(func(__e *ControlFlow) {
W2607 := __e.Get(1)
_ = W2607
tmp6310 := MakeNative(func(__e *ControlFlow) {
W2608 := __e.Get(1)
_ = W2608
tmp6317 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2608)


if True == tmp6317 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6311 := MakeNative(func(__e *ControlFlow) {
W2609 := __e.Get(1)
_ = W2609
tmp6312 := MakeNative(func(__e *ControlFlow) {
W2610 := __e.Get(1)
_ = W2610
tmp6313 := PrimNumberSubtract(MakeNumber(0), W2609)

__e.TailApply(PrimFunc(symshen_4comb), W2610, tmp6313)
return


}, 1)

tmp6314 := Call(__e, PrimFunc(symshen_4in_1_6), W2608)


__e.TailApply(tmp6312, tmp6314)
return


}, 1)

tmp6315 := Call(__e, PrimFunc(symshen_4_5_1out), W2608)


__e.TailApply(tmp6311, tmp6315)
return


}


}, 1)

tmp6318 := Call(__e, PrimFunc(symshen_4_5number_6), W2607)


__e.TailApply(tmp6310, tmp6318)
return


}, 1)

tmp6319 := Call(__e, PrimFunc(symshen_4in_1_6), W2606)


__e.TailApply(tmp6309, tmp6319)
return


}


}, 1)

tmp6322 := Call(__e, PrimFunc(symshen_4_5minus_6), V2604)


tmp6323 := Call(__e, tmp6308, tmp6322)


__e.TailApply(tmp6251, tmp6323)
return


}, 1)

tmp6324 := Call(__e, ns2_1set, symshen_4_5number_6, tmp6250)


_ = tmp6324

tmp6325 := MakeNative(func(__e *ControlFlow) {
V2629 := __e.Get(1)
_ = V2629
tmp6326 := MakeNative(func(__e *ControlFlow) {
W2630 := __e.Get(1)
_ = W2630
tmp6328 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2630)


if True == tmp6328 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2630)
return
}


}, 1)

tmp6334 := Call(__e, PrimFunc(symshen_4hds_a_2), V2629, MakeNumber(45))


var ifres6329 Obj

if True == tmp6334 {
tmp6330 := MakeNative(func(__e *ControlFlow) {
W2631 := __e.Get(1)
_ = W2631
__e.TailApply(PrimFunc(symshen_4comb), W2631, symshen_4skip)
return
}, 1)

tmp6331 := Call(__e, PrimFunc(symtail), V2629)


tmp6332 := Call(__e, tmp6330, tmp6331)


ifres6329 = tmp6332


} else {
tmp6333 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6329 = tmp6333


}

__e.TailApply(tmp6326, ifres6329)
return


}, 1)

tmp6335 := Call(__e, ns2_1set, symshen_4_5minus_6, tmp6325)


_ = tmp6335

tmp6336 := MakeNative(func(__e *ControlFlow) {
V2632 := __e.Get(1)
_ = V2632
tmp6337 := MakeNative(func(__e *ControlFlow) {
W2633 := __e.Get(1)
_ = W2633
tmp6339 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2633)


if True == tmp6339 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2633)
return
}


}, 1)

tmp6345 := Call(__e, PrimFunc(symshen_4hds_a_2), V2632, MakeNumber(43))


var ifres6340 Obj

if True == tmp6345 {
tmp6341 := MakeNative(func(__e *ControlFlow) {
W2634 := __e.Get(1)
_ = W2634
__e.TailApply(PrimFunc(symshen_4comb), W2634, symshen_4skip)
return
}, 1)

tmp6342 := Call(__e, PrimFunc(symtail), V2632)


tmp6343 := Call(__e, tmp6341, tmp6342)


ifres6340 = tmp6343


} else {
tmp6344 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6340 = tmp6344


}

__e.TailApply(tmp6337, ifres6340)
return


}, 1)

tmp6346 := Call(__e, ns2_1set, symshen_4_5plus_6, tmp6336)


_ = tmp6346

tmp6347 := MakeNative(func(__e *ControlFlow) {
V2635 := __e.Get(1)
_ = V2635
tmp6348 := MakeNative(func(__e *ControlFlow) {
W2636 := __e.Get(1)
_ = W2636
tmp6350 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2636)


if True == tmp6350 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2636)
return
}


}, 1)

tmp6351 := MakeNative(func(__e *ControlFlow) {
W2637 := __e.Get(1)
_ = W2637
tmp6358 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2637)


if True == tmp6358 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6352 := MakeNative(func(__e *ControlFlow) {
W2638 := __e.Get(1)
_ = W2638
tmp6353 := MakeNative(func(__e *ControlFlow) {
W2639 := __e.Get(1)
_ = W2639
tmp6354 := Call(__e, PrimFunc(symshen_4compute_1integer), W2638)


__e.TailApply(PrimFunc(symshen_4comb), W2639, tmp6354)
return


}, 1)

tmp6355 := Call(__e, PrimFunc(symshen_4in_1_6), W2637)


__e.TailApply(tmp6353, tmp6355)
return


}, 1)

tmp6356 := Call(__e, PrimFunc(symshen_4_5_1out), W2637)


__e.TailApply(tmp6352, tmp6356)
return


}


}, 1)

tmp6359 := Call(__e, PrimFunc(symshen_4_5digits_6), V2635)


tmp6360 := Call(__e, tmp6351, tmp6359)


__e.TailApply(tmp6348, tmp6360)
return


}, 1)

tmp6361 := Call(__e, ns2_1set, symshen_4_5integer_6, tmp6347)


_ = tmp6361

tmp6362 := MakeNative(func(__e *ControlFlow) {
V2640 := __e.Get(1)
_ = V2640
tmp6363 := MakeNative(func(__e *ControlFlow) {
W2641 := __e.Get(1)
_ = W2641
tmp6378 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2641)


if True == tmp6378 {
tmp6364 := MakeNative(func(__e *ControlFlow) {
W2648 := __e.Get(1)
_ = W2648
tmp6366 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2648)


if True == tmp6366 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2648)
return
}


}, 1)

tmp6367 := MakeNative(func(__e *ControlFlow) {
W2649 := __e.Get(1)
_ = W2649
tmp6374 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2649)


if True == tmp6374 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6368 := MakeNative(func(__e *ControlFlow) {
W2650 := __e.Get(1)
_ = W2650
tmp6369 := MakeNative(func(__e *ControlFlow) {
W2651 := __e.Get(1)
_ = W2651
tmp6370 := PrimCons(W2650, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W2651, tmp6370)
return


}, 1)

tmp6371 := Call(__e, PrimFunc(symshen_4in_1_6), W2649)


__e.TailApply(tmp6369, tmp6371)
return


}, 1)

tmp6372 := Call(__e, PrimFunc(symshen_4_5_1out), W2649)


__e.TailApply(tmp6368, tmp6372)
return


}


}, 1)

tmp6375 := Call(__e, PrimFunc(symshen_4_5digit_6), V2640)


tmp6376 := Call(__e, tmp6367, tmp6375)


__e.TailApply(tmp6364, tmp6376)
return


} else {
__e.Return(W2641)
return
}


}, 1)

tmp6379 := MakeNative(func(__e *ControlFlow) {
W2642 := __e.Get(1)
_ = W2642
tmp6394 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2642)


if True == tmp6394 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6380 := MakeNative(func(__e *ControlFlow) {
W2643 := __e.Get(1)
_ = W2643
tmp6381 := MakeNative(func(__e *ControlFlow) {
W2644 := __e.Get(1)
_ = W2644
tmp6382 := MakeNative(func(__e *ControlFlow) {
W2645 := __e.Get(1)
_ = W2645
tmp6389 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2645)


if True == tmp6389 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6383 := MakeNative(func(__e *ControlFlow) {
W2646 := __e.Get(1)
_ = W2646
tmp6384 := MakeNative(func(__e *ControlFlow) {
W2647 := __e.Get(1)
_ = W2647
tmp6385 := PrimCons(W2643, W2646)

__e.TailApply(PrimFunc(symshen_4comb), W2647, tmp6385)
return


}, 1)

tmp6386 := Call(__e, PrimFunc(symshen_4in_1_6), W2645)


__e.TailApply(tmp6384, tmp6386)
return


}, 1)

tmp6387 := Call(__e, PrimFunc(symshen_4_5_1out), W2645)


__e.TailApply(tmp6383, tmp6387)
return


}


}, 1)

tmp6390 := Call(__e, PrimFunc(symshen_4_5digits_6), W2644)


__e.TailApply(tmp6382, tmp6390)
return


}, 1)

tmp6391 := Call(__e, PrimFunc(symshen_4in_1_6), W2642)


__e.TailApply(tmp6381, tmp6391)
return


}, 1)

tmp6392 := Call(__e, PrimFunc(symshen_4_5_1out), W2642)


__e.TailApply(tmp6380, tmp6392)
return


}


}, 1)

tmp6395 := Call(__e, PrimFunc(symshen_4_5digit_6), V2640)


tmp6396 := Call(__e, tmp6379, tmp6395)


__e.TailApply(tmp6363, tmp6396)
return


}, 1)

tmp6397 := Call(__e, ns2_1set, symshen_4_5digits_6, tmp6362)


_ = tmp6397

tmp6398 := MakeNative(func(__e *ControlFlow) {
V2652 := __e.Get(1)
_ = V2652
tmp6399 := MakeNative(func(__e *ControlFlow) {
W2653 := __e.Get(1)
_ = W2653
tmp6401 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2653)


if True == tmp6401 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2653)
return
}


}, 1)

tmp6412 := PrimIsPair(V2652)

var ifres6402 Obj

if True == tmp6412 {
tmp6403 := MakeNative(func(__e *ControlFlow) {
W2654 := __e.Get(1)
_ = W2654
tmp6404 := MakeNative(func(__e *ControlFlow) {
W2655 := __e.Get(1)
_ = W2655
tmp6407 := Call(__e, PrimFunc(symshen_4digit_2), W2654)


if True == tmp6407 {
tmp6405 := Call(__e, PrimFunc(symshen_4byte_1_6digit), W2654)


__e.TailApply(PrimFunc(symshen_4comb), W2655, tmp6405)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6408 := Call(__e, PrimFunc(symtail), V2652)


__e.TailApply(tmp6404, tmp6408)
return


}, 1)

tmp6409 := Call(__e, PrimFunc(symhead), V2652)


tmp6410 := Call(__e, tmp6403, tmp6409)


ifres6402 = tmp6410


} else {
tmp6411 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6402 = tmp6411


}

__e.TailApply(tmp6399, ifres6402)
return


}, 1)

tmp6413 := Call(__e, ns2_1set, symshen_4_5digit_6, tmp6398)


_ = tmp6413

tmp6414 := MakeNative(func(__e *ControlFlow) {
V2656 := __e.Get(1)
_ = V2656
__e.Return(PrimNumberSubtract(V2656, MakeNumber(48)))
return
}, 1)

tmp6415 := Call(__e, ns2_1set, symshen_4byte_1_6digit, tmp6414)


_ = tmp6415

tmp6416 := MakeNative(func(__e *ControlFlow) {
V2657 := __e.Get(1)
_ = V2657
tmp6417 := Call(__e, PrimFunc(symreverse), V2657)


__e.TailApply(PrimFunc(symshen_4compute_1integer_1h), tmp6417, MakeNumber(0))
return


}, 1)

tmp6418 := Call(__e, ns2_1set, symshen_4compute_1integer, tmp6416)


_ = tmp6418

tmp6419 := MakeNative(func(__e *ControlFlow) {
V2660 := __e.Get(1)
_ = V2660
V2661 := __e.Get(2)
_ = V2661
tmp6429 := PrimEqual(Nil, V2660)

if True == tmp6429 {
__e.Return(MakeNumber(0))
return
} else {
tmp6427 := PrimIsPair(V2660)

if True == tmp6427 {
tmp6420 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V2661)


tmp6421 := PrimHead(V2660)

tmp6422 := PrimNumberMultiply(tmp6420, tmp6421)

tmp6423 := PrimTail(V2660)

tmp6424 := PrimNumberAdd(V2661, MakeNumber(1))

tmp6425 := Call(__e, PrimFunc(symshen_4compute_1integer_1h), tmp6423, tmp6424)


__e.Return(PrimNumberAdd(tmp6422, tmp6425))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.compute-integer-h")))
return
}


}


}, 2)

tmp6430 := Call(__e, ns2_1set, symshen_4compute_1integer_1h, tmp6419)


_ = tmp6430

tmp6431 := MakeNative(func(__e *ControlFlow) {
V2664 := __e.Get(1)
_ = V2664
V2665 := __e.Get(2)
_ = V2665
tmp6439 := PrimEqual(MakeNumber(0), V2665)

if True == tmp6439 {
__e.Return(MakeNumber(1))
return
} else {
tmp6437 := PrimGreatThan(V2665, MakeNumber(0))

if True == tmp6437 {
tmp6432 := PrimNumberSubtract(V2665, MakeNumber(1))

tmp6433 := Call(__e, PrimFunc(symshen_4expt), V2664, tmp6432)


__e.Return(PrimNumberMultiply(V2664, tmp6433))
return


} else {
tmp6434 := PrimNumberAdd(V2665, MakeNumber(1))

tmp6435 := Call(__e, PrimFunc(symshen_4expt), V2664, tmp6434)


__e.Return(PrimNumberDivide(tmp6435, V2664))
return


}


}


}, 2)

tmp6440 := Call(__e, ns2_1set, symshen_4expt, tmp6431)


_ = tmp6440

tmp6441 := MakeNative(func(__e *ControlFlow) {
V2666 := __e.Get(1)
_ = V2666
tmp6442 := MakeNative(func(__e *ControlFlow) {
W2667 := __e.Get(1)
_ = W2667
tmp6462 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2667)


if True == tmp6462 {
tmp6443 := MakeNative(func(__e *ControlFlow) {
W2676 := __e.Get(1)
_ = W2676
tmp6445 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2676)


if True == tmp6445 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2676)
return
}


}, 1)

tmp6446 := MakeNative(func(__e *ControlFlow) {
W2677 := __e.Get(1)
_ = W2677
tmp6458 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2677)


if True == tmp6458 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6447 := MakeNative(func(__e *ControlFlow) {
W2678 := __e.Get(1)
_ = W2678
tmp6448 := MakeNative(func(__e *ControlFlow) {
W2679 := __e.Get(1)
_ = W2679
tmp6454 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2679)


if True == tmp6454 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6449 := MakeNative(func(__e *ControlFlow) {
W2680 := __e.Get(1)
_ = W2680
tmp6450 := MakeNative(func(__e *ControlFlow) {
W2681 := __e.Get(1)
_ = W2681
__e.TailApply(PrimFunc(symshen_4comb), W2681, W2680)
return
}, 1)

tmp6451 := Call(__e, PrimFunc(symshen_4in_1_6), W2679)


__e.TailApply(tmp6450, tmp6451)
return


}, 1)

tmp6452 := Call(__e, PrimFunc(symshen_4_5_1out), W2679)


__e.TailApply(tmp6449, tmp6452)
return


}


}, 1)

tmp6455 := Call(__e, PrimFunc(symshen_4_5fraction_6), W2678)


__e.TailApply(tmp6448, tmp6455)
return


}, 1)

tmp6456 := Call(__e, PrimFunc(symshen_4in_1_6), W2677)


__e.TailApply(tmp6447, tmp6456)
return


}


}, 1)

tmp6459 := Call(__e, PrimFunc(symshen_4_5stop_6), V2666)


tmp6460 := Call(__e, tmp6446, tmp6459)


__e.TailApply(tmp6443, tmp6460)
return


} else {
__e.Return(W2667)
return
}


}, 1)

tmp6463 := MakeNative(func(__e *ControlFlow) {
W2668 := __e.Get(1)
_ = W2668
tmp6484 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2668)


if True == tmp6484 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6464 := MakeNative(func(__e *ControlFlow) {
W2669 := __e.Get(1)
_ = W2669
tmp6465 := MakeNative(func(__e *ControlFlow) {
W2670 := __e.Get(1)
_ = W2670
tmp6466 := MakeNative(func(__e *ControlFlow) {
W2671 := __e.Get(1)
_ = W2671
tmp6479 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2671)


if True == tmp6479 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6467 := MakeNative(func(__e *ControlFlow) {
W2672 := __e.Get(1)
_ = W2672
tmp6468 := MakeNative(func(__e *ControlFlow) {
W2673 := __e.Get(1)
_ = W2673
tmp6475 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2673)


if True == tmp6475 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6469 := MakeNative(func(__e *ControlFlow) {
W2674 := __e.Get(1)
_ = W2674
tmp6470 := MakeNative(func(__e *ControlFlow) {
W2675 := __e.Get(1)
_ = W2675
tmp6471 := PrimNumberAdd(W2669, W2674)

__e.TailApply(PrimFunc(symshen_4comb), W2675, tmp6471)
return


}, 1)

tmp6472 := Call(__e, PrimFunc(symshen_4in_1_6), W2673)


__e.TailApply(tmp6470, tmp6472)
return


}, 1)

tmp6473 := Call(__e, PrimFunc(symshen_4_5_1out), W2673)


__e.TailApply(tmp6469, tmp6473)
return


}


}, 1)

tmp6476 := Call(__e, PrimFunc(symshen_4_5fraction_6), W2672)


__e.TailApply(tmp6468, tmp6476)
return


}, 1)

tmp6477 := Call(__e, PrimFunc(symshen_4in_1_6), W2671)


__e.TailApply(tmp6467, tmp6477)
return


}


}, 1)

tmp6480 := Call(__e, PrimFunc(symshen_4_5stop_6), W2670)


__e.TailApply(tmp6466, tmp6480)
return


}, 1)

tmp6481 := Call(__e, PrimFunc(symshen_4in_1_6), W2668)


__e.TailApply(tmp6465, tmp6481)
return


}, 1)

tmp6482 := Call(__e, PrimFunc(symshen_4_5_1out), W2668)


__e.TailApply(tmp6464, tmp6482)
return


}


}, 1)

tmp6485 := Call(__e, PrimFunc(symshen_4_5integer_6), V2666)


tmp6486 := Call(__e, tmp6463, tmp6485)


__e.TailApply(tmp6442, tmp6486)
return


}, 1)

tmp6487 := Call(__e, ns2_1set, symshen_4_5float_6, tmp6441)


_ = tmp6487

tmp6488 := MakeNative(func(__e *ControlFlow) {
V2682 := __e.Get(1)
_ = V2682
tmp6489 := MakeNative(func(__e *ControlFlow) {
W2683 := __e.Get(1)
_ = W2683
tmp6491 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2683)


if True == tmp6491 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2683)
return
}


}, 1)

tmp6497 := Call(__e, PrimFunc(symshen_4hds_a_2), V2682, MakeNumber(46))


var ifres6492 Obj

if True == tmp6497 {
tmp6493 := MakeNative(func(__e *ControlFlow) {
W2684 := __e.Get(1)
_ = W2684
__e.TailApply(PrimFunc(symshen_4comb), W2684, symshen_4skip)
return
}, 1)

tmp6494 := Call(__e, PrimFunc(symtail), V2682)


tmp6495 := Call(__e, tmp6493, tmp6494)


ifres6492 = tmp6495


} else {
tmp6496 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6492 = tmp6496


}

__e.TailApply(tmp6489, ifres6492)
return


}, 1)

tmp6498 := Call(__e, ns2_1set, symshen_4_5stop_6, tmp6488)


_ = tmp6498

tmp6499 := MakeNative(func(__e *ControlFlow) {
V2685 := __e.Get(1)
_ = V2685
tmp6500 := MakeNative(func(__e *ControlFlow) {
W2686 := __e.Get(1)
_ = W2686
tmp6502 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2686)


if True == tmp6502 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2686)
return
}


}, 1)

tmp6503 := MakeNative(func(__e *ControlFlow) {
W2687 := __e.Get(1)
_ = W2687
tmp6510 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2687)


if True == tmp6510 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6504 := MakeNative(func(__e *ControlFlow) {
W2688 := __e.Get(1)
_ = W2688
tmp6505 := MakeNative(func(__e *ControlFlow) {
W2689 := __e.Get(1)
_ = W2689
tmp6506 := Call(__e, PrimFunc(symshen_4compute_1fraction), W2688)


__e.TailApply(PrimFunc(symshen_4comb), W2689, tmp6506)
return


}, 1)

tmp6507 := Call(__e, PrimFunc(symshen_4in_1_6), W2687)


__e.TailApply(tmp6505, tmp6507)
return


}, 1)

tmp6508 := Call(__e, PrimFunc(symshen_4_5_1out), W2687)


__e.TailApply(tmp6504, tmp6508)
return


}


}, 1)

tmp6511 := Call(__e, PrimFunc(symshen_4_5digits_6), V2685)


tmp6512 := Call(__e, tmp6503, tmp6511)


__e.TailApply(tmp6500, tmp6512)
return


}, 1)

tmp6513 := Call(__e, ns2_1set, symshen_4_5fraction_6, tmp6499)


_ = tmp6513

tmp6514 := MakeNative(func(__e *ControlFlow) {
V2690 := __e.Get(1)
_ = V2690
__e.TailApply(PrimFunc(symshen_4compute_1fraction_1h), V2690, MakeNumber(-1))
return
}, 1)

tmp6515 := Call(__e, ns2_1set, symshen_4compute_1fraction, tmp6514)


_ = tmp6515

tmp6516 := MakeNative(func(__e *ControlFlow) {
V2693 := __e.Get(1)
_ = V2693
V2694 := __e.Get(2)
_ = V2694
tmp6526 := PrimEqual(Nil, V2693)

if True == tmp6526 {
__e.Return(MakeNumber(0))
return
} else {
tmp6524 := PrimIsPair(V2693)

if True == tmp6524 {
tmp6517 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V2694)


tmp6518 := PrimHead(V2693)

tmp6519 := PrimNumberMultiply(tmp6517, tmp6518)

tmp6520 := PrimTail(V2693)

tmp6521 := PrimNumberSubtract(V2694, MakeNumber(1))

tmp6522 := Call(__e, PrimFunc(symshen_4compute_1fraction_1h), tmp6520, tmp6521)


__e.Return(PrimNumberAdd(tmp6519, tmp6522))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.compute-fraction-h")))
return
}


}


}, 2)

tmp6527 := Call(__e, ns2_1set, symshen_4compute_1fraction_1h, tmp6516)


_ = tmp6527

tmp6528 := MakeNative(func(__e *ControlFlow) {
V2695 := __e.Get(1)
_ = V2695
tmp6529 := MakeNative(func(__e *ControlFlow) {
W2696 := __e.Get(1)
_ = W2696
tmp6558 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2696)


if True == tmp6558 {
tmp6530 := MakeNative(func(__e *ControlFlow) {
W2705 := __e.Get(1)
_ = W2705
tmp6532 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2705)


if True == tmp6532 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2705)
return
}


}, 1)

tmp6533 := MakeNative(func(__e *ControlFlow) {
W2706 := __e.Get(1)
_ = W2706
tmp6554 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2706)


if True == tmp6554 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6534 := MakeNative(func(__e *ControlFlow) {
W2707 := __e.Get(1)
_ = W2707
tmp6535 := MakeNative(func(__e *ControlFlow) {
W2708 := __e.Get(1)
_ = W2708
tmp6536 := MakeNative(func(__e *ControlFlow) {
W2709 := __e.Get(1)
_ = W2709
tmp6549 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2709)


if True == tmp6549 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6537 := MakeNative(func(__e *ControlFlow) {
W2710 := __e.Get(1)
_ = W2710
tmp6538 := MakeNative(func(__e *ControlFlow) {
W2711 := __e.Get(1)
_ = W2711
tmp6545 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2711)


if True == tmp6545 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6539 := MakeNative(func(__e *ControlFlow) {
W2712 := __e.Get(1)
_ = W2712
tmp6540 := MakeNative(func(__e *ControlFlow) {
W2713 := __e.Get(1)
_ = W2713
tmp6541 := Call(__e, PrimFunc(symshen_4compute_1E), W2707, W2712)


__e.TailApply(PrimFunc(symshen_4comb), W2713, tmp6541)
return


}, 1)

tmp6542 := Call(__e, PrimFunc(symshen_4in_1_6), W2711)


__e.TailApply(tmp6540, tmp6542)
return


}, 1)

tmp6543 := Call(__e, PrimFunc(symshen_4_5_1out), W2711)


__e.TailApply(tmp6539, tmp6543)
return


}


}, 1)

tmp6546 := Call(__e, PrimFunc(symshen_4_5log10_6), W2710)


__e.TailApply(tmp6538, tmp6546)
return


}, 1)

tmp6547 := Call(__e, PrimFunc(symshen_4in_1_6), W2709)


__e.TailApply(tmp6537, tmp6547)
return


}


}, 1)

tmp6550 := Call(__e, PrimFunc(symshen_4_5lowE_6), W2708)


__e.TailApply(tmp6536, tmp6550)
return


}, 1)

tmp6551 := Call(__e, PrimFunc(symshen_4in_1_6), W2706)


__e.TailApply(tmp6535, tmp6551)
return


}, 1)

tmp6552 := Call(__e, PrimFunc(symshen_4_5_1out), W2706)


__e.TailApply(tmp6534, tmp6552)
return


}


}, 1)

tmp6555 := Call(__e, PrimFunc(symshen_4_5integer_6), V2695)


tmp6556 := Call(__e, tmp6533, tmp6555)


__e.TailApply(tmp6530, tmp6556)
return


} else {
__e.Return(W2696)
return
}


}, 1)

tmp6559 := MakeNative(func(__e *ControlFlow) {
W2697 := __e.Get(1)
_ = W2697
tmp6580 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2697)


if True == tmp6580 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6560 := MakeNative(func(__e *ControlFlow) {
W2698 := __e.Get(1)
_ = W2698
tmp6561 := MakeNative(func(__e *ControlFlow) {
W2699 := __e.Get(1)
_ = W2699
tmp6562 := MakeNative(func(__e *ControlFlow) {
W2700 := __e.Get(1)
_ = W2700
tmp6575 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2700)


if True == tmp6575 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6563 := MakeNative(func(__e *ControlFlow) {
W2701 := __e.Get(1)
_ = W2701
tmp6564 := MakeNative(func(__e *ControlFlow) {
W2702 := __e.Get(1)
_ = W2702
tmp6571 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2702)


if True == tmp6571 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6565 := MakeNative(func(__e *ControlFlow) {
W2703 := __e.Get(1)
_ = W2703
tmp6566 := MakeNative(func(__e *ControlFlow) {
W2704 := __e.Get(1)
_ = W2704
tmp6567 := Call(__e, PrimFunc(symshen_4compute_1E), W2698, W2703)


__e.TailApply(PrimFunc(symshen_4comb), W2704, tmp6567)
return


}, 1)

tmp6568 := Call(__e, PrimFunc(symshen_4in_1_6), W2702)


__e.TailApply(tmp6566, tmp6568)
return


}, 1)

tmp6569 := Call(__e, PrimFunc(symshen_4_5_1out), W2702)


__e.TailApply(tmp6565, tmp6569)
return


}


}, 1)

tmp6572 := Call(__e, PrimFunc(symshen_4_5log10_6), W2701)


__e.TailApply(tmp6564, tmp6572)
return


}, 1)

tmp6573 := Call(__e, PrimFunc(symshen_4in_1_6), W2700)


__e.TailApply(tmp6563, tmp6573)
return


}


}, 1)

tmp6576 := Call(__e, PrimFunc(symshen_4_5lowE_6), W2699)


__e.TailApply(tmp6562, tmp6576)
return


}, 1)

tmp6577 := Call(__e, PrimFunc(symshen_4in_1_6), W2697)


__e.TailApply(tmp6561, tmp6577)
return


}, 1)

tmp6578 := Call(__e, PrimFunc(symshen_4_5_1out), W2697)


__e.TailApply(tmp6560, tmp6578)
return


}


}, 1)

tmp6581 := Call(__e, PrimFunc(symshen_4_5float_6), V2695)


tmp6582 := Call(__e, tmp6559, tmp6581)


__e.TailApply(tmp6529, tmp6582)
return


}, 1)

tmp6583 := Call(__e, ns2_1set, symshen_4_5e_1number_6, tmp6528)


_ = tmp6583

tmp6584 := MakeNative(func(__e *ControlFlow) {
V2714 := __e.Get(1)
_ = V2714
tmp6585 := MakeNative(func(__e *ControlFlow) {
W2715 := __e.Get(1)
_ = W2715
tmp6618 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2715)


if True == tmp6618 {
tmp6586 := MakeNative(func(__e *ControlFlow) {
W2721 := __e.Get(1)
_ = W2721
tmp6600 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2721)


if True == tmp6600 {
tmp6587 := MakeNative(func(__e *ControlFlow) {
W2727 := __e.Get(1)
_ = W2727
tmp6589 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2727)


if True == tmp6589 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2727)
return
}


}, 1)

tmp6590 := MakeNative(func(__e *ControlFlow) {
W2728 := __e.Get(1)
_ = W2728
tmp6596 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2728)


if True == tmp6596 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6591 := MakeNative(func(__e *ControlFlow) {
W2729 := __e.Get(1)
_ = W2729
tmp6592 := MakeNative(func(__e *ControlFlow) {
W2730 := __e.Get(1)
_ = W2730
__e.TailApply(PrimFunc(symshen_4comb), W2730, W2729)
return
}, 1)

tmp6593 := Call(__e, PrimFunc(symshen_4in_1_6), W2728)


__e.TailApply(tmp6592, tmp6593)
return


}, 1)

tmp6594 := Call(__e, PrimFunc(symshen_4_5_1out), W2728)


__e.TailApply(tmp6591, tmp6594)
return


}


}, 1)

tmp6597 := Call(__e, PrimFunc(symshen_4_5integer_6), V2714)


tmp6598 := Call(__e, tmp6590, tmp6597)


__e.TailApply(tmp6587, tmp6598)
return


} else {
__e.Return(W2721)
return
}


}, 1)

tmp6601 := MakeNative(func(__e *ControlFlow) {
W2722 := __e.Get(1)
_ = W2722
tmp6614 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2722)


if True == tmp6614 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6602 := MakeNative(func(__e *ControlFlow) {
W2723 := __e.Get(1)
_ = W2723
tmp6603 := MakeNative(func(__e *ControlFlow) {
W2724 := __e.Get(1)
_ = W2724
tmp6610 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2724)


if True == tmp6610 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6604 := MakeNative(func(__e *ControlFlow) {
W2725 := __e.Get(1)
_ = W2725
tmp6605 := MakeNative(func(__e *ControlFlow) {
W2726 := __e.Get(1)
_ = W2726
tmp6606 := PrimNumberSubtract(MakeNumber(0), W2725)

__e.TailApply(PrimFunc(symshen_4comb), W2726, tmp6606)
return


}, 1)

tmp6607 := Call(__e, PrimFunc(symshen_4in_1_6), W2724)


__e.TailApply(tmp6605, tmp6607)
return


}, 1)

tmp6608 := Call(__e, PrimFunc(symshen_4_5_1out), W2724)


__e.TailApply(tmp6604, tmp6608)
return


}


}, 1)

tmp6611 := Call(__e, PrimFunc(symshen_4_5log10_6), W2723)


__e.TailApply(tmp6603, tmp6611)
return


}, 1)

tmp6612 := Call(__e, PrimFunc(symshen_4in_1_6), W2722)


__e.TailApply(tmp6602, tmp6612)
return


}


}, 1)

tmp6615 := Call(__e, PrimFunc(symshen_4_5minus_6), V2714)


tmp6616 := Call(__e, tmp6601, tmp6615)


__e.TailApply(tmp6586, tmp6616)
return


} else {
__e.Return(W2715)
return
}


}, 1)

tmp6619 := MakeNative(func(__e *ControlFlow) {
W2716 := __e.Get(1)
_ = W2716
tmp6631 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2716)


if True == tmp6631 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6620 := MakeNative(func(__e *ControlFlow) {
W2717 := __e.Get(1)
_ = W2717
tmp6621 := MakeNative(func(__e *ControlFlow) {
W2718 := __e.Get(1)
_ = W2718
tmp6627 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2718)


if True == tmp6627 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6622 := MakeNative(func(__e *ControlFlow) {
W2719 := __e.Get(1)
_ = W2719
tmp6623 := MakeNative(func(__e *ControlFlow) {
W2720 := __e.Get(1)
_ = W2720
__e.TailApply(PrimFunc(symshen_4comb), W2720, W2719)
return
}, 1)

tmp6624 := Call(__e, PrimFunc(symshen_4in_1_6), W2718)


__e.TailApply(tmp6623, tmp6624)
return


}, 1)

tmp6625 := Call(__e, PrimFunc(symshen_4_5_1out), W2718)


__e.TailApply(tmp6622, tmp6625)
return


}


}, 1)

tmp6628 := Call(__e, PrimFunc(symshen_4_5log10_6), W2717)


__e.TailApply(tmp6621, tmp6628)
return


}, 1)

tmp6629 := Call(__e, PrimFunc(symshen_4in_1_6), W2716)


__e.TailApply(tmp6620, tmp6629)
return


}


}, 1)

tmp6632 := Call(__e, PrimFunc(symshen_4_5plus_6), V2714)


tmp6633 := Call(__e, tmp6619, tmp6632)


__e.TailApply(tmp6585, tmp6633)
return


}, 1)

tmp6634 := Call(__e, ns2_1set, symshen_4_5log10_6, tmp6584)


_ = tmp6634

tmp6635 := MakeNative(func(__e *ControlFlow) {
V2731 := __e.Get(1)
_ = V2731
tmp6636 := MakeNative(func(__e *ControlFlow) {
W2732 := __e.Get(1)
_ = W2732
tmp6638 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2732)


if True == tmp6638 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2732)
return
}


}, 1)

tmp6644 := Call(__e, PrimFunc(symshen_4hds_a_2), V2731, MakeNumber(101))


var ifres6639 Obj

if True == tmp6644 {
tmp6640 := MakeNative(func(__e *ControlFlow) {
W2733 := __e.Get(1)
_ = W2733
__e.TailApply(PrimFunc(symshen_4comb), W2733, symshen_4skip)
return
}, 1)

tmp6641 := Call(__e, PrimFunc(symtail), V2731)


tmp6642 := Call(__e, tmp6640, tmp6641)


ifres6639 = tmp6642


} else {
tmp6643 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6639 = tmp6643


}

__e.TailApply(tmp6636, ifres6639)
return


}, 1)

tmp6645 := Call(__e, ns2_1set, symshen_4_5lowE_6, tmp6635)


_ = tmp6645

tmp6646 := MakeNative(func(__e *ControlFlow) {
V2734 := __e.Get(1)
_ = V2734
V2735 := __e.Get(2)
_ = V2735
tmp6647 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V2735)


__e.Return(PrimNumberMultiply(V2734, tmp6647))
return


}, 2)

tmp6648 := Call(__e, ns2_1set, symshen_4compute_1E, tmp6646)


_ = tmp6648

tmp6649 := MakeNative(func(__e *ControlFlow) {
V2736 := __e.Get(1)
_ = V2736
tmp6650 := MakeNative(func(__e *ControlFlow) {
W2737 := __e.Get(1)
_ = W2737
tmp6662 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2737)


if True == tmp6662 {
tmp6651 := MakeNative(func(__e *ControlFlow) {
W2742 := __e.Get(1)
_ = W2742
tmp6653 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2742)


if True == tmp6653 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2742)
return
}


}, 1)

tmp6654 := MakeNative(func(__e *ControlFlow) {
W2743 := __e.Get(1)
_ = W2743
tmp6658 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2743)


if True == tmp6658 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6655 := MakeNative(func(__e *ControlFlow) {
W2744 := __e.Get(1)
_ = W2744
__e.TailApply(PrimFunc(symshen_4comb), W2744, symshen_4skip)
return
}, 1)

tmp6656 := Call(__e, PrimFunc(symshen_4in_1_6), W2743)


__e.TailApply(tmp6655, tmp6656)
return


}


}, 1)

tmp6659 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V2736)


tmp6660 := Call(__e, tmp6654, tmp6659)


__e.TailApply(tmp6651, tmp6660)
return


} else {
__e.Return(W2737)
return
}


}, 1)

tmp6663 := MakeNative(func(__e *ControlFlow) {
W2738 := __e.Get(1)
_ = W2738
tmp6673 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2738)


if True == tmp6673 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6664 := MakeNative(func(__e *ControlFlow) {
W2739 := __e.Get(1)
_ = W2739
tmp6665 := MakeNative(func(__e *ControlFlow) {
W2740 := __e.Get(1)
_ = W2740
tmp6669 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2740)


if True == tmp6669 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6666 := MakeNative(func(__e *ControlFlow) {
W2741 := __e.Get(1)
_ = W2741
__e.TailApply(PrimFunc(symshen_4comb), W2741, symshen_4skip)
return
}, 1)

tmp6667 := Call(__e, PrimFunc(symshen_4in_1_6), W2740)


__e.TailApply(tmp6666, tmp6667)
return


}


}, 1)

tmp6670 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), W2739)


__e.TailApply(tmp6665, tmp6670)
return


}, 1)

tmp6671 := Call(__e, PrimFunc(symshen_4in_1_6), W2738)


__e.TailApply(tmp6664, tmp6671)
return


}


}, 1)

tmp6674 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V2736)


tmp6675 := Call(__e, tmp6663, tmp6674)


__e.TailApply(tmp6650, tmp6675)
return


}, 1)

tmp6676 := Call(__e, ns2_1set, symshen_4_5whitespaces_6, tmp6649)


_ = tmp6676

tmp6677 := MakeNative(func(__e *ControlFlow) {
V2745 := __e.Get(1)
_ = V2745
tmp6678 := MakeNative(func(__e *ControlFlow) {
W2746 := __e.Get(1)
_ = W2746
tmp6680 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2746)


if True == tmp6680 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2746)
return
}


}, 1)

tmp6690 := PrimIsPair(V2745)

var ifres6681 Obj

if True == tmp6690 {
tmp6682 := MakeNative(func(__e *ControlFlow) {
W2747 := __e.Get(1)
_ = W2747
tmp6683 := MakeNative(func(__e *ControlFlow) {
W2748 := __e.Get(1)
_ = W2748
tmp6685 := Call(__e, PrimFunc(symshen_4whitespace_2), W2747)


if True == tmp6685 {
__e.TailApply(PrimFunc(symshen_4comb), W2748, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6686 := Call(__e, PrimFunc(symtail), V2745)


__e.TailApply(tmp6683, tmp6686)
return


}, 1)

tmp6687 := Call(__e, PrimFunc(symhead), V2745)


tmp6688 := Call(__e, tmp6682, tmp6687)


ifres6681 = tmp6688


} else {
tmp6689 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6681 = tmp6689


}

__e.TailApply(tmp6678, ifres6681)
return


}, 1)

tmp6691 := Call(__e, ns2_1set, symshen_4_5whitespace_6, tmp6677)


_ = tmp6691

tmp6692 := MakeNative(func(__e *ControlFlow) {
V2751 := __e.Get(1)
_ = V2751
tmp6700 := PrimEqual(MakeNumber(32), V2751)

if True == tmp6700 {
__e.Return(True)
return
} else {
tmp6698 := PrimEqual(MakeNumber(13), V2751)

if True == tmp6698 {
__e.Return(True)
return
} else {
tmp6696 := PrimEqual(MakeNumber(10), V2751)

if True == tmp6696 {
__e.Return(True)
return
} else {
tmp6694 := PrimEqual(MakeNumber(9), V2751)

if True == tmp6694 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}


}


}, 1)

tmp6701 := Call(__e, ns2_1set, symshen_4whitespace_2, tmp6692)


_ = tmp6701

tmp6702 := MakeNative(func(__e *ControlFlow) {
V2752 := __e.Get(1)
_ = V2752
tmp6725 := PrimEqual(Nil, V2752)

if True == tmp6725 {
__e.Return(Nil)
return
} else {
tmp6723 := PrimIsPair(V2752)

var ifres6719 Obj

if True == tmp6723 {
tmp6721 := PrimHead(V2752)

tmp6722 := Call(__e, PrimFunc(symshen_4packaged_2), tmp6721)


var ifres6720 Obj

if True == tmp6722 {
ifres6720 = True


} else {
ifres6720 = False


}

ifres6719 = ifres6720


} else {
ifres6719 = False


}

if True == ifres6719 {
tmp6703 := PrimHead(V2752)

tmp6704 := Call(__e, PrimFunc(symshen_4unpackage), tmp6703)


tmp6705 := PrimTail(V2752)

tmp6706 := Call(__e, PrimFunc(symappend), tmp6704, tmp6705)


__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp6706)
return


} else {
tmp6717 := PrimIsPair(V2752)

if True == tmp6717 {
tmp6707 := MakeNative(func(__e *ControlFlow) {
W2753 := __e.Get(1)
_ = W2753
tmp6713 := Call(__e, PrimFunc(symshen_4packaged_2), W2753)


if True == tmp6713 {
tmp6708 := PrimTail(V2752)

tmp6709 := PrimCons(W2753, tmp6708)

__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp6709)
return


} else {
tmp6710 := PrimTail(V2752)

tmp6711 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), tmp6710)


__e.Return(PrimCons(W2753, tmp6711))
return


}


}, 1)

tmp6714 := PrimHead(V2752)

tmp6715 := Call(__e, PrimFunc(symmacroexpand), tmp6714)


__e.TailApply(tmp6707, tmp6715)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.unpackage&macroexpand")))
return
}


}


}


}, 1)

tmp6726 := Call(__e, ns2_1set, symshen_4unpackage_emacroexpand, tmp6702)


_ = tmp6726

tmp6727 := MakeNative(func(__e *ControlFlow) {
V2756 := __e.Get(1)
_ = V2756
tmp6742 := PrimIsPair(V2756)

var ifres6729 Obj

if True == tmp6742 {
tmp6740 := PrimHead(V2756)

tmp6741 := PrimEqual(sympackage, tmp6740)

var ifres6731 Obj

if True == tmp6741 {
tmp6738 := PrimTail(V2756)

tmp6739 := PrimIsPair(tmp6738)

var ifres6733 Obj

if True == tmp6739 {
tmp6735 := PrimTail(V2756)

tmp6736 := PrimTail(tmp6735)

tmp6737 := PrimIsPair(tmp6736)

var ifres6734 Obj

if True == tmp6737 {
ifres6734 = True


} else {
ifres6734 = False


}

ifres6733 = ifres6734


} else {
ifres6733 = False


}

var ifres6732 Obj

if True == ifres6733 {
ifres6732 = True


} else {
ifres6732 = False


}

ifres6731 = ifres6732


} else {
ifres6731 = False


}

var ifres6730 Obj

if True == ifres6731 {
ifres6730 = True


} else {
ifres6730 = False


}

ifres6729 = ifres6730


} else {
ifres6729 = False


}

if True == ifres6729 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp6743 := Call(__e, ns2_1set, symshen_4packaged_2, tmp6727)


_ = tmp6743

tmp6744 := MakeNative(func(__e *ControlFlow) {
V2759 := __e.Get(1)
_ = V2759
tmp6805 := PrimIsPair(V2759)

var ifres6787 Obj

if True == tmp6805 {
tmp6803 := PrimHead(V2759)

tmp6804 := PrimEqual(sympackage, tmp6803)

var ifres6789 Obj

if True == tmp6804 {
tmp6801 := PrimTail(V2759)

tmp6802 := PrimIsPair(tmp6801)

var ifres6791 Obj

if True == tmp6802 {
tmp6798 := PrimTail(V2759)

tmp6799 := PrimHead(tmp6798)

tmp6800 := PrimEqual(symnull, tmp6799)

var ifres6793 Obj

if True == tmp6800 {
tmp6795 := PrimTail(V2759)

tmp6796 := PrimTail(tmp6795)

tmp6797 := PrimIsPair(tmp6796)

var ifres6794 Obj

if True == tmp6797 {
ifres6794 = True


} else {
ifres6794 = False


}

ifres6793 = ifres6794


} else {
ifres6793 = False


}

var ifres6792 Obj

if True == ifres6793 {
ifres6792 = True


} else {
ifres6792 = False


}

ifres6791 = ifres6792


} else {
ifres6791 = False


}

var ifres6790 Obj

if True == ifres6791 {
ifres6790 = True


} else {
ifres6790 = False


}

ifres6789 = ifres6790


} else {
ifres6789 = False


}

var ifres6788 Obj

if True == ifres6789 {
ifres6788 = True


} else {
ifres6788 = False


}

ifres6787 = ifres6788


} else {
ifres6787 = False


}

if True == ifres6787 {
tmp6745 := PrimTail(V2759)

tmp6746 := PrimTail(tmp6745)

__e.Return(PrimTail(tmp6746))
return


} else {
tmp6785 := PrimIsPair(V2759)

var ifres6772 Obj

if True == tmp6785 {
tmp6783 := PrimHead(V2759)

tmp6784 := PrimEqual(sympackage, tmp6783)

var ifres6774 Obj

if True == tmp6784 {
tmp6781 := PrimTail(V2759)

tmp6782 := PrimIsPair(tmp6781)

var ifres6776 Obj

if True == tmp6782 {
tmp6778 := PrimTail(V2759)

tmp6779 := PrimTail(tmp6778)

tmp6780 := PrimIsPair(tmp6779)

var ifres6777 Obj

if True == tmp6780 {
ifres6777 = True


} else {
ifres6777 = False


}

ifres6776 = ifres6777


} else {
ifres6776 = False


}

var ifres6775 Obj

if True == ifres6776 {
ifres6775 = True


} else {
ifres6775 = False


}

ifres6774 = ifres6775


} else {
ifres6774 = False


}

var ifres6773 Obj

if True == ifres6774 {
ifres6773 = True


} else {
ifres6773 = False


}

ifres6772 = ifres6773


} else {
ifres6772 = False


}

if True == ifres6772 {
tmp6747 := MakeNative(func(__e *ControlFlow) {
W2760 := __e.Get(1)
_ = W2760
tmp6748 := MakeNative(func(__e *ControlFlow) {
W2761 := __e.Get(1)
_ = W2761
tmp6749 := MakeNative(func(__e *ControlFlow) {
W2762 := __e.Get(1)
_ = W2762
tmp6750 := MakeNative(func(__e *ControlFlow) {
W2763 := __e.Get(1)
_ = W2763
__e.Return(W2761)
return
}, 1)

tmp6751 := PrimTail(V2759)

tmp6752 := PrimHead(tmp6751)

tmp6753 := PrimTail(V2759)

tmp6754 := PrimTail(tmp6753)

tmp6755 := PrimTail(tmp6754)

tmp6756 := Call(__e, PrimFunc(symshen_4record_1internal), tmp6752, W2760, tmp6755)


__e.TailApply(tmp6750, tmp6756)
return


}, 1)

tmp6757 := PrimTail(V2759)

tmp6758 := PrimHead(tmp6757)

tmp6759 := Call(__e, PrimFunc(symshen_4record_1external), tmp6758, W2760)


__e.TailApply(tmp6749, tmp6759)
return


}, 1)

tmp6760 := PrimTail(V2759)

tmp6761 := PrimHead(tmp6760)

tmp6762 := PrimStr(tmp6761)

tmp6763 := PrimTail(V2759)

tmp6764 := PrimTail(tmp6763)

tmp6765 := PrimTail(tmp6764)

tmp6766 := Call(__e, PrimFunc(symshen_4package_1symbols), tmp6762, W2760, tmp6765)


__e.TailApply(tmp6748, tmp6766)
return


}, 1)

tmp6767 := PrimTail(V2759)

tmp6768 := PrimTail(tmp6767)

tmp6769 := PrimHead(tmp6768)

tmp6770 := Call(__e, PrimFunc(symeval), tmp6769)


__e.TailApply(tmp6747, tmp6770)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.unpackage")))
return
}


}


}, 1)

tmp6806 := Call(__e, ns2_1set, symshen_4unpackage, tmp6744)


_ = tmp6806

tmp6807 := MakeNative(func(__e *ControlFlow) {
V2764 := __e.Get(1)
_ = V2764
V2765 := __e.Get(2)
_ = V2765
V2766 := __e.Get(3)
_ = V2766
tmp6808 := MakeNative(func(__e *ControlFlow) {
W2767 := __e.Get(1)
_ = W2767
tmp6809 := MakeNative(func(__e *ControlFlow) {
W2769 := __e.Get(1)
_ = W2769
tmp6810 := Call(__e, PrimFunc(symunion), W2769, W2767)


tmp6811 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V2764, symshen_4internal_1symbols, tmp6810, tmp6811)
return


}, 1)

tmp6812 := PrimStr(V2764)

tmp6813 := Call(__e, PrimFunc(symshen_4internal_1symbols), tmp6812, V2765, V2766)


__e.TailApply(tmp6809, tmp6813)
return


}, 1)

tmp6814 := MakeNative(func(__e *ControlFlow) {
tmp6815 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V2764, symshen_4internal_1symbols, tmp6815)
return


}, 0)

tmp6816 := MakeNative(func(__e *ControlFlow) {
Z2768 := __e.Get(1)
_ = Z2768
__e.Return(Nil)
return
}, 1)

tmp6817 := Call(__e, try_1catch, tmp6814, tmp6816)


__e.TailApply(tmp6808, tmp6817)
return


}, 3)

tmp6818 := Call(__e, ns2_1set, symshen_4record_1internal, tmp6807)


_ = tmp6818

tmp6819 := MakeNative(func(__e *ControlFlow) {
V2776 := __e.Get(1)
_ = V2776
V2777 := __e.Get(2)
_ = V2777
V2778 := __e.Get(3)
_ = V2778
tmp6828 := PrimIsPair(V2778)

if True == tmp6828 {
tmp6820 := PrimHead(V2778)

tmp6821 := Call(__e, PrimFunc(symshen_4internal_1symbols), V2776, V2777, tmp6820)


tmp6822 := PrimTail(V2778)

tmp6823 := Call(__e, PrimFunc(symshen_4internal_1symbols), V2776, V2777, tmp6822)


__e.TailApply(PrimFunc(symunion), tmp6821, tmp6823)
return


} else {
tmp6826 := Call(__e, PrimFunc(symshen_4internal_2), V2778, V2776, V2777)


if True == tmp6826 {
tmp6824 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V2776, V2778)


__e.Return(PrimCons(tmp6824, Nil))
return


} else {
__e.Return(Nil)
return
}


}


}, 3)

tmp6829 := Call(__e, ns2_1set, symshen_4internal_1symbols, tmp6819)


_ = tmp6829

tmp6830 := MakeNative(func(__e *ControlFlow) {
V2779 := __e.Get(1)
_ = V2779
V2780 := __e.Get(2)
_ = V2780
tmp6831 := MakeNative(func(__e *ControlFlow) {
W2781 := __e.Get(1)
_ = W2781
tmp6832 := Call(__e, PrimFunc(symunion), V2780, W2781)


tmp6833 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V2779, symshen_4external_1symbols, tmp6832, tmp6833)
return


}, 1)

tmp6834 := MakeNative(func(__e *ControlFlow) {
tmp6835 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V2779, symshen_4external_1symbols, tmp6835)
return


}, 0)

tmp6836 := MakeNative(func(__e *ControlFlow) {
Z2782 := __e.Get(1)
_ = Z2782
__e.Return(Nil)
return
}, 1)

tmp6837 := Call(__e, try_1catch, tmp6834, tmp6836)


__e.TailApply(tmp6831, tmp6837)
return


}, 2)

tmp6838 := Call(__e, ns2_1set, symshen_4record_1external, tmp6830)


_ = tmp6838

tmp6839 := MakeNative(func(__e *ControlFlow) {
V2787 := __e.Get(1)
_ = V2787
V2788 := __e.Get(2)
_ = V2788
V2789 := __e.Get(3)
_ = V2789
tmp6844 := PrimIsPair(V2789)

if True == tmp6844 {
tmp6840 := MakeNative(func(__e *ControlFlow) {
Z2790 := __e.Get(1)
_ = Z2790
__e.TailApply(PrimFunc(symshen_4package_1symbols), V2787, V2788, Z2790)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp6840, V2789)
return


} else {
tmp6842 := Call(__e, PrimFunc(symshen_4internal_2), V2789, V2787, V2788)


if True == tmp6842 {
__e.TailApply(PrimFunc(symshen_4intern_1in_1package), V2787, V2789)
return
} else {
__e.Return(V2789)
return
}


}


}, 3)

tmp6845 := Call(__e, ns2_1set, symshen_4package_1symbols, tmp6839)


_ = tmp6845

tmp6846 := MakeNative(func(__e *ControlFlow) {
V2791 := __e.Get(1)
_ = V2791
V2792 := __e.Get(2)
_ = V2792
tmp6847 := PrimStr(V2792)

tmp6848 := Call(__e, PrimFunc(sym_8s), MakeString("."), tmp6847)


tmp6849 := Call(__e, PrimFunc(sym_8s), V2791, tmp6848)


__e.Return(PrimIntern(tmp6849))
return


}, 2)

tmp6850 := Call(__e, ns2_1set, symshen_4intern_1in_1package, tmp6846)


_ = tmp6850

tmp6851 := MakeNative(func(__e *ControlFlow) {
V2793 := __e.Get(1)
_ = V2793
V2794 := __e.Get(2)
_ = V2794
V2795 := __e.Get(3)
_ = V2795
tmp6881 := Call(__e, PrimFunc(symelement_2), V2793, V2795)


tmp6882 := PrimNot(tmp6881)

if True == tmp6882 {
tmp6878 := Call(__e, PrimFunc(symshen_4sng_2), V2793)


tmp6879 := PrimNot(tmp6878)

var ifres6853 Obj

if True == tmp6879 {
tmp6876 := Call(__e, PrimFunc(symshen_4dbl_2), V2793)


tmp6877 := PrimNot(tmp6876)

var ifres6855 Obj

if True == tmp6877 {
tmp6875 := PrimIsSymbol(V2793)

var ifres6857 Obj

if True == tmp6875 {
tmp6873 := Call(__e, PrimFunc(symshen_4sysfunc_2), V2793)


tmp6874 := PrimNot(tmp6873)

var ifres6859 Obj

if True == tmp6874 {
tmp6871 := PrimIsVariable(V2793)

tmp6872 := PrimNot(tmp6871)

var ifres6861 Obj

if True == tmp6872 {
tmp6868 := PrimStr(V2793)

tmp6869 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp6868)


tmp6870 := PrimNot(tmp6869)

var ifres6863 Obj

if True == tmp6870 {
tmp6865 := PrimStr(V2793)

tmp6866 := Call(__e, PrimFunc(symshen_4internal_1to_1P_2), V2794, tmp6865)


tmp6867 := PrimNot(tmp6866)

var ifres6864 Obj

if True == tmp6867 {
ifres6864 = True


} else {
ifres6864 = False


}

ifres6863 = ifres6864


} else {
ifres6863 = False


}

var ifres6862 Obj

if True == ifres6863 {
ifres6862 = True


} else {
ifres6862 = False


}

ifres6861 = ifres6862


} else {
ifres6861 = False


}

var ifres6860 Obj

if True == ifres6861 {
ifres6860 = True


} else {
ifres6860 = False


}

ifres6859 = ifres6860


} else {
ifres6859 = False


}

var ifres6858 Obj

if True == ifres6859 {
ifres6858 = True


} else {
ifres6858 = False


}

ifres6857 = ifres6858


} else {
ifres6857 = False


}

var ifres6856 Obj

if True == ifres6857 {
ifres6856 = True


} else {
ifres6856 = False


}

ifres6855 = ifres6856


} else {
ifres6855 = False


}

var ifres6854 Obj

if True == ifres6855 {
ifres6854 = True


} else {
ifres6854 = False


}

ifres6853 = ifres6854


} else {
ifres6853 = False


}

if True == ifres6853 {
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


}, 3)

tmp6883 := Call(__e, ns2_1set, symshen_4internal_2, tmp6851)


_ = tmp6883

tmp6884 := MakeNative(func(__e *ControlFlow) {
V2800 := __e.Get(1)
_ = V2800
tmp6938 := Call(__e, PrimFunc(symshen_4_7string_2), V2800)


var ifres6886 Obj

if True == tmp6938 {
tmp6936 := Call(__e, PrimFunc(symhdstr), V2800)


tmp6937 := PrimEqual(MakeString("s"), tmp6936)

var ifres6888 Obj

if True == tmp6937 {
tmp6934 := PrimTailString(V2800)

tmp6935 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6934)


var ifres6890 Obj

if True == tmp6935 {
tmp6931 := PrimTailString(V2800)

tmp6932 := Call(__e, PrimFunc(symhdstr), tmp6931)


tmp6933 := PrimEqual(MakeString("h"), tmp6932)

var ifres6892 Obj

if True == tmp6933 {
tmp6928 := PrimTailString(V2800)

tmp6929 := PrimTailString(tmp6928)

tmp6930 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6929)


var ifres6894 Obj

if True == tmp6930 {
tmp6924 := PrimTailString(V2800)

tmp6925 := PrimTailString(tmp6924)

tmp6926 := Call(__e, PrimFunc(symhdstr), tmp6925)


tmp6927 := PrimEqual(MakeString("e"), tmp6926)

var ifres6896 Obj

if True == tmp6927 {
tmp6920 := PrimTailString(V2800)

tmp6921 := PrimTailString(tmp6920)

tmp6922 := PrimTailString(tmp6921)

tmp6923 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6922)


var ifres6898 Obj

if True == tmp6923 {
tmp6915 := PrimTailString(V2800)

tmp6916 := PrimTailString(tmp6915)

tmp6917 := PrimTailString(tmp6916)

tmp6918 := Call(__e, PrimFunc(symhdstr), tmp6917)


tmp6919 := PrimEqual(MakeString("n"), tmp6918)

var ifres6900 Obj

if True == tmp6919 {
tmp6910 := PrimTailString(V2800)

tmp6911 := PrimTailString(tmp6910)

tmp6912 := PrimTailString(tmp6911)

tmp6913 := PrimTailString(tmp6912)

tmp6914 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6913)


var ifres6902 Obj

if True == tmp6914 {
tmp6904 := PrimTailString(V2800)

tmp6905 := PrimTailString(tmp6904)

tmp6906 := PrimTailString(tmp6905)

tmp6907 := PrimTailString(tmp6906)

tmp6908 := Call(__e, PrimFunc(symhdstr), tmp6907)


tmp6909 := PrimEqual(MakeString("."), tmp6908)

var ifres6903 Obj

if True == tmp6909 {
ifres6903 = True


} else {
ifres6903 = False


}

ifres6902 = ifres6903


} else {
ifres6902 = False


}

var ifres6901 Obj

if True == ifres6902 {
ifres6901 = True


} else {
ifres6901 = False


}

ifres6900 = ifres6901


} else {
ifres6900 = False


}

var ifres6899 Obj

if True == ifres6900 {
ifres6899 = True


} else {
ifres6899 = False


}

ifres6898 = ifres6899


} else {
ifres6898 = False


}

var ifres6897 Obj

if True == ifres6898 {
ifres6897 = True


} else {
ifres6897 = False


}

ifres6896 = ifres6897


} else {
ifres6896 = False


}

var ifres6895 Obj

if True == ifres6896 {
ifres6895 = True


} else {
ifres6895 = False


}

ifres6894 = ifres6895


} else {
ifres6894 = False


}

var ifres6893 Obj

if True == ifres6894 {
ifres6893 = True


} else {
ifres6893 = False


}

ifres6892 = ifres6893


} else {
ifres6892 = False


}

var ifres6891 Obj

if True == ifres6892 {
ifres6891 = True


} else {
ifres6891 = False


}

ifres6890 = ifres6891


} else {
ifres6890 = False


}

var ifres6889 Obj

if True == ifres6890 {
ifres6889 = True


} else {
ifres6889 = False


}

ifres6888 = ifres6889


} else {
ifres6888 = False


}

var ifres6887 Obj

if True == ifres6888 {
ifres6887 = True


} else {
ifres6887 = False


}

ifres6886 = ifres6887


} else {
ifres6886 = False


}

if True == ifres6886 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp6939 := Call(__e, ns2_1set, symshen_4internal_1to_1shen_2, tmp6884)


_ = tmp6939

tmp6940 := MakeNative(func(__e *ControlFlow) {
V2801 := __e.Get(1)
_ = V2801
tmp6941 := PrimValue(sym_dproperty_1vector_d)

tmp6942 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp6941)


__e.TailApply(PrimFunc(symelement_2), V2801, tmp6942)
return


}, 1)

tmp6943 := Call(__e, ns2_1set, symshen_4sysfunc_2, tmp6940)


_ = tmp6943

tmp6944 := MakeNative(func(__e *ControlFlow) {
V2809 := __e.Get(1)
_ = V2809
V2810 := __e.Get(2)
_ = V2810
tmp6965 := PrimEqual(MakeString(""), V2809)

var ifres6958 Obj

if True == tmp6965 {
tmp6964 := Call(__e, PrimFunc(symshen_4_7string_2), V2810)


var ifres6960 Obj

if True == tmp6964 {
tmp6962 := Call(__e, PrimFunc(symhdstr), V2810)


tmp6963 := PrimEqual(MakeString("."), tmp6962)

var ifres6961 Obj

if True == tmp6963 {
ifres6961 = True


} else {
ifres6961 = False


}

ifres6960 = ifres6961


} else {
ifres6960 = False


}

var ifres6959 Obj

if True == ifres6960 {
ifres6959 = True


} else {
ifres6959 = False


}

ifres6958 = ifres6959


} else {
ifres6958 = False


}

if True == ifres6958 {
__e.Return(True)
return
} else {
tmp6956 := Call(__e, PrimFunc(symshen_4_7string_2), V2809)


var ifres6948 Obj

if True == tmp6956 {
tmp6955 := Call(__e, PrimFunc(symshen_4_7string_2), V2810)


var ifres6950 Obj

if True == tmp6955 {
tmp6952 := Call(__e, PrimFunc(symhdstr), V2809)


tmp6953 := Call(__e, PrimFunc(symhdstr), V2810)


tmp6954 := PrimEqual(tmp6952, tmp6953)

var ifres6951 Obj

if True == tmp6954 {
ifres6951 = True


} else {
ifres6951 = False


}

ifres6950 = ifres6951


} else {
ifres6950 = False


}

var ifres6949 Obj

if True == ifres6950 {
ifres6949 = True


} else {
ifres6949 = False


}

ifres6948 = ifres6949


} else {
ifres6948 = False


}

if True == ifres6948 {
tmp6945 := PrimTailString(V2809)

tmp6946 := PrimTailString(V2810)

__e.TailApply(PrimFunc(symshen_4internal_1to_1P_2), tmp6945, tmp6946)
return


} else {
__e.Return(False)
return
}


}


}, 2)

tmp6966 := Call(__e, ns2_1set, symshen_4internal_1to_1P_2, tmp6944)


_ = tmp6966

tmp6967 := MakeNative(func(__e *ControlFlow) {
V2813 := __e.Get(1)
_ = V2813
V2814 := __e.Get(2)
_ = V2814
tmp6980 := Call(__e, PrimFunc(symelement_2), V2813, V2814)


if True == tmp6980 {
__e.Return(V2813)
return
} else {
tmp6978 := PrimIsPair(V2813)

var ifres6974 Obj

if True == tmp6978 {
tmp6976 := PrimHead(V2813)

tmp6977 := Call(__e, PrimFunc(symshen_4non_1application_2), tmp6976)


var ifres6975 Obj

if True == tmp6977 {
ifres6975 = True


} else {
ifres6975 = False


}

ifres6974 = ifres6975


} else {
ifres6974 = False


}

if True == ifres6974 {
tmp6968 := PrimHead(V2813)

__e.TailApply(PrimFunc(symshen_4special_1case), tmp6968, V2813, V2814)
return


} else {
tmp6972 := PrimIsPair(V2813)

if True == tmp6972 {
tmp6969 := MakeNative(func(__e *ControlFlow) {
Z2815 := __e.Get(1)
_ = Z2815
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2815, V2814)
return
}, 1)

tmp6970 := Call(__e, PrimFunc(symmap), tmp6969, V2813)


__e.TailApply(PrimFunc(symshen_4process_1application), tmp6970, V2814)
return


} else {
__e.Return(V2813)
return
}


}


}


}, 2)

tmp6981 := Call(__e, ns2_1set, symshen_4process_1applications, tmp6967)


_ = tmp6981

tmp6982 := MakeNative(func(__e *ControlFlow) {
V2818 := __e.Get(1)
_ = V2818
tmp6992 := PrimEqual(symdefine, V2818)

if True == tmp6992 {
__e.Return(True)
return
} else {
tmp6990 := PrimEqual(symdefun, V2818)

if True == tmp6990 {
__e.Return(True)
return
} else {
tmp6988 := PrimEqual(symsynonyms, V2818)

if True == tmp6988 {
__e.Return(True)
return
} else {
tmp6986 := Call(__e, PrimFunc(symshen_4special_2), V2818)


if True == tmp6986 {
__e.Return(True)
return
} else {
tmp6984 := Call(__e, PrimFunc(symshen_4extraspecial_2), V2818)


if True == tmp6984 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}


}


}


}, 1)

tmp6993 := Call(__e, ns2_1set, symshen_4non_1application_2, tmp6982)


_ = tmp6993

tmp6994 := MakeNative(func(__e *ControlFlow) {
V2823 := __e.Get(1)
_ = V2823
V2824 := __e.Get(2)
_ = V2824
V2825 := __e.Get(3)
_ = V2825
tmp7236 := PrimEqual(symlambda, V2823)

var ifres7214 Obj

if True == tmp7236 {
tmp7235 := PrimIsPair(V2824)

var ifres7216 Obj

if True == tmp7235 {
tmp7233 := PrimHead(V2824)

tmp7234 := PrimEqual(symlambda, tmp7233)

var ifres7218 Obj

if True == tmp7234 {
tmp7231 := PrimTail(V2824)

tmp7232 := PrimIsPair(tmp7231)

var ifres7220 Obj

if True == tmp7232 {
tmp7228 := PrimTail(V2824)

tmp7229 := PrimTail(tmp7228)

tmp7230 := PrimIsPair(tmp7229)

var ifres7222 Obj

if True == tmp7230 {
tmp7224 := PrimTail(V2824)

tmp7225 := PrimTail(tmp7224)

tmp7226 := PrimTail(tmp7225)

tmp7227 := PrimEqual(Nil, tmp7226)

var ifres7223 Obj

if True == tmp7227 {
ifres7223 = True


} else {
ifres7223 = False


}

ifres7222 = ifres7223


} else {
ifres7222 = False


}

var ifres7221 Obj

if True == ifres7222 {
ifres7221 = True


} else {
ifres7221 = False


}

ifres7220 = ifres7221


} else {
ifres7220 = False


}

var ifres7219 Obj

if True == ifres7220 {
ifres7219 = True


} else {
ifres7219 = False


}

ifres7218 = ifres7219


} else {
ifres7218 = False


}

var ifres7217 Obj

if True == ifres7218 {
ifres7217 = True


} else {
ifres7217 = False


}

ifres7216 = ifres7217


} else {
ifres7216 = False


}

var ifres7215 Obj

if True == ifres7216 {
ifres7215 = True


} else {
ifres7215 = False


}

ifres7214 = ifres7215


} else {
ifres7214 = False


}

if True == ifres7214 {
tmp6995 := PrimTail(V2824)

tmp6996 := PrimHead(tmp6995)

tmp6997 := PrimTail(V2824)

tmp6998 := PrimTail(tmp6997)

tmp6999 := PrimHead(tmp6998)

tmp7000 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6999, V2825)


tmp7001 := PrimCons(tmp7000, Nil)

tmp7002 := PrimCons(tmp6996, tmp7001)

__e.Return(PrimCons(symlambda, tmp7002))
return


} else {
tmp7212 := PrimEqual(symlet, V2823)

var ifres7183 Obj

if True == tmp7212 {
tmp7211 := PrimIsPair(V2824)

var ifres7185 Obj

if True == tmp7211 {
tmp7209 := PrimHead(V2824)

tmp7210 := PrimEqual(symlet, tmp7209)

var ifres7187 Obj

if True == tmp7210 {
tmp7207 := PrimTail(V2824)

tmp7208 := PrimIsPair(tmp7207)

var ifres7189 Obj

if True == tmp7208 {
tmp7204 := PrimTail(V2824)

tmp7205 := PrimTail(tmp7204)

tmp7206 := PrimIsPair(tmp7205)

var ifres7191 Obj

if True == tmp7206 {
tmp7200 := PrimTail(V2824)

tmp7201 := PrimTail(tmp7200)

tmp7202 := PrimTail(tmp7201)

tmp7203 := PrimIsPair(tmp7202)

var ifres7193 Obj

if True == tmp7203 {
tmp7195 := PrimTail(V2824)

tmp7196 := PrimTail(tmp7195)

tmp7197 := PrimTail(tmp7196)

tmp7198 := PrimTail(tmp7197)

tmp7199 := PrimEqual(Nil, tmp7198)

var ifres7194 Obj

if True == tmp7199 {
ifres7194 = True


} else {
ifres7194 = False


}

ifres7193 = ifres7194


} else {
ifres7193 = False


}

var ifres7192 Obj

if True == ifres7193 {
ifres7192 = True


} else {
ifres7192 = False


}

ifres7191 = ifres7192


} else {
ifres7191 = False


}

var ifres7190 Obj

if True == ifres7191 {
ifres7190 = True


} else {
ifres7190 = False


}

ifres7189 = ifres7190


} else {
ifres7189 = False


}

var ifres7188 Obj

if True == ifres7189 {
ifres7188 = True


} else {
ifres7188 = False


}

ifres7187 = ifres7188


} else {
ifres7187 = False


}

var ifres7186 Obj

if True == ifres7187 {
ifres7186 = True


} else {
ifres7186 = False


}

ifres7185 = ifres7186


} else {
ifres7185 = False


}

var ifres7184 Obj

if True == ifres7185 {
ifres7184 = True


} else {
ifres7184 = False


}

ifres7183 = ifres7184


} else {
ifres7183 = False


}

if True == ifres7183 {
tmp7003 := PrimTail(V2824)

tmp7004 := PrimHead(tmp7003)

tmp7005 := PrimTail(V2824)

tmp7006 := PrimTail(tmp7005)

tmp7007 := PrimHead(tmp7006)

tmp7008 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7007, V2825)


tmp7009 := PrimTail(V2824)

tmp7010 := PrimTail(tmp7009)

tmp7011 := PrimTail(tmp7010)

tmp7012 := PrimHead(tmp7011)

tmp7013 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7012, V2825)


tmp7014 := PrimCons(tmp7013, Nil)

tmp7015 := PrimCons(tmp7008, tmp7014)

tmp7016 := PrimCons(tmp7004, tmp7015)

__e.Return(PrimCons(symlet, tmp7016))
return


} else {
tmp7181 := PrimEqual(symdefun, V2823)

var ifres7152 Obj

if True == tmp7181 {
tmp7180 := PrimIsPair(V2824)

var ifres7154 Obj

if True == tmp7180 {
tmp7178 := PrimHead(V2824)

tmp7179 := PrimEqual(symdefun, tmp7178)

var ifres7156 Obj

if True == tmp7179 {
tmp7176 := PrimTail(V2824)

tmp7177 := PrimIsPair(tmp7176)

var ifres7158 Obj

if True == tmp7177 {
tmp7173 := PrimTail(V2824)

tmp7174 := PrimTail(tmp7173)

tmp7175 := PrimIsPair(tmp7174)

var ifres7160 Obj

if True == tmp7175 {
tmp7169 := PrimTail(V2824)

tmp7170 := PrimTail(tmp7169)

tmp7171 := PrimTail(tmp7170)

tmp7172 := PrimIsPair(tmp7171)

var ifres7162 Obj

if True == tmp7172 {
tmp7164 := PrimTail(V2824)

tmp7165 := PrimTail(tmp7164)

tmp7166 := PrimTail(tmp7165)

tmp7167 := PrimTail(tmp7166)

tmp7168 := PrimEqual(Nil, tmp7167)

var ifres7163 Obj

if True == tmp7168 {
ifres7163 = True


} else {
ifres7163 = False


}

ifres7162 = ifres7163


} else {
ifres7162 = False


}

var ifres7161 Obj

if True == ifres7162 {
ifres7161 = True


} else {
ifres7161 = False


}

ifres7160 = ifres7161


} else {
ifres7160 = False


}

var ifres7159 Obj

if True == ifres7160 {
ifres7159 = True


} else {
ifres7159 = False


}

ifres7158 = ifres7159


} else {
ifres7158 = False


}

var ifres7157 Obj

if True == ifres7158 {
ifres7157 = True


} else {
ifres7157 = False


}

ifres7156 = ifres7157


} else {
ifres7156 = False


}

var ifres7155 Obj

if True == ifres7156 {
ifres7155 = True


} else {
ifres7155 = False


}

ifres7154 = ifres7155


} else {
ifres7154 = False


}

var ifres7153 Obj

if True == ifres7154 {
ifres7153 = True


} else {
ifres7153 = False


}

ifres7152 = ifres7153


} else {
ifres7152 = False


}

if True == ifres7152 {
__e.Return(V2824)
return
} else {
tmp7150 := PrimEqual(symdefine, V2823)

var ifres7128 Obj

if True == tmp7150 {
tmp7149 := PrimIsPair(V2824)

var ifres7130 Obj

if True == tmp7149 {
tmp7147 := PrimHead(V2824)

tmp7148 := PrimEqual(symdefine, tmp7147)

var ifres7132 Obj

if True == tmp7148 {
tmp7145 := PrimTail(V2824)

tmp7146 := PrimIsPair(tmp7145)

var ifres7134 Obj

if True == tmp7146 {
tmp7142 := PrimTail(V2824)

tmp7143 := PrimTail(tmp7142)

tmp7144 := PrimIsPair(tmp7143)

var ifres7136 Obj

if True == tmp7144 {
tmp7138 := PrimTail(V2824)

tmp7139 := PrimTail(tmp7138)

tmp7140 := PrimHead(tmp7139)

tmp7141 := PrimEqual(sym_i, tmp7140)

var ifres7137 Obj

if True == tmp7141 {
ifres7137 = True


} else {
ifres7137 = False


}

ifres7136 = ifres7137


} else {
ifres7136 = False


}

var ifres7135 Obj

if True == ifres7136 {
ifres7135 = True


} else {
ifres7135 = False


}

ifres7134 = ifres7135


} else {
ifres7134 = False


}

var ifres7133 Obj

if True == ifres7134 {
ifres7133 = True


} else {
ifres7133 = False


}

ifres7132 = ifres7133


} else {
ifres7132 = False


}

var ifres7131 Obj

if True == ifres7132 {
ifres7131 = True


} else {
ifres7131 = False


}

ifres7130 = ifres7131


} else {
ifres7130 = False


}

var ifres7129 Obj

if True == ifres7130 {
ifres7129 = True


} else {
ifres7129 = False


}

ifres7128 = ifres7129


} else {
ifres7128 = False


}

if True == ifres7128 {
tmp7017 := PrimTail(V2824)

tmp7018 := PrimHead(tmp7017)

tmp7019 := PrimTail(V2824)

tmp7020 := PrimHead(tmp7019)

tmp7021 := PrimTail(V2824)

tmp7022 := PrimTail(tmp7021)

tmp7023 := PrimTail(tmp7022)

tmp7024 := Call(__e, PrimFunc(symshen_4process_1after_1type), tmp7020, tmp7023, V2825)


tmp7025 := PrimCons(sym_i, tmp7024)

tmp7026 := PrimCons(tmp7018, tmp7025)

__e.Return(PrimCons(symdefine, tmp7026))
return


} else {
tmp7126 := PrimEqual(symdefine, V2823)

var ifres7115 Obj

if True == tmp7126 {
tmp7125 := PrimIsPair(V2824)

var ifres7117 Obj

if True == tmp7125 {
tmp7123 := PrimHead(V2824)

tmp7124 := PrimEqual(symdefine, tmp7123)

var ifres7119 Obj

if True == tmp7124 {
tmp7121 := PrimTail(V2824)

tmp7122 := PrimIsPair(tmp7121)

var ifres7120 Obj

if True == tmp7122 {
ifres7120 = True


} else {
ifres7120 = False


}

ifres7119 = ifres7120


} else {
ifres7119 = False


}

var ifres7118 Obj

if True == ifres7119 {
ifres7118 = True


} else {
ifres7118 = False


}

ifres7117 = ifres7118


} else {
ifres7117 = False


}

var ifres7116 Obj

if True == ifres7117 {
ifres7116 = True


} else {
ifres7116 = False


}

ifres7115 = ifres7116


} else {
ifres7115 = False


}

if True == ifres7115 {
tmp7027 := PrimTail(V2824)

tmp7028 := PrimHead(tmp7027)

tmp7029 := MakeNative(func(__e *ControlFlow) {
Z2826 := __e.Get(1)
_ = Z2826
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2826, V2825)
return
}, 1)

tmp7030 := PrimTail(V2824)

tmp7031 := PrimTail(tmp7030)

tmp7032 := Call(__e, PrimFunc(symmap), tmp7029, tmp7031)


tmp7033 := PrimCons(tmp7028, tmp7032)

__e.Return(PrimCons(symdefine, tmp7033))
return


} else {
tmp7113 := PrimEqual(symsynonyms, V2823)

if True == tmp7113 {
__e.Return(PrimCons(symsynonyms, V2824))
return
} else {
tmp7111 := PrimEqual(symtype, V2823)

var ifres7089 Obj

if True == tmp7111 {
tmp7110 := PrimIsPair(V2824)

var ifres7091 Obj

if True == tmp7110 {
tmp7108 := PrimHead(V2824)

tmp7109 := PrimEqual(symtype, tmp7108)

var ifres7093 Obj

if True == tmp7109 {
tmp7106 := PrimTail(V2824)

tmp7107 := PrimIsPair(tmp7106)

var ifres7095 Obj

if True == tmp7107 {
tmp7103 := PrimTail(V2824)

tmp7104 := PrimTail(tmp7103)

tmp7105 := PrimIsPair(tmp7104)

var ifres7097 Obj

if True == tmp7105 {
tmp7099 := PrimTail(V2824)

tmp7100 := PrimTail(tmp7099)

tmp7101 := PrimTail(tmp7100)

tmp7102 := PrimEqual(Nil, tmp7101)

var ifres7098 Obj

if True == tmp7102 {
ifres7098 = True


} else {
ifres7098 = False


}

ifres7097 = ifres7098


} else {
ifres7097 = False


}

var ifres7096 Obj

if True == ifres7097 {
ifres7096 = True


} else {
ifres7096 = False


}

ifres7095 = ifres7096


} else {
ifres7095 = False


}

var ifres7094 Obj

if True == ifres7095 {
ifres7094 = True


} else {
ifres7094 = False


}

ifres7093 = ifres7094


} else {
ifres7093 = False


}

var ifres7092 Obj

if True == ifres7093 {
ifres7092 = True


} else {
ifres7092 = False


}

ifres7091 = ifres7092


} else {
ifres7091 = False


}

var ifres7090 Obj

if True == ifres7091 {
ifres7090 = True


} else {
ifres7090 = False


}

ifres7089 = ifres7090


} else {
ifres7089 = False


}

if True == ifres7089 {
tmp7034 := PrimTail(V2824)

tmp7035 := PrimHead(tmp7034)

tmp7036 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7035, V2825)


tmp7037 := PrimTail(V2824)

tmp7038 := PrimTail(tmp7037)

tmp7039 := PrimCons(tmp7036, tmp7038)

__e.Return(PrimCons(symtype, tmp7039))
return


} else {
tmp7087 := PrimEqual(syminput_7, V2823)

var ifres7065 Obj

if True == tmp7087 {
tmp7086 := PrimIsPair(V2824)

var ifres7067 Obj

if True == tmp7086 {
tmp7084 := PrimHead(V2824)

tmp7085 := PrimEqual(syminput_7, tmp7084)

var ifres7069 Obj

if True == tmp7085 {
tmp7082 := PrimTail(V2824)

tmp7083 := PrimIsPair(tmp7082)

var ifres7071 Obj

if True == tmp7083 {
tmp7079 := PrimTail(V2824)

tmp7080 := PrimTail(tmp7079)

tmp7081 := PrimIsPair(tmp7080)

var ifres7073 Obj

if True == tmp7081 {
tmp7075 := PrimTail(V2824)

tmp7076 := PrimTail(tmp7075)

tmp7077 := PrimTail(tmp7076)

tmp7078 := PrimEqual(Nil, tmp7077)

var ifres7074 Obj

if True == tmp7078 {
ifres7074 = True


} else {
ifres7074 = False


}

ifres7073 = ifres7074


} else {
ifres7073 = False


}

var ifres7072 Obj

if True == ifres7073 {
ifres7072 = True


} else {
ifres7072 = False


}

ifres7071 = ifres7072


} else {
ifres7071 = False


}

var ifres7070 Obj

if True == ifres7071 {
ifres7070 = True


} else {
ifres7070 = False


}

ifres7069 = ifres7070


} else {
ifres7069 = False


}

var ifres7068 Obj

if True == ifres7069 {
ifres7068 = True


} else {
ifres7068 = False


}

ifres7067 = ifres7068


} else {
ifres7067 = False


}

var ifres7066 Obj

if True == ifres7067 {
ifres7066 = True


} else {
ifres7066 = False


}

ifres7065 = ifres7066


} else {
ifres7065 = False


}

if True == ifres7065 {
tmp7040 := PrimTail(V2824)

tmp7041 := PrimHead(tmp7040)

tmp7042 := PrimTail(V2824)

tmp7043 := PrimTail(tmp7042)

tmp7044 := PrimHead(tmp7043)

tmp7045 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7044, V2825)


tmp7046 := PrimCons(tmp7045, Nil)

tmp7047 := PrimCons(tmp7041, tmp7046)

__e.Return(PrimCons(syminput_7, tmp7047))
return


} else {
tmp7063 := PrimIsPair(V2824)

var ifres7059 Obj

if True == tmp7063 {
tmp7061 := PrimHead(V2824)

tmp7062 := Call(__e, PrimFunc(symshen_4special_2), tmp7061)


var ifres7060 Obj

if True == tmp7062 {
ifres7060 = True


} else {
ifres7060 = False


}

ifres7059 = ifres7060


} else {
ifres7059 = False


}

if True == ifres7059 {
tmp7048 := PrimHead(V2824)

tmp7049 := MakeNative(func(__e *ControlFlow) {
Z2827 := __e.Get(1)
_ = Z2827
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2827, V2825)
return
}, 1)

tmp7050 := PrimTail(V2824)

tmp7051 := Call(__e, PrimFunc(symmap), tmp7049, tmp7050)


__e.Return(PrimCons(tmp7048, tmp7051))
return


} else {
tmp7057 := PrimIsPair(V2824)

var ifres7053 Obj

if True == tmp7057 {
tmp7055 := PrimHead(V2824)

tmp7056 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp7055)


var ifres7054 Obj

if True == tmp7056 {
ifres7054 = True


} else {
ifres7054 = False


}

ifres7053 = ifres7054


} else {
ifres7053 = False


}

if True == ifres7053 {
__e.Return(V2824)
return
} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.special-case")))
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


}, 3)

tmp7237 := Call(__e, ns2_1set, symshen_4special_1case, tmp6994)


_ = tmp7237

tmp7238 := MakeNative(func(__e *ControlFlow) {
V2830 := __e.Get(1)
_ = V2830
V2831 := __e.Get(2)
_ = V2831
V2832 := __e.Get(3)
_ = V2832
tmp7254 := PrimIsPair(V2831)

var ifres7250 Obj

if True == tmp7254 {
tmp7252 := PrimHead(V2831)

tmp7253 := PrimEqual(sym_j, tmp7252)

var ifres7251 Obj

if True == tmp7253 {
ifres7251 = True


} else {
ifres7251 = False


}

ifres7250 = ifres7251


} else {
ifres7250 = False


}

if True == ifres7250 {
tmp7239 := MakeNative(func(__e *ControlFlow) {
Z2833 := __e.Get(1)
_ = Z2833
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2833, V2832)
return
}, 1)

tmp7240 := PrimTail(V2831)

tmp7241 := Call(__e, PrimFunc(symmap), tmp7239, tmp7240)


__e.Return(PrimCons(sym_j, tmp7241))
return


} else {
tmp7248 := PrimIsPair(V2831)

if True == tmp7248 {
tmp7242 := PrimHead(V2831)

tmp7243 := PrimTail(V2831)

tmp7244 := Call(__e, PrimFunc(symshen_4process_1after_1type), V2830, tmp7243, V2832)


__e.Return(PrimCons(tmp7242, tmp7244))
return


} else {
tmp7245 := Call(__e, PrimFunc(symshen_4app), V2830, MakeString("\n"), symshen_4a)


tmp7246 := PrimStringConcat(MakeString("missing } in "), tmp7245)

__e.Return(PrimSimpleError(tmp7246))
return


}


}


}, 3)

tmp7255 := Call(__e, ns2_1set, symshen_4process_1after_1type, tmp7238)


_ = tmp7255

tmp7256 := MakeNative(func(__e *ControlFlow) {
V2834 := __e.Get(1)
_ = V2834
V2835 := __e.Get(2)
_ = V2835
tmp7301 := PrimIsPair(V2834)

if True == tmp7301 {
tmp7257 := MakeNative(func(__e *ControlFlow) {
W2836 := __e.Get(1)
_ = W2836
tmp7258 := MakeNative(func(__e *ControlFlow) {
W2837 := __e.Get(1)
_ = W2837
tmp7295 := Call(__e, PrimFunc(symelement_2), V2834, V2835)


if True == tmp7295 {
__e.Return(V2834)
return
} else {
tmp7292 := PrimHead(V2834)

tmp7293 := Call(__e, PrimFunc(symshen_4shen_1call_2), tmp7292)


if True == tmp7293 {
__e.Return(V2834)
return
} else {
tmp7290 := Call(__e, PrimFunc(symshen_4foreign_2), V2834)


if True == tmp7290 {
__e.TailApply(PrimFunc(symshen_4unpack_1foreign), V2834)
return
} else {
tmp7288 := Call(__e, PrimFunc(symshen_4fn_1call_2), V2834)


if True == tmp7288 {
__e.TailApply(PrimFunc(symshen_4fn_1call), V2834)
return
} else {
tmp7286 := Call(__e, PrimFunc(symshen_4zero_1place_2), V2834)


if True == tmp7286 {
__e.Return(V2834)
return
} else {
tmp7283 := PrimHead(V2834)

tmp7284 := Call(__e, PrimFunc(symshen_4undefined_1f_2), tmp7283, W2836)


if True == tmp7284 {
tmp7259 := PrimHead(V2834)

tmp7260 := PrimCons(tmp7259, Nil)

tmp7261 := PrimCons(symfn, tmp7260)

tmp7262 := PrimTail(V2834)

tmp7263 := PrimCons(tmp7261, tmp7262)

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7263)
return


} else {
tmp7280 := PrimHead(V2834)

tmp7281 := PrimIsVariable(tmp7280)

if True == tmp7281 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V2834)
return
} else {
tmp7277 := PrimHead(V2834)

tmp7278 := Call(__e, PrimFunc(symshen_4application_2), tmp7277)


if True == tmp7278 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V2834)
return
} else {
tmp7274 := PrimHead(V2834)

tmp7275 := Call(__e, PrimFunc(symshen_4partial_1application_d_2), tmp7274, W2836, W2837)


if True == tmp7275 {
tmp7264 := PrimNumberSubtract(W2836, W2837)

__e.TailApply(PrimFunc(symshen_4lambda_1function), V2834, tmp7264)
return


} else {
tmp7271 := PrimHead(V2834)

tmp7272 := Call(__e, PrimFunc(symshen_4overapplication_2), tmp7271, W2836, W2837)


if True == tmp7272 {
tmp7265 := PrimHead(V2834)

tmp7266 := PrimCons(tmp7265, Nil)

tmp7267 := PrimCons(symfn, tmp7266)

tmp7268 := PrimTail(V2834)

tmp7269 := PrimCons(tmp7267, tmp7268)

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7269)
return


} else {
__e.Return(V2834)
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


}, 1)

tmp7296 := PrimTail(V2834)

tmp7297 := Call(__e, PrimFunc(symlength), tmp7296)


__e.TailApply(tmp7258, tmp7297)
return


}, 1)

tmp7298 := PrimHead(V2834)

tmp7299 := Call(__e, PrimFunc(symarity), tmp7298)


__e.TailApply(tmp7257, tmp7299)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.process-application")))
return
}


}, 2)

tmp7302 := Call(__e, ns2_1set, symshen_4process_1application, tmp7256)


_ = tmp7302

tmp7303 := MakeNative(func(__e *ControlFlow) {
V2838 := __e.Get(1)
_ = V2838
tmp7329 := PrimIsPair(V2838)

var ifres7309 Obj

if True == tmp7329 {
tmp7327 := PrimHead(V2838)

tmp7328 := PrimIsPair(tmp7327)

var ifres7311 Obj

if True == tmp7328 {
tmp7324 := PrimHead(V2838)

tmp7325 := PrimHead(tmp7324)

tmp7326 := PrimEqual(symforeign, tmp7325)

var ifres7313 Obj

if True == tmp7326 {
tmp7321 := PrimHead(V2838)

tmp7322 := PrimTail(tmp7321)

tmp7323 := PrimIsPair(tmp7322)

var ifres7315 Obj

if True == tmp7323 {
tmp7317 := PrimHead(V2838)

tmp7318 := PrimTail(tmp7317)

tmp7319 := PrimTail(tmp7318)

tmp7320 := PrimEqual(Nil, tmp7319)

var ifres7316 Obj

if True == tmp7320 {
ifres7316 = True


} else {
ifres7316 = False


}

ifres7315 = ifres7316


} else {
ifres7315 = False


}

var ifres7314 Obj

if True == ifres7315 {
ifres7314 = True


} else {
ifres7314 = False


}

ifres7313 = ifres7314


} else {
ifres7313 = False


}

var ifres7312 Obj

if True == ifres7313 {
ifres7312 = True


} else {
ifres7312 = False


}

ifres7311 = ifres7312


} else {
ifres7311 = False


}

var ifres7310 Obj

if True == ifres7311 {
ifres7310 = True


} else {
ifres7310 = False


}

ifres7309 = ifres7310


} else {
ifres7309 = False


}

if True == ifres7309 {
tmp7304 := PrimHead(V2838)

tmp7305 := PrimTail(tmp7304)

tmp7306 := PrimHead(tmp7305)

tmp7307 := PrimTail(V2838)

__e.Return(PrimCons(tmp7306, tmp7307))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.unpack-foreign")))
return
}


}, 1)

tmp7330 := Call(__e, ns2_1set, symshen_4unpack_1foreign, tmp7303)


_ = tmp7330

tmp7331 := MakeNative(func(__e *ControlFlow) {
V2841 := __e.Get(1)
_ = V2841
tmp7353 := PrimIsPair(V2841)

var ifres7333 Obj

if True == tmp7353 {
tmp7351 := PrimHead(V2841)

tmp7352 := PrimIsPair(tmp7351)

var ifres7335 Obj

if True == tmp7352 {
tmp7348 := PrimHead(V2841)

tmp7349 := PrimHead(tmp7348)

tmp7350 := PrimEqual(symforeign, tmp7349)

var ifres7337 Obj

if True == tmp7350 {
tmp7345 := PrimHead(V2841)

tmp7346 := PrimTail(tmp7345)

tmp7347 := PrimIsPair(tmp7346)

var ifres7339 Obj

if True == tmp7347 {
tmp7341 := PrimHead(V2841)

tmp7342 := PrimTail(tmp7341)

tmp7343 := PrimTail(tmp7342)

tmp7344 := PrimEqual(Nil, tmp7343)

var ifres7340 Obj

if True == tmp7344 {
ifres7340 = True


} else {
ifres7340 = False


}

ifres7339 = ifres7340


} else {
ifres7339 = False


}

var ifres7338 Obj

if True == ifres7339 {
ifres7338 = True


} else {
ifres7338 = False


}

ifres7337 = ifres7338


} else {
ifres7337 = False


}

var ifres7336 Obj

if True == ifres7337 {
ifres7336 = True


} else {
ifres7336 = False


}

ifres7335 = ifres7336


} else {
ifres7335 = False


}

var ifres7334 Obj

if True == ifres7335 {
ifres7334 = True


} else {
ifres7334 = False


}

ifres7333 = ifres7334


} else {
ifres7333 = False


}

if True == ifres7333 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7354 := Call(__e, ns2_1set, symshen_4foreign_2, tmp7331)


_ = tmp7354

tmp7355 := MakeNative(func(__e *ControlFlow) {
V2844 := __e.Get(1)
_ = V2844
tmp7361 := PrimIsPair(V2844)

var ifres7357 Obj

if True == tmp7361 {
tmp7359 := PrimTail(V2844)

tmp7360 := PrimEqual(Nil, tmp7359)

var ifres7358 Obj

if True == tmp7360 {
ifres7358 = True


} else {
ifres7358 = False


}

ifres7357 = ifres7358


} else {
ifres7357 = False


}

if True == ifres7357 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7362 := Call(__e, ns2_1set, symshen_4zero_1place_2, tmp7355)


_ = tmp7362

tmp7363 := MakeNative(func(__e *ControlFlow) {
V2845 := __e.Get(1)
_ = V2845
tmp7368 := PrimIsSymbol(V2845)

if True == tmp7368 {
tmp7365 := PrimStr(V2845)

tmp7366 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp7365)


if True == tmp7366 {
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

tmp7369 := Call(__e, ns2_1set, symshen_4shen_1call_2, tmp7363)


_ = tmp7369

tmp7370 := MakeNative(func(__e *ControlFlow) {
V2850 := __e.Get(1)
_ = V2850
tmp7400 := PrimIsPair(V2850)

var ifres7387 Obj

if True == tmp7400 {
tmp7398 := PrimHead(V2850)

tmp7399 := PrimEqual(symprotect, tmp7398)

var ifres7389 Obj

if True == tmp7399 {
tmp7396 := PrimTail(V2850)

tmp7397 := PrimIsPair(tmp7396)

var ifres7391 Obj

if True == tmp7397 {
tmp7393 := PrimTail(V2850)

tmp7394 := PrimTail(tmp7393)

tmp7395 := PrimEqual(Nil, tmp7394)

var ifres7392 Obj

if True == tmp7395 {
ifres7392 = True


} else {
ifres7392 = False


}

ifres7391 = ifres7392


} else {
ifres7391 = False


}

var ifres7390 Obj

if True == ifres7391 {
ifres7390 = True


} else {
ifres7390 = False


}

ifres7389 = ifres7390


} else {
ifres7389 = False


}

var ifres7388 Obj

if True == ifres7389 {
ifres7388 = True


} else {
ifres7388 = False


}

ifres7387 = ifres7388


} else {
ifres7387 = False


}

if True == ifres7387 {
__e.Return(False)
return
} else {
tmp7385 := PrimIsPair(V2850)

var ifres7372 Obj

if True == tmp7385 {
tmp7383 := PrimHead(V2850)

tmp7384 := PrimEqual(symforeign, tmp7383)

var ifres7374 Obj

if True == tmp7384 {
tmp7381 := PrimTail(V2850)

tmp7382 := PrimIsPair(tmp7381)

var ifres7376 Obj

if True == tmp7382 {
tmp7378 := PrimTail(V2850)

tmp7379 := PrimTail(tmp7378)

tmp7380 := PrimEqual(Nil, tmp7379)

var ifres7377 Obj

if True == tmp7380 {
ifres7377 = True


} else {
ifres7377 = False


}

ifres7376 = ifres7377


} else {
ifres7376 = False


}

var ifres7375 Obj

if True == ifres7376 {
ifres7375 = True


} else {
ifres7375 = False


}

ifres7374 = ifres7375


} else {
ifres7374 = False


}

var ifres7373 Obj

if True == ifres7374 {
ifres7373 = True


} else {
ifres7373 = False


}

ifres7372 = ifres7373


} else {
ifres7372 = False


}

if True == ifres7372 {
__e.Return(False)
return
} else {
__e.Return(PrimIsPair(V2850))
return
}


}


}, 1)

tmp7401 := Call(__e, ns2_1set, symshen_4application_2, tmp7370)


_ = tmp7401

tmp7402 := MakeNative(func(__e *ControlFlow) {
V2855 := __e.Get(1)
_ = V2855
V2856 := __e.Get(2)
_ = V2856
tmp7410 := PrimEqual(MakeNumber(-1), V2856)

if True == tmp7410 {
tmp7408 := Call(__e, PrimFunc(symshen_4lowercase_1symbol_2), V2855)


if True == tmp7408 {
tmp7404 := Call(__e, PrimFunc(symexternal), symshen)


tmp7405 := Call(__e, PrimFunc(symelement_2), V2855, tmp7404)


tmp7406 := PrimNot(tmp7405)

if True == tmp7406 {
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


}, 2)

tmp7411 := Call(__e, ns2_1set, symshen_4undefined_1f_2, tmp7402)


_ = tmp7411

tmp7412 := MakeNative(func(__e *ControlFlow) {
V2857 := __e.Get(1)
_ = V2857
tmp7417 := PrimIsSymbol(V2857)

if True == tmp7417 {
tmp7414 := PrimIsVariable(V2857)

tmp7415 := PrimNot(tmp7414)

if True == tmp7415 {
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

tmp7418 := Call(__e, ns2_1set, symshen_4lowercase_1symbol_2, tmp7412)


_ = tmp7418

tmp7419 := MakeNative(func(__e *ControlFlow) {
V2858 := __e.Get(1)
_ = V2858
tmp7449 := PrimIsPair(V2858)

var ifres7440 Obj

if True == tmp7449 {
tmp7447 := PrimTail(V2858)

tmp7448 := PrimIsPair(tmp7447)

var ifres7442 Obj

if True == tmp7448 {
tmp7444 := PrimTail(V2858)

tmp7445 := PrimTail(tmp7444)

tmp7446 := PrimEqual(Nil, tmp7445)

var ifres7443 Obj

if True == tmp7446 {
ifres7443 = True


} else {
ifres7443 = False


}

ifres7442 = ifres7443


} else {
ifres7442 = False


}

var ifres7441 Obj

if True == ifres7442 {
ifres7441 = True


} else {
ifres7441 = False


}

ifres7440 = ifres7441


} else {
ifres7440 = False


}

if True == ifres7440 {
__e.Return(V2858)
return
} else {
tmp7438 := PrimIsPair(V2858)

var ifres7429 Obj

if True == tmp7438 {
tmp7436 := PrimTail(V2858)

tmp7437 := PrimIsPair(tmp7436)

var ifres7431 Obj

if True == tmp7437 {
tmp7433 := PrimTail(V2858)

tmp7434 := PrimTail(tmp7433)

tmp7435 := PrimIsPair(tmp7434)

var ifres7432 Obj

if True == tmp7435 {
ifres7432 = True


} else {
ifres7432 = False


}

ifres7431 = ifres7432


} else {
ifres7431 = False


}

var ifres7430 Obj

if True == ifres7431 {
ifres7430 = True


} else {
ifres7430 = False


}

ifres7429 = ifres7430


} else {
ifres7429 = False


}

if True == ifres7429 {
tmp7420 := PrimHead(V2858)

tmp7421 := PrimTail(V2858)

tmp7422 := PrimHead(tmp7421)

tmp7423 := PrimCons(tmp7422, Nil)

tmp7424 := PrimCons(tmp7420, tmp7423)

tmp7425 := PrimTail(V2858)

tmp7426 := PrimTail(tmp7425)

tmp7427 := PrimCons(tmp7424, tmp7426)

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7427)
return


} else {
__e.Return(V2858)
return
}


}


}, 1)

tmp7450 := Call(__e, ns2_1set, symshen_4simple_1curry, tmp7419)


_ = tmp7450

tmp7451 := MakeNative(func(__e *ControlFlow) {
V2859 := __e.Get(1)
_ = V2859
__e.TailApply(PrimFunc(symfn), V2859)
return
}, 1)

tmp7452 := Call(__e, ns2_1set, symfunction, tmp7451)


_ = tmp7452

tmp7453 := MakeNative(func(__e *ControlFlow) {
V2860 := __e.Get(1)
_ = V2860
tmp7462 := Call(__e, PrimFunc(symarity), V2860)


tmp7463 := PrimEqual(tmp7462, MakeNumber(0))

if True == tmp7463 {
__e.TailApply(V2860)
return
} else {
tmp7454 := MakeNative(func(__e *ControlFlow) {
W2861 := __e.Get(1)
_ = W2861
tmp7458 := Call(__e, PrimFunc(symempty_2), W2861)


if True == tmp7458 {
tmp7455 := Call(__e, PrimFunc(symshen_4app), V2860, MakeString(" is undefined\n"), symshen_4a)


tmp7456 := PrimStringConcat(MakeString("fn: "), tmp7455)

__e.Return(PrimSimpleError(tmp7456))
return


} else {
__e.Return(PrimTail(W2861))
return
}


}, 1)

tmp7459 := PrimValue(symshen_4_dlambdatable_d)

tmp7460 := Call(__e, PrimFunc(symassoc), V2860, tmp7459)


__e.TailApply(tmp7454, tmp7460)
return


}


}, 1)

tmp7464 := Call(__e, ns2_1set, symfn, tmp7453)


_ = tmp7464

tmp7465 := MakeNative(func(__e *ControlFlow) {
V2864 := __e.Get(1)
_ = V2864
tmp7495 := PrimIsPair(V2864)

var ifres7482 Obj

if True == tmp7495 {
tmp7493 := PrimHead(V2864)

tmp7494 := PrimEqual(symfn, tmp7493)

var ifres7484 Obj

if True == tmp7494 {
tmp7491 := PrimTail(V2864)

tmp7492 := PrimIsPair(tmp7491)

var ifres7486 Obj

if True == tmp7492 {
tmp7488 := PrimTail(V2864)

tmp7489 := PrimTail(tmp7488)

tmp7490 := PrimEqual(Nil, tmp7489)

var ifres7487 Obj

if True == tmp7490 {
ifres7487 = True


} else {
ifres7487 = False


}

ifres7486 = ifres7487


} else {
ifres7486 = False


}

var ifres7485 Obj

if True == ifres7486 {
ifres7485 = True


} else {
ifres7485 = False


}

ifres7484 = ifres7485


} else {
ifres7484 = False


}

var ifres7483 Obj

if True == ifres7484 {
ifres7483 = True


} else {
ifres7483 = False


}

ifres7482 = ifres7483


} else {
ifres7482 = False


}

if True == ifres7482 {
__e.Return(True)
return
} else {
tmp7480 := PrimIsPair(V2864)

var ifres7467 Obj

if True == tmp7480 {
tmp7478 := PrimHead(V2864)

tmp7479 := PrimEqual(symfunction, tmp7478)

var ifres7469 Obj

if True == tmp7479 {
tmp7476 := PrimTail(V2864)

tmp7477 := PrimIsPair(tmp7476)

var ifres7471 Obj

if True == tmp7477 {
tmp7473 := PrimTail(V2864)

tmp7474 := PrimTail(tmp7473)

tmp7475 := PrimEqual(Nil, tmp7474)

var ifres7472 Obj

if True == tmp7475 {
ifres7472 = True


} else {
ifres7472 = False


}

ifres7471 = ifres7472


} else {
ifres7471 = False


}

var ifres7470 Obj

if True == ifres7471 {
ifres7470 = True


} else {
ifres7470 = False


}

ifres7469 = ifres7470


} else {
ifres7469 = False


}

var ifres7468 Obj

if True == ifres7469 {
ifres7468 = True


} else {
ifres7468 = False


}

ifres7467 = ifres7468


} else {
ifres7467 = False


}

if True == ifres7467 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp7496 := Call(__e, ns2_1set, symshen_4fn_1call_2, tmp7465)


_ = tmp7496

tmp7497 := MakeNative(func(__e *ControlFlow) {
V2865 := __e.Get(1)
_ = V2865
tmp7538 := PrimIsPair(V2865)

var ifres7525 Obj

if True == tmp7538 {
tmp7536 := PrimHead(V2865)

tmp7537 := PrimEqual(symfunction, tmp7536)

var ifres7527 Obj

if True == tmp7537 {
tmp7534 := PrimTail(V2865)

tmp7535 := PrimIsPair(tmp7534)

var ifres7529 Obj

if True == tmp7535 {
tmp7531 := PrimTail(V2865)

tmp7532 := PrimTail(tmp7531)

tmp7533 := PrimEqual(Nil, tmp7532)

var ifres7530 Obj

if True == tmp7533 {
ifres7530 = True


} else {
ifres7530 = False


}

ifres7529 = ifres7530


} else {
ifres7529 = False


}

var ifres7528 Obj

if True == ifres7529 {
ifres7528 = True


} else {
ifres7528 = False


}

ifres7527 = ifres7528


} else {
ifres7527 = False


}

var ifres7526 Obj

if True == ifres7527 {
ifres7526 = True


} else {
ifres7526 = False


}

ifres7525 = ifres7526


} else {
ifres7525 = False


}

if True == ifres7525 {
tmp7498 := PrimTail(V2865)

tmp7499 := PrimCons(symfn, tmp7498)

__e.TailApply(PrimFunc(symshen_4fn_1call), tmp7499)
return


} else {
tmp7523 := PrimIsPair(V2865)

var ifres7510 Obj

if True == tmp7523 {
tmp7521 := PrimHead(V2865)

tmp7522 := PrimEqual(symfn, tmp7521)

var ifres7512 Obj

if True == tmp7522 {
tmp7519 := PrimTail(V2865)

tmp7520 := PrimIsPair(tmp7519)

var ifres7514 Obj

if True == tmp7520 {
tmp7516 := PrimTail(V2865)

tmp7517 := PrimTail(tmp7516)

tmp7518 := PrimEqual(Nil, tmp7517)

var ifres7515 Obj

if True == tmp7518 {
ifres7515 = True


} else {
ifres7515 = False


}

ifres7514 = ifres7515


} else {
ifres7514 = False


}

var ifres7513 Obj

if True == ifres7514 {
ifres7513 = True


} else {
ifres7513 = False


}

ifres7512 = ifres7513


} else {
ifres7512 = False


}

var ifres7511 Obj

if True == ifres7512 {
ifres7511 = True


} else {
ifres7511 = False


}

ifres7510 = ifres7511


} else {
ifres7510 = False


}

if True == ifres7510 {
tmp7500 := MakeNative(func(__e *ControlFlow) {
W2866 := __e.Get(1)
_ = W2866
tmp7505 := PrimEqual(W2866, MakeNumber(-1))

if True == tmp7505 {
__e.Return(V2865)
return
} else {
tmp7503 := PrimEqual(W2866, MakeNumber(0))

if True == tmp7503 {
__e.Return(PrimTail(V2865))
return
} else {
tmp7501 := PrimTail(V2865)

__e.TailApply(PrimFunc(symshen_4lambda_1function), tmp7501, W2866)
return


}


}


}, 1)

tmp7506 := PrimTail(V2865)

tmp7507 := PrimHead(tmp7506)

tmp7508 := Call(__e, PrimFunc(symarity), tmp7507)


__e.TailApply(tmp7500, tmp7508)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.fn-call")))
return
}


}


}, 1)

tmp7539 := Call(__e, ns2_1set, symshen_4fn_1call, tmp7497)


_ = tmp7539

tmp7540 := MakeNative(func(__e *ControlFlow) {
V2867 := __e.Get(1)
_ = V2867
V2868 := __e.Get(2)
_ = V2868
V2869 := __e.Get(3)
_ = V2869
tmp7541 := MakeNative(func(__e *ControlFlow) {
W2870 := __e.Get(1)
_ = W2870
tmp7542 := MakeNative(func(__e *ControlFlow) {
W2871 := __e.Get(1)
_ = W2871
__e.Return(W2870)
return
}, 1)

var ifres7548 Obj

if True == W2870 {
tmp7556 := Call(__e, PrimFunc(symshen_4loading_2))


var ifres7550 Obj

if True == tmp7556 {
tmp7552 := PrimCons(sym_1, Nil)

tmp7553 := PrimCons(sym_7, tmp7552)

tmp7554 := Call(__e, PrimFunc(symelement_2), V2867, tmp7553)


tmp7555 := PrimNot(tmp7554)

var ifres7551 Obj

if True == tmp7555 {
ifres7551 = True


} else {
ifres7551 = False


}

ifres7550 = ifres7551


} else {
ifres7550 = False


}

var ifres7549 Obj

if True == ifres7550 {
ifres7549 = True


} else {
ifres7549 = False


}

ifres7548 = ifres7549


} else {
ifres7548 = False


}

var ifres7543 Obj

if True == ifres7548 {
tmp7544 := Call(__e, PrimFunc(symshen_4app), V2867, MakeString("\n"), symshen_4a)


tmp7545 := PrimStringConcat(MakeString("partial application of "), tmp7544)

tmp7546 := Call(__e, PrimFunc(symstoutput))


tmp7547 := Call(__e, PrimFunc(sympr), tmp7545, tmp7546)


ifres7543 = tmp7547


} else {
ifres7543 = symshen_4skip


}

__e.TailApply(tmp7542, ifres7543)
return


}, 1)

tmp7557 := PrimGreatThan(V2868, V2869)

__e.TailApply(tmp7541, tmp7557)
return


}, 3)

tmp7558 := Call(__e, ns2_1set, symshen_4partial_1application_d_2, tmp7540)


_ = tmp7558

tmp7559 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dloading_2_d))
return
}, 0)

tmp7560 := Call(__e, ns2_1set, symshen_4loading_2, tmp7559)


_ = tmp7560

tmp7561 := MakeNative(func(__e *ControlFlow) {
V2876 := __e.Get(1)
_ = V2876
V2877 := __e.Get(2)
_ = V2877
V2878 := __e.Get(3)
_ = V2878
tmp7579 := PrimEqual(MakeNumber(-1), V2877)

if True == tmp7579 {
__e.Return(False)
return
} else {
tmp7562 := MakeNative(func(__e *ControlFlow) {
W2879 := __e.Get(1)
_ = W2879
tmp7563 := MakeNative(func(__e *ControlFlow) {
W2880 := __e.Get(1)
_ = W2880
__e.Return(W2879)
return
}, 1)

var ifres7574 Obj

if True == W2879 {
tmp7576 := Call(__e, PrimFunc(symshen_4loading_2))


var ifres7575 Obj

if True == tmp7576 {
ifres7575 = True


} else {
ifres7575 = False


}

ifres7574 = ifres7575


} else {
ifres7574 = False


}

var ifres7564 Obj

if True == ifres7574 {
tmp7566 := PrimEqual(V2878, MakeNumber(1))

var ifres7565 Obj

if True == tmp7566 {
ifres7565 = MakeString("")


} else {
ifres7565 = MakeString("s")


}

tmp7567 := Call(__e, PrimFunc(symshen_4app), ifres7565, MakeString("\n"), symshen_4a)


tmp7568 := PrimStringConcat(MakeString(" argument"), tmp7567)

tmp7569 := Call(__e, PrimFunc(symshen_4app), V2878, tmp7568, symshen_4a)


tmp7570 := PrimStringConcat(MakeString(" might not like "), tmp7569)

tmp7571 := Call(__e, PrimFunc(symshen_4app), V2876, tmp7570, symshen_4a)


tmp7572 := Call(__e, PrimFunc(symstoutput))


tmp7573 := Call(__e, PrimFunc(sympr), tmp7571, tmp7572)


ifres7564 = tmp7573


} else {
ifres7564 = symshen_4skip


}

__e.TailApply(tmp7563, ifres7564)
return


}, 1)

tmp7577 := PrimLessThan(V2877, V2878)

__e.TailApply(tmp7562, tmp7577)
return


}


}, 3)

__e.TailApply(ns2_1set, symshen_4overapplication_2, tmp7561)
return




}, 0)

