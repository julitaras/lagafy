# Seguridad
Para controlar el acceso y brindar seguridad a la información que manipulamos a continuación veremos como configurar un middleware para autenticación con JWT

## ¿Qué es JWT ?
JWT (o JSON WEB TOKEN) es un standard que nos permite hacer una comunicación entre dos partes de la identidad de un determinado de usuario de forma segura, permitiendonos llevar a cabo procesos de autenticación. 

Todo ello codificado en objetos json que se usan en el cuerpo (payload) de un mensaje firmado digitalmente. Se trata de una cadena de texto que tiene tres partes codificadas en base64 separadas por un punto.

Las tres partes: 

``Header``: Indica el algoritmo y el tipo del token(H256 y JWT).

``Payload``: Datos del usuario y privilegios: como nosotros generamos el token, podemos incluir todos los datos que estimemos convenientes.

``Signature``: Firma para verificar que el token es válido.

El proceso para usarlo sería el siguiente:
1. POST con usuario y contraseña
2. En el servidor se crea el token JWT con el secreto y se lo devuelve
3. Se envía el jwt en el header Autorization
4. En el servidor se comprueba la firma del token. Se obtiene el recurso protegido y se lo devuelve.

## OAUTH 2
Es un standard abierto para la autorización de APIs que permite compartir información entre sitios sin tener que compartir la identidad. 

- Implementa diferentes flujos de autenticación.
- Permite  la integración con aplicaciones de terceros.
- El usuario delega la capacidad de realizar ciertas acciones en su nombre.

``Roles que intervienen en el proceso de autenticación:``

#### Dueño del recurso

   Es el usuario que da la autorización a una aplicación para acceder  a su cuenta. Si bien la API no es suya, los datos que se manejan si lo son.

   El acceso a la cuenta del usuario se limita al “alcance” de la autorización otorgada

#### Cliente

   El cliente es la aplicación que desea acceder a la cuenta del usuario.

   Antes de que pueda hacerlo, el usuario debe darle la autorización y esta debe ser validada por la API.

#### Servidor de autorización

   Es el responsable de gestionar las peticiones de autorización.

   Verifica la identidad de los usuarios y emite tokens de acceso a la aplicación cliente.

#### Servidor de recursos

   Es el servicio que aloja el recurso protegido.

## PROCESO DE AUTENTICACIÓN CON MICROSOFT OAUTH 2

Azure Active Directory (Azure AD) usa OAuth 2.0 para permitirle autorizar el acceso a aplicaciones web y API web en su inquilino de Azure AD. Se utiliza para realizar autenticación y autorización en la mayoría de los tipos de aplicaciones.

Primero, registre su aplicación con su inquilino de Azure Active Directory (Azure AD). Esto le dará un ID de aplicación para su aplicación, así como también le permitirá recibir tokens.

El flujo del código de autorización comienza cuando el cliente dirige al usuario al /authorize. 

En este punto, se le pide al usuario que ingrese sus credenciales y consienta los permisos solicitados por la aplicación en el Portal de Azure. Una vez que el usuario se autentica y otorga su consentimiento, Azure AD envía una respuesta a su aplicación a la redirect_uri con el código.

Con el código de autorización y el permiso que otorgo el usuario, puede intercambiar el código por un token de acceso al recurso deseado, enviando una solicitud ``POST`` al ``/{inquilino}/oauth2/token``

Si la respuesta es exitosa, puede usar el token en las solicitudes a las API web, incluyéndose en el Authorization header.

Para más información, visitar la documentación oficial.

- [Microsoft OUAHT 2.0](https://docs.microsoft.com/en-us/azure/active-directory/azuread-dev/v1-protocols-oauth-code)

## JWT

Para empezar creamos la carpeta middleware a la misma altura que main.go y dentro de la carpeta creamos el archivo auth_middleware.go. Luego nomenclamos el package y hacemos los import's correspondientes:

```go
package middleware

import (
	"api-dashboard/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)
```

Creamos una función que retorne *gin.HandlerFunc*. Esa función de retorno será la que valide el token

```go
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
```

Dentro de la funcion que retornamos debemos obtener el token del header; para ello vamos a necesitar el package *strings*

```go
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		rawToken := strings.Trim(strings.TrimLeft(c.GetHeader("authorization"), "Bearer"), " ")

}
```

Ahora llamamos al método *Parse()* de jwt para validar el token y manejamos los posibles errores

```go
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		rawToken := strings.Trim(strings.TrimLeft(c.GetHeader("authorization"), "Bearer"), " ")

		tkn, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
			// El secret del jwt debe ir en un archivo .env al cual accedemos mediante 'setting.AppSetting'
			return []byte(setting.AppSetting.JwtSecret), nil
		})
	}
}
```

Luego nos encargamos de manejar los posibles errores y usamos el método *Next()* para que el flujo de datos salga del middleware y siga la ejecución del request 

```go
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		rawToken := strings.Trim(strings.TrimLeft(c.GetHeader("authorization"), "Bearer"), " ")

		tkn, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
			// El secret del jwt debe ir en un archivo .env al cual accedemos mediante 'setting.AppSetting'
			return []byte(setting.AppSetting.JwtSecret), nil
		})
		// Manejamos el error
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		// Verificamos la validez del token
		if !tkn.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
```

Para terminar, debemos ir al archivo router.go y dentro de la función *SetRoutes()* configurar el middleware en los groups

```go
group := r.Group("/apigroup")
group.Use(middleware.JwtAuth())
```

## Office 365

El comienzo es muy parecido a lo que hicimos para JWT.

En el archivo auth_middleware.go que se encuentra dentro de la carpeta middleware creamos el package y los import's

```go
package middleware

import (
	"api-dashboard/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
	oidc "github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
)
```

Seguidamente creamos una funcion que retorne *gin.HandlerFunc*; esa función de retorno validará el token

```go
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
```

Ahora usando el método *NewProvider()* de la librería de OpenID Connect *go oidc* generamos un nuevo provider, que no es otra cosa que la abstracción del servicio externo de autenticación. Aquí también debemos manejar los posibles errores

```go
provider, err := oidc.NewProvider(c, "https://login.microsoftonline.com/5ab9af9b-4534-4c31-8e50-1e098461481c/v2.0")
// Manejamos posibles errores
if err != nil {
	log.Println((err))
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"Auth": "Error getting provider",
	})
	return
}
```

Utilizamos el package *strings* con sus métodos *Trim()* y *TrimLeft()* para obtener el token

```go
rawIDToken := strings.Trim(strings.TrimLeft(c.GetHeader("authorization"), "Bearer"), " ")
```

Luego creamos un verificador utilizando el método *Verifier()* de *provider* pasando como parámetro la configuración con el ClientID del archivo .env

```go
// El ClientID debe ir en un archivo .env al cual accedemos mediante 'setting.AppSetting'
verifier := provider.Verifier(&oidc.Config{ClientID: setting.AppSetting.ClientID})
```

Mediante el método *Verify()* del *verifier* comprobamos la validez del token y obtenemos el payload

```go
idToken, err := verifier.Verify(c, rawIDToken)
// Manejamos posibles errores
if err != nil {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"Token": "Invalid Token",
	})
	c.Abort()
	return
}
```

Una vez hecha la validación y el parseo del token podemos obtener los claims, que son pares clave-valor que contiene la información (En este caso nombre y mail). Para esto necesitamos inicializar un struct para los datos y pasarle la dirección en memoria de este al método *Claims()* 

```go
var claims struct {
	Name  string `json:"name"`
	Email string `json:"preferred_username"`
}
// Obtenemos los claims y manejamos posibles errores
if err := idToken.Claims(&claims); err != nil {
	log.Println(err)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"Claims": "Error extracting custom claims",
	})
}
```

Finalmente guardamos los datos obtenidos en el contexto y mediante el método *Next()* dejamos atrás el middleware y continuamos con el flujo del request

```go
c.Set("userEmail", claims.Email)
c.Set("userName", claims.Name)

c.Next()
```

Recapitulando, todo junto seria de la siguiente forma:

```go
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {

		provider, err := oidc.NewProvider(c, "https://login.microsoftonline.com/5ab9af9b-4534-4c31-8e50-1e098461481c/v2.0")
		if err != nil {
			log.Println((err))
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"Auth": "Error getting provider",
			})
			return
		}

		rawIDToken := strings.Trim(strings.TrimLeft(c.GetHeader("authorization"), "Bearer"), " ")

		verifier := provider.Verifier(&oidc.Config{ClientID: setting.AppSetting.ClientID})

		idToken, err := verifier.Verify(c, rawIDToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"Token": "Invalid Token",
			})
			c.Abort()
			return
		}

		var claims struct {
			Name  string `json:"name"`
			Email string `json:"preferred_username"`
		}
		if err := idToken.Claims(&claims); err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"Claims": "Error extracting custom claims",
			})

		}

		c.Set("userEmail", claims.Email)
		c.Set("userName", claims.Name)

		c.Next()
	}
}
```

Para más información visitar el siguiente links:

- [Go oidc](https://godoc.org/github.com/coreos/go-oidc)
- [Auth0](https://auth0.com/docs/tokens/)
- [OpenID connect](https://openid.net/connect/)