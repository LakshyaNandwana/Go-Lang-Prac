## INTERFACES IN GOLang


#### Points to remember for interfaces in golang.

- Interfaces can bee a set of methods or set of types 
- They can even both (not very common though)
- Calling a method via interface is slower than calling the method via the value itself(approximately 4 times).
- If the requirement says that if the underlying type under the io writer is supporting sync, call the sync immediately to see the results.
- The bigger the interface, the weaker teh abstraction.
- Interfaces tells us what we need, not what we provide.




#### Advantages

- Interfaces allows us to separate mechanism from behavior.
- Interfaces allows us to be flexible with the types
- In Test cases, interfaces allows us to mock things/functions.
- 




#### Uses

- Keep interfaces small using type assertions.
- Rule of thumb: Accept interfaces, return types.