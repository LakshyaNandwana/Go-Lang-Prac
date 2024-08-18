# DEPENDENCY INJECTIONS


### Types

- Constructor Injection
- Prroperty Injection
- Method(Setter) Injection




#### Constructor Injection

The most common kind is the Constructor Injection. It allows you to make your implementation immutable, nothing can change the dependencies (if your properties are private). Also, it requires all dependencies to be ready to create something. If they aren’t, it usually will generate an error.


#### Property And Method Injection

Property and Method injection are pretty similar. Their adoption is a question of a language feature. In C# is more common to have Property Injection. In Go, we will see the usage of both. These kinds allow you to change dependencies in runtime, so by design, they aren’t immutable. But if you need to change the implementation of some dependency, you don’t need to recreate everything. You can just override what you need. It may be useful if you have a feature flag that changes an implementation inside your service.



