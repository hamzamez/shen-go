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

