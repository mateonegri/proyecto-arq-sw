# 📌**Arquitectura de Software I - Integrador 2023**
Este proyecto consiste en la creación de un sitio de reserva de habitaciones para una cadena de hoteles. Está compuesto por un backend desarrollado en Golang, que proporcionará las interfaces necesarias, y un frontend desarrollado en React, que representa la interfaz de usuario y consume los servicios del backend.

## ✒️Integrantes: 
#### 	- Mateo Negri
#### 	- Manuela Simes
#### 	- Valentino Gaggio
#### 	- Paz Vartanian

## 🛠️Construido con...
Para este proyecto como se comento anteriormente se utilizo:

#### - Golang
#### - React


### **Características**
**Autenticación de usuarios:** El sistema cuenta con un sistema de inicio de sesión y gestión de permisos de usuarios. Existen dos tipos de usuarios: viajante y administrador. Los administradores tienen la capacidad de agregar nuevos hoteles, incluyendo título, foto, descripción y habitaciones, así como ver el listado de reservas realizadas por los viajantes.

**Pantalla de bienvenida:** La página principal de la aplicación muestra un listado de hoteles disponibles en un menú desplegable (combo box), junto con las fechas de check-in y check-out seleccionadas.

**Detalle del hotel:** Al seleccionar un hotel, se muestra la disponibilidad de habitaciones para la reserva, así como las alternativas disponibles. Todos los hoteles tienen el mismo tipo de habitación y la disponibilidad se agota cuando se reservan todas las habitaciones disponibles para el período seleccionado.
Confirmación de reserva: Una vez seleccionada una habitación, se solicita al cliente que inicie sesión. Si el cliente no está registrado, se le pedirá que se registre proporcionando sus datos personales antes de confirmar la reserva.

**Página de listado de reservas:** Tanto los clientes como los administradores del sitio tienen acceso a una página donde pueden ver y filtrar las reservas realizadas. Los administradores tienen la capacidad de filtrar las reservas por hotel y día, y visualizar los datos de las personas que realizaron las reservas.

### **Requisitos adicionales**
No se requiere implementar procesos de pago. Se asume que el pago ha sido aprobado y se enfoca en el caso de uso de confirmación de la reserva.

**Puntos adicionales:** Se valorará la posibilidad de subir fotos de los hoteles y cargar comodidades (amenities) para los mismos.

### Condiciones de Regularidad y Examen Final
**Regularidad:** Para regularizar la materia se debe desarrollar el flujo de reserva.

**Examen Final:** Para el examen final se debe completar el proyecto en su totalidad, incluyendo la carga completa de hoteles con la posibilidad de subir fotos y comodidades.
