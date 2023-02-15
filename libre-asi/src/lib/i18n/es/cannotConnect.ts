const cannotConnect = {
	ERROR: 'Error.',
	PARAGRAPH_1:
		'Algo salió mal mientras se procesaba tu solicitud, intenta de nuevo más tarde, si el problema persiste. Contacta a tu administrador.',
	PARAGRAPH: 'Por tu seguridad, tu sesión ha sido invalidada.',
	CODE: 'Código: 503.',
	GO_BACK: 'Volver a la página de bienvenida',
	IF_ADMIN: 'Si eres el administrador, haz clic aquí para obtener ayuda',
	HELP: `El código 503 es retornado cuando Libre-ASI API rechazó la conexión, esto puede ser causado ya sea
    por una mala configuración de variables de entorno o detener la ejecución del servidor de la API, asegúrate
    que está en ejecución, revisa los logs del servidor y que su dirección está configurada conrrectamente en la
    aplicación e intenta nuevamente`
};

export default cannotConnect;
