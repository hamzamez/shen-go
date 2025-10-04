(define make-machine
    Register-names Ops Controller-text ->
        (let Machine (make-new-machine)
             Allocated (map (/. X ((Machine allocate-register) X)) Register-names)
             (do
                ((Machine install-operations) Ops)
                ((Machine install-instruction-sequence) (assemble Controller-text Machine))
                Machine)))

\\ define register
\\ usage :
\\ (make-register a)  will create a global variable a and returns a
\\ (set-contents! a 25)   (= (value a) 25)
\\ (get-contents a)       25

(define make-register
    Name -> (let _ (set Name *unassigned*)
                 Name))

(define get-contents 
    Register -> (value Register))

(define set-contents!
    Register Value -> (set Register Value))


\\ define a stack
\\ usage :
\\ (make-stack s)  will create a global variable s [] returns s
\\ (push s 45)    [45]
\\ (push s 78)    [78 45]
\\ (pop s)        [45]
\\ (initialize s) []

(define make-stack
    Name -> (let _ (set Name [])
                 Name))

(define pop
    Stack -> (poph Stack (value Stack)))
(define poph
    Stack [] -> (error "Empty stack: POP")
    Stack [Top | S] -> (do
                         (set Stack S)
                          Top))

(define push
    Stack Value -> (set Stack [Value | (value Stack)]))

(define initialize
    Stack -> (set Stack []))


\\ make-new-machine

(define make-symbol
    \\ usage :
    \\ (make-symbol a b)  returns a-b
    \\ (make-symbol m pc) returns m-pc
    \\ will use it to generate unique global variables
    MachineName Component -> (concat MachineName (concat - Component)))

(define set!
    Symbol Value Object -> (put Symbol value Value (value Object)))
(define get!
    Symbol Object -> (get Symbol value (value Object)))

(define create-object
    Name _          -> (error "Object defined" Name) where (bound? Name) 
    Name Attributes -> (do 
                         (set Name (vector (* (length Attributes) 2)))
                         (initialize-object Name Attributes)
                         Name))
(define initialize-object
    Name [] -> Name
    Name [[Key Value] | Attributes] -> (do (put Key value Value (value Name))
                                             (initialize-object Name Attributes))
    _    _  -> (error "Something wrong with your input"))

(define make-new-machine
    Self -> (do (create-opject 
                    [[pc *unassigned*]
                     [flag *unassigned*]
                     [stack []]
                     [the-instruction-sequence []]
                     [the-ops [[initialize-stack (/. X (set! stack [] Self))]]]
                     [register-table [pc flag]]
                    ])

                (let The-ops [[initialize-stack (initialize Stack)]]
                      Register-table [[pc Pc] [flag Flag]]
                      Allocate-register (/. Reg (if (assoc Reg Register-table)
                                                    (error "Muliply define register:" Reg)
                                                    )))


\\(define allocate-register
\\    )
\\
\\(define machine-dispatch
\\    Name start -> (do (set-contents! (make-symbol Name pc) ))


(define make-instruction
    Text -> [Text | []])
(define instruction-text 
    [Inst | Insts] -> Inst)
(define instruction-execution-proc
    [Inst | Insts] -> Insts)
(define set-instruction-execution-proc!
    Inst Proc -> (let CarInst (hd (value Inst))
                      CdrInst (tl (value Inst))
                      (set Inst [CarInst | Proc])))

(define make-label-entry
    Label-name Insts -> [Label-name | Insts])

