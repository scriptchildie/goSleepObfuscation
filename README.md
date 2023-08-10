# goSleepObfuscation
Work around the go runtime to obfuscate part of the executable code in golang executables

There are various tools to do sleep obfuscation in c & rust but it's nearly impossible to replicate in golang due to the go runtime. 

See : 
- https://github.com/Cracked5pider/Ekko
- https://github.com/memN0ps/ekko-rs


In order to get some sort of obfuscation i wrote this code that encrypts specific functions in memory when not in use.

In the code below the function toBEncrypted() is encrypted for 3 seconds during the sleep period. This is meant to be used as a PoC and it was not tested extensively.

## Caveats: 
This wouldn't work with goroutines
```
func main() {

	for {
		toBEncrypted()
		err := encryptFunc(reflect.ValueOf(toBEncrypted).Pointer(), 3) //encrypt function toBEncrypted(), sleep for 3 seconds , decrypt function
		if err != nil {
			log.Fatalf("%v", err)

		}
		toBEncrypted()
		fmt.Println("-----------------")
	}

}

// This function is encrypted during sleeping
func toBEncrypted() (uint32, error) {
	fmt.Println("is it working?")
	return 0, nil
}

```
