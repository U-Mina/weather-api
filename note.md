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

# array (compare with c)
    : a array va denote ENTIRE array, NOT a PTR to the 1st element (in C)
    : think array in Go as: indexed struct
# array literal
    : b := [2]string{"lucky", "ducky"}
    : or, compiler can count => b := [...]int{3, 2, 1, 5}

# slice => build on array
# slice literal
    : b := []string{"a", "n", "c"}
    : using 'make' function
    -> func make([]T, len ,cap) []T
        : parameter explain: take a type ([]T), a length (len), a optional capacity (cap)
            when called, 'make' allocates an ARRAY amd return a SLICE that refers to that array
    ```
    var s []byte
    s = make([]byte, 5) ===> s == []byte{0, 0, 0, 0, 0}
    (when 'cap' is omitted, it defaults to length!)
    ```
# form slice from array
    b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
    // b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
    // b[:2] == []byte{'g', 'o'}
    // b[2:] == []byte{'l', 'a', 'n', 'g'}
    // b[:] == b
# expand slice
    : func append(s []T, x ...T) []T =>> append elements x to the end of slice s
        a := make([]int, 1)
        // a == []int{0}
        a = append(a, 1, 2, 3)
        // a == []int{0, 1, 2, 3}
    : Append one slice to another, use '...'
        a := []string{"John", "Paul"}
        b := []string{"George", "Ringo", "Pete"}
        a = append(a, b...)  ==>>> equivalent to "append(a, b[0], b[1], b[2])"
        // a == []string{"John", "Paul", "George", "Ringo", "Pete"}
# slice as a Reference to arrays
    : slice Dose NOT save data, just describe a section of an underlying array
    : change elements of a slice modifies the ARRAY!!!
    : OTHER slice SHARE the same array will Change!
    : ==>
        names := [4]string{"hello", "thursday", "just", "pass"}
        a := names[:1]
        b := names[0:]
# slice can have ANY type, incl. other slice
# slice Range
    : '_' to skip index or value