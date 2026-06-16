1. ¿Cuántas líneas tiene tu función main al final del taller? Cuéntalas con cuidado.
Tiene un aproximado de 63 líneas de código funcional sin contar los comentarios.

2.	Si tuvieras que agregar 5 operaciones más (raíz cuadrada, logaritmo, seno, coseno, módulo), ¿qué tan grande se haría tu main? ¿Sería fácil de leer para alguien que vea el código por primera vez?
Si le agregamos más operaciones el código crecería considerablemente, sería complejo entenderlo y más si apenas te estas relacionando con las herramientas.

3.	Notaste que el código para 'pedir un número al usuario' o 'imprimir el resultado' se repite varias veces. ¿No sería mejor escribirlo una sola vez y reutilizarlo en muchos lugares?
Si lo note y claro que sería mejor crear funciones reutilizables para lograr una estructura de código más limpia.

4.	Tu historial es una variable string gigante. ¿Qué pasaría si quisieras: ordenarlo alfabéticamente, eliminar la operación número 2, ¿o contar cuántas veces se usó la operación de suma?
Al usar un string para el historial no es posible manipular los datos fácilmente, no se pueden ordenar, ni se puede eliminar elementos específicos. Revisando la documentación oficial nos dimos cuenta de que si sería mejor usar una estructura como un slice para manejar cada operación individualmente.

5.  Después de este taller, ¿qué fue lo más difícil de Go para ti? ¿Y lo más interesante?
Lo más difícil de Go fue entender el manejo de los datos, también el control del flujo con for y switch. Lo más interesante fue aprender a estructurar programas con funciones y ver cómo se pueden facilitar operaciones matemáticas de forma eficiente.