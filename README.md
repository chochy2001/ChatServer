# English Version
# Chat Server and Netcat Client in Go

This project implements a chat server and a netcat client using the Go programming language. The server allows multiple clients to connect and send messages, while the netcat client can be used to connect to the server and send messages.

## Project Structure

The project consists of two main files:

1. `chat.go`: Contains the code for the chat server.
2. `netcat.go`: Contains the code for the netcat client.

## Requirements

- Go 1.22 or higher

## Setup

1. Make sure you have Go installed on your system.

2. Clone this repository or download the `chat.go` and `netcat.go` files.

3. Open a terminal and navigate to the project directory.

4. Run the following command to download the dependencies:
   ```
   go mod download
   ```

## Running the Chat Server

1. In the terminal, run the following command to start the chat server:
   ```
   go run chat.go
   ```

2. By default, the server will run on `localhost` and port `8080`. You can specify a different host and port using the `-host` and `-port` flags, respectively. For example:
   ```
   go run chat.go -host 0.0.0.0 -port 9000
   ```

## Running the Netcat Client

1. Open another terminal and run the following command to start the netcat client and connect to the chat server:
   ```
   go run netcat.go
   ```

2. By default, the client will connect to `localhost` and port `8080`. You can specify a different host and port using the `-host` and `-port` flags, respectively. For example:
   ```
   go run netcat.go -host 192.168.0.100 -port 9000
   ```

3. Once connected, you will receive a welcome message that includes your IP address and port. For example:
   ```
   You are 127.0.0.1:52440 Welcome to the server
   ```

4. You can now send messages to the chat server. The messages will be broadcast to all connected clients, along with your client name (IP address and port). For example:
   ```
   Hello, I'm Josh
   127.0.0.1:52440: Hello, I'm Josh
   ```

5. When other clients send messages, you will receive them with their respective client names. For example:
   ```
   127.0.0.1:52438: Hi, my name is Jose
   ```

6. To exit the netcat client, press `Ctrl+C`. Upon exiting, a message will be sent indicating that you have left the chat. For example:
   ```
   127.0.0.1:52438 has left
   ```

## Interaction Examples

Here are examples of interaction between two clients:

Client 1:
```
./netcat
You are 127.0.0.1:52440 Welcome to the server
Hello, I'm Josh
127.0.0.1:52440: Hello, I'm Josh
What's your name?
127.0.0.1:52440: What's your name?
127.0.0.1:52438: Hi, my name is Jose
127.0.0.1:52438: Have a great day, I have to go.
127.0.0.1:52438: Bye.
127.0.0.1:52438:
127.0.0.1:52438 has left
```

Client 2:
```
./netcat
You are 127.0.0.1:52438 Welcome to the server
127.0.0.1:52440 has arrived
127.0.0.1:52440: Hello, I'm Josh
127.0.0.1:52440: What's your name?
Hi, my name is Jose
127.0.0.1:52438: Hi, my name is Jose
Have a great day, I have to go.
127.0.0.1:52438: Have a great day, I have to go.
Bye.
127.0.0.1:52438: Bye.
^C
127.0.0.1:52438:
^C
```

## Code Functionality

### Chat Server (`chat.go`)

- The server uses channels to manage clients and messages.
- Three channels are defined: `incomingClients`, `leavingClients`, and `messages`, representing incoming clients, leaving clients, and messages sent by clients, respectively.
- The `HandleConnection` function runs in a separate goroutine for each connected client. It handles the connection, sends welcome and closing messages, and forwards client messages to the `messages` channel.
- The `MessageWrite` function runs in a separate goroutine for each client and is responsible for sending messages from the `messages` channel to the corresponding client.
- The `Broadcaster` function runs in a goroutine and is responsible for receiving messages from the `messages` channel and sending them to all connected clients.
- In the `main` function, the server is started, new client connections are accepted, and they are handled in separate goroutines using the `HandleConnection` function.

### Netcat Client (`netcat.go`)

- The netcat client connects to the chat server using the `net.Dial` function.
- A control channel `done` is created to indicate when the copying of data from the connection to the standard output is completed.
- A goroutine is started that copies the received data from the connection to the standard output using `io.Copy`. When the copying is complete, a signal is sent to the `done` channel.
- The `CopyContent` function is used to copy data from the standard input to the connection, allowing the user to send messages to the server.
- In the `main` function, a connection to the server is established, goroutines are started to handle input and output data, and it waits for the data copying to complete or the program to be interrupted.

Enjoy chatting using the Go chat server and netcat client!


# Spanish Version

# Chat Server y Cliente Netcat en Go

Este proyecto implementa un servidor de chat y un cliente netcat utilizando el lenguaje de programación Go. El servidor permite que múltiples clientes se conecten y envíen mensajes, mientras que el cliente netcat se puede utilizar para conectarse al servidor y enviar mensajes.

## Estructura del proyecto

El proyecto consta de dos archivos principales:

1. `chat.go`: Contiene el código del servidor de chat.
2. `netcat.go`: Contiene el código del cliente netcat.

## Requisitos

- Go 1.22 o superior

## Configuración

1. Asegúrate de tener Go instalado en tu sistema.

2. Clona este repositorio o descarga los archivos `chat.go` y `netcat.go`.

3. Abre una terminal y navega hasta el directorio del proyecto.

4. Ejecuta el siguiente comando para descargar las dependencias:
   ```
   go mod download
   ```

## Ejecución del servidor de chat

1. En la terminal, ejecuta el siguiente comando para iniciar el servidor de chat:
   ```
   go run chat.go
   ```

2. Por defecto, el servidor se ejecutará en `localhost` en el puerto `8080`. Puedes especificar un host y puerto diferentes utilizando los flags `-host` y `-port`, respectivamente. Por ejemplo:
   ```
   go run chat.go -host 0.0.0.0 -port 9000
   ```

## Ejecución del cliente netcat

1. Abre otra terminal y ejecuta el siguiente comando para iniciar el cliente netcat y conectarte al servidor de chat:
   ```
   go run netcat.go
   ```

2. Por defecto, el cliente se conectará a `localhost` en el puerto `8080`. Puedes especificar un host y puerto diferentes utilizando los flags `-host` y `-port`, respectivamente. Por ejemplo:
   ```
   go run netcat.go -host 192.168.0.100 -port 9000
   ```

3. Una vez conectado, recibirás un mensaje de bienvenida que incluye tu dirección IP y puerto. Por ejemplo:
   ```
   You are 127.0.0.1:52440 Welcome to the server
   ```

4. Ahora puedes enviar mensajes al servidor de chat. Los mensajes se enviarán a todos los clientes conectados, junto con tu nombre de cliente (dirección IP y puerto). Por ejemplo:
   ```
   Hola, Soy Josh
   127.0.0.1:52440: Hola, Soy Josh
   ```

5. Cuando otros clientes envíen mensajes, los recibirás con su respectivo nombre de cliente. Por ejemplo:
   ```
   127.0.0.1:52438: Hola, mi nombre es Jose
   ```

6. Para salir del cliente netcat, presiona `Ctrl+C`. Al salir, se enviará un mensaje indicando que has abandonado el chat. Por ejemplo:
   ```
   127.0.0.1:52438 has left
   ```

## Ejemplos de interacción

A continuación se muestran ejemplos de interacción entre dos clientes:

Cliente 1:
```
./netcat
You are 127.0.0.1:52440 Welcome to the server
Hola, Soy Josh
127.0.0.1:52440: Hola, Soy Josh
Tu como te llamas?
127.0.0.1:52440: Tu como te llamas?
127.0.0.1:52438: Hola, mi nombre es Jose
127.0.0.1:52438: que tengas un excelente dia, me tengo que ir.
127.0.0.1:52438: bye.
127.0.0.1:52438:
127.0.0.1:52438 has left
```

Cliente 2:
```
./netcat
You are 127.0.0.1:52438 Welcome to the server
127.0.0.1:52440 has arrived
127.0.0.1:52440: Hola, Soy Josh
127.0.0.1:52440: Tu como te llamas?
Hola, mi nombre es Jose
127.0.0.1:52438: Hola, mi nombre es Jose
que tengas un excelente dia, me tengo que ir.
127.0.0.1:52438: que tengas un excelente dia, me tengo que ir.
bye.
127.0.0.1:52438: bye.
^C
127.0.0.1:52438:
^C
```

## Funcionamiento del código

### Servidor de chat (`chat.go`)

- El servidor utiliza canales para gestionar los clientes y los mensajes.
- Se definen tres canales: `incomingClients`, `leavingClients` y `messages`, que representan los clientes entrantes, los clientes salientes y los mensajes enviados por los clientes, respectivamente.
- La función `HandleConnection` se ejecuta en una goroutine separada para cada cliente que se conecta. Maneja la conexión, envía mensajes de bienvenida y cierre, y reenvía los mensajes del cliente al canal `messages`.
- La función `MessageWrite` se ejecuta en una goroutine separada para cada cliente y se encarga de enviar los mensajes del canal `messages` al cliente correspondiente.
- La función `Broadcaster` se ejecuta en una goroutine y se encarga de recibir los mensajes del canal `messages` y enviarlos a todos los clientes conectados.
- En la función `main`, se inicia el servidor, se aceptan nuevas conexiones de clientes y se manejan en goroutines separadas utilizando la función `HandleConnection`.

### Cliente Netcat (`netcat.go`)

- El cliente netcat se conecta al servidor de chat utilizando la función `net.Dial`.
- Se crea un canal de control `done` para indicar cuándo se ha completado la copia de datos desde la conexión a la salida estándar.
- Se inicia una goroutine que copia los datos recibidos desde la conexión a la salida estándar utilizando `io.Copy`. Cuando se completa la copia, se envía una señal al canal `done`.
- La función `CopyContent` se utiliza para copiar los datos desde la entrada estándar a la conexión, permitiendo al usuario enviar mensajes al servidor.
- En la función `main`, se establece la conexión con el servidor, se inician las goroutines para manejar la entrada y salida de datos, y se espera a que se complete la copia de datos o se interrumpa el programa.

¡Disfruta del chat utilizando el servidor y el cliente netcat en Go!





