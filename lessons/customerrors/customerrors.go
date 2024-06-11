package main

type CustomError struct {
	Msg string
}

func (c *CustomError) Error() string {
	return c.Msg
}

func PhoneWorking(pwr bool) (string, error) {
	if pwr {
		return "Phone is working", nil
	}
	return "", &CustomError{"Phone is not working"}
}

func main() {
	phone, err := PhoneWorking(false)
	if err != nil {
		switch err.(type) {
		case *CustomError:
			// Type assertion
			// Accessing the fields of the struct
			customErr := err.(*CustomError)
			println("Error: ", customErr.Msg)
		default:
			println("Error: ", err)
		}
	} else {
		println(phone)
	}
}
