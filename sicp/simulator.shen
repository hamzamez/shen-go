(define stringify
    []      -> "()"
    [skip]  -> ")"
    X       -> (str X) where (not (cons? X))
    [skip X | Y] -> (@s  (stringify X) " " (stringify [skip | Y]) ) 
    [X | Y] -> (@s "("  (stringify  X) " " (stringify [skip | Y]) ))

(defmacro quote
    [test X Y Z] -> [test X Y (stringify Z)])

(define test
    X Y Z -> [X Y Z])

(define gcd-machine ->
    (make-machine
    gcd-machine
    [a b c]
    [[rem remainder][= =]]
    (test-b (test (op =) (reg b) (const 0))
             (branch (label gcd-done))
             (assign t (op rem) (reg a) (reg b))
             (assign a (reg b))
             (assign b (reg t))
             (goto (label test-b))
             gcd-done)))


    


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
                (put the-ops content [[initialize-stack initialize]] (value Self))
                (put register-table content [pc flag] (value Self))
                Self))


(define install-instruction-sequence
    Machine Seq -> (put the-instruction-sequence content Seq (value Machine)))


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


(define get-register
    Machine Register -> (let Y (element? Register (get-register-table Machine))
                             (if Y
                                Register
                                (error "Unkown register: ~A" Register))))

(define get-register-table
    Machine -> (get register-table content (value Machine)))

(define install-operations
    Machine Ops -> (let PrevOps (get the-ops content (value Machine))
                        (put the-ops content (append PrevOps Ops) (value Machine))))
                        

(define stack
    Machine -> (get stack content (value Machine)))  

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
    assign Labels Machine Pc Flag Stack Ops -> (error "Implement make-execution-procedure"))

(define make-assign
    Inst Machine Labels Operations Pc -> 
    (let Target (get-register Machine (assign-reg-name Inst))
         Value-exp (assign-value-exp Inst)
         Value-proc (if (operation-exp? Value-exp)
                        (make-operation-exp Value-exp Machine Labels Operations)
                        (make-primitive-exp (hd Value-exp) Machine Labels))
         (/. Ignored (do 
                        (set-contents! Target (Value-proc Ignored))
                        (advance-pc Pc)))))

(define assign-reg-name
    [[_ | X ] | _] -> X)

(define  assign-value-exp
    [_ _ | Y] -> Y)

(define make-operation-exp
    Exp Machine Labels Operations
    -> (let Op (lookup-prim (operation-exp-op Exp) Operations)
            Aprocs (map (/. E (make-primitive-exp E Machine Labels)) 
                        (operation-exp-operands Exp))
            (/. Ignored (eval [Op (map (/. P (P Ignored)) Aprocs)]))))

(define make-primitive-exp
    E X Y -> (error "implement make-primitive-exp"))

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