def exeption_to_dict(exception: Exception) -> dict:
	return {
		"type": str(type(exception)),
		"args": exception.args,
		"message": str(exception),
	}

def exception_from_dict(exception: dict) -> Exception:
	exception_type = exception["type"]

    # Use eval carefully and consider security implications (see note below)
	exception_class = eval(exception_type.replace("'", ""))  # Remove quotes for eval
	args = exception["args"]
	message = exception["message"]

	# Create the exception instance
	return exception_class(*args, message=message)