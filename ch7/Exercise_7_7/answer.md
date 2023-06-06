“Exercise 7.7:
Explain why the help message contains °C when the default
value of 20.0 does not.

The help message in the code includes the unit "°C" because it is specified in the usage string passed to the
TemperatureFlag function when defining the temp flag. Here's the relevant line of code:

temp := TemperatureFlag("temp", 20.0, "the temperature")

The usage string provided is "the temperature", which doesn't explicitly mention the unit. However, the unit "°C" is
appended to the usage string internally when creating the flag help message. This behavior is implemented by the flag
package in Go.

When generating the help message, the flag package combines the flag name, default value, and usage string to form
the message. In this case, the default value is 20.0, and the unit is "°C". The flag package includes the default
value and the unit in the help message to provide additional context and information about the flag's expected format.

Therefore, even though the default value of 20.0 doesn't explicitly include the unit "°C" in the code,
the flag package adds it to the help message to provide a more complete description of the expected input
format for the temp flag.

