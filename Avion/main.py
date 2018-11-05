#!/usr/bin/env python
# -*- coding: utf-8 -*-

import grpc
import metadatos_pb2
import metadatos_pb2_grpc
from funciones import *
from cliente import *

# Instanciar Avion
datos = IniciarAvion()

# Crear Canal Comunicación con la Torre Actual
stub = Conectar_Torre(datos)

# Enviar Datos a Torre Inicial ; Imprime el mensaje de respueta.
print(datos_a_torre(datos,stub))

while True:

	# Despegue
	datos = Despegar(datos,stub)

	# Aterrizaje
	stub = Aterrizar(datos)