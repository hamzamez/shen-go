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
\\ (make-register a)  will create a global variable a
\\ (set-contents! a 25)   (= (value a) 25)
\\ (get-contents a)       25

(define make-register
    Name -> (set Name *unassigned*))

(define get-contents 
    Register -> (value Register))

(define set-contents!
    Register Value -> (set Register Value))


\\ define a stack
\\ usage :
\\ (make-stack s)  will create a global variable s []
\\ (push s 45)    [45]
\\ (push s 78)    [78 45]
\\ (pop s)        [45]
\\ (initialize s) []

(define make-stack
    Name -> (set Name []))

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

