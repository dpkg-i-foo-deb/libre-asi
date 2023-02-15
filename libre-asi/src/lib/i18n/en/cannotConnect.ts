const cannotConnect = {
	ERROR: 'Error.',
	PARAGRAPH_1:
		'Something went wrong while processing your request, try again later, if the error persists.	Contact your administrator.',
	PARAGRAPH_2: 'For your security, your session has been invalidated.',
	CODE: 'Code: 503.',
	GO_BACK: 'Go back to welcome page',
	IF_ADMIN: "If you're the administrator... Click here to see some help",
	HELP: `Status code 503 is thrown when Libre-ASI API refused the connection, this can be caused by
	either misconfiguring environment variables or stopping the API server. Make sure it is
	running, check server logs and its address is correctly configured on the server and try again`
};

export default cannotConnect;
