The current shen-go implementation by Arthur Mao does work, But it is slow.

By slow I mean it is slower than shen-scheme, on my machine shen-go runs the tests
in 6.49s while the shen-scheme port runs it in 0.89s.

The first goal is to make shen-go as fast or faster than shen-scheme.\
The second goal is make any program written in shen run as fast as the same
program written in go.

We can generalize this : any program written in Shen should be as fast of faster
then the same program written in the target language.

I will start by following the book [Structure and Interpretation of Computer Programs](https://web.mit.edu/6.001/6.037/sicp.pdf)\
part 4 : computing with register machines. Instead of using scheme I will be using Shen.

## Register machine simulator

The literal traslation of make register functions to shen 
```shen
(define make-register
    \\ the name is not needed it seems
    \\ but we will keep for now.
    Name -> (dispatch-register *unassigned*))


(define dispatch-register
    Contents get -> (value Contents)
    Contents set -> (/. Value (set Contents Value))
    _        Msg -> (error "Unknown request: REGISTER" Msg))

(define get-contents 
    Register -> (Register get))

(define set-contents!
    Register Value -> ((Register set) Value))
```

But because Shen is not lexically scoped, (set var value) is a global operation,
and var will be accessible from anywhere in the code.

```shen
(define make-register
    Name -> (set Name *unassigned*))

(define get-contents 
    Register -> (value Register))

(define set-contents!
    Register Value -> (set Register Value))
```

we use it as 
```shen
(make-register a)
(set-contents! a 25)
(get-contents a) \\ output 25
```

Mark Tarver pointed out that we can use vector to acheive the same result

```shen
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
```

This will work just like the scheme version
```shen
(let R (make-register r)
     Content (R get) \\ *unassigned*
     V  ((R set) 25) \\ <25>
     C  (R get)  \\ 25
     )
```

### Making a machine
To make a machine that calculates gcd, we call make-machine with 4 arguments
- name of the machine gcd-machine: global variable pointing to the machine model created
- list of registers needed by our machine  [a b t]
- list of primitive operations [[rem remainder][= =]] , rem : to be used in machine code , remainder : function on scheme (I think we will change that to mod in shen)
- list of instructions that will run on our simulator.

```shen
(define gcd-machine ->
    (make-machine
        gcd-machine
        [a b t]
        [[rem remainder][= =]]
        [test-b 
            [test [op =] [reg b] [const 0]] 
            [branch [label gcd-done]] 
            [assign t [op rem] [reg a] [reg b]] 
            [assign a [reg b]] 
            [assign b [reg t]] 
            [goto [label test-b]] 
         gcd-done]))
```


