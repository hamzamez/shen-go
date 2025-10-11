(define gcd-machine ->
    (make-machine
    gcd-machine
    [a b t]
    [[rem remainder][= =]]
    [test-b [test [op =] [reg b] [const 0]] 
            [branch [label gcd-done]] 
            [assign t [op rem] [reg a] [reg b]] 
            [assign a [reg b]] [assign b [reg t]] 
            [goto [label test-b]] 
            gcd-done]))


    


(define make-machine
    Name Register-names Ops Controller-text ->
        (let Machine (make-new-machine Name)
             Allocated (allocate-registers Machine Register-names)
             (do
                (install-operations Machine Ops)
                (install-instruction-sequence Machine (assemble Controller-text Machine))
                Machine)))



(define make-new-machine
    Self -> (error "Machine already defined: ~A" Self) where (bound? Self)
    Self -> (do (set Self (vector 10))
                (put registers pc *unassigned* (value Self))
                (put registers flag *unassigned* (value Self))
                (put stack content [] (value Self))
                (put the-instruction-sequence content [] (value Self))
                (put the-ops 
                     content 
                     [[initialize-stack (/. Ignored 
                                            (put stack content [] (value Self)))]] 
                     (value Self))
                (put register-table content [pc flag] (value Self))
                Self))


(define allocate-registers
    Machine Rs -> (map (/. X (allocate-register Machine X)) Rs))

(define allocate-register
    Machine Register -> 
        (let Register-table (get register-table content (value Machine)) 
             (if (element? Register Register-table)
                         (error "Multiply defined register: ~A" Register)
                         (do (put registers Register *unassigned* (value Machine))
                             (put register-table content [Register | Register-table] (value Machine))
                              Register))))


(define install-instruction-sequence
    Machine Seq -> (put the-instruction-sequence content Seq (value Machine)))





(define get-register
    Machine Register -> (let Y (element? Register (get-register-table Machine))
                             (if Y
                                Register
                                (error "Unkown register: ~A" Register))))

(define get-register-table
    Machine -> (get register-table content (value Machine)))

(define install-operations
    Machine Ops -> (let PrevOps (operations Machine)
                        (put the-ops content (append PrevOps Ops) (value Machine))))
                        

(define stack
    Machine -> (get stack content (value Machine)))  

(define push-stack
    Machine Value -> (let Prev (stack Machine)
                          (put stack content [Value | Prev] (value Machine))))

(define pop-stack
    Machine -> (let Stack (stack Machine)
                    Top (hd Stack)
                    (if (empty? Stack)
                        (error "Empty stack: POP")
                        (do (put stack content (tl Stack) (value Machine))
                            Top))))

(define operations
    Machine -> (get the-ops content (value Machine)))


(define start
    Machine -> (let Inst-seq (get the-instruction-sequence content (value Machine))
                    R        (put registers pc Inst-seq (value Machine))
                    (execute Machine R)))

(define execute
    Machine *unassigned* -> (error "Register Pc is : *unassigned*")
    Machine []           -> done
    Machine Insts        -> "Program this it is not programmed")


(define get-register-contents
    Machine Register-name -> (get registers Register-name (value Machine)))

(define set-register-contents!
    Machine Register-name Value -> 
        (do (put registers Register-name Value (value Machine))
            done))



(define assemble
    Controller-text Machine -> 
        (let Receive (/. Insts Labels (update-insts! Insts Labels Machine))
        (extract-labels Controller-text Receive)))


(define extract-labels
    []                 Receive -> (Receive [] [])
    [Next-inst | Text] Receive -> 
        (extract-labels Text (/. Insts Labels 
                                (if (symbol? Next-inst)
                                    \\ if symbol add to lables
                                    (Receive Insts [(make-label-entry Next-inst Insts) | Labels])
                                    \\ if not symbol add to Insts (instructions)
                                    (Receive [(make-instruction Next-inst) | Insts])))))

(define update-insts!
    Insts Labels Machine -> (erro "impliment update-insts!"))

(define make-instruction
    Text -> [Text []])

(define instruction-text
    [Inst | _] -> Inst)

(define instruction-execution-proc
    [_ | Insts] -> Insts)

(define set-instruction-execution-proc!
    Inst Proc -> (error "implement set-instruction-execution-proc!"))

(define make-label-entry
    Label-name Insts -> [Label-name | Insts])

(define lookup-label
    [] Label-name -> (error "Undefined label: ASSEMBLE ~A" Label-name)
    [[Label-name | Val] | _] Label-name -> Val)

(define make-execution-procedure
    [assign  | Inst] Labels Machine Ops -> (make-assign [assign | Inst] Machine Labels Ops)
    [test    | Inst] Labels Machine Ops -> (make-test    [test    | Inst] Machine Labels Ops)
    [branch  | Inst] Labels Machine Ops -> (make-branch  [branch  | Inst] Machine Labels Ops)
    [goto    | Inst] Labels Machine Ops -> (make-goto    [goto    | Inst] Machine Labels Ops)
    [save    | Inst] Labels Machine Ops -> (make-save    [save    | Inst] Machine Labels Ops)
    [restore | Inst] Labels Machine Ops -> (make-restore [restore | Inst] Machine Labels Ops)
    [perform | Inst] Labels Machine Ops -> (make-perform [perform | Inst] Machine Labels Ops)
    Inst _ _ _ -> (error "Unknown instruction type: ASSEMBLE ~A" Inst))

(define make-assign
    [assign Target | Exp] Machine Labels Ops ->
        (/. Ignored (let Value-proc (make-assign-proc Exp Machine Labels Ops)
                     (do (set-register-contents! Machine Target )
                          (advance-pc Machine)))))


(define make-assign-proc
    [[reg Reg]]       Machine Labels Ops  -> (/. Ignored (get-register-contents Machine Reg))
    [[const Value]]   Machine Labels Ops  -> (/. Ignored Value)
    [[op | Op]| Args] Machine Labels Ops  -> (make-operation-exp [[op | Op]| Args] Machine Labels Ops)
    [[label Name]]    Machine Labels Ops  -> (/. Ignored (lookup-label Labels Name))
    Exp               _       _      _    -> (error "Uknown expression type: ASSEMBLE ~A" Exp))

(define advance-pc
    Machine -> (let PrevPc (get-register-contents Machine pc)
                    (set-register-contents! Machine pc (tl PrevPc))))



(define make-test
    [test [op | Op]| Exp] Machine Labels Ops -> 
        (let Condition-proc (make-operation-exp [[op | Op] | Exp] Machine Labels Ops)
             (/. Ignored (do (set-register-contents! Machine flag (Condition-proc))
                 (advance-pc Machine))))
    Inst _ _ _ -> (error "Bad Test instruction: ASSEMBLE ~A" Inst))

(define make-branch
    [branch [label Dest]] Machine Labels Ops -> 
        (/. Ignored (do (get-register-contents Machine flag)
                        (set-regiser-contents! Machine pc (lookup-label Labels Dest))
                        (advance-pc Machine)))
    Inst _ _ _ -> (error "Bad BRANCH instruction: ASSEMBLE" Inst))

(define make-goto 
    [goto [label Dest]] Machine Labels Ops -> 
            (/. Ignored (lookup-label Labels Dest))
    [goto [reg Register]] Machine Labels Ops ->
            (/. Ignored (let Reg (get-register-contents Machine Register)
                             (set-register-contents! Machine pc Reg)))
    Inst _ _ _ -> (error "Bad GOTO instruction: ASSEMBLE ~A" Inst))

(define make-save
    [save [reg Register]] Machine Labels Ops ->
        (/. Ignored (do (push-stack Machine Register)
                        (advance-pc Machine))))

(define make-restore
    [restore [reg Register]] Machine Labels Ops ->
        (/. Ignored (do (set-register-contents! Machine Register (pop-stack Machine))
                        (advance-pc Machine))))



(define assign-reg-name
    [[_ | X ] | _] -> X)

(define assign-value-exp
    [_ _ | Y] -> Y)

(define make-operation-exp
    Exp Machine Labels Operations
    -> (let Op (lookup-prim (operation-exp-op Exp) Operations)
            Aprocs (map (/. E (make-primitive-exp E Machine Labels)) 
                        (operation-exp-operands Exp))
            (/. Ignored (eval [Op (map (/. P (P Ignored)) Aprocs)]))))

(define make-primitive-exp
    [const Value] Machine Labels -> (/. Ignored Value)
    [label Label] Machine Labels -> (/. Ignored (lookup-label Labels Label))
    [reg   Reg  ] Machine Labels -> (/. Ignored (get-register-contents Machine Reg))
    Exp           _       _      -> (error "Uknown expression type: ASSEMBLE ~A" Exp))
    

(define operation-exp?
    [[op | _]|_] -> true
    _ -> false)

(define operation-exp-op
    [[_ | Op] | _] -> Op)

(define operation-exp-operands
    [_ | Operand] -> Operand)

(define tagged-list?
    [Tag | _] Tag -> true
    _         _   -> false)

(define lookup-prim
    Symbol [] -> (error "Unkown operation: ASSEMPLE ~A" Symbol)
    Symbol [[Symbol | Op] | _] -> Op
    Symbol [_ | Y] -> (lookup-prim Symbol Y))