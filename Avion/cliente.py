#!/usr/bin/env python
# -*- coding: utf-8 -*-
import grpc
import metadatos_pb2
import metadatos_pb2_grpc

#IMPLEMENTACIÓN DE FUNCIONES RELATIVAS A LA COMUNICACIÓN AVION - TORRE DE CONTROL

def Conectar_Torre(datos):
	channel = grpc.insecure_channel(datos['ipTorreInicial'])
	stub = metadatos_pb2_grpc.APIStub(channel)
	return stub

# Función que entrega un diccionario de datos hacia la torre de control inicial
def datos_a_torre(datos,stub):
	respuesta = stub.DatosHaciatorre(metadatos_pb2.dataAvion(aerolinea=datos['aerolinea'],codigo=datos['codigo'],pesoMax=int(datos['pesoMax']),CapTanque=int(datos['CapTanque']),ipTorreInicial=datos['ipTorreInicial']))
	return respuesta

# Funcion que consulta a la torre actual si existe el destino en solicitud para despegar.
def comprobar_destino(destino,datos,stub):
	respuesta = stub.ComprobarDestino(metadatos_pb2.dataDestino(destino=destino))
	return respuesta

# Solicita a la torre disponibilidad de pista y altura para despegar
def permiso_despegar(destino,datos,stub):
	respuesta = stub.SolicitarDespegue(metadatos_pb2.dataDespegue(destino=destino,aerolinea=datos['aerolinea'],codigo=codigo['codigo']))
	return respuesta

# IMPLEMENTAR DE MUTUO ACUERDO CON DAVID
def esperar_confirmacion(datos,stub):
	return 1

def contactar_torre(datos,stub):
	respuesta = stub.SolicitarAterrizaje(metadatos_pb2.dataAterrizaje(aerolinea=datos['aerolinea'],codigo=datos['codigo'],pesoMax=int(datos['pesoMax']),CapTanque=int(datos['CapTanque']),ipTorreActual=datos['ipTorreInicial']))
	return respuesta


