(define make-machine
    Register-names Ops Controller-text ->
        (let Machine (make-new-machine)
             Allocated (map (/. X ((Machine allocate-register) X)) Register-names)
             (do
                ((Machine install-operations) Ops)
                ((Machine install-instruction-sequence) (assemble Controller-text Machine))
                Machine)))

(define make-register
    Name -> (set Name *unassigned*))

(define get-contents 
    Register -> (value Register))

(define set-contents!
    Register Value -> (set Register Value))