# Ejercicio Golang 

## Caracteristicas
1. Una tarjeta se identifica de acuerdo a:
* La marca
* Numero de tarjeta
* Cardholder
* Fecha de Vto
2. Una operación es válida si el consumo es menor a $1000
3. Una tarjeta es valida para operar si el vencimiento es posterior al corriente día
4. Existen 3 marcas de tarjetas: Visa, NARA y AMEX, cada marca, tiene un calculo de tasa diferente:
* Visa: <año>/<mes>
* Nara: <dia del mes> * 0.5
* AMEX: <mes>*0.1

## Ejercicios
En base a las características anteriores, consideramos los siguientes ejercicios:
1. Consultar y persistir la tarjeta en una base de datos
2. Informar si una operación es válida
3. Informar si una tarjeta es válida para operar
4. Obtener la tasa de una operación

## Ejercicios HTTP
1. Implementar una API para consultar los anteriores puntos.
2. Subir la solución a un host
3. Integrar la solución a una DB relacional