(define make-register
    \\ the name is not needed it seems
    \\ but we will keep for now.
    Name -> (let Contents (vector 1)
                 Init (vector-> Contents 1 *unassigned*)
                 (dispatch-register Contents)))


(define dispatch-register
    Contents get -> (<-vector Contents 1)
    Contents set -> (/. Value (vector-> Contents 1 Value))
    _        Msg -> (error "Unknown request: REGISTER" Msg))

(define get-contents 
    Register -> (Register get))

(define set-contents!
    Register Value -> ((Register set) Value))