\\ our machine starts by allocating registers
\\ allocating a register is just adding a pair [<register symbol> <register value>]
\\ to a list register-table  [[pc *unassigned*][flag *unassigned*]]
\\ *unassigned* is the default value when allocating a register for the first time.
\\ Usage :
\\ (allocate-registers [pc flag s m n])
\\ [[n *unassigned*] [m *unassigned*] [s *unassigned*] [flag *unassigned*] [pc *unassigned*]]

(define allocate-registers
    \\ take a list of register names and return a list of [registername value] pairs
    Registers -> (allocate-registers-loop Registers []))

(define allocate-registers-loop
    [] Allocated -> Allocated
    [R | Registers] Allocated -> 
        (allocate-registers-loop Registers (allocate-register R Allocated)))

(define allocate-register
    Name Table -> (can-allocate-register? Name (assoc Name Table) Table))

(define can-allocate-register?
    Name [] Table -> [[Name *unassigned*] | Table]
    Name [Name | Value] Table -> (error "Multiply defined register: ~A" Name))

\\ after allocating the register, our machine installs operations
\\ which is just appending a list of ops [[rem remainder]  [= =]]
\\ to the ops of the machine which has only one op initialize-stack
\\ we can check if the ops are duplicate like registers or if the primitives
\\ are not defined but we will leave that for later.
(define install-operations
    Ops -> [[initialize-stack (/. Stack [])] | Ops])


