# Interface and Dependency injection
- testable and flexible
- separate tasks (decoupling)
- 1. define a contract (interface)
    : decribe generaic behaviors
    : dont know about specific OBJs to deal with
- 2. create Concrete worker (Struct)
    : to fullfil the contract from step 1!!
    : know specific details (how to talk to openweather API)
- 3. Use contract (Dependency injection)
    : change 'weatherHandler() to ONLY know generaic contract
    : Inject (pass in) specific worker from s2

# Struct vs Interface

# defer is on STACK LIFO!!

# Pointers
    zero => nil
# pointer to struct
v := myStruc{1, 2};
p := &v;
    => '&' give the address of 'v'
    =>  'p' is a PTR to 'v'. it is of '*myStruc' type!
    => fmt.Printf("%T", p) -> '*main.mtStruc'!!!
NOTE:!!!
    IN Go, to access the fields of the struct, '*' is not needed!!!
    p.var === (*p).var

# Struct literal
    : a way to create a value of a struct type