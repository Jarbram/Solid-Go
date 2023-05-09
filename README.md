# Solid-Go
Los principios SOLID son un conjunto de cinco principios de diseño de software que se utilizan para crear código que es fácil de mantener, ampliar y modificar. Cada letra del acrónimo SOLID representa uno de estos principios:

- S: Principio de Responsabilidad Única (Single Responsibility Principle)
- O: Principio Abierto-Cerrado (Open-Closed Principle)
- L: Principio de Sustitución de Liskov (Liskov Substitution Principle)
- I: Principio de Segregación de Interfaces (Interface Segregation Principle)
- D: Principio de Inversión de Dependencia (Dependency Inversion Principle)

## Principio de Responsabilidad Única (SRP)

Establece que una clase o módulo debe tener una única responsabilidad, es decir, debe hacer sólo una cosa y hacerla bien. Esto significa que una clase o módulo debe tener una única razón para cambiar.

En términos más simples, SRP significa que cada clase debe tener una sola responsabilidad y no debe ser responsable de más de una cosa. Esto hace que el código sea más fácil de entender, de mantener y de ampliar.

La aplicación de SRP es importante cuando se desarrolla software porque hace que el código sea más fácil de manejar. Si una clase o módulo tiene demasiadas responsabilidades, es más difícil entender cómo funciona, más difícil de cambiar, y más propenso a errores. Además, si una clase tiene varias responsabilidades, es más probable que una modificación en una de ellas afecte a otra, lo que puede llevar a errores en el software.

- Un ejemplo de aplicación del principio SRP podría ser una clase que se encarga de manejar la lógica de negocio y también maneja la interacción con la base de datos. En este caso, la clase tiene dos responsabilidades, lo que hace que el código sea más difícil de entender y mantener. Para aplicar SRP, se podría dividir la clase en dos: una clase que se encarga de la lógica de negocio y otra clase que se encarga de la interacción con la base de datos.
- Otro ejemplo podría ser una clase que se encarga de enviar correos electrónicos y al mismo tiempo maneja la lógica de autenticación de usuarios. En este caso, la clase tiene dos responsabilidades diferentes, lo que hace que el código sea más difícil de entender y de mantener. Para aplicar SRP, se podría dividir la clase en dos: una clase que se encarga de enviar correos electrónicos y otra clase que se encarga de manejar la autenticación de usuarios.

En resumen, SRP es un principio importante que se debe aplicar para crear código limpio y fácil de entender. Se debe aplicar siempre que una clase o módulo tenga más de una responsabilidad. Al dividir una clase o módulo en varias partes más pequeñas y especializadas, el código se vuelve más fácil de entender, mantener y ampliar.

## Principio Abierto-Cerrado (OCP)
Es uno de los principios SOLID, y se refiere a la capacidad de un sistema para ser extendido sin tener que modificar su código fuente original. En otras palabras, el principio establece que un módulo o una clase debe estar cerrado para modificaciones pero abierto para extensiones.

Esto significa que, si necesitamos agregar nueva funcionalidad a nuestro programa, deberíamos poder hacerlo sin tener que modificar el código fuente original. En su lugar, deberíamos poder crear nuevas clases o módulos que extiendan la funcionalidad existente, manteniendo así la integridad del código original y minimizando el riesgo de introducir nuevos errores o problemas.

Para aplicar este principio en Go, es importante diseñar nuestros programas en módulos y paquetes separados, de tal manera que cada módulo o paquete tenga una única responsabilidad y pueda ser extendido sin afectar al resto del programa. También debemos utilizar interfaces en lugar de tipos concretos siempre que sea posible, ya que esto nos permite agregar nuevas implementaciones sin tener que modificar el código existente.
- Un ejemplo concreto de cómo aplicar este principio en Go podría ser una aplicación de venta en línea que permite a los usuarios buscar productos y realizar compras. Si en algún momento queremos agregar una nueva funcionalidad para permitir a los usuarios hacer comentarios sobre los productos, deberíamos poder hacerlo sin tener que modificar el código original de la aplicación.
Para lograr esto, podríamos crear un nuevo paquete "comments" que contenga todas las funciones necesarias para manejar comentarios. Este paquete podría implementar una interfaz definida en el paquete original de búsqueda de productos, lo que permitiría a la aplicación manejar comentarios sin tener que modificar el código fuente original. De esta manera, estamos aplicando el principio abierto-cerrado en Go y logrando una mayor flexibilidad y mantenibilidad en nuestro código.
